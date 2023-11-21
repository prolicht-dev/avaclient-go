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

// ExchangePhaseDto Indicates a ServiceSpecification's exchange phase, based on the GAEB exchange phases.
type ExchangePhaseDto string

// List of ExchangePhaseDto
const (
	EXCHANGEPHASEDTO_UNDEFINED     ExchangePhaseDto = "Undefined"
	EXCHANGEPHASEDTO_BASE          ExchangePhaseDto = "Base"
	EXCHANGEPHASEDTO_COST_ESTIMATE ExchangePhaseDto = "CostEstimate"
	EXCHANGEPHASEDTO_OFFER_REQUEST ExchangePhaseDto = "OfferRequest"
	EXCHANGEPHASEDTO_OFFER         ExchangePhaseDto = "Offer"
	EXCHANGEPHASEDTO_SIDE_OFFER    ExchangePhaseDto = "SideOffer"
	EXCHANGEPHASEDTO_GRANT         ExchangePhaseDto = "Grant"
)

// All allowed values of ExchangePhaseDto enum
var AllowedExchangePhaseDtoEnumValues = []ExchangePhaseDto{
	"Undefined",
	"Base",
	"CostEstimate",
	"OfferRequest",
	"Offer",
	"SideOffer",
	"Grant",
}

func (v *ExchangePhaseDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExchangePhaseDto(value)
	for _, existing := range AllowedExchangePhaseDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExchangePhaseDto", value)
}

// NewExchangePhaseDtoFromValue returns a pointer to a valid ExchangePhaseDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExchangePhaseDtoFromValue(v string) (*ExchangePhaseDto, error) {
	ev := ExchangePhaseDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExchangePhaseDto: valid values are %v", v, AllowedExchangePhaseDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExchangePhaseDto) IsValid() bool {
	for _, existing := range AllowedExchangePhaseDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExchangePhaseDto value
func (v ExchangePhaseDto) Ptr() *ExchangePhaseDto {
	return &v
}

type NullableExchangePhaseDto struct {
	value *ExchangePhaseDto
	isSet bool
}

func (v NullableExchangePhaseDto) Get() *ExchangePhaseDto {
	return v.value
}

func (v *NullableExchangePhaseDto) Set(val *ExchangePhaseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangePhaseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangePhaseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangePhaseDto(val *ExchangePhaseDto) *NullableExchangePhaseDto {
	return &NullableExchangePhaseDto{value: val, isSet: true}
}

func (v NullableExchangePhaseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangePhaseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
