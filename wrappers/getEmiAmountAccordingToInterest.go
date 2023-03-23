package wrappers

import (
	"net/url"
	"github.com/payu-india/web-sdk-go/http"
	"github.com/payu-india/web-sdk-go/utils"
	"strconv"
)

// VerifyPayment is a wrapper for the verify_payment command
// It takes in the credentials and the var1 value and returns the response as a dictionary
func GetEmiAmountAccordingToInterest(creds utils.Creds, apiEndPoint string, amount float64) (map[string]interface{}, error) {
	command := "getEmiAmountAccordingToInterest"
	// Create the payload
	var1 := strconv.FormatFloat(amount, 'f', 2, 64)
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
