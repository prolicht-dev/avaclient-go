# \GaebConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GaebConversionConvertToAva**](GaebConversionApi.md#GaebConversionConvertToAva) | **Post** /conversion/gaeb/ava | Converts GAEB files to Dangl.AVA projects
[**GaebConversionConvertToExcel**](GaebConversionApi.md#GaebConversionConvertToExcel) | **Post** /conversion/gaeb/excel | Converts GAEB files to Excel
[**GaebConversionConvertToGaeb**](GaebConversionApi.md#GaebConversionConvertToGaeb) | **Post** /conversion/gaeb/gaeb | Converts GAEB files to GAEB files. Used for example when transforming or repairing GAEB files.
[**GaebConversionConvertToOenorm**](GaebConversionApi.md#GaebConversionConvertToOenorm) | **Post** /conversion/gaeb/oenorm | Converts GAEB files to Oenorm files



## GaebConversionConvertToAva

> ProjectDto GaebConversionConvertToAva(ctx, optional)

Converts GAEB files to Dangl.AVA projects

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GaebConversionConvertToAvaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GaebConversionConvertToAvaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removePlainTextLongTexts** | **optional.Bool**| If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **optional.Bool**| If set to true, html long texts will be removed from the output to reduce response sizes | 
 **gaebFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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


## GaebConversionConvertToExcel

> *os.File GaebConversionConvertToExcel(ctx, optional)

Converts GAEB files to Excel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GaebConversionConvertToExcelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GaebConversionConvertToExcelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writePrices** | **optional.Bool**| Defaults to true | 
 **writeLongTexts** | **optional.Bool**| Defaults to true | 
 **conversionCulture** | **optional.String**| The culture that should be used for the conversion process, to have localized Excel files | 
 **gaebFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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


## GaebConversionConvertToGaeb

> *os.File GaebConversionConvertToGaeb(ctx, optional)

Converts GAEB files to GAEB files. Used for example when transforming or repairing GAEB files.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GaebConversionConvertToGaebOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GaebConversionConvertToGaebOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationGaebType** | **optional.String**| Defaults to GAEB XML V3.2 | 
 **targetExchangePhaseTransform** | **optional.String**| Defaults to none, meaning no transformation will be done | 
 **gaebFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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


## GaebConversionConvertToOenorm

> *os.File GaebConversionConvertToOenorm(ctx, optional)

Converts GAEB files to Oenorm files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GaebConversionConvertToOenormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GaebConversionConvertToOenormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationOenormType** | **optional.String**| Defaults to Lv2015 | 
 **tryRepairProjectStructure** | **optional.Bool**| Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 
 **gaebFile** | **optional.Interface of *os.File****optional.*os.File**| The input file | 

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

