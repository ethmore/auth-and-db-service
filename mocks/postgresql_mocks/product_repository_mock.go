package postgresql_mocks

import (
	"auth-and-db-service/repositories/postgresql"
	"errors"

	"github.com/stretchr/testify/mock"
)

type MockProductRepo struct {
	mock.Mock
}

func (mp *MockProductRepo) InsertProduct(sr postgresql.ISellerRepo, sellerMail, title, price, description, photo, stock string) error {
	if sellerMail == "" || title == "" || price == "" || description == "" || photo == "" || stock == "" {
		return errors.New("empty arg/s")
	}
	return nil
}

func (p *MockProductRepo) UpdateProduct(title, price, description, photo, stock, id string) error {
	if id == "" || title == "" || price == "" || description == "" || photo == "" || stock == "" {
		return errors.New("empty arg/s")
	}
	return nil
}

func (p *MockProductRepo) DeleteProduct(id string) error {
	if id == "" {
		return errors.New("empty arg/s")
	}
	return nil
}

func (p *MockProductRepo) GetSellerProducts(sr postgresql.ISellerRepo, eMail string) ([]postgresql.Product, error) {
	if eMail == "" {
		return nil, errors.New("empty arg/s")
	}

	var products = []postgresql.Product{
		{
			Id:          "123",
			Title:       "testTitle",
			Price:       "123",
			Description: "testDescription",
			Image:       "123",
			Stock:       "123",
			SellerID:    "123",
		}, {
			Id:          "123",
			Title:       "testTitle",
			Price:       "123",
			Description: "testDescription",
			Image:       "123",
			Stock:       "123",
			SellerID:    "123",
		},
	}
	return products, nil
}

func (p *MockProductRepo) GetAllProducts() ([]postgresql.Product, error) {
	var products = []postgresql.Product{
		{
			Id:          "123",
			Title:       "testTitle",
			Price:       "123",
			Description: "testDescription",
			Image:       "123",
			Stock:       "123",
			SellerID:    "123",
		}, {
			Id:          "123",
			Title:       "testTitle",
			Price:       "123",
			Description: "testDescription",
			Image:       "123",
			Stock:       "123",
			SellerID:    "123",
		}, {
			Id:          "123",
			Title:       "testTitle",
			Price:       "123",
			Description: "testDescription",
			Image:       "123",
			Stock:       "123",
			SellerID:    "123",
		},
	}
	return products, nil
}

func (p *MockProductRepo) GetProduct(id string) (*postgresql.Product, error) {
	if id == "" {
		return nil, errors.New("empty arg/s")
	}

	var product = postgresql.Product{
		Id:          "123",
		Title:       "testTitle",
		Price:       "123",
		Description: "testDescription",
		Image:       "123",
		Stock:       "123",
		SellerID:    "123",
	}

	return &product, nil
}
