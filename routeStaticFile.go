package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const cssFolderName = "css"
const cssFolder = cssFolderName + string(os.PathSeparator)
const cssPath = "/" + cssFolderName + "/"
const staticCSSFolder = staticFolder + cssFolder

const jsFolderName = "js"
const jsFolder = jsFolderName + string(os.PathSeparator)
const jsPath = "/" + jsFolderName + "/"

const robotsPath = "/robots.txt"
const humansPath = "/humans.txt"
const staticTxtFolder = "txt" + string(os.PathSeparator)
const staticTxtFolderPath = staticFolder + staticTxtFolder

const faviconFileName = "favicon.ico"
const faviconPath = "/" + faviconFileName
const staticImgFolder = "img" + string(os.PathSeparator)
const staticImgFolderPath = staticFolder + staticImgFolder

const idioFAIL = "404"

// switch path solutions to idiosyncracies of the web, like robots.txt.
func routeStaticFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if globalDebug {
		fmt.Println("routeIdioSync(", path, ")")
	}
	if strings.HasPrefix(path, jsPath) {
		// render js as text
		setCache(7, w)
		w.Header().Set("Content-Type", "text/javascript")
		renderStatic(staticFolder+jsFolder+path[len(jsPath):], false, w, r)
		return
	} else if strings.HasPrefix(path, cssPath) {
		// render css as text
		setCache(7, w)
		w.Header().Set("Content-Type", "text/css")
		renderStatic(staticCSSFolder+path[len(cssPath):], false, w, r)
		return
	} else {
		switch path {
		case faviconPath:
			// pull the favicon from the /images folder?
			//w.Header().Set("Content-Type", "image/ico")
			setCache(30, w)
			http.ServeFile(w, r, staticImgFolderPath+faviconFileName)
			return
		case robotsPath:
			// serve the robots.txt file from /static/txt/robots.txt
			w.Header().Set("Content-Type", "text/plain")
			renderStatic(staticTxtFolderPath+robotsPath, false, w, r)
			return
		case humansPath:
			// serve the humans.txt file from the same static location
			w.Header().Set("Content-Type", "text/plain")
			renderStatic(staticTxtFolderPath+humansPath, false, w, r)
			return
		default:
			errorShortCircuit(w, r, idioFAIL)
		}
	}
}

func setCache(days int64, w http.ResponseWriter) {
	intSeconds := days * 24 * 60 * 60
	cacheAge := "max-age=" + fmt.Sprintf("%v", intSeconds)
	fmt.Println(cacheAge)
	w.Header().Set("Cache-Control", cacheAge)
}
