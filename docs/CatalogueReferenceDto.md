# CatalogueReferenceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**CatalogueReferenceKey** | Pointer to **string** | This points to the item in the catalogue itself. This means that, for example when this quantity assignment references a \&quot;DIN 276\&quot; catalogue, this property indicates the number / identifier / key in DIN 276 that is referenced. | [optional] 
**CatalogueReferenceId** | **string** | The Id of the CatalogueReference that is targeted by this item. Set this property to set the referenced catalogue. | 
**ProjectCatalogues** | Pointer to [**[]CatalogueDto**](CatalogueDto.md) | These are CatalogueReference that are used within this ServiceSpecification. Catalogue references are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. | [optional] 
**Catalogue** | Pointer to [**CatalogueDto**](CatalogueDto.md) |  | [optional] 

## Methods

### NewCatalogueReferenceDto

`func NewCatalogueReferenceDto(id string, catalogueReferenceId string, ) *CatalogueReferenceDto`

NewCatalogueReferenceDto instantiates a new CatalogueReferenceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueReferenceDtoWithDefaults

`func NewCatalogueReferenceDtoWithDefaults() *CatalogueReferenceDto`

NewCatalogueReferenceDtoWithDefaults instantiates a new CatalogueReferenceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogueReferenceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogueReferenceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogueReferenceDto) SetId(v string)`

SetId sets Id field to given value.


### GetCatalogueReferenceKey

`func (o *CatalogueReferenceDto) GetCatalogueReferenceKey() string`

GetCatalogueReferenceKey returns the CatalogueReferenceKey field if non-nil, zero value otherwise.

### GetCatalogueReferenceKeyOk

`func (o *CatalogueReferenceDto) GetCatalogueReferenceKeyOk() (*string, bool)`

GetCatalogueReferenceKeyOk returns a tuple with the CatalogueReferenceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueReferenceKey

`func (o *CatalogueReferenceDto) SetCatalogueReferenceKey(v string)`

SetCatalogueReferenceKey sets CatalogueReferenceKey field to given value.

### HasCatalogueReferenceKey

`func (o *CatalogueReferenceDto) HasCatalogueReferenceKey() bool`

HasCatalogueReferenceKey returns a boolean if a field has been set.

### GetCatalogueReferenceId

`func (o *CatalogueReferenceDto) GetCatalogueReferenceId() string`

GetCatalogueReferenceId returns the CatalogueReferenceId field if non-nil, zero value otherwise.

### GetCatalogueReferenceIdOk

`func (o *CatalogueReferenceDto) GetCatalogueReferenceIdOk() (*string, bool)`

GetCatalogueReferenceIdOk returns a tuple with the CatalogueReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueReferenceId

`func (o *CatalogueReferenceDto) SetCatalogueReferenceId(v string)`

SetCatalogueReferenceId sets CatalogueReferenceId field to given value.


### GetProjectCatalogues

`func (o *CatalogueReferenceDto) GetProjectCatalogues() []CatalogueDto`

GetProjectCatalogues returns the ProjectCatalogues field if non-nil, zero value otherwise.

### GetProjectCataloguesOk

`func (o *CatalogueReferenceDto) GetProjectCataloguesOk() (*[]CatalogueDto, bool)`

GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCatalogues

`func (o *CatalogueReferenceDto) SetProjectCatalogues(v []CatalogueDto)`

SetProjectCatalogues sets ProjectCatalogues field to given value.

### HasProjectCatalogues

`func (o *CatalogueReferenceDto) HasProjectCatalogues() bool`

HasProjectCatalogues returns a boolean if a field has been set.

### GetCatalogue

`func (o *CatalogueReferenceDto) GetCatalogue() CatalogueDto`

GetCatalogue returns the Catalogue field if non-nil, zero value otherwise.

### GetCatalogueOk

`func (o *CatalogueReferenceDto) GetCatalogueOk() (*CatalogueDto, bool)`

GetCatalogueOk returns a tuple with the Catalogue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogue

`func (o *CatalogueReferenceDto) SetCatalogue(v CatalogueDto)`

SetCatalogue sets Catalogue field to given value.

### HasCatalogue

`func (o *CatalogueReferenceDto) HasCatalogue() bool`

HasCatalogue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


