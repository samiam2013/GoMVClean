package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// I have to figure out how to run https queries against the modeling
// and run them concurrently with the implemented model system
// this involves writing a regex so that model.go's creation ->
// 			/public/*table/*column/*row/schema
// becomes /public/*table/*column/*row/schema.json for the row schemaFile
// but is only the request path from the user, in reality.
// public/schema has links to the tables.
// public/table/schema has links to the rows
// public/table/row/schema has links to the data columns by hashed id

//functionally equivalent to
// if !stringContains(r.URL.Path,"/model/public/*tablePath/*rowPath/*columnPath/*hash(csrf_token)")
//    updateStream(w, func(r))

const apiPath = "/api/"
const apiCopies int32 = 0

func routeAPI(w http.ResponseWriter, r *http.Request) {
	// and at the end of the function call it needs to be albe to call
	// each and every other single running instance of this server that's
	// running on another machine,
	fmt.Fprint(w, "hello I am the api\n\n")
	header, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	headString := string(header)
	fmt.Fprint(w, headString)

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

	// this way the store is consistent and scalable as long as there
	// are no hash collisions amounst separate models (even md5 has no collision)
	// out to 512 bytes, it's been rainbow-tabled that far.
	return
}
