# InvoiceTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNet** | Pointer to **float32** | BT-106 Sum of all invoice items in net. BT-106 in XRechnung. | [optional] 
**TotalAllowances** | Pointer to **float32** | BT-107 This is the sum of all allowances (e.g. deductions or withheld tax) in the invoice BT-107 in XRechnung. | [optional] 
**TotalCharges** | Pointer to **float32** | BT-108 This is the sum of all charges in the invoice, except VAT. For example, late fees or surcharges. BT-108 in XRechnung. | [optional] 
**TotalAfterDeductions** | Pointer to **float32** | BT-109 This is the total net value of the invoice after applied allowances and charges, meaning the actual net sum to be paid. BT-109 in XRechnung. | [optional] 
**TotalVatAmount** | Pointer to **float32** | BT-110 The total amount of VAT in the invoice. BT-110 in XRechnung. | [optional] 
**TotalGross** | Pointer to **float32** | BT-112 The total gross amount of the invoice, including all taxes and deductions. BT-112 in XRechnung. | [optional] 
**AlreadyPaidTotal** | Pointer to **float32** | BT-113 The total amount that has already been paid for this invoice, e.g. used when previous invoice sums are included here that have already been paid. BT-113 in XRechnung. | [optional] 
**PayableRoundingAmount** | Pointer to **float32** | BT-114 If this is given, this specifies if any rounding amount is applied to the invoice, before calculating the TotalToBePaid value. BT-114 in XRechnung. | [optional] 
**TotalToBePaid** | Pointer to **float32** | BT-115 The sum that needs to be paid for this invoice. This is the total gross amount minus the already paid amount. BT-115 in XRechnung. | [optional] 

## Methods

### NewInvoiceTotals

`func NewInvoiceTotals() *InvoiceTotals`

NewInvoiceTotals instantiates a new InvoiceTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceTotalsWithDefaults

`func NewInvoiceTotalsWithDefaults() *InvoiceTotals`

NewInvoiceTotalsWithDefaults instantiates a new InvoiceTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNet

`func (o *InvoiceTotals) GetTotalNet() float32`

GetTotalNet returns the TotalNet field if non-nil, zero value otherwise.

### GetTotalNetOk

`func (o *InvoiceTotals) GetTotalNetOk() (*float32, bool)`

GetTotalNetOk returns a tuple with the TotalNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNet

`func (o *InvoiceTotals) SetTotalNet(v float32)`

SetTotalNet sets TotalNet field to given value.

### HasTotalNet

`func (o *InvoiceTotals) HasTotalNet() bool`

HasTotalNet returns a boolean if a field has been set.

### GetTotalAllowances

`func (o *InvoiceTotals) GetTotalAllowances() float32`

GetTotalAllowances returns the TotalAllowances field if non-nil, zero value otherwise.

### GetTotalAllowancesOk

`func (o *InvoiceTotals) GetTotalAllowancesOk() (*float32, bool)`

GetTotalAllowancesOk returns a tuple with the TotalAllowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAllowances

`func (o *InvoiceTotals) SetTotalAllowances(v float32)`

SetTotalAllowances sets TotalAllowances field to given value.

### HasTotalAllowances

`func (o *InvoiceTotals) HasTotalAllowances() bool`

HasTotalAllowances returns a boolean if a field has been set.

### GetTotalCharges

`func (o *InvoiceTotals) GetTotalCharges() float32`

GetTotalCharges returns the TotalCharges field if non-nil, zero value otherwise.

### GetTotalChargesOk

`func (o *InvoiceTotals) GetTotalChargesOk() (*float32, bool)`

GetTotalChargesOk returns a tuple with the TotalCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCharges

`func (o *InvoiceTotals) SetTotalCharges(v float32)`

SetTotalCharges sets TotalCharges field to given value.

### HasTotalCharges

`func (o *InvoiceTotals) HasTotalCharges() bool`

HasTotalCharges returns a boolean if a field has been set.

### GetTotalAfterDeductions

`func (o *InvoiceTotals) GetTotalAfterDeductions() float32`

GetTotalAfterDeductions returns the TotalAfterDeductions field if non-nil, zero value otherwise.

### GetTotalAfterDeductionsOk

`func (o *InvoiceTotals) GetTotalAfterDeductionsOk() (*float32, bool)`

GetTotalAfterDeductionsOk returns a tuple with the TotalAfterDeductions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAfterDeductions

`func (o *InvoiceTotals) SetTotalAfterDeductions(v float32)`

SetTotalAfterDeductions sets TotalAfterDeductions field to given value.

### HasTotalAfterDeductions

`func (o *InvoiceTotals) HasTotalAfterDeductions() bool`

HasTotalAfterDeductions returns a boolean if a field has been set.

### GetTotalVatAmount

`func (o *InvoiceTotals) GetTotalVatAmount() float32`

GetTotalVatAmount returns the TotalVatAmount field if non-nil, zero value otherwise.

### GetTotalVatAmountOk

`func (o *InvoiceTotals) GetTotalVatAmountOk() (*float32, bool)`

GetTotalVatAmountOk returns a tuple with the TotalVatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVatAmount

`func (o *InvoiceTotals) SetTotalVatAmount(v float32)`

SetTotalVatAmount sets TotalVatAmount field to given value.

### HasTotalVatAmount

`func (o *InvoiceTotals) HasTotalVatAmount() bool`

HasTotalVatAmount returns a boolean if a field has been set.

### GetTotalGross

`func (o *InvoiceTotals) GetTotalGross() float32`

GetTotalGross returns the TotalGross field if non-nil, zero value otherwise.

### GetTotalGrossOk

`func (o *InvoiceTotals) GetTotalGrossOk() (*float32, bool)`

GetTotalGrossOk returns a tuple with the TotalGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGross

`func (o *InvoiceTotals) SetTotalGross(v float32)`

SetTotalGross sets TotalGross field to given value.

### HasTotalGross

`func (o *InvoiceTotals) HasTotalGross() bool`

HasTotalGross returns a boolean if a field has been set.

### GetAlreadyPaidTotal

`func (o *InvoiceTotals) GetAlreadyPaidTotal() float32`

GetAlreadyPaidTotal returns the AlreadyPaidTotal field if non-nil, zero value otherwise.

### GetAlreadyPaidTotalOk

`func (o *InvoiceTotals) GetAlreadyPaidTotalOk() (*float32, bool)`

GetAlreadyPaidTotalOk returns a tuple with the AlreadyPaidTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyPaidTotal

`func (o *InvoiceTotals) SetAlreadyPaidTotal(v float32)`

SetAlreadyPaidTotal sets AlreadyPaidTotal field to given value.

### HasAlreadyPaidTotal

`func (o *InvoiceTotals) HasAlreadyPaidTotal() bool`

HasAlreadyPaidTotal returns a boolean if a field has been set.

### GetPayableRoundingAmount

`func (o *InvoiceTotals) GetPayableRoundingAmount() float32`

GetPayableRoundingAmount returns the PayableRoundingAmount field if non-nil, zero value otherwise.

### GetPayableRoundingAmountOk

`func (o *InvoiceTotals) GetPayableRoundingAmountOk() (*float32, bool)`

GetPayableRoundingAmountOk returns a tuple with the PayableRoundingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableRoundingAmount

`func (o *InvoiceTotals) SetPayableRoundingAmount(v float32)`

SetPayableRoundingAmount sets PayableRoundingAmount field to given value.

### HasPayableRoundingAmount

`func (o *InvoiceTotals) HasPayableRoundingAmount() bool`

HasPayableRoundingAmount returns a boolean if a field has been set.

### GetTotalToBePaid

`func (o *InvoiceTotals) GetTotalToBePaid() float32`

GetTotalToBePaid returns the TotalToBePaid field if non-nil, zero value otherwise.

### GetTotalToBePaidOk

`func (o *InvoiceTotals) GetTotalToBePaidOk() (*float32, bool)`

GetTotalToBePaidOk returns a tuple with the TotalToBePaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalToBePaid

`func (o *InvoiceTotals) SetTotalToBePaid(v float32)`

SetTotalToBePaid sets TotalToBePaid field to given value.

### HasTotalToBePaid

`func (o *InvoiceTotals) HasTotalToBePaid() bool`

HasTotalToBePaid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


