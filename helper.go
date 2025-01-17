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

// Initialize logger
// var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

// basicAuth ... will create the basicauth encoded string for the credentials
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// queryData ... will make REST verbs based on the url
func queryData(c *redfishProvider, call string, link string, data []byte) ([]byte, http.Header, int, error) {

	var err error
	for i := 0; i < 2; i++ {
		if c.Certificate != "" && i == 0 {
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
		client.CloseIdleConnections()
		if err != nil {
			r, _ := regexp.Compile("dial tcp")
			if r.MatchString(err.Error()) {
				err := errors.New(StatusUnreachable)
				return nil, nil, 0, err
			}
			// fmt.Printf("errored out in queryData for %s\n", link)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			if resp.StatusCode == 401 {
				err := errors.New(StatusUnauthorized)
				return nil, resp.Header, resp.StatusCode, err
			} else if resp.StatusCode == 400 {
				err := errors.New(StatusBadRequest)
				return nil, resp.Header, resp.StatusCode, err
			}

		}

		_body, _ := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, resp.StatusCode, err
		}

		return _body, resp.Header, resp.StatusCode, nil
	}
	return nil, nil, 0, err
}

// queryDataForce ... will make REST verbs based on the url without retrying
func queryDataForce(c *redfishProvider, call string, link string, data []byte) ([]byte, http.Header, int, error) {

	if c.Certificate != "" {
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM([]byte(c.Certificate))
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false, RootCAs: certPool}
		// fmt.Printf("testing queryDataForce with certificate for %s\n", link)
	} else {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		// fmt.Printf("testing queryDataForce without certificate for %s\n", link)
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
	client.CloseIdleConnections()
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) {
			err := errors.New(StatusUnreachable)
			return nil, nil, 0, err
		}
		return nil, nil, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err := errors.New(StatusUnauthorized)
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 400 {
			err := errors.New(StatusBadRequest)
			return nil, resp.Header, resp.StatusCode, err
		}

	}

	_body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, resp.StatusCode, err
	}

	return _body, resp.Header, resp.StatusCode, nil
}

// postForm ... will make REST POST request with form data
func postForm(c *redfishProvider, link string, form *bytes.Buffer, contentType string) ([]byte, http.Header, int, error) {

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
	client.CloseIdleConnections()
	if err != nil {
		r, _ := regexp.Compile("dial tcp")
		if r.MatchString(err.Error()) {
			err := fmt.Errorf("postForm: %s", strings.ToLower(StatusInternalServerError))
			return nil, nil, 0, err
		}
		return nil, nil, 0, err
	}

	// fmt.Printf("Response: %+v\n", resp)
	defer resp.Body.Close()
	if resp.StatusCode != 202 {
		if resp.StatusCode == 401 {
			err := fmt.Errorf("postForm: %s", strings.ToLower(StatusUnauthorized))
			return nil, resp.Header, resp.StatusCode, err
		} else if resp.StatusCode == 400 {
			err := fmt.Errorf("postForm: %s", strings.ToLower(StatusBadRequest))
			return nil, resp.Header, resp.StatusCode, err
		}

	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, resp.StatusCode, err
	}

	return respBody, resp.Header, resp.StatusCode, nil
}
