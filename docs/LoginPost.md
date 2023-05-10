# LoginPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**Password** | **string** |  | 
**StaySignedIn** | **bool** |  | 

## Methods

### NewLoginPost

`func NewLoginPost(identifier string, password string, staySignedIn bool, ) *LoginPost`

NewLoginPost instantiates a new LoginPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginPostWithDefaults

`func NewLoginPostWithDefaults() *LoginPost`

NewLoginPostWithDefaults instantiates a new LoginPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *LoginPost) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *LoginPost) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *LoginPost) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetPassword

`func (o *LoginPost) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginPost) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginPost) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetStaySignedIn

`func (o *LoginPost) GetStaySignedIn() bool`

GetStaySignedIn returns the StaySignedIn field if non-nil, zero value otherwise.

### GetStaySignedInOk

`func (o *LoginPost) GetStaySignedInOk() (*bool, bool)`

GetStaySignedInOk returns a tuple with the StaySignedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaySignedIn

`func (o *LoginPost) SetStaySignedIn(v bool)`

SetStaySignedIn sets StaySignedIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


