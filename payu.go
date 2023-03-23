package main

import (
	"errors"
	"github.com/payu-india/web-sdk-go/utils"
	"github.com/payu-india/web-sdk-go/wrappers"
)

type PayuStruct struct {
	creds utils.Creds
	urls  Url
	env   string
}

type Url struct {
	paymentURL string
	apiURL     string
}

func Payu(key string, salt string, env string) (*PayuStruct, error) {
	//initiaze url struct
	url := Url{}
	//set the urls based on the env
	switch env {
	case "TEST":
		url.paymentURL = "https://test.payu.in/_payment"
		url.apiURL = "https://test.payu.in/merchant/postservice.php?form=2"
	case "PROD":
		url.paymentURL = "https://secure.payu.in/_payment"
		url.apiURL = "https://info.payu.in/merchant/postservice.php?form=2"
	default:
		return nil, errors.New("Invalid ENV passed: possible value are TEST and PROD")
	}
	return &PayuStruct{creds: utils.Creds{Key: key, Salt: salt}, env: env, urls: url}, nil
}

func (p *PayuStruct) GeneratePaymentForm(params map[string]interface{}) (string, error) {
	return wrappers.GeneratePaymentForm(p.creds, p.urls.paymentURL, params)
}

func (p *PayuStruct) VerifyPayment(txnid string) (map[string]interface{}, error) {
	return wrappers.VerifyPayment(p.creds, p.urls.apiURL, txnid)
}

func (p *PayuStruct) GetTransactionDetails(startdate string, enddate string) (map[string]interface{}, error) {
	return wrappers.GetTransactionDetails(p.creds, p.urls.apiURL, startdate, enddate)
}

func (p *PayuStruct) ValidateVPA(vpa string) (map[string]interface{}, error) {
	return wrappers.ValidateVPA(p.creds, p.urls.apiURL, vpa)
}

func (p *PayuStruct) CancelRefundTransaction(payuid string, tokenID string, amount float64) (map[string]interface{}, error) {
	return wrappers.CancelRefundTransaction(p.creds, p.urls.apiURL, payuid, tokenID, amount)
}

func (p *PayuStruct) CheckActionStatus(requestID string) (map[string]interface{}, error) {
	return wrappers.CheckActionStatus(p.creds, p.urls.apiURL, requestID)
}

func (p *PayuStruct) GetIssuingBankStatus(bin int) (map[string]interface{}, error) {
	return wrappers.GetIssuingBankStatus(p.creds, p.urls.apiURL, bin)
}

func (p *PayuStruct) CheckIsDomestic(bin int) (map[string]interface{}, error) {
	return wrappers.CheckIsDomestic(p.creds, p.urls.apiURL, bin)
}

func (p *PayuStruct) EligibleBinsForEMI(bin int) (map[string]interface{}, error) {
	return wrappers.EligibleBinsForEMI(p.creds, p.urls.apiURL, bin)
}

func (p *PayuStruct) GetEmiAmountAccordingToInterest(amount float64) (map[string]interface{}, error) {
	return wrappers.GetEmiAmountAccordingToInterest(p.creds, p.urls.apiURL, amount)
}

func (p *PayuStruct) GetNetbankingStatus(pg string) (map[string]interface{}, error) {
	return wrappers.GetNetbankingStatus(p.creds, p.urls.apiURL, pg)
}

func (p *PayuStruct) CheckReversehash(params map[string]interface{}) (bool, error) {
	return utils.CheckReversehash(p.creds, params)
}

func (p *PayuStruct) ExpireInvoice(txnid string) (map[string]interface{}, error) {
	return wrappers.ExpireInvoice(p.creds, p.urls.apiURL, txnid)
}

func (p *PayuStruct) GetSettlementDetails(dateOrUTR string) (map[string]interface{}, error) {
	return wrappers.GetSettlementDetails(p.creds, p.urls.apiURL, dateOrUTR)
}

func (p *PayuStruct) CreateInvoice(params map[string]interface{}) (map[string]interface{}, error) {
	return wrappers.CreateInvoice(p.creds, p.urls.apiURL, params)
}

func (p *PayuStruct) GetCheckoutDetails(params map[string]interface{}) (map[string]interface{}, error) {
	return wrappers.GetCheckoutDetails(p.creds, p.urls.apiURL, params)
}
