package redfishapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	oemIDRACAttrPath = "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellAttributes/iDRAC.Embedded.1"
	stdIDRACAttrPath = "/redfish/v1/Managers/iDRAC.Embedded.1/Attributes"
)

const idracAttrBody = `{
		"Attributes": {
			"SerialCapture.1.Enable": "Enabled",
			"SerialCapture.1.SerialDataSize": 524288,
			"CurrentNIC.1.MACAddress": "aa:bb:cc:dd:ee:ff"
		}
	}`

// TestGetIDRACAttrDellOEMPreferred verifies that the OEM DellAttributes endpoint
// is queried first and its response is consumed when it returns 200.
func TestGetIDRACAttrDellOEMPreferred(t *testing.T) {
	var gotPaths []string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPaths = append(gotPaths, r.URL.Path)
		if r.URL.Path == oemIDRACAttrPath {
			fmt.Fprint(w, idracAttrBody)
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	attrs, err := provider.GetIDRACAttrDell()
	if err != nil {
		t.Fatalf("GetIDRACAttrDell returned error: %v", err)
	}
	if attrs.SerialCapture_1_Enable != "Enabled" {
		t.Fatalf("expected SerialCapture.1.Enable=Enabled, got %q", attrs.SerialCapture_1_Enable)
	}
	if attrs.SerialCapture_1_SerialDataSize != 524288 {
		t.Fatalf("expected SerialCapture.1.SerialDataSize=524288, got %d", attrs.SerialCapture_1_SerialDataSize)
	}
	if len(gotPaths) == 0 || gotPaths[0] != oemIDRACAttrPath {
		t.Fatalf("expected OEM path queried first, got %v", gotPaths)
	}
	for _, p := range gotPaths {
		if p == stdIDRACAttrPath {
			t.Fatalf("standard path should not be queried when OEM succeeds: %v", gotPaths)
		}
	}
}

// TestGetIDRACAttrDellFallbackToStandard verifies that when the OEM endpoint
// returns 404 (e.g. older firmware), the standard endpoint is queried as a
// fallback and its body is consumed.
func TestGetIDRACAttrDellFallbackToStandard(t *testing.T) {
	var gotPaths []string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPaths = append(gotPaths, r.URL.Path)
		if r.URL.Path == stdIDRACAttrPath {
			fmt.Fprint(w, idracAttrBody)
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	attrs, err := provider.GetIDRACAttrDell()
	if err != nil {
		t.Fatalf("GetIDRACAttrDell returned error: %v", err)
	}
	if attrs.SerialCapture_1_Enable != "Enabled" {
		t.Fatalf("expected SerialCapture.1.Enable=Enabled (via standard fallback), got %q", attrs.SerialCapture_1_Enable)
	}
	if len(gotPaths) < 2 || gotPaths[0] != oemIDRACAttrPath || gotPaths[1] != stdIDRACAttrPath {
		t.Fatalf("expected OEM tried then standard, got %v", gotPaths)
	}
}

// TestSetAttributesDellIDRACOEMPreferred verifies that PATCH targets the OEM
// DellAttributes endpoint first and stops there when it returns 200.
func TestSetAttributesDellIDRACOEMPreferred(t *testing.T) {
	var (
		gotMethod string
		gotPaths  []string
		gotBody   string
	)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPaths = append(gotPaths, r.URL.Path)
		buf := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(buf)
		gotBody = string(buf)
		if r.URL.Path == oemIDRACAttrPath {
			fmt.Fprint(w, `{"@Message.ExtendedInfo":[{"Message":"The request completed successfully.","MessageId":"Base.1.12.Success","Severity":"OK"}]}`)
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	msg, err := provider.SetAttributesDell("idrac", []byte(`{"Attributes":{"SerialCapture.1.Enable":"Enabled"}}`))
	if err != nil {
		t.Fatalf("SetAttributesDell returned error: %v", err)
	}
	if msg == "" {
		t.Fatalf("expected non-empty message")
	}
	if gotMethod != http.MethodPatch {
		t.Fatalf("expected PATCH, got %s", gotMethod)
	}
	if len(gotPaths) == 0 || gotPaths[0] != oemIDRACAttrPath {
		t.Fatalf("expected OEM path queried first, got %v", gotPaths)
	}
	for _, p := range gotPaths {
		if p == stdIDRACAttrPath {
			t.Fatalf("standard path should not be PATCHed when OEM succeeds: %v", gotPaths)
		}
	}
	if gotBody == "" {
		t.Fatalf("expected non-empty request body")
	}
}

// TestSetAttributesDellIDRACFallbackToStandard verifies that when the OEM
// endpoint is unavailable (404), the PATCH falls back to the standard path.
func TestSetAttributesDellIDRACFallbackToStandard(t *testing.T) {
	var gotPaths []string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPaths = append(gotPaths, r.URL.Path)
		if r.URL.Path == stdIDRACAttrPath {
			fmt.Fprint(w, `{"@Message.ExtendedInfo":[{"Message":"OK","MessageId":"Base.1.12.Success","Severity":"OK"}]}`)
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	_, err := provider.SetAttributesDell("idrac", []byte(`{"Attributes":{"SerialCapture.1.Enable":"Enabled"}}`))
	if err != nil {
		t.Fatalf("SetAttributesDell returned error: %v", err)
	}
	if len(gotPaths) < 2 || gotPaths[0] != oemIDRACAttrPath || gotPaths[1] != stdIDRACAttrPath {
		t.Fatalf("expected OEM tried then standard, got %v", gotPaths)
	}
}

// TestSetAttributesDellIDRACBothUnavailable verifies an error is returned when
// both endpoints respond with non-success status codes.
func TestSetAttributesDellIDRACBothUnavailable(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	_, err := provider.SetAttributesDell("idrac", []byte(`{"Attributes":{"SerialCapture.1.Enable":"Enabled"}}`))
	if err == nil {
		t.Fatalf("expected error when both endpoints return 404")
	}
}
