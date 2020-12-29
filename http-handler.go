package rest

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	client HTTPClient = &http.Client{}
)

//HTTPClient extends Do method of net/http for mock testing
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func executeAPI(method string, apiURL string, headers []*Headers, payload io.Reader) (int, string, error) {
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

	statusCode := response.StatusCode
	defer response.Body.Close()
	responseBody, _ := ioutil.ReadAll(response.Body)
	responseString := string(responseBody)

	return statusCode, responseString, err
}
