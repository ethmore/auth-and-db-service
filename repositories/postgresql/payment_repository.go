package postgresql

import "errors"

func (p *Postgresql) InsertPayment(buyerId, addressId, totalPrice string) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO payments (buyer_id, address_id, total_price) VALUES ($1, $2, $3) RETURNING id", buyerId, addressId, totalPrice).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *Postgresql) UpdatePaymentStatus(status string, paymentID int) error {
	updateStmt := `UPDATE payments SET "status"=$1 WHERE id=$2`
	_, err := db.Exec(updateStmt, status, paymentID)
	if err != nil {
		return err
	}
	return nil
}

func (mp *MockPostregql) InsertPayment(buyerId, addressId, totalPrice string) (int, error) {
	if buyerId == "" || addressId == "" || totalPrice == "" {
		return 0, errors.New("empty arg/s")
	}
	return 1, nil
}

func (mp *MockPostregql) UpdatePaymentStatus(status string, paymentID int) error {
	if status != "" || paymentID != 0 {
		return errors.New("empty arg/s")
	}
	return nil
}
