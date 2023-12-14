# \DanglIdentityApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DanglIdentityGetUserInfo**](DanglIdentityApi.md#DanglIdentityGetUserInfo) | **Get** /identity/user-info | 
[**DanglIdentityLoginAndReturnToken**](DanglIdentityApi.md#DanglIdentityLoginAndReturnToken) | **Post** /identity/token-login | 
[**DanglIdentityLoginWithCookie**](DanglIdentityApi.md#DanglIdentityLoginWithCookie) | **Post** /identity/login | 
[**DanglIdentityRefreshToken**](DanglIdentityApi.md#DanglIdentityRefreshToken) | **Post** /identity/token-refresh | 
[**DanglIdentityRegister**](DanglIdentityApi.md#DanglIdentityRegister) | **Post** /identity/register | 
[**DanglIdentityRequestPasswordReset**](DanglIdentityApi.md#DanglIdentityRequestPasswordReset) | **Post** /identity/password-forgotten | 
[**DanglIdentitySignOutWithSignInManager**](DanglIdentityApi.md#DanglIdentitySignOutWithSignInManager) | **Delete** /identity/login | 



## DanglIdentityGetUserInfo

> UserInfoGet DanglIdentityGetUserInfo(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prolicht-dev/avaclient-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DanglIdentityApi.DanglIdentityGetUserInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DanglIdentityApi.DanglIdentityGetUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DanglIdentityGetUserInfo`: UserInfoGet
    fmt.Fprintf(os.Stdout, "Response from `DanglIdentityApi.DanglIdentityGetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDanglIdentityGetUserInfoRequest struct via the builder pattern


### Return type

[**UserInfoGet**](UserInfoGet.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentityLoginAndReturnToken

> TokenResponseGet DanglIdentityLoginAndReturnToken(ctx).Model(model).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prolicht-dev/avaclient-go"
)

func main() {
    model := *openapiclient.NewTokenLoginPost("Identifier_example", "Password_example") // TokenLoginPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DanglIdentityApi.DanglIdentityLoginAndReturnToken(context.Background()).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DanglIdentityApi.DanglIdentityLoginAndReturnToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DanglIdentityLoginAndReturnToken`: TokenResponseGet
    fmt.Fprintf(os.Stdout, "Response from `DanglIdentityApi.DanglIdentityLoginAndReturnToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDanglIdentityLoginAndReturnTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**TokenLoginPost**](TokenLoginPost.md) |  | 

### Return type

[**TokenResponseGet**](TokenResponseGet.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentityLoginWithCookie

> DanglIdentityLoginWithCookie(ctx).Model(model).RedirectUrl(redirectUrl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prolicht-dev/avaclient-go"
)

func main() {
    model := *openapiclient.NewLoginPost("Identifier_example", "Password_example", false) // LoginPost | 
    redirectUrl := "redirectUrl_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DanglIdentityApi.DanglIdentityLoginWithCookie(context.Background()).Model(model).RedirectUrl(redirectUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DanglIdentityApi.DanglIdentityLoginWithCookie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDanglIdentityLoginWithCookieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**LoginPost**](LoginPost.md) |  | 
 **redirectUrl** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentityRefreshToken

> TokenResponseGet DanglIdentityRefreshToken(ctx).Model(model).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prolicht-dev/avaclient-go"
)

func main() {
    model := *openapiclient.NewTokenRefreshPost("RefreshToken_example") // TokenRefreshPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DanglIdentityApi.DanglIdentityRefreshToken(context.Background()).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DanglIdentityApi.DanglIdentityRefreshToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DanglIdentityRefreshToken`: TokenResponseGet
    fmt.Fprintf(os.Stdout, "Response from `DanglIdentityApi.DanglIdentityRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDanglIdentityRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**TokenRefreshPost**](TokenRefreshPost.md) |  | 

### Return type

[**TokenResponseGet**](TokenResponseGet.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentityRegister

> DanglIdentityRegister(ctx).RegisterModel(registerModel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prolicht-dev/avaclient-go"
)

func main() {
    registerModel := *openapiclient.NewRegisterPost("Username_example", "Email_example", "Password_example") // RegisterPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DanglIdentityApi.DanglIdentityRegister(context.Background()).RegisterModel(registerModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DanglIdentityApi.DanglIdentityRegister``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDanglIdentityRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerModel** | [**RegisterPost**](RegisterPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentityRequestPasswordReset

> DanglIdentityRequestPasswordReset(ctx).ForgotPasswordModel(forgotPasswordModel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prolicht-dev/avaclient-go"
)

func main() {
    forgotPasswordModel := *openapiclient.NewForgotPasswordPost("Identifier_example") // ForgotPasswordPost | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DanglIdentityApi.DanglIdentityRequestPasswordReset(context.Background()).ForgotPasswordModel(forgotPasswordModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DanglIdentityApi.DanglIdentityRequestPasswordReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDanglIdentityRequestPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordModel** | [**ForgotPasswordPost**](ForgotPasswordPost.md) |  | 

### Return type

 (empty response body)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DanglIdentitySignOutWithSignInManager

> DanglIdentitySignOutWithSignInManager(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prolicht-dev/avaclient-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DanglIdentityApi.DanglIdentitySignOutWithSignInManager(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DanglIdentityApi.DanglIdentitySignOutWithSignInManager``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDanglIdentitySignOutWithSignInManagerRequest struct via the builder pattern


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

