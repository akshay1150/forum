package models

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	//Token    string `json:"token" bson:"token"`
}

type Token struct {
	UserId uint
	jwt.StandardClaims
}

func (acc *Account) CreateAccount() (string, error) {

	resp := acc.validate()
	if resp != nil {
		return "", resp
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(acc.Password), 14)
	if err != nil {
		return "", errors.New("cannnot hash password")
	}
	acc.Password = string(hash)
	_, err = acc.Create()
	if err != nil {
		return "", errors.New("error in creating user")
	}

	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["exp"] = time.Now().Add(10 * time.Minute)
	// claims["authorized"] = true
	// claims["user"] = acc.Email
	// jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return jwtToken.SignedString([]byte("akshay"))
	//return id, err
	return "", nil

}

func (acc *Account) validate() error {

	if !strings.Contains(acc.Email, "@") {

		return errors.New("email invalid")
	}
	if len(acc.Password) < 6 {
		return errors.New("password length should be more than 6 chanracter")
	}
	_, err := acc.FindOne()
	if err != nil {
		return err
	}
	return nil
}
