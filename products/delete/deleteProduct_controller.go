package delete

import (
	"APIRest-Golang-MongoDB/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteProductByName(w http.ResponseWriter, req *http.Request) {

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
	filter := bson.D{primitive.E{Key: "name", Value: product.Name}}

	deleteResp, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println("Other error 500")
		log.Fatal(err)
	}
	if deleteResp.DeletedCount == 1 {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Product ", product.Name, " deleted")
	} else {
		w.WriteHeader(http.StatusPartialContent)
		fmt.Println("Product ", product.Name, "could not be deleted")
	}
}
