package main

import (
	"log"
	"net/http"

	"github.com/forum/controller"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.HomePage)
	router.HandleFunc("/signup", controller.CreateAccount)

	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
