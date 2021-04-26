package models

import (
	"log"
	"reflect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Id 				 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
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
	if reflect.ValueOf("id").IsValid() {
		// log.Println("a.Id", reflect.TypeOf(a.Id).Kind())
		log.Println("a.ID", a.Id)
		// id, err := primitive.ObjectIDFromHex(a.Id)
		queryObject = bson.M{ "_id": a.Id }
	} else if a.TaskType == "MRCValidation" {
		queryObject = bson.M{
			"userId": bson.M{"$ne": a.UserId},
			"status": "unverified",
			"taskType":  a.TaskType,
		}
	} else {
		log.Println("here")
		queryObject = bson.M{
			"articleId": a.ArticleId,
			"taskId":    a.TaskId,
			"taskType":  a.TaskType,
		}
	}
	return queryObject
}

type MRCValidation struct {
	LabelUserId 	 	 string `bson:"labelUserId" json:"labelUserId"`
	ValidationUserId string `bson:"validationUserId" json:"validationUserId"`
	OriginalId 			 primitive.ObjectID `bson:"originalId,omitempty" json:"originalId"`
	ValidationId		 primitive.ObjectID `bson:"validationId,omitempty" json:"validationId"`
	Status   		 		 string `bson:"status" json:"status"`
}

func (a *MRCValidation) ToQueryBson() bson.M {
	queryObject := bson.M{
		"status": a.Status,
	}
	return queryObject
}