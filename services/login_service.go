package services

import (
	"auth-and-db-service/bcrypt"
	"auth-and-db-service/dotEnv"
	"auth-and-db-service/repositories/mongodb"
	"auth-and-db-service/repositories/postgresql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type LoginBody struct {
	Email    string
	Password string
	Type     string
}

func UserLogin(ur mongodb.IUsersRepo, userBody LoginBody) (string, error) {
	user, mongoErr := ur.FindOneUser(userBody.Email)
	if mongoErr != nil {
		return "", mongoErr
	}
	if user == nil {
		return "", errors.New("email not registered")
	}

	salt := dotEnv.GoDotEnvVariable("SALT")
	saltedPassword := userBody.Password + salt
	match := bcrypt.CheckPasswordHash(saltedPassword, user.Password)

	if !match {
		return "", errors.New("wrong password")
	}

	secretToken := dotEnv.GoDotEnvVariable("TOKEN")
	hmacSampleSecret := []byte(secretToken)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mail": userBody.Email,
		"type": userBody.Type,
		"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, tokenErr := token.SignedString(hmacSampleSecret)
	if tokenErr != nil {
		return "", tokenErr
	}

	return tokenString, nil
}

func SellerLogin(sellerRepo postgresql.ISellerRepo, sellerBody LoginBody) (string, error) {
	seller, pqErr := sellerRepo.GetSeller(sellerBody.Email)
	if pqErr != nil {
		return "", pqErr
	}
	if seller.Email != sellerBody.Email {
		return "", errors.New("email not registered")
	}

	salt := dotEnv.GoDotEnvVariable("SALT")
	saltedPassword := sellerBody.Password + salt
	match := bcrypt.CheckPasswordHash(saltedPassword, seller.Password)
	if !match {
		return "", errors.New("wrong password")
	}

	secretToken := dotEnv.GoDotEnvVariable("TOKEN")
	hmacSampleSecret := []byte(secretToken)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mail": sellerBody.Email,
		"type": sellerBody.Type,
		"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, tokenErr := token.SignedString(hmacSampleSecret)
	if tokenErr != nil {
		return "", tokenErr
	}

	return tokenString, nil
}
