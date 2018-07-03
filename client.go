package redfishapi

import (
	"dell"
	"hp"
)

type IloClient struct {
	Hostname string
	Username string
	Password string
}

func NewIloClient(hostname string, username string, password string) *IloClient {

	return &IloClient{
		Hostname: hostname,
		Username: username,
		Password: password,
	}
}

func (c *IloClient) GetMacAddress(make string) (string, error) {

	if make == "dell" {
		return dell.GetMacAddress()

	} else if make == "hp" {
		return hp.GetMacAddress()

	}
}
