package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func renderTestPost(body []byte, w http.ResponseWriter, r *http.Request) string {
	myID := r.PostFormValue("textField")
	myInput := r.PostFormValue("textArea")
	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	rString := string(requestDump)

	fmt.Fprint(w, "<html><body><code style=\"white-space: pre-wrap;\">")
	fmt.Fprint(w, "HTTP Request: \n\n", rString)
	fmt.Fprint(w, "Post Values: \n\nmyID: ", myID)
	fmt.Fprint(w, "\nmyInput: ", myInput)
	fmt.Fprint(w, "</code></body></html>")

	//fmt.Println("renderTestPost() called!")
	return "200"
}
