/*
AVACloud API 1.51.0

AVACloud API specification

API version: 1.51.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
)

// checks if the ItemNumberSchemaTierDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemNumberSchemaTierDto{}

// ItemNumberSchemaTierDto Represents information about a single Tier in an ItemNumber.
type ItemNumberSchemaTierDto struct {
	// The (maximum) length for this tier. Will not accept a length less than 1. Defaults to 1 if length less than one is specified.
	Length   int32                       `json:"length"`
	Type     ItemNumberTypeDto           `json:"type"`
	TierType ItemNumberSchemaTierTypeDto `json:"tierType"`
	// Indicates if this tier represents a lot. See the documentation for more information about lots.
	IsLot bool `json:"isLot"`
	// This value is the increment, or step size, that should be used for new item numbers. It defaults to DEFAULT_INCREMENT, but can be changed to any other positive number greater than zero. Invalid values make this be set to one '1'
	Increment int32 `json:"increment"`
	// This is an optional name for the given tier
	TierName *string `json:"tierName,omitempty"`
}

// NewItemNumberSchemaTierDto instantiates a new ItemNumberSchemaTierDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemNumberSchemaTierDto(length int32, type_ ItemNumberTypeDto, tierType ItemNumberSchemaTierTypeDto, isLot bool, increment int32) *ItemNumberSchemaTierDto {
	this := ItemNumberSchemaTierDto{}
	this.Length = length
	this.Type = type_
	this.TierType = tierType
	this.IsLot = isLot
	this.Increment = increment
	return &this
}

// NewItemNumberSchemaTierDtoWithDefaults instantiates a new ItemNumberSchemaTierDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemNumberSchemaTierDtoWithDefaults() *ItemNumberSchemaTierDto {
	this := ItemNumberSchemaTierDto{}
	return &this
}

// GetLength returns the Length field value
func (o *ItemNumberSchemaTierDto) GetLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *ItemNumberSchemaTierDto) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *ItemNumberSchemaTierDto) SetLength(v int32) {
	o.Length = v
}

// GetType returns the Type field value
func (o *ItemNumberSchemaTierDto) GetType() ItemNumberTypeDto {
	if o == nil {
		var ret ItemNumberTypeDto
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ItemNumberSchemaTierDto) GetTypeOk() (*ItemNumberTypeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ItemNumberSchemaTierDto) SetType(v ItemNumberTypeDto) {
	o.Type = v
}

// GetTierType returns the TierType field value
func (o *ItemNumberSchemaTierDto) GetTierType() ItemNumberSchemaTierTypeDto {
	if o == nil {
		var ret ItemNumberSchemaTierTypeDto
		return ret
	}

	return o.TierType
}

// GetTierTypeOk returns a tuple with the TierType field value
// and a boolean to check if the value has been set.
func (o *ItemNumberSchemaTierDto) GetTierTypeOk() (*ItemNumberSchemaTierTypeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TierType, true
}

// SetTierType sets field value
func (o *ItemNumberSchemaTierDto) SetTierType(v ItemNumberSchemaTierTypeDto) {
	o.TierType = v
}

// GetIsLot returns the IsLot field value
func (o *ItemNumberSchemaTierDto) GetIsLot() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLot
}

// GetIsLotOk returns a tuple with the IsLot field value
// and a boolean to check if the value has been set.
func (o *ItemNumberSchemaTierDto) GetIsLotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLot, true
}

// SetIsLot sets field value
func (o *ItemNumberSchemaTierDto) SetIsLot(v bool) {
	o.IsLot = v
}

// GetIncrement returns the Increment field value
func (o *ItemNumberSchemaTierDto) GetIncrement() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value
// and a boolean to check if the value has been set.
func (o *ItemNumberSchemaTierDto) GetIncrementOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Increment, true
}

// SetIncrement sets field value
func (o *ItemNumberSchemaTierDto) SetIncrement(v int32) {
	o.Increment = v
}

// GetTierName returns the TierName field value if set, zero value otherwise.
func (o *ItemNumberSchemaTierDto) GetTierName() string {
	if o == nil || IsNil(o.TierName) {
		var ret string
		return ret
	}
	return *o.TierName
}

// GetTierNameOk returns a tuple with the TierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemNumberSchemaTierDto) GetTierNameOk() (*string, bool) {
	if o == nil || IsNil(o.TierName) {
		return nil, false
	}
	return o.TierName, true
}

// HasTierName returns a boolean if a field has been set.
func (o *ItemNumberSchemaTierDto) HasTierName() bool {
	if o != nil && !IsNil(o.TierName) {
		return true
	}

	return false
}

// SetTierName gets a reference to the given string and assigns it to the TierName field.
func (o *ItemNumberSchemaTierDto) SetTierName(v string) {
	o.TierName = &v
}

func (o ItemNumberSchemaTierDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemNumberSchemaTierDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["length"] = o.Length
	toSerialize["type"] = o.Type
	toSerialize["tierType"] = o.TierType
	// skip: isLot is readOnly
	toSerialize["increment"] = o.Increment
	if !IsNil(o.TierName) {
		toSerialize["tierName"] = o.TierName
	}
	return toSerialize, nil
}

type NullableItemNumberSchemaTierDto struct {
	value *ItemNumberSchemaTierDto
	isSet bool
}

func (v NullableItemNumberSchemaTierDto) Get() *ItemNumberSchemaTierDto {
	return v.value
}

func (v *NullableItemNumberSchemaTierDto) Set(val *ItemNumberSchemaTierDto) {
	v.value = val
	v.isSet = true
}

func (v NullableItemNumberSchemaTierDto) IsSet() bool {
	return v.isSet
}

func (v *NullableItemNumberSchemaTierDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemNumberSchemaTierDto(val *ItemNumberSchemaTierDto) *NullableItemNumberSchemaTierDto {
	return &NullableItemNumberSchemaTierDto{value: val, isSet: true}
}

func (v NullableItemNumberSchemaTierDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemNumberSchemaTierDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
