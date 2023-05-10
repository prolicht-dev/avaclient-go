# \StatusApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusGetStatus**](StatusApi.md#StatusGetStatus) | **Get** /status | Reports the health status of the AVACloud API



## StatusGetStatus

> GetStatus StatusGetStatus(ctx).Execute()

Reports the health status of the AVACloud API

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
    resp, r, err := apiClient.StatusApi.StatusGetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusApi.StatusGetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusGetStatus`: GetStatus
    fmt.Fprintf(os.Stdout, "Response from `StatusApi.StatusGetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusGetStatusRequest struct via the builder pattern


### Return type

[**GetStatus**](GetStatus.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

