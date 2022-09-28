package controllers

import (
	"bytes"

	postgresqlmocks "auth-and-db-service/mocks/postgresql_mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreatePayment(t *testing.T) {
	r := gin.Default()

	mock := &postgresqlmocks.MockPaymentRepo{}
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

func TestUpdatePaymentStatus(t *testing.T) {
	r := gin.Default()

	mock := &postgresqlmocks.MockPaymentRepo{}
	r.POST("/updatePaymentStatus", UpdatePaymentStatus(mock))
	recorder := httptest.NewRecorder()

	var jsonStr = []byte(`{"paymentID":123, "status":"success"}`)
	request, _ := http.NewRequest("POST", "/updatePaymentStatus", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	if http.StatusOK != recorder.Code {
		t.Errorf("error. expected: %d - got: %d", http.StatusOK, recorder.Code)
	}
}
