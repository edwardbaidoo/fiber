package collections

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsersCollections(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Fiber").Collection("Users")
	return collection
}
