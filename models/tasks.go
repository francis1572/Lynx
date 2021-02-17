package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MRCTask struct {
	TaskId    string `bson:"taskId" json:"taskId"`
	ArticleId string `bson:"articleId" json:"articleId"`
	TaskType  string `bson:"taskType" json:"taskType"`
	TaskTitle  string `bson:"taskTitle" json:"taskTitle"`
	Context   string `bson:"context" json:"context"`
	Answered  int    `bson:"answered" json:"answered"`
}

func (t *MRCTask) TableName() string {
	return "MRCTask"
}

func (t *MRCTask) ToQueryBson() bson.M {
	var queryObject bson.M
	if t.ArticleId != "" {
		queryObject = bson.M {
			"articleId": t.ArticleId,
		}
	} else {
		queryObject = bson.M{
			"taskId": t.TaskId,
		}
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
	var queryObject bson.M
	if a.UserId == "" {
		queryObject = bson.M {
			"taskId": a.TaskId,
		}
	} else {
		queryObject = bson.M{
			"userId": a.UserId,
			"taskId": a.TaskId,
		}
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

type QAPairModel struct {
	Question string `bson:"question" json:"question"`
	Answer   string `bson:"answer" json:"answer"`
}

type TaskViewModel struct {
	TaskId    string `bson:"taskId" json:"taskId"`
	TaskType  string `bson:"taskType" json:"taskType"`
	TaskTitle  string `bson:"taskTitle" json:"taskTitle"`
	Context   string `bson:"context" json:"context"`
	Answered  int    `bson:"answered" json:"answered"`
	QAPairs []QAPairModel `bson:"qaList" json:"qaList"`
}
