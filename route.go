package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Lynx/respond"
)

type RouteMux struct {
}

// example
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析參數，預設是不會解析的
	fmt.Println(r.Form) //這些資訊是輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello!") //這個寫入到 w 的是輸出到客戶端的
}

func (p *RouteMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		sayhelloName(w, r)
		return
	case "/getTask":
		respond.GetTaskByTaskId(collection, w, r)
		return
	case "/articles":
		respond.GetArticles(database, w, r)
		return
	case "/saveArticles":
		respond.SaveArticles(database, w, r)
		return
	case "/testArticles":
		respond.TestArticles(database, w, r)
		return
	default:
		break
	}
	http.NotFound(w, r)
	return
}
