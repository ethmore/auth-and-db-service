package postgresql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InsertUser(name, surname, email, password string) string {
	insertDynStmt := `INSERT INTO users (name, surname, email, password) values($1, $2, $3, $4)`
	_, e := db.Exec(insertDynStmt, name, surname, email, password)
	CheckError(e)
	return "Inserted Succesfully"
}

func UpdateUser(name, surname, email, password, id string) string {
	updateStmt := `update "users" set "name"=$1, "surname"=$2, "email"=$3, "password"=$4 where "id"=$5`
	_, e := db.Exec(updateStmt, name, surname, email, password, id)
	CheckError(e)
	return "Updated Succesfully"
}

func DeleteUser(id string) string {
	deleteStmt := `delete from "users" where id=$1`
	_, e := db.Exec(deleteStmt, id)
	CheckError(e)
	return "Deleted Succesfully"
}

type User struct {
	CompanyName   string
	Email         string
	Password      string
	PasswordAgain string
	Address       string
	PhoneNumber   string
}

func GetUser(email string) string {
	var seller Seller

	err := db.QueryRow("SELECT email FROM users WHERE email=$1", email).Scan(&seller.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return "not registered"
		} else {
			log.Fatal(err)
		}
	}
	return seller.Email
}
