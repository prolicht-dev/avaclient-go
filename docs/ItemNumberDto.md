# ItemNumberDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**StringRepresentation** | Pointer to **string** | Will return this ItemNumber as point delimited string. Will not distinguish between upper- and lowercase and return an all-lowercase representation. Will consider first numbers, then characters, e.g. 1a is considered preceding aa. Transformation to all lowercase can be configured in the ItemNumberSchema property. | [optional] 
**IsSchemaCompliant** | **bool** | Indicates if the characters and the length of the Identifiers match the current ItemNumberSchema. | [readonly] 
**ItemNumberSchema** | Pointer to [**ItemNumberSchemaDto**](ItemNumberSchemaDto.md) |  | [optional] 
**Identifiers** | Pointer to **[]string** | Collection of the single identifiers in this ItemNumber. P.e., \&quot;02.03.004\&quot; will have three elements \&quot;02\&quot;, \&quot;03\&quot;, and \&quot;004\&quot;. Since ReadOnlyObservableCollection&#x60;1 does have the event set to protected, it can be accessed like this: (itemNumber.Identifiers as INotifyCollectionChanged).CollectionChanged | [optional] 
**IsLot** | **bool** | This indicates if this item number is at the lot level. Find out more about lots in the documentation. | [readonly] 
**HierarchyLevel** | **int32** | This is a zero based hierarchy level. It&#39;s set automatically when used in the context of a Project, and can be used to identify the hierarchy level of the current element. | 
**IsAttachedToPosition** | **bool** | This property indicates if this ItemNumber is attached to an object of the Position type. | [readonly] 

## Methods

### NewItemNumberDto

`func NewItemNumberDto(id string, isSchemaCompliant bool, isLot bool, hierarchyLevel int32, isAttachedToPosition bool, ) *ItemNumberDto`

NewItemNumberDto instantiates a new ItemNumberDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemNumberDtoWithDefaults

`func NewItemNumberDtoWithDefaults() *ItemNumberDto`

NewItemNumberDtoWithDefaults instantiates a new ItemNumberDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemNumberDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemNumberDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemNumberDto) SetId(v string)`

SetId sets Id field to given value.


### GetStringRepresentation

`func (o *ItemNumberDto) GetStringRepresentation() string`

GetStringRepresentation returns the StringRepresentation field if non-nil, zero value otherwise.

### GetStringRepresentationOk

`func (o *ItemNumberDto) GetStringRepresentationOk() (*string, bool)`

GetStringRepresentationOk returns a tuple with the StringRepresentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringRepresentation

`func (o *ItemNumberDto) SetStringRepresentation(v string)`

SetStringRepresentation sets StringRepresentation field to given value.

### HasStringRepresentation

`func (o *ItemNumberDto) HasStringRepresentation() bool`

HasStringRepresentation returns a boolean if a field has been set.

### GetIsSchemaCompliant

`func (o *ItemNumberDto) GetIsSchemaCompliant() bool`

GetIsSchemaCompliant returns the IsSchemaCompliant field if non-nil, zero value otherwise.

### GetIsSchemaCompliantOk

`func (o *ItemNumberDto) GetIsSchemaCompliantOk() (*bool, bool)`

GetIsSchemaCompliantOk returns a tuple with the IsSchemaCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSchemaCompliant

`func (o *ItemNumberDto) SetIsSchemaCompliant(v bool)`

SetIsSchemaCompliant sets IsSchemaCompliant field to given value.


### GetItemNumberSchema

`func (o *ItemNumberDto) GetItemNumberSchema() ItemNumberSchemaDto`

GetItemNumberSchema returns the ItemNumberSchema field if non-nil, zero value otherwise.

### GetItemNumberSchemaOk

`func (o *ItemNumberDto) GetItemNumberSchemaOk() (*ItemNumberSchemaDto, bool)`

GetItemNumberSchemaOk returns a tuple with the ItemNumberSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemNumberSchema

`func (o *ItemNumberDto) SetItemNumberSchema(v ItemNumberSchemaDto)`

SetItemNumberSchema sets ItemNumberSchema field to given value.

### HasItemNumberSchema

`func (o *ItemNumberDto) HasItemNumberSchema() bool`

HasItemNumberSchema returns a boolean if a field has been set.

### GetIdentifiers

`func (o *ItemNumberDto) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ItemNumberDto) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ItemNumberDto) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *ItemNumberDto) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetIsLot

`func (o *ItemNumberDto) GetIsLot() bool`

GetIsLot returns the IsLot field if non-nil, zero value otherwise.

### GetIsLotOk

`func (o *ItemNumberDto) GetIsLotOk() (*bool, bool)`

GetIsLotOk returns a tuple with the IsLot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLot

`func (o *ItemNumberDto) SetIsLot(v bool)`

SetIsLot sets IsLot field to given value.


### GetHierarchyLevel

`func (o *ItemNumberDto) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *ItemNumberDto) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *ItemNumberDto) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.


### GetIsAttachedToPosition

`func (o *ItemNumberDto) GetIsAttachedToPosition() bool`

GetIsAttachedToPosition returns the IsAttachedToPosition field if non-nil, zero value otherwise.

### GetIsAttachedToPositionOk

`func (o *ItemNumberDto) GetIsAttachedToPositionOk() (*bool, bool)`

GetIsAttachedToPositionOk returns a tuple with the IsAttachedToPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAttachedToPosition

`func (o *ItemNumberDto) SetIsAttachedToPosition(v bool)`

SetIsAttachedToPosition sets IsAttachedToPosition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


