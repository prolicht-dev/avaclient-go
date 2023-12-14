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

// ValidationType This enumeration represents the possible conversion targets when performing a validation
type ValidationType string

// List of ValidationType
const (
	VALIDATIONTYPE_PROJECT ValidationType = "Project"
	VALIDATIONTYPE_GAEB    ValidationType = "Gaeb"
	VALIDATIONTYPE_OENORM  ValidationType = "Oenorm"
)

// All allowed values of ValidationType enum
var AllowedValidationTypeEnumValues = []ValidationType{
	"Project",
	"Gaeb",
	"Oenorm",
}

func (v *ValidationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ValidationType(value)
	for _, existing := range AllowedValidationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ValidationType", value)
}

// NewValidationTypeFromValue returns a pointer to a valid ValidationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewValidationTypeFromValue(v string) (*ValidationType, error) {
	ev := ValidationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ValidationType: valid values are %v", v, AllowedValidationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ValidationType) IsValid() bool {
	for _, existing := range AllowedValidationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValidationType value
func (v ValidationType) Ptr() *ValidationType {
	return &v
}

type NullableValidationType struct {
	value *ValidationType
	isSet bool
}

func (v NullableValidationType) Get() *ValidationType {
	return v.value
}

func (v *NullableValidationType) Set(val *ValidationType) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationType) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationType(val *ValidationType) *NullableValidationType {
	return &NullableValidationType{value: val, isSet: true}
}

func (v NullableValidationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
