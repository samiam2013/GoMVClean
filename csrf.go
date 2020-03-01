package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const csrfPath = apiPath + "csrf"
const csrfTimeoutMinutes = 10

func issueCSRF(w http.ResponseWriter, r *http.Request) {
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
		updateQuery("model/public/csrf/"+jsonFileName, jsonString, w, r)
	}
	return
}

func csrfToken(w http.ResponseWriter, r *http.Request) ([]byte, string) {
	// make a timeout, path, and userIP for jsonMap
	var timeStamp int64 = time.Now().Unix()
	csrfTimeoutUnix := int64((csrfTimeoutMinutes * 60) + timeStamp)
	timeoutString := fmt.Sprintf("%d", csrfTimeoutUnix)
	formPath := r.FormValue("form")
	userIP := r.RemoteAddr
	// make a hash string
	hashString := genRandHash(timeStamp) // found in crypto.go
	// map the json output
	jsonMap := map[string]string{
		"token":    hashString,
		"formPath": formPath,
		"timeout":  timeoutString,
		"ipPort":   userIP,
	}
	// marshall the JSON
	jsonData, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Fprint(w, "{'error':[500: 'json mapping failed in csrfToken()']}")
		return nil, string(jsonData)
	}
	return jsonData, hashString
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
