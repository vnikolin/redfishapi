package redfishapi

import (
	"encoding/json"
)

// TODO: Need to complete
// func (c *IloClient) GetFwVersion() (string, error) {

// 	url := c.Hostname + "/redfish/v1/Systems/1/FirmwareInventory/"
// 	return "yet to do", nil
// }

func (c *IloClient) GetSystemInfoHP() (*SystemData, error) {

	url := c.Hostname + "/redfish/v1/Systems/1"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return SystemData{}, err
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
