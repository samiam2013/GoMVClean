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
	fmt.Fprint(w, "Hi, welcome to view.go, here's what you sent me: \n\n")
	fmt.Fprint(w, "HTTP Request: \n\n")
	fmt.Fprint(w, headString)
	fmt.Fprint(w, "</code></body></html>")
	return

	//here's an idea, check if it's post or get
	// if post type
	//    then routePost(path)
	// else
	//    then routeGet(path)
}

//  routePost (path)
//  switch Path
// 		case "inputForm"
//			load javascript for pulling ajax result after data processed and uQuery
//			load new result into view with JS redirect for where data landed
//			return
// 		case "/model/public/"
//			load static index page for docs on the schema
//			load public schema in through get /model/public/schema/schema.json view
//			return
//		case
