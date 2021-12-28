package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ClientMongoDB() *mongo.Client {

	fmt.Println("Connecting!")
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://UGV2Udu4shHqbzkQ:UGV2Udu4shHqbzkQ@doriansproyect.2zla2.mongodb.net/DoriansProyect?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Congratulations, you're already connected to MongoDB!")
	return client
}
