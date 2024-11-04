package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func delFromId(collection *mongo.Collection, ctx context.Context, idHex string) {
	primitive, err := primitive.ObjectIDFromHex(idHex)
	check(err)
	res, err := collection.DeleteOne(ctx, bson.M{"_id": primitive})
	check(err)
	fmt.Println("Deleted: ", res.DeletedCount)
}
