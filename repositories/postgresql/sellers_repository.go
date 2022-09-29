package postgresql

import (
	"errors"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Seller struct {
	Id          int
	CompanyName string `gorm:"column:name"`
	Email       string
	Password    string
	Address     string
	PhoneNumber string `gorm:"column:phonenumber"`
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
	tx := db2.Create(&Seller{CompanyName: name, Email: email, Password: password, Address: address, PhoneNumber: phonenumber})
	return tx.Error
}

func (s *SellerRepo) Update(name, email, password, address, phonenumber, id string) error {
	intID, _ := strconv.Atoi(id)
	var seller = Seller{
		Id: intID,
	}
	tx := db2.Model(&seller).Updates(Seller{CompanyName: name, Email: email, Password: password, Address: address, PhoneNumber: phonenumber})
	return tx.Error
}

func (s *SellerRepo) Delete(id string) error {
	tx := db2.Delete(&Seller{}, id)
	return tx.Error
}

func (s *SellerRepo) GetSeller(email string) (*Seller, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	var seller Seller
	tx := db2.Model(&Seller{}).Select("id", "email", "password").Where("email = ?", email).First(&seller)

	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &seller, nil
}

func (s *SellerRepo) GetSellerNameByID(id string) (*SellerName, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var seller SellerName
	tx := db2.Model(&Seller{}).Select("id", "name").Where("id = ?", id).First(&seller)

	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &seller, nil
}
