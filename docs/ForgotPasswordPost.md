# ForgotPasswordPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**PreferredLanguages** | Pointer to **[]string** |  | [optional] 

## Methods

### NewForgotPasswordPost

`func NewForgotPasswordPost(identifier string, ) *ForgotPasswordPost`

NewForgotPasswordPost instantiates a new ForgotPasswordPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForgotPasswordPostWithDefaults

`func NewForgotPasswordPostWithDefaults() *ForgotPasswordPost`

NewForgotPasswordPostWithDefaults instantiates a new ForgotPasswordPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ForgotPasswordPost) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ForgotPasswordPost) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ForgotPasswordPost) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetPreferredLanguages

`func (o *ForgotPasswordPost) GetPreferredLanguages() []string`

GetPreferredLanguages returns the PreferredLanguages field if non-nil, zero value otherwise.

### GetPreferredLanguagesOk

`func (o *ForgotPasswordPost) GetPreferredLanguagesOk() (*[]string, bool)`

GetPreferredLanguagesOk returns a tuple with the PreferredLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguages

`func (o *ForgotPasswordPost) SetPreferredLanguages(v []string)`

SetPreferredLanguages sets PreferredLanguages field to given value.

### HasPreferredLanguages

`func (o *ForgotPasswordPost) HasPreferredLanguages() bool`

HasPreferredLanguages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


