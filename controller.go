package main

import (
	"fmt"
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
const hTTPSafe = true // default false, true once http2 is implemented

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

func runHTTPSequence(runHTTPS bool) {
	fmt.Println("run model() in TLS not done...")
	//log.Fatal("write the hekkin code to model in TLS ONLY...")
	go httpsModel() // found in routeModelHTTPS.go

	// then
	http.HandleFunc(indexPath, routeIndex)
	http.HandleFunc(errorsPath, routeError)  // found in errors.go
	http.HandleFunc(staticPath, routeStatic) // found in static.go
	http.HandleFunc(viewPath, routeView)     // found in view.go
	http.HandleFunc(modelPath, routeModel)   // found in model.go
	fmt.Println("GoMvClean v42 running...")
	log.Fatal(http.ListenAndServe(httpPort, nil))

}

func main() {
	runHTTPSequence(true)
	if gDebug {
		testEverything(true) //found in test.go
	}
	return
}
