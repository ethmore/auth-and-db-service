package postgresql_mocks

import (
	"auth-and-db-service/repositories/postgresql"
	"errors"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockOrderRepo struct {
	mock.Mock
}

func (mp *MockOrderRepo) InsertOrder(userID primitive.ObjectID, o postgresql.Order) (int, error) {
	return 123, nil
}

func (mp *MockOrderRepo) GetAllOrders(userID string) ([]postgresql.Order, error) {

	orders := []postgresql.Order{
		{ID: 123,
			Products: []postgresql.Product_{
				{
					Title:      "testTitle",
					Qty:        "123",
					Price:      "123",
					SellerName: "testName"},
				{
					Title:      "testTitle",
					Qty:        "123",
					Price:      "123",
					SellerName: "testName",
				},
			},
			TotalPrice:         "123",
			ShipmentAddressID:  "123",
			CardLastFourDigits: "1234",
			PaymentStatus:      "success",
			OrderStatus:        "success",
			OrderTime:          "00:00:00",
		}}
	return orders, nil
}

func (mp *MockOrderRepo) GetAllOrderProducts(orderID int) ([]postgresql.Product_, error) {
	if orderID == 0 {
		return nil, errors.New("empty orderID")
	}

	prods := []postgresql.Product_{
		{
			Title:      "testTitle",
			Qty:        "123",
			Price:      "123",
			SellerName: "testName"},
		{
			Title:      "testTitle",
			Qty:        "123",
			Price:      "123",
			SellerName: "testName",
		},
	}

	return prods, nil
}
