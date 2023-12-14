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

// checks if the ClaimGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimGet{}

// ClaimGet struct for ClaimGet
type ClaimGet struct {
	Issuer *string `json:"issuer,omitempty"`
	Type   *string `json:"type,omitempty"`
	Value  *string `json:"value,omitempty"`
}

// NewClaimGet instantiates a new ClaimGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimGet() *ClaimGet {
	this := ClaimGet{}
	return &this
}

// NewClaimGetWithDefaults instantiates a new ClaimGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimGetWithDefaults() *ClaimGet {
	this := ClaimGet{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *ClaimGet) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimGet) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *ClaimGet) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *ClaimGet) SetIssuer(v string) {
	o.Issuer = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClaimGet) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimGet) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClaimGet) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClaimGet) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ClaimGet) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimGet) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ClaimGet) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ClaimGet) SetValue(v string) {
	o.Value = &v
}

func (o ClaimGet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableClaimGet struct {
	value *ClaimGet
	isSet bool
}

func (v NullableClaimGet) Get() *ClaimGet {
	return v.value
}

func (v *NullableClaimGet) Set(val *ClaimGet) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimGet) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimGet(val *ClaimGet) *NullableClaimGet {
	return &NullableClaimGet{value: val, isSet: true}
}

func (v NullableClaimGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
