package postgresql

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product_ struct {
	Title      string
	Qty        string
	Price      string
	SellerName string
}

type OrderProduct struct {
	Title      string
	Qty        string
	Price      string
	SellerName string `json:"seller_name"`
}

type Order struct {
	Token              string
	ID                 int
	Products           []Product_
	TotalPrice         string
	ShipmentAddressID  string
	CardLastFourDigits string
	PaymentStatus      string
	OrderStatus        string
	OrderTime          string
}

type Orders struct {
	ID                 int
	TotalPrice         string
	ShipmentAddressID  string
	CardLastFourDigits string
	PaymentStatus      string
	OrderStatus        string
	OrderTime          string
	UserID             string
}

type OrderProducts struct {
	OrderID    int
	Title      string
	Qty        string
	Price      string
	SellerName string `json:"seller_name"`
}

type IOrderRepo interface {
	InsertOrder(userID primitive.ObjectID, o Order) (int, error)
	GetAllOrders(userID string) ([]Order, error)
	GetAllOrderProducts(orderID int) ([]Product_, error)
}

type OrderRepo struct{}

func (p *OrderRepo) InsertOrder(userID primitive.ObjectID, o Order) (int, error) {
	strUserID := userID.Hex()
	order := Orders{
		TotalPrice:         o.TotalPrice,
		ShipmentAddressID:  o.ShipmentAddressID,
		CardLastFourDigits: o.CardLastFourDigits,
		PaymentStatus:      o.PaymentStatus,
		OrderStatus:        o.OrderStatus,
		OrderTime:          o.OrderTime,
		UserID:             strUserID,
	}
	result := db2.Create(&order)
	if result.Error != nil {
		return 0, result.Error
	}

	for _, j := range o.Products {
		result2 := db2.Create(&OrderProducts{OrderID: order.ID, Title: j.Title, Qty: j.Qty, Price: j.Price, SellerName: j.SellerName})
		if result2.Error != nil {
			return 0, result2.Error
		}
	}

	return order.ID, nil
}

func (p *OrderRepo) GetAllOrders(userID string) ([]Order, error) {
	var orders []Order
	result := db2.Model(&Orders{}).Select("id", "total_price", "shipment_address_id", "card_last_four_digits", "payment_status", "order_status", "order_time").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (p *OrderRepo) GetAllOrderProducts(orderID int) ([]Product_, error) {
	var orderProducts []Product_

	result := db2.Model(&OrderProducts{}).Select("title", "qty", "price", "seller_name").Where("id = ?", orderID).Find(&orderProducts)
	if result.Error != nil {
		return nil, result.Error
	}

	return orderProducts, nil
}
