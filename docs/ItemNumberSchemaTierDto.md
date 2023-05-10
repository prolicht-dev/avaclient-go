# ItemNumberSchemaTierDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | **int32** | The (maximum) length for this tier. Will not accept a length less than 1. Defaults to 1 if length less than one is specified. | 
**Type** | [**ItemNumberTypeDto**](ItemNumberTypeDto.md) |  | 
**TierType** | [**ItemNumberSchemaTierTypeDto**](ItemNumberSchemaTierTypeDto.md) |  | 
**IsLot** | **bool** | Indicates if this tier represents a lot. See the documentation for more information about lots. | [readonly] 
**Increment** | **int32** | This value is the increment, or step size, that should be used for new item numbers. It defaults to DEFAULT_INCREMENT, but can be changed to any other positive number greater than zero. Invalid values make this be set to one &#39;1&#39; | 
**TierName** | Pointer to **string** | This is an optional name for the given tier | [optional] 

## Methods

### NewItemNumberSchemaTierDto

`func NewItemNumberSchemaTierDto(length int32, type_ ItemNumberTypeDto, tierType ItemNumberSchemaTierTypeDto, isLot bool, increment int32, ) *ItemNumberSchemaTierDto`

NewItemNumberSchemaTierDto instantiates a new ItemNumberSchemaTierDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemNumberSchemaTierDtoWithDefaults

`func NewItemNumberSchemaTierDtoWithDefaults() *ItemNumberSchemaTierDto`

NewItemNumberSchemaTierDtoWithDefaults instantiates a new ItemNumberSchemaTierDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *ItemNumberSchemaTierDto) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ItemNumberSchemaTierDto) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ItemNumberSchemaTierDto) SetLength(v int32)`

SetLength sets Length field to given value.


### GetType

`func (o *ItemNumberSchemaTierDto) GetType() ItemNumberTypeDto`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemNumberSchemaTierDto) GetTypeOk() (*ItemNumberTypeDto, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemNumberSchemaTierDto) SetType(v ItemNumberTypeDto)`

SetType sets Type field to given value.


### GetTierType

`func (o *ItemNumberSchemaTierDto) GetTierType() ItemNumberSchemaTierTypeDto`

GetTierType returns the TierType field if non-nil, zero value otherwise.

### GetTierTypeOk

`func (o *ItemNumberSchemaTierDto) GetTierTypeOk() (*ItemNumberSchemaTierTypeDto, bool)`

GetTierTypeOk returns a tuple with the TierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierType

`func (o *ItemNumberSchemaTierDto) SetTierType(v ItemNumberSchemaTierTypeDto)`

SetTierType sets TierType field to given value.


### GetIsLot

`func (o *ItemNumberSchemaTierDto) GetIsLot() bool`

GetIsLot returns the IsLot field if non-nil, zero value otherwise.

### GetIsLotOk

`func (o *ItemNumberSchemaTierDto) GetIsLotOk() (*bool, bool)`

GetIsLotOk returns a tuple with the IsLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLot

`func (o *ItemNumberSchemaTierDto) SetIsLot(v bool)`

SetIsLot sets IsLot field to given value.


### GetIncrement

`func (o *ItemNumberSchemaTierDto) GetIncrement() int32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *ItemNumberSchemaTierDto) GetIncrementOk() (*int32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *ItemNumberSchemaTierDto) SetIncrement(v int32)`

SetIncrement sets Increment field to given value.


### GetTierName

`func (o *ItemNumberSchemaTierDto) GetTierName() string`

GetTierName returns the TierName field if non-nil, zero value otherwise.

### GetTierNameOk

`func (o *ItemNumberSchemaTierDto) GetTierNameOk() (*string, bool)`

GetTierNameOk returns a tuple with the TierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierName

`func (o *ItemNumberSchemaTierDto) SetTierName(v string)`

SetTierName sets TierName field to given value.

### HasTierName

`func (o *ItemNumberSchemaTierDto) HasTierName() bool`

HasTierName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


