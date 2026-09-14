# LineItemAllowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAmount** | Pointer to **float32** | BT-136 The net amount of the allowance | [optional] 
**RelativeAllowanceBaseAmount** | Pointer to **float32** | BT-137 If this is a relative / percentage allowance, this should be set and give the base amount that is being used for the calculation | [optional] 
**RelativeAllowancePercentage** | Pointer to **float32** | BT-138 If this is a relative / percentage allowance, this should be set and be given in decimal representation, meaning 19% would be 0.19 | [optional] 
**Reason** | Pointer to **string** | BT-139 A description of the allowance. | [optional] 
**ReasonCode** | Pointer to **string** | BT-140 A reason code for the allowance according to UNTDID 5189. Should be a code from this list: https://unece.org/fileadmin/DAM/trade/untdid/d16b/tred/tred5189.htm | [optional] 
**ReasonType** | Pointer to [**AllowanceReasonType**](AllowanceReasonType.md) |  | [optional] 

## Methods

### NewLineItemAllowance

`func NewLineItemAllowance() *LineItemAllowance`

NewLineItemAllowance instantiates a new LineItemAllowance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemAllowanceWithDefaults

`func NewLineItemAllowanceWithDefaults() *LineItemAllowance`

NewLineItemAllowanceWithDefaults instantiates a new LineItemAllowance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetAmount

`func (o *LineItemAllowance) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *LineItemAllowance) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *LineItemAllowance) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *LineItemAllowance) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetRelativeAllowanceBaseAmount

`func (o *LineItemAllowance) GetRelativeAllowanceBaseAmount() float32`

GetRelativeAllowanceBaseAmount returns the RelativeAllowanceBaseAmount field if non-nil, zero value otherwise.

### GetRelativeAllowanceBaseAmountOk

`func (o *LineItemAllowance) GetRelativeAllowanceBaseAmountOk() (*float32, bool)`

GetRelativeAllowanceBaseAmountOk returns a tuple with the RelativeAllowanceBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeAllowanceBaseAmount

`func (o *LineItemAllowance) SetRelativeAllowanceBaseAmount(v float32)`

SetRelativeAllowanceBaseAmount sets RelativeAllowanceBaseAmount field to given value.

### HasRelativeAllowanceBaseAmount

`func (o *LineItemAllowance) HasRelativeAllowanceBaseAmount() bool`

HasRelativeAllowanceBaseAmount returns a boolean if a field has been set.

### GetRelativeAllowancePercentage

`func (o *LineItemAllowance) GetRelativeAllowancePercentage() float32`

GetRelativeAllowancePercentage returns the RelativeAllowancePercentage field if non-nil, zero value otherwise.

### GetRelativeAllowancePercentageOk

`func (o *LineItemAllowance) GetRelativeAllowancePercentageOk() (*float32, bool)`

GetRelativeAllowancePercentageOk returns a tuple with the RelativeAllowancePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeAllowancePercentage

`func (o *LineItemAllowance) SetRelativeAllowancePercentage(v float32)`

SetRelativeAllowancePercentage sets RelativeAllowancePercentage field to given value.

### HasRelativeAllowancePercentage

`func (o *LineItemAllowance) HasRelativeAllowancePercentage() bool`

HasRelativeAllowancePercentage returns a boolean if a field has been set.

### GetReason

`func (o *LineItemAllowance) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LineItemAllowance) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LineItemAllowance) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *LineItemAllowance) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *LineItemAllowance) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *LineItemAllowance) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *LineItemAllowance) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *LineItemAllowance) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReasonType

`func (o *LineItemAllowance) GetReasonType() AllowanceReasonType`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *LineItemAllowance) GetReasonTypeOk() (*AllowanceReasonType, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *LineItemAllowance) SetReasonType(v AllowanceReasonType)`

SetReasonType sets ReasonType field to given value.

### HasReasonType

`func (o *LineItemAllowance) HasReasonType() bool`

HasReasonType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


