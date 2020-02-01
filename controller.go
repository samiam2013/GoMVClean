package main

import (
	"log"
	"net/http"
)

//global debugger constant
const gDebug bool = true

const httpPort string = ":8080"
const indexPath string = "/"
const indexName string = "index"
const indexPathName string = indexPath + indexName
const indexFAIL = "404"

// if path == / , load homePage.html. otherwise 404
func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == indexPath || path == indexPathName {
		path = staticMarkupFolder + indexName + staticMarkupType
		renderStatic(path, w, r)
		return
	}
	errorShortCircuit(w, r, indexFAIL)
	return
}

func main() {
	http.HandleFunc(indexPath, homeHandler)
	http.HandleFunc(errorsPath, routeError)  // found in errors.go
	http.HandleFunc(staticPath, routeStatic) // found in static.go
	http.HandleFunc(viewPath, routeView)     // found in view.go
	http.HandleFunc(modelPath, routeModel)   // found in model.go
	if gDebug {
		testEverything(true) //found in test.go
	}
	log.Fatal(http.ListenAndServe(httpPort, nil))
	return
}
