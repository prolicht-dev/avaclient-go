# \ValidationApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidationValidateFile**](ValidationApi.md#ValidationValidateFile) | **Post** /validation/file | This endpoint validates AVA files, typically GAEB or ÖNorm. The type of file needs to be provided as a query parameter, since there is no auto detection of the uploaded file type.
[**ValidationValidateProject**](ValidationApi.md#ValidationValidateProject) | **Post** /validation/project | This endpoint provides a full validation of a provided ProjectDto. It will take the given exchange phase into account and do some general project validation. Optionally, a conversion to a desired target can also be done, in which case the target file will also be validated.



## ValidationValidateFile

> ValidationResult ValidationValidateFile(ctx).FileValidationSourceType(fileValidationSourceType).AvaFile(avaFile).Execute()

This endpoint validates AVA files, typically GAEB or ÖNorm. The type of file needs to be provided as a query parameter, since there is no auto detection of the uploaded file type.

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
    fileValidationSourceType := "fileValidationSourceType_example" // string | You need to indicate which type of file is being provided, there is no auto detection mechanism (optional)
    avaFile := os.NewFile(1234, "some_file") // *os.File | The file to validate (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ValidationApi.ValidationValidateFile(context.Background()).FileValidationSourceType(fileValidationSourceType).AvaFile(avaFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidationApi.ValidationValidateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidationValidateFile`: ValidationResult
    fmt.Fprintf(os.Stdout, "Response from `ValidationApi.ValidationValidateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidationValidateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileValidationSourceType** | **string** | You need to indicate which type of file is being provided, there is no auto detection mechanism | 
 **avaFile** | ***os.File** | The file to validate | 

### Return type

[**ValidationResult**](ValidationResult.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidationValidateProject

> ValidationResult ValidationValidateProject(ctx).AvaProjectValidationSourceOptions(avaProjectValidationSourceOptions).Execute()

This endpoint provides a full validation of a provided ProjectDto. It will take the given exchange phase into account and do some general project validation. Optionally, a conversion to a desired target can also be done, in which case the target file will also be validated.

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
    avaProjectValidationSourceOptions := *openapiclient.NewPostAvaProjectValidationSourceOptions(*openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")), openapiclient.ValidationType("Project")) // PostAvaProjectValidationSourceOptions | The options used for the validation operation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ValidationApi.ValidationValidateProject(context.Background()).AvaProjectValidationSourceOptions(avaProjectValidationSourceOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidationApi.ValidationValidateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidationValidateProject`: ValidationResult
    fmt.Fprintf(os.Stdout, "Response from `ValidationApi.ValidationValidateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidationValidateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProjectValidationSourceOptions** | [**PostAvaProjectValidationSourceOptions**](PostAvaProjectValidationSourceOptions.md) | The options used for the validation operation | 

### Return type

[**ValidationResult**](ValidationResult.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

