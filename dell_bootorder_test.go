package redfishapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestGetBootOrderDellFallsBackToOemDellBootSources(t *testing.T) {
	var genericPathHits atomic.Int32

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/redfish/v1/Systems/System.Embedded.1/BootSources":
			genericPathHits.Add(1)
			http.NotFound(w, r)
		case "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellBootSources":
			fmt.Fprint(w, `{"Attributes":{"UefiBootSeq":[{"Enabled":true,"Id":"BIOS.Setup.1-1#UefiBootSeq#Unknown.Unknown.1-1#abc","Index":0,"Name":"Unknown.Unknown.1-1"},{"Enabled":true,"Id":"BIOS.Setup.1-1#UefiBootSeq#NIC.PxeDevice.1-1#def","Index":1,"Name":"NIC.PxeDevice.1-1"}]}}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	result, err := provider.GetBootOrderDell()
	if err != nil {
		t.Fatalf("GetBootOrderDell returned error: %v", err)
	}
	if genericPathHits.Load() != 1 {
		t.Fatalf("expected generic boot sources path to be attempted once, got %d hits", genericPathHits.Load())
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 boot order entries, got %d", len(result))
	}
	if result[0].Name != "Unknown.Unknown.1-1" || result[0].Index != 0 || !result[0].Enabled {
		t.Fatalf("unexpected first boot entry: %+v", result[0])
	}
	if result[1].Name != "NIC.PxeDevice.1-1" || result[1].Index != 1 || !result[1].Enabled {
		t.Fatalf("unexpected second boot entry: %+v", result[1])
	}
}

func TestSetBootOrderDellTreatsNoContentAsSuccess(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/redfish/v1/Systems/System.Embedded.1/BootSources/Settings":
			w.WriteHeader(http.StatusNoContent)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	message, err := provider.SetBootOrderDell([]byte(`{"Attributes":{}}`))
	if err != nil {
		t.Fatalf("SetBootOrderDell returned error: %v", err)
	}
	if message != "" {
		t.Fatalf("expected empty message for 204 response, got %q", message)
	}
}
