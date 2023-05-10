# IElementDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**GaebXmlId** | Pointer to **string** | This is used to store the GAEB XML Id within this IElement. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization. | [optional] 
**AddendumNumber** | Pointer to **string** | This optional string property is shared by all IElements, and indicates if the element is part of an addendum, a &#39;Nachtrag&#39; in German. | [optional] 
**ProjectCatalogues** | Pointer to [**[]CatalogueDto**](CatalogueDto.md) |  | [optional] 
**CatalogueReferences** | Pointer to [**[]CatalogueReferenceDto**](CatalogueReferenceDto.md) |  | [optional] 
**ElementTypeDiscriminator** | **string** |  | 

## Methods

### NewIElementDto

`func NewIElementDto(id string, elementTypeDiscriminator string, ) *IElementDto`

NewIElementDto instantiates a new IElementDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIElementDtoWithDefaults

`func NewIElementDtoWithDefaults() *IElementDto`

NewIElementDtoWithDefaults instantiates a new IElementDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IElementDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IElementDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IElementDto) SetId(v string)`

SetId sets Id field to given value.


### GetGaebXmlId

`func (o *IElementDto) GetGaebXmlId() string`

GetGaebXmlId returns the GaebXmlId field if non-nil, zero value otherwise.

### GetGaebXmlIdOk

`func (o *IElementDto) GetGaebXmlIdOk() (*string, bool)`

GetGaebXmlIdOk returns a tuple with the GaebXmlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaebXmlId

`func (o *IElementDto) SetGaebXmlId(v string)`

SetGaebXmlId sets GaebXmlId field to given value.

### HasGaebXmlId

`func (o *IElementDto) HasGaebXmlId() bool`

HasGaebXmlId returns a boolean if a field has been set.

### GetAddendumNumber

`func (o *IElementDto) GetAddendumNumber() string`

GetAddendumNumber returns the AddendumNumber field if non-nil, zero value otherwise.

### GetAddendumNumberOk

`func (o *IElementDto) GetAddendumNumberOk() (*string, bool)`

GetAddendumNumberOk returns a tuple with the AddendumNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddendumNumber

`func (o *IElementDto) SetAddendumNumber(v string)`

SetAddendumNumber sets AddendumNumber field to given value.

### HasAddendumNumber

`func (o *IElementDto) HasAddendumNumber() bool`

HasAddendumNumber returns a boolean if a field has been set.

### GetProjectCatalogues

`func (o *IElementDto) GetProjectCatalogues() []CatalogueDto`

GetProjectCatalogues returns the ProjectCatalogues field if non-nil, zero value otherwise.

### GetProjectCataloguesOk

`func (o *IElementDto) GetProjectCataloguesOk() (*[]CatalogueDto, bool)`

GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCatalogues

`func (o *IElementDto) SetProjectCatalogues(v []CatalogueDto)`

SetProjectCatalogues sets ProjectCatalogues field to given value.

### HasProjectCatalogues

`func (o *IElementDto) HasProjectCatalogues() bool`

HasProjectCatalogues returns a boolean if a field has been set.

### GetCatalogueReferences

`func (o *IElementDto) GetCatalogueReferences() []CatalogueReferenceDto`

GetCatalogueReferences returns the CatalogueReferences field if non-nil, zero value otherwise.

### GetCatalogueReferencesOk

`func (o *IElementDto) GetCatalogueReferencesOk() (*[]CatalogueReferenceDto, bool)`

GetCatalogueReferencesOk returns a tuple with the CatalogueReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueReferences

`func (o *IElementDto) SetCatalogueReferences(v []CatalogueReferenceDto)`

SetCatalogueReferences sets CatalogueReferences field to given value.

### HasCatalogueReferences

`func (o *IElementDto) HasCatalogueReferences() bool`

HasCatalogueReferences returns a boolean if a field has been set.

### GetElementTypeDiscriminator

`func (o *IElementDto) GetElementTypeDiscriminator() string`

GetElementTypeDiscriminator returns the ElementTypeDiscriminator field if non-nil, zero value otherwise.

### GetElementTypeDiscriminatorOk

`func (o *IElementDto) GetElementTypeDiscriminatorOk() (*string, bool)`

GetElementTypeDiscriminatorOk returns a tuple with the ElementTypeDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementTypeDiscriminator

`func (o *IElementDto) SetElementTypeDiscriminator(v string)`

SetElementTypeDiscriminator sets ElementTypeDiscriminator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


