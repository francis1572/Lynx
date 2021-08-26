package viewModels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskListModel struct {
	TaskId     string `bson:"taskId" json:"taskId"`
	TaskTitle  string `bson:"taskTitle" json:"taskTitle"`
	Context    string `bson:"context" json:"context"`
	Answered   int    `bson:"answered" json:"answered"`
	IsAnswered bool   `bson:"isAnswered" json:"isAnswered"`
}

type SentiTaskListModel struct {
	TaskId     primitive.ObjectID `bson:"_id" json:"_id"`
	TaskTitle  string             `bson:"taskTitle" json:"taskTitle"`
	Context    string             `bson:"context" json:"context"`
	AspectPool []string           `bson:"aspectPool" json:"aspectPool"`
	IsAnswered bool               `bson:"isAnswered" json:"isAnswered"`
}
type TasksViewModel struct {
	ArticleId    primitive.ObjectID `bson:"articleId" json:"articleId"`
	ArticleTitle string             `bson:"articleTitle" json:"articleTitle"`
	TaskType     string             `bson:"taskType" json:"taskType"`
	TaskList     []TaskListModel    `bson:"taskList" json:"taskList"`
}

type SentiTasksViewModel struct {
	ArticleId    primitive.ObjectID   `bson:"articleId" json:"articleId"`
	ArticleTitle string               `bson:"articleTitle" json:"articleTitle"`
	TaskType     string               `bson:"taskType" json:"taskType"`
	TaskList     []SentiTaskListModel `bson:"taskList" json:"taskList"`
}

type QAPairModel struct {
	Question string `bson:"question" json:"question"`
}

type ValidationQAPairModel struct {
	OriginalId  string `bson:"originalId" json:"originalId"`
	ArticleId   string `bson:"articleId" json:"articleId"`
	TaskId      string `bson:"taskId" json:"taskId"`
	Question    string `bson:"question" json:"question"`
	TaskTitle   string `bson:"taskTitle" json:"taskTitle"`
	TaskContext string `bson:"taskContext" json:"taskContext"`
}

type TaskViewModel struct {
	TaskId    string        `bson:"taskId" json:"taskId"`
	TaskType  string        `bson:"taskType" json:"taskType"`
	TaskTitle string        `bson:"taskTitle" json:"taskTitle"`
	Context   string        `bson:"context" json:"context"`
	Answered  int           `bson:"answered" json:"answered"`
	QAPairs   []QAPairModel `bson:"qaList" json:"qaList"`
}

type SentiTaskViewModel struct {
	TaskId     primitive.ObjectID `bson:"taskId" json:"taskId"`
	ArticleId  primitive.ObjectID `bson:"articleId" json:"articleId"`
	ProjectId  primitive.ObjectID `bson:"projectId" json:"projectId"`
	TaskType   string             `bson:"taskType" json:"taskType"`
	TaskTitle  string             `bson:"taskTitle" json:"taskTitle"`
	Context    string             `bson:"context" json:"context"`
	AspectPool []string           `bson:"aspectPool" json:"aspectPool"`
	IsAnswered bool               `bson:"isAnswered" json:"isAnswered"`
}

type ValidationDataViewModel struct {
	UserId           string `bson:"userId" json:"userId"`
	OriginalId       string `bson:"originalId" json:"original"`
	ValidationAnswer string `bson:"validationAnswer" json:"validationAnswer"`
	StartIdx         int    `bson:"startIdx" json:"startIdx"`
}

type DecisionDataViewModel struct {
	ValidationStatusId  primitive.ObjectID `bson:"validationStatusId" json:"validationStatusId"`
	OriginalId          primitive.ObjectID `bson:"originalId,omitempty" json:"original"`
	ValidationId        primitive.ObjectID `bson:"validationId,omitempty" json:"validationId"`
	Question            string             `bson:"question" json:"question"`
	OriginalAnswer      string             `bson:"originalAnswer" json:"originalAnswer"`
	OriginalStartIdx    int                `bson:"originalStartIdx" json:"originalStart"`
	ValidationAnswer    string             `bson:"validationAnswer" json:"validationAnswer"`
	ValidationStartIdx  int                `bson:"validationStartIdx" json:"validationStartIdx"`
	OriginalTaskContext string             `bson:"originalTaskContext" json:"originalTaskContext"`
}
