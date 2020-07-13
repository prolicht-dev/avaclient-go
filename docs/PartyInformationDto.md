# PartyInformationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Name** | **string** | This party&#39;s name. | [optional] 
**Street** | **string** | This party&#39;s street. | [optional] 
**ZipCode** | **string** | This party&#39;s ZipCode. | [optional] 
**City** | **string** | This party&#39;s City. | [optional] 
**Country** | **string** | This party&#39;s Country. | [optional] 
**Identifier** | **string** | This party&#39;s Identifier. | [optional] 
**Remarks** | **string** | Remarks for this party. | [optional] 
**Email** | **string** | An email address for this party. | [optional] 
**Phone** | **string** | A phone number for this party. | [optional] 
**ContactPersonName** | **string** | The name of a contact person. | [optional] 
**AwardIdentifier** | **string** | This is an identifier related to this PartyInformation and their internal reference of the tender (or award). This might be used to assign an identifier (German \&quot;Vergabenummer\&quot;) for the current project. This is typically only used in Buyer and Bidder representations and should map to the concept of \&quot;Vergabenummer\&quot; or \&quot;AwardNo\&quot; in GAEB. | [optional] 
**IsInEuropeanEconomicArea** | **bool** | This property indicates if the party is registered within the European Economic Area. | 
**VatId** | **string** | If this is within the European Economic Area (see IsInEuropeanEconomicArea, then as a business entity it likely has an EU VAT Id. | [optional] 
**Fax** | **string** | The fax number for this party. | [optional] 
**CountryCode** | **string** | The two letter ISO country code, e.g. DE for Germany. | [optional] 
**CreditorOrDebtorIdentifier** | **string** | Depending on which party this class represents, it might have either a &#39;creditor&#39; or &#39;debtor&#39; number. This is often used in internal accounting systems. | [optional] 
**GlobalLocationNumber** | **string** | The Global Location Number (GLN) is issued by GS1 and is intended to be a unique identifier for the physical address of a party, e.g. a business office. | [optional] 
**BankingInformation** | [**[]BankingInformationDto**](BankingInformationDto.md) | This list contains information about bank accounts associated with this PartyInformation. It&#39;s typically used for buyers and bidders. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


