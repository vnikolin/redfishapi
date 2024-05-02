package redfishapi

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// basicAuth ... will create the basicauth encoded string for the credentials
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// queryData ... will make REST verbs based on the url
func queryData(c *redfishProvider, call string, link string, data []byte) ([]byte, http.Header, int, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest(call, link, bytes.NewBuffer(data))
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Second * 300,
	}
	resp, err := client.Do(req)
	// ablakmak... try to test status code here
	fmt.Println("ablakmak resp.StatusCode: ", resp.StatusCode)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) == true {
			err := errors.New(StatusInternalServerError)
			return nil, nil, 0, err
		}
		return nil, nil, 0, err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New(StatusUnauthorized)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 400 {
			err := errors.New(StatusBadRequest)
			return nil, resp.Header, resp.StatusCode, err
		}

	}
	defer resp.Body.Close()

	_body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, resp.StatusCode, err
	}

	return _body, resp.Header, resp.StatusCode, nil
}
