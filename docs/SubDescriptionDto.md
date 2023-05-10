# SubDescriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Quantity** | **float32** | Returns the total calculated sum of all quantity assignments. Will return the result rounded to three decimal places. | [readonly] 
**QuantityOverride** | Pointer to **float32** | You can use this property to directly set the quantity for this sub description. This will override any given QuantityComponents | [optional] 
**QuantityComponents** | Pointer to [**[]CalculationDto**](CalculationDto.md) | Holds quantity information for this sub description. Quantity is listening to changes here and is reporting the total sum of all quantity components. | [optional] 
**AmountToBeEnteredByBidder** | **bool** | Indicates if the bidder is asked to specify an amount. | 
**Identifier** | Pointer to **string** | Identifier for this SubDescription. | [optional] 
**ShortText** | Pointer to **string** | Short description for this DescriptionBase element. | [optional] 
**LongText** | Pointer to **string** | Detailed description for this DescriptionBase element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText. | [optional] 
**UnitTag** | Pointer to **string** | If this is given, then the sub description has a different unit tag than the parent position. | [optional] 
**HtmlLongText** | Pointer to **string** | This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out. | [optional] 
**AdditionType** | [**AdditionTypeDto**](AdditionTypeDto.md) |  | 
**StandardizedDescription** | Pointer to [**StandardizedDescriptionDto**](StandardizedDescriptionDto.md) |  | [optional] 
**ExecutionDescriptionReference** | Pointer to **string** | This identifier can be used to point to the Id of an ExecutionDescription in the same ServiceSpecification. ExecutionDescriptions act as a way to centrally describe how positions (or sub descriptions) should be executed in practice. Often, the position (or sub description) itself still has text of its own to highlight deviations from that or add more details. When working with import and export features, this property is only supported in GAEB 90 data exchange. | [optional] 
**ProjectCatalogues** | Pointer to [**[]CatalogueDto**](CatalogueDto.md) | These are Catalogues that are used within this Calculation. Catalogues are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. They are propagate to all child elements, e.g. other containers and QuantityAssignments. In the context of a ServiceSpecification, all elements share the same instance of the collection. | [optional] 
**DescriptionId** | Pointer to **string** | This is an identifier specific for this description. Some exchange formats, like GAEB XML, use it to identify descriptions. It&#39;s different to an elements identifier in that it should only apply to the description component, meaning the text itself. | [optional] 
**HasBidderCommentInHtmlLongText** | **bool** |  | [readonly] 
**ElementType** | Pointer to **string** |  | [optional] 
**ElementTypeDiscriminator** | Pointer to **string** |  | [optional] 

## Methods

### NewSubDescriptionDto

`func NewSubDescriptionDto(id string, quantity float32, amountToBeEnteredByBidder bool, additionType AdditionTypeDto, hasBidderCommentInHtmlLongText bool, ) *SubDescriptionDto`

NewSubDescriptionDto instantiates a new SubDescriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubDescriptionDtoWithDefaults

`func NewSubDescriptionDtoWithDefaults() *SubDescriptionDto`

NewSubDescriptionDtoWithDefaults instantiates a new SubDescriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubDescriptionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubDescriptionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubDescriptionDto) SetId(v string)`

SetId sets Id field to given value.


### GetQuantity

`func (o *SubDescriptionDto) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubDescriptionDto) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubDescriptionDto) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetQuantityOverride

`func (o *SubDescriptionDto) GetQuantityOverride() float32`

GetQuantityOverride returns the QuantityOverride field if non-nil, zero value otherwise.

### GetQuantityOverrideOk

`func (o *SubDescriptionDto) GetQuantityOverrideOk() (*float32, bool)`

GetQuantityOverrideOk returns a tuple with the QuantityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOverride

`func (o *SubDescriptionDto) SetQuantityOverride(v float32)`

SetQuantityOverride sets QuantityOverride field to given value.

### HasQuantityOverride

`func (o *SubDescriptionDto) HasQuantityOverride() bool`

HasQuantityOverride returns a boolean if a field has been set.

### GetQuantityComponents

`func (o *SubDescriptionDto) GetQuantityComponents() []CalculationDto`

GetQuantityComponents returns the QuantityComponents field if non-nil, zero value otherwise.

### GetQuantityComponentsOk

`func (o *SubDescriptionDto) GetQuantityComponentsOk() (*[]CalculationDto, bool)`

GetQuantityComponentsOk returns a tuple with the QuantityComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityComponents

`func (o *SubDescriptionDto) SetQuantityComponents(v []CalculationDto)`

SetQuantityComponents sets QuantityComponents field to given value.

### HasQuantityComponents

`func (o *SubDescriptionDto) HasQuantityComponents() bool`

HasQuantityComponents returns a boolean if a field has been set.

### GetAmountToBeEnteredByBidder

`func (o *SubDescriptionDto) GetAmountToBeEnteredByBidder() bool`

GetAmountToBeEnteredByBidder returns the AmountToBeEnteredByBidder field if non-nil, zero value otherwise.

### GetAmountToBeEnteredByBidderOk

`func (o *SubDescriptionDto) GetAmountToBeEnteredByBidderOk() (*bool, bool)`

GetAmountToBeEnteredByBidderOk returns a tuple with the AmountToBeEnteredByBidder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToBeEnteredByBidder

`func (o *SubDescriptionDto) SetAmountToBeEnteredByBidder(v bool)`

SetAmountToBeEnteredByBidder sets AmountToBeEnteredByBidder field to given value.


### GetIdentifier

`func (o *SubDescriptionDto) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SubDescriptionDto) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SubDescriptionDto) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *SubDescriptionDto) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetShortText

`func (o *SubDescriptionDto) GetShortText() string`

GetShortText returns the ShortText field if non-nil, zero value otherwise.

### GetShortTextOk

`func (o *SubDescriptionDto) GetShortTextOk() (*string, bool)`

GetShortTextOk returns a tuple with the ShortText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortText

`func (o *SubDescriptionDto) SetShortText(v string)`

SetShortText sets ShortText field to given value.

### HasShortText

`func (o *SubDescriptionDto) HasShortText() bool`

HasShortText returns a boolean if a field has been set.

### GetLongText

`func (o *SubDescriptionDto) GetLongText() string`

GetLongText returns the LongText field if non-nil, zero value otherwise.

### GetLongTextOk

`func (o *SubDescriptionDto) GetLongTextOk() (*string, bool)`

GetLongTextOk returns a tuple with the LongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongText

`func (o *SubDescriptionDto) SetLongText(v string)`

SetLongText sets LongText field to given value.

### HasLongText

`func (o *SubDescriptionDto) HasLongText() bool`

HasLongText returns a boolean if a field has been set.

### GetUnitTag

`func (o *SubDescriptionDto) GetUnitTag() string`

GetUnitTag returns the UnitTag field if non-nil, zero value otherwise.

### GetUnitTagOk

`func (o *SubDescriptionDto) GetUnitTagOk() (*string, bool)`

GetUnitTagOk returns a tuple with the UnitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitTag

`func (o *SubDescriptionDto) SetUnitTag(v string)`

SetUnitTag sets UnitTag field to given value.

### HasUnitTag

`func (o *SubDescriptionDto) HasUnitTag() bool`

HasUnitTag returns a boolean if a field has been set.

### GetHtmlLongText

`func (o *SubDescriptionDto) GetHtmlLongText() string`

GetHtmlLongText returns the HtmlLongText field if non-nil, zero value otherwise.

### GetHtmlLongTextOk

`func (o *SubDescriptionDto) GetHtmlLongTextOk() (*string, bool)`

GetHtmlLongTextOk returns a tuple with the HtmlLongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlLongText

`func (o *SubDescriptionDto) SetHtmlLongText(v string)`

SetHtmlLongText sets HtmlLongText field to given value.

### HasHtmlLongText

`func (o *SubDescriptionDto) HasHtmlLongText() bool`

HasHtmlLongText returns a boolean if a field has been set.

### GetAdditionType

`func (o *SubDescriptionDto) GetAdditionType() AdditionTypeDto`

GetAdditionType returns the AdditionType field if non-nil, zero value otherwise.

### GetAdditionTypeOk

`func (o *SubDescriptionDto) GetAdditionTypeOk() (*AdditionTypeDto, bool)`

GetAdditionTypeOk returns a tuple with the AdditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionType

`func (o *SubDescriptionDto) SetAdditionType(v AdditionTypeDto)`

SetAdditionType sets AdditionType field to given value.


### GetStandardizedDescription

`func (o *SubDescriptionDto) GetStandardizedDescription() StandardizedDescriptionDto`

GetStandardizedDescription returns the StandardizedDescription field if non-nil, zero value otherwise.

### GetStandardizedDescriptionOk

`func (o *SubDescriptionDto) GetStandardizedDescriptionOk() (*StandardizedDescriptionDto, bool)`

GetStandardizedDescriptionOk returns a tuple with the StandardizedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardizedDescription

`func (o *SubDescriptionDto) SetStandardizedDescription(v StandardizedDescriptionDto)`

SetStandardizedDescription sets StandardizedDescription field to given value.

### HasStandardizedDescription

`func (o *SubDescriptionDto) HasStandardizedDescription() bool`

HasStandardizedDescription returns a boolean if a field has been set.

### GetExecutionDescriptionReference

`func (o *SubDescriptionDto) GetExecutionDescriptionReference() string`

GetExecutionDescriptionReference returns the ExecutionDescriptionReference field if non-nil, zero value otherwise.

### GetExecutionDescriptionReferenceOk

`func (o *SubDescriptionDto) GetExecutionDescriptionReferenceOk() (*string, bool)`

GetExecutionDescriptionReferenceOk returns a tuple with the ExecutionDescriptionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDescriptionReference

`func (o *SubDescriptionDto) SetExecutionDescriptionReference(v string)`

SetExecutionDescriptionReference sets ExecutionDescriptionReference field to given value.

### HasExecutionDescriptionReference

`func (o *SubDescriptionDto) HasExecutionDescriptionReference() bool`

HasExecutionDescriptionReference returns a boolean if a field has been set.

### GetProjectCatalogues

`func (o *SubDescriptionDto) GetProjectCatalogues() []CatalogueDto`

GetProjectCatalogues returns the ProjectCatalogues field if non-nil, zero value otherwise.

### GetProjectCataloguesOk

`func (o *SubDescriptionDto) GetProjectCataloguesOk() (*[]CatalogueDto, bool)`

GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCatalogues

`func (o *SubDescriptionDto) SetProjectCatalogues(v []CatalogueDto)`

SetProjectCatalogues sets ProjectCatalogues field to given value.

### HasProjectCatalogues

`func (o *SubDescriptionDto) HasProjectCatalogues() bool`

HasProjectCatalogues returns a boolean if a field has been set.

### GetDescriptionId

`func (o *SubDescriptionDto) GetDescriptionId() string`

GetDescriptionId returns the DescriptionId field if non-nil, zero value otherwise.

### GetDescriptionIdOk

`func (o *SubDescriptionDto) GetDescriptionIdOk() (*string, bool)`

GetDescriptionIdOk returns a tuple with the DescriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionId

`func (o *SubDescriptionDto) SetDescriptionId(v string)`

SetDescriptionId sets DescriptionId field to given value.

### HasDescriptionId

`func (o *SubDescriptionDto) HasDescriptionId() bool`

HasDescriptionId returns a boolean if a field has been set.

### GetHasBidderCommentInHtmlLongText

`func (o *SubDescriptionDto) GetHasBidderCommentInHtmlLongText() bool`

GetHasBidderCommentInHtmlLongText returns the HasBidderCommentInHtmlLongText field if non-nil, zero value otherwise.

### GetHasBidderCommentInHtmlLongTextOk

`func (o *SubDescriptionDto) GetHasBidderCommentInHtmlLongTextOk() (*bool, bool)`

GetHasBidderCommentInHtmlLongTextOk returns a tuple with the HasBidderCommentInHtmlLongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBidderCommentInHtmlLongText

`func (o *SubDescriptionDto) SetHasBidderCommentInHtmlLongText(v bool)`

SetHasBidderCommentInHtmlLongText sets HasBidderCommentInHtmlLongText field to given value.


### GetElementType

`func (o *SubDescriptionDto) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *SubDescriptionDto) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *SubDescriptionDto) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *SubDescriptionDto) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetElementTypeDiscriminator

`func (o *SubDescriptionDto) GetElementTypeDiscriminator() string`

GetElementTypeDiscriminator returns the ElementTypeDiscriminator field if non-nil, zero value otherwise.

### GetElementTypeDiscriminatorOk

`func (o *SubDescriptionDto) GetElementTypeDiscriminatorOk() (*string, bool)`

GetElementTypeDiscriminatorOk returns a tuple with the ElementTypeDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementTypeDiscriminator

`func (o *SubDescriptionDto) SetElementTypeDiscriminator(v string)`

SetElementTypeDiscriminator sets ElementTypeDiscriminator field to given value.

### HasElementTypeDiscriminator

`func (o *SubDescriptionDto) HasElementTypeDiscriminator() bool`

HasElementTypeDiscriminator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


