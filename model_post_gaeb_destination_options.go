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

// checks if the PostGaebDestinationOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostGaebDestinationOptions{}

// PostGaebDestinationOptions Options for conversions to GAEB
type PostGaebDestinationOptions struct {
	DestinationGaebType          DestinationGaebType          `json:"destinationGaebType"`
	TargetExchangePhaseTransform DestinationGaebExchangePhase `json:"targetExchangePhaseTransform"`
	// Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export.
	EnforceStrictOfferPhaseLongTextOutput bool `json:"enforceStrictOfferPhaseLongTextOutput"`
	// Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the 'QtyDeterm' (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the 'QuantityComponents' property of positions. Please see the entry for 'Quantity Determination' in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the 'QuantityComponents' property.
	ExportQuantityDetermination bool `json:"exportQuantityDetermination"`
	// If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true.
	RemoveUnprintableCharactersFromTexts bool `json:"removeUnprintableCharactersFromTexts"`
	// If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions.
	ForceIncludeDescriptions bool `json:"forceIncludeDescriptions"`
	// When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons.
	TreatNullItemNumberSchemaAsInvalid bool `json:"treatNullItemNumberSchemaAsInvalid"`
}

// NewPostGaebDestinationOptions instantiates a new PostGaebDestinationOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostGaebDestinationOptions(destinationGaebType DestinationGaebType, targetExchangePhaseTransform DestinationGaebExchangePhase, enforceStrictOfferPhaseLongTextOutput bool, exportQuantityDetermination bool, removeUnprintableCharactersFromTexts bool, forceIncludeDescriptions bool, treatNullItemNumberSchemaAsInvalid bool) *PostGaebDestinationOptions {
	this := PostGaebDestinationOptions{}
	this.DestinationGaebType = destinationGaebType
	this.TargetExchangePhaseTransform = targetExchangePhaseTransform
	this.EnforceStrictOfferPhaseLongTextOutput = enforceStrictOfferPhaseLongTextOutput
	this.ExportQuantityDetermination = exportQuantityDetermination
	this.RemoveUnprintableCharactersFromTexts = removeUnprintableCharactersFromTexts
	this.ForceIncludeDescriptions = forceIncludeDescriptions
	this.TreatNullItemNumberSchemaAsInvalid = treatNullItemNumberSchemaAsInvalid
	return &this
}

// NewPostGaebDestinationOptionsWithDefaults instantiates a new PostGaebDestinationOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostGaebDestinationOptionsWithDefaults() *PostGaebDestinationOptions {
	this := PostGaebDestinationOptions{}
	return &this
}

// GetDestinationGaebType returns the DestinationGaebType field value
func (o *PostGaebDestinationOptions) GetDestinationGaebType() DestinationGaebType {
	if o == nil {
		var ret DestinationGaebType
		return ret
	}

	return o.DestinationGaebType
}

// GetDestinationGaebTypeOk returns a tuple with the DestinationGaebType field value
// and a boolean to check if the value has been set.
func (o *PostGaebDestinationOptions) GetDestinationGaebTypeOk() (*DestinationGaebType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationGaebType, true
}

// SetDestinationGaebType sets field value
func (o *PostGaebDestinationOptions) SetDestinationGaebType(v DestinationGaebType) {
	o.DestinationGaebType = v
}

// GetTargetExchangePhaseTransform returns the TargetExchangePhaseTransform field value
func (o *PostGaebDestinationOptions) GetTargetExchangePhaseTransform() DestinationGaebExchangePhase {
	if o == nil {
		var ret DestinationGaebExchangePhase
		return ret
	}

	return o.TargetExchangePhaseTransform
}

// GetTargetExchangePhaseTransformOk returns a tuple with the TargetExchangePhaseTransform field value
// and a boolean to check if the value has been set.
func (o *PostGaebDestinationOptions) GetTargetExchangePhaseTransformOk() (*DestinationGaebExchangePhase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetExchangePhaseTransform, true
}

// SetTargetExchangePhaseTransform sets field value
func (o *PostGaebDestinationOptions) SetTargetExchangePhaseTransform(v DestinationGaebExchangePhase) {
	o.TargetExchangePhaseTransform = v
}

// GetEnforceStrictOfferPhaseLongTextOutput returns the EnforceStrictOfferPhaseLongTextOutput field value
func (o *PostGaebDestinationOptions) GetEnforceStrictOfferPhaseLongTextOutput() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnforceStrictOfferPhaseLongTextOutput
}

// GetEnforceStrictOfferPhaseLongTextOutputOk returns a tuple with the EnforceStrictOfferPhaseLongTextOutput field value
// and a boolean to check if the value has been set.
func (o *PostGaebDestinationOptions) GetEnforceStrictOfferPhaseLongTextOutputOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnforceStrictOfferPhaseLongTextOutput, true
}

// SetEnforceStrictOfferPhaseLongTextOutput sets field value
func (o *PostGaebDestinationOptions) SetEnforceStrictOfferPhaseLongTextOutput(v bool) {
	o.EnforceStrictOfferPhaseLongTextOutput = v
}

// GetExportQuantityDetermination returns the ExportQuantityDetermination field value
func (o *PostGaebDestinationOptions) GetExportQuantityDetermination() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExportQuantityDetermination
}

// GetExportQuantityDeterminationOk returns a tuple with the ExportQuantityDetermination field value
// and a boolean to check if the value has been set.
func (o *PostGaebDestinationOptions) GetExportQuantityDeterminationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportQuantityDetermination, true
}

// SetExportQuantityDetermination sets field value
func (o *PostGaebDestinationOptions) SetExportQuantityDetermination(v bool) {
	o.ExportQuantityDetermination = v
}

// GetRemoveUnprintableCharactersFromTexts returns the RemoveUnprintableCharactersFromTexts field value
func (o *PostGaebDestinationOptions) GetRemoveUnprintableCharactersFromTexts() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RemoveUnprintableCharactersFromTexts
}

// GetRemoveUnprintableCharactersFromTextsOk returns a tuple with the RemoveUnprintableCharactersFromTexts field value
// and a boolean to check if the value has been set.
func (o *PostGaebDestinationOptions) GetRemoveUnprintableCharactersFromTextsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoveUnprintableCharactersFromTexts, true
}

// SetRemoveUnprintableCharactersFromTexts sets field value
func (o *PostGaebDestinationOptions) SetRemoveUnprintableCharactersFromTexts(v bool) {
	o.RemoveUnprintableCharactersFromTexts = v
}

// GetForceIncludeDescriptions returns the ForceIncludeDescriptions field value
func (o *PostGaebDestinationOptions) GetForceIncludeDescriptions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForceIncludeDescriptions
}

// GetForceIncludeDescriptionsOk returns a tuple with the ForceIncludeDescriptions field value
// and a boolean to check if the value has been set.
func (o *PostGaebDestinationOptions) GetForceIncludeDescriptionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForceIncludeDescriptions, true
}

// SetForceIncludeDescriptions sets field value
func (o *PostGaebDestinationOptions) SetForceIncludeDescriptions(v bool) {
	o.ForceIncludeDescriptions = v
}

// GetTreatNullItemNumberSchemaAsInvalid returns the TreatNullItemNumberSchemaAsInvalid field value
func (o *PostGaebDestinationOptions) GetTreatNullItemNumberSchemaAsInvalid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TreatNullItemNumberSchemaAsInvalid
}

// GetTreatNullItemNumberSchemaAsInvalidOk returns a tuple with the TreatNullItemNumberSchemaAsInvalid field value
// and a boolean to check if the value has been set.
func (o *PostGaebDestinationOptions) GetTreatNullItemNumberSchemaAsInvalidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TreatNullItemNumberSchemaAsInvalid, true
}

// SetTreatNullItemNumberSchemaAsInvalid sets field value
func (o *PostGaebDestinationOptions) SetTreatNullItemNumberSchemaAsInvalid(v bool) {
	o.TreatNullItemNumberSchemaAsInvalid = v
}

func (o PostGaebDestinationOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostGaebDestinationOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destinationGaebType"] = o.DestinationGaebType
	toSerialize["targetExchangePhaseTransform"] = o.TargetExchangePhaseTransform
	toSerialize["enforceStrictOfferPhaseLongTextOutput"] = o.EnforceStrictOfferPhaseLongTextOutput
	toSerialize["exportQuantityDetermination"] = o.ExportQuantityDetermination
	toSerialize["removeUnprintableCharactersFromTexts"] = o.RemoveUnprintableCharactersFromTexts
	toSerialize["forceIncludeDescriptions"] = o.ForceIncludeDescriptions
	toSerialize["treatNullItemNumberSchemaAsInvalid"] = o.TreatNullItemNumberSchemaAsInvalid
	return toSerialize, nil
}

type NullablePostGaebDestinationOptions struct {
	value *PostGaebDestinationOptions
	isSet bool
}

func (v NullablePostGaebDestinationOptions) Get() *PostGaebDestinationOptions {
	return v.value
}

func (v *NullablePostGaebDestinationOptions) Set(val *PostGaebDestinationOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePostGaebDestinationOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePostGaebDestinationOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostGaebDestinationOptions(val *PostGaebDestinationOptions) *NullablePostGaebDestinationOptions {
	return &NullablePostGaebDestinationOptions{value: val, isSet: true}
}

func (v NullablePostGaebDestinationOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostGaebDestinationOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
