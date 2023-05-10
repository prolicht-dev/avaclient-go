# OenormPositionPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginCode** | [**OenormOriginCodeDto**](OenormOriginCodeDto.md) |  | 
**OpeningTextIsFreeText** | **bool** | This marks if the opening texts within this element are considered free text. It corresponds to &#39;vorbemerkungskennzeichen&#39; in ÖNorm. | 
**IsMainPosition** | **bool** | This indicates if the ÖNorm &#39;wesentliche position&#39; mark is set | 
**IsUndividedPosition** | **bool** | This indicates if the ÖNorm position was a &#39;ungeteilteposition&#39; (undivided position). This will only be taken into account when the position is also the sole element inside it&#39;s parent group | 
**OenormShortText** | Pointer to **string** | In some ÖNorm formats, the short text can have it&#39;s own addition, so the text is split up in OenormShortText and OenormShortTextAddition To serialize this, either the ShortText property of the parent position needs to be null, or OenormShortText &#39; &#39; OenormShortTextAddition needs to match the ShortText. | [optional] 
**OenormShortTextAddition** | Pointer to **string** | In some ÖNorm formats, the short text can have it&#39;s own addition, so the text is split up in OenormShortText and OenormShortTextAddition To serialize this, either the ShortText property of the parent position needs to be null, or OenormShortText &#39; &#39; OenormShortTextAddition needs to match the ShortText. | [optional] 

## Methods

### NewOenormPositionPropertiesDto

`func NewOenormPositionPropertiesDto(originCode OenormOriginCodeDto, openingTextIsFreeText bool, isMainPosition bool, isUndividedPosition bool, ) *OenormPositionPropertiesDto`

NewOenormPositionPropertiesDto instantiates a new OenormPositionPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOenormPositionPropertiesDtoWithDefaults

`func NewOenormPositionPropertiesDtoWithDefaults() *OenormPositionPropertiesDto`

NewOenormPositionPropertiesDtoWithDefaults instantiates a new OenormPositionPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginCode

`func (o *OenormPositionPropertiesDto) GetOriginCode() OenormOriginCodeDto`

GetOriginCode returns the OriginCode field if non-nil, zero value otherwise.

### GetOriginCodeOk

`func (o *OenormPositionPropertiesDto) GetOriginCodeOk() (*OenormOriginCodeDto, bool)`

GetOriginCodeOk returns a tuple with the OriginCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCode

`func (o *OenormPositionPropertiesDto) SetOriginCode(v OenormOriginCodeDto)`

SetOriginCode sets OriginCode field to given value.


### GetOpeningTextIsFreeText

`func (o *OenormPositionPropertiesDto) GetOpeningTextIsFreeText() bool`

GetOpeningTextIsFreeText returns the OpeningTextIsFreeText field if non-nil, zero value otherwise.

### GetOpeningTextIsFreeTextOk

`func (o *OenormPositionPropertiesDto) GetOpeningTextIsFreeTextOk() (*bool, bool)`

GetOpeningTextIsFreeTextOk returns a tuple with the OpeningTextIsFreeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningTextIsFreeText

`func (o *OenormPositionPropertiesDto) SetOpeningTextIsFreeText(v bool)`

SetOpeningTextIsFreeText sets OpeningTextIsFreeText field to given value.


### GetIsMainPosition

`func (o *OenormPositionPropertiesDto) GetIsMainPosition() bool`

GetIsMainPosition returns the IsMainPosition field if non-nil, zero value otherwise.

### GetIsMainPositionOk

`func (o *OenormPositionPropertiesDto) GetIsMainPositionOk() (*bool, bool)`

GetIsMainPositionOk returns a tuple with the IsMainPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMainPosition

`func (o *OenormPositionPropertiesDto) SetIsMainPosition(v bool)`

SetIsMainPosition sets IsMainPosition field to given value.


### GetIsUndividedPosition

`func (o *OenormPositionPropertiesDto) GetIsUndividedPosition() bool`

GetIsUndividedPosition returns the IsUndividedPosition field if non-nil, zero value otherwise.

### GetIsUndividedPositionOk

`func (o *OenormPositionPropertiesDto) GetIsUndividedPositionOk() (*bool, bool)`

GetIsUndividedPositionOk returns a tuple with the IsUndividedPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUndividedPosition

`func (o *OenormPositionPropertiesDto) SetIsUndividedPosition(v bool)`

SetIsUndividedPosition sets IsUndividedPosition field to given value.


### GetOenormShortText

`func (o *OenormPositionPropertiesDto) GetOenormShortText() string`

GetOenormShortText returns the OenormShortText field if non-nil, zero value otherwise.

### GetOenormShortTextOk

`func (o *OenormPositionPropertiesDto) GetOenormShortTextOk() (*string, bool)`

GetOenormShortTextOk returns a tuple with the OenormShortText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOenormShortText

`func (o *OenormPositionPropertiesDto) SetOenormShortText(v string)`

SetOenormShortText sets OenormShortText field to given value.

### HasOenormShortText

`func (o *OenormPositionPropertiesDto) HasOenormShortText() bool`

HasOenormShortText returns a boolean if a field has been set.

### GetOenormShortTextAddition

`func (o *OenormPositionPropertiesDto) GetOenormShortTextAddition() string`

GetOenormShortTextAddition returns the OenormShortTextAddition field if non-nil, zero value otherwise.

### GetOenormShortTextAdditionOk

`func (o *OenormPositionPropertiesDto) GetOenormShortTextAdditionOk() (*string, bool)`

GetOenormShortTextAdditionOk returns a tuple with the OenormShortTextAddition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOenormShortTextAddition

`func (o *OenormPositionPropertiesDto) SetOenormShortTextAddition(v string)`

SetOenormShortTextAddition sets OenormShortTextAddition field to given value.

### HasOenormShortTextAddition

`func (o *OenormPositionPropertiesDto) HasOenormShortTextAddition() bool`

HasOenormShortTextAddition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


