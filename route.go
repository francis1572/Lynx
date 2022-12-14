package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	respond "Lynx/controller"
)

type RouteMux struct {
}

// example
func sayhelloName(w http.ResponseWriter, r *http.Request) error {
	var queryInfo map[string]string
	err := json.NewDecoder(r.Body).Decode(&queryInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	fmt.Println("query info:", queryInfo)
	fmt.Fprintf(w, queryInfo["data"])
	return nil
}

func (p *RouteMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "application/json")
	switch path {
	case "/":
		sayhelloName(w, r)
		return
	case "/login":
		respond.Login(Database, w, r)
		return
	case "/users":
		respond.GetUsers(Database, w, r)
		return
	case "/projectUsers":
		respond.GetUsersByProject(Database, w, r)
		return
	case "/saveAuth":
		respond.SaveAuth(Database, w, r)
		return
	// case "/saveProject":
	// 	respond.SaveProject(Database, w, r)
	// 	return
	// case "/projects":
	// 	respond.GetProjects(Database, w, r)
	// 	return
	// case "/articles":
	// 	respond.GetArticles(Database, w, r)
	// 	return
	case "/sentiArticles":
		respond.GetSentiArticles(Database, w, r)
		return
	// case "/saveArticles":
	// 	respond.SaveArticles(Database, w, r)
	// 	return
	// case "/tasks":
	// 	respond.GetTasksByArticleId(Database, w, r)
	// 	return
	case "/sentiTasks":
		respond.GetSentiTasksByArticleId(Database, w, r)
		return
	case "/getTask":
		log.Println("POST /getTask")
		respond.GetTaskById(Database, w, r)
		return
	case "/getSentiTask":
		log.Println("POST /getSentiTask")
		respond.GetSentiTaskById(Database, w, r)
		return
	case "/getSentiValidation":
		log.Println("POST /getSentiValidation")
		respond.GetSentiValidation(Database, w, r)
		return
	case "/saveAnswer":
		log.Println("POST /SaveAnswer")
		respond.SaveAnswer(Database, w, r)
		return
	case "/getValidation":
		log.Println("POST /GetValidation")
		respond.GetValidation(Database, w, r)
		return
	case "/getSentiAspects":
		log.Println("POST /getSentiAspects")
		respond.GetSentiAspects(Database, w, r)
		return
	case "/postSentiValidation":
		log.Println("POST /PostSentiValidation")
		respond.PostSentiValidation(Database, w, r)
		return
	case "/saveValidation":
		log.Println("POST /SaveValidation")
		respond.SaveValidation(Database, w, r)
		return
	case "/saveSentiAnswer":
		log.Println("POST /SaveSentiAnswer")
		respond.SaveSentiAnswer(Database, w, r)
		return
	case "/checkIsAnswered":
		log.Println("POST /CheckIsAnswered")
		respond.CheckIsAnswered(Database, w, r)
		return
	case "/checkIsValidated":
		log.Println("POST /CheckIsValidated")
		respond.CheckIsValidated(Database, w, r)
		return
	case "/discardSentiAnswer":
		log.Println("POST /discardSentiAnswer")
		respond.DiscardSentiAnswer(Database, w, r)
		return
	case "/getDecision":
		log.Println("POST /getRandomDecision")
		respond.GetDecision(Database, w, r)
		return
	case "/saveDecision":
		log.Println("POST /SaveDecision")
		respond.SaveDecision(Database, w, r)
		return
	case "/test":
		respond.Test(w, r)
		return
	default:
		break
	}
	http.NotFound(w, r)
	return
}
