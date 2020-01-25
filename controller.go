package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	//fmt.Println(path, "len:", len(path))
	if path == "/" {
		path = staticMarkupFolder + "homePage" + staticMarkupType
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
	testModelSchema(true)                    // found in model.go
	log.Fatal(http.ListenAndServe(":8080", nil))
	return
}
