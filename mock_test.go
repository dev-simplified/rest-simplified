package rest

import (
	"fmt"
	"testing"
)

func TestCreateAPIMockClient(t *testing.T) {
	fmt.Println("testing API mock client creation")
	mockClient := createAPIMockClient()
	fmt.Println(mockClient)
}

func TestEnableMock(t *testing.T) {
	EnableMock()
	if !enableMock {
		t.Error("expected mock to be enabled")
	} else {
		fmt.Println("mock enabled")
	}
}

func TestCreateMockResponse(t *testing.T) {
	CreateMockResponse(300, "testing", nil)
	if mockResponseCode != 300 {
		t.Error("expected response code to be 300 but got " + fmt.Sprint(mockResponseCode))
	}
	if mockRresponse != "testing" {
		t.Error("expected testing as response, but got " + mockRresponse)
	}
	if mockErr != nil {
		t.Error("expected no error, but got" + mockErr.Error())
	}
}

func TestMockExecuteAPI(t *testing.T) {
	client := &apiMockClient{}
	responseCode, response, err := client.ExecuteAPI("payload")
	if responseCode != 300 {
		t.Error("expected response code to be 300 but got " + fmt.Sprint(responseCode))
	}
	if response != "testing" {
		t.Error(`expected testing as response, but got ` + response)
	}
	if err != nil {
		t.Error("expected no error, but got" + err.Error())
	}
}

func TestMockAddAdditionalRequestHeader(t *testing.T) {
	fmt.Println("no validation test")
}
