# PostAvaSourceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TryAutoGenerateItemNumbersAndSchema** | **bool** | If this is set to true, AVACloud will try to generate item numbers and an item number schema automatically for the input project. The operation will not have any effect if either an item number schema is already present, or if any of the elements already has an item number. | 

## Methods

### NewPostAvaSourceOptions

`func NewPostAvaSourceOptions(tryAutoGenerateItemNumbersAndSchema bool, ) *PostAvaSourceOptions`

NewPostAvaSourceOptions instantiates a new PostAvaSourceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAvaSourceOptionsWithDefaults

`func NewPostAvaSourceOptionsWithDefaults() *PostAvaSourceOptions`

NewPostAvaSourceOptionsWithDefaults instantiates a new PostAvaSourceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTryAutoGenerateItemNumbersAndSchema

`func (o *PostAvaSourceOptions) GetTryAutoGenerateItemNumbersAndSchema() bool`

GetTryAutoGenerateItemNumbersAndSchema returns the TryAutoGenerateItemNumbersAndSchema field if non-nil, zero value otherwise.

### GetTryAutoGenerateItemNumbersAndSchemaOk

`func (o *PostAvaSourceOptions) GetTryAutoGenerateItemNumbersAndSchemaOk() (*bool, bool)`

GetTryAutoGenerateItemNumbersAndSchemaOk returns a tuple with the TryAutoGenerateItemNumbersAndSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryAutoGenerateItemNumbersAndSchema

`func (o *PostAvaSourceOptions) SetTryAutoGenerateItemNumbersAndSchema(v bool)`

SetTryAutoGenerateItemNumbersAndSchema sets TryAutoGenerateItemNumbersAndSchema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


