package respond

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lynx/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLabelsByTaskId(collection *mongo.Collection, w http.ResponseWriter, r *http.Request) error {
	// todo: change to request parse
	queryModel := models.Label{TaskId: "taskId6a458bc6800048a7"}
	var labelResult []models.Label
	cur, err := collection.Find(context.Background(), queryModel.ToQueryBson())
	if err != nil {
		log.Fatal(err)
	}
	// defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		result := models.Label{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		labelResult = append(labelResult, result)
	}
	jsondata, _ := json.Marshal(labelResult)
	w.Write(jsondata)
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}
