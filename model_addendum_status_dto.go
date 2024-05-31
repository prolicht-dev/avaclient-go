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

// AddendumStatusDto This enumeration represents the status of an addendum, 'Nachtrag' in German
type AddendumStatusDto string

// List of AddendumStatusDto
const (
	ADDENDUMSTATUSDTO_UNKNOWN               AddendumStatusDto = "Unknown"
	ADDENDUMSTATUSDTO_RECOGNIZED            AddendumStatusDto = "Recognized"
	ADDENDUMSTATUSDTO_FILED                 AddendumStatusDto = "Filed"
	ADDENDUMSTATUSDTO_OFFERED               AddendumStatusDto = "Offered"
	ADDENDUMSTATUSDTO_WITHDRAWN             AddendumStatusDto = "Withdrawn"
	ADDENDUMSTATUSDTO_REJECTED              AddendumStatusDto = "Rejected"
	ADDENDUMSTATUSDTO_REJECTION_OBJECTED    AddendumStatusDto = "RejectionObjected"
	ADDENDUMSTATUSDTO_FORMALLY_ACKNOWLEDGED AddendumStatusDto = "FormallyAcknowledged"
	ADDENDUMSTATUSDTO_APPROVED              AddendumStatusDto = "Approved"
)

// All allowed values of AddendumStatusDto enum
var AllowedAddendumStatusDtoEnumValues = []AddendumStatusDto{
	"Unknown",
	"Recognized",
	"Filed",
	"Offered",
	"Withdrawn",
	"Rejected",
	"RejectionObjected",
	"FormallyAcknowledged",
	"Approved",
}

func (v *AddendumStatusDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AddendumStatusDto(value)
	for _, existing := range AllowedAddendumStatusDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AddendumStatusDto", value)
}

// NewAddendumStatusDtoFromValue returns a pointer to a valid AddendumStatusDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAddendumStatusDtoFromValue(v string) (*AddendumStatusDto, error) {
	ev := AddendumStatusDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AddendumStatusDto: valid values are %v", v, AllowedAddendumStatusDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AddendumStatusDto) IsValid() bool {
	for _, existing := range AllowedAddendumStatusDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AddendumStatusDto value
func (v AddendumStatusDto) Ptr() *AddendumStatusDto {
	return &v
}

type NullableAddendumStatusDto struct {
	value *AddendumStatusDto
	isSet bool
}

func (v NullableAddendumStatusDto) Get() *AddendumStatusDto {
	return v.value
}

func (v *NullableAddendumStatusDto) Set(val *AddendumStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAddendumStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAddendumStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddendumStatusDto(val *AddendumStatusDto) *NullableAddendumStatusDto {
	return &NullableAddendumStatusDto{value: val, isSet: true}
}

func (v NullableAddendumStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddendumStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
