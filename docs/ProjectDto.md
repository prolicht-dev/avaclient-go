# ProjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**PriceAccuracy** | **int32** | This property controls the accuracy of all price properties, meaning how many decimal places are preserved in calculations. It defaults to DEFAULT_PRICE_ACCURACY. Please see the Dangl.AVA documentation for further information about decimal precision. If UnitPriceAccuracy is set, then this property is ignored for unit prices. | 
**UnitPriceAccuracy** | Pointer to **int32** | This property controls the accuracy of position unit price properties, meaning how many decimal places are preserved in calculations. Please see the Dangl.AVA documentation for further information about decimal precision. This can be separately set from PriceAccuracy, and it only controls the accuracy of the unit price of positions, not the total price. It defaults to null, which means the standard PriceAccuracy is used for unit prices. | [optional] 
**ForceStrictTotals** | **bool** | This forces total prices to be the strict product of quantities times unit price in positions. It is enabled by default. If this is disabled, both the unit price and the total price of positions is calculated from the non-rounded values. Please see the documentation for a more detailed explanation of this setting. | 
**PriceRoundingMode** | [**PriceRoundingModeDto**](PriceRoundingModeDto.md) |  | 
**ProjectInformation** | Pointer to [**ProjectInformationDto**](ProjectInformationDto.md) |  | [optional] 
**ServiceSpecifications** | Pointer to [**[]ServiceSpecificationDto**](ServiceSpecificationDto.md) | The ServiceSpecifications in this Project. | [optional] 
**GaebXmlId** | Pointer to **string** | This is used to store the GAEB XML Id within this Project. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization. | [optional] 

## Methods

### NewProjectDto

`func NewProjectDto(id string, priceAccuracy int32, forceStrictTotals bool, priceRoundingMode PriceRoundingModeDto, ) *ProjectDto`

NewProjectDto instantiates a new ProjectDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDtoWithDefaults

`func NewProjectDtoWithDefaults() *ProjectDto`

NewProjectDtoWithDefaults instantiates a new ProjectDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDto) SetId(v string)`

SetId sets Id field to given value.


### GetPriceAccuracy

`func (o *ProjectDto) GetPriceAccuracy() int32`

GetPriceAccuracy returns the PriceAccuracy field if non-nil, zero value otherwise.

### GetPriceAccuracyOk

`func (o *ProjectDto) GetPriceAccuracyOk() (*int32, bool)`

GetPriceAccuracyOk returns a tuple with the PriceAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAccuracy

`func (o *ProjectDto) SetPriceAccuracy(v int32)`

SetPriceAccuracy sets PriceAccuracy field to given value.


### GetUnitPriceAccuracy

`func (o *ProjectDto) GetUnitPriceAccuracy() int32`

GetUnitPriceAccuracy returns the UnitPriceAccuracy field if non-nil, zero value otherwise.

### GetUnitPriceAccuracyOk

`func (o *ProjectDto) GetUnitPriceAccuracyOk() (*int32, bool)`

GetUnitPriceAccuracyOk returns a tuple with the UnitPriceAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceAccuracy

`func (o *ProjectDto) SetUnitPriceAccuracy(v int32)`

SetUnitPriceAccuracy sets UnitPriceAccuracy field to given value.

### HasUnitPriceAccuracy

`func (o *ProjectDto) HasUnitPriceAccuracy() bool`

HasUnitPriceAccuracy returns a boolean if a field has been set.

### GetForceStrictTotals

`func (o *ProjectDto) GetForceStrictTotals() bool`

GetForceStrictTotals returns the ForceStrictTotals field if non-nil, zero value otherwise.

### GetForceStrictTotalsOk

`func (o *ProjectDto) GetForceStrictTotalsOk() (*bool, bool)`

GetForceStrictTotalsOk returns a tuple with the ForceStrictTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceStrictTotals

`func (o *ProjectDto) SetForceStrictTotals(v bool)`

SetForceStrictTotals sets ForceStrictTotals field to given value.


### GetPriceRoundingMode

`func (o *ProjectDto) GetPriceRoundingMode() PriceRoundingModeDto`

GetPriceRoundingMode returns the PriceRoundingMode field if non-nil, zero value otherwise.

### GetPriceRoundingModeOk

`func (o *ProjectDto) GetPriceRoundingModeOk() (*PriceRoundingModeDto, bool)`

GetPriceRoundingModeOk returns a tuple with the PriceRoundingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRoundingMode

`func (o *ProjectDto) SetPriceRoundingMode(v PriceRoundingModeDto)`

SetPriceRoundingMode sets PriceRoundingMode field to given value.


### GetProjectInformation

`func (o *ProjectDto) GetProjectInformation() ProjectInformationDto`

GetProjectInformation returns the ProjectInformation field if non-nil, zero value otherwise.

### GetProjectInformationOk

`func (o *ProjectDto) GetProjectInformationOk() (*ProjectInformationDto, bool)`

GetProjectInformationOk returns a tuple with the ProjectInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectInformation

`func (o *ProjectDto) SetProjectInformation(v ProjectInformationDto)`

SetProjectInformation sets ProjectInformation field to given value.

### HasProjectInformation

`func (o *ProjectDto) HasProjectInformation() bool`

HasProjectInformation returns a boolean if a field has been set.

### GetServiceSpecifications

`func (o *ProjectDto) GetServiceSpecifications() []ServiceSpecificationDto`

GetServiceSpecifications returns the ServiceSpecifications field if non-nil, zero value otherwise.

### GetServiceSpecificationsOk

`func (o *ProjectDto) GetServiceSpecificationsOk() (*[]ServiceSpecificationDto, bool)`

GetServiceSpecificationsOk returns a tuple with the ServiceSpecifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSpecifications

`func (o *ProjectDto) SetServiceSpecifications(v []ServiceSpecificationDto)`

SetServiceSpecifications sets ServiceSpecifications field to given value.

### HasServiceSpecifications

`func (o *ProjectDto) HasServiceSpecifications() bool`

HasServiceSpecifications returns a boolean if a field has been set.

### GetGaebXmlId

`func (o *ProjectDto) GetGaebXmlId() string`

GetGaebXmlId returns the GaebXmlId field if non-nil, zero value otherwise.

### GetGaebXmlIdOk

`func (o *ProjectDto) GetGaebXmlIdOk() (*string, bool)`

GetGaebXmlIdOk returns a tuple with the GaebXmlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaebXmlId

`func (o *ProjectDto) SetGaebXmlId(v string)`

SetGaebXmlId sets GaebXmlId field to given value.

### HasGaebXmlId

`func (o *ProjectDto) HasGaebXmlId() bool`

HasGaebXmlId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


