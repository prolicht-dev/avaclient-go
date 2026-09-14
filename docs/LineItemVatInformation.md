# LineItemVatInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VatCategory** | [**VatCategory**](VatCategory.md) |  | 
**VatRate** | Pointer to **float32** | BT-152 The actual tax rate as a decimal value, meaning 19% would be 0.19 | [optional] 

## Methods

### NewLineItemVatInformation

`func NewLineItemVatInformation(vatCategory VatCategory, ) *LineItemVatInformation`

NewLineItemVatInformation instantiates a new LineItemVatInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemVatInformationWithDefaults

`func NewLineItemVatInformationWithDefaults() *LineItemVatInformation`

NewLineItemVatInformationWithDefaults instantiates a new LineItemVatInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVatCategory

`func (o *LineItemVatInformation) GetVatCategory() VatCategory`

GetVatCategory returns the VatCategory field if non-nil, zero value otherwise.

### GetVatCategoryOk

`func (o *LineItemVatInformation) GetVatCategoryOk() (*VatCategory, bool)`

GetVatCategoryOk returns a tuple with the VatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCategory

`func (o *LineItemVatInformation) SetVatCategory(v VatCategory)`

SetVatCategory sets VatCategory field to given value.


### GetVatRate

`func (o *LineItemVatInformation) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *LineItemVatInformation) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *LineItemVatInformation) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *LineItemVatInformation) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


