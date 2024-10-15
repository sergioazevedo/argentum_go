package servertest

import (
	"net/http"
	"net/http/httptest"
)

func NewTestServer(statusCode int, responseData string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(responseData))
	}))

	return server
}
