package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

func PostForm(url string, payload url.Values) (map[string]interface{}, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Encode the payload as x-www-form-urlencoded
	requestBody := bytes.NewBufferString(payload.Encode())

	// Create a new HTTP request with the POST method and the request body
	request, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
			return nil, err
	}

	// Set the content type header to x-www-form-urlencoded
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request and get the response
	response, err := client.Do(request)
	if err != nil {
			return nil, err
	}

	// Close the response body when we're done
	defer response.Body.Close()

	// Decode the response body into a map[string]interface{} variable
	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
			return nil , err 
	}

	// Return the response body as a dictionary
	return responseBody, nil
}
