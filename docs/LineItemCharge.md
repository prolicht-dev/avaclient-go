# LineItemCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAmount** | Pointer to **float32** | BT-141 The net amount of the charge | [optional] 
**RelativeChargeBaseAmount** | Pointer to **float32** | BT-142 If this is a relative / percentage charge, this should be set and give the base amount that is being used for the calculation | [optional] 
**RelativeChargePercentage** | Pointer to **float32** | BT-143 If this is a relative / percentage charge, this should be set and be given in decimal representation, meaning 19% would be 0.19 | [optional] 
**Reason** | Pointer to **string** | BT-144 A description of the charge. | [optional] 
**ReasonCode** | Pointer to **string** | BT-145 A reason code for the charge according to UNTDID 7161. See this list: https://unece.org/fileadmin/DAM/trade/untdid/d16b/tred/tred7161.htm | [optional] 
**ReasonType** | Pointer to [**ChargeReasonType**](ChargeReasonType.md) |  | [optional] 

## Methods

### NewLineItemCharge

`func NewLineItemCharge() *LineItemCharge`

NewLineItemCharge instantiates a new LineItemCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemChargeWithDefaults

`func NewLineItemChargeWithDefaults() *LineItemCharge`

NewLineItemChargeWithDefaults instantiates a new LineItemCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetAmount

`func (o *LineItemCharge) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *LineItemCharge) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *LineItemCharge) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *LineItemCharge) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetRelativeChargeBaseAmount

`func (o *LineItemCharge) GetRelativeChargeBaseAmount() float32`

GetRelativeChargeBaseAmount returns the RelativeChargeBaseAmount field if non-nil, zero value otherwise.

### GetRelativeChargeBaseAmountOk

`func (o *LineItemCharge) GetRelativeChargeBaseAmountOk() (*float32, bool)`

GetRelativeChargeBaseAmountOk returns a tuple with the RelativeChargeBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeChargeBaseAmount

`func (o *LineItemCharge) SetRelativeChargeBaseAmount(v float32)`

SetRelativeChargeBaseAmount sets RelativeChargeBaseAmount field to given value.

### HasRelativeChargeBaseAmount

`func (o *LineItemCharge) HasRelativeChargeBaseAmount() bool`

HasRelativeChargeBaseAmount returns a boolean if a field has been set.

### GetRelativeChargePercentage

`func (o *LineItemCharge) GetRelativeChargePercentage() float32`

GetRelativeChargePercentage returns the RelativeChargePercentage field if non-nil, zero value otherwise.

### GetRelativeChargePercentageOk

`func (o *LineItemCharge) GetRelativeChargePercentageOk() (*float32, bool)`

GetRelativeChargePercentageOk returns a tuple with the RelativeChargePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeChargePercentage

`func (o *LineItemCharge) SetRelativeChargePercentage(v float32)`

SetRelativeChargePercentage sets RelativeChargePercentage field to given value.

### HasRelativeChargePercentage

`func (o *LineItemCharge) HasRelativeChargePercentage() bool`

HasRelativeChargePercentage returns a boolean if a field has been set.

### GetReason

`func (o *LineItemCharge) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LineItemCharge) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LineItemCharge) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *LineItemCharge) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *LineItemCharge) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *LineItemCharge) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *LineItemCharge) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *LineItemCharge) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetReasonType

`func (o *LineItemCharge) GetReasonType() ChargeReasonType`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *LineItemCharge) GetReasonTypeOk() (*ChargeReasonType, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *LineItemCharge) SetReasonType(v ChargeReasonType)`

SetReasonType sets ReasonType field to given value.

### HasReasonType

`func (o *LineItemCharge) HasReasonType() bool`

HasReasonType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


