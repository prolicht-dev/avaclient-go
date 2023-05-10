/*
AVACloud API 1.41.4

AVACloud API specification

API version: 1.41.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
)

// checks if the OenormPositionPropertiesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OenormPositionPropertiesDto{}

// OenormPositionPropertiesDto This class models special properties that only apply to some exchange scenarios where ÖNorm is used. It is special for Positions and extends the OenormProperties base class.
type OenormPositionPropertiesDto struct {
	OriginCode OenormOriginCodeDto `json:"originCode"`
	// This marks if the opening texts within this element are considered free text. It corresponds to 'vorbemerkungskennzeichen' in ÖNorm.
	OpeningTextIsFreeText bool `json:"openingTextIsFreeText"`
	// This indicates if the ÖNorm 'wesentliche position' mark is set
	IsMainPosition bool `json:"isMainPosition"`
	// This indicates if the ÖNorm position was a 'ungeteilteposition' (undivided position). This will only be taken into account when the position is also the sole element inside it's parent group
	IsUndividedPosition bool `json:"isUndividedPosition"`
	// In some ÖNorm formats, the short text can have it's own addition, so the text is split up in OenormShortText and OenormShortTextAddition To serialize this, either the ShortText property of the parent position needs to be null, or OenormShortText ' ' OenormShortTextAddition needs to match the ShortText.
	OenormShortText *string `json:"oenormShortText,omitempty"`
	// In some ÖNorm formats, the short text can have it's own addition, so the text is split up in OenormShortText and OenormShortTextAddition To serialize this, either the ShortText property of the parent position needs to be null, or OenormShortText ' ' OenormShortTextAddition needs to match the ShortText.
	OenormShortTextAddition *string `json:"oenormShortTextAddition,omitempty"`
}

// NewOenormPositionPropertiesDto instantiates a new OenormPositionPropertiesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOenormPositionPropertiesDto(originCode OenormOriginCodeDto, openingTextIsFreeText bool, isMainPosition bool, isUndividedPosition bool) *OenormPositionPropertiesDto {
	this := OenormPositionPropertiesDto{}
	this.OriginCode = originCode
	this.OpeningTextIsFreeText = openingTextIsFreeText
	this.IsMainPosition = isMainPosition
	this.IsUndividedPosition = isUndividedPosition
	return &this
}

// NewOenormPositionPropertiesDtoWithDefaults instantiates a new OenormPositionPropertiesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOenormPositionPropertiesDtoWithDefaults() *OenormPositionPropertiesDto {
	this := OenormPositionPropertiesDto{}
	return &this
}

// GetOriginCode returns the OriginCode field value
func (o *OenormPositionPropertiesDto) GetOriginCode() OenormOriginCodeDto {
	if o == nil {
		var ret OenormOriginCodeDto
		return ret
	}

	return o.OriginCode
}

// GetOriginCodeOk returns a tuple with the OriginCode field value
// and a boolean to check if the value has been set.
func (o *OenormPositionPropertiesDto) GetOriginCodeOk() (*OenormOriginCodeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginCode, true
}

// SetOriginCode sets field value
func (o *OenormPositionPropertiesDto) SetOriginCode(v OenormOriginCodeDto) {
	o.OriginCode = v
}

// GetOpeningTextIsFreeText returns the OpeningTextIsFreeText field value
func (o *OenormPositionPropertiesDto) GetOpeningTextIsFreeText() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OpeningTextIsFreeText
}

// GetOpeningTextIsFreeTextOk returns a tuple with the OpeningTextIsFreeText field value
// and a boolean to check if the value has been set.
func (o *OenormPositionPropertiesDto) GetOpeningTextIsFreeTextOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpeningTextIsFreeText, true
}

// SetOpeningTextIsFreeText sets field value
func (o *OenormPositionPropertiesDto) SetOpeningTextIsFreeText(v bool) {
	o.OpeningTextIsFreeText = v
}

// GetIsMainPosition returns the IsMainPosition field value
func (o *OenormPositionPropertiesDto) GetIsMainPosition() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMainPosition
}

// GetIsMainPositionOk returns a tuple with the IsMainPosition field value
// and a boolean to check if the value has been set.
func (o *OenormPositionPropertiesDto) GetIsMainPositionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMainPosition, true
}

// SetIsMainPosition sets field value
func (o *OenormPositionPropertiesDto) SetIsMainPosition(v bool) {
	o.IsMainPosition = v
}

// GetIsUndividedPosition returns the IsUndividedPosition field value
func (o *OenormPositionPropertiesDto) GetIsUndividedPosition() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUndividedPosition
}

// GetIsUndividedPositionOk returns a tuple with the IsUndividedPosition field value
// and a boolean to check if the value has been set.
func (o *OenormPositionPropertiesDto) GetIsUndividedPositionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUndividedPosition, true
}

// SetIsUndividedPosition sets field value
func (o *OenormPositionPropertiesDto) SetIsUndividedPosition(v bool) {
	o.IsUndividedPosition = v
}

// GetOenormShortText returns the OenormShortText field value if set, zero value otherwise.
func (o *OenormPositionPropertiesDto) GetOenormShortText() string {
	if o == nil || IsNil(o.OenormShortText) {
		var ret string
		return ret
	}
	return *o.OenormShortText
}

// GetOenormShortTextOk returns a tuple with the OenormShortText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OenormPositionPropertiesDto) GetOenormShortTextOk() (*string, bool) {
	if o == nil || IsNil(o.OenormShortText) {
		return nil, false
	}
	return o.OenormShortText, true
}

// HasOenormShortText returns a boolean if a field has been set.
func (o *OenormPositionPropertiesDto) HasOenormShortText() bool {
	if o != nil && !IsNil(o.OenormShortText) {
		return true
	}

	return false
}

// SetOenormShortText gets a reference to the given string and assigns it to the OenormShortText field.
func (o *OenormPositionPropertiesDto) SetOenormShortText(v string) {
	o.OenormShortText = &v
}

// GetOenormShortTextAddition returns the OenormShortTextAddition field value if set, zero value otherwise.
func (o *OenormPositionPropertiesDto) GetOenormShortTextAddition() string {
	if o == nil || IsNil(o.OenormShortTextAddition) {
		var ret string
		return ret
	}
	return *o.OenormShortTextAddition
}

// GetOenormShortTextAdditionOk returns a tuple with the OenormShortTextAddition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OenormPositionPropertiesDto) GetOenormShortTextAdditionOk() (*string, bool) {
	if o == nil || IsNil(o.OenormShortTextAddition) {
		return nil, false
	}
	return o.OenormShortTextAddition, true
}

// HasOenormShortTextAddition returns a boolean if a field has been set.
func (o *OenormPositionPropertiesDto) HasOenormShortTextAddition() bool {
	if o != nil && !IsNil(o.OenormShortTextAddition) {
		return true
	}

	return false
}

// SetOenormShortTextAddition gets a reference to the given string and assigns it to the OenormShortTextAddition field.
func (o *OenormPositionPropertiesDto) SetOenormShortTextAddition(v string) {
	o.OenormShortTextAddition = &v
}

func (o OenormPositionPropertiesDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OenormPositionPropertiesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["originCode"] = o.OriginCode
	toSerialize["openingTextIsFreeText"] = o.OpeningTextIsFreeText
	toSerialize["isMainPosition"] = o.IsMainPosition
	toSerialize["isUndividedPosition"] = o.IsUndividedPosition
	if !IsNil(o.OenormShortText) {
		toSerialize["oenormShortText"] = o.OenormShortText
	}
	if !IsNil(o.OenormShortTextAddition) {
		toSerialize["oenormShortTextAddition"] = o.OenormShortTextAddition
	}
	return toSerialize, nil
}

type NullableOenormPositionPropertiesDto struct {
	value *OenormPositionPropertiesDto
	isSet bool
}

func (v NullableOenormPositionPropertiesDto) Get() *OenormPositionPropertiesDto {
	return v.value
}

func (v *NullableOenormPositionPropertiesDto) Set(val *OenormPositionPropertiesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOenormPositionPropertiesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOenormPositionPropertiesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOenormPositionPropertiesDto(val *OenormPositionPropertiesDto) *NullableOenormPositionPropertiesDto {
	return &NullableOenormPositionPropertiesDto{value: val, isSet: true}
}

func (v NullableOenormPositionPropertiesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOenormPositionPropertiesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
