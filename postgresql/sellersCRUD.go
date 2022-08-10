package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Check() {
	err := db.Ping()
	CheckError(err)
	fmt.Println("Connection OK!")
}

func Insert(name, email, password, address, phonenumber string) int {
	insertDynStmt := `INSERT INTO sellers (name, email, password, address, phonenumber) values($1, $2, $3, $4, $5)`
	_, e := db.Exec(insertDynStmt, name, email, password, address, phonenumber)
	CheckError(e)
	return 200
}

func Update(name, email, password, address, phonenumber, id string) int {
	updateStmt := `update "sellers" set "name"=$1, "email"=$2, "password"=$3, "address"=$4, "phonenumber"=$5 where "id"=$6`
	_, e := db.Exec(updateStmt, name, email, password, address, phonenumber, id)
	CheckError(e)
	return 200
}

func Delete(id string) int {
	deleteStmt := `delete from "sellers" where id=$1`
	_, e := db.Exec(deleteStmt, id)
	CheckError(e)
	return 200
}

type Seller struct {
	CompanyName   string
	Email         string
	Password      string
	PasswordAgain string
	Address       string
	PhoneNumber   string
}

func GetSeller(email string) string {
	var seller Seller

	err := db.QueryRow("SELECT email FROM sellers WHERE email=$1", email).Scan(&seller.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return "not registered"
		} else {
			log.Fatal(err)
		}
	}
	return seller.Email
}
