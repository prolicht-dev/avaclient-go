# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | BT-46 (BT-29 for the Seller) The identifier of the organization | [optional] 
**IdentifierScheme** | Pointer to **string** | BT-46 (BT-29 for the Seller) If a Identifier is used, this indicates the scheme used for the identifier. This works in conjunction with the Identifier property. | [optional] 
**Name** | Pointer to **string** | BT-44 (BT-27 for the seller, BT-62 for the seller tax representative) The name of the organization | [optional] 
**RegistrationNumber** | Pointer to **string** | BT-30, for the seller, doesn&#39;t exist in the buyer actually A registration number, e.g. a Handelsregisternummer in Germany | [optional] 
**VatId** | Pointer to **string** | BT-48, BT-31 for the seller, BT-63 for the seller tax representative The VAT ID for the organization | [optional] 
**TaxId** | Pointer to **string** | BT-32 for the seller, does not exist for the buyer The tax ID for the organization | [optional] 
**Email** | Pointer to **string** | BT-49 An email address for the organization | [optional] 
**ContactPersonName** | Pointer to **string** | BT-56 The name of a contact person | [optional] 
**ContactPersonPhone** | Pointer to **string** | BT-57 The phone of a contact person | [optional] 
**ContactPersonEmail** | Pointer to **string** | BT-58 The email of a contact person | [optional] 
**Address** | Pointer to **string** | BT-50, BT-35 for the seller, BT-64 for the seller tax representative BT-51, BT-36 for the seller, BT-65 for the seller tax representative BT-163, BT-162 for the seller, BT-164 for the seller tax representative The address of the organization | [optional] 
**City** | Pointer to **string** | BT-52, BT-37 for the seller, BT-66 for the seller tax representative The city of the organization | [optional] 
**ZipCode** | Pointer to **string** | BT-53, BT-38 for the seller, BT-67 for the seller tax representative The zip code of the organization | [optional] 
**CountryCode** | Pointer to **string** | BT-55, BT-40 for the seller, BT-69 for the seller tax representative The ISO 3166-1 2 letter country code | [optional] 
**State** | Pointer to **string** | BT-54, BT-39 for the seller, BT-68 for the seller tax representative The state of the organization, e.g. a federal state in Germany | [optional] 
**AdditionalLegalInformation** | Pointer to **string** | BT-33 This provides some additional legal information. BT-33 in XRechnung, and this is only used for the seller. | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *Organization) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Organization) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Organization) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Organization) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIdentifierScheme

`func (o *Organization) GetIdentifierScheme() string`

GetIdentifierScheme returns the IdentifierScheme field if non-nil, zero value otherwise.

### GetIdentifierSchemeOk

`func (o *Organization) GetIdentifierSchemeOk() (*string, bool)`

GetIdentifierSchemeOk returns a tuple with the IdentifierScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierScheme

`func (o *Organization) SetIdentifierScheme(v string)`

SetIdentifierScheme sets IdentifierScheme field to given value.

### HasIdentifierScheme

`func (o *Organization) HasIdentifierScheme() bool`

HasIdentifierScheme returns a boolean if a field has been set.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organization) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *Organization) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *Organization) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *Organization) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *Organization) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### GetVatId

`func (o *Organization) GetVatId() string`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *Organization) GetVatIdOk() (*string, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *Organization) SetVatId(v string)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *Organization) HasVatId() bool`

HasVatId returns a boolean if a field has been set.

### GetTaxId

`func (o *Organization) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *Organization) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *Organization) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *Organization) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetEmail

`func (o *Organization) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Organization) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Organization) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Organization) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetContactPersonName

`func (o *Organization) GetContactPersonName() string`

GetContactPersonName returns the ContactPersonName field if non-nil, zero value otherwise.

### GetContactPersonNameOk

`func (o *Organization) GetContactPersonNameOk() (*string, bool)`

GetContactPersonNameOk returns a tuple with the ContactPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonName

`func (o *Organization) SetContactPersonName(v string)`

SetContactPersonName sets ContactPersonName field to given value.

### HasContactPersonName

`func (o *Organization) HasContactPersonName() bool`

HasContactPersonName returns a boolean if a field has been set.

### GetContactPersonPhone

`func (o *Organization) GetContactPersonPhone() string`

GetContactPersonPhone returns the ContactPersonPhone field if non-nil, zero value otherwise.

### GetContactPersonPhoneOk

`func (o *Organization) GetContactPersonPhoneOk() (*string, bool)`

GetContactPersonPhoneOk returns a tuple with the ContactPersonPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonPhone

`func (o *Organization) SetContactPersonPhone(v string)`

SetContactPersonPhone sets ContactPersonPhone field to given value.

### HasContactPersonPhone

`func (o *Organization) HasContactPersonPhone() bool`

HasContactPersonPhone returns a boolean if a field has been set.

### GetContactPersonEmail

`func (o *Organization) GetContactPersonEmail() string`

GetContactPersonEmail returns the ContactPersonEmail field if non-nil, zero value otherwise.

### GetContactPersonEmailOk

`func (o *Organization) GetContactPersonEmailOk() (*string, bool)`

GetContactPersonEmailOk returns a tuple with the ContactPersonEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonEmail

`func (o *Organization) SetContactPersonEmail(v string)`

SetContactPersonEmail sets ContactPersonEmail field to given value.

### HasContactPersonEmail

`func (o *Organization) HasContactPersonEmail() bool`

HasContactPersonEmail returns a boolean if a field has been set.

### GetAddress

`func (o *Organization) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Organization) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Organization) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Organization) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *Organization) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Organization) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Organization) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Organization) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZipCode

`func (o *Organization) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Organization) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Organization) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *Organization) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *Organization) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Organization) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Organization) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Organization) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetState

`func (o *Organization) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Organization) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Organization) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Organization) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAdditionalLegalInformation

`func (o *Organization) GetAdditionalLegalInformation() string`

GetAdditionalLegalInformation returns the AdditionalLegalInformation field if non-nil, zero value otherwise.

### GetAdditionalLegalInformationOk

`func (o *Organization) GetAdditionalLegalInformationOk() (*string, bool)`

GetAdditionalLegalInformationOk returns a tuple with the AdditionalLegalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalLegalInformation

`func (o *Organization) SetAdditionalLegalInformation(v string)`

SetAdditionalLegalInformation sets AdditionalLegalInformation field to given value.

### HasAdditionalLegalInformation

`func (o *Organization) HasAdditionalLegalInformation() bool`

HasAdditionalLegalInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


