# PositionSustainabilityIndicatorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Source** | Pointer to **string** | This is usually in the form of an URL that specifies the indicator. | [optional] 
**Comment** | Pointer to **string** | This is a text giving more information about the indicator, e.g. details about the calculation method, or other comments. | [optional] 
**LifeCycleStages** | Pointer to [**[]SustainabilityLifeCycleStageDto**](SustainabilityLifeCycleStageDto.md) | If present, this indicator may be divided up into different life cycle stages. | [optional] 
**IndicatorKey** | Pointer to **string** | This is a key that specifies the type of indicator. It is not used for any calculations but can be used to identify the indicator type, e.g. \&quot;carbonFootprint\&quot;, \&quot;waterConsumption\&quot; or similar. The actual value of this property is not defined in the library and can be set freely, but it is recommended to use a consistent naming scheme, e.g. by using URLs or IRIs to avoid conflicts. In GAEB XML, it has a max length of 50 characters. | [optional] 
**Value** | Pointer to **float32** | This is the actual value of the indicator. | [optional] 
**UnitTag** | Pointer to **string** | The unit of measurement for the value. | [optional] 
**ConversionFactor** | Pointer to **float32** | If present, this is a conversion factor used here in this indicator for the value. | [optional] 

## Methods

### NewPositionSustainabilityIndicatorDto

`func NewPositionSustainabilityIndicatorDto(id string, ) *PositionSustainabilityIndicatorDto`

NewPositionSustainabilityIndicatorDto instantiates a new PositionSustainabilityIndicatorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionSustainabilityIndicatorDtoWithDefaults

`func NewPositionSustainabilityIndicatorDtoWithDefaults() *PositionSustainabilityIndicatorDto`

NewPositionSustainabilityIndicatorDtoWithDefaults instantiates a new PositionSustainabilityIndicatorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PositionSustainabilityIndicatorDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PositionSustainabilityIndicatorDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PositionSustainabilityIndicatorDto) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *PositionSustainabilityIndicatorDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PositionSustainabilityIndicatorDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PositionSustainabilityIndicatorDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PositionSustainabilityIndicatorDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetComment

`func (o *PositionSustainabilityIndicatorDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PositionSustainabilityIndicatorDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PositionSustainabilityIndicatorDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PositionSustainabilityIndicatorDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLifeCycleStages

`func (o *PositionSustainabilityIndicatorDto) GetLifeCycleStages() []SustainabilityLifeCycleStageDto`

GetLifeCycleStages returns the LifeCycleStages field if non-nil, zero value otherwise.

### GetLifeCycleStagesOk

`func (o *PositionSustainabilityIndicatorDto) GetLifeCycleStagesOk() (*[]SustainabilityLifeCycleStageDto, bool)`

GetLifeCycleStagesOk returns a tuple with the LifeCycleStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeCycleStages

`func (o *PositionSustainabilityIndicatorDto) SetLifeCycleStages(v []SustainabilityLifeCycleStageDto)`

SetLifeCycleStages sets LifeCycleStages field to given value.

### HasLifeCycleStages

`func (o *PositionSustainabilityIndicatorDto) HasLifeCycleStages() bool`

HasLifeCycleStages returns a boolean if a field has been set.

### GetIndicatorKey

`func (o *PositionSustainabilityIndicatorDto) GetIndicatorKey() string`

GetIndicatorKey returns the IndicatorKey field if non-nil, zero value otherwise.

### GetIndicatorKeyOk

`func (o *PositionSustainabilityIndicatorDto) GetIndicatorKeyOk() (*string, bool)`

GetIndicatorKeyOk returns a tuple with the IndicatorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatorKey

`func (o *PositionSustainabilityIndicatorDto) SetIndicatorKey(v string)`

SetIndicatorKey sets IndicatorKey field to given value.

### HasIndicatorKey

`func (o *PositionSustainabilityIndicatorDto) HasIndicatorKey() bool`

HasIndicatorKey returns a boolean if a field has been set.

### GetValue

`func (o *PositionSustainabilityIndicatorDto) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PositionSustainabilityIndicatorDto) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PositionSustainabilityIndicatorDto) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *PositionSustainabilityIndicatorDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnitTag

`func (o *PositionSustainabilityIndicatorDto) GetUnitTag() string`

GetUnitTag returns the UnitTag field if non-nil, zero value otherwise.

### GetUnitTagOk

`func (o *PositionSustainabilityIndicatorDto) GetUnitTagOk() (*string, bool)`

GetUnitTagOk returns a tuple with the UnitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitTag

`func (o *PositionSustainabilityIndicatorDto) SetUnitTag(v string)`

SetUnitTag sets UnitTag field to given value.

### HasUnitTag

`func (o *PositionSustainabilityIndicatorDto) HasUnitTag() bool`

HasUnitTag returns a boolean if a field has been set.

### GetConversionFactor

`func (o *PositionSustainabilityIndicatorDto) GetConversionFactor() float32`

GetConversionFactor returns the ConversionFactor field if non-nil, zero value otherwise.

### GetConversionFactorOk

`func (o *PositionSustainabilityIndicatorDto) GetConversionFactorOk() (*float32, bool)`

GetConversionFactorOk returns a tuple with the ConversionFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFactor

`func (o *PositionSustainabilityIndicatorDto) SetConversionFactor(v float32)`

SetConversionFactor sets ConversionFactor field to given value.

### HasConversionFactor

`func (o *PositionSustainabilityIndicatorDto) HasConversionFactor() bool`

HasConversionFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


