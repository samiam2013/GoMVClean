package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

const viewPath = "/view/"

func routeView(w http.ResponseWriter, r *http.Request) {
	//this code, thanks to *someone* on stackoverflow
	// Save a copy of this request for debugging.
	header, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	headString := string(header)

	fmt.Fprint(w, "<html><body><code style=\"white-space: pre-wrap;\">")
	fmt.Fprint(w, "Hi, welcome to updateStream.go, here's what you sent me: \n\n")
	fmt.Fprint(w, "HTTP Request: \n\n")
	fmt.Fprint(w, headString)
	fmt.Fprint(w, "</code></body></html>")
	return
}
