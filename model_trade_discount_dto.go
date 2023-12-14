/*
AVACloud API 2.0.0

AVACloud API specification

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
)

// checks if the TradeDiscountDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TradeDiscountDto{}

// TradeDiscountDto This class holds information about offered trade discounts (Skonto in German)
type TradeDiscountDto struct {
	// Elements GUID identifier.
	Id string `json:"id"`
	// The amount of days for how long a declared trade discount applies. Must be a positive number, negative values will be ignored and not set.
	Deadline int32 `json:"deadline"`
	// The rate of the trade discount.
	Rate float32 `json:"rate"`
}

// NewTradeDiscountDto instantiates a new TradeDiscountDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeDiscountDto(id string, deadline int32, rate float32) *TradeDiscountDto {
	this := TradeDiscountDto{}
	this.Id = id
	this.Deadline = deadline
	this.Rate = rate
	return &this
}

// NewTradeDiscountDtoWithDefaults instantiates a new TradeDiscountDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeDiscountDtoWithDefaults() *TradeDiscountDto {
	this := TradeDiscountDto{}
	return &this
}

// GetId returns the Id field value
func (o *TradeDiscountDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TradeDiscountDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TradeDiscountDto) SetId(v string) {
	o.Id = v
}

// GetDeadline returns the Deadline field value
func (o *TradeDiscountDto) GetDeadline() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value
// and a boolean to check if the value has been set.
func (o *TradeDiscountDto) GetDeadlineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deadline, true
}

// SetDeadline sets field value
func (o *TradeDiscountDto) SetDeadline(v int32) {
	o.Deadline = v
}

// GetRate returns the Rate field value
func (o *TradeDiscountDto) GetRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *TradeDiscountDto) GetRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *TradeDiscountDto) SetRate(v float32) {
	o.Rate = v
}

func (o TradeDiscountDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeDiscountDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["deadline"] = o.Deadline
	toSerialize["rate"] = o.Rate
	return toSerialize, nil
}

type NullableTradeDiscountDto struct {
	value *TradeDiscountDto
	isSet bool
}

func (v NullableTradeDiscountDto) Get() *TradeDiscountDto {
	return v.value
}

func (v *NullableTradeDiscountDto) Set(val *TradeDiscountDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeDiscountDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeDiscountDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeDiscountDto(val *TradeDiscountDto) *NullableTradeDiscountDto {
	return &NullableTradeDiscountDto{value: val, isSet: true}
}

func (v NullableTradeDiscountDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeDiscountDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
