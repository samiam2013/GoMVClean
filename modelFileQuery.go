package GoMVClean

import (
	"io/ioutil"
	"log"
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

func uQuery(writePriority bool, path, schemaFileName string,
	w http.ResponseWriter, r *http.Request) bool {
	// use the query path to determine which data to update
	// then use the Request to figure out what to put there.
	// return said database query as a renderStatic(path)
	if query(path, modelSchemaPub, schemaFileName, w, r) {
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
	if schemaFileName == schemaFile {
		return modeljsonVerifyWrite(path, w, r)
	}
	defer log.Fatalln(schemaFailString)
	errorShortCircuit(w, r, modelFAIL)
	return false
}

func modeljsonVerifyWrite(path string,
	w http.ResponseWriter, r *http.Request) bool {
	// check one last time that we're writing to the request location
	urlPath := r.URL.Path[len(path):]
	if len(urlPath) > 1 { // needs to be changed to json check algo
		//let's write to the request location
		// [the royal let's] : * DUDE abides. *
		jsonString, jsonErr := loadStaticBody(urlPath)
		if jsonErr {
			defer log.Fatalln(fatalLoadString)
			errorShortCircuit(w, r, modelRenderFAIL)
			return false
		}
		ioutil.WriteFile(path, jsonString, 0642)
	}
	defer log.Fatalln()
	errorShortCircuit(w, r, modelFAIL)
	return true
}
