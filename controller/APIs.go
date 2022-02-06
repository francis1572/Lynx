package respond

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"Lynx/models"
	"Lynx/service"
	"Lynx/viewModels"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var user models.User
	var response = models.Success{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	userResult, err := service.GetUser(database, user.ToQueryBson())
	// if no user found then insert a new one and return
	if userResult == nil {
		_, err := service.SaveUser(database, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		// check if insert user
		_, err = service.GetUser(database, user.ToQueryBson())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		response.Success = true
		response.Message = "Add New User"
		jsondata, _ := json.Marshal(response)
		w.Write(jsondata)
		return nil
	}
	// if already has users
	response.Success = true
	response.Message = "User Login"
	jsondata, _ := json.Marshal(response)
	w.Write(jsondata)
	return nil
}

func GetUsersByProject(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo map[string]string
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	projectId, _ := strconv.Atoi(queryInfo["projectId"])
	var queryAuth = models.Auth{ProjectId: projectId}
	auths, err := service.GetAuthByProjectId(database, queryAuth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("Authes", auths)
	var userIds []string
	for _, auth := range auths {
		userIds = append(userIds, auth.UserId)
	}
	users, err := service.GetUsersByIds(database, userIds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	var result = []viewModels.ProjectManageViewModel{}
	for _, user := range users {
		var userAuth = models.Auth{}
		for _, auth := range auths {
			if auth.UserId == user.UserId {
				userAuth = auth
				break
			}
		}
		var userViewModel = viewModels.ProjectManageViewModel{
			ProjectId:  userAuth.ProjectId,
			UserId:     userAuth.UserId,
			CodeType:   userAuth.CodeType,
			StatusCode: userAuth.StatusCode,
			Name:       user.Name,
			Email:      user.Email,
			ImageUrl:   user.ImageUrl,
		}
		result = append(result, userViewModel)
	}

	jsondata, _ := json.Marshal(result)
	w.Write(jsondata)
	return nil
}

func SaveAuth(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.Auth
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	requestBody.CodeType = "1"
	res, err := service.SaveAuth(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response = models.Success{
		Success: true,
		Message: res.InsertedID.(primitive.ObjectID).Hex(),
	}
	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

// func SaveProject(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
// 	var addProjModel = viewModels.AddProjectViewModel{}
// 	err := json.NewDecoder(r.Body).Decode(&addProjModel)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}

// 	// assign new projectId to each member's authentication
// 	currentId, err := service.GetProjectCount(database)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}

// 	// check if csv format decode is correct first
// 	articles, tasks, err := convertToArticlesAndTasks(addProjModel.CsvFile, int(currentId))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}

// 	// then adding objects into db
// 	addProjModel.Project.ProjectId = int(currentId)
// 	projectResult, err := service.SaveProject(database, addProjModel.Project)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	for i := range addProjModel.Members {
// 		addProjModel.Members[i].ProjectId = int(currentId)
// 	}
// 	log.Println("addProj members", addProjModel.Members)
// 	_, err = service.SaveAuths(database, addProjModel.Members)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	_, err = service.SaveArticles(database, articles)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	_, err = service.SaveTasks(database, tasks)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	var response = models.Success{
// 		Success: true,
// 		Message: projectResult.InsertedID.(primitive.ObjectID).Hex(),
// 	}
// 	jsondata, _ := json.Marshal(response)
// 	_, _ = w.Write(jsondata)
// 	return nil
// }

// private func for decode articles and Tasks.
// func convertToArticlesAndTasks(csvFile [][]string, currentId int) ([]models.Article, []models.MRCTask, error) {
// 	var articles []models.Article
// 	var tasks []models.MRCTask
// 	articleSets := make(map[string]int)
// 	var ai int = -1 //point index of articles
// 	for _, pair := range csvFile {
// 		var contexts = pair[0]
// 		var indices = strings.Split(pair[1], "-")
// 		var articleId = "articleId" + indices[0]
// 		var taskId = "taskId" + pair[1]
// 		// basic check for the csv format
// 		if len(pair) != 2 || len(indices) != 2 {
// 			log.Println("錯誤的CSV檔案格式")
// 			return nil, nil, errors.New("錯誤的CSV檔案格式")
// 		}
// 		// check if the articles already exist else create one
// 		if _, ok := articleSets[articleId]; ok {
// 			articles[ai].TotalTasks += 1
// 		} else {
// 			articleSets[articleId] = 1
// 			var article = models.Article{
// 				ArticleId:    articleId,
// 				ProjectId:    currentId,
// 				ArticleTitle: contexts,
// 				TotalTasks:   1,
// 			}
// 			articles = append(articles, article)
// 			ai += 1
// 		}
// 		var task = models.MRCTask{
// 			TaskId:    taskId,
// 			ArticleId: articleId,
// 			TaskType:  "MRC",
// 			TaskTitle: contexts,
// 			Context:   contexts,
// 		}
// 		tasks = append(tasks, task)
// 	}
// 	return articles, tasks, nil
// }

func GetUsers(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	users, err := service.GetUsers(database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	jsondata, _ := json.Marshal(users)
	w.Write(jsondata)
	return nil
}

// func GetProjects(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
// 	var queryInfo map[string]string
// 	err := json.NewDecoder(r.Body).Decode(&queryInfo)
// 	var userId = queryInfo["userId"]
// 	var statusCode = queryInfo["statusCode"]
// 	var queryAuth = models.Auth{UserId: userId, StatusCode: statusCode}
// 	log.Println(userId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}

// 	projects, err := service.GetProjectsByAuth(database, queryAuth)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	jsondata, _ := json.Marshal(projects)
// 	w.Write(jsondata)
// 	return nil
// }

// func GetArticles(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
// 	var queryInfo map[string]string
// 	var result viewModels.ProjectViewModel
// 	var articles []models.Article
// 	err := json.NewDecoder(r.Body).Decode(&queryInfo)
// 	var userId = queryInfo["userId"]
// 	log.Println(userId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	projectId, _ := strconv.Atoi(queryInfo["projectId"])
// 	var queryProject = models.Project{ProjectId: projectId}
// 	projectResult, err := service.GetProjectByProjectId(database, queryProject)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	articles, err = service.GetArticlesByProjectId(database, projectId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	result.ProjectId = projectResult.ProjectId
// 	result.ProjectName = projectResult.ProjectName
// 	result.ProjectType = projectResult.ProjectType
// 	result.LabelInfo = projectResult.LabelInfo
// 	result.ArticleList = articles
// 	// [PENDING] Further information for articles
// 	// for i, article := range articles {
// 	// 	// get how many tasks that each article has
// 	// 	tasks, err := service.GetTasksByArticleId(database, article.ToQueryBson())
// 	// 	if err != nil {
// 	// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	// 		return err
// 	// 	}
// 	// 	articles[i].TotalTasks = len(tasks)
// 	// 	for _, task := range tasks {
// 	// 		answers, err := service.GetAnswers(database, models.MRCAnswer{UserId: userId, TaskId: task.TaskId})
// 	// 		if err != nil {
// 	// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 	// 			return err
// 	// 		}
// 	// 		// get how many tasks user has answered
// 	// 		articles[i].Answered = len(answers)
// 	// 	}
// 	// }
// 	jsondata, _ := json.Marshal(result)
// 	w.Write(jsondata)
// 	return nil
// }

// func GetTasksByArticleId(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
// 	var queryInfo models.SentiArticle
// 	var result viewModels.TasksViewModel
// 	// decode request condition to queryInfo
// 	err := json.NewDecoder(r.Body).Decode(&queryInfo)
// 	log.Println("GetTasksByArticleId queryInfo:", queryInfo)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	// get tasks by articles
// 	tasks, err := service.GetTasksByArticleId(database, bson.M{"articleId": queryInfo.ArticleId})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}

// 	// get ArticleInfo
// 	result.ArticleId = queryInfo.ArticleId
// 	result.TaskType = queryInfo.TaskType
// 	articleResult, err := service.GetArticleByArticleId(database, bson.M{"articleId": queryInfo.ArticleId})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	result.ArticleTitle = articleResult.ArticleTitle

// 	// get tasksInfo
// 	for _, task := range tasks {
// 		var t = viewModels.TaskListModel{
// 			TaskId:    task.TaskId,
// 			TaskTitle: task.TaskTitle,
// 			Context:   task.Context,
// 		}
// 		answers, err := service.GetAnswers(database, models.MRCAnswer{UserId: queryInfo.UserId, ArticleId: queryInfo["articleId"], TaskId: task.TaskId})
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return err
// 		}
// 		t.Answered = len(answers)
// 		// if user has answers the question
// 		if len(answers) != 0 {
// 			t.IsAnswered = true
// 		}
// 		result.TaskList = append(result.TaskList, t)
// 	}
// 	jsondata, _ := json.Marshal(result)
// 	w.Write(jsondata)
// 	return nil
// }

//[TODO] update MRCTask for answered + 1
// func SaveArticles(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
// 	collection := database.Collection("Articles")
// 	var articles []models.Article
// 	err := json.NewDecoder(r.Body).Decode(&articles)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return err
// 	}
// 	for i, _ := range articles {
// 		articleId, _ := uuid.NewV4()
// 		articles[i].ArticleId = "articleId" + articleId.String()
// 	}
// 	log.Println(articles)

// 	// to insert into db, need to convert struct to interface{}
// 	docs := make([]interface{}, len(articles))
// 	for i, u := range articles {
// 		docs[i] = u
// 	}
// 	articleResult, err := collection.InsertMany(context.Background(), docs)
// 	if err != nil {
// 		log.Fatal(err)
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}
// 	log.Printf("Inserted %v documents into articles collection!\n", len(articleResult.InsertedIDs))
// 	jsondata, _ := json.Marshal(models.InsertSuccess)
// 	_, _ = w.Write(jsondata)
// 	return nil
// }

func GetTaskById(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("GetTaskById queryInfo:", requestBody)
	answers, err := service.GetAnswers(database, models.MRCAnswer{ArticleId: requestBody["articleId"], TaskId: requestBody["taskId"], TaskType: requestBody["taskType"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	task, err := service.GetTaskById(database, models.MRCTask{ArticleId: requestBody["articleId"], TaskId: requestBody["taskId"], TaskType: requestBody["taskType"]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response viewModels.TaskViewModel
	response.TaskId = task.TaskId
	response.TaskType = task.TaskType
	response.TaskTitle = task.TaskTitle
	response.Answered = task.Answered
	response.Context = task.Context

	for _, answer := range answers {
		var QAPair = viewModels.QAPairModel{
			Question: answer.Question,
		}
		response.QAPairs = append(response.QAPairs, QAPair)
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func SaveAnswer(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.MRCAnswer
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	res, err := service.SaveAnswer(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response = models.Success{
		Success: true,
		Message: res.InsertedID.(primitive.ObjectID).Hex(),
	}
	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func Test(w http.ResponseWriter, r *http.Request) error {
	var requestModel []interface{}
	err := json.NewDecoder(r.Body).Decode(&requestModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println(requestModel)
	jsondata, _ := json.Marshal(models.InsertSuccess)
	_, _ = w.Write(jsondata)
	return nil
}

func GetValidation(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo map[string]string
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	var userId = queryInfo["userId"]
	log.Println("userId", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	questionPair, err := service.GetRandomValidationQuestion(database, models.MRCAnswer{UserId: queryInfo["userId"], TaskType: queryInfo["taskType"]})
	if questionPair == nil {
		var response = models.Success{
			Success: true,
			Message: "no validation",
		}
		jsondata, _ := json.Marshal(response)
		w.Write(jsondata)
		return nil
	}
	task, err := service.GetTaskById(database, models.MRCTask{ArticleId: questionPair.ArticleId, TaskId: questionPair.TaskId, TaskType: questionPair.TaskType})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response viewModels.ValidationQAPairModel
	response.OriginalId = questionPair.Id.Hex()
	response.Question = questionPair.Question
	response.ArticleId = questionPair.ArticleId
	response.TaskId = questionPair.TaskId
	response.TaskTitle = task.TaskTitle
	response.TaskContext = task.Context
	jsondata, _ := json.Marshal(response)
	w.Write(jsondata)
	return nil
}

func SaveValidation(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	// Decode
	var queryInfo map[string]string
	log.Println("originalId", queryInfo["originalId"])
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	log.Println("queryInfo validationAnswer", len(queryInfo["validationAnswer"]))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	// find original answer
	id, err := primitive.ObjectIDFromHex(queryInfo["originalId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	res, err := service.FindAnswerById(database, id)
	log.Println("original answer", res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	// save validation answer
	var validationAnswer models.MRCAnswer
	validationAnswer.UserId = queryInfo["userId"]
	validationAnswer.ArticleId = res.ArticleId
	validationAnswer.TaskId = res.TaskId
	validationAnswer.TaskType = "Validation"
	validationAnswer.Status = "unverified"
	validationAnswer.Question = res.Question
	validationAnswer.Answer = queryInfo["validationAnswer"]
	startIdx, err := strconv.Atoi(queryInfo["startIdx"])
	validationAnswer.StartIdx = startIdx
	result, err := service.SaveAnswer(database, validationAnswer)
	log.Println("save result", result.InsertedID)
	validationId := result.InsertedID.(primitive.ObjectID)

	// check validation
	var validationStatus models.MRCValidation
	validationStatus.LabelUserId = res.UserId
	validationStatus.ValidationUserId = queryInfo["userId"]
	validationStatus.OriginalId = id
	validationStatus.ValidationId = validationId
	if res.Answer == queryInfo["validationAnswer"] && queryInfo["startIdx"] == strconv.Itoa(res.StartIdx) {
		validationStatus.Status = "verified"
	} else {
		validationStatus.Status = "pending"
	}
	log.Println("status info", validationStatus)
	statusResult, err := service.SaveValidationStatus(database, validationStatus)
	log.Println("status result", statusResult)

	// update answer
	updateResult, err := service.UpdateAnswer(database, validationStatus)
	log.Println("update result", updateResult)

	// result and response
	var response = models.Success{
		Success: true,
		Message: statusResult.InsertedID.(primitive.ObjectID).Hex(),
	}
	jsondata, _ := json.Marshal(response)
	w.Write(jsondata)
	return nil
}

func GetDecision(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo map[string]string
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	var userId = queryInfo["userId"]
	log.Println("userId", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	decisionInfo, err := service.GetRandomDecisionInfo(database, userId)
	log.Println("decisionInfo", decisionInfo)
	if decisionInfo == nil {
		var response = models.Success{
			Success: true,
			Message: "no decision",
		}
		jsondata, _ := json.Marshal(response)
		w.Write(jsondata)
		return nil
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	originalAnswer, err := service.FindAnswerById(database, decisionInfo.OriginalId)
	log.Println("originalAnswer", originalAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	validationAnswer, err := service.FindAnswerById(database, decisionInfo.ValidationId)
	log.Println("validationAnswer", validationAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	task, err := service.GetTaskById(database, models.MRCTask{ArticleId: originalAnswer.ArticleId, TaskId: originalAnswer.TaskId, TaskType: "MRC"})
	log.Println("task", task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response viewModels.DecisionDataViewModel
	response.ValidationStatusId = decisionInfo.Id
	response.OriginalId = originalAnswer.Id
	response.ValidationId = validationAnswer.Id
	response.Question = originalAnswer.Question
	response.OriginalAnswer = originalAnswer.Answer
	response.OriginalStartIdx = originalAnswer.StartIdx
	response.ValidationAnswer = validationAnswer.Answer
	response.ValidationStartIdx = validationAnswer.StartIdx
	response.OriginalTaskContext = task.Context
	jsondata, _ := json.Marshal(response)
	w.Write(jsondata)
	return nil
}

func SaveDecision(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo map[string]string
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	// update answer
	var decisionStatus models.MRCValidation
	originalId, err := primitive.ObjectIDFromHex(queryInfo["originalId"])
	decisionStatus.OriginalId = originalId
	decisionStatus.Status = queryInfo["status"]
	updateAnswer, err := service.UpdateAnswer(database, decisionStatus)
	log.Println("updateAnswer", updateAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	// update status
	var validationStatus models.MRCValidation
	validationStatusId, err := primitive.ObjectIDFromHex(queryInfo["validationStatusId"])
	validationStatus.OriginalId = validationStatusId
	validationStatus.Status = queryInfo["status"]
	updateErr := service.UpdateValidationStatus(database, validationStatus)
	if updateErr != nil {
		http.Error(w, updateErr.Error(), http.StatusBadRequest)
		return updateErr
	}

	// save decision result
	var decisionResult models.MRCDecision
	decisionResult.UserId = queryInfo["userId"]
	decisionResult.OriginalId = originalId
	validationId, err := primitive.ObjectIDFromHex(queryInfo["validationId"])
	decisionResult.ValidationId = validationId
	decisionResult.ValidationStatusId = validationStatusId
	decisionResult.DecisionResult = queryInfo["decisionResult"]
	saveDecisionResult, err := service.SaveDecision(database, decisionResult)

	var response = models.Success{
		Success: true,
		Message: saveDecisionResult.InsertedID.(primitive.ObjectID).Hex(),
	}
	jsondata, _ := json.Marshal(response)
	w.Write(jsondata)
	return nil
}

//================================= sentiment API =================================
func GetSentiArticles(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo models.SentiArticle

	var articles []models.Article
	var result viewModels.ProjectViewModel
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	// var userId = queryInfo["userId"]
	// log.Println(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	projectId := queryInfo.ProjectId

	var queryProject = models.Project{ProjectId: projectId}
	// log.Print(queryProject)
	projectResult, err := service.GetProjectByProjectId(database, queryProject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	articles, err = service.GetSentiArticles(database, queryProject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	result.ProjectId = projectResult.ProjectId
	result.ProjectName = projectResult.ProjectName
	result.ProjectType = projectResult.ProjectType
	result.Rule = projectResult.Rule
	result.ManagerId = projectResult.ManagerId
	result.ArticleList = articles

	jsondata, _ := json.Marshal(result)
	w.Write(jsondata)
	return nil
}

func GetSentiTasksByArticleId(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo models.SentiArticle
	var result viewModels.SentiTasksViewModel
	// decode request condition to queryInfo
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	log.Println("GetSentiTasksByArticleId queryInfo:", bson.M{"articleId": queryInfo.ArticleId, "isAnswered": false})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	// get tasks by articles
	tasks, err := service.GetSentiTasksByArticleId(database, bson.M{"articleId": queryInfo.ArticleId, "isAnswered": false})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println(tasks)

	// get ArticleInfo
	result.ArticleId = queryInfo.ArticleId
	result.TaskType = queryInfo.TaskType
	articleResult, err := service.GetSentiArticleByArticleId(database, bson.M{"_id": queryInfo.ArticleId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	result.ArticleTitle = articleResult.ArticleTitle

	// get tasksInfo
	for _, task := range tasks {
		var t = viewModels.SentiTaskListModel{
			TaskId:     task.TaskId,
			TaskTitle:  task.TaskTitle,
			Context:    task.Context,
			AspectPool: task.AspectPool,
		}

		answers, err := service.GetAspectByTaskId(database, models.SentiAspect{TaskId: task.TaskId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		// if user has answers the question
		if len(answers) != 0 {
			t.IsAnswered = true
		}

		result.TaskList = append(result.TaskList, t)
	}
	jsondata, _ := json.Marshal(result)
	w.Write(jsondata)
	return nil
}

func GetSentiTaskById(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.SentiTask
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	aspects, err := service.GetAspectByTaskId(database, models.SentiAspect{TaskId: requestBody.TaskId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println("GetSentiTaskById queryInfo:", bson.M{"_id": requestBody.TaskId, "taskType": requestBody.TaskType})
	task, err := service.GetSentiTaskById(database, bson.M{"_id": requestBody.TaskId, "taskType": requestBody.TaskType})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	var response viewModels.SentiTaskViewModel
	// var response viewModels.SentiTasksViewModel
	response.TaskId = task.TaskId
	response.TaskType = task.TaskType
	response.TaskTitle = task.TaskTitle
	response.Context = task.Context
	response.AspectPool = task.AspectPool
	if len(aspects) != 0 {
		response.IsAnswered = true
	}

	// for _, answer := range answers {
	// 	var QAPair = viewModels.QAPairModel{
	// 		Question: answer.Question,
	// 		Answer:   answer.Answer,
	// 	}
	// 	response.QAPairs = append(response.QAPairs, QAPair)
	// }

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}

func SaveSentiAnswer(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.SentiAnswer

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	_, err = service.SaveSentiAnswer(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	// var response = models.Success{
	// 	Success: true,
	// 	// Message: res.InsertedID.(primitive.ObjectID).Hex(),
	// 	Message: res.InsertedIDs,
	// }
	// jsondata, _ := json.Marshal(response)
	// _, _ = w.Write(jsondata)
	return nil
}

// validation 才會用到的

func PostSentiValidation(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.SentiAnswer
	var finalAnswer models.SentiAnswer
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody.Task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	log.Println("GetSentiAnsById queryInfo:", requestBody)

	sentiVal := requestBody.Sentiment
	sentiList, err := service.GetSentiAnswer(database, models.SentiSentiment{TaskId: sentiVal[0].TaskId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	// log.Println("sentiList:", sentiList)
	allMatch := 1
	for _, sentiValItem := range sentiVal {
		isMatch := 0
		// log.Println("sentiVal:", sentiValItem)
		for _, sentiListItem := range sentiList {

			// log.Println("sentiList:", sentiListItem)
			if (sentiValItem.AspectId == sentiListItem.AspectId) && (sentiValItem.Sentiment == sentiValItem.Sentiment) && (sentiValItem.Dir == sentiListItem.Dir) && (sentiValItem.Offset == sentiListItem.Offset) {
				isMatch = 1
			}
		}
		if isMatch == 0 {
			allMatch = 0
			break
		}
	}
	finalAnswer.Task = requestBody.Task
	finalAnswer.Aspect = requestBody.Aspect
	finalAnswer.Sentiment = requestBody.Sentiment
	finalAnswer.ProjectId = requestBody.ProjectId

	if allMatch == 1 {
		log.Println("successful validation")
		finalAnswer.State = "All Match"

	} else {
		log.Println("failed validation")
		finalAnswer.State = "Not Match"
	}

	_, err = service.SaveFinalAnswer(database, finalAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	jsondata, _ := json.Marshal(allMatch)
	_, _ = w.Write(jsondata)

	return err
}

func CheckIsAnswered(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.SentiTask

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	cnt, err := service.CheckIsAnswered(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	jsondata, _ := json.Marshal(cnt)
	_, _ = w.Write(jsondata)

	return nil
}

func CheckIsValidated(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.SentiTask

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	cnt, err := service.CheckIsValidated(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	jsondata, _ := json.Marshal(cnt)
	_, _ = w.Write(jsondata)

	return nil
}

func GetSentiValidation(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo models.SentiTask
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	task, err := service.GetRandomSentiTask(database, models.SentiTask{TaskType: queryInfo.TaskType, ProjectId: queryInfo.ProjectId})
	if task == nil {
		var response = models.Success{
			Success: true,
			Message: "no validation",
		}
		jsondata, _ := json.Marshal(response)
		w.Write(jsondata)
		return nil
	}

	aspects, err := service.GetAspectByTaskId(database, models.SentiAspect{TaskId: task.TaskId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	var response viewModels.SentiTaskViewModel
	response.TaskId = task.TaskId
	response.TaskType = task.TaskType
	response.TaskTitle = task.TaskTitle
	response.ArticleId = task.ArticleId
	response.ProjectId = task.ProjectId
	response.Context = task.Context
	response.AspectPool = task.AspectPool

	if len(aspects) != 0 {
		response.IsAnswered = true
	}

	jsondata, _ := json.Marshal(response)
	_, _ = w.Write(jsondata)
	return nil
}
func GetSentiAspects(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var queryInfo models.SentiTask
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println(models.SentiAspect{TaskId: queryInfo.TaskId})
	aspects, err := service.GetAspectByTaskId(database, models.SentiAspect{TaskId: queryInfo.TaskId})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	log.Println(aspects)

	jsondata, _ := json.Marshal(aspects)
	_, _ = w.Write(jsondata)
	return nil
}

func DiscardSentiAnswer(database *mongo.Database, w http.ResponseWriter, r *http.Request) error {
	var requestBody models.SentiTask

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	// log.Println(requestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	cnt, err := service.DiscardSentiAnswer(database, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	jsondata, _ := json.Marshal(cnt)
	_, _ = w.Write(jsondata)

	return nil
}
