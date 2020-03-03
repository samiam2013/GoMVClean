package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const csrfTimeoutMinutes = 10
const csrfTokenPrefix = "csrfToken_"
const csrfTablePath = "model/public/csrf/"
const csrfTokenMarkup = ".json"

type csrfRequest struct {
	FormDestination string `json:"formDestination"`
}

type csrfStruct struct {
	Token    string `json:"token"`
	FormPath string `json:"formPath"`
	Timeout  int64  `json:"timeout"`
	IP       string `json:"ip"`
}

func issueCSRF(w http.ResponseWriter, r *http.Request) {
	// csrfToken() defined in csrf.go
	jsonData, hashString := csrfToken(w, r)

	//if it passes, make a string out of the data and print it.
	jsonString := string(jsonData)
	fmt.Fprint(w, jsonString)
	jsonFileName := csrfTokenPrefix + hashString + csrfTokenMarkup
	// run the corresponding model update query, have to confirm csrfs
	// uQuery is in modelFileQuery.go
	updateQuery(csrfTablePath+jsonFileName, jsonString, w, r)
	return
}

func csrfToken(w http.ResponseWriter, r *http.Request) ([]byte, string) {
	// make a timeout, path, and userIP for jsonMap
	var timeStamp int64 = time.Now().Unix()
	timeoutUnix := int64((csrfTimeoutMinutes * 60) + timeStamp)
	// get the form value from json body
	requestBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprint(w, "{'error[500]':'csrf json request could not be read'}")
	}

	var request csrfRequest
	err = json.Unmarshal(requestBody, &request)
	if err != nil {
		fmt.Fprint(w, "{'error[500]':'csrf json request could not be parsed'}")
	}

	// golang authors, WTF is a TYPE ASSERTION?
	formPath := request.FormDestination

	userIPPort := r.RemoteAddr
	//split the userIP into ip and port
	re := regexp.MustCompile(`(.*):\d+`)
	matches := re.FindStringSubmatch(userIPPort)
	userIP := string(matches[1])
	// make a hash string
	hashString := genRandHash(timeStamp) // found in crypto.go
	// map the json output
	jsonMap := csrfStruct{
		Token:    hashString,
		FormPath: formPath,
		Timeout:  timeoutUnix,
		IP:       userIP,
	}
	// marshall the JSON
	jsonData, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Fprint(w, "{'error':[500: 'json mapping failed in csrfToken()']}")
		return nil, string(jsonData)
	}
	return jsonData, hashString
}

func verifyCSRF(tokenSent string, w http.ResponseWriter, r *http.Request) bool {

	// map the JSON string to deconstruct it
	var JSONdataSent map[string]interface{}
	// try unmarshalling the JSON, if it fails send back false
	if err := json.Unmarshal([]byte(tokenSent), &JSONdataSent); err != nil {
		return false
	}
	tokenSentVal := fmt.Sprintf("%v", JSONdataSent["token"])
	tokenFindPath := csrfTablePath + csrfTokenPrefix + tokenSentVal + csrfTokenMarkup
	// try to get the token from the database
	byteTextFound, err := ioutil.ReadFile(tokenFindPath)
	if err != nil {
		// couldn't find or couldn't read CSRF
		return false
	}
	// make sure token found matches token sent
	if tokenSent == string(byteTextFound) {
		// make sure the requesting IP matches the token IP
		if strings.HasPrefix(r.RemoteAddr, fmt.Sprintf("%v", JSONdataSent["ip"])) {
			// make sure the timeout hasn't passed.
			currentTS := time.Now().Unix()
			timeoutString := fmt.Sprintf("%v", JSONdataSent["timeout"])
			toString, err := strconv.ParseInt(timeoutString, 10, 64)
			// delete the file after checks are done.
			defer os.Remove(tokenFindPath)
			if err != nil {
				fmt.Println("could not parse timeout string! csrf.go csrfVerify()")
				return false
			}
			if currentTS < toString {
				if globalDebug {
					fmt.Println("csrf token validated. ")
				}
				return true
			}
			fmt.Println("timeout past due on csrf token")
			return false
		}
		fmt.Println("user ip doesn't match token ip")
		return false
	}
	return false
}
