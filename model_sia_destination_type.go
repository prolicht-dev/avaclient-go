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

// SiaDestinationType Represents valid SIA destination types
type SiaDestinationType string

// List of SiaDestinationType
const (
	SIADESTINATIONTYPE_SIA451 SiaDestinationType = "Sia451"
	SIADESTINATIONTYPE_IFA18  SiaDestinationType = "Ifa18"
)

// All allowed values of SiaDestinationType enum
var AllowedSiaDestinationTypeEnumValues = []SiaDestinationType{
	"Sia451",
	"Ifa18",
}

func (v *SiaDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SiaDestinationType(value)
	for _, existing := range AllowedSiaDestinationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SiaDestinationType", value)
}

// NewSiaDestinationTypeFromValue returns a pointer to a valid SiaDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSiaDestinationTypeFromValue(v string) (*SiaDestinationType, error) {
	ev := SiaDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SiaDestinationType: valid values are %v", v, AllowedSiaDestinationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SiaDestinationType) IsValid() bool {
	for _, existing := range AllowedSiaDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SiaDestinationType value
func (v SiaDestinationType) Ptr() *SiaDestinationType {
	return &v
}

type NullableSiaDestinationType struct {
	value *SiaDestinationType
	isSet bool
}

func (v NullableSiaDestinationType) Get() *SiaDestinationType {
	return v.value
}

func (v *NullableSiaDestinationType) Set(val *SiaDestinationType) {
	v.value = val
	v.isSet = true
}

func (v NullableSiaDestinationType) IsSet() bool {
	return v.isSet
}

func (v *NullableSiaDestinationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiaDestinationType(val *SiaDestinationType) *NullableSiaDestinationType {
	return &NullableSiaDestinationType{value: val, isSet: true}
}

func (v NullableSiaDestinationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiaDestinationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
