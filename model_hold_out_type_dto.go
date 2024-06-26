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

// HoldOutTypeDto Types of retentions to be used in positions
type HoldOutTypeDto string

// List of HoldOutTypeDto
const (
	HOLDOUTTYPEDTO_UNKNOWN   HoldOutTypeDto = "Unknown"
	HOLDOUTTYPEDTO_BASE      HoldOutTypeDto = "Base"
	HOLDOUTTYPEDTO_EXTENSION HoldOutTypeDto = "Extension"
	HOLDOUTTYPEDTO_RENT      HoldOutTypeDto = "Rent"
)

// All allowed values of HoldOutTypeDto enum
var AllowedHoldOutTypeDtoEnumValues = []HoldOutTypeDto{
	"Unknown",
	"Base",
	"Extension",
	"Rent",
}

func (v *HoldOutTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HoldOutTypeDto(value)
	for _, existing := range AllowedHoldOutTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HoldOutTypeDto", value)
}

// NewHoldOutTypeDtoFromValue returns a pointer to a valid HoldOutTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHoldOutTypeDtoFromValue(v string) (*HoldOutTypeDto, error) {
	ev := HoldOutTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HoldOutTypeDto: valid values are %v", v, AllowedHoldOutTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HoldOutTypeDto) IsValid() bool {
	for _, existing := range AllowedHoldOutTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HoldOutTypeDto value
func (v HoldOutTypeDto) Ptr() *HoldOutTypeDto {
	return &v
}

type NullableHoldOutTypeDto struct {
	value *HoldOutTypeDto
	isSet bool
}

func (v NullableHoldOutTypeDto) Get() *HoldOutTypeDto {
	return v.value
}

func (v *NullableHoldOutTypeDto) Set(val *HoldOutTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldOutTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldOutTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldOutTypeDto(val *HoldOutTypeDto) *NullableHoldOutTypeDto {
	return &NullableHoldOutTypeDto{value: val, isSet: true}
}

func (v NullableHoldOutTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldOutTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
