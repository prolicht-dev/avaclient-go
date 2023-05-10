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

// CommerceInquiryTypeDto This enum represents types of price inquiries
type CommerceInquiryTypeDto string

// List of CommerceInquiryTypeDto
const (
	COMMERCEINQUIRYTYPEDTO_UNKNOWN    CommerceInquiryTypeDto = "Unknown"
	COMMERCEINQUIRYTYPEDTO_PROJECT    CommerceInquiryTypeDto = "Project"
	COMMERCEINQUIRYTYPEDTO_IMMEDIATE  CommerceInquiryTypeDto = "Immediate"
	COMMERCEINQUIRYTYPEDTO_EXHIBITION CommerceInquiryTypeDto = "Exhibition"
)

// All allowed values of CommerceInquiryTypeDto enum
var AllowedCommerceInquiryTypeDtoEnumValues = []CommerceInquiryTypeDto{
	"Unknown",
	"Project",
	"Immediate",
	"Exhibition",
}

func (v *CommerceInquiryTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommerceInquiryTypeDto(value)
	for _, existing := range AllowedCommerceInquiryTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommerceInquiryTypeDto", value)
}

// NewCommerceInquiryTypeDtoFromValue returns a pointer to a valid CommerceInquiryTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommerceInquiryTypeDtoFromValue(v string) (*CommerceInquiryTypeDto, error) {
	ev := CommerceInquiryTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommerceInquiryTypeDto: valid values are %v", v, AllowedCommerceInquiryTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommerceInquiryTypeDto) IsValid() bool {
	for _, existing := range AllowedCommerceInquiryTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CommerceInquiryTypeDto value
func (v CommerceInquiryTypeDto) Ptr() *CommerceInquiryTypeDto {
	return &v
}

type NullableCommerceInquiryTypeDto struct {
	value *CommerceInquiryTypeDto
	isSet bool
}

func (v NullableCommerceInquiryTypeDto) Get() *CommerceInquiryTypeDto {
	return v.value
}

func (v *NullableCommerceInquiryTypeDto) Set(val *CommerceInquiryTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCommerceInquiryTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCommerceInquiryTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommerceInquiryTypeDto(val *CommerceInquiryTypeDto) *NullableCommerceInquiryTypeDto {
	return &NullableCommerceInquiryTypeDto{value: val, isSet: true}
}

func (v NullableCommerceInquiryTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommerceInquiryTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}