package middlewaremocks

import (
	"auth-and-db-service/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockUserAuthenticator struct {
	mock.Mock
}

func (ua *MockUserAuthenticator) UserAuth(c *gin.Context) (*middleware.Authentication, error) {
	var a = middleware.Authentication{
		EMail: "registered@test.com",
		Type:  "user",
	}

	return &a, nil
}

type MockSellerAuthenticator struct {
	mock.Mock
}

func (ua *MockSellerAuthenticator) UserAuth(c *gin.Context) (*middleware.Authentication, error) {
	var a = middleware.Authentication{
		EMail: "registered@test.com",
		Type:  "seller",
	}

	return &a, nil
}
