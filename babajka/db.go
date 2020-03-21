package babajka

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "contentanalytics"
)

type dbAnalyticsDocument struct {
	Slug    string
	Metrics LocalizedMetric
}

func (cl *Client) pushMetricsToDB(metrics *Metrics) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err := getMongoCollection(ctx, cl.config.Mongodb.URL, cl.config.Mongodb.Options.DBName)
	if err != nil {
		return 0, err
	}

	err = collection.Drop(ctx)
	if err != nil {
		return 0, err
	}

	count := 0
	for slug, analyticsData := range *metrics {
		if _, err := collection.InsertOne(ctx,
			dbAnalyticsDocument{Slug: slug, Metrics: analyticsData}); err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}

func getMongoCollection(ctx context.Context, connectionString string, dbName string) (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to ", connectionString)
	return client.Database(dbName).Collection(collectionName), nil
}
