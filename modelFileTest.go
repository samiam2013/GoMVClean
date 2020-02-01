package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//see if this is modelPath/tableName/tableNameN
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
	//fmt.Fprintf(w, "{Database: '%s', table: '%s'}\n", path, wholePath)
	return tableQuery(path, wholePath, w, r)
}

// run rowQuery against DBPath
func tableQuery(path, wholePath string,
	w http.ResponseWriter, r *http.Request) bool {
	modelPrint("tableQuery( " + wholePath + " )")
	smallQL := sQLRead
	switch wholePath {
	case modelPubPath:
		//if head says POST, go to "update mode"
		contLenStr := r.Header.Get("Content-Length")
		contLen, _ := strconv.Atoi(contLenStr)
		//fmt.Println("post: ",contLen)
		if contLen == 0 {
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
	case modelSchemaRead:
		return rowQuery(modelSchemaRead, smallQL, w, r)
	default:
		return false
	}
}
