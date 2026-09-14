# VatBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetTaxableAmount** | Pointer to **float32** | BT-116 The total net amount of this VAT category | [optional] 
**VatAmount** | Pointer to **float32** | BT-117 The total VAT | [optional] 
**VatCategory** | [**VatCategory**](VatCategory.md) |  | 
**TaxRate** | Pointer to **float32** | BT-119 The actual tax rate as a decimal value, e.g. 19% would be 0.19 | [optional] 
**VatExemptionReasonText** | Pointer to **string** | BT-120 If the VAT is exempt, there should be a description here. For example, in reverse charge scenarios, this text should contain a reference to \&quot;Reverse charge according to ...\&quot;. | [optional] 
**VatExemptionReasonCode** | Pointer to **string** | BT-121 A code from the VAT exemption reason code list(VATEX)“ | [optional] 

## Methods

### NewVatBreakdown

`func NewVatBreakdown(vatCategory VatCategory, ) *VatBreakdown`

NewVatBreakdown instantiates a new VatBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVatBreakdownWithDefaults

`func NewVatBreakdownWithDefaults() *VatBreakdown`

NewVatBreakdownWithDefaults instantiates a new VatBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetTaxableAmount

`func (o *VatBreakdown) GetNetTaxableAmount() float32`

GetNetTaxableAmount returns the NetTaxableAmount field if non-nil, zero value otherwise.

### GetNetTaxableAmountOk

`func (o *VatBreakdown) GetNetTaxableAmountOk() (*float32, bool)`

GetNetTaxableAmountOk returns a tuple with the NetTaxableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTaxableAmount

`func (o *VatBreakdown) SetNetTaxableAmount(v float32)`

SetNetTaxableAmount sets NetTaxableAmount field to given value.

### HasNetTaxableAmount

`func (o *VatBreakdown) HasNetTaxableAmount() bool`

HasNetTaxableAmount returns a boolean if a field has been set.

### GetVatAmount

`func (o *VatBreakdown) GetVatAmount() float32`

GetVatAmount returns the VatAmount field if non-nil, zero value otherwise.

### GetVatAmountOk

`func (o *VatBreakdown) GetVatAmountOk() (*float32, bool)`

GetVatAmountOk returns a tuple with the VatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAmount

`func (o *VatBreakdown) SetVatAmount(v float32)`

SetVatAmount sets VatAmount field to given value.

### HasVatAmount

`func (o *VatBreakdown) HasVatAmount() bool`

HasVatAmount returns a boolean if a field has been set.

### GetVatCategory

`func (o *VatBreakdown) GetVatCategory() VatCategory`

GetVatCategory returns the VatCategory field if non-nil, zero value otherwise.

### GetVatCategoryOk

`func (o *VatBreakdown) GetVatCategoryOk() (*VatCategory, bool)`

GetVatCategoryOk returns a tuple with the VatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCategory

`func (o *VatBreakdown) SetVatCategory(v VatCategory)`

SetVatCategory sets VatCategory field to given value.


### GetTaxRate

`func (o *VatBreakdown) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *VatBreakdown) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *VatBreakdown) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *VatBreakdown) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### GetVatExemptionReasonText

`func (o *VatBreakdown) GetVatExemptionReasonText() string`

GetVatExemptionReasonText returns the VatExemptionReasonText field if non-nil, zero value otherwise.

### GetVatExemptionReasonTextOk

`func (o *VatBreakdown) GetVatExemptionReasonTextOk() (*string, bool)`

GetVatExemptionReasonTextOk returns a tuple with the VatExemptionReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatExemptionReasonText

`func (o *VatBreakdown) SetVatExemptionReasonText(v string)`

SetVatExemptionReasonText sets VatExemptionReasonText field to given value.

### HasVatExemptionReasonText

`func (o *VatBreakdown) HasVatExemptionReasonText() bool`

HasVatExemptionReasonText returns a boolean if a field has been set.

### GetVatExemptionReasonCode

`func (o *VatBreakdown) GetVatExemptionReasonCode() string`

GetVatExemptionReasonCode returns the VatExemptionReasonCode field if non-nil, zero value otherwise.

### GetVatExemptionReasonCodeOk

`func (o *VatBreakdown) GetVatExemptionReasonCodeOk() (*string, bool)`

GetVatExemptionReasonCodeOk returns a tuple with the VatExemptionReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatExemptionReasonCode

`func (o *VatBreakdown) SetVatExemptionReasonCode(v string)`

SetVatExemptionReasonCode sets VatExemptionReasonCode field to given value.

### HasVatExemptionReasonCode

`func (o *VatBreakdown) HasVatExemptionReasonCode() bool`

HasVatExemptionReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


