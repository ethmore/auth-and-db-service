package services

import (
	"e-comm/authService/bcrypt"
	"e-comm/authService/dotEnv"
	"e-comm/authService/repositories/mongodb"
	"errors"
)

type ChangePassword struct {
	Token            string
	OldPassword      string
	NewPassword      string
	NewPasswordAgain string
}

func ChangeUserPassword(passBody ChangePassword, email string) error {

	user, findErr := mongodb.FindOneUser(email)
	if findErr != nil {
		return findErr
	}

	if passBody.NewPassword != passBody.NewPasswordAgain {
		return errors.New("passwords does not match")
	}

	salt := dotEnv.GoDotEnvVariable("SALT")
	saltedPassword := passBody.OldPassword + salt
	match := bcrypt.CheckPasswordHash(saltedPassword, user.Password)
	if !match {
		return errors.New("old password does not match")
	}

	newSalt := dotEnv.GoDotEnvVariable("SALT")
	newSaltedPassword := passBody.NewPassword + newSalt
	newHash, _ := bcrypt.HashPassword(newSaltedPassword)

	updateErr := mongodb.ChangeUserPassword(email, newHash)
	if updateErr != nil {
		return updateErr
	}

	return nil
}
