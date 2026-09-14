# InvoicePeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **time.Time** | BT-134 BT-73 The start date | [optional] 
**End** | Pointer to **time.Time** | BT-135 BT-74 The end date | [optional] 

## Methods

### NewInvoicePeriod

`func NewInvoicePeriod() *InvoicePeriod`

NewInvoicePeriod instantiates a new InvoicePeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePeriodWithDefaults

`func NewInvoicePeriodWithDefaults() *InvoicePeriod`

NewInvoicePeriodWithDefaults instantiates a new InvoicePeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *InvoicePeriod) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *InvoicePeriod) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *InvoicePeriod) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *InvoicePeriod) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *InvoicePeriod) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *InvoicePeriod) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *InvoicePeriod) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *InvoicePeriod) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


