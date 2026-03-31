package redfishapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStorageControllersReturnsStorageMembers(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/redfish/v1/Systems/System.Embedded.1/Storage":
			fmt.Fprint(w, `{
				"@odata.context": "/redfish/v1/$metadata#StorageCollection.StorageCollection",
				"@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage",
				"@odata.type": "#StorageCollection.StorageCollection",
				"Description": "Collection Of Storage entities",
				"Members": [
					{"@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage/RAID.SL.3-1"},
					{"@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.2-1"},
					{"@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.1-1"},
					{"@odata.id": "/redfish/v1/Systems/System.Embedded.1/Storage/CPU.1"}
				],
				"Members@odata.count": 4,
				"Name": "Storage Collection"
			}`)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	provider := &redfishProvider{Hostname: server.URL, Username: "user", Password: "pass"}
	members, err := provider.GetStorageControllers()
	if err != nil {
		t.Fatalf("GetStorageControllers returned error: %v", err)
	}

	if len(members) != 4 {
		t.Fatalf("expected 4 storage members, got %d", len(members))
	}

	expected := []string{
		"/redfish/v1/Systems/System.Embedded.1/Storage/RAID.SL.3-1",
		"/redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.2-1",
		"/redfish/v1/Systems/System.Embedded.1/Storage/AHCI.Embedded.1-1",
		"/redfish/v1/Systems/System.Embedded.1/Storage/CPU.1",
	}

	for i, want := range expected {
		if members[i].OdataId != want {
			t.Fatalf("unexpected member %d: got %q want %q", i, members[i].OdataId, want)
		}
	}
}
