package viewModels

type TaskListModel struct {
	TaskId     string `bson:"taskId" json:"taskId"`
	TaskTitle  string `bson:"taskTitle" json:"taskTitle"`
	Context    string `bson:"context" json:"context"`
	Answered   int    `bson:"answered" json:"answered"`
	IsAnswered bool   `bson:"isAnswered" json:"isAnswered"`
}

type SentiTaskListModel struct {
	TaskId     string   `bson:"taskId" json:"taskId"`
	TaskTitle  string   `bson:"taskTitle" json:"taskTitle"`
	Context    string   `bson:"context" json:"context"`
	AspectPool []string `bson:"aspectPool" json:"aspectPool"`
	IsAnswered bool     `bson:"isAnswered" json:"isAnswered"`
}
type TasksViewModel struct {
	ArticleId    string          `bson:"articleId" json:"articleId"`
	ArticleTitle string          `bson:"articleTitle" json:"articleTitle"`
	TaskType     string          `bson:"taskType" json:"taskType"`
	TaskList     []TaskListModel `bson:"taskList" json:"taskList"`
}

type SentiTasksViewModel struct {
	ArticleId    string               `bson:"articleId" json:"articleId"`
	ArticleTitle string               `bson:"articleTitle" json:"articleTitle"`
	TaskType     string               `bson:"taskType" json:"taskType"`
	TaskList     []SentiTaskListModel `bson:"taskList" json:"taskList"`
}

type QAPairModel struct {
	Question string `bson:"question" json:"question"`
	Answer   string `bson:"answer" json:"answer"`
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
	TaskId     string   `bson:"taskId" json:"taskId"`
	TaskType   string   `bson:"taskType" json:"taskType"`
	TaskTitle  string   `bson:"taskTitle" json:"taskTitle"`
	Context    string   `bson:"context" json:"context"`
	AspectPool []string `bson:"aspectPool" json:"aspectPool"`
	IsAnswered bool     `bson:"isAnswered" json:"isAnswered"`
}
