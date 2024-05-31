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

// checks if the CommercePropertiesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommercePropertiesDto{}

// CommercePropertiesDto This class holds specialized information that is relevant to trade or commerce based data exchange scenarios, e.g. between customers, vendors, suppliers and distributors. It is used when exporting to GAEB XML 9x exchange phases.
type CommercePropertiesDto struct {
	// This maps to ArtNo in GAEB XML and represents an article number given by the supplier.
	ArticleNumber *string `json:"articleNumber,omitempty"`
	// This maps to EAN in GAEB XML and represents an GTIN (formerly EAN) article number.
	EanGtinArticleNumber *string `json:"eanGtinArticleNumber,omitempty"`
	// This maps to ArtNoID in GAEB XML and represents an ILN article number.
	IlnArticleNumber *string `json:"ilnArticleNumber,omitempty"`
	// This maps to CatalogNo in GAEB XML and represents an identifier of a specific catalogue. The referenced catalogue is usually a customer specific one, not related to CatalogueReferences.
	CatalogueNumber *string `json:"catalogueNumber,omitempty"`
	// This maps to CatalogArtNo in GAEB XML and represents a key that maps to an entry in a specific catalogue. The referenced catalogue is usually a customer specific one, not related to CatalogueReferences.
	CatalogueArticleNumber *string `json:"catalogueArticleNumber,omitempty"`
	// This optional property can be used to indicate the basis for prices for a single position. Price basis means that the price is given per unit of the basis, e.g. per a pack of five when this property is set to '5'.
	PriceBasis *float32 `json:"priceBasis,omitempty"`
	// This optional property can be used to indicate a sub position identifier specific for the commerce exchange
	SubPositionIdentifier *string `json:"subPositionIdentifier,omitempty"`
	// In a commerce exchange, this property is usually used to reference the base item number of an underlying phase 83 exchange file
	CustomerBaseItemNumber *string `json:"customerBaseItemNumber,omitempty"`
}

// NewCommercePropertiesDto instantiates a new CommercePropertiesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommercePropertiesDto() *CommercePropertiesDto {
	this := CommercePropertiesDto{}
	return &this
}

// NewCommercePropertiesDtoWithDefaults instantiates a new CommercePropertiesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommercePropertiesDtoWithDefaults() *CommercePropertiesDto {
	this := CommercePropertiesDto{}
	return &this
}

// GetArticleNumber returns the ArticleNumber field value if set, zero value otherwise.
func (o *CommercePropertiesDto) GetArticleNumber() string {
	if o == nil || IsNil(o.ArticleNumber) {
		var ret string
		return ret
	}
	return *o.ArticleNumber
}

// GetArticleNumberOk returns a tuple with the ArticleNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercePropertiesDto) GetArticleNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ArticleNumber) {
		return nil, false
	}
	return o.ArticleNumber, true
}

// HasArticleNumber returns a boolean if a field has been set.
func (o *CommercePropertiesDto) HasArticleNumber() bool {
	if o != nil && !IsNil(o.ArticleNumber) {
		return true
	}

	return false
}

// SetArticleNumber gets a reference to the given string and assigns it to the ArticleNumber field.
func (o *CommercePropertiesDto) SetArticleNumber(v string) {
	o.ArticleNumber = &v
}

// GetEanGtinArticleNumber returns the EanGtinArticleNumber field value if set, zero value otherwise.
func (o *CommercePropertiesDto) GetEanGtinArticleNumber() string {
	if o == nil || IsNil(o.EanGtinArticleNumber) {
		var ret string
		return ret
	}
	return *o.EanGtinArticleNumber
}

// GetEanGtinArticleNumberOk returns a tuple with the EanGtinArticleNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercePropertiesDto) GetEanGtinArticleNumberOk() (*string, bool) {
	if o == nil || IsNil(o.EanGtinArticleNumber) {
		return nil, false
	}
	return o.EanGtinArticleNumber, true
}

// HasEanGtinArticleNumber returns a boolean if a field has been set.
func (o *CommercePropertiesDto) HasEanGtinArticleNumber() bool {
	if o != nil && !IsNil(o.EanGtinArticleNumber) {
		return true
	}

	return false
}

// SetEanGtinArticleNumber gets a reference to the given string and assigns it to the EanGtinArticleNumber field.
func (o *CommercePropertiesDto) SetEanGtinArticleNumber(v string) {
	o.EanGtinArticleNumber = &v
}

// GetIlnArticleNumber returns the IlnArticleNumber field value if set, zero value otherwise.
func (o *CommercePropertiesDto) GetIlnArticleNumber() string {
	if o == nil || IsNil(o.IlnArticleNumber) {
		var ret string
		return ret
	}
	return *o.IlnArticleNumber
}

// GetIlnArticleNumberOk returns a tuple with the IlnArticleNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercePropertiesDto) GetIlnArticleNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IlnArticleNumber) {
		return nil, false
	}
	return o.IlnArticleNumber, true
}

// HasIlnArticleNumber returns a boolean if a field has been set.
func (o *CommercePropertiesDto) HasIlnArticleNumber() bool {
	if o != nil && !IsNil(o.IlnArticleNumber) {
		return true
	}

	return false
}

// SetIlnArticleNumber gets a reference to the given string and assigns it to the IlnArticleNumber field.
func (o *CommercePropertiesDto) SetIlnArticleNumber(v string) {
	o.IlnArticleNumber = &v
}

// GetCatalogueNumber returns the CatalogueNumber field value if set, zero value otherwise.
func (o *CommercePropertiesDto) GetCatalogueNumber() string {
	if o == nil || IsNil(o.CatalogueNumber) {
		var ret string
		return ret
	}
	return *o.CatalogueNumber
}

// GetCatalogueNumberOk returns a tuple with the CatalogueNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercePropertiesDto) GetCatalogueNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogueNumber) {
		return nil, false
	}
	return o.CatalogueNumber, true
}

// HasCatalogueNumber returns a boolean if a field has been set.
func (o *CommercePropertiesDto) HasCatalogueNumber() bool {
	if o != nil && !IsNil(o.CatalogueNumber) {
		return true
	}

	return false
}

// SetCatalogueNumber gets a reference to the given string and assigns it to the CatalogueNumber field.
func (o *CommercePropertiesDto) SetCatalogueNumber(v string) {
	o.CatalogueNumber = &v
}

// GetCatalogueArticleNumber returns the CatalogueArticleNumber field value if set, zero value otherwise.
func (o *CommercePropertiesDto) GetCatalogueArticleNumber() string {
	if o == nil || IsNil(o.CatalogueArticleNumber) {
		var ret string
		return ret
	}
	return *o.CatalogueArticleNumber
}

// GetCatalogueArticleNumberOk returns a tuple with the CatalogueArticleNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercePropertiesDto) GetCatalogueArticleNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogueArticleNumber) {
		return nil, false
	}
	return o.CatalogueArticleNumber, true
}

// HasCatalogueArticleNumber returns a boolean if a field has been set.
func (o *CommercePropertiesDto) HasCatalogueArticleNumber() bool {
	if o != nil && !IsNil(o.CatalogueArticleNumber) {
		return true
	}

	return false
}

// SetCatalogueArticleNumber gets a reference to the given string and assigns it to the CatalogueArticleNumber field.
func (o *CommercePropertiesDto) SetCatalogueArticleNumber(v string) {
	o.CatalogueArticleNumber = &v
}

// GetPriceBasis returns the PriceBasis field value if set, zero value otherwise.
func (o *CommercePropertiesDto) GetPriceBasis() float32 {
	if o == nil || IsNil(o.PriceBasis) {
		var ret float32
		return ret
	}
	return *o.PriceBasis
}

// GetPriceBasisOk returns a tuple with the PriceBasis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercePropertiesDto) GetPriceBasisOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceBasis) {
		return nil, false
	}
	return o.PriceBasis, true
}

// HasPriceBasis returns a boolean if a field has been set.
func (o *CommercePropertiesDto) HasPriceBasis() bool {
	if o != nil && !IsNil(o.PriceBasis) {
		return true
	}

	return false
}

// SetPriceBasis gets a reference to the given float32 and assigns it to the PriceBasis field.
func (o *CommercePropertiesDto) SetPriceBasis(v float32) {
	o.PriceBasis = &v
}

// GetSubPositionIdentifier returns the SubPositionIdentifier field value if set, zero value otherwise.
func (o *CommercePropertiesDto) GetSubPositionIdentifier() string {
	if o == nil || IsNil(o.SubPositionIdentifier) {
		var ret string
		return ret
	}
	return *o.SubPositionIdentifier
}

// GetSubPositionIdentifierOk returns a tuple with the SubPositionIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercePropertiesDto) GetSubPositionIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.SubPositionIdentifier) {
		return nil, false
	}
	return o.SubPositionIdentifier, true
}

// HasSubPositionIdentifier returns a boolean if a field has been set.
func (o *CommercePropertiesDto) HasSubPositionIdentifier() bool {
	if o != nil && !IsNil(o.SubPositionIdentifier) {
		return true
	}

	return false
}

// SetSubPositionIdentifier gets a reference to the given string and assigns it to the SubPositionIdentifier field.
func (o *CommercePropertiesDto) SetSubPositionIdentifier(v string) {
	o.SubPositionIdentifier = &v
}

// GetCustomerBaseItemNumber returns the CustomerBaseItemNumber field value if set, zero value otherwise.
func (o *CommercePropertiesDto) GetCustomerBaseItemNumber() string {
	if o == nil || IsNil(o.CustomerBaseItemNumber) {
		var ret string
		return ret
	}
	return *o.CustomerBaseItemNumber
}

// GetCustomerBaseItemNumberOk returns a tuple with the CustomerBaseItemNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercePropertiesDto) GetCustomerBaseItemNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerBaseItemNumber) {
		return nil, false
	}
	return o.CustomerBaseItemNumber, true
}

// HasCustomerBaseItemNumber returns a boolean if a field has been set.
func (o *CommercePropertiesDto) HasCustomerBaseItemNumber() bool {
	if o != nil && !IsNil(o.CustomerBaseItemNumber) {
		return true
	}

	return false
}

// SetCustomerBaseItemNumber gets a reference to the given string and assigns it to the CustomerBaseItemNumber field.
func (o *CommercePropertiesDto) SetCustomerBaseItemNumber(v string) {
	o.CustomerBaseItemNumber = &v
}

func (o CommercePropertiesDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommercePropertiesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArticleNumber) {
		toSerialize["articleNumber"] = o.ArticleNumber
	}
	if !IsNil(o.EanGtinArticleNumber) {
		toSerialize["eanGtinArticleNumber"] = o.EanGtinArticleNumber
	}
	if !IsNil(o.IlnArticleNumber) {
		toSerialize["ilnArticleNumber"] = o.IlnArticleNumber
	}
	if !IsNil(o.CatalogueNumber) {
		toSerialize["catalogueNumber"] = o.CatalogueNumber
	}
	if !IsNil(o.CatalogueArticleNumber) {
		toSerialize["catalogueArticleNumber"] = o.CatalogueArticleNumber
	}
	if !IsNil(o.PriceBasis) {
		toSerialize["priceBasis"] = o.PriceBasis
	}
	if !IsNil(o.SubPositionIdentifier) {
		toSerialize["subPositionIdentifier"] = o.SubPositionIdentifier
	}
	if !IsNil(o.CustomerBaseItemNumber) {
		toSerialize["customerBaseItemNumber"] = o.CustomerBaseItemNumber
	}
	return toSerialize, nil
}

type NullableCommercePropertiesDto struct {
	value *CommercePropertiesDto
	isSet bool
}

func (v NullableCommercePropertiesDto) Get() *CommercePropertiesDto {
	return v.value
}

func (v *NullableCommercePropertiesDto) Set(val *CommercePropertiesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCommercePropertiesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCommercePropertiesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommercePropertiesDto(val *CommercePropertiesDto) *NullableCommercePropertiesDto {
	return &NullableCommercePropertiesDto{value: val, isSet: true}
}

func (v NullableCommercePropertiesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommercePropertiesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
