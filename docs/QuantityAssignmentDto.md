# QuantityAssignmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Quantity** | **float32** | The total quantity in this quantity assignment | 
**ProjectCatalogues** | Pointer to [**[]CatalogueDto**](CatalogueDto.md) | These are Catalogue that are used within this ServiceSpecification. Catalogue references are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. | [optional] 
**CatalogueReferences** | Pointer to [**[]CatalogueReferenceDto**](CatalogueReferenceDto.md) | Referenced catalogues for this QuantityAssignment. | [optional] 

## Methods

### NewQuantityAssignmentDto

`func NewQuantityAssignmentDto(id string, quantity float32, ) *QuantityAssignmentDto`

NewQuantityAssignmentDto instantiates a new QuantityAssignmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuantityAssignmentDtoWithDefaults

`func NewQuantityAssignmentDtoWithDefaults() *QuantityAssignmentDto`

NewQuantityAssignmentDtoWithDefaults instantiates a new QuantityAssignmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuantityAssignmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuantityAssignmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuantityAssignmentDto) SetId(v string)`

SetId sets Id field to given value.


### GetQuantity

`func (o *QuantityAssignmentDto) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuantityAssignmentDto) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuantityAssignmentDto) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetProjectCatalogues

`func (o *QuantityAssignmentDto) GetProjectCatalogues() []CatalogueDto`

GetProjectCatalogues returns the ProjectCatalogues field if non-nil, zero value otherwise.

### GetProjectCataloguesOk

`func (o *QuantityAssignmentDto) GetProjectCataloguesOk() (*[]CatalogueDto, bool)`

GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCatalogues

`func (o *QuantityAssignmentDto) SetProjectCatalogues(v []CatalogueDto)`

SetProjectCatalogues sets ProjectCatalogues field to given value.

### HasProjectCatalogues

`func (o *QuantityAssignmentDto) HasProjectCatalogues() bool`

HasProjectCatalogues returns a boolean if a field has been set.

### GetCatalogueReferences

`func (o *QuantityAssignmentDto) GetCatalogueReferences() []CatalogueReferenceDto`

GetCatalogueReferences returns the CatalogueReferences field if non-nil, zero value otherwise.

### GetCatalogueReferencesOk

`func (o *QuantityAssignmentDto) GetCatalogueReferencesOk() (*[]CatalogueReferenceDto, bool)`

GetCatalogueReferencesOk returns a tuple with the CatalogueReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueReferences

`func (o *QuantityAssignmentDto) SetCatalogueReferences(v []CatalogueReferenceDto)`

SetCatalogueReferences sets CatalogueReferences field to given value.

### HasCatalogueReferences

`func (o *QuantityAssignmentDto) HasCatalogueReferences() bool`

HasCatalogueReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


