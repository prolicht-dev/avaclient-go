/*
AVACloud API 1.45.0

AVACloud API specification

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
	"fmt"
)

// PriceComponentTypeDto This enumeration models types of price components for better classification.
type PriceComponentTypeDto string

// List of PriceComponentTypeDto
const (
	PRICECOMPONENTTYPEDTO_UNKNOWN       PriceComponentTypeDto = "Unknown"
	PRICECOMPONENTTYPEDTO_WAGES         PriceComponentTypeDto = "Wages"
	PRICECOMPONENTTYPEDTO_MATERIALS     PriceComponentTypeDto = "Materials"
	PRICECOMPONENTTYPEDTO_PLANT         PriceComponentTypeDto = "Plant"
	PRICECOMPONENTTYPEDTO_MISCELLANEOUS PriceComponentTypeDto = "Miscellaneous"
)

// All allowed values of PriceComponentTypeDto enum
var AllowedPriceComponentTypeDtoEnumValues = []PriceComponentTypeDto{
	"Unknown",
	"Wages",
	"Materials",
	"Plant",
	"Miscellaneous",
}

func (v *PriceComponentTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriceComponentTypeDto(value)
	for _, existing := range AllowedPriceComponentTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriceComponentTypeDto", value)
}

// NewPriceComponentTypeDtoFromValue returns a pointer to a valid PriceComponentTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriceComponentTypeDtoFromValue(v string) (*PriceComponentTypeDto, error) {
	ev := PriceComponentTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriceComponentTypeDto: valid values are %v", v, AllowedPriceComponentTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriceComponentTypeDto) IsValid() bool {
	for _, existing := range AllowedPriceComponentTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PriceComponentTypeDto value
func (v PriceComponentTypeDto) Ptr() *PriceComponentTypeDto {
	return &v
}

type NullablePriceComponentTypeDto struct {
	value *PriceComponentTypeDto
	isSet bool
}

func (v NullablePriceComponentTypeDto) Get() *PriceComponentTypeDto {
	return v.value
}

func (v *NullablePriceComponentTypeDto) Set(val *PriceComponentTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceComponentTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceComponentTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceComponentTypeDto(val *PriceComponentTypeDto) *NullablePriceComponentTypeDto {
	return &NullablePriceComponentTypeDto{value: val, isSet: true}
}

func (v NullablePriceComponentTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceComponentTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
