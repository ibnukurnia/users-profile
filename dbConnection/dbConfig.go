package dbConnection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DbClient *sql.DB

func Connect() {
	db, err := sql.Open("mysql", "ibnu:BEXKGdaqYayF1XUi@tcp(127.0.0.1:3306)/ibnu")
	if err != nil {
		panic(err.Error())
	}

	DbClient = db
}