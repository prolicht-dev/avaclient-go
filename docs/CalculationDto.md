# CalculationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Descriptive text for this calculation. | [optional] 
**Formula** | **string** | This Calculation&#39;s mathematical expression. Please note that thousands separators are not supported. Both comma and point will be treated as decimal separators. | [optional] 
**Result** | **float64** | The calculated result from the formula, 0 if invalid. | [readonly] 
**Valid** | **bool** | Whether the Formula is a valid expression. | [readonly] 
**ErrorPositionInLine** | **int32** | Will be -1 if the Formula is correct, else it will show the position in the formula where an error was encountered. This is a zero based index. | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


