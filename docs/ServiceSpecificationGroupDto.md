# ServiceSpecificationGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectHourlyWage** | **float32** | The hourly wage that is used within this ElementContainerBase. Will be propagated to child elements. | [readonly] 
**ProjectTaxRate** | **float32** | The tax rate that is used within this ElementContainerBase. Will be propagated to child elements. | 
**ProjectPriceComponents** | Pointer to **[]string** | The price components that are used within this project. They are ignored during Json deserialization because they will be set from the parent project. | [optional] 
**ProjectItemNumberSchema** | Pointer to [**ItemNumberSchemaDto**](ItemNumberSchemaDto.md) |  | [optional] 
**Elements** | Pointer to [**[]IElementDto**](IElementDto.md) | The IElements within this ElementContainerBase. | [optional] 
**ProjectLabourTimeLabel** | Pointer to **string** | The label used in the parent Project to mark labour time, e.g. \&quot;Hours\&quot; or \&quot;Stunden\&quot;. | [optional] [readonly] 
**ContainsDuplicateItemNumbers** | **bool** | Indicates if there are child IElements that have conflicting, duplicated ItemNumbers or if any child ElementContainerBase elements themselves contain duplicate ItemNumber s. Will always indicate false when told to ignore duplicate item numbers. | [readonly] 
**ContainsDuplicateElementIds** | **bool** | Indicates if there are child IElements that have conflicting, duplicated Ids or if any child ElementContainerBase elements themselves contain duplicate Id s. Will always indicate false when told to ignore duplicate item numbers. | [readonly] 
**IgnoreDuplicateItemNumbers** | **bool** | Indicate if duplicated ItemNumbers within child elements are to be ignored. Will not perform checks for duplicates if yes. | 
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
**ShortText** | Pointer to **string** | Description for this ServiceSpecificationGroup. | [optional] 
**ComissionStatus** | [**ComissionStatusDto**](ComissionStatusDto.md) |  | 
**ItemNumber** | Pointer to [**ItemNumberDto**](ItemNumberDto.md) |  | [optional] 
**ElementType** | Pointer to **string** |  | [optional] 
**IsLot** | **bool** | This indicates if this group is the parent of a lot. See the documentation for more information about working with lots. | [readonly] 
**AlternativeTo** | Pointer to **string** | If this group is an alternative for a base group, then this property should point to the id of the group in this service specification that it can replace. If this is an alternative group to a base group, the PriceType should typically be set to \&quot;WithoutTotal\&quot; so this group does not factor into total costs. The PriceType is not changed when this property is set | [optional] 
**OenormProperties** | Pointer to [**OenormPropertiesDto**](OenormPropertiesDto.md) |  | [optional] 
**HierarchyLevel** | **int32** | This is a zero based hierarchy level. It&#39;s set automatically when used in the context of a Project, and can be used to identify the hierarchy level of the current element. | 
**AddendumStatus** | Pointer to [**AddendumStatusDto**](AddendumStatusDto.md) |  | [optional] 
**AlternativeIdentifier** | Pointer to **int32** | This is an optional property that can be used together with AlternativeTo. If this is set, you can indicate which alternative group a specific group is assigned to. That way, if you specifiy multiple alternative ServiceSpecificationGroups with the same AlternativeIdentifier, you can indicate that to replace a single base ServiceSpecificationGroup, multiple alternative ServiceSpecificationGroups should be used. This property is not checked or managed automatically, so it is possible for this property to become invalid, by for example setting this property but not setting a base group via AlternativeTo. | [optional] 
**AlternativeGroupIdentifier** | Pointer to **int32** | This is an optional property that can be used together with AlternativeTo and AlternativeGroupIdentifier. If this is set, you can indicate which alternative group a specific group is assigned to. That way, you can specify the id (in integer format) for the alternative group this group belongs to. It&#39;s different to AlternativeIdentifier in that the other property describes the id of the group, while this property here describes the group itself. If a group only has set AlternativeGroupIdentifier but not AlternativeIdentifier, then it likely is a base group for a specific group. This property is not checked or managed automatically, so it is possible for this property to become invalid, by for example setting this property but not setting a base group via AlternativeTo. | [optional] 

## Methods

### NewServiceSpecificationGroupDto

`func NewServiceSpecificationGroupDto(projectHourlyWage float32, projectTaxRate float32, containsDuplicateItemNumbers bool, containsDuplicateElementIds bool, ignoreDuplicateItemNumbers bool, ignoreDuplicateElementIds bool, ignoreChildPriceUpdates bool, deductedPrice float32, deductionFactor float32, totalPrice float32, totalPriceGross float32, totalPriceGrossDeducted float32, priceType PriceTypeDto, comissionStatus ComissionStatusDto, isLot bool, hierarchyLevel int32, ) *ServiceSpecificationGroupDto`

NewServiceSpecificationGroupDto instantiates a new ServiceSpecificationGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSpecificationGroupDtoWithDefaults

`func NewServiceSpecificationGroupDtoWithDefaults() *ServiceSpecificationGroupDto`

NewServiceSpecificationGroupDtoWithDefaults instantiates a new ServiceSpecificationGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectHourlyWage

`func (o *ServiceSpecificationGroupDto) GetProjectHourlyWage() float32`

GetProjectHourlyWage returns the ProjectHourlyWage field if non-nil, zero value otherwise.

### GetProjectHourlyWageOk

`func (o *ServiceSpecificationGroupDto) GetProjectHourlyWageOk() (*float32, bool)`

GetProjectHourlyWageOk returns a tuple with the ProjectHourlyWage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectHourlyWage

`func (o *ServiceSpecificationGroupDto) SetProjectHourlyWage(v float32)`

SetProjectHourlyWage sets ProjectHourlyWage field to given value.


### GetProjectTaxRate

`func (o *ServiceSpecificationGroupDto) GetProjectTaxRate() float32`

GetProjectTaxRate returns the ProjectTaxRate field if non-nil, zero value otherwise.

### GetProjectTaxRateOk

`func (o *ServiceSpecificationGroupDto) GetProjectTaxRateOk() (*float32, bool)`

GetProjectTaxRateOk returns a tuple with the ProjectTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTaxRate

`func (o *ServiceSpecificationGroupDto) SetProjectTaxRate(v float32)`

SetProjectTaxRate sets ProjectTaxRate field to given value.


### GetProjectPriceComponents

`func (o *ServiceSpecificationGroupDto) GetProjectPriceComponents() []string`

GetProjectPriceComponents returns the ProjectPriceComponents field if non-nil, zero value otherwise.

### GetProjectPriceComponentsOk

`func (o *ServiceSpecificationGroupDto) GetProjectPriceComponentsOk() (*[]string, bool)`

GetProjectPriceComponentsOk returns a tuple with the ProjectPriceComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPriceComponents

`func (o *ServiceSpecificationGroupDto) SetProjectPriceComponents(v []string)`

SetProjectPriceComponents sets ProjectPriceComponents field to given value.

### HasProjectPriceComponents

`func (o *ServiceSpecificationGroupDto) HasProjectPriceComponents() bool`

HasProjectPriceComponents returns a boolean if a field has been set.

### GetProjectItemNumberSchema

`func (o *ServiceSpecificationGroupDto) GetProjectItemNumberSchema() ItemNumberSchemaDto`

GetProjectItemNumberSchema returns the ProjectItemNumberSchema field if non-nil, zero value otherwise.

### GetProjectItemNumberSchemaOk

`func (o *ServiceSpecificationGroupDto) GetProjectItemNumberSchemaOk() (*ItemNumberSchemaDto, bool)`

GetProjectItemNumberSchemaOk returns a tuple with the ProjectItemNumberSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectItemNumberSchema

`func (o *ServiceSpecificationGroupDto) SetProjectItemNumberSchema(v ItemNumberSchemaDto)`

SetProjectItemNumberSchema sets ProjectItemNumberSchema field to given value.

### HasProjectItemNumberSchema

`func (o *ServiceSpecificationGroupDto) HasProjectItemNumberSchema() bool`

HasProjectItemNumberSchema returns a boolean if a field has been set.

### GetElements

`func (o *ServiceSpecificationGroupDto) GetElements() []IElementDto`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *ServiceSpecificationGroupDto) GetElementsOk() (*[]IElementDto, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *ServiceSpecificationGroupDto) SetElements(v []IElementDto)`

SetElements sets Elements field to given value.

### HasElements

`func (o *ServiceSpecificationGroupDto) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetProjectLabourTimeLabel

`func (o *ServiceSpecificationGroupDto) GetProjectLabourTimeLabel() string`

GetProjectLabourTimeLabel returns the ProjectLabourTimeLabel field if non-nil, zero value otherwise.

### GetProjectLabourTimeLabelOk

`func (o *ServiceSpecificationGroupDto) GetProjectLabourTimeLabelOk() (*string, bool)`

GetProjectLabourTimeLabelOk returns a tuple with the ProjectLabourTimeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLabourTimeLabel

`func (o *ServiceSpecificationGroupDto) SetProjectLabourTimeLabel(v string)`

SetProjectLabourTimeLabel sets ProjectLabourTimeLabel field to given value.

### HasProjectLabourTimeLabel

`func (o *ServiceSpecificationGroupDto) HasProjectLabourTimeLabel() bool`

HasProjectLabourTimeLabel returns a boolean if a field has been set.

### GetContainsDuplicateItemNumbers

`func (o *ServiceSpecificationGroupDto) GetContainsDuplicateItemNumbers() bool`

GetContainsDuplicateItemNumbers returns the ContainsDuplicateItemNumbers field if non-nil, zero value otherwise.

### GetContainsDuplicateItemNumbersOk

`func (o *ServiceSpecificationGroupDto) GetContainsDuplicateItemNumbersOk() (*bool, bool)`

GetContainsDuplicateItemNumbersOk returns a tuple with the ContainsDuplicateItemNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsDuplicateItemNumbers

`func (o *ServiceSpecificationGroupDto) SetContainsDuplicateItemNumbers(v bool)`

SetContainsDuplicateItemNumbers sets ContainsDuplicateItemNumbers field to given value.


### GetContainsDuplicateElementIds

`func (o *ServiceSpecificationGroupDto) GetContainsDuplicateElementIds() bool`

GetContainsDuplicateElementIds returns the ContainsDuplicateElementIds field if non-nil, zero value otherwise.

### GetContainsDuplicateElementIdsOk

`func (o *ServiceSpecificationGroupDto) GetContainsDuplicateElementIdsOk() (*bool, bool)`

GetContainsDuplicateElementIdsOk returns a tuple with the ContainsDuplicateElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsDuplicateElementIds

`func (o *ServiceSpecificationGroupDto) SetContainsDuplicateElementIds(v bool)`

SetContainsDuplicateElementIds sets ContainsDuplicateElementIds field to given value.


### GetIgnoreDuplicateItemNumbers

`func (o *ServiceSpecificationGroupDto) GetIgnoreDuplicateItemNumbers() bool`

GetIgnoreDuplicateItemNumbers returns the IgnoreDuplicateItemNumbers field if non-nil, zero value otherwise.

### GetIgnoreDuplicateItemNumbersOk

`func (o *ServiceSpecificationGroupDto) GetIgnoreDuplicateItemNumbersOk() (*bool, bool)`

GetIgnoreDuplicateItemNumbersOk returns a tuple with the IgnoreDuplicateItemNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDuplicateItemNumbers

`func (o *ServiceSpecificationGroupDto) SetIgnoreDuplicateItemNumbers(v bool)`

SetIgnoreDuplicateItemNumbers sets IgnoreDuplicateItemNumbers field to given value.


### GetIgnoreDuplicateElementIds

`func (o *ServiceSpecificationGroupDto) GetIgnoreDuplicateElementIds() bool`

GetIgnoreDuplicateElementIds returns the IgnoreDuplicateElementIds field if non-nil, zero value otherwise.

### GetIgnoreDuplicateElementIdsOk

`func (o *ServiceSpecificationGroupDto) GetIgnoreDuplicateElementIdsOk() (*bool, bool)`

GetIgnoreDuplicateElementIdsOk returns a tuple with the IgnoreDuplicateElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDuplicateElementIds

`func (o *ServiceSpecificationGroupDto) SetIgnoreDuplicateElementIds(v bool)`

SetIgnoreDuplicateElementIds sets IgnoreDuplicateElementIds field to given value.


### GetTotalPriceGrossByTaxRate

`func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossByTaxRate() []GrossPriceComponentDto`

GetTotalPriceGrossByTaxRate returns the TotalPriceGrossByTaxRate field if non-nil, zero value otherwise.

### GetTotalPriceGrossByTaxRateOk

`func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossByTaxRateOk() (*[]GrossPriceComponentDto, bool)`

GetTotalPriceGrossByTaxRateOk returns a tuple with the TotalPriceGrossByTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceGrossByTaxRate

`func (o *ServiceSpecificationGroupDto) SetTotalPriceGrossByTaxRate(v []GrossPriceComponentDto)`

SetTotalPriceGrossByTaxRate sets TotalPriceGrossByTaxRate field to given value.

### HasTotalPriceGrossByTaxRate

`func (o *ServiceSpecificationGroupDto) HasTotalPriceGrossByTaxRate() bool`

HasTotalPriceGrossByTaxRate returns a boolean if a field has been set.

### GetIgnoreChildPriceUpdates

`func (o *ServiceSpecificationGroupDto) GetIgnoreChildPriceUpdates() bool`

GetIgnoreChildPriceUpdates returns the IgnoreChildPriceUpdates field if non-nil, zero value otherwise.

### GetIgnoreChildPriceUpdatesOk

`func (o *ServiceSpecificationGroupDto) GetIgnoreChildPriceUpdatesOk() (*bool, bool)`

GetIgnoreChildPriceUpdatesOk returns a tuple with the IgnoreChildPriceUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreChildPriceUpdates

`func (o *ServiceSpecificationGroupDto) SetIgnoreChildPriceUpdates(v bool)`

SetIgnoreChildPriceUpdates sets IgnoreChildPriceUpdates field to given value.


### GetDeductedPrice

`func (o *ServiceSpecificationGroupDto) GetDeductedPrice() float32`

GetDeductedPrice returns the DeductedPrice field if non-nil, zero value otherwise.

### GetDeductedPriceOk

`func (o *ServiceSpecificationGroupDto) GetDeductedPriceOk() (*float32, bool)`

GetDeductedPriceOk returns a tuple with the DeductedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductedPrice

`func (o *ServiceSpecificationGroupDto) SetDeductedPrice(v float32)`

SetDeductedPrice sets DeductedPrice field to given value.


### GetDeductionFactor

`func (o *ServiceSpecificationGroupDto) GetDeductionFactor() float32`

GetDeductionFactor returns the DeductionFactor field if non-nil, zero value otherwise.

### GetDeductionFactorOk

`func (o *ServiceSpecificationGroupDto) GetDeductionFactorOk() (*float32, bool)`

GetDeductionFactorOk returns a tuple with the DeductionFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductionFactor

`func (o *ServiceSpecificationGroupDto) SetDeductionFactor(v float32)`

SetDeductionFactor sets DeductionFactor field to given value.


### GetAbsoluteDeduction

`func (o *ServiceSpecificationGroupDto) GetAbsoluteDeduction() float32`

GetAbsoluteDeduction returns the AbsoluteDeduction field if non-nil, zero value otherwise.

### GetAbsoluteDeductionOk

`func (o *ServiceSpecificationGroupDto) GetAbsoluteDeductionOk() (*float32, bool)`

GetAbsoluteDeductionOk returns a tuple with the AbsoluteDeduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteDeduction

`func (o *ServiceSpecificationGroupDto) SetAbsoluteDeduction(v float32)`

SetAbsoluteDeduction sets AbsoluteDeduction field to given value.

### HasAbsoluteDeduction

`func (o *ServiceSpecificationGroupDto) HasAbsoluteDeduction() bool`

HasAbsoluteDeduction returns a boolean if a field has been set.

### GetTotalPrice

`func (o *ServiceSpecificationGroupDto) GetTotalPrice() float32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *ServiceSpecificationGroupDto) GetTotalPriceOk() (*float32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *ServiceSpecificationGroupDto) SetTotalPrice(v float32)`

SetTotalPrice sets TotalPrice field to given value.


### GetTotalPriceGross

`func (o *ServiceSpecificationGroupDto) GetTotalPriceGross() float32`

GetTotalPriceGross returns the TotalPriceGross field if non-nil, zero value otherwise.

### GetTotalPriceGrossOk

`func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossOk() (*float32, bool)`

GetTotalPriceGrossOk returns a tuple with the TotalPriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceGross

`func (o *ServiceSpecificationGroupDto) SetTotalPriceGross(v float32)`

SetTotalPriceGross sets TotalPriceGross field to given value.


### GetTotalPriceGrossDeducted

`func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossDeducted() float32`

GetTotalPriceGrossDeducted returns the TotalPriceGrossDeducted field if non-nil, zero value otherwise.

### GetTotalPriceGrossDeductedOk

`func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossDeductedOk() (*float32, bool)`

GetTotalPriceGrossDeductedOk returns a tuple with the TotalPriceGrossDeducted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceGrossDeducted

`func (o *ServiceSpecificationGroupDto) SetTotalPriceGrossDeducted(v float32)`

SetTotalPriceGrossDeducted sets TotalPriceGrossDeducted field to given value.


### GetPriceType

`func (o *ServiceSpecificationGroupDto) GetPriceType() PriceTypeDto`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *ServiceSpecificationGroupDto) GetPriceTypeOk() (*PriceTypeDto, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *ServiceSpecificationGroupDto) SetPriceType(v PriceTypeDto)`

SetPriceType sets PriceType field to given value.


### GetShortText

`func (o *ServiceSpecificationGroupDto) GetShortText() string`

GetShortText returns the ShortText field if non-nil, zero value otherwise.

### GetShortTextOk

`func (o *ServiceSpecificationGroupDto) GetShortTextOk() (*string, bool)`

GetShortTextOk returns a tuple with the ShortText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortText

`func (o *ServiceSpecificationGroupDto) SetShortText(v string)`

SetShortText sets ShortText field to given value.

### HasShortText

`func (o *ServiceSpecificationGroupDto) HasShortText() bool`

HasShortText returns a boolean if a field has been set.

### GetComissionStatus

`func (o *ServiceSpecificationGroupDto) GetComissionStatus() ComissionStatusDto`

GetComissionStatus returns the ComissionStatus field if non-nil, zero value otherwise.

### GetComissionStatusOk

`func (o *ServiceSpecificationGroupDto) GetComissionStatusOk() (*ComissionStatusDto, bool)`

GetComissionStatusOk returns a tuple with the ComissionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComissionStatus

`func (o *ServiceSpecificationGroupDto) SetComissionStatus(v ComissionStatusDto)`

SetComissionStatus sets ComissionStatus field to given value.


### GetItemNumber

`func (o *ServiceSpecificationGroupDto) GetItemNumber() ItemNumberDto`

GetItemNumber returns the ItemNumber field if non-nil, zero value otherwise.

### GetItemNumberOk

`func (o *ServiceSpecificationGroupDto) GetItemNumberOk() (*ItemNumberDto, bool)`

GetItemNumberOk returns a tuple with the ItemNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemNumber

`func (o *ServiceSpecificationGroupDto) SetItemNumber(v ItemNumberDto)`

SetItemNumber sets ItemNumber field to given value.

### HasItemNumber

`func (o *ServiceSpecificationGroupDto) HasItemNumber() bool`

HasItemNumber returns a boolean if a field has been set.

### GetElementType

`func (o *ServiceSpecificationGroupDto) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *ServiceSpecificationGroupDto) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *ServiceSpecificationGroupDto) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *ServiceSpecificationGroupDto) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetIsLot

`func (o *ServiceSpecificationGroupDto) GetIsLot() bool`

GetIsLot returns the IsLot field if non-nil, zero value otherwise.

### GetIsLotOk

`func (o *ServiceSpecificationGroupDto) GetIsLotOk() (*bool, bool)`

GetIsLotOk returns a tuple with the IsLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLot

`func (o *ServiceSpecificationGroupDto) SetIsLot(v bool)`

SetIsLot sets IsLot field to given value.


### GetAlternativeTo

`func (o *ServiceSpecificationGroupDto) GetAlternativeTo() string`

GetAlternativeTo returns the AlternativeTo field if non-nil, zero value otherwise.

### GetAlternativeToOk

`func (o *ServiceSpecificationGroupDto) GetAlternativeToOk() (*string, bool)`

GetAlternativeToOk returns a tuple with the AlternativeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeTo

`func (o *ServiceSpecificationGroupDto) SetAlternativeTo(v string)`

SetAlternativeTo sets AlternativeTo field to given value.

### HasAlternativeTo

`func (o *ServiceSpecificationGroupDto) HasAlternativeTo() bool`

HasAlternativeTo returns a boolean if a field has been set.

### GetOenormProperties

`func (o *ServiceSpecificationGroupDto) GetOenormProperties() OenormPropertiesDto`

GetOenormProperties returns the OenormProperties field if non-nil, zero value otherwise.

### GetOenormPropertiesOk

`func (o *ServiceSpecificationGroupDto) GetOenormPropertiesOk() (*OenormPropertiesDto, bool)`

GetOenormPropertiesOk returns a tuple with the OenormProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOenormProperties

`func (o *ServiceSpecificationGroupDto) SetOenormProperties(v OenormPropertiesDto)`

SetOenormProperties sets OenormProperties field to given value.

### HasOenormProperties

`func (o *ServiceSpecificationGroupDto) HasOenormProperties() bool`

HasOenormProperties returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *ServiceSpecificationGroupDto) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *ServiceSpecificationGroupDto) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *ServiceSpecificationGroupDto) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.


### GetAddendumStatus

`func (o *ServiceSpecificationGroupDto) GetAddendumStatus() AddendumStatusDto`

GetAddendumStatus returns the AddendumStatus field if non-nil, zero value otherwise.

### GetAddendumStatusOk

`func (o *ServiceSpecificationGroupDto) GetAddendumStatusOk() (*AddendumStatusDto, bool)`

GetAddendumStatusOk returns a tuple with the AddendumStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddendumStatus

`func (o *ServiceSpecificationGroupDto) SetAddendumStatus(v AddendumStatusDto)`

SetAddendumStatus sets AddendumStatus field to given value.

### HasAddendumStatus

`func (o *ServiceSpecificationGroupDto) HasAddendumStatus() bool`

HasAddendumStatus returns a boolean if a field has been set.

### GetAlternativeIdentifier

`func (o *ServiceSpecificationGroupDto) GetAlternativeIdentifier() int32`

GetAlternativeIdentifier returns the AlternativeIdentifier field if non-nil, zero value otherwise.

### GetAlternativeIdentifierOk

`func (o *ServiceSpecificationGroupDto) GetAlternativeIdentifierOk() (*int32, bool)`

GetAlternativeIdentifierOk returns a tuple with the AlternativeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeIdentifier

`func (o *ServiceSpecificationGroupDto) SetAlternativeIdentifier(v int32)`

SetAlternativeIdentifier sets AlternativeIdentifier field to given value.

### HasAlternativeIdentifier

`func (o *ServiceSpecificationGroupDto) HasAlternativeIdentifier() bool`

HasAlternativeIdentifier returns a boolean if a field has been set.

### GetAlternativeGroupIdentifier

`func (o *ServiceSpecificationGroupDto) GetAlternativeGroupIdentifier() int32`

GetAlternativeGroupIdentifier returns the AlternativeGroupIdentifier field if non-nil, zero value otherwise.

### GetAlternativeGroupIdentifierOk

`func (o *ServiceSpecificationGroupDto) GetAlternativeGroupIdentifierOk() (*int32, bool)`

GetAlternativeGroupIdentifierOk returns a tuple with the AlternativeGroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeGroupIdentifier

`func (o *ServiceSpecificationGroupDto) SetAlternativeGroupIdentifier(v int32)`

SetAlternativeGroupIdentifier sets AlternativeGroupIdentifier field to given value.

### HasAlternativeGroupIdentifier

`func (o *ServiceSpecificationGroupDto) HasAlternativeGroupIdentifier() bool`

HasAlternativeGroupIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


