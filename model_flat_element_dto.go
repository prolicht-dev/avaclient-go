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

// checks if the FlatElementDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlatElementDto{}

// FlatElementDto This model contains a flattened represention of an AVA project, which
type FlatElementDto struct {
	// If this is not null, then this contains the id of the previous element in the hierarchy on the same level.
	PreviousElementId *string `json:"previousElementId,omitempty"`
	// If this is not null, then this contains the id of the parent element.
	ParentElementId *string      `json:"parentElementId,omitempty"`
	Element         *IElementDto `json:"element,omitempty"`
}

// NewFlatElementDto instantiates a new FlatElementDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlatElementDto() *FlatElementDto {
	this := FlatElementDto{}
	return &this
}

// NewFlatElementDtoWithDefaults instantiates a new FlatElementDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlatElementDtoWithDefaults() *FlatElementDto {
	this := FlatElementDto{}
	return &this
}

// GetPreviousElementId returns the PreviousElementId field value if set, zero value otherwise.
func (o *FlatElementDto) GetPreviousElementId() string {
	if o == nil || IsNil(o.PreviousElementId) {
		var ret string
		return ret
	}
	return *o.PreviousElementId
}

// GetPreviousElementIdOk returns a tuple with the PreviousElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatElementDto) GetPreviousElementIdOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousElementId) {
		return nil, false
	}
	return o.PreviousElementId, true
}

// HasPreviousElementId returns a boolean if a field has been set.
func (o *FlatElementDto) HasPreviousElementId() bool {
	if o != nil && !IsNil(o.PreviousElementId) {
		return true
	}

	return false
}

// SetPreviousElementId gets a reference to the given string and assigns it to the PreviousElementId field.
func (o *FlatElementDto) SetPreviousElementId(v string) {
	o.PreviousElementId = &v
}

// GetParentElementId returns the ParentElementId field value if set, zero value otherwise.
func (o *FlatElementDto) GetParentElementId() string {
	if o == nil || IsNil(o.ParentElementId) {
		var ret string
		return ret
	}
	return *o.ParentElementId
}

// GetParentElementIdOk returns a tuple with the ParentElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatElementDto) GetParentElementIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentElementId) {
		return nil, false
	}
	return o.ParentElementId, true
}

// HasParentElementId returns a boolean if a field has been set.
func (o *FlatElementDto) HasParentElementId() bool {
	if o != nil && !IsNil(o.ParentElementId) {
		return true
	}

	return false
}

// SetParentElementId gets a reference to the given string and assigns it to the ParentElementId field.
func (o *FlatElementDto) SetParentElementId(v string) {
	o.ParentElementId = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *FlatElementDto) GetElement() IElementDto {
	if o == nil || IsNil(o.Element) {
		var ret IElementDto
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatElementDto) GetElementOk() (*IElementDto, bool) {
	if o == nil || IsNil(o.Element) {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *FlatElementDto) HasElement() bool {
	if o != nil && !IsNil(o.Element) {
		return true
	}

	return false
}

// SetElement gets a reference to the given IElementDto and assigns it to the Element field.
func (o *FlatElementDto) SetElement(v IElementDto) {
	o.Element = &v
}

func (o FlatElementDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlatElementDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreviousElementId) {
		toSerialize["previousElementId"] = o.PreviousElementId
	}
	if !IsNil(o.ParentElementId) {
		toSerialize["parentElementId"] = o.ParentElementId
	}
	if !IsNil(o.Element) {
		toSerialize["element"] = o.Element
	}
	return toSerialize, nil
}

type NullableFlatElementDto struct {
	value *FlatElementDto
	isSet bool
}

func (v NullableFlatElementDto) Get() *FlatElementDto {
	return v.value
}

func (v *NullableFlatElementDto) Set(val *FlatElementDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFlatElementDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFlatElementDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlatElementDto(val *FlatElementDto) *NullableFlatElementDto {
	return &NullableFlatElementDto{value: val, isSet: true}
}

func (v NullableFlatElementDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlatElementDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}