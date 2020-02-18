package GoMVClean

import (
	"net/http"
)

func rowQuery(path, smallQL string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("rowQuery( " + path + " )")
	switch path {
	case modelSchemaPub:
		//found in breakStuff.go, I bring you query()
		query(modelSchemaPub, smallQL, schemaFile, w, r)
		return true
	case modelNoGo:
		modelPrint("modelNoGo hit. rowquery(" + path + ")")
		return false
	default:
		return false
	}
}

func query(path, smallQl, schema string, w http.ResponseWriter, r *http.Request) bool {
	modelPrint("query(" + path + ")")
	switch path {
	case modelSchemaPub:
		renderStatic(modelSchemaPub+"/"+schema+modelMarkup, false, w, r)
		return true
	case modelNoGo:
		if pubAccessPrivDB {
			//code for go-around

		}
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
		return uQuery(true, path, schemaFile, w, r)
	}
	return true
}
