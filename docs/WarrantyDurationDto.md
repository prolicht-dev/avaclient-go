# WarrantyDurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The scalar value of the duration. This value must be equal to or bigger than zero (&gt;&#x3D; 0). Negative values can not be set and will be ignored. | 
**Unit** | [**DurationUnitDto**](DurationUnitDto.md) |  | 

## Methods

### NewWarrantyDurationDto

`func NewWarrantyDurationDto(duration int32, unit DurationUnitDto, ) *WarrantyDurationDto`

NewWarrantyDurationDto instantiates a new WarrantyDurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarrantyDurationDtoWithDefaults

`func NewWarrantyDurationDtoWithDefaults() *WarrantyDurationDto`

NewWarrantyDurationDtoWithDefaults instantiates a new WarrantyDurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *WarrantyDurationDto) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WarrantyDurationDto) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WarrantyDurationDto) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetUnit

`func (o *WarrantyDurationDto) GetUnit() DurationUnitDto`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *WarrantyDurationDto) GetUnitOk() (*DurationUnitDto, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *WarrantyDurationDto) SetUnit(v DurationUnitDto)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


