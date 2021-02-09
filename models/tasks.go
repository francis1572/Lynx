package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MRCTask struct {
	TaskId    string `bson:"taskId" json:"taskId"`
	ArticleId string `bson:"articleId" json:"articleId"`
	TaskType  string `bson:"taskType" json:"taskType"`
	Context   string `bson:"context" json:"context"`
	Answered  int    `bson:"answered" json:"answered"`
}

func (t *MRCTask) TableName() string {
	return "MRCTask"
}

func (t *MRCTask) ToQueryBson() bson.M {
	queryObject := bson.M{
		"articleId": t.ArticleId,
	}
	return queryObject
}

type MRCAnswer struct {
	UserId     string `bson:"userId" json:"userId"`
	TaskId     string `bson:"taskId" json:"taskId"`
	TaskType   string `bson:"taskType" json:"taskType"`
	IsValidate bool   `bson:"isValidate" json:"isValidate"`
	Question   string `bson:"question" json:"question"`
	Answer     string `bson:"answer" json:"answer"`
}

func (a *MRCAnswer) ToQueryBson() bson.M {
	queryObject := bson.M{
		"userId": a.UserId,
		"taskId": a.TaskId,
	}
	return queryObject
}

type TaskListModel struct {
	TaskId     string `bson:"taskId" json:"taskId"`
	Context    string `bson:"context" json:"context"`
	Answered   int    `bson:"answered" json:"answered"`
	IsAnswered bool   `bson:"isAnswered" json:"isAnswered"`
}

type TasksViewModel struct {
	ArticleId    string          `bson:"articleId" json:"articleId"`
	ArticleTitle string          `bson:"articleTitle" json:"articleTitle"`
	TaskType     string          `bson:"taskType" json:"taskType"`
	TaskList     []TaskListModel `bson:"taskList" json:"taskList"`
}
