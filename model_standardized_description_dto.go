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

// checks if the StandardizedDescriptionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardizedDescriptionDto{}

// StandardizedDescriptionDto This class represents a standardized description. This means that instead of solely relying on texts to describe a service, external standards and definitions are referenced for a common understanding.
type StandardizedDescriptionDto struct {
	StandardReferenceType StandardReferenceTypeDto `json:"standardReferenceType"`
	// This string property is the identifier to map to the references standard. Its type is given in the StandardReferenceType
	StandardReference *string           `json:"standardReference,omitempty"`
	StlbReference     *STLBReferenceDto `json:"stlbReference,omitempty"`
}

// NewStandardizedDescriptionDto instantiates a new StandardizedDescriptionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardizedDescriptionDto(standardReferenceType StandardReferenceTypeDto) *StandardizedDescriptionDto {
	this := StandardizedDescriptionDto{}
	this.StandardReferenceType = standardReferenceType
	return &this
}

// NewStandardizedDescriptionDtoWithDefaults instantiates a new StandardizedDescriptionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardizedDescriptionDtoWithDefaults() *StandardizedDescriptionDto {
	this := StandardizedDescriptionDto{}
	return &this
}

// GetStandardReferenceType returns the StandardReferenceType field value
func (o *StandardizedDescriptionDto) GetStandardReferenceType() StandardReferenceTypeDto {
	if o == nil {
		var ret StandardReferenceTypeDto
		return ret
	}

	return o.StandardReferenceType
}

// GetStandardReferenceTypeOk returns a tuple with the StandardReferenceType field value
// and a boolean to check if the value has been set.
func (o *StandardizedDescriptionDto) GetStandardReferenceTypeOk() (*StandardReferenceTypeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StandardReferenceType, true
}

// SetStandardReferenceType sets field value
func (o *StandardizedDescriptionDto) SetStandardReferenceType(v StandardReferenceTypeDto) {
	o.StandardReferenceType = v
}

// GetStandardReference returns the StandardReference field value if set, zero value otherwise.
func (o *StandardizedDescriptionDto) GetStandardReference() string {
	if o == nil || IsNil(o.StandardReference) {
		var ret string
		return ret
	}
	return *o.StandardReference
}

// GetStandardReferenceOk returns a tuple with the StandardReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardizedDescriptionDto) GetStandardReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.StandardReference) {
		return nil, false
	}
	return o.StandardReference, true
}

// HasStandardReference returns a boolean if a field has been set.
func (o *StandardizedDescriptionDto) HasStandardReference() bool {
	if o != nil && !IsNil(o.StandardReference) {
		return true
	}

	return false
}

// SetStandardReference gets a reference to the given string and assigns it to the StandardReference field.
func (o *StandardizedDescriptionDto) SetStandardReference(v string) {
	o.StandardReference = &v
}

// GetStlbReference returns the StlbReference field value if set, zero value otherwise.
func (o *StandardizedDescriptionDto) GetStlbReference() STLBReferenceDto {
	if o == nil || IsNil(o.StlbReference) {
		var ret STLBReferenceDto
		return ret
	}
	return *o.StlbReference
}

// GetStlbReferenceOk returns a tuple with the StlbReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardizedDescriptionDto) GetStlbReferenceOk() (*STLBReferenceDto, bool) {
	if o == nil || IsNil(o.StlbReference) {
		return nil, false
	}
	return o.StlbReference, true
}

// HasStlbReference returns a boolean if a field has been set.
func (o *StandardizedDescriptionDto) HasStlbReference() bool {
	if o != nil && !IsNil(o.StlbReference) {
		return true
	}

	return false
}

// SetStlbReference gets a reference to the given STLBReferenceDto and assigns it to the StlbReference field.
func (o *StandardizedDescriptionDto) SetStlbReference(v STLBReferenceDto) {
	o.StlbReference = &v
}

func (o StandardizedDescriptionDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardizedDescriptionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["standardReferenceType"] = o.StandardReferenceType
	if !IsNil(o.StandardReference) {
		toSerialize["standardReference"] = o.StandardReference
	}
	if !IsNil(o.StlbReference) {
		toSerialize["stlbReference"] = o.StlbReference
	}
	return toSerialize, nil
}

type NullableStandardizedDescriptionDto struct {
	value *StandardizedDescriptionDto
	isSet bool
}

func (v NullableStandardizedDescriptionDto) Get() *StandardizedDescriptionDto {
	return v.value
}

func (v *NullableStandardizedDescriptionDto) Set(val *StandardizedDescriptionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardizedDescriptionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardizedDescriptionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardizedDescriptionDto(val *StandardizedDescriptionDto) *NullableStandardizedDescriptionDto {
	return &NullableStandardizedDescriptionDto{value: val, isSet: true}
}

func (v NullableStandardizedDescriptionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardizedDescriptionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
