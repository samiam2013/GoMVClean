package main

import (
	"log"
	"net/http"
)

// if path == / , load homePage.html. otherwise 404
func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" || path == "/index" {
		path = staticMarkupFolder + "index" + staticMarkupType
		renderStatic(path, w, r)
		return
	}
	errorShortCircuit(w, r, "404")
	return
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc(errorsPath, routeError)  // found in errors.go
	http.HandleFunc(staticPath, routeStatic) // found in static.go
	http.HandleFunc(viewPath, routeView)     // found in view.go
	http.HandleFunc(modelPath, routeModel)   // found in model.go
	log.Fatal(http.ListenAndServe(":8080", nil))
	return
}
