package http_request

import (
	"fmt"
	"net/http"
)

func Perform(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		return resp, fmt.Errorf("%s", resp.Body)
	}

	return resp, nil
}
