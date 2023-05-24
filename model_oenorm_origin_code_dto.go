/*
AVACloud API 1.41.8

AVACloud API specification

API version: 1.41.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
	"fmt"
)

// OenormOriginCodeDto This indicates where the content of this element originates, if set. It corresponds to 'herkunftskennzeichen' in ÖNorm
type OenormOriginCodeDto string

// List of OenormOriginCodeDto
const (
	OENORMORIGINCODEDTO_UNKNOWN              OenormOriginCodeDto = "Unknown"
	OENORMORIGINCODEDTO_FREE_TEXT            OenormOriginCodeDto = "FreeText"
	OENORMORIGINCODEDTO_ADDITION_DESCRIPTION OenormOriginCodeDto = "AdditionDescription"
)

// All allowed values of OenormOriginCodeDto enum
var AllowedOenormOriginCodeDtoEnumValues = []OenormOriginCodeDto{
	"Unknown",
	"FreeText",
	"AdditionDescription",
}

func (v *OenormOriginCodeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OenormOriginCodeDto(value)
	for _, existing := range AllowedOenormOriginCodeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OenormOriginCodeDto", value)
}

// NewOenormOriginCodeDtoFromValue returns a pointer to a valid OenormOriginCodeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOenormOriginCodeDtoFromValue(v string) (*OenormOriginCodeDto, error) {
	ev := OenormOriginCodeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OenormOriginCodeDto: valid values are %v", v, AllowedOenormOriginCodeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OenormOriginCodeDto) IsValid() bool {
	for _, existing := range AllowedOenormOriginCodeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OenormOriginCodeDto value
func (v OenormOriginCodeDto) Ptr() *OenormOriginCodeDto {
	return &v
}

type NullableOenormOriginCodeDto struct {
	value *OenormOriginCodeDto
	isSet bool
}

func (v NullableOenormOriginCodeDto) Get() *OenormOriginCodeDto {
	return v.value
}

func (v *NullableOenormOriginCodeDto) Set(val *OenormOriginCodeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOenormOriginCodeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOenormOriginCodeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOenormOriginCodeDto(val *OenormOriginCodeDto) *NullableOenormOriginCodeDto {
	return &NullableOenormOriginCodeDto{value: val, isSet: true}
}

func (v NullableOenormOriginCodeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOenormOriginCodeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
