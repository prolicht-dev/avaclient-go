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

// checks if the ProductDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDataDto{}

// ProductDataDto This class represents data about products and their vendor
type ProductDataDto struct {
	// Elements GUID identifier.
	Id     string               `json:"id"`
	Vendor *PartyInformationDto `json:"vendor,omitempty"`
	// The collection of ArticleData that describe this product, e.g. a pipe product could be composed out of multiple pipe segments and fittings.
	Articles []ArticleDataDto `json:"articles,omitempty"`
	// Short description for this ITextElement element.
	ShortText *string `json:"shortText,omitempty"`
	// Detailed description for this ITextElement element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText.
	LongText *string `json:"longText,omitempty"`
	// This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out.
	HtmlLongText *string `json:"htmlLongText,omitempty"`
}

// NewProductDataDto instantiates a new ProductDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDataDto(id string) *ProductDataDto {
	this := ProductDataDto{}
	this.Id = id
	return &this
}

// NewProductDataDtoWithDefaults instantiates a new ProductDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDataDtoWithDefaults() *ProductDataDto {
	this := ProductDataDto{}
	return &this
}

// GetId returns the Id field value
func (o *ProductDataDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductDataDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductDataDto) SetId(v string) {
	o.Id = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ProductDataDto) GetVendor() PartyInformationDto {
	if o == nil || IsNil(o.Vendor) {
		var ret PartyInformationDto
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataDto) GetVendorOk() (*PartyInformationDto, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ProductDataDto) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given PartyInformationDto and assigns it to the Vendor field.
func (o *ProductDataDto) SetVendor(v PartyInformationDto) {
	o.Vendor = &v
}

// GetArticles returns the Articles field value if set, zero value otherwise.
func (o *ProductDataDto) GetArticles() []ArticleDataDto {
	if o == nil || IsNil(o.Articles) {
		var ret []ArticleDataDto
		return ret
	}
	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataDto) GetArticlesOk() ([]ArticleDataDto, bool) {
	if o == nil || IsNil(o.Articles) {
		return nil, false
	}
	return o.Articles, true
}

// HasArticles returns a boolean if a field has been set.
func (o *ProductDataDto) HasArticles() bool {
	if o != nil && !IsNil(o.Articles) {
		return true
	}

	return false
}

// SetArticles gets a reference to the given []ArticleDataDto and assigns it to the Articles field.
func (o *ProductDataDto) SetArticles(v []ArticleDataDto) {
	o.Articles = v
}

// GetShortText returns the ShortText field value if set, zero value otherwise.
func (o *ProductDataDto) GetShortText() string {
	if o == nil || IsNil(o.ShortText) {
		var ret string
		return ret
	}
	return *o.ShortText
}

// GetShortTextOk returns a tuple with the ShortText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataDto) GetShortTextOk() (*string, bool) {
	if o == nil || IsNil(o.ShortText) {
		return nil, false
	}
	return o.ShortText, true
}

// HasShortText returns a boolean if a field has been set.
func (o *ProductDataDto) HasShortText() bool {
	if o != nil && !IsNil(o.ShortText) {
		return true
	}

	return false
}

// SetShortText gets a reference to the given string and assigns it to the ShortText field.
func (o *ProductDataDto) SetShortText(v string) {
	o.ShortText = &v
}

// GetLongText returns the LongText field value if set, zero value otherwise.
func (o *ProductDataDto) GetLongText() string {
	if o == nil || IsNil(o.LongText) {
		var ret string
		return ret
	}
	return *o.LongText
}

// GetLongTextOk returns a tuple with the LongText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataDto) GetLongTextOk() (*string, bool) {
	if o == nil || IsNil(o.LongText) {
		return nil, false
	}
	return o.LongText, true
}

// HasLongText returns a boolean if a field has been set.
func (o *ProductDataDto) HasLongText() bool {
	if o != nil && !IsNil(o.LongText) {
		return true
	}

	return false
}

// SetLongText gets a reference to the given string and assigns it to the LongText field.
func (o *ProductDataDto) SetLongText(v string) {
	o.LongText = &v
}

// GetHtmlLongText returns the HtmlLongText field value if set, zero value otherwise.
func (o *ProductDataDto) GetHtmlLongText() string {
	if o == nil || IsNil(o.HtmlLongText) {
		var ret string
		return ret
	}
	return *o.HtmlLongText
}

// GetHtmlLongTextOk returns a tuple with the HtmlLongText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDataDto) GetHtmlLongTextOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlLongText) {
		return nil, false
	}
	return o.HtmlLongText, true
}

// HasHtmlLongText returns a boolean if a field has been set.
func (o *ProductDataDto) HasHtmlLongText() bool {
	if o != nil && !IsNil(o.HtmlLongText) {
		return true
	}

	return false
}

// SetHtmlLongText gets a reference to the given string and assigns it to the HtmlLongText field.
func (o *ProductDataDto) SetHtmlLongText(v string) {
	o.HtmlLongText = &v
}

func (o ProductDataDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.Articles) {
		toSerialize["articles"] = o.Articles
	}
	if !IsNil(o.ShortText) {
		toSerialize["shortText"] = o.ShortText
	}
	if !IsNil(o.LongText) {
		toSerialize["longText"] = o.LongText
	}
	if !IsNil(o.HtmlLongText) {
		toSerialize["htmlLongText"] = o.HtmlLongText
	}
	return toSerialize, nil
}

type NullableProductDataDto struct {
	value *ProductDataDto
	isSet bool
}

func (v NullableProductDataDto) Get() *ProductDataDto {
	return v.value
}

func (v *NullableProductDataDto) Set(val *ProductDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDataDto(val *ProductDataDto) *NullableProductDataDto {
	return &NullableProductDataDto{value: val, isSet: true}
}

func (v NullableProductDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
