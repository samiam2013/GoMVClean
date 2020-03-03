package main

import (
	"fmt"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) bool {
	fmt.Println("starting user registration routine...")
	//pull fields from registration form, validate data
	// relocate (302) to form if data's not validate
	// 		implement Javascript to validate beforehand
	//			so that this doesn't have to be hit often.
	// relocate to landing page if valid,
	return false
}
