# ArticleDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Name** | Pointer to **string** | The name (or brand name) for this article, usually given by the supplier or vendor. | [optional] 
**ArticleNumber** | Pointer to **string** | An article number that describes it, useful when integrating other systems. | [optional] 
**Quantity** | **float32** | Quantity for this article. If this is used within a Position, the quantity here should be the quantity required for the full quantity of the position, not for a single unit. | 
**UnitTag** | Pointer to **string** | The unit tag for this single ArticleData. | [optional] 
**Description** | Pointer to **string** | This is an optional text element that can be used to further describe the ArticleData. | [optional] 
**ShortText** | Pointer to **string** | Short description for this ITextElement element. | [optional] 
**LongText** | Pointer to **string** | Detailed description for this ITextElement element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText. | [optional] 
**HtmlLongText** | Pointer to **string** | This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out. | [optional] 

## Methods

### NewArticleDataDto

`func NewArticleDataDto(id string, quantity float32, ) *ArticleDataDto`

NewArticleDataDto instantiates a new ArticleDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArticleDataDtoWithDefaults

`func NewArticleDataDtoWithDefaults() *ArticleDataDto`

NewArticleDataDtoWithDefaults instantiates a new ArticleDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArticleDataDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArticleDataDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArticleDataDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ArticleDataDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArticleDataDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArticleDataDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArticleDataDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArticleNumber

`func (o *ArticleDataDto) GetArticleNumber() string`

GetArticleNumber returns the ArticleNumber field if non-nil, zero value otherwise.

### GetArticleNumberOk

`func (o *ArticleDataDto) GetArticleNumberOk() (*string, bool)`

GetArticleNumberOk returns a tuple with the ArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleNumber

`func (o *ArticleDataDto) SetArticleNumber(v string)`

SetArticleNumber sets ArticleNumber field to given value.

### HasArticleNumber

`func (o *ArticleDataDto) HasArticleNumber() bool`

HasArticleNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *ArticleDataDto) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ArticleDataDto) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ArticleDataDto) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitTag

`func (o *ArticleDataDto) GetUnitTag() string`

GetUnitTag returns the UnitTag field if non-nil, zero value otherwise.

### GetUnitTagOk

`func (o *ArticleDataDto) GetUnitTagOk() (*string, bool)`

GetUnitTagOk returns a tuple with the UnitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitTag

`func (o *ArticleDataDto) SetUnitTag(v string)`

SetUnitTag sets UnitTag field to given value.

### HasUnitTag

`func (o *ArticleDataDto) HasUnitTag() bool`

HasUnitTag returns a boolean if a field has been set.

### GetDescription

`func (o *ArticleDataDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArticleDataDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArticleDataDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArticleDataDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortText

`func (o *ArticleDataDto) GetShortText() string`

GetShortText returns the ShortText field if non-nil, zero value otherwise.

### GetShortTextOk

`func (o *ArticleDataDto) GetShortTextOk() (*string, bool)`

GetShortTextOk returns a tuple with the ShortText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortText

`func (o *ArticleDataDto) SetShortText(v string)`

SetShortText sets ShortText field to given value.

### HasShortText

`func (o *ArticleDataDto) HasShortText() bool`

HasShortText returns a boolean if a field has been set.

### GetLongText

`func (o *ArticleDataDto) GetLongText() string`

GetLongText returns the LongText field if non-nil, zero value otherwise.

### GetLongTextOk

`func (o *ArticleDataDto) GetLongTextOk() (*string, bool)`

GetLongTextOk returns a tuple with the LongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongText

`func (o *ArticleDataDto) SetLongText(v string)`

SetLongText sets LongText field to given value.

### HasLongText

`func (o *ArticleDataDto) HasLongText() bool`

HasLongText returns a boolean if a field has been set.

### GetHtmlLongText

`func (o *ArticleDataDto) GetHtmlLongText() string`

GetHtmlLongText returns the HtmlLongText field if non-nil, zero value otherwise.

### GetHtmlLongTextOk

`func (o *ArticleDataDto) GetHtmlLongTextOk() (*string, bool)`

GetHtmlLongTextOk returns a tuple with the HtmlLongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlLongText

`func (o *ArticleDataDto) SetHtmlLongText(v string)`

SetHtmlLongText sets HtmlLongText field to given value.

### HasHtmlLongText

`func (o *ArticleDataDto) HasHtmlLongText() bool`

HasHtmlLongText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


