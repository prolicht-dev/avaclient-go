# ServiceSpecificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
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
**Bidder** | [**PartyInformationDto**](PartyInformationDto.md) |  | [optional] 
**BidderDiscriminator** | **string** | This discriminator is used to identify different bidders in a project. It is different from the Identifier property in the Bidder in that the BidderDiscriminator is intended to be a numerical identifier within a project, while the Identifier does uniquely identify a bidder in the system independent of a specific project. This property should map to \&quot;Bieternummer\&quot; or \&quot;BidderNo\&quot; in GAEB. | [optional] 
**GaebXmlId** | **string** | This is used to store the GAEB XML Id within this ServiceSpecification. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization. | [optional] 
**ProjectInformation** | [**ProjectInformationDto**](ProjectInformationDto.md) |  | [optional] 
**ExchangePhase** | [**ExchangePhaseDto**](ExchangePhaseDto.md) |  | 
**Origin** | [**OriginDto**](OriginDto.md) |  | 
**CreationDate** | [**time.Time**](time.Time.md) | Creation date of this ServiceSpecification. | [optional] 
**OfferByDate** | [**time.Time**](time.Time.md) | Date indicating until when an offer has to be submitted. In German, this is often called the \&quot;Eröffnungstermin\&quot; | [optional] 
**DecisionDate** | [**time.Time**](time.Time.md) | Date indicating by when the buyer will select a contractor. | [optional] 
**SubmissionLocation** | **string** | String indicating where the physical submission of the offer is taking place. | [optional] 
**Description** | **string** | Description of this ServiceSpecification. | [optional] 
**Name** | **string** | The name of this ServiceSpecification. | [optional] 
**PriceInformation** | [**PriceInformationDto**](PriceInformationDto.md) |  | [optional] 
**ProjectCatalogues** | [**[]CatalogueDto**](CatalogueDto.md) | These are Catalogue that are used within this ElementContainerBase. Catalogue references are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. | [optional] 
**CatalogueReferences** | [**[]CatalogueReferenceDto**](CatalogueReferenceDto.md) | Referenced catalogues for this QuantityAssignment. | [optional] 
**PlannedExecutionStart** | [**time.Time**](time.Time.md) | The date when the execution of the services is scheduled to start | [optional] 
**PlannedExecutionEnd** | [**time.Time**](time.Time.md) | The date then the execution of the services is scheduled to be finished | [optional] 
**ContractDate** | [**time.Time**](time.Time.md) | The date on which the contract has been awarded. This matches \&quot;Auftragsdatum\&quot; in GAEB | [optional] 
**ContractIdentifier** | **string** | This value can be used to indicate the number or identifier of the contract. It matches \&quot;Auftragsnummer\&quot; in GAEB | [optional] 
**WarrantyDuration** | [**WarrantyDurationDto**](WarrantyDurationDto.md) |  | [optional] 
**WarrantyEnd** | [**time.Time**](time.Time.md) | The date on which the warranty period ends | [optional] 
**ApprovalDate** | [**time.Time**](time.Time.md) | The date on which the services rendered by the bidder are scheduled to be approved by the buyer | [optional] 
**TypeOfApproval** | **string** | This should specify how the approval is performed by the buyer. This matches \&quot;AcceptType\&quot; in GAEB | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


