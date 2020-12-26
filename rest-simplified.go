package rest

import (
	"bytes"
	"io"
	"net/url"
)

type Headers struct {
	Key   string
	Value string
}

type APIClient struct {
	APIURL               string
	AdditionalAPIHeaders []*Headers
	ContentType          string
	Authorization        string
	APIMethod            string
}

func CreateBearerAuth(token string) string {
	authToken := "Bearer " + token
	return authToken
}

func CreateBasicAuth(userName string, password string) string {
	authToken := "Basic " + userName + ":" + url.QueryEscape(password)
	return authToken
}

func CreateAPIClient(apiUrl string, apiMethod string, authorization string, contentType string) *APIClient {
	apiClient := &APIClient{}
	apiClient.APIURL = apiUrl
	apiClient.APIMethod = apiMethod
	apiClient.Authorization = authorization
	apiClient.ContentType = contentType
	return apiClient
}

func (client *APIClient) AddAdditionalRequestHeader(headerName string, headerValue string) *APIClient {
	header := &Headers{}
	header.Key = headerName
	header.Value = headerValue
	clientAPIHeaders := client.AdditionalAPIHeaders
	clientAPIHeaders = append(clientAPIHeaders, header)
	client.AdditionalAPIHeaders = clientAPIHeaders
	return client
}

func (client *APIClient) ExecuteAPI(JSONPayload string) (int, string, error) {
	var apiMethod string
	var apiURL string
	var contentType string
	var authorization string
	var payload io.Reader = nil

	apiMethod = client.APIMethod
	apiURL = client.APIURL
	contentType = client.ContentType
	authorization = client.Authorization
	headers := client.AdditionalAPIHeaders

	contentTypeHeader := &Headers{}
	contentTypeHeader.Key = "content-type"
	contentTypeHeader.Value = contentType
	headers = append(headers, contentTypeHeader)

	if authorization != "none" || authorization != "" {
		authorizationHeader := &Headers{}
		authorizationHeader.Key = "Authorization"
		authorizationHeader.Value = authorization
		headers = append(headers, authorizationHeader)
	}

	if JSONPayload != "" {
		payload = bytes.NewBuffer([]byte(JSONPayload))
	}

	responseCode, responseBody, err := executeAPI(apiMethod, apiURL, headers, payload)

	return responseCode, responseBody, err
}
