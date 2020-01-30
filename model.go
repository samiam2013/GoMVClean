package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const modelPath string = "/model/"
const publicDBPath string = "public/"
const privateDBPath string = "private/"
const modelTestFolder string = "schema"
const modelMarkup string = ".json"

const modelTestPath string = modelPath + publicDBPath + modelTestFolder
const modelNoGo string = privateDBPath + modelTestFolder
const modelPubPath string = modelPath + publicDBPath
const modelPrivPath string = modelPath + privateDBPath
const modelPubQueryUpdatePath string = modelPubPath + modelTestFolder
const modelSchemaRead string = publicDBPath + modelTestFolder

const pubAccessPrivDB bool = false //false by default

const sQLUpdate string = "update"
const sQLRead string = "read"

const modelDEBUG = true

func modelPrint(prStr string) {
	if modelDEBUG {
		fmt.Println(prStr)
	}
}

func routeModel(w http.ResponseWriter, r *http.Request) {
	//url path is the modelQueryPath
	modelPrint("routing model!...")
	path := r.URL.Path[len(modelPath):]
	renderModel(path, w, r)
	return
}

func renderModel(path string, w http.ResponseWriter, r *http.Request) {
	modelPrint("routeModel(" + path + ")")
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
	errorShortCircuit(w, r, "403")
	return
}

func tryPath(outerPath, innerPath, wholePath string) bool {
	modelPrint("tryPath(" + outerPath + "," + innerPath + "," + wholePath + ")")
	queryInnerPathStr := outerPath[len(innerPath):]
	//modelPrint(queryInnerPathStr)
	queryPathLen := len(queryInnerPathStr)
	//modelPrint(queryPathLen)
	iPathTrailLen := len(outerPath[len(innerPath):])
	//modelPrint("iPathTrailLen:", iPathTrailLen)
	//modelPrint("queryPathLen:", queryPathLen)
	pathEqTrailLen := (iPathTrailLen == queryPathLen)
	return pathEqTrailLen
}

func tryPathRunFunc(w http.ResponseWriter, r *http.Request,
	outerPath, innerPath string,
	runFunction func(string, string, http.ResponseWriter, *http.Request) bool) bool {
	modelPrint("tryPathRunFunc( " + outerPath + ", " + innerPath + ", func())")
	if len(outerPath) > len(innerPath) {
		runFunction(innerPath, outerPath, w, r)
		return true
	}
	return false
}

func privateQuery(path, wholePath string, w http.ResponseWriter,
	r *http.Request) bool {
	modelPrint("privateQuery(path: " + path + " )")
	fmt.Fprintf(w, "{{Private : {Database: '%s', table: '%s'}, queried},\n", path, wholePath)
	rVal := tableQuery(path, wholePath, w, r)
	defer fmt.Fprintf(w, " returned: \"%v\"}", rVal)
	return rVal
}

func publicQuery(path, wholePath string, w http.ResponseWriter,
	r *http.Request) bool {
	modelPrint("publicQuery(path: " + path + " ) called.")
	//fmt.Fprintf(w, "{Database: '%s', table: '%s'}\n", path, wholePath)
	return tableQuery(path, wholePath, w, r)
}

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
		defer testModelSchema(wholePath, smallQL, w, r)
		return true
	case modelPrivPath:
		return testModelSchema(wholePath, smallQL, w, r)
	case modelTestPath:
		smallQL := sQLUpdate
		return testModelSchema(wholePath, smallQL, w, r)
	case modelSchemaRead:
		modelPrint("public/schema tableQuery( " + wholePath + " )...")
		return testModelSchema(wholePath, smallQL, w, r)
	default:
		modelPrint("error: path not in tableQuery( " + wholePath + " )")
		errorShortCircuit(w, r, "403")
	}
	return false
}

func testModelSchema(path, smallQL string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("testModelSchema(" + path + ", " + smallQL + " )")
	switch path {
	case modelTestPath:
		defer rowQuery(modelPubQueryUpdatePath, smallQL, w, r)
		return true
	case modelNoGo:
		return rowQuery(modelNoGo, smallQL, w, r)
	case modelSchemaRead:
		return rowQuery(modelSchemaRead, "read", w, r)
	default:
		return false
	}
}

func rowQuery(path, smallQL string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("rowQuery(" + path + ")")
	switch path {
	case modelTestPath:
		query(modelTestPath, smallQL, "0", w, r)
		return true
	case modelSchemaRead:
		query(modelSchemaRead, smallQL, "0", w, r)
		return true
	case modelNoGo:
		modelPrint("modelNoGo hit. rowquery(" + path + ")")
		return false
	default:
		return false
	}
}

func query(path, smallQl, id string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("query(" + path + ")")
	switch path {
	case modelTestPath:
		//now it's finally safe to reach straight into the public files, ideally.
		//that's where hashing comes in
		switch smallQl {
		case "read":
			renderStatic(path+"/"+id+modelMarkup, w, r)
			return true
		case "update":
			updateQuery(path, id, w, r)
			renderStatic(path+modelMarkup, w, r)
			return true
		default:
			return false
		}
	case modelSchemaRead:
		renderStatic(modelSchemaRead+"/"+id+modelMarkup, w, r)
		return true
	case modelNoGo:
		return false
	default:
		return false //woah wtf
	}
}

func updateQuery(path, id string, w http.ResponseWriter, r *http.Request) bool {
	//test and make sure the path exists
	if testModelSchema(path, "update", w, r) {
		//write into the Database
		modelPrint("updateQuery(" + path + ") runnning...")
		return false
	}
	return true
}
