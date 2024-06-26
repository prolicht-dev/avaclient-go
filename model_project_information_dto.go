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

// checks if the ProjectInformationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectInformationDto{}

// ProjectInformationDto General information about a Project.
type ProjectInformationDto struct {
	Buyer *PartyInformationDto `json:"buyer,omitempty"`
	// Description for the project.
	Description *string `json:"description,omitempty"`
	// Short description for the project.
	DescriptionShort *string `json:"descriptionShort,omitempty"`
	// Name of the project.
	Name             *string              `json:"name,omitempty"`
	Site             *PartyInformationDto `json:"site,omitempty"`
	ItemNumberSchema *ItemNumberSchemaDto `json:"itemNumberSchema,omitempty"`
	// Short label for the currency used.
	CurrencyShort *string `json:"currencyShort,omitempty"`
	// Full label of the currency used.
	CurrencyLong *string `json:"currencyLong,omitempty"`
	// Label for the labour time part of prices used in the project.
	LabourTimeLabel *string `json:"labourTimeLabel,omitempty"`
	// Labels for the price components used in the project. Caution: Removal of a price component will trigger to have dependent price informations be deleted in IElements that use this property. If this property is changed or altered after the project has already been used, it is strongly advised to do operations step by step, e.g. if renaming and reordering multiple price components, this should be done one by one. Otherwise, a combination of rename and reordering will not be correctly propagated downwards to child objects in this Project.
	PriceComponents []string `json:"priceComponents,omitempty"`
	// This dictionary specifies actual types used for the PriceComponents. For example, a single price component might have the name 'Material' as a string, ans this dictionary would then have a key 'Material' and a value of Materials. You can not add keys here that are not also present in the PriceComponents property, and removing price components will also remove them from this dictionary here.
	PriceComponentTypes *map[string]PriceComponentTypeDto `json:"priceComponentTypes,omitempty"`
	// This bool indicates if this project allows the bidder to add bidder comments. Bidder comments are a way to attach clarifying information when submitting an offer.
	BidderCommentAllowed bool `json:"bidderCommentAllowed"`
	// This property indicates if the project should allow side offers from contractors. In GAEB, a side offer would typically be in exchange phase 85.
	SideOffersAllowed bool                `json:"sideOffersAllowed"`
	AwardType         AwardTypeDto        `json:"awardType"`
	SpecialAwardKind  SpecialAwardKindDto `json:"specialAwardKind"`
	// Requesters in a construction project, in German also called 'Bedarfsträger', are parties connected to the building process, e.g. architects or planners.
	Requesters []PartyInformationDto `json:"requesters,omitempty"`
	// Notification sites are addresses that are used for delivering messages in the context of construction projects.
	NotificationSites []PartyInformationDto `json:"notificationSites,omitempty"`
}

// NewProjectInformationDto instantiates a new ProjectInformationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectInformationDto(bidderCommentAllowed bool, sideOffersAllowed bool, awardType AwardTypeDto, specialAwardKind SpecialAwardKindDto) *ProjectInformationDto {
	this := ProjectInformationDto{}
	this.BidderCommentAllowed = bidderCommentAllowed
	this.SideOffersAllowed = sideOffersAllowed
	this.AwardType = awardType
	this.SpecialAwardKind = specialAwardKind
	return &this
}

// NewProjectInformationDtoWithDefaults instantiates a new ProjectInformationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectInformationDtoWithDefaults() *ProjectInformationDto {
	this := ProjectInformationDto{}
	return &this
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetBuyer() PartyInformationDto {
	if o == nil || IsNil(o.Buyer) {
		var ret PartyInformationDto
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetBuyerOk() (*PartyInformationDto, bool) {
	if o == nil || IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasBuyer() bool {
	if o != nil && !IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given PartyInformationDto and assigns it to the Buyer field.
func (o *ProjectInformationDto) SetBuyer(v PartyInformationDto) {
	o.Buyer = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectInformationDto) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionShort returns the DescriptionShort field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetDescriptionShort() string {
	if o == nil || IsNil(o.DescriptionShort) {
		var ret string
		return ret
	}
	return *o.DescriptionShort
}

// GetDescriptionShortOk returns a tuple with the DescriptionShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetDescriptionShortOk() (*string, bool) {
	if o == nil || IsNil(o.DescriptionShort) {
		return nil, false
	}
	return o.DescriptionShort, true
}

// HasDescriptionShort returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasDescriptionShort() bool {
	if o != nil && !IsNil(o.DescriptionShort) {
		return true
	}

	return false
}

// SetDescriptionShort gets a reference to the given string and assigns it to the DescriptionShort field.
func (o *ProjectInformationDto) SetDescriptionShort(v string) {
	o.DescriptionShort = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectInformationDto) SetName(v string) {
	o.Name = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetSite() PartyInformationDto {
	if o == nil || IsNil(o.Site) {
		var ret PartyInformationDto
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetSiteOk() (*PartyInformationDto, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given PartyInformationDto and assigns it to the Site field.
func (o *ProjectInformationDto) SetSite(v PartyInformationDto) {
	o.Site = &v
}

// GetItemNumberSchema returns the ItemNumberSchema field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetItemNumberSchema() ItemNumberSchemaDto {
	if o == nil || IsNil(o.ItemNumberSchema) {
		var ret ItemNumberSchemaDto
		return ret
	}
	return *o.ItemNumberSchema
}

// GetItemNumberSchemaOk returns a tuple with the ItemNumberSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetItemNumberSchemaOk() (*ItemNumberSchemaDto, bool) {
	if o == nil || IsNil(o.ItemNumberSchema) {
		return nil, false
	}
	return o.ItemNumberSchema, true
}

// HasItemNumberSchema returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasItemNumberSchema() bool {
	if o != nil && !IsNil(o.ItemNumberSchema) {
		return true
	}

	return false
}

// SetItemNumberSchema gets a reference to the given ItemNumberSchemaDto and assigns it to the ItemNumberSchema field.
func (o *ProjectInformationDto) SetItemNumberSchema(v ItemNumberSchemaDto) {
	o.ItemNumberSchema = &v
}

// GetCurrencyShort returns the CurrencyShort field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetCurrencyShort() string {
	if o == nil || IsNil(o.CurrencyShort) {
		var ret string
		return ret
	}
	return *o.CurrencyShort
}

// GetCurrencyShortOk returns a tuple with the CurrencyShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetCurrencyShortOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyShort) {
		return nil, false
	}
	return o.CurrencyShort, true
}

// HasCurrencyShort returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasCurrencyShort() bool {
	if o != nil && !IsNil(o.CurrencyShort) {
		return true
	}

	return false
}

// SetCurrencyShort gets a reference to the given string and assigns it to the CurrencyShort field.
func (o *ProjectInformationDto) SetCurrencyShort(v string) {
	o.CurrencyShort = &v
}

// GetCurrencyLong returns the CurrencyLong field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetCurrencyLong() string {
	if o == nil || IsNil(o.CurrencyLong) {
		var ret string
		return ret
	}
	return *o.CurrencyLong
}

// GetCurrencyLongOk returns a tuple with the CurrencyLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetCurrencyLongOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyLong) {
		return nil, false
	}
	return o.CurrencyLong, true
}

// HasCurrencyLong returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasCurrencyLong() bool {
	if o != nil && !IsNil(o.CurrencyLong) {
		return true
	}

	return false
}

// SetCurrencyLong gets a reference to the given string and assigns it to the CurrencyLong field.
func (o *ProjectInformationDto) SetCurrencyLong(v string) {
	o.CurrencyLong = &v
}

// GetLabourTimeLabel returns the LabourTimeLabel field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetLabourTimeLabel() string {
	if o == nil || IsNil(o.LabourTimeLabel) {
		var ret string
		return ret
	}
	return *o.LabourTimeLabel
}

// GetLabourTimeLabelOk returns a tuple with the LabourTimeLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetLabourTimeLabelOk() (*string, bool) {
	if o == nil || IsNil(o.LabourTimeLabel) {
		return nil, false
	}
	return o.LabourTimeLabel, true
}

// HasLabourTimeLabel returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasLabourTimeLabel() bool {
	if o != nil && !IsNil(o.LabourTimeLabel) {
		return true
	}

	return false
}

// SetLabourTimeLabel gets a reference to the given string and assigns it to the LabourTimeLabel field.
func (o *ProjectInformationDto) SetLabourTimeLabel(v string) {
	o.LabourTimeLabel = &v
}

// GetPriceComponents returns the PriceComponents field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetPriceComponents() []string {
	if o == nil || IsNil(o.PriceComponents) {
		var ret []string
		return ret
	}
	return o.PriceComponents
}

// GetPriceComponentsOk returns a tuple with the PriceComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetPriceComponentsOk() ([]string, bool) {
	if o == nil || IsNil(o.PriceComponents) {
		return nil, false
	}
	return o.PriceComponents, true
}

// HasPriceComponents returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasPriceComponents() bool {
	if o != nil && !IsNil(o.PriceComponents) {
		return true
	}

	return false
}

// SetPriceComponents gets a reference to the given []string and assigns it to the PriceComponents field.
func (o *ProjectInformationDto) SetPriceComponents(v []string) {
	o.PriceComponents = v
}

// GetPriceComponentTypes returns the PriceComponentTypes field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetPriceComponentTypes() map[string]PriceComponentTypeDto {
	if o == nil || IsNil(o.PriceComponentTypes) {
		var ret map[string]PriceComponentTypeDto
		return ret
	}
	return *o.PriceComponentTypes
}

// GetPriceComponentTypesOk returns a tuple with the PriceComponentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetPriceComponentTypesOk() (*map[string]PriceComponentTypeDto, bool) {
	if o == nil || IsNil(o.PriceComponentTypes) {
		return nil, false
	}
	return o.PriceComponentTypes, true
}

// HasPriceComponentTypes returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasPriceComponentTypes() bool {
	if o != nil && !IsNil(o.PriceComponentTypes) {
		return true
	}

	return false
}

// SetPriceComponentTypes gets a reference to the given map[string]PriceComponentTypeDto and assigns it to the PriceComponentTypes field.
func (o *ProjectInformationDto) SetPriceComponentTypes(v map[string]PriceComponentTypeDto) {
	o.PriceComponentTypes = &v
}

// GetBidderCommentAllowed returns the BidderCommentAllowed field value
func (o *ProjectInformationDto) GetBidderCommentAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BidderCommentAllowed
}

// GetBidderCommentAllowedOk returns a tuple with the BidderCommentAllowed field value
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetBidderCommentAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BidderCommentAllowed, true
}

// SetBidderCommentAllowed sets field value
func (o *ProjectInformationDto) SetBidderCommentAllowed(v bool) {
	o.BidderCommentAllowed = v
}

// GetSideOffersAllowed returns the SideOffersAllowed field value
func (o *ProjectInformationDto) GetSideOffersAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SideOffersAllowed
}

// GetSideOffersAllowedOk returns a tuple with the SideOffersAllowed field value
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetSideOffersAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SideOffersAllowed, true
}

// SetSideOffersAllowed sets field value
func (o *ProjectInformationDto) SetSideOffersAllowed(v bool) {
	o.SideOffersAllowed = v
}

// GetAwardType returns the AwardType field value
func (o *ProjectInformationDto) GetAwardType() AwardTypeDto {
	if o == nil {
		var ret AwardTypeDto
		return ret
	}

	return o.AwardType
}

// GetAwardTypeOk returns a tuple with the AwardType field value
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetAwardTypeOk() (*AwardTypeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwardType, true
}

// SetAwardType sets field value
func (o *ProjectInformationDto) SetAwardType(v AwardTypeDto) {
	o.AwardType = v
}

// GetSpecialAwardKind returns the SpecialAwardKind field value
func (o *ProjectInformationDto) GetSpecialAwardKind() SpecialAwardKindDto {
	if o == nil {
		var ret SpecialAwardKindDto
		return ret
	}

	return o.SpecialAwardKind
}

// GetSpecialAwardKindOk returns a tuple with the SpecialAwardKind field value
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetSpecialAwardKindOk() (*SpecialAwardKindDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecialAwardKind, true
}

// SetSpecialAwardKind sets field value
func (o *ProjectInformationDto) SetSpecialAwardKind(v SpecialAwardKindDto) {
	o.SpecialAwardKind = v
}

// GetRequesters returns the Requesters field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetRequesters() []PartyInformationDto {
	if o == nil || IsNil(o.Requesters) {
		var ret []PartyInformationDto
		return ret
	}
	return o.Requesters
}

// GetRequestersOk returns a tuple with the Requesters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetRequestersOk() ([]PartyInformationDto, bool) {
	if o == nil || IsNil(o.Requesters) {
		return nil, false
	}
	return o.Requesters, true
}

// HasRequesters returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasRequesters() bool {
	if o != nil && !IsNil(o.Requesters) {
		return true
	}

	return false
}

// SetRequesters gets a reference to the given []PartyInformationDto and assigns it to the Requesters field.
func (o *ProjectInformationDto) SetRequesters(v []PartyInformationDto) {
	o.Requesters = v
}

// GetNotificationSites returns the NotificationSites field value if set, zero value otherwise.
func (o *ProjectInformationDto) GetNotificationSites() []PartyInformationDto {
	if o == nil || IsNil(o.NotificationSites) {
		var ret []PartyInformationDto
		return ret
	}
	return o.NotificationSites
}

// GetNotificationSitesOk returns a tuple with the NotificationSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectInformationDto) GetNotificationSitesOk() ([]PartyInformationDto, bool) {
	if o == nil || IsNil(o.NotificationSites) {
		return nil, false
	}
	return o.NotificationSites, true
}

// HasNotificationSites returns a boolean if a field has been set.
func (o *ProjectInformationDto) HasNotificationSites() bool {
	if o != nil && !IsNil(o.NotificationSites) {
		return true
	}

	return false
}

// SetNotificationSites gets a reference to the given []PartyInformationDto and assigns it to the NotificationSites field.
func (o *ProjectInformationDto) SetNotificationSites(v []PartyInformationDto) {
	o.NotificationSites = v
}

func (o ProjectInformationDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectInformationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Buyer) {
		toSerialize["buyer"] = o.Buyer
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DescriptionShort) {
		toSerialize["descriptionShort"] = o.DescriptionShort
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.ItemNumberSchema) {
		toSerialize["itemNumberSchema"] = o.ItemNumberSchema
	}
	if !IsNil(o.CurrencyShort) {
		toSerialize["currencyShort"] = o.CurrencyShort
	}
	if !IsNil(o.CurrencyLong) {
		toSerialize["currencyLong"] = o.CurrencyLong
	}
	if !IsNil(o.LabourTimeLabel) {
		toSerialize["labourTimeLabel"] = o.LabourTimeLabel
	}
	if !IsNil(o.PriceComponents) {
		toSerialize["priceComponents"] = o.PriceComponents
	}
	// skip: priceComponentTypes is readOnly
	toSerialize["bidderCommentAllowed"] = o.BidderCommentAllowed
	toSerialize["sideOffersAllowed"] = o.SideOffersAllowed
	toSerialize["awardType"] = o.AwardType
	toSerialize["specialAwardKind"] = o.SpecialAwardKind
	if !IsNil(o.Requesters) {
		toSerialize["requesters"] = o.Requesters
	}
	if !IsNil(o.NotificationSites) {
		toSerialize["notificationSites"] = o.NotificationSites
	}
	return toSerialize, nil
}

type NullableProjectInformationDto struct {
	value *ProjectInformationDto
	isSet bool
}

func (v NullableProjectInformationDto) Get() *ProjectInformationDto {
	return v.value
}

func (v *NullableProjectInformationDto) Set(val *ProjectInformationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectInformationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectInformationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectInformationDto(val *ProjectInformationDto) *NullableProjectInformationDto {
	return &NullableProjectInformationDto{value: val, isSet: true}
}

func (v NullableProjectInformationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectInformationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
