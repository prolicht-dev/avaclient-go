# \UglConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UglConversionConvertToAva**](UglConversionApi.md#UglConversionConvertToAva) | **Post** /conversion/ugl/ava | Converts Ugl files to Dangl.AVA projects



## UglConversionConvertToAva

> ProjectDto UglConversionConvertToAva(ctx).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).UglFile(uglFile).Execute()

Converts Ugl files to Dangl.AVA projects

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
    removePlainTextLongTexts := true // bool | If set to true, plain text long texts will be removed from the output to reduce response sizes (optional)
    removeHtmlLongTexts := true // bool | If set to true, html long texts will be removed from the output to reduce response sizes (optional)
    uglFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UglConversionApi.UglConversionConvertToAva(context.Background()).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).UglFile(uglFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UglConversionApi.UglConversionConvertToAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UglConversionConvertToAva`: ProjectDto
    fmt.Fprintf(os.Stdout, "Response from `UglConversionApi.UglConversionConvertToAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUglConversionConvertToAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removePlainTextLongTexts** | **bool** | If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **bool** | If set to true, html long texts will be removed from the output to reduce response sizes | 
 **uglFile** | ***os.File** | The input file | 

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

