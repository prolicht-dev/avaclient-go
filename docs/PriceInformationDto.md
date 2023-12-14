# PriceInformationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**HourlyWage** | **float32** | The amount of currency per one hour of manual labour work in this project. | 
**DeductionFactor** | **float32** | The final, total price will be deducted by this rate. | 
**FlatSum** | **float32** | This is given when there is only one flat price for the whole service specification. | 
**TaxRate** | **float32** | Global tax rate for the project. Note that certain elements may have a different, specific tax rate. | 
**HasUnsetTaxRate** | **bool** | This property indicates if the TaxRate property is actually zero (0) or if it&#39;s just unset. The TaxRate is a non-nullable property, so it can&#39;t be null and therefore a distinction between unset and zero relies on this additional property. When exporting to GAEB, this can be used e.g. to decide whether or not the tax rate should be serialized. This property is automtically set to false when the TaxRate is set to any value, even zero. If the TaxRate is not set to zero, this property will always return false and can not be set to true. | 
**TradeDiscounts** | Pointer to [**[]TradeDiscountDto**](TradeDiscountDto.md) | Trade discounts for offered in this ServiceSpecification. | [optional] 

## Methods

### NewPriceInformationDto

`func NewPriceInformationDto(id string, hourlyWage float32, deductionFactor float32, flatSum float32, taxRate float32, hasUnsetTaxRate bool, ) *PriceInformationDto`

NewPriceInformationDto instantiates a new PriceInformationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceInformationDtoWithDefaults

`func NewPriceInformationDtoWithDefaults() *PriceInformationDto`

NewPriceInformationDtoWithDefaults instantiates a new PriceInformationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PriceInformationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceInformationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceInformationDto) SetId(v string)`

SetId sets Id field to given value.


### GetHourlyWage

`func (o *PriceInformationDto) GetHourlyWage() float32`

GetHourlyWage returns the HourlyWage field if non-nil, zero value otherwise.

### GetHourlyWageOk

`func (o *PriceInformationDto) GetHourlyWageOk() (*float32, bool)`

GetHourlyWageOk returns a tuple with the HourlyWage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyWage

`func (o *PriceInformationDto) SetHourlyWage(v float32)`

SetHourlyWage sets HourlyWage field to given value.


### GetDeductionFactor

`func (o *PriceInformationDto) GetDeductionFactor() float32`

GetDeductionFactor returns the DeductionFactor field if non-nil, zero value otherwise.

### GetDeductionFactorOk

`func (o *PriceInformationDto) GetDeductionFactorOk() (*float32, bool)`

GetDeductionFactorOk returns a tuple with the DeductionFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductionFactor

`func (o *PriceInformationDto) SetDeductionFactor(v float32)`

SetDeductionFactor sets DeductionFactor field to given value.


### GetFlatSum

`func (o *PriceInformationDto) GetFlatSum() float32`

GetFlatSum returns the FlatSum field if non-nil, zero value otherwise.

### GetFlatSumOk

`func (o *PriceInformationDto) GetFlatSumOk() (*float32, bool)`

GetFlatSumOk returns a tuple with the FlatSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatSum

`func (o *PriceInformationDto) SetFlatSum(v float32)`

SetFlatSum sets FlatSum field to given value.


### GetTaxRate

`func (o *PriceInformationDto) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *PriceInformationDto) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *PriceInformationDto) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.


### GetHasUnsetTaxRate

`func (o *PriceInformationDto) GetHasUnsetTaxRate() bool`

GetHasUnsetTaxRate returns the HasUnsetTaxRate field if non-nil, zero value otherwise.

### GetHasUnsetTaxRateOk

`func (o *PriceInformationDto) GetHasUnsetTaxRateOk() (*bool, bool)`

GetHasUnsetTaxRateOk returns a tuple with the HasUnsetTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUnsetTaxRate

`func (o *PriceInformationDto) SetHasUnsetTaxRate(v bool)`

SetHasUnsetTaxRate sets HasUnsetTaxRate field to given value.


### GetTradeDiscounts

`func (o *PriceInformationDto) GetTradeDiscounts() []TradeDiscountDto`

GetTradeDiscounts returns the TradeDiscounts field if non-nil, zero value otherwise.

### GetTradeDiscountsOk

`func (o *PriceInformationDto) GetTradeDiscountsOk() (*[]TradeDiscountDto, bool)`

GetTradeDiscountsOk returns a tuple with the TradeDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeDiscounts

`func (o *PriceInformationDto) SetTradeDiscounts(v []TradeDiscountDto)`

SetTradeDiscounts sets TradeDiscounts field to given value.

### HasTradeDiscounts

`func (o *PriceInformationDto) HasTradeDiscounts() bool`

HasTradeDiscounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


