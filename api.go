package main

import (
	"encoding/json"
	"fmt"
	"log"
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

const csrfPath = apiPath + "csrf"
const csrfTimeoutMinutes = 10

// RouteAPI uses the url path to switch to the most efficient
//  and functionally decomposed JSON input/output scheme I can muster.
func routeAPI(w http.ResponseWriter, r *http.Request) {
	// functionality for api/csrf/server/0..n ?

	path := r.URL.Path
	switch path {
	case csrfPath:
		// csrfToken() defined in csrf.go
		jsonData, hashString := csrfToken(w, r)
		// testJSON makes sure responses will parse as json
		if testJSON("stringValuePairs", jsonData, w, r) {
			//if it passes, make a string out of the data and print it.
			jsonString := string(jsonData)
			fmt.Fprint(w, jsonString)
			jsonFileName := "csrfToken_" + hashString + ".json"
			// run the corresponding model update query, have to confirm csrfs
			// uQuery is in modelFileQuery.go
			uQuery(true, "public/updateStream/blobs/"+jsonFileName,
				jsonString, "schema", w, r)
		}
		return
	// default case: Internal error 500
	default:
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprint(w, "{'error [501]':'", path, " not implemented.'}")
		return
	}
}

func testJSON(mapSelect string, jsonData []byte,
	w http.ResponseWriter, r *http.Request) bool {
	// verifies that json is good and returns true/false
	//	allows you to automatically test json against a set of map interfaces
	switch mapSelect {
	case "stringValuePairs":
		// this is the map of the interface go will use to un-marshall
		var dat map[string]interface{}
		// try unmarshalling the JSON, if it fails send back JSON error string
		if err := json.Unmarshal(jsonData, &dat); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{'error [500]':'malformed json stopped by API'}")
			log.Fatal(err)
			return false
		}
		return true
	default:
		fmt.Fprint(w,
			"{'error [501]':'JSON map case ", mapSelect, " not implemented'}",
			mapSelect, "* unknown!']}")
		return false
	}
}
