package repo

import (
	"context"
	"g-case-study/clients"
	"g-case-study/dto/record"
	"g-case-study/globals"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Record struct {
	Key       string    `bson:"key"`
	CreatedAt time.Time `bson:"createdAt"`
	Counts    []int     `bson:"counts"`
}

type RecordsRepository struct {
	clientOptions *options.ClientOptions
}

func CreateRecordsRepository() *RecordsRepository {
	// Set up MongoDB client options
	return &RecordsRepository{clientOptions: options.Client().ApplyURI(globals.ApplicationSettings.MongoDbAddress)}
}

func (repo *RecordsRepository) FindRecordSummaries(startDate time.Time, endDate time.Time,
	countStart int, countEnd int) ([]record.RecordSummaryDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := clients.CreateMongoClient(repo.clientOptions)
	client.Connect()
	collection := client.Client.Database(globals.ApplicationSettings.MongoDbName).Collection("records")
	defer client.Disconnect()

	// Define pipeline stages
	pipeline := []bson.M{
		{"$addFields": bson.M{"dateColumn": bson.M{"$toDate": "$createdAt"}}},
		{"$match": bson.M{"dateColumn": bson.M{"$gte": startDate, "$lte": endDate}}},
		{"$project": bson.M{"key": "$key", "createdAt": "$createdAt", "totalCount": bson.M{"$sum": "$counts"}}},
		{"$match": bson.M{"totalCount": bson.M{"$gte": countStart, "$lte": countEnd}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var summaries []record.RecordSummaryDto
	if err := cursor.All(ctx, &summaries); err != nil {
		return nil, err
	}

	return summaries, nil
}
