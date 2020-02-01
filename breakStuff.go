package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func uQuery(writePriority bool, path, schemaFileName string,
	w http.ResponseWriter, r *http.Request) bool {
	// use the query path to determine which data to update
	// then use the Request to figure out what to put there.
	// return said database query as a renderStatic(path)
	if query(path, modelSchemaRead, schemaFileName, w, r) {
		//the data exists and can be read
		if writePriority {
			//here's where whe write post data
			// we're assuming the uQuery Programmers know what this means
			// here, you can write uQuery(true, "public/schema/uploadStream",
			//																			"allPaths(), w, r)
			// and you'd recursively duplicate the site OUT OF EXISTENCE
			// SO BE CAREFUL
			return modelWrite(path, schemaFileName, w, r)
		}
	}
	return true
}

func modelWrite(path, schemaFileName string,
	w http.ResponseWriter, r *http.Request) bool {
	//try to write:Static to the model
	if schemaFileName == schemasFilesName {
		return modelStaticWrite(path, w, r)
	}
	defer log.Fatalln("hey sysadmins, SOMEONE tried to write to SOMEWHERE \n" +
		"that's not in GoMVClean -> model.go's -> const schemaFileName string")
	errorShortCircuit(w, r, "403")
	return false
}

func modelStaticWrite(path string,
	w http.ResponseWriter, r *http.Request) bool {
	//let's check one last time that we're writing to the request location
	urlPath := r.URL.Path[len(path):]
	if len(urlPath) > 1 {
		//let's write to the request location
		jsonString, jsonErr := loadStaticBody(urlPath)
		if jsonErr {
			defer log.Fatalln("hey sysadmins, \n"+
				"GoMVClean-> modelQuery.go-> modelStaticWrite(", urlPath, ") \n"+
				"-> loadStaticBody(", urlPath, ")\n"+
				"-> static.go ->  ioutil.Readfile(", urlPath, ") \n"+
				"That failed with a 404...")
			errorShortCircuit(w, r, "404")
			return false
		}
		ioutil.WriteFile(path, jsonString, 0642)
	}
	defer log.Fatalln("hey sysadmins, BIG PROBLEM, someone is using \n" +
		"modelStaticWrite() in GoMVClean -> modelQuery.go \n" +
		"THIS IS ONLY SUPPOSED TO BE USED BY modelWrite() \n" +
		"because of safety mechnisms put in place. \n" +
		"GODSPEED MAN, 73s W9USI over and OUT\n")
	errorShortCircuit(w, r, "403")
	return true
}
