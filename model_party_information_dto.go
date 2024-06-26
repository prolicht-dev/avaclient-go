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

// checks if the PartyInformationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartyInformationDto{}

// PartyInformationDto Represents information about a party (a site or an organization).
type PartyInformationDto struct {
	// Elements GUID identifier.
	Id string `json:"id"`
	// This party's name.
	Name *string `json:"name,omitempty"`
	// This party's street.
	Street *string `json:"street,omitempty"`
	// This party's ZipCode.
	ZipCode *string `json:"zipCode,omitempty"`
	// This party's City.
	City *string `json:"city,omitempty"`
	// This party's Country.
	Country *string `json:"country,omitempty"`
	// This party's Identifier.
	Identifier *string `json:"identifier,omitempty"`
	// Remarks for this party.
	Remarks *string `json:"remarks,omitempty"`
	// An email address for this party.
	Email *string `json:"email,omitempty"`
	// A phone number for this party.
	Phone *string `json:"phone,omitempty"`
	// The name of a contact person.
	ContactPersonName *string `json:"contactPersonName,omitempty"`
	// This is an identifier related to this PartyInformation and their internal reference of the tender (or award). This might be used to assign an identifier (German \"Vergabenummer\") for the current project. This is typically only used in Buyer and Bidder representations and should map to the concept of \"Vergabenummer\" or \"AwardNo\" in GAEB.
	AwardIdentifier *string `json:"awardIdentifier,omitempty"`
	// This property indicates if the party is registered within the European Economic Area.
	IsInEuropeanEconomicArea bool `json:"isInEuropeanEconomicArea"`
	// If this is within the European Economic Area (see IsInEuropeanEconomicArea, then as a business entity it likely has an EU VAT Id.
	VatId *string `json:"vatId,omitempty"`
	// The fax number for this party.
	Fax *string `json:"fax,omitempty"`
	// The two letter ISO country code, e.g. DE for Germany.
	CountryCode *string `json:"countryCode,omitempty"`
	// Depending on which party this class represents, it might have either a 'creditor' or 'debtor' number. This is often used in internal accounting systems.
	CreditorOrDebtorIdentifier *string `json:"creditorOrDebtorIdentifier,omitempty"`
	// The Global Location Number (GLN) is issued by GS1 and is intended to be a unique identifier for the physical address of a party, e.g. a business office.
	GlobalLocationNumber *string `json:"globalLocationNumber,omitempty"`
	// This list contains information about bank accounts associated with this PartyInformation. It's typically used for buyers and bidders.
	BankingInformation []BankingInformationDto `json:"bankingInformation,omitempty"`
	// This can be used to specify a registration number for this party, e.g. a company registration number. It is usually used in GAEB XML commerce exchanges for bidders to supply their WEEE registration number, which is used to identify them as a registered WEEE company.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
}

// NewPartyInformationDto instantiates a new PartyInformationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartyInformationDto(id string, isInEuropeanEconomicArea bool) *PartyInformationDto {
	this := PartyInformationDto{}
	this.Id = id
	this.IsInEuropeanEconomicArea = isInEuropeanEconomicArea
	return &this
}

// NewPartyInformationDtoWithDefaults instantiates a new PartyInformationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyInformationDtoWithDefaults() *PartyInformationDto {
	this := PartyInformationDto{}
	return &this
}

// GetId returns the Id field value
func (o *PartyInformationDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PartyInformationDto) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartyInformationDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartyInformationDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartyInformationDto) SetName(v string) {
	o.Name = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *PartyInformationDto) GetStreet() string {
	if o == nil || IsNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetStreetOk() (*string, bool) {
	if o == nil || IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *PartyInformationDto) HasStreet() bool {
	if o != nil && !IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *PartyInformationDto) SetStreet(v string) {
	o.Street = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *PartyInformationDto) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode) {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *PartyInformationDto) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *PartyInformationDto) SetZipCode(v string) {
	o.ZipCode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *PartyInformationDto) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *PartyInformationDto) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *PartyInformationDto) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PartyInformationDto) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PartyInformationDto) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PartyInformationDto) SetCountry(v string) {
	o.Country = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PartyInformationDto) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PartyInformationDto) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PartyInformationDto) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *PartyInformationDto) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *PartyInformationDto) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *PartyInformationDto) SetRemarks(v string) {
	o.Remarks = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PartyInformationDto) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PartyInformationDto) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PartyInformationDto) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *PartyInformationDto) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *PartyInformationDto) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *PartyInformationDto) SetPhone(v string) {
	o.Phone = &v
}

// GetContactPersonName returns the ContactPersonName field value if set, zero value otherwise.
func (o *PartyInformationDto) GetContactPersonName() string {
	if o == nil || IsNil(o.ContactPersonName) {
		var ret string
		return ret
	}
	return *o.ContactPersonName
}

// GetContactPersonNameOk returns a tuple with the ContactPersonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetContactPersonNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactPersonName) {
		return nil, false
	}
	return o.ContactPersonName, true
}

// HasContactPersonName returns a boolean if a field has been set.
func (o *PartyInformationDto) HasContactPersonName() bool {
	if o != nil && !IsNil(o.ContactPersonName) {
		return true
	}

	return false
}

// SetContactPersonName gets a reference to the given string and assigns it to the ContactPersonName field.
func (o *PartyInformationDto) SetContactPersonName(v string) {
	o.ContactPersonName = &v
}

// GetAwardIdentifier returns the AwardIdentifier field value if set, zero value otherwise.
func (o *PartyInformationDto) GetAwardIdentifier() string {
	if o == nil || IsNil(o.AwardIdentifier) {
		var ret string
		return ret
	}
	return *o.AwardIdentifier
}

// GetAwardIdentifierOk returns a tuple with the AwardIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetAwardIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.AwardIdentifier) {
		return nil, false
	}
	return o.AwardIdentifier, true
}

// HasAwardIdentifier returns a boolean if a field has been set.
func (o *PartyInformationDto) HasAwardIdentifier() bool {
	if o != nil && !IsNil(o.AwardIdentifier) {
		return true
	}

	return false
}

// SetAwardIdentifier gets a reference to the given string and assigns it to the AwardIdentifier field.
func (o *PartyInformationDto) SetAwardIdentifier(v string) {
	o.AwardIdentifier = &v
}

// GetIsInEuropeanEconomicArea returns the IsInEuropeanEconomicArea field value
func (o *PartyInformationDto) GetIsInEuropeanEconomicArea() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInEuropeanEconomicArea
}

// GetIsInEuropeanEconomicAreaOk returns a tuple with the IsInEuropeanEconomicArea field value
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetIsInEuropeanEconomicAreaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInEuropeanEconomicArea, true
}

// SetIsInEuropeanEconomicArea sets field value
func (o *PartyInformationDto) SetIsInEuropeanEconomicArea(v bool) {
	o.IsInEuropeanEconomicArea = v
}

// GetVatId returns the VatId field value if set, zero value otherwise.
func (o *PartyInformationDto) GetVatId() string {
	if o == nil || IsNil(o.VatId) {
		var ret string
		return ret
	}
	return *o.VatId
}

// GetVatIdOk returns a tuple with the VatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetVatIdOk() (*string, bool) {
	if o == nil || IsNil(o.VatId) {
		return nil, false
	}
	return o.VatId, true
}

// HasVatId returns a boolean if a field has been set.
func (o *PartyInformationDto) HasVatId() bool {
	if o != nil && !IsNil(o.VatId) {
		return true
	}

	return false
}

// SetVatId gets a reference to the given string and assigns it to the VatId field.
func (o *PartyInformationDto) SetVatId(v string) {
	o.VatId = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *PartyInformationDto) GetFax() string {
	if o == nil || IsNil(o.Fax) {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetFaxOk() (*string, bool) {
	if o == nil || IsNil(o.Fax) {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *PartyInformationDto) HasFax() bool {
	if o != nil && !IsNil(o.Fax) {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *PartyInformationDto) SetFax(v string) {
	o.Fax = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PartyInformationDto) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PartyInformationDto) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PartyInformationDto) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCreditorOrDebtorIdentifier returns the CreditorOrDebtorIdentifier field value if set, zero value otherwise.
func (o *PartyInformationDto) GetCreditorOrDebtorIdentifier() string {
	if o == nil || IsNil(o.CreditorOrDebtorIdentifier) {
		var ret string
		return ret
	}
	return *o.CreditorOrDebtorIdentifier
}

// GetCreditorOrDebtorIdentifierOk returns a tuple with the CreditorOrDebtorIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetCreditorOrDebtorIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.CreditorOrDebtorIdentifier) {
		return nil, false
	}
	return o.CreditorOrDebtorIdentifier, true
}

// HasCreditorOrDebtorIdentifier returns a boolean if a field has been set.
func (o *PartyInformationDto) HasCreditorOrDebtorIdentifier() bool {
	if o != nil && !IsNil(o.CreditorOrDebtorIdentifier) {
		return true
	}

	return false
}

// SetCreditorOrDebtorIdentifier gets a reference to the given string and assigns it to the CreditorOrDebtorIdentifier field.
func (o *PartyInformationDto) SetCreditorOrDebtorIdentifier(v string) {
	o.CreditorOrDebtorIdentifier = &v
}

// GetGlobalLocationNumber returns the GlobalLocationNumber field value if set, zero value otherwise.
func (o *PartyInformationDto) GetGlobalLocationNumber() string {
	if o == nil || IsNil(o.GlobalLocationNumber) {
		var ret string
		return ret
	}
	return *o.GlobalLocationNumber
}

// GetGlobalLocationNumberOk returns a tuple with the GlobalLocationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetGlobalLocationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalLocationNumber) {
		return nil, false
	}
	return o.GlobalLocationNumber, true
}

// HasGlobalLocationNumber returns a boolean if a field has been set.
func (o *PartyInformationDto) HasGlobalLocationNumber() bool {
	if o != nil && !IsNil(o.GlobalLocationNumber) {
		return true
	}

	return false
}

// SetGlobalLocationNumber gets a reference to the given string and assigns it to the GlobalLocationNumber field.
func (o *PartyInformationDto) SetGlobalLocationNumber(v string) {
	o.GlobalLocationNumber = &v
}

// GetBankingInformation returns the BankingInformation field value if set, zero value otherwise.
func (o *PartyInformationDto) GetBankingInformation() []BankingInformationDto {
	if o == nil || IsNil(o.BankingInformation) {
		var ret []BankingInformationDto
		return ret
	}
	return o.BankingInformation
}

// GetBankingInformationOk returns a tuple with the BankingInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetBankingInformationOk() ([]BankingInformationDto, bool) {
	if o == nil || IsNil(o.BankingInformation) {
		return nil, false
	}
	return o.BankingInformation, true
}

// HasBankingInformation returns a boolean if a field has been set.
func (o *PartyInformationDto) HasBankingInformation() bool {
	if o != nil && !IsNil(o.BankingInformation) {
		return true
	}

	return false
}

// SetBankingInformation gets a reference to the given []BankingInformationDto and assigns it to the BankingInformation field.
func (o *PartyInformationDto) SetBankingInformation(v []BankingInformationDto) {
	o.BankingInformation = v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *PartyInformationDto) GetRegistrationNumber() string {
	if o == nil || IsNil(o.RegistrationNumber) {
		var ret string
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyInformationDto) GetRegistrationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationNumber) {
		return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *PartyInformationDto) HasRegistrationNumber() bool {
	if o != nil && !IsNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given string and assigns it to the RegistrationNumber field.
func (o *PartyInformationDto) SetRegistrationNumber(v string) {
	o.RegistrationNumber = &v
}

func (o PartyInformationDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartyInformationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zipCode"] = o.ZipCode
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.ContactPersonName) {
		toSerialize["contactPersonName"] = o.ContactPersonName
	}
	if !IsNil(o.AwardIdentifier) {
		toSerialize["awardIdentifier"] = o.AwardIdentifier
	}
	toSerialize["isInEuropeanEconomicArea"] = o.IsInEuropeanEconomicArea
	if !IsNil(o.VatId) {
		toSerialize["vatId"] = o.VatId
	}
	if !IsNil(o.Fax) {
		toSerialize["fax"] = o.Fax
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.CreditorOrDebtorIdentifier) {
		toSerialize["creditorOrDebtorIdentifier"] = o.CreditorOrDebtorIdentifier
	}
	if !IsNil(o.GlobalLocationNumber) {
		toSerialize["globalLocationNumber"] = o.GlobalLocationNumber
	}
	if !IsNil(o.BankingInformation) {
		toSerialize["bankingInformation"] = o.BankingInformation
	}
	if !IsNil(o.RegistrationNumber) {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	return toSerialize, nil
}

type NullablePartyInformationDto struct {
	value *PartyInformationDto
	isSet bool
}

func (v NullablePartyInformationDto) Get() *PartyInformationDto {
	return v.value
}

func (v *NullablePartyInformationDto) Set(val *PartyInformationDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyInformationDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyInformationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyInformationDto(val *PartyInformationDto) *NullablePartyInformationDto {
	return &NullablePartyInformationDto{value: val, isSet: true}
}

func (v NullablePartyInformationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyInformationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
