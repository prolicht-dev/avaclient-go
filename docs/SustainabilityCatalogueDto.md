# SustainabilityCatalogueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Key** | Pointer to **string** | The key must be a unique identifier for the sustainability catalogue. It is used to link sustainability data in positions to a catalogue. The Dangl.AVA module does not enforce uniqueness, but it is recommended to use a unique key for each catalogue. | [optional] 
**Description** | Pointer to **string** | The description property can be used to provide further information about the sustainability catalogue. | [optional] 

## Methods

### NewSustainabilityCatalogueDto

`func NewSustainabilityCatalogueDto(id string, ) *SustainabilityCatalogueDto`

NewSustainabilityCatalogueDto instantiates a new SustainabilityCatalogueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSustainabilityCatalogueDtoWithDefaults

`func NewSustainabilityCatalogueDtoWithDefaults() *SustainabilityCatalogueDto`

NewSustainabilityCatalogueDtoWithDefaults instantiates a new SustainabilityCatalogueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SustainabilityCatalogueDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SustainabilityCatalogueDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SustainabilityCatalogueDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *SustainabilityCatalogueDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SustainabilityCatalogueDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SustainabilityCatalogueDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SustainabilityCatalogueDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *SustainabilityCatalogueDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SustainabilityCatalogueDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SustainabilityCatalogueDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SustainabilityCatalogueDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


