package postgresql

import (
	"fmt"

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

// type Address struct {
// 	Id              string `bson:"_id"`
// 	Title           string
// 	Name            string
// 	Surname         string
// 	PhoneNumber     string
// 	Province        string
// 	County          string
// 	DetailedAddress string
// }

// type Order struct {
// 	Token              string
// 	Products           []Product_
// 	TotalPrice         string
// 	ShipmentAddress    Address
// 	CardLastFourDigits string
// 	PaymentStatus      string
// 	OrderStatus        string
// 	OrderTime          string
// }

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

func InsertOrder(userID primitive.ObjectID, o Order) (int, error) {
	var id int
	strUserID := userID.Hex()
	err := db.QueryRow("INSERT INTO orders (user_id, total_price, shipment_address_id, card_last_four_digits, payment_status, order_status, order_time) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", strUserID, o.TotalPrice, o.ShipmentAddressID, o.CardLastFourDigits, o.PaymentStatus, o.OrderStatus, o.OrderTime).Scan(&id)
	if err != nil {
		fmt.Println("order")
		return 0, err
	}

	for _, j := range o.Products {
		insertStmt := `INSERT INTO order_products (order_id, title, qty, price, seller_name) VALUES ($1, $2, $3, $4, $5)`
		_, err := db.Exec(insertStmt, id, j.Title, j.Qty, j.Price, j.SellerName)
		if err != nil {
			fmt.Println("order_prod")
			return 0, err
		}
	}

	return id, nil
}

func GetAllOrders(userID string) ([]Order, error) {
	rows, err := db.Query("SELECT id, total_price, shipment_address_id, card_last_four_digits, payment_status, order_status, order_time FROM orders WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var order Order
	var orders []Order
	for rows.Next() {
		err := rows.Scan(&order.ID, &order.TotalPrice, &order.ShipmentAddressID, &order.CardLastFourDigits, &order.PaymentStatus, &order.OrderStatus, &order.OrderTime)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func GetAllOrderProducts(orderID int) ([]Product_, error) {
	rows, err := db.Query("SELECT title, qty, price, seller_name FROM order_products WHERE order_id=$1", orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderProducts []Product_
	var orderProduct Product_
	for rows.Next() {
		err := rows.Scan(&orderProduct.Title, &orderProduct.Qty, &orderProduct.Price, &orderProduct.SellerName)
		if err != nil {
			return nil, err
		}

		orderProducts = append(orderProducts, orderProduct)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orderProducts, nil
}
