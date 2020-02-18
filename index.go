package main

import "net/http"

// IndexPath is exported to the controller for routing
const IndexPath string = "/"
const indexName string = "index"
const indexPathName string = IndexPath + indexName
const indexFAIL = "404"

// if path ("/"|"/indexName") serve indexName+staticMarkupType otherwise 404
func routeIndex(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == IndexPath || path == indexPathName {
		path = staticFolder + indexName + staticMarkupType
		renderStatic(path, true, w, r)
		return
	}
	errorShortCircuit(w, r, indexFAIL)
	return
}
