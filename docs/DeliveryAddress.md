# DeliveryAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | BT-75, BT-76 and BT-165 The address. BT-75, BT-76 and BT-165 in XRechnung. | [optional] 
**City** | Pointer to **string** | BT-77 The city. BT-77 in XRechnung. | [optional] 
**ZipCode** | Pointer to **string** | BT-78 The zip code. BT-78 in XRechnung. | [optional] 
**CountryCode** | Pointer to **string** | BT-80 The ISO 3166-1 2 letter country code. BT-80 in XRechnung. | [optional] 

## Methods

### NewDeliveryAddress

`func NewDeliveryAddress() *DeliveryAddress`

NewDeliveryAddress instantiates a new DeliveryAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryAddressWithDefaults

`func NewDeliveryAddressWithDefaults() *DeliveryAddress`

NewDeliveryAddressWithDefaults instantiates a new DeliveryAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DeliveryAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DeliveryAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DeliveryAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DeliveryAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *DeliveryAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DeliveryAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DeliveryAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *DeliveryAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetZipCode

`func (o *DeliveryAddress) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *DeliveryAddress) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *DeliveryAddress) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *DeliveryAddress) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *DeliveryAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DeliveryAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DeliveryAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DeliveryAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


