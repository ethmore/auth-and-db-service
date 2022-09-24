package controllers

import (
	"bytes"

	"e-comm/authService/repositories/postgresql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreatePayment(t *testing.T) {
	r := gin.Default()

	mock := new(postgresql.MockPostregql)
	r.POST("/createPayment", CreatePayment(mock))
	recorder := httptest.NewRecorder()

	var jsonStr = []byte(`{"buyerID":"123", "addressID":"234", "totalPrice":"345"}`)
	request, _ := http.NewRequest("POST", "/createPayment", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	if http.StatusOK != recorder.Code {
		t.Errorf("error. expected: %d - got: %d", http.StatusOK, recorder.Code)
	}
}
