package main

import (
	"log"
	"net/http"

	"github.com/Lynx/db"
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
	http.ListenAndServe(":9090", mux)
}
