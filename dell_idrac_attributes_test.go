package redfishapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIDRACAttrDellSerialCaptureFields(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/redfish/v1/Managers/iDRAC.Embedded.1/Attributes" {
			fmt.Fprint(w, `{
				"Attributes": {
					"SerialCapture.1.Enable": "Enabled",
					"SerialCapture.1.SerialDataSize": 524288,
					"CurrentNIC.1.MACAddress": "aa:bb:cc:dd:ee:ff"
				}
			}`)
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
}

func TestSetAttributesDellIDRACPath(t *testing.T) {
	var (
		gotMethod string
		gotPath   string
		gotBody   string
	)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		buf := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(buf)
		gotBody = string(buf)
		fmt.Fprint(w, `{"@Message.ExtendedInfo":[{"Message":"The request completed successfully.","MessageId":"Base.1.12.Success","Severity":"OK"}]}`)
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
	if gotPath != "/redfish/v1/Managers/iDRAC.Embedded.1/Attributes" {
		t.Fatalf("unexpected path %q", gotPath)
	}
	if gotBody == "" {
		t.Fatalf("expected non-empty request body")
	}
}
