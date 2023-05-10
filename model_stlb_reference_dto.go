/*
AVACloud API 1.41.4

AVACloud API specification

API version: 1.41.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
	"time"
)

// checks if the STLBReferenceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &STLBReferenceDto{}

// STLBReferenceDto This class represents a specialized reference to the German STLB \"Standardleistungsbuch Bau\". The STLB is a commercial offering and describes common services in the construction sector. When this is used, this describes the exact type of a service via a reference to this standard
type STLBReferenceDto struct {
	// The date of the STLB version. Typically, only the Year and Month are used
	VersionDate *time.Time `json:"versionDate,omitempty"`
	// The name of the catalogue within the STLB
	CatalogueName *string `json:"catalogueName,omitempty"`
	// The name of the group in STLB
	Group *string `json:"group,omitempty"`
	// The cost group this service is associated with
	CostGroup *string `json:"costGroup,omitempty"`
	// The service area (or type) in the STLB
	ServiceArea *string `json:"serviceArea,omitempty"`
	// These keys may optionally be used to further reference multiple, specific items within the STLB
	Keys []STLBKeyDto `json:"keys,omitempty"`
}

// NewSTLBReferenceDto instantiates a new STLBReferenceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSTLBReferenceDto() *STLBReferenceDto {
	this := STLBReferenceDto{}
	return &this
}

// NewSTLBReferenceDtoWithDefaults instantiates a new STLBReferenceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSTLBReferenceDtoWithDefaults() *STLBReferenceDto {
	this := STLBReferenceDto{}
	return &this
}

// GetVersionDate returns the VersionDate field value if set, zero value otherwise.
func (o *STLBReferenceDto) GetVersionDate() time.Time {
	if o == nil || IsNil(o.VersionDate) {
		var ret time.Time
		return ret
	}
	return *o.VersionDate
}

// GetVersionDateOk returns a tuple with the VersionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STLBReferenceDto) GetVersionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VersionDate) {
		return nil, false
	}
	return o.VersionDate, true
}

// HasVersionDate returns a boolean if a field has been set.
func (o *STLBReferenceDto) HasVersionDate() bool {
	if o != nil && !IsNil(o.VersionDate) {
		return true
	}

	return false
}

// SetVersionDate gets a reference to the given time.Time and assigns it to the VersionDate field.
func (o *STLBReferenceDto) SetVersionDate(v time.Time) {
	o.VersionDate = &v
}

// GetCatalogueName returns the CatalogueName field value if set, zero value otherwise.
func (o *STLBReferenceDto) GetCatalogueName() string {
	if o == nil || IsNil(o.CatalogueName) {
		var ret string
		return ret
	}
	return *o.CatalogueName
}

// GetCatalogueNameOk returns a tuple with the CatalogueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STLBReferenceDto) GetCatalogueNameOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogueName) {
		return nil, false
	}
	return o.CatalogueName, true
}

// HasCatalogueName returns a boolean if a field has been set.
func (o *STLBReferenceDto) HasCatalogueName() bool {
	if o != nil && !IsNil(o.CatalogueName) {
		return true
	}

	return false
}

// SetCatalogueName gets a reference to the given string and assigns it to the CatalogueName field.
func (o *STLBReferenceDto) SetCatalogueName(v string) {
	o.CatalogueName = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *STLBReferenceDto) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STLBReferenceDto) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *STLBReferenceDto) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *STLBReferenceDto) SetGroup(v string) {
	o.Group = &v
}

// GetCostGroup returns the CostGroup field value if set, zero value otherwise.
func (o *STLBReferenceDto) GetCostGroup() string {
	if o == nil || IsNil(o.CostGroup) {
		var ret string
		return ret
	}
	return *o.CostGroup
}

// GetCostGroupOk returns a tuple with the CostGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STLBReferenceDto) GetCostGroupOk() (*string, bool) {
	if o == nil || IsNil(o.CostGroup) {
		return nil, false
	}
	return o.CostGroup, true
}

// HasCostGroup returns a boolean if a field has been set.
func (o *STLBReferenceDto) HasCostGroup() bool {
	if o != nil && !IsNil(o.CostGroup) {
		return true
	}

	return false
}

// SetCostGroup gets a reference to the given string and assigns it to the CostGroup field.
func (o *STLBReferenceDto) SetCostGroup(v string) {
	o.CostGroup = &v
}

// GetServiceArea returns the ServiceArea field value if set, zero value otherwise.
func (o *STLBReferenceDto) GetServiceArea() string {
	if o == nil || IsNil(o.ServiceArea) {
		var ret string
		return ret
	}
	return *o.ServiceArea
}

// GetServiceAreaOk returns a tuple with the ServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STLBReferenceDto) GetServiceAreaOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceArea) {
		return nil, false
	}
	return o.ServiceArea, true
}

// HasServiceArea returns a boolean if a field has been set.
func (o *STLBReferenceDto) HasServiceArea() bool {
	if o != nil && !IsNil(o.ServiceArea) {
		return true
	}

	return false
}

// SetServiceArea gets a reference to the given string and assigns it to the ServiceArea field.
func (o *STLBReferenceDto) SetServiceArea(v string) {
	o.ServiceArea = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *STLBReferenceDto) GetKeys() []STLBKeyDto {
	if o == nil || IsNil(o.Keys) {
		var ret []STLBKeyDto
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STLBReferenceDto) GetKeysOk() ([]STLBKeyDto, bool) {
	if o == nil || IsNil(o.Keys) {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *STLBReferenceDto) HasKeys() bool {
	if o != nil && !IsNil(o.Keys) {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []STLBKeyDto and assigns it to the Keys field.
func (o *STLBReferenceDto) SetKeys(v []STLBKeyDto) {
	o.Keys = v
}

func (o STLBReferenceDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o STLBReferenceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VersionDate) {
		toSerialize["versionDate"] = o.VersionDate
	}
	if !IsNil(o.CatalogueName) {
		toSerialize["catalogueName"] = o.CatalogueName
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.CostGroup) {
		toSerialize["costGroup"] = o.CostGroup
	}
	if !IsNil(o.ServiceArea) {
		toSerialize["serviceArea"] = o.ServiceArea
	}
	if !IsNil(o.Keys) {
		toSerialize["keys"] = o.Keys
	}
	return toSerialize, nil
}

type NullableSTLBReferenceDto struct {
	value *STLBReferenceDto
	isSet bool
}

func (v NullableSTLBReferenceDto) Get() *STLBReferenceDto {
	return v.value
}

func (v *NullableSTLBReferenceDto) Set(val *STLBReferenceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSTLBReferenceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSTLBReferenceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSTLBReferenceDto(val *STLBReferenceDto) *NullableSTLBReferenceDto {
	return &NullableSTLBReferenceDto{value: val, isSet: true}
}

func (v NullableSTLBReferenceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSTLBReferenceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
