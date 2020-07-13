# \OenormConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OenormConversionConvertToAva**](OenormConversionApi.md#OenormConversionConvertToAva) | **Post** /conversion/oenorm/ava | Converts ÖNorm files to Dangl.AVA projects
[**OenormConversionConvertToExcel**](OenormConversionApi.md#OenormConversionConvertToExcel) | **Post** /conversion/oenorm/excel | Converts ÖNorm files to Excel
[**OenormConversionConvertToGaeb**](OenormConversionApi.md#OenormConversionConvertToGaeb) | **Post** /conversion/oenorm/gaeb | Converts ÖNorm files to GAEB files.
[**OenormConversionConvertToOenorm**](OenormConversionApi.md#OenormConversionConvertToOenorm) | **Post** /conversion/oenorm/oenorm | Converts ÖNorm files to Oenorm files.



## OenormConversionConvertToAva

> ProjectDto OenormConversionConvertToAva(ctx, optional)

Converts ÖNorm files to Dangl.AVA projects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OenormConversionConvertToAvaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OenormConversionConvertToAvaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removePlainTextLongTexts** | **optional.Bool**| If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **optional.Bool**| If set to true, html long texts will be removed from the output to reduce response sizes | 
 **oenormFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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


## OenormConversionConvertToExcel

> *os.File OenormConversionConvertToExcel(ctx, optional)

Converts ÖNorm files to Excel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OenormConversionConvertToExcelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OenormConversionConvertToExcelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writePrices** | **optional.Bool**| Defaults to true | 
 **writeLongTexts** | **optional.Bool**| Defaults to true | 
 **conversionCulture** | **optional.String**| The culture that should be used for the conversion process, to have localized Excel files | 
 **oenormFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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


## OenormConversionConvertToGaeb

> *os.File OenormConversionConvertToGaeb(ctx, optional)

Converts ÖNorm files to GAEB files.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OenormConversionConvertToGaebOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OenormConversionConvertToGaebOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationGaebType** | **optional.String**| Defaults to GAEB XML V3.2 | 
 **targetExchangePhaseTransform** | **optional.String**| Defaults to none, meaning no transformation will be done | 
 **oenormFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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


## OenormConversionConvertToOenorm

> *os.File OenormConversionConvertToOenorm(ctx, optional)

Converts ÖNorm files to Oenorm files.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OenormConversionConvertToOenormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OenormConversionConvertToOenormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationOenormType** | **optional.String**| Defaults to Lv2015 | 
 **tryRepairProjectStructure** | **optional.Bool**| Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 
 **oenormFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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

