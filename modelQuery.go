package main

import (
	"net/http"
)

func rowQuery(path, smallQL string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("rowQuery( " + path + " )")
	switch path {
	case modelSchemaRead:
		//found in breakStuff.go
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
		//found in breakStuff.go
		return uQuery(true, path, schemasFilesName, w, r)
	}
	return true
}
