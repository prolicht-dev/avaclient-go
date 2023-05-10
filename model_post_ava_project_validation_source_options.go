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

// checks if the PostAvaProjectValidationSourceOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAvaProjectValidationSourceOptions{}

// PostAvaProjectValidationSourceOptions Options for validating an AVA project
type PostAvaProjectValidationSourceOptions struct {
	AvaProject ProjectDto `json:"avaProject"`
	// The index of the ServiceSpecification that should be validated. If not given, will default to the first one in the project.
	ServiceSpecificationIndex *int32                        `json:"serviceSpecificationIndex,omitempty"`
	ValidationType            ValidationType                `json:"validationType"`
	ExchangePhase             *ExchangePhaseDto             `json:"exchangePhase,omitempty"`
	AvaSourceOptions          *PostAvaSourceOptions         `json:"avaSourceOptions,omitempty"`
	OenormDestinationOptions  *PostOenormDestinationOptions `json:"oenormDestinationOptions,omitempty"`
	GaebDestinationOptions    *PostGaebDestinationOptions   `json:"gaebDestinationOptions,omitempty"`
}

// NewPostAvaProjectValidationSourceOptions instantiates a new PostAvaProjectValidationSourceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAvaProjectValidationSourceOptions(avaProject ProjectDto, validationType ValidationType) *PostAvaProjectValidationSourceOptions {
	this := PostAvaProjectValidationSourceOptions{}
	this.AvaProject = avaProject
	this.ValidationType = validationType
	return &this
}

// NewPostAvaProjectValidationSourceOptionsWithDefaults instantiates a new PostAvaProjectValidationSourceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAvaProjectValidationSourceOptionsWithDefaults() *PostAvaProjectValidationSourceOptions {
	this := PostAvaProjectValidationSourceOptions{}
	return &this
}

// GetAvaProject returns the AvaProject field value
func (o *PostAvaProjectValidationSourceOptions) GetAvaProject() ProjectDto {
	if o == nil {
		var ret ProjectDto
		return ret
	}

	return o.AvaProject
}

// GetAvaProjectOk returns a tuple with the AvaProject field value
// and a boolean to check if the value has been set.
func (o *PostAvaProjectValidationSourceOptions) GetAvaProjectOk() (*ProjectDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvaProject, true
}

// SetAvaProject sets field value
func (o *PostAvaProjectValidationSourceOptions) SetAvaProject(v ProjectDto) {
	o.AvaProject = v
}

// GetServiceSpecificationIndex returns the ServiceSpecificationIndex field value if set, zero value otherwise.
func (o *PostAvaProjectValidationSourceOptions) GetServiceSpecificationIndex() int32 {
	if o == nil || IsNil(o.ServiceSpecificationIndex) {
		var ret int32
		return ret
	}
	return *o.ServiceSpecificationIndex
}

// GetServiceSpecificationIndexOk returns a tuple with the ServiceSpecificationIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAvaProjectValidationSourceOptions) GetServiceSpecificationIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceSpecificationIndex) {
		return nil, false
	}
	return o.ServiceSpecificationIndex, true
}

// HasServiceSpecificationIndex returns a boolean if a field has been set.
func (o *PostAvaProjectValidationSourceOptions) HasServiceSpecificationIndex() bool {
	if o != nil && !IsNil(o.ServiceSpecificationIndex) {
		return true
	}

	return false
}

// SetServiceSpecificationIndex gets a reference to the given int32 and assigns it to the ServiceSpecificationIndex field.
func (o *PostAvaProjectValidationSourceOptions) SetServiceSpecificationIndex(v int32) {
	o.ServiceSpecificationIndex = &v
}

// GetValidationType returns the ValidationType field value
func (o *PostAvaProjectValidationSourceOptions) GetValidationType() ValidationType {
	if o == nil {
		var ret ValidationType
		return ret
	}

	return o.ValidationType
}

// GetValidationTypeOk returns a tuple with the ValidationType field value
// and a boolean to check if the value has been set.
func (o *PostAvaProjectValidationSourceOptions) GetValidationTypeOk() (*ValidationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationType, true
}

// SetValidationType sets field value
func (o *PostAvaProjectValidationSourceOptions) SetValidationType(v ValidationType) {
	o.ValidationType = v
}

// GetExchangePhase returns the ExchangePhase field value if set, zero value otherwise.
func (o *PostAvaProjectValidationSourceOptions) GetExchangePhase() ExchangePhaseDto {
	if o == nil || IsNil(o.ExchangePhase) {
		var ret ExchangePhaseDto
		return ret
	}
	return *o.ExchangePhase
}

// GetExchangePhaseOk returns a tuple with the ExchangePhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAvaProjectValidationSourceOptions) GetExchangePhaseOk() (*ExchangePhaseDto, bool) {
	if o == nil || IsNil(o.ExchangePhase) {
		return nil, false
	}
	return o.ExchangePhase, true
}

// HasExchangePhase returns a boolean if a field has been set.
func (o *PostAvaProjectValidationSourceOptions) HasExchangePhase() bool {
	if o != nil && !IsNil(o.ExchangePhase) {
		return true
	}

	return false
}

// SetExchangePhase gets a reference to the given ExchangePhaseDto and assigns it to the ExchangePhase field.
func (o *PostAvaProjectValidationSourceOptions) SetExchangePhase(v ExchangePhaseDto) {
	o.ExchangePhase = &v
}

// GetAvaSourceOptions returns the AvaSourceOptions field value if set, zero value otherwise.
func (o *PostAvaProjectValidationSourceOptions) GetAvaSourceOptions() PostAvaSourceOptions {
	if o == nil || IsNil(o.AvaSourceOptions) {
		var ret PostAvaSourceOptions
		return ret
	}
	return *o.AvaSourceOptions
}

// GetAvaSourceOptionsOk returns a tuple with the AvaSourceOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAvaProjectValidationSourceOptions) GetAvaSourceOptionsOk() (*PostAvaSourceOptions, bool) {
	if o == nil || IsNil(o.AvaSourceOptions) {
		return nil, false
	}
	return o.AvaSourceOptions, true
}

// HasAvaSourceOptions returns a boolean if a field has been set.
func (o *PostAvaProjectValidationSourceOptions) HasAvaSourceOptions() bool {
	if o != nil && !IsNil(o.AvaSourceOptions) {
		return true
	}

	return false
}

// SetAvaSourceOptions gets a reference to the given PostAvaSourceOptions and assigns it to the AvaSourceOptions field.
func (o *PostAvaProjectValidationSourceOptions) SetAvaSourceOptions(v PostAvaSourceOptions) {
	o.AvaSourceOptions = &v
}

// GetOenormDestinationOptions returns the OenormDestinationOptions field value if set, zero value otherwise.
func (o *PostAvaProjectValidationSourceOptions) GetOenormDestinationOptions() PostOenormDestinationOptions {
	if o == nil || IsNil(o.OenormDestinationOptions) {
		var ret PostOenormDestinationOptions
		return ret
	}
	return *o.OenormDestinationOptions
}

// GetOenormDestinationOptionsOk returns a tuple with the OenormDestinationOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAvaProjectValidationSourceOptions) GetOenormDestinationOptionsOk() (*PostOenormDestinationOptions, bool) {
	if o == nil || IsNil(o.OenormDestinationOptions) {
		return nil, false
	}
	return o.OenormDestinationOptions, true
}

// HasOenormDestinationOptions returns a boolean if a field has been set.
func (o *PostAvaProjectValidationSourceOptions) HasOenormDestinationOptions() bool {
	if o != nil && !IsNil(o.OenormDestinationOptions) {
		return true
	}

	return false
}

// SetOenormDestinationOptions gets a reference to the given PostOenormDestinationOptions and assigns it to the OenormDestinationOptions field.
func (o *PostAvaProjectValidationSourceOptions) SetOenormDestinationOptions(v PostOenormDestinationOptions) {
	o.OenormDestinationOptions = &v
}

// GetGaebDestinationOptions returns the GaebDestinationOptions field value if set, zero value otherwise.
func (o *PostAvaProjectValidationSourceOptions) GetGaebDestinationOptions() PostGaebDestinationOptions {
	if o == nil || IsNil(o.GaebDestinationOptions) {
		var ret PostGaebDestinationOptions
		return ret
	}
	return *o.GaebDestinationOptions
}

// GetGaebDestinationOptionsOk returns a tuple with the GaebDestinationOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAvaProjectValidationSourceOptions) GetGaebDestinationOptionsOk() (*PostGaebDestinationOptions, bool) {
	if o == nil || IsNil(o.GaebDestinationOptions) {
		return nil, false
	}
	return o.GaebDestinationOptions, true
}

// HasGaebDestinationOptions returns a boolean if a field has been set.
func (o *PostAvaProjectValidationSourceOptions) HasGaebDestinationOptions() bool {
	if o != nil && !IsNil(o.GaebDestinationOptions) {
		return true
	}

	return false
}

// SetGaebDestinationOptions gets a reference to the given PostGaebDestinationOptions and assigns it to the GaebDestinationOptions field.
func (o *PostAvaProjectValidationSourceOptions) SetGaebDestinationOptions(v PostGaebDestinationOptions) {
	o.GaebDestinationOptions = &v
}

func (o PostAvaProjectValidationSourceOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAvaProjectValidationSourceOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["avaProject"] = o.AvaProject
	if !IsNil(o.ServiceSpecificationIndex) {
		toSerialize["serviceSpecificationIndex"] = o.ServiceSpecificationIndex
	}
	toSerialize["validationType"] = o.ValidationType
	if !IsNil(o.ExchangePhase) {
		toSerialize["exchangePhase"] = o.ExchangePhase
	}
	if !IsNil(o.AvaSourceOptions) {
		toSerialize["avaSourceOptions"] = o.AvaSourceOptions
	}
	if !IsNil(o.OenormDestinationOptions) {
		toSerialize["oenormDestinationOptions"] = o.OenormDestinationOptions
	}
	if !IsNil(o.GaebDestinationOptions) {
		toSerialize["gaebDestinationOptions"] = o.GaebDestinationOptions
	}
	return toSerialize, nil
}

type NullablePostAvaProjectValidationSourceOptions struct {
	value *PostAvaProjectValidationSourceOptions
	isSet bool
}

func (v NullablePostAvaProjectValidationSourceOptions) Get() *PostAvaProjectValidationSourceOptions {
	return v.value
}

func (v *NullablePostAvaProjectValidationSourceOptions) Set(val *PostAvaProjectValidationSourceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAvaProjectValidationSourceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAvaProjectValidationSourceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAvaProjectValidationSourceOptions(val *PostAvaProjectValidationSourceOptions) *NullablePostAvaProjectValidationSourceOptions {
	return &NullablePostAvaProjectValidationSourceOptions{value: val, isSet: true}
}

func (v NullablePostAvaProjectValidationSourceOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAvaProjectValidationSourceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}