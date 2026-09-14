# OenormNoteTextPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginCode** | [**OenormOriginCodeDto**](OenormOriginCodeDto.md) |  | 
**OpeningTextIsFreeText** | **bool** | This marks if the opening texts within this element are considered free text. It corresponds to &#39;vorbemerkungskennzeichen&#39; in ÖNorm. | 
**IsUndividedPosition** | **bool** | This indicates if the ÖNorm position / note text was a &#39;ungeteilteposition&#39; (undivided position). This will only be taken into account when the position / note text is also the sole element inside it&#39;s parent group | 

## Methods

### NewOenormNoteTextPropertiesDto

`func NewOenormNoteTextPropertiesDto(originCode OenormOriginCodeDto, openingTextIsFreeText bool, isUndividedPosition bool, ) *OenormNoteTextPropertiesDto`

NewOenormNoteTextPropertiesDto instantiates a new OenormNoteTextPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOenormNoteTextPropertiesDtoWithDefaults

`func NewOenormNoteTextPropertiesDtoWithDefaults() *OenormNoteTextPropertiesDto`

NewOenormNoteTextPropertiesDtoWithDefaults instantiates a new OenormNoteTextPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginCode

`func (o *OenormNoteTextPropertiesDto) GetOriginCode() OenormOriginCodeDto`

GetOriginCode returns the OriginCode field if non-nil, zero value otherwise.

### GetOriginCodeOk

`func (o *OenormNoteTextPropertiesDto) GetOriginCodeOk() (*OenormOriginCodeDto, bool)`

GetOriginCodeOk returns a tuple with the OriginCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCode

`func (o *OenormNoteTextPropertiesDto) SetOriginCode(v OenormOriginCodeDto)`

SetOriginCode sets OriginCode field to given value.


### GetOpeningTextIsFreeText

`func (o *OenormNoteTextPropertiesDto) GetOpeningTextIsFreeText() bool`

GetOpeningTextIsFreeText returns the OpeningTextIsFreeText field if non-nil, zero value otherwise.

### GetOpeningTextIsFreeTextOk

`func (o *OenormNoteTextPropertiesDto) GetOpeningTextIsFreeTextOk() (*bool, bool)`

GetOpeningTextIsFreeTextOk returns a tuple with the OpeningTextIsFreeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningTextIsFreeText

`func (o *OenormNoteTextPropertiesDto) SetOpeningTextIsFreeText(v bool)`

SetOpeningTextIsFreeText sets OpeningTextIsFreeText field to given value.


### GetIsUndividedPosition

`func (o *OenormNoteTextPropertiesDto) GetIsUndividedPosition() bool`

GetIsUndividedPosition returns the IsUndividedPosition field if non-nil, zero value otherwise.

### GetIsUndividedPositionOk

`func (o *OenormNoteTextPropertiesDto) GetIsUndividedPositionOk() (*bool, bool)`

GetIsUndividedPositionOk returns a tuple with the IsUndividedPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUndividedPosition

`func (o *OenormNoteTextPropertiesDto) SetIsUndividedPosition(v bool)`

SetIsUndividedPosition sets IsUndividedPosition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


