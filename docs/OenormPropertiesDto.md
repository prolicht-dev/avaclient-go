# OenormPropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginCode** | [**OenormOriginCodeDto**](OenormOriginCodeDto.md) |  | 
**OpeningTextIsFreeText** | **bool** | This marks if the opening texts within this element are considered free text. It corresponds to &#39;vorbemerkungskennzeichen&#39; in ÖNorm. | 

## Methods

### NewOenormPropertiesDto

`func NewOenormPropertiesDto(originCode OenormOriginCodeDto, openingTextIsFreeText bool, ) *OenormPropertiesDto`

NewOenormPropertiesDto instantiates a new OenormPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOenormPropertiesDtoWithDefaults

`func NewOenormPropertiesDtoWithDefaults() *OenormPropertiesDto`

NewOenormPropertiesDtoWithDefaults instantiates a new OenormPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginCode

`func (o *OenormPropertiesDto) GetOriginCode() OenormOriginCodeDto`

GetOriginCode returns the OriginCode field if non-nil, zero value otherwise.

### GetOriginCodeOk

`func (o *OenormPropertiesDto) GetOriginCodeOk() (*OenormOriginCodeDto, bool)`

GetOriginCodeOk returns a tuple with the OriginCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCode

`func (o *OenormPropertiesDto) SetOriginCode(v OenormOriginCodeDto)`

SetOriginCode sets OriginCode field to given value.


### GetOpeningTextIsFreeText

`func (o *OenormPropertiesDto) GetOpeningTextIsFreeText() bool`

GetOpeningTextIsFreeText returns the OpeningTextIsFreeText field if non-nil, zero value otherwise.

### GetOpeningTextIsFreeTextOk

`func (o *OenormPropertiesDto) GetOpeningTextIsFreeTextOk() (*bool, bool)`

GetOpeningTextIsFreeTextOk returns a tuple with the OpeningTextIsFreeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningTextIsFreeText

`func (o *OenormPropertiesDto) SetOpeningTextIsFreeText(v bool)`

SetOpeningTextIsFreeText sets OpeningTextIsFreeText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


