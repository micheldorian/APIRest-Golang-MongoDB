package get

import (
	"APIRest-Golang-MongoDB/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProductByName(w http.ResponseWriter, req *http.Request) {
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
	params := mux.Vars(req)
	name := params["name"]
	fmt.Println(name)
	filter := bson.M{"name": name}

	err = collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("enter error")
	}
	if product.Name != "" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusPartialContent)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}
