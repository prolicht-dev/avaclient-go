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

// UglDestinationVersion Represents valid UGL versions for export
type UglDestinationVersion string

// List of UglDestinationVersion
const (
	UGLDESTINATIONVERSION_V4 UglDestinationVersion = "V4"
	UGLDESTINATIONVERSION_V5 UglDestinationVersion = "V5"
)

// All allowed values of UglDestinationVersion enum
var AllowedUglDestinationVersionEnumValues = []UglDestinationVersion{
	"V4",
	"V5",
}

func (v *UglDestinationVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UglDestinationVersion(value)
	for _, existing := range AllowedUglDestinationVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UglDestinationVersion", value)
}

// NewUglDestinationVersionFromValue returns a pointer to a valid UglDestinationVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUglDestinationVersionFromValue(v string) (*UglDestinationVersion, error) {
	ev := UglDestinationVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UglDestinationVersion: valid values are %v", v, AllowedUglDestinationVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UglDestinationVersion) IsValid() bool {
	for _, existing := range AllowedUglDestinationVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UglDestinationVersion value
func (v UglDestinationVersion) Ptr() *UglDestinationVersion {
	return &v
}

type NullableUglDestinationVersion struct {
	value *UglDestinationVersion
	isSet bool
}

func (v NullableUglDestinationVersion) Get() *UglDestinationVersion {
	return v.value
}

func (v *NullableUglDestinationVersion) Set(val *UglDestinationVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableUglDestinationVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableUglDestinationVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUglDestinationVersion(val *UglDestinationVersion) *NullableUglDestinationVersion {
	return &NullableUglDestinationVersion{value: val, isSet: true}
}

func (v NullableUglDestinationVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUglDestinationVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
