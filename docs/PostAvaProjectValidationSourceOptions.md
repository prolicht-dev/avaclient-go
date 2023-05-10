# PostAvaProjectValidationSourceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvaProject** | [**ProjectDto**](ProjectDto.md) |  | 
**ServiceSpecificationIndex** | Pointer to **int32** | The index of the ServiceSpecification that should be validated. If not given, will default to the first one in the project. | [optional] 
**ValidationType** | [**ValidationType**](ValidationType.md) |  | 
**ExchangePhase** | Pointer to [**ExchangePhaseDto**](ExchangePhaseDto.md) |  | [optional] 
**AvaSourceOptions** | Pointer to [**PostAvaSourceOptions**](PostAvaSourceOptions.md) |  | [optional] 
**OenormDestinationOptions** | Pointer to [**PostOenormDestinationOptions**](PostOenormDestinationOptions.md) |  | [optional] 
**GaebDestinationOptions** | Pointer to [**PostGaebDestinationOptions**](PostGaebDestinationOptions.md) |  | [optional] 

## Methods

### NewPostAvaProjectValidationSourceOptions

`func NewPostAvaProjectValidationSourceOptions(avaProject ProjectDto, validationType ValidationType, ) *PostAvaProjectValidationSourceOptions`

NewPostAvaProjectValidationSourceOptions instantiates a new PostAvaProjectValidationSourceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAvaProjectValidationSourceOptionsWithDefaults

`func NewPostAvaProjectValidationSourceOptionsWithDefaults() *PostAvaProjectValidationSourceOptions`

NewPostAvaProjectValidationSourceOptionsWithDefaults instantiates a new PostAvaProjectValidationSourceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvaProject

`func (o *PostAvaProjectValidationSourceOptions) GetAvaProject() ProjectDto`

GetAvaProject returns the AvaProject field if non-nil, zero value otherwise.

### GetAvaProjectOk

`func (o *PostAvaProjectValidationSourceOptions) GetAvaProjectOk() (*ProjectDto, bool)`

GetAvaProjectOk returns a tuple with the AvaProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvaProject

`func (o *PostAvaProjectValidationSourceOptions) SetAvaProject(v ProjectDto)`

SetAvaProject sets AvaProject field to given value.


### GetServiceSpecificationIndex

`func (o *PostAvaProjectValidationSourceOptions) GetServiceSpecificationIndex() int32`

GetServiceSpecificationIndex returns the ServiceSpecificationIndex field if non-nil, zero value otherwise.

### GetServiceSpecificationIndexOk

`func (o *PostAvaProjectValidationSourceOptions) GetServiceSpecificationIndexOk() (*int32, bool)`

GetServiceSpecificationIndexOk returns a tuple with the ServiceSpecificationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSpecificationIndex

`func (o *PostAvaProjectValidationSourceOptions) SetServiceSpecificationIndex(v int32)`

SetServiceSpecificationIndex sets ServiceSpecificationIndex field to given value.

### HasServiceSpecificationIndex

`func (o *PostAvaProjectValidationSourceOptions) HasServiceSpecificationIndex() bool`

HasServiceSpecificationIndex returns a boolean if a field has been set.

### GetValidationType

`func (o *PostAvaProjectValidationSourceOptions) GetValidationType() ValidationType`

GetValidationType returns the ValidationType field if non-nil, zero value otherwise.

### GetValidationTypeOk

`func (o *PostAvaProjectValidationSourceOptions) GetValidationTypeOk() (*ValidationType, bool)`

GetValidationTypeOk returns a tuple with the ValidationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationType

`func (o *PostAvaProjectValidationSourceOptions) SetValidationType(v ValidationType)`

SetValidationType sets ValidationType field to given value.


### GetExchangePhase

`func (o *PostAvaProjectValidationSourceOptions) GetExchangePhase() ExchangePhaseDto`

GetExchangePhase returns the ExchangePhase field if non-nil, zero value otherwise.

### GetExchangePhaseOk

`func (o *PostAvaProjectValidationSourceOptions) GetExchangePhaseOk() (*ExchangePhaseDto, bool)`

GetExchangePhaseOk returns a tuple with the ExchangePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangePhase

`func (o *PostAvaProjectValidationSourceOptions) SetExchangePhase(v ExchangePhaseDto)`

SetExchangePhase sets ExchangePhase field to given value.

### HasExchangePhase

`func (o *PostAvaProjectValidationSourceOptions) HasExchangePhase() bool`

HasExchangePhase returns a boolean if a field has been set.

### GetAvaSourceOptions

`func (o *PostAvaProjectValidationSourceOptions) GetAvaSourceOptions() PostAvaSourceOptions`

GetAvaSourceOptions returns the AvaSourceOptions field if non-nil, zero value otherwise.

### GetAvaSourceOptionsOk

`func (o *PostAvaProjectValidationSourceOptions) GetAvaSourceOptionsOk() (*PostAvaSourceOptions, bool)`

GetAvaSourceOptionsOk returns a tuple with the AvaSourceOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvaSourceOptions

`func (o *PostAvaProjectValidationSourceOptions) SetAvaSourceOptions(v PostAvaSourceOptions)`

SetAvaSourceOptions sets AvaSourceOptions field to given value.

### HasAvaSourceOptions

`func (o *PostAvaProjectValidationSourceOptions) HasAvaSourceOptions() bool`

HasAvaSourceOptions returns a boolean if a field has been set.

### GetOenormDestinationOptions

`func (o *PostAvaProjectValidationSourceOptions) GetOenormDestinationOptions() PostOenormDestinationOptions`

GetOenormDestinationOptions returns the OenormDestinationOptions field if non-nil, zero value otherwise.

### GetOenormDestinationOptionsOk

`func (o *PostAvaProjectValidationSourceOptions) GetOenormDestinationOptionsOk() (*PostOenormDestinationOptions, bool)`

GetOenormDestinationOptionsOk returns a tuple with the OenormDestinationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOenormDestinationOptions

`func (o *PostAvaProjectValidationSourceOptions) SetOenormDestinationOptions(v PostOenormDestinationOptions)`

SetOenormDestinationOptions sets OenormDestinationOptions field to given value.

### HasOenormDestinationOptions

`func (o *PostAvaProjectValidationSourceOptions) HasOenormDestinationOptions() bool`

HasOenormDestinationOptions returns a boolean if a field has been set.

### GetGaebDestinationOptions

`func (o *PostAvaProjectValidationSourceOptions) GetGaebDestinationOptions() PostGaebDestinationOptions`

GetGaebDestinationOptions returns the GaebDestinationOptions field if non-nil, zero value otherwise.

### GetGaebDestinationOptionsOk

`func (o *PostAvaProjectValidationSourceOptions) GetGaebDestinationOptionsOk() (*PostGaebDestinationOptions, bool)`

GetGaebDestinationOptionsOk returns a tuple with the GaebDestinationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaebDestinationOptions

`func (o *PostAvaProjectValidationSourceOptions) SetGaebDestinationOptions(v PostGaebDestinationOptions)`

SetGaebDestinationOptions sets GaebDestinationOptions field to given value.

### HasGaebDestinationOptions

`func (o *PostAvaProjectValidationSourceOptions) HasGaebDestinationOptions() bool`

HasGaebDestinationOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


