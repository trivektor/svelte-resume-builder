package db

import (
  "context"
  "go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

var Context context.Context

var CancelFunc context.CancelFunc
