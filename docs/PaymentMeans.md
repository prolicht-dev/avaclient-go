# PaymentMeans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIdentifier** | Pointer to **string** | BT-84 This is typically the IBAN or account number BT-84 in XRechnung. | [optional] 
**AccountName** | Pointer to **string** | BT-85 The name for the account, typically the name of the owner for a bank account. BT-85 in XRechnung. | [optional] 
**ServiceProviderIdentifier** | Pointer to **string** | BT-86 This can be the BIC BT-86 in XRechnung. This will be saved in BG-17. | [optional] 
**DirectDebitMandateId** | Pointer to **string** | BT-89 If used, this is the mandate id for direct debit under which the direct debit will be performed BT-89 in XRechnung. This will be saved in BG-19. | [optional] 
**DirectDebitIban** | Pointer to **string** | BT-91 If used, this is the IBAN for direct debit BT-91 in XRechnung. | [optional] 
**PaymentId** | Pointer to **string** | BT-83 The payment id, e.g. a reference number for the payment. In German, this would be called \&quot;Verwendungszweck\&quot;. BT-83 in XRechnung. | [optional] 
**PaymentCardHolderName** | Pointer to **string** | BT-88 If this is a card payment, then this is the name of the cardholder. BT-88 in XRechnung. | [optional] 
**PaymentCardPrimaryAccountNumber** | Pointer to **string** | BT-87 If this is a card payment, then this is the type of card. BT-87 in XRechnung. This will be saved in BG-18. | [optional] 
**PaymentCardNetworkId** | Pointer to **string** | If this is a card payment, then this is the type of card. This has no mapping yet in XRechnung, see: https://github.com/itplr-kosit/validator-configuration-xrechnung/issues/53 | [optional] 

## Methods

### NewPaymentMeans

`func NewPaymentMeans() *PaymentMeans`

NewPaymentMeans instantiates a new PaymentMeans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMeansWithDefaults

`func NewPaymentMeansWithDefaults() *PaymentMeans`

NewPaymentMeansWithDefaults instantiates a new PaymentMeans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIdentifier

`func (o *PaymentMeans) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *PaymentMeans) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *PaymentMeans) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.

### HasAccountIdentifier

`func (o *PaymentMeans) HasAccountIdentifier() bool`

HasAccountIdentifier returns a boolean if a field has been set.

### GetAccountName

`func (o *PaymentMeans) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentMeans) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentMeans) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PaymentMeans) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetServiceProviderIdentifier

`func (o *PaymentMeans) GetServiceProviderIdentifier() string`

GetServiceProviderIdentifier returns the ServiceProviderIdentifier field if non-nil, zero value otherwise.

### GetServiceProviderIdentifierOk

`func (o *PaymentMeans) GetServiceProviderIdentifierOk() (*string, bool)`

GetServiceProviderIdentifierOk returns a tuple with the ServiceProviderIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderIdentifier

`func (o *PaymentMeans) SetServiceProviderIdentifier(v string)`

SetServiceProviderIdentifier sets ServiceProviderIdentifier field to given value.

### HasServiceProviderIdentifier

`func (o *PaymentMeans) HasServiceProviderIdentifier() bool`

HasServiceProviderIdentifier returns a boolean if a field has been set.

### GetDirectDebitMandateId

`func (o *PaymentMeans) GetDirectDebitMandateId() string`

GetDirectDebitMandateId returns the DirectDebitMandateId field if non-nil, zero value otherwise.

### GetDirectDebitMandateIdOk

`func (o *PaymentMeans) GetDirectDebitMandateIdOk() (*string, bool)`

GetDirectDebitMandateIdOk returns a tuple with the DirectDebitMandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebitMandateId

`func (o *PaymentMeans) SetDirectDebitMandateId(v string)`

SetDirectDebitMandateId sets DirectDebitMandateId field to given value.

### HasDirectDebitMandateId

`func (o *PaymentMeans) HasDirectDebitMandateId() bool`

HasDirectDebitMandateId returns a boolean if a field has been set.

### GetDirectDebitIban

`func (o *PaymentMeans) GetDirectDebitIban() string`

GetDirectDebitIban returns the DirectDebitIban field if non-nil, zero value otherwise.

### GetDirectDebitIbanOk

`func (o *PaymentMeans) GetDirectDebitIbanOk() (*string, bool)`

GetDirectDebitIbanOk returns a tuple with the DirectDebitIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebitIban

`func (o *PaymentMeans) SetDirectDebitIban(v string)`

SetDirectDebitIban sets DirectDebitIban field to given value.

### HasDirectDebitIban

`func (o *PaymentMeans) HasDirectDebitIban() bool`

HasDirectDebitIban returns a boolean if a field has been set.

### GetPaymentId

`func (o *PaymentMeans) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentMeans) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentMeans) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PaymentMeans) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentCardHolderName

`func (o *PaymentMeans) GetPaymentCardHolderName() string`

GetPaymentCardHolderName returns the PaymentCardHolderName field if non-nil, zero value otherwise.

### GetPaymentCardHolderNameOk

`func (o *PaymentMeans) GetPaymentCardHolderNameOk() (*string, bool)`

GetPaymentCardHolderNameOk returns a tuple with the PaymentCardHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCardHolderName

`func (o *PaymentMeans) SetPaymentCardHolderName(v string)`

SetPaymentCardHolderName sets PaymentCardHolderName field to given value.

### HasPaymentCardHolderName

`func (o *PaymentMeans) HasPaymentCardHolderName() bool`

HasPaymentCardHolderName returns a boolean if a field has been set.

### GetPaymentCardPrimaryAccountNumber

`func (o *PaymentMeans) GetPaymentCardPrimaryAccountNumber() string`

GetPaymentCardPrimaryAccountNumber returns the PaymentCardPrimaryAccountNumber field if non-nil, zero value otherwise.

### GetPaymentCardPrimaryAccountNumberOk

`func (o *PaymentMeans) GetPaymentCardPrimaryAccountNumberOk() (*string, bool)`

GetPaymentCardPrimaryAccountNumberOk returns a tuple with the PaymentCardPrimaryAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCardPrimaryAccountNumber

`func (o *PaymentMeans) SetPaymentCardPrimaryAccountNumber(v string)`

SetPaymentCardPrimaryAccountNumber sets PaymentCardPrimaryAccountNumber field to given value.

### HasPaymentCardPrimaryAccountNumber

`func (o *PaymentMeans) HasPaymentCardPrimaryAccountNumber() bool`

HasPaymentCardPrimaryAccountNumber returns a boolean if a field has been set.

### GetPaymentCardNetworkId

`func (o *PaymentMeans) GetPaymentCardNetworkId() string`

GetPaymentCardNetworkId returns the PaymentCardNetworkId field if non-nil, zero value otherwise.

### GetPaymentCardNetworkIdOk

`func (o *PaymentMeans) GetPaymentCardNetworkIdOk() (*string, bool)`

GetPaymentCardNetworkIdOk returns a tuple with the PaymentCardNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCardNetworkId

`func (o *PaymentMeans) SetPaymentCardNetworkId(v string)`

SetPaymentCardNetworkId sets PaymentCardNetworkId field to given value.

### HasPaymentCardNetworkId

`func (o *PaymentMeans) HasPaymentCardNetworkId() bool`

HasPaymentCardNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


