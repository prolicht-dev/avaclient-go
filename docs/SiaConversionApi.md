# \SiaConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SiaConversionConvertToAva**](SiaConversionApi.md#SiaConversionConvertToAva) | **Post** /conversion/sia/ava | Converts SIA 451 files to Dangl.AVA projects
[**SiaConversionConvertToExcel**](SiaConversionApi.md#SiaConversionConvertToExcel) | **Post** /conversion/sia/excel | Converts SIA 451 files to Excel
[**SiaConversionConvertToFlatAva**](SiaConversionApi.md#SiaConversionConvertToFlatAva) | **Post** /conversion/sia/flat-ava | Converts SIA 451 files to Dangl.AVA projects
[**SiaConversionConvertToGaeb**](SiaConversionApi.md#SiaConversionConvertToGaeb) | **Post** /conversion/sia/gaeb | Converts SIA 451 files to GAEB files
[**SiaConversionConvertToOenorm**](SiaConversionApi.md#SiaConversionConvertToOenorm) | **Post** /conversion/sia/oenorm | Converts SIA 451 files to Oenorm files



## SiaConversionConvertToAva

> ProjectDto SiaConversionConvertToAva(ctx).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).SiaFile(siaFile).Execute()

Converts SIA 451 files to Dangl.AVA projects

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
    siaFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiaConversionApi.SiaConversionConvertToAva(context.Background()).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).SiaFile(siaFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiaConversionApi.SiaConversionConvertToAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiaConversionConvertToAva`: ProjectDto
    fmt.Fprintf(os.Stdout, "Response from `SiaConversionApi.SiaConversionConvertToAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiaConversionConvertToAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removePlainTextLongTexts** | **bool** | If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **bool** | If set to true, html long texts will be removed from the output to reduce response sizes | 
 **siaFile** | ***os.File** | The input file | 

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


## SiaConversionConvertToExcel

> *os.File SiaConversionConvertToExcel(ctx).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).SiaFile(siaFile).Execute()

Converts SIA 451 files to Excel

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
    siaFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiaConversionApi.SiaConversionConvertToExcel(context.Background()).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).SiaFile(siaFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiaConversionApi.SiaConversionConvertToExcel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiaConversionConvertToExcel`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SiaConversionApi.SiaConversionConvertToExcel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiaConversionConvertToExcelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **writePrices** | **bool** | Defaults to true | 
 **writeLongTexts** | **bool** | Defaults to true | 
 **conversionCulture** | **string** | The culture that should be used for the conversion process, to have localized Excel files | 
 **includeArticleNumbers** | **bool** | If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from &#39;position.commerceProperties.articleNumber&#39; | 
 **siaFile** | ***os.File** | The input file | 

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


## SiaConversionConvertToFlatAva

> FlatAvaProject SiaConversionConvertToFlatAva(ctx).SiaFile(siaFile).Execute()

Converts SIA 451 files to Dangl.AVA projects

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
    siaFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiaConversionApi.SiaConversionConvertToFlatAva(context.Background()).SiaFile(siaFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiaConversionApi.SiaConversionConvertToFlatAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiaConversionConvertToFlatAva`: FlatAvaProject
    fmt.Fprintf(os.Stdout, "Response from `SiaConversionApi.SiaConversionConvertToFlatAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiaConversionConvertToFlatAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siaFile** | ***os.File** | The input file | 

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


## SiaConversionConvertToGaeb

> *os.File SiaConversionConvertToGaeb(ctx).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).SiaFile(siaFile).Execute()

Converts SIA 451 files to GAEB files

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
    siaFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiaConversionApi.SiaConversionConvertToGaeb(context.Background()).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).SiaFile(siaFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiaConversionApi.SiaConversionConvertToGaeb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiaConversionConvertToGaeb`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SiaConversionApi.SiaConversionConvertToGaeb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiaConversionConvertToGaebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationGaebType** | **string** | Defaults to GAEB XML V3.2 | 
 **targetExchangePhaseTransform** | **string** | Defaults to none, meaning no transformation will be done. The phases are: Base &#x3D; 81 CostEstimate &#x3D; 82 OfferRequest &#x3D; 83 Offer &#x3D; 84 SideOffer &#x3D; 85 Grant &#x3D; 86 | 
 **enforceStrictOfferPhaseLongTextOutput** | **bool** | Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export. | 
 **exportQuantityDetermination** | **bool** | Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the &#39;QtyDeterm&#39; (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the &#39;QuantityComponents&#39; property of positions. Please see the entry for &#39;Quantity Determination&#39; in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the &#39;QuantityComponents&#39; property. | 
 **removeUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 
 **forceIncludeDescriptions** | **bool** | If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions. | 
 **treatNullItemNumberSchemaAsInvalid** | **bool** | When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons. | 
 **siaFile** | ***os.File** | The input file | 

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


## SiaConversionConvertToOenorm

> *os.File SiaConversionConvertToOenorm(ctx).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).SiaFile(siaFile).Execute()

Converts SIA 451 files to Oenorm files

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
    siaFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SiaConversionApi.SiaConversionConvertToOenorm(context.Background()).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).SiaFile(siaFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiaConversionApi.SiaConversionConvertToOenorm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiaConversionConvertToOenorm`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SiaConversionApi.SiaConversionConvertToOenorm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiaConversionConvertToOenormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationOenormType** | **string** | Defaults to Lv2015 | 
 **tryRepairProjectStructure** | **bool** | Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 
 **skipTryEnforceSchemaCompliantXmlOutput** | **bool** | If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option. | 
 **removeUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 
 **siaFile** | ***os.File** | The input file | 

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

