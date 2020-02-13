package main

import "net/http"

const indexPath string = "/"
const indexName string = "index"
const indexPathName string = indexPath + indexName
const indexFAIL = "404"

// if path ("/"|"/indexName") serve indexName+staticMarkupType otherwise 404
func routeIndex(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == indexPath || path == indexPathName {
		path = staticFolder + indexName + staticMarkupType
		renderStatic(path, true, w, r)
		return
	}
	errorShortCircuit(w, r, indexFAIL)
	return
}
