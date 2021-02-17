package respond

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lynx/models"
	"github.com/Lynx/service"
	uuid "github.com/nu7hatch/gouuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetArticles(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo map[string]string
	var articles []models.Article
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	var userId = queryInfo["userId"]
	log.Println(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	articles, err = service.GetArticles(database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	for i, article := range articles {
		// get how many tasks that each article has
		tasks, err := service.GetTasksByArticleId(database, article.ToQueryBson())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		articles[i].TotalTasks = len(tasks)
		for _, task := range tasks {
			answers, err := service.GetAnswers(database, models.MRCAnswer{UserId: userId, TaskId: task.TaskId})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return err
			}
			// get how many tasks user has answered
			articles[i].Answered = len(answers)
		}
	}
	jsondata, _ := json.Marshal(articles)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsondata)
	return nil
}

func GetTasksByArticleId(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	// [TODO]: Change userId to real Login user
	var userId = "tester01"
	var queryTask models.MRCTask
	var result models.TasksViewModel
	// decode request condition to queryTask
	err := json.NewDecoder(r.Body).Decode(&queryTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	// get tasks by articles
	tasks, err := service.GetTasksByArticleId(database, queryTask.ToQueryBson())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("tasks", tasks)

	// get ArticleInfo
	result.ArticleId = queryTask.ArticleId
	result.TaskType = queryTask.TaskType
	articleResult, err := service.GetArticleByArticleId(database, bson.M{"articleId": queryTask.ArticleId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	result.ArticleTitle = articleResult.ArticleTitle

	// get tasksInfo
	for _, task := range tasks {
		var t = models.TaskListModel{
			TaskId:   task.TaskId,
			Context:  task.Context,
			Answered: task.Answered,
		}
		answers, err := service.GetAnswers(database, models.MRCAnswer{UserId: userId, TaskId: task.TaskId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		// if user has answers the question
		if len(answers) != 0 {
			t.IsAnswered = true
		}
		result.TaskList = append(result.TaskList, t)
	}
	jsondata, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsondata)
	return nil
}

func SaveArticles(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	collection := database.Collection("Articles")
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

	// to insert into db, need to convert struct to interface{}
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
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsondata)
	return nil
}

func GetTaskById(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	answers, err := service.GetAnswers(database, models.MRCAnswer{UserId: requestBody["userId"], TaskId: requestBody["taskId"]})
	task, err := service.GetTaskById(database, models.MRCTask{TaskId: requestBody["taskId"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response models.TaskViewModel
	response.TaskId = task.TaskId
	response.TaskType = task.TaskType
	response.TaskTitle = task.TaskTitle
	response.Answered = task.Answered
	response.Context = task.Context

	for _, answer := range answers {
		var QAPair = models.QAPairModel {
			Question: answer.Question,
			Answer: answer.Answer,
		}
		response.QAPairs = append(response.QAPairs, QAPair)
	}

	jsondata, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsondata)
	return nil
}

func SaveAnswer(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.MRCAnswer
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	res, err := service.SaveAnswer(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response = models.Success {
		Success: true,
		Message: res.InsertedID.(primitive.ObjectID).Hex(),
	}
	jsondata, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsondata)
	return nil
}

func Test(w http.ResponseWriter, r *http.Request) error {
	var requestModel []interface{}
	err := json.NewDecoder(r.Body).Decode(&requestModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println(requestModel)
	jsondata, _ := json.Marshal(models.InsertSuccess)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsondata)
	return nil
}
