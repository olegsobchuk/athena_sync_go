package athena

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiURL = "https://api.athenahealth.com"
)

// Connection api connection
type Connection struct {
	basePath     string
	clientID     string
	clientSecret string
	version      string
	Token        string
	practiceID   string
	client       http.Client
}

// New initialize connection
func (c *Connection) New(clientID, clientSecret, practiceID string) (err error) {
	c.basePath = apiURL
	c.clientID = clientID
	c.clientSecret = clientSecret
	c.practiceID = practiceID
	c.version = "preview1"
	c.client = http.Client{}

	err = c.auth()
	return err
}

func (c *Connection) auth() (err error) {
	tokenPath := strings.Join([]string{c.basePath, "oauthpreview", "token"}, "/")
	vals := url.Values{}
	vals.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", tokenPath, strings.NewReader(vals.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.clientID, c.clientSecret)

	response, err := c.call(req)
	if err != nil {
		return err
	}
	keys := response.(map[string]interface{})
	c.Token = keys["access_token"].(string)
	return nil
}

func (c *Connection) call(req *http.Request) (response interface{}, err error) {
	req.Header.Add("Accept-Encoding", "deflate")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil && resp.StatusCode != 400 { // HTTP 400 Not Authorized
		return nil, err
	}
	var decoded interface{}
	err = json.Unmarshal(body, &decoded)
	if err != nil {
		return nil, err
	}
	return decoded, err
}

// GET get request
func (c *Connection) GET(path string, params map[string]string) (response interface{}, err error) {
	vals := url.Values{}
	for k, v := range params {
		vals.Add(k, v)
	}
	targetURL := strings.Join([]string{c.basePath, c.version, c.practiceID, path}, "/")
	fmt.Println(targetURL)
	req, err := http.NewRequest("GET", targetURL, strings.NewReader(vals.Encode()))
	req.Header.Add("Authorization", "Bearer "+c.Token)
	if err != nil {
		return nil, err
	}
	response, err = c.call(req)
	return response, err
}
