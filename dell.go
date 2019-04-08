package redfishapi

import (
	"encoding/json"
	"regexp"

	ver "github.com/hashicorp/go-version"
)

// Declaring the Constant Values
const (
	StatusUnauthorized        = "Unauthorized"
	StatusInternalServerError = "Server Error"
)

//StartServerDell ...
// ResetType@Redfish.AllowableValues
// 0	"On"
// 1	"ForceOff"
// 2	"GracefulRestart"
// 3	"GracefulShutdown"
// 4	"PushPowerButton"
// 5	"Nmi"
// target: "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"
func (c *IloClient) StartServerDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"

	var jsonStr = []byte(`{"ResetType": "On"}`)
	_, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Started", nil
}

//StopServerDell ... Will Request to stop the server
func (c *IloClient) StopServerDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"

	var jsonStr = []byte(`{"ResetType": "ForceOff"}`)
	_, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Stopped", nil
}

//GetServerPowerStateDell ... Will fetch the current state of the Server
func (c *IloClient) GetServerPowerStateDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1"
	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}

	var data SystemViewDell

	json.Unmarshal(resp, &data)

	return data.PowerState, nil

}

// func (c *IloClient) SetBiosAttributes() (string, error) {}

//CheckLoginDell ... Will check the credentials of the Server
func (c *IloClient) CheckLoginDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1"
	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}
	var data SystemViewDell
	json.Unmarshal(resp, &data)
	return string(data.Status.Health), nil
}

//GetMacAddressDell ... Will fetch all the mac address of a particular Server
func (c *IloClient) GetMacAddressDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/EthernetInterfaces/"
	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}
	var x MemberCountDell
	var Macs []MACData
	json.Unmarshal(resp, &x)
	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return "", err
		}
		var y GetMacAddressDell
		json.Unmarshal(resp, &y)
		macData := MACData{
			Name:        y.Name,
			Description: y.Description,
			MacAddress:  y.MACAddress,
			Status:      y.Status.Health,
			State:       y.Status.State,
			Vlan:        y.VLAN,
		}
		Macs = append(Macs, macData)
	}
	output, _ := json.Marshal(Macs)
	return string(output), nil
}

//GetProcessorHealthDell ... Will Fetch the Processor Health Details
func (c *IloClient) GetProcessorHealthDell() ([]HealthList, error) {
	///redfish/v1/Systems/System.Embedded.1/Processors

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Processors"
	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x             ProcessorsListDataDell
		processHealth []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}

		var y ProcessorDataDell

		json.Unmarshal(resp, &y)

		procHealth := HealthList{
			Name:   y.ID,
			Health: y.Status.Health,
			State:  y.Status.State,
		}
		processHealth = append(processHealth, procHealth)
	}

	return processHealth, nil

}

// func (c *IloClient) GetMemoryHealthDell() (string, error) {}

//GetPowerHealthDell ... Will Fetch the Power Health Details
func (c *IloClient) GetPowerHealthDell() ([]HealthList, error) {
	url := c.Hostname + "/redfish/v1/Chassis/System.Embedded.1/Power"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x             PowerDataDell
		powerSupplies []HealthList
	)

	json.Unmarshal(resp, &x)

	if x.PowerSuppliescount != 0 {
		for i := range x.PowerSupplies {
			powerControlHealth := HealthList{
				Name:   x.PowerSupplies[i].MemberID,
				Health: x.PowerSupplies[i].Status.Health,
				State:  x.PowerSupplies[i].Status.State,
			}
			powerSupplies = append(powerSupplies, powerControlHealth)
		}
	}

	if x.Redundancycount != 0 {
		for i := range x.Redundancy {
			redundHealth := HealthList{
				Name:   x.Redundancy[i].Name,
				Health: x.Redundancy[i].Status.Health,
				State:  x.Redundancy[i].Status.State,
			}
			powerSupplies = append(powerSupplies, redundHealth)
		}
	}

	if x.Voltagescount != 0 {
		for i := range x.Voltages {
			voltageHealth := HealthList{
				Name:   x.Voltages[i].Name,
				Health: x.Voltages[i].Status.Health,
				State:  x.Voltages[i].Status.State,
			}
			powerSupplies = append(powerSupplies, voltageHealth)
		}
	}

	return powerSupplies, nil
}

//GetSensorsHealthDell ... Will Fetch the Sensors Health Details
func (c *IloClient) GetSensorsHealthDell() ([]HealthList, error) {

	url := c.Hostname + "/redfish/v1/Chassis/System.Embedded.1/Thermal"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x             ThermalHealthListDell
		thermalHealth []HealthList
	)

	json.Unmarshal(resp, &x)

	// Fetching the Redundancy health info
	if x.Redundancycount != 0 {
		for i := range x.Redundancy {
			redundHealth := HealthList{
				Name:   x.Redundancy[i].Name,
				Health: x.Redundancy[i].Status.Health,
				State:  x.Redundancy[i].Status.State,
			}
			thermalHealth = append(thermalHealth, redundHealth)
		}
	}

	if x.Fanscount != 0 {
		for i := range x.Fans {
			fanHealth := HealthList{
				Name:   x.Fans[i].Name,
				Health: x.Fans[i].Status.Health,
				State:  x.Fans[i].Status.State,
			}
			thermalHealth = append(thermalHealth, fanHealth)
		}
	}

	if x.Temperaturesount != 0 {
		for i := range x.Temperatures {
			tempData := HealthList{
				Name:   x.Temperatures[i].Name,
				Health: x.Temperatures[i].Status.Health,
				State:  x.Temperatures[i].Status.State,
			}
			thermalHealth = append(thermalHealth, tempData)
		}
	}

	return thermalHealth, nil

}

//GetStorageHealthDell ... Will Fetch the Storage Health Details
func (c *IloClient) GetStorageHealthDell() ([]StorageHealthList, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Storage"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x           StorageCollectionDell
		_healthdata []StorageHealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {

		_url := c.Hostname + x.Members[i].OdataId
		resp, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}

		var y StorageDetailsDell

		json.Unmarshal(resp, &y)

		storageHealth := StorageHealthList{
			Name:   y.ID,
			Health: y.Status.Health,
			State:  y.Status.State,
			Space:  0,
		}
		_healthdata = append(_healthdata, storageHealth)

		if y.Devicescount != 0 {
			for k := range y.Devices {
				storageHealth := StorageHealthList{
					Name:   y.Devices[k].Name,
					Health: y.Devices[k].Status.Health,
					State:  y.Devices[k].Status.State,
					Space:  y.Devices[k].CapacityBytes,
				}
				_healthdata = append(_healthdata, storageHealth)
			}

		} else {
			continue
		}

	}
	return _healthdata, nil

}

//GetAggHealthDataDell ... will fetch the data related to all components health(aggregated view)
func (c *IloClient) GetAggHealthDataDell() ([]HealthList, error) {

	url := c.Hostname + "/redfish/v1/UpdateService/FirmwareInventory"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x           MemberCountDell
		_healthdata []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		r, _ := regexp.Compile("Installed")
		if r.MatchString(x.Members[i].OdataId) == true {
			_url := c.Hostname + x.Members[i].OdataId
			resp, err := queryData(c, "GET", _url, nil)
			if err != nil {
				return nil, err
			}

			var y FirmwareDataDell

			json.Unmarshal(resp, &y)

			healthData := HealthList{
				Name:   y.Name,
				State:  y.Status.State,
				Health: y.Status.Health,
			}

			_healthdata = append(_healthdata, healthData)

		}
	}

	return _healthdata, nil

}

//GetFirmwareDell ... will fetch the Firmware details
func (c *IloClient) GetFirmwareDell() ([]FirmwareData, error) {

	url := c.Hostname + "/redfish/v1/UpdateService/FirmwareInventory"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x         MemberCountDell
		_firmdata []FirmwareData
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		r, _ := regexp.Compile("Installed")
		if r.MatchString(x.Members[i].OdataId) == true {
			_url := c.Hostname + x.Members[i].OdataId
			resp, err := queryData(c, "GET", _url, nil)
			if err != nil {
				return nil, err
			}

			var y FirmwareDataDell

			json.Unmarshal(resp, &y)

			firmData := FirmwareData{
				Name:       y.Name,
				Id:         y.ID,
				Version:    y.Version,
				Updateable: y.Updateable,
			}

			_firmdata = append(_firmdata, firmData)

		}
	}

	return _firmdata, nil

}

//GetBiosDataDell ... will fetch the Bios Details
func (c *IloClient) GetBiosDataDell() (BiosData, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Bios"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return BiosData{}, err
	}

	var x BiosAttrDell

	json.Unmarshal(resp, &x)

	_data := x.Attributes

	_BiosData := BiosData{
		BootMode:          _data.BootMode,
		BootSeqRetry:      _data.BootSeqRetry,
		InternalUsb:       _data.InternalUsb,
		SriovGlobalEnable: _data.SriovGlobalEnable,
		SysProfile:        _data.SysProfile,
		AcPwrRcvry:        _data.AcPwrRcvry,
		AcPwrRcvryDelay:   _data.AcPwrRcvryDelay,
		Serial:            _data.SystemServiceTag,
	}

	return _BiosData, nil

}

//GetLifecycleAttrDell ... will fetch the lifecycle attributes
func (c *IloClient) GetLifecycleAttrDell() (LifeCycleData, error) {

	url := c.Hostname + "/redfish/v1/Managers/LifecycleController.Embedded.1/Attributes"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return LifeCycleData{}, err
	}

	var x LifeCycleAttrDell

	json.Unmarshal(resp, &x)

	_data := x.Attributes

	_LfcycleDat := LifeCycleData{
		AutoBackup:                          _data.LCAttributes_1_AutoBackup,
		AutoDiscovery:                       _data.LCAttributes_1_AutoDiscovery,
		BIOSRTDRequested:                    _data.LCAttributes_1_BIOSRTDRequested,
		CollectSystemInventoryOnRestart:     _data.LCAttributes_1_CollectSystemInventoryOnRestart,
		DiscoveryFactoryDefaults:            _data.LCAttributes_1_DiscoveryFactoryDefaults,
		IPAddress:                           _data.LCAttributes_1_IPAddress,
		IPChangeNotifyPS:                    _data.LCAttributes_1_IPChangeNotifyPS,
		IgnoreCertWarning:                   _data.LCAttributes_1_IgnoreCertWarning,
		Licensed:                            _data.LCAttributes_1_Licensed,
		LifecycleControllerState:            _data.LCAttributes_1_LifecycleControllerState,
		PartConfigurationUpdate:             _data.LCAttributes_1_PartConfigurationUpdate,
		PartFirmwareUpdate:                  _data.LCAttributes_1_PartFirmwareUpdate,
		ProvisioningServer:                  _data.LCAttributes_1_ProvisioningServer,
		StorageHealthRollupStatus:           _data.LCAttributes_1_StorageHealthRollupStatus,
		SystemID:                            _data.LCAttributes_1_SystemID,
		UserProxyPassword:                   _data.LCAttributes_1_UserProxyPassword,
		UserProxyPort:                       _data.LCAttributes_1_UserProxyPort,
		UserProxyServer:                     _data.LCAttributes_1_UserProxyServer,
		UserProxyType:                       _data.LCAttributes_1_UserProxyType,
		UserProxyUserName:                   _data.LCAttributes_1_UserProxyUserName,
		VirtualAddressManagementApplication: _data.LCAttributes_1_VirtualAddressManagementApplication,
	}

	return _LfcycleDat, nil

}

//GetIDRACAttrDell ... will fetch the Idrac attributes
func (c *IloClient) GetIDRACAttrDell() (IDRACData, error) {

	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Attributes"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return IDRACData{}, err
	}

	var x IDRACAttrDell

	json.Unmarshal(resp, &x)

	_data := x.Attributes

	_idracData := IDRACData{
		VirtualConsoleMaxSessions: _data.VirtualConsole_1_MaxSessions,
		VirtualConsolePluginType:  _data.VirtualConsole_1_PluginType,
		WebServerSSLEncryption:    _data.WebServer_1_SSLEncryptionBitLength,
		IPMILanEnable:             _data.IPMILan_1_Enable,
		DNSDomainName:             _data.NIC_1_DNSDomainName,
		SnmpAgentStatus:           _data.SNMP_1_AgentEnable,
		SnmpAgentCommunity:        _data.SNMP_1_AgentCommunity,
	}

	return _idracData, nil

}

//GetSysAttrDell ... will fetch the System Attributes
func (c *IloClient) GetSysAttrDell() (SysAttrData, error) {

	url := c.Hostname + "/redfish/v1/Managers/System.Embedded.1/Attributes"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return SysAttrData{}, err
	}

	var x SysAttrDell

	json.Unmarshal(resp, &x)

	_data := x.Attributes

	_sysData := SysAttrData{
		ServerPwrPSRedPolicy: _data.ServerPwr_1_PSRedPolicy,
		ServerPwrPSRapidOn:   _data.ServerPwr_1_PSRapidOn,
	}

	return _sysData, nil

}

//GetBootOrderDell ... will fetch the BootOrder Details
func (c *IloClient) GetBootOrderDell() ([]BootOrderData, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/BootSources"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var x BootOrderDell

	json.Unmarshal(resp, &x)

	var _bootOrder []BootOrderData

	for i := range x.Attributes.BootSeq {

		_result := BootOrderData{
			Enabled: x.Attributes.BootSeq[i].Enabled,
			Index:   x.Attributes.BootSeq[i].Index,
			Name:    x.Attributes.BootSeq[i].Name,
		}

		_bootOrder = append(_bootOrder, _result)
	}

	return _bootOrder, nil

}

//GetSystemEventLogsDell ... Fetch the System Event Logs from the Idrac
func (c *IloClient) GetSystemEventLogsDell(version string) ([]SystemEventLogRes, error) {

	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Logs/Sel"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// v1, err := ver.NewVersion("3.15.17.15")
	v1, _ := ver.NewConstraint("<= 3.15.17.15")
	v2, _ := ver.NewConstraint("<= 3.21.26.22")
	v3, _ := ver.NewConstraint("> 3.21.26.22")
	v4, _ := ver.NewVersion(version)

	if v1.Check(v4) {

		var x SystemEventLogsV1Dell

		json.Unmarshal(resp, &x)

		var _systemEventLogs []SystemEventLogRes

		for i := range x.Members {

			_result := SystemEventLogRes{
				EntryCode:  x.Members[i].EntryCode[0].Member,
				Message:    x.Members[i].Message,
				Name:       x.Members[i].Name,
				SensorType: x.Members[i].SensorType[0].Member,
				Severity:   x.Members[i].Severity,
			}

			_systemEventLogs = append(_systemEventLogs, _result)
		}

		return _systemEventLogs, nil

	} else if v2.Check(v4) || v3.Check(v4) {

		var x SystemEventLogsV2Dell

		json.Unmarshal(resp, &x)

		var _systemEventLogs []SystemEventLogRes

		for i := range x.Members {

			_result := SystemEventLogRes{
				EntryCode:  x.Members[i].EntryCode,
				Message:    x.Members[i].Message,
				Name:       x.Members[i].Name,
				SensorType: x.Members[i].SensorType,
				Severity:   x.Members[i].Severity,
			}

			_systemEventLogs = append(_systemEventLogs, _result)
		}

		return _systemEventLogs, nil
	}
	return nil, err
}

//GetUserAccountsDell ... Fetch the current users created
func (c *IloClient) GetUserAccountsDell() ([]Accounts, error) {

	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Accounts"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var x MemberCountDell
	var users []Accounts

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}

		var y AccountsInfoDell

		json.Unmarshal(resp, &y)

		user := Accounts{
			Name:     y.Name,
			Enabled:  y.Enabled,
			Locked:   y.Locked,
			RoleId:   y.RoleID,
			Username: y.UserName,
		}
		users = append(users, user)

	}

	return users, nil

}

//GetSystemInfoDell ... Will fetch the system info
func (c *IloClient) GetSystemInfoDell() (SystemData, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1"

	resp, err := queryData(c, "GET", url, nil)
	if err != nil {
		return SystemData{}, err
	}

	var x SystemViewDell

	json.Unmarshal(resp, &x)

	_result := SystemData{Health: x.Status.Health,
		Memory:          x.MemorySummary.TotalSystemMemoryGiB,
		Model:           x.Model,
		PowerState:      x.PowerState,
		Processors:      x.ProcessorSummary.Count,
		ProcessorFamily: x.ProcessorSummary.Model,
		SerialNumber:    x.SerialNumber,
	}

	return _result, nil

}
