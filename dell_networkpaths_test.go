package redfishapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIDRACAttrDellFallsBackToManagerEthernetInterface(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/redfish/v1/Managers/iDRAC.Embedded.1/Attributes":
			http.NotFound(w, r)
		case "/redfish/v1/Managers/iDRAC.Embedded.1/EthernetInterfaces/NIC.1":
			fmt.Fprint(w, `{"Id":"NIC.1","MACAddress":"aa:bb:cc:dd:ee:ff","PermanentMACAddress":"aa:bb:cc:dd:ee:ff"}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	attrs, err := provider.GetIDRACAttrDell()
	if err != nil {
		t.Fatalf("GetIDRACAttrDell returned error: %v", err)
	}
	if attrs.CurrentNIC_1_MACAddress != "aa:bb:cc:dd:ee:ff" {
		t.Fatalf("unexpected CurrentNIC_1_MACAddress: %q", attrs.CurrentNIC_1_MACAddress)
	}
}

func TestGetNetworkPortsDellPrefersSystemEthernetInterfaces(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/redfish/v1/Systems/System.Embedded.1/EthernetInterfaces":
			fmt.Fprint(w, `{"Members":[{"@odata.id":"/redfish/v1/Systems/System.Embedded.1/EthernetInterfaces/NIC.Slot.10-1-1"}]}`)
		case "/redfish/v1/Systems/System.Embedded.1/EthernetInterfaces/NIC.Slot.10-1-1":
			fmt.Fprint(w, `{"Id":"NIC.Slot.10-1-1","Description":"Ethernet Interface","MACAddress":"8C:84:74:FA:21:84","PermanentMACAddress":"8C:84:74:FA:21:84","LinkStatus":"LinkUp","SpeedMbps":1000,"Status":{"Health":"OK","State":"Enabled"}}`)
		case "/redfish/v1/Chassis/System.Embedded.1/NetworkAdapters":
			t.Fatalf("legacy chassis network adapters path should not be queried when system EthernetInterfaces works")
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	ports, err := provider.GetNetworkPortsDell()
	if err != nil {
		t.Fatalf("GetNetworkPortsDell returned error: %v", err)
	}
	if len(ports) != 1 {
		t.Fatalf("expected 1 port, got %d", len(ports))
	}
	if ports[0].Name != "NIC.Slot.10-1-1" {
		t.Fatalf("unexpected port name: %q", ports[0].Name)
	}
	if ports[0].MacAddress != "8C:84:74:FA:21:84" {
		t.Fatalf("unexpected mac address: %q", ports[0].MacAddress)
	}
	if ports[0].State != "Up" {
		t.Fatalf("unexpected link state: %q", ports[0].State)
	}
	if ports[0].CurrentLinkSpeedMbps != 1000 {
		t.Fatalf("unexpected link speed: %d", ports[0].CurrentLinkSpeedMbps)
	}
}

func TestGetMacAddressModelDellFallsBackToChassisNetworkAdapters(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/redfish/v1/Systems/System.Embedded.1/NetworkAdapters/":
			http.NotFound(w, r)
		case "/redfish/v1/Chassis/System.Embedded.1/NetworkAdapters":
			fmt.Fprint(w, `{"Members":[{"@odata.id":"/redfish/v1/Chassis/System.Embedded.1/NetworkAdapters/NIC.Slot.10"},{"@odata.id":"/redfish/v1/Chassis/System.Embedded.1/NetworkAdapters/NIC.Slot.9"}]}`)
		case "/redfish/v1/Chassis/System.Embedded.1/NetworkAdapters/NIC.Slot.10":
			fmt.Fprint(w, `{"Id":"NIC.Slot.10","Manufacturer":"Broadcom","Model":"Broadcom BCM5719 4x1GbT OCP Ethernet NIC","Controllers":[{"Links":{"NetworkDeviceFunctions":[{"@odata.id":"/redfish/v1/Chassis/System.Embedded.1/NetworkAdapters/NIC.Slot.10/NetworkDeviceFunctions/NIC.Slot.10-1-1"}]}}]}`)
		case "/redfish/v1/Chassis/System.Embedded.1/NetworkAdapters/NIC.Slot.9":
			fmt.Fprint(w, `{"Id":"NIC.Slot.9","Manufacturer":"Mellanox Technologies","Model":null,"Controllers":[{"Links":{"NetworkDeviceFunctions":[{"@odata.id":"/redfish/v1/Chassis/System.Embedded.1/NetworkAdapters/NIC.Slot.9/NetworkDeviceFunctions/NIC.Slot.9-1-1"}]}}]}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	models, err := provider.GetMacAddressModelDell()
	if err != nil {
		t.Fatalf("GetMacAddressModelDell returned error: %v", err)
	}
	if len(models) != 2 {
		t.Fatalf("expected 2 models, got %d", len(models))
	}
	if models[0].MacName != "NIC.Slot.10-1-1" || models[0].MacModel != "Broadcom BCM5719 4x1GbT OCP Ethernet NIC" {
		t.Fatalf("unexpected first model: %#v", models[0])
	}
	if models[1].MacName != "NIC.Slot.9-1-1" || models[1].MacManufacturer != "Mellanox Technologies" {
		t.Fatalf("unexpected second model: %#v", models[1])
	}
}
