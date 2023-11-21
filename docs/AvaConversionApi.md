# \AvaConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AvaConversionConvertToAva**](AvaConversionApi.md#AvaConversionConvertToAva) | **Post** /conversion/ava/ava | Converts Dangl.AVA projects to Dangl.AVA. This is useful when you want to generate the calculated properties.
[**AvaConversionConvertToDatanorm**](AvaConversionApi.md#AvaConversionConvertToDatanorm) | **Post** /conversion/ava/datanorm | Converts Dangl.AVA projects to Datanorm
[**AvaConversionConvertToExcel**](AvaConversionApi.md#AvaConversionConvertToExcel) | **Post** /conversion/ava/excel | Converts Dangl.AVA projects to Excel
[**AvaConversionConvertToGaeb**](AvaConversionApi.md#AvaConversionConvertToGaeb) | **Post** /conversion/ava/gaeb | Converts Dangl.AVA projects to GAEB
[**AvaConversionConvertToIdsConnect**](AvaConversionApi.md#AvaConversionConvertToIdsConnect) | **Post** /conversion/ava/ids-connect | Converts Dangl.AVA projects to IDS Connect files
[**AvaConversionConvertToOenorm**](AvaConversionApi.md#AvaConversionConvertToOenorm) | **Post** /conversion/ava/oenorm | Converts Dangl.AVA projects to Oenorm
[**AvaConversionConvertToReb**](AvaConversionApi.md#AvaConversionConvertToReb) | **Post** /conversion/ava/reb | Converts Dangl.AVA projects to REB
[**AvaConversionConvertToSia**](AvaConversionApi.md#AvaConversionConvertToSia) | **Post** /conversion/ava/sia | Converts Dangl.AVA projects to SIA 451
[**AvaConversionConvertToUgl**](AvaConversionApi.md#AvaConversionConvertToUgl) | **Post** /conversion/ava/ugl | Converts Dangl.AVA projects to UGL



## AvaConversionConvertToAva

> ProjectDto AvaConversionConvertToAva(ctx).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).Execute()

Converts Dangl.AVA projects to Dangl.AVA. This is useful when you want to generate the calculated properties.

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
    avaProject := *openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")) // ProjectDto | The Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    removePlainTextLongTexts := true // bool | If set to true, plain text long texts will be removed from the output to reduce response sizes (optional)
    removeHtmlLongTexts := true // bool | If set to true, html long texts will be removed from the output to reduce response sizes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvaConversionApi.AvaConversionConvertToAva(context.Background()).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaConversionApi.AvaConversionConvertToAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaConversionConvertToAva`: ProjectDto
    fmt.Fprintf(os.Stdout, "Response from `AvaConversionApi.AvaConversionConvertToAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaConversionConvertToAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProject** | [**ProjectDto**](ProjectDto.md) | The Dangl.AVA project | 
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


## AvaConversionConvertToDatanorm

> *os.File AvaConversionConvertToDatanorm(ctx).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DatanormDestinationVersion(datanormDestinationVersion).Execute()

Converts Dangl.AVA projects to Datanorm

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
    avaProject := *openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")) // ProjectDto | The Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    datanormDestinationVersion := "datanormDestinationVersion_example" // string | The Datanorm version to convert to. Defaults to V4. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvaConversionApi.AvaConversionConvertToDatanorm(context.Background()).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DatanormDestinationVersion(datanormDestinationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaConversionApi.AvaConversionConvertToDatanorm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaConversionConvertToDatanorm`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AvaConversionApi.AvaConversionConvertToDatanorm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaConversionConvertToDatanormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProject** | [**ProjectDto**](ProjectDto.md) | The Dangl.AVA project | 
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


## AvaConversionConvertToExcel

> *os.File AvaConversionConvertToExcel(ctx).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).Execute()

Converts Dangl.AVA projects to Excel

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
    avaProject := *openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")) // ProjectDto | The Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    writePrices := true // bool | Defaults to true (optional)
    writeLongTexts := true // bool | Defaults to true (optional)
    conversionCulture := "conversionCulture_example" // string | The culture that should be used for the conversion process, to have localized Excel files (optional)
    includeArticleNumbers := true // bool | If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from 'position.commerceProperties.articleNumber' (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvaConversionApi.AvaConversionConvertToExcel(context.Background()).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaConversionApi.AvaConversionConvertToExcel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaConversionConvertToExcel`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AvaConversionApi.AvaConversionConvertToExcel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaConversionConvertToExcelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProject** | [**ProjectDto**](ProjectDto.md) | The Dangl.AVA project | 
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


## AvaConversionConvertToGaeb

> *os.File AvaConversionConvertToGaeb(ctx).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).Execute()

Converts Dangl.AVA projects to GAEB

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
    avaProject := *openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")) // ProjectDto | The Dangl.AVA project
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
    resp, r, err := apiClient.AvaConversionApi.AvaConversionConvertToGaeb(context.Background()).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaConversionApi.AvaConversionConvertToGaeb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaConversionConvertToGaeb`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AvaConversionApi.AvaConversionConvertToGaeb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaConversionConvertToGaebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProject** | [**ProjectDto**](ProjectDto.md) | The Dangl.AVA project | 
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


## AvaConversionConvertToIdsConnect

> *os.File AvaConversionConvertToIdsConnect(ctx).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).Version(version).Execute()

Converts Dangl.AVA projects to IDS Connect files

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
    avaProject := *openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")) // ProjectDto | The Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    version := "version_example" // string | The IDS Connect version to convert to. Defaults to V2_5. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvaConversionApi.AvaConversionConvertToIdsConnect(context.Background()).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaConversionApi.AvaConversionConvertToIdsConnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaConversionConvertToIdsConnect`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AvaConversionApi.AvaConversionConvertToIdsConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaConversionConvertToIdsConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProject** | [**ProjectDto**](ProjectDto.md) | The Dangl.AVA project | 
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


## AvaConversionConvertToOenorm

> *os.File AvaConversionConvertToOenorm(ctx).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).Execute()

Converts Dangl.AVA projects to Oenorm

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
    avaProject := *openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")) // ProjectDto | The Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    destinationOenormType := "destinationOenormType_example" // string | Defaults to Lv2015 (optional)
    tryRepairProjectStructure := true // bool | Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target (optional)
    skipTryEnforceSchemaCompliantXmlOutput := true // bool | If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option. (optional)
    removeUnprintableCharactersFromTexts := true // bool | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvaConversionApi.AvaConversionConvertToOenorm(context.Background()).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaConversionApi.AvaConversionConvertToOenorm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaConversionConvertToOenorm`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AvaConversionApi.AvaConversionConvertToOenorm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaConversionConvertToOenormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProject** | [**ProjectDto**](ProjectDto.md) | The Dangl.AVA project | 
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


## AvaConversionConvertToReb

> *os.File AvaConversionConvertToReb(ctx).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationRebType(destinationRebType).Execute()

Converts Dangl.AVA projects to REB

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
    avaProject := *openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")) // ProjectDto | The Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    destinationRebType := "destinationRebType_example" // string | Defaults to D11 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvaConversionApi.AvaConversionConvertToReb(context.Background()).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).DestinationRebType(destinationRebType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaConversionApi.AvaConversionConvertToReb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaConversionConvertToReb`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AvaConversionApi.AvaConversionConvertToReb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaConversionConvertToRebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProject** | [**ProjectDto**](ProjectDto.md) | The Dangl.AVA project | 
 **tryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 
 **destinationRebType** | **string** | Defaults to D11 | 

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


## AvaConversionConvertToSia

> *os.File AvaConversionConvertToSia(ctx).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).SiaDestinationType(siaDestinationType).Execute()

Converts Dangl.AVA projects to SIA 451

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
    avaProject := *openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")) // ProjectDto | The Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    siaDestinationType := "siaDestinationType_example" // string | Defaults to Sia451 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvaConversionApi.AvaConversionConvertToSia(context.Background()).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).SiaDestinationType(siaDestinationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaConversionApi.AvaConversionConvertToSia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaConversionConvertToSia`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AvaConversionApi.AvaConversionConvertToSia`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaConversionConvertToSiaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProject** | [**ProjectDto**](ProjectDto.md) | The Dangl.AVA project | 
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


## AvaConversionConvertToUgl

> *os.File AvaConversionConvertToUgl(ctx).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).UglDestinationVersion(uglDestinationVersion).Execute()

Converts Dangl.AVA projects to UGL

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
    avaProject := *openapiclient.NewProjectDto("Id_example", int32(123), false, openapiclient.PriceRoundingModeDto("Normal")) // ProjectDto | The Dangl.AVA project
    tryAutoGenerateItemNumbersAndSchema := true // bool | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. (optional)
    uglDestinationVersion := "uglDestinationVersion_example" // string | The UGL version to convert to. Defaults to V1. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvaConversionApi.AvaConversionConvertToUgl(context.Background()).AvaProject(avaProject).TryAutoGenerateItemNumbersAndSchema(tryAutoGenerateItemNumbersAndSchema).UglDestinationVersion(uglDestinationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvaConversionApi.AvaConversionConvertToUgl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvaConversionConvertToUgl`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AvaConversionApi.AvaConversionConvertToUgl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvaConversionConvertToUglRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avaProject** | [**ProjectDto**](ProjectDto.md) | The Dangl.AVA project | 
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

