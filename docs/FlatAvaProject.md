# FlatAvaProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**ProjectDto**](ProjectDto.md) |  | [optional] 
**Elements** | Pointer to [**[]FlatElementDto**](FlatElementDto.md) | The flattened elements of the project. | [optional] 

## Methods

### NewFlatAvaProject

`func NewFlatAvaProject() *FlatAvaProject`

NewFlatAvaProject instantiates a new FlatAvaProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatAvaProjectWithDefaults

`func NewFlatAvaProjectWithDefaults() *FlatAvaProject`

NewFlatAvaProjectWithDefaults instantiates a new FlatAvaProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *FlatAvaProject) GetProject() ProjectDto`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *FlatAvaProject) GetProjectOk() (*ProjectDto, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *FlatAvaProject) SetProject(v ProjectDto)`

SetProject sets Project field to given value.

### HasProject

`func (o *FlatAvaProject) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetElements

`func (o *FlatAvaProject) GetElements() []FlatElementDto`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *FlatAvaProject) GetElementsOk() (*[]FlatElementDto, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *FlatAvaProject) SetElements(v []FlatElementDto)`

SetElements sets Elements field to given value.

### HasElements

`func (o *FlatAvaProject) HasElements() bool`

HasElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


