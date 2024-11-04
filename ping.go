package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ping(client *mongo.Client, ctx context.Context) {
	err := client.Ping(ctx, readpref.Primary())
	check(err)
}
