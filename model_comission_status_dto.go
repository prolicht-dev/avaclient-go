/*
AVACloud API 2.0.0

AVACloud API specification

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
	"fmt"
)

// ComissionStatusDto Indicates if this service specification item is commissioned (and therefore should be executed), postponed for later or undefined.
type ComissionStatusDto string

// List of ComissionStatusDto
const (
	COMISSIONSTATUSDTO_UNDEFINED    ComissionStatusDto = "Undefined"
	COMISSIONSTATUSDTO_COMMISSIONED ComissionStatusDto = "Commissioned"
	COMISSIONSTATUSDTO_POSTPONED    ComissionStatusDto = "Postponed"
	COMISSIONSTATUSDTO_REMOVED      ComissionStatusDto = "Removed"
)

// All allowed values of ComissionStatusDto enum
var AllowedComissionStatusDtoEnumValues = []ComissionStatusDto{
	"Undefined",
	"Commissioned",
	"Postponed",
	"Removed",
}

func (v *ComissionStatusDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ComissionStatusDto(value)
	for _, existing := range AllowedComissionStatusDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ComissionStatusDto", value)
}

// NewComissionStatusDtoFromValue returns a pointer to a valid ComissionStatusDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComissionStatusDtoFromValue(v string) (*ComissionStatusDto, error) {
	ev := ComissionStatusDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ComissionStatusDto: valid values are %v", v, AllowedComissionStatusDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComissionStatusDto) IsValid() bool {
	for _, existing := range AllowedComissionStatusDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ComissionStatusDto value
func (v ComissionStatusDto) Ptr() *ComissionStatusDto {
	return &v
}

type NullableComissionStatusDto struct {
	value *ComissionStatusDto
	isSet bool
}

func (v NullableComissionStatusDto) Get() *ComissionStatusDto {
	return v.value
}

func (v *NullableComissionStatusDto) Set(val *ComissionStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableComissionStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableComissionStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComissionStatusDto(val *ComissionStatusDto) *NullableComissionStatusDto {
	return &NullableComissionStatusDto{value: val, isSet: true}
}

func (v NullableComissionStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComissionStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
