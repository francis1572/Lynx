package main

import (
	"log"
	"net/http"

	"Lynx/db"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client   = db.GetDBCli()
	Database *mongo.Database
)

func main() {
	Database = client.Database("LINE_LABEL")
	mux := &RouteMux{}
	log.Println("Server Launched on port 9090")
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":9090", handler)
}
