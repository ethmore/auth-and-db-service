package controllers

import (
	middlewaremocks "auth-and-db-service/mocks/middleware_mocks"
	mongodbmocks "auth-and-db-service/mocks/mongodb_mocks"
	postgresqlmocks "auth-and-db-service/mocks/postgresql_mocks"
	"auth-and-db-service/repositories/postgresql"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInsertOrder(t *testing.T) {
	r := gin.Default()

	mockOrder := &postgresqlmocks.MockOrderRepo{}
	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}

	r.POST("/insertOrder", InsertOrder(mockAuth, mockOrder, mockUser))
	recorder := httptest.NewRecorder()

	var orderBody = postgresql.Order{
		Token: "",
		ID:    123,
		Products: []postgresql.Product_{
			{
				Title:      "testTitle",
				Qty:        "1",
				Price:      "123",
				SellerName: "testName",
			},
			{
				Title:      "testTitle",
				Qty:        "1",
				Price:      "123",
				SellerName: "testName",
			},
		},
		TotalPrice:         "testDescription",
		ShipmentAddressID:  "123",
		CardLastFourDigits: "1234",
		PaymentStatus:      "success",
		OrderStatus:        "success",
		OrderTime:          "00:00:00",
	}

	var jsonStr, _ = json.Marshal(orderBody)

	request, _ := http.NewRequest("POST", "/insertOrder", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAllOrders(t *testing.T) {
	r := gin.Default()

	mockOrder := &postgresqlmocks.MockOrderRepo{}
	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}

	r.POST("/getAllOrders", GetAllOrders(mockAuth, mockOrder, mockUser))
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/getAllOrders", bytes.NewBuffer([]byte(``)))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
