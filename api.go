package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"
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
const apiCopies int64 = 0

const csrfPath = apiPath + "csrf"
const csrfTimeoutMinutes = 10

func routeAPI(w http.ResponseWriter, r *http.Request) {
	// and at the end of the function call it needs to be albe to call
	// each and every other single running instance of this server that's
	// running on another machine,

	path := r.URL.Path
	switch path {
	case csrfPath:
		// !!!!!!!!!!!
		//	THIS NEEDS RATE-LIMITING AGAINST IP, SESSION COOKIE!
		//
		// 	THIS NEEDS TO REQUIRE A DECLARATION OF NEW SESSION OR
		//  THE LAST CSRF TOKEN THE USER WAS ISSUED.
		//
		// seed the random number generator with the time.
		rand.Seed(time.Now().UnixNano())
		// get the current time as int64 unix time in seconds.
		timeStamp := time.Now().Unix()
		// make a random 64 bit number
		randN := rand.Int63n(timeStamp)
		// make an 8*8 bit byte-slice
		randB := make([]byte, 8)
		// force that int64 into that 8 byte slice
		binary.LittleEndian.PutUint64(randB, uint64(randN))
		// generate the byte form random hash
		hashN := sha256.Sum256(randB)
		// generate the string from the byte form
		hashString := fmt.Sprintf("%x", hashN)

		// make an expiration for the token in unix seconds
		csrfTimeoutUnix := int64((csrfTimeoutMinutes * 60) + timeStamp)

		// format the int64 timeout for Go, it's a picky language
		timeoutString := fmt.Sprintf("%d", csrfTimeoutUnix)

		// figure out which form called us
		formPath := r.URL.Path

		// map the json output
		jsonMap := map[string]string{
			"token":    hashString,
			"formPath": formPath,
			"timeout":  timeoutString,
		}

		// marshall the JSON
		jsonData, err := json.Marshal(jsonMap)
		if err != nil {
			errorShortCircuit(w, r, "500")
		}

		// make a string from the json
		jsonString := string(jsonData)

		// try to unmarshal it, you could print dat
		var dat map[string]interface{}
		if err := json.Unmarshal(jsonData, &dat); err != nil {
			errorShortCircuit(w, r, "500")
			log.Fatal(err)
		} else {
			// but print this weird stuff instead, see what it means.
			fmt.Fprint(w, jsonString)
			// run the corresponding model query, have to confirm csrfs
			uQuery(true, "./public/updateStream/blobs/"+hashString+".json", jsonString, "schema", w, r)
		}

	// this happens if the api doesn't know what to do
	// default case: echo the input.
	default:
		header, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		headString := string(header)
		fmt.Fprint(w, headString)
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

	// this way the store is consistent and scalable as long as there
	// are no hash collisions amounst separate models (even md5 has no collision)
	// out to 512 bytes, it's been rainbow-tabled that far.
	return
}
