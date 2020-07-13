# \RebConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RebConversionConvertToAva**](RebConversionApi.md#RebConversionConvertToAva) | **Post** /conversion/reb/ava | Converts REB files to Dangl.AVA projects
[**RebConversionConvertToExcel**](RebConversionApi.md#RebConversionConvertToExcel) | **Post** /conversion/reb/excel | Converts REB files to Excel
[**RebConversionConvertToGaeb**](RebConversionApi.md#RebConversionConvertToGaeb) | **Post** /conversion/reb/gaeb | Converts REB files to GAEB files
[**RebConversionConvertToOenorm**](RebConversionApi.md#RebConversionConvertToOenorm) | **Post** /conversion/reb/oenorm | Converts REB files to Oenorm



## RebConversionConvertToAva

> ProjectDto RebConversionConvertToAva(ctx, optional)

Converts REB files to Dangl.AVA projects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RebConversionConvertToAvaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RebConversionConvertToAvaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removePlainTextLongTexts** | **optional.Bool**| If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **optional.Bool**| If set to true, html long texts will be removed from the output to reduce response sizes | 
 **rebFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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


## RebConversionConvertToExcel

> *os.File RebConversionConvertToExcel(ctx, optional)

Converts REB files to Excel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RebConversionConvertToExcelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RebConversionConvertToExcelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writePrices** | **optional.Bool**| Defaults to true | 
 **writeLongTexts** | **optional.Bool**| Defaults to true | 
 **conversionCulture** | **optional.String**| The culture that should be used for the conversion process, to have localized Excel files | 
 **rebFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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


## RebConversionConvertToGaeb

> *os.File RebConversionConvertToGaeb(ctx, optional)

Converts REB files to GAEB files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RebConversionConvertToGaebOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RebConversionConvertToGaebOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationGaebType** | **optional.String**| Defaults to GAEB XML V3.2 | 
 **targetExchangePhaseTransform** | **optional.String**| Defaults to none, meaning no transformation will be done | 
 **rebFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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


## RebConversionConvertToOenorm

> *os.File RebConversionConvertToOenorm(ctx, optional)

Converts REB files to Oenorm

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RebConversionConvertToOenormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RebConversionConvertToOenormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationOenormType** | **optional.String**| Defaults to Lv2015 | 
 **tryRepairProjectStructure** | **optional.Bool**| Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 
 **rebFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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

