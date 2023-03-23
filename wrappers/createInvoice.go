package wrappers

import (
	"encoding/json"
	"net/url"
	"github.com/payu-india/web-sdk-go/http"
	"github.com/payu-india/web-sdk-go/utils"
)

// VerifyPayment is a wrapper for the verify_payment command
// It takes in the credentials and the var1 value and returns the response as a dictionary
func CreateInvoice(creds utils.Creds, apiEndPoint string, params map[string]interface{}) (map[string]interface{}, error) {
	command := "create_invoice"
	//Marshal map to jsons string
	mJson, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var1 := string(mJson)
	// Create the payload

	payload := url.Values{
		"key":     {creds.Key},
		"command": {command},
		"var1":    {var1},
		"hash":    {utils.ApiHasher(creds, utils.ApiStruct{Command: command, Var1: var1})},
	}

	// Send the request and get the response
	response, err := http.PostForm(apiEndPoint, payload)
	if err != nil {
		return nil, err
	}

	// Return the response
	return response, nil
}
