/*
AVACloud API 1.45.0

AVACloud API specification

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package avaclient

import (
	"encoding/json"
)

// checks if the ArticleDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArticleDataDto{}

// ArticleDataDto This class represents a single article, usually used within ProductData
type ArticleDataDto struct {
	// Elements GUID identifier.
	Id string `json:"id"`
	// The name (or brand name) for this article, usually given by the supplier or vendor.
	Name *string `json:"name,omitempty"`
	// An article number that describes it, useful when integrating other systems.
	ArticleNumber *string `json:"articleNumber,omitempty"`
	// Quantity for this article. If this is used within a Position, the quantity here should be the quantity required for the full quantity of the position, not for a single unit.
	Quantity float32 `json:"quantity"`
	// The unit tag for this single ArticleData.
	UnitTag *string `json:"unitTag,omitempty"`
	// This is an optional text element that can be used to further describe the ArticleData.
	Description *string `json:"description,omitempty"`
	// Short description for this ITextElement element.
	ShortText *string `json:"shortText,omitempty"`
	// Detailed description for this ITextElement element. When the HtmlLongText is set, this is automatically overwritten and filled with the appropriate plain text representation of the Html text. Vice versa, setting this property overrides the HtmlLongText.
	LongText *string `json:"longText,omitempty"`
	// This contains the Html representation of the Longtext. When the LongText is set, this is automatically overwritten and filled with the appropriate Html representation of the plaintext. Vice versa, setting this property overrides the LongText. GAEB 90 and GAEB 2000 exports do not support any image functionality. In GAEB XML, only images that use an embedded Base64 data uri are exported, regular url references are cleared before written out.
	HtmlLongText *string `json:"htmlLongText,omitempty"`
}

// NewArticleDataDto instantiates a new ArticleDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleDataDto(id string, quantity float32) *ArticleDataDto {
	this := ArticleDataDto{}
	this.Id = id
	this.Quantity = quantity
	return &this
}

// NewArticleDataDtoWithDefaults instantiates a new ArticleDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleDataDtoWithDefaults() *ArticleDataDto {
	this := ArticleDataDto{}
	return &this
}

// GetId returns the Id field value
func (o *ArticleDataDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArticleDataDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArticleDataDto) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ArticleDataDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleDataDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ArticleDataDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ArticleDataDto) SetName(v string) {
	o.Name = &v
}

// GetArticleNumber returns the ArticleNumber field value if set, zero value otherwise.
func (o *ArticleDataDto) GetArticleNumber() string {
	if o == nil || IsNil(o.ArticleNumber) {
		var ret string
		return ret
	}
	return *o.ArticleNumber
}

// GetArticleNumberOk returns a tuple with the ArticleNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleDataDto) GetArticleNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ArticleNumber) {
		return nil, false
	}
	return o.ArticleNumber, true
}

// HasArticleNumber returns a boolean if a field has been set.
func (o *ArticleDataDto) HasArticleNumber() bool {
	if o != nil && !IsNil(o.ArticleNumber) {
		return true
	}

	return false
}

// SetArticleNumber gets a reference to the given string and assigns it to the ArticleNumber field.
func (o *ArticleDataDto) SetArticleNumber(v string) {
	o.ArticleNumber = &v
}

// GetQuantity returns the Quantity field value
func (o *ArticleDataDto) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ArticleDataDto) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ArticleDataDto) SetQuantity(v float32) {
	o.Quantity = v
}

// GetUnitTag returns the UnitTag field value if set, zero value otherwise.
func (o *ArticleDataDto) GetUnitTag() string {
	if o == nil || IsNil(o.UnitTag) {
		var ret string
		return ret
	}
	return *o.UnitTag
}

// GetUnitTagOk returns a tuple with the UnitTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleDataDto) GetUnitTagOk() (*string, bool) {
	if o == nil || IsNil(o.UnitTag) {
		return nil, false
	}
	return o.UnitTag, true
}

// HasUnitTag returns a boolean if a field has been set.
func (o *ArticleDataDto) HasUnitTag() bool {
	if o != nil && !IsNil(o.UnitTag) {
		return true
	}

	return false
}

// SetUnitTag gets a reference to the given string and assigns it to the UnitTag field.
func (o *ArticleDataDto) SetUnitTag(v string) {
	o.UnitTag = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArticleDataDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleDataDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArticleDataDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArticleDataDto) SetDescription(v string) {
	o.Description = &v
}

// GetShortText returns the ShortText field value if set, zero value otherwise.
func (o *ArticleDataDto) GetShortText() string {
	if o == nil || IsNil(o.ShortText) {
		var ret string
		return ret
	}
	return *o.ShortText
}

// GetShortTextOk returns a tuple with the ShortText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleDataDto) GetShortTextOk() (*string, bool) {
	if o == nil || IsNil(o.ShortText) {
		return nil, false
	}
	return o.ShortText, true
}

// HasShortText returns a boolean if a field has been set.
func (o *ArticleDataDto) HasShortText() bool {
	if o != nil && !IsNil(o.ShortText) {
		return true
	}

	return false
}

// SetShortText gets a reference to the given string and assigns it to the ShortText field.
func (o *ArticleDataDto) SetShortText(v string) {
	o.ShortText = &v
}

// GetLongText returns the LongText field value if set, zero value otherwise.
func (o *ArticleDataDto) GetLongText() string {
	if o == nil || IsNil(o.LongText) {
		var ret string
		return ret
	}
	return *o.LongText
}

// GetLongTextOk returns a tuple with the LongText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleDataDto) GetLongTextOk() (*string, bool) {
	if o == nil || IsNil(o.LongText) {
		return nil, false
	}
	return o.LongText, true
}

// HasLongText returns a boolean if a field has been set.
func (o *ArticleDataDto) HasLongText() bool {
	if o != nil && !IsNil(o.LongText) {
		return true
	}

	return false
}

// SetLongText gets a reference to the given string and assigns it to the LongText field.
func (o *ArticleDataDto) SetLongText(v string) {
	o.LongText = &v
}

// GetHtmlLongText returns the HtmlLongText field value if set, zero value otherwise.
func (o *ArticleDataDto) GetHtmlLongText() string {
	if o == nil || IsNil(o.HtmlLongText) {
		var ret string
		return ret
	}
	return *o.HtmlLongText
}

// GetHtmlLongTextOk returns a tuple with the HtmlLongText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleDataDto) GetHtmlLongTextOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlLongText) {
		return nil, false
	}
	return o.HtmlLongText, true
}

// HasHtmlLongText returns a boolean if a field has been set.
func (o *ArticleDataDto) HasHtmlLongText() bool {
	if o != nil && !IsNil(o.HtmlLongText) {
		return true
	}

	return false
}

// SetHtmlLongText gets a reference to the given string and assigns it to the HtmlLongText field.
func (o *ArticleDataDto) SetHtmlLongText(v string) {
	o.HtmlLongText = &v
}

func (o ArticleDataDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArticleDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ArticleNumber) {
		toSerialize["articleNumber"] = o.ArticleNumber
	}
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.UnitTag) {
		toSerialize["unitTag"] = o.UnitTag
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

type NullableArticleDataDto struct {
	value *ArticleDataDto
	isSet bool
}

func (v NullableArticleDataDto) Get() *ArticleDataDto {
	return v.value
}

func (v *NullableArticleDataDto) Set(val *ArticleDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleDataDto(val *ArticleDataDto) *NullableArticleDataDto {
	return &NullableArticleDataDto{value: val, isSet: true}
}

func (v NullableArticleDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
