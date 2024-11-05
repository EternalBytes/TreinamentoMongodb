package main

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getOneDoc(collection *mongo.Collection, ctx context.Context, id string) {
	if len(id) == 0 {
		panic(errors.New("id cannot be empty string"))
	}
	prim, err := primitive.ObjectIDFromHex(id)
	check(err)
	podcasts := bson.M{}
	// IF THE BSON D OR M PASSED IS EMPTY THE LEST DOCUMENT ADDED WILL BE RETURNED
	err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: prim}}).Decode(&podcasts)
	check(err)
	fmt.Println(podcasts["_id"])
}
