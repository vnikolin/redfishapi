package hp

import (
	"client"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FirmwareInventory struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Current        struct {
		One03c3239103c21c0 []struct {
			ImageSizeBytes  int      `json:"ImageSizeBytes"`
			Key             string   `json:"Key"`
			Location        string   `json:"Location"`
			Name            string   `json:"Name"`
			UEFIDevicePaths []string `json:"UEFIDevicePaths"`
			Updateable      bool     `json:"Updateable"`
			VersionString   string   `json:"VersionString"`
		} `json:"103c3239103c21c0"`
		One0df0720103c220a []struct {
			ImageSizeBytes  int      `json:"ImageSizeBytes"`
			InUse           bool     `json:"InUse"`
			Key             string   `json:"Key"`
			Location        string   `json:"Location"`
			Name            string   `json:"Name"`
			ResetRequired   bool     `json:"ResetRequired"`
			UEFIDevicePaths []string `json:"UEFIDevicePaths"`
			Updateable      bool     `json:"Updateable"`
			VersionString   string   `json:"VersionString"`
		} `json:"10df0720103c220a"`
		One4e41657103c22be []struct {
			ImageSizeBytes  int      `json:"ImageSizeBytes"`
			Key             string   `json:"Key"`
			Location        string   `json:"Location"`
			Name            string   `json:"Name"`
			UEFIDevicePaths []string `json:"UEFIDevicePaths"`
			Updateable      bool     `json:"Updateable"`
			VersionString   string   `json:"VersionString"`
		} `json:"14e41657103c22be"`
		Eight08610fb103c17d3 []struct {
			ImageSizeBytes  int      `json:"ImageSizeBytes"`
			Key             string   `json:"Key"`
			Location        string   `json:"Location"`
			Name            string   `json:"Name"`
			ResetRequired   bool     `json:"ResetRequired"`
			UEFIDevicePaths []string `json:"UEFIDevicePaths"`
			Updateable      bool     `json:"Updateable"`
			VersionString   string   `json:"VersionString"`
		} `json:"808610fb103c17d3"`
		PlatformDefinitionTable []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"PlatformDefinitionTable"`
		PowerManagementController []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"PowerManagementController"`
		PowerManagementControllerBootloader []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"PowerManagementControllerBootloader"`
		PowerSupplies []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"PowerSupplies"`
		SASProgrammableLogicDevice []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SASProgrammableLogicDevice"`
		SPSFirmwareVersionData []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SPSFirmwareVersionData"`
		StorageBattery []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"StorageBattery"`
		SystemBMC []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SystemBMC"`
		SystemProgrammableLogicDevice []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SystemProgrammableLogicDevice"`
		SystemRomActive []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SystemRomActive"`
		SystemRomBackup []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SystemRomBackup"`
		TrustedPlatformModule12 []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"TrustedPlatformModule12"`
	} `json:"Current"`
	ID    string `json:"Id"`
	Name  string `json:"Name"`
	Type  string `json:"Type"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//SystemInfo is a struct which fetches the Overall System High Level info and its a Singleton Resource
type SystemInfo struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Actions      struct {
		ComputerSystemReset struct {
			ResetTypeRedfishAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
			Target                          string   `json:"target"`
		} `json:"#ComputerSystem.Reset"`
	} `json:"Actions"`
	AssetTag         string `json:"AssetTag"`
	AvailableActions []struct {
		Action       string `json:"Action"`
		Capabilities []struct {
			AllowableValues []string `json:"AllowableValues"`
			PropertyName    string   `json:"PropertyName"`
		} `json:"Capabilities"`
	} `json:"AvailableActions"`
	Bios struct {
		Current struct {
			VersionString string `json:"VersionString"`
		} `json:"Current"`
	} `json:"Bios"`
	BiosVersion string `json:"BiosVersion"`
	Boot        struct {
		BootSourceOverrideEnabled   string   `json:"BootSourceOverrideEnabled"`
		BootSourceOverrideSupported []string `json:"BootSourceOverrideSupported"`
		BootSourceOverrideTarget    string   `json:"BootSourceOverrideTarget"`
	} `json:"Boot"`
	Description     string `json:"Description"`
	HostCorrelation struct {
		HostMACAddress []string `json:"HostMACAddress"`
		HostName       string   `json:"HostName"`
		IPAddress      []string `json:"IPAddress"`
	} `json:"HostCorrelation"`
	HostName     string `json:"HostName"`
	ID           string `json:"Id"`
	IndicatorLED string `json:"IndicatorLED"`
	LogServices  struct {
		OdataID string `json:"@odata.id"`
	} `json:"LogServices"`
	Manufacturer string `json:"Manufacturer"`
	Memory       struct {
		Status struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
		TotalSystemMemoryGB int `json:"TotalSystemMemoryGB"`
	} `json:"Memory"`
	MemorySummary struct {
		Status struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
		TotalSystemMemoryGiB int `json:"TotalSystemMemoryGiB"`
	} `json:"MemorySummary"`
	Model string `json:"Model"`
	Name  string `json:"Name"`
	Oem   struct {
		Hp struct {
			OdataType string `json:"@odata.type"`
			Actions   struct {
				HpComputerSystemExtPowerButton struct {
					PushTypeRedfishAllowableValues []string `json:"PushType@Redfish.AllowableValues"`
					Target                         string   `json:"target"`
				} `json:"#HpComputerSystemExt.PowerButton"`
				HpComputerSystemExtSystemReset struct {
					ResetTypeRedfishAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
					Target                          string   `json:"target"`
				} `json:"#HpComputerSystemExt.SystemReset"`
			} `json:"Actions"`
			AvailableActions []struct {
				Action       string `json:"Action"`
				Capabilities []struct {
					AllowableValues []string `json:"AllowableValues"`
					PropertyName    string   `json:"PropertyName"`
				} `json:"Capabilities"`
			} `json:"AvailableActions"`
			Battery []struct {
				Condition       string `json:"Condition"`
				ErrorCode       int    `json:"ErrorCode"`
				FirmwareVersion string `json:"FirmwareVersion"`
				Index           int    `json:"Index"`
				MaxCapWatts     int    `json:"MaxCapWatts"`
				Model           string `json:"Model"`
				Present         string `json:"Present"`
				ProductName     string `json:"ProductName"`
				SerialNumber    string `json:"SerialNumber"`
				Spare           string `json:"Spare"`
			} `json:"Battery"`
			Bios struct {
				Backup struct {
					Date          string `json:"Date"`
					Family        string `json:"Family"`
					VersionString string `json:"VersionString"`
				} `json:"Backup"`
				Current struct {
					Date          string `json:"Date"`
					Family        string `json:"Family"`
					VersionString string `json:"VersionString"`
				} `json:"Current"`
				UefiClass int `json:"UefiClass"`
			} `json:"Bios"`
			DeviceDiscoveryComplete struct {
				AMSDeviceDiscovery  string `json:"AMSDeviceDiscovery"`
				DeviceDiscovery     string `json:"DeviceDiscovery"`
				SmartArrayDiscovery string `json:"SmartArrayDiscovery"`
			} `json:"DeviceDiscoveryComplete"`
			IntelligentProvisioningIndex    int      `json:"IntelligentProvisioningIndex"`
			IntelligentProvisioningLocation string   `json:"IntelligentProvisioningLocation"`
			IntelligentProvisioningVersion  string   `json:"IntelligentProvisioningVersion"`
			PostState                       string   `json:"PostState"`
			PowerAllocationLimit            int      `json:"PowerAllocationLimit"`
			PowerAutoOn                     string   `json:"PowerAutoOn"`
			PowerOnDelay                    string   `json:"PowerOnDelay"`
			PowerRegulatorMode              string   `json:"PowerRegulatorMode"`
			PowerRegulatorModesSupported    []string `json:"PowerRegulatorModesSupported"`
			TrustedModules                  []struct {
				FWVersion struct {
					Current struct {
						MajorVersion int `json:"MajorVersion"`
						MinorVersion int `json:"MinorVersion"`
					} `json:"Current"`
				} `json:"FWVersion"`
				ModuleType string `json:"ModuleType"`
				Status     string `json:"Status"`
			} `json:"TrustedModules"`
			Type           string `json:"Type"`
			VirtualProfile string `json:"VirtualProfile"`
			Links          struct {
				BIOS struct {
					Href string `json:"href"`
				} `json:"BIOS"`
				EthernetInterfaces struct {
					Href string `json:"href"`
				} `json:"EthernetInterfaces"`
				FirmwareInventory struct {
					Href string `json:"href"`
				} `json:"FirmwareInventory"`
				Memory struct {
					Href string `json:"href"`
				} `json:"Memory"`
				NetworkAdapters struct {
					Href string `json:"href"`
				} `json:"NetworkAdapters"`
				PCIDevices struct {
					Href string `json:"href"`
				} `json:"PCIDevices"`
				PCISlots struct {
					Href string `json:"href"`
				} `json:"PCISlots"`
				SecureBoot struct {
					Href string `json:"href"`
				} `json:"SecureBoot"`
				SmartStorage struct {
					Href string `json:"href"`
				} `json:"SmartStorage"`
				SoftwareInventory struct {
					Href string `json:"href"`
				} `json:"SoftwareInventory"`
			} `json:"links"`
		} `json:"Hp"`
	} `json:"Oem"`
	Power            string `json:"Power"`
	PowerState       string `json:"PowerState"`
	ProcessorSummary struct {
		Count  int    `json:"Count"`
		Model  string `json:"Model"`
		Status struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
	} `json:"ProcessorSummary"`
	Processors struct {
		Count           int    `json:"Count"`
		ProcessorFamily string `json:"ProcessorFamily"`
		Status          struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
	} `json:"Processors"`
	SKU          string `json:"SKU"`
	SerialNumber string `json:"SerialNumber"`
	Status       struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	SystemType string `json:"SystemType"`
	Type       string `json:"Type"`
	UUID       string `json:"UUID"`
	Links      struct {
		Chassis []struct {
			Href string `json:"href"`
		} `json:"Chassis"`
		Logs struct {
			Href string `json:"href"`
		} `json:"Logs"`
		ManagedBy []struct {
			Href string `json:"href"`
		} `json:"ManagedBy"`
		Processors struct {
			Href string `json:"href"`
		} `json:"Processors"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

type MacAdresssList struct {
	// FactoryMacs []string `json:"factory_macs"`
	HostMacs []string `json:"host_macs"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// TODO: Need to complete
func (c *client.IloClient) GetFwVersion() (string, error) {

	url := c.Hostname + "/redfish/v1/Systems/1/FirmwareInventory/"
}

func (c *IloClient) GetMacAddress() (string, error) {

	url := c.Hostname + "/redfish/v1/Systems/1"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)

	var x SystemInfo

	json.Unmarshal(_body, &x)

	macAddress := MacAdresssList{
		HostMacs: x.HostCorrelation.HostMACAddress,
	}

	output, _ := json.Marshal(macAddress)
	return string(output), nil

}
