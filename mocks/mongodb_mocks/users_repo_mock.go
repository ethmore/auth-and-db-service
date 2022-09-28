package mongodbmocks

import (
	"auth-and-db-service/repositories/mongodb"
	"errors"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUsersRepo struct {
	mock.Mock
}

func (mup *MockUsersRepo) InsertOneUser(name, surname, email, password string) error {
	if name == "" || surname == "" || email == "" || password == "" {
		return errors.New("empty arg/s")
	}
	return nil
}

func (mup *MockUsersRepo) UpdateOneUser(name, surname, email, password string) error {
	if name == "" || surname == "" || email == "" || password == "" {
		return errors.New("empty arg/s")
	}
	return nil
}

func (mup *MockUsersRepo) DeleteOneUser(id string) error {
	if id == "" {
		return errors.New("empty arg/s")
	}
	return nil
}

func (mup *MockUsersRepo) FindOneUser(email string) (*mongodb.User, error) {
	if email == "" {
		return nil, errors.New("empty arg/s")
	}

	if email == "registered@test.com" {
		user := mongodb.User{
			Id:       primitive.NewObjectID(),
			Name:     "testName",
			Surname:  "testSurname",
			Email:    "registered@test.com",
			Password: "$2a$14$r0VvNArMYwf3O.Tq1Hhg9uEGHRSGcjEOU6GO3UxRZvsrAdkP3tkua",
		}

		return &user, nil
	}

	return nil, nil
}

func (mup *MockUsersRepo) ChangeUserPassword(userMail, newPassword string) error {
	if userMail == "" || newPassword == "" {
		return errors.New("empty arg/s")
	}

	return nil
}
