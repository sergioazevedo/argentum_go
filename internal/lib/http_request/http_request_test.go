package http_request_test

import (
	"net/http"
	"testing"

	"github.com/sergioazevedo/argentum_go/internal/lib/http_request"
	"github.com/sergioazevedo/argentum_go/internal/lib/servertest"
	"github.com/stretchr/testify/assert"
)

func Test_Perform_SucessfullRequest(t *testing.T) {
	server := servertest.NewTestServer(http.StatusOK, `{"key":"value"}`)
	defer server.Close()

	request, _ := http.NewRequest("GET", server.URL, nil)
	response, err := http_request.Perform(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 200, response.StatusCode)
	assert.NotEmpty(t, response.Body)
}

func Test_Perform_FailedRequest(t *testing.T) {
	server := servertest.NewTestServer(http.StatusInternalServerError, `Some Error`)
	defer server.Close()

	request, _ := http.NewRequest("GET", server.URL, nil)
	response, err := http_request.Perform(request)

	assert.NotNil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 500, response.StatusCode)
	assert.NotEmpty(t, response.Body)
}
