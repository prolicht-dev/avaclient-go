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

// checks if the LoginPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginPost{}

// LoginPost struct for LoginPost
type LoginPost struct {
	Identifier   string `json:"identifier"`
	Password     string `json:"password"`
	StaySignedIn bool   `json:"staySignedIn"`
}

// NewLoginPost instantiates a new LoginPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginPost(identifier string, password string, staySignedIn bool) *LoginPost {
	this := LoginPost{}
	this.Identifier = identifier
	this.Password = password
	this.StaySignedIn = staySignedIn
	return &this
}

// NewLoginPostWithDefaults instantiates a new LoginPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginPostWithDefaults() *LoginPost {
	this := LoginPost{}
	return &this
}

// GetIdentifier returns the Identifier field value
func (o *LoginPost) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *LoginPost) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *LoginPost) SetIdentifier(v string) {
	o.Identifier = v
}

// GetPassword returns the Password field value
func (o *LoginPost) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *LoginPost) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *LoginPost) SetPassword(v string) {
	o.Password = v
}

// GetStaySignedIn returns the StaySignedIn field value
func (o *LoginPost) GetStaySignedIn() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StaySignedIn
}

// GetStaySignedInOk returns a tuple with the StaySignedIn field value
// and a boolean to check if the value has been set.
func (o *LoginPost) GetStaySignedInOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StaySignedIn, true
}

// SetStaySignedIn sets field value
func (o *LoginPost) SetStaySignedIn(v bool) {
	o.StaySignedIn = v
}

func (o LoginPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identifier"] = o.Identifier
	toSerialize["password"] = o.Password
	toSerialize["staySignedIn"] = o.StaySignedIn
	return toSerialize, nil
}

type NullableLoginPost struct {
	value *LoginPost
	isSet bool
}

func (v NullableLoginPost) Get() *LoginPost {
	return v.value
}

func (v *NullableLoginPost) Set(val *LoginPost) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginPost) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginPost(val *LoginPost) *NullableLoginPost {
	return &NullableLoginPost{value: val, isSet: true}
}

func (v NullableLoginPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
