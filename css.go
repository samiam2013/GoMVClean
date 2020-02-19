package main

import (
	"net/http"
	"os"
)

const cssFolderName = "css"
const cssFolder = cssFolderName + string(os.PathSeparator)
const cssPath = "/" + cssFolderName + "/"
const staticCSSFolder = staticFolder + cssFolder

// RouteCSS files and sets the text/css content-type
func routeCSS(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "text/css")
	renderStatic(staticCSSFolder+path[len(cssPath):], false, w, r)
}
