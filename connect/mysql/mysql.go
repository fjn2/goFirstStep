package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetInstance() *sql.DB {
	db, err := sql.Open("mysql", "root@/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("CONNECTED")

	return db
}
