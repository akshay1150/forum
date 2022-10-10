package controller

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/forum/models"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Header.Get("token")
		// if err != nil {
		// 	if err == http.ErrNoCookie {
		// 		// If the cookie is not set, return an unauthorized status
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		return
		// 	}
		// 	// For any other type of error, return a bad request status
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// Get the JWT string from the cookie
		tknStr := c

		//fmt.Println("*****token****")
		// Initialize a new instance of `Claims`
		claims := &models.Claims{}

		//fmt.Println("***1*****", tknStr)

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		var jwtKey = []byte("mysecretkey")
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		//fmt.Printf("%+v", tkn)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Finally, return the welcome message to the user, along with their
		// username given in the token
		w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
	}
}
