package main

import (
	"net/http"
	"os"
)

const ErrorsPath = "/errors/"
const errorsMarkupFolder = "errors" + string(os.PathSeparator)
const errorsMarkupType = ".html"

func RouteError(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(ErrorsPath):]
	path := ErrorsPath + title + errorsMarkupType
	renderStatic(path, true, w, r)
	return
}

func errorShortCircuit(w http.ResponseWriter,
	r *http.Request, errNumString string) {
	path := errorsMarkupFolder + errNumString + errorsMarkupType
	//fmt.Println("errors.go:errorShortCircuit(,", path, ",)")
	switch errNumString {
	case "404":
		w.WriteHeader(http.StatusNotFound)
		renderStatic(path, true, w, r)
		return
	case "403":
		w.WriteHeader(http.StatusForbidden)
		renderStatic(path, true, w, r)
		return
	default:
		path = errorsMarkupFolder + "500" + errorsMarkupType
		w.WriteHeader(http.StatusInternalServerError)
		renderStatic(path, true, w, r)
		return
	}
}