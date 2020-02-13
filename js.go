package main

import (
	"fmt"
	"net/http"
	"os"
)

const jsDEBUG = false

//jquery minified file path
const javascriptPath = "/js/"
const jsFolder = "js" + string(os.PathSeparator)
const jQueryFileName = "jquery.js"
const jQueryFilePath = staticFolder + jsFolder + jQueryFileName

//manually load jquery if it's called.
func routeJS(w http.ResponseWriter, r *http.Request) {
	if jsDEBUG {
		fmt.Println("routing js path...")
	}
	renderStatic(jQueryFilePath, false, w, r)
}
