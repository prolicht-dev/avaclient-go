# SubDescriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Quantity** | **float32** | Returns the total calculated sum of all quantity assignments. Will return the result rounded to three decimal places. | [readonly] 
**QuantityComponents** | [**[]CalculationDto**](CalculationDto.md) | Holds quantity information for this sub description. Quantity is listening to changes here and is reporting the total sum of all quantity components. | [optional] 
**AmountToBeEnteredByBidder** | **bool** | Indicates if the bidder is asked to specify an amount. | 
**Identifier** | **string** | Identifier for this SubDescription. | [optional] 
**ShortText** | **string** | Short description for this DescriptionBase element. | [optional] 
**LongText** | **string** | Detailed description for this DescriptionBase element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText. | [optional] 
**HtmlLongText** | **string** | This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out. | [optional] 
**AdditionType** | [**AdditionTypeDto**](AdditionTypeDto.md) |  | 
**StandardizedDescription** | [**StandardizedDescriptionDto**](StandardizedDescriptionDto.md) |  | [optional] 
**ExecutionDescriptionReference** | **string** | This identifier can be used to point to the Id of an ExecutionDescription in the same ServiceSpecification. ExecutionDescriptions act as a way to centrally describe how positions (or sub descriptions) should be executed in practice. Often, the position (or sub description) itself still has text of its own to highlight deviations from that or add more details. When working with import and export features, this property is only supported in GAEB 90 data exchange. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


