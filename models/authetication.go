package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

//User structure
type Auth struct {
	ProjectId  int    `bson:"projectId" json:"projectId"`
	UserId     string `bson:"userId" json:"userId"`
	CodeType   string `bson:"codeType" json:"codeType"`
	StatusCode string `bson:"statusCode" json:"statusCode"`
}

type Auths []Auth

//TableName return name of database table
func (u *Auth) TableName() string {
	return "Authentication"
}

func (u *Auth) ToQueryBson() bson.M {
	queryObject := bson.M{
		"userId":     u.UserId,
		"statusCode": u.StatusCode,
	}
	return queryObject
}

func (a Auths) SelectProjectIdList() []int {
	var list []int
	for _, user := range a {
		list = append(list, user.ProjectId)
	}
	return list
}
