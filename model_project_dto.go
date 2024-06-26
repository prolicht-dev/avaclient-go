/*
AVACloud API 1.51.0

AVACloud API specification

API version: 1.51.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
)

// checks if the ProjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectDto{}

// ProjectDto A Project contains all relevant information for a construction project.
type ProjectDto struct {
	// Elements GUID identifier.
	Id string `json:"id"`
	// This property controls the accuracy of all price properties, meaning how many decimal places are preserved in calculations. It defaults to DEFAULT_PRICE_ACCURACY. Please see the Dangl.AVA documentation for further information about decimal precision. If UnitPriceAccuracy is set, then this property is ignored for unit prices.
	PriceAccuracy int32 `json:"priceAccuracy"`
	// This property controls the accuracy of position unit price properties, meaning how many decimal places are preserved in calculations. Please see the Dangl.AVA documentation for further information about decimal precision. This can be separately set from PriceAccuracy, and it only controls the accuracy of the unit price of positions, not the total price. It defaults to null, which means the standard PriceAccuracy is used for unit prices.
	UnitPriceAccuracy *int32 `json:"unitPriceAccuracy,omitempty"`
	// This forces total prices to be the strict product of quantities times unit price in positions. It is enabled by default. If this is disabled, both the unit price and the total price of positions is calculated from the non-rounded values. Please see the documentation for a more detailed explanation of this setting.
	ForceStrictTotals  bool                   `json:"forceStrictTotals"`
	PriceRoundingMode  PriceRoundingModeDto   `json:"priceRoundingMode"`
	ProjectInformation *ProjectInformationDto `json:"projectInformation,omitempty"`
	// The ServiceSpecifications in this Project.
	ServiceSpecifications []ServiceSpecificationDto `json:"serviceSpecifications,omitempty"`
	// This is used to store the GAEB XML Id within this Project. This data is not used for any calculations or evaluations but only for GAEB serialization and deserialization.
	GaebXmlId *string `json:"gaebXmlId,omitempty"`
}

// NewProjectDto instantiates a new ProjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDto(id string, priceAccuracy int32, forceStrictTotals bool, priceRoundingMode PriceRoundingModeDto) *ProjectDto {
	this := ProjectDto{}
	this.Id = id
	this.PriceAccuracy = priceAccuracy
	this.ForceStrictTotals = forceStrictTotals
	this.PriceRoundingMode = priceRoundingMode
	return &this
}

// NewProjectDtoWithDefaults instantiates a new ProjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDtoWithDefaults() *ProjectDto {
	this := ProjectDto{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectDto) SetId(v string) {
	o.Id = v
}

// GetPriceAccuracy returns the PriceAccuracy field value
func (o *ProjectDto) GetPriceAccuracy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PriceAccuracy
}

// GetPriceAccuracyOk returns a tuple with the PriceAccuracy field value
// and a boolean to check if the value has been set.
func (o *ProjectDto) GetPriceAccuracyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceAccuracy, true
}

// SetPriceAccuracy sets field value
func (o *ProjectDto) SetPriceAccuracy(v int32) {
	o.PriceAccuracy = v
}

// GetUnitPriceAccuracy returns the UnitPriceAccuracy field value if set, zero value otherwise.
func (o *ProjectDto) GetUnitPriceAccuracy() int32 {
	if o == nil || IsNil(o.UnitPriceAccuracy) {
		var ret int32
		return ret
	}
	return *o.UnitPriceAccuracy
}

// GetUnitPriceAccuracyOk returns a tuple with the UnitPriceAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDto) GetUnitPriceAccuracyOk() (*int32, bool) {
	if o == nil || IsNil(o.UnitPriceAccuracy) {
		return nil, false
	}
	return o.UnitPriceAccuracy, true
}

// HasUnitPriceAccuracy returns a boolean if a field has been set.
func (o *ProjectDto) HasUnitPriceAccuracy() bool {
	if o != nil && !IsNil(o.UnitPriceAccuracy) {
		return true
	}

	return false
}

// SetUnitPriceAccuracy gets a reference to the given int32 and assigns it to the UnitPriceAccuracy field.
func (o *ProjectDto) SetUnitPriceAccuracy(v int32) {
	o.UnitPriceAccuracy = &v
}

// GetForceStrictTotals returns the ForceStrictTotals field value
func (o *ProjectDto) GetForceStrictTotals() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForceStrictTotals
}

// GetForceStrictTotalsOk returns a tuple with the ForceStrictTotals field value
// and a boolean to check if the value has been set.
func (o *ProjectDto) GetForceStrictTotalsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForceStrictTotals, true
}

// SetForceStrictTotals sets field value
func (o *ProjectDto) SetForceStrictTotals(v bool) {
	o.ForceStrictTotals = v
}

// GetPriceRoundingMode returns the PriceRoundingMode field value
func (o *ProjectDto) GetPriceRoundingMode() PriceRoundingModeDto {
	if o == nil {
		var ret PriceRoundingModeDto
		return ret
	}

	return o.PriceRoundingMode
}

// GetPriceRoundingModeOk returns a tuple with the PriceRoundingMode field value
// and a boolean to check if the value has been set.
func (o *ProjectDto) GetPriceRoundingModeOk() (*PriceRoundingModeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceRoundingMode, true
}

// SetPriceRoundingMode sets field value
func (o *ProjectDto) SetPriceRoundingMode(v PriceRoundingModeDto) {
	o.PriceRoundingMode = v
}

// GetProjectInformation returns the ProjectInformation field value if set, zero value otherwise.
func (o *ProjectDto) GetProjectInformation() ProjectInformationDto {
	if o == nil || IsNil(o.ProjectInformation) {
		var ret ProjectInformationDto
		return ret
	}
	return *o.ProjectInformation
}

// GetProjectInformationOk returns a tuple with the ProjectInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDto) GetProjectInformationOk() (*ProjectInformationDto, bool) {
	if o == nil || IsNil(o.ProjectInformation) {
		return nil, false
	}
	return o.ProjectInformation, true
}

// HasProjectInformation returns a boolean if a field has been set.
func (o *ProjectDto) HasProjectInformation() bool {
	if o != nil && !IsNil(o.ProjectInformation) {
		return true
	}

	return false
}

// SetProjectInformation gets a reference to the given ProjectInformationDto and assigns it to the ProjectInformation field.
func (o *ProjectDto) SetProjectInformation(v ProjectInformationDto) {
	o.ProjectInformation = &v
}

// GetServiceSpecifications returns the ServiceSpecifications field value if set, zero value otherwise.
func (o *ProjectDto) GetServiceSpecifications() []ServiceSpecificationDto {
	if o == nil || IsNil(o.ServiceSpecifications) {
		var ret []ServiceSpecificationDto
		return ret
	}
	return o.ServiceSpecifications
}

// GetServiceSpecificationsOk returns a tuple with the ServiceSpecifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDto) GetServiceSpecificationsOk() ([]ServiceSpecificationDto, bool) {
	if o == nil || IsNil(o.ServiceSpecifications) {
		return nil, false
	}
	return o.ServiceSpecifications, true
}

// HasServiceSpecifications returns a boolean if a field has been set.
func (o *ProjectDto) HasServiceSpecifications() bool {
	if o != nil && !IsNil(o.ServiceSpecifications) {
		return true
	}

	return false
}

// SetServiceSpecifications gets a reference to the given []ServiceSpecificationDto and assigns it to the ServiceSpecifications field.
func (o *ProjectDto) SetServiceSpecifications(v []ServiceSpecificationDto) {
	o.ServiceSpecifications = v
}

// GetGaebXmlId returns the GaebXmlId field value if set, zero value otherwise.
func (o *ProjectDto) GetGaebXmlId() string {
	if o == nil || IsNil(o.GaebXmlId) {
		var ret string
		return ret
	}
	return *o.GaebXmlId
}

// GetGaebXmlIdOk returns a tuple with the GaebXmlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDto) GetGaebXmlIdOk() (*string, bool) {
	if o == nil || IsNil(o.GaebXmlId) {
		return nil, false
	}
	return o.GaebXmlId, true
}

// HasGaebXmlId returns a boolean if a field has been set.
func (o *ProjectDto) HasGaebXmlId() bool {
	if o != nil && !IsNil(o.GaebXmlId) {
		return true
	}

	return false
}

// SetGaebXmlId gets a reference to the given string and assigns it to the GaebXmlId field.
func (o *ProjectDto) SetGaebXmlId(v string) {
	o.GaebXmlId = &v
}

func (o ProjectDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["priceAccuracy"] = o.PriceAccuracy
	if !IsNil(o.UnitPriceAccuracy) {
		toSerialize["unitPriceAccuracy"] = o.UnitPriceAccuracy
	}
	toSerialize["forceStrictTotals"] = o.ForceStrictTotals
	toSerialize["priceRoundingMode"] = o.PriceRoundingMode
	if !IsNil(o.ProjectInformation) {
		toSerialize["projectInformation"] = o.ProjectInformation
	}
	if !IsNil(o.ServiceSpecifications) {
		toSerialize["serviceSpecifications"] = o.ServiceSpecifications
	}
	if !IsNil(o.GaebXmlId) {
		toSerialize["gaebXmlId"] = o.GaebXmlId
	}
	return toSerialize, nil
}

type NullableProjectDto struct {
	value *ProjectDto
	isSet bool
}

func (v NullableProjectDto) Get() *ProjectDto {
	return v.value
}

func (v *NullableProjectDto) Set(val *ProjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDto(val *ProjectDto) *NullableProjectDto {
	return &NullableProjectDto{value: val, isSet: true}
}

func (v NullableProjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
