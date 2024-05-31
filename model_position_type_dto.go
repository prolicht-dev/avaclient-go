/*
AVACloud API 1.51.0

AVACloud API specification

API version: 1.51.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
	"fmt"
)

// PositionTypeDto This is a classification for Position elements.
type PositionTypeDto string

// List of PositionTypeDto
const (
	POSITIONTYPEDTO_REGULAR     PositionTypeDto = "Regular"
	POSITIONTYPEDTO_OPTIONAL    PositionTypeDto = "Optional"
	POSITIONTYPEDTO_ALTERNATIVE PositionTypeDto = "Alternative"
)

// All allowed values of PositionTypeDto enum
var AllowedPositionTypeDtoEnumValues = []PositionTypeDto{
	"Regular",
	"Optional",
	"Alternative",
}

func (v *PositionTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PositionTypeDto(value)
	for _, existing := range AllowedPositionTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PositionTypeDto", value)
}

// NewPositionTypeDtoFromValue returns a pointer to a valid PositionTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPositionTypeDtoFromValue(v string) (*PositionTypeDto, error) {
	ev := PositionTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PositionTypeDto: valid values are %v", v, AllowedPositionTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PositionTypeDto) IsValid() bool {
	for _, existing := range AllowedPositionTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PositionTypeDto value
func (v PositionTypeDto) Ptr() *PositionTypeDto {
	return &v
}

type NullablePositionTypeDto struct {
	value *PositionTypeDto
	isSet bool
}

func (v NullablePositionTypeDto) Get() *PositionTypeDto {
	return v.value
}

func (v *NullablePositionTypeDto) Set(val *PositionTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePositionTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePositionTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositionTypeDto(val *PositionTypeDto) *NullablePositionTypeDto {
	return &NullablePositionTypeDto{value: val, isSet: true}
}

func (v NullablePositionTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositionTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
