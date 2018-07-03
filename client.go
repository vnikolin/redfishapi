package redfishapi

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
		return GetMacAddressDell(), nil

	} else if make == "hp" {
		return GetMacAddressHp(), nil

	}
}
