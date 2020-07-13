# \ExcelConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExcelConversionConvertToAva**](ExcelConversionApi.md#ExcelConversionConvertToAva) | **Post** /conversion/excel/ava | Converts Excel files to Dangl.AVA projects.
[**ExcelConversionConvertToExcel**](ExcelConversionApi.md#ExcelConversionConvertToExcel) | **Post** /conversion/excel/excel | Converts Excel files to Excel files. Used, for example, when elements were added in excel to generate or modify a project. The Excel file can then be shared containing the full project with all formattings, formulas and styles applied.
[**ExcelConversionConvertToGaeb**](ExcelConversionApi.md#ExcelConversionConvertToGaeb) | **Post** /conversion/excel/gaeb | Converts Excel files to GAEB files.
[**ExcelConversionConvertToOenorm**](ExcelConversionApi.md#ExcelConversionConvertToOenorm) | **Post** /conversion/excel/oenorm | Converts Excel files to Oenorm files.



## ExcelConversionConvertToAva

> ProjectDto ExcelConversionConvertToAva(ctx, optional)

Converts Excel files to Dangl.AVA projects.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExcelConversionConvertToAvaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExcelConversionConvertToAvaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNewElements** | **optional.Bool**| Defaults to false | 
 **rebuildItemNumberSchema** | **optional.Bool**| When importing new elements from Excel, sometimes the ItemNumberSchema in the file is not in compliance with the GAEB requirements. Enabling this option tries to repair the ItemNumberSchema. Defaults to false. | 
 **removePlainTextLongTexts** | **optional.Bool**| If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **optional.Bool**| If set to true, html long texts will be removed from the output to reduce response sizes | 
 **excelFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

### Return type

[**ProjectDto**](ProjectDto.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/vnd.com.dangl-it.ProjectDto.v1+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExcelConversionConvertToExcel

> *os.File ExcelConversionConvertToExcel(ctx, optional)

Converts Excel files to Excel files. Used, for example, when elements were added in excel to generate or modify a project. The Excel file can then be shared containing the full project with all formattings, formulas and styles applied.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExcelConversionConvertToExcelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExcelConversionConvertToExcelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNewElements** | **optional.Bool**| Defaults to false | 
 **rebuildItemNumberSchema** | **optional.Bool**| When importing new elements from Excel, sometimes the ItemNumberSchema in the file is not in compliance with the GAEB requirements. Enabling this option tries to repair the ItemNumberSchema. Defaults to false. | 
 **writePrices** | **optional.Bool**| Defaults to true | 
 **writeLongTexts** | **optional.Bool**| Defaults to true | 
 **conversionCulture** | **optional.String**| The culture that should be used for the conversion process, to have localized Excel files | 
 **excelFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExcelConversionConvertToGaeb

> *os.File ExcelConversionConvertToGaeb(ctx, optional)

Converts Excel files to GAEB files.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExcelConversionConvertToGaebOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExcelConversionConvertToGaebOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNewElements** | **optional.Bool**| Defaults to false | 
 **rebuildItemNumberSchema** | **optional.Bool**| When importing new elements from Excel, sometimes the ItemNumberSchema in the file is not in compliance with the GAEB requirements. Enabling this option tries to repair the ItemNumberSchema. Defaults to false. | 
 **destinationGaebType** | **optional.String**| Defaults to GAEB XML V3.2 | 
 **targetExchangePhaseTransform** | **optional.String**| Defaults to none, meaning no transformation will be done | 
 **excelFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExcelConversionConvertToOenorm

> *os.File ExcelConversionConvertToOenorm(ctx, optional)

Converts Excel files to Oenorm files.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExcelConversionConvertToOenormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExcelConversionConvertToOenormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readNewElements** | **optional.Bool**| Defaults to false | 
 **rebuildItemNumberSchema** | **optional.Bool**| When importing new elements from Excel, sometimes the ItemNumberSchema in the file is not in compliance with the GAEB requirements. Enabling this option tries to repair the ItemNumberSchema. Defaults to false. | 
 **destinationOenormType** | **optional.String**| Defaults to Lv2015 | 
 **tryRepairProjectStructure** | **optional.Bool**| Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 
 **excelFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

