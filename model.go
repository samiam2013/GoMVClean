package main

import (
	"fmt"
	"net/http"
)

const modelPath string = "/model/"
const publicDBPath string = "public/"
const privateDBPath string = "private/"
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

func tryPathRunFunc(w http.ResponseWriter, r *http.Request, outerPath, innerPath string,
	runFunction func(string, string, http.ResponseWriter, *http.Request) bool) bool {
	//fmt.Println("tryPathRunFunc(", outerPath, ",", innerPath, ", func())")
	if len(outerPath) > len(innerPath) {
		return runFunction(innerPath, outerPath, w, r)
	}
	return false
}

func privateQuery(path, wholePath string, w http.ResponseWriter, r *http.Request) bool {
	fmt.Fprintf(w, "{Database: '%s', table: '%s'}\n", path, wholePath)
	return pubAccessPrivDB
}

func publicQuery(path, wholePath string, w http.ResponseWriter, r *http.Request) bool {
	//fmt.Println("publicQuery(path:", path, ") called.")
	fmt.Fprintf(w, "{Database: '%s', table: '%s'}\n", path, wholePath)

	return true
}

func testModelSchema(verbose bool) {

}
