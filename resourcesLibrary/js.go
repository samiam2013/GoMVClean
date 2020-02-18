package resourceLibrary

import (
	"fmt"
	"net/http"
	"os"
)

const jsDEBUG = false

// JSPath is Exported to controller.go for routing
const JSPath = "/js/"
const jsFolder = "js" + string(os.PathSeparator)
const jQueryFileName = "jquery.js"
const jQueryFilePath = staticFolder + jsFolder + jQueryFileName

//manually load jquery if it's called.
func routeJS(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == jQueryFilePath {
		renderStatic(jQueryFilePath, false, w, r)
	} else {
		renderStatic(staticFolder+jsFolder+path[len(JSPath):], false, w, r)
	}
	if jsDEBUG {
		fmt.Println("routing js path...")
	}
}
