package redfishapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestGetNetworkSwitchInfoDellPrefersNewSystemPath(t *testing.T) {
	var oldPathHits atomic.Int32

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSwitchConnections":
			fmt.Fprint(w, `{"Members":[{"@odata.id":"/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSwitchConnections/NIC.Slot.10-1-1"}]}`)
		case "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSwitchConnections/NIC.Slot.10-1-1":
			fmt.Fprint(w, `{"Id":"NIC.Slot.10-1-1","Description":"DellSwitchConnection","StaleData":"NotStale","SwitchConnectionID":"Leaf1","SwitchPortConnectionID":"Ethernet1"}`)
		case "/redfish/v1/Systems/System.Embedded.1/NetworkPorts/Oem/Dell/DellSwitchConnections/":
			oldPathHits.Add(1)
			http.NotFound(w, r)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	result, err := provider.GetNetworkSwitchInfoDell()
	if err != nil {
		t.Fatalf("GetNetworkSwitchInfoDell returned error: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 switch result, got %d", len(result))
	}
	if result[0].Name != "NIC.Slot.10-1-1" {
		t.Fatalf("unexpected switch entry name: %s", result[0].Name)
	}
	if result[0].SwitchConnectionID != "Leaf1" || result[0].SwitchPortConnectionID != "Ethernet1" {
		t.Fatalf("unexpected switch entry values: %+v", result[0])
	}
	if oldPathHits.Load() != 0 {
		t.Fatalf("expected old path not to be queried when new path works, got %d hits", oldPathHits.Load())
	}
}

func TestGetNetworkSwitchInfoDellFallsBackToLegacyPath(t *testing.T) {
	var newPathHits atomic.Int32

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSwitchConnections":
			newPathHits.Add(1)
			http.NotFound(w, r)
		case "/redfish/v1/Systems/System.Embedded.1/NetworkPorts/Oem/Dell/DellSwitchConnections/":
			fmt.Fprint(w, `{"Members":[{"@odata.id":"/redfish/v1/Systems/System.Embedded.1/NetworkPorts/Oem/Dell/DellSwitchConnections/NIC.Integrated.1-1-1"}]}`)
		case "/redfish/v1/Systems/System.Embedded.1/NetworkPorts/Oem/Dell/DellSwitchConnections/NIC.Integrated.1-1-1":
			fmt.Fprint(w, `{"Id":"NIC.Integrated.1-1-1","Description":"DellSwitchConnection","StaleData":"NotStale","SwitchConnectionID":"aa:bb:cc:dd:ee:ff","SwitchPortConnectionID":"Ethernet49"}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	result, err := provider.GetNetworkSwitchInfoDell()
	if err != nil {
		t.Fatalf("GetNetworkSwitchInfoDell returned error: %v", err)
	}
	if newPathHits.Load() != 1 {
		t.Fatalf("expected new path to be attempted once, got %d hits", newPathHits.Load())
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 switch result, got %d", len(result))
	}
	if result[0].Name != "NIC.Integrated.1-1-1" {
		t.Fatalf("unexpected switch entry name: %s", result[0].Name)
	}
	if result[0].SwitchConnectionID != "aa:bb:cc:dd:ee:ff" || result[0].SwitchPortConnectionID != "Ethernet49" {
		t.Fatalf("unexpected switch entry values: %+v", result[0])
	}
}
