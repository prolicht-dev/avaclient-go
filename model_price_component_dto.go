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

// checks if the PriceComponentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceComponentDto{}

// PriceComponentDto This class provides means to store information about a price and it's composition. Note that this is referencing to a single price component, so for example a Position would have a list of PriceComponents, one for Material, one for Labour etc.
type PriceComponentDto struct {
	// The total, calculated price of this component.
	Price float32 `json:"price"`
	// The label associated with this price component. Will be taken from the parent Projects ProjectInformation.
	Label *string `json:"label,omitempty"`
	// The single Calculation elements this price component is composed of.
	Values []CalculationDto `json:"values,omitempty"`
	// These are Catalogues that are used within this PriceComponent. Catalogues are used to describe catalogues, or collections, that can be used to describe elements with commonly known properties. For example, QuantityAssignments use these to categorize themselves. They are propagate to all child elements, e.g. other containers and QuantityAssignments. In the context of a ServiceSpecification, all elements share the same instance of the collection.
	ProjectCatalogues []CatalogueDto `json:"projectCatalogues,omitempty"`
}

// NewPriceComponentDto instantiates a new PriceComponentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceComponentDto(price float32) *PriceComponentDto {
	this := PriceComponentDto{}
	this.Price = price
	return &this
}

// NewPriceComponentDtoWithDefaults instantiates a new PriceComponentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceComponentDtoWithDefaults() *PriceComponentDto {
	this := PriceComponentDto{}
	return &this
}

// GetPrice returns the Price field value
func (o *PriceComponentDto) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *PriceComponentDto) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *PriceComponentDto) SetPrice(v float32) {
	o.Price = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PriceComponentDto) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceComponentDto) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PriceComponentDto) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PriceComponentDto) SetLabel(v string) {
	o.Label = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *PriceComponentDto) GetValues() []CalculationDto {
	if o == nil || IsNil(o.Values) {
		var ret []CalculationDto
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceComponentDto) GetValuesOk() ([]CalculationDto, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *PriceComponentDto) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []CalculationDto and assigns it to the Values field.
func (o *PriceComponentDto) SetValues(v []CalculationDto) {
	o.Values = v
}

// GetProjectCatalogues returns the ProjectCatalogues field value if set, zero value otherwise.
func (o *PriceComponentDto) GetProjectCatalogues() []CatalogueDto {
	if o == nil || IsNil(o.ProjectCatalogues) {
		var ret []CatalogueDto
		return ret
	}
	return o.ProjectCatalogues
}

// GetProjectCataloguesOk returns a tuple with the ProjectCatalogues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceComponentDto) GetProjectCataloguesOk() ([]CatalogueDto, bool) {
	if o == nil || IsNil(o.ProjectCatalogues) {
		return nil, false
	}
	return o.ProjectCatalogues, true
}

// HasProjectCatalogues returns a boolean if a field has been set.
func (o *PriceComponentDto) HasProjectCatalogues() bool {
	if o != nil && !IsNil(o.ProjectCatalogues) {
		return true
	}

	return false
}

// SetProjectCatalogues gets a reference to the given []CatalogueDto and assigns it to the ProjectCatalogues field.
func (o *PriceComponentDto) SetProjectCatalogues(v []CatalogueDto) {
	o.ProjectCatalogues = v
}

func (o PriceComponentDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceComponentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: price is readOnly
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.ProjectCatalogues) {
		toSerialize["projectCatalogues"] = o.ProjectCatalogues
	}
	return toSerialize, nil
}

type NullablePriceComponentDto struct {
	value *PriceComponentDto
	isSet bool
}

func (v NullablePriceComponentDto) Get() *PriceComponentDto {
	return v.value
}

func (v *NullablePriceComponentDto) Set(val *PriceComponentDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceComponentDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceComponentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceComponentDto(val *PriceComponentDto) *NullablePriceComponentDto {
	return &NullablePriceComponentDto{value: val, isSet: true}
}

func (v NullablePriceComponentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceComponentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
