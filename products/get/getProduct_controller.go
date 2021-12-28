package get

import (
	"APIRest-Golang-MongoDB/database"
	"APIRest-Golang-MongoDB/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProductByName(w http.ResponseWriter, req *http.Request) {
	collection := database.ClientMongoDB().Database("DoriansProjectDB").Collection("products")

	var product models.Product
	params := mux.Vars(req)
	name := params["name"]
	fmt.Println(name)
	filter := bson.M{"name": name}

	err := collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("enter error")
	}
	w.Header().Set("Content-Type", "application/json")
	if product.Name != "" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusPartialContent)
	}
	json.NewEncoder(w).Encode(product)
}
