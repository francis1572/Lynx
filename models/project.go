package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Project struct {
	ProjectId   int    `bson:"projectId" json:"projectId"`
	ProjectName string `bson:"projectName" json:"projectName"`
	ProjectType string `bson:"projectType" json:"projectType"`
	LabelInfo   string `bson:"labelInfo" json:"labelInfo"`
	Manager     string `bson:"manager" json:"manager"`
}

func (p *Project) TableName() string {
	return "Project"
}

func (p *Project) ToQueryBson() bson.M {
	var queryObject bson.M
	if p.ProjectId != 0 {
		queryObject = bson.M{
			"projectId": p.ProjectId,
		}
	} else {
		queryObject = bson.M{
			"projectName": p.ProjectName,
		}
	}
	return queryObject
}
