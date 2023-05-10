# STLBReferenceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionDate** | Pointer to **time.Time** | The date of the STLB version. Typically, only the Year and Month are used | [optional] 
**CatalogueName** | Pointer to **string** | The name of the catalogue within the STLB | [optional] 
**Group** | Pointer to **string** | The name of the group in STLB | [optional] 
**CostGroup** | Pointer to **string** | The cost group this service is associated with | [optional] 
**ServiceArea** | Pointer to **string** | The service area (or type) in the STLB | [optional] 
**Keys** | Pointer to [**[]STLBKeyDto**](STLBKeyDto.md) | These keys may optionally be used to further reference multiple, specific items within the STLB | [optional] 

## Methods

### NewSTLBReferenceDto

`func NewSTLBReferenceDto() *STLBReferenceDto`

NewSTLBReferenceDto instantiates a new STLBReferenceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTLBReferenceDtoWithDefaults

`func NewSTLBReferenceDtoWithDefaults() *STLBReferenceDto`

NewSTLBReferenceDtoWithDefaults instantiates a new STLBReferenceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionDate

`func (o *STLBReferenceDto) GetVersionDate() time.Time`

GetVersionDate returns the VersionDate field if non-nil, zero value otherwise.

### GetVersionDateOk

`func (o *STLBReferenceDto) GetVersionDateOk() (*time.Time, bool)`

GetVersionDateOk returns a tuple with the VersionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionDate

`func (o *STLBReferenceDto) SetVersionDate(v time.Time)`

SetVersionDate sets VersionDate field to given value.

### HasVersionDate

`func (o *STLBReferenceDto) HasVersionDate() bool`

HasVersionDate returns a boolean if a field has been set.

### GetCatalogueName

`func (o *STLBReferenceDto) GetCatalogueName() string`

GetCatalogueName returns the CatalogueName field if non-nil, zero value otherwise.

### GetCatalogueNameOk

`func (o *STLBReferenceDto) GetCatalogueNameOk() (*string, bool)`

GetCatalogueNameOk returns a tuple with the CatalogueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueName

`func (o *STLBReferenceDto) SetCatalogueName(v string)`

SetCatalogueName sets CatalogueName field to given value.

### HasCatalogueName

`func (o *STLBReferenceDto) HasCatalogueName() bool`

HasCatalogueName returns a boolean if a field has been set.

### GetGroup

`func (o *STLBReferenceDto) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *STLBReferenceDto) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *STLBReferenceDto) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *STLBReferenceDto) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCostGroup

`func (o *STLBReferenceDto) GetCostGroup() string`

GetCostGroup returns the CostGroup field if non-nil, zero value otherwise.

### GetCostGroupOk

`func (o *STLBReferenceDto) GetCostGroupOk() (*string, bool)`

GetCostGroupOk returns a tuple with the CostGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostGroup

`func (o *STLBReferenceDto) SetCostGroup(v string)`

SetCostGroup sets CostGroup field to given value.

### HasCostGroup

`func (o *STLBReferenceDto) HasCostGroup() bool`

HasCostGroup returns a boolean if a field has been set.

### GetServiceArea

`func (o *STLBReferenceDto) GetServiceArea() string`

GetServiceArea returns the ServiceArea field if non-nil, zero value otherwise.

### GetServiceAreaOk

`func (o *STLBReferenceDto) GetServiceAreaOk() (*string, bool)`

GetServiceAreaOk returns a tuple with the ServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceArea

`func (o *STLBReferenceDto) SetServiceArea(v string)`

SetServiceArea sets ServiceArea field to given value.

### HasServiceArea

`func (o *STLBReferenceDto) HasServiceArea() bool`

HasServiceArea returns a boolean if a field has been set.

### GetKeys

`func (o *STLBReferenceDto) GetKeys() []STLBKeyDto`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *STLBReferenceDto) GetKeysOk() (*[]STLBKeyDto, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *STLBReferenceDto) SetKeys(v []STLBKeyDto)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *STLBReferenceDto) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


