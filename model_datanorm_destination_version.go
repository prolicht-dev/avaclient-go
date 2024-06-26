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

// DatanormDestinationVersion Represents valid Datanorm versions for export
type DatanormDestinationVersion string

// List of DatanormDestinationVersion
const (
	DATANORMDESTINATIONVERSION_V4 DatanormDestinationVersion = "V4"
	DATANORMDESTINATIONVERSION_V5 DatanormDestinationVersion = "V5"
)

// All allowed values of DatanormDestinationVersion enum
var AllowedDatanormDestinationVersionEnumValues = []DatanormDestinationVersion{
	"V4",
	"V5",
}

func (v *DatanormDestinationVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatanormDestinationVersion(value)
	for _, existing := range AllowedDatanormDestinationVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatanormDestinationVersion", value)
}

// NewDatanormDestinationVersionFromValue returns a pointer to a valid DatanormDestinationVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatanormDestinationVersionFromValue(v string) (*DatanormDestinationVersion, error) {
	ev := DatanormDestinationVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatanormDestinationVersion: valid values are %v", v, AllowedDatanormDestinationVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatanormDestinationVersion) IsValid() bool {
	for _, existing := range AllowedDatanormDestinationVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatanormDestinationVersion value
func (v DatanormDestinationVersion) Ptr() *DatanormDestinationVersion {
	return &v
}

type NullableDatanormDestinationVersion struct {
	value *DatanormDestinationVersion
	isSet bool
}

func (v NullableDatanormDestinationVersion) Get() *DatanormDestinationVersion {
	return v.value
}

func (v *NullableDatanormDestinationVersion) Set(val *DatanormDestinationVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableDatanormDestinationVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableDatanormDestinationVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatanormDestinationVersion(val *DatanormDestinationVersion) *NullableDatanormDestinationVersion {
	return &NullableDatanormDestinationVersion{value: val, isSet: true}
}

func (v NullableDatanormDestinationVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatanormDestinationVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
