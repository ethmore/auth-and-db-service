package services

import (
	mongodbmocks "auth-and-db-service/mocks/mongodb_mocks"
	"auth-and-db-service/mocks/postgresql_mocks"
	"testing"
)

func TestUserRegister(t *testing.T) {
	userBody := UserRegisterBody{
		Name:          "testName",
		Surname:       "testSurname",
		Email:         "testTest@test.com",
		Password:      "test",
		PasswordAgain: "test",
	}

	mock := &mongodbmocks.MockUsersRepo{}
	err := UserRegister(mock, userBody)
	// assert.Equal(t, err, nil)
	if err != nil {
		t.Errorf("expected 'nil' got %s", err)
	}
}

func TestSellerRegister(t *testing.T) {
	sellerBody := SellerRegisterBody{
		CompanyName:   "testName",
		Email:         "testTest@test.com",
		Password:      "test",
		PasswordAgain: "test",
		Address:       "testAddress",
		PhoneNumber:   "5441111111",
	}

	mock := &postgresql_mocks.MockSellerRepo{}
	err := SellerRegister(mock, sellerBody)
	// assert.Equal(t, err, nil)
	if err != nil {
		t.Errorf("expected 'nil' got %s", err)
	}
}
