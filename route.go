package main

import (
	"net/http"

	"github.com/Lynx/respond"
)

type RouteMux struct {
}

func (p *RouteMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	if r.URL.Path == "/labels" {
		respond.GetLabelsByTaskId(collection, w, r)
		return
	}
	http.NotFound(w, r)
	return
}
