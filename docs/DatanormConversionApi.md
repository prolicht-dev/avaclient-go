# \DatanormConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DatanormConversionConvertToAva**](DatanormConversionApi.md#DatanormConversionConvertToAva) | **Post** /conversion/datanorm/ava | Converts Datanorm files to Dangl.AVA projects
[**DatanormConversionConvertToFlatAva**](DatanormConversionApi.md#DatanormConversionConvertToFlatAva) | **Post** /conversion/datanorm/flat-ava | Converts Datanorm files to Dangl.AVA projects



## DatanormConversionConvertToAva

> ProjectDto DatanormConversionConvertToAva(ctx).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).DatanormFile(datanormFile).Execute()

Converts Datanorm files to Dangl.AVA projects

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
    datanormFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatanormConversionApi.DatanormConversionConvertToAva(context.Background()).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).DatanormFile(datanormFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatanormConversionApi.DatanormConversionConvertToAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatanormConversionConvertToAva`: ProjectDto
    fmt.Fprintf(os.Stdout, "Response from `DatanormConversionApi.DatanormConversionConvertToAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatanormConversionConvertToAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removePlainTextLongTexts** | **bool** | If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **bool** | If set to true, html long texts will be removed from the output to reduce response sizes | 
 **datanormFile** | ***os.File** | The input file | 

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


## DatanormConversionConvertToFlatAva

> FlatAvaProject DatanormConversionConvertToFlatAva(ctx).DatanormFile(datanormFile).Execute()

Converts Datanorm files to Dangl.AVA projects

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
    datanormFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DatanormConversionApi.DatanormConversionConvertToFlatAva(context.Background()).DatanormFile(datanormFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatanormConversionApi.DatanormConversionConvertToFlatAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DatanormConversionConvertToFlatAva`: FlatAvaProject
    fmt.Fprintf(os.Stdout, "Response from `DatanormConversionApi.DatanormConversionConvertToFlatAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatanormConversionConvertToFlatAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datanormFile** | ***os.File** | The input file | 

### Return type

[**FlatAvaProject**](FlatAvaProject.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

