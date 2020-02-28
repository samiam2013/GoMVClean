package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// verbosity switch
const staticDEBUG = globalDebug && false

// StaticPath is exported to controller.go for routing
const staticPath = "/static/"

//break out that auto-backslash because windows is broken
const staticFolder = "static" + string(os.PathSeparator)
const staticMarkupType = ".html"
const staticFAIL = "404"

//make sure the software knows where the html header/footer are located
const headerName = "header"
const headerPath = staticFolder + headerName + staticMarkupType
const footerName = "footer"
const footerPath = staticFolder + footerName + staticMarkupType

// pull the staticMarkupFolder out of path and render it
func routeStatic(w http.ResponseWriter, r *http.Request) {
	pageName := r.URL.Path[len(staticPath):]
	path := staticFolder + pageName
	if filepath.Ext(pageName) == staticMarkupType {
		fmt.Println("special case: routstatic Markup type(", staticMarkupType, ")..", path)
		renderStatic(path, true, w, r)
	} else {
		renderStatic(path, false, w, r)
	}
	if staticDEBUG {
		fmt.Println("routeStatic()...path:", path, ", pageName", pageName)
	}
	return
}

// render a static path to the http.ResponseWriter
func renderStatic(path string, isMARKUP bool,
	w http.ResponseWriter, r *http.Request) {
	if staticDEBUG {
		fmt.Println("renderStatic(", path, ")")
	}
	body, err := loadStaticBody(path)
	if err != true {
		if isMARKUP {
			renderStatic(headerPath, false, w, r)
		}
		fmt.Fprintf(w, string(body))
		if isMARKUP {
			renderStatic(footerPath, false, w, r)
		}
	} else {
		errorShortCircuit(w, r, staticFAIL)
	}
	return
}

// this gets re-used in modelQuery.go query()
// and again in modelBreakStuff.go through query()
// DANGEROUS STUFF HERE
// load a static body given the relative path
func loadStaticBody(path string) ([]byte, bool) {
	if staticDEBUG {
		fmt.Println("loadStaticBody(", path, ")")
	}
	body, err := ioutil.ReadFile(path)
	if err == nil {
		return body, false
	}
	return nil, true
}
