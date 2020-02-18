package GoMVClean

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samiam2013/GoMVClean/resourceLibrary/resourceLibrary"
)

//global debugger constant
const gDebug bool = true

const hostFQDN = "localhost"
const httpsPort string = ":443"
const tlsPath = "TLS"
const tlsFolder = tlsPath + string(os.PathSeparator)
const tlsKey = tlsFolder + "snakeoil.key"
const tlsCert = tlsFolder + "snakeoil.cert"

// const CSSPath set in resourceLibrary/css.go

func main() {
	if gDebug {
		testEverything(true) //found in test.go
	}
	http.HandleFunc(IndexPath, RouteIndex)             // found in index.go
	http.HandleFunc(ErrorsPath, RouteError)            // found in errors.go
	http.HandleFunc(StaticPath, RouteStatic)           // found in static.go
	http.HandleFunc(JsPath, RouteJS)                   // found in js.go
	http.HandleFunc(CSSPath, resourceLibrary.RouteCSS) // found in css.go
	http.HandleFunc(viewPath, routeView)               // found in view.go
	http.HandleFunc(modelPath, routeModel)             // found in model.go
	http.HandleFunc(apiPath, routeAPI)                 // found in api.go

	http.HandleFunc(faviconPath, routeIdiosync) // found in idioSync.go
	http.HandleFunc(robotsPath, routeIdiosync)  // ...^
	http.HandleFunc(humansPath, routeIdiosync)  // ...^

	fmt.Println("GoMvClean beta running...")

	err := http.ListenAndServeTLS(httpsPort, tlsCert, tlsKey, nil)
	if err != nil {
		log.Fatal("TLS Error: ", err)
	}
	return
}
