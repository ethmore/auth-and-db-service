package controllers

import (
	middlewaremocks "auth-and-db-service/mocks/middleware_mocks"
	mongodbmocks "auth-and-db-service/mocks/mongodb_mocks"
	postgresqlmocks "auth-and-db-service/mocks/postgresql_mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddProductToCart(t *testing.T) {
	r := gin.Default()

	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockProd := &postgresqlmocks.MockProductRepo{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserCart := &mongodbmocks.MockUserCartRepo{}

	r.POST("/addProductToCart", AddProductToCart(mockAuth, mockProd, mockUser, mockUserCart))
	recorder := httptest.NewRecorder()

	var cartRequest = CartRequest{
		Id:  "123",
		Qty: "1",
	}
	var jsonStr, _ = json.Marshal(cartRequest)

	request, _ := http.NewRequest("POST", "/addProductToCart", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetCartInfo(t *testing.T) {
	r := gin.Default()

	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockProd := &postgresqlmocks.MockProductRepo{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserCart := &mongodbmocks.MockUserCartRepo{}

	r.POST("/getCartInfo", GetCartInfo(mockAuth, mockProd, mockUser, mockUserCart))
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/getCartInfo", bytes.NewBuffer([]byte(``)))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetCartProducts(t *testing.T) {
	r := gin.Default()

	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserCart := &mongodbmocks.MockUserCartRepo{}

	r.POST("/getCartProducts", GetCartProducts(mockAuth, mockUser, mockUserCart))
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/getCartProducts", bytes.NewBuffer([]byte(``)))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRemoveProductFromCart(t *testing.T) {
	r := gin.Default()

	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserCart := &mongodbmocks.MockUserCartRepo{}

	r.POST("/removeProductFromCart", RemoveProductFromCart(mockAuth, mockUser, mockUserCart))
	recorder := httptest.NewRecorder()

	var cartRequest = CartRequest{
		Id:  "123",
		Qty: "1",
	}
	var jsonStr, _ = json.Marshal(cartRequest)

	request, _ := http.NewRequest("POST", "/removeProductFromCart", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestChangeProductQty(t *testing.T) {
	r := gin.Default()

	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserCart := &mongodbmocks.MockUserCartRepo{}

	r.POST("/changeProductQty", ChangeProductQty(mockAuth, mockUser, mockUserCart))
	recorder := httptest.NewRecorder()

	var cartRequest = CartRequest{
		Id:  "123",
		Qty: "1",
	}
	var jsonStr, _ = json.Marshal(cartRequest)

	request, _ := http.NewRequest("POST", "/changeProductQty", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestAddTotalToCart(t *testing.T) {
	r := gin.Default()

	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserCart := &mongodbmocks.MockUserCartRepo{}

	r.POST("/addTotalToCart", AddTotalToCart(mockAuth, mockUser, mockUserCart))
	recorder := httptest.NewRecorder()

	var purchaseBody = PurchaseBody{
		TotalPrice: "123",
	}
	var jsonStr, _ = json.Marshal(purchaseBody)

	request, _ := http.NewRequest("POST", "/addTotalToCart", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetTotalPrice(t *testing.T) {
	r := gin.Default()

	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserCart := &mongodbmocks.MockUserCartRepo{}

	r.POST("/getTotalPrice", GetTotalPrice(mockAuth, mockUser, mockUserCart))
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/getTotalPrice", bytes.NewBuffer([]byte(``)))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestClearCart(t *testing.T) {
	r := gin.Default()

	mockAuth := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserCart := &mongodbmocks.MockUserCartRepo{}

	r.POST("/clearCart", ClearCart(mockAuth, mockUser, mockUserCart))
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/clearCart", bytes.NewBuffer([]byte(``)))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
