# Payee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | BT-59 The name of the payee | [optional] 
**Identifier** | Pointer to **string** | BT-60 The identifier of the payee | [optional] 
**LegalRegistrationIdentifier** | Pointer to **string** | BT-61 A registered identifier for the payee, e.g. a VAT ID or Tax ID | [optional] 

## Methods

### NewPayee

`func NewPayee() *Payee`

NewPayee instantiates a new Payee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeWithDefaults

`func NewPayeeWithDefaults() *Payee`

NewPayeeWithDefaults instantiates a new Payee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Payee) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Payee) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Payee) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Payee) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdentifier

`func (o *Payee) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Payee) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Payee) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Payee) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetLegalRegistrationIdentifier

`func (o *Payee) GetLegalRegistrationIdentifier() string`

GetLegalRegistrationIdentifier returns the LegalRegistrationIdentifier field if non-nil, zero value otherwise.

### GetLegalRegistrationIdentifierOk

`func (o *Payee) GetLegalRegistrationIdentifierOk() (*string, bool)`

GetLegalRegistrationIdentifierOk returns a tuple with the LegalRegistrationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalRegistrationIdentifier

`func (o *Payee) SetLegalRegistrationIdentifier(v string)`

SetLegalRegistrationIdentifier sets LegalRegistrationIdentifier field to given value.

### HasLegalRegistrationIdentifier

`func (o *Payee) HasLegalRegistrationIdentifier() bool`

HasLegalRegistrationIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


