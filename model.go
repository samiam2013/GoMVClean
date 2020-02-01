package main

import (
	"fmt"
	"net/http"
)

//where do I sit on the website tree?
const modelPath string = "/model/"

//where do I get my data from?
const publicDBPath string = "public/"
const privateDBPath string = "private/"

//what's the folder for a schema called?
const modelSchemaFolder string = "schema"

//what's the file for a folder schema called?
const schemasFilesName string = "schema"
const modelMarkup string = ".json"

//const modelTestPath string = modelPath + publicDBPath + modelTestFolder
const modelNoGo string = privateDBPath + modelSchemaFolder
const modelPubPath string = modelPath + publicDBPath
const modelPrivPath string = modelPath + privateDBPath
const modelPubQueryUpdatePath string = modelPubPath + modelSchemaFolder
const modelSchemaRead string = publicDBPath + modelSchemaFolder

//can the public see the private things? it's a boolean value!
const pubAccessPrivDB bool = false //false by default

//smallQL language commands
const sQLUpdate string = "update"
const sQLRead string = "read"

//model verbosity switch
const modelDEBUG = false

//model failure Error Constant
const modelFAIL = "403"

func modelPrint(prStr string) {
	if modelDEBUG {
		fmt.Println(prStr)
	}
}

//get the path, send it along to render
func routeModel(w http.ResponseWriter, r *http.Request) {
	modelPrint("routing model!...")
	path := r.URL.Path[len(modelPath):]
	renderModel(path, w, r)
	return
}

//see if the path asked for is available to be tried with private/pub permission
// then if tryPath(), then tryPathRunFunc()
func renderModel(path string, w http.ResponseWriter, r *http.Request) {
	modelPrint("routeModel( " + path + " )")
	wholePath := r.URL.Path
	if pubAccessPrivDB {
		if tryPath(path, privateDBPath, wholePath) {
			modelPrint("tryPath(" + path + ") returned early, privateQuery() trigd.")
			tryPathRunFunc(w, r, path, privateDBPath, privateQuery)
			return
		}
	} else if tryPath(path, publicDBPath, wholePath) {
		tryPathRunFunc(w, r, path, publicDBPath, publicQuery)
		return
	}
	//tell them explicitly that's forbidden
	errorShortCircuit(w, r, modelFAIL)
	return
}
