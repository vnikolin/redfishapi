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
	MacAddress  string `json:"macaddress"`
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

func (c *IloClient) GetBiosDataDell() (string, error) {

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
