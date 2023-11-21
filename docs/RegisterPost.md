# RegisterPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Email** | **string** |  | 
**Password** | **string** |  | 
**PreferredLanguages** | Pointer to **[]string** |  | [optional] 
**ServiceLoginRedirectUri** | Pointer to **string** |  | [optional] 

## Methods

### NewRegisterPost

`func NewRegisterPost(username string, email string, password string, ) *RegisterPost`

NewRegisterPost instantiates a new RegisterPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterPostWithDefaults

`func NewRegisterPostWithDefaults() *RegisterPost`

NewRegisterPostWithDefaults instantiates a new RegisterPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *RegisterPost) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterPost) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterPost) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *RegisterPost) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterPost) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterPost) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *RegisterPost) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterPost) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterPost) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPreferredLanguages

`func (o *RegisterPost) GetPreferredLanguages() []string`

GetPreferredLanguages returns the PreferredLanguages field if non-nil, zero value otherwise.

### GetPreferredLanguagesOk

`func (o *RegisterPost) GetPreferredLanguagesOk() (*[]string, bool)`

GetPreferredLanguagesOk returns a tuple with the PreferredLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguages

`func (o *RegisterPost) SetPreferredLanguages(v []string)`

SetPreferredLanguages sets PreferredLanguages field to given value.

### HasPreferredLanguages

`func (o *RegisterPost) HasPreferredLanguages() bool`

HasPreferredLanguages returns a boolean if a field has been set.

### GetServiceLoginRedirectUri

`func (o *RegisterPost) GetServiceLoginRedirectUri() string`

GetServiceLoginRedirectUri returns the ServiceLoginRedirectUri field if non-nil, zero value otherwise.

### GetServiceLoginRedirectUriOk

`func (o *RegisterPost) GetServiceLoginRedirectUriOk() (*string, bool)`

GetServiceLoginRedirectUriOk returns a tuple with the ServiceLoginRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLoginRedirectUri

`func (o *RegisterPost) SetServiceLoginRedirectUri(v string)`

SetServiceLoginRedirectUri sets ServiceLoginRedirectUri field to given value.

### HasServiceLoginRedirectUri

`func (o *RegisterPost) HasServiceLoginRedirectUri() bool`

HasServiceLoginRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


