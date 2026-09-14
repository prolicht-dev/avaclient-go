# LineItemInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | BT-153 A name for the line item | [optional] 
**Description** | Pointer to **string** | BT-154 A description for the line item | [optional] 
**SellerIdentifier** | Pointer to **string** | BT-155 An identifier assigned by the seller, e.g. an internal article number | [optional] 
**BuyerIdentifier** | Pointer to **string** | BT-156 An identifier assigned by the buyer, e.g. an internal article number | [optional] 
**StandardIdentifier** | Pointer to **string** | BT-157 A standard identifier for the line item, e.g. an EAN or a GTIN | [optional] 
**StandardSchemeId** | Pointer to **string** | BT-157 If used together with StandardIdentifier, this is the scheme identifier for which standard is used. | [optional] 
**ClassificationIdentifiers** | Pointer to **[]string** | BT-158 A list of classification identifiers for the line item. You can further use ClassificationTypes to specify the scheme identifier for each classification identifier. | [optional] 
**ClassificationTypes** | Pointer to **[]string** | BT-158 This is used in connection with ClassificationIdentifiers. It is the scheme identifier for each identifier. If this list is used, it should have the same length as ClassificationIdentifiers, with null or empty values used for non-existing scheme identifiers. | [optional] 
**CountryOfOrigin** | Pointer to **string** | BT-159 The origin country code according to ISO 3166-1 in 2 letter format | [optional] 
**Attributes** | Pointer to [**[]LineItemAttribute**](LineItemAttribute.md) | BG-32 Optional additional attributes for this line item | [optional] 

## Methods

### NewLineItemInformation

`func NewLineItemInformation() *LineItemInformation`

NewLineItemInformation instantiates a new LineItemInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemInformationWithDefaults

`func NewLineItemInformationWithDefaults() *LineItemInformation`

NewLineItemInformationWithDefaults instantiates a new LineItemInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LineItemInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LineItemInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LineItemInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LineItemInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LineItemInformation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LineItemInformation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LineItemInformation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LineItemInformation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSellerIdentifier

`func (o *LineItemInformation) GetSellerIdentifier() string`

GetSellerIdentifier returns the SellerIdentifier field if non-nil, zero value otherwise.

### GetSellerIdentifierOk

`func (o *LineItemInformation) GetSellerIdentifierOk() (*string, bool)`

GetSellerIdentifierOk returns a tuple with the SellerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerIdentifier

`func (o *LineItemInformation) SetSellerIdentifier(v string)`

SetSellerIdentifier sets SellerIdentifier field to given value.

### HasSellerIdentifier

`func (o *LineItemInformation) HasSellerIdentifier() bool`

HasSellerIdentifier returns a boolean if a field has been set.

### GetBuyerIdentifier

`func (o *LineItemInformation) GetBuyerIdentifier() string`

GetBuyerIdentifier returns the BuyerIdentifier field if non-nil, zero value otherwise.

### GetBuyerIdentifierOk

`func (o *LineItemInformation) GetBuyerIdentifierOk() (*string, bool)`

GetBuyerIdentifierOk returns a tuple with the BuyerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerIdentifier

`func (o *LineItemInformation) SetBuyerIdentifier(v string)`

SetBuyerIdentifier sets BuyerIdentifier field to given value.

### HasBuyerIdentifier

`func (o *LineItemInformation) HasBuyerIdentifier() bool`

HasBuyerIdentifier returns a boolean if a field has been set.

### GetStandardIdentifier

`func (o *LineItemInformation) GetStandardIdentifier() string`

GetStandardIdentifier returns the StandardIdentifier field if non-nil, zero value otherwise.

### GetStandardIdentifierOk

`func (o *LineItemInformation) GetStandardIdentifierOk() (*string, bool)`

GetStandardIdentifierOk returns a tuple with the StandardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardIdentifier

`func (o *LineItemInformation) SetStandardIdentifier(v string)`

SetStandardIdentifier sets StandardIdentifier field to given value.

### HasStandardIdentifier

`func (o *LineItemInformation) HasStandardIdentifier() bool`

HasStandardIdentifier returns a boolean if a field has been set.

### GetStandardSchemeId

`func (o *LineItemInformation) GetStandardSchemeId() string`

GetStandardSchemeId returns the StandardSchemeId field if non-nil, zero value otherwise.

### GetStandardSchemeIdOk

`func (o *LineItemInformation) GetStandardSchemeIdOk() (*string, bool)`

GetStandardSchemeIdOk returns a tuple with the StandardSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardSchemeId

`func (o *LineItemInformation) SetStandardSchemeId(v string)`

SetStandardSchemeId sets StandardSchemeId field to given value.

### HasStandardSchemeId

`func (o *LineItemInformation) HasStandardSchemeId() bool`

HasStandardSchemeId returns a boolean if a field has been set.

### GetClassificationIdentifiers

`func (o *LineItemInformation) GetClassificationIdentifiers() []string`

GetClassificationIdentifiers returns the ClassificationIdentifiers field if non-nil, zero value otherwise.

### GetClassificationIdentifiersOk

`func (o *LineItemInformation) GetClassificationIdentifiersOk() (*[]string, bool)`

GetClassificationIdentifiersOk returns a tuple with the ClassificationIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationIdentifiers

`func (o *LineItemInformation) SetClassificationIdentifiers(v []string)`

SetClassificationIdentifiers sets ClassificationIdentifiers field to given value.

### HasClassificationIdentifiers

`func (o *LineItemInformation) HasClassificationIdentifiers() bool`

HasClassificationIdentifiers returns a boolean if a field has been set.

### GetClassificationTypes

`func (o *LineItemInformation) GetClassificationTypes() []string`

GetClassificationTypes returns the ClassificationTypes field if non-nil, zero value otherwise.

### GetClassificationTypesOk

`func (o *LineItemInformation) GetClassificationTypesOk() (*[]string, bool)`

GetClassificationTypesOk returns a tuple with the ClassificationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationTypes

`func (o *LineItemInformation) SetClassificationTypes(v []string)`

SetClassificationTypes sets ClassificationTypes field to given value.

### HasClassificationTypes

`func (o *LineItemInformation) HasClassificationTypes() bool`

HasClassificationTypes returns a boolean if a field has been set.

### GetCountryOfOrigin

`func (o *LineItemInformation) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *LineItemInformation) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *LineItemInformation) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *LineItemInformation) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### GetAttributes

`func (o *LineItemInformation) GetAttributes() []LineItemAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LineItemInformation) GetAttributesOk() (*[]LineItemAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LineItemInformation) SetAttributes(v []LineItemAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LineItemInformation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


