/*
AVACloud API 2.0.0

AVACloud API specification

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
)

// checks if the FlatAvaProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlatAvaProject{}

// FlatAvaProject This model contains a flattened represention of an AVA project, which makes it easy to use in systems that expect elements in a list form.
type FlatAvaProject struct {
	Project *ProjectDto `json:"project,omitempty"`
	// The flattened elements of the project.
	Elements []FlatElementDto `json:"elements,omitempty"`
}

// NewFlatAvaProject instantiates a new FlatAvaProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlatAvaProject() *FlatAvaProject {
	this := FlatAvaProject{}
	return &this
}

// NewFlatAvaProjectWithDefaults instantiates a new FlatAvaProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlatAvaProjectWithDefaults() *FlatAvaProject {
	this := FlatAvaProject{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *FlatAvaProject) GetProject() ProjectDto {
	if o == nil || IsNil(o.Project) {
		var ret ProjectDto
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatAvaProject) GetProjectOk() (*ProjectDto, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *FlatAvaProject) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given ProjectDto and assigns it to the Project field.
func (o *FlatAvaProject) SetProject(v ProjectDto) {
	o.Project = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *FlatAvaProject) GetElements() []FlatElementDto {
	if o == nil || IsNil(o.Elements) {
		var ret []FlatElementDto
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatAvaProject) GetElementsOk() ([]FlatElementDto, bool) {
	if o == nil || IsNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *FlatAvaProject) HasElements() bool {
	if o != nil && !IsNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []FlatElementDto and assigns it to the Elements field.
func (o *FlatAvaProject) SetElements(v []FlatElementDto) {
	o.Elements = v
}

func (o FlatAvaProject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlatAvaProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Elements) {
		toSerialize["elements"] = o.Elements
	}
	return toSerialize, nil
}

type NullableFlatAvaProject struct {
	value *FlatAvaProject
	isSet bool
}

func (v NullableFlatAvaProject) Get() *FlatAvaProject {
	return v.value
}

func (v *NullableFlatAvaProject) Set(val *FlatAvaProject) {
	v.value = val
	v.isSet = true
}

func (v NullableFlatAvaProject) IsSet() bool {
	return v.isSet
}

func (v *NullableFlatAvaProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlatAvaProject(val *FlatAvaProject) *NullableFlatAvaProject {
	return &NullableFlatAvaProject{value: val, isSet: true}
}

func (v NullableFlatAvaProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlatAvaProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}