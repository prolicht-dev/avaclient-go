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

// StandardReferenceTypeDto This enumeration identifies a pre-known standard used for referencing standardized descriptions.
type StandardReferenceTypeDto string

// List of StandardReferenceTypeDto
const (
	STANDARDREFERENCETYPEDTO_UNKNOWN    StandardReferenceTypeDto = "Unknown"
	STANDARDREFERENCETYPEDTO_ST_LB      StandardReferenceTypeDto = "StLB"
	STANDARDREFERENCETYPEDTO_ST_LK      StandardReferenceTypeDto = "StLK"
	STANDARDREFERENCETYPEDTO_STLB_BAU_Z StandardReferenceTypeDto = "STLBBauZ"
)

// All allowed values of StandardReferenceTypeDto enum
var AllowedStandardReferenceTypeDtoEnumValues = []StandardReferenceTypeDto{
	"Unknown",
	"StLB",
	"StLK",
	"STLBBauZ",
}

func (v *StandardReferenceTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StandardReferenceTypeDto(value)
	for _, existing := range AllowedStandardReferenceTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StandardReferenceTypeDto", value)
}

// NewStandardReferenceTypeDtoFromValue returns a pointer to a valid StandardReferenceTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStandardReferenceTypeDtoFromValue(v string) (*StandardReferenceTypeDto, error) {
	ev := StandardReferenceTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StandardReferenceTypeDto: valid values are %v", v, AllowedStandardReferenceTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StandardReferenceTypeDto) IsValid() bool {
	for _, existing := range AllowedStandardReferenceTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StandardReferenceTypeDto value
func (v StandardReferenceTypeDto) Ptr() *StandardReferenceTypeDto {
	return &v
}

type NullableStandardReferenceTypeDto struct {
	value *StandardReferenceTypeDto
	isSet bool
}

func (v NullableStandardReferenceTypeDto) Get() *StandardReferenceTypeDto {
	return v.value
}

func (v *NullableStandardReferenceTypeDto) Set(val *StandardReferenceTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardReferenceTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardReferenceTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardReferenceTypeDto(val *StandardReferenceTypeDto) *NullableStandardReferenceTypeDto {
	return &NullableStandardReferenceTypeDto{value: val, isSet: true}
}

func (v NullableStandardReferenceTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardReferenceTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
