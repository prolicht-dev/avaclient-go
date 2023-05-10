# ProductDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Vendor** | Pointer to [**PartyInformationDto**](PartyInformationDto.md) |  | [optional] 
**Articles** | Pointer to [**[]ArticleDataDto**](ArticleDataDto.md) | The collection of ArticleData that describe this product, e.g. a pipe product could be composed out of multiple pipe segments and fittings. | [optional] 
**ShortText** | Pointer to **string** | Short description for this ITextElement element. | [optional] 
**LongText** | Pointer to **string** | Detailed description for this ITextElement element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText. | [optional] 
**HtmlLongText** | Pointer to **string** | This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out. | [optional] 

## Methods

### NewProductDataDto

`func NewProductDataDto(id string, ) *ProductDataDto`

NewProductDataDto instantiates a new ProductDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDataDtoWithDefaults

`func NewProductDataDtoWithDefaults() *ProductDataDto`

NewProductDataDtoWithDefaults instantiates a new ProductDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductDataDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductDataDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductDataDto) SetId(v string)`

SetId sets Id field to given value.


### GetVendor

`func (o *ProductDataDto) GetVendor() PartyInformationDto`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ProductDataDto) GetVendorOk() (*PartyInformationDto, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ProductDataDto) SetVendor(v PartyInformationDto)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ProductDataDto) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetArticles

`func (o *ProductDataDto) GetArticles() []ArticleDataDto`

GetArticles returns the Articles field if non-nil, zero value otherwise.

### GetArticlesOk

`func (o *ProductDataDto) GetArticlesOk() (*[]ArticleDataDto, bool)`

GetArticlesOk returns a tuple with the Articles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticles

`func (o *ProductDataDto) SetArticles(v []ArticleDataDto)`

SetArticles sets Articles field to given value.

### HasArticles

`func (o *ProductDataDto) HasArticles() bool`

HasArticles returns a boolean if a field has been set.

### GetShortText

`func (o *ProductDataDto) GetShortText() string`

GetShortText returns the ShortText field if non-nil, zero value otherwise.

### GetShortTextOk

`func (o *ProductDataDto) GetShortTextOk() (*string, bool)`

GetShortTextOk returns a tuple with the ShortText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortText

`func (o *ProductDataDto) SetShortText(v string)`

SetShortText sets ShortText field to given value.

### HasShortText

`func (o *ProductDataDto) HasShortText() bool`

HasShortText returns a boolean if a field has been set.

### GetLongText

`func (o *ProductDataDto) GetLongText() string`

GetLongText returns the LongText field if non-nil, zero value otherwise.

### GetLongTextOk

`func (o *ProductDataDto) GetLongTextOk() (*string, bool)`

GetLongTextOk returns a tuple with the LongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongText

`func (o *ProductDataDto) SetLongText(v string)`

SetLongText sets LongText field to given value.

### HasLongText

`func (o *ProductDataDto) HasLongText() bool`

HasLongText returns a boolean if a field has been set.

### GetHtmlLongText

`func (o *ProductDataDto) GetHtmlLongText() string`

GetHtmlLongText returns the HtmlLongText field if non-nil, zero value otherwise.

### GetHtmlLongTextOk

`func (o *ProductDataDto) GetHtmlLongTextOk() (*string, bool)`

GetHtmlLongTextOk returns a tuple with the HtmlLongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlLongText

`func (o *ProductDataDto) SetHtmlLongText(v string)`

SetHtmlLongText sets HtmlLongText field to given value.

### HasHtmlLongText

`func (o *ProductDataDto) HasHtmlLongText() bool`

HasHtmlLongText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


