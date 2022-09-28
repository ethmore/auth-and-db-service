package services

import (
	mongodbmocks "auth-and-db-service/mocks/mongodb_mocks"
	"auth-and-db-service/mocks/postgresql_mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	loginBody := LoginBody{
		Email:    "registered@test.com",
		Password: "test",
		Type:     "user",
	}

	mock := &mongodbmocks.MockUsersRepo{}
	_, err := UserLogin(mock, loginBody)
	assert.Equal(t, err, nil)
}

func TestSellerLogin(t *testing.T) {
	loginBody := LoginBody{
		Email:    "registered@test.com",
		Password: "test",
		Type:     "user",
	}

	mock := &postgresql_mocks.MockSellerRepo{}
	_, err := SellerLogin(mock, loginBody)
	assert.Equal(t, nil, err)
}
