# ExecutionDescriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | Pointer to [**[]NoteTextDto**](NoteTextDto.md) | Blocks within an ExecutionDescription contain the actual information. | [optional] 
**Label** | Pointer to **string** | Labels this ExecutionDescription. | [optional] 
**Identifier** | Pointer to **string** | Uniquely identifies this ExecutionDescription. | [optional] 
**ElementType** | Pointer to **string** |  | [optional] 

## Methods

### NewExecutionDescriptionDto

`func NewExecutionDescriptionDto() *ExecutionDescriptionDto`

NewExecutionDescriptionDto instantiates a new ExecutionDescriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionDescriptionDtoWithDefaults

`func NewExecutionDescriptionDtoWithDefaults() *ExecutionDescriptionDto`

NewExecutionDescriptionDtoWithDefaults instantiates a new ExecutionDescriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *ExecutionDescriptionDto) GetBlocks() []NoteTextDto`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *ExecutionDescriptionDto) GetBlocksOk() (*[]NoteTextDto, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *ExecutionDescriptionDto) SetBlocks(v []NoteTextDto)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *ExecutionDescriptionDto) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetLabel

`func (o *ExecutionDescriptionDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExecutionDescriptionDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExecutionDescriptionDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ExecutionDescriptionDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIdentifier

`func (o *ExecutionDescriptionDto) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ExecutionDescriptionDto) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ExecutionDescriptionDto) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ExecutionDescriptionDto) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetElementType

`func (o *ExecutionDescriptionDto) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *ExecutionDescriptionDto) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *ExecutionDescriptionDto) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *ExecutionDescriptionDto) HasElementType() bool`

HasElementType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


