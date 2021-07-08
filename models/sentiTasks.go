package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type SentiTask struct {
	ArticleId  string   `bson:"articleId" json:"articleId"`
	TaskId     string   `bson:"taskId" json:"taskId"`
	AspectPool []string `bson:"aspectPool" json:"aspectPool"`
	TaskTitle  string   `bson:"taskTitle" json:"taskTitle"`
	Context    string   `bson:"context" json:"context"`
	TaskType   string   `bson:"taskType" json:"taskType"`
	IsAnswered bool     `bson:"isAnswered" json:"isAnswered"`
	IsValidate bool     `bson:"isValidate" json:"isValidate"`
}

func (t *SentiTask) TableName() string {
	return "SentiTask"
}

func (t *SentiTask) ToQueryBson() bson.M {
	var queryObject bson.M
	if t.TaskId != "" {
		queryObject = bson.M{
			"articleId": t.ArticleId,
			"taskId":    t.TaskId,
			"taskType":  t.TaskType,
		}
	} else {
		queryObject = bson.M{
			"articleId": t.ArticleId,
		}
	}
	return queryObject
}

type SentiSentiment struct {
	TaskId    string `bson:"taskId" json:"taskId"`
	AspectId  string `bson:"aspectId" json:"aspectId"`
	Offset    int    `bson:"offset" json:"offset"`
	Sentiment string `bson:"sentiment" json:"sentiment"`
	Dir       string `bson:"dir" json:"dir"`
}

func (t *SentiSentiment) TableName() string {
	return "SentiSentiment"
}

func (a *SentiSentiment) ToQueryBson() bson.M {
	var queryObject bson.M
	if a.AspectId != "" {
		queryObject = bson.M{
			"taskId":   a.TaskId,
			"aspectId": a.AspectId,
		}
	} else {
		queryObject = bson.M{
			"taskId": a.TaskId,
		}
	}

	return queryObject
}

type SentiAspect struct {
	TaskId      string `bson:"taskId" json:"taskId"`
	AspectId    string `bson:"aspectId" json:"aspectId"`
	MajorAspect string `bson:"majorAspect" json:"majorAspect"`
	MinorAspect string `bson:"minorAspect" json:"minorAspect"`
	Offset      int    `bson:"offset" json:"offset"`
	// SentimentList []SentiSentiment `bson:"sentimentList" json:"sentimentList"`
}

func (t *SentiAspect) TableName() string {
	return "SentiAspect"
}

func (a *SentiAspect) ToQueryBson() bson.M {
	queryObject := bson.M{
		"taskId": a.TaskId,
	}
	return queryObject
}

// type SentiAnswer struct {
// 	UserId     string        `bson:"userId" json:"userId"`
// 	ArticleId  string        `bson:"articleId" json:"articleId"`
// 	TaskId     string        `bson:"taskId" json:"taskId"`
// 	TaskType   string        `bson:"taskType" json:"taskType"`
// 	AnswerList []SentiAspect `bson:"answerList" json:"answerList"`
// 	IsValidate bool          `bson:"isValidate" json:"isValidate"`
// }
type SentiAnswer struct {
	Aspect    []SentiAspect    `bson:"aspect" json:"aspect"`
	Sentiment []SentiSentiment `bson:"sentimentt" json:"sentiment"`
}

// func (a *SentiAnswer) ToQueryBson() bson.M {
// 	var queryObject bson.M
// 	if a.UserId == "" {
// 		queryObject = bson.M{
// 			"articleId": a.ArticleId,
// 			"taskId":    a.TaskId,
// 			"taskType":  a.TaskType,
// 		}
// 	} else {
// 		queryObject = bson.M{
// 			"userId":    a.UserId,
// 			"articleId": a.ArticleId,
// 			"taskId":    a.TaskId,
// 		}
// 	}
// 	return queryObject
// }
