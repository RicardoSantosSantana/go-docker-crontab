package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func leitura() {

	// response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	// responseData, err := ioutil.ReadAll(response.Body)

	// var responseObject Response
	// json.Unmarshal(responseData, &responseObject)

	// var count int = len(responseObject.Pokemon)
	// for i := 0; i < count; i++ {
	// 	fmt.Println(responseObject.Pokemon[i])
	// }
	// fmt.Println(&responseObject)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println(string(responseData))

	db, err := sql.Open("mysql", "meli:meli123@tcp(127.0.0.1:3306)/dbmeli")

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("SELECT * FROM tasks")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var task Tasks
		err := res.Scan(&task.Id, &task.title)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", task)
	}

}
