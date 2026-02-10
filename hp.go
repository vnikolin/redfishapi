package redfishapi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// StartServerHP ...
// ResetType@Redfish.AllowableValues
// 0	"On"
// 1	"ForceOff"
// 2	"ForceRestart",
// 3	"Nmi",
// 4	"PushPowerButton"
// target: "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset/"
func (c *redfishProvider) StartServerHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset/"
	var jsonStr = []byte(`{"ResetType": "On"}`)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Started", nil
}

// StopServerHP ... Will Request to stop the server
func (c *redfishProvider) StopServerHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset/"
	var jsonStr = []byte(`{"ResetType": "ForceOff"}`)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Stopped", nil
}

// GetSystemInfoHP ... Will fetch the system info
func (c *redfishProvider) GetSystemInfoHP() (SystemData, error) {

	url := c.Hostname + "/redfish/v1/Systems/1"

	resp, _, _, err := queryData(c, "GET", url, nil)
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

// GetServerPowerStateHP ... Will fetch the current state of the Server
func (c *redfishProvider) GetServerPowerStateHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}

	var data SystemInfoHP

	json.Unmarshal(resp, &data)

	return data.Power, nil

}

// CheckLoginHP ... Will check the credentials of the Server
func (c *redfishProvider) CheckLoginHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}
	var data SystemInfoHP
	json.Unmarshal(resp, &data)
	return string(data.Status.Health), nil
}

// GetFirmwareHP ... will fetch the Firmware details
func (c *redfishProvider) GetFirmwareHP() ([]FirmwareData, error) {

	url := c.Hostname + "/redfish/v1/Systems/1/FirmwareInventory/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var (
		x         FirmwareInventoryHP
		_firmdata []FirmwareData
	)
	json.Unmarshal(resp, &x)

	for i := range x.Current.One03c3239103c21c0 {
		_result := FirmwareData{
			Id:         x.Current.One03c3239103c21c0[i].Key,
			Name:       x.Current.One03c3239103c21c0[i].Name,
			Updateable: x.Current.One03c3239103c21c0[i].Updateable,
			Version:    x.Current.One03c3239103c21c0[i].VersionString,
		}
		_firmdata = append(_firmdata, _result)
	}

	for i := range x.Current.One4e41657103c22be {
		_result := FirmwareData{
			Id:         x.Current.One4e41657103c22be[i].Key,
			Name:       x.Current.One4e41657103c22be[i].Name,
			Updateable: x.Current.One4e41657103c22be[i].Updateable,
			Version:    x.Current.One4e41657103c22be[i].VersionString,
		}
		_firmdata = append(_firmdata, _result)
	}

	for i := range x.Current.Eight08610fb103c17d0 {
		_result := FirmwareData{
			Id:         x.Current.Eight08610fb103c17d0[i].Key,
			Name:       x.Current.Eight08610fb103c17d0[i].Name,
			Updateable: x.Current.Eight08610fb103c17d0[i].Updateable,
			Version:    x.Current.Eight08610fb103c17d0[i].VersionString,
		}
		_firmdata = append(_firmdata, _result)
	}

	for i := range x.Current.Eight08610fb103c17d3 {
		_result := FirmwareData{
			Id:         x.Current.Eight08610fb103c17d3[i].Key,
			Name:       x.Current.Eight08610fb103c17d3[i].Name,
			Updateable: x.Current.Eight08610fb103c17d3[i].Updateable,
			Version:    x.Current.Eight08610fb103c17d3[i].VersionString,
		}
		_firmdata = append(_firmdata, _result)
	}

	for i := range x.Current.SystemBMC {
		_result := FirmwareData{
			Id:         x.Current.SystemBMC[i].Key,
			Name:       x.Current.SystemBMC[i].Name,
			Updateable: false, //Adding False as default because of redfish have no field
			Version:    x.Current.SystemBMC[i].VersionString,
		}
		_firmdata = append(_firmdata, _result)
	}

	for i := range x.Current.SystemRomActive {
		_result := FirmwareData{
			Id:         x.Current.SystemRomActive[i].Key,
			Name:       x.Current.SystemRomActive[i].Name,
			Updateable: false, //Adding False as default because of redfish have no field
			Version:    x.Current.SystemRomActive[i].VersionString,
		}
		_firmdata = append(_firmdata, _result)
	}

	for i := range x.Current.SystemRomBackup {
		_result := FirmwareData{
			Id:         x.Current.SystemRomBackup[i].Key,
			Name:       x.Current.SystemRomBackup[i].Name,
			Updateable: false, //Adding False as default because of redfish have no field
			Version:    x.Current.SystemRomBackup[i].VersionString,
		}
		_firmdata = append(_firmdata, _result)
	}

	return _firmdata, nil
}

// GetThermalHealthHP ... will fetch the Thermal Health
func (c *redfishProvider) GetThermalHealthHP() ([]HealthList, error) {
	url := c.Hostname + "/redfish/v1/Chassis/1/Thermal/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x       ThermalHealthListHP
		_health []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.Fans {
		_result := HealthList{Name: x.Fans[i].FanName,
			Health: x.Fans[i].Status.Health,
			State:  x.Fans[i].Status.State}
		_health = append(_health, _result)
	}

	for i := range x.Temperatures {
		_result := HealthList{Name: x.Temperatures[i].Name,
			Health: x.Temperatures[i].Status.Health,
			State:  x.Temperatures[i].Status.State}
		_health = append(_health, _result)
	}

	return _health, nil
}

// GetPowerHealthHP ... will fetch the Power Health
func (c *redfishProvider) GetPowerHealthHP() ([]HealthList, error) {
	url := c.Hostname + "/redfish/v1/Chassis/1/Power/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x       PowerDataHP
		_health []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.PowerSupplies {
		_name := fmt.Sprintf("%s_%d", x.PowerSupplies[i].Name, i)
		_result := HealthList{Name: _name,
			Health: x.PowerSupplies[i].Status.Health,
			State:  x.PowerSupplies[i].Status.State}
		_health = append(_health, _result)
	}

	return _health, nil
}

// GetInterfaceHealthHP ... will fetch the Interface Health
func (c *redfishProvider) GetInterfaceHealthHP() ([]HealthList, error) {
	url := c.Hostname + "/redfish/v1/Managers/1/EthernetInterfaces/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x       EthernetInterfacesHP
		_health []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.Items {
		_result := HealthList{Name: x.Items[i].Name,
			Health: x.Items[i].Status.Health,
			State:  x.Items[i].Status.State}
		_health = append(_health, _result)
	}

	return _health, nil
}

// GetProcessorHealthHP ... will Fetch the Processor Health Details
func (c *redfishProvider) GetProcessorInfoHP() ([]ProcessorInfoHP, error) {

	url := c.Hostname + "/redfish/v1/Systems/1/Processors/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x           MemberCountHP
		processData []ProcessorInfoHP
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataID
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}

		var y ProcessorInfoHP

		json.Unmarshal(resp, &y)

		processData = append(processData, y)
	}

	return processData, nil

}

// GetProcessorHealthHP ... will Fetch the Processor Health Details
func (c *redfishProvider) GetProcessorHealthHP() ([]HealthList, error) {

	url := c.Hostname + "/redfish/v1/Systems/1/Processors/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x             MemberCountHP
		processHealth []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataID
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}

		var y ProcessorInfoHP

		json.Unmarshal(resp, &y)

		procHealth := HealthList{
			Name:   y.ID,
			Health: y.Status.Health,
			State:  y.Oem.Hp.ConfigStatus.State,
		}
		processHealth = append(processHealth, procHealth)
	}

	return processHealth, nil

}

// GetUserAccountsHP ... will fetch the current User Accounts
func (c *redfishProvider) GetUserAccountsHP() ([]Accounts, error) {

	url := c.Hostname + "/redfish/v1/AccountService/Accounts"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x       AccountsInfoHP
		users   []Accounts
		_locked bool
	)

	json.Unmarshal(resp, &x)

	for i := range x.Items {

		if x.Items[i].Oem.Hp.Privileges.LoginPriv {
			_locked = false
		} else {
			_locked = true
		}

		user := Accounts{
			Name:     x.Items[i].Name,
			Enabled:  x.Items[i].Oem.Hp.Privileges.LoginPriv,
			Locked:   _locked,
			RoleId:   x.Items[i].ID,
			Username: x.Items[i].UserName,
		}
		users = append(users, user)

	}

	return users, nil

}

// GetSystemEventLogsHP ... will fetch the SystemEvent Logs
func (c *redfishProvider) GetSystemEventLogsHP() ([]SystemEventLogRes, error) {

	url := c.Hostname + "/redfish/v1/Managers/1/LogServices/IEL/Entries/"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var x SystemEventLogsHP

	json.Unmarshal(resp, &x)

	var _systemEventLogs []SystemEventLogRes

	for i := range x.Items {

		_result := SystemEventLogRes{
			EntryCode:  x.Items[i].EntryType,
			Message:    x.Items[i].Message,
			Name:       x.Items[i].Name,
			SensorType: x.Items[i].Type,
			Severity:   x.Items[i].Severity,
		}

		_systemEventLogs = append(_systemEventLogs, _result)
	}

	return _systemEventLogs, nil

}

// GetBiosDataHP ... will fetch the Bios Details
func (c *redfishProvider) GetBiosDataHP() (BiosDataHP, error) {

	url := c.Hostname + "/redfish/v1/systems/1/bios/settings/"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return BiosDataHP{}, err
	}

	var biosData BiosDataHP

	json.Unmarshal(resp, &biosData)

	return biosData, nil
}

// GetLicenseInfoHP ... will fetch the current License Details
func (c *redfishProvider) GetLicenseInfoHP() (LicenseInfo, error) {

	url := c.Hostname + "/redfish/v1/Managers/1/LicenseService/"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return LicenseInfo{}, err
	}

	var x LicenseInfoHP

	json.Unmarshal(resp, &x)

	if len(x.Items) == 0 {
		return LicenseInfo{}, fmt.Errorf("GetLicenseInfoHP: no license items returned")
	}

	_result := LicenseInfo{
		Name:        x.Name,
		LicenseKey:  x.Items[0].LicenseKey,
		LicenseType: x.Items[0].LicenseType,
	}

	return _result, nil
}

// GetPCISlotsHP ... will fetch the PCI Slots Details
func (c *redfishProvider) GetPCISlotsHP() ([]PCISlotsInfo, error) {

	url := c.Hostname + "/redfish/v1/Systems/1/PCISlots/"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var x PCISlotsInfoHP

	json.Unmarshal(resp, &x)

	var _pciSlots []PCISlotsInfo

	for i := range x.Items {
		_result := PCISlotsInfo{
			Name:   x.Items[i].Name,
			Status: x.Items[i].Status.OperationalStatus[0].Status,
		}
		_pciSlots = append(_pciSlots, _result)
	}

	return _pciSlots, nil

}

// GetEthernetInterfacesHP ... will fetch the EthernetInterfaces Details
func (c *redfishProvider) GetEthernetInterfacesHP() ([]MACData, error) {

	url := c.Hostname + "/redfish/v1/Managers/1/EthernetInterfaces/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x        EthernetInterfacesHP
		_macData []MACData
	)

	json.Unmarshal(resp, &x)

	for i := range x.Items {
		_result := MACData{
			Name:        x.Items[i].Name,
			Description: x.Items[i].Description,
			MacAddress:  x.Items[i].MacAddress,
			State:       x.Items[i].Status.State,
			Status:      strconv.FormatBool(x.Items[i].Oem.Hp.NICEnabled),
			Vlan:        "Null",
		}
		_macData = append(_macData, _result)
	}
	return _macData, nil

}
