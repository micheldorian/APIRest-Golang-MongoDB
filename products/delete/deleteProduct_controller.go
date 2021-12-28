package delete

import (
	"APIRest-Golang-MongoDB/database"
	"APIRest-Golang-MongoDB/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProductByName(w http.ResponseWriter, req *http.Request) {
	collection := database.ClientMongoDB().Database("DoriansProjectDB").Collection("products")
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
