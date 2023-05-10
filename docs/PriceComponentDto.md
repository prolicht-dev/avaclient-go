# PriceComponentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **float32** | The total, calculated price of this component. | [readonly] 
**Label** | Pointer to **string** | The label associated with this price component. Will be taken from the parent Projects ProjectInformation. | [optional] 
**Values** | Pointer to [**[]CalculationDto**](CalculationDto.md) | The single Calculation elements this price component is composed of. | [optional] 
**ProjectCatalogues** | Pointer to [**[]CatalogueDto**](CatalogueDto.md) | These are Catalogues that are used within this PriceComponent. Catalogues are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. They are propagate to all child elements, e.g. other containers and QuantityAssignments. In the context of a ServiceSpecification, all elements share the same instance of the collection. | [optional] 

## Methods

### NewPriceComponentDto

`func NewPriceComponentDto(price float32, ) *PriceComponentDto`

NewPriceComponentDto instantiates a new PriceComponentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceComponentDtoWithDefaults

`func NewPriceComponentDtoWithDefaults() *PriceComponentDto`

NewPriceComponentDtoWithDefaults instantiates a new PriceComponentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *PriceComponentDto) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceComponentDto) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceComponentDto) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetLabel

`func (o *PriceComponentDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PriceComponentDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PriceComponentDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PriceComponentDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValues

`func (o *PriceComponentDto) GetValues() []CalculationDto`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PriceComponentDto) GetValuesOk() (*[]CalculationDto, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PriceComponentDto) SetValues(v []CalculationDto)`

SetValues sets Values field to given value.

### HasValues

`func (o *PriceComponentDto) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetProjectCatalogues

`func (o *PriceComponentDto) GetProjectCatalogues() []CatalogueDto`

GetProjectCatalogues returns the ProjectCatalogues field if non-nil, zero value otherwise.

### GetProjectCataloguesOk

`func (o *PriceComponentDto) GetProjectCataloguesOk() (*[]CatalogueDto, bool)`

GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCatalogues

`func (o *PriceComponentDto) SetProjectCatalogues(v []CatalogueDto)`

SetProjectCatalogues sets ProjectCatalogues field to given value.

### HasProjectCatalogues

`func (o *PriceComponentDto) HasProjectCatalogues() bool`

HasProjectCatalogues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


