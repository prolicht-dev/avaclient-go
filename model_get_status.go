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

// checks if the GetStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatus{}

// GetStatus Indicates the status of the AVACloud service
type GetStatus struct {
	// If any problems in the service health are known, this is set to false
	IsHealthy bool `json:"isHealthy"`
	// The current version of the AVACloud service
	Version *string `json:"version,omitempty"`
	// The environment of the current instance
	Environment *string `json:"environment,omitempty"`
}

// NewGetStatus instantiates a new GetStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus(isHealthy bool) *GetStatus {
	this := GetStatus{}
	this.IsHealthy = isHealthy
	return &this
}

// NewGetStatusWithDefaults instantiates a new GetStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatusWithDefaults() *GetStatus {
	this := GetStatus{}
	return &this
}

// GetIsHealthy returns the IsHealthy field value
func (o *GetStatus) GetIsHealthy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHealthy
}

// GetIsHealthyOk returns a tuple with the IsHealthy field value
// and a boolean to check if the value has been set.
func (o *GetStatus) GetIsHealthyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHealthy, true
}

// SetIsHealthy sets field value
func (o *GetStatus) SetIsHealthy(v bool) {
	o.IsHealthy = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetStatus) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetStatus) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetStatus) SetVersion(v string) {
	o.Version = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *GetStatus) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *GetStatus) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *GetStatus) SetEnvironment(v string) {
	o.Environment = &v
}

func (o GetStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isHealthy"] = o.IsHealthy
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullableGetStatus struct {
	value *GetStatus
	isSet bool
}

func (v NullableGetStatus) Get() *GetStatus {
	return v.value
}

func (v *NullableGetStatus) Set(val *GetStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus(val *GetStatus) *NullableGetStatus {
	return &NullableGetStatus{value: val, isSet: true}
}

func (v NullableGetStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
