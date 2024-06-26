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

// ItemNumberTypeDto Determines the type of an ItemNumberSchema
type ItemNumberTypeDto string

// List of ItemNumberTypeDto
const (
	ITEMNUMBERTYPEDTO_NUMERIC      ItemNumberTypeDto = "Numeric"
	ITEMNUMBERTYPEDTO_ALPHANUMERIC ItemNumberTypeDto = "Alphanumeric"
)

// All allowed values of ItemNumberTypeDto enum
var AllowedItemNumberTypeDtoEnumValues = []ItemNumberTypeDto{
	"Numeric",
	"Alphanumeric",
}

func (v *ItemNumberTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ItemNumberTypeDto(value)
	for _, existing := range AllowedItemNumberTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ItemNumberTypeDto", value)
}

// NewItemNumberTypeDtoFromValue returns a pointer to a valid ItemNumberTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewItemNumberTypeDtoFromValue(v string) (*ItemNumberTypeDto, error) {
	ev := ItemNumberTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ItemNumberTypeDto: valid values are %v", v, AllowedItemNumberTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ItemNumberTypeDto) IsValid() bool {
	for _, existing := range AllowedItemNumberTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ItemNumberTypeDto value
func (v ItemNumberTypeDto) Ptr() *ItemNumberTypeDto {
	return &v
}

type NullableItemNumberTypeDto struct {
	value *ItemNumberTypeDto
	isSet bool
}

func (v NullableItemNumberTypeDto) Get() *ItemNumberTypeDto {
	return v.value
}

func (v *NullableItemNumberTypeDto) Set(val *ItemNumberTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableItemNumberTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableItemNumberTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemNumberTypeDto(val *ItemNumberTypeDto) *NullableItemNumberTypeDto {
	return &NullableItemNumberTypeDto{value: val, isSet: true}
}

func (v NullableItemNumberTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemNumberTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
