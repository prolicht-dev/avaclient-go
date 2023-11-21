# PriceCatalogueDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**UnitPriceFrom** | Pointer to **float32** | The lower bound of the unit price | [optional] 
**UnitPriceTo** | Pointer to **float32** | The upper bound of the unit price | [optional] 
**UnitPriceAverage** | Pointer to **float32** | The average unit price | [optional] 
**WagePriceFrom** | Pointer to **float32** | The lower bound of the wage price | [optional] 
**WagePriceTo** | Pointer to **float32** | The upper bound of the wage price | [optional] 
**WagePriceAverage** | Pointer to **float32** | The average wage price | [optional] 

## Methods

### NewPriceCatalogueDataDto

`func NewPriceCatalogueDataDto(id string, ) *PriceCatalogueDataDto`

NewPriceCatalogueDataDto instantiates a new PriceCatalogueDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceCatalogueDataDtoWithDefaults

`func NewPriceCatalogueDataDtoWithDefaults() *PriceCatalogueDataDto`

NewPriceCatalogueDataDtoWithDefaults instantiates a new PriceCatalogueDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriceCatalogueDataDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceCatalogueDataDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceCatalogueDataDto) SetId(v string)`

SetId sets Id field to given value.


### GetUnitPriceFrom

`func (o *PriceCatalogueDataDto) GetUnitPriceFrom() float32`

GetUnitPriceFrom returns the UnitPriceFrom field if non-nil, zero value otherwise.

### GetUnitPriceFromOk

`func (o *PriceCatalogueDataDto) GetUnitPriceFromOk() (*float32, bool)`

GetUnitPriceFromOk returns a tuple with the UnitPriceFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceFrom

`func (o *PriceCatalogueDataDto) SetUnitPriceFrom(v float32)`

SetUnitPriceFrom sets UnitPriceFrom field to given value.

### HasUnitPriceFrom

`func (o *PriceCatalogueDataDto) HasUnitPriceFrom() bool`

HasUnitPriceFrom returns a boolean if a field has been set.

### GetUnitPriceTo

`func (o *PriceCatalogueDataDto) GetUnitPriceTo() float32`

GetUnitPriceTo returns the UnitPriceTo field if non-nil, zero value otherwise.

### GetUnitPriceToOk

`func (o *PriceCatalogueDataDto) GetUnitPriceToOk() (*float32, bool)`

GetUnitPriceToOk returns a tuple with the UnitPriceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceTo

`func (o *PriceCatalogueDataDto) SetUnitPriceTo(v float32)`

SetUnitPriceTo sets UnitPriceTo field to given value.

### HasUnitPriceTo

`func (o *PriceCatalogueDataDto) HasUnitPriceTo() bool`

HasUnitPriceTo returns a boolean if a field has been set.

### GetUnitPriceAverage

`func (o *PriceCatalogueDataDto) GetUnitPriceAverage() float32`

GetUnitPriceAverage returns the UnitPriceAverage field if non-nil, zero value otherwise.

### GetUnitPriceAverageOk

`func (o *PriceCatalogueDataDto) GetUnitPriceAverageOk() (*float32, bool)`

GetUnitPriceAverageOk returns a tuple with the UnitPriceAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceAverage

`func (o *PriceCatalogueDataDto) SetUnitPriceAverage(v float32)`

SetUnitPriceAverage sets UnitPriceAverage field to given value.

### HasUnitPriceAverage

`func (o *PriceCatalogueDataDto) HasUnitPriceAverage() bool`

HasUnitPriceAverage returns a boolean if a field has been set.

### GetWagePriceFrom

`func (o *PriceCatalogueDataDto) GetWagePriceFrom() float32`

GetWagePriceFrom returns the WagePriceFrom field if non-nil, zero value otherwise.

### GetWagePriceFromOk

`func (o *PriceCatalogueDataDto) GetWagePriceFromOk() (*float32, bool)`

GetWagePriceFromOk returns a tuple with the WagePriceFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagePriceFrom

`func (o *PriceCatalogueDataDto) SetWagePriceFrom(v float32)`

SetWagePriceFrom sets WagePriceFrom field to given value.

### HasWagePriceFrom

`func (o *PriceCatalogueDataDto) HasWagePriceFrom() bool`

HasWagePriceFrom returns a boolean if a field has been set.

### GetWagePriceTo

`func (o *PriceCatalogueDataDto) GetWagePriceTo() float32`

GetWagePriceTo returns the WagePriceTo field if non-nil, zero value otherwise.

### GetWagePriceToOk

`func (o *PriceCatalogueDataDto) GetWagePriceToOk() (*float32, bool)`

GetWagePriceToOk returns a tuple with the WagePriceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagePriceTo

`func (o *PriceCatalogueDataDto) SetWagePriceTo(v float32)`

SetWagePriceTo sets WagePriceTo field to given value.

### HasWagePriceTo

`func (o *PriceCatalogueDataDto) HasWagePriceTo() bool`

HasWagePriceTo returns a boolean if a field has been set.

### GetWagePriceAverage

`func (o *PriceCatalogueDataDto) GetWagePriceAverage() float32`

GetWagePriceAverage returns the WagePriceAverage field if non-nil, zero value otherwise.

### GetWagePriceAverageOk

`func (o *PriceCatalogueDataDto) GetWagePriceAverageOk() (*float32, bool)`

GetWagePriceAverageOk returns a tuple with the WagePriceAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagePriceAverage

`func (o *PriceCatalogueDataDto) SetWagePriceAverage(v float32)`

SetWagePriceAverage sets WagePriceAverage field to given value.

### HasWagePriceAverage

`func (o *PriceCatalogueDataDto) HasWagePriceAverage() bool`

HasWagePriceAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


