# ProjectDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**PriceAccuracy** | **int32** | This property controls the accuracy of all price properties, meaning how many decimal places are preserved in calculations. It defaults to DEFAULT_PRICE_ACCURACY. Please see the Dangl.AVA documentation for further information about decimal precision. | 
**ForceStrictTotals** | **bool** | This forces total prices to be the strict product of quantities times unit price in positions. It is enabled by default. If this is disabled, both the unit price and the total price of positions is calculated from the non-rounded values. Please see the documentation for a more detailed explanation of this setting. | 
**PriceRoundingMode** | [**PriceRoundingModeDto**](PriceRoundingModeDto.md) |  | 
**ProjectInformation** | [**ProjectInformationDto**](ProjectInformationDto.md) |  | [optional] 
**ServiceSpecifications** | [**[]ServiceSpecificationDto**](ServiceSpecificationDto.md) | The ServiceSpecifications in this Project. | [optional] 
**GaebXmlId** | **string** | This is used to store the GAEB XML Id within this Project. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


