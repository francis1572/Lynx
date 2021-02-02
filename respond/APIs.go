package respond

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lynx/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTaskByTaskId(collection *mongo.Collection, w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	if r.Form["taskId"] == nil {
		http.Error(w, "No taskId", http.StatusInternalServerError)
		return nil
	}
	var taskId = r.FormValue("taskId")
	queryModel := models.Label{TaskId: taskId}
	var labelResult []models.Label
	cur, err := collection.Find(context.Background(), queryModel.ToQueryBson())
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	// defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		result := models.Label{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return err
		}
		labelResult = append(labelResult, result)
	}
	jsondata, _ := json.Marshal(labelResult)
	w.Write(jsondata)
	return nil
}

func GetArticles(collection *mongo.Collection, w http.ResponseWriter, r *http.Request) error {
	// todo: change to request parse
	queryModel := models.Label{TaskId: "taskId6a458bc6800048a7"}
	var labelResult []models.Label
	cur, err := collection.Find(context.Background(), queryModel.ToQueryBson())
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	// defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		result := models.Label{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return err
		}
		labelResult = append(labelResult, result)
	}
	jsondata, _ := json.Marshal(labelResult)
	w.Write(jsondata)
	return nil
}
