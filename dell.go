package redfishapi

import (
	"bytes"
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

type SysAttrData struct {
	ServerPwrPSRedPolicy string `json:"redundancy_policy"`
	ServerPwrPSRapidOn   string `json:"hot_spare"`
}

type SysAttrDell struct {
	_Redfish_Settings struct {
		_odata_context string `json:"@odata.context"`
		_odata_type    string `json:"@odata.type"`
		SettingsObject struct {
			_odata_id string `json:"@odata.id"`
		} `json:"SettingsObject"`
	} `json:"@Redfish.Settings"`
	_odata_context    string `json:"@odata.context"`
	_odata_id         string `json:"@odata.id"`
	_odata_type       string `json:"@odata.type"`
	AttributeRegistry string `json:"AttributeRegistry"`
	Attributes        struct {
		Backplane_1_BackplaneBusMode                         string `json:"Backplane.1.BackplaneBusMode"`
		Backplane_1_BackplaneSplitMode                       int    `json:"Backplane.1.BackplaneSplitMode"`
		Diagnostics_1_OSAppCollectionTime                    string `json:"Diagnostics.1.OSAppCollectionTime"`
		LCD_1_ChassisIdentifyDuration                        int    `json:"LCD.1.ChassisIdentifyDuration"`
		LCD_1_Configuration                                  string `json:"LCD.1.Configuration"`
		LCD_1_CurrentDisplay                                 string `json:"LCD.1.CurrentDisplay"`
		LCD_1_ErrorDisplayMode                               string `json:"LCD.1.ErrorDisplayMode"`
		LCD_1_FrontPanelLocking                              string `json:"LCD.1.FrontPanelLocking"`
		LCD_1_HideErrs                                       string `json:"LCD.1.HideErrs"`
		LCD_1_LicenseMsgEnable                               string `json:"LCD.1.LicenseMsgEnable"`
		LCD_1_NMIResetOverride                               string `json:"LCD.1.NMIResetOverride"`
		LCD_1_NumberErrsHidden                               int    `json:"LCD.1.NumberErrsHidden"`
		LCD_1_NumberErrsVisible                              int    `json:"LCD.1.NumberErrsVisible"`
		LCD_1_QualifierTemp                                  string `json:"LCD.1.QualifierTemp"`
		LCD_1_QualifierWatt                                  string `json:"LCD.1.QualifierWatt"`
		LCD_1_UserDefinedString                              string `json:"LCD.1.UserDefinedString"`
		LCD_1_vConsoleIndication                             string `json:"LCD.1.vConsoleIndication"`
		PCIeSlotLFM_1_3rdPartyCard                           string `json:"PCIeSlotLFM.1.3rdPartyCard"`
		PCIeSlotLFM_1_CardType                               string `json:"PCIeSlotLFM.1.CardType"`
		PCIeSlotLFM_1_CustomLFM                              int    `json:"PCIeSlotLFM.1.CustomLFM"`
		PCIeSlotLFM_1_LFMMode                                string `json:"PCIeSlotLFM.1.LFMMode"`
		PCIeSlotLFM_1_MaxLFM                                 int    `json:"PCIeSlotLFM.1.MaxLFM"`
		PCIeSlotLFM_1_SlotState                              string `json:"PCIeSlotLFM.1.SlotState"`
		PCIeSlotLFM_1_TargetLFM                              string `json:"PCIeSlotLFM.1.TargetLFM"`
		PCIeSlotLFM_10_3rdPartyCard                          string `json:"PCIeSlotLFM.10.3rdPartyCard"`
		PCIeSlotLFM_10_CardType                              string `json:"PCIeSlotLFM.10.CardType"`
		PCIeSlotLFM_10_CustomLFM                             int    `json:"PCIeSlotLFM.10.CustomLFM"`
		PCIeSlotLFM_10_LFMMode                               string `json:"PCIeSlotLFM.10.LFMMode"`
		PCIeSlotLFM_10_MaxLFM                                int    `json:"PCIeSlotLFM.10.MaxLFM"`
		PCIeSlotLFM_10_SlotState                             string `json:"PCIeSlotLFM.10.SlotState"`
		PCIeSlotLFM_10_TargetLFM                             string `json:"PCIeSlotLFM.10.TargetLFM"`
		PCIeSlotLFM_11_3rdPartyCard                          string `json:"PCIeSlotLFM.11.3rdPartyCard"`
		PCIeSlotLFM_11_CardType                              string `json:"PCIeSlotLFM.11.CardType"`
		PCIeSlotLFM_11_CustomLFM                             int    `json:"PCIeSlotLFM.11.CustomLFM"`
		PCIeSlotLFM_11_LFMMode                               string `json:"PCIeSlotLFM.11.LFMMode"`
		PCIeSlotLFM_11_MaxLFM                                int    `json:"PCIeSlotLFM.11.MaxLFM"`
		PCIeSlotLFM_11_SlotState                             string `json:"PCIeSlotLFM.11.SlotState"`
		PCIeSlotLFM_11_TargetLFM                             string `json:"PCIeSlotLFM.11.TargetLFM"`
		PCIeSlotLFM_12_3rdPartyCard                          string `json:"PCIeSlotLFM.12.3rdPartyCard"`
		PCIeSlotLFM_12_CardType                              string `json:"PCIeSlotLFM.12.CardType"`
		PCIeSlotLFM_12_CustomLFM                             int    `json:"PCIeSlotLFM.12.CustomLFM"`
		PCIeSlotLFM_12_LFMMode                               string `json:"PCIeSlotLFM.12.LFMMode"`
		PCIeSlotLFM_12_MaxLFM                                int    `json:"PCIeSlotLFM.12.MaxLFM"`
		PCIeSlotLFM_12_SlotState                             string `json:"PCIeSlotLFM.12.SlotState"`
		PCIeSlotLFM_12_TargetLFM                             string `json:"PCIeSlotLFM.12.TargetLFM"`
		PCIeSlotLFM_13_3rdPartyCard                          string `json:"PCIeSlotLFM.13.3rdPartyCard"`
		PCIeSlotLFM_13_CardType                              string `json:"PCIeSlotLFM.13.CardType"`
		PCIeSlotLFM_13_CustomLFM                             int    `json:"PCIeSlotLFM.13.CustomLFM"`
		PCIeSlotLFM_13_LFMMode                               string `json:"PCIeSlotLFM.13.LFMMode"`
		PCIeSlotLFM_13_MaxLFM                                int    `json:"PCIeSlotLFM.13.MaxLFM"`
		PCIeSlotLFM_13_SlotState                             string `json:"PCIeSlotLFM.13.SlotState"`
		PCIeSlotLFM_13_TargetLFM                             string `json:"PCIeSlotLFM.13.TargetLFM"`
		PCIeSlotLFM_14_3rdPartyCard                          string `json:"PCIeSlotLFM.14.3rdPartyCard"`
		PCIeSlotLFM_14_CardType                              string `json:"PCIeSlotLFM.14.CardType"`
		PCIeSlotLFM_14_CustomLFM                             int    `json:"PCIeSlotLFM.14.CustomLFM"`
		PCIeSlotLFM_14_LFMMode                               string `json:"PCIeSlotLFM.14.LFMMode"`
		PCIeSlotLFM_14_MaxLFM                                int    `json:"PCIeSlotLFM.14.MaxLFM"`
		PCIeSlotLFM_14_SlotState                             string `json:"PCIeSlotLFM.14.SlotState"`
		PCIeSlotLFM_14_TargetLFM                             string `json:"PCIeSlotLFM.14.TargetLFM"`
		PCIeSlotLFM_15_3rdPartyCard                          string `json:"PCIeSlotLFM.15.3rdPartyCard"`
		PCIeSlotLFM_15_CardType                              string `json:"PCIeSlotLFM.15.CardType"`
		PCIeSlotLFM_15_CustomLFM                             int    `json:"PCIeSlotLFM.15.CustomLFM"`
		PCIeSlotLFM_15_LFMMode                               string `json:"PCIeSlotLFM.15.LFMMode"`
		PCIeSlotLFM_15_MaxLFM                                int    `json:"PCIeSlotLFM.15.MaxLFM"`
		PCIeSlotLFM_15_SlotState                             string `json:"PCIeSlotLFM.15.SlotState"`
		PCIeSlotLFM_15_TargetLFM                             string `json:"PCIeSlotLFM.15.TargetLFM"`
		PCIeSlotLFM_2_3rdPartyCard                           string `json:"PCIeSlotLFM.2.3rdPartyCard"`
		PCIeSlotLFM_2_CardType                               string `json:"PCIeSlotLFM.2.CardType"`
		PCIeSlotLFM_2_CustomLFM                              int    `json:"PCIeSlotLFM.2.CustomLFM"`
		PCIeSlotLFM_2_LFMMode                                string `json:"PCIeSlotLFM.2.LFMMode"`
		PCIeSlotLFM_2_MaxLFM                                 int    `json:"PCIeSlotLFM.2.MaxLFM"`
		PCIeSlotLFM_2_SlotState                              string `json:"PCIeSlotLFM.2.SlotState"`
		PCIeSlotLFM_2_TargetLFM                              string `json:"PCIeSlotLFM.2.TargetLFM"`
		PCIeSlotLFM_3_3rdPartyCard                           string `json:"PCIeSlotLFM.3.3rdPartyCard"`
		PCIeSlotLFM_3_CardType                               string `json:"PCIeSlotLFM.3.CardType"`
		PCIeSlotLFM_3_CustomLFM                              int    `json:"PCIeSlotLFM.3.CustomLFM"`
		PCIeSlotLFM_3_LFMMode                                string `json:"PCIeSlotLFM.3.LFMMode"`
		PCIeSlotLFM_3_MaxLFM                                 int    `json:"PCIeSlotLFM.3.MaxLFM"`
		PCIeSlotLFM_3_SlotState                              string `json:"PCIeSlotLFM.3.SlotState"`
		PCIeSlotLFM_3_TargetLFM                              string `json:"PCIeSlotLFM.3.TargetLFM"`
		PCIeSlotLFM_4_3rdPartyCard                           string `json:"PCIeSlotLFM.4.3rdPartyCard"`
		PCIeSlotLFM_4_CardType                               string `json:"PCIeSlotLFM.4.CardType"`
		PCIeSlotLFM_4_CustomLFM                              int    `json:"PCIeSlotLFM.4.CustomLFM"`
		PCIeSlotLFM_4_LFMMode                                string `json:"PCIeSlotLFM.4.LFMMode"`
		PCIeSlotLFM_4_MaxLFM                                 int    `json:"PCIeSlotLFM.4.MaxLFM"`
		PCIeSlotLFM_4_SlotState                              string `json:"PCIeSlotLFM.4.SlotState"`
		PCIeSlotLFM_4_TargetLFM                              string `json:"PCIeSlotLFM.4.TargetLFM"`
		PCIeSlotLFM_5_3rdPartyCard                           string `json:"PCIeSlotLFM.5.3rdPartyCard"`
		PCIeSlotLFM_5_CardType                               string `json:"PCIeSlotLFM.5.CardType"`
		PCIeSlotLFM_5_CustomLFM                              int    `json:"PCIeSlotLFM.5.CustomLFM"`
		PCIeSlotLFM_5_LFMMode                                string `json:"PCIeSlotLFM.5.LFMMode"`
		PCIeSlotLFM_5_MaxLFM                                 int    `json:"PCIeSlotLFM.5.MaxLFM"`
		PCIeSlotLFM_5_SlotState                              string `json:"PCIeSlotLFM.5.SlotState"`
		PCIeSlotLFM_5_TargetLFM                              string `json:"PCIeSlotLFM.5.TargetLFM"`
		PCIeSlotLFM_6_3rdPartyCard                           string `json:"PCIeSlotLFM.6.3rdPartyCard"`
		PCIeSlotLFM_6_CardType                               string `json:"PCIeSlotLFM.6.CardType"`
		PCIeSlotLFM_6_CustomLFM                              int    `json:"PCIeSlotLFM.6.CustomLFM"`
		PCIeSlotLFM_6_LFMMode                                string `json:"PCIeSlotLFM.6.LFMMode"`
		PCIeSlotLFM_6_MaxLFM                                 int    `json:"PCIeSlotLFM.6.MaxLFM"`
		PCIeSlotLFM_6_SlotState                              string `json:"PCIeSlotLFM.6.SlotState"`
		PCIeSlotLFM_6_TargetLFM                              string `json:"PCIeSlotLFM.6.TargetLFM"`
		PCIeSlotLFM_7_3rdPartyCard                           string `json:"PCIeSlotLFM.7.3rdPartyCard"`
		PCIeSlotLFM_7_CardType                               string `json:"PCIeSlotLFM.7.CardType"`
		PCIeSlotLFM_7_CustomLFM                              int    `json:"PCIeSlotLFM.7.CustomLFM"`
		PCIeSlotLFM_7_LFMMode                                string `json:"PCIeSlotLFM.7.LFMMode"`
		PCIeSlotLFM_7_MaxLFM                                 int    `json:"PCIeSlotLFM.7.MaxLFM"`
		PCIeSlotLFM_7_SlotState                              string `json:"PCIeSlotLFM.7.SlotState"`
		PCIeSlotLFM_7_TargetLFM                              string `json:"PCIeSlotLFM.7.TargetLFM"`
		PCIeSlotLFM_8_3rdPartyCard                           string `json:"PCIeSlotLFM.8.3rdPartyCard"`
		PCIeSlotLFM_8_CardType                               string `json:"PCIeSlotLFM.8.CardType"`
		PCIeSlotLFM_8_CustomLFM                              int    `json:"PCIeSlotLFM.8.CustomLFM"`
		PCIeSlotLFM_8_LFMMode                                string `json:"PCIeSlotLFM.8.LFMMode"`
		PCIeSlotLFM_8_MaxLFM                                 int    `json:"PCIeSlotLFM.8.MaxLFM"`
		PCIeSlotLFM_8_SlotState                              string `json:"PCIeSlotLFM.8.SlotState"`
		PCIeSlotLFM_8_TargetLFM                              string `json:"PCIeSlotLFM.8.TargetLFM"`
		PCIeSlotLFM_9_3rdPartyCard                           string `json:"PCIeSlotLFM.9.3rdPartyCard"`
		PCIeSlotLFM_9_CardType                               string `json:"PCIeSlotLFM.9.CardType"`
		PCIeSlotLFM_9_CustomLFM                              int    `json:"PCIeSlotLFM.9.CustomLFM"`
		PCIeSlotLFM_9_LFMMode                                string `json:"PCIeSlotLFM.9.LFMMode"`
		PCIeSlotLFM_9_MaxLFM                                 int    `json:"PCIeSlotLFM.9.MaxLFM"`
		PCIeSlotLFM_9_SlotState                              string `json:"PCIeSlotLFM.9.SlotState"`
		PCIeSlotLFM_9_TargetLFM                              string `json:"PCIeSlotLFM.9.TargetLFM"`
		PowerHistorical_1_IntervalInSeconds                  int    `json:"PowerHistorical.1.IntervalInSeconds"`
		QuickSync_1_Access                                   string `json:"QuickSync.1.Access"`
		QuickSync_1_InactivityTimeout                        int    `json:"QuickSync.1.InactivityTimeout"`
		QuickSync_1_InactivityTimerEnable                    string `json:"QuickSync.1.InactivityTimerEnable"`
		QuickSync_1_Presence                                 string `json:"QuickSync.1.Presence"`
		QuickSync_1_ReadAuthentication                       string `json:"QuickSync.1.ReadAuthentication"`
		QuickSync_1_WifiEnable                               string `json:"QuickSync.1.WifiEnable"`
		SC_BMC_1_PowerMonitoring                             int    `json:"SC-BMC.1.PowerMonitoring"`
		ServerInfo_1_NodeID                                  string `json:"ServerInfo.1.NodeID"`
		ServerInfo_1_RChassisServiceTag                      string `json:"ServerInfo.1.RChassisServiceTag"`
		ServerInfo_1_ServiceTag                              string `json:"ServerInfo.1.ServiceTag"`
		ServerOS_1_HostName                                  string `json:"ServerOS.1.HostName"`
		ServerOS_1_OEMOSVersion                              string `json:"ServerOS.1.OEMOSVersion"`
		ServerOS_1_OSName                                    string `json:"ServerOS.1.OSName"`
		ServerOS_1_OSVersion                                 string `json:"ServerOS.1.OSVersion"`
		ServerOS_1_ServerPoweredOnTime                       int    `json:"ServerOS.1.ServerPoweredOnTime"`
		ServerPwr_1_ActivePolicyName                         string `json:"ServerPwr.1.ActivePolicyName"`
		ServerPwr_1_ActivePowerCapVal                        int    `json:"ServerPwr.1.ActivePowerCapVal"`
		ServerPwr_1_PSPFCEnabled                             string `json:"ServerPwr.1.PSPFCEnabled"`
		ServerPwr_1_PSRapidOn                                string `json:"ServerPwr.1.PSRapidOn"`
		ServerPwr_1_PSRedPolicy                              string `json:"ServerPwr.1.PSRedPolicy"`
		ServerPwr_1_PowerCapMaxThres                         int    `json:"ServerPwr.1.PowerCapMaxThres"`
		ServerPwr_1_PowerCapMinThres                         int    `json:"ServerPwr.1.PowerCapMinThres"`
		ServerPwr_1_PowerCapSetting                          string `json:"ServerPwr.1.PowerCapSetting"`
		ServerPwr_1_PowerCapValue                            int    `json:"ServerPwr.1.PowerCapValue"`
		ServerPwr_1_RapidOnPrimaryPSU                        string `json:"ServerPwr.1.RapidOnPrimaryPSU"`
		ServerPwrMon_1_AccumulativePower                     int    `json:"ServerPwrMon.1.AccumulativePower"`
		ServerPwrMon_1_CumulativePowerStartTime              int    `json:"ServerPwrMon.1.CumulativePowerStartTime"`
		ServerPwrMon_1_CumulativePowerStartTimeStr           string `json:"ServerPwrMon.1.CumulativePowerStartTimeStr"`
		ServerPwrMon_1_MinPowerTime                          int    `json:"ServerPwrMon.1.MinPowerTime"`
		ServerPwrMon_1_MinPowerTimeStr                       string `json:"ServerPwrMon.1.MinPowerTimeStr"`
		ServerPwrMon_1_MinPowerWatts                         int    `json:"ServerPwrMon.1.MinPowerWatts"`
		ServerPwrMon_1_PeakCurrentTime                       int    `json:"ServerPwrMon.1.PeakCurrentTime"`
		ServerPwrMon_1_PeakCurrentTimeStr                    string `json:"ServerPwrMon.1.PeakCurrentTimeStr"`
		ServerPwrMon_1_PeakPowerStartTime                    int    `json:"ServerPwrMon.1.PeakPowerStartTime"`
		ServerPwrMon_1_PeakPowerStartTimeStr                 string `json:"ServerPwrMon.1.PeakPowerStartTimeStr"`
		ServerPwrMon_1_PeakPowerTime                         int    `json:"ServerPwrMon.1.PeakPowerTime"`
		ServerPwrMon_1_PeakPowerTimeStr                      string `json:"ServerPwrMon.1.PeakPowerTimeStr"`
		ServerPwrMon_1_PeakPowerWatts                        int    `json:"ServerPwrMon.1.PeakPowerWatts"`
		ServerPwrMon_1_PowerConfigReset                      string `json:"ServerPwrMon.1.PowerConfigReset"`
		ServerTopology_1_AisleName                           string `json:"ServerTopology.1.AisleName"`
		ServerTopology_1_DataCenterName                      string `json:"ServerTopology.1.DataCenterName"`
		ServerTopology_1_RackName                            string `json:"ServerTopology.1.RackName"`
		ServerTopology_1_RackSlot                            int    `json:"ServerTopology.1.RackSlot"`
		ServerTopology_1_RoomName                            string `json:"ServerTopology.1.RoomName"`
		ServerTopology_1_SizeOfManagedSystemInU              int    `json:"ServerTopology.1.SizeOfManagedSystemInU"`
		Storage_1_AvailableSpareAlertThreshold               int    `json:"Storage.1.AvailableSpareAlertThreshold"`
		Storage_1_RemainingRatedWriteEnduranceAlertThreshold int    `json:"Storage.1.RemainingRatedWriteEnduranceAlertThreshold"`
		ThermalConfig_1_CriticalEventGenerationInterval      int    `json:"ThermalConfig.1.CriticalEventGenerationInterval"`
		ThermalConfig_1_EventGenerationInterval              int    `json:"ThermalConfig.1.EventGenerationInterval"`
		ThermalConfig_1_FreshAirCompliantConfiguration       string `json:"ThermalConfig.1.FreshAirCompliantConfiguration"`
		ThermalConfig_1_ValidFanConfiguration                string `json:"ThermalConfig.1.ValidFanConfiguration"`
		ThermalHistorical_1_IntervalInSeconds                int    `json:"ThermalHistorical.1.IntervalInSeconds"`
		ThermalSettings_1_AirExhaustTemp                     string `json:"ThermalSettings.1.AirExhaustTemp"`
		ThermalSettings_1_AirExhaustTempSupport              string `json:"ThermalSettings.1.AirExhaustTempSupport"`
		ThermalSettings_1_FanSpeedHighOffsetVal              int    `json:"ThermalSettings.1.FanSpeedHighOffsetVal"`
		ThermalSettings_1_FanSpeedLowOffsetVal               int    `json:"ThermalSettings.1.FanSpeedLowOffsetVal"`
		ThermalSettings_1_FanSpeedMaxOffsetVal               int    `json:"ThermalSettings.1.FanSpeedMaxOffsetVal"`
		ThermalSettings_1_FanSpeedMediumOffsetVal            int    `json:"ThermalSettings.1.FanSpeedMediumOffsetVal"`
		ThermalSettings_1_FanSpeedOffset                     string `json:"ThermalSettings.1.FanSpeedOffset"`
		ThermalSettings_1_MFSMaximumLimit                    int    `json:"ThermalSettings.1.MFSMaximumLimit"`
		ThermalSettings_1_MFSMinimumLimit                    int    `json:"ThermalSettings.1.MFSMinimumLimit"`
		ThermalSettings_1_MinimumFanSpeed                    int    `json:"ThermalSettings.1.MinimumFanSpeed"`
		ThermalSettings_1_PCIeSlotLFMSupport                 string `json:"ThermalSettings.1.PCIeSlotLFMSupport"`
		ThermalSettings_1_SystemCFMSupport                   string `json:"ThermalSettings.1.SystemCFMSupport"`
		ThermalSettings_1_ThermalProfile                     string `json:"ThermalSettings.1.ThermalProfile"`
		USBFront_1_Enable                                    string `json:"USBFront.1.Enable"`
	} `json:"Attributes"`
	Description string `json:"Description"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
}

type IDRACData struct {
	VirtualConsoleMaxSessions int    `json:"virtualconsole_maxsessions"`
	VirtualConsolePluginType  string `json:"virtualconsole_plugintype"`
	WebServerSSLEncryption    string `json:"webserver_ssl_encryption"`
	IPMILanEnable             string `json:"ipmi_lan_enable"`
	DNSDomainName             string `json:"dns_domainname"`
}

type IDRACAttrDell struct {
	_Redfish_Settings struct {
		_odata_context string `json:"@odata.context"`
		_odata_type    string `json:"@odata.type"`
		SettingsObject struct {
			_odata_id string `json:"@odata.id"`
		} `json:"SettingsObject"`
	} `json:"@Redfish.Settings"`
	_odata_context    string `json:"@odata.context"`
	_odata_id         string `json:"@odata.id"`
	_odata_type       string `json:"@odata.type"`
	AttributeRegistry string `json:"AttributeRegistry"`
	Attributes        struct {
		ADGroup_1_Domain                                                     string      `json:"ADGroup.1.Domain"`
		ADGroup_1_Name                                                       string      `json:"ADGroup.1.Name"`
		ADGroup_1_Privilege                                                  int         `json:"ADGroup.1.Privilege"`
		ADGroup_2_Domain                                                     string      `json:"ADGroup.2.Domain"`
		ADGroup_2_Name                                                       string      `json:"ADGroup.2.Name"`
		ADGroup_2_Privilege                                                  int         `json:"ADGroup.2.Privilege"`
		ADGroup_3_Domain                                                     string      `json:"ADGroup.3.Domain"`
		ADGroup_3_Name                                                       string      `json:"ADGroup.3.Name"`
		ADGroup_3_Privilege                                                  int         `json:"ADGroup.3.Privilege"`
		ADGroup_4_Domain                                                     string      `json:"ADGroup.4.Domain"`
		ADGroup_4_Name                                                       string      `json:"ADGroup.4.Name"`
		ADGroup_4_Privilege                                                  int         `json:"ADGroup.4.Privilege"`
		ADGroup_5_Domain                                                     string      `json:"ADGroup.5.Domain"`
		ADGroup_5_Name                                                       string      `json:"ADGroup.5.Name"`
		ADGroup_5_Privilege                                                  int         `json:"ADGroup.5.Privilege"`
		ASRConfig_1_Enable                                                   string      `json:"ASRConfig.1.Enable"`
		ActiveDirectory_1_AuthTimeout                                        int         `json:"ActiveDirectory.1.AuthTimeout"`
		ActiveDirectory_1_CertValidationEnable                               string      `json:"ActiveDirectory.1.CertValidationEnable"`
		ActiveDirectory_1_DCLookupByUserDomain                               string      `json:"ActiveDirectory.1.DCLookupByUserDomain"`
		ActiveDirectory_1_DCLookupDomainName                                 string      `json:"ActiveDirectory.1.DCLookupDomainName"`
		ActiveDirectory_1_DCLookupEnable                                     string      `json:"ActiveDirectory.1.DCLookupEnable"`
		ActiveDirectory_1_DomainController1                                  string      `json:"ActiveDirectory.1.DomainController1"`
		ActiveDirectory_1_DomainController2                                  string      `json:"ActiveDirectory.1.DomainController2"`
		ActiveDirectory_1_DomainController3                                  string      `json:"ActiveDirectory.1.DomainController3"`
		ActiveDirectory_1_Enable                                             string      `json:"ActiveDirectory.1.Enable"`
		ActiveDirectory_1_GCLookupEnable                                     string      `json:"ActiveDirectory.1.GCLookupEnable"`
		ActiveDirectory_1_GCRootDomain                                       string      `json:"ActiveDirectory.1.GCRootDomain"`
		ActiveDirectory_1_GlobalCatalog1                                     string      `json:"ActiveDirectory.1.GlobalCatalog1"`
		ActiveDirectory_1_GlobalCatalog2                                     string      `json:"ActiveDirectory.1.GlobalCatalog2"`
		ActiveDirectory_1_GlobalCatalog3                                     string      `json:"ActiveDirectory.1.GlobalCatalog3"`
		ActiveDirectory_1_RacDomain                                          string      `json:"ActiveDirectory.1.RacDomain"`
		ActiveDirectory_1_RacName                                            string      `json:"ActiveDirectory.1.RacName"`
		ActiveDirectory_1_SSOEnable                                          string      `json:"ActiveDirectory.1.SSOEnable"`
		ActiveDirectory_1_Schema                                             string      `json:"ActiveDirectory.1.Schema"`
		AutoOSLockGroup_1_AutoOSLockState                                    string      `json:"AutoOSLockGroup.1.AutoOSLockState"`
		CurrentIPv4_1_Address                                                string      `json:"CurrentIPv4.1.Address"`
		CurrentIPv4_1_DHCPEnable                                             string      `json:"CurrentIPv4.1.DHCPEnable"`
		CurrentIPv4_1_DNS1                                                   string      `json:"CurrentIPv4.1.DNS1"`
		CurrentIPv4_1_DNS2                                                   string      `json:"CurrentIPv4.1.DNS2"`
		CurrentIPv4_1_DNSFromDHCP                                            string      `json:"CurrentIPv4.1.DNSFromDHCP"`
		CurrentIPv4_1_DupAddrDetected                                        string      `json:"CurrentIPv4.1.DupAddrDetected"`
		CurrentIPv4_1_Enable                                                 string      `json:"CurrentIPv4.1.Enable"`
		CurrentIPv4_1_Gateway                                                string      `json:"CurrentIPv4.1.Gateway"`
		CurrentIPv4_1_Netmask                                                string      `json:"CurrentIPv4.1.Netmask"`
		CurrentIPv6_1_Address1                                               string      `json:"CurrentIPv6.1.Address1"`
		CurrentIPv6_1_Address10                                              string      `json:"CurrentIPv6.1.Address10"`
		CurrentIPv6_1_Address11                                              string      `json:"CurrentIPv6.1.Address11"`
		CurrentIPv6_1_Address12                                              string      `json:"CurrentIPv6.1.Address12"`
		CurrentIPv6_1_Address13                                              string      `json:"CurrentIPv6.1.Address13"`
		CurrentIPv6_1_Address14                                              string      `json:"CurrentIPv6.1.Address14"`
		CurrentIPv6_1_Address15                                              string      `json:"CurrentIPv6.1.Address15"`
		CurrentIPv6_1_Address2                                               string      `json:"CurrentIPv6.1.Address2"`
		CurrentIPv6_1_Address3                                               string      `json:"CurrentIPv6.1.Address3"`
		CurrentIPv6_1_Address4                                               string      `json:"CurrentIPv6.1.Address4"`
		CurrentIPv6_1_Address5                                               string      `json:"CurrentIPv6.1.Address5"`
		CurrentIPv6_1_Address6                                               string      `json:"CurrentIPv6.1.Address6"`
		CurrentIPv6_1_Address7                                               string      `json:"CurrentIPv6.1.Address7"`
		CurrentIPv6_1_Address8                                               string      `json:"CurrentIPv6.1.Address8"`
		CurrentIPv6_1_Address9                                               string      `json:"CurrentIPv6.1.Address9"`
		CurrentIPv6_1_AutoConfig                                             string      `json:"CurrentIPv6.1.AutoConfig"`
		CurrentIPv6_1_DNS1                                                   string      `json:"CurrentIPv6.1.DNS1"`
		CurrentIPv6_1_DNS2                                                   string      `json:"CurrentIPv6.1.DNS2"`
		CurrentIPv6_1_DNSFromDHCP6                                           string      `json:"CurrentIPv6.1.DNSFromDHCP6"`
		CurrentIPv6_1_Enable                                                 string      `json:"CurrentIPv6.1.Enable"`
		CurrentIPv6_1_Gateway                                                string      `json:"CurrentIPv6.1.Gateway"`
		CurrentIPv6_1_IPV6NumOfExtAddress                                    int         `json:"CurrentIPv6.1.IPV6NumOfExtAddress"`
		CurrentIPv6_1_LinkLocalAddress                                       string      `json:"CurrentIPv6.1.LinkLocalAddress"`
		CurrentIPv6_1_PrefixLength                                           int         `json:"CurrentIPv6.1.PrefixLength"`
		CurrentNIC_1_ActiveNIC                                               string      `json:"CurrentNIC.1.ActiveNIC"`
		CurrentNIC_1_ActiveSharedLOM                                         string      `json:"CurrentNIC.1.ActiveSharedLOM"`
		CurrentNIC_1_AutoDetect                                              string      `json:"CurrentNIC.1.AutoDetect"`
		CurrentNIC_1_Autoneg                                                 string      `json:"CurrentNIC.1.Autoneg"`
		CurrentNIC_1_DNSDomainFromDHCP                                       string      `json:"CurrentNIC.1.DNSDomainFromDHCP"`
		CurrentNIC_1_DNSDomainName                                           string      `json:"CurrentNIC.1.DNSDomainName"`
		CurrentNIC_1_DNSRacName                                              string      `json:"CurrentNIC.1.DNSRacName"`
		CurrentNIC_1_DNSRegister                                             string      `json:"CurrentNIC.1.DNSRegister"`
		CurrentNIC_1_DedicatedNICScanTime                                    int         `json:"CurrentNIC.1.DedicatedNICScanTime"`
		CurrentNIC_1_Duplex                                                  string      `json:"CurrentNIC.1.Duplex"`
		CurrentNIC_1_Enable                                                  string      `json:"CurrentNIC.1.Enable"`
		CurrentNIC_1_Failover                                                string      `json:"CurrentNIC.1.Failover"`
		CurrentNIC_1_LinkStatus                                              string      `json:"CurrentNIC.1.LinkStatus"`
		CurrentNIC_1_MACAddress                                              string      `json:"CurrentNIC.1.MACAddress"`
		CurrentNIC_1_MACAddress2                                             string      `json:"CurrentNIC.1.MACAddress2"`
		CurrentNIC_1_MTU                                                     int         `json:"CurrentNIC.1.MTU"`
		CurrentNIC_1_MgmtIfaceName                                           string      `json:"CurrentNIC.1.MgmtIfaceName"`
		CurrentNIC_1_NumberOfLOM                                             int         `json:"CurrentNIC.1.NumberOfLOM"`
		CurrentNIC_1_Selection                                               string      `json:"CurrentNIC.1.Selection"`
		CurrentNIC_1_SharedNICScanTime                                       int         `json:"CurrentNIC.1.SharedNICScanTime"`
		CurrentNIC_1_Speed                                                   string      `json:"CurrentNIC.1.Speed"`
		CurrentNIC_1_VLanEnable                                              string      `json:"CurrentNIC.1.VLanEnable"`
		CurrentNIC_1_VLanID                                                  int         `json:"CurrentNIC.1.VLanID"`
		CurrentNIC_1_VLanPriority                                            int         `json:"CurrentNIC.1.VLanPriority"`
		CurrentNIC_1_VLanSetting                                             string      `json:"CurrentNIC.1.VLanSetting"`
		DefaultCredentialMitigationConfigGroup_1_DefaultCredentialMitigation string      `json:"DefaultCredentialMitigationConfigGroup.1.DefaultCredentialMitigation"`
		EmailAlert_1_Address                                                 string      `json:"EmailAlert.1.Address"`
		EmailAlert_1_CustomMsg                                               string      `json:"EmailAlert.1.CustomMsg"`
		EmailAlert_1_Enable                                                  string      `json:"EmailAlert.1.Enable"`
		EmailAlert_2_Address                                                 string      `json:"EmailAlert.2.Address"`
		EmailAlert_2_CustomMsg                                               string      `json:"EmailAlert.2.CustomMsg"`
		EmailAlert_2_Enable                                                  string      `json:"EmailAlert.2.Enable"`
		EmailAlert_3_Address                                                 string      `json:"EmailAlert.3.Address"`
		EmailAlert_3_CustomMsg                                               string      `json:"EmailAlert.3.CustomMsg"`
		EmailAlert_3_Enable                                                  string      `json:"EmailAlert.3.Enable"`
		EmailAlert_4_Address                                                 string      `json:"EmailAlert.4.Address"`
		EmailAlert_4_CustomMsg                                               string      `json:"EmailAlert.4.CustomMsg"`
		EmailAlert_4_Enable                                                  string      `json:"EmailAlert.4.Enable"`
		GUI_1_SecurityPolicyMessage                                          string      `json:"GUI.1.SecurityPolicyMessage"`
		GroupManager_1_GroupName                                             string      `json:"GroupManager.1.GroupName"`
		GroupManager_1_GroupUUID                                             string      `json:"GroupManager.1.GroupUUID"`
		GroupManager_1_Status                                                string      `json:"GroupManager.1.Status"`
		IOIDOpt_1_IOIDOptEnable                                              string      `json:"IOIDOpt.1.IOIDOptEnable"`
		IOIDOpt_1_InitiatorPersistencePolicy                                 string      `json:"IOIDOpt.1.InitiatorPersistencePolicy"`
		IOIDOpt_1_StorageTargetPersistencePolicy                             string      `json:"IOIDOpt.1.StorageTargetPersistencePolicy"`
		IOIDOpt_1_VirtualAddressPersistencePolicyAuxPwrd                     string      `json:"IOIDOpt.1.VirtualAddressPersistencePolicyAuxPwrd"`
		IOIDOpt_1_VirtualAddressPersistencePolicyNonAuxPwrd                  string      `json:"IOIDOpt.1.VirtualAddressPersistencePolicyNonAuxPwrd"`
		IPBlocking_1_BlockEnable                                             string      `json:"IPBlocking.1.BlockEnable"`
		IPBlocking_1_FailCount                                               int         `json:"IPBlocking.1.FailCount"`
		IPBlocking_1_FailWindow                                              int         `json:"IPBlocking.1.FailWindow"`
		IPBlocking_1_PenaltyTime                                             int         `json:"IPBlocking.1.PenaltyTime"`
		IPBlocking_1_RangeAddr                                               string      `json:"IPBlocking.1.RangeAddr"`
		IPBlocking_1_RangeEnable                                             string      `json:"IPBlocking.1.RangeEnable"`
		IPBlocking_1_RangeMask                                               string      `json:"IPBlocking.1.RangeMask"`
		IPMILan_1_AlertEnable                                                string      `json:"IPMILan.1.AlertEnable"`
		IPMILan_1_CommunityName                                              string      `json:"IPMILan.1.CommunityName"`
		IPMILan_1_Enable                                                     string      `json:"IPMILan.1.Enable"`
		IPMILan_1_EncryptionKey                                              string      `json:"IPMILan.1.EncryptionKey"`
		IPMILan_1_PrivLimit                                                  string      `json:"IPMILan.1.PrivLimit"`
		IPMISOL_1_AccumulateInterval                                         int         `json:"IPMISOL.1.AccumulateInterval"`
		IPMISOL_1_BaudRate                                                   string      `json:"IPMISOL.1.BaudRate"`
		IPMISOL_1_Enable                                                     string      `json:"IPMISOL.1.Enable"`
		IPMISOL_1_MinPrivilege                                               string      `json:"IPMISOL.1.MinPrivilege"`
		IPMISOL_1_SendThreshold                                              int         `json:"IPMISOL.1.SendThreshold"`
		IPMISerial_1_BaudRate                                                string      `json:"IPMISerial.1.BaudRate"`
		IPMISerial_1_ChanPrivLimit                                           string      `json:"IPMISerial.1.ChanPrivLimit"`
		IPMISerial_1_ConnectionMode                                          string      `json:"IPMISerial.1.ConnectionMode"`
		IPMISerial_1_DeleteControl                                           string      `json:"IPMISerial.1.DeleteControl"`
		IPMISerial_1_EchoControl                                             string      `json:"IPMISerial.1.EchoControl"`
		IPMISerial_1_FlowControl                                             string      `json:"IPMISerial.1.FlowControl"`
		IPMISerial_1_HandshakeControl                                        string      `json:"IPMISerial.1.HandshakeControl"`
		IPMISerial_1_InputNewLineSeq                                         string      `json:"IPMISerial.1.InputNewLineSeq"`
		IPMISerial_1_LineEdit                                                string      `json:"IPMISerial.1.LineEdit"`
		IPMISerial_1_NewLineSeq                                              string      `json:"IPMISerial.1.NewLineSeq"`
		IPv4_1_Address                                                       string      `json:"IPv4.1.Address"`
		IPv4_1_DHCPEnable                                                    string      `json:"IPv4.1.DHCPEnable"`
		IPv4_1_DNS1                                                          string      `json:"IPv4.1.DNS1"`
		IPv4_1_DNS2                                                          string      `json:"IPv4.1.DNS2"`
		IPv4_1_DNSFromDHCP                                                   string      `json:"IPv4.1.DNSFromDHCP"`
		IPv4_1_Enable                                                        string      `json:"IPv4.1.Enable"`
		IPv4_1_Gateway                                                       string      `json:"IPv4.1.Gateway"`
		IPv4_1_Netmask                                                       string      `json:"IPv4.1.Netmask"`
		IPv4Static_1_Address                                                 string      `json:"IPv4Static.1.Address"`
		IPv4Static_1_DNS1                                                    string      `json:"IPv4Static.1.DNS1"`
		IPv4Static_1_DNS2                                                    string      `json:"IPv4Static.1.DNS2"`
		IPv4Static_1_DNSFromDHCP                                             string      `json:"IPv4Static.1.DNSFromDHCP"`
		IPv4Static_1_Gateway                                                 string      `json:"IPv4Static.1.Gateway"`
		IPv4Static_1_Netmask                                                 string      `json:"IPv4Static.1.Netmask"`
		IPv6_1_Address1                                                      string      `json:"IPv6.1.Address1"`
		IPv6_1_Address10                                                     string      `json:"IPv6.1.Address10"`
		IPv6_1_Address11                                                     string      `json:"IPv6.1.Address11"`
		IPv6_1_Address12                                                     string      `json:"IPv6.1.Address12"`
		IPv6_1_Address13                                                     string      `json:"IPv6.1.Address13"`
		IPv6_1_Address14                                                     string      `json:"IPv6.1.Address14"`
		IPv6_1_Address15                                                     string      `json:"IPv6.1.Address15"`
		IPv6_1_Address2                                                      string      `json:"IPv6.1.Address2"`
		IPv6_1_Address3                                                      string      `json:"IPv6.1.Address3"`
		IPv6_1_Address4                                                      string      `json:"IPv6.1.Address4"`
		IPv6_1_Address5                                                      string      `json:"IPv6.1.Address5"`
		IPv6_1_Address6                                                      string      `json:"IPv6.1.Address6"`
		IPv6_1_Address7                                                      string      `json:"IPv6.1.Address7"`
		IPv6_1_Address8                                                      string      `json:"IPv6.1.Address8"`
		IPv6_1_Address9                                                      string      `json:"IPv6.1.Address9"`
		IPv6_1_AddressState                                                  string      `json:"IPv6.1.AddressState"`
		IPv6_1_AutoConfig                                                    string      `json:"IPv6.1.AutoConfig"`
		IPv6_1_DNS1                                                          string      `json:"IPv6.1.DNS1"`
		IPv6_1_DNS2                                                          string      `json:"IPv6.1.DNS2"`
		IPv6_1_DNSFromDHCP6                                                  string      `json:"IPv6.1.DNSFromDHCP6"`
		IPv6_1_Enable                                                        string      `json:"IPv6.1.Enable"`
		IPv6_1_Gateway                                                       string      `json:"IPv6.1.Gateway"`
		IPv6_1_IPV6NumOfExtAddress                                           int         `json:"IPv6.1.IPV6NumOfExtAddress"`
		IPv6_1_LinkLocalAddress                                              string      `json:"IPv6.1.LinkLocalAddress"`
		IPv6_1_PrefixLength                                                  int         `json:"IPv6.1.PrefixLength"`
		IPv6Static_1_Address1                                                string      `json:"IPv6Static.1.Address1"`
		IPv6Static_1_DNS1                                                    string      `json:"IPv6Static.1.DNS1"`
		IPv6Static_1_DNS2                                                    string      `json:"IPv6Static.1.DNS2"`
		IPv6Static_1_DNSFromDHCP6                                            string      `json:"IPv6Static.1.DNSFromDHCP6"`
		IPv6Static_1_Gateway                                                 string      `json:"IPv6Static.1.Gateway"`
		IPv6Static_1_PrefixLength                                            int         `json:"IPv6Static.1.PrefixLength"`
		IPv6URL_1_URL                                                        string      `json:"IPv6URL.1.URL"`
		Info_1_Build                                                         string      `json:"Info.1.Build"`
		Info_1_CPLDVersion                                                   string      `json:"Info.1.CPLDVersion"`
		Info_1_Description                                                   string      `json:"Info.1.Description"`
		Info_1_HWRev                                                         string      `json:"Info.1.HWRev"`
		Info_1_IPMIVersion                                                   string      `json:"Info.1.IPMIVersion"`
		Info_1_Name                                                          string      `json:"Info.1.Name"`
		Info_1_Product                                                       string      `json:"Info.1.Product"`
		Info_1_RollbackBuild                                                 string      `json:"Info.1.RollbackBuild"`
		Info_1_RollbackVersion                                               string      `json:"Info.1.RollbackVersion"`
		Info_1_ServerGen                                                     string      `json:"Info.1.ServerGen"`
		Info_1_Type                                                          string      `json:"Info.1.Type"`
		Info_1_Version                                                       string      `json:"Info.1.Version"`
		IntegratedDatacenter_1_DiscoveryEnable                               string      `json:"IntegratedDatacenter.1.DiscoveryEnable"`
		LDAP_1_BaseDN                                                        string      `json:"LDAP.1.BaseDN"`
		LDAP_1_BindDN                                                        string      `json:"LDAP.1.BindDN"`
		LDAP_1_BindPassword                                                  interface{} `json:"LDAP.1.BindPassword"`
		LDAP_1_CertValidationEnable                                          string      `json:"LDAP.1.CertValidationEnable"`
		LDAP_1_Enable                                                        string      `json:"LDAP.1.Enable"`
		LDAP_1_GroupAttribute                                                string      `json:"LDAP.1.GroupAttribute"`
		LDAP_1_GroupAttributeIsDN                                            string      `json:"LDAP.1.GroupAttributeIsDN"`
		LDAP_1_Port                                                          int         `json:"LDAP.1.Port"`
		LDAP_1_SearchFilter                                                  string      `json:"LDAP.1.SearchFilter"`
		LDAP_1_Server                                                        string      `json:"LDAP.1.Server"`
		LDAP_1_UserAttribute                                                 string      `json:"LDAP.1.UserAttribute"`
		LDAPRoleGroup_1_DN                                                   string      `json:"LDAPRoleGroup.1.DN"`
		LDAPRoleGroup_1_Privilege                                            int         `json:"LDAPRoleGroup.1.Privilege"`
		LDAPRoleGroup_2_DN                                                   string      `json:"LDAPRoleGroup.2.DN"`
		LDAPRoleGroup_2_Privilege                                            int         `json:"LDAPRoleGroup.2.Privilege"`
		LDAPRoleGroup_3_DN                                                   string      `json:"LDAPRoleGroup.3.DN"`
		LDAPRoleGroup_3_Privilege                                            int         `json:"LDAPRoleGroup.3.Privilege"`
		LDAPRoleGroup_4_DN                                                   string      `json:"LDAPRoleGroup.4.DN"`
		LDAPRoleGroup_4_Privilege                                            int         `json:"LDAPRoleGroup.4.Privilege"`
		LDAPRoleGroup_5_DN                                                   string      `json:"LDAPRoleGroup.5.DN"`
		LDAPRoleGroup_5_Privilege                                            int         `json:"LDAPRoleGroup.5.Privilege"`
		LocalSecurity_1_KCSAccess                                            string      `json:"LocalSecurity.1.KCSAccess"`
		LocalSecurity_1_LocalConfig                                          string      `json:"LocalSecurity.1.LocalConfig"`
		LocalSecurity_1_PrebootConfig                                        string      `json:"LocalSecurity.1.PrebootConfig"`
		Lockdown_1_SystemLockdown                                            string      `json:"Lockdown.1.SystemLockdown"`
		Logging_1_SELOEMEventFilterEnable                                    string      `json:"Logging.1.SELOEMEventFilterEnable"`
		NIC_1_AutoConfig                                                     string      `json:"NIC.1.AutoConfig"`
		NIC_1_AutoConfigIPV6                                                 string      `json:"NIC.1.AutoConfigIPV6"`
		NIC_1_AutoDetect                                                     string      `json:"NIC.1.AutoDetect"`
		NIC_1_Autoneg                                                        string      `json:"NIC.1.Autoneg"`
		NIC_1_DNSDomainFromDHCP                                              string      `json:"NIC.1.DNSDomainFromDHCP"`
		NIC_1_DNSDomainName                                                  string      `json:"NIC.1.DNSDomainName"`
		NIC_1_DNSDomainNameFromDHCP                                          string      `json:"NIC.1.DNSDomainNameFromDHCP"`
		NIC_1_DNSRacName                                                     string      `json:"NIC.1.DNSRacName"`
		NIC_1_DNSRegister                                                    string      `json:"NIC.1.DNSRegister"`
		NIC_1_DedicatedNICScanTime                                           int         `json:"NIC.1.DedicatedNICScanTime"`
		NIC_1_Duplex                                                         string      `json:"NIC.1.Duplex"`
		NIC_1_Enable                                                         string      `json:"NIC.1.Enable"`
		NIC_1_Failover                                                       string      `json:"NIC.1.Failover"`
		NIC_1_MACAddress                                                     string      `json:"NIC.1.MACAddress"`
		NIC_1_MTU                                                            int         `json:"NIC.1.MTU"`
		NIC_1_Selection                                                      string      `json:"NIC.1.Selection"`
		NIC_1_SharedNICScanTime                                              int         `json:"NIC.1.SharedNICScanTime"`
		NIC_1_Speed                                                          string      `json:"NIC.1.Speed"`
		NIC_1_SwitchConnection                                               string      `json:"NIC.1.SwitchConnection"`
		NIC_1_SwitchPortConnection                                           string      `json:"NIC.1.SwitchPortConnection"`
		NIC_1_VLanEnable                                                     string      `json:"NIC.1.VLanEnable"`
		NIC_1_VLanID                                                         int         `json:"NIC.1.VLanID"`
		NIC_1_VLanPort                                                       string      `json:"NIC.1.VLanPort"`
		NIC_1_VLanPriority                                                   int         `json:"NIC.1.VLanPriority"`
		NIC_1_VLanSetting                                                    string      `json:"NIC.1.VLanSetting"`
		NICStatic_1_DNSDomainFromDHCP                                        string      `json:"NICStatic.1.DNSDomainFromDHCP"`
		NICStatic_1_DNSDomainName                                            string      `json:"NICStatic.1.DNSDomainName"`
		NTPConfigGroup_1_NTP1                                                string      `json:"NTPConfigGroup.1.NTP1"`
		NTPConfigGroup_1_NTP2                                                string      `json:"NTPConfigGroup.1.NTP2"`
		NTPConfigGroup_1_NTP3                                                string      `json:"NTPConfigGroup.1.NTP3"`
		NTPConfigGroup_1_NTPEnable                                           string      `json:"NTPConfigGroup.1.NTPEnable"`
		NTPConfigGroup_1_NTPMaxDist                                          int         `json:"NTPConfigGroup.1.NTPMaxDist"`
		OS_BMC_1_AdminState                                                  string      `json:"OS-BMC.1.AdminState"`
		OS_BMC_1_OsIPAddress                                                 string      `json:"OS-BMC.1.OsIpAddress"`
		OS_BMC_1_PTCapability                                                string      `json:"OS-BMC.1.PTCapability"`
		OS_BMC_1_PTMode                                                      string      `json:"OS-BMC.1.PTMode"`
		OS_BMC_1_UsbNicIPAddress                                             string      `json:"OS-BMC.1.UsbNicIpAddress"`
		PlatformCapability_1_FrontPortUSBConfiguration                       string      `json:"PlatformCapability.1.FrontPortUSBConfiguration"`
		RFS_1_AttachMode                                                     string      `json:"RFS.1.AttachMode"`
		RFS_1_Enable                                                         string      `json:"RFS.1.Enable"`
		RFS_1_Image                                                          string      `json:"RFS.1.Image"`
		RFS_1_MediaAttachState                                               string      `json:"RFS.1.MediaAttachState"`
		RFS_1_Password                                                       interface{} `json:"RFS.1.Password"`
		RFS_1_Status                                                         string      `json:"RFS.1.Status"`
		RFS_1_User                                                           string      `json:"RFS.1.User"`
		RFS_1_WriteProtected                                                 string      `json:"RFS.1.WriteProtected"`
		Racadm_1_Enable                                                      string      `json:"Racadm.1.Enable"`
		Racadm_1_MaxSessions                                                 int         `json:"Racadm.1.MaxSessions"`
		Racadm_1_Timeout                                                     int         `json:"Racadm.1.Timeout"`
		Redfish_1_Enable                                                     string      `json:"Redfish.1.Enable"`
		RedfishEventing_1_DeliveryRetryAttempts                              int         `json:"RedfishEventing.1.DeliveryRetryAttempts"`
		RedfishEventing_1_DeliveryRetryIntervalInSeconds                     int         `json:"RedfishEventing.1.DeliveryRetryIntervalInSeconds"`
		RedfishEventing_1_IgnoreCertificateErrors                            string      `json:"RedfishEventing.1.IgnoreCertificateErrors"`
		RemoteHosts_1_SMTPAuthentication                                     string      `json:"RemoteHosts.1.SMTPAuthentication"`
		RemoteHosts_1_SMTPPassword                                           interface{} `json:"RemoteHosts.1.SMTPPassword"`
		RemoteHosts_1_SMTPPort                                               int         `json:"RemoteHosts.1.SMTPPort"`
		RemoteHosts_1_SMTPServerIPAddress                                    string      `json:"RemoteHosts.1.SMTPServerIPAddress"`
		RemoteHosts_1_SMTPUserName                                           string      `json:"RemoteHosts.1.SMTPUserName"`
		SNMP_1_AgentCommunity                                                string      `json:"SNMP.1.AgentCommunity"`
		SNMP_1_AgentEnable                                                   string      `json:"SNMP.1.AgentEnable"`
		SNMP_1_AlertPort                                                     int         `json:"SNMP.1.AlertPort"`
		SNMP_1_DiscoveryPort                                                 int         `json:"SNMP.1.DiscoveryPort"`
		SNMP_1_EngineID                                                      string      `json:"SNMP.1.EngineID"`
		SNMP_1_SNMPProtocol                                                  string      `json:"SNMP.1.SNMPProtocol"`
		SNMP_1_TrapFormat                                                    string      `json:"SNMP.1.TrapFormat"`
		SNMPAlert_1_Destination                                              string      `json:"SNMPAlert.1.Destination"`
		SNMPAlert_1_SNMPv3UserID                                             int         `json:"SNMPAlert.1.SNMPv3UserID"`
		SNMPAlert_1_SNMPv3Username                                           string      `json:"SNMPAlert.1.SNMPv3Username"`
		SNMPAlert_1_State                                                    string      `json:"SNMPAlert.1.State"`
		SNMPAlert_2_Destination                                              string      `json:"SNMPAlert.2.Destination"`
		SNMPAlert_2_SNMPv3UserID                                             int         `json:"SNMPAlert.2.SNMPv3UserID"`
		SNMPAlert_2_SNMPv3Username                                           string      `json:"SNMPAlert.2.SNMPv3Username"`
		SNMPAlert_2_State                                                    string      `json:"SNMPAlert.2.State"`
		SNMPAlert_3_Destination                                              string      `json:"SNMPAlert.3.Destination"`
		SNMPAlert_3_SNMPv3UserID                                             int         `json:"SNMPAlert.3.SNMPv3UserID"`
		SNMPAlert_3_SNMPv3Username                                           string      `json:"SNMPAlert.3.SNMPv3Username"`
		SNMPAlert_3_State                                                    string      `json:"SNMPAlert.3.State"`
		SNMPAlert_4_Destination                                              string      `json:"SNMPAlert.4.Destination"`
		SNMPAlert_4_SNMPv3UserID                                             int         `json:"SNMPAlert.4.SNMPv3UserID"`
		SNMPAlert_4_SNMPv3Username                                           string      `json:"SNMPAlert.4.SNMPv3Username"`
		SNMPAlert_4_State                                                    string      `json:"SNMPAlert.4.State"`
		SNMPAlert_5_Destination                                              string      `json:"SNMPAlert.5.Destination"`
		SNMPAlert_5_SNMPv3UserID                                             int         `json:"SNMPAlert.5.SNMPv3UserID"`
		SNMPAlert_5_SNMPv3Username                                           string      `json:"SNMPAlert.5.SNMPv3Username"`
		SNMPAlert_5_State                                                    string      `json:"SNMPAlert.5.State"`
		SNMPAlert_6_Destination                                              string      `json:"SNMPAlert.6.Destination"`
		SNMPAlert_6_SNMPv3UserID                                             int         `json:"SNMPAlert.6.SNMPv3UserID"`
		SNMPAlert_6_SNMPv3Username                                           string      `json:"SNMPAlert.6.SNMPv3Username"`
		SNMPAlert_6_State                                                    string      `json:"SNMPAlert.6.State"`
		SNMPAlert_7_Destination                                              string      `json:"SNMPAlert.7.Destination"`
		SNMPAlert_7_SNMPv3UserID                                             int         `json:"SNMPAlert.7.SNMPv3UserID"`
		SNMPAlert_7_SNMPv3Username                                           string      `json:"SNMPAlert.7.SNMPv3Username"`
		SNMPAlert_7_State                                                    string      `json:"SNMPAlert.7.State"`
		SNMPAlert_8_Destination                                              string      `json:"SNMPAlert.8.Destination"`
		SNMPAlert_8_SNMPv3UserID                                             int         `json:"SNMPAlert.8.SNMPv3UserID"`
		SNMPAlert_8_SNMPv3Username                                           string      `json:"SNMPAlert.8.SNMPv3Username"`
		SNMPAlert_8_State                                                    string      `json:"SNMPAlert.8.State"`
		SNMPTrapIPv4_1_DestIPv4Addr                                          string      `json:"SNMPTrapIPv4.1.DestIPv4Addr"`
		SNMPTrapIPv4_1_DestinationNum                                        int         `json:"SNMPTrapIPv4.1.DestinationNum"`
		SNMPTrapIPv4_1_State                                                 string      `json:"SNMPTrapIPv4.1.State"`
		SNMPTrapIPv4_2_DestIPv4Addr                                          string      `json:"SNMPTrapIPv4.2.DestIPv4Addr"`
		SNMPTrapIPv4_2_DestinationNum                                        int         `json:"SNMPTrapIPv4.2.DestinationNum"`
		SNMPTrapIPv4_2_State                                                 string      `json:"SNMPTrapIPv4.2.State"`
		SNMPTrapIPv4_3_DestIPv4Addr                                          string      `json:"SNMPTrapIPv4.3.DestIPv4Addr"`
		SNMPTrapIPv4_3_DestinationNum                                        int         `json:"SNMPTrapIPv4.3.DestinationNum"`
		SNMPTrapIPv4_3_State                                                 string      `json:"SNMPTrapIPv4.3.State"`
		SNMPTrapIPv4_4_DestIPv4Addr                                          string      `json:"SNMPTrapIPv4.4.DestIPv4Addr"`
		SNMPTrapIPv4_4_DestinationNum                                        int         `json:"SNMPTrapIPv4.4.DestinationNum"`
		SNMPTrapIPv4_4_State                                                 string      `json:"SNMPTrapIPv4.4.State"`
		SNMPTrapIPv6_1_DestIPv6Addr                                          string      `json:"SNMPTrapIPv6.1.DestIPv6Addr"`
		SNMPTrapIPv6_1_DestinationNum                                        int         `json:"SNMPTrapIPv6.1.DestinationNum"`
		SNMPTrapIPv6_1_State                                                 string      `json:"SNMPTrapIPv6.1.State"`
		SNMPTrapIPv6_2_DestIPv6Addr                                          string      `json:"SNMPTrapIPv6.2.DestIPv6Addr"`
		SNMPTrapIPv6_2_DestinationNum                                        int         `json:"SNMPTrapIPv6.2.DestinationNum"`
		SNMPTrapIPv6_2_State                                                 string      `json:"SNMPTrapIPv6.2.State"`
		SNMPTrapIPv6_3_DestIPv6Addr                                          string      `json:"SNMPTrapIPv6.3.DestIPv6Addr"`
		SNMPTrapIPv6_3_DestinationNum                                        int         `json:"SNMPTrapIPv6.3.DestinationNum"`
		SNMPTrapIPv6_3_State                                                 string      `json:"SNMPTrapIPv6.3.State"`
		SSH_1_Enable                                                         string      `json:"SSH.1.Enable"`
		SSH_1_MaxSessions                                                    int         `json:"SSH.1.MaxSessions"`
		SSH_1_Port                                                           int         `json:"SSH.1.Port"`
		SSH_1_Timeout                                                        int         `json:"SSH.1.Timeout"`
		Security_1_CsrCommonName                                             string      `json:"Security.1.CsrCommonName"`
		Security_1_CsrCountryCode                                            string      `json:"Security.1.CsrCountryCode"`
		Security_1_CsrEmailAddr                                              string      `json:"Security.1.CsrEmailAddr"`
		Security_1_CsrKeySize                                                string      `json:"Security.1.CsrKeySize"`
		Security_1_CsrLocalityName                                           string      `json:"Security.1.CsrLocalityName"`
		Security_1_CsrOrganizationName                                       string      `json:"Security.1.CsrOrganizationName"`
		Security_1_CsrOrganizationUnit                                       string      `json:"Security.1.CsrOrganizationUnit"`
		Security_1_CsrStateName                                              string      `json:"Security.1.CsrStateName"`
		Security_1_FIPSMode                                                  string      `json:"Security.1.FIPSMode"`
		Serial_1_BaudRate                                                    string      `json:"Serial.1.BaudRate"`
		Serial_1_Command                                                     string      `json:"Serial.1.Command"`
		Serial_1_Enable                                                      string      `json:"Serial.1.Enable"`
		Serial_1_HistorySize                                                 int         `json:"Serial.1.HistorySize"`
		Serial_1_IdleTimeout                                                 int         `json:"Serial.1.IdleTimeout"`
		Serial_1_NoAuth                                                      string      `json:"Serial.1.NoAuth"`
		SerialRedirection_1_Enable                                           string      `json:"SerialRedirection.1.Enable"`
		SerialRedirection_1_QuitKey                                          string      `json:"SerialRedirection.1.QuitKey"`
		ServerBoot_1_BootOnce                                                string      `json:"ServerBoot.1.BootOnce"`
		ServerBoot_1_FirstBootDevice                                         string      `json:"ServerBoot.1.FirstBootDevice"`
		ServiceModule_1_HostSNMPAlert                                        string      `json:"ServiceModule.1.HostSNMPAlert"`
		ServiceModule_1_HostSNMPGet                                          string      `json:"ServiceModule.1.HostSNMPGet"`
		ServiceModule_1_LCLReplication                                       string      `json:"ServiceModule.1.LCLReplication"`
		ServiceModule_1_OMSAPresence                                         string      `json:"ServiceModule.1.OMSAPresence"`
		ServiceModule_1_OSInfo                                               string      `json:"ServiceModule.1.OSInfo"`
		ServiceModule_1_SNMPOnHostOS                                         string      `json:"ServiceModule.1.SNMPOnHostOS"`
		ServiceModule_1_ServiceModuleEnable                                  string      `json:"ServiceModule.1.ServiceModuleEnable"`
		ServiceModule_1_ServiceModuleState                                   string      `json:"ServiceModule.1.ServiceModuleState"`
		ServiceModule_1_ServiceModuleVersion                                 string      `json:"ServiceModule.1.ServiceModuleVersion"`
		ServiceModule_1_WMIInfo                                              string      `json:"ServiceModule.1.WMIInfo"`
		ServiceModule_1_WatchdogRecoveryAction                               string      `json:"ServiceModule.1.WatchdogRecoveryAction"`
		ServiceModule_1_WatchdogResetTime                                    int         `json:"ServiceModule.1.WatchdogResetTime"`
		ServiceModule_1_WatchdogState                                        string      `json:"ServiceModule.1.WatchdogState"`
		ServiceModule_1_iDRACHardReset                                       string      `json:"ServiceModule.1.iDRACHardReset"`
		SmartCard_1_SmartCardCRLEnable                                       string      `json:"SmartCard.1.SmartCardCRLEnable"`
		SmartCard_1_SmartCardLogonEnable                                     string      `json:"SmartCard.1.SmartCardLogonEnable"`
		SupportAssist_1_DefaultIPAddress                                     string      `json:"SupportAssist.1.DefaultIPAddress"`
		SupportAssist_1_DefaultPassword                                      interface{} `json:"SupportAssist.1.DefaultPassword"`
		SupportAssist_1_DefaultProtocol                                      string      `json:"SupportAssist.1.DefaultProtocol"`
		SupportAssist_1_DefaultShareName                                     string      `json:"SupportAssist.1.DefaultShareName"`
		SupportAssist_1_DefaultUserName                                      string      `json:"SupportAssist.1.DefaultUserName"`
		SupportAssist_1_DefaultWorkgroupName                                 string      `json:"SupportAssist.1.DefaultWorkgroupName"`
		SupportAssist_1_EmailOptIn                                           string      `json:"SupportAssist.1.EmailOptIn"`
		SupportAssist_1_EventBasedAutoCollection                             string      `json:"SupportAssist.1.EventBasedAutoCollection"`
		SupportAssist_1_FilterAutoCollections                                string      `json:"SupportAssist.1.FilterAutoCollections"`
		SupportAssist_1_HostOSProxyAddress                                   string      `json:"SupportAssist.1.HostOSProxyAddress"`
		SupportAssist_1_HostOSProxyConfigured                                string      `json:"SupportAssist.1.HostOSProxyConfigured"`
		SupportAssist_1_HostOSProxyPassword                                  interface{} `json:"SupportAssist.1.HostOSProxyPassword"`
		SupportAssist_1_HostOSProxyPort                                      int         `json:"SupportAssist.1.HostOSProxyPort"`
		SupportAssist_1_HostOSProxyUserName                                  string      `json:"SupportAssist.1.HostOSProxyUserName"`
		SupportAssist_1_NativeOSLogsCollectionSupported                      string      `json:"SupportAssist.1.NativeOSLogsCollectionSupported"`
		SupportAssist_1_PreferredLanguage                                    string      `json:"SupportAssist.1.PreferredLanguage"`
		SupportAssist_1_ProSupportPlusRecommendationsReport                  string      `json:"SupportAssist.1.ProSupportPlusRecommendationsReport"`
		SupportAssist_1_RegistrationID                                       string      `json:"SupportAssist.1.RegistrationID"`
		SupportAssist_1_SupportAssistEnableState                             string      `json:"SupportAssist.1.SupportAssistEnableState"`
		SwitchConnectionView_1_Enable                                        string      `json:"SwitchConnectionView.1.Enable"`
		SysInfo_1_LocalConsoleLockOut                                        int         `json:"SysInfo.1.LocalConsoleLockOut"`
		SysInfo_1_SystemRev                                                  int         `json:"SysInfo.1.SystemRev"`
		SysLog_1_Port                                                        int         `json:"SysLog.1.Port"`
		SysLog_1_PowerLogEnable                                              string      `json:"SysLog.1.PowerLogEnable"`
		SysLog_1_PowerLogInterval                                            int         `json:"SysLog.1.PowerLogInterval"`
		SysLog_1_Server1                                                     string      `json:"SysLog.1.Server1"`
		SysLog_1_Server2                                                     string      `json:"SysLog.1.Server2"`
		SysLog_1_Server3                                                     string      `json:"SysLog.1.Server3"`
		SysLog_1_SysLogEnable                                                string      `json:"SysLog.1.SysLogEnable"`
		Telnet_1_Enable                                                      string      `json:"Telnet.1.Enable"`
		Telnet_1_MaxSessions                                                 int         `json:"Telnet.1.MaxSessions"`
		Telnet_1_Port                                                        int         `json:"Telnet.1.Port"`
		Telnet_1_Timeout                                                     int         `json:"Telnet.1.Timeout"`
		Time_1_DayLightOffset                                                int         `json:"Time.1.DayLightOffset"`
		Time_1_TimeZoneOffset                                                int         `json:"Time.1.TimeZoneOffset"`
		Time_1_Timezone                                                      string      `json:"Time.1.Timezone"`
		USB_1_ConfigurationXML                                               string      `json:"USB.1.ConfigurationXML"`
		USB_1_ManagementPortMode                                             string      `json:"USB.1.ManagementPortMode"`
		USB_1_PortStatus                                                     string      `json:"USB.1.PortStatus"`
		USB_1_ZipPassword                                                    interface{} `json:"USB.1.ZipPassword"`
		Update_1_FwUpdateIPAddr                                              string      `json:"Update.1.FwUpdateIPAddr"`
		Update_1_FwUpdatePath                                                string      `json:"Update.1.FwUpdatePath"`
		Update_1_FwUpdateTFTPEnable                                          string      `json:"Update.1.FwUpdateTFTPEnable"`
		UserDomain_1_Name                                                    string      `json:"UserDomain.1.Name"`
		UserDomain_10_Name                                                   string      `json:"UserDomain.10.Name"`
		UserDomain_11_Name                                                   string      `json:"UserDomain.11.Name"`
		UserDomain_12_Name                                                   string      `json:"UserDomain.12.Name"`
		UserDomain_13_Name                                                   string      `json:"UserDomain.13.Name"`
		UserDomain_14_Name                                                   string      `json:"UserDomain.14.Name"`
		UserDomain_15_Name                                                   string      `json:"UserDomain.15.Name"`
		UserDomain_16_Name                                                   string      `json:"UserDomain.16.Name"`
		UserDomain_17_Name                                                   string      `json:"UserDomain.17.Name"`
		UserDomain_18_Name                                                   string      `json:"UserDomain.18.Name"`
		UserDomain_19_Name                                                   string      `json:"UserDomain.19.Name"`
		UserDomain_2_Name                                                    string      `json:"UserDomain.2.Name"`
		UserDomain_20_Name                                                   string      `json:"UserDomain.20.Name"`
		UserDomain_21_Name                                                   string      `json:"UserDomain.21.Name"`
		UserDomain_22_Name                                                   string      `json:"UserDomain.22.Name"`
		UserDomain_23_Name                                                   string      `json:"UserDomain.23.Name"`
		UserDomain_24_Name                                                   string      `json:"UserDomain.24.Name"`
		UserDomain_25_Name                                                   string      `json:"UserDomain.25.Name"`
		UserDomain_26_Name                                                   string      `json:"UserDomain.26.Name"`
		UserDomain_27_Name                                                   string      `json:"UserDomain.27.Name"`
		UserDomain_28_Name                                                   string      `json:"UserDomain.28.Name"`
		UserDomain_29_Name                                                   string      `json:"UserDomain.29.Name"`
		UserDomain_3_Name                                                    string      `json:"UserDomain.3.Name"`
		UserDomain_30_Name                                                   string      `json:"UserDomain.30.Name"`
		UserDomain_31_Name                                                   string      `json:"UserDomain.31.Name"`
		UserDomain_32_Name                                                   string      `json:"UserDomain.32.Name"`
		UserDomain_33_Name                                                   string      `json:"UserDomain.33.Name"`
		UserDomain_34_Name                                                   string      `json:"UserDomain.34.Name"`
		UserDomain_35_Name                                                   string      `json:"UserDomain.35.Name"`
		UserDomain_36_Name                                                   string      `json:"UserDomain.36.Name"`
		UserDomain_37_Name                                                   string      `json:"UserDomain.37.Name"`
		UserDomain_38_Name                                                   string      `json:"UserDomain.38.Name"`
		UserDomain_39_Name                                                   string      `json:"UserDomain.39.Name"`
		UserDomain_4_Name                                                    string      `json:"UserDomain.4.Name"`
		UserDomain_40_Name                                                   string      `json:"UserDomain.40.Name"`
		UserDomain_5_Name                                                    string      `json:"UserDomain.5.Name"`
		UserDomain_6_Name                                                    string      `json:"UserDomain.6.Name"`
		UserDomain_7_Name                                                    string      `json:"UserDomain.7.Name"`
		UserDomain_8_Name                                                    string      `json:"UserDomain.8.Name"`
		UserDomain_9_Name                                                    string      `json:"UserDomain.9.Name"`
		Users_1_AuthenticationProtocol                                       string      `json:"Users.1.AuthenticationProtocol"`
		Users_1_Enable                                                       string      `json:"Users.1.Enable"`
		Users_1_IPMIKey                                                      string      `json:"Users.1.IPMIKey"`
		Users_1_IpmiLanPrivilege                                             string      `json:"Users.1.IpmiLanPrivilege"`
		Users_1_IpmiSerialPrivilege                                          string      `json:"Users.1.IpmiSerialPrivilege"`
		Users_1_MD5v3Key                                                     string      `json:"Users.1.MD5v3Key"`
		Users_1_Password                                                     interface{} `json:"Users.1.Password"`
		Users_1_PrivacyProtocol                                              string      `json:"Users.1.PrivacyProtocol"`
		Users_1_Privilege                                                    int         `json:"Users.1.Privilege"`
		Users_1_ProtocolEnable                                               string      `json:"Users.1.ProtocolEnable"`
		Users_1_SHA1v3Key                                                    string      `json:"Users.1.SHA1v3Key"`
		Users_1_SHA256Password                                               string      `json:"Users.1.SHA256Password"`
		Users_1_SHA256PasswordSalt                                           string      `json:"Users.1.SHA256PasswordSalt"`
		Users_1_SolEnable                                                    string      `json:"Users.1.SolEnable"`
		Users_1_UserName                                                     string      `json:"Users.1.UserName"`
		Users_10_AuthenticationProtocol                                      string      `json:"Users.10.AuthenticationProtocol"`
		Users_10_Enable                                                      string      `json:"Users.10.Enable"`
		Users_10_IPMIKey                                                     string      `json:"Users.10.IPMIKey"`
		Users_10_IpmiLanPrivilege                                            string      `json:"Users.10.IpmiLanPrivilege"`
		Users_10_IpmiSerialPrivilege                                         string      `json:"Users.10.IpmiSerialPrivilege"`
		Users_10_MD5v3Key                                                    string      `json:"Users.10.MD5v3Key"`
		Users_10_Password                                                    interface{} `json:"Users.10.Password"`
		Users_10_PrivacyProtocol                                             string      `json:"Users.10.PrivacyProtocol"`
		Users_10_Privilege                                                   int         `json:"Users.10.Privilege"`
		Users_10_ProtocolEnable                                              string      `json:"Users.10.ProtocolEnable"`
		Users_10_SHA1v3Key                                                   string      `json:"Users.10.SHA1v3Key"`
		Users_10_SHA256Password                                              string      `json:"Users.10.SHA256Password"`
		Users_10_SHA256PasswordSalt                                          string      `json:"Users.10.SHA256PasswordSalt"`
		Users_10_SolEnable                                                   string      `json:"Users.10.SolEnable"`
		Users_10_UserName                                                    string      `json:"Users.10.UserName"`
		Users_11_AuthenticationProtocol                                      string      `json:"Users.11.AuthenticationProtocol"`
		Users_11_Enable                                                      string      `json:"Users.11.Enable"`
		Users_11_IPMIKey                                                     string      `json:"Users.11.IPMIKey"`
		Users_11_IpmiLanPrivilege                                            string      `json:"Users.11.IpmiLanPrivilege"`
		Users_11_IpmiSerialPrivilege                                         string      `json:"Users.11.IpmiSerialPrivilege"`
		Users_11_MD5v3Key                                                    string      `json:"Users.11.MD5v3Key"`
		Users_11_Password                                                    interface{} `json:"Users.11.Password"`
		Users_11_PrivacyProtocol                                             string      `json:"Users.11.PrivacyProtocol"`
		Users_11_Privilege                                                   int         `json:"Users.11.Privilege"`
		Users_11_ProtocolEnable                                              string      `json:"Users.11.ProtocolEnable"`
		Users_11_SHA1v3Key                                                   string      `json:"Users.11.SHA1v3Key"`
		Users_11_SHA256Password                                              string      `json:"Users.11.SHA256Password"`
		Users_11_SHA256PasswordSalt                                          string      `json:"Users.11.SHA256PasswordSalt"`
		Users_11_SolEnable                                                   string      `json:"Users.11.SolEnable"`
		Users_11_UserName                                                    string      `json:"Users.11.UserName"`
		Users_12_AuthenticationProtocol                                      string      `json:"Users.12.AuthenticationProtocol"`
		Users_12_Enable                                                      string      `json:"Users.12.Enable"`
		Users_12_IPMIKey                                                     string      `json:"Users.12.IPMIKey"`
		Users_12_IpmiLanPrivilege                                            string      `json:"Users.12.IpmiLanPrivilege"`
		Users_12_IpmiSerialPrivilege                                         string      `json:"Users.12.IpmiSerialPrivilege"`
		Users_12_MD5v3Key                                                    string      `json:"Users.12.MD5v3Key"`
		Users_12_Password                                                    interface{} `json:"Users.12.Password"`
		Users_12_PrivacyProtocol                                             string      `json:"Users.12.PrivacyProtocol"`
		Users_12_Privilege                                                   int         `json:"Users.12.Privilege"`
		Users_12_ProtocolEnable                                              string      `json:"Users.12.ProtocolEnable"`
		Users_12_SHA1v3Key                                                   string      `json:"Users.12.SHA1v3Key"`
		Users_12_SHA256Password                                              string      `json:"Users.12.SHA256Password"`
		Users_12_SHA256PasswordSalt                                          string      `json:"Users.12.SHA256PasswordSalt"`
		Users_12_SolEnable                                                   string      `json:"Users.12.SolEnable"`
		Users_12_UserName                                                    string      `json:"Users.12.UserName"`
		Users_13_AuthenticationProtocol                                      string      `json:"Users.13.AuthenticationProtocol"`
		Users_13_Enable                                                      string      `json:"Users.13.Enable"`
		Users_13_IPMIKey                                                     string      `json:"Users.13.IPMIKey"`
		Users_13_IpmiLanPrivilege                                            string      `json:"Users.13.IpmiLanPrivilege"`
		Users_13_IpmiSerialPrivilege                                         string      `json:"Users.13.IpmiSerialPrivilege"`
		Users_13_MD5v3Key                                                    string      `json:"Users.13.MD5v3Key"`
		Users_13_Password                                                    interface{} `json:"Users.13.Password"`
		Users_13_PrivacyProtocol                                             string      `json:"Users.13.PrivacyProtocol"`
		Users_13_Privilege                                                   int         `json:"Users.13.Privilege"`
		Users_13_ProtocolEnable                                              string      `json:"Users.13.ProtocolEnable"`
		Users_13_SHA1v3Key                                                   string      `json:"Users.13.SHA1v3Key"`
		Users_13_SHA256Password                                              string      `json:"Users.13.SHA256Password"`
		Users_13_SHA256PasswordSalt                                          string      `json:"Users.13.SHA256PasswordSalt"`
		Users_13_SolEnable                                                   string      `json:"Users.13.SolEnable"`
		Users_13_UserName                                                    string      `json:"Users.13.UserName"`
		Users_14_AuthenticationProtocol                                      string      `json:"Users.14.AuthenticationProtocol"`
		Users_14_Enable                                                      string      `json:"Users.14.Enable"`
		Users_14_IPMIKey                                                     string      `json:"Users.14.IPMIKey"`
		Users_14_IpmiLanPrivilege                                            string      `json:"Users.14.IpmiLanPrivilege"`
		Users_14_IpmiSerialPrivilege                                         string      `json:"Users.14.IpmiSerialPrivilege"`
		Users_14_MD5v3Key                                                    string      `json:"Users.14.MD5v3Key"`
		Users_14_Password                                                    interface{} `json:"Users.14.Password"`
		Users_14_PrivacyProtocol                                             string      `json:"Users.14.PrivacyProtocol"`
		Users_14_Privilege                                                   int         `json:"Users.14.Privilege"`
		Users_14_ProtocolEnable                                              string      `json:"Users.14.ProtocolEnable"`
		Users_14_SHA1v3Key                                                   string      `json:"Users.14.SHA1v3Key"`
		Users_14_SHA256Password                                              string      `json:"Users.14.SHA256Password"`
		Users_14_SHA256PasswordSalt                                          string      `json:"Users.14.SHA256PasswordSalt"`
		Users_14_SolEnable                                                   string      `json:"Users.14.SolEnable"`
		Users_14_UserName                                                    string      `json:"Users.14.UserName"`
		Users_15_AuthenticationProtocol                                      string      `json:"Users.15.AuthenticationProtocol"`
		Users_15_Enable                                                      string      `json:"Users.15.Enable"`
		Users_15_IPMIKey                                                     string      `json:"Users.15.IPMIKey"`
		Users_15_IpmiLanPrivilege                                            string      `json:"Users.15.IpmiLanPrivilege"`
		Users_15_IpmiSerialPrivilege                                         string      `json:"Users.15.IpmiSerialPrivilege"`
		Users_15_MD5v3Key                                                    string      `json:"Users.15.MD5v3Key"`
		Users_15_Password                                                    interface{} `json:"Users.15.Password"`
		Users_15_PrivacyProtocol                                             string      `json:"Users.15.PrivacyProtocol"`
		Users_15_Privilege                                                   int         `json:"Users.15.Privilege"`
		Users_15_ProtocolEnable                                              string      `json:"Users.15.ProtocolEnable"`
		Users_15_SHA1v3Key                                                   string      `json:"Users.15.SHA1v3Key"`
		Users_15_SHA256Password                                              string      `json:"Users.15.SHA256Password"`
		Users_15_SHA256PasswordSalt                                          string      `json:"Users.15.SHA256PasswordSalt"`
		Users_15_SolEnable                                                   string      `json:"Users.15.SolEnable"`
		Users_15_UserName                                                    string      `json:"Users.15.UserName"`
		Users_16_AuthenticationProtocol                                      string      `json:"Users.16.AuthenticationProtocol"`
		Users_16_Enable                                                      string      `json:"Users.16.Enable"`
		Users_16_IPMIKey                                                     string      `json:"Users.16.IPMIKey"`
		Users_16_IpmiLanPrivilege                                            string      `json:"Users.16.IpmiLanPrivilege"`
		Users_16_IpmiSerialPrivilege                                         string      `json:"Users.16.IpmiSerialPrivilege"`
		Users_16_MD5v3Key                                                    string      `json:"Users.16.MD5v3Key"`
		Users_16_Password                                                    interface{} `json:"Users.16.Password"`
		Users_16_PrivacyProtocol                                             string      `json:"Users.16.PrivacyProtocol"`
		Users_16_Privilege                                                   int         `json:"Users.16.Privilege"`
		Users_16_ProtocolEnable                                              string      `json:"Users.16.ProtocolEnable"`
		Users_16_SHA1v3Key                                                   string      `json:"Users.16.SHA1v3Key"`
		Users_16_SHA256Password                                              string      `json:"Users.16.SHA256Password"`
		Users_16_SHA256PasswordSalt                                          string      `json:"Users.16.SHA256PasswordSalt"`
		Users_16_SolEnable                                                   string      `json:"Users.16.SolEnable"`
		Users_16_UserName                                                    string      `json:"Users.16.UserName"`
		Users_2_AuthenticationProtocol                                       string      `json:"Users.2.AuthenticationProtocol"`
		Users_2_Enable                                                       string      `json:"Users.2.Enable"`
		Users_2_IPMIKey                                                      string      `json:"Users.2.IPMIKey"`
		Users_2_IpmiLanPrivilege                                             string      `json:"Users.2.IpmiLanPrivilege"`
		Users_2_IpmiSerialPrivilege                                          string      `json:"Users.2.IpmiSerialPrivilege"`
		Users_2_MD5v3Key                                                     string      `json:"Users.2.MD5v3Key"`
		Users_2_Password                                                     interface{} `json:"Users.2.Password"`
		Users_2_PrivacyProtocol                                              string      `json:"Users.2.PrivacyProtocol"`
		Users_2_Privilege                                                    int         `json:"Users.2.Privilege"`
		Users_2_ProtocolEnable                                               string      `json:"Users.2.ProtocolEnable"`
		Users_2_SHA1v3Key                                                    string      `json:"Users.2.SHA1v3Key"`
		Users_2_SHA256Password                                               string      `json:"Users.2.SHA256Password"`
		Users_2_SHA256PasswordSalt                                           string      `json:"Users.2.SHA256PasswordSalt"`
		Users_2_SolEnable                                                    string      `json:"Users.2.SolEnable"`
		Users_2_UserName                                                     string      `json:"Users.2.UserName"`
		Users_3_AuthenticationProtocol                                       string      `json:"Users.3.AuthenticationProtocol"`
		Users_3_Enable                                                       string      `json:"Users.3.Enable"`
		Users_3_IPMIKey                                                      string      `json:"Users.3.IPMIKey"`
		Users_3_IpmiLanPrivilege                                             string      `json:"Users.3.IpmiLanPrivilege"`
		Users_3_IpmiSerialPrivilege                                          string      `json:"Users.3.IpmiSerialPrivilege"`
		Users_3_MD5v3Key                                                     string      `json:"Users.3.MD5v3Key"`
		Users_3_Password                                                     interface{} `json:"Users.3.Password"`
		Users_3_PrivacyProtocol                                              string      `json:"Users.3.PrivacyProtocol"`
		Users_3_Privilege                                                    int         `json:"Users.3.Privilege"`
		Users_3_ProtocolEnable                                               string      `json:"Users.3.ProtocolEnable"`
		Users_3_SHA1v3Key                                                    string      `json:"Users.3.SHA1v3Key"`
		Users_3_SHA256Password                                               string      `json:"Users.3.SHA256Password"`
		Users_3_SHA256PasswordSalt                                           string      `json:"Users.3.SHA256PasswordSalt"`
		Users_3_SolEnable                                                    string      `json:"Users.3.SolEnable"`
		Users_3_UserName                                                     string      `json:"Users.3.UserName"`
		Users_4_AuthenticationProtocol                                       string      `json:"Users.4.AuthenticationProtocol"`
		Users_4_Enable                                                       string      `json:"Users.4.Enable"`
		Users_4_IPMIKey                                                      string      `json:"Users.4.IPMIKey"`
		Users_4_IpmiLanPrivilege                                             string      `json:"Users.4.IpmiLanPrivilege"`
		Users_4_IpmiSerialPrivilege                                          string      `json:"Users.4.IpmiSerialPrivilege"`
		Users_4_MD5v3Key                                                     string      `json:"Users.4.MD5v3Key"`
		Users_4_Password                                                     interface{} `json:"Users.4.Password"`
		Users_4_PrivacyProtocol                                              string      `json:"Users.4.PrivacyProtocol"`
		Users_4_Privilege                                                    int         `json:"Users.4.Privilege"`
		Users_4_ProtocolEnable                                               string      `json:"Users.4.ProtocolEnable"`
		Users_4_SHA1v3Key                                                    string      `json:"Users.4.SHA1v3Key"`
		Users_4_SHA256Password                                               string      `json:"Users.4.SHA256Password"`
		Users_4_SHA256PasswordSalt                                           string      `json:"Users.4.SHA256PasswordSalt"`
		Users_4_SolEnable                                                    string      `json:"Users.4.SolEnable"`
		Users_4_UserName                                                     string      `json:"Users.4.UserName"`
		Users_5_AuthenticationProtocol                                       string      `json:"Users.5.AuthenticationProtocol"`
		Users_5_Enable                                                       string      `json:"Users.5.Enable"`
		Users_5_IPMIKey                                                      string      `json:"Users.5.IPMIKey"`
		Users_5_IpmiLanPrivilege                                             string      `json:"Users.5.IpmiLanPrivilege"`
		Users_5_IpmiSerialPrivilege                                          string      `json:"Users.5.IpmiSerialPrivilege"`
		Users_5_MD5v3Key                                                     string      `json:"Users.5.MD5v3Key"`
		Users_5_Password                                                     interface{} `json:"Users.5.Password"`
		Users_5_PrivacyProtocol                                              string      `json:"Users.5.PrivacyProtocol"`
		Users_5_Privilege                                                    int         `json:"Users.5.Privilege"`
		Users_5_ProtocolEnable                                               string      `json:"Users.5.ProtocolEnable"`
		Users_5_SHA1v3Key                                                    string      `json:"Users.5.SHA1v3Key"`
		Users_5_SHA256Password                                               string      `json:"Users.5.SHA256Password"`
		Users_5_SHA256PasswordSalt                                           string      `json:"Users.5.SHA256PasswordSalt"`
		Users_5_SolEnable                                                    string      `json:"Users.5.SolEnable"`
		Users_5_UserName                                                     string      `json:"Users.5.UserName"`
		Users_6_AuthenticationProtocol                                       string      `json:"Users.6.AuthenticationProtocol"`
		Users_6_Enable                                                       string      `json:"Users.6.Enable"`
		Users_6_IPMIKey                                                      string      `json:"Users.6.IPMIKey"`
		Users_6_IpmiLanPrivilege                                             string      `json:"Users.6.IpmiLanPrivilege"`
		Users_6_IpmiSerialPrivilege                                          string      `json:"Users.6.IpmiSerialPrivilege"`
		Users_6_MD5v3Key                                                     string      `json:"Users.6.MD5v3Key"`
		Users_6_Password                                                     interface{} `json:"Users.6.Password"`
		Users_6_PrivacyProtocol                                              string      `json:"Users.6.PrivacyProtocol"`
		Users_6_Privilege                                                    int         `json:"Users.6.Privilege"`
		Users_6_ProtocolEnable                                               string      `json:"Users.6.ProtocolEnable"`
		Users_6_SHA1v3Key                                                    string      `json:"Users.6.SHA1v3Key"`
		Users_6_SHA256Password                                               string      `json:"Users.6.SHA256Password"`
		Users_6_SHA256PasswordSalt                                           string      `json:"Users.6.SHA256PasswordSalt"`
		Users_6_SolEnable                                                    string      `json:"Users.6.SolEnable"`
		Users_6_UserName                                                     string      `json:"Users.6.UserName"`
		Users_7_AuthenticationProtocol                                       string      `json:"Users.7.AuthenticationProtocol"`
		Users_7_Enable                                                       string      `json:"Users.7.Enable"`
		Users_7_IPMIKey                                                      string      `json:"Users.7.IPMIKey"`
		Users_7_IpmiLanPrivilege                                             string      `json:"Users.7.IpmiLanPrivilege"`
		Users_7_IpmiSerialPrivilege                                          string      `json:"Users.7.IpmiSerialPrivilege"`
		Users_7_MD5v3Key                                                     string      `json:"Users.7.MD5v3Key"`
		Users_7_Password                                                     interface{} `json:"Users.7.Password"`
		Users_7_PrivacyProtocol                                              string      `json:"Users.7.PrivacyProtocol"`
		Users_7_Privilege                                                    int         `json:"Users.7.Privilege"`
		Users_7_ProtocolEnable                                               string      `json:"Users.7.ProtocolEnable"`
		Users_7_SHA1v3Key                                                    string      `json:"Users.7.SHA1v3Key"`
		Users_7_SHA256Password                                               string      `json:"Users.7.SHA256Password"`
		Users_7_SHA256PasswordSalt                                           string      `json:"Users.7.SHA256PasswordSalt"`
		Users_7_SolEnable                                                    string      `json:"Users.7.SolEnable"`
		Users_7_UserName                                                     string      `json:"Users.7.UserName"`
		Users_8_AuthenticationProtocol                                       string      `json:"Users.8.AuthenticationProtocol"`
		Users_8_Enable                                                       string      `json:"Users.8.Enable"`
		Users_8_IPMIKey                                                      string      `json:"Users.8.IPMIKey"`
		Users_8_IpmiLanPrivilege                                             string      `json:"Users.8.IpmiLanPrivilege"`
		Users_8_IpmiSerialPrivilege                                          string      `json:"Users.8.IpmiSerialPrivilege"`
		Users_8_MD5v3Key                                                     string      `json:"Users.8.MD5v3Key"`
		Users_8_Password                                                     interface{} `json:"Users.8.Password"`
		Users_8_PrivacyProtocol                                              string      `json:"Users.8.PrivacyProtocol"`
		Users_8_Privilege                                                    int         `json:"Users.8.Privilege"`
		Users_8_ProtocolEnable                                               string      `json:"Users.8.ProtocolEnable"`
		Users_8_SHA1v3Key                                                    string      `json:"Users.8.SHA1v3Key"`
		Users_8_SHA256Password                                               string      `json:"Users.8.SHA256Password"`
		Users_8_SHA256PasswordSalt                                           string      `json:"Users.8.SHA256PasswordSalt"`
		Users_8_SolEnable                                                    string      `json:"Users.8.SolEnable"`
		Users_8_UserName                                                     string      `json:"Users.8.UserName"`
		Users_9_AuthenticationProtocol                                       string      `json:"Users.9.AuthenticationProtocol"`
		Users_9_Enable                                                       string      `json:"Users.9.Enable"`
		Users_9_IPMIKey                                                      string      `json:"Users.9.IPMIKey"`
		Users_9_IpmiLanPrivilege                                             string      `json:"Users.9.IpmiLanPrivilege"`
		Users_9_IpmiSerialPrivilege                                          string      `json:"Users.9.IpmiSerialPrivilege"`
		Users_9_MD5v3Key                                                     string      `json:"Users.9.MD5v3Key"`
		Users_9_Password                                                     interface{} `json:"Users.9.Password"`
		Users_9_PrivacyProtocol                                              string      `json:"Users.9.PrivacyProtocol"`
		Users_9_Privilege                                                    int         `json:"Users.9.Privilege"`
		Users_9_ProtocolEnable                                               string      `json:"Users.9.ProtocolEnable"`
		Users_9_SHA1v3Key                                                    string      `json:"Users.9.SHA1v3Key"`
		Users_9_SHA256Password                                               string      `json:"Users.9.SHA256Password"`
		Users_9_SHA256PasswordSalt                                           string      `json:"Users.9.SHA256PasswordSalt"`
		Users_9_SolEnable                                                    string      `json:"Users.9.SolEnable"`
		Users_9_UserName                                                     string      `json:"Users.9.UserName"`
		VNCServer_1_ActiveSessions                                           int         `json:"VNCServer.1.ActiveSessions"`
		VNCServer_1_Enable                                                   string      `json:"VNCServer.1.Enable"`
		VNCServer_1_LowerEncryptionBitLength                                 string      `json:"VNCServer.1.LowerEncryptionBitLength"`
		VNCServer_1_MaxSessions                                              int         `json:"VNCServer.1.MaxSessions"`
		VNCServer_1_Password                                                 interface{} `json:"VNCServer.1.Password"`
		VNCServer_1_Port                                                     int         `json:"VNCServer.1.Port"`
		VNCServer_1_SSLEncryptionBitLength                                   string      `json:"VNCServer.1.SSLEncryptionBitLength"`
		VNCServer_1_Timeout                                                  int         `json:"VNCServer.1.Timeout"`
		VirtualConsole_1_AccessPrivilege                                     string      `json:"VirtualConsole.1.AccessPrivilege"`
		VirtualConsole_1_AttachState                                         string      `json:"VirtualConsole.1.AttachState"`
		VirtualConsole_1_Enable                                              string      `json:"VirtualConsole.1.Enable"`
		VirtualConsole_1_EncryptEnable                                       string      `json:"VirtualConsole.1.EncryptEnable"`
		VirtualConsole_1_LocalDisable                                        string      `json:"VirtualConsole.1.LocalDisable"`
		VirtualConsole_1_LocalVideo                                          string      `json:"VirtualConsole.1.LocalVideo"`
		VirtualConsole_1_MaxSessions                                         int         `json:"VirtualConsole.1.MaxSessions"`
		VirtualConsole_1_PluginType                                          string      `json:"VirtualConsole.1.PluginType"`
		VirtualConsole_1_Port                                                int         `json:"VirtualConsole.1.Port"`
		VirtualConsole_1_Timeout                                             int         `json:"VirtualConsole.1.Timeout"`
		VirtualMedia_1_ActiveSessions                                        int         `json:"VirtualMedia.1.ActiveSessions"`
		VirtualMedia_1_Attached                                              string      `json:"VirtualMedia.1.Attached"`
		VirtualMedia_1_BootOnce                                              string      `json:"VirtualMedia.1.BootOnce"`
		VirtualMedia_1_Enable                                                string      `json:"VirtualMedia.1.Enable"`
		VirtualMedia_1_EncryptEnable                                         string      `json:"VirtualMedia.1.EncryptEnable"`
		VirtualMedia_1_FloppyEmulation                                       string      `json:"VirtualMedia.1.FloppyEmulation"`
		VirtualMedia_1_KeyEnable                                             string      `json:"VirtualMedia.1.KeyEnable"`
		VirtualMedia_1_MaxSessions                                           int         `json:"VirtualMedia.1.MaxSessions"`
		WebServer_1_Enable                                                   string      `json:"WebServer.1.Enable"`
		WebServer_1_HttpPort                                                 int         `json:"WebServer.1.HttpPort"`
		WebServer_1_HttpsPort                                                int         `json:"WebServer.1.HttpsPort"`
		WebServer_1_HttpsRedirection                                         string      `json:"WebServer.1.HttpsRedirection"`
		WebServer_1_LowerEncryptionBitLength                                 string      `json:"WebServer.1.LowerEncryptionBitLength"`
		WebServer_1_MaxNumberOfSessions                                      int         `json:"WebServer.1.MaxNumberOfSessions"`
		WebServer_1_SSLEncryptionBitLength                                   string      `json:"WebServer.1.SSLEncryptionBitLength"`
		WebServer_1_TLSProtocol                                              string      `json:"WebServer.1.TLSProtocol"`
		WebServer_1_Timeout                                                  int         `json:"WebServer.1.Timeout"`
		WebServer_1_TitleBarOption                                           string      `json:"WebServer.1.TitleBarOption"`
		WebServer_1_TitleBarOptionCustom                                     string      `json:"WebServer.1.TitleBarOptionCustom"`
		VFlashSD_1_Enable                                                    string      `json:"vFlashSD.1.Enable"`
		VFlashSD_1_Presence                                                  string      `json:"vFlashSD.1.Presence"`
	} `json:"Attributes"`
	Description string `json:"Description"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
}

type LifeCycleData struct {
	AutoBackup                          string      `json:"autobackup"`
	AutoDiscovery                       string      `json:"autodiscovery"`
	BIOSRTDRequested                    string      `json:"bios_rtd"`
	CollectSystemInventoryOnRestart     string      `json:"collect_system_inventory_on_restart"`
	DiscoveryFactoryDefaults            string      `json:"discovery_factory_defaults"`
	IPAddress                           string      `json:"ipaddress"`
	IPChangeNotifyPS                    string      `json:"ip_change_notify"`
	IgnoreCertWarning                   string      `json:"ignore_cert_warning"`
	Licensed                            string      `json:"licensed"`
	LifecycleControllerState            string      `json:"lifecycle_controller_state"`
	PartConfigurationUpdate             string      `json:"part_config_update"`
	PartFirmwareUpdate                  string      `json:"part_firmware_update"`
	ProvisioningServer                  string      `json:"provisioning_server"`
	StorageHealthRollupStatus           int         `json:"storage_health_rollup_status"`
	SystemID                            string      `json:"systemid"`
	UserProxyPassword                   interface{} `json:"user_proxy_password"`
	UserProxyPort                       string      `json:"user_proxy_port"`
	UserProxyServer                     string      `json:"user_proxy_server"`
	UserProxyType                       string      `json:"user_proxy_type"`
	UserProxyUserName                   string      `json:"user_proxy_username"`
	VirtualAddressManagementApplication string      `json:"virtual_addr_mgmt_application"`
}

type LifeCycleAttrDell struct {
	_Redfish_Settings struct {
		_odata_context string `json:"@odata.context"`
		_odata_type    string `json:"@odata.type"`
		SettingsObject struct {
			_odata_id string `json:"@odata.id"`
		} `json:"SettingsObject"`
	} `json:"@Redfish.Settings"`
	_odata_context    string `json:"@odata.context"`
	_odata_id         string `json:"@odata.id"`
	_odata_type       string `json:"@odata.type"`
	AttributeRegistry string `json:"AttributeRegistry"`
	Attributes        struct {
		LCAttributes_1_AutoBackup                          string      `json:"LCAttributes.1.AutoBackup"`
		LCAttributes_1_AutoDiscovery                       string      `json:"LCAttributes.1.AutoDiscovery"`
		LCAttributes_1_AutoUpdate                          string      `json:"LCAttributes.1.AutoUpdate"`
		LCAttributes_1_BIOSRTDRequested                    string      `json:"LCAttributes.1.BIOSRTDRequested"`
		LCAttributes_1_CollectSystemInventoryOnRestart     string      `json:"LCAttributes.1.CollectSystemInventoryOnRestart"`
		LCAttributes_1_DiscoveryFactoryDefaults            string      `json:"LCAttributes.1.DiscoveryFactoryDefaults"`
		LCAttributes_1_IPAddress                           string      `json:"LCAttributes.1.IPAddress"`
		LCAttributes_1_IPChangeNotifyPS                    string      `json:"LCAttributes.1.IPChangeNotifyPS"`
		LCAttributes_1_IgnoreCertWarning                   string      `json:"LCAttributes.1.IgnoreCertWarning"`
		LCAttributes_1_Licensed                            string      `json:"LCAttributes.1.Licensed"`
		LCAttributes_1_LifecycleControllerState            string      `json:"LCAttributes.1.LifecycleControllerState"`
		LCAttributes_1_PartConfigurationUpdate             string      `json:"LCAttributes.1.PartConfigurationUpdate"`
		LCAttributes_1_PartFirmwareUpdate                  string      `json:"LCAttributes.1.PartFirmwareUpdate"`
		LCAttributes_1_ProvisioningServer                  string      `json:"LCAttributes.1.ProvisioningServer"`
		LCAttributes_1_StorageHealthRollupStatus           int         `json:"LCAttributes.1.StorageHealthRollupStatus"`
		LCAttributes_1_SystemID                            string      `json:"LCAttributes.1.SystemID"`
		LCAttributes_1_UserProxyPassword                   interface{} `json:"LCAttributes.1.UserProxyPassword"`
		LCAttributes_1_UserProxyPort                       string      `json:"LCAttributes.1.UserProxyPort"`
		LCAttributes_1_UserProxyServer                     string      `json:"LCAttributes.1.UserProxyServer"`
		LCAttributes_1_UserProxyType                       string      `json:"LCAttributes.1.UserProxyType"`
		LCAttributes_1_UserProxyUserName                   string      `json:"LCAttributes.1.UserProxyUserName"`
		LCAttributes_1_VirtualAddressManagementApplication string      `json:"LCAttributes.1.VirtualAddressManagementApplication"`
	} `json:"Attributes"`
	Description string `json:"Description"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
}

type BiosData struct {
	BootMode          string `json:"bootmode"`
	BootSeqRetry      string `json:"boot_sequence_retry"`
	InternalUsb       string `json:"internal_usb"`
	SriovGlobalEnable string `json:"sriov_enable"`
	SysProfile        string `json:"sys_profile"`
	AcPwrRcvry        string `json:"pwr_rcvry"`
	AcPwrRcvryDelay   string `json:"pwr_rcvry_delay"`
	Serial            string `json:"serial"`
}

type BiosDell struct {
	_Redfish_Settings struct {
		_odata_context string `json:"@odata.context"`
		_odata_type    string `json:"@odata.type"`
		SettingsObject struct {
			_odata_id string `json:"@odata.id"`
		} `json:"SettingsObject"`
	} `json:"@Redfish.Settings"`
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Actions        struct {
		Bios_ChangePassword struct {
			Target string `json:"target"`
		} `json:"#Bios.ChangePassword"`
		Bios_ResetBios struct {
			Target string `json:"target"`
		} `json:"#Bios.ResetBios"`
	} `json:"Actions"`
	AttributeRegistry string `json:"AttributeRegistry"`
	Attributes        struct {
		AcPwrRcvry                   string `json:"AcPwrRcvry"`
		AcPwrRcvryDelay              string `json:"AcPwrRcvryDelay"`
		AcPwrRcvryUserDelay          int    `json:"AcPwrRcvryUserDelay"`
		AesNi                        string `json:"AesNi"`
		AssetTag                     string `json:"AssetTag"`
		AuthorizeDeviceFirmware      string `json:"AuthorizeDeviceFirmware"`
		BootMode                     string `json:"BootMode"`
		BootSeqRetry                 string `json:"BootSeqRetry"`
		ConTermType                  string `json:"ConTermType"`
		ControlledTurbo              string `json:"ControlledTurbo"`
		CorrEccSmi                   string `json:"CorrEccSmi"`
		CPUInterconnectBusLinkPower  string `json:"CpuInterconnectBusLinkPower"`
		CPUInterconnectBusSpeed      string `json:"CpuInterconnectBusSpeed"`
		CurrentEmbVideoState         string `json:"CurrentEmbVideoState"`
		CurrentMemOpModeState        string `json:"CurrentMemOpModeState"`
		DcuIPPrefetcher              string `json:"DcuIpPrefetcher"`
		DcuStreamerPrefetcher        string `json:"DcuStreamerPrefetcher"`
		DellAutoDiscovery            string `json:"DellAutoDiscovery"`
		DellWyseP25BIOSAccess        string `json:"DellWyseP25BIOSAccess"`
		DynamicCoreAllocation        string `json:"DynamicCoreAllocation"`
		EmbSata                      string `json:"EmbSata"`
		EmbVideo                     string `json:"EmbVideo"`
		EnergyPerformanceBias        string `json:"EnergyPerformanceBias"`
		ErrPrompt                    string `json:"ErrPrompt"`
		ExtSerialConnector           string `json:"ExtSerialConnector"`
		FailSafeBaud                 string `json:"FailSafeBaud"`
		ForceInt10                   string `json:"ForceInt10"`
		HddFailover                  string `json:"HddFailover"`
		InBandManageabilityInterface string `json:"InBandManageabilityInterface"`
		IntegratedNetwork1           string `json:"IntegratedNetwork1"`
		IntegratedRaid               string `json:"IntegratedRaid"`
		IntelTxt                     string `json:"IntelTxt"`
		InternalUsb                  string `json:"InternalUsb"`
		IoatEngine                   string `json:"IoatEngine"`
		LogicalProc                  string `json:"LogicalProc"`
		MemFrequency                 string `json:"MemFrequency"`
		MemOpMode                    string `json:"MemOpMode"`
		MemPatrolScrub               string `json:"MemPatrolScrub"`
		MemRefreshRate               string `json:"MemRefreshRate"`
		MemTest                      string `json:"MemTest"`
		MemoryMappedIOH              string `json:"MemoryMappedIOH"`
		MmioAbove4Gb                 string `json:"MmioAbove4Gb"`
		MonitorMwait                 string `json:"MonitorMwait"`
		NodeInterleave               string `json:"NodeInterleave"`
		NumLock                      string `json:"NumLock"`
		NvmeMode                     string `json:"NvmeMode"`
		OneTimeBootMode              string `json:"OneTimeBootMode"`
		OneTimeBootSeqDev            string `json:"OneTimeBootSeqDev"`
		OneTimeHddSeqDev             string `json:"OneTimeHddSeqDev"`
		OppSrefEn                    string `json:"OppSrefEn"`
		OsWatchdogTimer              string `json:"OsWatchdogTimer"`
		PasswordStatus               string `json:"PasswordStatus"`
		PcieAspmL1                   string `json:"PcieAspmL1"`
		PowerCycleRequest            string `json:"PowerCycleRequest"`
		Proc1Brand                   string `json:"Proc1Brand"`
		Proc1Id                      string `json:"Proc1Id"`
		Proc1L2Cache                 string `json:"Proc1L2Cache"`
		Proc1L3Cache                 string `json:"Proc1L3Cache"`
		Proc1NumCores                int    `json:"Proc1NumCores"`
		Proc1TurboCoreNum            string `json:"Proc1TurboCoreNum"`
		Proc2Brand                   string `json:"Proc2Brand"`
		Proc2Id                      string `json:"Proc2Id"`
		Proc2L2Cache                 string `json:"Proc2L2Cache"`
		Proc2L3Cache                 string `json:"Proc2L3Cache"`
		Proc2NumCores                int    `json:"Proc2NumCores"`
		Proc2TurboCoreNum            string `json:"Proc2TurboCoreNum"`
		ProcAdjCacheLine             string `json:"ProcAdjCacheLine"`
		ProcBusSpeed                 string `json:"ProcBusSpeed"`
		ProcC1E                      string `json:"ProcC1E"`
		ProcCStates                  string `json:"ProcCStates"`
		ProcCoreSpeed                string `json:"ProcCoreSpeed"`
		ProcCores                    string `json:"ProcCores"`
		ProcHwPrefetcher             string `json:"ProcHwPrefetcher"`
		ProcPwrPerf                  string `json:"ProcPwrPerf"`
		ProcTurboMode                string `json:"ProcTurboMode"`
		ProcVirtualization           string `json:"ProcVirtualization"`
		ProcX2Apic                   string `json:"ProcX2Apic"`
		PwrButton                    string `json:"PwrButton"`
		RedirAfterBoot               string `json:"RedirAfterBoot"`
		RedundantOsLocation          string `json:"RedundantOsLocation"`
		SHA256SetupPassword          string `json:"SHA256SetupPassword"`
		SHA256SetupPasswordSalt      string `json:"SHA256SetupPasswordSalt"`
		SHA256SystemPassword         string `json:"SHA256SystemPassword"`
		SHA256SystemPasswordSalt     string `json:"SHA256SystemPasswordSalt"`
		SataPortA                    string `json:"SataPortA"`
		SataPortACapacity            string `json:"SataPortACapacity"`
		SataPortADriveType           string `json:"SataPortADriveType"`
		SataPortAModel               string `json:"SataPortAModel"`
		SataPortB                    string `json:"SataPortB"`
		SataPortBCapacity            string `json:"SataPortBCapacity"`
		SataPortBDriveType           string `json:"SataPortBDriveType"`
		SataPortBModel               string `json:"SataPortBModel"`
		SataPortC                    string `json:"SataPortC"`
		SataPortCCapacity            string `json:"SataPortCCapacity"`
		SataPortCDriveType           string `json:"SataPortCDriveType"`
		SataPortCModel               string `json:"SataPortCModel"`
		SataPortD                    string `json:"SataPortD"`
		SataPortDCapacity            string `json:"SataPortDCapacity"`
		SataPortDDriveType           string `json:"SataPortDDriveType"`
		SataPortDModel               string `json:"SataPortDModel"`
		SataPortE                    string `json:"SataPortE"`
		SataPortECapacity            string `json:"SataPortECapacity"`
		SataPortEDriveType           string `json:"SataPortEDriveType"`
		SataPortEModel               string `json:"SataPortEModel"`
		SataPortF                    string `json:"SataPortF"`
		SataPortFCapacity            string `json:"SataPortFCapacity"`
		SataPortFDriveType           string `json:"SataPortFDriveType"`
		SataPortFModel               string `json:"SataPortFModel"`
		SataPortG                    string `json:"SataPortG"`
		SataPortGCapacity            string `json:"SataPortGCapacity"`
		SataPortGDriveType           string `json:"SataPortGDriveType"`
		SataPortGModel               string `json:"SataPortGModel"`
		SataPortH                    string `json:"SataPortH"`
		SataPortHCapacity            string `json:"SataPortHCapacity"`
		SataPortHDriveType           string `json:"SataPortHDriveType"`
		SataPortHModel               string `json:"SataPortHModel"`
		SataPortI                    string `json:"SataPortI"`
		SataPortICapacity            string `json:"SataPortICapacity"`
		SataPortIDriveType           string `json:"SataPortIDriveType"`
		SataPortIModel               string `json:"SataPortIModel"`
		SataPortJ                    string `json:"SataPortJ"`
		SataPortJCapacity            string `json:"SataPortJCapacity"`
		SataPortJDriveType           string `json:"SataPortJDriveType"`
		SataPortJModel               string `json:"SataPortJModel"`
		SataPortK                    string `json:"SataPortK"`
		SataPortKCapacity            string `json:"SataPortKCapacity"`
		SataPortKDriveType           string `json:"SataPortKDriveType"`
		SataPortKModel               string `json:"SataPortKModel"`
		SataPortL                    string `json:"SataPortL"`
		SataPortLCapacity            string `json:"SataPortLCapacity"`
		SataPortLDriveType           string `json:"SataPortLDriveType"`
		SataPortLModel               string `json:"SataPortLModel"`
		SataPortM                    string `json:"SataPortM"`
		SataPortMCapacity            string `json:"SataPortMCapacity"`
		SataPortMDriveType           string `json:"SataPortMDriveType"`
		SataPortMModel               string `json:"SataPortMModel"`
		SataPortN                    string `json:"SataPortN"`
		SataPortNCapacity            string `json:"SataPortNCapacity"`
		SataPortNDriveType           string `json:"SataPortNDriveType"`
		SataPortNModel               string `json:"SataPortNModel"`
		SecureBoot                   string `json:"SecureBoot"`
		SecureBootMode               string `json:"SecureBootMode"`
		SecureBootPolicy             string `json:"SecureBootPolicy"`
		SecurityFreezeLock           string `json:"SecurityFreezeLock"`
		SerialComm                   string `json:"SerialComm"`
		SerialPortAddress            string `json:"SerialPortAddress"`
		SetBootOrderFqdd1            string `json:"SetBootOrderFqdd1"`
		SetBootOrderFqdd10           string `json:"SetBootOrderFqdd10"`
		SetBootOrderFqdd11           string `json:"SetBootOrderFqdd11"`
		SetBootOrderFqdd12           string `json:"SetBootOrderFqdd12"`
		SetBootOrderFqdd13           string `json:"SetBootOrderFqdd13"`
		SetBootOrderFqdd14           string `json:"SetBootOrderFqdd14"`
		SetBootOrderFqdd15           string `json:"SetBootOrderFqdd15"`
		SetBootOrderFqdd16           string `json:"SetBootOrderFqdd16"`
		SetBootOrderFqdd2            string `json:"SetBootOrderFqdd2"`
		SetBootOrderFqdd3            string `json:"SetBootOrderFqdd3"`
		SetBootOrderFqdd4            string `json:"SetBootOrderFqdd4"`
		SetBootOrderFqdd5            string `json:"SetBootOrderFqdd5"`
		SetBootOrderFqdd6            string `json:"SetBootOrderFqdd6"`
		SetBootOrderFqdd7            string `json:"SetBootOrderFqdd7"`
		SetBootOrderFqdd8            string `json:"SetBootOrderFqdd8"`
		SetBootOrderFqdd9            string `json:"SetBootOrderFqdd9"`
		SetLegacyHddOrderFqdd1       string `json:"SetLegacyHddOrderFqdd1"`
		SetLegacyHddOrderFqdd10      string `json:"SetLegacyHddOrderFqdd10"`
		SetLegacyHddOrderFqdd11      string `json:"SetLegacyHddOrderFqdd11"`
		SetLegacyHddOrderFqdd12      string `json:"SetLegacyHddOrderFqdd12"`
		SetLegacyHddOrderFqdd13      string `json:"SetLegacyHddOrderFqdd13"`
		SetLegacyHddOrderFqdd14      string `json:"SetLegacyHddOrderFqdd14"`
		SetLegacyHddOrderFqdd15      string `json:"SetLegacyHddOrderFqdd15"`
		SetLegacyHddOrderFqdd16      string `json:"SetLegacyHddOrderFqdd16"`
		SetLegacyHddOrderFqdd2       string `json:"SetLegacyHddOrderFqdd2"`
		SetLegacyHddOrderFqdd3       string `json:"SetLegacyHddOrderFqdd3"`
		SetLegacyHddOrderFqdd4       string `json:"SetLegacyHddOrderFqdd4"`
		SetLegacyHddOrderFqdd5       string `json:"SetLegacyHddOrderFqdd5"`
		SetLegacyHddOrderFqdd6       string `json:"SetLegacyHddOrderFqdd6"`
		SetLegacyHddOrderFqdd7       string `json:"SetLegacyHddOrderFqdd7"`
		SetLegacyHddOrderFqdd8       string `json:"SetLegacyHddOrderFqdd8"`
		SetLegacyHddOrderFqdd9       string `json:"SetLegacyHddOrderFqdd9"`
		SetupPassword                string `json:"SetupPassword"`
		Slot1                        string `json:"Slot1"`
		Slot1Bif                     string `json:"Slot1Bif"`
		Slot2                        string `json:"Slot2"`
		Slot2Bif                     string `json:"Slot2Bif"`
		Slot3                        string `json:"Slot3"`
		Slot3Bif                     string `json:"Slot3Bif"`
		SriovGlobalEnable            string `json:"SriovGlobalEnable"`
		SubNumaCluster               string `json:"SubNumaCluster"`
		SysMemSize                   string `json:"SysMemSize"`
		SysMemSpeed                  string `json:"SysMemSpeed"`
		SysMemType                   string `json:"SysMemType"`
		SysMemVolt                   string `json:"SysMemVolt"`
		SysMfrContactInfo            string `json:"SysMfrContactInfo"`
		SysPassword                  string `json:"SysPassword"`
		SysProfile                   string `json:"SysProfile"`
		SystemBiosVersion            string `json:"SystemBiosVersion"`
		SystemCpldVersion            string `json:"SystemCpldVersion"`
		SystemManufacturer           string `json:"SystemManufacturer"`
		SystemMeVersion              string `json:"SystemMeVersion"`
		SystemModelName              string `json:"SystemModelName"`
		SystemServiceTag             string `json:"SystemServiceTag"`
		TpmFirmware                  string `json:"TpmFirmware"`
		TpmInfo                      string `json:"TpmInfo"`
		TpmPpiBypassClear            string `json:"TpmPpiBypassClear"`
		TpmPpiBypassProvision        string `json:"TpmPpiBypassProvision"`
		TpmSecurity                  string `json:"TpmSecurity"`
		UefiComplianceVersion        string `json:"UefiComplianceVersion"`
		UefiVariableAccess           string `json:"UefiVariableAccess"`
		UncoreFrequency              string `json:"UncoreFrequency"`
		UpiPrefetch                  string `json:"UpiPrefetch"`
		UsbManagedPort               string `json:"UsbManagedPort"`
		UsbPorts                     string `json:"UsbPorts"`
		VideoMem                     string `json:"VideoMem"`
		WorkloadProfile              string `json:"WorkloadProfile"`
		WriteCache                   string `json:"WriteCache"`
		WriteDataCrc                 string `json:"WriteDataCrc"`
	} `json:"Attributes"`
	Description string `json:"Description"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
}

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
			OdataId string `json:"@odata.id"`
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
type BootOrder struct {
	_Redfish_Settings struct {
		_odata_context string `json:"@odata.context"`
		_odata_type    string `json:"@odata.type"`
		SettingsObject struct {
			_odata_id string `json:"@odata.id"`
		} `json:"SettingsObject"`
	} `json:"@Redfish.Settings"`
	_odata_context    string `json:"@odata.context"`
	_odata_id         string `json:"@odata.id"`
	_odata_type       string `json:"@odata.type"`
	AttributeRegistry string `json:"AttributeRegistry"`
	Attributes        struct {
		BootSeq []struct {
			Enabled bool   `json:"Enabled"`
			ID      string `json:"Id"`
			Index   int    `json:"Index"`
			Name    string `json:"Name"`
		} `json:"BootSeq"`
		HddSeq []struct {
			ID    string `json:"Id"`
			Index int    `json:"Index"`
			Name  string `json:"Name"`
		} `json:"HddSeq"`
	} `json:"Attributes"`
	Description string `json:"Description"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
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
	MacAddress  string `json:"macaddress"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	State       string `json:"state"`
	Vlan        string `json:"vlan"`
}
type FirmwareData struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Version    string `json:"version"`
	Updateable bool   `json:"updateable"`
}
type BootOrderData struct {
	Enabled bool   `json:"enabled"`
	Index   int    `json:"index"`
	Name    string `json:"name"`
}
type ProcessorData struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Description    string `json:"Description"`
	ID             string `json:"Id"`
	InstructionSet []struct {
		Member string `json:"Member"`
	} `json:"InstructionSet"`
	Manufacturer          string `json:"Manufacturer"`
	MaxSpeedMHz           int    `json:"MaxSpeedMHz"`
	Model                 string `json:"Model"`
	Name                  string `json:"Name"`
	ProcessorArchitecture []struct {
		Member string `json:"Member"`
	} `json:"ProcessorArchitecture"`
	ProcessorID struct {
		EffectiveFamily         string `json:"EffectiveFamily"`
		EffectiveModel          string `json:"EffectiveModel"`
		IdentificationRegisters string `json:"IdentificationRegisters"`
		MicrocodeInfo           string `json:"MicrocodeInfo"`
		Step                    string `json:"Step"`
		VendorID                string `json:"VendorId"`
	} `json:"ProcessorId"`
	ProcessorType string `json:"ProcessorType"`
	Socket        string `json:"Socket"`
	Status        struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	TotalCores   int `json:"TotalCores"`
	TotalThreads int `json:"TotalThreads"`
}
type ProcessorsListData struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Description    string `json:"Description"`
	Members        []struct {
		OdataId string `json:"@odata.id"`
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
}
type PowerData struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Description    string `json:"Description"`
	ID             string `json:"Id"`
	Name           string `json:"Name"`
	PowerControl   []struct {
		_odata_id           string `json:"@odata.id"`
		MemberID            string `json:"MemberId"`
		Name                string `json:"Name"`
		PowerAllocatedWatts int    `json:"PowerAllocatedWatts"`
		PowerAvailableWatts int    `json:"PowerAvailableWatts"`
		PowerCapacityWatts  int    `json:"PowerCapacityWatts"`
		PowerConsumedWatts  int    `json:"PowerConsumedWatts"`
		PowerLimit          struct {
			CorrectionInMs int    `json:"CorrectionInMs"`
			LimitException string `json:"LimitException"`
			LimitInWatts   int    `json:"LimitInWatts"`
		} `json:"PowerLimit"`
		PowerMetrics struct {
			AverageConsumedWatts int `json:"AverageConsumedWatts"`
			IntervalInMin        int `json:"IntervalInMin"`
			MaxConsumedWatts     int `json:"MaxConsumedWatts"`
			MinConsumedWatts     int `json:"MinConsumedWatts"`
		} `json:"PowerMetrics"`
		PowerRequestedWatts int `json:"PowerRequestedWatts"`
		RelatedItem         []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"RelatedItem"`
		RelatedItem_odata_count int `json:"RelatedItem@odata.count"`
	} `json:"PowerControl"`
	PowerControlcount int `json:"PowerControl@odata.count"`
	PowerSupplies     []struct {
		_odata_id       string `json:"@odata.id"`
		FirmwareVersion string `json:"FirmwareVersion"`
		InputRanges     []struct {
			InputType          string `json:"InputType"`
			MaximumFrequencyHz int    `json:"MaximumFrequencyHz"`
			MaximumVoltage     int    `json:"MaximumVoltage"`
			MinimumFrequencyHz int    `json:"MinimumFrequencyHz"`
			MinimumVoltage     int    `json:"MinimumVoltage"`
			OutputWattage      int    `json:"OutputWattage"`
		} `json:"InputRanges"`
		InputRanges_odata_count int    `json:"InputRanges@odata.count"`
		LastPowerOutputWatts    int    `json:"LastPowerOutputWatts"`
		LineInputVoltage        int    `json:"LineInputVoltage"`
		LineInputVoltageType    string `json:"LineInputVoltageType"`
		Manufacturer            string `json:"Manufacturer"`
		MemberID                string `json:"MemberId"`
		Model                   string `json:"Model"`
		Name                    string `json:"Name"`
		PartNumber              string `json:"PartNumber"`
		PowerCapacityWatts      int    `json:"PowerCapacityWatts"`
		PowerSupplyType         string `json:"PowerSupplyType"`
		Redundancy              []struct {
			_odata_id       string `json:"@odata.id"`
			MaxNumSupported int    `json:"MaxNumSupported"`
			MemberID        string `json:"MemberId"`
			MinNumNeeded    int    `json:"MinNumNeeded"`
			Mode            []struct {
				Member string `json:"Member"`
			} `json:"Mode"`
			Name          string `json:"Name"`
			RedundancySet []struct {
				_odata_id string `json:"@odata.id"`
			} `json:"RedundancySet"`
			RedundancySet_odata_count int `json:"RedundancySet@odata.count"`
			Status                    struct {
				Health string `json:"Health"`
				State  string `json:"State"`
			} `json:"Status"`
		} `json:"Redundancy"`
		Redundancycount int `json:"Redundancy@odata.count"`
		RelatedItem     []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"RelatedItem"`
		RelatedItem_odata_count int    `json:"RelatedItem@odata.count"`
		SerialNumber            string `json:"SerialNumber"`
		SparePartNumber         string `json:"SparePartNumber"`
		Status                  struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
	} `json:"PowerSupplies"`
	PowerSuppliescount int `json:"PowerSupplies@odata.count"`
	Redundancy         []struct {
		_odata_id       string `json:"@odata.id"`
		MaxNumSupported int    `json:"MaxNumSupported"`
		MemberID        string `json:"MemberId"`
		MinNumNeeded    int    `json:"MinNumNeeded"`
		Mode            []struct {
			Member string `json:"Member"`
		} `json:"Mode"`
		Name          string `json:"Name"`
		RedundancySet []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"RedundancySet"`
		RedundancySet_odata_count int `json:"RedundancySet@odata.count"`
		Status                    struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
	} `json:"Redundancy"`
	Redundancycount int `json:"Redundancy@odata.count"`
	Voltages        []struct {
		_odata_id                 string      `json:"@odata.id"`
		LowerThresholdCritical    interface{} `json:"LowerThresholdCritical"`
		LowerThresholdFatal       interface{} `json:"LowerThresholdFatal"`
		LowerThresholdNonCritical interface{} `json:"LowerThresholdNonCritical"`
		MaxReadingRange           int         `json:"MaxReadingRange"`
		MemberID                  string      `json:"MemberId"`
		MinReadingRange           int         `json:"MinReadingRange"`
		Name                      string      `json:"Name"`
		PhysicalContext           string      `json:"PhysicalContext"`
		ReadingVolts              int         `json:"ReadingVolts"`
		RelatedItem               []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"RelatedItem"`
		RelatedItem_odata_count int `json:"RelatedItem@odata.count"`
		SensorNumber            int `json:"SensorNumber"`
		Status                  struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		UpperThresholdCritical    interface{} `json:"UpperThresholdCritical"`
		UpperThresholdFatal       interface{} `json:"UpperThresholdFatal"`
		UpperThresholdNonCritical interface{} `json:"UpperThresholdNonCritical"`
	} `json:"Voltages"`
	Voltagescount int `json:"Voltages@odata.count"`
}

type HealthList struct {
	Name   string `json:"name"`
	Health string `json:"health"`
	State  string `json:"state"`
}
type StorageHealthList struct {
	Name   string `json:"name"`
	Health string `json:"health"`
	State  string `json:"state"`
	Space  int    `json:"space"`
}
type StorageCollection struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Description    string `json:"Description"`
	Members        []struct {
		OdataId string `json:"@odata.id"`
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
}
type StorageDetails struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Description    string `json:"Description"`
	Devices        []struct {
		CapacityBytes int    `json:"CapacityBytes"`
		Manufacturer  string `json:"Manufacturer"`
		Model         string `json:"Model"`
		Name          string `json:"Name"`
		Status        struct {
			Health       string `json:"Health"`
			HealthRollup string `json:"HealthRollup"`
			State        string `json:"State"`
		} `json:"Status"`
	} `json:"Devices"`
	Devicescount int    `json:"Devices@odata.count"`
	ID           string `json:"Id"`
	Name         string `json:"Name"`
	Status       struct {
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
		State        string `json:"State"`
	} `json:"Status"`
	UefiDevicePath string `json:"UefiDevicePath"`
}

type ThermalHealthList struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Description    string `json:"Description"`
	Fans           []struct {
		_odata_id                 string        `json:"@odata.id"`
		FanName                   string        `json:"FanName"`
		LowerThresholdCritical    int           `json:"LowerThresholdCritical"`
		LowerThresholdFatal       int           `json:"LowerThresholdFatal"`
		LowerThresholdNonCritical int           `json:"LowerThresholdNonCritical"`
		MaxReadingRange           int           `json:"MaxReadingRange"`
		MemberID                  string        `json:"MemberId"`
		MinReadingRange           int           `json:"MinReadingRange"`
		Name                      string        `json:"Name"`
		PhysicalContext           string        `json:"PhysicalContext"`
		Reading                   int           `json:"Reading"`
		ReadingUnits              string        `json:"ReadingUnits"`
		Redundancy                []interface{} `json:"Redundancy"`
		Redundancy_odata_count    int           `json:"Redundancy@odata.count"`
		RelatedItem               []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"RelatedItem"`
		RelatedItem_odata_count int `json:"RelatedItem@odata.count"`
		Status                  struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		UpperThresholdCritical    interface{} `json:"UpperThresholdCritical"`
		UpperThresholdFatal       interface{} `json:"UpperThresholdFatal"`
		UpperThresholdNonCritical interface{} `json:"UpperThresholdNonCritical"`
	} `json:"Fans"`
	Fanscount  int    `json:"Fans@odata.count"`
	ID         string `json:"Id"`
	Name       string `json:"Name"`
	Redundancy []struct {
		_odata_id       string      `json:"@odata.id"`
		MaxNumSupported interface{} `json:"MaxNumSupported"`
		MemberID        string      `json:"MemberId"`
		MinNumNeeded    interface{} `json:"MinNumNeeded"`
		Mode            []struct {
			Member string `json:"Member"`
		} `json:"Mode"`
		Name                      string        `json:"Name"`
		RedundancySet             []interface{} `json:"RedundancySet"`
		RedundancySet_odata_count int           `json:"RedundancySet@odata.count"`
		Status                    struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
	} `json:"Redundancy"`
	Redundancycount int `json:"Redundancy@odata.count"`
	Temperatures    []struct {
		_odata_id                 string      `json:"@odata.id"`
		LowerThresholdCritical    int         `json:"LowerThresholdCritical"`
		LowerThresholdFatal       int         `json:"LowerThresholdFatal"`
		LowerThresholdNonCritical interface{} `json:"LowerThresholdNonCritical"`
		MaxReadingRangeTemp       int         `json:"MaxReadingRangeTemp"`
		MemberID                  string      `json:"MemberId"`
		MinReadingRangeTemp       int         `json:"MinReadingRangeTemp"`
		Name                      string      `json:"Name"`
		PhysicalContext           string      `json:"PhysicalContext"`
		ReadingCelsius            int         `json:"ReadingCelsius"`
		RelatedItem               []struct {
			_odata_id string `json:"@odata.id"`
		} `json:"RelatedItem"`
		RelatedItem_odata_count int `json:"RelatedItem@odata.count"`
		SensorNumber            int `json:"SensorNumber"`
		Status                  struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		UpperThresholdCritical    int         `json:"UpperThresholdCritical"`
		UpperThresholdFatal       int         `json:"UpperThresholdFatal"`
		UpperThresholdNonCritical interface{} `json:"UpperThresholdNonCritical"`
	} `json:"Temperatures"`
	Temperaturesount int `json:"Temperatures@odata.count"`
}

// ResetType@Redfish.AllowableValues
// 0	"On"
// 1	"ForceOff"
// 2	"GracefulRestart"
// 3	"GracefulShutdown"
// 4	"PushPowerButton"
// 5	"Nmi"
// target: "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"
func (c *IloClient) StartServer() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var jsonStr = []byte(`{"ResetType": "On"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))
	req.Header.Set("Content-Type", "application/json")

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

	return "Server Started", nil
}

func (c *IloClient) StopServer() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var jsonStr = []byte(`{"ResetType": "ForceOff"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))
	req.Header.Set("Content-Type", "application/json")

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

	return "Server Stopped", nil
}

func (c *IloClient) GetServerPowerState() (string, error) {
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
			return "nil", err
		} else {
			return "nil", err
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

	return data.PowerState, nil

}

// func (c *IloClient) SetBiosAttributes() (string, error) {}

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

func (c *IloClient) GetProcessorHealthDell() (string, error) {
	///redfish/v1/Systems/System.Embedded.1/Processors

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Processors"

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
		x             ProcessorsListData
		processHealth []HealthList
	)

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

		var y ProcessorData

		json.Unmarshal(_body, &y)

		procHealth := HealthList{
			Name:   y.ID,
			Health: y.Status.Health,
			State:  y.Status.State,
		}
		processHealth = append(processHealth, procHealth)
	}

	output, _ := json.Marshal(processHealth)
	return string(output), nil

}

// func (c *IloClient) GetMemoryHealthDell() (string, error) {}

func (c *IloClient) GetPowerHealthDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Chassis/System.Embedded.1/Power"

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
		x             PowerData
		powerSupplies []HealthList
	)

	json.Unmarshal(_body, &x)

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

	output, _ := json.Marshal(powerSupplies)
	return string(output), nil
}

func (c *IloClient) GetSensorsHealthDell() (string, error) {

	url := c.Hostname + "/redfish/v1/Chassis/System.Embedded.1/Thermal"

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
		x             ThermalHealthList
		thermalHealth []HealthList
	)

	json.Unmarshal(_body, &x)

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

	output, _ := json.Marshal(thermalHealth)
	return string(output), nil

}

func (c *IloClient) GetStorageHealthDell() (string, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Storage/Controllers"

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
		x           StorageCollection
		_healthdata []StorageHealthList
	)

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

		var y StorageDetails

		json.Unmarshal(_body, &y)

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
	output, _ := json.Marshal(_healthdata)
	return string(output), nil

}

// GetAggHealthDataDell will fetch the data related to all components health(aggregated view)
func (c *IloClient) GetAggHealthDataDell() (string, error) {

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
		_healthdata []HealthList
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

			healthData := HealthList{
				Name:   y.Name,
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

func (c *IloClient) GetBiosDataDell() (string, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Bios"

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

	var x BiosDell

	json.Unmarshal(_body, &x)

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

	output, _ := json.Marshal(_BiosData)

	return string(output), nil

}

func (c *IloClient) GetLifecycleAttrDell() (string, error) {

	url := c.Hostname + "/redfish/v1/Managers/LifecycleController.Embedded.1/Attributes"

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

	var x LifeCycleAttrDell

	json.Unmarshal(_body, &x)

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

	output, _ := json.Marshal(_LfcycleDat)

	return string(output), nil

}

func (c *IloClient) GetIDRACAttrDell() (string, error) {

	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Attributes"

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

	var x IDRACAttrDell

	json.Unmarshal(_body, &x)

	_data := x.Attributes

	_idracData := IDRACData{
		VirtualConsoleMaxSessions: _data.VirtualConsole_1_MaxSessions,
		VirtualConsolePluginType:  _data.VirtualConsole_1_PluginType,
		WebServerSSLEncryption:    _data.WebServer_1_SSLEncryptionBitLength,
		IPMILanEnable:             _data.IPMILan_1_Enable,
		DNSDomainName:             _data.NIC_1_DNSDomainName,
	}

	output, _ := json.Marshal(_idracData)

	return string(output), nil

}

func (c *IloClient) GetSysAttrDell() (string, error) {

	url := c.Hostname + "/redfish/v1/Managers/System.Embedded.1/Attributes"

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

	var x SysAttrDell

	json.Unmarshal(_body, &x)

	_data := x.Attributes

	_sysData := SysAttrData{
		ServerPwrPSRedPolicy: _data.ServerPwr_1_PSRedPolicy,
		ServerPwrPSRapidOn:   _data.ServerPwr_1_PSRapidOn,
	}

	output, _ := json.Marshal(_sysData)

	return string(output), nil

}

func (c *IloClient) GetBootOrderDell() (string, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/BootSources"

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

	var x BootOrder

	json.Unmarshal(_body, &x)

	var _bootOrder []BootOrderData

	for i := range x.Attributes.BootSeq {

		_result := BootOrderData{
			Enabled: x.Attributes.BootSeq[i].Enabled,
			Index:   x.Attributes.BootSeq[i].Index,
			Name:    x.Attributes.BootSeq[i].Name,
		}

		_bootOrder = append(_bootOrder, _result)
	}

	output, _ := json.Marshal(_bootOrder)

	return string(output), nil

}
