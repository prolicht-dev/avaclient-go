# ExecutionDescriptionDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | Pointer to [**[]NoteTextDto**](NoteTextDto.md) | Blocks within an ExecutionDescription contain the actual information. | [optional] 
**Label** | Pointer to **string** | Labels this ExecutionDescription. | [optional] 
**Identifier** | Pointer to **string** | Uniquely identifies this ExecutionDescription. | [optional] 
**ElementType** | Pointer to **string** |  | [optional] 

## Methods

### NewExecutionDescriptionDtoAllOf

`func NewExecutionDescriptionDtoAllOf() *ExecutionDescriptionDtoAllOf`

NewExecutionDescriptionDtoAllOf instantiates a new ExecutionDescriptionDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionDescriptionDtoAllOfWithDefaults

`func NewExecutionDescriptionDtoAllOfWithDefaults() *ExecutionDescriptionDtoAllOf`

NewExecutionDescriptionDtoAllOfWithDefaults instantiates a new ExecutionDescriptionDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *ExecutionDescriptionDtoAllOf) GetBlocks() []NoteTextDto`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *ExecutionDescriptionDtoAllOf) GetBlocksOk() (*[]NoteTextDto, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *ExecutionDescriptionDtoAllOf) SetBlocks(v []NoteTextDto)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *ExecutionDescriptionDtoAllOf) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetLabel

`func (o *ExecutionDescriptionDtoAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExecutionDescriptionDtoAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExecutionDescriptionDtoAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ExecutionDescriptionDtoAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIdentifier

`func (o *ExecutionDescriptionDtoAllOf) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ExecutionDescriptionDtoAllOf) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ExecutionDescriptionDtoAllOf) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ExecutionDescriptionDtoAllOf) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetElementType

`func (o *ExecutionDescriptionDtoAllOf) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *ExecutionDescriptionDtoAllOf) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *ExecutionDescriptionDtoAllOf) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *ExecutionDescriptionDtoAllOf) HasElementType() bool`

HasElementType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


