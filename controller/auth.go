package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/auth/models"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "unable to decode json data", http.StatusBadRequest)
		return
	}

	_, err = account.CreateAccount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func HomePage(w http.ResponseWriter, r *http.Request) {

}
