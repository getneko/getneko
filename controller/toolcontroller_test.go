package controller_test

import (
	"getneko/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChackLowVersion(t *testing.T) {
	router := router.Router()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/chacklowversion", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"code":0,"message":"","data":"1.0"}`, w.Body.String())
}
