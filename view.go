package main

import (
	"fmt"
	"net/http"
	"os"
)

const viewName = "view"
const viewPath = "/" + viewName + "/"
const viewFolder = viewName + string(os.PathSeparator)
const viewType = ".html"

func routeView(w http.ResponseWriter, r *http.Request) {
	fmt.Println("routeView()...")
	pageName := r.URL.Path[len(viewFolder):]
	path := viewFolder + pageName + viewType
	dynamicView(path, w, r)
	return
}

func renderView(path string, body []byte,
	w http.ResponseWriter, r *http.Request) string {
	//so that we can check the errCode
	pageName := r.URL.Path[len(viewPath):]
	//fmt.Println("routeTemplate() pageName:", pageName)
	switch string(pageName) {
	case "testPost":
		return renderTestPost(body, w, r) // found in testPost.go
	default:
		fmt.Println("error: ", pageName)
		return "500"
	}
}

func dynamicView(path string, w http.ResponseWriter, r *http.Request) {
	//fmt.Println("dynamicView(", path, ")...")
	body, err := loadStaticBody(path)
	if err {
		//run template $body
		statusCode := renderView(path, body, w, r)
		if statusCode != "200" {
			errorShortCircuit(w, r, statusCode)
		}
	} else {
		errorShortCircuit(w, r, "404")
	}
	return
}
