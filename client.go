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
