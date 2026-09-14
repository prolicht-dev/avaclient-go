# LineItemAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | BT-160 The key / name of the attribute | 
**Value** | Pointer to **string** | BT-161 The value of the attribute | [optional] 

## Methods

### NewLineItemAttribute

`func NewLineItemAttribute(name string, ) *LineItemAttribute`

NewLineItemAttribute instantiates a new LineItemAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemAttributeWithDefaults

`func NewLineItemAttributeWithDefaults() *LineItemAttribute`

NewLineItemAttributeWithDefaults instantiates a new LineItemAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LineItemAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LineItemAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LineItemAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *LineItemAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LineItemAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LineItemAttribute) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *LineItemAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


