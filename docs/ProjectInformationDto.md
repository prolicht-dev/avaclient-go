# ProjectInformationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buyer** | [**PartyInformationDto**](PartyInformationDto.md) |  | [optional] 
**Description** | **string** | Description for the project. | [optional] 
**DescriptionShort** | **string** | Short description for the project. | [optional] 
**Name** | **string** | Name of the project. | [optional] 
**Site** | [**PartyInformationDto**](PartyInformationDto.md) |  | [optional] 
**ItemNumberSchema** | [**ItemNumberSchemaDto**](ItemNumberSchemaDto.md) |  | [optional] 
**CurrencyShort** | **string** | Short label for the currency used. | [optional] 
**CurrencyLong** | **string** | Full label of the currency used. | [optional] 
**LabourTimeLabel** | **string** | Label for the labour time part of prices used in the project. | [optional] 
**PriceComponents** | **[]string** | Labels for the price components used in the project. Caution: Removal of a price component will trigger to have dependent price informations be deleted in IElements that use this property. | [optional] 
**BidderCommentAllowed** | **bool** | This bool indicates if this project allows the bidder to add bidder comments. Bidder comments are a way to attach clarifying information when submitting an offer. | 
**SideOffersAllowed** | **bool** | This property indicates if the project should allow side offers from contractors. In GAEB, a side offer would typically be in exchange phase 85. | 
**AwardType** | [**AwardTypeDto**](AwardTypeDto.md) |  | 
**SpecialAwardKind** | [**SpecialAwardKindDto**](SpecialAwardKindDto.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


