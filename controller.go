package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//global debugger constant
const globalDebug bool = true

const hostFQDN = "localhost"
const httpsPort = ":443"
const tlsFolder = "TLS" + string(os.PathSeparator)
const tlsKey = tlsFolder + "snakeoil.key"
const tlsCert = tlsFolder + "snakeoil.cert"

// const CSSPath set in resourceLibrary/css.go

func main() {

	http.HandleFunc(indexPath, routeIndex) // found in index.go
	http.HandleFunc(viewPath, routeView)   // found in view.go
	http.HandleFunc(modelPath, routeModel) // found in model.go
	http.HandleFunc(apiPath, routeAPI)     // found in api.go

	http.HandleFunc(jsPath, routeStaticFile)      // found in routeStaticFile.go
	http.HandleFunc(cssPath, routeStaticFile)     // ...^
	http.HandleFunc(faviconPath, routeStaticFile) // ...^
	http.HandleFunc(robotsPath, routeStaticFile)  // ...^
	http.HandleFunc(humansPath, routeStaticFile)  // ...^

	// warn users that the snakeoil keys need to be replaced
	fmt.Println("GoMvClean beta running...")
	if tlsKey == tlsFolder+"snakeoil.key" ||
		tlsCert == tlsFolder+"snakeoil.cert" {
		if globalDebug {
			fmt.Println("You need to replace the snakeoil TLS keys in ./TLS/" +
				"\n You also need to add the name of your new keys in controller.go")
		}
	}

	if globalDebug {
		testEverything(true) //found in test.go
	}

	// start https server
	err := http.ListenAndServeTLS(httpsPort, tlsCert, tlsKey, nil)
	if err != nil {
		log.Fatal("Server Error: ", err)
	}

	return
}
