package viewModels

import (
	"Lynx/models"
)

type ProjectViewModel struct {
	ProjectId   int              `bson:"projectId" json:"projectId"`
	ProjectName string           `bson:"projectName" json:"projectName"`
	ProjectType string           `bson:"projectType" json:"projectType"`
	LabelInfo   string           `bson:"labelInfo" json:"labelInfo"`
	ArticleList []models.Article `bson:"articleList" json:"articleList"`
}

type ProjectManageViewModel struct {
	ProjectId  int    `bson:"projectId" json:"projectId"`
	UserId     string `bson:"userId" json:"userId"`
	CodeType   string `bson:"codeType" json:"codeType"`
	StatusCode string `bson:"statusCode" json:"statusCode"`
	Name       string `bson:"name" json:"name"`
	Email      string `bson:"email" json:"email"`
	ImageUrl   string `bson:"imageUrl" json:"imageUrl"`
}
type AddProjectViewModel struct {
	Project models.Project `bson:"project" json:"project"`
	Members []models.Auth  `bson:"members" json:"members"`
	CsvFile [][]string     `bson:"csvFile" json:"csvFile"`
}
