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
	case "/articles":
		respond.GetArticles(Database, w, r)
		return
	case "/sentiArticles":
		respond.GetSentiArticles(Database, w, r)
		return
	case "/saveArticles":
		respond.SaveArticles(Database, w, r)
		return
	case "/tasks":
		respond.GetTasksByArticleId(Database, w, r)
		return
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
	case "/saveAnswer":
		log.Println("POST /SaveAnswer")
		respond.SaveAnswer(Database, w, r)
		return
	case "/saveSentiAnswer":
		log.Println("POST /SaveSentiAnswer")
		respond.SaveSentiAnswer(Database, w, r)
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
