package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func DBConnection() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://angela:angyeah6@cluster0.gxqmm.mongodb.net/LINE_LABEL?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}

// getter for client
func GetDBCli() *mongo.Client {
	if client == nil {
		DBConnection()
	}
	// fmt.Println("client", client)
	return client
}
