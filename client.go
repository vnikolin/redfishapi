package redfishapi

// //IloClient ... Contstructor required Variables
// type IloClient struct {
// 	Hostname string
// 	Username string
// 	Password string
// }

// //NewIloClient ... Initializes the Constructor with the above variables
// func NewIloClient(hostname string, username string, password string) *IloClient {

// 	return &IloClient{
// 		Hostname: hostname,
// 		Username: username,
// 		Password: password,
// 	}
// }

// redfishProvider ... Contstructor required Variables
type redfishProvider struct {
	Hostname    string
	Username    string
	Password    string
	Certificate string
}

// NewRedfishProvider ... Initializes the Constructor with the above variables
func NewRedfishProvider(hostname string, username string, password string, certificate string, insecureSkipVerify bool) RedfishProvider {

	return &redfishProvider{
		Hostname:    hostname,
		Username:    username,
		Password:    password,
		Certificate: certificate,
	}
}
