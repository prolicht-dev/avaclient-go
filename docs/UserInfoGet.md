# UserInfoGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIsAuthenticated** | **bool** |  | 
**ClientIsAuthenticated** | **bool** |  | 
**CurrentUserId** | Pointer to **string** |  | [optional] 
**CurrentUserIdenticonId** | Pointer to **string** |  | [optional] 
**CurrentUserName** | Pointer to **string** |  | [optional] 
**CurrentUserEmail** | Pointer to **string** |  | [optional] 
**CurrentClientId** | Pointer to **string** |  | [optional] 
**UserClaims** | Pointer to [**[]ClaimGet**](ClaimGet.md) |  | [optional] 
**ClientClaims** | Pointer to [**[]ClaimGet**](ClaimGet.md) |  | [optional] 

## Methods

### NewUserInfoGet

`func NewUserInfoGet(userIsAuthenticated bool, clientIsAuthenticated bool, ) *UserInfoGet`

NewUserInfoGet instantiates a new UserInfoGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoGetWithDefaults

`func NewUserInfoGetWithDefaults() *UserInfoGet`

NewUserInfoGetWithDefaults instantiates a new UserInfoGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIsAuthenticated

`func (o *UserInfoGet) GetUserIsAuthenticated() bool`

GetUserIsAuthenticated returns the UserIsAuthenticated field if non-nil, zero value otherwise.

### GetUserIsAuthenticatedOk

`func (o *UserInfoGet) GetUserIsAuthenticatedOk() (*bool, bool)`

GetUserIsAuthenticatedOk returns a tuple with the UserIsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIsAuthenticated

`func (o *UserInfoGet) SetUserIsAuthenticated(v bool)`

SetUserIsAuthenticated sets UserIsAuthenticated field to given value.


### GetClientIsAuthenticated

`func (o *UserInfoGet) GetClientIsAuthenticated() bool`

GetClientIsAuthenticated returns the ClientIsAuthenticated field if non-nil, zero value otherwise.

### GetClientIsAuthenticatedOk

`func (o *UserInfoGet) GetClientIsAuthenticatedOk() (*bool, bool)`

GetClientIsAuthenticatedOk returns a tuple with the ClientIsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIsAuthenticated

`func (o *UserInfoGet) SetClientIsAuthenticated(v bool)`

SetClientIsAuthenticated sets ClientIsAuthenticated field to given value.


### GetCurrentUserId

`func (o *UserInfoGet) GetCurrentUserId() string`

GetCurrentUserId returns the CurrentUserId field if non-nil, zero value otherwise.

### GetCurrentUserIdOk

`func (o *UserInfoGet) GetCurrentUserIdOk() (*string, bool)`

GetCurrentUserIdOk returns a tuple with the CurrentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserId

`func (o *UserInfoGet) SetCurrentUserId(v string)`

SetCurrentUserId sets CurrentUserId field to given value.

### HasCurrentUserId

`func (o *UserInfoGet) HasCurrentUserId() bool`

HasCurrentUserId returns a boolean if a field has been set.

### GetCurrentUserIdenticonId

`func (o *UserInfoGet) GetCurrentUserIdenticonId() string`

GetCurrentUserIdenticonId returns the CurrentUserIdenticonId field if non-nil, zero value otherwise.

### GetCurrentUserIdenticonIdOk

`func (o *UserInfoGet) GetCurrentUserIdenticonIdOk() (*string, bool)`

GetCurrentUserIdenticonIdOk returns a tuple with the CurrentUserIdenticonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserIdenticonId

`func (o *UserInfoGet) SetCurrentUserIdenticonId(v string)`

SetCurrentUserIdenticonId sets CurrentUserIdenticonId field to given value.

### HasCurrentUserIdenticonId

`func (o *UserInfoGet) HasCurrentUserIdenticonId() bool`

HasCurrentUserIdenticonId returns a boolean if a field has been set.

### GetCurrentUserName

`func (o *UserInfoGet) GetCurrentUserName() string`

GetCurrentUserName returns the CurrentUserName field if non-nil, zero value otherwise.

### GetCurrentUserNameOk

`func (o *UserInfoGet) GetCurrentUserNameOk() (*string, bool)`

GetCurrentUserNameOk returns a tuple with the CurrentUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserName

`func (o *UserInfoGet) SetCurrentUserName(v string)`

SetCurrentUserName sets CurrentUserName field to given value.

### HasCurrentUserName

`func (o *UserInfoGet) HasCurrentUserName() bool`

HasCurrentUserName returns a boolean if a field has been set.

### GetCurrentUserEmail

`func (o *UserInfoGet) GetCurrentUserEmail() string`

GetCurrentUserEmail returns the CurrentUserEmail field if non-nil, zero value otherwise.

### GetCurrentUserEmailOk

`func (o *UserInfoGet) GetCurrentUserEmailOk() (*string, bool)`

GetCurrentUserEmailOk returns a tuple with the CurrentUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserEmail

`func (o *UserInfoGet) SetCurrentUserEmail(v string)`

SetCurrentUserEmail sets CurrentUserEmail field to given value.

### HasCurrentUserEmail

`func (o *UserInfoGet) HasCurrentUserEmail() bool`

HasCurrentUserEmail returns a boolean if a field has been set.

### GetCurrentClientId

`func (o *UserInfoGet) GetCurrentClientId() string`

GetCurrentClientId returns the CurrentClientId field if non-nil, zero value otherwise.

### GetCurrentClientIdOk

`func (o *UserInfoGet) GetCurrentClientIdOk() (*string, bool)`

GetCurrentClientIdOk returns a tuple with the CurrentClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentClientId

`func (o *UserInfoGet) SetCurrentClientId(v string)`

SetCurrentClientId sets CurrentClientId field to given value.

### HasCurrentClientId

`func (o *UserInfoGet) HasCurrentClientId() bool`

HasCurrentClientId returns a boolean if a field has been set.

### GetUserClaims

`func (o *UserInfoGet) GetUserClaims() []ClaimGet`

GetUserClaims returns the UserClaims field if non-nil, zero value otherwise.

### GetUserClaimsOk

`func (o *UserInfoGet) GetUserClaimsOk() (*[]ClaimGet, bool)`

GetUserClaimsOk returns a tuple with the UserClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaims

`func (o *UserInfoGet) SetUserClaims(v []ClaimGet)`

SetUserClaims sets UserClaims field to given value.

### HasUserClaims

`func (o *UserInfoGet) HasUserClaims() bool`

HasUserClaims returns a boolean if a field has been set.

### GetClientClaims

`func (o *UserInfoGet) GetClientClaims() []ClaimGet`

GetClientClaims returns the ClientClaims field if non-nil, zero value otherwise.

### GetClientClaimsOk

`func (o *UserInfoGet) GetClientClaimsOk() (*[]ClaimGet, bool)`

GetClientClaimsOk returns a tuple with the ClientClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientClaims

`func (o *UserInfoGet) SetClientClaims(v []ClaimGet)`

SetClientClaims sets ClientClaims field to given value.

### HasClientClaims

`func (o *UserInfoGet) HasClientClaims() bool`

HasClientClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


