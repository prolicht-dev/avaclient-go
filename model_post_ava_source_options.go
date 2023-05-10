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

// checks if the PostAvaSourceOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAvaSourceOptions{}

// PostAvaSourceOptions Options for conversions from AVA
type PostAvaSourceOptions struct {
	// If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number.
	TryAutoGenerateItemNumbersAndSchema bool `json:"tryAutoGenerateItemNumbersAndSchema"`
}

// NewPostAvaSourceOptions instantiates a new PostAvaSourceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAvaSourceOptions(tryAutoGenerateItemNumbersAndSchema bool) *PostAvaSourceOptions {
	this := PostAvaSourceOptions{}
	this.TryAutoGenerateItemNumbersAndSchema = tryAutoGenerateItemNumbersAndSchema
	return &this
}

// NewPostAvaSourceOptionsWithDefaults instantiates a new PostAvaSourceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAvaSourceOptionsWithDefaults() *PostAvaSourceOptions {
	this := PostAvaSourceOptions{}
	return &this
}

// GetTryAutoGenerateItemNumbersAndSchema returns the TryAutoGenerateItemNumbersAndSchema field value
func (o *PostAvaSourceOptions) GetTryAutoGenerateItemNumbersAndSchema() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TryAutoGenerateItemNumbersAndSchema
}

// GetTryAutoGenerateItemNumbersAndSchemaOk returns a tuple with the TryAutoGenerateItemNumbersAndSchema field value
// and a boolean to check if the value has been set.
func (o *PostAvaSourceOptions) GetTryAutoGenerateItemNumbersAndSchemaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TryAutoGenerateItemNumbersAndSchema, true
}

// SetTryAutoGenerateItemNumbersAndSchema sets field value
func (o *PostAvaSourceOptions) SetTryAutoGenerateItemNumbersAndSchema(v bool) {
	o.TryAutoGenerateItemNumbersAndSchema = v
}

func (o PostAvaSourceOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAvaSourceOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tryAutoGenerateItemNumbersAndSchema"] = o.TryAutoGenerateItemNumbersAndSchema
	return toSerialize, nil
}

type NullablePostAvaSourceOptions struct {
	value *PostAvaSourceOptions
	isSet bool
}

func (v NullablePostAvaSourceOptions) Get() *PostAvaSourceOptions {
	return v.value
}

func (v *NullablePostAvaSourceOptions) Set(val *PostAvaSourceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAvaSourceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAvaSourceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAvaSourceOptions(val *PostAvaSourceOptions) *NullablePostAvaSourceOptions {
	return &NullablePostAvaSourceOptions{value: val, isSet: true}
}

func (v NullablePostAvaSourceOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAvaSourceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}