package rest

import (
	"fmt"
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
