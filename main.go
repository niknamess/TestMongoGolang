package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoField struct {
	FieldStr  string `json: "Filed Str"`
	FieldInt  int    `json: "Filed Int"`
	FieldBool bool   `json: "Filed Bool"`
}

func main() {
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOption Type:", reflect.TypeOf(clientOption))
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		fmt.Println("Mongo.connect() ERROR", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("Fiest_database").Collection("First Collection")
	fmt.Println("Collection type", reflect.TypeOf(col))
	oneDoc := MongoField{
		FieldStr:  "String Info Text",
		FieldInt:  82648,
		FieldBool: true,
	}
	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc))

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() api result type: ", result)

		newID := result.InsertedID
		fmt.Println("InsertedOne(), newID", newID)
		fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	}
}
