# PositionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**GaebXmlId** | **string** | This is used to store the GAEB XML Id within this IElement. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization. | [optional] 
**ElementTypeDiscriminator** | **string** |  | 
**UnitPrice** | **float32** | Will return the price per unit, rounded according to the project settings or the default value of three decimal places | [readonly] 
**UnitPriceOverride** | **float32** | You can use this property to directly set the unit price for this position. This will override any given PriceComponents | [optional] 
**Quantity** | **float32** | Will return this Position&#39;s total quantity, rounded to three decimal places. | [readonly] 
**QuantityOverride** | **float32** | You can use this property to directly set the quantity for this position. This will override any given QuantityComponents | [optional] 
**UnitTag** | **string** | The tag of the unit used for this positions quantity. | [optional] 
**LabourComponents** | [**LabourPriceComponentDto**](LabourPriceComponentDto.md) |  | [optional] 
**PriceComponents** | [**[]PriceComponentDto**](PriceComponentDto.md) | The single price components in this Position. Price components should not be handled directly on a per-position basis but set globally in the parent Projects ProjectInformation. | [optional] 
**QuantityComponents** | [**[]CalculationDto**](CalculationDto.md) | The quantity components of this Position. | [optional] 
**SubDescriptions** | [**[]SubDescriptionDto**](SubDescriptionDto.md) | Further structuring of this Position. | [optional] 
**ComissionStatus** | [**ComissionStatusDto**](ComissionStatusDto.md) |  | 
**ComplementedBy** | **[]string** | A list of positions that complement this Position. The positions are referenced by their GUIDs. It might be used together with ComplementedByQuantities in case that only a given quantity is complemented by positions. | [optional] 
**Complemented** | **bool** | Will indicate if this Position is complemented in this ServiceSpecification by other Positions. It can not be set to false when there are entries in the ComplementedBy property. | 
**AmountToBeEnteredByBidder** | **bool** | Indicates that the amount for this Position is to be set by the bidder. | 
**PriceCompositionRequired** | **bool** | Indicates if the bidder demands for prices to be broken up into their price components. | 
**UseDifferentTaxRate** | **bool** | Indicates if this Position should use a different TaxRate than what is the default for the Project. | 
**TaxRate** | **float32** | Will return either the parent ServiceSpecification&#39;s TaxRate or it&#39;s own if it has a different value. (For example, in Germany Water has a different TaxRate than regular Positions) | 
**ItemNumber** | [**ItemNumberDto**](ItemNumberDto.md) |  | [optional] 
**DeductionFactor** | **float32** | The rate of deductions, i.e. 0.12m means 12% price deduction. | 
**TotalPrice** | **float32** | Returns the product of UnitPrice times Quantity. | [readonly] 
**TotalPriceGross** | **float32** | The total gross price for this Position. | [readonly] 
**TotalPriceGrossDeducted** | **float32** | Total gross price after applied deductions. | [readonly] 
**DeductedPrice** | **float32** | Net price after applied deductions. | [readonly] 
**PositionType** | [**PositionTypeDto**](PositionTypeDto.md) |  | 
**PriceType** | [**PriceTypeDto**](PriceTypeDto.md) |  | 
**ServiceType** | [**ServiceTypeDto**](ServiceTypeDto.md) |  | 
**ProductData** | [**ProductDataDto**](ProductDataDto.md) |  | [optional] 
**ShortText** | **string** | Short description for this DescriptionBase element. | [optional] 
**LongText** | **string** | Detailed description for this DescriptionBase element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText. | [optional] 
**HtmlLongText** | **string** | This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out. | [optional] 
**AdditionType** | [**AdditionTypeDto**](AdditionTypeDto.md) |  | 
**ElementType** | **string** |  | [optional] 
**QuantityAssignments** | [**[]QuantityAssignmentDto**](QuantityAssignmentDto.md) | Quantity assignments are, in contrast to SubDescriptions, used to categorize parts of this Position. For example, it could be categorized by cost group - e.g. a Position describing concrete walls could follow the German DIN 276 Cost Groups Standard and specify that of the total 1.000m² wall, 500m² are classified as exterior walls and 500m² are classified as interior walls. They would then have different cost groups associated, e.g. for accounting purposes. | [optional] 
**CommerceProperties** | [**CommercePropertiesDto**](CommercePropertiesDto.md) |  | [optional] 
**AlternativeTo** | **string** | If this position is an Alternative, then this property should point to the id of the position in this service specification that it can replace. | [optional] 
**IsLumpSum** | **bool** | If this is true, it indicates that the position is intended to be a lump sum, \&quot;Pauschal\&quot; in German. This means the position total price that is being invoiced should not depend on the actual quantity. In practice, partial invoices are still often used for partial services rendered. This property does not affect the price calculation for this position. Typically, the Quantity should be set to 1.0 when this flag is used. | 
**RepetitionTo** | **string** | This identifier can be used to point to the Id of a position in the same ServiceSpecification that acts as a base position. It matches \&quot;Bezugsposition\&quot; in GAEB. This can be used for positions that repeat partially or are linked together | [optional] 
**ProjectCatalogues** | [**[]CatalogueDto**](CatalogueDto.md) | These are Catalogues that are used within this Position. Catalogues are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. They are propagate to all child elements, e.g. other containers and QuantityAssignments. In the context of a ServiceSpecification, all elements share the same instance of the collection. | [optional] 
**CatalogueReferences** | [**[]CatalogueReferenceDto**](CatalogueReferenceDto.md) | Referenced catalogues for this Position. | [optional] 
**Type** | **string** |  | [optional] 
**StandardizedDescription** | [**StandardizedDescriptionDto**](StandardizedDescriptionDto.md) |  | [optional] 
**ComplementedByQuantities** | [**[]ComplementedByQuantityDto**](ComplementedByQuantityDto.md) | This list contains references to positions that complement this one, additionally also specifying a quantity for which the addition is intended. This does not replace the ComplementedBy property and there are no automatic checks being done between these two properties, so it&#39;s up to the user code to ensure deletions (and additions, if desired) are performed for both properties. When copying withing keeping Ids, this list will not be part of the copy process, since it would only contain quantities without actual position references. Containers, however, will rebuild the list with updated position references when copying positions that contain entries here. | [optional] 
**ExecutionDescriptionReference** | **string** | This identifier can be used to point to the Id of an ExecutionDescription in the same ServiceSpecification. ExecutionDescriptions act as a way to centrally describe how positions should be executed in practice. Often, the position itself still has text of its own to highlight deviations from that or add more details. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


