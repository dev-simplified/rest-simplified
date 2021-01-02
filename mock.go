package rest

var (
	mockResponseCode int    = 200
	mockRresponse    string = `{"message": "success"}`
	mockErr          error  = nil
)

//APIMockClient is used to mock API calls
type APIMockClient struct{}

func createAPIMockClient() *APIMockClient {
	apiMockClient := &APIMockClient{}
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
func (*APIMockClient) ExecuteAPI(APIPayload string) (responseCode int, responseBody string, err error) {
	return mockResponseCode, mockRresponse, mockErr
}