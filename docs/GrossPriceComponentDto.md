# GrossPriceComponentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxRate** | **float32** | This components tax rate. | 
**DeductionFactor** | **float32** | This is the factor of applied deductions for this component | 
**TotalNet** | **float32** | The total net price for all components of a given tax rate. | 
**TotalDeducted** | **float32** | The resulting price component after applied deductions | [readonly] 
**TotalTax** | **float32** | The total tax amount for all components of a given tax rate. | [readonly] 
**TotalGross** | **float32** | The total gross price for all components of a given tax rate. | [readonly] 
**TotalGrossDeducted** | **float32** | The total gross price for all components of a given tax rate, after applied deductions. | [readonly] 
**TotalTaxDeducted** | **float32** | The total tax amount for all components of a given tax rate, after applied deductions. | [readonly] 

## Methods

### NewGrossPriceComponentDto

`func NewGrossPriceComponentDto(taxRate float32, deductionFactor float32, totalNet float32, totalDeducted float32, totalTax float32, totalGross float32, totalGrossDeducted float32, totalTaxDeducted float32, ) *GrossPriceComponentDto`

NewGrossPriceComponentDto instantiates a new GrossPriceComponentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrossPriceComponentDtoWithDefaults

`func NewGrossPriceComponentDtoWithDefaults() *GrossPriceComponentDto`

NewGrossPriceComponentDtoWithDefaults instantiates a new GrossPriceComponentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxRate

`func (o *GrossPriceComponentDto) GetTaxRate() float32`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GrossPriceComponentDto) GetTaxRateOk() (*float32, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GrossPriceComponentDto) SetTaxRate(v float32)`

SetTaxRate sets TaxRate field to given value.


### GetDeductionFactor

`func (o *GrossPriceComponentDto) GetDeductionFactor() float32`

GetDeductionFactor returns the DeductionFactor field if non-nil, zero value otherwise.

### GetDeductionFactorOk

`func (o *GrossPriceComponentDto) GetDeductionFactorOk() (*float32, bool)`

GetDeductionFactorOk returns a tuple with the DeductionFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeductionFactor

`func (o *GrossPriceComponentDto) SetDeductionFactor(v float32)`

SetDeductionFactor sets DeductionFactor field to given value.


### GetTotalNet

`func (o *GrossPriceComponentDto) GetTotalNet() float32`

GetTotalNet returns the TotalNet field if non-nil, zero value otherwise.

### GetTotalNetOk

`func (o *GrossPriceComponentDto) GetTotalNetOk() (*float32, bool)`

GetTotalNetOk returns a tuple with the TotalNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNet

`func (o *GrossPriceComponentDto) SetTotalNet(v float32)`

SetTotalNet sets TotalNet field to given value.


### GetTotalDeducted

`func (o *GrossPriceComponentDto) GetTotalDeducted() float32`

GetTotalDeducted returns the TotalDeducted field if non-nil, zero value otherwise.

### GetTotalDeductedOk

`func (o *GrossPriceComponentDto) GetTotalDeductedOk() (*float32, bool)`

GetTotalDeductedOk returns a tuple with the TotalDeducted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeducted

`func (o *GrossPriceComponentDto) SetTotalDeducted(v float32)`

SetTotalDeducted sets TotalDeducted field to given value.


### GetTotalTax

`func (o *GrossPriceComponentDto) GetTotalTax() float32`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *GrossPriceComponentDto) GetTotalTaxOk() (*float32, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *GrossPriceComponentDto) SetTotalTax(v float32)`

SetTotalTax sets TotalTax field to given value.


### GetTotalGross

`func (o *GrossPriceComponentDto) GetTotalGross() float32`

GetTotalGross returns the TotalGross field if non-nil, zero value otherwise.

### GetTotalGrossOk

`func (o *GrossPriceComponentDto) GetTotalGrossOk() (*float32, bool)`

GetTotalGrossOk returns a tuple with the TotalGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGross

`func (o *GrossPriceComponentDto) SetTotalGross(v float32)`

SetTotalGross sets TotalGross field to given value.


### GetTotalGrossDeducted

`func (o *GrossPriceComponentDto) GetTotalGrossDeducted() float32`

GetTotalGrossDeducted returns the TotalGrossDeducted field if non-nil, zero value otherwise.

### GetTotalGrossDeductedOk

`func (o *GrossPriceComponentDto) GetTotalGrossDeductedOk() (*float32, bool)`

GetTotalGrossDeductedOk returns a tuple with the TotalGrossDeducted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossDeducted

`func (o *GrossPriceComponentDto) SetTotalGrossDeducted(v float32)`

SetTotalGrossDeducted sets TotalGrossDeducted field to given value.


### GetTotalTaxDeducted

`func (o *GrossPriceComponentDto) GetTotalTaxDeducted() float32`

GetTotalTaxDeducted returns the TotalTaxDeducted field if non-nil, zero value otherwise.

### GetTotalTaxDeductedOk

`func (o *GrossPriceComponentDto) GetTotalTaxDeductedOk() (*float32, bool)`

GetTotalTaxDeductedOk returns a tuple with the TotalTaxDeducted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxDeducted

`func (o *GrossPriceComponentDto) SetTotalTaxDeducted(v float32)`

SetTotalTaxDeducted sets TotalTaxDeducted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


