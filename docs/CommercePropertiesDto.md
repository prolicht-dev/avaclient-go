# CommercePropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArticleNumber** | Pointer to **string** | This maps to ArtNo in GAEB XML and represents an article number given by the supplier. | [optional] 
**EanGtinArticleNumber** | Pointer to **string** | This maps to EAN in GAEB XML and represents an GTIN (formerly EAN) article number. | [optional] 
**IlnArticleNumber** | Pointer to **string** | This maps to ArtNoID in GAEB XML and represents an ILN article number. | [optional] 
**CatalogueNumber** | Pointer to **string** | This maps to CatalogNo in GAEB XML and represents an identifier of a specific catalogue. The referenced catalogue is usually a customer specific one, not related to CatalogueReferences. | [optional] 
**CatalogueArticleNumber** | Pointer to **string** | This maps to CatalogArtNo in GAEB XML and represents a key that maps to an entry in a specific catalogue. The referenced catalogue is usually a customer specific one, not related to CatalogueReferences. | [optional] 
**PriceBasis** | Pointer to **float32** | This optional property can be used to indicate the basis for prices for a single position. Price basis means that the price is given per unit of the basis, e.g. per a pack of five when this property is set to &#39;5&#39;. | [optional] 
**SubPositionIdentifier** | Pointer to **string** | This optional property can be used to indicate a sub position identifier specific for the commerce exchange | [optional] 

## Methods

### NewCommercePropertiesDto

`func NewCommercePropertiesDto() *CommercePropertiesDto`

NewCommercePropertiesDto instantiates a new CommercePropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommercePropertiesDtoWithDefaults

`func NewCommercePropertiesDtoWithDefaults() *CommercePropertiesDto`

NewCommercePropertiesDtoWithDefaults instantiates a new CommercePropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticleNumber

`func (o *CommercePropertiesDto) GetArticleNumber() string`

GetArticleNumber returns the ArticleNumber field if non-nil, zero value otherwise.

### GetArticleNumberOk

`func (o *CommercePropertiesDto) GetArticleNumberOk() (*string, bool)`

GetArticleNumberOk returns a tuple with the ArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleNumber

`func (o *CommercePropertiesDto) SetArticleNumber(v string)`

SetArticleNumber sets ArticleNumber field to given value.

### HasArticleNumber

`func (o *CommercePropertiesDto) HasArticleNumber() bool`

HasArticleNumber returns a boolean if a field has been set.

### GetEanGtinArticleNumber

`func (o *CommercePropertiesDto) GetEanGtinArticleNumber() string`

GetEanGtinArticleNumber returns the EanGtinArticleNumber field if non-nil, zero value otherwise.

### GetEanGtinArticleNumberOk

`func (o *CommercePropertiesDto) GetEanGtinArticleNumberOk() (*string, bool)`

GetEanGtinArticleNumberOk returns a tuple with the EanGtinArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEanGtinArticleNumber

`func (o *CommercePropertiesDto) SetEanGtinArticleNumber(v string)`

SetEanGtinArticleNumber sets EanGtinArticleNumber field to given value.

### HasEanGtinArticleNumber

`func (o *CommercePropertiesDto) HasEanGtinArticleNumber() bool`

HasEanGtinArticleNumber returns a boolean if a field has been set.

### GetIlnArticleNumber

`func (o *CommercePropertiesDto) GetIlnArticleNumber() string`

GetIlnArticleNumber returns the IlnArticleNumber field if non-nil, zero value otherwise.

### GetIlnArticleNumberOk

`func (o *CommercePropertiesDto) GetIlnArticleNumberOk() (*string, bool)`

GetIlnArticleNumberOk returns a tuple with the IlnArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIlnArticleNumber

`func (o *CommercePropertiesDto) SetIlnArticleNumber(v string)`

SetIlnArticleNumber sets IlnArticleNumber field to given value.

### HasIlnArticleNumber

`func (o *CommercePropertiesDto) HasIlnArticleNumber() bool`

HasIlnArticleNumber returns a boolean if a field has been set.

### GetCatalogueNumber

`func (o *CommercePropertiesDto) GetCatalogueNumber() string`

GetCatalogueNumber returns the CatalogueNumber field if non-nil, zero value otherwise.

### GetCatalogueNumberOk

`func (o *CommercePropertiesDto) GetCatalogueNumberOk() (*string, bool)`

GetCatalogueNumberOk returns a tuple with the CatalogueNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueNumber

`func (o *CommercePropertiesDto) SetCatalogueNumber(v string)`

SetCatalogueNumber sets CatalogueNumber field to given value.

### HasCatalogueNumber

`func (o *CommercePropertiesDto) HasCatalogueNumber() bool`

HasCatalogueNumber returns a boolean if a field has been set.

### GetCatalogueArticleNumber

`func (o *CommercePropertiesDto) GetCatalogueArticleNumber() string`

GetCatalogueArticleNumber returns the CatalogueArticleNumber field if non-nil, zero value otherwise.

### GetCatalogueArticleNumberOk

`func (o *CommercePropertiesDto) GetCatalogueArticleNumberOk() (*string, bool)`

GetCatalogueArticleNumberOk returns a tuple with the CatalogueArticleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueArticleNumber

`func (o *CommercePropertiesDto) SetCatalogueArticleNumber(v string)`

SetCatalogueArticleNumber sets CatalogueArticleNumber field to given value.

### HasCatalogueArticleNumber

`func (o *CommercePropertiesDto) HasCatalogueArticleNumber() bool`

HasCatalogueArticleNumber returns a boolean if a field has been set.

### GetPriceBasis

`func (o *CommercePropertiesDto) GetPriceBasis() float32`

GetPriceBasis returns the PriceBasis field if non-nil, zero value otherwise.

### GetPriceBasisOk

`func (o *CommercePropertiesDto) GetPriceBasisOk() (*float32, bool)`

GetPriceBasisOk returns a tuple with the PriceBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBasis

`func (o *CommercePropertiesDto) SetPriceBasis(v float32)`

SetPriceBasis sets PriceBasis field to given value.

### HasPriceBasis

`func (o *CommercePropertiesDto) HasPriceBasis() bool`

HasPriceBasis returns a boolean if a field has been set.

### GetSubPositionIdentifier

`func (o *CommercePropertiesDto) GetSubPositionIdentifier() string`

GetSubPositionIdentifier returns the SubPositionIdentifier field if non-nil, zero value otherwise.

### GetSubPositionIdentifierOk

`func (o *CommercePropertiesDto) GetSubPositionIdentifierOk() (*string, bool)`

GetSubPositionIdentifierOk returns a tuple with the SubPositionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPositionIdentifier

`func (o *CommercePropertiesDto) SetSubPositionIdentifier(v string)`

SetSubPositionIdentifier sets SubPositionIdentifier field to given value.

### HasSubPositionIdentifier

`func (o *CommercePropertiesDto) HasSubPositionIdentifier() bool`

HasSubPositionIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


