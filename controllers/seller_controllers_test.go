package controllers

import (
	middlewaremocks "auth-and-db-service/mocks/middleware_mocks"
	postgresqlmocks "auth-and-db-service/mocks/postgresql_mocks"
	"encoding/json"

	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSellerRegisterPostHandler(t *testing.T) {
	r := gin.Default()

	mock := &postgresqlmocks.MockSellerRepo{}
	r.POST("/sellerRegister", SellerRegisterPostHandler(mock))
	recorder := httptest.NewRecorder()

	var jsonStr = []byte(`{"companyName":"testName", "email":"test@test.com", "password":"test", "passwordAgain":"test", "address":"testAddress", "phoneNumber":"5441111111"}`)
	request, _ := http.NewRequest("POST", "/sellerRegister", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	if http.StatusOK != recorder.Code {
		t.Errorf("error. expected: %d - got: %d", http.StatusOK, recorder.Code)
	}
}

func TestSellerLoginPostHandler(t *testing.T) {
	r := gin.Default()

	mock := &postgresqlmocks.MockSellerRepo{}
	r.POST("/sellerLogin", SellerLoginPostHandler(mock))
	recorder := httptest.NewRecorder()

	var jsonStr = []byte(`{"email":"registered@test.com", "password":"test", "type":"user"}`)
	request, _ := http.NewRequest("POST", "/sellerLogin", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	if http.StatusOK != recorder.Code {
		t.Errorf("error. expected: %d - got: %d", http.StatusOK, recorder.Code)
	}
}

func TestAddProduct(t *testing.T) {
	r := gin.Default()

	mockProduct := &postgresqlmocks.MockProductRepo{}
	mockSeller := &postgresqlmocks.MockSellerRepo{}
	mockAuth := &middlewaremocks.MockUserAuthenticator{}

	r.POST("/addProduct", AddProduct(mockAuth, mockProduct, mockSeller))
	recorder := httptest.NewRecorder()

	var requestBody = Product{
		Token:       "",
		Id:          "123",
		Title:       "testTitle",
		Description: "testDescription",
		Price:       "123",
		Stock:       "123",
		Photo:       "test",
	}

	var jsonStr, _ = json.Marshal(requestBody)

	request, _ := http.NewRequest("POST", "/addProduct", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestEditProduct(t *testing.T) {
	r := gin.Default()

	mockProduct := &postgresqlmocks.MockProductRepo{}
	mockAuth := &middlewaremocks.MockUserAuthenticator{}

	r.POST("/editProduct", EditProduct(mockAuth, mockProduct))
	recorder := httptest.NewRecorder()

	var requestBody = Product{
		Token:       "",
		Id:          "123",
		Title:       "testTitle",
		Description: "testDescription",
		Price:       "123",
		Stock:       "123",
		Photo:       "test",
	}

	var jsonStr, _ = json.Marshal(requestBody)

	request, _ := http.NewRequest("POST", "/editProduct", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetSellerProducts(t *testing.T) {
	r := gin.Default()

	mockProduct := &postgresqlmocks.MockProductRepo{}
	mockAuth := &middlewaremocks.MockSellerAuthenticator{}
	mockSeller := &postgresqlmocks.MockSellerRepo{}
	r.POST("/getProducts", GetSellerProducts(mockAuth, mockProduct, mockSeller))
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/getProducts", bytes.NewBuffer([]byte(``)))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestDeleteProduct(t *testing.T) {
	r := gin.Default()

	mockProduct := &postgresqlmocks.MockProductRepo{}
	mockAuth := &middlewaremocks.MockSellerAuthenticator{}
	r.POST("/deleteProduct", DeleteProduct(mockAuth, mockProduct))
	recorder := httptest.NewRecorder()

	var requestBody = Product{
		Token:       "",
		Id:          "123",
		Title:       "testTitle",
		Description: "testDescription",
		Price:       "123",
		Stock:       "123",
		Photo:       "test",
	}

	var jsonStr, _ = json.Marshal(requestBody)

	request, _ := http.NewRequest("POST", "/deleteProduct", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetProductsSellers(t *testing.T) {
	r := gin.Default()

	mockProduct := &postgresqlmocks.MockProductRepo{}
	mockAuth := &middlewaremocks.MockSellerAuthenticator{}
	mockSeller := &postgresqlmocks.MockSellerRepo{}
	r.POST("/getProductsSellers", GetProductsSellers(mockAuth, mockProduct, mockSeller))
	recorder := httptest.NewRecorder()

	var requestBody = ProductsSeller{
		ProductIDs: []string{
			"1",
			"2",
			"3",
		},
	}

	var jsonStr, _ = json.Marshal(requestBody)

	request, _ := http.NewRequest("POST", "/getProductsSellers", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
