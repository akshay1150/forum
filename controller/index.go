package controller

import (
	"fmt"
	"net/http"

	"github.com/forum/utils"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("****1******")
	utils.GenerateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

func Login(w http.ResponseWriter, r *http.Request) {
	utils.GenerateHTML(w, nil, "login.layout", "public.navbar", "login")
}
