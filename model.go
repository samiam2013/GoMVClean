package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//model verbosity switch
const modelDEBUG = true

//where do I sit on the website tree?
const modelPath string = "/model/"

//where do I get my data from?
const publicDBPath string = "public/"
const privateDBPath string = "private/"
const modelMarkup string = ".json"
const modelPubPath string = modelPath + publicDBPath
const modelPrivPath string = modelPath + privateDBPath

func updateQuery(path, data string, w http.ResponseWriter, r *http.Request) bool {
	dataByte := []byte(data)
	err := ioutil.WriteFile(path, dataByte, 0642)
	if err != nil {
		fmt.Println(path, " couldn't be written to")
		return false
	}
	w.WriteHeader(http.StatusInternalServerError)
	return true
}
