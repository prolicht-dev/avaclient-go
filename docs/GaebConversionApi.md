# \GaebConversionApi

All URIs are relative to *https://avacloud-api.dangl-it.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GaebConversionConvertToAva**](GaebConversionApi.md#GaebConversionConvertToAva) | **Post** /conversion/gaeb/ava | Converts GAEB files to Dangl.AVA projects
[**GaebConversionConvertToExcel**](GaebConversionApi.md#GaebConversionConvertToExcel) | **Post** /conversion/gaeb/excel | Converts GAEB files to Excel
[**GaebConversionConvertToGaeb**](GaebConversionApi.md#GaebConversionConvertToGaeb) | **Post** /conversion/gaeb/gaeb | Converts GAEB files to GAEB files. Used for example when transforming or repairing GAEB files.
[**GaebConversionConvertToOenorm**](GaebConversionApi.md#GaebConversionConvertToOenorm) | **Post** /conversion/gaeb/oenorm | Converts GAEB files to Oenorm files



## GaebConversionConvertToAva

> ProjectDto GaebConversionConvertToAva(ctx).SupportSkippedItemNumberLevelsInPositions(supportSkippedItemNumberLevelsInPositions).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).OutputHtmlAsXml(outputHtmlAsXml).KeepEmptyHtmlText(keepEmptyHtmlText).AllowUpperCaseItemNumbers(allowUpperCaseItemNumbers).AllowLumpSumItemsWithDifferingQuantities(allowLumpSumItemsWithDifferingQuantities).GaebFile(gaebFile).Execute()

Converts GAEB files to Dangl.AVA projects

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
    supportSkippedItemNumberLevelsInPositions := true // bool | Defaults to 'false'. This controls if, when deserializing GAEB files, skipped levels in position item numbers should be supported. For example, if an ItemNumberSchema defines three levels - two group levels and one position level - but the ItemNumber of the position is just '01.02', then it will be displayed as '01.__.02' if this is set to true. (optional)
    removePlainTextLongTexts := true // bool | If set to true, plain text long texts will be removed from the output to reduce response sizes (optional)
    removeHtmlLongTexts := true // bool | If set to true, html long texts will be removed from the output to reduce response sizes (optional)
    outputHtmlAsXml := true // bool | Defaults to 'false'. If this is enabled, then Html text will be output as XML in the output. This means that e.g. void Html tags will always be output with their closing tag, e.g. it will produce '<br />' instead of '<br>'. (optional)
    keepEmptyHtmlText := true // bool | Defaults to 'false'. If this is enabled, then Html text that is empty will be kept in the output. Otherwise, Html text without any plain text will be removed. This is useful for example if you want to keep texts that only consist of empty paragraphs in the output. (optional)
    allowUpperCaseItemNumbers := true // bool | Defaults to 'false'. If this is enabled, then the ItemNumber of positions will be in uppercase format if the source file has them. By default, all item numbers will be converted to lowercase, but this option will enable the option to support uppercase item numbers as well. (optional)
    allowLumpSumItemsWithDifferingQuantities := true // bool | Defaults to 'false'. By default, the GAEB standard requires lump sum items ('Pauschalpositionen' in German) to have a quantity of exactly 1. AVACloud does enforce this convention, but if you set this property to 'true', then differing quantities will be kept during the import. (optional)
    gaebFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaebConversionApi.GaebConversionConvertToAva(context.Background()).SupportSkippedItemNumberLevelsInPositions(supportSkippedItemNumberLevelsInPositions).RemovePlainTextLongTexts(removePlainTextLongTexts).RemoveHtmlLongTexts(removeHtmlLongTexts).OutputHtmlAsXml(outputHtmlAsXml).KeepEmptyHtmlText(keepEmptyHtmlText).AllowUpperCaseItemNumbers(allowUpperCaseItemNumbers).AllowLumpSumItemsWithDifferingQuantities(allowLumpSumItemsWithDifferingQuantities).GaebFile(gaebFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaebConversionApi.GaebConversionConvertToAva``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GaebConversionConvertToAva`: ProjectDto
    fmt.Fprintf(os.Stdout, "Response from `GaebConversionApi.GaebConversionConvertToAva`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGaebConversionConvertToAvaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportSkippedItemNumberLevelsInPositions** | **bool** | Defaults to &#39;false&#39;. This controls if, when deserializing GAEB files, skipped levels in position item numbers should be supported. For example, if an ItemNumberSchema defines three levels - two group levels and one position level - but the ItemNumber of the position is just &#39;01.02&#39;, then it will be displayed as &#39;01.__.02&#39; if this is set to true. | 
 **removePlainTextLongTexts** | **bool** | If set to true, plain text long texts will be removed from the output to reduce response sizes | 
 **removeHtmlLongTexts** | **bool** | If set to true, html long texts will be removed from the output to reduce response sizes | 
 **outputHtmlAsXml** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then Html text will be output as XML in the output. This means that e.g. void Html tags will always be output with their closing tag, e.g. it will produce &#39;&lt;br /&gt;&#39; instead of &#39;&lt;br&gt;&#39;. | 
 **keepEmptyHtmlText** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then Html text that is empty will be kept in the output. Otherwise, Html text without any plain text will be removed. This is useful for example if you want to keep texts that only consist of empty paragraphs in the output. | 
 **allowUpperCaseItemNumbers** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then the ItemNumber of positions will be in uppercase format if the source file has them. By default, all item numbers will be converted to lowercase, but this option will enable the option to support uppercase item numbers as well. | 
 **allowLumpSumItemsWithDifferingQuantities** | **bool** | Defaults to &#39;false&#39;. By default, the GAEB standard requires lump sum items (&#39;Pauschalpositionen&#39; in German) to have a quantity of exactly 1. AVACloud does enforce this convention, but if you set this property to &#39;true&#39;, then differing quantities will be kept during the import. | 
 **gaebFile** | ***os.File** | The input file | 

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


## GaebConversionConvertToExcel

> *os.File GaebConversionConvertToExcel(ctx).SupportSkippedItemNumberLevelsInPositions(supportSkippedItemNumberLevelsInPositions).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).OutputHtmlAsXml(outputHtmlAsXml).KeepEmptyHtmlText(keepEmptyHtmlText).AllowUpperCaseItemNumbers(allowUpperCaseItemNumbers).AllowLumpSumItemsWithDifferingQuantities(allowLumpSumItemsWithDifferingQuantities).GaebFile(gaebFile).Execute()

Converts GAEB files to Excel

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
    supportSkippedItemNumberLevelsInPositions := true // bool | Defaults to 'false'. This controls if, when deserializing GAEB files, skipped levels in position item numbers should be supported. For example, if an ItemNumberSchema defines three levels - two group levels and one position level - but the ItemNumber of the position is just '01.02', then it will be displayed as '01.__.02' if this is set to true. (optional)
    writePrices := true // bool | Defaults to true (optional)
    writeLongTexts := true // bool | Defaults to true (optional)
    conversionCulture := "conversionCulture_example" // string | The culture that should be used for the conversion process, to have localized Excel files (optional)
    includeArticleNumbers := true // bool | If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from 'position.commerceProperties.articleNumber' (optional)
    outputHtmlAsXml := true // bool | Defaults to 'false'. If this is enabled, then Html text will be output as XML in the output. This means that e.g. void Html tags will always be output with their closing tag, e.g. it will produce '<br />' instead of '<br>'. (optional)
    keepEmptyHtmlText := true // bool | Defaults to 'false'. If this is enabled, then Html text that is empty will be kept in the output. Otherwise, Html text without any plain text will be removed. This is useful for example if you want to keep texts that only consist of empty paragraphs in the output. (optional)
    allowUpperCaseItemNumbers := true // bool | Defaults to 'false'. If this is enabled, then the ItemNumber of positions will be in uppercase format if the source file has them. By default, all item numbers will be converted to lowercase, but this option will enable the option to support uppercase item numbers as well. (optional)
    allowLumpSumItemsWithDifferingQuantities := true // bool | Defaults to 'false'. By default, the GAEB standard requires lump sum items ('Pauschalpositionen' in German) to have a quantity of exactly 1. AVACloud does enforce this convention, but if you set this property to 'true', then differing quantities will be kept during the import. (optional)
    gaebFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaebConversionApi.GaebConversionConvertToExcel(context.Background()).SupportSkippedItemNumberLevelsInPositions(supportSkippedItemNumberLevelsInPositions).WritePrices(writePrices).WriteLongTexts(writeLongTexts).ConversionCulture(conversionCulture).IncludeArticleNumbers(includeArticleNumbers).OutputHtmlAsXml(outputHtmlAsXml).KeepEmptyHtmlText(keepEmptyHtmlText).AllowUpperCaseItemNumbers(allowUpperCaseItemNumbers).AllowLumpSumItemsWithDifferingQuantities(allowLumpSumItemsWithDifferingQuantities).GaebFile(gaebFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaebConversionApi.GaebConversionConvertToExcel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GaebConversionConvertToExcel`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GaebConversionApi.GaebConversionConvertToExcel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGaebConversionConvertToExcelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportSkippedItemNumberLevelsInPositions** | **bool** | Defaults to &#39;false&#39;. This controls if, when deserializing GAEB files, skipped levels in position item numbers should be supported. For example, if an ItemNumberSchema defines three levels - two group levels and one position level - but the ItemNumber of the position is just &#39;01.02&#39;, then it will be displayed as &#39;01.__.02&#39; if this is set to true. | 
 **writePrices** | **bool** | Defaults to true | 
 **writeLongTexts** | **bool** | Defaults to true | 
 **conversionCulture** | **string** | The culture that should be used for the conversion process, to have localized Excel files | 
 **includeArticleNumbers** | **bool** | If this is enabled, then a new column will be created in the overview worksheet that contains the article numbers for positions. Article numbers will be read from &#39;position.commerceProperties.articleNumber&#39; | 
 **outputHtmlAsXml** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then Html text will be output as XML in the output. This means that e.g. void Html tags will always be output with their closing tag, e.g. it will produce &#39;&lt;br /&gt;&#39; instead of &#39;&lt;br&gt;&#39;. | 
 **keepEmptyHtmlText** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then Html text that is empty will be kept in the output. Otherwise, Html text without any plain text will be removed. This is useful for example if you want to keep texts that only consist of empty paragraphs in the output. | 
 **allowUpperCaseItemNumbers** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then the ItemNumber of positions will be in uppercase format if the source file has them. By default, all item numbers will be converted to lowercase, but this option will enable the option to support uppercase item numbers as well. | 
 **allowLumpSumItemsWithDifferingQuantities** | **bool** | Defaults to &#39;false&#39;. By default, the GAEB standard requires lump sum items (&#39;Pauschalpositionen&#39; in German) to have a quantity of exactly 1. AVACloud does enforce this convention, but if you set this property to &#39;true&#39;, then differing quantities will be kept during the import. | 
 **gaebFile** | ***os.File** | The input file | 

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


## GaebConversionConvertToGaeb

> *os.File GaebConversionConvertToGaeb(ctx).SupportSkippedItemNumberLevelsInPositions(supportSkippedItemNumberLevelsInPositions).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).OutputHtmlAsXml(outputHtmlAsXml).KeepEmptyHtmlText(keepEmptyHtmlText).AllowUpperCaseItemNumbers(allowUpperCaseItemNumbers).AllowLumpSumItemsWithDifferingQuantities(allowLumpSumItemsWithDifferingQuantities).GaebFile(gaebFile).Execute()

Converts GAEB files to GAEB files. Used for example when transforming or repairing GAEB files.

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
    supportSkippedItemNumberLevelsInPositions := true // bool | Defaults to 'false'. This controls if, when deserializing GAEB files, skipped levels in position item numbers should be supported. For example, if an ItemNumberSchema defines three levels - two group levels and one position level - but the ItemNumber of the position is just '01.02', then it will be displayed as '01.__.02' if this is set to true. (optional)
    destinationGaebType := "destinationGaebType_example" // string | Defaults to GAEB XML V3.2 (optional)
    targetExchangePhaseTransform := "targetExchangePhaseTransform_example" // string | Defaults to none, meaning no transformation will be done. The phases are: Base = 81 CostEstimate = 82 OfferRequest = 83 Offer = 84 SideOffer = 85 Grant = 86 (optional)
    enforceStrictOfferPhaseLongTextOutput := true // bool | Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export. (optional)
    exportQuantityDetermination := true // bool | Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the 'QtyDeterm' (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the 'QuantityComponents' property of positions. Please see the entry for 'Quantity Determination' in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the 'QuantityComponents' property. (optional)
    removeUnprintableCharactersFromTexts := true // bool | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. (optional)
    forceIncludeDescriptions := true // bool | If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions. (optional)
    treatNullItemNumberSchemaAsInvalid := true // bool | When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons. (optional)
    outputHtmlAsXml := true // bool | Defaults to 'false'. If this is enabled, then Html text will be output as XML in the output. This means that e.g. void Html tags will always be output with their closing tag, e.g. it will produce '<br />' instead of '<br>'. (optional)
    keepEmptyHtmlText := true // bool | Defaults to 'false'. If this is enabled, then Html text that is empty will be kept in the output. Otherwise, Html text without any plain text will be removed. This is useful for example if you want to keep texts that only consist of empty paragraphs in the output. (optional)
    allowUpperCaseItemNumbers := true // bool | Defaults to 'false'. If this is enabled, then the ItemNumber of positions will be in uppercase format if the source file has them. By default, all item numbers will be converted to lowercase, but this option will enable the option to support uppercase item numbers as well. (optional)
    allowLumpSumItemsWithDifferingQuantities := true // bool | Defaults to 'false'. By default, the GAEB standard requires lump sum items ('Pauschalpositionen' in German) to have a quantity of exactly 1. AVACloud does enforce this convention, but if you set this property to 'true', then differing quantities will be kept during the import. (optional)
    gaebFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaebConversionApi.GaebConversionConvertToGaeb(context.Background()).SupportSkippedItemNumberLevelsInPositions(supportSkippedItemNumberLevelsInPositions).DestinationGaebType(destinationGaebType).TargetExchangePhaseTransform(targetExchangePhaseTransform).EnforceStrictOfferPhaseLongTextOutput(enforceStrictOfferPhaseLongTextOutput).ExportQuantityDetermination(exportQuantityDetermination).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).ForceIncludeDescriptions(forceIncludeDescriptions).TreatNullItemNumberSchemaAsInvalid(treatNullItemNumberSchemaAsInvalid).OutputHtmlAsXml(outputHtmlAsXml).KeepEmptyHtmlText(keepEmptyHtmlText).AllowUpperCaseItemNumbers(allowUpperCaseItemNumbers).AllowLumpSumItemsWithDifferingQuantities(allowLumpSumItemsWithDifferingQuantities).GaebFile(gaebFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaebConversionApi.GaebConversionConvertToGaeb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GaebConversionConvertToGaeb`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GaebConversionApi.GaebConversionConvertToGaeb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGaebConversionConvertToGaebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportSkippedItemNumberLevelsInPositions** | **bool** | Defaults to &#39;false&#39;. This controls if, when deserializing GAEB files, skipped levels in position item numbers should be supported. For example, if an ItemNumberSchema defines three levels - two group levels and one position level - but the ItemNumber of the position is just &#39;01.02&#39;, then it will be displayed as &#39;01.__.02&#39; if this is set to true. | 
 **destinationGaebType** | **string** | Defaults to GAEB XML V3.2 | 
 **targetExchangePhaseTransform** | **string** | Defaults to none, meaning no transformation will be done. The phases are: Base &#x3D; 81 CostEstimate &#x3D; 82 OfferRequest &#x3D; 83 Offer &#x3D; 84 SideOffer &#x3D; 85 Grant &#x3D; 86 | 
 **enforceStrictOfferPhaseLongTextOutput** | **bool** | Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export. | 
 **exportQuantityDetermination** | **bool** | Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the &#39;QtyDeterm&#39; (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the &#39;QuantityComponents&#39; property of positions. Please see the entry for &#39;Quantity Determination&#39; in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the &#39;QuantityComponents&#39; property. | 
 **removeUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 
 **forceIncludeDescriptions** | **bool** | If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions. | 
 **treatNullItemNumberSchemaAsInvalid** | **bool** | When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons. | 
 **outputHtmlAsXml** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then Html text will be output as XML in the output. This means that e.g. void Html tags will always be output with their closing tag, e.g. it will produce &#39;&lt;br /&gt;&#39; instead of &#39;&lt;br&gt;&#39;. | 
 **keepEmptyHtmlText** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then Html text that is empty will be kept in the output. Otherwise, Html text without any plain text will be removed. This is useful for example if you want to keep texts that only consist of empty paragraphs in the output. | 
 **allowUpperCaseItemNumbers** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then the ItemNumber of positions will be in uppercase format if the source file has them. By default, all item numbers will be converted to lowercase, but this option will enable the option to support uppercase item numbers as well. | 
 **allowLumpSumItemsWithDifferingQuantities** | **bool** | Defaults to &#39;false&#39;. By default, the GAEB standard requires lump sum items (&#39;Pauschalpositionen&#39; in German) to have a quantity of exactly 1. AVACloud does enforce this convention, but if you set this property to &#39;true&#39;, then differing quantities will be kept during the import. | 
 **gaebFile** | ***os.File** | The input file | 

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


## GaebConversionConvertToOenorm

> *os.File GaebConversionConvertToOenorm(ctx).SupportSkippedItemNumberLevelsInPositions(supportSkippedItemNumberLevelsInPositions).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).OutputHtmlAsXml(outputHtmlAsXml).KeepEmptyHtmlText(keepEmptyHtmlText).AllowUpperCaseItemNumbers(allowUpperCaseItemNumbers).AllowLumpSumItemsWithDifferingQuantities(allowLumpSumItemsWithDifferingQuantities).GaebFile(gaebFile).Execute()

Converts GAEB files to Oenorm files

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
    supportSkippedItemNumberLevelsInPositions := true // bool | Defaults to 'false'. This controls if, when deserializing GAEB files, skipped levels in position item numbers should be supported. For example, if an ItemNumberSchema defines three levels - two group levels and one position level - but the ItemNumber of the position is just '01.02', then it will be displayed as '01.__.02' if this is set to true. (optional)
    destinationOenormType := "destinationOenormType_example" // string | Defaults to Lv2015 (optional)
    tryRepairProjectStructure := true // bool | Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target (optional)
    skipTryEnforceSchemaCompliantXmlOutput := true // bool | If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option. (optional)
    removeUnprintableCharactersFromTexts := true // bool | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. (optional)
    outputHtmlAsXml := true // bool | Defaults to 'false'. If this is enabled, then Html text will be output as XML in the output. This means that e.g. void Html tags will always be output with their closing tag, e.g. it will produce '<br />' instead of '<br>'. (optional)
    keepEmptyHtmlText := true // bool | Defaults to 'false'. If this is enabled, then Html text that is empty will be kept in the output. Otherwise, Html text without any plain text will be removed. This is useful for example if you want to keep texts that only consist of empty paragraphs in the output. (optional)
    allowUpperCaseItemNumbers := true // bool | Defaults to 'false'. If this is enabled, then the ItemNumber of positions will be in uppercase format if the source file has them. By default, all item numbers will be converted to lowercase, but this option will enable the option to support uppercase item numbers as well. (optional)
    allowLumpSumItemsWithDifferingQuantities := true // bool | Defaults to 'false'. By default, the GAEB standard requires lump sum items ('Pauschalpositionen' in German) to have a quantity of exactly 1. AVACloud does enforce this convention, but if you set this property to 'true', then differing quantities will be kept during the import. (optional)
    gaebFile := os.NewFile(1234, "some_file") // *os.File | The input file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GaebConversionApi.GaebConversionConvertToOenorm(context.Background()).SupportSkippedItemNumberLevelsInPositions(supportSkippedItemNumberLevelsInPositions).DestinationOenormType(destinationOenormType).TryRepairProjectStructure(tryRepairProjectStructure).SkipTryEnforceSchemaCompliantXmlOutput(skipTryEnforceSchemaCompliantXmlOutput).RemoveUnprintableCharactersFromTexts(removeUnprintableCharactersFromTexts).OutputHtmlAsXml(outputHtmlAsXml).KeepEmptyHtmlText(keepEmptyHtmlText).AllowUpperCaseItemNumbers(allowUpperCaseItemNumbers).AllowLumpSumItemsWithDifferingQuantities(allowLumpSumItemsWithDifferingQuantities).GaebFile(gaebFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GaebConversionApi.GaebConversionConvertToOenorm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GaebConversionConvertToOenorm`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GaebConversionApi.GaebConversionConvertToOenorm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGaebConversionConvertToOenormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportSkippedItemNumberLevelsInPositions** | **bool** | Defaults to &#39;false&#39;. This controls if, when deserializing GAEB files, skipped levels in position item numbers should be supported. For example, if an ItemNumberSchema defines three levels - two group levels and one position level - but the ItemNumber of the position is just &#39;01.02&#39;, then it will be displayed as &#39;01.__.02&#39; if this is set to true. | 
 **destinationOenormType** | **string** | Defaults to Lv2015 | 
 **tryRepairProjectStructure** | **bool** | Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 
 **skipTryEnforceSchemaCompliantXmlOutput** | **bool** | If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option. | 
 **removeUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 
 **outputHtmlAsXml** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then Html text will be output as XML in the output. This means that e.g. void Html tags will always be output with their closing tag, e.g. it will produce &#39;&lt;br /&gt;&#39; instead of &#39;&lt;br&gt;&#39;. | 
 **keepEmptyHtmlText** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then Html text that is empty will be kept in the output. Otherwise, Html text without any plain text will be removed. This is useful for example if you want to keep texts that only consist of empty paragraphs in the output. | 
 **allowUpperCaseItemNumbers** | **bool** | Defaults to &#39;false&#39;. If this is enabled, then the ItemNumber of positions will be in uppercase format if the source file has them. By default, all item numbers will be converted to lowercase, but this option will enable the option to support uppercase item numbers as well. | 
 **allowLumpSumItemsWithDifferingQuantities** | **bool** | Defaults to &#39;false&#39;. By default, the GAEB standard requires lump sum items (&#39;Pauschalpositionen&#39; in German) to have a quantity of exactly 1. AVACloud does enforce this convention, but if you set this property to &#39;true&#39;, then differing quantities will be kept during the import. | 
 **gaebFile** | ***os.File** | The input file | 

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

