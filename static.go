package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const staticPath = "/static/"
const staticMarkupFolder = "static" + string(os.PathSeparator)
const staticMarkupType = ".html"

func routeStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("routeStatic()...")
	pageName := r.URL.Path[len(staticMarkupFolder):]
	path := staticMarkupFolder + pageName + staticMarkupType
	renderStatic(path, w, r)
	return
}

func renderStatic(path string, w http.ResponseWriter, r *http.Request) {
	fmt.Println("renderStatic(", path, ")")
	body, err := loadStaticBody(path)
	if err != true {
		fmt.Fprintf(w, string(body))
	} else {
		errorShortCircuit(w, r, "404")
	}
	return
}

func loadStaticBody(path string) ([]byte, bool) {
	body, err := ioutil.ReadFile(path)
	if err == nil {
		return body, false
	}
	return nil, true
}
