# PositionHoldOutPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**HoldOutType** | [**HoldOutTypeDto**](HoldOutTypeDto.md) |  | 
**HoldOutDuration** | **float32** | The duration of the period | 
**HoldOutDurationUnit** | [**HoldOutDurationUnitDto**](HoldOutDurationUnitDto.md) |  | 
**BasePositionReferences** | Pointer to **[]string** | References to base positions | [optional] 

## Methods

### NewPositionHoldOutPropertiesDto

`func NewPositionHoldOutPropertiesDto(id string, holdOutType HoldOutTypeDto, holdOutDuration float32, holdOutDurationUnit HoldOutDurationUnitDto, ) *PositionHoldOutPropertiesDto`

NewPositionHoldOutPropertiesDto instantiates a new PositionHoldOutPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionHoldOutPropertiesDtoWithDefaults

`func NewPositionHoldOutPropertiesDtoWithDefaults() *PositionHoldOutPropertiesDto`

NewPositionHoldOutPropertiesDtoWithDefaults instantiates a new PositionHoldOutPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PositionHoldOutPropertiesDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PositionHoldOutPropertiesDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PositionHoldOutPropertiesDto) SetId(v string)`

SetId sets Id field to given value.


### GetHoldOutType

`func (o *PositionHoldOutPropertiesDto) GetHoldOutType() HoldOutTypeDto`

GetHoldOutType returns the HoldOutType field if non-nil, zero value otherwise.

### GetHoldOutTypeOk

`func (o *PositionHoldOutPropertiesDto) GetHoldOutTypeOk() (*HoldOutTypeDto, bool)`

GetHoldOutTypeOk returns a tuple with the HoldOutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldOutType

`func (o *PositionHoldOutPropertiesDto) SetHoldOutType(v HoldOutTypeDto)`

SetHoldOutType sets HoldOutType field to given value.


### GetHoldOutDuration

`func (o *PositionHoldOutPropertiesDto) GetHoldOutDuration() float32`

GetHoldOutDuration returns the HoldOutDuration field if non-nil, zero value otherwise.

### GetHoldOutDurationOk

`func (o *PositionHoldOutPropertiesDto) GetHoldOutDurationOk() (*float32, bool)`

GetHoldOutDurationOk returns a tuple with the HoldOutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldOutDuration

`func (o *PositionHoldOutPropertiesDto) SetHoldOutDuration(v float32)`

SetHoldOutDuration sets HoldOutDuration field to given value.


### GetHoldOutDurationUnit

`func (o *PositionHoldOutPropertiesDto) GetHoldOutDurationUnit() HoldOutDurationUnitDto`

GetHoldOutDurationUnit returns the HoldOutDurationUnit field if non-nil, zero value otherwise.

### GetHoldOutDurationUnitOk

`func (o *PositionHoldOutPropertiesDto) GetHoldOutDurationUnitOk() (*HoldOutDurationUnitDto, bool)`

GetHoldOutDurationUnitOk returns a tuple with the HoldOutDurationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldOutDurationUnit

`func (o *PositionHoldOutPropertiesDto) SetHoldOutDurationUnit(v HoldOutDurationUnitDto)`

SetHoldOutDurationUnit sets HoldOutDurationUnit field to given value.


### GetBasePositionReferences

`func (o *PositionHoldOutPropertiesDto) GetBasePositionReferences() []string`

GetBasePositionReferences returns the BasePositionReferences field if non-nil, zero value otherwise.

### GetBasePositionReferencesOk

`func (o *PositionHoldOutPropertiesDto) GetBasePositionReferencesOk() (*[]string, bool)`

GetBasePositionReferencesOk returns a tuple with the BasePositionReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePositionReferences

`func (o *PositionHoldOutPropertiesDto) SetBasePositionReferences(v []string)`

SetBasePositionReferences sets BasePositionReferences field to given value.

### HasBasePositionReferences

`func (o *PositionHoldOutPropertiesDto) HasBasePositionReferences() bool`

HasBasePositionReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


