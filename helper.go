package redfishapi

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// basicAuth ... will create the basicauth encoded string for the credentials
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// queryData ... will make REST verbs based on the url
func queryData(c *redfishProvider, call string, link string, data []byte) ([]byte, http.Header, int, error) {
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// add certificat check logic here
	if c.Certificate != "" {
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM([]byte(c.Certificate))
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false, RootCAs: certPool}
	} else {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	req, err := http.NewRequest(call, link, bytes.NewBuffer(data))
	if err != nil {
		return nil, nil, 0, err
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Second * 300,
	}
	resp, err := client.Do(req)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) {
			err := errors.New(StatusUnreachable)
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

// postForm ... will make REST POST request with form data
func postForm(c *redfishProvider, link string, form *bytes.Buffer, contentType string) ([]byte, http.Header, int, error) {
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// add certificat check logic here
	if c.Certificate != "" {
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM([]byte(c.Certificate))
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false, RootCAs: certPool}
	} else {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	req, err := http.NewRequest("POST", link, form)
	if err != nil {
		return nil, nil, 0, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

	client := &http.Client{
		Timeout: time.Second * 300,
	}
	resp, err := client.Do(req)
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) {
			err := fmt.Errorf("postForm: %s", strings.ToLower(StatusInternalServerError))
			return nil, nil, 0, err
		}
		return nil, nil, 0, err
	}

	// fmt.Printf("Response: %+v\n", resp)

	if resp.StatusCode != 202 {
		if resp.StatusCode == 401 {
			err := fmt.Errorf("postForm: %s", strings.ToLower(StatusUnauthorized))
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 400 {
			err := fmt.Errorf("postForm: %s", strings.ToLower(StatusBadRequest))
			return nil, resp.Header, resp.StatusCode, err
		}

	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, resp.StatusCode, err
	}

	return respBody, resp.Header, resp.StatusCode, nil
}
