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
const hTTPSafe = false // default false, true once http2 is implemented

// if path ("/"|"/indexName") serve indexName+staticMarkupType otherwise 404
func routeIndex(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == indexPath || path == indexPathName {
		path = staticMarkupFolder + indexName + staticMarkupType
		renderStatic(path, w, r)
		return
	}
	errorShortCircuit(w, r, indexFAIL)
	return
}

func runHTTP() {
	http.HandleFunc(indexPath, routeIndex)
	http.HandleFunc(errorsPath, routeError)  // found in errors.go
	http.HandleFunc(staticPath, routeStatic) // found in static.go
	http.HandleFunc(viewPath, routeView)     // found in view.go
	http.HandleFunc(modelPath, routeModel)   // found in model.go
	log.Fatal(http.ListenAndServe(httpPort, nil))
}

func main() {
	if hTTPSafe {
		runHTTP()
	} else if hTTPSafe {
		//configure the https through the http server and run it here
	}
	if gDebug {
		testEverything(true) //found in test.go
	}
	return
}
