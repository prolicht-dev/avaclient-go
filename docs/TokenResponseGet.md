# TokenResponseGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**ExpiresIn** | **int64** |  | 
**HttpErrorReason** | Pointer to **string** |  | [optional] 
**HttpStatusCode** | [**HttpStatusCode**](HttpStatusCode.md) |  | 
**IdentityToken** | Pointer to **string** |  | [optional] 
**IsError** | **bool** |  | 
**RefreshToken** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**ErrorType** | [**ResponseErrorType**](ResponseErrorType.md) |  | 

## Methods

### NewTokenResponseGet

`func NewTokenResponseGet(expiresIn int64, httpStatusCode HttpStatusCode, isError bool, errorType ResponseErrorType, ) *TokenResponseGet`

NewTokenResponseGet instantiates a new TokenResponseGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseGetWithDefaults

`func NewTokenResponseGetWithDefaults() *TokenResponseGet`

NewTokenResponseGetWithDefaults instantiates a new TokenResponseGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *TokenResponseGet) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenResponseGet) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenResponseGet) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenResponseGet) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetError

`func (o *TokenResponseGet) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TokenResponseGet) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TokenResponseGet) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TokenResponseGet) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorDescription

`func (o *TokenResponseGet) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *TokenResponseGet) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *TokenResponseGet) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *TokenResponseGet) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetExpiresIn

`func (o *TokenResponseGet) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *TokenResponseGet) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *TokenResponseGet) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.


### GetHttpErrorReason

`func (o *TokenResponseGet) GetHttpErrorReason() string`

GetHttpErrorReason returns the HttpErrorReason field if non-nil, zero value otherwise.

### GetHttpErrorReasonOk

`func (o *TokenResponseGet) GetHttpErrorReasonOk() (*string, bool)`

GetHttpErrorReasonOk returns a tuple with the HttpErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpErrorReason

`func (o *TokenResponseGet) SetHttpErrorReason(v string)`

SetHttpErrorReason sets HttpErrorReason field to given value.

### HasHttpErrorReason

`func (o *TokenResponseGet) HasHttpErrorReason() bool`

HasHttpErrorReason returns a boolean if a field has been set.

### GetHttpStatusCode

`func (o *TokenResponseGet) GetHttpStatusCode() HttpStatusCode`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *TokenResponseGet) GetHttpStatusCodeOk() (*HttpStatusCode, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *TokenResponseGet) SetHttpStatusCode(v HttpStatusCode)`

SetHttpStatusCode sets HttpStatusCode field to given value.


### GetIdentityToken

`func (o *TokenResponseGet) GetIdentityToken() string`

GetIdentityToken returns the IdentityToken field if non-nil, zero value otherwise.

### GetIdentityTokenOk

`func (o *TokenResponseGet) GetIdentityTokenOk() (*string, bool)`

GetIdentityTokenOk returns a tuple with the IdentityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityToken

`func (o *TokenResponseGet) SetIdentityToken(v string)`

SetIdentityToken sets IdentityToken field to given value.

### HasIdentityToken

`func (o *TokenResponseGet) HasIdentityToken() bool`

HasIdentityToken returns a boolean if a field has been set.

### GetIsError

`func (o *TokenResponseGet) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *TokenResponseGet) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *TokenResponseGet) SetIsError(v bool)`

SetIsError sets IsError field to given value.


### GetRefreshToken

`func (o *TokenResponseGet) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenResponseGet) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenResponseGet) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenResponseGet) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenResponseGet) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenResponseGet) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenResponseGet) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TokenResponseGet) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetErrorType

`func (o *TokenResponseGet) GetErrorType() ResponseErrorType`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *TokenResponseGet) GetErrorTypeOk() (*ResponseErrorType, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *TokenResponseGet) SetErrorType(v ResponseErrorType)`

SetErrorType sets ErrorType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


