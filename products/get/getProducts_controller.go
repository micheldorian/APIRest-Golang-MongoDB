package get

import (
	"APIRest-Golang-MongoDB/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProducts(w http.ResponseWriter, req *http.Request) {

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

	fmt.Println("collection")
	collection := client.Database("DoriansProjectDB").Collection("products")
	limitString := req.URL.Query().Get("Limit")
	limit, err := strconv.ParseInt(limitString, 0, 64)
	findOpts := options.Find()
	findOpts.SetLimit(limit)
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOpts)
	if err != nil {
		log.Fatal(err)
	}

	var products []*models.Product

	if cur.RemainingBatchLength() == 0 {
		w.WriteHeader(http.StatusPartialContent)
	} else {

		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s models.Product
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
			log.Fatal("Cur del next")
		}

		products = append(products, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		log.Fatal("Cur")
	}

	cur.Close(context.TODO())
	json.NewEncoder(w).Encode(products)

}
