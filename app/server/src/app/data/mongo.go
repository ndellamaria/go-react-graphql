package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(
  MONGO_URI,
))

