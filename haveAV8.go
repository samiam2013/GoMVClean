package main

import (
	"fmt"
	"net/http"
)

func haveAV8(wholePath string,
	w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "heres where V8's console will be summoned")
	return
}
