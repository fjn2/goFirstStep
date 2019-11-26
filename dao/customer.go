package dao

import (
	"test/connect/mysql"
)

type Customer struct {
	Id   int
	Name string
}

func (_ *Customer) GetAll(channel chan Customer) {
	var db = mysql.GetInstance()
	rows, err := db.Query("SELECT * from test.Example")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		newCustomer := Customer{}
		rows.Scan(&newCustomer.Id, &newCustomer.Name)
		channel <- newCustomer
	}
	close(channel)
}
