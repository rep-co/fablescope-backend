package iamgenerator

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: when it will be decided how to deploy refactor this mess
// You are entering nasty dirty code
// I promise to refactor when i will understand what exactly do i need

type IAMTokenGenerator interface {
	GenerateToken() (string, error)
}

const (
	protocol      = "https://"
	apiGatewayURL = ".apigw.yandexcloud.net"
	serverlessURL = "http://169.254.169.254/computeMetadata/v1/instance/service-accounts/default/token"
)

type IAMTokenHTTPResponse struct {
	AccessToken string `json:"access_token"`
}

type IAMTokenHTTP struct {
	client       *http.Client
	apiGatewayID string
}

func NewIAMTokenHTTP(apiGatewayID string) *IAMTokenHTTP {
	return &IAMTokenHTTP{
		client:       &http.Client{},
		apiGatewayID: apiGatewayID,
	}
}

func (iam *IAMTokenHTTP) GenerateToken() (string, error) {
	request, err := iam.newRequest()
	if err != nil {
		return "", err
	}
	response := IAMTokenHTTPResponse{}
	err = iam.sendRequest(request, &response)
	if err != nil {
		return "", err
	}

	return response.AccessToken, nil
}

func (iam *IAMTokenHTTP) newRequest() (*http.Request, error) {
	url := fmt.Sprintf("%s%s%s/iam", protocol, iam.apiGatewayID, apiGatewayURL)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (iam *IAMTokenHTTP) sendRequest(request *http.Request, v any) error {
	response, err := iam.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(v)
}

type IAMTokenServerless struct {
	client *http.Client
}

func NewIAMTokenServerless() *IAMTokenHTTP {
	return &IAMTokenHTTP{
		client: &http.Client{},
	}
}

func (iam *IAMTokenServerless) GenerateToken() (string, error) {
	request, err := iam.newRequest()
	if err != nil {
		return "", err
	}
	response := IAMTokenHTTPResponse{}
	err = iam.sendRequest(request, &response)
	if err != nil {
		return "", err
	}

	return response.AccessToken, nil
}

func (iam *IAMTokenServerless) newRequest() (*http.Request, error) {
	url := serverlessURL
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (iam *IAMTokenServerless) sendRequest(request *http.Request, v any) error {
	request.Header.Set("Metadata-Flavor", "Google")

	response, err := iam.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(v)
}
