package rest

import (
	"net/url"
)

type Headers struct {
	Key   string
	Value string
}

type APIClient struct {
	APIURL           string
	CustomAPIHeaders []*Headers
	ContentType      string
	Authorization    string
	APIType          string
	AuthType         string
	AuthToken        string
}

func CreateBearerAuth(token string) string {
	authToken := "Bearer " + token
	return authToken
}

func CreateBasicAuth(userName string, password string) string {
	authToken := "Basic " + userName + ":" + url.QueryEscape(password)
	return authToken
}
