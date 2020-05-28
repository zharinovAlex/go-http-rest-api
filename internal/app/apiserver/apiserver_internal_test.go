package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleStatusCheck(t *testing.T) {
	server := New(NewConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/status", nil)

	server.handleStatusCheck().ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "OK")
}
