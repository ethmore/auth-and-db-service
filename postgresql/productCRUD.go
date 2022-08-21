package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func InsertProduct(sellerMail, title, price, description, photo, stock string) int {
	sellerId, _, _ := GetSeller(sellerMail)

	if sellerId != 0 {
		insertStmt := `INSERT INTO products (sellerid, title, price, description, photo, stock) values ($1, $2, $3, $4, $5, $6)`
		_, e := db.Exec(insertStmt, sellerId, title, price, description, photo, stock)
		CheckError(e)
		return 200
	} else {
		return 400
	}
}

func UpdateProduct(title, price, description, photo, stock, id string) int {
	updateStmt := `update "products" set "title"=$1, "price"=$2, "description"=$3, "photo"=$4, "stock"=$5 where "id"=$6`
	_, e := db.Exec(updateStmt, title, price, description, photo, stock, id)
	CheckError(e)
	return 200
}

func DeleteProduct(id string) int {
	deleteStmt := `delete from "products" where id=$1`
	_, e := db.Exec(deleteStmt, id)
	CheckError(e)
	return 200
}

type Product struct {
	Id          string
	Title       string
	Price       string
	Description string
	Image       string
	Stock       string
}

func GetSellerProducts(eMail string) ([]Product, error) {
	id, _, _ := GetSeller(eMail)

	rows, err := db.Query("SELECT id, title, price, description, photo, stock FROM products WHERE sellerid=$1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Fatal(err)
		}
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Title, &product.Price, &product.Description, &product.Image, &product.Stock)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)
	return products, nil
}

func GetAllProducts() ([]Product, error) {
	fmt.Println("A")

	rows, err := db.Query("SELECT id, title, price, description, photo, stock FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Fatal(err)
		}
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Title, &product.Price, &product.Description, &product.Image, &product.Stock)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)
	return products, nil
}

func GetProduct(id string) (*Product, error) {
	var product Product
	idInt, _ := strconv.Atoi(id)

	err := db.QueryRow("SELECT id, title, price, description, photo, stock FROM products WHERE id=$1", idInt).Scan(&product.Id, &product.Title, &product.Price, &product.Description, &product.Image, &product.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println(product)
	return &product, nil
}
