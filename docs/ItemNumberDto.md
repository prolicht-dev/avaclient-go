# ItemNumberDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**StringRepresentation** | **string** | Will return this ItemNumber as point delimited string. Will not distinguish between upper- and lowercase and return an all-lowercase representation. Will consider first numbers, then characters, e.g. 1a is considered preceding aa. | [optional] 
**IsSchemaCompliant** | **bool** | Indicates if the characters and the length of the Identifiers match the current ItemNumberSchema. | [readonly] 
**ItemNumberSchema** | [**ItemNumberSchemaDto**](ItemNumberSchemaDto.md) |  | [optional] 
**Identifiers** | **[]string** | Collection of the single identifiers in this ItemNumber. P.e., \&quot;02.03.004\&quot; will have three elements \&quot;02\&quot;, \&quot;03\&quot;, and \&quot;004\&quot;. Since ReadOnlyObservableCollection&#x60;1 does have the event set to protected, it can be accessed like this: (itemNumber.Identifiers as INotifyCollectionChanged).CollectionChanged | [optional] 
**IsLot** | **bool** | This indicates if this item number is at the lot level. Find out more about lots in the documentation. | [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


