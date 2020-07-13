# PriceInformationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**HourlyWage** | **float32** | The amount of currency per one hour of manual labour work in this project. | 
**DeductionFactor** | **float32** | The final, total price will be deducted by this rate. | 
**FlatSum** | **float32** | This is given when there is only one flat price for the whole service specification. | 
**TaxRate** | **float32** | Global tax rate for the project. Note that certain elements may have a different, specific tax rate. | 
**TradeDiscounts** | [**[]TradeDiscountDto**](TradeDiscountDto.md) | Trade discounts for offered in this ServiceSpecification. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


