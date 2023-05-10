/*
AVACloud API 1.41.4

AVACloud API specification

API version: 1.41.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
)

// checks if the ObjectValidationCheckDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectValidationCheckDetails{}

// ObjectValidationCheckDetails Basic information about Object check validation details, commonly used when checking GAEB 2000 files
type ObjectValidationCheckDetails struct {
	// The path under which the element with error was found.
	ElementPath *string `json:"elementPath,omitempty"`
}

// NewObjectValidationCheckDetails instantiates a new ObjectValidationCheckDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectValidationCheckDetails() *ObjectValidationCheckDetails {
	this := ObjectValidationCheckDetails{}
	return &this
}

// NewObjectValidationCheckDetailsWithDefaults instantiates a new ObjectValidationCheckDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectValidationCheckDetailsWithDefaults() *ObjectValidationCheckDetails {
	this := ObjectValidationCheckDetails{}
	return &this
}

// GetElementPath returns the ElementPath field value if set, zero value otherwise.
func (o *ObjectValidationCheckDetails) GetElementPath() string {
	if o == nil || IsNil(o.ElementPath) {
		var ret string
		return ret
	}
	return *o.ElementPath
}

// GetElementPathOk returns a tuple with the ElementPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectValidationCheckDetails) GetElementPathOk() (*string, bool) {
	if o == nil || IsNil(o.ElementPath) {
		return nil, false
	}
	return o.ElementPath, true
}

// HasElementPath returns a boolean if a field has been set.
func (o *ObjectValidationCheckDetails) HasElementPath() bool {
	if o != nil && !IsNil(o.ElementPath) {
		return true
	}

	return false
}

// SetElementPath gets a reference to the given string and assigns it to the ElementPath field.
func (o *ObjectValidationCheckDetails) SetElementPath(v string) {
	o.ElementPath = &v
}

func (o ObjectValidationCheckDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectValidationCheckDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ElementPath) {
		toSerialize["elementPath"] = o.ElementPath
	}
	return toSerialize, nil
}

type NullableObjectValidationCheckDetails struct {
	value *ObjectValidationCheckDetails
	isSet bool
}

func (v NullableObjectValidationCheckDetails) Get() *ObjectValidationCheckDetails {
	return v.value
}

func (v *NullableObjectValidationCheckDetails) Set(val *ObjectValidationCheckDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectValidationCheckDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectValidationCheckDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectValidationCheckDetails(val *ObjectValidationCheckDetails) *NullableObjectValidationCheckDetails {
	return &NullableObjectValidationCheckDetails{value: val, isSet: true}
}

func (v NullableObjectValidationCheckDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectValidationCheckDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
