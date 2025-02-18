package redfishapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	ver "github.com/Masterminds/semver/v3"
)

// Declaring the Constant Values
const (
	StatusUnauthorized        = "Unauthorized"
	StatusInternalServerError = "Server Error"
	StatusBadRequest          = "Bad Request"
	StatusUnreachable         = "Unreachable"
	JobStateCompleted         = "Completed"
	JobStateFailed            = "Failed"
	JobStateRunning           = "Running"
	JobStateScheduled         = "Scheduled"
	TaskStateStarting         = "Starting"
	TaskStateRunning          = "Running"
	TaskStateCompleted        = "Completed"
	TaskStatusOK              = "OK"
	TaskStatusCritical        = "Critical"
)

type RedfishProvider interface {
	StartServerDell() (string, error)
	StopServerDell() (string, error)
	GracefulRestartDell() (string, error)
	ResetSSLConfig() (string, error)
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
	// DeleteUserDell(num int, role string, status bool) (string, error)
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

// ResetType@Redfish.AllowableValues
// "On"
// "ForceOff"
// "ForceRestart"
// "GracefulRestart"
// "GracefulShutdown"
// "PushPowerButton"
// "Nmi"
// "PowerCycle"
// target: "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"
// works: R730xd,R740xd
func (c *redfishProvider) PowerActionServerDell(powerAction string) (string, error) {

	allowableActions := []string{"On", "ForceOff", "ForceRestart", "GracefulRestart", "GracefulShutdown", "PushPowerButton", "Nmi", "PowerCycle"}
	// check if the action is valid
	if !slices.Contains(allowableActions, powerAction) {
		return "", fmt.Errorf("invalid power action: %s", powerAction)
	}
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"

	var jsonStr = []byte(`{"ResetType": "` + powerAction + `"}`)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server " + powerAction, nil
}

// StartServerDell ...
func (c *redfishProvider) StartServerDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"

	var jsonStr = []byte(`{"ResetType": "On"}`)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Started", nil
}

// StopServerDell ... Will Request to stop the server
// works: R730xd,R740xd
func (c *redfishProvider) StopServerDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset"

	var jsonStr = []byte(`{"ResetType": "ForceOff"}`)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Stopped", nil
}

// GracefulRestartDell ... Will Reset Idrac and will take some time to come up
func (c *redfishProvider) GracefulRestartDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Manager.Reset"

	var jsonStr = []byte(`{"ResetType": "GracefulRestart"}`)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Idrac Reset", nil

}

// ResetSSLConfig ... Will reset SSL configuration to factory default
func (c *redfishProvider) ResetSSLConfig() (string, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DelliDRACCardService/Actions/DelliDRACCardService.SSLResetCfg"

	var jsonStr = []byte(`{}`)
	// payload := "{}"
	// []byte(payload)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "SSL Reset", nil

}

// GetServerPowerStateDell ... Will fetch the current state of the Server
// works: R730xd,R740xd
func (c *redfishProvider) GetServerPowerStateDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}

	var data SystemViewDell

	json.Unmarshal(resp, &data)

	return data.PowerState, nil

}

// CheckLoginDell ... Will check the credentials of the Server
// works: R730xd,R740xd,R650,R750,R660,R760
func (c *redfishProvider) CheckLoginDell() (string, bool, error) {
	var (
		resp []byte
		data SystemViewDell
		url  = c.Hostname + "/redfish/v1/Systems/System.Embedded.1"
	)

	// if c.Certificate is there check if the certificate works
	if c.Certificate != "" {
		resp, _, _, err := queryDataForce(c, "GET", url, nil)
		if err == nil {
			json.Unmarshal(resp, &data)
			return string(data.Status.Health), true, nil
		}
		// check if error is StatusUnreachable, Unauthorized or BadRequest
		if err.Error() == StatusUnauthorized || err.Error() == StatusBadRequest {
			return "", true, err
		}
		if err.Error() == StatusUnreachable {
			return "", false, err
		}
		// if we made it here certificate is not valid
		// zero value of c.Certificate to avoid further certificate check falling back to insecureSkipVerify
		c.Certificate = ""
	}

	resp, _, _, err := queryDataForce(c, "GET", url, nil)
	if err != nil {
		return "", false, err
	}

	json.Unmarshal(resp, &data)
	return string(data.Status.Health), false, nil
}

//ImportConfigDell ... Importing the configurations to Server
/* Payload
{
    "ShareParameters": {
        "Target": "ALL"
    },
    "ImportBuffer": "<SystemConfiguration><Component FQDD=\"NIC.Integrated.1-3-1\"><Attribute Name=\"LegacyBootProto\">PXE</Attribute></Component><Component FQDD=\"NIC.Integrated.1-2-1\"><Attribute Name=\"LegacyBootProto\">PXE</Attribute></Component></SystemConfiguration>"
}
*/
func (c *redfishProvider) ImportConfigDell(jsonData []byte) (string, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ImportSystemConfiguration"
	_, _, status, err := queryData(c, "POST", url, jsonData)
	if err != nil {
		return "", err
	}
	if status == 201 || status == 202 {
		return "Accepted", nil
	}
	return "Not Accepted", nil

}

//CreateJobDell ... Create a Job based on the changed bios settings
/* Payload
   {"TargetSettingsURI":"/redfish/v1/Systems/System.Embedded.1/Bios/Settings"}
*/
func (c *redfishProvider) CreateJobDell(jsonData []byte) (string, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Jobs"
	resp, _, _, err := queryData(c, "POST", url, jsonData)
	if err != nil {
		return "", err
	}
	var k JobResponseDell
	json.Unmarshal(resp, &k)
	return k.MessageExtendedInfo[0].Message, nil
}

func (c *redfishProvider) GetJobsStatusDell() ([]JobStatusDell, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Jobs"
	var jobs []JobStatusDell
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return jobs, err
	}
	var k MemberCountDell
	json.Unmarshal(resp, &k)
	for i := range k.Members {
		_url := c.Hostname + k.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return jobs, err
		}
		var output JobStatusDell
		json.Unmarshal(resp, &output)
		jobs = append(jobs, output)
	}
	return jobs, nil
}

// GetJobStatusDell ... Get the status of the Job
func (c *redfishProvider) GetJobStatusDell(jobID string) (JobStatusDell, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Jobs/" + jobID
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return JobStatusDell{}, err
	}
	var output JobStatusDell
	json.Unmarshal(resp, &output)
	return output, nil
}

func (c *redfishProvider) GetAllJobsDell() ([]Members, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Jobs"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var k MemberCountDell
	json.Unmarshal(resp, &k)
	return k.Members, nil
}

//SetBiosSettingsDell ... Set Bios Settings
/* Payload
{"Attributes":{"BootMode": "Bios"}}
*/
func (c *redfishProvider) SetBiosSettingsDell(jsonData []byte) (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Bios/Settings"

	// var jsonStr = []byte(`{"Attributes": {"PowerCycleRequest": "FullPowerCycle"}, "@Redfish.SettingsApplyTime": {"ApplyTime": "OnReset"}}`)
	_, header, status, err := queryData(c, "PATCH", url, jsonData)

	if err != nil {
		return "", err
	}

	if status != http.StatusAccepted {
		return "", fmt.Errorf("unexpected status code: %d", status)
	}

	return header.Get("Location"), nil
}

// ClearJobsDell ... Deletes all the Jobs in the jobs queue
func (c *redfishProvider) ClearJobsDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Jobs"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}
	var k MemberCountDell
	json.Unmarshal(resp, &k)
	for i := range k.Members {
		_url := c.Hostname + k.Members[i].OdataId
		_, _, _, err := queryData(c, "DELETE", _url, nil)
		if err != nil {
			return "", err
		}
	}
	return "Jobs Deleted", nil
}

// ClearJobsDellForce ... Forces the deletion of all the Jobs in the jobs queue
func (c *redfishProvider) ClearJobsDellForce() (string, error) {

	url := c.Hostname + "/redfish/v1/Dell/Managers/iDRAC.Embedded.1/DellJobService/Actions/DellJobService.DeleteJobQueue"
	var jsonStr = []byte(`{"JobID": "JID_CLEARALL_FORCE"}`)

	_, _, status, err := queryData(c, "POST", url, jsonStr)

	if err != nil {
		return "failure", err
	}

	if status != http.StatusOK {
		return "failure", fmt.Errorf("unexpected status code: %d", status)
	}

	return "success", nil

}

//SetAttributesDell ... Will set the Attributes for IDRAC,Lifecycle Attributes and System
/* Payload
{"Attributes":{"LCAttributes.1.AutoUpdate": "1"}}
*/
func (c *redfishProvider) SetAttributesDell(service string, jsonData []byte) (string, error) {
	var url string
	if service == "idrac" {
		url = c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Attributes"
	} else if service == "lc" {
		url = c.Hostname + "/redfish/v1/Managers/LifecycleController.Embedded.1/Attributes"
	} else if service == "system" {
		url = c.Hostname + "/redfish/v1/Managers/System.Embedded.1/Attributes"
	}
	resp, _, _, err := queryData(c, "PATCH", url, jsonData)
	if err != nil {
		return "", err
	}
	var k JobResponseDell
	json.Unmarshal(resp, &k)
	return k.MessageExtendedInfo[0].Message, nil
}

// ClearStorageControllerRaidDell ... Clears Raid of the Storage Controller and returns the jub URL
func (c *redfishProvider) ClearStorageControllerRaidDell(controllerID string) (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellRaidService/Actions/DellRaidService.ResetConfig"

	data, _ := json.Marshal(map[string]interface{}{
		"TargetFQDD": controllerID,
	})

	_, header, status, err := queryData(c, "POST", url, []byte(data))

	if err != nil {
		return "", err
	}

	// check if the status is 202
	if status != http.StatusAccepted {
		return "", fmt.Errorf("unexpected status code: %d", status)
	}

	return header.Get("Location"), nil
}

// FleaDrainDell ... Will Flea Drain the Server at next reboot
func (c *redfishProvider) FleaDrainDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Bios/Settings"

	var jsonStr = []byte(`{"Attributes": {"PowerCycleRequest": "FullPowerCycle"}, "@Redfish.SettingsApplyTime": {"ApplyTime": "OnReset"}}`)
	_, header, status, err := queryData(c, "PATCH", url, jsonStr)

	if err != nil {
		return "", err
	}

	if status != http.StatusAccepted {
		return "", fmt.Errorf("unexpected status code: %d", status)
	}

	return header.Get("Location"), nil
}

// GetStorageRaidDell ... Will Fetch the Storage Raid Details
func (c *redfishProvider) GetStorageRaidDell() ([]StorageRaidDetailsDell, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Storage"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x         MemberCountDell
		_raiddata []StorageRaidDetailsDell
	)

	json.Unmarshal(resp, &x)
	for i := range x.Members {

		// check controller
		controllerUrl := c.Hostname + x.Members[i].OdataId
		respController, _, _, err := queryData(c, "GET", controllerUrl, nil)
		if err != nil {
			return nil, err
		}
		var storageControllerDell StorageControllerDell
		json.Unmarshal(respController, &storageControllerDell)

		// check volumes
		_url := c.Hostname + x.Members[i].OdataId + "/Volumes"
		respVolumes, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}
		var y MemberCountDell
		json.Unmarshal(respVolumes, &y)

		for i := range y.Members {

			_url := c.Hostname + y.Members[i].OdataId
			resp, _, _, err := queryData(c, "GET", _url, nil)
			if err != nil {
				return nil, err
			}
			//fmt.Println(string([]byte(resp)))

			var z StorageRaidRawDell
			json.Unmarshal(resp, &z)
			raidDevice := StorageRaidDetailsDell{
				Name:             z.Name,
				Id:               z.Id,
				ControllerId:     storageControllerDell.ID,
				Layout:           z.RAIDType,
				MediaType:        z.Oem.Dell.DellVirtualDisk.MediaType,
				DrivesCount:      strconv.Itoa(z.Links.DrivesCount),
				ReadCachePolicy:  z.Oem.Dell.DellVirtualDisk.ReadCachePolicy,
				CapacityBytes:    strconv.Itoa(z.CapacityBytes),
				StripeSize:       z.Oem.Dell.DellVirtualDisk.StripeSize,
				WriteCachePolicy: z.Oem.Dell.DellVirtualDisk.WriteCachePolicy,
			}

			_raiddata = append(_raiddata, raidDevice)

		}

	}
	return _raiddata, nil

}

// GetNetworkSwitchInfoDell ... Will fetch the Network Switch Info
func (c *redfishProvider) GetNetworkSwitchInfoDell() ([]SwitchData, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/NetworkPorts/Oem/Dell/DellSwitchConnections/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var x MemberCountDell
	var SwitchInfo []SwitchData
	json.Unmarshal(resp, &x)
	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}
		var y GetSwitchInfoDell
		json.Unmarshal(resp, &y)
		switchData := SwitchData{
			Name:                   y.ID,
			Description:            y.Description,
			StaleData:              y.StaleData,
			SwitchConnectionID:     y.SwitchConnectionID,
			SwitchPortConnectionID: y.SwitchPortConnectionID,
		}
		SwitchInfo = append(SwitchInfo, switchData)
	}
	// output, _ := json.Marshal(SwitchInfo)
	return SwitchInfo, nil
}

// GetNetworkPortsDell .... Will fetch network port info
func (c *redfishProvider) GetNetworkPortsDell() ([]MACData, error) {
	url := c.Hostname + "/redfish/v1/Chassis/System.Embedded.1/NetworkAdapters"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var x MemberCountDell
	var Macs []MACData
	json.Unmarshal(resp, &x)
	for i := range x.Members {

		_url := c.Hostname + x.Members[i].OdataId + "/NetworkPorts"
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}
		var y MemberCountDell
		json.Unmarshal(resp, &y)

		for i := range y.Members {

			_url := c.Hostname + y.Members[i].OdataId
			resp, _, _, err := queryData(c, "GET", _url, nil)
			if err != nil {
				return nil, err
			}
			var z NetworkPortsDell
			json.Unmarshal(resp, &z)
			var macData MACData
			if len(z.AssociatedNetworkAddresses) > 0 {
				macData = MACData{
					Name:         z.ID,
					Description:  z.Description,
					MacAddress:   z.AssociatedNetworkAddresses[0],
					Status:       z.Status.Health,
					State:        z.LinkStatus,
					PartNumber:   z.Oem.Dell.DellNetworkTransceiver.PartNumber,
					SerialNumber: z.Oem.Dell.DellNetworkTransceiver.SerialNumber,
					VendorName:   z.Oem.Dell.DellNetworkTransceiver.VendorName,
					Vlan:         "NULL",
				}
			} else {
				continue
			}

			macData.UpdateEmpty()
			Macs = append(Macs, macData)
		}

	}
	return Macs, nil
}

// GetMacAddressDell ... Will fetch all the mac address of a particular Server
func (c *redfishProvider) GetMacAddressDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/EthernetInterfaces/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}
	var x MemberCountDell
	var Macs []MACData
	json.Unmarshal(resp, &x)
	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return "", err
		}
		var y GetMacAddressDell
		json.Unmarshal(resp, &y)
		macData := MACData{
			Name:        y.ID,
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

// GetIdracLicenses ... Will fetch all iDRAC licenses
func (c *redfishProvider) GetIdracLicenses() ([]LicenseData, error) {
	url := c.Hostname + "/redfish/v1/LicenseService/Licenses/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return []LicenseData{}, err
	}
	var x MemberCountDell
	var Licenses []LicenseData
	json.Unmarshal(resp, &x)
	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return []LicenseData{}, err
		}
		var y GetLicenseDell
		json.Unmarshal(resp, &y)
		licenseData := LicenseData{
			Description: y.Description,
			Id:          y.ID,
			LicenseType: y.LicenseType,
			Status:      y.Status.Health,
		}
		Licenses = append(Licenses, licenseData)
	}

	if len(Licenses) > 0 {
		return Licenses, nil
	}

	// fall back on OEM API if generic API doesn't work
	url = c.Hostname + "/redfish/v1/Dell/Managers/iDRAC.Embedded.1/DellLicenseCollection"
	resp, _, _, err = queryData(c, "GET", url, nil)
	if err != nil {
		return []LicenseData{}, err
	}
	json.Unmarshal(resp, &x)
	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return []LicenseData{}, err
		}
		var (
			y                  GetLicenseDellOEM
			licenseDescription string
		)
		json.Unmarshal(resp, &y)
		if len(y.LicenseDescription) > 0 {
			licenseDescription = y.LicenseDescription[0]
		}
		licenseData := LicenseData{
			Description: licenseDescription,
			Id:          y.ID,
			LicenseType: y.LicenseType,
			Status:      y.LicensePrimaryStatus,
		}
		Licenses = append(Licenses, licenseData)
	}

	return Licenses, nil
}

// GetMacAddressModelDell ... Will fetch the Nic Model
func (c *redfishProvider) GetMacAddressModelDell() ([]MACModelDell, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/NetworkAdapters/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var x MemberCountDell
	var Macs []MACModelDell
	json.Unmarshal(resp, &x)
	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}
		var y NetworkDeviceDell
		json.Unmarshal(resp, &y)

		for _, k := range y.Controllers {
			for _, z := range k.Links.NetworkDeviceFunctions {
				firmName := strings.Split(z.OdataId, "/")
				result := MACModelDell{
					MacName:         firmName[len(firmName)-1],
					MacModel:        y.Model,
					MacManufacturer: y.Manufacturer,
				}
				Macs = append(Macs, result)
			}
		}
	}
	return Macs, nil

}

// GetProcessorHealthDell ... Will Fetch the Processor Health Details
// works: R730xd,R740xd
func (c *redfishProvider) GetProcessorHealthDell() ([]HealthList, error) {
	///redfish/v1/Systems/System.Embedded.1/Processors

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Processors"
	resp, _, _, err := queryData(c, "GET", url, nil)
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
		resp, _, _, err := queryData(c, "GET", _url, nil)
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

// func (c *redfishProvider) GetMemoryHealthDell() (string, error) {}

// GetPowerHealthDell ... Will Fetch the Power Health Details
// works: R730xd,R740xd
func (c *redfishProvider) GetPowerHealthDell() ([]HealthList, error) {
	url := c.Hostname + "/redfish/v1/Chassis/System.Embedded.1/Power"

	resp, _, _, err := queryData(c, "GET", url, nil)
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

// GetSensorsHealthDell ... Will Fetch the Sensors Health Details
// works: R730xd,R740xd
func (c *redfishProvider) GetSensorsHealthDell() ([]HealthList, error) {

	url := c.Hostname + "/redfish/v1/Chassis/System.Embedded.1/Thermal"

	resp, _, _, err := queryData(c, "GET", url, nil)
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

	if x.Temperaturescount != 0 {
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

// GetStorageDriveDetailsDell ... Will Fetch the Storage Drive Details
func (c *redfishProvider) GetStorageDriveDetailsDell() ([]StorageDriveDetailsDell, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Storage"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x          StorageCollectionDell
		_drivedata []StorageDriveDetailsDell
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {

		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}

		var y StorageDetailsDell

		json.Unmarshal(resp, &y)

		if y.Drivescount != 0 {
			for k := range y.Drives {
				_url := c.Hostname + y.Drives[k].OdataId
				resp, _, _, err := queryData(c, "GET", _url, nil)
				if err != nil {
					return nil, err
				}
				var z StorageDriveDetailsDell

				json.Unmarshal(resp, &z)

				_drivedata = append(_drivedata, z)
			}

		} else {
			continue
		}

	}
	return _drivedata, nil

}

// GetStorageHealthDell ... Will Fetch the Storage Health Details
// works: R730xd,R740xd
func (c *redfishProvider) GetStorageHealthDell() ([]StorageHealthList, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Storage"

	resp, _, _, err := queryData(c, "GET", url, nil)
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
		resp, _, _, err := queryData(c, "GET", _url, nil)
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

		if y.Drivescount != 0 {
			for k := range y.Drives {
				_url := c.Hostname + y.Drives[k].OdataId
				resp, _, _, err := queryData(c, "GET", _url, nil)
				if err != nil {
					return nil, err
				}
				var z StorageDriveDetailsDell

				json.Unmarshal(resp, &z)

				storageHealth := StorageHealthList{
					Name:   z.Name,
					Health: z.Status.Health,
					State:  z.Status.State,
					Space:  z.CapacityBytes,
				}
				_healthdata = append(_healthdata, storageHealth)
			}

		} else {
			continue
		}

	}
	return _healthdata, nil

}

// GetAggHealthDataDell ... will fetch the data related to all components health(aggregated view)
func (c *redfishProvider) GetAggHealthDataDell(model string) ([]HealthList, error) {

	if strings.ToLower(model) == "r730xd" {

		return nil, nil

	} else if strings.ToLower(model) == "r740xd" || strings.ToLower(model) == "r650" || strings.ToLower(model) == "r750" {
		url := c.Hostname + "/redfish/v1/UpdateService/FirmwareInventory"

		resp, _, _, err := queryData(c, "GET", url, nil)
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
				resp, _, _, err := queryData(c, "GET", _url, nil)
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
	return nil, nil
}

// GetFirmwareDell ... will fetch the Firmware details
func (c *redfishProvider) GetFirmwareDell() ([]FirmwareData, error) {

	url := c.Hostname + "/redfish/v1/UpdateService/FirmwareInventory"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x         MemberCountDell
		_firmdata []FirmwareData
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {

		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
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

	return _firmdata, nil

}

// FirmwareUpdateDell ... will create a job plan for firmware update
func (c *redfishProvider) FirmwareUpdateDell() (string, error) {
	url := c.Hostname + "/redfish/v1/UpdateService/FirmwareInventory"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}

	var (
		x         MemberCountDell
		firmLinks []string
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		r, _ := regexp.Compile("Available")
		if r.MatchString(x.Members[i].OdataId) == true {

			firmLinks = append(firmLinks, x.Members[i].OdataId)

		}
	}

	data, _ := json.Marshal(map[string]interface{}{
		"SoftwareIdentityURIs": firmLinks,
		"InstallUpon":          "NowAndReboot",
	})

	firmUrl := c.Hostname + "/redfish/v1/UpdateService/Actions/Oem/DellUpdateService.Install"
	_, header, _, errr := queryData(c, "POST", firmUrl, []byte(data))
	if errr != nil {
		return "", err
	}

	var taskURL string

	for k, v := range header {
		if k == "Location" {
			taskURL = v[0]
			break
		}
	}

	if len(firmLinks) == 0 {
		return "No Firmware Available", nil
	} else {
		return taskURL, nil
	}

}

// FirmwareUploadDell ... will fetch the payload from remote repo
func (c *redfishProvider) FirmwareUploadDell(repoUrl string) (string, error) {

	url := c.Hostname + "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate"

	data, _ := json.Marshal(map[string]interface{}{
		"ImageURI": repoUrl,
	})

	_, headers, _, err := queryData(c, "POST", url, []byte(data))
	if err != nil {
		return "", err
	}

	return headers["Location"][0], nil
}

func (c *redfishProvider) TaskStatusDell(taskUrl string) (ExportConfigStatus, error) {
	url := c.Hostname + taskUrl

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return ExportConfigStatus{}, err
	}

	var x ExportConfigStatus

	json.Unmarshal(resp, &x)

	return x, nil

}

// GetBiosDataDell ... will fetch the Bios Details
func (c *redfishProvider) GetBiosDataDell() (BiosAttributesData, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/Bios"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return BiosAttributesData{}, err
	}

	var x BiosAttrDell

	json.Unmarshal(resp, &x)

	return x.Attributes, nil

}

// GetLifecycleAttrDell ... will fetch the lifecycle attributes
func (c *redfishProvider) GetLifecycleAttrDell() (LifeCycleData, error) {

	url := c.Hostname + "/redfish/v1/Managers/LifecycleController.Embedded.1/Attributes"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return LifeCycleData{}, err
	}

	var x LifeCycleAttrDell

	json.Unmarshal(resp, &x)

	_data := x.Attributes

	_LfcycleDat := LifeCycleData{
		AutoBackup:                          _data.AutoBackup,
		AutoDiscovery:                       _data.AutoDiscovery,
		AutoUpdate:                          _data.AutoUpdate,
		BIOSRTDRequested:                    _data.BIOSRTDRequested,
		CollectSystemInventoryOnRestart:     _data.CollectSystemInventoryOnRestart,
		DiscoveryFactoryDefaults:            _data.DiscoveryFactoryDefaults,
		IPAddress:                           _data.IPAddress,
		IPChangeNotifyPS:                    _data.IPChangeNotifyPS,
		IgnoreCertWarning:                   _data.IgnoreCertWarning,
		Licensed:                            _data.Licensed,
		LifecycleControllerState:            _data.LifecycleControllerState,
		PartConfigurationUpdate:             _data.PartConfigurationUpdate,
		PartFirmwareUpdate:                  _data.PartFirmwareUpdate,
		ProvisioningServer:                  _data.ProvisioningServer,
		StorageHealthRollupStatus:           _data.StorageHealthRollupStatus,
		SystemID:                            _data.SystemID,
		UserProxyPassword:                   _data.UserProxyPassword,
		UserProxyPort:                       _data.UserProxyPort,
		UserProxyServer:                     _data.UserProxyServer,
		UserProxyType:                       _data.UserProxyType,
		UserProxyUserName:                   _data.UserProxyUserName,
		VirtualAddressManagementApplication: _data.VirtualAddressManagementApplication,
	}

	return _LfcycleDat, nil

}

// ListUsersDell ...
func (c *redfishProvider) ListUsersDell() ([]UserListDell, error) {

	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Accounts"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x         MemberCountDell
		_userdata []UserListDell
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}
		var y UserListResponseDell

		json.Unmarshal(resp, &y)

		userData := UserListDell{
			UserName: y.UserName,
			RoleID:   y.RoleID,
			Enabled:  y.Enabled,
			Locked:   y.Locked,
		}

		_userdata = append(_userdata, userData)
	}
	return _userdata, nil
}

// CreateUserDell ... will create a new user
func (c *redfishProvider) CreateUserDell(num int, username string, password string, role string, status bool) (string, error) {
	url := fmt.Sprintf("%s/redfish/v1/Managers/iDRAC.Embedded.1/Accounts/%d", c.Hostname, num)
	data, _ := json.Marshal(map[string]interface{}{
		"UserName": username,
		"Password": password,
		"Enabled":  status,
		"RoleId":   role,
	})

	resp, _, _, err := queryData(c, "PATCH", url, []byte(data))
	if err != nil {
		return "", err
	}
	var k JobResponseDell
	json.Unmarshal(resp, &k)
	return k.MessageExtendedInfo[0].Message, nil
}

// DeleteUserDell ... will delete a user
// func (c *redfishProvider) DeleteUserDell(num int, role string, status bool) (string, error) {
func (c *redfishProvider) DeleteUserDell(num int) (string, error) {
	url := fmt.Sprintf("%s/redfish/v1/Managers/iDRAC.Embedded.1/Accounts/%d", c.Hostname, num)
	// data, _ := json.Marshal(map[string]interface{}{
	// 	"Enabled":  status,
	// })
	data, _ := json.Marshal(map[string]interface{}{
		"Enabled":  false,
		"UserName": "",
	})

	resp, _, _, err := queryData(c, "PATCH", url, []byte(data))
	if err != nil {
		return "", err
	}
	var k JobResponseDell
	json.Unmarshal(resp, &k)
	return k.MessageExtendedInfo[0].Message, nil
}

// GetIDRACAttrDell ... will fetch the Idrac attributes
func (c *redfishProvider) GetIDRACAttrDell() (IDRACAttributesData, error) {

	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Attributes"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return IDRACAttributesData{}, err
	}

	var x IDRACAttrDell

	json.Unmarshal(resp, &x)

	return x.Attributes, nil

}

// GetSysAttrDell ... will fetch the System Attributes
func (c *redfishProvider) GetSysAttrDell() (SysAttributesData, error) {

	url := c.Hostname + "/redfish/v1/Managers/System.Embedded.1/Attributes"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return SysAttributesData{}, err
	}

	var x SysAttrDell

	json.Unmarshal(resp, &x)

	return x.Attributes, nil

}

// GetBootOrderDell ... will fetch the BootOrder Details
func (c *redfishProvider) GetBootOrderDell() ([]BootOrderData, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/BootSources"

	resp, _, _, err := queryData(c, "GET", url, nil)
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
			ID:      x.Attributes.BootSeq[i].ID,
		}

		_bootOrder = append(_bootOrder, _result)
	}

	for i := range x.Attributes.UefiBootSeq {

		_result := BootOrderData{
			Enabled: x.Attributes.UefiBootSeq[i].Enabled,
			Index:   x.Attributes.UefiBootSeq[i].Index,
			Name:    x.Attributes.UefiBootSeq[i].Name,
			ID:      x.Attributes.UefiBootSeq[i].ID,
		}

		_bootOrder = append(_bootOrder, _result)
	}

	return _bootOrder, nil

}

// SetBootOrderDell ... Set the Boot Order f
func (c *redfishProvider) SetBootOrderDell(jsonData []byte) (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1/BootSources/Settings"
	resp, _, _, err := queryData(c, "PATCH", url, jsonData)
	if err != nil {
		return "", err
	}

	var k JobResponseDell
	json.Unmarshal(resp, &k)
	if len(k.MessageExtendedInfo) > 0 {
		return k.MessageExtendedInfo[0].Message, nil
	} else {
		return "", nil
	}

}

// GetSystemEventLogsDell ... Fetch the System Event Logs from the Idrac
func (c *redfishProvider) GetSystemEventLogsDell(version string) ([]SystemEventLogRes, error) {

	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Logs/Sel"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	v1, verErr := ver.NewConstraint("<= 3.15.17.15")
	if verErr != nil {
		return nil, verErr
	}
	v2, verErr := ver.NewConstraint("<= 3.21.26.22")
	if verErr != nil {
		return nil, verErr
	}
	v3, verErr := ver.NewConstraint("> 3.21.26.22")
	if verErr != nil {
		return nil, verErr
	}
	v4, verErr := ver.NewVersion(version)
	if verErr != nil {
		return nil, verErr
	}

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

// GetLifeCycleEventLogsDell ... Fetch the LifeCycle Event Logs from the Idrac
func (c *redfishProvider) GetLifeCycleEventLogsDell(totalPages int) ([]LifeCycleEventLogRes, error) {

	// totalPages corresponds to total pages of 50 logs
	var _lfyCycleEventLogs []LifeCycleEventLogRes
	var skipCount int

	size := make([]struct{}, totalPages)
	for range size {

		// fmt.Println("i is:", i)
		// fmt.Println("skipCount is:", skipCount)
		url := fmt.Sprintf("%s/%s%d", c.Hostname, "redfish/v1/Managers/iDRAC.Embedded.1/LogServices/Lclog/Entries?$skip=", skipCount)

		resp, _, _, err := queryData(c, "GET", url, nil)
		if err != nil {
			return nil, err
		}

		var x LifeCycleLogsV1Dell

		json.Unmarshal(resp, &x)

		for i := range x.Members {

			_result := LifeCycleEventLogRes{
				Created:     x.Members[i].Created,
				Description: x.Members[i].Description,
				EntryType:   x.Members[i].EntryType,
				ID:          x.Members[i].ID,
				Message:     x.Members[i].Message,
				MessageID:   x.Members[i].MessageID,
				Name:        x.Members[i].Name,
				Severity:    x.Members[i].Severity,
			}

			_lfyCycleEventLogs = append(_lfyCycleEventLogs, _result)
		}

		if len(x.Members) < 50 {
			return _lfyCycleEventLogs, nil
		}

		skipCount += len(x.Members)
	}

	return _lfyCycleEventLogs, nil
}

// WriteLCLog ... Will write entry to LC Log
func (c *redfishProvider) WriteLCLog(messageDesctiption string) (string, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/DellLCService/Actions/DellLCService.InsertCommentInLCLog"

	// var jsonStr = []byte(`{"Comment": "dummyentry"}`)
	var jsonStr = []byte(`{"Comment": "` + messageDesctiption + `"}`)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "LCLog Entry Written", nil
}

// GetUserAccountsDell ... Fetch the current users created
func (c *redfishProvider) GetUserAccountsDell() ([]Accounts, error) {

	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Accounts"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var x MemberCountDell
	var users []Accounts

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
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

// GetSystemInfoDell ... Will fetch the system info
func (c *redfishProvider) GetSystemInfoDell() (SystemData, error) {

	url := c.Hostname + "/redfish/v1/Systems/System.Embedded.1"

	resp, _, _, err := queryData(c, "GET", url, nil)
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
		ServiceTag:      x.SKU,
	}

	return _result, nil

}

// GetComponentAttr ... Will fetch all the component level attributes
// Supported values are: ALL, System, BIOS, IDRAC, NIC, FC, LifecycleController, RAID.
func (c *redfishProvider) GetComponentAttr(comp string) (ExportConfigResponse, error) {

	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/Actions/Oem/EID_674_Manager.ExportSystemConfiguration"
	data, _ := json.Marshal(map[string]interface{}{
		"ExportFormat": "JSON",
		"ShareParameters": map[string]interface{}{
			"Target": comp,
		},
	})

	_, header, _, err := queryData(c, "POST", url, []byte(data))
	if err != nil {
		return ExportConfigResponse{}, err
	}
	var taskURL string

	for k, v := range header {
		if k == "Location" {
			taskURL = v[0]
			break
		}
	}

	for {
		taskUrl := c.Hostname + taskURL

		resp, _, _, err := queryData(c, "GET", taskUrl, nil)
		if err != nil {
			return ExportConfigResponse{}, err
		}

		var x ExportConfigStatus

		json.Unmarshal(resp, &x)

		if x.TaskState == "Running" {
			x = ExportConfigStatus{}
			time.Sleep(time.Minute)
		} else {
			var y ExportConfigResponse
			json.Unmarshal(resp, &y)
			return y, nil
		}
	}

	return ExportConfigResponse{}, nil
}

// MountImageDell ... Will mount a image over http share
// Supports for 4.x Firmware
func (c *redfishProvider) MountImageDell(image string) (string, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.InsertMedia"

	data, _ := json.Marshal(map[string]interface{}{
		"Image":          image,
		"Inserted":       true,
		"WriteProtected": true,
	})

	_, _, status, err := queryData(c, "POST", url, []byte(data))
	if err != nil {
		return "", err
	}

	if status == 204 {
		return "Image Uploaded", nil
	} else if status == 500 {
		return "Image Not Uploaded", err
	}

	return "", nil
}

// UnMountImageDell ... Will unmount a imoge
// Supports for 4.x Firmware
func (c *redfishProvider) UnMountImageDell() (string, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD/Actions/VirtualMedia.EjectMedia"
	payload := "{}"
	_, _, _, err := queryData(c, "POST", url, []byte(payload))
	if err != nil {
		return "", err
	}
	return "Image Unmounted", nil
}

// GetRemoteImageStatusDell ... Get remote image status
func (c *redfishProvider) GetRemoteImageStatusDell() (ImageStatusDell, error) {
	url := c.Hostname + "/redfish/v1/Managers/iDRAC.Embedded.1/VirtualMedia/CD"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return ImageStatusDell{}, err
	}

	var x ImageStatusDell

	json.Unmarshal(resp, &x)

	return x, nil
}

// UpdateFirmwareDell ... will update Dell server firmware
func (c *redfishProvider) UpdateFirmwareDell(firmwareDir string, firmwareFile string) (string, error) {

	url := c.Hostname + "/redfish/v1/UpdateService/MultipartUpload"
	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)

	// added the below to create custom form field since CreateFormField doesn't set the Content-Type
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="UpdateParameters"`)
	h.Set("Content-Type", "application/json")
	formField, err := writer.CreatePart(h)
	if err != nil {
		return "", err
	}

	_, err = formField.Write([]byte(`{"@Redfish.OperationApplyTime":"Immediate"}`))
	if err != nil {
		return "", err
	}

	fw, err := writer.CreateFormFile("UpdateFile", firmwareFile)
	if err != nil {
		return "", err
	}
	fd, err := os.Open(firmwareDir + "/" + firmwareFile)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	_, err = io.Copy(fw, fd)
	if err != nil {
		return "", err
	}

	writer.Close()

	_, header, status, err := postForm(c, url, form, writer.FormDataContentType())

	if err != nil {
		return "", err
	}

	if status != http.StatusAccepted {
		return "", fmt.Errorf("unexpected status code: %d", status)
	}

	return header.Get("Location"), nil
}
