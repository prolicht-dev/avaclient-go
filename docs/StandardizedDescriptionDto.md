# StandardizedDescriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandardReferenceType** | [**StandardReferenceTypeDto**](StandardReferenceTypeDto.md) |  | 
**StandardReference** | Pointer to **string** | This string property is the identifier to map to the references standard. Its type is given in the StandardReferenceType | [optional] 
**StlbReference** | Pointer to [**STLBReferenceDto**](STLBReferenceDto.md) |  | [optional] 

## Methods

### NewStandardizedDescriptionDto

`func NewStandardizedDescriptionDto(standardReferenceType StandardReferenceTypeDto, ) *StandardizedDescriptionDto`

NewStandardizedDescriptionDto instantiates a new StandardizedDescriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardizedDescriptionDtoWithDefaults

`func NewStandardizedDescriptionDtoWithDefaults() *StandardizedDescriptionDto`

NewStandardizedDescriptionDtoWithDefaults instantiates a new StandardizedDescriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandardReferenceType

`func (o *StandardizedDescriptionDto) GetStandardReferenceType() StandardReferenceTypeDto`

GetStandardReferenceType returns the StandardReferenceType field if non-nil, zero value otherwise.

### GetStandardReferenceTypeOk

`func (o *StandardizedDescriptionDto) GetStandardReferenceTypeOk() (*StandardReferenceTypeDto, bool)`

GetStandardReferenceTypeOk returns a tuple with the StandardReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardReferenceType

`func (o *StandardizedDescriptionDto) SetStandardReferenceType(v StandardReferenceTypeDto)`

SetStandardReferenceType sets StandardReferenceType field to given value.


### GetStandardReference

`func (o *StandardizedDescriptionDto) GetStandardReference() string`

GetStandardReference returns the StandardReference field if non-nil, zero value otherwise.

### GetStandardReferenceOk

`func (o *StandardizedDescriptionDto) GetStandardReferenceOk() (*string, bool)`

GetStandardReferenceOk returns a tuple with the StandardReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardReference

`func (o *StandardizedDescriptionDto) SetStandardReference(v string)`

SetStandardReference sets StandardReference field to given value.

### HasStandardReference

`func (o *StandardizedDescriptionDto) HasStandardReference() bool`

HasStandardReference returns a boolean if a field has been set.

### GetStlbReference

`func (o *StandardizedDescriptionDto) GetStlbReference() STLBReferenceDto`

GetStlbReference returns the StlbReference field if non-nil, zero value otherwise.

### GetStlbReferenceOk

`func (o *StandardizedDescriptionDto) GetStlbReferenceOk() (*STLBReferenceDto, bool)`

GetStlbReferenceOk returns a tuple with the StlbReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStlbReference

`func (o *StandardizedDescriptionDto) SetStlbReference(v STLBReferenceDto)`

SetStlbReference sets StlbReference field to given value.

### HasStlbReference

`func (o *StandardizedDescriptionDto) HasStlbReference() bool`

HasStlbReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


