package database

import (
	"context"
	"log"
	"time"

	"github.com/Deatsilence/lawai-go/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDB string

func DBinstance() *mongo.Client {
	config.LoadEnvVariables()
	mongoUri := config.GetMongoURI()
	mongoDB = config.GetMongoDBName()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))

	if err != nil {
		log.Fatalf("Error connecting to database `%v`", err)
	}

	log.Println("Connected to MongoDB!")

	defer cancel()
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(mongoDB).Collection(collectionName)
	return collection
}
