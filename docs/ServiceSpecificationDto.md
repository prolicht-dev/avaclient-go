# ServiceSpecificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**ProjectHourlyWage** | **float32** | The hourly wage that is used within this ElementContainerBase. Will be propagated to child elements. | [readonly] 
**ProjectTaxRate** | **float32** | The tax rate that is used within this ElementContainerBase. Will be propagated to child elements. | 
**ProjectPriceComponents** | Pointer to **[]string** | The price components that are used within this project. They are ignored during Json deserialization because they will be set from the parent project. | [optional] 
**ProjectItemNumberSchema** | Pointer to [**ItemNumberSchemaDto**](ItemNumberSchemaDto.md) |  | [optional] 
**Elements** | Pointer to [**[]IElementDto**](IElementDto.md) | The IElements within this ElementContainerBase. | [optional] 
**ProjectLabourTimeLabel** | Pointer to **string** | The label used in the parent Project to mark labour time, e.g. \&quot;Hours\&quot; or \&quot;Stunden\&quot;. | [optional] [readonly] 
**ContainsDuplicateItemNumbers** | **bool** | Indicates if there are child IElements that have conflicting, duplicated ItemNumbers or if any child ElementContainerBase elements themselves contain duplicate ItemNumber s. Will always indicate false when told to ignore duplicate item numbers. | [readonly] 
**ContainsDuplicateElementIds** | **bool** | Indicates if there are child IElements that have conflicting, duplicated Ids or if any child ElementContainerBase elements themselves contain duplicate Id s. Will always indicate false when told to ignore duplicate item numbers. | [readonly] 
**IgnoreDuplicateItemNumbers** | **bool** | Indicate if duplicated ItemNumbers within child elements are to be ignored. Will not perform checks for duplicates if yes. | 
**IgnoreProjectCataloguePropagation** | **bool** | If this is set to true, the ProjectCatalogues property will not be propagated to child elements. This is useful in mapping scenarios, where you want to disable propagation for multiple changes, and only enable it once you have mapped all properties | 
**IgnoreDuplicateElementIds** | **bool** | Indicate if duplicated Ids within child elements are to be ignored. Will not perform checks for duplicates if yes. | 
**TotalPriceGrossByTaxRate** | Pointer to [**[]GrossPriceComponentDto**](GrossPriceComponentDto.md) | Price composition by tax rate. | [optional] 
**IgnoreChildPriceUpdates** | **bool** | Internally used to indicate that a propagation is currently done, this is done to not recalculate every single result from a lot of changes when it is sufficient to calculate the total price at once. | 
**DeductedPrice** | **float32** | Net price after applied deductions. | [readonly] 
**DeductionFactor** | **float32** | Factor of applied deductions to the total price. For example, \&quot;0.03\&quot; means that a 3% deduction is to be applied. | 
**AbsoluteDeduction** | Pointer to **float32** | The exact amount of the discount as an absolute value. For backwards compatibility reasons, setting this value will also set a calculated value to DeductionFactor, which will also be updated in case the total price is changed to reflect a relative value of the absolute discount sum. | [optional] 
**TotalPrice** | **float32** | Will return this ElementContainerBase&#39;s total price. | [readonly] 
**TotalPriceGross** | **float32** | The total gross price for this ElementContainerBase including all child elements. | [readonly] 
**TotalPriceGrossDeducted** | **float32** | Total gross price after applied deductions. | [readonly] 
**PriceType** | [**PriceTypeDto**](PriceTypeDto.md) |  | 
**Bidder** | Pointer to [**PartyInformationDto**](PartyInformationDto.md) |  | [optional] 
**BidderDiscriminator** | Pointer to **string** | This discriminator is used to identify different bidders in a project. It is different from the Identifier property in the Bidder in that the BidderDiscriminator is intended to be a numerical identifier within a project, while the Identifier does uniquely identify a bidder in the system independent of a specific project. This property should map to \&quot;Bieternummer\&quot; or \&quot;BidderNo\&quot; in GAEB. | [optional] 
**GaebXmlId** | Pointer to **string** | This is used to store the GAEB XML Id within this ServiceSpecification. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization. | [optional] 
**ProjectInformation** | Pointer to [**ProjectInformationDto**](ProjectInformationDto.md) |  | [optional] 
**ExchangePhase** | [**ExchangePhaseDto**](ExchangePhaseDto.md) |  | 
**Origin** | [**OriginDto**](OriginDto.md) |  | 
**OriginDetail** | Pointer to **string** | This property complements the Origin property. Some formats, e.g. GaebXml might have additional information attached, e.g. GAEB XML V3.2 oder GAEB XML V3.1. This string property is used to hold such data. The Dangl.AVA module does not have a list of possible values, since this is defined in the native format assemblies, e.g. Dangl.GAEB or Dangl.Oenorm. | [optional] 
**CreationDate** | Pointer to **time.Time** | Creation date of this ServiceSpecification. | [optional] 
**OfferByDate** | Pointer to **time.Time** | Date indicating until when an offer has to be submitted. In German, this is often called the \&quot;Eröffnungstermin\&quot; | [optional] 
**DecisionDate** | Pointer to **time.Time** | Date indicating by when the buyer will select a contractor. | [optional] 
**BidDate** | Pointer to **time.Time** | Date of when the offer / bid was created | [optional] 
**WarrantyBondPercentage** | **float32** | If a construction project requires a warranty bond, this property indicates the amount as a percentage, e.g. &#39;0.15m&#39; means 15% of the construction total amount as a warranty bond. | 
**ExecutionGuaranteePercentage** | **float32** | If a construction project requires an execution guarantty, this property indicates the amount as a percentage, e.g. &#39;0.15m&#39; means 15% of the construction total amount as an execution guarantee. In practice, this percentage is usually deducted from intermediate invoices and only billed in the final invoice. | 
**SubmissionLocation** | Pointer to **string** | String indicating where the physical submission of the offer is taking place. | [optional] 
**Description** | Pointer to **string** | Description of this ServiceSpecification. | [optional] 
**Name** | Pointer to **string** | The name of this ServiceSpecification. | [optional] 
**PriceInformation** | Pointer to [**PriceInformationDto**](PriceInformationDto.md) |  | [optional] 
**ProjectCatalogues** | Pointer to [**[]CatalogueDto**](CatalogueDto.md) | These are Catalogue that are used within this ElementContainerBase. Catalogue references are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. | [optional] 
**CatalogueReferences** | Pointer to [**[]CatalogueReferenceDto**](CatalogueReferenceDto.md) | Referenced catalogues for this ElementContainerBase. | [optional] 
**PlannedExecutionStart** | Pointer to **time.Time** | The date when the execution of the services is scheduled to start | [optional] 
**PlannedExecutionEnd** | Pointer to **time.Time** | The date then the execution of the services is scheduled to be finished | [optional] 
**ContractDate** | Pointer to **time.Time** | The date on which the contract has been awarded. This matches \&quot;Auftragsdatum\&quot; in GAEB | [optional] 
**ContractIdentifier** | Pointer to **string** | This value can be used to indicate the number or identifier of the contract. It matches \&quot;Auftragsnummer\&quot; in GAEB | [optional] 
**WarrantyDuration** | Pointer to [**WarrantyDurationDto**](WarrantyDurationDto.md) |  | [optional] 
**WarrantyEnd** | Pointer to **time.Time** | The date on which the warranty period ends | [optional] 
**ApprovalDate** | Pointer to **time.Time** | The date on which the services rendered by the bidder are scheduled to be approved by the buyer | [optional] 
**TypeOfApproval** | Pointer to **string** | This should specify how the approval is performed by the buyer. This matches \&quot;AcceptType\&quot; in GAEB | [optional] 
**AddendumNumber** | Pointer to **string** | This optional string property is shared by all IElements, and indicates if the element is part of an addendum, a &#39;Nachtrag&#39; in German. | [optional] 
**AddendumStatus** | Pointer to [**AddendumStatusDto**](AddendumStatusDto.md) |  | [optional] 
**CommerceProperties** | Pointer to [**ServiceSpecificationCommercePropertiesDto**](ServiceSpecificationCommercePropertiesDto.md) |  | [optional] 

## Methods

### NewServiceSpecificationDto

`func NewServiceSpecificationDto(id string, projectHourlyWage float32, projectTaxRate float32, containsDuplicateItemNumbers bool, containsDuplicateElementIds bool, ignoreDuplicateItemNumbers bool, ignoreProjectCataloguePropagation bool, ignoreDuplicateElementIds bool, ignoreChildPriceUpdates bool, deductedPrice float32, deductionFactor float32, totalPrice float32, totalPriceGross float32, totalPriceGrossDeducted float32, priceType PriceTypeDto, exchangePhase ExchangePhaseDto, origin OriginDto, warrantyBondPercentage float32, executionGuaranteePercentage float32, ) *ServiceSpecificationDto`

NewServiceSpecificationDto instantiates a new ServiceSpecificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSpecificationDtoWithDefaults

`func NewServiceSpecificationDtoWithDefaults() *ServiceSpecificationDto`

NewServiceSpecificationDtoWithDefaults instantiates a new ServiceSpecificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceSpecificationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceSpecificationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceSpecificationDto) SetId(v string)`

SetId sets Id field to given value.


### GetProjectHourlyWage

`func (o *ServiceSpecificationDto) GetProjectHourlyWage() float32`

GetProjectHourlyWage returns the ProjectHourlyWage field if non-nil, zero value otherwise.

### GetProjectHourlyWageOk

`func (o *ServiceSpecificationDto) GetProjectHourlyWageOk() (*float32, bool)`

GetProjectHourlyWageOk returns a tuple with the ProjectHourlyWage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectHourlyWage

`func (o *ServiceSpecificationDto) SetProjectHourlyWage(v float32)`

SetProjectHourlyWage sets ProjectHourlyWage field to given value.


### GetProjectTaxRate

`func (o *ServiceSpecificationDto) GetProjectTaxRate() float32`

GetProjectTaxRate returns the ProjectTaxRate field if non-nil, zero value otherwise.

### GetProjectTaxRateOk

`func (o *ServiceSpecificationDto) GetProjectTaxRateOk() (*float32, bool)`

GetProjectTaxRateOk returns a tuple with the ProjectTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTaxRate

`func (o *ServiceSpecificationDto) SetProjectTaxRate(v float32)`

SetProjectTaxRate sets ProjectTaxRate field to given value.


### GetProjectPriceComponents

`func (o *ServiceSpecificationDto) GetProjectPriceComponents() []string`

GetProjectPriceComponents returns the ProjectPriceComponents field if non-nil, zero value otherwise.

### GetProjectPriceComponentsOk

`func (o *ServiceSpecificationDto) GetProjectPriceComponentsOk() (*[]string, bool)`

GetProjectPriceComponentsOk returns a tuple with the ProjectPriceComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPriceComponents

`func (o *ServiceSpecificationDto) SetProjectPriceComponents(v []string)`

SetProjectPriceComponents sets ProjectPriceComponents field to given value.

### HasProjectPriceComponents

`func (o *ServiceSpecificationDto) HasProjectPriceComponents() bool`

HasProjectPriceComponents returns a boolean if a field has been set.

### GetProjectItemNumberSchema

`func (o *ServiceSpecificationDto) GetProjectItemNumberSchema() ItemNumberSchemaDto`

GetProjectItemNumberSchema returns the ProjectItemNumberSchema field if non-nil, zero value otherwise.

### GetProjectItemNumberSchemaOk

`func (o *ServiceSpecificationDto) GetProjectItemNumberSchemaOk() (*ItemNumberSchemaDto, bool)`

GetProjectItemNumberSchemaOk returns a tuple with the ProjectItemNumberSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectItemNumberSchema

`func (o *ServiceSpecificationDto) SetProjectItemNumberSchema(v ItemNumberSchemaDto)`

SetProjectItemNumberSchema sets ProjectItemNumberSchema field to given value.

### HasProjectItemNumberSchema

`func (o *ServiceSpecificationDto) HasProjectItemNumberSchema() bool`

HasProjectItemNumberSchema returns a boolean if a field has been set.

### GetElements

`func (o *ServiceSpecificationDto) GetElements() []IElementDto`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *ServiceSpecificationDto) GetElementsOk() (*[]IElementDto, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *ServiceSpecificationDto) SetElements(v []IElementDto)`

SetElements sets Elements field to given value.

### HasElements

`func (o *ServiceSpecificationDto) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetProjectLabourTimeLabel

`func (o *ServiceSpecificationDto) GetProjectLabourTimeLabel() string`

GetProjectLabourTimeLabel returns the ProjectLabourTimeLabel field if non-nil, zero value otherwise.

### GetProjectLabourTimeLabelOk

`func (o *ServiceSpecificationDto) GetProjectLabourTimeLabelOk() (*string, bool)`

GetProjectLabourTimeLabelOk returns a tuple with the ProjectLabourTimeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLabourTimeLabel

`func (o *ServiceSpecificationDto) SetProjectLabourTimeLabel(v string)`

SetProjectLabourTimeLabel sets ProjectLabourTimeLabel field to given value.

### HasProjectLabourTimeLabel

`func (o *ServiceSpecificationDto) HasProjectLabourTimeLabel() bool`

HasProjectLabourTimeLabel returns a boolean if a field has been set.

### GetContainsDuplicateItemNumbers

`func (o *ServiceSpecificationDto) GetContainsDuplicateItemNumbers() bool`

GetContainsDuplicateItemNumbers returns the ContainsDuplicateItemNumbers field if non-nil, zero value otherwise.

### GetContainsDuplicateItemNumbersOk

`func (o *ServiceSpecificationDto) GetContainsDuplicateItemNumbersOk() (*bool, bool)`

GetContainsDuplicateItemNumbersOk returns a tuple with the ContainsDuplicateItemNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsDuplicateItemNumbers

`func (o *ServiceSpecificationDto) SetContainsDuplicateItemNumbers(v bool)`

SetContainsDuplicateItemNumbers sets ContainsDuplicateItemNumbers field to given value.


### GetContainsDuplicateElementIds

`func (o *ServiceSpecificationDto) GetContainsDuplicateElementIds() bool`

GetContainsDuplicateElementIds returns the ContainsDuplicateElementIds field if non-nil, zero value otherwise.

### GetContainsDuplicateElementIdsOk

`func (o *ServiceSpecificationDto) GetContainsDuplicateElementIdsOk() (*bool, bool)`

GetContainsDuplicateElementIdsOk returns a tuple with the ContainsDuplicateElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsDuplicateElementIds

`func (o *ServiceSpecificationDto) SetContainsDuplicateElementIds(v bool)`

SetContainsDuplicateElementIds sets ContainsDuplicateElementIds field to given value.


### GetIgnoreDuplicateItemNumbers

`func (o *ServiceSpecificationDto) GetIgnoreDuplicateItemNumbers() bool`

GetIgnoreDuplicateItemNumbers returns the IgnoreDuplicateItemNumbers field if non-nil, zero value otherwise.

### GetIgnoreDuplicateItemNumbersOk

`func (o *ServiceSpecificationDto) GetIgnoreDuplicateItemNumbersOk() (*bool, bool)`

GetIgnoreDuplicateItemNumbersOk returns a tuple with the IgnoreDuplicateItemNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDuplicateItemNumbers

`func (o *ServiceSpecificationDto) SetIgnoreDuplicateItemNumbers(v bool)`

SetIgnoreDuplicateItemNumbers sets IgnoreDuplicateItemNumbers field to given value.


### GetIgnoreProjectCataloguePropagation

`func (o *ServiceSpecificationDto) GetIgnoreProjectCataloguePropagation() bool`

GetIgnoreProjectCataloguePropagation returns the IgnoreProjectCataloguePropagation field if non-nil, zero value otherwise.

### GetIgnoreProjectCataloguePropagationOk

`func (o *ServiceSpecificationDto) GetIgnoreProjectCataloguePropagationOk() (*bool, bool)`

GetIgnoreProjectCataloguePropagationOk returns a tuple with the IgnoreProjectCataloguePropagation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreProjectCataloguePropagation

`func (o *ServiceSpecificationDto) SetIgnoreProjectCataloguePropagation(v bool)`

SetIgnoreProjectCataloguePropagation sets IgnoreProjectCataloguePropagation field to given value.


### GetIgnoreDuplicateElementIds

`func (o *ServiceSpecificationDto) GetIgnoreDuplicateElementIds() bool`

GetIgnoreDuplicateElementIds returns the IgnoreDuplicateElementIds field if non-nil, zero value otherwise.

### GetIgnoreDuplicateElementIdsOk

`func (o *ServiceSpecificationDto) GetIgnoreDuplicateElementIdsOk() (*bool, bool)`

GetIgnoreDuplicateElementIdsOk returns a tuple with the IgnoreDuplicateElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDuplicateElementIds

`func (o *ServiceSpecificationDto) SetIgnoreDuplicateElementIds(v bool)`

SetIgnoreDuplicateElementIds sets IgnoreDuplicateElementIds field to given value.


### GetTotalPriceGrossByTaxRate

`func (o *ServiceSpecificationDto) GetTotalPriceGrossByTaxRate() []GrossPriceComponentDto`

GetTotalPriceGrossByTaxRate returns the TotalPriceGrossByTaxRate field if non-nil, zero value otherwise.

### GetTotalPriceGrossByTaxRateOk

`func (o *ServiceSpecificationDto) GetTotalPriceGrossByTaxRateOk() (*[]GrossPriceComponentDto, bool)`

GetTotalPriceGrossByTaxRateOk returns a tuple with the TotalPriceGrossByTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceGrossByTaxRate

`func (o *ServiceSpecificationDto) SetTotalPriceGrossByTaxRate(v []GrossPriceComponentDto)`

SetTotalPriceGrossByTaxRate sets TotalPriceGrossByTaxRate field to given value.

### HasTotalPriceGrossByTaxRate

`func (o *ServiceSpecificationDto) HasTotalPriceGrossByTaxRate() bool`

HasTotalPriceGrossByTaxRate returns a boolean if a field has been set.

### GetIgnoreChildPriceUpdates

`func (o *ServiceSpecificationDto) GetIgnoreChildPriceUpdates() bool`

GetIgnoreChildPriceUpdates returns the IgnoreChildPriceUpdates field if non-nil, zero value otherwise.

### GetIgnoreChildPriceUpdatesOk

`func (o *ServiceSpecificationDto) GetIgnoreChildPriceUpdatesOk() (*bool, bool)`

GetIgnoreChildPriceUpdatesOk returns a tuple with the IgnoreChildPriceUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreChildPriceUpdates

`func (o *ServiceSpecificationDto) SetIgnoreChildPriceUpdates(v bool)`

SetIgnoreChildPriceUpdates sets IgnoreChildPriceUpdates field to given value.


### GetDeductedPrice

`func (o *ServiceSpecificationDto) GetDeductedPrice() float32`

GetDeductedPrice returns the DeductedPrice field if non-nil, zero value otherwise.

### GetDeductedPriceOk

`func (o *ServiceSpecificationDto) GetDeductedPriceOk() (*float32, bool)`

GetDeductedPriceOk returns a tuple with the DeductedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductedPrice

`func (o *ServiceSpecificationDto) SetDeductedPrice(v float32)`

SetDeductedPrice sets DeductedPrice field to given value.


### GetDeductionFactor

`func (o *ServiceSpecificationDto) GetDeductionFactor() float32`

GetDeductionFactor returns the DeductionFactor field if non-nil, zero value otherwise.

### GetDeductionFactorOk

`func (o *ServiceSpecificationDto) GetDeductionFactorOk() (*float32, bool)`

GetDeductionFactorOk returns a tuple with the DeductionFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductionFactor

`func (o *ServiceSpecificationDto) SetDeductionFactor(v float32)`

SetDeductionFactor sets DeductionFactor field to given value.


### GetAbsoluteDeduction

`func (o *ServiceSpecificationDto) GetAbsoluteDeduction() float32`

GetAbsoluteDeduction returns the AbsoluteDeduction field if non-nil, zero value otherwise.

### GetAbsoluteDeductionOk

`func (o *ServiceSpecificationDto) GetAbsoluteDeductionOk() (*float32, bool)`

GetAbsoluteDeductionOk returns a tuple with the AbsoluteDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteDeduction

`func (o *ServiceSpecificationDto) SetAbsoluteDeduction(v float32)`

SetAbsoluteDeduction sets AbsoluteDeduction field to given value.

### HasAbsoluteDeduction

`func (o *ServiceSpecificationDto) HasAbsoluteDeduction() bool`

HasAbsoluteDeduction returns a boolean if a field has been set.

### GetTotalPrice

`func (o *ServiceSpecificationDto) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *ServiceSpecificationDto) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *ServiceSpecificationDto) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.


### GetTotalPriceGross

`func (o *ServiceSpecificationDto) GetTotalPriceGross() float32`

GetTotalPriceGross returns the TotalPriceGross field if non-nil, zero value otherwise.

### GetTotalPriceGrossOk

`func (o *ServiceSpecificationDto) GetTotalPriceGrossOk() (*float32, bool)`

GetTotalPriceGrossOk returns a tuple with the TotalPriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceGross

`func (o *ServiceSpecificationDto) SetTotalPriceGross(v float32)`

SetTotalPriceGross sets TotalPriceGross field to given value.


### GetTotalPriceGrossDeducted

`func (o *ServiceSpecificationDto) GetTotalPriceGrossDeducted() float32`

GetTotalPriceGrossDeducted returns the TotalPriceGrossDeducted field if non-nil, zero value otherwise.

### GetTotalPriceGrossDeductedOk

`func (o *ServiceSpecificationDto) GetTotalPriceGrossDeductedOk() (*float32, bool)`

GetTotalPriceGrossDeductedOk returns a tuple with the TotalPriceGrossDeducted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceGrossDeducted

`func (o *ServiceSpecificationDto) SetTotalPriceGrossDeducted(v float32)`

SetTotalPriceGrossDeducted sets TotalPriceGrossDeducted field to given value.


### GetPriceType

`func (o *ServiceSpecificationDto) GetPriceType() PriceTypeDto`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *ServiceSpecificationDto) GetPriceTypeOk() (*PriceTypeDto, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *ServiceSpecificationDto) SetPriceType(v PriceTypeDto)`

SetPriceType sets PriceType field to given value.


### GetBidder

`func (o *ServiceSpecificationDto) GetBidder() PartyInformationDto`

GetBidder returns the Bidder field if non-nil, zero value otherwise.

### GetBidderOk

`func (o *ServiceSpecificationDto) GetBidderOk() (*PartyInformationDto, bool)`

GetBidderOk returns a tuple with the Bidder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidder

`func (o *ServiceSpecificationDto) SetBidder(v PartyInformationDto)`

SetBidder sets Bidder field to given value.

### HasBidder

`func (o *ServiceSpecificationDto) HasBidder() bool`

HasBidder returns a boolean if a field has been set.

### GetBidderDiscriminator

`func (o *ServiceSpecificationDto) GetBidderDiscriminator() string`

GetBidderDiscriminator returns the BidderDiscriminator field if non-nil, zero value otherwise.

### GetBidderDiscriminatorOk

`func (o *ServiceSpecificationDto) GetBidderDiscriminatorOk() (*string, bool)`

GetBidderDiscriminatorOk returns a tuple with the BidderDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidderDiscriminator

`func (o *ServiceSpecificationDto) SetBidderDiscriminator(v string)`

SetBidderDiscriminator sets BidderDiscriminator field to given value.

### HasBidderDiscriminator

`func (o *ServiceSpecificationDto) HasBidderDiscriminator() bool`

HasBidderDiscriminator returns a boolean if a field has been set.

### GetGaebXmlId

`func (o *ServiceSpecificationDto) GetGaebXmlId() string`

GetGaebXmlId returns the GaebXmlId field if non-nil, zero value otherwise.

### GetGaebXmlIdOk

`func (o *ServiceSpecificationDto) GetGaebXmlIdOk() (*string, bool)`

GetGaebXmlIdOk returns a tuple with the GaebXmlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaebXmlId

`func (o *ServiceSpecificationDto) SetGaebXmlId(v string)`

SetGaebXmlId sets GaebXmlId field to given value.

### HasGaebXmlId

`func (o *ServiceSpecificationDto) HasGaebXmlId() bool`

HasGaebXmlId returns a boolean if a field has been set.

### GetProjectInformation

`func (o *ServiceSpecificationDto) GetProjectInformation() ProjectInformationDto`

GetProjectInformation returns the ProjectInformation field if non-nil, zero value otherwise.

### GetProjectInformationOk

`func (o *ServiceSpecificationDto) GetProjectInformationOk() (*ProjectInformationDto, bool)`

GetProjectInformationOk returns a tuple with the ProjectInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectInformation

`func (o *ServiceSpecificationDto) SetProjectInformation(v ProjectInformationDto)`

SetProjectInformation sets ProjectInformation field to given value.

### HasProjectInformation

`func (o *ServiceSpecificationDto) HasProjectInformation() bool`

HasProjectInformation returns a boolean if a field has been set.

### GetExchangePhase

`func (o *ServiceSpecificationDto) GetExchangePhase() ExchangePhaseDto`

GetExchangePhase returns the ExchangePhase field if non-nil, zero value otherwise.

### GetExchangePhaseOk

`func (o *ServiceSpecificationDto) GetExchangePhaseOk() (*ExchangePhaseDto, bool)`

GetExchangePhaseOk returns a tuple with the ExchangePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangePhase

`func (o *ServiceSpecificationDto) SetExchangePhase(v ExchangePhaseDto)`

SetExchangePhase sets ExchangePhase field to given value.


### GetOrigin

`func (o *ServiceSpecificationDto) GetOrigin() OriginDto`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ServiceSpecificationDto) GetOriginOk() (*OriginDto, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ServiceSpecificationDto) SetOrigin(v OriginDto)`

SetOrigin sets Origin field to given value.


### GetOriginDetail

`func (o *ServiceSpecificationDto) GetOriginDetail() string`

GetOriginDetail returns the OriginDetail field if non-nil, zero value otherwise.

### GetOriginDetailOk

`func (o *ServiceSpecificationDto) GetOriginDetailOk() (*string, bool)`

GetOriginDetailOk returns a tuple with the OriginDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDetail

`func (o *ServiceSpecificationDto) SetOriginDetail(v string)`

SetOriginDetail sets OriginDetail field to given value.

### HasOriginDetail

`func (o *ServiceSpecificationDto) HasOriginDetail() bool`

HasOriginDetail returns a boolean if a field has been set.

### GetCreationDate

`func (o *ServiceSpecificationDto) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ServiceSpecificationDto) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ServiceSpecificationDto) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ServiceSpecificationDto) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetOfferByDate

`func (o *ServiceSpecificationDto) GetOfferByDate() time.Time`

GetOfferByDate returns the OfferByDate field if non-nil, zero value otherwise.

### GetOfferByDateOk

`func (o *ServiceSpecificationDto) GetOfferByDateOk() (*time.Time, bool)`

GetOfferByDateOk returns a tuple with the OfferByDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferByDate

`func (o *ServiceSpecificationDto) SetOfferByDate(v time.Time)`

SetOfferByDate sets OfferByDate field to given value.

### HasOfferByDate

`func (o *ServiceSpecificationDto) HasOfferByDate() bool`

HasOfferByDate returns a boolean if a field has been set.

### GetDecisionDate

`func (o *ServiceSpecificationDto) GetDecisionDate() time.Time`

GetDecisionDate returns the DecisionDate field if non-nil, zero value otherwise.

### GetDecisionDateOk

`func (o *ServiceSpecificationDto) GetDecisionDateOk() (*time.Time, bool)`

GetDecisionDateOk returns a tuple with the DecisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDate

`func (o *ServiceSpecificationDto) SetDecisionDate(v time.Time)`

SetDecisionDate sets DecisionDate field to given value.

### HasDecisionDate

`func (o *ServiceSpecificationDto) HasDecisionDate() bool`

HasDecisionDate returns a boolean if a field has been set.

### GetBidDate

`func (o *ServiceSpecificationDto) GetBidDate() time.Time`

GetBidDate returns the BidDate field if non-nil, zero value otherwise.

### GetBidDateOk

`func (o *ServiceSpecificationDto) GetBidDateOk() (*time.Time, bool)`

GetBidDateOk returns a tuple with the BidDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidDate

`func (o *ServiceSpecificationDto) SetBidDate(v time.Time)`

SetBidDate sets BidDate field to given value.

### HasBidDate

`func (o *ServiceSpecificationDto) HasBidDate() bool`

HasBidDate returns a boolean if a field has been set.

### GetWarrantyBondPercentage

`func (o *ServiceSpecificationDto) GetWarrantyBondPercentage() float32`

GetWarrantyBondPercentage returns the WarrantyBondPercentage field if non-nil, zero value otherwise.

### GetWarrantyBondPercentageOk

`func (o *ServiceSpecificationDto) GetWarrantyBondPercentageOk() (*float32, bool)`

GetWarrantyBondPercentageOk returns a tuple with the WarrantyBondPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyBondPercentage

`func (o *ServiceSpecificationDto) SetWarrantyBondPercentage(v float32)`

SetWarrantyBondPercentage sets WarrantyBondPercentage field to given value.


### GetExecutionGuaranteePercentage

`func (o *ServiceSpecificationDto) GetExecutionGuaranteePercentage() float32`

GetExecutionGuaranteePercentage returns the ExecutionGuaranteePercentage field if non-nil, zero value otherwise.

### GetExecutionGuaranteePercentageOk

`func (o *ServiceSpecificationDto) GetExecutionGuaranteePercentageOk() (*float32, bool)`

GetExecutionGuaranteePercentageOk returns a tuple with the ExecutionGuaranteePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionGuaranteePercentage

`func (o *ServiceSpecificationDto) SetExecutionGuaranteePercentage(v float32)`

SetExecutionGuaranteePercentage sets ExecutionGuaranteePercentage field to given value.


### GetSubmissionLocation

`func (o *ServiceSpecificationDto) GetSubmissionLocation() string`

GetSubmissionLocation returns the SubmissionLocation field if non-nil, zero value otherwise.

### GetSubmissionLocationOk

`func (o *ServiceSpecificationDto) GetSubmissionLocationOk() (*string, bool)`

GetSubmissionLocationOk returns a tuple with the SubmissionLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionLocation

`func (o *ServiceSpecificationDto) SetSubmissionLocation(v string)`

SetSubmissionLocation sets SubmissionLocation field to given value.

### HasSubmissionLocation

`func (o *ServiceSpecificationDto) HasSubmissionLocation() bool`

HasSubmissionLocation returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceSpecificationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceSpecificationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceSpecificationDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceSpecificationDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ServiceSpecificationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceSpecificationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceSpecificationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceSpecificationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriceInformation

`func (o *ServiceSpecificationDto) GetPriceInformation() PriceInformationDto`

GetPriceInformation returns the PriceInformation field if non-nil, zero value otherwise.

### GetPriceInformationOk

`func (o *ServiceSpecificationDto) GetPriceInformationOk() (*PriceInformationDto, bool)`

GetPriceInformationOk returns a tuple with the PriceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInformation

`func (o *ServiceSpecificationDto) SetPriceInformation(v PriceInformationDto)`

SetPriceInformation sets PriceInformation field to given value.

### HasPriceInformation

`func (o *ServiceSpecificationDto) HasPriceInformation() bool`

HasPriceInformation returns a boolean if a field has been set.

### GetProjectCatalogues

`func (o *ServiceSpecificationDto) GetProjectCatalogues() []CatalogueDto`

GetProjectCatalogues returns the ProjectCatalogues field if non-nil, zero value otherwise.

### GetProjectCataloguesOk

`func (o *ServiceSpecificationDto) GetProjectCataloguesOk() (*[]CatalogueDto, bool)`

GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCatalogues

`func (o *ServiceSpecificationDto) SetProjectCatalogues(v []CatalogueDto)`

SetProjectCatalogues sets ProjectCatalogues field to given value.

### HasProjectCatalogues

`func (o *ServiceSpecificationDto) HasProjectCatalogues() bool`

HasProjectCatalogues returns a boolean if a field has been set.

### GetCatalogueReferences

`func (o *ServiceSpecificationDto) GetCatalogueReferences() []CatalogueReferenceDto`

GetCatalogueReferences returns the CatalogueReferences field if non-nil, zero value otherwise.

### GetCatalogueReferencesOk

`func (o *ServiceSpecificationDto) GetCatalogueReferencesOk() (*[]CatalogueReferenceDto, bool)`

GetCatalogueReferencesOk returns a tuple with the CatalogueReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogueReferences

`func (o *ServiceSpecificationDto) SetCatalogueReferences(v []CatalogueReferenceDto)`

SetCatalogueReferences sets CatalogueReferences field to given value.

### HasCatalogueReferences

`func (o *ServiceSpecificationDto) HasCatalogueReferences() bool`

HasCatalogueReferences returns a boolean if a field has been set.

### GetPlannedExecutionStart

`func (o *ServiceSpecificationDto) GetPlannedExecutionStart() time.Time`

GetPlannedExecutionStart returns the PlannedExecutionStart field if non-nil, zero value otherwise.

### GetPlannedExecutionStartOk

`func (o *ServiceSpecificationDto) GetPlannedExecutionStartOk() (*time.Time, bool)`

GetPlannedExecutionStartOk returns a tuple with the PlannedExecutionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedExecutionStart

`func (o *ServiceSpecificationDto) SetPlannedExecutionStart(v time.Time)`

SetPlannedExecutionStart sets PlannedExecutionStart field to given value.

### HasPlannedExecutionStart

`func (o *ServiceSpecificationDto) HasPlannedExecutionStart() bool`

HasPlannedExecutionStart returns a boolean if a field has been set.

### GetPlannedExecutionEnd

`func (o *ServiceSpecificationDto) GetPlannedExecutionEnd() time.Time`

GetPlannedExecutionEnd returns the PlannedExecutionEnd field if non-nil, zero value otherwise.

### GetPlannedExecutionEndOk

`func (o *ServiceSpecificationDto) GetPlannedExecutionEndOk() (*time.Time, bool)`

GetPlannedExecutionEndOk returns a tuple with the PlannedExecutionEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedExecutionEnd

`func (o *ServiceSpecificationDto) SetPlannedExecutionEnd(v time.Time)`

SetPlannedExecutionEnd sets PlannedExecutionEnd field to given value.

### HasPlannedExecutionEnd

`func (o *ServiceSpecificationDto) HasPlannedExecutionEnd() bool`

HasPlannedExecutionEnd returns a boolean if a field has been set.

### GetContractDate

`func (o *ServiceSpecificationDto) GetContractDate() time.Time`

GetContractDate returns the ContractDate field if non-nil, zero value otherwise.

### GetContractDateOk

`func (o *ServiceSpecificationDto) GetContractDateOk() (*time.Time, bool)`

GetContractDateOk returns a tuple with the ContractDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDate

`func (o *ServiceSpecificationDto) SetContractDate(v time.Time)`

SetContractDate sets ContractDate field to given value.

### HasContractDate

`func (o *ServiceSpecificationDto) HasContractDate() bool`

HasContractDate returns a boolean if a field has been set.

### GetContractIdentifier

`func (o *ServiceSpecificationDto) GetContractIdentifier() string`

GetContractIdentifier returns the ContractIdentifier field if non-nil, zero value otherwise.

### GetContractIdentifierOk

`func (o *ServiceSpecificationDto) GetContractIdentifierOk() (*string, bool)`

GetContractIdentifierOk returns a tuple with the ContractIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractIdentifier

`func (o *ServiceSpecificationDto) SetContractIdentifier(v string)`

SetContractIdentifier sets ContractIdentifier field to given value.

### HasContractIdentifier

`func (o *ServiceSpecificationDto) HasContractIdentifier() bool`

HasContractIdentifier returns a boolean if a field has been set.

### GetWarrantyDuration

`func (o *ServiceSpecificationDto) GetWarrantyDuration() WarrantyDurationDto`

GetWarrantyDuration returns the WarrantyDuration field if non-nil, zero value otherwise.

### GetWarrantyDurationOk

`func (o *ServiceSpecificationDto) GetWarrantyDurationOk() (*WarrantyDurationDto, bool)`

GetWarrantyDurationOk returns a tuple with the WarrantyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyDuration

`func (o *ServiceSpecificationDto) SetWarrantyDuration(v WarrantyDurationDto)`

SetWarrantyDuration sets WarrantyDuration field to given value.

### HasWarrantyDuration

`func (o *ServiceSpecificationDto) HasWarrantyDuration() bool`

HasWarrantyDuration returns a boolean if a field has been set.

### GetWarrantyEnd

`func (o *ServiceSpecificationDto) GetWarrantyEnd() time.Time`

GetWarrantyEnd returns the WarrantyEnd field if non-nil, zero value otherwise.

### GetWarrantyEndOk

`func (o *ServiceSpecificationDto) GetWarrantyEndOk() (*time.Time, bool)`

GetWarrantyEndOk returns a tuple with the WarrantyEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyEnd

`func (o *ServiceSpecificationDto) SetWarrantyEnd(v time.Time)`

SetWarrantyEnd sets WarrantyEnd field to given value.

### HasWarrantyEnd

`func (o *ServiceSpecificationDto) HasWarrantyEnd() bool`

HasWarrantyEnd returns a boolean if a field has been set.

### GetApprovalDate

`func (o *ServiceSpecificationDto) GetApprovalDate() time.Time`

GetApprovalDate returns the ApprovalDate field if non-nil, zero value otherwise.

### GetApprovalDateOk

`func (o *ServiceSpecificationDto) GetApprovalDateOk() (*time.Time, bool)`

GetApprovalDateOk returns a tuple with the ApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalDate

`func (o *ServiceSpecificationDto) SetApprovalDate(v time.Time)`

SetApprovalDate sets ApprovalDate field to given value.

### HasApprovalDate

`func (o *ServiceSpecificationDto) HasApprovalDate() bool`

HasApprovalDate returns a boolean if a field has been set.

### GetTypeOfApproval

`func (o *ServiceSpecificationDto) GetTypeOfApproval() string`

GetTypeOfApproval returns the TypeOfApproval field if non-nil, zero value otherwise.

### GetTypeOfApprovalOk

`func (o *ServiceSpecificationDto) GetTypeOfApprovalOk() (*string, bool)`

GetTypeOfApprovalOk returns a tuple with the TypeOfApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfApproval

`func (o *ServiceSpecificationDto) SetTypeOfApproval(v string)`

SetTypeOfApproval sets TypeOfApproval field to given value.

### HasTypeOfApproval

`func (o *ServiceSpecificationDto) HasTypeOfApproval() bool`

HasTypeOfApproval returns a boolean if a field has been set.

### GetAddendumNumber

`func (o *ServiceSpecificationDto) GetAddendumNumber() string`

GetAddendumNumber returns the AddendumNumber field if non-nil, zero value otherwise.

### GetAddendumNumberOk

`func (o *ServiceSpecificationDto) GetAddendumNumberOk() (*string, bool)`

GetAddendumNumberOk returns a tuple with the AddendumNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddendumNumber

`func (o *ServiceSpecificationDto) SetAddendumNumber(v string)`

SetAddendumNumber sets AddendumNumber field to given value.

### HasAddendumNumber

`func (o *ServiceSpecificationDto) HasAddendumNumber() bool`

HasAddendumNumber returns a boolean if a field has been set.

### GetAddendumStatus

`func (o *ServiceSpecificationDto) GetAddendumStatus() AddendumStatusDto`

GetAddendumStatus returns the AddendumStatus field if non-nil, zero value otherwise.

### GetAddendumStatusOk

`func (o *ServiceSpecificationDto) GetAddendumStatusOk() (*AddendumStatusDto, bool)`

GetAddendumStatusOk returns a tuple with the AddendumStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddendumStatus

`func (o *ServiceSpecificationDto) SetAddendumStatus(v AddendumStatusDto)`

SetAddendumStatus sets AddendumStatus field to given value.

### HasAddendumStatus

`func (o *ServiceSpecificationDto) HasAddendumStatus() bool`

HasAddendumStatus returns a boolean if a field has been set.

### GetCommerceProperties

`func (o *ServiceSpecificationDto) GetCommerceProperties() ServiceSpecificationCommercePropertiesDto`

GetCommerceProperties returns the CommerceProperties field if non-nil, zero value otherwise.

### GetCommercePropertiesOk

`func (o *ServiceSpecificationDto) GetCommercePropertiesOk() (*ServiceSpecificationCommercePropertiesDto, bool)`

GetCommercePropertiesOk returns a tuple with the CommerceProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommerceProperties

`func (o *ServiceSpecificationDto) SetCommerceProperties(v ServiceSpecificationCommercePropertiesDto)`

SetCommerceProperties sets CommerceProperties field to given value.

### HasCommerceProperties

`func (o *ServiceSpecificationDto) HasCommerceProperties() bool`

HasCommerceProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


