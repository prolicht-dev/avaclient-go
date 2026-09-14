# InvoiceLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | BT-126 An unique identifier for the line item | [optional] 
**Note** | Pointer to **string** | BT-127 An unstructured note to classify line items | [optional] 
**ObjectIdentifier** | Pointer to **string** | BT-128 A seller-defined object identifier for this line item, meaning an identifier for the actual line. This could be a numbering sequence or a reference to a product. | [optional] 
**Quantity** | Pointer to **float32** | BT-129 The actual invoiced quantity of the line item | [optional] 
**UnitCode** | Pointer to **string** | BT-130 A code for the unit of the quantity. This should be a code as defined in UN/ECE Recommendation No. 20 „Codes for Units of Measure Used in  International Trade“ und UN/ECE Recommendation No 21 „Codes for Passengers, Types of Cargo, Packages and Packaging Materials(with Complementary Codes for Package Names). See here: https://ec.europa.eu/digital-building-blocks/sites/display/DIGITAL/Registry+of+supporting+artefacts+to+implement+EN16931 | [optional] 
**NetAmount** | Pointer to **float32** | BT-131 The net amount for this line item. This is the total price for positions | [optional] 
**PurchaseOrderReference** | Pointer to **string** | BT-132 A reference to an item in the buyers purachase order | [optional] 
**LineItemPeriod** | Pointer to [**InvoicePeriod**](InvoicePeriod.md) |  | [optional] 
**Allowances** | Pointer to [**[]LineItemAllowance**](LineItemAllowance.md) | BG-27 Allowances specific to this line item | [optional] 
**Charges** | Pointer to [**[]LineItemCharge**](LineItemCharge.md) | BG-28 charges specific to this line item | [optional] 
**PriceDetails** | Pointer to [**LineItemPriceDetail**](LineItemPriceDetail.md) |  | [optional] 
**VatInformation** | Pointer to [**[]LineItemVatInformation**](LineItemVatInformation.md) | BG-30 VAT information for this line item | [optional] 
**ItemInformation** | Pointer to [**LineItemInformation**](LineItemInformation.md) |  | [optional] 
**SubItems** | Pointer to [**[]InvoiceLineItem**](InvoiceLineItem.md) | BG-DEX-01 Line items can be nested, in which case this property holds the sub items | [optional] 

## Methods

### NewInvoiceLineItem

`func NewInvoiceLineItem() *InvoiceLineItem`

NewInvoiceLineItem instantiates a new InvoiceLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineItemWithDefaults

`func NewInvoiceLineItemWithDefaults() *InvoiceLineItem`

NewInvoiceLineItemWithDefaults instantiates a new InvoiceLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *InvoiceLineItem) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *InvoiceLineItem) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *InvoiceLineItem) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *InvoiceLineItem) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetNote

`func (o *InvoiceLineItem) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *InvoiceLineItem) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *InvoiceLineItem) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *InvoiceLineItem) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetObjectIdentifier

`func (o *InvoiceLineItem) GetObjectIdentifier() string`

GetObjectIdentifier returns the ObjectIdentifier field if non-nil, zero value otherwise.

### GetObjectIdentifierOk

`func (o *InvoiceLineItem) GetObjectIdentifierOk() (*string, bool)`

GetObjectIdentifierOk returns a tuple with the ObjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIdentifier

`func (o *InvoiceLineItem) SetObjectIdentifier(v string)`

SetObjectIdentifier sets ObjectIdentifier field to given value.

### HasObjectIdentifier

`func (o *InvoiceLineItem) HasObjectIdentifier() bool`

HasObjectIdentifier returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceLineItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLineItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLineItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitCode

`func (o *InvoiceLineItem) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *InvoiceLineItem) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *InvoiceLineItem) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *InvoiceLineItem) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetNetAmount

`func (o *InvoiceLineItem) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *InvoiceLineItem) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *InvoiceLineItem) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *InvoiceLineItem) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetPurchaseOrderReference

`func (o *InvoiceLineItem) GetPurchaseOrderReference() string`

GetPurchaseOrderReference returns the PurchaseOrderReference field if non-nil, zero value otherwise.

### GetPurchaseOrderReferenceOk

`func (o *InvoiceLineItem) GetPurchaseOrderReferenceOk() (*string, bool)`

GetPurchaseOrderReferenceOk returns a tuple with the PurchaseOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderReference

`func (o *InvoiceLineItem) SetPurchaseOrderReference(v string)`

SetPurchaseOrderReference sets PurchaseOrderReference field to given value.

### HasPurchaseOrderReference

`func (o *InvoiceLineItem) HasPurchaseOrderReference() bool`

HasPurchaseOrderReference returns a boolean if a field has been set.

### GetLineItemPeriod

`func (o *InvoiceLineItem) GetLineItemPeriod() InvoicePeriod`

GetLineItemPeriod returns the LineItemPeriod field if non-nil, zero value otherwise.

### GetLineItemPeriodOk

`func (o *InvoiceLineItem) GetLineItemPeriodOk() (*InvoicePeriod, bool)`

GetLineItemPeriodOk returns a tuple with the LineItemPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemPeriod

`func (o *InvoiceLineItem) SetLineItemPeriod(v InvoicePeriod)`

SetLineItemPeriod sets LineItemPeriod field to given value.

### HasLineItemPeriod

`func (o *InvoiceLineItem) HasLineItemPeriod() bool`

HasLineItemPeriod returns a boolean if a field has been set.

### GetAllowances

`func (o *InvoiceLineItem) GetAllowances() []LineItemAllowance`

GetAllowances returns the Allowances field if non-nil, zero value otherwise.

### GetAllowancesOk

`func (o *InvoiceLineItem) GetAllowancesOk() (*[]LineItemAllowance, bool)`

GetAllowancesOk returns a tuple with the Allowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowances

`func (o *InvoiceLineItem) SetAllowances(v []LineItemAllowance)`

SetAllowances sets Allowances field to given value.

### HasAllowances

`func (o *InvoiceLineItem) HasAllowances() bool`

HasAllowances returns a boolean if a field has been set.

### GetCharges

`func (o *InvoiceLineItem) GetCharges() []LineItemCharge`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *InvoiceLineItem) GetChargesOk() (*[]LineItemCharge, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *InvoiceLineItem) SetCharges(v []LineItemCharge)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *InvoiceLineItem) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetPriceDetails

`func (o *InvoiceLineItem) GetPriceDetails() LineItemPriceDetail`

GetPriceDetails returns the PriceDetails field if non-nil, zero value otherwise.

### GetPriceDetailsOk

`func (o *InvoiceLineItem) GetPriceDetailsOk() (*LineItemPriceDetail, bool)`

GetPriceDetailsOk returns a tuple with the PriceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDetails

`func (o *InvoiceLineItem) SetPriceDetails(v LineItemPriceDetail)`

SetPriceDetails sets PriceDetails field to given value.

### HasPriceDetails

`func (o *InvoiceLineItem) HasPriceDetails() bool`

HasPriceDetails returns a boolean if a field has been set.

### GetVatInformation

`func (o *InvoiceLineItem) GetVatInformation() []LineItemVatInformation`

GetVatInformation returns the VatInformation field if non-nil, zero value otherwise.

### GetVatInformationOk

`func (o *InvoiceLineItem) GetVatInformationOk() (*[]LineItemVatInformation, bool)`

GetVatInformationOk returns a tuple with the VatInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatInformation

`func (o *InvoiceLineItem) SetVatInformation(v []LineItemVatInformation)`

SetVatInformation sets VatInformation field to given value.

### HasVatInformation

`func (o *InvoiceLineItem) HasVatInformation() bool`

HasVatInformation returns a boolean if a field has been set.

### GetItemInformation

`func (o *InvoiceLineItem) GetItemInformation() LineItemInformation`

GetItemInformation returns the ItemInformation field if non-nil, zero value otherwise.

### GetItemInformationOk

`func (o *InvoiceLineItem) GetItemInformationOk() (*LineItemInformation, bool)`

GetItemInformationOk returns a tuple with the ItemInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemInformation

`func (o *InvoiceLineItem) SetItemInformation(v LineItemInformation)`

SetItemInformation sets ItemInformation field to given value.

### HasItemInformation

`func (o *InvoiceLineItem) HasItemInformation() bool`

HasItemInformation returns a boolean if a field has been set.

### GetSubItems

`func (o *InvoiceLineItem) GetSubItems() []InvoiceLineItem`

GetSubItems returns the SubItems field if non-nil, zero value otherwise.

### GetSubItemsOk

`func (o *InvoiceLineItem) GetSubItemsOk() (*[]InvoiceLineItem, bool)`

GetSubItemsOk returns a tuple with the SubItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubItems

`func (o *InvoiceLineItem) SetSubItems(v []InvoiceLineItem)`

SetSubItems sets SubItems field to given value.

### HasSubItems

`func (o *InvoiceLineItem) HasSubItems() bool`

HasSubItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


