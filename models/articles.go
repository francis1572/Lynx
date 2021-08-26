package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User structure
type Article struct {
	ArticleId    primitive.ObjectID `bson:"_id" json:"_id"`
	ProjectId    primitive.ObjectID `bson:"projectId" json:"projectId"`
	ArticleTitle string             `bson:"articleTitle" json:"articleTitle"`
	TotalTasks   int                `bson:"totalTasks" json:"totalTasks"`
	Answered     int                `bson:"answered" json:"answered"`
}

//TableName return name of database table
func (a *Article) TableName() string {
	return "Articles"
}

func (a *Article) ToQueryBson() bson.M {
	if a.ArticleId.Hex() != "000000000000000000000000" {
		queryObject := bson.M{
			"articleId": a.ArticleId,
		}
		return queryObject
	} else {
		queryObject := bson.M{
			"projectId": a.ProjectId,
		}
		return queryObject
	}

}

type SentiArticle struct {
	ArticleId    primitive.ObjectID `bson:"_id" json:"_id"`
	ProjectId    primitive.ObjectID `bson:"projectId" json:"projectId"`
	TaskType     string             `bson:"taskType" json:"taskType"`
	ArticleTitle string             `bson:"articleTitle" json:"articleTitle"`
	TotalTasks   int                `bson:"totalTasks" json:"totalTasks"`
	IsAnswered   bool               `bson:"isAnswered" json:"isAnswered"`
	IsValidated  bool               `bson:"isValidated " json:"isValidated"`
}

//TableName return name of database table
func (a *SentiArticle) TableName() string {
	return "Articles"
}

func (a *SentiArticle) ToQueryBson() bson.M {
	if a.ArticleId.Hex() != "000000000000000000000000" {
		queryObject := bson.M{
			"articleId": a.ArticleId,
		}
		return queryObject
	} else {
		queryObject := bson.M{
			"projectId": a.ProjectId,
		}
		return queryObject
	}

}
