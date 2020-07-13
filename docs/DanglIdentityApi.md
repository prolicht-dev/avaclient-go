# \DanglIdentityApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DanglIdentityLoginAndReturnToken**](DanglIdentityApi.md#DanglIdentityLoginAndReturnToken) | **Post** /identity/token-login | 
[**DanglIdentityLoginWithCookie**](DanglIdentityApi.md#DanglIdentityLoginWithCookie) | **Post** /identity/login | 
[**DanglIdentityRefreshToken**](DanglIdentityApi.md#DanglIdentityRefreshToken) | **Post** /identity/token-refresh | 
[**DanglIdentityRegister**](DanglIdentityApi.md#DanglIdentityRegister) | **Post** /identity/register | 
[**DanglIdentityRequestPasswordReset**](DanglIdentityApi.md#DanglIdentityRequestPasswordReset) | **Post** /identity/password-forgotten | 
[**DanglIdentitySignOutWithSignInManager**](DanglIdentityApi.md#DanglIdentitySignOutWithSignInManager) | **Delete** /identity/login | 



## DanglIdentityLoginAndReturnToken

> TokenResponseGet DanglIdentityLoginAndReturnToken(ctx, model)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | [**TokenLoginPost**](TokenLoginPost.md)|  | 

### Return type

[**TokenResponseGet**](TokenResponseGet.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentityLoginWithCookie

> DanglIdentityLoginWithCookie(ctx, model, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | [**LoginPost**](LoginPost.md)|  | 
 **optional** | ***DanglIdentityLoginWithCookieOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DanglIdentityLoginWithCookieOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redirectUrl** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentityRefreshToken

> TokenResponseGet DanglIdentityRefreshToken(ctx, model)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | [**TokenRefreshPost**](TokenRefreshPost.md)|  | 

### Return type

[**TokenResponseGet**](TokenResponseGet.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentityRegister

> DanglIdentityRegister(ctx, registerModel)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registerModel** | [**RegisterPost**](RegisterPost.md)|  | 

### Return type

 (empty response body)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentityRequestPasswordReset

> DanglIdentityRequestPasswordReset(ctx, forgotPasswordModel)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forgotPasswordModel** | [**ForgotPasswordPost**](ForgotPasswordPost.md)|  | 

### Return type

 (empty response body)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentitySignOutWithSignInManager

> DanglIdentitySignOutWithSignInManager(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

