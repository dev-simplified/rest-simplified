package rest

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	client httpClient = &http.Client{}
)

type netMockStruct struct{}

//httpClient extends Do method of net/http for mock testing
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func executeAPI(method string, apiURL string, headers []*headers, payload io.Reader) (int, string, error) {
	request, err := http.NewRequest(method, apiURL, payload)
	if err != nil {
		log.Panic(err)
		return 0, "", err
	}

	for i := 0; i < len(headers); i++ {
		header := headers[i]
		request.Header.Add(header.Key, header.Value)
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
		return 0, "", err
	}

	statusCode := response.StatusCode
	defer response.Body.Close()
	responseBody, _ := ioutil.ReadAll(response.Body)
	responseString := string(responseBody)

	return statusCode, responseString, err
}
