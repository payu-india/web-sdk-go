package wrappers

import (
	"github.com/payu-india/web-sdk-go/http"
	"github.com/payu-india/web-sdk-go/utils"
	"net/url"
)

// VerifyPayment is a wrapper for the verify_payment command
// It takes in the credentials and the var1 value and returns the response as a dictionary
func GetTransactionDetails(creds utils.Creds,apiEndPoint string, var1 string, var2 string ) (map[string]interface{}, error) {
	command := "get_Transaction_Details"
	// Create the payload
	payload := url.Values{
		"key": {creds.Key},
		"command": {command},
		"var1": {var1},
		"var2": {var2},
		"hash": {utils.ApiHasher(creds, utils.ApiStruct{Command: command, Var1: var1})},
	}

	// Send the request and get the response
	response, err := http.PostForm(apiEndPoint, payload)
	if err != nil {
		return nil, err
	}

	// Return the response
	return response, nil
}