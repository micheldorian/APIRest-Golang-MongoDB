package get

import (
	"APIRest-Golang-MongoDB/database"
	"APIRest-Golang-MongoDB/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProducts(w http.ResponseWriter, req *http.Request) {

	collection := database.ClientMongoDB().Database("DoriansProjectDB").Collection("products")

	limitString := req.URL.Query().Get("Limit")
	limit, err := strconv.ParseInt(limitString, 0, 64)
	findOpts := options.Find()
	findOpts.SetLimit(limit)
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOpts)
	if err != nil {
		log.Fatal(err)
	}

	var products []*models.Product

	w.Header().Set("Content-Type", "application/json")
	if cur.RemainingBatchLength() == 0 {
		w.WriteHeader(http.StatusPartialContent)
	} else {
		w.WriteHeader(http.StatusOK)

	}

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
