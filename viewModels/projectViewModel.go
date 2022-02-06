package viewModels

import (
	"Lynx/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectViewModel struct {
	ProjectId   primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	ProjectName string             `bson:"name" json:"name"`
	ProjectType string             `bson:"type" json:"type"`
	Rule        string             `bson:"rule" json:"rule"`
	ManagerId   string             `bson:"managerId" json:"managerId"`
	ArticleList []models.Article   `bson:"articleList" json:"articleList"`
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
