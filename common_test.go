package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func testHTTPResponse(
	t *testing.T, 
	router *gin.Engine, 
	req *http.Request, 
	f func(recorder *httptest.ResponseRecorder) bool,
) {
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	if !f(recorder) {
		t.Fail()
	}
}