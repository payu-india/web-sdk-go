package wrappers

import (
	"bytes"
	"fmt"
	"github.com/payu-india/web-sdk-go/utils"
)


func GeneratePaymentForm(creds utils.Creds, apiEndPoint string, params map[string]interface{}) (string, error) {
	mandatoryParams := []string{"txnid", "amount", "productinfo", "firstname", "email", "phone", "surl", "furl"}
	for _, paramName := range mandatoryParams {
			if _, ok := params[paramName]; !ok {
					return "", fmt.Errorf("missing mandatory parameter %q", paramName)
			}
	}
	hash := utils.PaymentHasher(creds, params)
	params["hash"] = hash
	params["key"] = creds.Key

	var inputs bytes.Buffer
	inputs.WriteString("<form name=\"payment_post\" id=\"payment_post\" action=\"" + apiEndPoint + "\" method=\"post\">")
	for key, value := range params {
		switch v := value.(type) {
    case int:
        inputs.WriteString(fmt.Sprintf("<input hidden type=\"text\" name=\"%s\" value=\"%d\"/>", key, v))
    case float64:
        inputs.WriteString(fmt.Sprintf("<input hidden type=\"text\" name=\"%s\" value=\"%f\"/>", key, v))
    default:
        inputs.WriteString(fmt.Sprintf("<input hidden type=\"text\" name=\"%s\" value=\"%s\"/>", key, value))
    }
	}
	inputs.WriteString("</form><script type=\"text/javascript\"> window.onload=function(){document.forms['payment_post'].submit();}</script>")

	return inputs.String(), nil
}