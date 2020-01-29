package main

import (
	"fmt"
	"net/http"
)

const modelPath string = "/model/"
const publicDBPath string = "public/"
const privateDBPath string = "private/"
const modelTestPath string = "public/test"
const modelNoGo string = "private/test"
const modelPubPath string = modelPath + publicDBPath
const modelPrivPath string = modelPath + privateDBPath
const pubAccessPrivDB bool = false

func routeModel(w http.ResponseWriter, r *http.Request) {
	//url path is the modelQueryPath
	fmt.Println("routing model!...")
	path := r.URL.Path[len(modelPath):]
	renderModel(path, w, r)
	return
}

func renderModel(path string, w http.ResponseWriter, r *http.Request) {
	//fmt.Println("routeModel(", path, ")")
	wholePath := r.URL.Path
	if pubAccessPrivDB {
		if tryPath(path, privateDBPath, wholePath) {
			fmt.Println("tryPath(", path, ") returned false")
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
	//fmt.Println("tryPath(", outerPath, ",", innerPath, ",", wholePath, ")")
	queryInnerPathStr := outerPath[len(innerPath):]
	//fmt.Println(queryInnerPathStr)
	queryPathLen := len(queryInnerPathStr)
	//fmt.Println(queryPathLen)
	iPathTrailLen := len(outerPath[len(innerPath):])
	//fmt.Println(iPathTrailLen)
	return (iPathTrailLen == queryPathLen)
}

func tryPathRunFunc(w http.ResponseWriter, r *http.Request,
	outerPath, innerPath string,
	runFunction func(string, string, http.ResponseWriter, *http.Request) bool) bool {
	//fmt.Println("tryPathRunFunc(", outerPath, ",", innerPath, ", func())")
	if len(outerPath) > len(innerPath) {
		return runFunction(innerPath, outerPath, w, r)
	}
	return false
}

func privateQuery(path, wholePath string, w http.ResponseWriter,
	r *http.Request) bool {
	fmt.Println("privateQuery(path: ", path, ")")
	fmt.Fprintf(w, "{Database: '%s', table: '%s'}\n", path, wholePath)
	return pubAccessPrivDB
}

func publicQuery(path, wholePath string, w http.ResponseWriter,
	r *http.Request) bool {
	//fmt.Println("publicQuery(path:", path, ") called.")
	//fmt.Fprintf(w, "{Database: '%s', table: '%s'}\n", path, wholePath)
	return tableQuery(path, wholePath, w, r)
}

func tableQuery(path, wholePath string, w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("tableQuery(", wholePath, ")")
	switch wholePath {
	case "public/":
		fmt.Fprint(w, "{'query found, public model test case!''}")
		return true
	case modelTestPath:
		return testModelSchema(wholePath, true, w, r)
	default:
		fmt.Println("error: path not in tableQuery(", wholePath, ")")
		errorShortCircuit(w, r, "403")
	}
	return false
}

func testModelSchema(path string, verbose bool, w http.ResponseWriter, r *http.Request) bool {
	switch path {
	case modelTestPath:
		return rowQuery(modelTestPath, w, r)
	case modelNoGo:
		return rowQuery(modelNoGo, w, r)
	default:
		return false
	}
}

func rowQuery(path string, w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("rowQuery(", path, ")")
	switch path {
	case modelTestPath:
		query(modelTestPath, w, r)
		return true
	case modelNoGo:
		return false
	default:
		return false
	}
}

func query(path string, w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("query( ", path, " )")
	switch path {
	case modelTestPath:
		//now it's finally safe to reach straight into the public files, ideally.
		//that's where hashing comes in
		renderStatic(path+"/0.json", w, r)
		return true //not necessarily :()
	case modelNoGo:
		return false
	default:
		return false //woah wtf
	}
}
