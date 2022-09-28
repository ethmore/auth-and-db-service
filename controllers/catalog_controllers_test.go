package controllers

import (
	postgresqlmocks "auth-and-db-service/mocks/postgresql_mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {
	r := gin.Default()

	mockProd := &postgresqlmocks.MockProductRepo{}

	r.POST("/getProduct", GetProduct(mockProd))
	recorder := httptest.NewRecorder()

	var body = ProductResponse{
		Id: "123",
	}
	var jsonStr, _ = json.Marshal(body)

	request, _ := http.NewRequest("POST", "/getProduct", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAllProducts(t *testing.T) {
	r := gin.Default()

	mockProd := &postgresqlmocks.MockProductRepo{}

	r.POST("/getAllProducts", GetAllProducts(mockProd))
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/getAllProducts", bytes.NewBuffer([]byte(``)))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
