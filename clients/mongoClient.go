package clients

import (
	"context"
	"g-case-study/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoClient struct {
	Client *mongo.Client
}

func CreateMongoClient(options *options.ClientOptions) MongoClient {
	// Create a new MongoDB client
	client, err := mongo.NewClient(options)
	if err != nil {
		logging.Log.Error(err)
	}

	return MongoClient{Client: client}
}

func (t *MongoClient) Connect() {
	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err := t.Client.Connect(ctx)
	if err != nil {
		logging.Log.Error(err)
	}
}

func (t *MongoClient) Disconnect() {
	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Close the MongoDB client when we're done
	err := t.Client.Disconnect(ctx)
	if err != nil {
		logging.Log.Error(err)
	}
}
