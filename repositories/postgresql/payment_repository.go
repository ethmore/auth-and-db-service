package postgresql

type IPaymentRepo interface {
	InsertPayment(buyerId, addressId, totalPrice string) (int, error)
	UpdatePaymentStatus(status string, paymentID int) error
}

type PaymentRepo struct{}

func (p *PaymentRepo) InsertPayment(buyerId, addressId, totalPrice string) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO payments (buyer_id, address_id, total_price) VALUES ($1, $2, $3) RETURNING id", buyerId, addressId, totalPrice).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PaymentRepo) UpdatePaymentStatus(status string, paymentID int) error {
	updateStmt := `UPDATE payments SET "status"=$1 WHERE id=$2`
	_, err := db.Exec(updateStmt, status, paymentID)
	if err != nil {
		return err
	}
	return nil
}
