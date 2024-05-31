# \FlatAvaConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlatAvaConversionConvertToAva**](FlatAvaConversionApi.md#FlatAvaConversionConvertToAva) | **Post** /conversion/flat-ava/ava | Converts flat Dangl.AVA projects to Dangl.AVA. This is useful when you want to generate the calculated properties.
[**FlatAvaConversionConvertToDatanorm**](FlatAvaConversionApi.md#FlatAvaConversionConvertToDatanorm) | **Post** /conversion/flat-ava/datanorm | Converts flat Dangl.AVA projects to Datanorm
[**FlatAvaConversionConvertToExcel**](FlatAvaConversionApi.md#FlatAvaConversionConvertToExcel) | **Post** /conversion/flat-ava/excel | Converts flat Dangl.AVA projects to Excel
[**FlatAvaConversionConvertToFlatAva**](FlatAvaConversionApi.md#FlatAvaConversionConvertToFlatAva) | **Post** /conversion/flat-ava/flat-ava | Converts flat Dangl.AVA projects to Dangl.AVA. This is useful when you want to generate the calculated properties.
[**FlatAvaConversionConvertToGaeb**](FlatAvaConversionApi.md#FlatAvaConversionConvertToGaeb) | **Post** /conversion/flat-ava/gaeb | Converts flat Dangl.AVA projects to GAEB
[**FlatAvaConversionConvertToIdsConnect**](FlatAvaConversionApi.md#FlatAvaConversionConvertToIdsConnect) | **Post** /conversion/flat-ava/ids-connect | Converts flat Dangl.AVA projects to IDS Connect files
[**FlatAvaConversionConvertToOenorm**](FlatAvaConversionApi.md#FlatAvaConversionConvertToOenorm) | **Post** /conversion/flat-ava/oenorm | Converts flat Dangl.AVA projects to Oenorm
[**FlatAvaConversionConvertToReb**](FlatAvaConversionApi.md#FlatAvaConversionConvertToReb) | **Post** /conversion/flat-ava/reb | Converts flat Dangl.AVA projects to REB
[**FlatAvaConversionConvertToSia**](FlatAvaConversionApi.md#FlatAvaConversionConvertToSia) | **Post** /conversion/flat-ava/sia | Converts flat Dangl.AVA projects to SIA 451
[**FlatAvaConversionConvertToUgl**](FlatAvaConversionApi.md#FlatAvaConversionConvertToUgl) | **Post** /conversion/flat-ava/ugl | Converts flat Dangl.AVA projects to UGL



## FlatAvaConversionConvertToAva

> ProjectDto FlatAvaConversionConvertToAva(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).Execute()

Converts flat Dangl.AVA projects to Dangl.AVA. This is useful when you want to generate the calculated properties.

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    removePlainTextLongTexts := true // bool | If set to true, plain text long texts will be removed from the output to reduce response sizes (optional)
    removeHtmlLongTexts := true // bool | If set to true, html long texts will be removed from the output to reduce response sizes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToAva(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToAva`: ProjectDto
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **removePlainTextLongTexts** | **bool** | If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **bool** | If set to true, html long texts will be removed from the output to reduce response sizes | 

### Return type

[**ProjectDto**](ProjectDto.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/vnd.com.dangl-it.ProjectDto.v1+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlatAvaConversionConvertToDatanorm

> *os.File FlatAvaConversionConvertToDatanorm(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DatanormDestinationVersion(datanormDestinationVersion).Execute()

Converts flat Dangl.AVA projects to Datanorm

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    datanormDestinationVersion := "datanormDestinationVersion_example" // string | The Datanorm version to convert to. Defaults to V4. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToDatanorm(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DatanormDestinationVersion(datanormDestinationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToDatanorm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToDatanorm`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToDatanorm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToDatanormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **datanormDestinationVersion** | **string** | The Datanorm version to convert to. Defaults to V4. | 

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


## FlatAvaConversionConvertToExcel

> *os.File FlatAvaConversionConvertToExcel(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).Execute()

Converts flat Dangl.AVA projects to Excel

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    writePrices := true // bool | Defaults to true (optional)
    writeLongTexts := true // bool | Defaults to true (optional)
    conversionCulture := "conversionCulture_example" // string | The culture that should be used for the conversion process, to have localized Excel files (optional)
    includeArticleNumbers := true // bool | If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from 'position.commerceProperties.articleNumber' (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToExcel(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToExcel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToExcel`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToExcel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToExcelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **writePrices** | **bool** | Defaults to true | 
 **writeLongTexts** | **bool** | Defaults to true | 
 **conversionCulture** | **string** | The culture that should be used for the conversion process, to have localized Excel files | 
 **includeArticleNumbers** | **bool** | If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from &#39;position.commerceProperties.articleNumber&#39; | 

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


## FlatAvaConversionConvertToFlatAva

> FlatAvaProject FlatAvaConversionConvertToFlatAva(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).Execute()

Converts flat Dangl.AVA projects to Dangl.AVA. This is useful when you want to generate the calculated properties.

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToFlatAva(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToFlatAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToFlatAva`: FlatAvaProject
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToFlatAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToFlatAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 

### Return type

[**FlatAvaProject**](FlatAvaProject.md)

### Authorization

[Dangl.Identity](../README.md#Dangl.Identity)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlatAvaConversionConvertToGaeb

> *os.File FlatAvaConversionConvertToGaeb(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).Execute()

Converts flat Dangl.AVA projects to GAEB

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    destinationGaebType := "destinationGaebType_example" // string | Defaults to GAEB XML V3.2 (optional)
    targetExchangePhaseTransform := "targetExchangePhaseTransform_example" // string | Defaults to none, meaning no transformation will be done. The phases are: Base = 81 CostEstimate = 82 OfferRequest = 83 Offer = 84 SideOffer = 85 Grant = 86 (optional)
    enforceStrictOfferPhaseLongTextOutput := true // bool | Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export. (optional)
    exportQuantityDetermination := true // bool | Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the 'QtyDeterm' (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the 'QuantityComponents' property of positions. Please see the entry for 'Quantity Determination' in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the 'QuantityComponents' property. (optional)
    removeUnprintableCharactersFromTexts := true // bool | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. (optional)
    forceIncludeDescriptions := true // bool | If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions. (optional)
    treatNullItemNumberSchemaAsInvalid := true // bool | When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToGaeb(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToGaeb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToGaeb`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToGaeb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToGaebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **destinationGaebType** | **string** | Defaults to GAEB XML V3.2 | 
 **targetExchangePhaseTransform** | **string** | Defaults to none, meaning no transformation will be done. The phases are: Base &#x3D; 81 CostEstimate &#x3D; 82 OfferRequest &#x3D; 83 Offer &#x3D; 84 SideOffer &#x3D; 85 Grant &#x3D; 86 | 
 **enforceStrictOfferPhaseLongTextOutput** | **bool** | Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export. | 
 **exportQuantityDetermination** | **bool** | Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the &#39;QtyDeterm&#39; (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the &#39;QuantityComponents&#39; property of positions. Please see the entry for &#39;Quantity Determination&#39; in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the &#39;QuantityComponents&#39; property. | 
 **removeUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 
 **forceIncludeDescriptions** | **bool** | If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions. | 
 **treatNullItemNumberSchemaAsInvalid** | **bool** | When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons. | 

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


## FlatAvaConversionConvertToIdsConnect

> *os.File FlatAvaConversionConvertToIdsConnect(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).Version(version).Execute()

Converts flat Dangl.AVA projects to IDS Connect files

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    version := "version_example" // string | The IDS Connect version to convert to. Defaults to V2_5. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToIdsConnect(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToIdsConnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToIdsConnect`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToIdsConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToIdsConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **version** | **string** | The IDS Connect version to convert to. Defaults to V2_5. | 

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


## FlatAvaConversionConvertToOenorm

> *os.File FlatAvaConversionConvertToOenorm(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).Execute()

Converts flat Dangl.AVA projects to Oenorm

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    destinationOenormType := "destinationOenormType_example" // string | Defaults to Lv2015 (optional)
    tryRepairProjectStructure := true // bool | Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target (optional)
    skipTryEnforceSchemaCompliantXmlOutput := true // bool | If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option. (optional)
    removeUnprintableCharactersFromTexts := true // bool | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToOenorm(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToOenorm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToOenorm`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToOenorm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToOenormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **destinationOenormType** | **string** | Defaults to Lv2015 | 
 **tryRepairProjectStructure** | **bool** | Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 
 **skipTryEnforceSchemaCompliantXmlOutput** | **bool** | If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option. | 
 **removeUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 

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


## FlatAvaConversionConvertToReb

> *os.File FlatAvaConversionConvertToReb(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationRebType(destinationRebType).LastRowAddress(lastRowAddress).Execute()

Converts flat Dangl.AVA projects to REB

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    destinationRebType := "destinationRebType_example" // string | Defaults to D11 (optional)
    lastRowAddress := "lastRowAddress_example" // string | If this is present, the export to REB will be started from the next available row address after the given one. This must be a valid 6 character address, e.g. \"1234A0\" (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToReb(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationRebType(destinationRebType).LastRowAddress(lastRowAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToReb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToReb`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToReb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToRebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **destinationRebType** | **string** | Defaults to D11 | 
 **lastRowAddress** | **string** | If this is present, the export to REB will be started from the next available row address after the given one. This must be a valid 6 character address, e.g. \&quot;1234A0\&quot; | 

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


## FlatAvaConversionConvertToSia

> *os.File FlatAvaConversionConvertToSia(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).SiaDestinationType(siaDestinationType).Execute()

Converts flat Dangl.AVA projects to SIA 451

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    siaDestinationType := "siaDestinationType_example" // string | Defaults to Sia451 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToSia(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).SiaDestinationType(siaDestinationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToSia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToSia`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToSia`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToSiaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **siaDestinationType** | **string** | Defaults to Sia451 | 

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


## FlatAvaConversionConvertToUgl

> *os.File FlatAvaConversionConvertToUgl(ctx).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).UglDestinationVersion(uglDestinationVersion).Execute()

Converts flat Dangl.AVA projects to UGL

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
    flatAvaProject := *openapiclient.NewFlatAvaProject() // FlatAvaProject | The flat Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    uglDestinationVersion := "uglDestinationVersion_example" // string | The UGL version to convert to. Defaults to V1. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlatAvaConversionApi.FlatAvaConversionConvertToUgl(context.Background()).FlatAvaProject(flatAvaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).UglDestinationVersion(uglDestinationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlatAvaConversionApi.FlatAvaConversionConvertToUgl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FlatAvaConversionConvertToUgl`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FlatAvaConversionApi.FlatAvaConversionConvertToUgl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlatAvaConversionConvertToUglRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flatAvaProject** | [**FlatAvaProject**](FlatAvaProject.md) | The flat Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **uglDestinationVersion** | **string** | The UGL version to convert to. Defaults to V1. | 

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

