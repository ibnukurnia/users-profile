package dbConnection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DbClient *sql.DB

func Connect() {
	db, err := sql.Open("mysql", "ibnu:@BEXKGdaqYayF1XUitcp(http://167.172.73.163/phpmyadmin)/user_profile")
	if err != nil {
		panic(err.Error())
	}

	DbClient = db
}