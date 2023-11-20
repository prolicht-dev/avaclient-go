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

// checks if the OenormNoteTextPropertiesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OenormNoteTextPropertiesDto{}

// OenormNoteTextPropertiesDto This class models special properties that only apply to some exchange scenarios where ÖNorm is used. It is special for NoteTexts.
type OenormNoteTextPropertiesDto struct {
	OriginCode OenormOriginCodeDto `json:"originCode"`
}

// NewOenormNoteTextPropertiesDto instantiates a new OenormNoteTextPropertiesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOenormNoteTextPropertiesDto(originCode OenormOriginCodeDto) *OenormNoteTextPropertiesDto {
	this := OenormNoteTextPropertiesDto{}
	this.OriginCode = originCode
	return &this
}

// NewOenormNoteTextPropertiesDtoWithDefaults instantiates a new OenormNoteTextPropertiesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOenormNoteTextPropertiesDtoWithDefaults() *OenormNoteTextPropertiesDto {
	this := OenormNoteTextPropertiesDto{}
	return &this
}

// GetOriginCode returns the OriginCode field value
func (o *OenormNoteTextPropertiesDto) GetOriginCode() OenormOriginCodeDto {
	if o == nil {
		var ret OenormOriginCodeDto
		return ret
	}

	return o.OriginCode
}

// GetOriginCodeOk returns a tuple with the OriginCode field value
// and a boolean to check if the value has been set.
func (o *OenormNoteTextPropertiesDto) GetOriginCodeOk() (*OenormOriginCodeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginCode, true
}

// SetOriginCode sets field value
func (o *OenormNoteTextPropertiesDto) SetOriginCode(v OenormOriginCodeDto) {
	o.OriginCode = v
}

func (o OenormNoteTextPropertiesDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OenormNoteTextPropertiesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["originCode"] = o.OriginCode
	return toSerialize, nil
}

type NullableOenormNoteTextPropertiesDto struct {
	value *OenormNoteTextPropertiesDto
	isSet bool
}

func (v NullableOenormNoteTextPropertiesDto) Get() *OenormNoteTextPropertiesDto {
	return v.value
}

func (v *NullableOenormNoteTextPropertiesDto) Set(val *OenormNoteTextPropertiesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOenormNoteTextPropertiesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOenormNoteTextPropertiesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOenormNoteTextPropertiesDto(val *OenormNoteTextPropertiesDto) *NullableOenormNoteTextPropertiesDto {
	return &NullableOenormNoteTextPropertiesDto{value: val, isSet: true}
}

func (v NullableOenormNoteTextPropertiesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOenormNoteTextPropertiesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
