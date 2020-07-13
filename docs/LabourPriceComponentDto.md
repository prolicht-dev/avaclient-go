# LabourPriceComponentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The label associated with this price component. Will be taken from the parent Projects ProjectInformation. | [optional] 
**Price** | **float32** | The total, calculated price of this component. Will multiply the calculated amount of hours with the ServiceSpecifications hourly wage rate. | [readonly] 
**HourlyWage** | **float32** | The cost per hour of manual labor. | 
**Values** | [**[]CalculationDto**](CalculationDto.md) | The single Calculation elements this price component is composed of. | [optional] 
**UseOwnHourlyWage** | **bool** | Indicates if the ServiceSpecification&#39;s standard HourlyWage is to be used or a custom value. | 
**TotalTime** | **float32** | The total, calculated time of this component. Will return the result rounded to three decimal places. | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


