package jamf_pro_go

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	APIPathClassic = "/JSSResource/"
	APIPathV1      = "/uapi/"
)

type Config struct {
	BaseURL          string
	Log              Logger
	v1ApiToken       string
	classicApiToken  string
}

// V1Token is the response to the Jamf API Token request.
type V1Token struct {
	Token    string `json:"token"`
	Expires  uint64 `json:"expires"`
}

func NewConfig(url, userName, password string) (*Config, error) {
	if len(url) == 0 {
		return nil, errors.New("[Err] missing URL")
	}

	if len(userName) == 0 {
		return nil, errors.New("[Err] missing username")
	}

	if len(password) == 0 {
		return nil, errors.New("[Err] missing password")
	}

	var config Config

	config.BaseURL = url

	// generate Jamf Pro Classic API Token
	credentials := []byte(userName + ":" + password)
	encodedCredentials := base64.StdEncoding.EncodeToString(credentials)
	config.classicApiToken = encodedCredentials

	// request Jamf Pro API Token
	req, err := http.NewRequest(http.MethodPost, config.BaseURL + "/uapi/auth/tokens", nil)
	req.Header.Set("Authorization", "Basic " + encodedCredentials)
	if err != nil{
		fmt.Println("[Err] ", err.Error())
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[Err] ", err.Error())
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return &config, errors.New("[Err] Request Jamf Pro API Token: HTTP Status is " + resp.Status)
	}

	var r io.Reader = resp.Body
	var v1Token V1Token
	err = json.NewDecoder(r).Decode(&v1Token)
	if err != nil {
		fmt.Println("[Err] ", err.Error())
	}
	config.v1ApiToken = v1Token.Token
	return &config, nil
}


type Client struct {
	httpClient *http.Client
	config *Config
}

func NewClient(config *Config) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		config: config,
	}
}


func (c *Client) call(apiPath, method, apiVersion string,
	queryParams url.Values, postBody interface{}, res interface{},
) error {

	var (
		contentType string
		body        io.Reader
	)

	if method != http.MethodDelete {
		if apiVersion == "v1" {
			contentType = "application/json"
			jsonParams, err := json.Marshal(postBody)
			if err != nil {
				return err
			}
			body = bytes.NewBuffer(jsonParams)
		} else if apiVersion == "classic" {
			contentType = "application/xml"
			xmlParams, err := xml.Marshal(postBody)
			if err != nil {
				return err
			}
			body = bytes.NewBuffer(xmlParams)
		}
	}

	req, err := c.newRequest(apiPath, method, contentType, apiVersion, queryParams, body)
	if err != nil {
		return err
	}
	return c.do(req, apiVersion, res)
}


func (c *Client) newRequest(
	apiPath, method, contentType, apiVersion string,
	queryParams url.Values,
	body io.Reader,
) (*http.Request, error) {

	// construct url
	u, err := url.Parse(c.config.BaseURL)
	if err != nil {
		return nil, err
	}
	if apiVersion == "v1" {
		u.Path = path.Join(u.Path, APIPathV1, apiPath)
	} else if apiVersion == "classic" {
		u.Path = path.Join(u.Path, APIPathClassic, apiPath)
	}

	u.RawQuery = queryParams.Encode()
	// request with context
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	//req = req.WithContext(ctx)

	// set http headers
	if apiVersion == "v1" {
		req.Header.Set("Authorization", "Bearer " + c.config.v1ApiToken)
	} else if apiVersion == "classic" {
		req.Header.Set("Authorization", "Basic " + c.config.classicApiToken)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Accept", contentType)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, apiVersion string, res interface{}) error {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	c.logf("[jamf-pro-go] %s: %v %v%v", response.Status, req.Method, req.URL.Host, req.URL.Path)

	var r io.Reader = response.Body
	// r = io.TeeReader(r, os.Stderr)

	// parse Jamf Pro (classic) API errors
	code := response.StatusCode
	if code >= http.StatusBadRequest {
		byt, err := ioutil.ReadAll(r)
		if err != nil {
			// error occured, but ignored.
			c.logf("[jamf-pro-go] HTTP response body: %v", err)
		}
		res := &Error{
			StatusCode: code,
			RawError:   string(byt),
		}
		return res
	}

	if apiVersion == "v1" {
		return json.NewDecoder(r).Decode(&res)
	} else if apiVersion == "classic" {
		return xml.NewDecoder(r).Decode(&res)
	}

	return errors.New("[jamf-pro-go] apiVersion value is invalid")
}
