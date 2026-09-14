# LineItemPriceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetPrice** | Pointer to **float32** | BT-146 The net price of this line item. This is the unit price | [optional] 
**AbsoluteDiscountPerItem** | Pointer to **float32** | BT-147 A discount for this line item, given as an absolute value. It is essentially the difference between NetPrice (or the list price) and the actual invoiced price GrossPrice. | [optional] 
**GrossPrice** | Pointer to **float32** | BT-148 The gross price of this line item, meaning the price without taxes that will be invoiced | [optional] 
**BaseQuantity** | Pointer to **float32** | BT-149 The quantity for which this price is valid. In combination with the NetPrice, you basically specify the denominator of the unit price for a line item. | [optional] 
**BaseQuantityUnitCode** | Pointer to **string** | BT-150 A code representing the unit for the base quantity | [optional] 

## Methods

### NewLineItemPriceDetail

`func NewLineItemPriceDetail() *LineItemPriceDetail`

NewLineItemPriceDetail instantiates a new LineItemPriceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemPriceDetailWithDefaults

`func NewLineItemPriceDetailWithDefaults() *LineItemPriceDetail`

NewLineItemPriceDetailWithDefaults instantiates a new LineItemPriceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetPrice

`func (o *LineItemPriceDetail) GetNetPrice() float32`

GetNetPrice returns the NetPrice field if non-nil, zero value otherwise.

### GetNetPriceOk

`func (o *LineItemPriceDetail) GetNetPriceOk() (*float32, bool)`

GetNetPriceOk returns a tuple with the NetPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPrice

`func (o *LineItemPriceDetail) SetNetPrice(v float32)`

SetNetPrice sets NetPrice field to given value.

### HasNetPrice

`func (o *LineItemPriceDetail) HasNetPrice() bool`

HasNetPrice returns a boolean if a field has been set.

### GetAbsoluteDiscountPerItem

`func (o *LineItemPriceDetail) GetAbsoluteDiscountPerItem() float32`

GetAbsoluteDiscountPerItem returns the AbsoluteDiscountPerItem field if non-nil, zero value otherwise.

### GetAbsoluteDiscountPerItemOk

`func (o *LineItemPriceDetail) GetAbsoluteDiscountPerItemOk() (*float32, bool)`

GetAbsoluteDiscountPerItemOk returns a tuple with the AbsoluteDiscountPerItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteDiscountPerItem

`func (o *LineItemPriceDetail) SetAbsoluteDiscountPerItem(v float32)`

SetAbsoluteDiscountPerItem sets AbsoluteDiscountPerItem field to given value.

### HasAbsoluteDiscountPerItem

`func (o *LineItemPriceDetail) HasAbsoluteDiscountPerItem() bool`

HasAbsoluteDiscountPerItem returns a boolean if a field has been set.

### GetGrossPrice

`func (o *LineItemPriceDetail) GetGrossPrice() float32`

GetGrossPrice returns the GrossPrice field if non-nil, zero value otherwise.

### GetGrossPriceOk

`func (o *LineItemPriceDetail) GetGrossPriceOk() (*float32, bool)`

GetGrossPriceOk returns a tuple with the GrossPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossPrice

`func (o *LineItemPriceDetail) SetGrossPrice(v float32)`

SetGrossPrice sets GrossPrice field to given value.

### HasGrossPrice

`func (o *LineItemPriceDetail) HasGrossPrice() bool`

HasGrossPrice returns a boolean if a field has been set.

### GetBaseQuantity

`func (o *LineItemPriceDetail) GetBaseQuantity() float32`

GetBaseQuantity returns the BaseQuantity field if non-nil, zero value otherwise.

### GetBaseQuantityOk

`func (o *LineItemPriceDetail) GetBaseQuantityOk() (*float32, bool)`

GetBaseQuantityOk returns a tuple with the BaseQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseQuantity

`func (o *LineItemPriceDetail) SetBaseQuantity(v float32)`

SetBaseQuantity sets BaseQuantity field to given value.

### HasBaseQuantity

`func (o *LineItemPriceDetail) HasBaseQuantity() bool`

HasBaseQuantity returns a boolean if a field has been set.

### GetBaseQuantityUnitCode

`func (o *LineItemPriceDetail) GetBaseQuantityUnitCode() string`

GetBaseQuantityUnitCode returns the BaseQuantityUnitCode field if non-nil, zero value otherwise.

### GetBaseQuantityUnitCodeOk

`func (o *LineItemPriceDetail) GetBaseQuantityUnitCodeOk() (*string, bool)`

GetBaseQuantityUnitCodeOk returns a tuple with the BaseQuantityUnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseQuantityUnitCode

`func (o *LineItemPriceDetail) SetBaseQuantityUnitCode(v string)`

SetBaseQuantityUnitCode sets BaseQuantityUnitCode field to given value.

### HasBaseQuantityUnitCode

`func (o *LineItemPriceDetail) HasBaseQuantityUnitCode() bool`

HasBaseQuantityUnitCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


