# CatalogueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**GaebXmlId** | Pointer to **string** | This is used to store the GAEB XML Id within this Catalogue. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization. | [optional] 
**Name** | Pointer to **string** | The name that is given for this catalogue. | [optional] 
**Description** | Pointer to **string** | Additional information about this catalogue. | [optional] 
**CatalogueType** | [**CatalogueTypeDto**](CatalogueTypeDto.md) |  | 
**CatalogueTypeDetail** | Pointer to **string** | This property may hold additional information about the catalogue type. It is currently only used in GAEB XML exchange to carry detailed information about a catalogue type, but it&#39;s otherwise just a free text field. | [optional] 

## Methods

### NewCatalogueDto

`func NewCatalogueDto(id string, catalogueType CatalogueTypeDto, ) *CatalogueDto`

NewCatalogueDto instantiates a new CatalogueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogueDtoWithDefaults

`func NewCatalogueDtoWithDefaults() *CatalogueDto`

NewCatalogueDtoWithDefaults instantiates a new CatalogueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogueDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogueDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogueDto) SetId(v string)`

SetId sets Id field to given value.


### GetGaebXmlId

`func (o *CatalogueDto) GetGaebXmlId() string`

GetGaebXmlId returns the GaebXmlId field if non-nil, zero value otherwise.

### GetGaebXmlIdOk

`func (o *CatalogueDto) GetGaebXmlIdOk() (*string, bool)`

GetGaebXmlIdOk returns a tuple with the GaebXmlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaebXmlId

`func (o *CatalogueDto) SetGaebXmlId(v string)`

SetGaebXmlId sets GaebXmlId field to given value.

### HasGaebXmlId

`func (o *CatalogueDto) HasGaebXmlId() bool`

HasGaebXmlId returns a boolean if a field has been set.

### GetName

`func (o *CatalogueDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogueDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogueDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogueDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogueDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogueDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogueDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogueDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCatalogueType

`func (o *CatalogueDto) GetCatalogueType() CatalogueTypeDto`

GetCatalogueType returns the CatalogueType field if non-nil, zero value otherwise.

### GetCatalogueTypeOk

`func (o *CatalogueDto) GetCatalogueTypeOk() (*CatalogueTypeDto, bool)`

GetCatalogueTypeOk returns a tuple with the CatalogueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueType

`func (o *CatalogueDto) SetCatalogueType(v CatalogueTypeDto)`

SetCatalogueType sets CatalogueType field to given value.


### GetCatalogueTypeDetail

`func (o *CatalogueDto) GetCatalogueTypeDetail() string`

GetCatalogueTypeDetail returns the CatalogueTypeDetail field if non-nil, zero value otherwise.

### GetCatalogueTypeDetailOk

`func (o *CatalogueDto) GetCatalogueTypeDetailOk() (*string, bool)`

GetCatalogueTypeDetailOk returns a tuple with the CatalogueTypeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueTypeDetail

`func (o *CatalogueDto) SetCatalogueTypeDetail(v string)`

SetCatalogueTypeDetail sets CatalogueTypeDetail field to given value.

### HasCatalogueTypeDetail

`func (o *CatalogueDto) HasCatalogueTypeDetail() bool`

HasCatalogueTypeDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


