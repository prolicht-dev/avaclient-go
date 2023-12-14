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

// checks if the GrossPriceComponentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrossPriceComponentDto{}

// GrossPriceComponentDto This is used in an ElementContainerBase to hold the price composition.
type GrossPriceComponentDto struct {
	// This components tax rate.
	TaxRate float32 `json:"taxRate"`
	// This is the factor of applied deductions for this component
	DeductionFactor float32 `json:"deductionFactor"`
	// The total net price for all components of a given tax rate.
	TotalNet float32 `json:"totalNet"`
	// The resulting price component after applied deductions
	TotalDeducted float32 `json:"totalDeducted"`
	// The total tax amount for all components of a given tax rate.
	TotalTax float32 `json:"totalTax"`
	// The total gross price for all components of a given tax rate.
	TotalGross float32 `json:"totalGross"`
	// The total gross price for all components of a given tax rate, after applied deductions.
	TotalGrossDeducted float32 `json:"totalGrossDeducted"`
	// The total tax amount for all components of a given tax rate, after applied deductions.
	TotalTaxDeducted float32 `json:"totalTaxDeducted"`
}

// NewGrossPriceComponentDto instantiates a new GrossPriceComponentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrossPriceComponentDto(taxRate float32, deductionFactor float32, totalNet float32, totalDeducted float32, totalTax float32, totalGross float32, totalGrossDeducted float32, totalTaxDeducted float32) *GrossPriceComponentDto {
	this := GrossPriceComponentDto{}
	this.TaxRate = taxRate
	this.DeductionFactor = deductionFactor
	this.TotalNet = totalNet
	this.TotalDeducted = totalDeducted
	this.TotalTax = totalTax
	this.TotalGross = totalGross
	this.TotalGrossDeducted = totalGrossDeducted
	this.TotalTaxDeducted = totalTaxDeducted
	return &this
}

// NewGrossPriceComponentDtoWithDefaults instantiates a new GrossPriceComponentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrossPriceComponentDtoWithDefaults() *GrossPriceComponentDto {
	this := GrossPriceComponentDto{}
	return &this
}

// GetTaxRate returns the TaxRate field value
func (o *GrossPriceComponentDto) GetTaxRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TaxRate
}

// GetTaxRateOk returns a tuple with the TaxRate field value
// and a boolean to check if the value has been set.
func (o *GrossPriceComponentDto) GetTaxRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxRate, true
}

// SetTaxRate sets field value
func (o *GrossPriceComponentDto) SetTaxRate(v float32) {
	o.TaxRate = v
}

// GetDeductionFactor returns the DeductionFactor field value
func (o *GrossPriceComponentDto) GetDeductionFactor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DeductionFactor
}

// GetDeductionFactorOk returns a tuple with the DeductionFactor field value
// and a boolean to check if the value has been set.
func (o *GrossPriceComponentDto) GetDeductionFactorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeductionFactor, true
}

// SetDeductionFactor sets field value
func (o *GrossPriceComponentDto) SetDeductionFactor(v float32) {
	o.DeductionFactor = v
}

// GetTotalNet returns the TotalNet field value
func (o *GrossPriceComponentDto) GetTotalNet() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalNet
}

// GetTotalNetOk returns a tuple with the TotalNet field value
// and a boolean to check if the value has been set.
func (o *GrossPriceComponentDto) GetTotalNetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNet, true
}

// SetTotalNet sets field value
func (o *GrossPriceComponentDto) SetTotalNet(v float32) {
	o.TotalNet = v
}

// GetTotalDeducted returns the TotalDeducted field value
func (o *GrossPriceComponentDto) GetTotalDeducted() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalDeducted
}

// GetTotalDeductedOk returns a tuple with the TotalDeducted field value
// and a boolean to check if the value has been set.
func (o *GrossPriceComponentDto) GetTotalDeductedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDeducted, true
}

// SetTotalDeducted sets field value
func (o *GrossPriceComponentDto) SetTotalDeducted(v float32) {
	o.TotalDeducted = v
}

// GetTotalTax returns the TotalTax field value
func (o *GrossPriceComponentDto) GetTotalTax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value
// and a boolean to check if the value has been set.
func (o *GrossPriceComponentDto) GetTotalTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTax, true
}

// SetTotalTax sets field value
func (o *GrossPriceComponentDto) SetTotalTax(v float32) {
	o.TotalTax = v
}

// GetTotalGross returns the TotalGross field value
func (o *GrossPriceComponentDto) GetTotalGross() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalGross
}

// GetTotalGrossOk returns a tuple with the TotalGross field value
// and a boolean to check if the value has been set.
func (o *GrossPriceComponentDto) GetTotalGrossOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalGross, true
}

// SetTotalGross sets field value
func (o *GrossPriceComponentDto) SetTotalGross(v float32) {
	o.TotalGross = v
}

// GetTotalGrossDeducted returns the TotalGrossDeducted field value
func (o *GrossPriceComponentDto) GetTotalGrossDeducted() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalGrossDeducted
}

// GetTotalGrossDeductedOk returns a tuple with the TotalGrossDeducted field value
// and a boolean to check if the value has been set.
func (o *GrossPriceComponentDto) GetTotalGrossDeductedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalGrossDeducted, true
}

// SetTotalGrossDeducted sets field value
func (o *GrossPriceComponentDto) SetTotalGrossDeducted(v float32) {
	o.TotalGrossDeducted = v
}

// GetTotalTaxDeducted returns the TotalTaxDeducted field value
func (o *GrossPriceComponentDto) GetTotalTaxDeducted() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalTaxDeducted
}

// GetTotalTaxDeductedOk returns a tuple with the TotalTaxDeducted field value
// and a boolean to check if the value has been set.
func (o *GrossPriceComponentDto) GetTotalTaxDeductedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTaxDeducted, true
}

// SetTotalTaxDeducted sets field value
func (o *GrossPriceComponentDto) SetTotalTaxDeducted(v float32) {
	o.TotalTaxDeducted = v
}

func (o GrossPriceComponentDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrossPriceComponentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["taxRate"] = o.TaxRate
	toSerialize["deductionFactor"] = o.DeductionFactor
	toSerialize["totalNet"] = o.TotalNet
	// skip: totalDeducted is readOnly
	// skip: totalTax is readOnly
	// skip: totalGross is readOnly
	// skip: totalGrossDeducted is readOnly
	// skip: totalTaxDeducted is readOnly
	return toSerialize, nil
}

type NullableGrossPriceComponentDto struct {
	value *GrossPriceComponentDto
	isSet bool
}

func (v NullableGrossPriceComponentDto) Get() *GrossPriceComponentDto {
	return v.value
}

func (v *NullableGrossPriceComponentDto) Set(val *GrossPriceComponentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGrossPriceComponentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGrossPriceComponentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrossPriceComponentDto(val *GrossPriceComponentDto) *NullableGrossPriceComponentDto {
	return &NullableGrossPriceComponentDto{value: val, isSet: true}
}

func (v NullableGrossPriceComponentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrossPriceComponentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
