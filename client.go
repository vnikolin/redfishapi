package hpilo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
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
		UUID:        x.UUID,
		ServerModel: x.SPN,
		Serial:      x.SBSN,
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
