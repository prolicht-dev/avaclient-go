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

// PriceTypeDto This indicates the price type of a IPricedElement.
type PriceTypeDto string

// List of PriceTypeDto
const (
	PRICETYPEDTO_WITH_TOTAL    PriceTypeDto = "WithTotal"
	PRICETYPEDTO_WITHOUT_TOTAL PriceTypeDto = "WithoutTotal"
)

// All allowed values of PriceTypeDto enum
var AllowedPriceTypeDtoEnumValues = []PriceTypeDto{
	"WithTotal",
	"WithoutTotal",
}

func (v *PriceTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriceTypeDto(value)
	for _, existing := range AllowedPriceTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriceTypeDto", value)
}

// NewPriceTypeDtoFromValue returns a pointer to a valid PriceTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriceTypeDtoFromValue(v string) (*PriceTypeDto, error) {
	ev := PriceTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriceTypeDto: valid values are %v", v, AllowedPriceTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriceTypeDto) IsValid() bool {
	for _, existing := range AllowedPriceTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PriceTypeDto value
func (v PriceTypeDto) Ptr() *PriceTypeDto {
	return &v
}

type NullablePriceTypeDto struct {
	value *PriceTypeDto
	isSet bool
}

func (v NullablePriceTypeDto) Get() *PriceTypeDto {
	return v.value
}

func (v *NullablePriceTypeDto) Set(val *PriceTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTypeDto(val *PriceTypeDto) *NullablePriceTypeDto {
	return &NullablePriceTypeDto{value: val, isSet: true}
}

func (v NullablePriceTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
