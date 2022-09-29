package postgresql

import (
	"errors"
	"strconv"
)

type Product struct {
	Id          string
	Title       string
	Price       string
	Description string
	Image       string `gorm:"column:photo"`
	Stock       string
	SellerID    string `gorm:"column:sellerid"`
}

type IProductRepo interface {
	InsertProduct(sr ISellerRepo, sellerMail, title, price, description, photo, stock string) error
	UpdateProduct(title, price, description, photo, stock, id string) error
	DeleteProduct(id string) error
	GetSellerProducts(sr ISellerRepo, eMail string) ([]Product, error)
	GetAllProducts() ([]Product, error)
	GetProduct(id string) (*Product, error)
}

type ProductRepo struct{}

func (p *ProductRepo) InsertProduct(sr ISellerRepo, sellerMail, title, price, description, photo, stock string) error {
	seller, err := sr.GetSeller(sellerMail)
	if err != nil {
		return err
	}

	if seller.Id != 0 {
		strSellerID := strconv.Itoa(seller.Id)
		tx := db2.Create(&Product{SellerID: strSellerID, Title: title, Price: price, Description: description, Image: photo, Stock: stock})
		return tx.Error
	} else {
		err := errors.New("postgressql: Insert id is empty")
		return err
	}
}

func (p *ProductRepo) UpdateProduct(title, price, description, photo, stock, id string) error {
	product := Product{
		Id: id,
	}
	tx := db2.Model(&product).Updates(Product{Title: title, Price: price, Description: description, Image: photo, Stock: stock})
	return tx.Error
}

func (p *ProductRepo) DeleteProduct(id string) error {
	tx := db2.Delete(&Product{}, id)
	return tx.Error
}

func (p *ProductRepo) GetSellerProducts(sr ISellerRepo, eMail string) ([]Product, error) {
	seller, err := sr.GetSeller(eMail)
	if err != nil {
		return nil, err
	}

	var products []Product
	tx := db2.Model(&Product{}).Select("id", "title", "price", "description", "photo", "stock").Where("sellerid = ?", seller.Id).Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (p *ProductRepo) GetAllProducts() ([]Product, error) {
	var products []Product
	tx := db2.Model(&Product{}).Select("id", "title", "price", "description", "photo", "stock").Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (p *ProductRepo) GetProduct(id string) (*Product, error) {
	var product Product
	tx := db2.Model(&Product{}).Where("id = ?", id).Find(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &product, nil
}

func T() {
	sr := &PaymentRepo{}
	// sr.InsertPayment("26", "2342", "12345")
	sr.UpdatePaymentStatus("success", 139)
}
