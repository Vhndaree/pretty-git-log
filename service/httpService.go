package service

import (
	"log"
	"net/http"
)

// GetWithAuthorization - HTTP GET request with authorization token
func GetWithAuthorization(url, token string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return resp
}
