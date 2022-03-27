package dbConnection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DbClient *sql.DB

func Connect() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/user_profile")
	if err != nil {
		panic(err.Error())
	}

	DbClient = db
}