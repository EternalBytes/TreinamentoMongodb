package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func listDBs(client *mongo.Client, ctx context.Context) {
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	check(err)

	fmt.Println(databases)
}
