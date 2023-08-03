package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestValidNumber(t *testing.T) {
	r := gin.Default()
	r.GET("/validate", validateHandler)
	
	req, _ := http.NewRequest("GET", "/validate", nil)
	q := req.URL.Query()
	q.Add("number", "1234943")
	req.URL.RawQuery = q.Encode()
	testHTTPResponse(t, r, req, func (w *httptest.ResponseRecorder) bool {
		// validate the response format
		if w.Code != http.StatusOK || !json.Valid(w.Body.Bytes()) {
			return false
		}

		// check the result
		var result ValidateResponse
		err := json.Unmarshal(w.Body.Bytes(), &result)
		if err != nil || !result.Valid {
			return false
		}
		return true
	})
}

func TestInvalidNumber(t *testing.T) {
	r := gin.Default()
	r.GET("/validate", validateHandler)
	
	req, _ := http.NewRequest("GET", "/validate", nil)
	q := req.URL.Query()
	q.Add("number", "1234")
	req.URL.RawQuery = q.Encode()
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		if w.Code != http.StatusOK || !json.Valid(w.Body.Bytes()) {
			return false
		}

		var result ValidateResponse
		json.Unmarshal(w.Body.Bytes(), &result)
		return !result.Valid
	}) 
}

func TestInvalidRequest(t *testing.T) {
	r := gin.Default()
	r.GET("/validate", validateHandler)
	
	req, _ := http.NewRequest("GET", "/validate", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code != http.StatusOK 
	}) 
}
