package main

import (
	"fmt"
	"net/http"
)

func uploadStream(wholePath string,
	w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "heres where we do an abritrary thing with arbitrary data")
	return
}
