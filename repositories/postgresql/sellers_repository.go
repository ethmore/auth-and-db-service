package postgresql

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type Seller struct {
	Id          int
	CompanyName string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
}

type SellerName struct {
	Id          int
	CompanyName string
}

type ISellerRepo interface {
	Insert(name, email, password, address, phonenumber string) error
	Update(name, email, password, address, phonenumber, id string) error
	Delete(id string) error
	GetSeller(email string) (*Seller, error)
	GetSellerNameByID(id string) (*SellerName, error)
}

type SellerRepo struct{}

func (s *SellerRepo) Insert(name, email, password, address, phonenumber string) error {
	insertDynStmt := `INSERT INTO sellers (name, email, password, address, phonenumber) values($1, $2, $3, $4, $5)`
	_, err := db.Exec(insertDynStmt, name, email, password, address, phonenumber)
	return err
}

func (s *SellerRepo) Update(name, email, password, address, phonenumber, id string) error {
	updateStmt := `update "sellers" set "name"=$1, "email"=$2, "password"=$3, "address"=$4, "phonenumber"=$5 where "id"=$6`
	_, err := db.Exec(updateStmt, name, email, password, address, phonenumber, id)
	return err
}

func (s *SellerRepo) Delete(id string) error {
	deleteStmt := `delete from "sellers" where id=$1`
	_, err := db.Exec(deleteStmt, id)
	return err
}

func (s *SellerRepo) GetSeller(email string) (*Seller, error) {
	var seller Seller

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	err := db.QueryRow("SELECT id, email, password FROM sellers WHERE email=$1", email).Scan(&seller.Id, &seller.Email, &seller.Password)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &seller, nil
}

func (s *SellerRepo) GetSellerNameByID(id string) (*SellerName, error) {
	var seller SellerName

	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	err := db.QueryRow("SELECT id, name FROM sellers WHERE id=$1", id).Scan(&seller.Id, &seller.CompanyName)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &seller, nil
}
