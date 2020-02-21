package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func csrfToken(w http.ResponseWriter, r *http.Request) ([]byte, string) {
	// make a timeout, path, and userIP for jsonMap
	var timeStamp int64 = time.Now().Unix()
	csrfTimeoutUnix := int64((csrfTimeoutMinutes * 60) + timeStamp)
	timeoutString := fmt.Sprintf("%d", csrfTimeoutUnix)
	formPath := r.URL.Path
	userIP := r.RemoteAddr

	// make a hash string
	hashString := sha256Token(timeStamp)

	// map the json output
	jsonMap := map[string]string{
		"token":    hashString,
		"formPath": formPath,
		"timeout":  timeoutString,
		"ipPort":   userIP,
	}

	// marshall the JSON
	jsonData, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Fprint(w, "{'error':[500: 'json mapping failed in csrfToken()']}")
		return nil, string(jsonData)
	}
	return jsonData, hashString
}
