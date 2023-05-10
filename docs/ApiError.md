# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **map[string][]string** | This dictionary contains a set of all errors and their messages | [optional] 

## Methods

### NewApiError

`func NewApiError() *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ApiError) GetErrors() map[string][]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApiError) GetErrorsOk() (*map[string][]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApiError) SetErrors(v map[string][]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ApiError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


