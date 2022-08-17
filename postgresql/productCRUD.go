package postgresql

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
