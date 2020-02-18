package main

import (
	"net/http"
	"os"
)

const robotsPath string = "/robots.txt"
const humansPath string = "/humans.txt"
const staticTxtFolder = "txt" + string(os.PathSeparator)
const staticTxtFolderPath = staticFolder + staticTxtFolder

const faviconFileName = "favicon.ico"
const faviconPath string = "/" + faviconFileName
const staticImgFolder = "img" + string(os.PathSeparator)
const staticImgFolderPath = staticFolder + staticImgFolder

const idioFAIL = "404"

// switch path solutions to idiosyncracies of the web, like robots.txt.
func routeIdiosync(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case faviconPath:
		// pull the favicon from the /images folder?
		w.Header().Set("Content-Type", "image/ico")
		loadStaticBody(staticImgFolderPath + faviconFileName)
	case robotsPath:
		// serve the robots.txt file from /static/txt/robots.txt
		renderStatic(staticTxtFolderPath+robotsPath, false, w, r)
	case humansPath:
		// serve the humans.txt file from the same static location
		renderStatic(staticTxtFolderPath+humansPath, false, w, r)
	default:
		errorShortCircuit(w, r, idioFAIL)
	}
}
