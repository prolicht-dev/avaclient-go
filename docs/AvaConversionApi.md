# \AvaConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvaConversionConvertToAva**](AvaConversionApi.md#AvaConversionConvertToAva) | **Post** /conversion/ava/ava | Converts Dangl.AVA projects to Dangl.AVA. This is useful when you want to generate the calculated properties.
[**AvaConversionConvertToExcel**](AvaConversionApi.md#AvaConversionConvertToExcel) | **Post** /conversion/ava/excel | Converts Dangl.AVA projects to Excel
[**AvaConversionConvertToGaeb**](AvaConversionApi.md#AvaConversionConvertToGaeb) | **Post** /conversion/ava/gaeb | Converts Dangl.AVA projects to GAEB
[**AvaConversionConvertToOenorm**](AvaConversionApi.md#AvaConversionConvertToOenorm) | **Post** /conversion/ava/oenorm | Converts Dangl.AVA projects to Oenorm
[**AvaConversionConvertToReb**](AvaConversionApi.md#AvaConversionConvertToReb) | **Post** /conversion/ava/reb | Converts Dangl.AVA projects to REB
[**AvaConversionConvertToSia**](AvaConversionApi.md#AvaConversionConvertToSia) | **Post** /conversion/ava/sia | Converts Dangl.AVA projects to SIA 451



## AvaConversionConvertToAva

> ProjectDto AvaConversionConvertToAva(ctx, avaProject, optional)

Converts Dangl.AVA projects to Dangl.AVA. This is useful when you want to generate the calculated properties.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avaProject** | [**ProjectDto**](ProjectDto.md)| The Dangl.AVA project | 
 **optional** | ***AvaConversionConvertToAvaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AvaConversionConvertToAvaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removePlainTextLongTexts** | **optional.Bool**| If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **optional.Bool**| If set to true, html long texts will be removed from the output to reduce response sizes | 

### Return type

[**ProjectDto**](ProjectDto.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: application/vnd.com.dangl-it.ProjectDto.v1+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvaConversionConvertToExcel

> *os.File AvaConversionConvertToExcel(ctx, avaProject, optional)

Converts Dangl.AVA projects to Excel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avaProject** | [**ProjectDto**](ProjectDto.md)| The Dangl.AVA project | 
 **optional** | ***AvaConversionConvertToExcelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AvaConversionConvertToExcelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **writePrices** | **optional.Bool**| Defaults to true | 
 **writeLongTexts** | **optional.Bool**| Defaults to true | 
 **conversionCulture** | **optional.String**| The culture that should be used for the conversion process, to have localized Excel files | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvaConversionConvertToGaeb

> *os.File AvaConversionConvertToGaeb(ctx, avaProject, optional)

Converts Dangl.AVA projects to GAEB

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avaProject** | [**ProjectDto**](ProjectDto.md)| The Dangl.AVA project | 
 **optional** | ***AvaConversionConvertToGaebOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AvaConversionConvertToGaebOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationGaebType** | **optional.String**| Defaults to GAEB XML V3.2 | 
 **targetExchangePhaseTransform** | **optional.String**| Defaults to none, meaning no transformation will be done | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvaConversionConvertToOenorm

> *os.File AvaConversionConvertToOenorm(ctx, avaProject, optional)

Converts Dangl.AVA projects to Oenorm

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avaProject** | [**ProjectDto**](ProjectDto.md)| The Dangl.AVA project | 
 **optional** | ***AvaConversionConvertToOenormOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AvaConversionConvertToOenormOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationOenormType** | **optional.String**| Defaults to Lv2015 | 
 **tryRepairProjectStructure** | **optional.Bool**| Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvaConversionConvertToReb

> *os.File AvaConversionConvertToReb(ctx, avaProject)

Converts Dangl.AVA projects to REB

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avaProject** | [**ProjectDto**](ProjectDto.md)| The Dangl.AVA project | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvaConversionConvertToSia

> *os.File AvaConversionConvertToSia(ctx, avaProject)

Converts Dangl.AVA projects to SIA 451

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avaProject** | [**ProjectDto**](ProjectDto.md)| The Dangl.AVA project | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

