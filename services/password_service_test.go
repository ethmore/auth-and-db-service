package services

import (
	mongodbmocks "auth-and-db-service/mocks/mongodb_mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeUserPassword(t *testing.T) {
	mock := &mongodbmocks.MockUsersRepo{}
	passBody := ChangePassword{
		OldPassword:      "test",
		NewPassword:      "testNew",
		NewPasswordAgain: "testNew",
	}

	err := ChangeUserPassword(mock, passBody, "registered@test.com")
	assert.Equal(t, nil, err, "")
}
