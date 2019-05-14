package redfishapi

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

//basicAuth ... will create the basicauth encoded string for the credentials
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//queryData ... will make REST verbs based on the url
func queryData(c *IloClient, call string, link string, data []byte) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest(call, link, bytes.NewBuffer(data))
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) == true {
			err := errors.New(StatusInternalServerError)
			return nil, nil, err
		}
		return nil, nil, err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New(StatusUnauthorized)
			return nil, nil, err
		} else if resp.StatusCode == 400 {
			err := errors.New(StatusBadRequest)
			return nil, nil, err
		}

	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return nil, _body, nil
}
