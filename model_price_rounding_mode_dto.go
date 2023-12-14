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

// PriceRoundingModeDto This enum configures rounding modes for price calculations in projects
type PriceRoundingModeDto string

// List of PriceRoundingModeDto
const (
	PRICEROUNDINGMODEDTO_NORMAL  PriceRoundingModeDto = "Normal"
	PRICEROUNDINGMODEDTO_FLOOR   PriceRoundingModeDto = "Floor"
	PRICEROUNDINGMODEDTO_CEILING PriceRoundingModeDto = "Ceiling"
)

// All allowed values of PriceRoundingModeDto enum
var AllowedPriceRoundingModeDtoEnumValues = []PriceRoundingModeDto{
	"Normal",
	"Floor",
	"Ceiling",
}

func (v *PriceRoundingModeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriceRoundingModeDto(value)
	for _, existing := range AllowedPriceRoundingModeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriceRoundingModeDto", value)
}

// NewPriceRoundingModeDtoFromValue returns a pointer to a valid PriceRoundingModeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriceRoundingModeDtoFromValue(v string) (*PriceRoundingModeDto, error) {
	ev := PriceRoundingModeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriceRoundingModeDto: valid values are %v", v, AllowedPriceRoundingModeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriceRoundingModeDto) IsValid() bool {
	for _, existing := range AllowedPriceRoundingModeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PriceRoundingModeDto value
func (v PriceRoundingModeDto) Ptr() *PriceRoundingModeDto {
	return &v
}

type NullablePriceRoundingModeDto struct {
	value *PriceRoundingModeDto
	isSet bool
}

func (v NullablePriceRoundingModeDto) Get() *PriceRoundingModeDto {
	return v.value
}

func (v *NullablePriceRoundingModeDto) Set(val *PriceRoundingModeDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceRoundingModeDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceRoundingModeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceRoundingModeDto(val *PriceRoundingModeDto) *NullablePriceRoundingModeDto {
	return &NullablePriceRoundingModeDto{value: val, isSet: true}
}

func (v NullablePriceRoundingModeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceRoundingModeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
