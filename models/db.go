package models

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
