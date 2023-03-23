# PayU SDK for GO

The PayU SDK for GO enables you to easily work with APIs of PayU by easily integrating this SDK within your base system.
With our SDK, you do not need to worry about low level details for doing API integration and with few lines of code and a function calling, it can be done within few minutes.

## Features Supported
Following features are supported in the PayU GO web SDK:
 - Create Payment form.
 - Verify transactions, check the transaction status, transaction details, or discount rate for a transaction
 - Initiated refunds, cancel refund, check refund status.
 - Retrieve settlement details which the bank has to settle you.
 - Get information of the payment options, offers, recommendations, and downtime details.
 - Check the customer’s eligibility for EMI and get the EMI amount according to interest 
 - Pay by link genration

To get started with PayU, visit our [Developer Guide](https://devguide.payu.in/low-code-web-sdk/getting-started-low-code-web-sdk/)

# Table of Contents
 1. [Usage](#Usage)
 2. [Getting Started](#getting-started)
 3. [Documentation for various Methods](#documentation-for-various-methods)

## Usage

```shell
import (
  Payu "https://github.com/payu-india/web-sdk-go"
)
```


## Getting Started

Please follow the [installation](#installation) instruction and execute the following JS code for creating the instance of PayU Object:

```go
import (
  PayU "https://github.com/payu-india/web-sdk-go"
)
payu, err := Payu(
  <YOUR_MERCHANT_KEY>,
  <YOUR_MERCHANT_SALT>,
  <ENVIRONMENT>                 // Possible value  = TEST/LIVE
)     

```
## Documentation for various Methods 


Method |  Description
------------- | -------------
[**GeneratePaymentForm**](docs/initiate.md)  | Genereate auto submit HTML form to intitiate transaction 
[**VerifyPayment**](docs/verifyPayment.md) | Provides the details of a transaction  
[**GetTransactionDetails**](docs/getTransactionDetails.md) | Provides the details of a transactions for a specfic timeperiod
[**ValidateVPA**](docs/validateVPA.md) | Used to validate VPA of a user. 
[**CancelRefundTransaction**](docs/cancelRefundTransaction.md) | Initiate refunds. 
[**CheckActionStatus**](docs/checkActionStatus.md) | Check the status of a refund.  
[**GetNetbankingStatus**](docs/getNetbankingStatus.md) | Check downtime status of PGs. 
[**GetIssuingBankStatus**](docs/getIssuingBankStatus.md) | Check downtime through bin number. 
[**CheckIsDomestic**](docs/checkIsDomestic.md) | Check the bin information
[**CreateInvoice**](docs/createInvoice.md) |  Used to create email and SMS invoice ( Pay by link ).
[**ExpireInvoice**](docs/expireInvoice.md) | Used to expire an existing invoice.
[**EligibleBinsForEMI**](docs/eligibleBinsForEMI.md) |  Used for checking the card eligibilty for EMI through the bin number.
[**GetEmiAmountAccordingToInterest**](docs/getEmiAmountAccordingToInterest.md) | Used to fetch interest accordign to Banks and tenure.
[**GetSettlementDetails**](docs/getSettlementDetails.md) |  Used to fetch settlement details for a particular date.
[**GetCheckoutDetails**](docs/getCheckoutDetails.md) |  Used to fetch payment options, offers, eligibility, recommendations, and downtime details.