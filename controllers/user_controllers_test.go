package controllers

import (
	middlewaremocks "auth-and-db-service/mocks/middleware_mocks"
	mongodbmocks "auth-and-db-service/mocks/mongodb_mocks"
	"auth-and-db-service/services"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserRegisterPostHandler(t *testing.T) {
	r := gin.Default()

	mock := &mongodbmocks.MockUsersRepo{}
	r.POST("/userRegister", UserRegisterPostHandler(mock))
	recorder := httptest.NewRecorder()

	var body = services.UserRegisterBody{
		Name:          "testName",
		Surname:       "testSurname",
		Email:         "notRegistered@test.com",
		Password:      "test",
		PasswordAgain: "test",
	}

	var jsonStr, _ = json.Marshal(body)
	request, _ := http.NewRequest("POST", "/userRegister", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUserLoginPostHandler(t *testing.T) {
	r := gin.Default()

	mock := &mongodbmocks.MockUsersRepo{}
	r.POST("/userLogin", UserLoginPostHandler(mock))
	recorder := httptest.NewRecorder()

	var body = services.LoginBody{
		Email:    "registered@test.com",
		Password: "test",
		Type:     "user",
	}

	var jsonStr, _ = json.Marshal(body)
	request, _ := http.NewRequest("POST", "/userLogin", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUserProfile(t *testing.T) {
	r := gin.Default()

	mock := &middlewaremocks.MockUserAuthenticator{}
	r.POST("/profile", UserProfile(mock))
	recorder := httptest.NewRecorder()

	var body = services.LoginBody{
		Email:    "registered@test.com",
		Password: "test",
		Type:     "user",
	}

	var jsonStr, _ = json.Marshal(body)
	request, _ := http.NewRequest("POST", "/profile", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetUserInfo(t *testing.T) {
	r := gin.Default()

	mock := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}

	r.POST("/getUserInfo", GetUserInfo(mock, mockUser))
	recorder := httptest.NewRecorder()

	var body = Body{
		Token: "",
	}

	var jsonStr, _ = json.Marshal(body)
	request, _ := http.NewRequest("POST", "/getUserInfo", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestChangeUserPassword(t *testing.T) {
	r := gin.Default()

	mock := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}

	r.POST("/changeUserPassword", ChangeUserPassword(mock, mockUser))
	recorder := httptest.NewRecorder()

	var body = services.ChangePassword{
		OldPassword:      "test",
		NewPassword:      "testNew",
		NewPasswordAgain: "testNew",
	}

	var jsonStr, _ = json.Marshal(body)
	request, _ := http.NewRequest("POST", "/changeUserPassword", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestNewUserAddress(t *testing.T) {
	r := gin.Default()

	mock := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserAddress := &mongodbmocks.MockUserAddressesRepo{}

	r.POST("/newUserAddress", NewUserAddress(mock, mockUser, mockUserAddress))
	recorder := httptest.NewRecorder()

	var body = AddressBody{
		Title:           "testTitle",
		Name:            "testName",
		Surname:         "testSurname",
		PhoneNumber:     "5441111111",
		Province:        "testProvince",
		County:          "testCounty",
		DetailedAddress: "testAddress",
	}

	var jsonStr, _ = json.Marshal(body)
	request, _ := http.NewRequest("POST", "/newUserAddress", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetUserAddressById(t *testing.T) {
	r := gin.Default()

	mock := &middlewaremocks.MockUserAuthenticator{}
	mockUserAddress := &mongodbmocks.MockUserAddressesRepo{}

	r.POST("/getUserAddressById", GetUserAddressById(mock, mockUserAddress))
	recorder := httptest.NewRecorder()

	var body = PaymentAddress{
		AddressId: "123",
	}

	var jsonStr, _ = json.Marshal(body)
	request, _ := http.NewRequest("POST", "/getUserAddressById", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetUserAddresses(t *testing.T) {
	r := gin.Default()

	mock := &middlewaremocks.MockUserAuthenticator{}
	mockUser := &mongodbmocks.MockUsersRepo{}
	mockUserAddress := &mongodbmocks.MockUserAddressesRepo{}

	r.POST("/getUserAddressById", GetUserAddresses(mock, mockUser, mockUserAddress))
	recorder := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/getUserAddressById", bytes.NewBuffer([]byte(``)))
	request.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
