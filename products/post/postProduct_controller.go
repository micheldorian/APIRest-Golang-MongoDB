package post

import (
	"APIRest-Golang-MongoDB/database"
	"APIRest-Golang-MongoDB/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, req *http.Request) {
	collection := database.ClientMongoDB().Database("DoriansProjectDB").Collection("products")
	var product models.Product
	_ = json.NewDecoder(req.Body).Decode(&product)

	insertResult, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Fatal("Other error 500")
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Product ", product.Name, " created, id:", insertResult.InsertedID)
}
