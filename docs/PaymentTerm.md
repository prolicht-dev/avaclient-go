# PaymentTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** | The number of days until the discount is valid. | 
**Percent** | **float32** | The percentage of the discount. | 
**BaseAmount** | Pointer to **float32** | The base amount for the discount, if applicable. | [optional] 

## Methods

### NewPaymentTerm

`func NewPaymentTerm(days int32, percent float32, ) *PaymentTerm`

NewPaymentTerm instantiates a new PaymentTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTermWithDefaults

`func NewPaymentTermWithDefaults() *PaymentTerm`

NewPaymentTermWithDefaults instantiates a new PaymentTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *PaymentTerm) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PaymentTerm) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PaymentTerm) SetDays(v int32)`

SetDays sets Days field to given value.


### GetPercent

`func (o *PaymentTerm) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *PaymentTerm) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *PaymentTerm) SetPercent(v float32)`

SetPercent sets Percent field to given value.


### GetBaseAmount

`func (o *PaymentTerm) GetBaseAmount() float32`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *PaymentTerm) GetBaseAmountOk() (*float32, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *PaymentTerm) SetBaseAmount(v float32)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *PaymentTerm) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


