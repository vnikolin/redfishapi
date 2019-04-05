package redfishapi

import (
	"encoding/json"
)

// ResetType@Redfish.AllowableValues
// 0	"On"
// 1	"ForceOff"
// 2	"ForceRestart",
// 3	"Nmi",
// 4	"PushPowerButton"
// target: "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset/"
func (c *IloClient) StartServerHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset/"
	var jsonStr = []byte(`{"ResetType": "On"}`)
	_, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Started", nil
}

func (c *IloClient) StopServerHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset/"
	var jsonStr = []byte(`{"ResetType": "ForceOff"}`)
	_, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Stopped", nil
}

func (c *IloClient) GetSystemInfoHP() (SystemData, error) {

	url := c.Hostname + "/redfish/v1/Systems/1"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var x SystemInfoHP

	json.Unmarshal(resp, &x)

	_result := SystemData{Health: x.Status.Health,
		Memory:          x.Memory.TotalSystemMemoryGB,
		Model:           x.Model,
		PowerState:      x.PowerState,
		Processors:      x.Processors.Count,
		ProcessorFamily: x.Processors.ProcessorFamily,
		SerialNumber:    x.SerialNumber,
	}

	return _result, nil

}

// TODO: Need to complete
// func (c *IloClient) GetFwVersion() (string, error) {

// 	url := c.Hostname + "/redfish/v1/Systems/1/FirmwareInventory/"
// 	return "yet to do", nil
// }
