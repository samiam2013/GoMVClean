package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//global debugger constant
const gDebug bool = true

const httpsPort string = ":443"
const indexPath string = "/"
const indexName string = "index"
const indexPathName string = indexPath + indexName
const indexFAIL = "404"

const tlsPath = "TLS"
const tlsFolder = tlsPath + string(os.PathSeparator)
const tlsKey = tlsFolder + "snakeoil.key"
const tlsCert = tlsFolder + "snakeoil.cert"

// if path ("/"|"/indexName") serve indexName+staticMarkupType otherwise 404
func routeIndex(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == indexPath || path == indexPathName {
		path = staticMarkupFolder + indexName + staticMarkupType
		renderStatic(path, w, r)
		return
	}
	errorShortCircuit(w, r, indexFAIL)
	return
}

func main() {
	http.HandleFunc(indexPath, routeIndex)
	http.HandleFunc(errorsPath, routeError)  // found in errors.go
	http.HandleFunc(staticPath, routeStatic) // found in static.go
	http.HandleFunc(viewPath, routeView)     // found in view.go
	http.HandleFunc(modelPath, routeModel)   // found in model.go
	http.HandleFunc(apiPath, routeAPI)       // found in api.go
	fmt.Println("GoMvClean v42 running...")
	err := http.ListenAndServeTLS(httpsPort, tlsCert, tlsKey, nil)
	if err != nil {
		fmt.Println("priv key path: ", tlsKey)
		fmt.Println("pub key path: ", tlsCert)
		log.Fatal("TLS Error: ", err)
	}
	if gDebug {
		testEverything(true) //found in test.go
	}
	return
}
