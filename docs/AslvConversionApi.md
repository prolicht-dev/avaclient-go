# \AslvConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AslvConversionConvertToAva**](AslvConversionApi.md#AslvConversionConvertToAva) | **Post** /conversion/aslv/ava | Converts Aslv files to Dangl.AVA projects
[**AslvConversionConvertToExcel**](AslvConversionApi.md#AslvConversionConvertToExcel) | **Post** /conversion/aslv/excel | Converts Aslv files to Excel
[**AslvConversionConvertToFlatAva**](AslvConversionApi.md#AslvConversionConvertToFlatAva) | **Post** /conversion/aslv/flat-ava | Converts Aslv files to Dangl.AVA projects
[**AslvConversionConvertToGaeb**](AslvConversionApi.md#AslvConversionConvertToGaeb) | **Post** /conversion/aslv/gaeb | Converts Aslv files to GAEB files
[**AslvConversionConvertToOenorm**](AslvConversionApi.md#AslvConversionConvertToOenorm) | **Post** /conversion/aslv/oenorm | Converts Aslv files to Oenorm files



## AslvConversionConvertToAva

> ProjectDto AslvConversionConvertToAva(ctx).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).AslvFile(aslvFile).Execute()

Converts Aslv files to Dangl.AVA projects

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
    aslvFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AslvConversionApi.AslvConversionConvertToAva(context.Background()).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).AslvFile(aslvFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AslvConversionApi.AslvConversionConvertToAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AslvConversionConvertToAva`: ProjectDto
    fmt.Fprintf(os.Stdout, "Response from `AslvConversionApi.AslvConversionConvertToAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAslvConversionConvertToAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removePlainTextLongTexts** | **bool** | If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **bool** | If set to true, html long texts will be removed from the output to reduce response sizes | 
 **aslvFile** | ***os.File** | The input file | 

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


## AslvConversionConvertToExcel

> *os.File AslvConversionConvertToExcel(ctx).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).AslvFile(aslvFile).Execute()

Converts Aslv files to Excel

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
    writePrices := true // bool | Defaults to true (optional)
    writeLongTexts := true // bool | Defaults to true (optional)
    conversionCulture := "conversionCulture_example" // string | The culture that should be used for the conversion process, to have localized Excel files (optional)
    includeArticleNumbers := true // bool | If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from 'position.commerceProperties.articleNumber' (optional)
    aslvFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AslvConversionApi.AslvConversionConvertToExcel(context.Background()).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).AslvFile(aslvFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AslvConversionApi.AslvConversionConvertToExcel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AslvConversionConvertToExcel`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AslvConversionApi.AslvConversionConvertToExcel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAslvConversionConvertToExcelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writePrices** | **bool** | Defaults to true | 
 **writeLongTexts** | **bool** | Defaults to true | 
 **conversionCulture** | **string** | The culture that should be used for the conversion process, to have localized Excel files | 
 **includeArticleNumbers** | **bool** | If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from &#39;position.commerceProperties.articleNumber&#39; | 
 **aslvFile** | ***os.File** | The input file | 

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


## AslvConversionConvertToFlatAva

> FlatAvaProject AslvConversionConvertToFlatAva(ctx).AslvFile(aslvFile).Execute()

Converts Aslv files to Dangl.AVA projects

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
    aslvFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AslvConversionApi.AslvConversionConvertToFlatAva(context.Background()).AslvFile(aslvFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AslvConversionApi.AslvConversionConvertToFlatAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AslvConversionConvertToFlatAva`: FlatAvaProject
    fmt.Fprintf(os.Stdout, "Response from `AslvConversionApi.AslvConversionConvertToFlatAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAslvConversionConvertToFlatAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aslvFile** | ***os.File** | The input file | 

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


## AslvConversionConvertToGaeb

> *os.File AslvConversionConvertToGaeb(ctx).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).AslvFile(aslvFile).Execute()

Converts Aslv files to GAEB files

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
    destinationGaebType := "destinationGaebType_example" // string | Defaults to GAEB XML V3.2 (optional)
    targetExchangePhaseTransform := "targetExchangePhaseTransform_example" // string | Defaults to none, meaning no transformation will be done. The phases are: Base = 81 CostEstimate = 82 OfferRequest = 83 Offer = 84 SideOffer = 85 Grant = 86 (optional)
    enforceStrictOfferPhaseLongTextOutput := true // bool | Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export. (optional)
    exportQuantityDetermination := true // bool | Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the 'QtyDeterm' (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the 'QuantityComponents' property of positions. Please see the entry for 'Quantity Determination' in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the 'QuantityComponents' property. (optional)
    removeUnprintableCharactersFromTexts := true // bool | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. (optional)
    forceIncludeDescriptions := true // bool | If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions. (optional)
    treatNullItemNumberSchemaAsInvalid := true // bool | When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons. (optional)
    aslvFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AslvConversionApi.AslvConversionConvertToGaeb(context.Background()).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).AslvFile(aslvFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AslvConversionApi.AslvConversionConvertToGaeb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AslvConversionConvertToGaeb`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AslvConversionApi.AslvConversionConvertToGaeb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAslvConversionConvertToGaebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationGaebType** | **string** | Defaults to GAEB XML V3.2 | 
 **targetExchangePhaseTransform** | **string** | Defaults to none, meaning no transformation will be done. The phases are: Base &#x3D; 81 CostEstimate &#x3D; 82 OfferRequest &#x3D; 83 Offer &#x3D; 84 SideOffer &#x3D; 85 Grant &#x3D; 86 | 
 **enforceStrictOfferPhaseLongTextOutput** | **bool** | Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export. | 
 **exportQuantityDetermination** | **bool** | Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the &#39;QtyDeterm&#39; (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the &#39;QuantityComponents&#39; property of positions. Please see the entry for &#39;Quantity Determination&#39; in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the &#39;QuantityComponents&#39; property. | 
 **removeUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 
 **forceIncludeDescriptions** | **bool** | If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions. | 
 **treatNullItemNumberSchemaAsInvalid** | **bool** | When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons. | 
 **aslvFile** | ***os.File** | The input file | 

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


## AslvConversionConvertToOenorm

> *os.File AslvConversionConvertToOenorm(ctx).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).AslvFile(aslvFile).Execute()

Converts Aslv files to Oenorm files

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
    destinationOenormType := "destinationOenormType_example" // string | Defaults to Lv2015 (optional)
    tryRepairProjectStructure := true // bool | Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target (optional)
    skipTryEnforceSchemaCompliantXmlOutput := true // bool | If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option. (optional)
    removeUnprintableCharactersFromTexts := true // bool | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. (optional)
    aslvFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AslvConversionApi.AslvConversionConvertToOenorm(context.Background()).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).AslvFile(aslvFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AslvConversionApi.AslvConversionConvertToOenorm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AslvConversionConvertToOenorm`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AslvConversionApi.AslvConversionConvertToOenorm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAslvConversionConvertToOenormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationOenormType** | **string** | Defaults to Lv2015 | 
 **tryRepairProjectStructure** | **bool** | Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 
 **skipTryEnforceSchemaCompliantXmlOutput** | **bool** | If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option. | 
 **removeUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 
 **aslvFile** | ***os.File** | The input file | 

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

