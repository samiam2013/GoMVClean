package main

import "fmt"

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

const debugHTTPS = true

func httpsModelAPI() {
	//this runs in a go routine from controller.go
	if debugHTTPS {
		fmt.Println("method httpsModelAPI needs implementation.")
	}
	// and at the end of the function call it needs to be albe to call
	// each and every other single running instance of this server that's
	// running on another machine,

	// this way the store is consistent and scalable as long as there
	// are no hash collisions amounst separate models (even md5 has no collision)
	// out to 512 bytes, it's been rainbow-tabled that far.
	return
}
