package respond

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lynx/models"
	"github.com/mitchellh/mapstructure"
	uuid "github.com/nu7hatch/gouuid"
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

func GetArticles(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	collection := database.Collection("Articles")
	var article models.Article
	var articles []models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	cur, err := collection.Find(context.Background(), article.ToQueryBson())
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	// defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		result := models.Article{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return err
		}
		articles = append(articles, result)
	}
	jsondata, _ := json.Marshal(articles)
	w.Write(jsondata)
	return nil
}

// redundunt transition
func SaveArticlesV1(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	collection := database.Collection("Articles")
	var dataList models.Enumerable
	var articles []models.Article
	err := json.NewDecoder(r.Body).Decode(&dataList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	for _, item := range dataList.DataList {
		var article models.Article
		mapstructure.Decode(item, &article)
		articleId, _ := uuid.NewV4()
		article.ArticleId = "articleId" + articleId.String()
		articles = append(articles, article)
	}
	log.Println(articles)

	docs := make([]interface{}, len(articles))
	for i, u := range articles {
		docs[i] = u
	}
	articleResult, err := collection.InsertMany(context.Background(), docs)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("Inserted %v documents into articles collection!\n", len(articleResult.InsertedIDs))
	jsondata, _ := json.Marshal(models.Success{Success: 1, Message: "insert success"})
	w.Write(jsondata)
	return nil
}

func SaveArticles(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	collection := database.Collection("Articles")
	// var dataList models.Enumerable
	var articles []models.Article
	err := json.NewDecoder(r.Body).Decode(&articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	for i, _ := range articles {
		articleId, _ := uuid.NewV4()
		articles[i].ArticleId = "articleId" + articleId.String()
	}
	log.Println(articles)

	docs := make([]interface{}, len(articles))
	for i, u := range articles {
		docs[i] = u
	}
	articleResult, err := collection.InsertMany(context.Background(), docs)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("Inserted %v documents into articles collection!\n", len(articleResult.InsertedIDs))
	jsondata, _ := json.Marshal(models.InsertSuccess)
	_, _ = w.Write(jsondata)
	return nil
}
