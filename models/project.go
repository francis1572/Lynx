package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ProjectId   primitive.ObjectID `bson:"_id" json:"_id"`
	ProjectName string             `bson:"name" json:"name"`
	ProjectType string             `bson:"type" json:"type"`
	Rule        string             `bson:"rule" json:"rule"`
	ManagerId   string             `bson:"managerId" json:"managerId"`
}

func (p *Project) TableName() string {
	return "Project"
}

func (p *Project) ToQueryBson() bson.M {
	var queryObject bson.M
	if p.ProjectId.Hex() != "000000000000000000000000" {
		queryObject = bson.M{
			"_id": p.ProjectId,
		}
	} else {
		queryObject = bson.M{
			"projectName": p.ProjectName,
		}
	}
	return queryObject
}
