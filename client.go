package redfishapi

//IloClient ... Contstructor required Variables
type IloClient struct {
	Hostname string
	Username string
	Password string
}

//NewIloClient ... Initializes the Constructor with the above variables
func NewIloClient(hostname string, username string, password string) *IloClient {

	return &IloClient{
		Hostname: hostname,
		Username: username,
		Password: password,
	}
}
