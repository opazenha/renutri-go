package db

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance     *mongo.Client
	clientInstanceErr  error
	mongoOnce          sync.Once
	mongoDatabaseName  = "renutri"
)

// GetMongoClient returns a singleton MongoDB client instance
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		_ = godotenv.Load()
		uri := os.Getenv("MONGO_URL")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			clientInstanceErr = err
			return
		}
		clientInstance = client
		clientInstanceErr = nil
	})
	return clientInstance, clientInstanceErr
}

// GetDatabase returns the renutri database instance
func GetDatabase() (*mongo.Database, error) {
	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}
	return client.Database(mongoDatabaseName), nil
}

// GetCollection returns a collection from renutri DB by name
func GetCollection(collectionName string) (*mongo.Collection, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}
	return db.Collection(collectionName), nil
}
