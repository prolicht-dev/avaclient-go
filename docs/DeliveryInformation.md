# DeliveryInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliverToName** | Pointer to **string** | BT-70 BT-70 in XRechnung | [optional] 
**DeliverToLocationIdentifier** | Pointer to **string** | BT-71 BT-71 in XRechnung | [optional] 
**DeliveryDate** | Pointer to **time.Time** | BT-72 BT-72 in XRechnung | [optional] 
**DeliveryAddress** | Pointer to [**DeliveryAddress**](DeliveryAddress.md) |  | [optional] 

## Methods

### NewDeliveryInformation

`func NewDeliveryInformation() *DeliveryInformation`

NewDeliveryInformation instantiates a new DeliveryInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryInformationWithDefaults

`func NewDeliveryInformationWithDefaults() *DeliveryInformation`

NewDeliveryInformationWithDefaults instantiates a new DeliveryInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliverToName

`func (o *DeliveryInformation) GetDeliverToName() string`

GetDeliverToName returns the DeliverToName field if non-nil, zero value otherwise.

### GetDeliverToNameOk

`func (o *DeliveryInformation) GetDeliverToNameOk() (*string, bool)`

GetDeliverToNameOk returns a tuple with the DeliverToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverToName

`func (o *DeliveryInformation) SetDeliverToName(v string)`

SetDeliverToName sets DeliverToName field to given value.

### HasDeliverToName

`func (o *DeliveryInformation) HasDeliverToName() bool`

HasDeliverToName returns a boolean if a field has been set.

### GetDeliverToLocationIdentifier

`func (o *DeliveryInformation) GetDeliverToLocationIdentifier() string`

GetDeliverToLocationIdentifier returns the DeliverToLocationIdentifier field if non-nil, zero value otherwise.

### GetDeliverToLocationIdentifierOk

`func (o *DeliveryInformation) GetDeliverToLocationIdentifierOk() (*string, bool)`

GetDeliverToLocationIdentifierOk returns a tuple with the DeliverToLocationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverToLocationIdentifier

`func (o *DeliveryInformation) SetDeliverToLocationIdentifier(v string)`

SetDeliverToLocationIdentifier sets DeliverToLocationIdentifier field to given value.

### HasDeliverToLocationIdentifier

`func (o *DeliveryInformation) HasDeliverToLocationIdentifier() bool`

HasDeliverToLocationIdentifier returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *DeliveryInformation) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *DeliveryInformation) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *DeliveryInformation) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *DeliveryInformation) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *DeliveryInformation) GetDeliveryAddress() DeliveryAddress`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *DeliveryInformation) GetDeliveryAddressOk() (*DeliveryAddress, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *DeliveryInformation) SetDeliveryAddress(v DeliveryAddress)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *DeliveryInformation) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


