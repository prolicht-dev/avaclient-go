# InvoiceCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAmount** | Pointer to **float32** | BT-99 The net amount of the charge | [optional] 
**RelativeChargeBaseAmount** | Pointer to **float32** | BT-100 If this is a relative / percentage charge, this should be set and give the base amount that is being used for the calculation | [optional] 
**RelativeChargePercentage** | Pointer to **float32** | BT-101 If this is a relative / percentage charge, this should be set and be given in decimal representation, meaning 19% would be 0.19 | [optional] 
**VatCategory** | [**VatCategory**](VatCategory.md) |  | 
**VatRate** | Pointer to **float32** | BT-103 The actual tax rate as a decimal value, meaning 19% would be 0.19 | [optional] 
**Reason** | Pointer to **string** | BT-104 A description of the charge. This should not be set it ReasonCode is provided, otherwise this one takes precedence. | [optional] 
**ReasonCode** | Pointer to **string** | BT-105 A reason code for the charge according to UNTDID 7161. See this list: https://unece.org/fileadmin/DAM/trade/untdid/d16b/tred/tred7161.htm | [optional] 
**ReasonType** | Pointer to [**ChargeReasonType**](ChargeReasonType.md) |  | [optional] 

## Methods

### NewInvoiceCharge

`func NewInvoiceCharge(vatCategory VatCategory, ) *InvoiceCharge`

NewInvoiceCharge instantiates a new InvoiceCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceChargeWithDefaults

`func NewInvoiceChargeWithDefaults() *InvoiceCharge`

NewInvoiceChargeWithDefaults instantiates a new InvoiceCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetAmount

`func (o *InvoiceCharge) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *InvoiceCharge) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *InvoiceCharge) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *InvoiceCharge) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetRelativeChargeBaseAmount

`func (o *InvoiceCharge) GetRelativeChargeBaseAmount() float32`

GetRelativeChargeBaseAmount returns the RelativeChargeBaseAmount field if non-nil, zero value otherwise.

### GetRelativeChargeBaseAmountOk

`func (o *InvoiceCharge) GetRelativeChargeBaseAmountOk() (*float32, bool)`

GetRelativeChargeBaseAmountOk returns a tuple with the RelativeChargeBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeChargeBaseAmount

`func (o *InvoiceCharge) SetRelativeChargeBaseAmount(v float32)`

SetRelativeChargeBaseAmount sets RelativeChargeBaseAmount field to given value.

### HasRelativeChargeBaseAmount

`func (o *InvoiceCharge) HasRelativeChargeBaseAmount() bool`

HasRelativeChargeBaseAmount returns a boolean if a field has been set.

### GetRelativeChargePercentage

`func (o *InvoiceCharge) GetRelativeChargePercentage() float32`

GetRelativeChargePercentage returns the RelativeChargePercentage field if non-nil, zero value otherwise.

### GetRelativeChargePercentageOk

`func (o *InvoiceCharge) GetRelativeChargePercentageOk() (*float32, bool)`

GetRelativeChargePercentageOk returns a tuple with the RelativeChargePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeChargePercentage

`func (o *InvoiceCharge) SetRelativeChargePercentage(v float32)`

SetRelativeChargePercentage sets RelativeChargePercentage field to given value.

### HasRelativeChargePercentage

`func (o *InvoiceCharge) HasRelativeChargePercentage() bool`

HasRelativeChargePercentage returns a boolean if a field has been set.

### GetVatCategory

`func (o *InvoiceCharge) GetVatCategory() VatCategory`

GetVatCategory returns the VatCategory field if non-nil, zero value otherwise.

### GetVatCategoryOk

`func (o *InvoiceCharge) GetVatCategoryOk() (*VatCategory, bool)`

GetVatCategoryOk returns a tuple with the VatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCategory

`func (o *InvoiceCharge) SetVatCategory(v VatCategory)`

SetVatCategory sets VatCategory field to given value.


### GetVatRate

`func (o *InvoiceCharge) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *InvoiceCharge) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *InvoiceCharge) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *InvoiceCharge) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetReason

`func (o *InvoiceCharge) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoiceCharge) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoiceCharge) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InvoiceCharge) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *InvoiceCharge) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *InvoiceCharge) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *InvoiceCharge) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *InvoiceCharge) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReasonType

`func (o *InvoiceCharge) GetReasonType() ChargeReasonType`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *InvoiceCharge) GetReasonTypeOk() (*ChargeReasonType, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *InvoiceCharge) SetReasonType(v ChargeReasonType)`

SetReasonType sets ReasonType field to given value.

### HasReasonType

`func (o *InvoiceCharge) HasReasonType() bool`

HasReasonType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


