package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// verbosity switch
const staticDEBUG = false

// who you gonna call?
const staticPath = "/static/"

//break out that auto-backslash because windows is broken
const staticMarkupFolder = "static" + string(os.PathSeparator)
const staticMarkupType = ".html"
const staticFAIL = "404"

// pull the staticMarkupFolder out of path and render it
func routeStatic(w http.ResponseWriter, r *http.Request) {
	if staticDEBUG {
		fmt.Println("routeStatic()...")
	}
	pageName := r.URL.Path[len(staticPath):]
	path := staticMarkupFolder + pageName + staticMarkupType
	renderStatic(path, w, r)
	return
}

// render a static path to the http.ResponseWriter
func renderStatic(path string, w http.ResponseWriter, r *http.Request) {
	body, err := loadStaticBody(path)
	if err != true {
		fmt.Fprintf(w, string(body))
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
