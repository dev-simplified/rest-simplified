package rest

import (
	"bytes"
	"io"
	"net/url"
)

//Headers is a key value pair for rest API headers
type Headers struct {
	//Key is the name of the request header
	//Value is the value corresponding to the request header key
	Key   string
	Value string
}

//APIClient holds all request properties needed for making rest API calls
type APIClient struct {
	//APIURL: the URL of the rest APi to be invoked
	//APIMethod: API method to be used (GET/ POST/ PUT)
	//ContentType: The Content-type (eg application/json) of the data being used for invoking the rest API. Currently support has been created for application/json
	//Authorization: Basic Auth or Bearer Auth token created for API Authorization (use CreateBearerAuth() or CreateBasicAuth() methods)
	//AdditionalAOIHeaders: Array of additional headers other than Authorization and Content-Type
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
	contentTypeHeader.Key = "Content-type"
	contentTypeHeader.Value = contentType
	headers = append(headers, contentTypeHeader)

	if authorization != "none" && authorization != "" {
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
