# GetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsHealthy** | **bool** | If any problems in the service health are known, this is set to false | 
**Version** | Pointer to **string** | The current version of the AVACloud service | [optional] 
**Environment** | Pointer to **string** | The environment of the current instance | [optional] 

## Methods

### NewGetStatus

`func NewGetStatus(isHealthy bool, ) *GetStatus`

NewGetStatus instantiates a new GetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatusWithDefaults

`func NewGetStatusWithDefaults() *GetStatus`

NewGetStatusWithDefaults instantiates a new GetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsHealthy

`func (o *GetStatus) GetIsHealthy() bool`

GetIsHealthy returns the IsHealthy field if non-nil, zero value otherwise.

### GetIsHealthyOk

`func (o *GetStatus) GetIsHealthyOk() (*bool, bool)`

GetIsHealthyOk returns a tuple with the IsHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHealthy

`func (o *GetStatus) SetIsHealthy(v bool)`

SetIsHealthy sets IsHealthy field to given value.


### GetVersion

`func (o *GetStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnvironment

`func (o *GetStatus) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetStatus) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetStatus) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GetStatus) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


