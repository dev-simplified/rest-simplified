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
