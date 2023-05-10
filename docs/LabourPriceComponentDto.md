# LabourPriceComponentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The label associated with this price component. Will be taken from the parent Projects ProjectInformation. | [optional] 
**Price** | **float32** | The total, calculated price of this component. Will multiply the calculated amount of hours with the ServiceSpecifications hourly wage rate. | [readonly] 
**HourlyWage** | **float32** | The cost per hour of manual labor. | 
**Values** | Pointer to [**[]CalculationDto**](CalculationDto.md) | The single Calculation elements this price component is composed of. | [optional] 
**UseOwnHourlyWage** | **bool** | Indicates if the ServiceSpecification&#39;s standard HourlyWage is to be used or a custom value. | 
**TotalTime** | **float32** | The total, calculated time of this component. Will return the result rounded to three decimal places. | [readonly] 
**ProjectCatalogues** | Pointer to [**[]CatalogueDto**](CatalogueDto.md) | These are Catalogues that are used within this PriceComponent. Catalogues are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. They are propagate to all child elements, e.g. other containers and QuantityAssignments. In the context of a ServiceSpecification, all elements share the same instance of the collection. | [optional] 

## Methods

### NewLabourPriceComponentDto

`func NewLabourPriceComponentDto(price float32, hourlyWage float32, useOwnHourlyWage bool, totalTime float32, ) *LabourPriceComponentDto`

NewLabourPriceComponentDto instantiates a new LabourPriceComponentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabourPriceComponentDtoWithDefaults

`func NewLabourPriceComponentDtoWithDefaults() *LabourPriceComponentDto`

NewLabourPriceComponentDtoWithDefaults instantiates a new LabourPriceComponentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *LabourPriceComponentDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LabourPriceComponentDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LabourPriceComponentDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LabourPriceComponentDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPrice

`func (o *LabourPriceComponentDto) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *LabourPriceComponentDto) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *LabourPriceComponentDto) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetHourlyWage

`func (o *LabourPriceComponentDto) GetHourlyWage() float32`

GetHourlyWage returns the HourlyWage field if non-nil, zero value otherwise.

### GetHourlyWageOk

`func (o *LabourPriceComponentDto) GetHourlyWageOk() (*float32, bool)`

GetHourlyWageOk returns a tuple with the HourlyWage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyWage

`func (o *LabourPriceComponentDto) SetHourlyWage(v float32)`

SetHourlyWage sets HourlyWage field to given value.


### GetValues

`func (o *LabourPriceComponentDto) GetValues() []CalculationDto`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LabourPriceComponentDto) GetValuesOk() (*[]CalculationDto, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LabourPriceComponentDto) SetValues(v []CalculationDto)`

SetValues sets Values field to given value.

### HasValues

`func (o *LabourPriceComponentDto) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetUseOwnHourlyWage

`func (o *LabourPriceComponentDto) GetUseOwnHourlyWage() bool`

GetUseOwnHourlyWage returns the UseOwnHourlyWage field if non-nil, zero value otherwise.

### GetUseOwnHourlyWageOk

`func (o *LabourPriceComponentDto) GetUseOwnHourlyWageOk() (*bool, bool)`

GetUseOwnHourlyWageOk returns a tuple with the UseOwnHourlyWage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOwnHourlyWage

`func (o *LabourPriceComponentDto) SetUseOwnHourlyWage(v bool)`

SetUseOwnHourlyWage sets UseOwnHourlyWage field to given value.


### GetTotalTime

`func (o *LabourPriceComponentDto) GetTotalTime() float32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *LabourPriceComponentDto) GetTotalTimeOk() (*float32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *LabourPriceComponentDto) SetTotalTime(v float32)`

SetTotalTime sets TotalTime field to given value.


### GetProjectCatalogues

`func (o *LabourPriceComponentDto) GetProjectCatalogues() []CatalogueDto`

GetProjectCatalogues returns the ProjectCatalogues field if non-nil, zero value otherwise.

### GetProjectCataloguesOk

`func (o *LabourPriceComponentDto) GetProjectCataloguesOk() (*[]CatalogueDto, bool)`

GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCatalogues

`func (o *LabourPriceComponentDto) SetProjectCatalogues(v []CatalogueDto)`

SetProjectCatalogues sets ProjectCatalogues field to given value.

### HasProjectCatalogues

`func (o *LabourPriceComponentDto) HasProjectCatalogues() bool`

HasProjectCatalogues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


