# ValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSchemaCompliant** | **bool** | Indicates if the validation produced no errors | 
**CheckResults** | Pointer to [**[]ValidationCheckResult**](ValidationCheckResult.md) | List of validation check results, which are usually errors | [optional] 
**ValidationType** | [**ValidationType**](ValidationType.md) |  | 
**FileName** | Pointer to **string** | Name of the validated file | [optional] 

## Methods

### NewValidationResult

`func NewValidationResult(isSchemaCompliant bool, validationType ValidationType, ) *ValidationResult`

NewValidationResult instantiates a new ValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResultWithDefaults

`func NewValidationResultWithDefaults() *ValidationResult`

NewValidationResultWithDefaults instantiates a new ValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSchemaCompliant

`func (o *ValidationResult) GetIsSchemaCompliant() bool`

GetIsSchemaCompliant returns the IsSchemaCompliant field if non-nil, zero value otherwise.

### GetIsSchemaCompliantOk

`func (o *ValidationResult) GetIsSchemaCompliantOk() (*bool, bool)`

GetIsSchemaCompliantOk returns a tuple with the IsSchemaCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSchemaCompliant

`func (o *ValidationResult) SetIsSchemaCompliant(v bool)`

SetIsSchemaCompliant sets IsSchemaCompliant field to given value.


### GetCheckResults

`func (o *ValidationResult) GetCheckResults() []ValidationCheckResult`

GetCheckResults returns the CheckResults field if non-nil, zero value otherwise.

### GetCheckResultsOk

`func (o *ValidationResult) GetCheckResultsOk() (*[]ValidationCheckResult, bool)`

GetCheckResultsOk returns a tuple with the CheckResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResults

`func (o *ValidationResult) SetCheckResults(v []ValidationCheckResult)`

SetCheckResults sets CheckResults field to given value.

### HasCheckResults

`func (o *ValidationResult) HasCheckResults() bool`

HasCheckResults returns a boolean if a field has been set.

### GetValidationType

`func (o *ValidationResult) GetValidationType() ValidationType`

GetValidationType returns the ValidationType field if non-nil, zero value otherwise.

### GetValidationTypeOk

`func (o *ValidationResult) GetValidationTypeOk() (*ValidationType, bool)`

GetValidationTypeOk returns a tuple with the ValidationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationType

`func (o *ValidationResult) SetValidationType(v ValidationType)`

SetValidationType sets ValidationType field to given value.


### GetFileName

`func (o *ValidationResult) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ValidationResult) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ValidationResult) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ValidationResult) HasFileName() bool`

HasFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


