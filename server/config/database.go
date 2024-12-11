package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

// GetMongoDatabase
func GetMongoDatabase(ctx context.Context) *mongo.Database {
	
	var database *mongo.Database

	return database 

}


// ConnectDB connects to the MongoDB database

func ConnectDB (uri string) {

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection 
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{bson.E{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. Connected to MongoDB!")

}