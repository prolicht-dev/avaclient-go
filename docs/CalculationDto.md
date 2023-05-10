# CalculationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Descriptive text for this calculation. | [optional] 
**Formula** | Pointer to **string** | This Calculation&#39;s mathematical expression. Please note that thousands separators are not supported. Both comma and point will be treated as decimal separators. | [optional] 
**Result** | **float64** | The calculated result from the formula, 0 if invalid. | [readonly] 
**Valid** | **bool** | Whether the Formula is a valid expression. | [readonly] 
**ErrorPositionInLine** | **int32** | Will be -1 if the Formula is correct, else it will show the position in the formula where an error was encountered. This is a zero based index. | [readonly] 
**ProjectCatalogues** | Pointer to [**[]CatalogueDto**](CatalogueDto.md) | These are Catalogues that are used within this Calculation. Catalogues are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. They are propagate to all child elements, e.g. other containers and QuantityAssignments. In the context of a ServiceSpecification, all elements share the same instance of the collection. | [optional] 
**CatalogueReferences** | Pointer to [**[]CatalogueReferenceDto**](CatalogueReferenceDto.md) | Referenced catalogues for this Calculation. | [optional] 

## Methods

### NewCalculationDto

`func NewCalculationDto(result float64, valid bool, errorPositionInLine int32, ) *CalculationDto`

NewCalculationDto instantiates a new CalculationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculationDtoWithDefaults

`func NewCalculationDtoWithDefaults() *CalculationDto`

NewCalculationDtoWithDefaults instantiates a new CalculationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CalculationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CalculationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CalculationDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CalculationDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormula

`func (o *CalculationDto) GetFormula() string`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *CalculationDto) GetFormulaOk() (*string, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *CalculationDto) SetFormula(v string)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *CalculationDto) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetResult

`func (o *CalculationDto) GetResult() float64`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CalculationDto) GetResultOk() (*float64, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CalculationDto) SetResult(v float64)`

SetResult sets Result field to given value.


### GetValid

`func (o *CalculationDto) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CalculationDto) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CalculationDto) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetErrorPositionInLine

`func (o *CalculationDto) GetErrorPositionInLine() int32`

GetErrorPositionInLine returns the ErrorPositionInLine field if non-nil, zero value otherwise.

### GetErrorPositionInLineOk

`func (o *CalculationDto) GetErrorPositionInLineOk() (*int32, bool)`

GetErrorPositionInLineOk returns a tuple with the ErrorPositionInLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPositionInLine

`func (o *CalculationDto) SetErrorPositionInLine(v int32)`

SetErrorPositionInLine sets ErrorPositionInLine field to given value.


### GetProjectCatalogues

`func (o *CalculationDto) GetProjectCatalogues() []CatalogueDto`

GetProjectCatalogues returns the ProjectCatalogues field if non-nil, zero value otherwise.

### GetProjectCataloguesOk

`func (o *CalculationDto) GetProjectCataloguesOk() (*[]CatalogueDto, bool)`

GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCatalogues

`func (o *CalculationDto) SetProjectCatalogues(v []CatalogueDto)`

SetProjectCatalogues sets ProjectCatalogues field to given value.

### HasProjectCatalogues

`func (o *CalculationDto) HasProjectCatalogues() bool`

HasProjectCatalogues returns a boolean if a field has been set.

### GetCatalogueReferences

`func (o *CalculationDto) GetCatalogueReferences() []CatalogueReferenceDto`

GetCatalogueReferences returns the CatalogueReferences field if non-nil, zero value otherwise.

### GetCatalogueReferencesOk

`func (o *CalculationDto) GetCatalogueReferencesOk() (*[]CatalogueReferenceDto, bool)`

GetCatalogueReferencesOk returns a tuple with the CatalogueReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueReferences

`func (o *CalculationDto) SetCatalogueReferences(v []CatalogueReferenceDto)`

SetCatalogueReferences sets CatalogueReferences field to given value.

### HasCatalogueReferences

`func (o *CalculationDto) HasCatalogueReferences() bool`

HasCatalogueReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


