# \XRechnungConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XRechnungConversionConvertAvaToPdfInvoice**](XRechnungConversionApi.md#XRechnungConversionConvertAvaToPdfInvoice) | **Post** /conversion/xrechnung/ava-wrapper/pdf | This will convert an AVA wrapper object to a PDF file
[**XRechnungConversionConvertAvaToXRechnung**](XRechnungConversionApi.md#XRechnungConversionConvertAvaToXRechnung) | **Post** /conversion/xrechnung/ava-wrapper/xrechnung | This will convert an AVA wrapper object to an XRechnung file
[**XRechnungConversionConvertInvoiceToPdfInvoice**](XRechnungConversionApi.md#XRechnungConversionConvertInvoiceToPdfInvoice) | **Post** /conversion/xrechnung/invoice/pdf | This will convert an Invoice object to a PDF file
[**XRechnungConversionConvertInvoiceToXRechnung**](XRechnungConversionApi.md#XRechnungConversionConvertInvoiceToXRechnung) | **Post** /conversion/xrechnung/invoice/xrechnung | This will convert an Invoice object to an XRechnung file
[**XRechnungConversionConvertXRechnungToAva**](XRechnungConversionApi.md#XRechnungConversionConvertXRechnungToAva) | **Post** /conversion/xrechnung/ava-wrapper | This will read an XRechnung file and convert it to an AVA wrapper object
[**XRechnungConversionConvertXRechnungToInvoice**](XRechnungConversionApi.md#XRechnungConversionConvertXRechnungToInvoice) | **Post** /conversion/xrechnung/invoice | This will read an XRechnung file and convert it to an Invoice object
[**XRechnungConversionMergeInvoiceIntoPdfFile**](XRechnungConversionApi.md#XRechnungConversionMergeInvoiceIntoPdfFile) | **Post** /conversion/xrechnung/merge | This will embedd an XRechnung file into an existing PDF file as an attachment. If Factur-X is chosen as target format, AVACloud will also insert the required Factur-X XML metadata into the PDF file. AVACloud does not, however, ensure the PDF itself is meeting all the Factur-X requirements.



## XRechnungConversionConvertAvaToPdfInvoice

> *os.File XRechnungConversionConvertAvaToPdfInvoice(ctx).AvaWrapper(avaWrapper).Execute()

This will convert an AVA wrapper object to a PDF file

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
    avaWrapper := *openapiclient.NewAvaProjectWrapper(*openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")), openapiclient.InvoiceType("PartialInvoice"), false, float32(123)) // AvaProjectWrapper | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XRechnungConversionApi.XRechnungConversionConvertAvaToPdfInvoice(context.Background()).AvaWrapper(avaWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRechnungConversionApi.XRechnungConversionConvertAvaToPdfInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XRechnungConversionConvertAvaToPdfInvoice`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `XRechnungConversionApi.XRechnungConversionConvertAvaToPdfInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXRechnungConversionConvertAvaToPdfInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaWrapper** | [**AvaProjectWrapper**](AvaProjectWrapper.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XRechnungConversionConvertAvaToXRechnung

> *os.File XRechnungConversionConvertAvaToXRechnung(ctx).AvaWrapper(avaWrapper).TargetFormat(targetFormat).DisablePriceRounding(disablePriceRounding).DisableBrandingComment(disableBrandingComment).Execute()

This will convert an AVA wrapper object to an XRechnung file

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
    avaWrapper := *openapiclient.NewAvaProjectWrapper(*openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")), openapiclient.InvoiceType("PartialInvoice"), false, float32(123)) // AvaProjectWrapper | 
    targetFormat := "targetFormat_example" // string | The target format for the XRechnung export. Defaults to Ubl. Options are: Ubl, CrossIndustryInvoice and FacturX (ZUGFerd) (optional)
    disablePriceRounding := true // bool | If this is enabled, then the price values in the XRechnung export will not be rounded. Otherwise, prices will by default be rounded to two decimal places. (optional)
    disableBrandingComment := true // bool | If this is set to true, exported XRechnung files will not contain an XML comment indicating their origin. This is false by default, meaning the comment will be included. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XRechnungConversionApi.XRechnungConversionConvertAvaToXRechnung(context.Background()).AvaWrapper(avaWrapper).TargetFormat(targetFormat).DisablePriceRounding(disablePriceRounding).DisableBrandingComment(disableBrandingComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRechnungConversionApi.XRechnungConversionConvertAvaToXRechnung``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XRechnungConversionConvertAvaToXRechnung`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `XRechnungConversionApi.XRechnungConversionConvertAvaToXRechnung`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXRechnungConversionConvertAvaToXRechnungRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaWrapper** | [**AvaProjectWrapper**](AvaProjectWrapper.md) |  | 
 **targetFormat** | **string** | The target format for the XRechnung export. Defaults to Ubl. Options are: Ubl, CrossIndustryInvoice and FacturX (ZUGFerd) | 
 **disablePriceRounding** | **bool** | If this is enabled, then the price values in the XRechnung export will not be rounded. Otherwise, prices will by default be rounded to two decimal places. | 
 **disableBrandingComment** | **bool** | If this is set to true, exported XRechnung files will not contain an XML comment indicating their origin. This is false by default, meaning the comment will be included. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XRechnungConversionConvertInvoiceToPdfInvoice

> *os.File XRechnungConversionConvertInvoiceToPdfInvoice(ctx).Invoice(invoice).Execute()

This will convert an Invoice object to a PDF file

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
    invoice := *openapiclient.NewInvoice(openapiclient.SourceType("Self")) // Invoice | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XRechnungConversionApi.XRechnungConversionConvertInvoiceToPdfInvoice(context.Background()).Invoice(invoice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRechnungConversionApi.XRechnungConversionConvertInvoiceToPdfInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XRechnungConversionConvertInvoiceToPdfInvoice`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `XRechnungConversionApi.XRechnungConversionConvertInvoiceToPdfInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXRechnungConversionConvertInvoiceToPdfInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoice** | [**Invoice**](Invoice.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XRechnungConversionConvertInvoiceToXRechnung

> *os.File XRechnungConversionConvertInvoiceToXRechnung(ctx).Invoice(invoice).TargetFormat(targetFormat).DisablePriceRounding(disablePriceRounding).DisableBrandingComment(disableBrandingComment).Execute()

This will convert an Invoice object to an XRechnung file

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
    invoice := *openapiclient.NewInvoice(openapiclient.SourceType("Self")) // Invoice | 
    targetFormat := "targetFormat_example" // string | The target format for the XRechnung export. Defaults to Ubl. Options are: Ubl, CrossIndustryInvoice and FacturX (ZUGFerd) (optional)
    disablePriceRounding := true // bool | If this is enabled, then the price values in the XRechnung export will not be rounded. Otherwise, prices will by default be rounded to two decimal places. (optional)
    disableBrandingComment := true // bool | If this is set to true, exported XRechnung files will not contain an XML comment indicating their origin. This is false by default, meaning the comment will be included. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XRechnungConversionApi.XRechnungConversionConvertInvoiceToXRechnung(context.Background()).Invoice(invoice).TargetFormat(targetFormat).DisablePriceRounding(disablePriceRounding).DisableBrandingComment(disableBrandingComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRechnungConversionApi.XRechnungConversionConvertInvoiceToXRechnung``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XRechnungConversionConvertInvoiceToXRechnung`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `XRechnungConversionApi.XRechnungConversionConvertInvoiceToXRechnung`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXRechnungConversionConvertInvoiceToXRechnungRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoice** | [**Invoice**](Invoice.md) |  | 
 **targetFormat** | **string** | The target format for the XRechnung export. Defaults to Ubl. Options are: Ubl, CrossIndustryInvoice and FacturX (ZUGFerd) | 
 **disablePriceRounding** | **bool** | If this is enabled, then the price values in the XRechnung export will not be rounded. Otherwise, prices will by default be rounded to two decimal places. | 
 **disableBrandingComment** | **bool** | If this is set to true, exported XRechnung files will not contain an XML comment indicating their origin. This is false by default, meaning the comment will be included. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XRechnungConversionConvertXRechnungToAva

> AvaProjectWrapper XRechnungConversionConvertXRechnungToAva(ctx).XrechnungFile(xrechnungFile).Execute()

This will read an XRechnung file and convert it to an AVA wrapper object

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
    xrechnungFile := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XRechnungConversionApi.XRechnungConversionConvertXRechnungToAva(context.Background()).XrechnungFile(xrechnungFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRechnungConversionApi.XRechnungConversionConvertXRechnungToAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XRechnungConversionConvertXRechnungToAva`: AvaProjectWrapper
    fmt.Fprintf(os.Stdout, "Response from `XRechnungConversionApi.XRechnungConversionConvertXRechnungToAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXRechnungConversionConvertXRechnungToAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xrechnungFile** | ***os.File** |  | 

### Return type

[**AvaProjectWrapper**](AvaProjectWrapper.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XRechnungConversionConvertXRechnungToInvoice

> Invoice XRechnungConversionConvertXRechnungToInvoice(ctx).XrechnungFile(xrechnungFile).Execute()

This will read an XRechnung file and convert it to an Invoice object

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
    xrechnungFile := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XRechnungConversionApi.XRechnungConversionConvertXRechnungToInvoice(context.Background()).XrechnungFile(xrechnungFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRechnungConversionApi.XRechnungConversionConvertXRechnungToInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XRechnungConversionConvertXRechnungToInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `XRechnungConversionApi.XRechnungConversionConvertXRechnungToInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXRechnungConversionConvertXRechnungToInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xrechnungFile** | ***os.File** |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## XRechnungConversionMergeInvoiceIntoPdfFile

> *os.File XRechnungConversionMergeInvoiceIntoPdfFile(ctx).PdfEmbeddModel(pdfEmbeddModel).TargetFormat(targetFormat).DisablePriceRounding(disablePriceRounding).DisableBrandingComment(disableBrandingComment).Execute()

This will embedd an XRechnung file into an existing PDF file as an attachment. If Factur-X is chosen as target format, AVACloud will also insert the required Factur-X XML metadata into the PDF file. AVACloud does not, however, ensure the PDF itself is meeting all the Factur-X requirements.

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
    pdfEmbeddModel := *openapiclient.NewPostXRechnungPdfEmbedd(*openapiclient.NewInvoice(openapiclient.SourceType("Self")), "PdfFileBase64_example") // PostXRechnungPdfEmbedd | 
    targetFormat := "targetFormat_example" // string | The target format for the XRechnung export. Defaults to Ubl. Options are: Ubl, CrossIndustryInvoice and FacturX (ZUGFerd) (optional)
    disablePriceRounding := true // bool | If this is enabled, then the price values in the XRechnung export will not be rounded. Otherwise, prices will by default be rounded to two decimal places. (optional)
    disableBrandingComment := true // bool | If this is set to true, exported XRechnung files will not contain an XML comment indicating their origin. This is false by default, meaning the comment will be included. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.XRechnungConversionApi.XRechnungConversionMergeInvoiceIntoPdfFile(context.Background()).PdfEmbeddModel(pdfEmbeddModel).TargetFormat(targetFormat).DisablePriceRounding(disablePriceRounding).DisableBrandingComment(disableBrandingComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `XRechnungConversionApi.XRechnungConversionMergeInvoiceIntoPdfFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `XRechnungConversionMergeInvoiceIntoPdfFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `XRechnungConversionApi.XRechnungConversionMergeInvoiceIntoPdfFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiXRechnungConversionMergeInvoiceIntoPdfFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdfEmbeddModel** | [**PostXRechnungPdfEmbedd**](PostXRechnungPdfEmbedd.md) |  | 
 **targetFormat** | **string** | The target format for the XRechnung export. Defaults to Ubl. Options are: Ubl, CrossIndustryInvoice and FacturX (ZUGFerd) | 
 **disablePriceRounding** | **bool** | If this is enabled, then the price values in the XRechnung export will not be rounded. Otherwise, prices will by default be rounded to two decimal places. | 
 **disableBrandingComment** | **bool** | If this is set to true, exported XRechnung files will not contain an XML comment indicating their origin. This is false by default, meaning the comment will be included. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

