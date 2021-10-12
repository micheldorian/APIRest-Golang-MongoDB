package main

import (
	"APIRest-Golang-MongoDB/products/delete"
	"APIRest-Golang-MongoDB/products/get"
	"APIRest-Golang-MongoDB/products/post"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	serveWeb()
}

func serveWeb() {
	router := mux.NewRouter()

	//Endpointsgo get -v -u github.com/gorilla/mux
	router.HandleFunc("/product/{name}", get.GetProductByName).Methods("GET")
	router.HandleFunc("/product", get.GetProducts).Queries("Limit", "{[0-9]*?}").Methods("GET")
	router.HandleFunc("/product", post.CreateProduct).Methods("POST")
	router.HandleFunc("/product", delete.DeleteProductByName).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
