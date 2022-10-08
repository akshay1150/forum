package controller

import (
	"net/http"

	"github.com/forum/models"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	// err := json.NewDecoder(r.Body).Decode(&account)
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, "unable to decode ", http.StatusBadRequest)
	// 	return
	// }
	_ = r.ParseForm()
	account.Email = r.FormValue("email")
	account.Password = r.FormValue("password")

	_, err := account.CreateAccount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//	w.WriteHeader(http.StatusCreated)
	//	resp := fmt.Sprintf("user created with email id %s", account.Email)
	//http.Error()
	//w.Write([]byte(resp))
	http.Redirect(w, r, "/login", http.StatusCreated)
}

func Auth(w http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()
	var ac models.Account
	ac.Email = r.FormValue("email")
	ac.Password = r.FormValue("password")

	tk, err := ac.FindUser()
	if err != nil {
		//w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tk.Value,
		Expires: tk.Expires,
	})
	w.Write([]byte("user logged in succesfully"))
}
