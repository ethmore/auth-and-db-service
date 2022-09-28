package postgresql_mocks

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type MockPaymentRepo struct {
	mock.Mock
}

func (mp *MockPaymentRepo) InsertPayment(buyerId, addressId, totalPrice string) (int, error) {
	if buyerId == "" || addressId == "" || totalPrice == "" {
		return 0, errors.New("empty arg/s")
	}
	return 1, nil
}

func (mp *MockPaymentRepo) UpdatePaymentStatus(status string, paymentID int) error {
	if status == "" || paymentID == 0 {
		return errors.New("empty arg/s")
	}
	return nil
}
