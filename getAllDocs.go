package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getAllDocs(collection *mongo.Collection, ctx context.Context) {
	cursor, err := collection.Find(ctx, bson.D{})
	check(err)
	defer func() {
		err := cursor.Close(ctx)
		check(err)
	}()

	// CAN OVERFLOW MEMORY IF COLECTION IS BIG
	/* episodes := new([]bson.M)
	err = cursor.All(ctx, episodes)
	check(err)

	for _, v := range *episodes {
		fmt.Println(v)
	} */

	// AVOID OVERFLOW MEMORY ON GETTING ALL DOCUMENTS OF A COLLECTION
	for cursor.Next(ctx) {
		episodes := bson.M{}
		err := cursor.Decode(&episodes)
		check(err)
		fmt.Println(episodes["_id"])
	}
}
