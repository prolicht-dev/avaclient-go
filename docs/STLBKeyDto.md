# STLBKeyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtIdentifier** | **int32** | This identifier is required and uniquely describes a single reference within the STLB standard. It maps to \&quot;ArtChrIdent\&quot; in GAEB XML | 
**ArtIndex** | Pointer to **int32** | This optional index property further categorizes a single reference within the STLB standard. It maps to \&quot;ArtChIdx\&quot; in GAEB XML | [optional] 
**KindIdentifier** | Pointer to **int32** | This optional identifier further specifies the execution kind of the reference in the STLB standard. It maps to \&quot;ChVIdent\&quot; in GAEB XML | [optional] 

## Methods

### NewSTLBKeyDto

`func NewSTLBKeyDto(artIdentifier int32, ) *STLBKeyDto`

NewSTLBKeyDto instantiates a new STLBKeyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTLBKeyDtoWithDefaults

`func NewSTLBKeyDtoWithDefaults() *STLBKeyDto`

NewSTLBKeyDtoWithDefaults instantiates a new STLBKeyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtIdentifier

`func (o *STLBKeyDto) GetArtIdentifier() int32`

GetArtIdentifier returns the ArtIdentifier field if non-nil, zero value otherwise.

### GetArtIdentifierOk

`func (o *STLBKeyDto) GetArtIdentifierOk() (*int32, bool)`

GetArtIdentifierOk returns a tuple with the ArtIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtIdentifier

`func (o *STLBKeyDto) SetArtIdentifier(v int32)`

SetArtIdentifier sets ArtIdentifier field to given value.


### GetArtIndex

`func (o *STLBKeyDto) GetArtIndex() int32`

GetArtIndex returns the ArtIndex field if non-nil, zero value otherwise.

### GetArtIndexOk

`func (o *STLBKeyDto) GetArtIndexOk() (*int32, bool)`

GetArtIndexOk returns a tuple with the ArtIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtIndex

`func (o *STLBKeyDto) SetArtIndex(v int32)`

SetArtIndex sets ArtIndex field to given value.

### HasArtIndex

`func (o *STLBKeyDto) HasArtIndex() bool`

HasArtIndex returns a boolean if a field has been set.

### GetKindIdentifier

`func (o *STLBKeyDto) GetKindIdentifier() int32`

GetKindIdentifier returns the KindIdentifier field if non-nil, zero value otherwise.

### GetKindIdentifierOk

`func (o *STLBKeyDto) GetKindIdentifierOk() (*int32, bool)`

GetKindIdentifierOk returns a tuple with the KindIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKindIdentifier

`func (o *STLBKeyDto) SetKindIdentifier(v int32)`

SetKindIdentifier sets KindIdentifier field to given value.

### HasKindIdentifier

`func (o *STLBKeyDto) HasKindIdentifier() bool`

HasKindIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


