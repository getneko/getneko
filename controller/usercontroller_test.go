package controller_test

import (
	"encoding/json"
	"getneko/router"
	"getneko/structtypes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUerregcontroller(t *testing.T) {
	router := router.Router()
	//正常注册
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/userreg", strings.NewReader(`{
		"email": "test@gamil.com",
		"language": "en-US",
		"password": "test",
		"username": "test"
	  }`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var res structtypes.JSONResult
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, 0, res.Code)
	assert.Equal(t, "", res.Message)
	//用户名已经存在
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/v1/userreg", strings.NewReader(`{
		"email": "test@gamil.com",
		"language": "en-US",
		"password": "test",
		"username": "test"
	  }`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, -2, res.Code)
	assert.Equal(t, "Username or email already exists", res.Message)
	//参数不符合规范
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/v1/userreg", strings.NewReader(`{
		
	  }`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, -1, res.Code)
	assert.Equal(t, "illegal request", res.Message)
}

func TestUerLogincontroller(t *testing.T) {
	router := router.Router()
	//正常登录
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/userlogin", strings.NewReader(`{
		"language": "en-US",
		"password": "test",
		"username": "test"
	  }`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var res structtypes.JSONResult
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, 0, res.Code)
	assert.Equal(t, "", res.Message)
	//密码错误
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/v1/userlogin", strings.NewReader(`{
		"language": "en-US",
		"password": "test",
		"username": "test1"
	  }`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, -3, res.Code)
	assert.Equal(t, "Username does not exist or password is wrong", res.Message)
	//参数不符合规范
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/v1/userlogin", strings.NewReader(`{
		
	  }`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, -1, res.Code)
	assert.Equal(t, "illegal request", res.Message)
}
