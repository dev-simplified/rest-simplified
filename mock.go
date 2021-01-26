package rest

var (
	mockResponseCode int    = 200
	mockRresponse    string = `{"message": "success"}`
	mockErr          error  = nil
)

//apiMockClient is used to mock API calls
type apiMockClient struct{}

func createAPIMockClient() *apiMockClient {
	apiMockClient := &apiMockClient{}
	return apiMockClient
}

//EnableMock activates mock client creation for unit testing
func EnableMock() string {
	enableMock = true
	return "switching over to mock client for execution"
}

//CreateMockResponse creates a custom response to be returned during mock tests
func CreateMockResponse(responseCode int, response string, err error) string {
	mockResponseCode = responseCode
	mockRresponse = response
	mockErr = err
	return "mock response created"
}

//ExecuteAPI method here will provide a mock API response
func (client *apiMockClient) ExecuteAPI(APIPayload string) (responseCode int, responseBody string, err error) {
	return mockResponseCode, mockRresponse, mockErr
}

//AddAdditionalRequestHeader is for mocking the actual method
func (client *apiMockClient) AddAdditionalRequestHeader(headerName string, headerValue string) APIClientInterface {
	return client
}
