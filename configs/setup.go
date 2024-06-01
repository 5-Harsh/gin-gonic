package configs

import (
	"context"

	"fmt"

	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI(LoadEnvVariable("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {

		log.Fatal(err)

	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {

		log.Fatal(err)

	}
	fmt.Println("Connected To MongoDB")

	return client
}

var DB *mongo.Client = ConnectDB()
