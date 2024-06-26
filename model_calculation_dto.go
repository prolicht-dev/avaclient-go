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

// checks if the CalculationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalculationDto{}

// CalculationDto This class holds means to calculate mathematical expressions from given strings.
type CalculationDto struct {
	// Descriptive text for this calculation.
	Description *string `json:"description,omitempty"`
	// This Calculation's mathematical expression. Please note that thousands separators are not supported. Both comma and point will be treated as decimal separators.
	Formula *string `json:"formula,omitempty"`
	// The calculated result from the formula, 0 if invalid.
	Result float64 `json:"result"`
	// Whether the Formula is a valid expression.
	Valid bool `json:"valid"`
	// Will be -1 if the Formula is correct, else it will show the position in the formula where an error was encountered. This is a zero based index.
	ErrorPositionInLine int32 `json:"errorPositionInLine"`
	// These are Catalogues that are used within this Calculation. Catalogues are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. They are propagate to all child elements, e.g. other containers and QuantityAssignments. In the context of a ServiceSpecification, all elements share the same instance of the collection.
	ProjectCatalogues []CatalogueDto `json:"projectCatalogues,omitempty"`
	// Referenced catalogues for this Calculation.
	CatalogueReferences []CatalogueReferenceDto `json:"catalogueReferences,omitempty"`
}

// NewCalculationDto instantiates a new CalculationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalculationDto(result float64, valid bool, errorPositionInLine int32) *CalculationDto {
	this := CalculationDto{}
	this.Result = result
	this.Valid = valid
	this.ErrorPositionInLine = errorPositionInLine
	return &this
}

// NewCalculationDtoWithDefaults instantiates a new CalculationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalculationDtoWithDefaults() *CalculationDto {
	this := CalculationDto{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CalculationDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculationDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CalculationDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CalculationDto) SetDescription(v string) {
	o.Description = &v
}

// GetFormula returns the Formula field value if set, zero value otherwise.
func (o *CalculationDto) GetFormula() string {
	if o == nil || IsNil(o.Formula) {
		var ret string
		return ret
	}
	return *o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculationDto) GetFormulaOk() (*string, bool) {
	if o == nil || IsNil(o.Formula) {
		return nil, false
	}
	return o.Formula, true
}

// HasFormula returns a boolean if a field has been set.
func (o *CalculationDto) HasFormula() bool {
	if o != nil && !IsNil(o.Formula) {
		return true
	}

	return false
}

// SetFormula gets a reference to the given string and assigns it to the Formula field.
func (o *CalculationDto) SetFormula(v string) {
	o.Formula = &v
}

// GetResult returns the Result field value
func (o *CalculationDto) GetResult() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CalculationDto) GetResultOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CalculationDto) SetResult(v float64) {
	o.Result = v
}

// GetValid returns the Valid field value
func (o *CalculationDto) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *CalculationDto) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *CalculationDto) SetValid(v bool) {
	o.Valid = v
}

// GetErrorPositionInLine returns the ErrorPositionInLine field value
func (o *CalculationDto) GetErrorPositionInLine() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ErrorPositionInLine
}

// GetErrorPositionInLineOk returns a tuple with the ErrorPositionInLine field value
// and a boolean to check if the value has been set.
func (o *CalculationDto) GetErrorPositionInLineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorPositionInLine, true
}

// SetErrorPositionInLine sets field value
func (o *CalculationDto) SetErrorPositionInLine(v int32) {
	o.ErrorPositionInLine = v
}

// GetProjectCatalogues returns the ProjectCatalogues field value if set, zero value otherwise.
func (o *CalculationDto) GetProjectCatalogues() []CatalogueDto {
	if o == nil || IsNil(o.ProjectCatalogues) {
		var ret []CatalogueDto
		return ret
	}
	return o.ProjectCatalogues
}

// GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculationDto) GetProjectCataloguesOk() ([]CatalogueDto, bool) {
	if o == nil || IsNil(o.ProjectCatalogues) {
		return nil, false
	}
	return o.ProjectCatalogues, true
}

// HasProjectCatalogues returns a boolean if a field has been set.
func (o *CalculationDto) HasProjectCatalogues() bool {
	if o != nil && !IsNil(o.ProjectCatalogues) {
		return true
	}

	return false
}

// SetProjectCatalogues gets a reference to the given []CatalogueDto and assigns it to the ProjectCatalogues field.
func (o *CalculationDto) SetProjectCatalogues(v []CatalogueDto) {
	o.ProjectCatalogues = v
}

// GetCatalogueReferences returns the CatalogueReferences field value if set, zero value otherwise.
func (o *CalculationDto) GetCatalogueReferences() []CatalogueReferenceDto {
	if o == nil || IsNil(o.CatalogueReferences) {
		var ret []CatalogueReferenceDto
		return ret
	}
	return o.CatalogueReferences
}

// GetCatalogueReferencesOk returns a tuple with the CatalogueReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculationDto) GetCatalogueReferencesOk() ([]CatalogueReferenceDto, bool) {
	if o == nil || IsNil(o.CatalogueReferences) {
		return nil, false
	}
	return o.CatalogueReferences, true
}

// HasCatalogueReferences returns a boolean if a field has been set.
func (o *CalculationDto) HasCatalogueReferences() bool {
	if o != nil && !IsNil(o.CatalogueReferences) {
		return true
	}

	return false
}

// SetCatalogueReferences gets a reference to the given []CatalogueReferenceDto and assigns it to the CatalogueReferences field.
func (o *CalculationDto) SetCatalogueReferences(v []CatalogueReferenceDto) {
	o.CatalogueReferences = v
}

func (o CalculationDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalculationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Formula) {
		toSerialize["formula"] = o.Formula
	}
	// skip: result is readOnly
	// skip: valid is readOnly
	// skip: errorPositionInLine is readOnly
	if !IsNil(o.ProjectCatalogues) {
		toSerialize["projectCatalogues"] = o.ProjectCatalogues
	}
	if !IsNil(o.CatalogueReferences) {
		toSerialize["catalogueReferences"] = o.CatalogueReferences
	}
	return toSerialize, nil
}

type NullableCalculationDto struct {
	value *CalculationDto
	isSet bool
}

func (v NullableCalculationDto) Get() *CalculationDto {
	return v.value
}

func (v *NullableCalculationDto) Set(val *CalculationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCalculationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCalculationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalculationDto(val *CalculationDto) *NullableCalculationDto {
	return &NullableCalculationDto{value: val, isSet: true}
}

func (v NullableCalculationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalculationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
