/*
AVACloud API 1.41.4

AVACloud API specification

API version: 1.41.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
	"fmt"
)

// ValidationCheckType Enumeration for the different types of checks performed
type ValidationCheckType string

// List of ValidationCheckType
const (
	VALIDATIONCHECKTYPE_GENERAL            ValidationCheckType = "General"
	VALIDATIONCHECKTYPE_XML_SCHEMA_CHECK   ValidationCheckType = "XmlSchemaCheck"
	VALIDATIONCHECKTYPE_OBJECT_VALIDATION  ValidationCheckType = "ObjectValidation"
	VALIDATIONCHECKTYPE_PROJECT_VALIDATION ValidationCheckType = "ProjectValidation"
)

// All allowed values of ValidationCheckType enum
var AllowedValidationCheckTypeEnumValues = []ValidationCheckType{
	"General",
	"XmlSchemaCheck",
	"ObjectValidation",
	"ProjectValidation",
}

func (v *ValidationCheckType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ValidationCheckType(value)
	for _, existing := range AllowedValidationCheckTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ValidationCheckType", value)
}

// NewValidationCheckTypeFromValue returns a pointer to a valid ValidationCheckType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewValidationCheckTypeFromValue(v string) (*ValidationCheckType, error) {
	ev := ValidationCheckType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ValidationCheckType: valid values are %v", v, AllowedValidationCheckTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ValidationCheckType) IsValid() bool {
	for _, existing := range AllowedValidationCheckTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValidationCheckType value
func (v ValidationCheckType) Ptr() *ValidationCheckType {
	return &v
}

type NullableValidationCheckType struct {
	value *ValidationCheckType
	isSet bool
}

func (v NullableValidationCheckType) Get() *ValidationCheckType {
	return v.value
}

func (v *NullableValidationCheckType) Set(val *ValidationCheckType) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationCheckType) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationCheckType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationCheckType(val *ValidationCheckType) *NullableValidationCheckType {
	return &NullableValidationCheckType{value: val, isSet: true}
}

func (v NullableValidationCheckType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationCheckType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
