package main

import (
	"fmt"
	"net/http"
	"os"
)

const jsDEBUG = false

//jquery minified file path
const jsPath = "/js/"
const jsFolder = "js" + string(os.PathSeparator)
const jQueryFileName = "jquery.js"
const jQueryFilePath = staticFolder + jsFolder + jQueryFileName

//manually load jquery if it's called.
func routeJS(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == jQueryFilePath {
		renderStatic(jQueryFilePath, false, w, r)
	} else {
		renderStatic(staticFolder+jsFolder+path[len(jsPath):], false, w, r)
	}
	if jsDEBUG {
		fmt.Println("routing js path...")
	}
}
