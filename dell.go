package client

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
	UefiDevicePath string      `json:"UefiDevicePath"`
	VLAN           interface{} `json:"VLAN"`
}

type EtherInterfaces struct {
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

type MACData struct {
	MacAddresss []string `json:"macaddress"`
}

func (c *IloClient) GetMacAddressDell() (string, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/EthernetInterfaces/"

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

	var x EtherInterfaces
	var Macs []string

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

		Macs = append(Macs, y.MACAddress)

	}

	mac_address := MACData{
		MacAddresss: Macs,
	}

	output, _ := json.Marshal(mac_address)

	return string(output), nil

}
