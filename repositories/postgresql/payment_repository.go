package postgresql

type Payment struct {
	ID         int
	BuyerID    string `gorm:"column:buyer_id"`
	AddressID  string `gorm:"column:address_id"`
	TotalPrice string `gorm:"column:total_price"`
	Status     string
}

type IPaymentRepo interface {
	InsertPayment(buyerId, addressId, totalPrice string) (int, error)
	UpdatePaymentStatus(status string, paymentID int) error
}

type PaymentRepo struct{}

func (p *PaymentRepo) InsertPayment(buyerId, addressId, totalPrice string) (int, error) {
	payment := Payment{
		BuyerID:    buyerId,
		AddressID:  addressId,
		TotalPrice: totalPrice,
	}
	result := db2.Create(&payment)
	if result.Error != nil {
		return 0, result.Error
	}

	return payment.ID, nil
}

func (p *PaymentRepo) UpdatePaymentStatus(status string, paymentID int) error {
	payment := Payment{
		ID: paymentID,
	}
	result := db2.Model(&payment).Update("Status", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
