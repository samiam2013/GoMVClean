package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samiam2013/GoMVClean/modelconfig"
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
	if globalDebug {
		testEverything(true) //found in test.go
	}
	http.HandleFunc(indexPath, routeIndex)   // found in index.go
	http.HandleFunc(errorsPath, routeError)  // found in errors.go
	http.HandleFunc(staticPath, routeStatic) // found in view.go
	http.HandleFunc(jsPath, routeJS)         // found in js.go
	http.HandleFunc(cssPath, routeCSS)       // found in css.go
	http.HandleFunc(modelPath, routeModel)   // found in model.go
	http.HandleFunc(apiPath, routeAPI)       // found in api.go

	http.HandleFunc(faviconPath, routeIdiosync) // found in idioSync.go
	http.HandleFunc(robotsPath, routeIdiosync)  // ...^
	http.HandleFunc(humansPath, routeIdiosync)  // ...^

	// warn users that the snakeoil keys need to be replaced
	fmt.Println("GoMvClean beta running...")
	if tlsKey == tlsFolder+"snakeoil.key" ||
		tlsCert == tlsFolder+"snakeoil.cert" {
		if globalDebug {
			fmt.Println("You need to replace the snakeoil TLS keys in ./TLS/" +
				"\n You also need to add the name of your new keys in controller.go")
		}
	}

	// call the model configure script
	modelConfig.ModelInit()

	// start https server
	err := http.ListenAndServeTLS(httpsPort, tlsCert, tlsKey, nil)
	if err != nil {
		log.Fatal("TLS Error: ", err)
	}

	return
}
