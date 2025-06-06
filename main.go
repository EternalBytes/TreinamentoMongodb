package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// https://www.mongodb.com/pt-br/docs/manual/installation/
// INSTALL ON UBUNTU https://www.mongodb.com/pt-br/docs/manual/tutorial/install-mongodb-on-ubuntu/

// CREATE OR UPDATE USER
// https://www.mongodb.com/pt-br/docs/manual/tutorial/change-own-password-and-custom-data/

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// SERVER API AND CLIENT OTIONS
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	
	opts := options.Client().ApplyURI("mongodb://mayconpr:012345@localhost:27017/?authMechanism=SCRAM-SHA-1&authSource=quickstart").SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, opts)
	check(err)

	defer func() {
		cancel() // context cancel func
		err := client.Disconnect(ctx)
		check(err)
	}()

	// create db, collections and documents
	quickstartDatabase := client.Database("quickstart")
	podecastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")

	getAllDocs(episodesCollection, ctx)
	getOneDoc(podecastsCollection, ctx, "67292905b4f97f60eacb27c1")

	/* podcastResult, err := podecastsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The polyglot developer podcast"},
		{Key: "author", Value: "Nic Raboy"},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})
	check(err)

	fmt.Println(podcastResult.InsertedID)

	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{Key: "podcasts", Value: podcastResult.InsertedID},
			{Key: "title", Value: "Episode #1"},
			{Key: "description", Value: "This is the first episode."},
			{Key: "duration", Value: 30},
		},
		bson.D{
			{Key: "podcasts", Value: podcastResult.InsertedID},
			{Key: "title", Value: "Episode #2"},
			{Key: "description", Value: "This is the second episode."},
			{Key: "duration", Value: 40},
		},
	})
	check(err)

	fmt.Println(episodeResult.InsertedIDs...)
	*/
	// https://www.youtube.com/watch?v=WEYtAKYbB6k&list=PL7nSMdzIjUOr27j_kTjask5hyOD1Fh8OB&index=3
}
