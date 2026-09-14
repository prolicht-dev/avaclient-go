# FacturXPaymentTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A text description of the payment discount term. BT-20 in XRechnung, but this parent object is only used for Factur-X / ZUGFeRD and holds more detailled information about the payment term. | [optional] 
**DueDate** | Pointer to **time.Time** | The date until which the payment term applies. | [optional] 
**BaseAmount** | Pointer to **float32** | The base amount unto which the discount applies. | [optional] 
**ActualDiscountAmount** | Pointer to **float32** | The actual discount amount that will be deducted if the payment is done within the specified period. | [optional] 
**BaseDate** | Pointer to **time.Time** | The base date for calculating the discount period. This is typically the invoice date. | [optional] 
**PeriodMeasure** | Pointer to **float32** | The duration of the discount period, e.g. \&quot;30\&quot; for 30 days. | [optional] 
**PeriodMeasureUnitCode** | Pointer to **string** | The unit in which the period is specified, e.g. \&quot;DAY\&quot;. | [optional] 
**CalculationPercent** | Pointer to **float32** | The percentage of the discount, e.g. \&quot;.02\&quot; for 2%. | [optional] 

## Methods

### NewFacturXPaymentTerm

`func NewFacturXPaymentTerm() *FacturXPaymentTerm`

NewFacturXPaymentTerm instantiates a new FacturXPaymentTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacturXPaymentTermWithDefaults

`func NewFacturXPaymentTermWithDefaults() *FacturXPaymentTerm`

NewFacturXPaymentTermWithDefaults instantiates a new FacturXPaymentTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FacturXPaymentTerm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FacturXPaymentTerm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FacturXPaymentTerm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FacturXPaymentTerm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *FacturXPaymentTerm) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *FacturXPaymentTerm) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *FacturXPaymentTerm) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *FacturXPaymentTerm) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetBaseAmount

`func (o *FacturXPaymentTerm) GetBaseAmount() float32`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *FacturXPaymentTerm) GetBaseAmountOk() (*float32, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *FacturXPaymentTerm) SetBaseAmount(v float32)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *FacturXPaymentTerm) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetActualDiscountAmount

`func (o *FacturXPaymentTerm) GetActualDiscountAmount() float32`

GetActualDiscountAmount returns the ActualDiscountAmount field if non-nil, zero value otherwise.

### GetActualDiscountAmountOk

`func (o *FacturXPaymentTerm) GetActualDiscountAmountOk() (*float32, bool)`

GetActualDiscountAmountOk returns a tuple with the ActualDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDiscountAmount

`func (o *FacturXPaymentTerm) SetActualDiscountAmount(v float32)`

SetActualDiscountAmount sets ActualDiscountAmount field to given value.

### HasActualDiscountAmount

`func (o *FacturXPaymentTerm) HasActualDiscountAmount() bool`

HasActualDiscountAmount returns a boolean if a field has been set.

### GetBaseDate

`func (o *FacturXPaymentTerm) GetBaseDate() time.Time`

GetBaseDate returns the BaseDate field if non-nil, zero value otherwise.

### GetBaseDateOk

`func (o *FacturXPaymentTerm) GetBaseDateOk() (*time.Time, bool)`

GetBaseDateOk returns a tuple with the BaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDate

`func (o *FacturXPaymentTerm) SetBaseDate(v time.Time)`

SetBaseDate sets BaseDate field to given value.

### HasBaseDate

`func (o *FacturXPaymentTerm) HasBaseDate() bool`

HasBaseDate returns a boolean if a field has been set.

### GetPeriodMeasure

`func (o *FacturXPaymentTerm) GetPeriodMeasure() float32`

GetPeriodMeasure returns the PeriodMeasure field if non-nil, zero value otherwise.

### GetPeriodMeasureOk

`func (o *FacturXPaymentTerm) GetPeriodMeasureOk() (*float32, bool)`

GetPeriodMeasureOk returns a tuple with the PeriodMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodMeasure

`func (o *FacturXPaymentTerm) SetPeriodMeasure(v float32)`

SetPeriodMeasure sets PeriodMeasure field to given value.

### HasPeriodMeasure

`func (o *FacturXPaymentTerm) HasPeriodMeasure() bool`

HasPeriodMeasure returns a boolean if a field has been set.

### GetPeriodMeasureUnitCode

`func (o *FacturXPaymentTerm) GetPeriodMeasureUnitCode() string`

GetPeriodMeasureUnitCode returns the PeriodMeasureUnitCode field if non-nil, zero value otherwise.

### GetPeriodMeasureUnitCodeOk

`func (o *FacturXPaymentTerm) GetPeriodMeasureUnitCodeOk() (*string, bool)`

GetPeriodMeasureUnitCodeOk returns a tuple with the PeriodMeasureUnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodMeasureUnitCode

`func (o *FacturXPaymentTerm) SetPeriodMeasureUnitCode(v string)`

SetPeriodMeasureUnitCode sets PeriodMeasureUnitCode field to given value.

### HasPeriodMeasureUnitCode

`func (o *FacturXPaymentTerm) HasPeriodMeasureUnitCode() bool`

HasPeriodMeasureUnitCode returns a boolean if a field has been set.

### GetCalculationPercent

`func (o *FacturXPaymentTerm) GetCalculationPercent() float32`

GetCalculationPercent returns the CalculationPercent field if non-nil, zero value otherwise.

### GetCalculationPercentOk

`func (o *FacturXPaymentTerm) GetCalculationPercentOk() (*float32, bool)`

GetCalculationPercentOk returns a tuple with the CalculationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculationPercent

`func (o *FacturXPaymentTerm) SetCalculationPercent(v float32)`

SetCalculationPercent sets CalculationPercent field to given value.

### HasCalculationPercent

`func (o *FacturXPaymentTerm) HasCalculationPercent() bool`

HasCalculationPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


