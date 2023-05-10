# ProjectInformationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buyer** | Pointer to [**PartyInformationDto**](PartyInformationDto.md) |  | [optional] 
**Description** | Pointer to **string** | Description for the project. | [optional] 
**DescriptionShort** | Pointer to **string** | Short description for the project. | [optional] 
**Name** | Pointer to **string** | Name of the project. | [optional] 
**Site** | Pointer to [**PartyInformationDto**](PartyInformationDto.md) |  | [optional] 
**ItemNumberSchema** | Pointer to [**ItemNumberSchemaDto**](ItemNumberSchemaDto.md) |  | [optional] 
**CurrencyShort** | Pointer to **string** | Short label for the currency used. | [optional] 
**CurrencyLong** | Pointer to **string** | Full label of the currency used. | [optional] 
**LabourTimeLabel** | Pointer to **string** | Label for the labour time part of prices used in the project. | [optional] 
**PriceComponents** | Pointer to **[]string** | Labels for the price components used in the project. Caution: Removal of a price component will trigger to have dependent price informations be deleted in IElements that use this property. If this property is changed or altered after the project has already been used, it is strongly advised to do operations step by step, e.g. if renaming and reordering multiple price components, this should be done one by one. Otherwise, a combination of rename and reordering will not be correctly propagated downwards to child objects in this Project. | [optional] 
**PriceComponentTypes** | Pointer to [**map[string]PriceComponentTypeDto**](PriceComponentTypeDto.md) | This dictionary specifies actual types used for the PriceComponents. For example, a single price component might have the name &#39;Material&#39; as a string, ans this dictionary would then have a key &#39;Material&#39; and a value of Materials. You can not add keys here that are not also present in the PriceComponents property, and removing price components will also remove them from this dictionary here. | [optional] [readonly] 
**BidderCommentAllowed** | **bool** | This bool indicates if this project allows the bidder to add bidder comments. Bidder comments are a way to attach clarifying information when submitting an offer. | 
**SideOffersAllowed** | **bool** | This property indicates if the project should allow side offers from contractors. In GAEB, a side offer would typically be in exchange phase 85. | 
**AwardType** | [**AwardTypeDto**](AwardTypeDto.md) |  | 
**SpecialAwardKind** | [**SpecialAwardKindDto**](SpecialAwardKindDto.md) |  | 
**Requesters** | Pointer to [**[]PartyInformationDto**](PartyInformationDto.md) | Requesters in a construction project, in German also called &#39;Bedarfsträger&#39;, are parties connected to the building process, e.g. architects or planners. | [optional] 
**NotificationSites** | Pointer to [**[]PartyInformationDto**](PartyInformationDto.md) | Notification sites are addresses that are used for delivering messages in the context of construction projects. | [optional] 

## Methods

### NewProjectInformationDto

`func NewProjectInformationDto(bidderCommentAllowed bool, sideOffersAllowed bool, awardType AwardTypeDto, specialAwardKind SpecialAwardKindDto, ) *ProjectInformationDto`

NewProjectInformationDto instantiates a new ProjectInformationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectInformationDtoWithDefaults

`func NewProjectInformationDtoWithDefaults() *ProjectInformationDto`

NewProjectInformationDtoWithDefaults instantiates a new ProjectInformationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyer

`func (o *ProjectInformationDto) GetBuyer() PartyInformationDto`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *ProjectInformationDto) GetBuyerOk() (*PartyInformationDto, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *ProjectInformationDto) SetBuyer(v PartyInformationDto)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *ProjectInformationDto) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectInformationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectInformationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectInformationDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectInformationDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionShort

`func (o *ProjectInformationDto) GetDescriptionShort() string`

GetDescriptionShort returns the DescriptionShort field if non-nil, zero value otherwise.

### GetDescriptionShortOk

`func (o *ProjectInformationDto) GetDescriptionShortOk() (*string, bool)`

GetDescriptionShortOk returns a tuple with the DescriptionShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionShort

`func (o *ProjectInformationDto) SetDescriptionShort(v string)`

SetDescriptionShort sets DescriptionShort field to given value.

### HasDescriptionShort

`func (o *ProjectInformationDto) HasDescriptionShort() bool`

HasDescriptionShort returns a boolean if a field has been set.

### GetName

`func (o *ProjectInformationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectInformationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectInformationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectInformationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSite

`func (o *ProjectInformationDto) GetSite() PartyInformationDto`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ProjectInformationDto) GetSiteOk() (*PartyInformationDto, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ProjectInformationDto) SetSite(v PartyInformationDto)`

SetSite sets Site field to given value.

### HasSite

`func (o *ProjectInformationDto) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetItemNumberSchema

`func (o *ProjectInformationDto) GetItemNumberSchema() ItemNumberSchemaDto`

GetItemNumberSchema returns the ItemNumberSchema field if non-nil, zero value otherwise.

### GetItemNumberSchemaOk

`func (o *ProjectInformationDto) GetItemNumberSchemaOk() (*ItemNumberSchemaDto, bool)`

GetItemNumberSchemaOk returns a tuple with the ItemNumberSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemNumberSchema

`func (o *ProjectInformationDto) SetItemNumberSchema(v ItemNumberSchemaDto)`

SetItemNumberSchema sets ItemNumberSchema field to given value.

### HasItemNumberSchema

`func (o *ProjectInformationDto) HasItemNumberSchema() bool`

HasItemNumberSchema returns a boolean if a field has been set.

### GetCurrencyShort

`func (o *ProjectInformationDto) GetCurrencyShort() string`

GetCurrencyShort returns the CurrencyShort field if non-nil, zero value otherwise.

### GetCurrencyShortOk

`func (o *ProjectInformationDto) GetCurrencyShortOk() (*string, bool)`

GetCurrencyShortOk returns a tuple with the CurrencyShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyShort

`func (o *ProjectInformationDto) SetCurrencyShort(v string)`

SetCurrencyShort sets CurrencyShort field to given value.

### HasCurrencyShort

`func (o *ProjectInformationDto) HasCurrencyShort() bool`

HasCurrencyShort returns a boolean if a field has been set.

### GetCurrencyLong

`func (o *ProjectInformationDto) GetCurrencyLong() string`

GetCurrencyLong returns the CurrencyLong field if non-nil, zero value otherwise.

### GetCurrencyLongOk

`func (o *ProjectInformationDto) GetCurrencyLongOk() (*string, bool)`

GetCurrencyLongOk returns a tuple with the CurrencyLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyLong

`func (o *ProjectInformationDto) SetCurrencyLong(v string)`

SetCurrencyLong sets CurrencyLong field to given value.

### HasCurrencyLong

`func (o *ProjectInformationDto) HasCurrencyLong() bool`

HasCurrencyLong returns a boolean if a field has been set.

### GetLabourTimeLabel

`func (o *ProjectInformationDto) GetLabourTimeLabel() string`

GetLabourTimeLabel returns the LabourTimeLabel field if non-nil, zero value otherwise.

### GetLabourTimeLabelOk

`func (o *ProjectInformationDto) GetLabourTimeLabelOk() (*string, bool)`

GetLabourTimeLabelOk returns a tuple with the LabourTimeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabourTimeLabel

`func (o *ProjectInformationDto) SetLabourTimeLabel(v string)`

SetLabourTimeLabel sets LabourTimeLabel field to given value.

### HasLabourTimeLabel

`func (o *ProjectInformationDto) HasLabourTimeLabel() bool`

HasLabourTimeLabel returns a boolean if a field has been set.

### GetPriceComponents

`func (o *ProjectInformationDto) GetPriceComponents() []string`

GetPriceComponents returns the PriceComponents field if non-nil, zero value otherwise.

### GetPriceComponentsOk

`func (o *ProjectInformationDto) GetPriceComponentsOk() (*[]string, bool)`

GetPriceComponentsOk returns a tuple with the PriceComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceComponents

`func (o *ProjectInformationDto) SetPriceComponents(v []string)`

SetPriceComponents sets PriceComponents field to given value.

### HasPriceComponents

`func (o *ProjectInformationDto) HasPriceComponents() bool`

HasPriceComponents returns a boolean if a field has been set.

### GetPriceComponentTypes

`func (o *ProjectInformationDto) GetPriceComponentTypes() map[string]PriceComponentTypeDto`

GetPriceComponentTypes returns the PriceComponentTypes field if non-nil, zero value otherwise.

### GetPriceComponentTypesOk

`func (o *ProjectInformationDto) GetPriceComponentTypesOk() (*map[string]PriceComponentTypeDto, bool)`

GetPriceComponentTypesOk returns a tuple with the PriceComponentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceComponentTypes

`func (o *ProjectInformationDto) SetPriceComponentTypes(v map[string]PriceComponentTypeDto)`

SetPriceComponentTypes sets PriceComponentTypes field to given value.

### HasPriceComponentTypes

`func (o *ProjectInformationDto) HasPriceComponentTypes() bool`

HasPriceComponentTypes returns a boolean if a field has been set.

### GetBidderCommentAllowed

`func (o *ProjectInformationDto) GetBidderCommentAllowed() bool`

GetBidderCommentAllowed returns the BidderCommentAllowed field if non-nil, zero value otherwise.

### GetBidderCommentAllowedOk

`func (o *ProjectInformationDto) GetBidderCommentAllowedOk() (*bool, bool)`

GetBidderCommentAllowedOk returns a tuple with the BidderCommentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidderCommentAllowed

`func (o *ProjectInformationDto) SetBidderCommentAllowed(v bool)`

SetBidderCommentAllowed sets BidderCommentAllowed field to given value.


### GetSideOffersAllowed

`func (o *ProjectInformationDto) GetSideOffersAllowed() bool`

GetSideOffersAllowed returns the SideOffersAllowed field if non-nil, zero value otherwise.

### GetSideOffersAllowedOk

`func (o *ProjectInformationDto) GetSideOffersAllowedOk() (*bool, bool)`

GetSideOffersAllowedOk returns a tuple with the SideOffersAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSideOffersAllowed

`func (o *ProjectInformationDto) SetSideOffersAllowed(v bool)`

SetSideOffersAllowed sets SideOffersAllowed field to given value.


### GetAwardType

`func (o *ProjectInformationDto) GetAwardType() AwardTypeDto`

GetAwardType returns the AwardType field if non-nil, zero value otherwise.

### GetAwardTypeOk

`func (o *ProjectInformationDto) GetAwardTypeOk() (*AwardTypeDto, bool)`

GetAwardTypeOk returns a tuple with the AwardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardType

`func (o *ProjectInformationDto) SetAwardType(v AwardTypeDto)`

SetAwardType sets AwardType field to given value.


### GetSpecialAwardKind

`func (o *ProjectInformationDto) GetSpecialAwardKind() SpecialAwardKindDto`

GetSpecialAwardKind returns the SpecialAwardKind field if non-nil, zero value otherwise.

### GetSpecialAwardKindOk

`func (o *ProjectInformationDto) GetSpecialAwardKindOk() (*SpecialAwardKindDto, bool)`

GetSpecialAwardKindOk returns a tuple with the SpecialAwardKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialAwardKind

`func (o *ProjectInformationDto) SetSpecialAwardKind(v SpecialAwardKindDto)`

SetSpecialAwardKind sets SpecialAwardKind field to given value.


### GetRequesters

`func (o *ProjectInformationDto) GetRequesters() []PartyInformationDto`

GetRequesters returns the Requesters field if non-nil, zero value otherwise.

### GetRequestersOk

`func (o *ProjectInformationDto) GetRequestersOk() (*[]PartyInformationDto, bool)`

GetRequestersOk returns a tuple with the Requesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesters

`func (o *ProjectInformationDto) SetRequesters(v []PartyInformationDto)`

SetRequesters sets Requesters field to given value.

### HasRequesters

`func (o *ProjectInformationDto) HasRequesters() bool`

HasRequesters returns a boolean if a field has been set.

### GetNotificationSites

`func (o *ProjectInformationDto) GetNotificationSites() []PartyInformationDto`

GetNotificationSites returns the NotificationSites field if non-nil, zero value otherwise.

### GetNotificationSitesOk

`func (o *ProjectInformationDto) GetNotificationSitesOk() (*[]PartyInformationDto, bool)`

GetNotificationSitesOk returns a tuple with the NotificationSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationSites

`func (o *ProjectInformationDto) SetNotificationSites(v []PartyInformationDto)`

SetNotificationSites sets NotificationSites field to given value.

### HasNotificationSites

`func (o *ProjectInformationDto) HasNotificationSites() bool`

HasNotificationSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


