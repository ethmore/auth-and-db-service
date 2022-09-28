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
	Image       string
	Stock       string
	SellerID    string
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
		insertStmt := `INSERT INTO products (sellerid, title, price, description, photo, stock) values ($1, $2, $3, $4, $5, $6)`
		_, err := db.Exec(insertStmt, seller.Id, title, price, description, photo, stock)
		return err
	} else {
		err := errors.New("postgressql: Insert id is empty")
		return err
	}
}

func (p *ProductRepo) UpdateProduct(title, price, description, photo, stock, id string) error {
	updateStmt := `update "products" set "title"=$1, "price"=$2, "description"=$3, "photo"=$4, "stock"=$5 where id=$6`
	_, err := db.Exec(updateStmt, title, price, description, photo, stock, id)
	return err
}

func (p *ProductRepo) DeleteProduct(id string) error {
	deleteStmt := `delete from "products" where id=$1`
	_, err := db.Exec(deleteStmt, id)
	return err
}

func (p *ProductRepo) GetSellerProducts(sr ISellerRepo, eMail string) ([]Product, error) {
	seller, err := sr.GetSeller(eMail)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, title, price, description, photo, stock FROM products WHERE sellerid=$1", seller.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Title, &product.Price, &product.Description, &product.Image, &product.Stock)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepo) GetAllProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, title, price, description, photo, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Title, &product.Price, &product.Description, &product.Image, &product.Stock)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepo) GetProduct(id string) (*Product, error) {
	var product Product
	idInt, _ := strconv.Atoi(id)

	err := db.QueryRow("SELECT id, title, price, description, photo, stock, sellerid FROM products WHERE id=$1", idInt).Scan(&product.Id, &product.Title, &product.Price, &product.Description, &product.Image, &product.Stock, &product.SellerID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
