package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func GetDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()
	//fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")
	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	//ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	return client
}

func (ac *Account) Create() (string, error) {

	col := GetDB().Database("employee").Collection("user")
	_, err := col.InsertOne(context.TODO(), ac)
	if err != nil {
		return "", errors.New("errro in insert operation")
	}
	return "", nil
}

func (ac *Account) FindOne() (string, error) {
	col := GetDB().Database("employee").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result := col.FindOne(ctx, bson.M{"email": ac.Email})
	// if err != nil {
	// 	return "", errors.New("errro in insert operation")
	// }
	var acc Account
	_ = result.Decode(&acc)
	if acc.Email == ac.Email {
		return "", errors.New("email already exists")
	}
	return "", nil
}

func (ac *Account) FindUser() (TokenData, error) {
	col := GetDB().Database("employee").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result := col.FindOne(ctx, bson.M{"email": ac.Email})
	// if err != nil {
	// 	return "", errors.New("errro in insert operation")
	// }
	var acc Account
	var tk TokenData
	_ = result.Decode(&acc)
	if acc.Email != ac.Email {
		return tk, errors.New(" user does not exist")
	}
	err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(ac.Password))
	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		return tk, errors.New("invalid credentials")
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: ac.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte("mysecretkey"))
	fmt.Println("*****token****", tokenString)
	log.Println(err)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		// w.WriteHeader(http.StatusInternalServerError)
		// return
		return tk, errors.New("error in creating token")
	}
	tk = TokenData{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}
	return tk, nil
}

type TokenData struct {
	Name    string
	Value   string
	Expires time.Time
}
