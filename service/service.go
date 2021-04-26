package service

import (
	"context"
	"log"
	"time"
	"reflect"

	"Lynx/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAuthByProjectId(db *mongo.Database, auth models.Auth) ([]models.Auth, error) {
	collection := db.Collection(auth.TableName())
	var serviceResult = []models.Auth{}
	cur, err := collection.Find(context.Background(), bson.M{"projectId": auth.ProjectId})
	if err != nil {
		log.Println("Find Project Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Auth{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Auth Error", err)
			return nil, err
		}
		serviceResult = append(serviceResult, result)
	}
	return serviceResult, nil
}

func GetProjectsByAuth(db *mongo.Database, auth models.Auth) ([]models.Project, error) {
	collection := db.Collection(auth.TableName())
	var authList = models.Auths{}
	cur, err := collection.Find(context.Background(), auth.ToQueryBson())
	log.Println("cur", cur)
	log.Println("auth", auth.ToQueryBson())
	if err != nil {
		log.Println("Find Project Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Auth{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Project Error", err)
			return nil, err
		}
		authList = append(authList, result)
	}
	var projectIds = authList.SelectProjectIdList()
	var serviceResult = []models.Project{}
	for _, projectId := range projectIds {
		project, err := GetProjectByProjectId(db, models.Project{ProjectId: projectId})
		if err != nil {
			return nil, err
		}
		serviceResult = append(serviceResult, *project)
	}

	return serviceResult, nil
}

func SaveAuth(db *mongo.Database, auth models.Auth) (*mongo.InsertOneResult, error) {
	AuthCollection := db.Collection(auth.TableName())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := AuthCollection.InsertOne(ctx, auth)
	if err != nil {
		log.Println("Insert auth Error", err)
		return nil, err
	}
	return res, nil
}

func GetProjects(db *mongo.Database, userId string) ([]models.Project, error) {
	collection := db.Collection("Project")
	var serviceResult = []models.Project{}
	cur, err := collection.Find(context.Background(), bson.M{"manager": userId})
	if err != nil {
		log.Println("Find Project Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Project{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Project Error", err)
			return nil, err
		}
		serviceResult = append(serviceResult, result)
	}
	return serviceResult, nil
}

func GetProjectByProjectId(db *mongo.Database, project models.Project) (*models.Project, error) {
	collection := db.Collection("Project")
	var serviceResult = models.Project{}
	cur := collection.FindOne(context.Background(), project.ToQueryBson())
	// if no project then return nil
	if cur.Err() != nil {
		log.Println("Can't find project in DB")
		return nil, cur.Err()
	}
	// if has project then return
	err := cur.Decode(&serviceResult)
	if err != nil {
		log.Println("Decode project Error", err)
		return nil, err
	}
	log.Println("Get project:", serviceResult)
	return &serviceResult, nil
}

func GetUsersByIds(db *mongo.Database, userIds []string) ([]models.User, error) {
	collection := db.Collection("GUser")
	var serviceResult = []models.User{}
	log.Println("userIds", userIds)
	cur, err := collection.Find(context.Background(), bson.M{"userId": bson.M{"$in": userIds}})
	if err != nil {
		log.Println("Find Users Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.User{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Users Error", err)
			return nil, err
		}
		serviceResult = append(serviceResult, result)
	}
	return serviceResult, nil
}

func GetUser(db *mongo.Database, queryBson bson.M) (*models.User, error) {
	collection := db.Collection("GUser")
	var serviceResult = models.User{}
	cur := collection.FindOne(context.Background(), queryBson)
	// if no user then return nil
	if cur.Err() != nil {
		log.Println("Can't find user in DB")
		return nil, cur.Err()
	}
	// if has user then return
	err := cur.Decode(&serviceResult)
	if err != nil {
		log.Println("Decode user Error", err)
		return nil, err
	}
	log.Println("Get user:", serviceResult)
	return &serviceResult, nil
}

func GetUsers(db *mongo.Database) ([]models.User, error) {
	collection := db.Collection("GUser")
	var serviceResult = []models.User{}
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Find user Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.User{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Article Error", err)
			return nil, err
		}
		serviceResult = append(serviceResult, result)
	}
	return serviceResult, nil
}

func SaveUser(db *mongo.Database, user models.User) (*mongo.InsertOneResult, error) {
	UserCollection := db.Collection("GUser")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Insert user Error", err)
		return nil, err
	}
	log.Println("Insert user Success", user)
	return res, nil
}

func GetArticles(db *mongo.Database) ([]models.Article, error) {
	collection := db.Collection("Articles")
	var articles = []models.Article{}
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Find Articles Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Article{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Article Error", err)
			return nil, err
		}
		articles = append(articles, result)
	}
	return articles, nil
}

// return only one article
func GetArticleByArticleId(db *mongo.Database, queryBson bson.M) (*models.Article, error) {
	ArticleCollection := db.Collection("Articles")
	var serviceResult models.Article
	cur := ArticleCollection.FindOne(context.Background(), queryBson)
	err := cur.Decode(&serviceResult)
	if err != nil {
		log.Println("Decode articles Error", err)
		return nil, err
	}
	return &serviceResult, nil
}

func GetTasksByArticleId(db *mongo.Database, queryBson bson.M) ([]models.MRCTask, error) {
	TaskCollection := db.Collection("MRCTask")
	var tasks []models.MRCTask
	cur, err := TaskCollection.Find(context.Background(), queryBson)
	if err != nil {
		log.Println("Find tasks Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.MRCTask{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode tasks Error", err)
			return nil, err
		}
		tasks = append(tasks, result)
	}
	return tasks, nil
}

func GetAnswers(db *mongo.Database, task models.MRCAnswer) ([]*models.MRCAnswer, error) {
	AnswerCollection := db.Collection("MRCAnswer")
	var answers []*models.MRCAnswer
	cur, err := AnswerCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.MRCAnswer{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		answers = append(answers, &result)
	}
	return answers, nil
}

func GetTaskById(db *mongo.Database, taskModel models.MRCTask) (*models.MRCTask, error) {
	TaskCollection := db.Collection("MRCTask")
	var task models.MRCTask
	result := TaskCollection.FindOne(context.Background(), taskModel.ToQueryBson())
	err := result.Decode(&task)
	if err != nil {
		log.Println("Decode task Error", err)
		return nil, err
	}
	return &task, nil
}

func SaveAnswer(db *mongo.Database, answer models.MRCAnswer) (*mongo.InsertOneResult, error) {
	AnswerCollection := db.Collection("MRCAnswer")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := AnswerCollection.InsertOne(ctx, answer)
	if err != nil {
		log.Println("Insert answers Error", err)
		return nil, err
	}
	return res, nil
}

func UpdateAnswer(db *mongo.Database, answer models.MRCValidation) (*mongo.UpdateResult, error) {
	AnswerCollection := db.Collection("MRCAnswer")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": bson.M{"$eq": answer.OriginalId}}
	update := bson.M{"$set": bson.M{"status": answer.Status}}
	res, err := AnswerCollection.UpdateOne(ctx, filter, update)	
	log.Println("res type", reflect.TypeOf(res).Kind())
	if err != nil {
		log.Println("update answer error", err)
		return nil, err
	}
	return res, nil
}

func SaveValidationStatus(db *mongo.Database, validationAnswer models.MRCValidation) (*mongo.InsertOneResult, error) {
	log.Println("validation answer save:", validationAnswer)
	ValidationCollection := db.Collection("MRCValidation")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := ValidationCollection.InsertOne(ctx, validationAnswer)
	if err != nil {
		log.Println("Insert answers Error", err)
		return nil, err
	}
	return res, nil
}

func GetRandomValidationQuestion(db *mongo.Database, question models.MRCAnswer) (*models.MRCAnswer, error) {
	AnswerCollection := db.Collection("MRCAnswer")
	var questionPair models.MRCAnswer
	res := AnswerCollection.FindOne(context.Background(), question.ToQueryBson())
	err := res.Decode(&questionPair)
	if err != nil {
		log.Println("Decode task Error", err)
		return nil, err
	}
	return &questionPair, nil
}

func FindAnswerById(db *mongo.Database, id primitive.ObjectID) (*models.MRCAnswer, error) {
	AnswerCollection := db.Collection("MRCAnswer")
	var originalAnswerInfo models.MRCAnswer
	log.Println("Find original", id)
	res := AnswerCollection.FindOne(context.Background(), bson.M{"_id": id})
	err := res.Decode(&originalAnswerInfo)
	if err != nil {
		log.Println("Decode original info error", err)
		return nil, err
	}
	return &originalAnswerInfo, nil
}

func GetRandomDecisionInfo(db *mongo.Database, userId string) (*models.MRCValidation, error) {
	ValidationCollection := db.Collection("MRCValidation")
	var decisionInfo models.MRCValidation
	id, _ := primitive.ObjectIDFromHex(userId)
	res := ValidationCollection.FindOne(context.Background(), bson.M{"status": "unverified", "validationUserId": bson.M{"$ne": id}, "labelUserId": bson.M{"$ne": id}})
	resErr := res.Decode(&decisionInfo)
	if resErr != nil {
		log.Println("Decode decisionInfo error", resErr)
		return nil, resErr
	}
	return &decisionInfo, nil
}

//================================= sentiment API =================================
func GetSentiArticles(db *mongo.Database) ([]models.Article, error) {
	collection := db.Collection("SentiArticles")
	var articles = []models.Article{}
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Find Articles Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.Article{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode Article Error", err)
			return nil, err
		}
		articles = append(articles, result)
	}
	return articles, nil
}

func GetSentiArticleByArticleId(db *mongo.Database, queryBson bson.M) (*models.Article, error) {
	ArticleCollection := db.Collection("SentiArticles")
	var serviceResult models.Article
	cur := ArticleCollection.FindOne(context.Background(), queryBson)
	err := cur.Decode(&serviceResult)
	if err != nil {
		log.Println("Decode articles Error", err)
		return nil, err
	}
	return &serviceResult, nil
}

func GetSentiTasksByArticleId(db *mongo.Database, queryBson bson.M) ([]models.SentiTask, error) {
	TaskCollection := db.Collection("SentiTask")
	var tasks []models.SentiTask
	cur, err := TaskCollection.Find(context.Background(), queryBson)
	if err != nil {
		log.Println("Find tasks Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.SentiTask{}
		err := cur.Decode(&result)
		// log.Println(result)
		if err != nil {
			log.Println("Decode tasks Error", err)
			return nil, err
		}
		tasks = append(tasks, result)

	}
	return tasks, nil
}

func GetAspectByTaskId(db *mongo.Database, task models.SentiAspect) ([]*models.SentiAspect, error) {
	AnswerCollection := db.Collection("SentiAspect")
	var aspects []*models.SentiAspect
	cur, err := AnswerCollection.Find(context.Background(), task.ToQueryBson())
	if err != nil {
		log.Println("Find answers Error", err)
		return nil, err
	}
	for cur.Next(context.Background()) {
		result := models.SentiAspect{}
		err := cur.Decode(&result)
		if err != nil {
			log.Println("Decode answer Error", err)
			return nil, err
		}
		aspects = append(aspects, &result)
	}
	return aspects, nil
}

func GetSentiTaskById(db *mongo.Database, taskModel models.SentiTask) (*models.SentiTask, error) {
	TaskCollection := db.Collection("SentiTask")
	var task models.SentiTask
	result := TaskCollection.FindOne(context.Background(), taskModel.ToQueryBson())
	err := result.Decode(&task)
	if err != nil {
		log.Println("Decode task Error", err)
		return nil, err
	}
	return &task, nil
}

func SaveSentiAnswer(db *mongo.Database, answer models.SentiAnswer) (*mongo.InsertManyResult, error) {
	AspectCollection := db.Collection("SentiAspect")
	aspectList := make([]interface{}, len(answer.Aspect))
	for i := range answer.Aspect {
		aspectList[i] = answer.Aspect[i]
	}
	SentiCollection := db.Collection("SentiSentiment")
	sentiList := make([]interface{}, len(answer.Sentiment))
	for i := range answer.Sentiment {
		sentiList[i] = answer.Sentiment[i]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := AspectCollection.InsertMany(ctx, aspectList)
	if err != nil {
		log.Println("Insert aspects Error", err)
		return nil, err
	}
	res, err = SentiCollection.InsertMany(ctx, sentiList)
	if err != nil {
		log.Println("Insert sentiments Error", err)
		return nil, err
	}
	return res, nil
}

//這邊要等到 validation 的時候才會用到
// func GetSentiAnswer(db *mongo.Database, task models.MRCAnswer) ([]*models.MRCAnswer, error) {
// 	AnswerCollection := db.Collection("MRCAnswer")
// 	var answers []*models.MRCAnswer
// 	cur, err := AnswerCollection.Find(context.Background(), task.ToQueryBson())
// 	if err != nil {
// 		log.Println("Find answers Error", err)
// 		return nil, err
// 	}
// 	for cur.Next(context.Background()) {
// 		result := models.MRCAnswer{}
// 		err := cur.Decode(&result)
// 		if err != nil {
// 			log.Println("Decode answer Error", err)
// 			return nil, err
// 		}
// 		answers = append(answers, &result)
// 	}
// 	return answers, nil
// }
