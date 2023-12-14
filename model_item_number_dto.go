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

// checks if the ItemNumberDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemNumberDto{}

// ItemNumberDto This class represents as identifier of a certain service specification's element and is uniquely within the service specification.
type ItemNumberDto struct {
	// Elements GUID identifier.
	Id string `json:"id"`
	// Will return this ItemNumber as point delimited string. Will not distinguish between upper- and lowercase and return an all-lowercase representation. Will consider first numbers, then characters, e.g. 1a is considered preceding aa. Transformation to all lowercase can be configured in the ItemNumberSchema property.
	StringRepresentation *string `json:"stringRepresentation,omitempty"`
	// Indicates if the characters and the length of the Identifiers match the current ItemNumberSchema.
	IsSchemaCompliant bool                 `json:"isSchemaCompliant"`
	ItemNumberSchema  *ItemNumberSchemaDto `json:"itemNumberSchema,omitempty"`
	// Collection of the single identifiers in this ItemNumber. P.e., \"02.03.004\" will have three elements \"02\", \"03\", and \"004\". Since ReadOnlyObservableCollection`1 does have the event set to protected, it can be accessed like this: (itemNumber.Identifiers as INotifyCollectionChanged).CollectionChanged
	Identifiers []string `json:"identifiers,omitempty"`
	// This indicates if this item number is at the lot level. Find out more about lots in the documentation.
	IsLot bool `json:"isLot"`
	// This is a zero based hierarchy level. It's set automatically when used in the context of a Project, and can be used to identify the hierarchy level of the current element.
	HierarchyLevel int32 `json:"hierarchyLevel"`
	// This property indicates if this ItemNumber is attached to an object of the Position type.
	IsAttachedToPosition bool `json:"isAttachedToPosition"`
}

// NewItemNumberDto instantiates a new ItemNumberDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemNumberDto(id string, isSchemaCompliant bool, isLot bool, hierarchyLevel int32, isAttachedToPosition bool) *ItemNumberDto {
	this := ItemNumberDto{}
	this.Id = id
	this.IsSchemaCompliant = isSchemaCompliant
	this.IsLot = isLot
	this.HierarchyLevel = hierarchyLevel
	this.IsAttachedToPosition = isAttachedToPosition
	return &this
}

// NewItemNumberDtoWithDefaults instantiates a new ItemNumberDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemNumberDtoWithDefaults() *ItemNumberDto {
	this := ItemNumberDto{}
	return &this
}

// GetId returns the Id field value
func (o *ItemNumberDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ItemNumberDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ItemNumberDto) SetId(v string) {
	o.Id = v
}

// GetStringRepresentation returns the StringRepresentation field value if set, zero value otherwise.
func (o *ItemNumberDto) GetStringRepresentation() string {
	if o == nil || IsNil(o.StringRepresentation) {
		var ret string
		return ret
	}
	return *o.StringRepresentation
}

// GetStringRepresentationOk returns a tuple with the StringRepresentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemNumberDto) GetStringRepresentationOk() (*string, bool) {
	if o == nil || IsNil(o.StringRepresentation) {
		return nil, false
	}
	return o.StringRepresentation, true
}

// HasStringRepresentation returns a boolean if a field has been set.
func (o *ItemNumberDto) HasStringRepresentation() bool {
	if o != nil && !IsNil(o.StringRepresentation) {
		return true
	}

	return false
}

// SetStringRepresentation gets a reference to the given string and assigns it to the StringRepresentation field.
func (o *ItemNumberDto) SetStringRepresentation(v string) {
	o.StringRepresentation = &v
}

// GetIsSchemaCompliant returns the IsSchemaCompliant field value
func (o *ItemNumberDto) GetIsSchemaCompliant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSchemaCompliant
}

// GetIsSchemaCompliantOk returns a tuple with the IsSchemaCompliant field value
// and a boolean to check if the value has been set.
func (o *ItemNumberDto) GetIsSchemaCompliantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSchemaCompliant, true
}

// SetIsSchemaCompliant sets field value
func (o *ItemNumberDto) SetIsSchemaCompliant(v bool) {
	o.IsSchemaCompliant = v
}

// GetItemNumberSchema returns the ItemNumberSchema field value if set, zero value otherwise.
func (o *ItemNumberDto) GetItemNumberSchema() ItemNumberSchemaDto {
	if o == nil || IsNil(o.ItemNumberSchema) {
		var ret ItemNumberSchemaDto
		return ret
	}
	return *o.ItemNumberSchema
}

// GetItemNumberSchemaOk returns a tuple with the ItemNumberSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemNumberDto) GetItemNumberSchemaOk() (*ItemNumberSchemaDto, bool) {
	if o == nil || IsNil(o.ItemNumberSchema) {
		return nil, false
	}
	return o.ItemNumberSchema, true
}

// HasItemNumberSchema returns a boolean if a field has been set.
func (o *ItemNumberDto) HasItemNumberSchema() bool {
	if o != nil && !IsNil(o.ItemNumberSchema) {
		return true
	}

	return false
}

// SetItemNumberSchema gets a reference to the given ItemNumberSchemaDto and assigns it to the ItemNumberSchema field.
func (o *ItemNumberDto) SetItemNumberSchema(v ItemNumberSchemaDto) {
	o.ItemNumberSchema = &v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *ItemNumberDto) GetIdentifiers() []string {
	if o == nil || IsNil(o.Identifiers) {
		var ret []string
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemNumberDto) GetIdentifiersOk() ([]string, bool) {
	if o == nil || IsNil(o.Identifiers) {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *ItemNumberDto) HasIdentifiers() bool {
	if o != nil && !IsNil(o.Identifiers) {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []string and assigns it to the Identifiers field.
func (o *ItemNumberDto) SetIdentifiers(v []string) {
	o.Identifiers = v
}

// GetIsLot returns the IsLot field value
func (o *ItemNumberDto) GetIsLot() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLot
}

// GetIsLotOk returns a tuple with the IsLot field value
// and a boolean to check if the value has been set.
func (o *ItemNumberDto) GetIsLotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLot, true
}

// SetIsLot sets field value
func (o *ItemNumberDto) SetIsLot(v bool) {
	o.IsLot = v
}

// GetHierarchyLevel returns the HierarchyLevel field value
func (o *ItemNumberDto) GetHierarchyLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HierarchyLevel
}

// GetHierarchyLevelOk returns a tuple with the HierarchyLevel field value
// and a boolean to check if the value has been set.
func (o *ItemNumberDto) GetHierarchyLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HierarchyLevel, true
}

// SetHierarchyLevel sets field value
func (o *ItemNumberDto) SetHierarchyLevel(v int32) {
	o.HierarchyLevel = v
}

// GetIsAttachedToPosition returns the IsAttachedToPosition field value
func (o *ItemNumberDto) GetIsAttachedToPosition() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAttachedToPosition
}

// GetIsAttachedToPositionOk returns a tuple with the IsAttachedToPosition field value
// and a boolean to check if the value has been set.
func (o *ItemNumberDto) GetIsAttachedToPositionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAttachedToPosition, true
}

// SetIsAttachedToPosition sets field value
func (o *ItemNumberDto) SetIsAttachedToPosition(v bool) {
	o.IsAttachedToPosition = v
}

func (o ItemNumberDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemNumberDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.StringRepresentation) {
		toSerialize["stringRepresentation"] = o.StringRepresentation
	}
	// skip: isSchemaCompliant is readOnly
	if !IsNil(o.ItemNumberSchema) {
		toSerialize["itemNumberSchema"] = o.ItemNumberSchema
	}
	if !IsNil(o.Identifiers) {
		toSerialize["identifiers"] = o.Identifiers
	}
	// skip: isLot is readOnly
	toSerialize["hierarchyLevel"] = o.HierarchyLevel
	// skip: isAttachedToPosition is readOnly
	return toSerialize, nil
}

type NullableItemNumberDto struct {
	value *ItemNumberDto
	isSet bool
}

func (v NullableItemNumberDto) Get() *ItemNumberDto {
	return v.value
}

func (v *NullableItemNumberDto) Set(val *ItemNumberDto) {
	v.value = val
	v.isSet = true
}

func (v NullableItemNumberDto) IsSet() bool {
	return v.isSet
}

func (v *NullableItemNumberDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemNumberDto(val *ItemNumberDto) *NullableItemNumberDto {
	return &NullableItemNumberDto{value: val, isSet: true}
}

func (v NullableItemNumberDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemNumberDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
