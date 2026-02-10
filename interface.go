package redfishapi

// DellProvider defines all Dell iDRAC Redfish operations.
type DellProvider interface {
	StartServerDell() (string, error)
	StopServerDell() (string, error)
	GracefulRestartDell() (string, error)
	ResetSSLConfigDell() (string, error)
	GetServerPowerStateDell() (string, error)
	CheckLoginDell() (string, bool, error)
	ImportConfigDell(jsonData []byte) (string, error)
	CreateJobDell(jsonData []byte) (string, error)
	GetJobsStatusDell() ([]JobStatusDell, error)
	GetAllJobsDell() ([]Members, error)
	SetBiosSettingsDell(jsonData []byte) (string, error)
	ClearJobsDell() (string, error)
	SetAttributesDell(service string, jsonData []byte) (string, error)
	GetStorageRaidDell() ([]StorageRaidDetailsDell, error)
	GetNetworkSwitchInfoDell() ([]SwitchData, error)
	GetNetworkPortsDell() ([]MACData, error)
	GetMacAddressDell() (string, error)
	GetIdracLicenses() ([]LicenseData, error)
	GetMacAddressModelDell() ([]MACModelDell, error)
	GetProcessorHealthDell() ([]HealthList, error)
	GetPowerHealthDell() ([]HealthList, error)
	GetSensorsHealthDell() ([]HealthList, error)
	GetStorageDriveDetailsDell() ([]StorageDriveDetailsDell, error)
	GetStorageHealthDell() ([]StorageHealthList, error)
	GetAggHealthDataDell(model string) ([]HealthList, error)
	GetFirmwareDell() ([]FirmwareData, error)
	FirmwareUpdateDell() (string, error)
	FirmwareUploadDell(repoUrl string) (string, error)
	TaskStatusDell(taskUrl string) (ExportConfigStatus, error)
	GetBiosDataDell() (BiosAttributesData, error)
	GetLifecycleAttrDell() (LifeCycleData, error)
	ListUsersDell() ([]UserListDell, error)
	CreateUserDell(num int, username string, password string, role string, status bool) (string, error)
	DeleteUserDell(num int) (string, error)
	GetIDRACAttrDell() (IDRACAttributesData, error)
	GetSysAttrDell() (SysAttributesData, error)
	GetBootOrderDell() ([]BootOrderData, error)
	SetBootOrderDell(jsonData []byte) (string, error)
	GetSystemEventLogsDell(version string) ([]SystemEventLogRes, error)
	GetLifeCycleEventLogsDell(totalPages int) ([]LifeCycleEventLogRes, error)
	WriteLCLog(messageDesctiption string) (string, error)
	GetUserAccountsDell() ([]Accounts, error)
	GetSystemInfoDell() (SystemData, error)
	GetComponentAttr(comp string) (ExportConfigResponse, error)
	MountImageDell(image string) (string, error)
	UnMountImageDell() (string, error)
	GetRemoteImageStatusDell() (ImageStatusDell, error)
	ClearStorageControllerRaidDell(controllerID string) (string, error)
	GetJobStatusDell(jobID string) (JobStatusDell, error)
	ClearJobsDellForce() (string, error)
	FleaDrainDell() (string, error)
	PowerActionServerDell(powerAction string) (string, error)
	UpdateFirmwareDell(firmwareDir string, firmwareFile string) (string, error)
}

// HPProvider defines all HP iLO Redfish operations.
type HPProvider interface {
	StartServerHP() (string, error)
	StopServerHP() (string, error)
	GetSystemInfoHP() (SystemData, error)
	GetServerPowerStateHP() (string, error)
	CheckLoginHP() (string, error)
	GetFirmwareHP() ([]FirmwareData, error)
	GetThermalHealthHP() ([]HealthList, error)
	GetPowerHealthHP() ([]HealthList, error)
	GetInterfaceHealthHP() ([]HealthList, error)
	GetProcessorInfoHP() ([]ProcessorInfoHP, error)
	GetProcessorHealthHP() ([]HealthList, error)
	GetUserAccountsHP() ([]Accounts, error)
	GetSystemEventLogsHP() ([]SystemEventLogRes, error)
	GetBiosDataHP() (BiosDataHP, error)
	GetLicenseInfoHP() (LicenseInfo, error)
	GetPCISlotsHP() ([]PCISlotsInfo, error)
	GetEthernetInterfacesHP() ([]MACData, error)
}

// RedfishProvider composes both Dell and HP interfaces.
// The redfishProvider struct satisfies this interface.
type RedfishProvider interface {
	DellProvider
	HPProvider
}

// BMCProvider defines vendor-neutral BMC operations shared by both Dell and HP.
// Use NewDellBMC or NewHPBMC to adapt a vendor-specific provider to this interface.
type BMCProvider interface {
	GetSystemInfo() (SystemData, error)
	GetFirmware() ([]FirmwareData, error)
	StartServer() (string, error)
	StopServer() (string, error)
	GetServerPowerState() (string, error)
	GetProcessorHealth() ([]HealthList, error)
	GetPowerHealth() ([]HealthList, error)
	GetThermalHealth() ([]HealthList, error)
	GetUserAccounts() ([]Accounts, error)
}
