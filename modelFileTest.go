package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// who you gonna go visit?
const uStreamPath string = "updateStream"
const updateStreamPath = publicDBPath + uStreamPath

// how are you going to get there?
const dStreamPath string = publicDBPath + "downloadStream/"
const dSchemaFolder string = dStreamPath + schemaFolder + "/"
const dSchemaFilePath string = dSchemaFolder + schemaFile + modelMarkup

// what happens if they're not home?
const fileFAIL = "500"

//see if this is modelPath/tableName/column
func tryPath(outerPath, innerPath, wholePath string) bool {
	modelPrint(
		"tryPath( " + outerPath + ", " + innerPath + ", " + wholePath + " )")
	queryInnerPathStr := outerPath[len(innerPath):]
	queryPathLen := len(queryInnerPathStr)
	iPathTrailLen := len(outerPath[len(innerPath):])
	pathEqTrailLen := (iPathTrailLen == queryPathLen)
	return pathEqTrailLen
}

// try to run the privateQuery or publicQuery
func tryPathRunFunc(w http.ResponseWriter, r *http.Request,
	outerPath, innerPath string,
	runFunction func(string, string, http.ResponseWriter,
		*http.Request) bool) bool {
	modelPrint("tryPathRunFunc( " + outerPath + ", " + innerPath + ", func() )")
	if len(outerPath) > len(innerPath) {
		runFunction(innerPath, outerPath, w, r)
		return true
	}
	return false
}

// run tableQuery query against privateDBPath
func privateQuery(path, wholePath string,
	w http.ResponseWriter, r *http.Request) bool {
	modelPrint("privateQuery( " + path + " )")
	fmt.Fprintf(w,
		"{ { Private : { Database: \"%s\", table: \"%s\" }, \"queried\" } \n",
		path, wholePath)
	rVal := tableQuery(path, wholePath, w, r)
	defer fmt.Fprintf(w, " returned: \"%v\"}", rVal)
	return rVal
}

// run tableQuery against publicDBPath
func publicQuery(path, wholePath string,
	w http.ResponseWriter, r *http.Request) bool {
	modelPrint("publicQuery( " + path + " )")
	return tableQuery(path, wholePath, w, r)
}

// run rowQuery against DBPath
func tableQuery(path, wholePath string,
	w http.ResponseWriter, r *http.Request) bool {
	modelPrint("tableQuery( " + wholePath + " )")
	modelPrint("dSchemaFilePath: " + dSchemaFilePath)
	smallQL := sQLRead
	switch wholePath {
	case updateStreamPath:
		routeView(w, r)
		return true
	case dSchemaFilePath:
		//here we just serve the schema for the downloadStreamPath
		renderStatic(dSchemaFilePath, false, w, r)
		return true
	case modelPubPath:
		//if head says POST, go to "update mode"
		contLenStr := r.Header.Get("Content-Length")
		contLen, _ := strconv.Atoi(contLenStr)
		if contLen == 0 {
			smallQL = sQLRead
		} else {
			smallQL = sQLUpdate
		}
		testModelSchema(wholePath, smallQL, w, r)
		return true
	case modelPrivPath:
		return testModelSchema(wholePath, smallQL, w, r)
	case modelSchemaPub:
		return testModelSchema(wholePath, smallQL, w, r)
	default:
		modelPrint("error: path not in tableQuery( " + wholePath + " )")
		errorShortCircuit(w, r, modelFAIL)
	}
	return false
}

//test the schema to make sure you're still safe
func testModelSchema(path, smallQL string,
	w http.ResponseWriter, r *http.Request) bool {
	modelPrint("testModelSchema( " + path + ", " + smallQL + " )")
	switch path {
	case modelNoGo:
		return rowQuery(modelNoGo, smallQL, w, r)
	case modelSchemaPub:
		return rowQuery(modelSchemaPub, smallQL, w, r)
	default:
		return false
	}
}
