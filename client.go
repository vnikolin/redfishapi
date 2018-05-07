package hpilo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

type IloClient struct {
	Hostname string
	Username string
	Password string
}

type FirmwareVersion struct {
	FirmwareDate        string `json:"firmware_date"`
	FirmwareVersion     string `json:"firmware_version"`
	LicenseType         string `json:"license_type"`
	ManagementProcessor string `json:"management_processor"`
}

type MediaResponse struct {
	Response string `json:"response"`
}

type HostInfo struct {
	Serial      string `json:"serial"`
	ServerModel string `json:"server_model"`
	UUID        string `json:"uuid"`
}

//VersionRIBCL struct is to fetch the response from the Firmware Version
type VersionRIBCL struct {
	GET_FW_VERSION GET_FW_VERSION `xml:"GET_FW_VERSION"`
	VERSION        string         `xml:"VERSION,attr"`
}

type RESPONSE struct {
	STATUS  string `xml:"STATUS,attr"`
	MESSAGE string `xml:"MESSAGE,attr"`
}

type GET_FW_VERSION struct {
	FIRMWARE_VERSION     string `xml:"FIRMWARE_VERSION,attr"`
	FIRMWARE_DATE        string `xml:"FIRMWARE_DATE,attr"`
	MANAGEMENT_PROCESSOR string `xml:"MANAGEMENT_PROCESSOR,attr"`
	LICENSE_TYPE         string `xml:"LICENSE_TYPE,attr"`
}

type HostRIBCL struct {
	VERSION       string          `xml:"VERSION,attr"`
	SMBIOS_RECORD []SMBIOS_RECORD `xml:"GET_HOST_DATA>SMBIOS_RECORD,omitempty"`
	FIELD         [][]FIELD       `xml:"GET_HOST_DATA>SMBIOS_RECORD>FIELD,omitempty"`
}

type FIELD struct {
	VALUE string `xml:"VALUE,attr"`
	NAME  string `xml:"NAME,attr"`
}
type SMBIOS_RECORD struct {
	TYPE     string `xml:"TYPE,attr"`
	B64_DATA string `xml:"B64_DATA,attr"`
}

//RIMP will fetch the data from the xmldata api
type RIMP struct {
	DESCRIPTION        []string `xml:"HSI>NICS>NIC>DESCRIPTION"`
	CUUIDVIDVIRTUALHSI string   `xml:"HSI>VIRTUAL>VID>cUUID"`
	PWRM               string   `xml:"MP>PWRM"`
	EALERT             string   `xml:"MP>EALERT"`
	CUUID              string   `xml:"HSI>cUUID"`
	STATE              string   `xml:"HSI>VIRTUAL>STATE"`
	UUIDMP             string   `xml:"MP>UUID"`
	ERS                string   `xml:"MP>ERS"`
	LOCATION           []string `xml:"HSI>NICS>NIC>LOCATION"`
	PORT               []string `xml:"HSI>NICS>NIC>PORT"`
	ST                 string   `xml:"MP>ST"`
	SSO                string   `xml:"MP>SSO"`
	UUID               string   `xml:"HSI>UUID"`
	FWRI               string   `xml:"MP>FWRI"`
	BBLK               string   `xml:"MP>BBLK"`
	STATUSHEALTH       string   `xml:"HEALTH>STATUS"`
	SBSN               string   `xml:"HSI>SBSN"`
	SPN                string   `xml:"HSI>SPN"`
	PRODUCTID          string   `xml:"HSI>PRODUCTID"`
	MACADDR            []string `xml:"HSI>NICS>NIC>MACADDR"`
	HWRI               string   `xml:"MP>HWRI"`
	SP                 string   `xml:"HSI>SP"`
	PN                 string   `xml:"MP>PN"`
	SN                 string   `xml:"MP>SN"`
	IPADDR             []string `xml:"HSI>NICS>NIC>IPADDR"`
	STATUS             []string `xml:"HSI>NICS>NIC>STATUS"`
	BSN                string   `xml:"HSI>VIRTUAL>VID>BSN"`
	IPM                string   `xml:"MP>IPM"`
}

type HealthData struct {
	XMLName                  xml.Name                  `xml:"RIBCL,omitempty" json:"RIBCL,omitempty"`
	VERSION                  string                    `xml:"VERSION,attr"  json:",omitempty"`
	GET_EMBEDDED_HEALTH_DATA *GET_EMBEDDED_HEALTH_DATA `xml:"GET_EMBEDDED_HEALTH_DATA,omitempty" json:"GET_EMBEDDED_HEALTH_DATA,omitempty"`
}

type GET_EMBEDDED_HEALTH_DATA struct {
	XMLName            xml.Name            `xml:"GET_EMBEDDED_HEALTH_DATA,omitempty" json:"GET_EMBEDDED_HEALTH_DATA,omitempty"`
	HEALTH_AT_A_GLANCE *HEALTH_AT_A_GLANCE `xml:"HEALTH_AT_A_GLANCE,omitempty" json:"HEALTH_AT_A_GLANCE,omitempty"`
}

type HEALTH_AT_A_GLANCE struct {
	XMLName        xml.Name        `xml:"HEALTH_AT_A_GLANCE,omitempty" json:"HEALTH_AT_A_GLANCE,omitempty"`
	BATTERY        *BATTERY        `xml:"BATTERY,omitempty" json:"BATTERY,omitempty"`
	BIOS_HARDWARE  *BIOS_HARDWARE  `xml:"BIOS_HARDWARE,omitempty" json:"BIOS_HARDWARE,omitempty"`
	FANS           *FANS           `xml:"FANS,omitempty" json:"FANS,omitempty"`
	MEMORY         *MEMORY         `xml:"MEMORY,omitempty" json:"MEMORY,omitempty"`
	NETWORK        *NETWORK        `xml:"NETWORK,omitempty" json:"NETWORK,omitempty"`
	POWER_SUPPLIES *POWER_SUPPLIES `xml:"POWER_SUPPLIES,omitempty" json:"POWER_SUPPLIES,omitempty"`
	PROCESSOR      *PROCESSOR      `xml:"PROCESSOR,omitempty" json:"PROCESSOR,omitempty"`
	STORAGE        *STORAGE        `xml:"STORAGE,omitempty" json:"STORAGE,omitempty"`
	TEMPERATURE    *TEMPERATURE    `xml:"TEMPERATURE,omitempty" json:"TEMPERATURE,omitempty"`
}

type FANS struct {
	XMLName        xml.Name `xml:"FANS,omitempty" json:"FANS,omitempty"`
	AttrREDUNDANCY string   `xml:"REDUNDANCY,attr"  json:",omitempty"`
	AttrSTATUS     string   `xml:"STATUS,attr"  json:",omitempty"`
}

type MEMORY struct {
	XMLName    xml.Name `xml:"MEMORY,omitempty" json:"MEMORY,omitempty"`
	AttrSTATUS string   `xml:"STATUS,attr"  json:",omitempty"`
}

type BATTERY struct {
	XMLName    xml.Name `xml:"BATTERY,omitempty" json:"BATTERY,omitempty"`
	AttrSTATUS string   `xml:"STATUS,attr"  json:",omitempty"`
}

type BIOS_HARDWARE struct {
	XMLName    xml.Name `xml:"BIOS_HARDWARE,omitempty" json:"BIOS_HARDWARE,omitempty"`
	AttrSTATUS string   `xml:"STATUS,attr"  json:",omitempty"`
}

type NETWORK struct {
	XMLName    xml.Name `xml:"NETWORK,omitempty" json:"NETWORK,omitempty"`
	AttrSTATUS string   `xml:"STATUS,attr"  json:",omitempty"`
}

type POWER_SUPPLIES struct {
	XMLName        xml.Name `xml:"POWER_SUPPLIES,omitempty" json:"POWER_SUPPLIES,omitempty"`
	AttrREDUNDANCY string   `xml:"REDUNDANCY,attr"  json:",omitempty"`
	AttrSTATUS     string   `xml:"STATUS,attr"  json:",omitempty"`
}

type PROCESSOR struct {
	XMLName    xml.Name `xml:"PROCESSOR,omitempty" json:"PROCESSOR,omitempty"`
	AttrSTATUS string   `xml:"STATUS,attr"  json:",omitempty"`
}

type CRIBCL struct {
	XMLName                  xml.Name                  `xml:"RIBCL,omitempty" json:"RIBCL,omitempty"`
	VERSION                  string                    `xml:"VERSION,attr"  json:",omitempty"`
	GET_EMBEDDED_HEALTH_DATA *GET_EMBEDDED_HEALTH_DATA `xml:"GET_EMBEDDED_HEALTH_DATA,omitempty" json:"GET_EMBEDDED_HEALTH_DATA,omitempty"`
}

type STORAGE struct {
	XMLName    xml.Name `xml:"STORAGE,omitempty" json:"STORAGE,omitempty"`
	AttrSTATUS string   `xml:"STATUS,attr"  json:",omitempty"`
}

type TEMPERATURE struct {
	XMLName    xml.Name `xml:"TEMPERATURE,omitempty" json:"TEMPERATURE,omitempty"`
	AttrSTATUS string   `xml:"STATUS,attr"  json:",omitempty"`
}

type HealthInfo struct {
	//BatteryStatus           string `json:"battery_status"`
	BiosStatus              string `json:"bios_status"`
	FanRedundancy           string `json:"fan_redundancy"`
	FanStatus               string `json:"fan_status"`
	MemoryStatus            string `json:"memory_status"`
	NetworkStatus           string `json:"network_status"`
	PowerSupliesStatus      string `json:"power_suplies_status"`
	PowerSuppliesRedundancy string `json:"power_supplies_redundancy"`
	ProcessorStatus         string `json:"processor_status"`
	StorageStatus           string `json:"storage_status"`
	TemperatureStatus       string `json:"temperature_status"`
}

type HardwareInfo struct {
	ExtraData string `json:"extra_data"`
}

func NewIloClient(hostname string, username string, password string) *IloClient {

	return &IloClient{
		Hostname: hostname,
		Username: username,
		Password: password,
	}
}

func (c *IloClient) GetFwVersion() (string, error) {

	url := c.Hostname + "/ribcl"

	body := "<?xml version=\"1.0\" ?>" +
		"<?xmlilo output-format=\"xml\"?>" +
		"<?ilo entity-processing=\"standard\"?>" +
		"<RIBCL VERSION=\"2.0\">" +
		"<LOGIN PASSWORD=\"" + c.Password + "\" USER_LOGIN=\"" + c.Username + "\">" +
		"<RIB_INFO MODE=\"read\">" +
		"<GET_FW_VERSION/>" +
		"</RIB_INFO>" +
		"</LOGIN>" +
		"</RIBCL>"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)

	var x VersionRIBCL

	xml.Unmarshal(_body, &x)

	fwVersion := FirmwareVersion{
		FirmwareDate:        x.GET_FW_VERSION.FIRMWARE_DATE,
		FirmwareVersion:     x.GET_FW_VERSION.FIRMWARE_VERSION,
		LicenseType:         x.GET_FW_VERSION.LICENSE_TYPE,
		ManagementProcessor: x.GET_FW_VERSION.MANAGEMENT_PROCESSOR,
	}

	output, _ := json.Marshal(fwVersion)

	return string(output), nil

}

func (c *IloClient) GetHostData() (string, error) {

	url := c.Hostname + "/xmldata?item=All"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)

	var x RIMP

	xml.Unmarshal(_body, &x)

	hostData := HostInfo{
		UUID:        strings.TrimSpace(x.UUID),
		ServerModel: strings.TrimSpace(x.SPN),
		Serial:      strings.TrimSpace(x.SBSN),
	}

	output, _ := json.Marshal(hostData)

	return string(output), nil

}

func (c *IloClient) InsertMedia(image string) (string, error) {

	url := c.Hostname + "/ribcl"

	_imageBody := "<?xml version=\"1.0\" ?>" +
		"<?xmlilo output-format=\"xml\"?>" +
		"<?ilo entity-processing=\"standard\"?>" +
		"<RIBCL VERSION=\"2.0\">" +
		"<LOGIN PASSWORD=\"" + c.Password + "\" USER_LOGIN=\"" + c.Username + "\">" +
		"<RIB_INFO MODE=\"write\">" +
		"<INSERT_VIRTUAL_MEDIA DEVICE=\"CDROM\" IMAGE_URL=\"" + image + "\" />" +
		"</RIB_INFO>" +
		"</LOGIN>" +
		"</RIBCL>"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(_imageBody)))
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// _body, _ := ioutil.ReadAll(resp.Body)

	_response := MediaResponse{
		Response: "Image url Added to the Media",
	}

	output, _ := json.Marshal(_response)

	return string(output), nil

}

func (c *IloClient) SetVmStatus() (string, error) {

	url := c.Hostname + "/ribcl"

	_imageAccess := "<?xml version=\"1.0\" ?>" +
		"<?xmlilo output-format=\"xml\"?>" +
		"<?ilo entity-processing=\"standard\"?>" +
		"<RIBCL VERSION=\"2.0\">" +
		"<LOGIN PASSWORD=\"" + c.Password + "\" USER_LOGIN=\"" + c.Username + "\">" +
		"<RIB_INFO MODE=\"write\">" +
		"<SET_VM_STATUS DEVICE=\"CDROM\">" +
		"<VM_BOOT_OPTION VALUE=\"BOOT_ONCE\"/>" +
		"<VM_WRITE_PROTECT VALUE=\"YES\" />" +
		"</SET_VM_STATUS>" +
		"</RIB_INFO>" +
		"</LOGIN>" +
		"</RIBCL>"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(_imageAccess)))
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_response := MediaResponse{
		Response: "Media Set to Boot Once",
	}

	output, _ := json.Marshal(_response)

	return string(output), nil
}

func (c *IloClient) RestartServer() (string, error) {

	url := c.Hostname + "/ribcl"

	_serverRestart := "<?xml version=\"1.0\" ?>" +
		"<?xmlilo output-format=\"xml\"?>" +
		"<?ilo entity-processing=\"standard\"?>" +
		"<RIBCL VERSION=\"2.0\">" +
		"<LOGIN PASSWORD=\"" + c.Password + "\" USER_LOGIN=\"" + c.Username + "\">" +
		"<SERVER_INFO MODE=\"write\">" +
		"<RESET_SERVER/>" +
		"</SERVER_INFO>" +
		"</LOGIN>" +
		"</RIBCL>"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(_serverRestart)))
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_response := MediaResponse{
		Response: "Server Sent for a restart",
	}

	output, _ := json.Marshal(_response)

	return string(output), nil
}

func (c *IloClient) GetHealthData() (string, error) {

	url := c.Hostname + "/ribcl"

	body := "<?xml version=\"1.0\"?>" +
		"<?ilo entity-processing=\"standard\"?>" +
		"<?xmlilo output-format=\"xml\"?>" +
		"<RIBCL VERSION=\"2.0\">" +
		"<LOGIN PASSWORD=\"" + c.Password + "\" USER_LOGIN=\"" + c.Username + "\">" +
		"<SERVER_INFO MODE=\"read\">" +
		"<GET_EMBEDDED_HEALTH>" +
		"</GET_EMBEDDED_HEALTH>" +
		"</SERVER_INFO>" +
		"</LOGIN>" +
		"</RIBCL>"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)

	var healthData CRIBCL
	xml.Unmarshal(_body, &healthData)

	_healthData, _ := json.Marshal(HealthInfo{
		BiosStatus:              string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.BIOS_HARDWARE.AttrSTATUS),
		FanStatus:               string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.FANS.AttrSTATUS),
		FanRedundancy:           string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.FANS.AttrREDUNDANCY),
		TemperatureStatus:       string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.TEMPERATURE.AttrSTATUS),
		PowerSupliesStatus:      string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.POWER_SUPPLIES.AttrSTATUS),
		PowerSuppliesRedundancy: string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.POWER_SUPPLIES.AttrREDUNDANCY),
		//BatteryStatus:           string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.BATTERY.AttrSTATUS),
		ProcessorStatus: string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.PROCESSOR.AttrSTATUS),
		MemoryStatus:    string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.MEMORY.AttrSTATUS),
		NetworkStatus:   string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.NETWORK.AttrSTATUS),
		StorageStatus:   string(healthData.GET_EMBEDDED_HEALTH_DATA.HEALTH_AT_A_GLANCE.STORAGE.AttrSTATUS),
	},
	)

	return string(_healthData), nil
}

//GetHardwareData will fetch all the data in raw xaml format
//will be Depreciated soon to get the list of fields in different structs
func (c *IloClient) GetHardwareData() (string, error) {

	url := c.Hostname + "/ribcl"

	body := "<?xml version=\"1.0\"?>" +
		"<?ilo entity-processing=\"standard\"?>" +
		"<?xmlilo output-format=\"xml\"?>" +
		"<RIBCL VERSION=\"2.0\">" +
		"<LOGIN PASSWORD=\"" + c.Password + "\" USER_LOGIN=\"" + c.Username + "\">" +
		"<SERVER_INFO MODE=\"read\">" +
		"<GET_EMBEDDED_HEALTH>" +
		"</GET_EMBEDDED_HEALTH>" +
		"</SERVER_INFO>" +
		"</LOGIN>" +
		"</RIBCL>"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)

	return string(_body), nil
}
