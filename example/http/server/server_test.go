package server

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWenHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, "HELLO WORLD", string(body))
}
