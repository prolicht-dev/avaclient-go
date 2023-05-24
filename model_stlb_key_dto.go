/*
AVACloud API 1.41.8

AVACloud API specification

API version: 1.41.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
)

// checks if the STLBKeyDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &STLBKeyDto{}

// STLBKeyDto This class represents a single key reference within the German STLB \"Standardleistungsbuch Bau\"
type STLBKeyDto struct {
	// This identifier is required and uniquely describes a single reference within the STLB standard. It maps to \"ArtChrIdent\" in GAEB XML
	ArtIdentifier int32 `json:"artIdentifier"`
	// This optional index property further categorizes a single reference within the STLB standard. It maps to \"ArtChIdx\" in GAEB XML
	ArtIndex *int32 `json:"artIndex,omitempty"`
	// This optional identifier further specifies the execution kind of the reference in the STLB standard. It maps to \"ChVIdent\" in GAEB XML
	KindIdentifier *int32 `json:"kindIdentifier,omitempty"`
}

// NewSTLBKeyDto instantiates a new STLBKeyDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSTLBKeyDto(artIdentifier int32) *STLBKeyDto {
	this := STLBKeyDto{}
	this.ArtIdentifier = artIdentifier
	return &this
}

// NewSTLBKeyDtoWithDefaults instantiates a new STLBKeyDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSTLBKeyDtoWithDefaults() *STLBKeyDto {
	this := STLBKeyDto{}
	return &this
}

// GetArtIdentifier returns the ArtIdentifier field value
func (o *STLBKeyDto) GetArtIdentifier() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ArtIdentifier
}

// GetArtIdentifierOk returns a tuple with the ArtIdentifier field value
// and a boolean to check if the value has been set.
func (o *STLBKeyDto) GetArtIdentifierOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtIdentifier, true
}

// SetArtIdentifier sets field value
func (o *STLBKeyDto) SetArtIdentifier(v int32) {
	o.ArtIdentifier = v
}

// GetArtIndex returns the ArtIndex field value if set, zero value otherwise.
func (o *STLBKeyDto) GetArtIndex() int32 {
	if o == nil || IsNil(o.ArtIndex) {
		var ret int32
		return ret
	}
	return *o.ArtIndex
}

// GetArtIndexOk returns a tuple with the ArtIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STLBKeyDto) GetArtIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.ArtIndex) {
		return nil, false
	}
	return o.ArtIndex, true
}

// HasArtIndex returns a boolean if a field has been set.
func (o *STLBKeyDto) HasArtIndex() bool {
	if o != nil && !IsNil(o.ArtIndex) {
		return true
	}

	return false
}

// SetArtIndex gets a reference to the given int32 and assigns it to the ArtIndex field.
func (o *STLBKeyDto) SetArtIndex(v int32) {
	o.ArtIndex = &v
}

// GetKindIdentifier returns the KindIdentifier field value if set, zero value otherwise.
func (o *STLBKeyDto) GetKindIdentifier() int32 {
	if o == nil || IsNil(o.KindIdentifier) {
		var ret int32
		return ret
	}
	return *o.KindIdentifier
}

// GetKindIdentifierOk returns a tuple with the KindIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STLBKeyDto) GetKindIdentifierOk() (*int32, bool) {
	if o == nil || IsNil(o.KindIdentifier) {
		return nil, false
	}
	return o.KindIdentifier, true
}

// HasKindIdentifier returns a boolean if a field has been set.
func (o *STLBKeyDto) HasKindIdentifier() bool {
	if o != nil && !IsNil(o.KindIdentifier) {
		return true
	}

	return false
}

// SetKindIdentifier gets a reference to the given int32 and assigns it to the KindIdentifier field.
func (o *STLBKeyDto) SetKindIdentifier(v int32) {
	o.KindIdentifier = &v
}

func (o STLBKeyDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o STLBKeyDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["artIdentifier"] = o.ArtIdentifier
	if !IsNil(o.ArtIndex) {
		toSerialize["artIndex"] = o.ArtIndex
	}
	if !IsNil(o.KindIdentifier) {
		toSerialize["kindIdentifier"] = o.KindIdentifier
	}
	return toSerialize, nil
}

type NullableSTLBKeyDto struct {
	value *STLBKeyDto
	isSet bool
}

func (v NullableSTLBKeyDto) Get() *STLBKeyDto {
	return v.value
}

func (v *NullableSTLBKeyDto) Set(val *STLBKeyDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSTLBKeyDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSTLBKeyDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSTLBKeyDto(val *STLBKeyDto) *NullableSTLBKeyDto {
	return &NullableSTLBKeyDto{value: val, isSet: true}
}

func (v NullableSTLBKeyDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSTLBKeyDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
