# PaymentInstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | BT-82 Additional text to describe the payment instructions. | [optional] 
**PaymentTypeCode** | Pointer to **string** | BT-81 This should be a code according to UNCL 4461 to identify the type of payment. For example, code \&quot;58\&quot; means SEPA Credit Transfer, and code \&quot;30\&quot; means non-SEPA Credit Transfer. See this list: https://docs.peppol.eu/poacc/billing/3.0/codelist/UNCL4461/ | [optional] 
**PaymentType** | Pointer to [**PaymentType**](PaymentType.md) |  | [optional] 
**PaymentCreditorIdentifier** | Pointer to **string** | BT-90 If a SEPA direct debit payment mean is used, then this reference should be provided. It&#39;s the creditor identifier in SEPA, in German \&quot;Gläubiger-Identifikationsnummer\&quot;. Field BT-90 in XRechnung. | [optional] 
**PaymentMeans** | Pointer to [**[]PaymentMeans**](PaymentMeans.md) | BG-16 A list of payment means that can be used to pay the invoice. | [optional] 

## Methods

### NewPaymentInstructions

`func NewPaymentInstructions() *PaymentInstructions`

NewPaymentInstructions instantiates a new PaymentInstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstructionsWithDefaults

`func NewPaymentInstructionsWithDefaults() *PaymentInstructions`

NewPaymentInstructionsWithDefaults instantiates a new PaymentInstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PaymentInstructions) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentInstructions) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentInstructions) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentInstructions) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPaymentTypeCode

`func (o *PaymentInstructions) GetPaymentTypeCode() string`

GetPaymentTypeCode returns the PaymentTypeCode field if non-nil, zero value otherwise.

### GetPaymentTypeCodeOk

`func (o *PaymentInstructions) GetPaymentTypeCodeOk() (*string, bool)`

GetPaymentTypeCodeOk returns a tuple with the PaymentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypeCode

`func (o *PaymentInstructions) SetPaymentTypeCode(v string)`

SetPaymentTypeCode sets PaymentTypeCode field to given value.

### HasPaymentTypeCode

`func (o *PaymentInstructions) HasPaymentTypeCode() bool`

HasPaymentTypeCode returns a boolean if a field has been set.

### GetPaymentType

`func (o *PaymentInstructions) GetPaymentType() PaymentType`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PaymentInstructions) GetPaymentTypeOk() (*PaymentType, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PaymentInstructions) SetPaymentType(v PaymentType)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PaymentInstructions) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentCreditorIdentifier

`func (o *PaymentInstructions) GetPaymentCreditorIdentifier() string`

GetPaymentCreditorIdentifier returns the PaymentCreditorIdentifier field if non-nil, zero value otherwise.

### GetPaymentCreditorIdentifierOk

`func (o *PaymentInstructions) GetPaymentCreditorIdentifierOk() (*string, bool)`

GetPaymentCreditorIdentifierOk returns a tuple with the PaymentCreditorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCreditorIdentifier

`func (o *PaymentInstructions) SetPaymentCreditorIdentifier(v string)`

SetPaymentCreditorIdentifier sets PaymentCreditorIdentifier field to given value.

### HasPaymentCreditorIdentifier

`func (o *PaymentInstructions) HasPaymentCreditorIdentifier() bool`

HasPaymentCreditorIdentifier returns a boolean if a field has been set.

### GetPaymentMeans

`func (o *PaymentInstructions) GetPaymentMeans() []PaymentMeans`

GetPaymentMeans returns the PaymentMeans field if non-nil, zero value otherwise.

### GetPaymentMeansOk

`func (o *PaymentInstructions) GetPaymentMeansOk() (*[]PaymentMeans, bool)`

GetPaymentMeansOk returns a tuple with the PaymentMeans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMeans

`func (o *PaymentInstructions) SetPaymentMeans(v []PaymentMeans)`

SetPaymentMeans sets PaymentMeans field to given value.

### HasPaymentMeans

`func (o *PaymentInstructions) HasPaymentMeans() bool`

HasPaymentMeans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


