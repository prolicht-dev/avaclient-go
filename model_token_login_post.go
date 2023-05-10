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

// checks if the TokenLoginPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenLoginPost{}

// TokenLoginPost struct for TokenLoginPost
type TokenLoginPost struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

// NewTokenLoginPost instantiates a new TokenLoginPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenLoginPost(identifier string, password string) *TokenLoginPost {
	this := TokenLoginPost{}
	this.Identifier = identifier
	this.Password = password
	return &this
}

// NewTokenLoginPostWithDefaults instantiates a new TokenLoginPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenLoginPostWithDefaults() *TokenLoginPost {
	this := TokenLoginPost{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *TokenLoginPost) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *TokenLoginPost) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *TokenLoginPost) SetIdentifier(v string) {
	o.Identifier = v
}

// GetPassword returns the Password field value
func (o *TokenLoginPost) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *TokenLoginPost) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *TokenLoginPost) SetPassword(v string) {
	o.Password = v
}

func (o TokenLoginPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenLoginPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableTokenLoginPost struct {
	value *TokenLoginPost
	isSet bool
}

func (v NullableTokenLoginPost) Get() *TokenLoginPost {
	return v.value
}

func (v *NullableTokenLoginPost) Set(val *TokenLoginPost) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenLoginPost) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenLoginPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenLoginPost(val *TokenLoginPost) *NullableTokenLoginPost {
	return &NullableTokenLoginPost{value: val, isSet: true}
}

func (v NullableTokenLoginPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenLoginPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
