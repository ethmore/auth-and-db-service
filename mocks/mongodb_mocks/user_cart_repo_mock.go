package mongodbmocks

import (
	"auth-and-db-service/repositories/mongodb"
	"auth-and-db-service/repositories/postgresql"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserCartRepo struct{}

func (muc MockUserCartRepo) NewCart(up mongodb.IUsersRepo, userMail string) error {
	return nil
}

func (muc MockUserCartRepo) AddProductToCart(postgresqlRepo postgresql.IProductRepo, up mongodb.IUsersRepo, userMail, productId, qty string) error {
	return nil
}

func (muc MockUserCartRepo) FindAllCartProducts(up mongodb.IUsersRepo, userMail string) (*mongodb.Cart, error) {
	cart := mongodb.Cart{
		Id:     "123",
		UserId: primitive.NewObjectID(),
		Products: []mongodb.Product{
			{
				Id:  "123",
				Qty: "123",
			}, {
				Id:  "123",
				Qty: "123",
			},
		},
		TotalPrice: "123",
	}

	return &cart, nil
}

func (muc MockUserCartRepo) RemoveProductFromCart(up mongodb.IUsersRepo, userMail, productId string) error {
	return nil
}

func (muc MockUserCartRepo) UpdateCartProducts(userId primitive.ObjectID, newProducts []mongodb.Product) error {
	return nil
}

func (muc MockUserCartRepo) ChangeProductQty(up mongodb.IUsersRepo, userMail, productId, productQty string) error {
	return nil
}

func (muc MockUserCartRepo) AddTotalToCart(up mongodb.IUsersRepo, userMail, totalPrice string) error {
	return nil
}

func (muc MockUserCartRepo) GetTotalPrice(up mongodb.IUsersRepo, userMail string) (string, error) {
	return "123", nil
}

func (muc MockUserCartRepo) ClearCart(up mongodb.IUsersRepo, userMail string) error {
	return nil
}
