# FlatElementDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviousElementId** | Pointer to **string** | If this is not null, then this contains the id of the previous element in the hierarchy on the same level. | [optional] 
**ParentElementId** | Pointer to **string** | If this is not null, then this contains the id of the parent element. | [optional] 
**Element** | Pointer to [**IElementDto**](IElementDto.md) |  | [optional] 

## Methods

### NewFlatElementDto

`func NewFlatElementDto() *FlatElementDto`

NewFlatElementDto instantiates a new FlatElementDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatElementDtoWithDefaults

`func NewFlatElementDtoWithDefaults() *FlatElementDto`

NewFlatElementDtoWithDefaults instantiates a new FlatElementDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviousElementId

`func (o *FlatElementDto) GetPreviousElementId() string`

GetPreviousElementId returns the PreviousElementId field if non-nil, zero value otherwise.

### GetPreviousElementIdOk

`func (o *FlatElementDto) GetPreviousElementIdOk() (*string, bool)`

GetPreviousElementIdOk returns a tuple with the PreviousElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousElementId

`func (o *FlatElementDto) SetPreviousElementId(v string)`

SetPreviousElementId sets PreviousElementId field to given value.

### HasPreviousElementId

`func (o *FlatElementDto) HasPreviousElementId() bool`

HasPreviousElementId returns a boolean if a field has been set.

### GetParentElementId

`func (o *FlatElementDto) GetParentElementId() string`

GetParentElementId returns the ParentElementId field if non-nil, zero value otherwise.

### GetParentElementIdOk

`func (o *FlatElementDto) GetParentElementIdOk() (*string, bool)`

GetParentElementIdOk returns a tuple with the ParentElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentElementId

`func (o *FlatElementDto) SetParentElementId(v string)`

SetParentElementId sets ParentElementId field to given value.

### HasParentElementId

`func (o *FlatElementDto) HasParentElementId() bool`

HasParentElementId returns a boolean if a field has been set.

### GetElement

`func (o *FlatElementDto) GetElement() IElementDto`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *FlatElementDto) GetElementOk() (*IElementDto, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *FlatElementDto) SetElement(v IElementDto)`

SetElement sets Element field to given value.

### HasElement

`func (o *FlatElementDto) HasElement() bool`

HasElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


