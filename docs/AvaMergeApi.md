# \AvaMergeApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvaMergeConvertToAva**](AvaMergeApi.md#AvaMergeConvertToAva) | **Post** /ava-merge | This will merge two AVA projects into one. Typically this is used for merging offers into a base project, which preserves the base projects structure and adds prices and text additions from the offer project. This is currently in a preview phase, and the endpoint and its models are subject to change in the next versions of AVACloud.



## AvaMergeConvertToAva

> GetProjectMergeResult AvaMergeConvertToAva(ctx).MergeConfiguration(mergeConfiguration).Execute()

This will merge two AVA projects into one. Typically this is used for merging offers into a base project, which preserves the base projects structure and adds prices and text additions from the offer project. This is currently in a preview phase, and the endpoint and its models are subject to change in the next versions of AVACloud.

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
    mergeConfiguration := *openapiclient.NewPostAvaProjectMerge() // PostAvaProjectMerge | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvaMergeApi.AvaMergeConvertToAva(context.Background()).MergeConfiguration(mergeConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaMergeApi.AvaMergeConvertToAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaMergeConvertToAva`: GetProjectMergeResult
    fmt.Fprintf(os.Stdout, "Response from `AvaMergeApi.AvaMergeConvertToAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaMergeConvertToAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mergeConfiguration** | [**PostAvaProjectMerge**](PostAvaProjectMerge.md) |  | 

### Return type

[**GetProjectMergeResult**](GetProjectMergeResult.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/vnd.com.dangl-it.ProjectDto.v1+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

