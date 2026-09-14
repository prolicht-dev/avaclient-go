# GetProjectMergeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergedProject** | Pointer to [**ProjectDto**](ProjectDto.md) |  | [optional] 
**NonMatchedSourceElements** | Pointer to [**[]IElementDto**](IElementDto.md) | Elements that were present in the source project, but could not be matched to any element in the target project. | [optional] 
**NonMatchedTargetElements** | Pointer to [**[]IElementDto**](IElementDto.md) | Elements that were present in the target project, but could not be matched to any element in the source project. | [optional] 

## Methods

### NewGetProjectMergeResult

`func NewGetProjectMergeResult() *GetProjectMergeResult`

NewGetProjectMergeResult instantiates a new GetProjectMergeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectMergeResultWithDefaults

`func NewGetProjectMergeResultWithDefaults() *GetProjectMergeResult`

NewGetProjectMergeResultWithDefaults instantiates a new GetProjectMergeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergedProject

`func (o *GetProjectMergeResult) GetMergedProject() ProjectDto`

GetMergedProject returns the MergedProject field if non-nil, zero value otherwise.

### GetMergedProjectOk

`func (o *GetProjectMergeResult) GetMergedProjectOk() (*ProjectDto, bool)`

GetMergedProjectOk returns a tuple with the MergedProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedProject

`func (o *GetProjectMergeResult) SetMergedProject(v ProjectDto)`

SetMergedProject sets MergedProject field to given value.

### HasMergedProject

`func (o *GetProjectMergeResult) HasMergedProject() bool`

HasMergedProject returns a boolean if a field has been set.

### GetNonMatchedSourceElements

`func (o *GetProjectMergeResult) GetNonMatchedSourceElements() []IElementDto`

GetNonMatchedSourceElements returns the NonMatchedSourceElements field if non-nil, zero value otherwise.

### GetNonMatchedSourceElementsOk

`func (o *GetProjectMergeResult) GetNonMatchedSourceElementsOk() (*[]IElementDto, bool)`

GetNonMatchedSourceElementsOk returns a tuple with the NonMatchedSourceElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonMatchedSourceElements

`func (o *GetProjectMergeResult) SetNonMatchedSourceElements(v []IElementDto)`

SetNonMatchedSourceElements sets NonMatchedSourceElements field to given value.

### HasNonMatchedSourceElements

`func (o *GetProjectMergeResult) HasNonMatchedSourceElements() bool`

HasNonMatchedSourceElements returns a boolean if a field has been set.

### GetNonMatchedTargetElements

`func (o *GetProjectMergeResult) GetNonMatchedTargetElements() []IElementDto`

GetNonMatchedTargetElements returns the NonMatchedTargetElements field if non-nil, zero value otherwise.

### GetNonMatchedTargetElementsOk

`func (o *GetProjectMergeResult) GetNonMatchedTargetElementsOk() (*[]IElementDto, bool)`

GetNonMatchedTargetElementsOk returns a tuple with the NonMatchedTargetElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonMatchedTargetElements

`func (o *GetProjectMergeResult) SetNonMatchedTargetElements(v []IElementDto)`

SetNonMatchedTargetElements sets NonMatchedTargetElements field to given value.

### HasNonMatchedTargetElements

`func (o *GetProjectMergeResult) HasNonMatchedTargetElements() bool`

HasNonMatchedTargetElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


