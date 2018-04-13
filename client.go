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
		return nil, err
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
