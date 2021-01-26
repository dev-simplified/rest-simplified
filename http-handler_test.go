package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestHTTPExecuteAPI(t *testing.T) {
	fmt.Println("running TestexecuteAPI")
	mockClient := &netMockStruct{}
	client = mockClient
	var mockHeaders []*headers
	statusCode, response, err := executeAPI("GET", "http://url", mockHeaders, bytes.NewBuffer([]byte(`payload`)))
	fmt.Println(statusCode)
	fmt.Println(response)
	fmt.Println(err)
}

func (client *netMockStruct) Do(req *http.Request) (*http.Response, error) {
	var response *http.Response = new(http.Response)
	var err error

	response.StatusCode = 200
	response.Body = ioutil.NopCloser(strings.NewReader(`{"message": "success"}`))
	return response, err
}

func (client *netMockStruct) run() {
	fmt.Println("run works")
}
