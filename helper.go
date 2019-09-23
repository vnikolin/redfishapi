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
func queryData(c *IloClient, call string, link string, data []byte) ([]byte, http.Header, int, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest(call, link, bytes.NewBuffer(data))
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) == true {
			err := errors.New(StatusInternalServerError)
			return _body, resp.Header, resp.StatusCode, err
		}
		return _body, resp.Header, resp.StatusCode, err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New(StatusUnauthorized)
			return _body, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 400 {
			err := errors.New(StatusBadRequest)
			return _body, resp.Header, resp.StatusCode, err
		}

	}
	defer resp.Body.Close()

	return _body, resp.Header, resp.StatusCode, nil
}
