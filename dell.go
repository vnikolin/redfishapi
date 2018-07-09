package redfishapi

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	StatusUnauthorized        = "Unauthorized"
	StatusInternalServerError = "Server Error"
)

type SystemView struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Actions        struct {
		_ComputerSystem_Reset struct {
			ResetType_Redfish_AllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
			Target                            string   `json:"target"`
		} `json:"#ComputerSystem.Reset"`
	} `json:"Actions"`
	AssetTag string `json:"AssetTag"`
	Bios     struct {
		_odata_id string `json:"@odata.id"`
	} `json:"Bios"`
	BiosVersion string `json:"BiosVersion"`
	Boot        struct {
		BootSourceOverrideEnabled                        string   `json:"BootSourceOverrideEnabled"`
		BootSourceOverrideMode                           string   `json:"BootSourceOverrideMode"`
		BootSourceOverrideTarget                         string   `json:"BootSourceOverrideTarget"`
		BootSourceOverrideTarget_Redfish_AllowableValues []string `json:"BootSourceOverrideTarget@Redfish.AllowableValues"`
		UefiTargetBootSourceOverride                     string   `json:"UefiTargetBootSourceOverride"`
	} `json:"Boot"`
	Description        string `json:"Description"`
	EthernetInterfaces struct {
		_odata_id string `json:"@odata.id"`
	} `json:"EthernetInterfaces"`
	HostName     string `json:"HostName"`
	ID           string `json:"Id"`
	IndicatorLED string `json:"IndicatorLED"`
	Links        struct {
		Chassis []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"Chassis"`
		Chassis_odata_count int `json:"Chassis@odata.count"`
		CooledBy            []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"CooledBy"`
		CooledBy_odata_count int `json:"CooledBy@odata.count"`
		ManagedBy            []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"ManagedBy"`
		ManagedBy_odata_count int `json:"ManagedBy@odata.count"`
		Oem                   struct {
			Dell struct {
				_odata_type string `json:"@odata.type"`
				BootOrder   struct {
					_odata_id string `json:"@odata.id"`
				} `json:"BootOrder"`
			} `json:"Dell"`
		} `json:"Oem"`
		PoweredBy []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"PoweredBy"`
		PoweredBy_odata_count int `json:"PoweredBy@odata.count"`
	} `json:"Links"`
	Manufacturer  string `json:"Manufacturer"`
	MemorySummary struct {
		MemoryMirroring string `json:"MemoryMirroring"`
		Status          struct {
			Health       string `json:"Health"`
			HealthRollup string `json:"HealthRollup"`
			State        string `json:"State"`
		} `json:"Status"`
		TotalSystemMemoryGiB int `json:"TotalSystemMemoryGiB"`
	} `json:"MemorySummary"`
	Model            string `json:"Model"`
	Name             string `json:"Name"`
	PartNumber       string `json:"PartNumber"`
	PowerState       string `json:"PowerState"`
	ProcessorSummary struct {
		Count  int    `json:"Count"`
		Model  string `json:"Model"`
		Status struct {
			Health       string `json:"Health"`
			HealthRollup string `json:"HealthRollup"`
			State        string `json:"State"`
		} `json:"Status"`
	} `json:"ProcessorSummary"`
	Processors struct {
		_odata_id string `json:"@odata.id"`
	} `json:"Processors"`
	SKU        string `json:"SKU"`
	SecureBoot struct {
		_odata_id string `json:"@odata.id"`
	} `json:"SecureBoot"`
	SerialNumber  string `json:"SerialNumber"`
	SimpleStorage struct {
		_odata_id string `json:"@odata.id"`
	} `json:"SimpleStorage"`
	Status struct {
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
		State        string `json:"State"`
	} `json:"Status"`
	SystemType     string `json:"SystemType"`
	TrustedModules []struct {
		InterfaceType string `json:"InterfaceType"`
		Status        struct {
			State string `json:"State"`
		} `json:"Status"`
	} `json:"TrustedModules"`
	UUID string `json:"UUID"`
}

type GetMacAddress struct {
	_odata_context                     string        `json:"@odata.context"`
	_odata_id                          string        `json:"@odata.id"`
	_odata_type                        string        `json:"@odata.type"`
	AutoNeg                            bool          `json:"AutoNeg"`
	Description                        string        `json:"Description"`
	FQDN                               interface{}   `json:"FQDN"`
	FullDuplex                         bool          `json:"FullDuplex"`
	HostName                           interface{}   `json:"HostName"`
	IPv4Addresses                      []interface{} `json:"IPv4Addresses"`
	IPv4Addresses_odata_count          int           `json:"IPv4Addresses@odata.count"`
	IPv6AddressPolicyTable             []interface{} `json:"IPv6AddressPolicyTable"`
	IPv6AddressPolicyTable_odata_count int           `json:"IPv6AddressPolicyTable@odata.count"`
	IPv6Addresses                      []interface{} `json:"IPv6Addresses"`
	IPv6Addresses_odata_count          int           `json:"IPv6Addresses@odata.count"`
	IPv6DefaultGateway                 interface{}   `json:"IPv6DefaultGateway"`
	IPv6StaticAddresses                []interface{} `json:"IPv6StaticAddresses"`
	IPv6StaticAddresses_odata_count    int           `json:"IPv6StaticAddresses@odata.count"`
	ID                                 string        `json:"Id"`
	InterfaceEnabled                   interface{}   `json:"InterfaceEnabled"`
	MACAddress                         string        `json:"MACAddress"`
	MTUSize                            interface{}   `json:"MTUSize"`
	MaxIPv6StaticAddresses             interface{}   `json:"MaxIPv6StaticAddresses"`
	Name                               string        `json:"Name"`
	NameServers                        []interface{} `json:"NameServers"`
	NameServers_odata_count            int           `json:"NameServers@odata.count"`
	PermanentMACAddress                string        `json:"PermanentMACAddress"`
	SpeedMbps                          int           `json:"SpeedMbps"`
	Status                             struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	UefiDevicePath string `json:"UefiDevicePath"`
	VLAN           string `json:"VLAN"`
}

type MemberCount struct {
	OdataContext string `json:"@odata.context"`
	OdataId      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Description  string `json:"Description"`
	Members      []struct {
		OdataId string `json:"@odata.id"`
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
}

type FirmwareDataDell struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Description    string `json:"Description"`
	ID             string `json:"Id"`
	Name           string `json:"Name"`
	Status         struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	Updateable bool   `json:"Updateable"`
	Version    string `json:"Version"`
}

type MACData struct {
	MacAddresss string `json:"macaddress"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	State       string `json:"state"`
	Vlan        string `json:"vlan"`
}

type HealthData struct {
	Name   string `json:"name"`
	Id     string `json:"id"`
	State  string `json:"state"`
	Health string `json:"health"`
}

type FirmwareData struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Version    string `json:"version"`
	Updateable bool   `json:"updateable"`
}

func (c *IloClient) CheckLoginDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) == true {
			err := errors.New(StatusInternalServerError)
			return "", err
		} else {
			return "", err
		}
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New(StatusUnauthorized)
			return "", err
		}

	}

	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)

	var data SystemView

	json.Unmarshal(_body, &data)

	return string(data.Status.Health), nil
}

func (c *IloClient) GetMacAddressDell() (string, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/EthernetInterfaces/"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) == true {
			err := errors.New(StatusInternalServerError)
			return "", err
		} else {
			return "", err
		}
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New(StatusUnauthorized)
			return "", err
		}

	}

	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)

	var x MemberCount
	var Macs []MACData

	json.Unmarshal(_body, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		req, err := http.NewRequest("GET", _url, nil)
		req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		_body, _ := ioutil.ReadAll(resp.Body)

		var y GetMacAddress

		json.Unmarshal(_body, &y)

		macData := MACData{
			Name:        y.Name,
			Description: y.Description,
			MacAddresss: y.MACAddress,
			Status:      y.Status.Health,
			State:       y.Status.State,
			Vlan:        y.VLAN,
		}
		Macs = append(Macs, macData)

	}

	output, _ := json.Marshal(Macs)

	return string(output), nil

}

func (c *IloClient) GetHealthDataDell() (string, error) {

	url := c.Hostname + "/redfish/v1/UpdateService/FirmwareInventory"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) == true {
			err := errors.New("Server Error")
			return "", err
		} else {
			return "", err
		}
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New("Unauthorized")
			return "", err
		}

	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)

	var (
		x           MemberCount
		_healthdata []HealthData
	)

	json.Unmarshal(_body, &x)

	for i := range x.Members {
		r, _ := regexp.Compile("Installed")
		if r.MatchString(x.Members[i].OdataId) == true {
			_url := c.Hostname + x.Members[i].OdataId
			req, err := http.NewRequest("GET", _url, nil)
			req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()

			_body, _ := ioutil.ReadAll(resp.Body)

			var y FirmwareDataDell

			json.Unmarshal(_body, &y)

			healthData := HealthData{
				Name:   y.Name,
				Id:     y.ID,
				State:  y.Status.State,
				Health: y.Status.Health,
			}

			_healthdata = append(_healthdata, healthData)

		}
	}

	output, _ := json.Marshal(_healthdata)

	return string(output), nil

}

func (c *IloClient) GetFirmwareDell() (string, error) {

	url := c.Hostname + "/redfish/v1/UpdateService/FirmwareInventory"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) == true {
			err := errors.New("Server Error")
			return "", err
		} else {
			return "", err
		}
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New("Unauthorized")
			return "", err
		}

	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)

	var (
		x         MemberCount
		_firmdata []FirmwareData
	)

	json.Unmarshal(_body, &x)

	for i := range x.Members {
		r, _ := regexp.Compile("Installed")
		if r.MatchString(x.Members[i].OdataId) == true {
			_url := c.Hostname + x.Members[i].OdataId
			req, err := http.NewRequest("GET", _url, nil)
			req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()

			_body, _ := ioutil.ReadAll(resp.Body)

			var y FirmwareDataDell

			json.Unmarshal(_body, &y)

			firmData := FirmwareData{
				Name:       y.Name,
				Id:         y.ID,
				Version:    y.Version,
				Updateable: y.Updateable,
			}

			_firmdata = append(_firmdata, firmData)

		}
	}

	output, _ := json.Marshal(_firmdata)

	return string(output), nil

}
