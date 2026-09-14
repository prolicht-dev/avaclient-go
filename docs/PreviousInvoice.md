# PreviousInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceDate** | Pointer to **time.Time** | BT-26 The date of the previous invoice | [optional] 
**InvoiceNumber** | **string** | BT-25 The number of the previous invoice | 

## Methods

### NewPreviousInvoice

`func NewPreviousInvoice(invoiceNumber string, ) *PreviousInvoice`

NewPreviousInvoice instantiates a new PreviousInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviousInvoiceWithDefaults

`func NewPreviousInvoiceWithDefaults() *PreviousInvoice`

NewPreviousInvoiceWithDefaults instantiates a new PreviousInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceDate

`func (o *PreviousInvoice) GetInvoiceDate() time.Time`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *PreviousInvoice) GetInvoiceDateOk() (*time.Time, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *PreviousInvoice) SetInvoiceDate(v time.Time)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *PreviousInvoice) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *PreviousInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *PreviousInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *PreviousInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


