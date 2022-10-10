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
	router.HandleFunc("/login", controller.Login)
	router.HandleFunc("/auth", controller.Auth)
	router.HandleFunc("/welcome", controller.IsAuthorized(controller.Welcome))

	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
