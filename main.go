package main

import (
	"net/http"

	"github.com/Lynx/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client     = db.GetDBCli()
	database   *mongo.Database
	collection *mongo.Collection
)

func main() {
	database = client.Database("LINE_LABEL")
	collection = database.Collection("Label")
	mux := &RouteMux{}
	http.ListenAndServe(":9090", mux)
}
