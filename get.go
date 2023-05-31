package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	requestURLs := []string{"http://www.amazon.com", "http://www.google.com"}

	for _, requestURL := range requestURLs {
		fmt.Printf("client: GET REQUEST TO %s\n", requestURL)

		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			fmt.Printf("client: could not create request: %s\n", err)
			os.Exit(1)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("client: got response!\n")
		fmt.Printf("client: status code: %d\n", res.StatusCode)

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("client: could not read response body: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("client: response body length: %d\n", len(resBody))
	}
}
