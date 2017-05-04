package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"reflect"

	"github.com/new-black/eva-cli/messages"
)

type Client interface {
	Login(email string, password string, organizationUnitID int, applicationID int) (*messages.LoginResponse, error)
	GetApplications() (*messages.ListApplicationsResponse, error)
	GetCurrentUser() (*messages.User, error)
	Send(message interface{}, result interface{}) error
	StoreToken(token string) error
}

type HttpClient struct {
	Host string

	client *http.Client
}

func NewHttpClient(host string) *HttpClient {
	return &HttpClient{
		Host:   host,
		client: &http.Client{},
	}
}

func (c *HttpClient) GetApplications() (*messages.ListApplicationsResponse, error) {
	var result messages.ListApplicationsResponse

	res, err := c.client.Get(fmt.Sprintf("%s/api/v1/application", c.Host))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &result, json.NewDecoder(res.Body).Decode(&result)
}

func (c *HttpClient) Login(email string, password string, organizationUnitID int, applicationID int) (*messages.LoginResponse, error) {
	var result messages.LoginResponse

	message := messages.Login{
		EmailAddress:       email,
		Password:           password,
		OrganizationUnitID: organizationUnitID,
		ApplicationID:      applicationID,
	}

	return &result, c.Send(message, &result)
}

func (c *HttpClient) GetCurrentUser() (*messages.User, error) {
	var result messages.GetCurrentUserResponse

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/user/current", c.Host), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if pathToToken, err := getPathToToken(); err == nil {
		token, err := ioutil.ReadFile(pathToToken)

		if err == nil {
			req.Header.Set("Authorization", string(token))
		}
	}

	res, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized: %s", res.Status)
	}

	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result.User, nil
}

func (c *HttpClient) Send(message interface{}, result interface{}) error {
	b := &bytes.Buffer{}

	json.NewEncoder(b).Encode(message)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/message/%s", c.Host, reflect.TypeOf(message).Name()), b)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if pathToToken, err := getPathToToken(); err == nil {
		token, err := ioutil.ReadFile(pathToToken)

		if err == nil {
			req.Header.Set("Authorization", string(token))
		}
	}

	res, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get response %d", res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(&result)
}

func (c *HttpClient) StoreToken(token string) error {
	pathToToken, err := getPathToToken()

	if err != nil {
		return err
	}

	return ioutil.WriteFile(pathToToken, []byte(token), 0644)
}

func getPathToToken() (string, error) {
	u, err := user.Current()

	if err != nil {
		return "", err
	}

	pathToDir := filepath.Join(u.HomeDir, ".eva-cli")

	if err := os.MkdirAll(pathToDir, os.ModePerm); err != nil {
		return "", err
	}

	pathToToken := filepath.Join(pathToDir, ".token")

	return pathToToken, nil
}
