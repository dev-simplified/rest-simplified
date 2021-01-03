# rest-simplified

Simplified rest API client created as a wrapper over net/http. This can be used in other Golang projects for accessing rest APIs.

## usage

### To download the package

    go get github.com/dev-simplified/rest-simplified

### How to use the package

#### import the package

```go
import rest "github.com/dev-simplified/rest-simplified"
```

#### create authorization if required

```go
//if basic auth is to be used
auth := rest.CreateBasicAuth("username", "password")
//if bearer token is to be used
auth := rest.CreateBearerAuth("bearer token")
//for no auth
auth := "" or auth := "none"
```

#### create API client

```go
auth := "authorization created in previous step"
apiClient := rest.CreateAPIClient("http://api-url", "GET/POST/PUT/...", auth, "application/json")
```

#### Add additional request headers if needed

```go
//the below line adds an additional header (Accept in this case) to the headers handled by default during client creation (Authorization, Content-type) 
apiClient = apiclient.AddAdditionalRequestHeader("Accept", "application/json")
```

#### Execute the API

```go
payload := `json string as payload or request body` 

payload = "" //in case of no payload

//responseCode is the http response code (200/400/...)
//responseBody is the api response in string format
//err is any error that may have occured during API execution 
responseCode, responseBody, err := apiClient.ExecuteAPI(payload)

```

### Mock Testing

The package comes with a mock testing framework built into it. Mock tests can be enabled and custom responses can be created using the package methods. Below are the steps to perform API mock tests using rest-simplified.

#### Enable Mock

```go
//This method needs to be used inside a unit test function
status := rest.EnableMock()
//status is simply a string confirming the switch over to mock client
```

#### Create Custom Mock Response

```go
status = rest.CreateMockResponse(200, `{mock response}`, nil)
//status is a string confirming mock response creation
//CreateMockResponse takes mockResponseCode, mockResponseBody and mockError as input parameters
```

#### Call the API function to test

Now the rest API function to be unit tested can be executed from the test method and the response returned will be the mock response created.
