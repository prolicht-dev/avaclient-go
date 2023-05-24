/*
AVACloud API 1.41.8

AVACloud API specification

API version: 1.41.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
	"fmt"
)

// ServiceTypeDto The service type describes the type of service a Position represents.
type ServiceTypeDto string

// List of ServiceTypeDto
const (
	SERVICETYPEDTO_REGULAR          ServiceTypeDto = "Regular"
	SERVICETYPEDTO_HOURLY_PAID_WORK ServiceTypeDto = "HourlyPaidWork"
)

// All allowed values of ServiceTypeDto enum
var AllowedServiceTypeDtoEnumValues = []ServiceTypeDto{
	"Regular",
	"HourlyPaidWork",
}

func (v *ServiceTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceTypeDto(value)
	for _, existing := range AllowedServiceTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceTypeDto", value)
}

// NewServiceTypeDtoFromValue returns a pointer to a valid ServiceTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceTypeDtoFromValue(v string) (*ServiceTypeDto, error) {
	ev := ServiceTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceTypeDto: valid values are %v", v, AllowedServiceTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceTypeDto) IsValid() bool {
	for _, existing := range AllowedServiceTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceTypeDto value
func (v ServiceTypeDto) Ptr() *ServiceTypeDto {
	return &v
}

type NullableServiceTypeDto struct {
	value *ServiceTypeDto
	isSet bool
}

func (v NullableServiceTypeDto) Get() *ServiceTypeDto {
	return v.value
}

func (v *NullableServiceTypeDto) Set(val *ServiceTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTypeDto(val *ServiceTypeDto) *NullableServiceTypeDto {
	return &NullableServiceTypeDto{value: val, isSet: true}
}

func (v NullableServiceTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
