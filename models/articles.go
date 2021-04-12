package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

//User structure
type Article struct {
	ArticleId    string `bson:"articleId" json:"articleId"`
	ArticleTitle string `bson:"articleTitle" json:"articleTitle"`
	TotalTasks   int    `bson:"totalTasks" json:"totalTasks"`
	Answered     int    `bson:"answered" json:"answered"`
	LabelInfo    string `bson:"labelInfo" json:"labelInfo"`
}

//TableName return name of database table
func (a *Article) TableName() string {
	return "Articles"
}

func (a *Article) ToQueryBson() bson.M {
	queryObject := bson.M{
		"articleId": a.ArticleId,
	}
	return queryObject
}
