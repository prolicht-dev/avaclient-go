# ItemNumberSchemaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalLength** | **int32** | The count of tiers in the ItemNumberSchema | [readonly] 
**Tiers** | Pointer to [**[]ItemNumberSchemaTierDto**](ItemNumberSchemaTierDto.md) | The collection of tiers for this ItemNumberSchema. | [optional] 
**Separator** | Pointer to **string** | The separator to use for separiting the different levels in an ItemNumber. Defaults to DEFAULT_SEPARATOR, which is a point &#39;.&#39;. Setting this to a space or other whitespaces is discouraged, as this might not work correct in all situations and item numbers could be displayed not as intended. This can not be set to an empty or null string, trying that will default to the DEFAULT_SEPARATOR. If a value is set that has a different length than one &#39;1&#39;, the DEFAULT_SEPARATOR will be used instead. You should also not use values for the separator that are also valid for the item numbers themselves, as that might also lead to incorrect results | [optional] 
**Filler** | Pointer to **string** | This string is used to fill (left-pad) item numbers. For example, if a tier has a length of &#39;4&#39; but the given item number is &#39;12&#39;, with a Filler of &#39;0&#39;, then the final item number will be represented as &#39;0&#39;. This must be a single character string, if a value is given where the Length property does not evaluate to &#39;1&#39;, the DEFAULT_FILLER &#39;0&#39; is used. A space is fine to use. You should ensure that you use a value different than Separator, as that might produce unexpected results. No attempt is done by the code to recover from such ambiguous configurations. | [optional] 
**Identifier** | Pointer to **string** | This is just a string property that can optionally be used to store additional data for this ItemNumberSchema, e.g. an identification or a type. It does not have any influence over how item numbers are generated, and is not supported in most exchange formats. However, it is used to store ÖNorm service specification structure types. | [optional] 
**SkipNonExistingLevelsInPositionItemNumbers** | **bool** | This property indicates if ItemNumbers using this ItemNumberSchema should skip empty group levels. This is commonly only used in GAEB files, where there might be gaps in the hierarchy of elements and position identifiers should be placed at the end of the string representation. | 
**SkippedTiersFiller** | Pointer to **string** | This string is used only when the property SkipNonExistingLevelsInPositionItemNumbers in this ItemNumberSchema is also set to true. It defaults to DEFAULT_SKIPPED_TIERS_FILLER, but can be set to any string with a lenght of one. Null values or values with a longer length will lead to this property reverting back to the default value. This is used to fill skipped tiers in item numbers where a position is placed in a higher hierarchy level than what would be defined in the Tiers. For example, it could produce an item number like &#39;01.__.02&#39;, which would indicate a skipped second level. This should be using different values than Filler and Separator, since that could cause ambiguities in the code that generates the actual item numbers. No attempt is done by the code to recover from such ambiguous configurations. | [optional] 
**SchemaIsCorrectlyDefined** | **bool** | This is a read-only property that indicates if this schema has a valid structure. It internally just returns the result from IsCorrectlyDefined. This will return if the ItemNumberSchema is correctly defined. For it to be correctly defined, the following conditions must be true: There may only be one lot group, if there is one, it must be at the top. Following lot levels, there may be at least one group level. After the group levels, there must be one position level. After the position level, there may be one index level. If no tiers are defined at all, this will also return false. | [readonly] 
**AllowUpperCaseLettering** | **bool** | Defaults to false. If this is disabled, all letters in the ItemNumber string representations will be transformed to their lowercase representation. | 

## Methods

### NewItemNumberSchemaDto

`func NewItemNumberSchemaDto(totalLength int32, skipNonExistingLevelsInPositionItemNumbers bool, schemaIsCorrectlyDefined bool, allowUpperCaseLettering bool, ) *ItemNumberSchemaDto`

NewItemNumberSchemaDto instantiates a new ItemNumberSchemaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemNumberSchemaDtoWithDefaults

`func NewItemNumberSchemaDtoWithDefaults() *ItemNumberSchemaDto`

NewItemNumberSchemaDtoWithDefaults instantiates a new ItemNumberSchemaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalLength

`func (o *ItemNumberSchemaDto) GetTotalLength() int32`

GetTotalLength returns the TotalLength field if non-nil, zero value otherwise.

### GetTotalLengthOk

`func (o *ItemNumberSchemaDto) GetTotalLengthOk() (*int32, bool)`

GetTotalLengthOk returns a tuple with the TotalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLength

`func (o *ItemNumberSchemaDto) SetTotalLength(v int32)`

SetTotalLength sets TotalLength field to given value.


### GetTiers

`func (o *ItemNumberSchemaDto) GetTiers() []ItemNumberSchemaTierDto`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *ItemNumberSchemaDto) GetTiersOk() (*[]ItemNumberSchemaTierDto, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *ItemNumberSchemaDto) SetTiers(v []ItemNumberSchemaTierDto)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *ItemNumberSchemaDto) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetSeparator

`func (o *ItemNumberSchemaDto) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *ItemNumberSchemaDto) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *ItemNumberSchemaDto) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *ItemNumberSchemaDto) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetFiller

`func (o *ItemNumberSchemaDto) GetFiller() string`

GetFiller returns the Filler field if non-nil, zero value otherwise.

### GetFillerOk

`func (o *ItemNumberSchemaDto) GetFillerOk() (*string, bool)`

GetFillerOk returns a tuple with the Filler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiller

`func (o *ItemNumberSchemaDto) SetFiller(v string)`

SetFiller sets Filler field to given value.

### HasFiller

`func (o *ItemNumberSchemaDto) HasFiller() bool`

HasFiller returns a boolean if a field has been set.

### GetIdentifier

`func (o *ItemNumberSchemaDto) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ItemNumberSchemaDto) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ItemNumberSchemaDto) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ItemNumberSchemaDto) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSkipNonExistingLevelsInPositionItemNumbers

`func (o *ItemNumberSchemaDto) GetSkipNonExistingLevelsInPositionItemNumbers() bool`

GetSkipNonExistingLevelsInPositionItemNumbers returns the SkipNonExistingLevelsInPositionItemNumbers field if non-nil, zero value otherwise.

### GetSkipNonExistingLevelsInPositionItemNumbersOk

`func (o *ItemNumberSchemaDto) GetSkipNonExistingLevelsInPositionItemNumbersOk() (*bool, bool)`

GetSkipNonExistingLevelsInPositionItemNumbersOk returns a tuple with the SkipNonExistingLevelsInPositionItemNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipNonExistingLevelsInPositionItemNumbers

`func (o *ItemNumberSchemaDto) SetSkipNonExistingLevelsInPositionItemNumbers(v bool)`

SetSkipNonExistingLevelsInPositionItemNumbers sets SkipNonExistingLevelsInPositionItemNumbers field to given value.


### GetSkippedTiersFiller

`func (o *ItemNumberSchemaDto) GetSkippedTiersFiller() string`

GetSkippedTiersFiller returns the SkippedTiersFiller field if non-nil, zero value otherwise.

### GetSkippedTiersFillerOk

`func (o *ItemNumberSchemaDto) GetSkippedTiersFillerOk() (*string, bool)`

GetSkippedTiersFillerOk returns a tuple with the SkippedTiersFiller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedTiersFiller

`func (o *ItemNumberSchemaDto) SetSkippedTiersFiller(v string)`

SetSkippedTiersFiller sets SkippedTiersFiller field to given value.

### HasSkippedTiersFiller

`func (o *ItemNumberSchemaDto) HasSkippedTiersFiller() bool`

HasSkippedTiersFiller returns a boolean if a field has been set.

### GetSchemaIsCorrectlyDefined

`func (o *ItemNumberSchemaDto) GetSchemaIsCorrectlyDefined() bool`

GetSchemaIsCorrectlyDefined returns the SchemaIsCorrectlyDefined field if non-nil, zero value otherwise.

### GetSchemaIsCorrectlyDefinedOk

`func (o *ItemNumberSchemaDto) GetSchemaIsCorrectlyDefinedOk() (*bool, bool)`

GetSchemaIsCorrectlyDefinedOk returns a tuple with the SchemaIsCorrectlyDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaIsCorrectlyDefined

`func (o *ItemNumberSchemaDto) SetSchemaIsCorrectlyDefined(v bool)`

SetSchemaIsCorrectlyDefined sets SchemaIsCorrectlyDefined field to given value.


### GetAllowUpperCaseLettering

`func (o *ItemNumberSchemaDto) GetAllowUpperCaseLettering() bool`

GetAllowUpperCaseLettering returns the AllowUpperCaseLettering field if non-nil, zero value otherwise.

### GetAllowUpperCaseLetteringOk

`func (o *ItemNumberSchemaDto) GetAllowUpperCaseLetteringOk() (*bool, bool)`

GetAllowUpperCaseLetteringOk returns a tuple with the AllowUpperCaseLettering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpperCaseLettering

`func (o *ItemNumberSchemaDto) SetAllowUpperCaseLettering(v bool)`

SetAllowUpperCaseLettering sets AllowUpperCaseLettering field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


