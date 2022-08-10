package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

func api() {
	router := chi.NewRouter()
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
	 
		hah, err := ioutil.ReadAll(r.Body);
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		}
	
		fmt.Fprintf(w,"%s",hah)

	})

	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Printf("Error starting server: %s", err)
	}
}