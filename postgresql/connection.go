package postgresql

import (
	"database/sql"
	"fmt"

	"strconv"
	"sync"

	"e-comm/authService/dotEnv"

	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

var db *sql.DB
var err error

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Connect() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			host := dotEnv.GoDotEnvVariable("HOST")
			port := dotEnv.GoDotEnvVariable("PORT")
			user := dotEnv.GoDotEnvVariable("USER")
			password := dotEnv.GoDotEnvVariable("PASSWORD")
			dbname := dotEnv.GoDotEnvVariable("DBNAME")

			fmt.Println("Creating PostgreSQL connection.")

			// connection string

			intPort, _ := strconv.Atoi(port)
			psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, intPort, user, password, dbname)
			// open database
			db, err = sql.Open("postgres", psqlconn)
			CheckError(err)

			// check db
			err = db.Ping()
			CheckError(err)

			fmt.Println("Connected!")

			singleInstance = &single{}
		} else {
			fmt.Println("PostgreSQL connection already created.")
		}
	} else {
		fmt.Println("PostgreSQL connection already created.")
	}

	return singleInstance
}
