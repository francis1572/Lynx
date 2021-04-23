package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MRCTask struct {
	TaskId    string `bson:"taskId" json:"taskId"`
	ArticleId string `bson:"articleId" json:"articleId"`
	TaskType  string `bson:"taskType" json:"taskType"`
	TaskTitle string `bson:"taskTitle" json:"taskTitle"`
	Context   string `bson:"context" json:"context"`
	Answered  int    `bson:"answered" json:"answered"`
}

func (t *MRCTask) TableName() string {
	return "MRCTask"
}

func (t *MRCTask) ToQueryBson() bson.M {
	var queryObject bson.M
	if t.TaskId != "" {
		queryObject = bson.M{
			"articleId": t.ArticleId,
			"taskId":    t.TaskId,
			"taskType":  "MRC",
		}
	} else {
		queryObject = bson.M{
			"articleId": t.ArticleId,
		}
	}
	return queryObject
}

type MRCAnswer struct {
	UserId     string `bson:"userId" json:"userId"`
	ArticleId  string `bson:"articleId" json:"articleId"`
	TaskId     string `bson:"taskId" json:"taskId"`
	TaskType   string `bson:"taskType" json:"taskType"`
	Status 	   string `bson:"status" json:"status"`
	Question   string `bson:"question" json:"question"`
	Answer     string `bson:"answer" json:"answer"`
	StartIdx   int    `bson:"startIdx" json:"startIdx"`
}

func (a *MRCAnswer) ToQueryBson() bson.M {
	var queryObject bson.M
	if a.TaskType == "MRCValidation" {
		queryObject = bson.M{
			"userId": bson.M{"$ne": a.UserId},
			"status": "unverified",
			"taskType":  a.TaskType,
		}
	} else {
		queryObject = bson.M{
			"articleId": a.ArticleId,
			"taskId":    a.TaskId,
			"taskType":  a.TaskType,
		}
	}
	return queryObject
}
