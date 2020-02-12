package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//global debugger constant
const gDebug bool = true

const hostFQDN = "localhost"
const httpsPort string = ":443"
const tlsPath = "TLS"
const tlsFolder = tlsPath + string(os.PathSeparator)
const tlsKey = tlsFolder + "snakeoil.key"
const tlsCert = tlsFolder + "snakeoil.cert"

func main() {
	http.HandleFunc(indexPath, routeIndex)   // found in index.go
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
	go concurrentRedir() 
	return
}

func redirHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+hostFQDN+httpsPort+r.RequestURI,
		http.StatusMovedPermanently)
	return
}

func concurrentRedir() {
	err := http.ListenAndServe("httpPort", http.HandlerFunc(redirHTTPS))
	if err != nil {
		log.Fatal("HTTP -> HTTPS redirect FAIL:", err)
	}
}
