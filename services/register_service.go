package services

import (
	"e-comm/authService/bcrypt"
	"e-comm/authService/dotEnv"
	"e-comm/authService/repositories/mongodb"
	"e-comm/authService/repositories/postgresql"
	"errors"
)

type UserRegisterBody struct {
	Name          string
	Surname       string
	Email         string
	Password      string
	PasswordAgain string
}

type SellerRegisterBody struct {
	CompanyName   string
	Email         string
	Password      string
	PasswordAgain string
	Address       string
	PhoneNumber   string
}

func UserRegister(userBody UserRegisterBody) error {
	if userBody.Password != userBody.PasswordAgain {
		return errors.New("passwords does not match")
	}

	user, mongoErr := mongodb.FindOneUser(userBody.Email)
	if mongoErr != nil {
		return mongoErr
	}
	if user != nil {
		return errors.New("email already registered")
	}

	salt := dotEnv.GoDotEnvVariable("SALT")
	saltedPassword := userBody.Password + salt
	hash, _ := bcrypt.HashPassword(saltedPassword)

	insertErr := mongodb.InsertOneUser(userBody.Name, userBody.Surname, userBody.Email, hash)
	if insertErr != nil {
		return insertErr
	}

	return nil
}

func SellerRegister(postgresqlRepo postgresql.IPostgreSQL, sellerBody SellerRegisterBody) error {
	if sellerBody.Password != sellerBody.PasswordAgain {
		return errors.New("passwords does not match")
	}

	sellerFromDB, getErr := postgresqlRepo.GetSeller(sellerBody.Email)
	if getErr != nil {
		return getErr
	}
	if sellerFromDB.Email == sellerBody.Email {
		return errors.New("email already registered")
	}

	salt := dotEnv.GoDotEnvVariable("SALT")
	saltedPassword := sellerBody.Password + salt
	hash, _ := bcrypt.HashPassword(saltedPassword)

	insertErr := postgresqlRepo.Insert(sellerBody.CompanyName, sellerBody.Email, hash, sellerBody.Address, sellerBody.PhoneNumber)
	if insertErr != nil {
		return insertErr
	}

	return nil
}
