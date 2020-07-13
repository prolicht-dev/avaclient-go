# ServiceSpecificationGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**GaebXmlId** | **string** | This is used to store the GAEB XML Id within this IElement. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization. | [optional] 
**ElementTypeDiscriminator** | **string** |  | 
**ProjectHourlyWage** | **float32** | The hourly wage that is used within this ElementContainerBase. Will be propagated to child elements. | [readonly] 
**ProjectTaxRate** | **float32** | The tax rate that is used within this ElementContainerBase. Will be propagated to child elements. | 
**ProjectPriceComponents** | **[]string** | The price components that are used within this project. They are ignored during Json deserialization because they will be set from the parent project. | [optional] 
**ProjectItemNumberSchema** | [**ItemNumberSchemaDto**](ItemNumberSchemaDto.md) |  | [optional] 
**Elements** | [**[]IElementDto**](IElementDto.md) | The IElements within this ElementContainerBase. | [optional] 
**ProjectLabourTimeLabel** | **string** | The label used in the parent Project to mark labour time, e.g. \&quot;Hours\&quot; or \&quot;Stunden\&quot;. | [optional] [readonly] 
**ContainsDuplicateItemNumbers** | **bool** | Indicates if there are child IElements that have conflicting, duplicated ItemNumbers or if any child ElementContainerBase elements themselves contain duplicate ItemNumber s. Will always indicate false when told to ignore duplicate item numbers. | [readonly] 
**ContainsDuplicateElementIds** | **bool** | Indicates if there are child IElements that have conflicting, duplicated Ids or if any child ElementContainerBase elements themselves contain duplicate Id s. Will always indicate false when told to ignore duplicate item numbers. | [readonly] 
**IgnoreDuplicateItemNumbers** | **bool** | Indicate if duplicated ItemNumbers within child elements are to be ignored. Will not perform checks for duplicates if yes. | 
**IgnoreDuplicateElementIds** | **bool** | Indicate if duplicated Ids within child elements are to be ignored. Will not perform checks for duplicates if yes. | 
**TotalPriceGrossByTaxRate** | [**[]GrossPriceComponentDto**](GrossPriceComponentDto.md) | Price composition by tax rate. | [optional] 
**IgnoreChildPriceUpdates** | **bool** | Internally used to indicate that a propagation is currently done, this is done to not recalculate every single result from a lot of changes when it is sufficient to calculate the total price at once. | 
**DeductedPrice** | **float32** | Net price after applied deductions. | [readonly] 
**DeductionFactor** | **float32** | Factor of applied deductions to the total price. For example, \&quot;0.03\&quot; means that a 3% deduction is to be applied. | 
**TotalPrice** | **float32** | Will return this ElementContainerBase&#39;s total price. | [readonly] 
**TotalPriceGross** | **float32** | The total gross price for this ElementContainerBase including all child elements. | [readonly] 
**TotalPriceGrossDeducted** | **float32** | Total gross price after applied deductions. | [readonly] 
**PriceType** | [**PriceTypeDto**](PriceTypeDto.md) |  | 
**ShortText** | **string** | Description for this ServiceSpecificationGroup. | [optional] 
**ComissionStatus** | [**ComissionStatusDto**](ComissionStatusDto.md) |  | 
**ItemNumber** | [**ItemNumberDto**](ItemNumberDto.md) |  | [optional] 
**ElementType** | **string** |  | [optional] 
**IsLot** | **bool** | This indicates if this group is the parent of a lot. See the documentation for more information about working with lots. | [readonly] 
**AlternativeTo** | **string** | If this group is an alternative for a base group, then this property should point to the id of the group in this service specification that it can replace. If this is an alternative group to a base group, the PriceType should typically be set to \&quot;WithoutTotal\&quot; so this group does not factor into total costs. The PriceType is not changed when this property is set | [optional] 
**Type** | **string** |  | [optional] 
**ProjectCatalogues** | [**[]CatalogueDto**](CatalogueDto.md) | These are Catalogue that are used within this ElementContainerBase. Catalogue references are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. | [optional] 
**CatalogueReferences** | [**[]CatalogueReferenceDto**](CatalogueReferenceDto.md) | Referenced catalogues for this QuantityAssignment. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


