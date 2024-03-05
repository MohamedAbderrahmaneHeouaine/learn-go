package auth

import (
		"net/http"
	    "strings"
		)

func GetAPIKey(headers http.Header)(string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no aithentification info found")
	}

	vals := strings.Split(val, " ")
	if len(vals)!= 2 {
        return "", errors.New("invalid aithentification info")
    }	
	if vals[0]!= "ApiKey" {
		return "", errors.New("malformed fist part of auth header")
	}
	return vals[1], nil
}