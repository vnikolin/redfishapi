package redfishapi

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	legacyDeleteJobQueuePath = "/redfish/v1/Dell/Managers/iDRAC.Embedded.1/DellJobService/Actions/DellJobService.DeleteJobQueue"
	oemDeleteJobQueuePath    = "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellJobService/Actions/DellJobService.DeleteJobQueue"
	clearAllJobsPayload      = `{"JobID": "JID_CLEARALL_FORCE"}`
)

func TestClearJobsDellForceLegacyPathPreferred(t *testing.T) {
	var gotPaths []string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPaths = append(gotPaths, r.URL.Path)
		assertClearJobsDellForceRequest(t, r)

		if r.URL.Path == legacyDeleteJobQueuePath {
			fmt.Fprint(w, `{"Message":"The specified job was deleted."}`)
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	result, err := provider.ClearJobsDellForce()
	if err != nil {
		t.Fatalf("ClearJobsDellForce returned error: %v", err)
	}
	if result != "success" {
		t.Fatalf("expected success, got %q", result)
	}
	if len(gotPaths) != 1 || gotPaths[0] != legacyDeleteJobQueuePath {
		t.Fatalf("expected only legacy path, got %v", gotPaths)
	}
}

func TestClearJobsDellForceFallbackToOEMPath(t *testing.T) {
	var gotPaths []string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPaths = append(gotPaths, r.URL.Path)
		assertClearJobsDellForceRequest(t, r)

		switch r.URL.Path {
		case legacyDeleteJobQueuePath:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		case oemDeleteJobQueuePath:
			fmt.Fprint(w, `{"Message":"The specified job was deleted."}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	result, err := provider.ClearJobsDellForce()
	if err != nil {
		t.Fatalf("ClearJobsDellForce returned error: %v", err)
	}
	if result != "success" {
		t.Fatalf("expected success, got %q", result)
	}
	if len(gotPaths) != 2 || gotPaths[0] != legacyDeleteJobQueuePath || gotPaths[1] != oemDeleteJobQueuePath {
		t.Fatalf("expected legacy path then OEM path, got %v", gotPaths)
	}
}

func assertClearJobsDellForceRequest(t *testing.T, r *http.Request) {
	t.Helper()

	if r.Method != http.MethodPost {
		t.Fatalf("expected POST, got %s", r.Method)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("unable to read request body: %v", err)
	}
	if string(body) != clearAllJobsPayload {
		t.Fatalf("expected body %s, got %s", clearAllJobsPayload, string(body))
	}
}
