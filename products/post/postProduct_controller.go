package post

import (
	"DoriansProyect/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateProduct(w http.ResponseWriter, req *http.Request) {

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
	var product models.Product
	_ = json.NewDecoder(req.Body).Decode(&product)

	insertResult, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal("Other error 500")
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Product ", product.Name, " created, id:", insertResult.InsertedID)
}
