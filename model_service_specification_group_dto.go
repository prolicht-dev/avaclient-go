/*
AVACloud API 1.45.0

AVACloud API specification

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
)

// checks if the ServiceSpecificationGroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceSpecificationGroupDto{}

// ServiceSpecificationGroupDto struct for ServiceSpecificationGroupDto
type ServiceSpecificationGroupDto struct {
	IElementDto
	// The hourly wage that is used within this ElementContainerBase. Will be propagated to child elements.
	ProjectHourlyWage float32 `json:"projectHourlyWage"`
	// The tax rate that is used within this ElementContainerBase. Will be propagated to child elements.
	ProjectTaxRate float32 `json:"projectTaxRate"`
	// The price components that are used within this project. They are ignored during Json deserialization because they will be set from the parent project.
	ProjectPriceComponents  []string             `json:"projectPriceComponents,omitempty"`
	ProjectItemNumberSchema *ItemNumberSchemaDto `json:"projectItemNumberSchema,omitempty"`
	// The IElements within this ElementContainerBase.
	Elements []IElementDtoHolder `json:"elements,omitempty"`
	// The label used in the parent Project to mark labour time, e.g. \"Hours\" or \"Stunden\".
	ProjectLabourTimeLabel *string `json:"projectLabourTimeLabel,omitempty"`
	// Indicates if there are child IElements that have conflicting, duplicated ItemNumbers or if any child ElementContainerBase elements themselves contain duplicate ItemNumber s. Will always indicate false when told to ignore duplicate item numbers.
	ContainsDuplicateItemNumbers bool `json:"containsDuplicateItemNumbers"`
	// Indicates if there are child IElements that have conflicting, duplicated Ids or if any child ElementContainerBase elements themselves contain duplicate Id s. Will always indicate false when told to ignore duplicate item numbers.
	ContainsDuplicateElementIds bool `json:"containsDuplicateElementIds"`
	// Indicate if duplicated ItemNumbers within child elements are to be ignored. Will not perform checks for duplicates if yes.
	IgnoreDuplicateItemNumbers bool `json:"ignoreDuplicateItemNumbers"`
	// If this is set to true, the ProjectCatalogues property will not be propagated to child elements. This is useful in mapping scenarios, where you want to disable propagation for multiple changes, and only enable it once you have mapped all properties
	IgnoreProjectCataloguePropagation bool `json:"ignoreProjectCataloguePropagation"`
	// Indicate if duplicated Ids within child elements are to be ignored. Will not perform checks for duplicates if yes.
	IgnoreDuplicateElementIds bool `json:"ignoreDuplicateElementIds"`
	// Price composition by tax rate.
	TotalPriceGrossByTaxRate []GrossPriceComponentDto `json:"totalPriceGrossByTaxRate,omitempty"`
	// Internally used to indicate that a propagation is currently done, this is done to not recalculate every single result from a lot of changes when it is sufficient to calculate the total price at once.
	IgnoreChildPriceUpdates bool `json:"ignoreChildPriceUpdates"`
	// Net price after applied deductions.
	DeductedPrice float32 `json:"deductedPrice"`
	// Factor of applied deductions to the total price. For example, \"0.03\" means that a 3% deduction is to be applied.
	DeductionFactor float32 `json:"deductionFactor"`
	// The exact amount of the discount as an absolute value. For backwards compatibility reasons, setting this value will also set a calculated value to DeductionFactor, which will also be updated in case the total price is changed to reflect a relative value of the absolute discount sum.
	AbsoluteDeduction *float32 `json:"absoluteDeduction,omitempty"`
	// Will return this ElementContainerBase's total price.
	TotalPrice float32 `json:"totalPrice"`
	// The total gross price for this ElementContainerBase including all child elements.
	TotalPriceGross float32 `json:"totalPriceGross"`
	// Total gross price after applied deductions.
	TotalPriceGrossDeducted float32      `json:"totalPriceGrossDeducted"`
	PriceType               PriceTypeDto `json:"priceType"`
	// Description for this ServiceSpecificationGroup.
	ShortText       *string            `json:"shortText,omitempty"`
	ComissionStatus ComissionStatusDto `json:"comissionStatus"`
	ItemNumber      *ItemNumberDto     `json:"itemNumber,omitempty"`
	ElementType     *string            `json:"elementType,omitempty"`
	// This indicates if this group is the parent of a lot. See the documentation for more information about working with lots.
	IsLot bool `json:"isLot"`
	// If this group is an alternative for a base group, then this property should point to the id of the group in this service specification that it can replace. If this is an alternative group to a base group, the PriceType should typically be set to \"WithoutTotal\" so this group does not factor into total costs. The PriceType is not changed when this property is set
	AlternativeTo    *string              `json:"alternativeTo,omitempty"`
	OenormProperties *OenormPropertiesDto `json:"oenormProperties,omitempty"`
	// This is a zero based hierarchy level. It's set automatically when used in the context of a Project, and can be used to identify the hierarchy level of the current element.
	HierarchyLevel int32              `json:"hierarchyLevel"`
	AddendumStatus *AddendumStatusDto `json:"addendumStatus,omitempty"`
	// This is an optional property that can be used together with AlternativeTo. If this is set, you can indicate which alternative group a specific group is assigned to. That way, if you specifiy multiple alternative ServiceSpecificationGroups with the same AlternativeIdentifier, you can indicate that to replace a single base ServiceSpecificationGroup, multiple alternative ServiceSpecificationGroups should be used. This property is not checked or managed automatically, so it is possible for this property to become invalid, by for example setting this property but not setting a base group via AlternativeTo.
	AlternativeIdentifier *int32 `json:"alternativeIdentifier,omitempty"`
	// This is an optional property that can be used together with AlternativeTo and AlternativeGroupIdentifier. If this is set, you can indicate which alternative group a specific group is assigned to. That way, you can specify the id (in integer format) for the alternative group this group belongs to. It's different to AlternativeIdentifier in that the other property describes the id of the group, while this property here describes the group itself. If a group only has set AlternativeGroupIdentifier but not AlternativeIdentifier, then it likely is a base group for a specific group. This property is not checked or managed automatically, so it is possible for this property to become invalid, by for example setting this property but not setting a base group via AlternativeTo.
	AlternativeGroupIdentifier *int32 `json:"alternativeGroupIdentifier,omitempty"`
}

// NewServiceSpecificationGroupDto instantiates a new ServiceSpecificationGroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSpecificationGroupDto(projectHourlyWage float32, projectTaxRate float32, containsDuplicateItemNumbers bool, containsDuplicateElementIds bool, ignoreDuplicateItemNumbers bool, ignoreProjectCataloguePropagation bool, ignoreDuplicateElementIds bool, ignoreChildPriceUpdates bool, deductedPrice float32, deductionFactor float32, totalPrice float32, totalPriceGross float32, totalPriceGrossDeducted float32, priceType PriceTypeDto, comissionStatus ComissionStatusDto, isLot bool, hierarchyLevel int32, id string, elementTypeDiscriminator string) *ServiceSpecificationGroupDto {
	this := ServiceSpecificationGroupDto{}
	this.Id = id
	this.ElementTypeDiscriminator = elementTypeDiscriminator
	this.ProjectHourlyWage = projectHourlyWage
	this.ProjectTaxRate = projectTaxRate
	this.ContainsDuplicateItemNumbers = containsDuplicateItemNumbers
	this.ContainsDuplicateElementIds = containsDuplicateElementIds
	this.IgnoreDuplicateItemNumbers = ignoreDuplicateItemNumbers
	this.IgnoreProjectCataloguePropagation = ignoreProjectCataloguePropagation
	this.IgnoreDuplicateElementIds = ignoreDuplicateElementIds
	this.IgnoreChildPriceUpdates = ignoreChildPriceUpdates
	this.DeductedPrice = deductedPrice
	this.DeductionFactor = deductionFactor
	this.TotalPrice = totalPrice
	this.TotalPriceGross = totalPriceGross
	this.TotalPriceGrossDeducted = totalPriceGrossDeducted
	this.PriceType = priceType
	this.ComissionStatus = comissionStatus
	this.IsLot = isLot
	this.HierarchyLevel = hierarchyLevel
	return &this
}

// NewServiceSpecificationGroupDtoWithDefaults instantiates a new ServiceSpecificationGroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSpecificationGroupDtoWithDefaults() *ServiceSpecificationGroupDto {
	this := ServiceSpecificationGroupDto{}
	return &this
}

// GetProjectHourlyWage returns the ProjectHourlyWage field value
func (o *ServiceSpecificationGroupDto) GetProjectHourlyWage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProjectHourlyWage
}

// GetProjectHourlyWageOk returns a tuple with the ProjectHourlyWage field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetProjectHourlyWageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectHourlyWage, true
}

// SetProjectHourlyWage sets field value
func (o *ServiceSpecificationGroupDto) SetProjectHourlyWage(v float32) {
	o.ProjectHourlyWage = v
}

// GetProjectTaxRate returns the ProjectTaxRate field value
func (o *ServiceSpecificationGroupDto) GetProjectTaxRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProjectTaxRate
}

// GetProjectTaxRateOk returns a tuple with the ProjectTaxRate field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetProjectTaxRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectTaxRate, true
}

// SetProjectTaxRate sets field value
func (o *ServiceSpecificationGroupDto) SetProjectTaxRate(v float32) {
	o.ProjectTaxRate = v
}

// GetProjectPriceComponents returns the ProjectPriceComponents field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetProjectPriceComponents() []string {
	if o == nil || IsNil(o.ProjectPriceComponents) {
		var ret []string
		return ret
	}
	return o.ProjectPriceComponents
}

// GetProjectPriceComponentsOk returns a tuple with the ProjectPriceComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetProjectPriceComponentsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectPriceComponents) {
		return nil, false
	}
	return o.ProjectPriceComponents, true
}

// HasProjectPriceComponents returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasProjectPriceComponents() bool {
	if o != nil && !IsNil(o.ProjectPriceComponents) {
		return true
	}

	return false
}

// SetProjectPriceComponents gets a reference to the given []string and assigns it to the ProjectPriceComponents field.
func (o *ServiceSpecificationGroupDto) SetProjectPriceComponents(v []string) {
	o.ProjectPriceComponents = v
}

// GetProjectItemNumberSchema returns the ProjectItemNumberSchema field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetProjectItemNumberSchema() ItemNumberSchemaDto {
	if o == nil || IsNil(o.ProjectItemNumberSchema) {
		var ret ItemNumberSchemaDto
		return ret
	}
	return *o.ProjectItemNumberSchema
}

// GetProjectItemNumberSchemaOk returns a tuple with the ProjectItemNumberSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetProjectItemNumberSchemaOk() (*ItemNumberSchemaDto, bool) {
	if o == nil || IsNil(o.ProjectItemNumberSchema) {
		return nil, false
	}
	return o.ProjectItemNumberSchema, true
}

// HasProjectItemNumberSchema returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasProjectItemNumberSchema() bool {
	if o != nil && !IsNil(o.ProjectItemNumberSchema) {
		return true
	}

	return false
}

// SetProjectItemNumberSchema gets a reference to the given ItemNumberSchemaDto and assigns it to the ProjectItemNumberSchema field.
func (o *ServiceSpecificationGroupDto) SetProjectItemNumberSchema(v ItemNumberSchemaDto) {
	o.ProjectItemNumberSchema = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetElements() []IElementDtoHolder {
	if o == nil || IsNil(o.Elements) {
		var ret []IElementDtoHolder
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetElementsOk() ([]IElementDtoHolder, bool) {
	if o == nil || IsNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasElements() bool {
	if o != nil && !IsNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []IElementDtoHolder and assigns it to the Elements field.
func (o *ServiceSpecificationGroupDto) SetElements(v []IElementDtoHolder) {
	o.Elements = v
}

// GetProjectLabourTimeLabel returns the ProjectLabourTimeLabel field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetProjectLabourTimeLabel() string {
	if o == nil || IsNil(o.ProjectLabourTimeLabel) {
		var ret string
		return ret
	}
	return *o.ProjectLabourTimeLabel
}

// GetProjectLabourTimeLabelOk returns a tuple with the ProjectLabourTimeLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetProjectLabourTimeLabelOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectLabourTimeLabel) {
		return nil, false
	}
	return o.ProjectLabourTimeLabel, true
}

// HasProjectLabourTimeLabel returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasProjectLabourTimeLabel() bool {
	if o != nil && !IsNil(o.ProjectLabourTimeLabel) {
		return true
	}

	return false
}

// SetProjectLabourTimeLabel gets a reference to the given string and assigns it to the ProjectLabourTimeLabel field.
func (o *ServiceSpecificationGroupDto) SetProjectLabourTimeLabel(v string) {
	o.ProjectLabourTimeLabel = &v
}

// GetContainsDuplicateItemNumbers returns the ContainsDuplicateItemNumbers field value
func (o *ServiceSpecificationGroupDto) GetContainsDuplicateItemNumbers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ContainsDuplicateItemNumbers
}

// GetContainsDuplicateItemNumbersOk returns a tuple with the ContainsDuplicateItemNumbers field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetContainsDuplicateItemNumbersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainsDuplicateItemNumbers, true
}

// SetContainsDuplicateItemNumbers sets field value
func (o *ServiceSpecificationGroupDto) SetContainsDuplicateItemNumbers(v bool) {
	o.ContainsDuplicateItemNumbers = v
}

// GetContainsDuplicateElementIds returns the ContainsDuplicateElementIds field value
func (o *ServiceSpecificationGroupDto) GetContainsDuplicateElementIds() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ContainsDuplicateElementIds
}

// GetContainsDuplicateElementIdsOk returns a tuple with the ContainsDuplicateElementIds field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetContainsDuplicateElementIdsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainsDuplicateElementIds, true
}

// SetContainsDuplicateElementIds sets field value
func (o *ServiceSpecificationGroupDto) SetContainsDuplicateElementIds(v bool) {
	o.ContainsDuplicateElementIds = v
}

// GetIgnoreDuplicateItemNumbers returns the IgnoreDuplicateItemNumbers field value
func (o *ServiceSpecificationGroupDto) GetIgnoreDuplicateItemNumbers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgnoreDuplicateItemNumbers
}

// GetIgnoreDuplicateItemNumbersOk returns a tuple with the IgnoreDuplicateItemNumbers field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetIgnoreDuplicateItemNumbersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoreDuplicateItemNumbers, true
}

// SetIgnoreDuplicateItemNumbers sets field value
func (o *ServiceSpecificationGroupDto) SetIgnoreDuplicateItemNumbers(v bool) {
	o.IgnoreDuplicateItemNumbers = v
}

// GetIgnoreProjectCataloguePropagation returns the IgnoreProjectCataloguePropagation field value
func (o *ServiceSpecificationGroupDto) GetIgnoreProjectCataloguePropagation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgnoreProjectCataloguePropagation
}

// GetIgnoreProjectCataloguePropagationOk returns a tuple with the IgnoreProjectCataloguePropagation field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetIgnoreProjectCataloguePropagationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoreProjectCataloguePropagation, true
}

// SetIgnoreProjectCataloguePropagation sets field value
func (o *ServiceSpecificationGroupDto) SetIgnoreProjectCataloguePropagation(v bool) {
	o.IgnoreProjectCataloguePropagation = v
}

// GetIgnoreDuplicateElementIds returns the IgnoreDuplicateElementIds field value
func (o *ServiceSpecificationGroupDto) GetIgnoreDuplicateElementIds() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgnoreDuplicateElementIds
}

// GetIgnoreDuplicateElementIdsOk returns a tuple with the IgnoreDuplicateElementIds field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetIgnoreDuplicateElementIdsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoreDuplicateElementIds, true
}

// SetIgnoreDuplicateElementIds sets field value
func (o *ServiceSpecificationGroupDto) SetIgnoreDuplicateElementIds(v bool) {
	o.IgnoreDuplicateElementIds = v
}

// GetTotalPriceGrossByTaxRate returns the TotalPriceGrossByTaxRate field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossByTaxRate() []GrossPriceComponentDto {
	if o == nil || IsNil(o.TotalPriceGrossByTaxRate) {
		var ret []GrossPriceComponentDto
		return ret
	}
	return o.TotalPriceGrossByTaxRate
}

// GetTotalPriceGrossByTaxRateOk returns a tuple with the TotalPriceGrossByTaxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossByTaxRateOk() ([]GrossPriceComponentDto, bool) {
	if o == nil || IsNil(o.TotalPriceGrossByTaxRate) {
		return nil, false
	}
	return o.TotalPriceGrossByTaxRate, true
}

// HasTotalPriceGrossByTaxRate returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasTotalPriceGrossByTaxRate() bool {
	if o != nil && !IsNil(o.TotalPriceGrossByTaxRate) {
		return true
	}

	return false
}

// SetTotalPriceGrossByTaxRate gets a reference to the given []GrossPriceComponentDto and assigns it to the TotalPriceGrossByTaxRate field.
func (o *ServiceSpecificationGroupDto) SetTotalPriceGrossByTaxRate(v []GrossPriceComponentDto) {
	o.TotalPriceGrossByTaxRate = v
}

// GetIgnoreChildPriceUpdates returns the IgnoreChildPriceUpdates field value
func (o *ServiceSpecificationGroupDto) GetIgnoreChildPriceUpdates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgnoreChildPriceUpdates
}

// GetIgnoreChildPriceUpdatesOk returns a tuple with the IgnoreChildPriceUpdates field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetIgnoreChildPriceUpdatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoreChildPriceUpdates, true
}

// SetIgnoreChildPriceUpdates sets field value
func (o *ServiceSpecificationGroupDto) SetIgnoreChildPriceUpdates(v bool) {
	o.IgnoreChildPriceUpdates = v
}

// GetDeductedPrice returns the DeductedPrice field value
func (o *ServiceSpecificationGroupDto) GetDeductedPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DeductedPrice
}

// GetDeductedPriceOk returns a tuple with the DeductedPrice field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetDeductedPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeductedPrice, true
}

// SetDeductedPrice sets field value
func (o *ServiceSpecificationGroupDto) SetDeductedPrice(v float32) {
	o.DeductedPrice = v
}

// GetDeductionFactor returns the DeductionFactor field value
func (o *ServiceSpecificationGroupDto) GetDeductionFactor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DeductionFactor
}

// GetDeductionFactorOk returns a tuple with the DeductionFactor field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetDeductionFactorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeductionFactor, true
}

// SetDeductionFactor sets field value
func (o *ServiceSpecificationGroupDto) SetDeductionFactor(v float32) {
	o.DeductionFactor = v
}

// GetAbsoluteDeduction returns the AbsoluteDeduction field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetAbsoluteDeduction() float32 {
	if o == nil || IsNil(o.AbsoluteDeduction) {
		var ret float32
		return ret
	}
	return *o.AbsoluteDeduction
}

// GetAbsoluteDeductionOk returns a tuple with the AbsoluteDeduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetAbsoluteDeductionOk() (*float32, bool) {
	if o == nil || IsNil(o.AbsoluteDeduction) {
		return nil, false
	}
	return o.AbsoluteDeduction, true
}

// HasAbsoluteDeduction returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasAbsoluteDeduction() bool {
	if o != nil && !IsNil(o.AbsoluteDeduction) {
		return true
	}

	return false
}

// SetAbsoluteDeduction gets a reference to the given float32 and assigns it to the AbsoluteDeduction field.
func (o *ServiceSpecificationGroupDto) SetAbsoluteDeduction(v float32) {
	o.AbsoluteDeduction = &v
}

// GetTotalPrice returns the TotalPrice field value
func (o *ServiceSpecificationGroupDto) GetTotalPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetTotalPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *ServiceSpecificationGroupDto) SetTotalPrice(v float32) {
	o.TotalPrice = v
}

// GetTotalPriceGross returns the TotalPriceGross field value
func (o *ServiceSpecificationGroupDto) GetTotalPriceGross() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPriceGross
}

// GetTotalPriceGrossOk returns a tuple with the TotalPriceGross field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPriceGross, true
}

// SetTotalPriceGross sets field value
func (o *ServiceSpecificationGroupDto) SetTotalPriceGross(v float32) {
	o.TotalPriceGross = v
}

// GetTotalPriceGrossDeducted returns the TotalPriceGrossDeducted field value
func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossDeducted() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPriceGrossDeducted
}

// GetTotalPriceGrossDeductedOk returns a tuple with the TotalPriceGrossDeducted field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetTotalPriceGrossDeductedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPriceGrossDeducted, true
}

// SetTotalPriceGrossDeducted sets field value
func (o *ServiceSpecificationGroupDto) SetTotalPriceGrossDeducted(v float32) {
	o.TotalPriceGrossDeducted = v
}

// GetPriceType returns the PriceType field value
func (o *ServiceSpecificationGroupDto) GetPriceType() PriceTypeDto {
	if o == nil {
		var ret PriceTypeDto
		return ret
	}

	return o.PriceType
}

// GetPriceTypeOk returns a tuple with the PriceType field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetPriceTypeOk() (*PriceTypeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceType, true
}

// SetPriceType sets field value
func (o *ServiceSpecificationGroupDto) SetPriceType(v PriceTypeDto) {
	o.PriceType = v
}

// GetShortText returns the ShortText field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetShortText() string {
	if o == nil || IsNil(o.ShortText) {
		var ret string
		return ret
	}
	return *o.ShortText
}

// GetShortTextOk returns a tuple with the ShortText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetShortTextOk() (*string, bool) {
	if o == nil || IsNil(o.ShortText) {
		return nil, false
	}
	return o.ShortText, true
}

// HasShortText returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasShortText() bool {
	if o != nil && !IsNil(o.ShortText) {
		return true
	}

	return false
}

// SetShortText gets a reference to the given string and assigns it to the ShortText field.
func (o *ServiceSpecificationGroupDto) SetShortText(v string) {
	o.ShortText = &v
}

// GetComissionStatus returns the ComissionStatus field value
func (o *ServiceSpecificationGroupDto) GetComissionStatus() ComissionStatusDto {
	if o == nil {
		var ret ComissionStatusDto
		return ret
	}

	return o.ComissionStatus
}

// GetComissionStatusOk returns a tuple with the ComissionStatus field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetComissionStatusOk() (*ComissionStatusDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComissionStatus, true
}

// SetComissionStatus sets field value
func (o *ServiceSpecificationGroupDto) SetComissionStatus(v ComissionStatusDto) {
	o.ComissionStatus = v
}

// GetItemNumber returns the ItemNumber field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetItemNumber() ItemNumberDto {
	if o == nil || IsNil(o.ItemNumber) {
		var ret ItemNumberDto
		return ret
	}
	return *o.ItemNumber
}

// GetItemNumberOk returns a tuple with the ItemNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetItemNumberOk() (*ItemNumberDto, bool) {
	if o == nil || IsNil(o.ItemNumber) {
		return nil, false
	}
	return o.ItemNumber, true
}

// HasItemNumber returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasItemNumber() bool {
	if o != nil && !IsNil(o.ItemNumber) {
		return true
	}

	return false
}

// SetItemNumber gets a reference to the given ItemNumberDto and assigns it to the ItemNumber field.
func (o *ServiceSpecificationGroupDto) SetItemNumber(v ItemNumberDto) {
	o.ItemNumber = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetElementType() string {
	if o == nil || IsNil(o.ElementType) {
		var ret string
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetElementTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ElementType) {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasElementType() bool {
	if o != nil && !IsNil(o.ElementType) {
		return true
	}

	return false
}

// SetElementType gets a reference to the given string and assigns it to the ElementType field.
func (o *ServiceSpecificationGroupDto) SetElementType(v string) {
	o.ElementType = &v
}

// GetIsLot returns the IsLot field value
func (o *ServiceSpecificationGroupDto) GetIsLot() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLot
}

// GetIsLotOk returns a tuple with the IsLot field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetIsLotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLot, true
}

// SetIsLot sets field value
func (o *ServiceSpecificationGroupDto) SetIsLot(v bool) {
	o.IsLot = v
}

// GetAlternativeTo returns the AlternativeTo field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetAlternativeTo() string {
	if o == nil || IsNil(o.AlternativeTo) {
		var ret string
		return ret
	}
	return *o.AlternativeTo
}

// GetAlternativeToOk returns a tuple with the AlternativeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetAlternativeToOk() (*string, bool) {
	if o == nil || IsNil(o.AlternativeTo) {
		return nil, false
	}
	return o.AlternativeTo, true
}

// HasAlternativeTo returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasAlternativeTo() bool {
	if o != nil && !IsNil(o.AlternativeTo) {
		return true
	}

	return false
}

// SetAlternativeTo gets a reference to the given string and assigns it to the AlternativeTo field.
func (o *ServiceSpecificationGroupDto) SetAlternativeTo(v string) {
	o.AlternativeTo = &v
}

// GetOenormProperties returns the OenormProperties field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetOenormProperties() OenormPropertiesDto {
	if o == nil || IsNil(o.OenormProperties) {
		var ret OenormPropertiesDto
		return ret
	}
	return *o.OenormProperties
}

// GetOenormPropertiesOk returns a tuple with the OenormProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetOenormPropertiesOk() (*OenormPropertiesDto, bool) {
	if o == nil || IsNil(o.OenormProperties) {
		return nil, false
	}
	return o.OenormProperties, true
}

// HasOenormProperties returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasOenormProperties() bool {
	if o != nil && !IsNil(o.OenormProperties) {
		return true
	}

	return false
}

// SetOenormProperties gets a reference to the given OenormPropertiesDto and assigns it to the OenormProperties field.
func (o *ServiceSpecificationGroupDto) SetOenormProperties(v OenormPropertiesDto) {
	o.OenormProperties = &v
}

// GetHierarchyLevel returns the HierarchyLevel field value
func (o *ServiceSpecificationGroupDto) GetHierarchyLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HierarchyLevel
}

// GetHierarchyLevelOk returns a tuple with the HierarchyLevel field value
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetHierarchyLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HierarchyLevel, true
}

// SetHierarchyLevel sets field value
func (o *ServiceSpecificationGroupDto) SetHierarchyLevel(v int32) {
	o.HierarchyLevel = v
}

// GetAddendumStatus returns the AddendumStatus field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetAddendumStatus() AddendumStatusDto {
	if o == nil || IsNil(o.AddendumStatus) {
		var ret AddendumStatusDto
		return ret
	}
	return *o.AddendumStatus
}

// GetAddendumStatusOk returns a tuple with the AddendumStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetAddendumStatusOk() (*AddendumStatusDto, bool) {
	if o == nil || IsNil(o.AddendumStatus) {
		return nil, false
	}
	return o.AddendumStatus, true
}

// HasAddendumStatus returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasAddendumStatus() bool {
	if o != nil && !IsNil(o.AddendumStatus) {
		return true
	}

	return false
}

// SetAddendumStatus gets a reference to the given AddendumStatusDto and assigns it to the AddendumStatus field.
func (o *ServiceSpecificationGroupDto) SetAddendumStatus(v AddendumStatusDto) {
	o.AddendumStatus = &v
}

// GetAlternativeIdentifier returns the AlternativeIdentifier field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetAlternativeIdentifier() int32 {
	if o == nil || IsNil(o.AlternativeIdentifier) {
		var ret int32
		return ret
	}
	return *o.AlternativeIdentifier
}

// GetAlternativeIdentifierOk returns a tuple with the AlternativeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetAlternativeIdentifierOk() (*int32, bool) {
	if o == nil || IsNil(o.AlternativeIdentifier) {
		return nil, false
	}
	return o.AlternativeIdentifier, true
}

// HasAlternativeIdentifier returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasAlternativeIdentifier() bool {
	if o != nil && !IsNil(o.AlternativeIdentifier) {
		return true
	}

	return false
}

// SetAlternativeIdentifier gets a reference to the given int32 and assigns it to the AlternativeIdentifier field.
func (o *ServiceSpecificationGroupDto) SetAlternativeIdentifier(v int32) {
	o.AlternativeIdentifier = &v
}

// GetAlternativeGroupIdentifier returns the AlternativeGroupIdentifier field value if set, zero value otherwise.
func (o *ServiceSpecificationGroupDto) GetAlternativeGroupIdentifier() int32 {
	if o == nil || IsNil(o.AlternativeGroupIdentifier) {
		var ret int32
		return ret
	}
	return *o.AlternativeGroupIdentifier
}

// GetAlternativeGroupIdentifierOk returns a tuple with the AlternativeGroupIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSpecificationGroupDto) GetAlternativeGroupIdentifierOk() (*int32, bool) {
	if o == nil || IsNil(o.AlternativeGroupIdentifier) {
		return nil, false
	}
	return o.AlternativeGroupIdentifier, true
}

// HasAlternativeGroupIdentifier returns a boolean if a field has been set.
func (o *ServiceSpecificationGroupDto) HasAlternativeGroupIdentifier() bool {
	if o != nil && !IsNil(o.AlternativeGroupIdentifier) {
		return true
	}

	return false
}

// SetAlternativeGroupIdentifier gets a reference to the given int32 and assigns it to the AlternativeGroupIdentifier field.
func (o *ServiceSpecificationGroupDto) SetAlternativeGroupIdentifier(v int32) {
	o.AlternativeGroupIdentifier = &v
}

func (o ServiceSpecificationGroupDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceSpecificationGroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedIElementDto, errIElementDto := json.Marshal(o.IElementDto)
	if errIElementDto != nil {
		return map[string]interface{}{}, errIElementDto
	}
	errIElementDto = json.Unmarshal([]byte(serializedIElementDto), &toSerialize)
	if errIElementDto != nil {
		return map[string]interface{}{}, errIElementDto
	}
	// skip: projectHourlyWage is readOnly
	toSerialize["projectTaxRate"] = o.ProjectTaxRate
	if !IsNil(o.ProjectPriceComponents) {
		toSerialize["projectPriceComponents"] = o.ProjectPriceComponents
	}
	if !IsNil(o.ProjectItemNumberSchema) {
		toSerialize["projectItemNumberSchema"] = o.ProjectItemNumberSchema
	}
	if !IsNil(o.Elements) {
		toSerialize["elements"] = o.Elements
	}
	// skip: projectLabourTimeLabel is readOnly
	// skip: containsDuplicateItemNumbers is readOnly
	// skip: containsDuplicateElementIds is readOnly
	toSerialize["ignoreDuplicateItemNumbers"] = o.IgnoreDuplicateItemNumbers
	toSerialize["ignoreProjectCataloguePropagation"] = o.IgnoreProjectCataloguePropagation
	toSerialize["ignoreDuplicateElementIds"] = o.IgnoreDuplicateElementIds
	if !IsNil(o.TotalPriceGrossByTaxRate) {
		toSerialize["totalPriceGrossByTaxRate"] = o.TotalPriceGrossByTaxRate
	}
	toSerialize["ignoreChildPriceUpdates"] = o.IgnoreChildPriceUpdates
	// skip: deductedPrice is readOnly
	toSerialize["deductionFactor"] = o.DeductionFactor
	if !IsNil(o.AbsoluteDeduction) {
		toSerialize["absoluteDeduction"] = o.AbsoluteDeduction
	}
	// skip: totalPrice is readOnly
	// skip: totalPriceGross is readOnly
	// skip: totalPriceGrossDeducted is readOnly
	toSerialize["priceType"] = o.PriceType
	if !IsNil(o.ShortText) {
		toSerialize["shortText"] = o.ShortText
	}
	toSerialize["comissionStatus"] = o.ComissionStatus
	if !IsNil(o.ItemNumber) {
		toSerialize["itemNumber"] = o.ItemNumber
	}
	if !IsNil(o.ElementType) {
		toSerialize["elementType"] = o.ElementType
	}
	// skip: isLot is readOnly
	if !IsNil(o.AlternativeTo) {
		toSerialize["alternativeTo"] = o.AlternativeTo
	}
	if !IsNil(o.OenormProperties) {
		toSerialize["oenormProperties"] = o.OenormProperties
	}
	toSerialize["hierarchyLevel"] = o.HierarchyLevel
	if !IsNil(o.AddendumStatus) {
		toSerialize["addendumStatus"] = o.AddendumStatus
	}
	if !IsNil(o.AlternativeIdentifier) {
		toSerialize["alternativeIdentifier"] = o.AlternativeIdentifier
	}
	if !IsNil(o.AlternativeGroupIdentifier) {
		toSerialize["alternativeGroupIdentifier"] = o.AlternativeGroupIdentifier
	}
	return toSerialize, nil
}

type NullableServiceSpecificationGroupDto struct {
	value *ServiceSpecificationGroupDto
	isSet bool
}

func (v NullableServiceSpecificationGroupDto) Get() *ServiceSpecificationGroupDto {
	return v.value
}

func (v *NullableServiceSpecificationGroupDto) Set(val *ServiceSpecificationGroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSpecificationGroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSpecificationGroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSpecificationGroupDto(val *ServiceSpecificationGroupDto) *NullableServiceSpecificationGroupDto {
	return &NullableServiceSpecificationGroupDto{value: val, isSet: true}
}

func (v NullableServiceSpecificationGroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSpecificationGroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
