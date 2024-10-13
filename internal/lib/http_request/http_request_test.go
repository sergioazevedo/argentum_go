package http_request_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sergioazevedo/argentum_go/internal/lib/http_request"
	"github.com/stretchr/testify/assert"
)

func Test_Perform_SucessfullRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"key":"value"}`))
	}))
	defer server.Close()

	request, _ := http.NewRequest("GET", server.URL, nil)
	response, err := http_request.Perform(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotEmpty(t, response.Body)
}

func Test_Perform_FailedRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Some Error`))
	}))
	defer server.Close()

	request, _ := http.NewRequest("GET", server.URL, nil)
	response, err := http_request.Perform(request)

	assert.NotNil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 500, response.StatusCode)
	assert.NotEmpty(t, response.Body)
}
