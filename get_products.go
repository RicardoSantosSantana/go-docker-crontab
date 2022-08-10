package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func get_products() {

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	var count int = len(responseObject.Pokemon)
	for i := 0; i < count; i++ {
		fmt.Println(responseObject.Pokemon[i])
	}
	fmt.Println(&responseObject)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
