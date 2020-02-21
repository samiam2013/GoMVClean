package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//model render fail constatnt
const modelRenderFAIL = "404"

const fatalLoadString = "hey sysadmins, \n" +
	"GoMVClean-> modelQuery.go-> modelStaticWrite( urlPath ) \n" +
	"-> loadStaticBody( urlPath )\n" +
	"-> static.go ->  ioutil.Readfile( urlPathv ) \n" +
	"That failed with a " + modelRenderFAIL + "..."

const fatalWriteString = "hey sysadmins, BIG PROBLEM, someone is using \n" +
	"modelStaticWrite() in GoMVClean -> modelQuery.go \n" +
	"THIS IS ONLY SUPPOSED TO BE USED BY modelWrite() \n" +
	"because of safety mechnisms put in place. \n" +
	"GODSPEED MAN, 73s W9USI over and OUT \n" +
	"That failed with a " + modelFAIL + "..."

const schemaFailString = "hey sysadmins, SOMEONE tried to write to \n" +
	"SOMEWHERE that's not in GoMVClean -> model.go's -> const schemaFileName string" +
	"That failed with a " + modelFAIL + "..."

func uQuery(writePriority bool, path, data, schemaFileName string,
	w http.ResponseWriter, r *http.Request) bool {
	// use the query path to determine which data to update
	// then use the Request to figure out what to put there.
	// return said database query as a renderStatic(path)
	if writePriority {
		//here's where whe write post data
		// we're assuming the uQuery Programmers know what this means
		// here, you can write uQuery(true, "public/schema/uploadStream",
		//																			"allPaths(), w, r)
		// and you'd recursively duplicate the site OUT OF EXISTENCE
		// SO BE CAREFUL
		return modelWrite(path, schemaFileName, data, w, r)
	}
	return query(path, "read", "schema.json", w, r)
}

func modelWrite(path, schemaFileName, data string,
	w http.ResponseWriter, r *http.Request) bool {
	//try to write:Static to the model
	if schemaFileName == schemaFile {
		return modeljsonVerifyWrite(path, data, w, r)
	}
	//errorShortCircuit(w, r, modelFAIL)
	fmt.Println("modelWrite(", path, ",", schemaFileName, ", w , r ) failed.")
	return false
}

func modeljsonVerifyWrite(path, data string,
	w http.ResponseWriter, r *http.Request) bool {
	dataByte := []byte(data)
	err := ioutil.WriteFile(path, dataByte, 0642)
	if err != nil {
		fmt.Println(path, " couldn't be written to")
		return false
	}
	//fmt.Println(fatalLoadString)
	//errorShortCircuit(w, r, modelFAIL)
	return true
}
