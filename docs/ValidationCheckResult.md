# ValidationCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | [**ValidationSeverity**](ValidationSeverity.md) |  | 
**CheckType** | [**ValidationCheckType**](ValidationCheckType.md) |  | 
**Message** | Pointer to **string** | A human readable message describing the result of the check. | [optional] 
**ObjectValidationCheckDetails** | Pointer to [**ObjectValidationCheckDetails**](ObjectValidationCheckDetails.md) |  | [optional] 
**XmlSchemaValidationCheckDetails** | Pointer to [**XmlSchemaValidationCheckDetails**](XmlSchemaValidationCheckDetails.md) |  | [optional] 
**ProjectValidationCheckDetails** | Pointer to [**ProjectValidationCheckDetails**](ProjectValidationCheckDetails.md) |  | [optional] 

## Methods

### NewValidationCheckResult

`func NewValidationCheckResult(severity ValidationSeverity, checkType ValidationCheckType, ) *ValidationCheckResult`

NewValidationCheckResult instantiates a new ValidationCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationCheckResultWithDefaults

`func NewValidationCheckResultWithDefaults() *ValidationCheckResult`

NewValidationCheckResultWithDefaults instantiates a new ValidationCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *ValidationCheckResult) GetSeverity() ValidationSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ValidationCheckResult) GetSeverityOk() (*ValidationSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ValidationCheckResult) SetSeverity(v ValidationSeverity)`

SetSeverity sets Severity field to given value.


### GetCheckType

`func (o *ValidationCheckResult) GetCheckType() ValidationCheckType`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *ValidationCheckResult) GetCheckTypeOk() (*ValidationCheckType, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *ValidationCheckResult) SetCheckType(v ValidationCheckType)`

SetCheckType sets CheckType field to given value.


### GetMessage

`func (o *ValidationCheckResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationCheckResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationCheckResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationCheckResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetObjectValidationCheckDetails

`func (o *ValidationCheckResult) GetObjectValidationCheckDetails() ObjectValidationCheckDetails`

GetObjectValidationCheckDetails returns the ObjectValidationCheckDetails field if non-nil, zero value otherwise.

### GetObjectValidationCheckDetailsOk

`func (o *ValidationCheckResult) GetObjectValidationCheckDetailsOk() (*ObjectValidationCheckDetails, bool)`

GetObjectValidationCheckDetailsOk returns a tuple with the ObjectValidationCheckDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectValidationCheckDetails

`func (o *ValidationCheckResult) SetObjectValidationCheckDetails(v ObjectValidationCheckDetails)`

SetObjectValidationCheckDetails sets ObjectValidationCheckDetails field to given value.

### HasObjectValidationCheckDetails

`func (o *ValidationCheckResult) HasObjectValidationCheckDetails() bool`

HasObjectValidationCheckDetails returns a boolean if a field has been set.

### GetXmlSchemaValidationCheckDetails

`func (o *ValidationCheckResult) GetXmlSchemaValidationCheckDetails() XmlSchemaValidationCheckDetails`

GetXmlSchemaValidationCheckDetails returns the XmlSchemaValidationCheckDetails field if non-nil, zero value otherwise.

### GetXmlSchemaValidationCheckDetailsOk

`func (o *ValidationCheckResult) GetXmlSchemaValidationCheckDetailsOk() (*XmlSchemaValidationCheckDetails, bool)`

GetXmlSchemaValidationCheckDetailsOk returns a tuple with the XmlSchemaValidationCheckDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlSchemaValidationCheckDetails

`func (o *ValidationCheckResult) SetXmlSchemaValidationCheckDetails(v XmlSchemaValidationCheckDetails)`

SetXmlSchemaValidationCheckDetails sets XmlSchemaValidationCheckDetails field to given value.

### HasXmlSchemaValidationCheckDetails

`func (o *ValidationCheckResult) HasXmlSchemaValidationCheckDetails() bool`

HasXmlSchemaValidationCheckDetails returns a boolean if a field has been set.

### GetProjectValidationCheckDetails

`func (o *ValidationCheckResult) GetProjectValidationCheckDetails() ProjectValidationCheckDetails`

GetProjectValidationCheckDetails returns the ProjectValidationCheckDetails field if non-nil, zero value otherwise.

### GetProjectValidationCheckDetailsOk

`func (o *ValidationCheckResult) GetProjectValidationCheckDetailsOk() (*ProjectValidationCheckDetails, bool)`

GetProjectValidationCheckDetailsOk returns a tuple with the ProjectValidationCheckDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectValidationCheckDetails

`func (o *ValidationCheckResult) SetProjectValidationCheckDetails(v ProjectValidationCheckDetails)`

SetProjectValidationCheckDetails sets ProjectValidationCheckDetails field to given value.

### HasProjectValidationCheckDetails

`func (o *ValidationCheckResult) HasProjectValidationCheckDetails() bool`

HasProjectValidationCheckDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


