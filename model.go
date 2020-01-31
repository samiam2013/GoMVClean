package main

import (
	"fmt"
	"net/http"
	"strconv"
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
			modelPrint("tryPath(" + path + ") returned false")
			tryPathRunFunc(w, r, path, privateDBPath, privateQuery)
			return
		}
	} else if tryPath(path, publicDBPath, wholePath) {
		tryPathRunFunc(w, r, path, publicDBPath, publicQuery)
		return
	}
	//tell them explicitly that's forbidden
	errorShortCircuit(w, r, "403")
	return
}

//see if this is modelPath/tableName/tableNameN
func tryPath(outerPath, innerPath, wholePath string) bool {
	modelPrint("tryPath( " + outerPath + ", " + innerPath + ", " + wholePath + " )")
	queryInnerPathStr := outerPath[len(innerPath):]
	queryPathLen := len(queryInnerPathStr)
	iPathTrailLen := len(outerPath[len(innerPath):])
	pathEqTrailLen := (iPathTrailLen == queryPathLen)
	return pathEqTrailLen
}

// try to run the privateQuery or publicQuery
func tryPathRunFunc(w http.ResponseWriter, r *http.Request,
	outerPath, innerPath string,
	runFunction func(string, string, http.ResponseWriter, *http.Request) bool) bool {
	modelPrint("tryPathRunFunc( " + outerPath + ", " + innerPath + ", func() )")
	if len(outerPath) > len(innerPath) {
		runFunction(innerPath, outerPath, w, r)
		return true
	}
	return false
}

// run tableQuery query against privateDBPath
func privateQuery(path, wholePath string, w http.ResponseWriter,
	r *http.Request) bool {
	modelPrint("privateQuery( " + path + " )")
	fmt.Fprintf(w, "{{Private : {Database: '%s', table: '%s'}, queried},\n", path, wholePath)
	rVal := tableQuery(path, wholePath, w, r)
	defer fmt.Fprintf(w, " returned: \"%v\"}", rVal)
	return rVal
}

// run tableQuery against publicDBPath
func publicQuery(path, wholePath string, w http.ResponseWriter,
	r *http.Request) bool {
	modelPrint("publicQuery( " + path + " )")
	//fmt.Fprintf(w, "{Database: '%s', table: '%s'}\n", path, wholePath)
	return tableQuery(path, wholePath, w, r)
}

// run rowQuery against DBPath
func tableQuery(path, wholePath string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("tableQuery( " + wholePath + " )")
	smallQL := sQLRead
	switch wholePath {
	case modelPubPath:
		//if head says POST, go to "update mode"
		contLen := r.Header.Get("Content-Length")
		contLenStr, _ := strconv.Atoi(contLen)
		//fmt.Println("post: ",contLen)
		if contLenStr == 0 {
			smallQL = sQLRead
		} else {
			smallQL = sQLUpdate
		}
		testModelSchema(wholePath, smallQL, w, r)
		return true
	case modelPrivPath:
		return testModelSchema(wholePath, smallQL, w, r)
	case modelSchemaRead:
		//modelPrint("public/schema tableQuery( " + wholePath + " )...")
		return testModelSchema(wholePath, smallQL, w, r)
	default:
		modelPrint("error: path not in tableQuery( " + wholePath + " )")
		errorShortCircuit(w, r, "403")
	}
	return false
}

//test the schema to make sure you're still safe
func testModelSchema(path, smallQL string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("testModelSchema( " + path + ", " + smallQL + " )")
	switch path {
	case modelNoGo:
		return rowQuery(modelNoGo, smallQL, w, r)
	case modelSchemaRead:
		return rowQuery(modelSchemaRead, smallQL, w, r)
	default:
		return false
	}
}

func rowQuery(path, smallQL string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("rowQuery( " + path + " )")
	switch path {
	case modelSchemaRead:
		query(modelSchemaRead, smallQL, schemasFilesName, w, r)
		return true
	case modelNoGo:
		modelPrint("modelNoGo hit. rowquery(" + path + ")")
		return false
	default:
		return false
	}
}

func query(path, smallQl, id string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("query( " + path + " )")
	switch path {
	case modelSchemaRead:
		renderStatic(modelSchemaRead+"/"+schemasFilesName+modelMarkup, w, r)
		return true
	case modelNoGo:
		return false
	default:
		return false //woah wtf why
	}
}

func updateQuery(path, id string, w http.ResponseWriter, r *http.Request) bool {
	//test and make sure the path exists
	if testModelSchema(path, "update", w, r) {
		//write into the Database
		modelPrint("updateQuery( " + path + " ) runnning...")
		return uQuery(path, schemasFilesName, w, r)
	}
	return true
}

func uQuery(path, schemasFileName string,
	w http.ResponseWriter, r *http.Request) bool {
	// use the query path to determine which data to update
	// then use the Request to figure out what to put there.
	// return said database query as a renderStatic(path)

	return true
}
