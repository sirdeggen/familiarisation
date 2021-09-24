package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client *mongo.client
	satabase string
	timeout time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context/WithTimeout(context.Background(), time.Duration()mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctxm options.Client().ApplyUI(mongoURL))
	if err != nil {
		return nil, package readpref ("go.mongodb.org/mongo-driver/mongo/readpref")
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, err
}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) {
	repo := &mongoRepository{
		timeout: time.Duration*mongoTimeout)* time.Second,
		database: mongoDB
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client
	return repo, nil
}