package rest

import (
	"bytes"
	"io"
	"net/url"
)

var (
	enableMock bool = false
)

//headers is a key value pair for rest API headers
type headers struct {
	//Key is the name of the request header
	//Value is the value corresponding to the request header key
	Key   string
	Value string
}

//APIClientInterface allows for mocking of ExecuteAPI method
type APIClientInterface interface {
	ExecuteAPI(APIPayload string) (responseCode int, responseBody string, err error)
	AddAdditionalRequestHeader(headerName string, headerValue string) APIClientInterface
}

//APIClient holds all request properties needed for making rest API calls
type APIClient struct {
	//APIURL: the URL of the rest APi to be invoked
	//APIMethod: API method to be used (GET/ POST/ PUT)
	//ContentType: The Content-type (eg application/json) of the data being used for invoking the rest API. Currently support has been created for application/json
	//Authorization: Basic Auth or Bearer Auth token created for API Authorization (use CreateBearerAuth() or CreateBasicAuth() methods)
	//AdditionalAOIHeaders: Array of additional headers other than Authorization and Content-Type
	APIURL               string
	AdditionalAPIHeaders []*headers
	ContentType          string
	Authorization        string
	APIMethod            string
}

//CreateBearerAuth accepts token string as input and returns bearer token to be used for rest API access
func CreateBearerAuth(token string) string {
	authToken := "Bearer " + token
	return authToken
}

//CreateBasicAuth accepts user name and password for API authentication and returns Basic auth string
func CreateBasicAuth(userName string, password string) string {
	authToken := "Basic " + url.QueryEscape(userName) + ":" + url.QueryEscape(password)
	return authToken
}

//CreateAPIClient accepts the apiURL, apiMethod (GET/POST/etc), authorization token and contentType(application/json) and creates the client which can call the API.
//Authorization token can be created using CreateBasicAuth and CreateBearerAuth functions
func CreateAPIClient(apiURL string, apiMethod string, authorization string, contentType string) APIClientInterface {
	var apiClient APIClientInterface
	if !enableMock {
		apiClient = createAPIAccessClient(apiURL, apiMethod, authorization, contentType)
	} else {
		apiClient = createAPIMockClient()
	}
	return apiClient
}

func createAPIAccessClient(apiURL string, apiMethod string, authorization string, contentType string) *APIClient {
	apiClient := &APIClient{}
	apiClient.APIURL = apiURL
	apiClient.APIMethod = apiMethod
	apiClient.Authorization = authorization
	apiClient.ContentType = contentType
	return apiClient
}

//AddAdditionalRequestHeader method takes the header name and header value as input and appends these headers to Content-type and Authorization headers created by default.
//Example headerName: Accept, headerValue: application/json
func (client *APIClient) AddAdditionalRequestHeader(headerName string, headerValue string) APIClientInterface {
	header := &headers{}
	header.Key = headerName
	header.Value = headerValue
	clientAPIHeaders := client.AdditionalAPIHeaders
	clientAPIHeaders = append(clientAPIHeaders, header)
	client.AdditionalAPIHeaders = clientAPIHeaders
	return client
}

//ExecuteAPI method executes the api client (client contains all api details like URL, method,etc) with the APIPayload passed as input and returns the API response code, response body as a string and any error that may have occured.
func (client *APIClient) ExecuteAPI(APIPayload string) (responseCode int, responseBody string, err error) {
	var apiMethod string
	var apiURL string
	var contentType string
	var authorization string
	var payload io.Reader = nil

	apiMethod = client.APIMethod
	apiURL = client.APIURL
	contentType = client.ContentType
	authorization = client.Authorization
	apiHeaders := client.AdditionalAPIHeaders

	contentTypeHeader := &headers{}
	contentTypeHeader.Key = "Content-type"
	contentTypeHeader.Value = contentType
	apiHeaders = append(apiHeaders, contentTypeHeader)

	if authorization != "none" && authorization != "" {
		authorizationHeader := &headers{}
		authorizationHeader.Key = "Authorization"
		authorizationHeader.Value = authorization
		apiHeaders = append(apiHeaders, authorizationHeader)
	}

	if APIPayload != "" {
		payload = bytes.NewBuffer([]byte(APIPayload))
	}

	responseCode, responseBody, err = executeAPI(apiMethod, apiURL, apiHeaders, payload)

	return responseCode, responseBody, err
}
