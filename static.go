package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// verbosity switch
const staticDEBUG = true

// who you gonna call?
const staticPath = "/static/"

//break out that auto-backslash because windows is broken
const staticMarkupFolder = "static" + string(os.PathSeparator)
const staticMarkupType = ".html"
const staticFAIL = "404"

//make sure the software knows where the html header/footer are located
const headerName = "header"
const headerPath = staticMarkupFolder + headerName + staticMarkupType
const footerName = "footer"
const footerPath = staticMarkupFolder + footerName + staticMarkupType

//jquery minified file path
const jsFolderPath = "js" + string(os.PathSeparator)
const jQueryRelativePath = staticMarkupFolder + jsFolderPath + "jquery.js"
const jQueryPath = "/static/js/jquery.js"

// pull the staticMarkupFolder out of path and render it
func routeStatic(w http.ResponseWriter, r *http.Request) {
	pageName := r.URL.Path[len(staticPath):]
	path := staticMarkupFolder + pageName
	if filepath.Ext(pageName) != ".js" {
		path = path + staticMarkupType
	}
	renderStatic(path, true, w, r)
	if staticDEBUG {
		fmt.Println("routeStatic()...path:", path, ", pageName", pageName)
	}
	return
}

// render a static path to the http.ResponseWriter
func renderStatic(path string, isHTML bool,
	w http.ResponseWriter, r *http.Request) {
	body, err := loadStaticBody(path)
	if err != true {
		if isHTML {
			renderStatic(headerPath, false, w, r)
		}
		fmt.Fprintf(w, string(body))
		if isHTML {
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
	body, err := ioutil.ReadFile(path)
	if err == nil {
		return body, false
	}
	return nil, true
}

//manually load jquery if it's called.
func routeJQuery(w http.ResponseWriter, r *http.Request) {
	renderStatic(jQueryRelativePath, false, w, r)
}
