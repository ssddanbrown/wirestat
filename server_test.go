package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpointResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	result := makeRequest(req, Response{}, "")
	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
}

func TestAlertingEndpointResponse(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	result := makeRequest(req, Response{
		Alerts: []string{"Server is angry"},
	}, "")
	assert.Equal(t, 500, result.StatusCode)
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
}

func TestAccessKeyRequired(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	result := makeRequest(req, Response{}, "donkey")
	body, _ := ioutil.ReadAll(result.Body)
	assert.Equal(t, 401, result.StatusCode)
	assert.Equal(t, "Invalid access key provided", string(body))
}

func TestAccessKeyViaQuery(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?key=donkey", nil)
	result := makeRequest(req, Response{}, "donkey")
	assert.Equal(t, 200, result.StatusCode)
}

func TestAccessKeyViaHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("X-Access-Key", "donkey")
	result := makeRequest(req, Response{}, "donkey")
	assert.Equal(t, 200, result.StatusCode)
}

func makeRequest(req *http.Request, respData Response, key string) *http.Response {
	rr := httptest.NewRecorder()
	respBuilder := func() *Response {
		return &respData
	}

	handler := getServerHandler(respBuilder, key)

	// Test standard response
	handler.ServeHTTP(rr, req)

	return rr.Result()
}
