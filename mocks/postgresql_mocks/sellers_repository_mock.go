package postgresql_mocks

import (
	"auth-and-db-service/repositories/postgresql"
	"errors"

	"github.com/stretchr/testify/mock"
)

type MockSellerRepo struct {
	mock.Mock
}

func (ms *MockSellerRepo) Insert(name, email, password, address, phonenumber string) error {
	if name == "" || email == "" || password == "" || address == "" || phonenumber == "" {
		return errors.New("empty arg/s")
	}
	return nil
}

func (ms *MockSellerRepo) Update(name, email, password, address, phonenumber, id string) error {
	if name == "" || email == "" || password == "" || address == "" || phonenumber == "" {
		return errors.New("empty arg/s")
	}
	return nil
}

func (ms *MockSellerRepo) Delete(id string) error {
	if id == "" {
		return errors.New("empty arg/s")
	}
	return nil
}

func (ms *MockSellerRepo) GetSeller(email string) (*postgresql.Seller, error) {
	if email == "" {
		return nil, errors.New("empty email")
	}

	var mock = postgresql.Seller{
		Id:       123,
		Email:    "registered@test.com",
		Password: "$2a$14$r0VvNArMYwf3O.Tq1Hhg9uEGHRSGcjEOU6GO3UxRZvsrAdkP3tkua",
	}

	return &mock, nil
}

func (ms *MockSellerRepo) GetSellerNameByID(id string) (*postgresql.SellerName, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	var SellerNameMock = postgresql.SellerName{
		Id:          123,
		CompanyName: "testName",
	}
	return &SellerNameMock, nil
}
