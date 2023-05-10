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

// checks if the ForgotPasswordPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForgotPasswordPost{}

// ForgotPasswordPost struct for ForgotPasswordPost
type ForgotPasswordPost struct {
	Identifier         string   `json:"identifier"`
	PreferredLanguages []string `json:"preferredLanguages,omitempty"`
}

// NewForgotPasswordPost instantiates a new ForgotPasswordPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForgotPasswordPost(identifier string) *ForgotPasswordPost {
	this := ForgotPasswordPost{}
	this.Identifier = identifier
	return &this
}

// NewForgotPasswordPostWithDefaults instantiates a new ForgotPasswordPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForgotPasswordPostWithDefaults() *ForgotPasswordPost {
	this := ForgotPasswordPost{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *ForgotPasswordPost) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *ForgotPasswordPost) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *ForgotPasswordPost) SetIdentifier(v string) {
	o.Identifier = v
}

// GetPreferredLanguages returns the PreferredLanguages field value if set, zero value otherwise.
func (o *ForgotPasswordPost) GetPreferredLanguages() []string {
	if o == nil || IsNil(o.PreferredLanguages) {
		var ret []string
		return ret
	}
	return o.PreferredLanguages
}

// GetPreferredLanguagesOk returns a tuple with the PreferredLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForgotPasswordPost) GetPreferredLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.PreferredLanguages) {
		return nil, false
	}
	return o.PreferredLanguages, true
}

// HasPreferredLanguages returns a boolean if a field has been set.
func (o *ForgotPasswordPost) HasPreferredLanguages() bool {
	if o != nil && !IsNil(o.PreferredLanguages) {
		return true
	}

	return false
}

// SetPreferredLanguages gets a reference to the given []string and assigns it to the PreferredLanguages field.
func (o *ForgotPasswordPost) SetPreferredLanguages(v []string) {
	o.PreferredLanguages = v
}

func (o ForgotPasswordPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForgotPasswordPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	if !IsNil(o.PreferredLanguages) {
		toSerialize["preferredLanguages"] = o.PreferredLanguages
	}
	return toSerialize, nil
}

type NullableForgotPasswordPost struct {
	value *ForgotPasswordPost
	isSet bool
}

func (v NullableForgotPasswordPost) Get() *ForgotPasswordPost {
	return v.value
}

func (v *NullableForgotPasswordPost) Set(val *ForgotPasswordPost) {
	v.value = val
	v.isSet = true
}

func (v NullableForgotPasswordPost) IsSet() bool {
	return v.isSet
}

func (v *NullableForgotPasswordPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForgotPasswordPost(val *ForgotPasswordPost) *NullableForgotPasswordPost {
	return &NullableForgotPasswordPost{value: val, isSet: true}
}

func (v NullableForgotPasswordPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForgotPasswordPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
