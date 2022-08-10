package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func insert_product() {

	const (
		DB_NAME = "dbmeli"
		DB_HOST = "127.0.0.1"
		DB_USER = "meli"
		DB_PASS = "meli123"
		DB_PORT = "3306"
	)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME))

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now().String()
	valor := "insert into tasks(title) values('" + currentTime + "')"

	res, err := db.Query(valor)

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

}
