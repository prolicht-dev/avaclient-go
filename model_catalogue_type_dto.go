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

// CatalogueTypeDto This enumeration describes the type of a catalogue. Catalogues, or collections, hold information to categorize and describe items. For example, the German DIN 276 cost group standards describe different types of costs for building projects. When referencing the DIN 276 catalogue and providing an item key or identifier, it is possible to reference data in this catalogue.
type CatalogueTypeDto string

// List of CatalogueTypeDto
const (
	CATALOGUETYPEDTO_UNKNOWN                 CatalogueTypeDto = "Unknown"
	CATALOGUETYPEDTO_LOCATION                CatalogueTypeDto = "Location"
	CATALOGUETYPEDTO_DIN276                  CatalogueTypeDto = "DIN276"
	CATALOGUETYPEDTO_COST_UNIT               CatalogueTypeDto = "CostUnit"
	CATALOGUETYPEDTO_WORK_CATEGORY           CatalogueTypeDto = "WorkCategory"
	CATALOGUETYPEDTO_OENORM_B1801_COST_GROUP CatalogueTypeDto = "OenormB1801CostGroup"
	CATALOGUETYPEDTO_BIM                     CatalogueTypeDto = "BIM"
	CATALOGUETYPEDTO_ATTACHMENT              CatalogueTypeDto = "Attachment"
)

// All allowed values of CatalogueTypeDto enum
var AllowedCatalogueTypeDtoEnumValues = []CatalogueTypeDto{
	"Unknown",
	"Location",
	"DIN276",
	"CostUnit",
	"WorkCategory",
	"OenormB1801CostGroup",
	"BIM",
	"Attachment",
}

func (v *CatalogueTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CatalogueTypeDto(value)
	for _, existing := range AllowedCatalogueTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CatalogueTypeDto", value)
}

// NewCatalogueTypeDtoFromValue returns a pointer to a valid CatalogueTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCatalogueTypeDtoFromValue(v string) (*CatalogueTypeDto, error) {
	ev := CatalogueTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CatalogueTypeDto: valid values are %v", v, AllowedCatalogueTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CatalogueTypeDto) IsValid() bool {
	for _, existing := range AllowedCatalogueTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CatalogueTypeDto value
func (v CatalogueTypeDto) Ptr() *CatalogueTypeDto {
	return &v
}

type NullableCatalogueTypeDto struct {
	value *CatalogueTypeDto
	isSet bool
}

func (v NullableCatalogueTypeDto) Get() *CatalogueTypeDto {
	return v.value
}

func (v *NullableCatalogueTypeDto) Set(val *CatalogueTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogueTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogueTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogueTypeDto(val *CatalogueTypeDto) *NullableCatalogueTypeDto {
	return &NullableCatalogueTypeDto{value: val, isSet: true}
}

func (v NullableCatalogueTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogueTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
