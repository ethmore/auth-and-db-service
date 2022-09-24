package postgresql

import "go.mongodb.org/mongo-driver/bson/primitive"

type IPostgreSQL interface {
	InsertOrder(userID primitive.ObjectID, o Order) (int, error)
	GetAllOrders(userID string) ([]Order, error)
	GetAllOrderProducts(orderID int) ([]Product_, error)

	InsertPayment(buyerId, addressId, totalPrice string) (int, error)
	UpdatePaymentStatus(status string, paymentID int) error

	InsertProduct(sellerMail, title, price, description, photo, stock string) error
	UpdateProduct(title, price, description, photo, stock, id string) error
	DeleteProduct(id string) error
	GetSellerProducts(eMail string) ([]Product, error)
	GetAllProducts() ([]Product, error)
	GetProduct(id string) (*Product, error)

	Insert(name, email, password, address, phonenumber string) error
	Update(name, email, password, address, phonenumber, id string) error
	Delete(id string) error
	GetSeller(email string) (*Seller, error)
	GetSellerNameByID(id string) (*SellerName, error)
}

type Postgresql struct{}

type MockPostregql struct{}

func (mp *MockPostregql) InsertOrder(userID primitive.ObjectID, o Order) (int, error) {
	return 1, nil
}
func (mp *MockPostregql) GetAllOrders(userID string) ([]Order, error) {
	return nil, nil
}
func (mp *MockPostregql) GetAllOrderProducts(orderID int) ([]Product_, error) {
	return nil, nil
}

func (mp *MockPostregql) InsertProduct(sellerMail, title, price, description, photo, stock string) error {
	return nil
}
func (mp *MockPostregql) UpdateProduct(title, price, description, photo, stock, id string) error {
	return nil
}
func (mp *MockPostregql) DeleteProduct(id string) error {
	return nil
}
func (mp *MockPostregql) GetSellerProducts(eMail string) ([]Product, error) {
	return nil, nil
}
func (mp *MockPostregql) GetAllProducts() ([]Product, error) {
	return nil, nil
}

func (mp *MockPostregql) GetProduct(id string) (*Product, error) {
	return nil, nil
}
func (mp *MockPostregql) Insert(name, email, password, address, phonenumber string) error {
	return nil
}
func (mp *MockPostregql) Update(name, email, password, address, phonenumber, id string) error {
	return nil
}
func (mp *MockPostregql) Delete(id string) error {
	return nil
}
func (mp *MockPostregql) GetSeller(email string) (*Seller, error) {
	return nil, nil
}
func (mp *MockPostregql) GetSellerNameByID(id string) (*SellerName, error) {
	return nil, nil
}
