package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

//User structure
type Label struct {
	TaskId     string
	TaskType   string
	ImagePath  string
	TrueAnswer string
	Example    bool
	LabelId    string
}

//TableName return name of database table
func (u *Label) TableName() string {
	return "Label"
}

func (u *Label) ToQueryBson() bson.M {
	queryObject := bson.M{
		"taskId": u.TaskId,
	}
	return queryObject
}
