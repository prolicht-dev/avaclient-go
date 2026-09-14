# SustainabilityLifeCycleStageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Key** | Pointer to **string** | The key is the name or identifier of the stage. | [optional] 
**Value** | Pointer to **float32** | The value for this given stage. | [optional] 

## Methods

### NewSustainabilityLifeCycleStageDto

`func NewSustainabilityLifeCycleStageDto(id string, ) *SustainabilityLifeCycleStageDto`

NewSustainabilityLifeCycleStageDto instantiates a new SustainabilityLifeCycleStageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSustainabilityLifeCycleStageDtoWithDefaults

`func NewSustainabilityLifeCycleStageDtoWithDefaults() *SustainabilityLifeCycleStageDto`

NewSustainabilityLifeCycleStageDtoWithDefaults instantiates a new SustainabilityLifeCycleStageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SustainabilityLifeCycleStageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SustainabilityLifeCycleStageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SustainabilityLifeCycleStageDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *SustainabilityLifeCycleStageDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SustainabilityLifeCycleStageDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SustainabilityLifeCycleStageDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SustainabilityLifeCycleStageDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *SustainabilityLifeCycleStageDto) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SustainabilityLifeCycleStageDto) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SustainabilityLifeCycleStageDto) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *SustainabilityLifeCycleStageDto) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


