# InvoiceAllowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAmount** | Pointer to **float32** | BT-92 The net amount of the allowance | [optional] 
**RelativeAllowanceBaseAmount** | Pointer to **float32** | BT-93 If this is a relative / percentage allowance, this should be set and give the base amount that is being used for the calculation | [optional] 
**RelativeAllowancePercentage** | Pointer to **float32** | BT-94 If this is a relative / percentage allowance, this should be set and be given in decimal representation, meaning 19% would be 0.19 | [optional] 
**VatCategory** | [**VatCategory**](VatCategory.md) |  | 
**VatRate** | Pointer to **float32** | BT-96 The actual tax rate as a decimal value, meaning 19% would be 0.19 | [optional] 
**Reason** | Pointer to **string** | BT-97 A description of the allowance. This should not be set it ReasonCode is provided, otherwise this one takes precedence. | [optional] 
**ReasonCode** | Pointer to **string** | BT-98 A reason code for the allowance according to UNTDID 5189. See this list: https://unece.org/fileadmin/DAM/trade/untdid/d16b/tred/tred5189.htm | [optional] 
**ReasonType** | Pointer to [**AllowanceReasonType**](AllowanceReasonType.md) |  | [optional] 

## Methods

### NewInvoiceAllowance

`func NewInvoiceAllowance(vatCategory VatCategory, ) *InvoiceAllowance`

NewInvoiceAllowance instantiates a new InvoiceAllowance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAllowanceWithDefaults

`func NewInvoiceAllowanceWithDefaults() *InvoiceAllowance`

NewInvoiceAllowanceWithDefaults instantiates a new InvoiceAllowance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetAmount

`func (o *InvoiceAllowance) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *InvoiceAllowance) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *InvoiceAllowance) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *InvoiceAllowance) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetRelativeAllowanceBaseAmount

`func (o *InvoiceAllowance) GetRelativeAllowanceBaseAmount() float32`

GetRelativeAllowanceBaseAmount returns the RelativeAllowanceBaseAmount field if non-nil, zero value otherwise.

### GetRelativeAllowanceBaseAmountOk

`func (o *InvoiceAllowance) GetRelativeAllowanceBaseAmountOk() (*float32, bool)`

GetRelativeAllowanceBaseAmountOk returns a tuple with the RelativeAllowanceBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeAllowanceBaseAmount

`func (o *InvoiceAllowance) SetRelativeAllowanceBaseAmount(v float32)`

SetRelativeAllowanceBaseAmount sets RelativeAllowanceBaseAmount field to given value.

### HasRelativeAllowanceBaseAmount

`func (o *InvoiceAllowance) HasRelativeAllowanceBaseAmount() bool`

HasRelativeAllowanceBaseAmount returns a boolean if a field has been set.

### GetRelativeAllowancePercentage

`func (o *InvoiceAllowance) GetRelativeAllowancePercentage() float32`

GetRelativeAllowancePercentage returns the RelativeAllowancePercentage field if non-nil, zero value otherwise.

### GetRelativeAllowancePercentageOk

`func (o *InvoiceAllowance) GetRelativeAllowancePercentageOk() (*float32, bool)`

GetRelativeAllowancePercentageOk returns a tuple with the RelativeAllowancePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeAllowancePercentage

`func (o *InvoiceAllowance) SetRelativeAllowancePercentage(v float32)`

SetRelativeAllowancePercentage sets RelativeAllowancePercentage field to given value.

### HasRelativeAllowancePercentage

`func (o *InvoiceAllowance) HasRelativeAllowancePercentage() bool`

HasRelativeAllowancePercentage returns a boolean if a field has been set.

### GetVatCategory

`func (o *InvoiceAllowance) GetVatCategory() VatCategory`

GetVatCategory returns the VatCategory field if non-nil, zero value otherwise.

### GetVatCategoryOk

`func (o *InvoiceAllowance) GetVatCategoryOk() (*VatCategory, bool)`

GetVatCategoryOk returns a tuple with the VatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCategory

`func (o *InvoiceAllowance) SetVatCategory(v VatCategory)`

SetVatCategory sets VatCategory field to given value.


### GetVatRate

`func (o *InvoiceAllowance) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *InvoiceAllowance) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *InvoiceAllowance) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *InvoiceAllowance) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetReason

`func (o *InvoiceAllowance) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoiceAllowance) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoiceAllowance) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InvoiceAllowance) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *InvoiceAllowance) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *InvoiceAllowance) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *InvoiceAllowance) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *InvoiceAllowance) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReasonType

`func (o *InvoiceAllowance) GetReasonType() AllowanceReasonType`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *InvoiceAllowance) GetReasonTypeOk() (*AllowanceReasonType, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *InvoiceAllowance) SetReasonType(v AllowanceReasonType)`

SetReasonType sets ReasonType field to given value.

### HasReasonType

`func (o *InvoiceAllowance) HasReasonType() bool`

HasReasonType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


