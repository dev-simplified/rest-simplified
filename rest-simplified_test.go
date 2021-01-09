package rest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExecuteAPI(t *testing.T) {
	fmt.Println("executing test")
	apiClient := CreateAPIClient("https://api.github.com/orgs/dev-simplified/repos", "GET", "none", "application/json")
	status, response, err := apiClient.ExecuteAPI("")
	fmt.Println(status)
	fmt.Println(response)
	fmt.Println(err)
}

func TestCreateBearerAuth(t *testing.T) {
	t.Log("testing bearer auth creation")
	auth := CreateBearerAuth("token")
	t.Log("auth token created is : " + auth)
	if auth != "Bearer token" {
		t.Error("expected 'Bearer token' but got : " + auth)
	}
}

func TestCreateBasicAuth(t *testing.T) {
	auth := ""

	t.Log("testing first set")
	auth = CreateBasicAuth("jerry", "password")
	if auth != `Basic jerry:password` {
		t.Error("expected Basic jerry:password but got " + auth)
	}

	t.Log("testing second set")
	auth = CreateBasicAuth("jerry@domain.com", "password")
	if auth != `Basic jerry%40domain.com:password` {
		t.Error("expected Basic jerry%40domain.com:password but got " + auth)
	}

	t.Log("testing third set")
	auth = CreateBasicAuth("jerry", "password@123")
	if auth != `Basic jerry:password%40123` {
		t.Error("expected Basic jerry:password%40123 but got " + auth)
	}

	t.Log("testing fourth set")
	auth = CreateBasicAuth("jerry@domain.com", "password@123")
	if auth != `Basic jerry%40domain.com:password%40123` {
		t.Error("expected Basic jerry%40domain.com:password%40123 but got " + auth)
	}
}

func TestCreateAPIClient(t *testing.T) {
	apiClient := CreateAPIClient("url", "GET", "Basic usr:pwd", "application/json")
	clientType := reflect.TypeOf(apiClient)
	t.Log(clientType.String())
	if clientType.String() != `*rest.APIClient` {
		t.Error("expected *rest.APIClient but got " + clientType.String())
	}
}

func TestCreateMockClient(t *testing.T) {
	EnableMock()
	if !enableMock {
		t.Log("expected mock to be enabled")
	} else {
		t.Log("mock enabled")
	}
	apiClient := CreateAPIClient("url", "GET", "Basic usr:pwd", "application/json")
	clientType := reflect.TypeOf(apiClient)
	t.Log(clientType.String())
	if clientType.String() != `*rest.APIMockClient` {
		t.Error("expected *rest.APIMockClient but got " + clientType.String())
	}
}
