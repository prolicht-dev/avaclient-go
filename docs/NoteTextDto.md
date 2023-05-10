# NoteTextDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOpeningText** | **bool** | If this is set to true, this text is meant to not be seen as part of the regular elements hierarchy but as a special opening text at the beginning of the project. For example, in GAEB XML, this would map to the GAEB.Award.AddText. Typically, such texts describe project-wide contractual definitions. If this is set to true, this NoteText should be placed at the top of the elements hierarchy directly in the ServiceSpecification.Elements group, otherwise it will likely not be treated correctly when exporting to GAEB. You can only set IsOpeningText or IsClosingText to true. | 
**IsClosingText** | **bool** | If this is set to true, this text is meant to not be seen as part of the regular elements hierarchy but as a special closing text at the end of the project. For Example, in GAEB XML, this would map to the GAEB.AddText. Typically, such texts are used to describe project wide finishing descriptions. If this is set to true, this NoteText should be placed at the top of the elements hierarchy directly in the ServiceSpecification.Elements group, otherwise it will likely not be treated correctly when exporting to GAEB. You can only set IsOpeningText or IsClosingText to true. | 
**ShortText** | Pointer to **string** | Short description for this DescriptionBase element. | [optional] 
**AdditionType** | [**AdditionTypeDto**](AdditionTypeDto.md) |  | 
**LongText** | Pointer to **string** | Detailed description for this DescriptionBase element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText. | [optional] 
**HtmlLongText** | Pointer to **string** | This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out. | [optional] 
**Identifier** | Pointer to **string** | This is an optional internal identifier that may be used to add additional information to this NoteText. It is not supported in GAEB import or export. | [optional] 
**StandardizedDescription** | Pointer to [**StandardizedDescriptionDto**](StandardizedDescriptionDto.md) |  | [optional] 
**ElementType** | Pointer to **string** |  | [optional] 
**DescriptionId** | Pointer to **string** | This is an identifier specific for this description. Some exchange formats, like GAEB XML, use it to identify descriptions. It&#39;s different to an elements identifier in that it should only apply to the description component, meaning the text itself. | [optional] 
**OenormNoteTextProperties** | Pointer to [**OenormNoteTextPropertiesDto**](OenormNoteTextPropertiesDto.md) |  | [optional] 
**HasBidderCommentInHtmlLongText** | **bool** |  | [readonly] 

## Methods

### NewNoteTextDto

`func NewNoteTextDto(isOpeningText bool, isClosingText bool, additionType AdditionTypeDto, hasBidderCommentInHtmlLongText bool, ) *NoteTextDto`

NewNoteTextDto instantiates a new NoteTextDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteTextDtoWithDefaults

`func NewNoteTextDtoWithDefaults() *NoteTextDto`

NewNoteTextDtoWithDefaults instantiates a new NoteTextDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOpeningText

`func (o *NoteTextDto) GetIsOpeningText() bool`

GetIsOpeningText returns the IsOpeningText field if non-nil, zero value otherwise.

### GetIsOpeningTextOk

`func (o *NoteTextDto) GetIsOpeningTextOk() (*bool, bool)`

GetIsOpeningTextOk returns a tuple with the IsOpeningText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpeningText

`func (o *NoteTextDto) SetIsOpeningText(v bool)`

SetIsOpeningText sets IsOpeningText field to given value.


### GetIsClosingText

`func (o *NoteTextDto) GetIsClosingText() bool`

GetIsClosingText returns the IsClosingText field if non-nil, zero value otherwise.

### GetIsClosingTextOk

`func (o *NoteTextDto) GetIsClosingTextOk() (*bool, bool)`

GetIsClosingTextOk returns a tuple with the IsClosingText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosingText

`func (o *NoteTextDto) SetIsClosingText(v bool)`

SetIsClosingText sets IsClosingText field to given value.


### GetShortText

`func (o *NoteTextDto) GetShortText() string`

GetShortText returns the ShortText field if non-nil, zero value otherwise.

### GetShortTextOk

`func (o *NoteTextDto) GetShortTextOk() (*string, bool)`

GetShortTextOk returns a tuple with the ShortText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortText

`func (o *NoteTextDto) SetShortText(v string)`

SetShortText sets ShortText field to given value.

### HasShortText

`func (o *NoteTextDto) HasShortText() bool`

HasShortText returns a boolean if a field has been set.

### GetAdditionType

`func (o *NoteTextDto) GetAdditionType() AdditionTypeDto`

GetAdditionType returns the AdditionType field if non-nil, zero value otherwise.

### GetAdditionTypeOk

`func (o *NoteTextDto) GetAdditionTypeOk() (*AdditionTypeDto, bool)`

GetAdditionTypeOk returns a tuple with the AdditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionType

`func (o *NoteTextDto) SetAdditionType(v AdditionTypeDto)`

SetAdditionType sets AdditionType field to given value.


### GetLongText

`func (o *NoteTextDto) GetLongText() string`

GetLongText returns the LongText field if non-nil, zero value otherwise.

### GetLongTextOk

`func (o *NoteTextDto) GetLongTextOk() (*string, bool)`

GetLongTextOk returns a tuple with the LongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongText

`func (o *NoteTextDto) SetLongText(v string)`

SetLongText sets LongText field to given value.

### HasLongText

`func (o *NoteTextDto) HasLongText() bool`

HasLongText returns a boolean if a field has been set.

### GetHtmlLongText

`func (o *NoteTextDto) GetHtmlLongText() string`

GetHtmlLongText returns the HtmlLongText field if non-nil, zero value otherwise.

### GetHtmlLongTextOk

`func (o *NoteTextDto) GetHtmlLongTextOk() (*string, bool)`

GetHtmlLongTextOk returns a tuple with the HtmlLongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlLongText

`func (o *NoteTextDto) SetHtmlLongText(v string)`

SetHtmlLongText sets HtmlLongText field to given value.

### HasHtmlLongText

`func (o *NoteTextDto) HasHtmlLongText() bool`

HasHtmlLongText returns a boolean if a field has been set.

### GetIdentifier

`func (o *NoteTextDto) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *NoteTextDto) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *NoteTextDto) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *NoteTextDto) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetStandardizedDescription

`func (o *NoteTextDto) GetStandardizedDescription() StandardizedDescriptionDto`

GetStandardizedDescription returns the StandardizedDescription field if non-nil, zero value otherwise.

### GetStandardizedDescriptionOk

`func (o *NoteTextDto) GetStandardizedDescriptionOk() (*StandardizedDescriptionDto, bool)`

GetStandardizedDescriptionOk returns a tuple with the StandardizedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardizedDescription

`func (o *NoteTextDto) SetStandardizedDescription(v StandardizedDescriptionDto)`

SetStandardizedDescription sets StandardizedDescription field to given value.

### HasStandardizedDescription

`func (o *NoteTextDto) HasStandardizedDescription() bool`

HasStandardizedDescription returns a boolean if a field has been set.

### GetElementType

`func (o *NoteTextDto) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *NoteTextDto) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *NoteTextDto) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *NoteTextDto) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetDescriptionId

`func (o *NoteTextDto) GetDescriptionId() string`

GetDescriptionId returns the DescriptionId field if non-nil, zero value otherwise.

### GetDescriptionIdOk

`func (o *NoteTextDto) GetDescriptionIdOk() (*string, bool)`

GetDescriptionIdOk returns a tuple with the DescriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionId

`func (o *NoteTextDto) SetDescriptionId(v string)`

SetDescriptionId sets DescriptionId field to given value.

### HasDescriptionId

`func (o *NoteTextDto) HasDescriptionId() bool`

HasDescriptionId returns a boolean if a field has been set.

### GetOenormNoteTextProperties

`func (o *NoteTextDto) GetOenormNoteTextProperties() OenormNoteTextPropertiesDto`

GetOenormNoteTextProperties returns the OenormNoteTextProperties field if non-nil, zero value otherwise.

### GetOenormNoteTextPropertiesOk

`func (o *NoteTextDto) GetOenormNoteTextPropertiesOk() (*OenormNoteTextPropertiesDto, bool)`

GetOenormNoteTextPropertiesOk returns a tuple with the OenormNoteTextProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOenormNoteTextProperties

`func (o *NoteTextDto) SetOenormNoteTextProperties(v OenormNoteTextPropertiesDto)`

SetOenormNoteTextProperties sets OenormNoteTextProperties field to given value.

### HasOenormNoteTextProperties

`func (o *NoteTextDto) HasOenormNoteTextProperties() bool`

HasOenormNoteTextProperties returns a boolean if a field has been set.

### GetHasBidderCommentInHtmlLongText

`func (o *NoteTextDto) GetHasBidderCommentInHtmlLongText() bool`

GetHasBidderCommentInHtmlLongText returns the HasBidderCommentInHtmlLongText field if non-nil, zero value otherwise.

### GetHasBidderCommentInHtmlLongTextOk

`func (o *NoteTextDto) GetHasBidderCommentInHtmlLongTextOk() (*bool, bool)`

GetHasBidderCommentInHtmlLongTextOk returns a tuple with the HasBidderCommentInHtmlLongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBidderCommentInHtmlLongText

`func (o *NoteTextDto) SetHasBidderCommentInHtmlLongText(v bool)`

SetHasBidderCommentInHtmlLongText sets HasBidderCommentInHtmlLongText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


