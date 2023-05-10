# BankingInformationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Name** | Pointer to **string** | The name of the bank. | [optional] 
**Iban** | Pointer to **string** | The international identifier for the bank account. | [optional] 
**AccountNumber** | Pointer to **string** | The account number. Typically no longer used since the introduction of IBAN within the SEPA area. | [optional] 
**Bic** | Pointer to **string** | The international bank identifier. | [optional] 
**RoutingNumber** | Pointer to **string** | The routing number for the bank. Typically no longer used since the introduction of IBAN within the SEPA area. | [optional] 

## Methods

### NewBankingInformationDto

`func NewBankingInformationDto(id string, ) *BankingInformationDto`

NewBankingInformationDto instantiates a new BankingInformationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankingInformationDtoWithDefaults

`func NewBankingInformationDtoWithDefaults() *BankingInformationDto`

NewBankingInformationDtoWithDefaults instantiates a new BankingInformationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankingInformationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankingInformationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankingInformationDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BankingInformationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankingInformationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankingInformationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BankingInformationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIban

`func (o *BankingInformationDto) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *BankingInformationDto) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *BankingInformationDto) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *BankingInformationDto) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetAccountNumber

`func (o *BankingInformationDto) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankingInformationDto) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankingInformationDto) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *BankingInformationDto) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetBic

`func (o *BankingInformationDto) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *BankingInformationDto) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *BankingInformationDto) SetBic(v string)`

SetBic sets Bic field to given value.

### HasBic

`func (o *BankingInformationDto) HasBic() bool`

HasBic returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *BankingInformationDto) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankingInformationDto) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankingInformationDto) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *BankingInformationDto) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


