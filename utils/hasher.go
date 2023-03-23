package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func HashString(rawString string) string {
	sha := sha512.New()
	sha.Write([]byte(rawString))
	return hex.EncodeToString(sha.Sum(nil))
}

func PaymentHasher(credes Creds, params map[string]interface{}) string {
	var udfFields string
	for i := 1; i <= 5; i++ {
		if val, ok := params[fmt.Sprintf("udf%d", i)]; ok {
			udfFields += fmt.Sprintf("%v|", val)
		} else {
			udfFields += "|"
		}
	}
	hashString := fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v|||||%v", credes.Key, params["txnid"], params["amount"], params["productinfo"], params["firstname"], params["email"], udfFields, credes.Salt)

	if _, ok := params["additionalCharges"]; ok {
		hashString += fmt.Sprintf("|%v", params["additionalCharges"])
	}
	return HashString(hashString)
}

func ReverseHasher(credes Creds, params map[string]interface{}) (string,error) {
	mandatoryParams := []string{"status","txnid", "amount", "productinfo", "firstname"}
	for _, paramName := range mandatoryParams {
			if _, ok := params[paramName]; !ok {
					return "", fmt.Errorf("missing mandatory parameter %q", paramName)
			}
	}
	udf1, udf2, udf3, udf4, udf5 := "", "", "", "", ""
	if value, ok := params["udf1"].(string); ok {
		udf1 = value
	}
	if value, ok := params["udf2"].(string); ok {
		udf2 = value
	}
	if value, ok := params["udf3"].(string); ok {
		udf3 = value
	}
	if value, ok := params["udf4"].(string); ok {
		udf4 = value
	}
	if value, ok := params["udf5"].(string); ok {
		udf5 = value
	}
	hashString := fmt.Sprintf("%v|%v||||||%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v", credes.Salt, params["status"], udf5, udf4, udf3, udf2, udf1, params["email"], params["firstname"], params["productinfo"], params["amount"], params["txnid"], credes.Key)
	if _, ok := params["additionalCharges"]; ok {
		hashString = fmt.Sprintf("%v|%v", params["additionalCharges"], hashString)
	}
	return HashString(hashString),nil
}

func CheckReversehash(credes Creds, params map[string]interface{}) (bool, error) {
	hash, err := ReverseHasher(credes, params)
	if err != nil {
		return false, err
	}
	if hash == params["hash"] {
		return true, nil
	}
	return false, nil
}

func ApiHasher(cred Creds, params ApiStruct) string {
	rawString := fmt.Sprintf("%v|%v|%v|%v", cred.Key, params.Command, params.Var1, cred.Salt)
	return HashString(rawString)
}
