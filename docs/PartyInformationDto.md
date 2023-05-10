# PartyInformationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Name** | Pointer to **string** | This party&#39;s name. | [optional] 
**Street** | Pointer to **string** | This party&#39;s street. | [optional] 
**ZipCode** | Pointer to **string** | This party&#39;s ZipCode. | [optional] 
**City** | Pointer to **string** | This party&#39;s City. | [optional] 
**Country** | Pointer to **string** | This party&#39;s Country. | [optional] 
**Identifier** | Pointer to **string** | This party&#39;s Identifier. | [optional] 
**Remarks** | Pointer to **string** | Remarks for this party. | [optional] 
**Email** | Pointer to **string** | An email address for this party. | [optional] 
**Phone** | Pointer to **string** | A phone number for this party. | [optional] 
**ContactPersonName** | Pointer to **string** | The name of a contact person. | [optional] 
**AwardIdentifier** | Pointer to **string** | This is an identifier related to this PartyInformation and their internal reference of the tender (or award). This might be used to assign an identifier (German \&quot;Vergabenummer\&quot;) for the current project. This is typically only used in Buyer and Bidder representations and should map to the concept of \&quot;Vergabenummer\&quot; or \&quot;AwardNo\&quot; in GAEB. | [optional] 
**IsInEuropeanEconomicArea** | **bool** | This property indicates if the party is registered within the European Economic Area. | 
**VatId** | Pointer to **string** | If this is within the European Economic Area (see IsInEuropeanEconomicArea, then as a business entity it likely has an EU VAT Id. | [optional] 
**Fax** | Pointer to **string** | The fax number for this party. | [optional] 
**CountryCode** | Pointer to **string** | The two letter ISO country code, e.g. DE for Germany. | [optional] 
**CreditorOrDebtorIdentifier** | Pointer to **string** | Depending on which party this class represents, it might have either a &#39;creditor&#39; or &#39;debtor&#39; number. This is often used in internal accounting systems. | [optional] 
**GlobalLocationNumber** | Pointer to **string** | The Global Location Number (GLN) is issued by GS1 and is intended to be a unique identifier for the physical address of a party, e.g. a business office. | [optional] 
**BankingInformation** | Pointer to [**[]BankingInformationDto**](BankingInformationDto.md) | This list contains information about bank accounts associated with this PartyInformation. It&#39;s typically used for buyers and bidders. | [optional] 

## Methods

### NewPartyInformationDto

`func NewPartyInformationDto(id string, isInEuropeanEconomicArea bool, ) *PartyInformationDto`

NewPartyInformationDto instantiates a new PartyInformationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartyInformationDtoWithDefaults

`func NewPartyInformationDtoWithDefaults() *PartyInformationDto`

NewPartyInformationDtoWithDefaults instantiates a new PartyInformationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartyInformationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartyInformationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartyInformationDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PartyInformationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartyInformationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartyInformationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartyInformationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStreet

`func (o *PartyInformationDto) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *PartyInformationDto) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *PartyInformationDto) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *PartyInformationDto) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetZipCode

`func (o *PartyInformationDto) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *PartyInformationDto) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *PartyInformationDto) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *PartyInformationDto) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCity

`func (o *PartyInformationDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PartyInformationDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PartyInformationDto) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PartyInformationDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *PartyInformationDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PartyInformationDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PartyInformationDto) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PartyInformationDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetIdentifier

`func (o *PartyInformationDto) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PartyInformationDto) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PartyInformationDto) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PartyInformationDto) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRemarks

`func (o *PartyInformationDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *PartyInformationDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *PartyInformationDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *PartyInformationDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetEmail

`func (o *PartyInformationDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PartyInformationDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PartyInformationDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PartyInformationDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *PartyInformationDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PartyInformationDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PartyInformationDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PartyInformationDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetContactPersonName

`func (o *PartyInformationDto) GetContactPersonName() string`

GetContactPersonName returns the ContactPersonName field if non-nil, zero value otherwise.

### GetContactPersonNameOk

`func (o *PartyInformationDto) GetContactPersonNameOk() (*string, bool)`

GetContactPersonNameOk returns a tuple with the ContactPersonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPersonName

`func (o *PartyInformationDto) SetContactPersonName(v string)`

SetContactPersonName sets ContactPersonName field to given value.

### HasContactPersonName

`func (o *PartyInformationDto) HasContactPersonName() bool`

HasContactPersonName returns a boolean if a field has been set.

### GetAwardIdentifier

`func (o *PartyInformationDto) GetAwardIdentifier() string`

GetAwardIdentifier returns the AwardIdentifier field if non-nil, zero value otherwise.

### GetAwardIdentifierOk

`func (o *PartyInformationDto) GetAwardIdentifierOk() (*string, bool)`

GetAwardIdentifierOk returns a tuple with the AwardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardIdentifier

`func (o *PartyInformationDto) SetAwardIdentifier(v string)`

SetAwardIdentifier sets AwardIdentifier field to given value.

### HasAwardIdentifier

`func (o *PartyInformationDto) HasAwardIdentifier() bool`

HasAwardIdentifier returns a boolean if a field has been set.

### GetIsInEuropeanEconomicArea

`func (o *PartyInformationDto) GetIsInEuropeanEconomicArea() bool`

GetIsInEuropeanEconomicArea returns the IsInEuropeanEconomicArea field if non-nil, zero value otherwise.

### GetIsInEuropeanEconomicAreaOk

`func (o *PartyInformationDto) GetIsInEuropeanEconomicAreaOk() (*bool, bool)`

GetIsInEuropeanEconomicAreaOk returns a tuple with the IsInEuropeanEconomicArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInEuropeanEconomicArea

`func (o *PartyInformationDto) SetIsInEuropeanEconomicArea(v bool)`

SetIsInEuropeanEconomicArea sets IsInEuropeanEconomicArea field to given value.


### GetVatId

`func (o *PartyInformationDto) GetVatId() string`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *PartyInformationDto) GetVatIdOk() (*string, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *PartyInformationDto) SetVatId(v string)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *PartyInformationDto) HasVatId() bool`

HasVatId returns a boolean if a field has been set.

### GetFax

`func (o *PartyInformationDto) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *PartyInformationDto) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *PartyInformationDto) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *PartyInformationDto) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetCountryCode

`func (o *PartyInformationDto) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PartyInformationDto) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PartyInformationDto) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PartyInformationDto) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreditorOrDebtorIdentifier

`func (o *PartyInformationDto) GetCreditorOrDebtorIdentifier() string`

GetCreditorOrDebtorIdentifier returns the CreditorOrDebtorIdentifier field if non-nil, zero value otherwise.

### GetCreditorOrDebtorIdentifierOk

`func (o *PartyInformationDto) GetCreditorOrDebtorIdentifierOk() (*string, bool)`

GetCreditorOrDebtorIdentifierOk returns a tuple with the CreditorOrDebtorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditorOrDebtorIdentifier

`func (o *PartyInformationDto) SetCreditorOrDebtorIdentifier(v string)`

SetCreditorOrDebtorIdentifier sets CreditorOrDebtorIdentifier field to given value.

### HasCreditorOrDebtorIdentifier

`func (o *PartyInformationDto) HasCreditorOrDebtorIdentifier() bool`

HasCreditorOrDebtorIdentifier returns a boolean if a field has been set.

### GetGlobalLocationNumber

`func (o *PartyInformationDto) GetGlobalLocationNumber() string`

GetGlobalLocationNumber returns the GlobalLocationNumber field if non-nil, zero value otherwise.

### GetGlobalLocationNumberOk

`func (o *PartyInformationDto) GetGlobalLocationNumberOk() (*string, bool)`

GetGlobalLocationNumberOk returns a tuple with the GlobalLocationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalLocationNumber

`func (o *PartyInformationDto) SetGlobalLocationNumber(v string)`

SetGlobalLocationNumber sets GlobalLocationNumber field to given value.

### HasGlobalLocationNumber

`func (o *PartyInformationDto) HasGlobalLocationNumber() bool`

HasGlobalLocationNumber returns a boolean if a field has been set.

### GetBankingInformation

`func (o *PartyInformationDto) GetBankingInformation() []BankingInformationDto`

GetBankingInformation returns the BankingInformation field if non-nil, zero value otherwise.

### GetBankingInformationOk

`func (o *PartyInformationDto) GetBankingInformationOk() (*[]BankingInformationDto, bool)`

GetBankingInformationOk returns a tuple with the BankingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingInformation

`func (o *PartyInformationDto) SetBankingInformation(v []BankingInformationDto)`

SetBankingInformation sets BankingInformation field to given value.

### HasBankingInformation

`func (o *PartyInformationDto) HasBankingInformation() bool`

HasBankingInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


