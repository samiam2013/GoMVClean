package main

import (
	"fmt"
	"net/http"
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
// if !stringContains(r.URL.Path,
//			"/model/public/*tablePath/*rowPath/*columnPath/*hash(csrf_token)")
//    			updateStream(w, func(r))

const apiPath = "/api/" // for controller.go
const apiCopies int64 = 0

// RouteAPI uses the url path to switch to the most efficient
//  and functionally decomposed JSON input/output scheme I can muster.
func routeAPI(w http.ResponseWriter, r *http.Request) {
	// functionality for api/csrf/server/0..n ?
	path := r.URL.Path
	switch path {
	// found in csrf.go
	case csrfPath:
		issueCSRF(w, r)
	// default case: Internal error 501
	default:
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, "{'error [501]':'", path, " not implemented. ",
			"implement this in api.go by adding to the path switch ",
			"in routeAPI.go'}")
		return
	}

}
