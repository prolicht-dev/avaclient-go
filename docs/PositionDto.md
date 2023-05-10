# PositionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitPrice** | **float32** | Will return the price per unit, rounded according to the project settings or the default value of three decimal places | [readonly] 
**UnitPriceOverride** | Pointer to **float32** | You can use this property to directly set the unit price for this position. This will override any given PriceComponents | [optional] 
**Quantity** | **float32** | Will return this Position&#39;s total quantity, rounded to three decimal places. | [readonly] 
**QuantityOverride** | Pointer to **float32** | You can use this property to directly set the quantity for this position. This will override any given QuantityComponents | [optional] 
**IsComplementingPosition** | **bool** | This indicates true if this specific position is specified as a complementing position for any base position. | 
**ComplementsPositions** | Pointer to **[]string** | If IsComplementingPosition is set to true, this will indicate which base positions are targeted by this complementing position. | [optional] 
**ComplementingPricePercentageOverride** | Pointer to **float32** | You can use this property to directly specify the total price of this position as a percentage of the sum of the total prices of base positions that this position complements. It will essentially set the total price for this position to the sum of all total prices of the positions specified in ComplementsPositions multiplied by the percentage here. If present, this has precedence over UnitPriceOverride as well as QuantityOverride. | [optional] 
**ComplementingPricePercentage** | Pointer to **float32** | This is a read only property showing the total price of this positions as a percentage of the sum of the total prices of all base positions, in case this position is a complementing positions. If this position is not a complementing position, this will be null. Also, if no prices are present, this will also be null. | [optional] 
**UnitTag** | Pointer to **string** | The tag of the unit used for this positions quantity. | [optional] 
**LabourComponents** | Pointer to [**LabourPriceComponentDto**](LabourPriceComponentDto.md) |  | [optional] 
**PriceComponents** | Pointer to [**[]PriceComponentDto**](PriceComponentDto.md) | The single price components in this Position. Price components should not be handled directly on a per-position basis but set globally in the parent Projects ProjectInformation. | [optional] 
**QuantityComponents** | Pointer to [**[]CalculationDto**](CalculationDto.md) | The quantity components of this Position. | [optional] 
**SubDescriptions** | Pointer to [**[]SubDescriptionDto**](SubDescriptionDto.md) | Further structuring of this Position. | [optional] 
**ComissionStatus** | [**ComissionStatusDto**](ComissionStatusDto.md) |  | 
**ComplementedBy** | Pointer to **[]string** | A list of positions that complement this Position. The positions are referenced by their GUIDs. It might be used together with ComplementedByQuantities in case that only a given quantity is complemented by positions. | [optional] 
**Complemented** | **bool** | Will indicate if this Position is complemented in this ServiceSpecification by other Positions. It can not be set to false when there are entries in the ComplementedBy property. | 
**AmountToBeEnteredByBidder** | **bool** | Indicates that the amount for this Position is to be set by the bidder. | 
**PriceCompositionRequired** | **bool** | Indicates if the bidder demands for prices to be broken up into their price components. | 
**UseDifferentTaxRate** | **bool** | Indicates if this Position should use a different TaxRate than what is the default for the Project. | 
**TaxRate** | **float32** | Will return either the parent ServiceSpecification&#39;s TaxRate or it&#39;s own if it has a different value. (For example, in Germany Water has a different TaxRate than regular Positions) | 
**ItemNumber** | Pointer to [**ItemNumberDto**](ItemNumberDto.md) |  | [optional] 
**DeductionFactor** | **float32** | The rate of deductions, i.e. 0.12m means 12% price deduction. | 
**TotalPrice** | **float32** | Returns the product of UnitPrice times Quantity. | [readonly] 
**TotalPriceGross** | **float32** | The total gross price for this Position. | [readonly] 
**TotalPriceGrossDeducted** | **float32** | Total gross price after applied deductions. | [readonly] 
**DeductedPrice** | **float32** | Net price after applied deductions. Please be aware that this is the total price, from TotalPrice, not the unit price of the position. | [readonly] 
**PositionType** | [**PositionTypeDto**](PositionTypeDto.md) |  | 
**PriceType** | [**PriceTypeDto**](PriceTypeDto.md) |  | 
**ServiceType** | [**ServiceTypeDto**](ServiceTypeDto.md) |  | 
**ProductData** | Pointer to [**ProductDataDto**](ProductDataDto.md) |  | [optional] 
**ShortText** | Pointer to **string** | Short description for this DescriptionBase element. | [optional] 
**LongText** | Pointer to **string** | Detailed description for this DescriptionBase element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText. | [optional] 
**HtmlLongText** | Pointer to **string** | This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out. | [optional] 
**AdditionType** | [**AdditionTypeDto**](AdditionTypeDto.md) |  | 
**ElementType** | Pointer to **string** |  | [optional] 
**QuantityAssignments** | Pointer to [**[]QuantityAssignmentDto**](QuantityAssignmentDto.md) | Quantity assignments are, in contrast to SubDescriptions, used to categorize parts of this Position. For example, it could be categorized by cost group - e.g. a Position describing concrete walls could follow the German DIN 276 Cost Groups Standard and specify that of the total 1.000m² wall, 500m² are classified as exterior walls and 500m² are classified as interior walls. They would then have different cost groups associated, e.g. for accounting purposes. | [optional] 
**CommerceProperties** | Pointer to [**CommercePropertiesDto**](CommercePropertiesDto.md) |  | [optional] 
**AlternativeTo** | Pointer to **string** | If this position is an Alternative, then this property should point to the id of the position in this service specification that it can replace. If this is set to a value, you can optionally also specify an identifier via AlternativeIdentifier to specifiy multiple positions that must be used together to be an alternative to a single base position. | [optional] 
**AlternativeIdentifier** | Pointer to **int32** | This is an optional property that can be used together with AlternativeTo. If this is set, you can indicate which alternative group a specific position is assigned to. That way, if you specifiy multiple alternative Positions with the same AlternativeIdentifier, you can indicate that to replace a single base Position, multiple alternative Positions should be used. This property is not checked or managed automatically, so it is possible for this property to become invalid, by for example setting this property but not setting a base position via AlternativeTo. | [optional] 
**AlternativeGroupIdentifier** | Pointer to **int32** | This is an optional property that can be used together with AlternativeTo and AlternativeGroupIdentifier. If this is set, you can indicate which alternative group a specific position is assigned to. That way, you can specify the id (in integer format) for the alternative group this position belongs to. It&#39;s different to AlternativeIdentifier in that the other property describes the id of the group, while this property here describes the group itself. If a position only has set AlternativeGroupIdentifier but not AlternativeIdentifier, then it likely is a base position for a specific group. This was introduced in version v2.9.0 to be able to model both position and group ids for alternative group combinations. This property is not checked or managed automatically, so it is possible for this property to become invalid, by for example setting this property but not setting a base position via AlternativeTo. | [optional] 
**IsLumpSum** | **bool** | If this is true, it indicates that the position is intended to be a lump sum, \&quot;Pauschal\&quot; in German. This means the position total price that is being invoiced should not depend on the actual quantity. In practice, partial invoices are still often used for partial services rendered. This property does not affect the price calculation for this position. Typically, the Quantity should be set to 1.0 when this flag is used. | 
**RepetitionTo** | Pointer to **string** | This identifier can be used to point to the Id of a position in the same ServiceSpecification that acts as a base position. It matches \&quot;Bezugsposition\&quot; in GAEB. This can be used for positions that repeat partially or are linked together | [optional] 
**StandardizedDescription** | Pointer to [**StandardizedDescriptionDto**](StandardizedDescriptionDto.md) |  | [optional] 
**ComplementedByQuantities** | Pointer to [**[]ComplementedByQuantityDto**](ComplementedByQuantityDto.md) | This list contains references to positions that complement this one, additionally also specifying a quantity for which the addition is intended. This does not replace the ComplementedBy property and there are no automatic checks being done between these two properties, so it&#39;s up to the user code to ensure deletions (and additions, if desired) are performed for both properties. When copying withing keeping Ids, this list will not be part of the copy process, since it would only contain quantities without actual position references. Containers, however, will rebuild the list with updated position references when copying positions that contain entries here. | [optional] 
**ExecutionDescriptionReference** | Pointer to **string** | This identifier can be used to point to the Id of an ExecutionDescription in the same ServiceSpecification. ExecutionDescriptions act as a way to centrally describe how positions should be executed in practice. Often, the position itself still has text of its own to highlight deviations from that or add more details. | [optional] 
**NotOffered** | **bool** | This indicates if a position has not been offered. This is typically only expected to be true when the exchange phase of the parent ServiceSpecification is Offer, and it means that the position has not been offered at all. | 
**OenormPositionProperties** | Pointer to [**OenormPositionPropertiesDto**](OenormPositionPropertiesDto.md) |  | [optional] 
**DescriptionId** | Pointer to **string** | This is an identifier specific for this description. Some exchange formats, like GAEB XML, use it to identify descriptions. It&#39;s different to an elements identifier in that it should only apply to the description component, meaning the text itself. | [optional] 
**HierarchyLevel** | **int32** | This is a zero based hierarchy level. It&#39;s set automatically when used in the context of a Project, and can be used to identify the hierarchy level of the current element. | 
**AddendumStatus** | Pointer to [**AddendumStatusDto**](AddendumStatusDto.md) |  | [optional] 
**HasBidderCommentInHtmlLongText** | **bool** |  | [readonly] 
**GaebComplementingType** | [**PositionComplementingTypeDto**](PositionComplementingTypeDto.md) |  | 
**HoldOutProperties** | Pointer to [**PositionHoldOutPropertiesDto**](PositionHoldOutPropertiesDto.md) |  | [optional] 
**EstimatedQuantity** | Pointer to **float32** | This is an informational property, which directly holds a numerical value for an estimated quantity. It is not used for any price calculations. | [optional] 

## Methods

### NewPositionDto

`func NewPositionDto(unitPrice float32, quantity float32, isComplementingPosition bool, comissionStatus ComissionStatusDto, complemented bool, amountToBeEnteredByBidder bool, priceCompositionRequired bool, useDifferentTaxRate bool, taxRate float32, deductionFactor float32, totalPrice float32, totalPriceGross float32, totalPriceGrossDeducted float32, deductedPrice float32, positionType PositionTypeDto, priceType PriceTypeDto, serviceType ServiceTypeDto, additionType AdditionTypeDto, isLumpSum bool, notOffered bool, hierarchyLevel int32, hasBidderCommentInHtmlLongText bool, gaebComplementingType PositionComplementingTypeDto, ) *PositionDto`

NewPositionDto instantiates a new PositionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionDtoWithDefaults

`func NewPositionDtoWithDefaults() *PositionDto`

NewPositionDtoWithDefaults instantiates a new PositionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitPrice

`func (o *PositionDto) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *PositionDto) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *PositionDto) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetUnitPriceOverride

`func (o *PositionDto) GetUnitPriceOverride() float32`

GetUnitPriceOverride returns the UnitPriceOverride field if non-nil, zero value otherwise.

### GetUnitPriceOverrideOk

`func (o *PositionDto) GetUnitPriceOverrideOk() (*float32, bool)`

GetUnitPriceOverrideOk returns a tuple with the UnitPriceOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceOverride

`func (o *PositionDto) SetUnitPriceOverride(v float32)`

SetUnitPriceOverride sets UnitPriceOverride field to given value.

### HasUnitPriceOverride

`func (o *PositionDto) HasUnitPriceOverride() bool`

HasUnitPriceOverride returns a boolean if a field has been set.

### GetQuantity

`func (o *PositionDto) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PositionDto) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PositionDto) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetQuantityOverride

`func (o *PositionDto) GetQuantityOverride() float32`

GetQuantityOverride returns the QuantityOverride field if non-nil, zero value otherwise.

### GetQuantityOverrideOk

`func (o *PositionDto) GetQuantityOverrideOk() (*float32, bool)`

GetQuantityOverrideOk returns a tuple with the QuantityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOverride

`func (o *PositionDto) SetQuantityOverride(v float32)`

SetQuantityOverride sets QuantityOverride field to given value.

### HasQuantityOverride

`func (o *PositionDto) HasQuantityOverride() bool`

HasQuantityOverride returns a boolean if a field has been set.

### GetIsComplementingPosition

`func (o *PositionDto) GetIsComplementingPosition() bool`

GetIsComplementingPosition returns the IsComplementingPosition field if non-nil, zero value otherwise.

### GetIsComplementingPositionOk

`func (o *PositionDto) GetIsComplementingPositionOk() (*bool, bool)`

GetIsComplementingPositionOk returns a tuple with the IsComplementingPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplementingPosition

`func (o *PositionDto) SetIsComplementingPosition(v bool)`

SetIsComplementingPosition sets IsComplementingPosition field to given value.


### GetComplementsPositions

`func (o *PositionDto) GetComplementsPositions() []string`

GetComplementsPositions returns the ComplementsPositions field if non-nil, zero value otherwise.

### GetComplementsPositionsOk

`func (o *PositionDto) GetComplementsPositionsOk() (*[]string, bool)`

GetComplementsPositionsOk returns a tuple with the ComplementsPositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementsPositions

`func (o *PositionDto) SetComplementsPositions(v []string)`

SetComplementsPositions sets ComplementsPositions field to given value.

### HasComplementsPositions

`func (o *PositionDto) HasComplementsPositions() bool`

HasComplementsPositions returns a boolean if a field has been set.

### GetComplementingPricePercentageOverride

`func (o *PositionDto) GetComplementingPricePercentageOverride() float32`

GetComplementingPricePercentageOverride returns the ComplementingPricePercentageOverride field if non-nil, zero value otherwise.

### GetComplementingPricePercentageOverrideOk

`func (o *PositionDto) GetComplementingPricePercentageOverrideOk() (*float32, bool)`

GetComplementingPricePercentageOverrideOk returns a tuple with the ComplementingPricePercentageOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementingPricePercentageOverride

`func (o *PositionDto) SetComplementingPricePercentageOverride(v float32)`

SetComplementingPricePercentageOverride sets ComplementingPricePercentageOverride field to given value.

### HasComplementingPricePercentageOverride

`func (o *PositionDto) HasComplementingPricePercentageOverride() bool`

HasComplementingPricePercentageOverride returns a boolean if a field has been set.

### GetComplementingPricePercentage

`func (o *PositionDto) GetComplementingPricePercentage() float32`

GetComplementingPricePercentage returns the ComplementingPricePercentage field if non-nil, zero value otherwise.

### GetComplementingPricePercentageOk

`func (o *PositionDto) GetComplementingPricePercentageOk() (*float32, bool)`

GetComplementingPricePercentageOk returns a tuple with the ComplementingPricePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementingPricePercentage

`func (o *PositionDto) SetComplementingPricePercentage(v float32)`

SetComplementingPricePercentage sets ComplementingPricePercentage field to given value.

### HasComplementingPricePercentage

`func (o *PositionDto) HasComplementingPricePercentage() bool`

HasComplementingPricePercentage returns a boolean if a field has been set.

### GetUnitTag

`func (o *PositionDto) GetUnitTag() string`

GetUnitTag returns the UnitTag field if non-nil, zero value otherwise.

### GetUnitTagOk

`func (o *PositionDto) GetUnitTagOk() (*string, bool)`

GetUnitTagOk returns a tuple with the UnitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitTag

`func (o *PositionDto) SetUnitTag(v string)`

SetUnitTag sets UnitTag field to given value.

### HasUnitTag

`func (o *PositionDto) HasUnitTag() bool`

HasUnitTag returns a boolean if a field has been set.

### GetLabourComponents

`func (o *PositionDto) GetLabourComponents() LabourPriceComponentDto`

GetLabourComponents returns the LabourComponents field if non-nil, zero value otherwise.

### GetLabourComponentsOk

`func (o *PositionDto) GetLabourComponentsOk() (*LabourPriceComponentDto, bool)`

GetLabourComponentsOk returns a tuple with the LabourComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabourComponents

`func (o *PositionDto) SetLabourComponents(v LabourPriceComponentDto)`

SetLabourComponents sets LabourComponents field to given value.

### HasLabourComponents

`func (o *PositionDto) HasLabourComponents() bool`

HasLabourComponents returns a boolean if a field has been set.

### GetPriceComponents

`func (o *PositionDto) GetPriceComponents() []PriceComponentDto`

GetPriceComponents returns the PriceComponents field if non-nil, zero value otherwise.

### GetPriceComponentsOk

`func (o *PositionDto) GetPriceComponentsOk() (*[]PriceComponentDto, bool)`

GetPriceComponentsOk returns a tuple with the PriceComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceComponents

`func (o *PositionDto) SetPriceComponents(v []PriceComponentDto)`

SetPriceComponents sets PriceComponents field to given value.

### HasPriceComponents

`func (o *PositionDto) HasPriceComponents() bool`

HasPriceComponents returns a boolean if a field has been set.

### GetQuantityComponents

`func (o *PositionDto) GetQuantityComponents() []CalculationDto`

GetQuantityComponents returns the QuantityComponents field if non-nil, zero value otherwise.

### GetQuantityComponentsOk

`func (o *PositionDto) GetQuantityComponentsOk() (*[]CalculationDto, bool)`

GetQuantityComponentsOk returns a tuple with the QuantityComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityComponents

`func (o *PositionDto) SetQuantityComponents(v []CalculationDto)`

SetQuantityComponents sets QuantityComponents field to given value.

### HasQuantityComponents

`func (o *PositionDto) HasQuantityComponents() bool`

HasQuantityComponents returns a boolean if a field has been set.

### GetSubDescriptions

`func (o *PositionDto) GetSubDescriptions() []SubDescriptionDto`

GetSubDescriptions returns the SubDescriptions field if non-nil, zero value otherwise.

### GetSubDescriptionsOk

`func (o *PositionDto) GetSubDescriptionsOk() (*[]SubDescriptionDto, bool)`

GetSubDescriptionsOk returns a tuple with the SubDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubDescriptions

`func (o *PositionDto) SetSubDescriptions(v []SubDescriptionDto)`

SetSubDescriptions sets SubDescriptions field to given value.

### HasSubDescriptions

`func (o *PositionDto) HasSubDescriptions() bool`

HasSubDescriptions returns a boolean if a field has been set.

### GetComissionStatus

`func (o *PositionDto) GetComissionStatus() ComissionStatusDto`

GetComissionStatus returns the ComissionStatus field if non-nil, zero value otherwise.

### GetComissionStatusOk

`func (o *PositionDto) GetComissionStatusOk() (*ComissionStatusDto, bool)`

GetComissionStatusOk returns a tuple with the ComissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComissionStatus

`func (o *PositionDto) SetComissionStatus(v ComissionStatusDto)`

SetComissionStatus sets ComissionStatus field to given value.


### GetComplementedBy

`func (o *PositionDto) GetComplementedBy() []string`

GetComplementedBy returns the ComplementedBy field if non-nil, zero value otherwise.

### GetComplementedByOk

`func (o *PositionDto) GetComplementedByOk() (*[]string, bool)`

GetComplementedByOk returns a tuple with the ComplementedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementedBy

`func (o *PositionDto) SetComplementedBy(v []string)`

SetComplementedBy sets ComplementedBy field to given value.

### HasComplementedBy

`func (o *PositionDto) HasComplementedBy() bool`

HasComplementedBy returns a boolean if a field has been set.

### GetComplemented

`func (o *PositionDto) GetComplemented() bool`

GetComplemented returns the Complemented field if non-nil, zero value otherwise.

### GetComplementedOk

`func (o *PositionDto) GetComplementedOk() (*bool, bool)`

GetComplementedOk returns a tuple with the Complemented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplemented

`func (o *PositionDto) SetComplemented(v bool)`

SetComplemented sets Complemented field to given value.


### GetAmountToBeEnteredByBidder

`func (o *PositionDto) GetAmountToBeEnteredByBidder() bool`

GetAmountToBeEnteredByBidder returns the AmountToBeEnteredByBidder field if non-nil, zero value otherwise.

### GetAmountToBeEnteredByBidderOk

`func (o *PositionDto) GetAmountToBeEnteredByBidderOk() (*bool, bool)`

GetAmountToBeEnteredByBidderOk returns a tuple with the AmountToBeEnteredByBidder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToBeEnteredByBidder

`func (o *PositionDto) SetAmountToBeEnteredByBidder(v bool)`

SetAmountToBeEnteredByBidder sets AmountToBeEnteredByBidder field to given value.


### GetPriceCompositionRequired

`func (o *PositionDto) GetPriceCompositionRequired() bool`

GetPriceCompositionRequired returns the PriceCompositionRequired field if non-nil, zero value otherwise.

### GetPriceCompositionRequiredOk

`func (o *PositionDto) GetPriceCompositionRequiredOk() (*bool, bool)`

GetPriceCompositionRequiredOk returns a tuple with the PriceCompositionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCompositionRequired

`func (o *PositionDto) SetPriceCompositionRequired(v bool)`

SetPriceCompositionRequired sets PriceCompositionRequired field to given value.


### GetUseDifferentTaxRate

`func (o *PositionDto) GetUseDifferentTaxRate() bool`

GetUseDifferentTaxRate returns the UseDifferentTaxRate field if non-nil, zero value otherwise.

### GetUseDifferentTaxRateOk

`func (o *PositionDto) GetUseDifferentTaxRateOk() (*bool, bool)`

GetUseDifferentTaxRateOk returns a tuple with the UseDifferentTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDifferentTaxRate

`func (o *PositionDto) SetUseDifferentTaxRate(v bool)`

SetUseDifferentTaxRate sets UseDifferentTaxRate field to given value.


### GetTaxRate

`func (o *PositionDto) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *PositionDto) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *PositionDto) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.


### GetItemNumber

`func (o *PositionDto) GetItemNumber() ItemNumberDto`

GetItemNumber returns the ItemNumber field if non-nil, zero value otherwise.

### GetItemNumberOk

`func (o *PositionDto) GetItemNumberOk() (*ItemNumberDto, bool)`

GetItemNumberOk returns a tuple with the ItemNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemNumber

`func (o *PositionDto) SetItemNumber(v ItemNumberDto)`

SetItemNumber sets ItemNumber field to given value.

### HasItemNumber

`func (o *PositionDto) HasItemNumber() bool`

HasItemNumber returns a boolean if a field has been set.

### GetDeductionFactor

`func (o *PositionDto) GetDeductionFactor() float32`

GetDeductionFactor returns the DeductionFactor field if non-nil, zero value otherwise.

### GetDeductionFactorOk

`func (o *PositionDto) GetDeductionFactorOk() (*float32, bool)`

GetDeductionFactorOk returns a tuple with the DeductionFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductionFactor

`func (o *PositionDto) SetDeductionFactor(v float32)`

SetDeductionFactor sets DeductionFactor field to given value.


### GetTotalPrice

`func (o *PositionDto) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *PositionDto) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *PositionDto) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.


### GetTotalPriceGross

`func (o *PositionDto) GetTotalPriceGross() float32`

GetTotalPriceGross returns the TotalPriceGross field if non-nil, zero value otherwise.

### GetTotalPriceGrossOk

`func (o *PositionDto) GetTotalPriceGrossOk() (*float32, bool)`

GetTotalPriceGrossOk returns a tuple with the TotalPriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceGross

`func (o *PositionDto) SetTotalPriceGross(v float32)`

SetTotalPriceGross sets TotalPriceGross field to given value.


### GetTotalPriceGrossDeducted

`func (o *PositionDto) GetTotalPriceGrossDeducted() float32`

GetTotalPriceGrossDeducted returns the TotalPriceGrossDeducted field if non-nil, zero value otherwise.

### GetTotalPriceGrossDeductedOk

`func (o *PositionDto) GetTotalPriceGrossDeductedOk() (*float32, bool)`

GetTotalPriceGrossDeductedOk returns a tuple with the TotalPriceGrossDeducted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceGrossDeducted

`func (o *PositionDto) SetTotalPriceGrossDeducted(v float32)`

SetTotalPriceGrossDeducted sets TotalPriceGrossDeducted field to given value.


### GetDeductedPrice

`func (o *PositionDto) GetDeductedPrice() float32`

GetDeductedPrice returns the DeductedPrice field if non-nil, zero value otherwise.

### GetDeductedPriceOk

`func (o *PositionDto) GetDeductedPriceOk() (*float32, bool)`

GetDeductedPriceOk returns a tuple with the DeductedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductedPrice

`func (o *PositionDto) SetDeductedPrice(v float32)`

SetDeductedPrice sets DeductedPrice field to given value.


### GetPositionType

`func (o *PositionDto) GetPositionType() PositionTypeDto`

GetPositionType returns the PositionType field if non-nil, zero value otherwise.

### GetPositionTypeOk

`func (o *PositionDto) GetPositionTypeOk() (*PositionTypeDto, bool)`

GetPositionTypeOk returns a tuple with the PositionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionType

`func (o *PositionDto) SetPositionType(v PositionTypeDto)`

SetPositionType sets PositionType field to given value.


### GetPriceType

`func (o *PositionDto) GetPriceType() PriceTypeDto`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *PositionDto) GetPriceTypeOk() (*PriceTypeDto, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *PositionDto) SetPriceType(v PriceTypeDto)`

SetPriceType sets PriceType field to given value.


### GetServiceType

`func (o *PositionDto) GetServiceType() ServiceTypeDto`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PositionDto) GetServiceTypeOk() (*ServiceTypeDto, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PositionDto) SetServiceType(v ServiceTypeDto)`

SetServiceType sets ServiceType field to given value.


### GetProductData

`func (o *PositionDto) GetProductData() ProductDataDto`

GetProductData returns the ProductData field if non-nil, zero value otherwise.

### GetProductDataOk

`func (o *PositionDto) GetProductDataOk() (*ProductDataDto, bool)`

GetProductDataOk returns a tuple with the ProductData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductData

`func (o *PositionDto) SetProductData(v ProductDataDto)`

SetProductData sets ProductData field to given value.

### HasProductData

`func (o *PositionDto) HasProductData() bool`

HasProductData returns a boolean if a field has been set.

### GetShortText

`func (o *PositionDto) GetShortText() string`

GetShortText returns the ShortText field if non-nil, zero value otherwise.

### GetShortTextOk

`func (o *PositionDto) GetShortTextOk() (*string, bool)`

GetShortTextOk returns a tuple with the ShortText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortText

`func (o *PositionDto) SetShortText(v string)`

SetShortText sets ShortText field to given value.

### HasShortText

`func (o *PositionDto) HasShortText() bool`

HasShortText returns a boolean if a field has been set.

### GetLongText

`func (o *PositionDto) GetLongText() string`

GetLongText returns the LongText field if non-nil, zero value otherwise.

### GetLongTextOk

`func (o *PositionDto) GetLongTextOk() (*string, bool)`

GetLongTextOk returns a tuple with the LongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongText

`func (o *PositionDto) SetLongText(v string)`

SetLongText sets LongText field to given value.

### HasLongText

`func (o *PositionDto) HasLongText() bool`

HasLongText returns a boolean if a field has been set.

### GetHtmlLongText

`func (o *PositionDto) GetHtmlLongText() string`

GetHtmlLongText returns the HtmlLongText field if non-nil, zero value otherwise.

### GetHtmlLongTextOk

`func (o *PositionDto) GetHtmlLongTextOk() (*string, bool)`

GetHtmlLongTextOk returns a tuple with the HtmlLongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlLongText

`func (o *PositionDto) SetHtmlLongText(v string)`

SetHtmlLongText sets HtmlLongText field to given value.

### HasHtmlLongText

`func (o *PositionDto) HasHtmlLongText() bool`

HasHtmlLongText returns a boolean if a field has been set.

### GetAdditionType

`func (o *PositionDto) GetAdditionType() AdditionTypeDto`

GetAdditionType returns the AdditionType field if non-nil, zero value otherwise.

### GetAdditionTypeOk

`func (o *PositionDto) GetAdditionTypeOk() (*AdditionTypeDto, bool)`

GetAdditionTypeOk returns a tuple with the AdditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionType

`func (o *PositionDto) SetAdditionType(v AdditionTypeDto)`

SetAdditionType sets AdditionType field to given value.


### GetElementType

`func (o *PositionDto) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *PositionDto) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *PositionDto) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *PositionDto) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetQuantityAssignments

`func (o *PositionDto) GetQuantityAssignments() []QuantityAssignmentDto`

GetQuantityAssignments returns the QuantityAssignments field if non-nil, zero value otherwise.

### GetQuantityAssignmentsOk

`func (o *PositionDto) GetQuantityAssignmentsOk() (*[]QuantityAssignmentDto, bool)`

GetQuantityAssignmentsOk returns a tuple with the QuantityAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAssignments

`func (o *PositionDto) SetQuantityAssignments(v []QuantityAssignmentDto)`

SetQuantityAssignments sets QuantityAssignments field to given value.

### HasQuantityAssignments

`func (o *PositionDto) HasQuantityAssignments() bool`

HasQuantityAssignments returns a boolean if a field has been set.

### GetCommerceProperties

`func (o *PositionDto) GetCommerceProperties() CommercePropertiesDto`

GetCommerceProperties returns the CommerceProperties field if non-nil, zero value otherwise.

### GetCommercePropertiesOk

`func (o *PositionDto) GetCommercePropertiesOk() (*CommercePropertiesDto, bool)`

GetCommercePropertiesOk returns a tuple with the CommerceProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommerceProperties

`func (o *PositionDto) SetCommerceProperties(v CommercePropertiesDto)`

SetCommerceProperties sets CommerceProperties field to given value.

### HasCommerceProperties

`func (o *PositionDto) HasCommerceProperties() bool`

HasCommerceProperties returns a boolean if a field has been set.

### GetAlternativeTo

`func (o *PositionDto) GetAlternativeTo() string`

GetAlternativeTo returns the AlternativeTo field if non-nil, zero value otherwise.

### GetAlternativeToOk

`func (o *PositionDto) GetAlternativeToOk() (*string, bool)`

GetAlternativeToOk returns a tuple with the AlternativeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeTo

`func (o *PositionDto) SetAlternativeTo(v string)`

SetAlternativeTo sets AlternativeTo field to given value.

### HasAlternativeTo

`func (o *PositionDto) HasAlternativeTo() bool`

HasAlternativeTo returns a boolean if a field has been set.

### GetAlternativeIdentifier

`func (o *PositionDto) GetAlternativeIdentifier() int32`

GetAlternativeIdentifier returns the AlternativeIdentifier field if non-nil, zero value otherwise.

### GetAlternativeIdentifierOk

`func (o *PositionDto) GetAlternativeIdentifierOk() (*int32, bool)`

GetAlternativeIdentifierOk returns a tuple with the AlternativeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeIdentifier

`func (o *PositionDto) SetAlternativeIdentifier(v int32)`

SetAlternativeIdentifier sets AlternativeIdentifier field to given value.

### HasAlternativeIdentifier

`func (o *PositionDto) HasAlternativeIdentifier() bool`

HasAlternativeIdentifier returns a boolean if a field has been set.

### GetAlternativeGroupIdentifier

`func (o *PositionDto) GetAlternativeGroupIdentifier() int32`

GetAlternativeGroupIdentifier returns the AlternativeGroupIdentifier field if non-nil, zero value otherwise.

### GetAlternativeGroupIdentifierOk

`func (o *PositionDto) GetAlternativeGroupIdentifierOk() (*int32, bool)`

GetAlternativeGroupIdentifierOk returns a tuple with the AlternativeGroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeGroupIdentifier

`func (o *PositionDto) SetAlternativeGroupIdentifier(v int32)`

SetAlternativeGroupIdentifier sets AlternativeGroupIdentifier field to given value.

### HasAlternativeGroupIdentifier

`func (o *PositionDto) HasAlternativeGroupIdentifier() bool`

HasAlternativeGroupIdentifier returns a boolean if a field has been set.

### GetIsLumpSum

`func (o *PositionDto) GetIsLumpSum() bool`

GetIsLumpSum returns the IsLumpSum field if non-nil, zero value otherwise.

### GetIsLumpSumOk

`func (o *PositionDto) GetIsLumpSumOk() (*bool, bool)`

GetIsLumpSumOk returns a tuple with the IsLumpSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLumpSum

`func (o *PositionDto) SetIsLumpSum(v bool)`

SetIsLumpSum sets IsLumpSum field to given value.


### GetRepetitionTo

`func (o *PositionDto) GetRepetitionTo() string`

GetRepetitionTo returns the RepetitionTo field if non-nil, zero value otherwise.

### GetRepetitionToOk

`func (o *PositionDto) GetRepetitionToOk() (*string, bool)`

GetRepetitionToOk returns a tuple with the RepetitionTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionTo

`func (o *PositionDto) SetRepetitionTo(v string)`

SetRepetitionTo sets RepetitionTo field to given value.

### HasRepetitionTo

`func (o *PositionDto) HasRepetitionTo() bool`

HasRepetitionTo returns a boolean if a field has been set.

### GetStandardizedDescription

`func (o *PositionDto) GetStandardizedDescription() StandardizedDescriptionDto`

GetStandardizedDescription returns the StandardizedDescription field if non-nil, zero value otherwise.

### GetStandardizedDescriptionOk

`func (o *PositionDto) GetStandardizedDescriptionOk() (*StandardizedDescriptionDto, bool)`

GetStandardizedDescriptionOk returns a tuple with the StandardizedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardizedDescription

`func (o *PositionDto) SetStandardizedDescription(v StandardizedDescriptionDto)`

SetStandardizedDescription sets StandardizedDescription field to given value.

### HasStandardizedDescription

`func (o *PositionDto) HasStandardizedDescription() bool`

HasStandardizedDescription returns a boolean if a field has been set.

### GetComplementedByQuantities

`func (o *PositionDto) GetComplementedByQuantities() []ComplementedByQuantityDto`

GetComplementedByQuantities returns the ComplementedByQuantities field if non-nil, zero value otherwise.

### GetComplementedByQuantitiesOk

`func (o *PositionDto) GetComplementedByQuantitiesOk() (*[]ComplementedByQuantityDto, bool)`

GetComplementedByQuantitiesOk returns a tuple with the ComplementedByQuantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplementedByQuantities

`func (o *PositionDto) SetComplementedByQuantities(v []ComplementedByQuantityDto)`

SetComplementedByQuantities sets ComplementedByQuantities field to given value.

### HasComplementedByQuantities

`func (o *PositionDto) HasComplementedByQuantities() bool`

HasComplementedByQuantities returns a boolean if a field has been set.

### GetExecutionDescriptionReference

`func (o *PositionDto) GetExecutionDescriptionReference() string`

GetExecutionDescriptionReference returns the ExecutionDescriptionReference field if non-nil, zero value otherwise.

### GetExecutionDescriptionReferenceOk

`func (o *PositionDto) GetExecutionDescriptionReferenceOk() (*string, bool)`

GetExecutionDescriptionReferenceOk returns a tuple with the ExecutionDescriptionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDescriptionReference

`func (o *PositionDto) SetExecutionDescriptionReference(v string)`

SetExecutionDescriptionReference sets ExecutionDescriptionReference field to given value.

### HasExecutionDescriptionReference

`func (o *PositionDto) HasExecutionDescriptionReference() bool`

HasExecutionDescriptionReference returns a boolean if a field has been set.

### GetNotOffered

`func (o *PositionDto) GetNotOffered() bool`

GetNotOffered returns the NotOffered field if non-nil, zero value otherwise.

### GetNotOfferedOk

`func (o *PositionDto) GetNotOfferedOk() (*bool, bool)`

GetNotOfferedOk returns a tuple with the NotOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotOffered

`func (o *PositionDto) SetNotOffered(v bool)`

SetNotOffered sets NotOffered field to given value.


### GetOenormPositionProperties

`func (o *PositionDto) GetOenormPositionProperties() OenormPositionPropertiesDto`

GetOenormPositionProperties returns the OenormPositionProperties field if non-nil, zero value otherwise.

### GetOenormPositionPropertiesOk

`func (o *PositionDto) GetOenormPositionPropertiesOk() (*OenormPositionPropertiesDto, bool)`

GetOenormPositionPropertiesOk returns a tuple with the OenormPositionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOenormPositionProperties

`func (o *PositionDto) SetOenormPositionProperties(v OenormPositionPropertiesDto)`

SetOenormPositionProperties sets OenormPositionProperties field to given value.

### HasOenormPositionProperties

`func (o *PositionDto) HasOenormPositionProperties() bool`

HasOenormPositionProperties returns a boolean if a field has been set.

### GetDescriptionId

`func (o *PositionDto) GetDescriptionId() string`

GetDescriptionId returns the DescriptionId field if non-nil, zero value otherwise.

### GetDescriptionIdOk

`func (o *PositionDto) GetDescriptionIdOk() (*string, bool)`

GetDescriptionIdOk returns a tuple with the DescriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionId

`func (o *PositionDto) SetDescriptionId(v string)`

SetDescriptionId sets DescriptionId field to given value.

### HasDescriptionId

`func (o *PositionDto) HasDescriptionId() bool`

HasDescriptionId returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *PositionDto) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *PositionDto) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *PositionDto) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.


### GetAddendumStatus

`func (o *PositionDto) GetAddendumStatus() AddendumStatusDto`

GetAddendumStatus returns the AddendumStatus field if non-nil, zero value otherwise.

### GetAddendumStatusOk

`func (o *PositionDto) GetAddendumStatusOk() (*AddendumStatusDto, bool)`

GetAddendumStatusOk returns a tuple with the AddendumStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddendumStatus

`func (o *PositionDto) SetAddendumStatus(v AddendumStatusDto)`

SetAddendumStatus sets AddendumStatus field to given value.

### HasAddendumStatus

`func (o *PositionDto) HasAddendumStatus() bool`

HasAddendumStatus returns a boolean if a field has been set.

### GetHasBidderCommentInHtmlLongText

`func (o *PositionDto) GetHasBidderCommentInHtmlLongText() bool`

GetHasBidderCommentInHtmlLongText returns the HasBidderCommentInHtmlLongText field if non-nil, zero value otherwise.

### GetHasBidderCommentInHtmlLongTextOk

`func (o *PositionDto) GetHasBidderCommentInHtmlLongTextOk() (*bool, bool)`

GetHasBidderCommentInHtmlLongTextOk returns a tuple with the HasBidderCommentInHtmlLongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBidderCommentInHtmlLongText

`func (o *PositionDto) SetHasBidderCommentInHtmlLongText(v bool)`

SetHasBidderCommentInHtmlLongText sets HasBidderCommentInHtmlLongText field to given value.


### GetGaebComplementingType

`func (o *PositionDto) GetGaebComplementingType() PositionComplementingTypeDto`

GetGaebComplementingType returns the GaebComplementingType field if non-nil, zero value otherwise.

### GetGaebComplementingTypeOk

`func (o *PositionDto) GetGaebComplementingTypeOk() (*PositionComplementingTypeDto, bool)`

GetGaebComplementingTypeOk returns a tuple with the GaebComplementingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaebComplementingType

`func (o *PositionDto) SetGaebComplementingType(v PositionComplementingTypeDto)`

SetGaebComplementingType sets GaebComplementingType field to given value.


### GetHoldOutProperties

`func (o *PositionDto) GetHoldOutProperties() PositionHoldOutPropertiesDto`

GetHoldOutProperties returns the HoldOutProperties field if non-nil, zero value otherwise.

### GetHoldOutPropertiesOk

`func (o *PositionDto) GetHoldOutPropertiesOk() (*PositionHoldOutPropertiesDto, bool)`

GetHoldOutPropertiesOk returns a tuple with the HoldOutProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldOutProperties

`func (o *PositionDto) SetHoldOutProperties(v PositionHoldOutPropertiesDto)`

SetHoldOutProperties sets HoldOutProperties field to given value.

### HasHoldOutProperties

`func (o *PositionDto) HasHoldOutProperties() bool`

HasHoldOutProperties returns a boolean if a field has been set.

### GetEstimatedQuantity

`func (o *PositionDto) GetEstimatedQuantity() float32`

GetEstimatedQuantity returns the EstimatedQuantity field if non-nil, zero value otherwise.

### GetEstimatedQuantityOk

`func (o *PositionDto) GetEstimatedQuantityOk() (*float32, bool)`

GetEstimatedQuantityOk returns a tuple with the EstimatedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedQuantity

`func (o *PositionDto) SetEstimatedQuantity(v float32)`

SetEstimatedQuantity sets EstimatedQuantity field to given value.

### HasEstimatedQuantity

`func (o *PositionDto) HasEstimatedQuantity() bool`

HasEstimatedQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


