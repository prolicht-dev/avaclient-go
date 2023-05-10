# ComplementedByQuantityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **float32** | The quantity that is complemented. E.g., if the base position has a quantity of 100 m² of a brick wall, a complementing position might be for &#39;walls that exceed 3 m height&#39; and for a total quantity of 10 m². | 
**ComplementingPositionId** | Pointer to **string** | This is the reference to the Id of the Position that is complementing. This means it does NOT reference the base position but the one that contains the actual addition. | [optional] 

## Methods

### NewComplementedByQuantityDto

`func NewComplementedByQuantityDto(quantity float32, ) *ComplementedByQuantityDto`

NewComplementedByQuantityDto instantiates a new ComplementedByQuantityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplementedByQuantityDtoWithDefaults

`func NewComplementedByQuantityDtoWithDefaults() *ComplementedByQuantityDto`

NewComplementedByQuantityDtoWithDefaults instantiates a new ComplementedByQuantityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *ComplementedByQuantityDto) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ComplementedByQuantityDto) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ComplementedByQuantityDto) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetComplementingPositionId

`func (o *ComplementedByQuantityDto) GetComplementingPositionId() string`

GetComplementingPositionId returns the ComplementingPositionId field if non-nil, zero value otherwise.

### GetComplementingPositionIdOk

`func (o *ComplementedByQuantityDto) GetComplementingPositionIdOk() (*string, bool)`

GetComplementingPositionIdOk returns a tuple with the ComplementingPositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementingPositionId

`func (o *ComplementedByQuantityDto) SetComplementingPositionId(v string)`

SetComplementingPositionId sets ComplementingPositionId field to given value.

### HasComplementingPositionId

`func (o *ComplementedByQuantityDto) HasComplementingPositionId() bool`

HasComplementingPositionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


