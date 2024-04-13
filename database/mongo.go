package database

import (
	"context"
	"log"
	"yumtrip/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI(constants.MongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the global client variable
	constants.MongoClient = client

	// Set the collections
	constants.MongoCollections["users"] = client.Database(constants.MongoDBName).Collection("users")
	constants.MongoCollections["menus"] = client.Database(constants.MongoDBName).Collection("menus")
	constants.MongoCollections["items"] = client.Database(constants.MongoDBName).Collection("items")
	constants.MongoCollections["orders"] = client.Database(constants.MongoDBName).Collection("orders")
	constants.MongoCollections["stores"] = client.Database(constants.MongoDBName).Collection("stores")
	constants.MongoCollections["coupons"] = client.Database(constants.MongoDBName).Collection("coupons")
	constants.MongoCollections["notifications"] = client.Database(constants.MongoDBName).Collection("notifications")
	constants.MongoCollections["analytics"] = client.Database(constants.MongoDBName).Collection("analytics")
	constants.MongoCollections["roles"] = client.Database(constants.MongoDBName).Collection("roles")
	constants.MongoCollections["permissions"] = client.Database(constants.MongoDBName).Collection("permissions")

}

func GetCollection(collectionName string) *mongo.Collection {
	return constants.MongoCollections[collectionName]
}