package main

import "net/http"

// IndexPath is exported to the controller for routing
const indexPath string = "/"
const indexName string = "index"
const indexPathName string = indexPath + indexName
const indexFAIL = "404"

// if path ("/"|"/indexName") serve indexName+staticMarkupType otherwise 404
func routeIndex(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == indexPath || path == indexPathName {
		path = viewFolder + indexName + viewMarkupType
		renderStatic(path, true, w, r)
		return
	}
	errorShortCircuit(w, r, indexFAIL)
	return
}
