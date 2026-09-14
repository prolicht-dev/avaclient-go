# References

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerReference** | Pointer to **string** | BT-10 An internal reference used by the buyer, this can also be used as the &#39;Leitweg-ID&#39; in Germany BT-10 in XRechnung. | [optional] 
**ProjectReference** | Pointer to **string** | BT-11 Reference to a project BT-11 in XRechnung. | [optional] 
**ContractReference** | Pointer to **string** | BT-12 Reference to a contract BT-12 in XRechnung. | [optional] 
**PurchaseOrderReference** | Pointer to **string** | BT-13 Reference to a purchase order created by the buyer BT-13 in XRechnung. | [optional] 
**SalesOrderReference** | Pointer to **string** | BT-14 Reference to a sales order created by the seller BT-14 in XRechnung. | [optional] 
**ReceivingAdviceReference** | Pointer to **string** | BT-15 Reference to a delivery note BT-15 in XRechnung. | [optional] 
**DispatchAdviceReference** | Pointer to **string** | BT-16 Reference to a dispatch advice BT-16 in XRechnung. | [optional] 
**TenderOrLotReference** | Pointer to **string** | BT-17 Reference to a tender or lot BT-17 in XRechnung. | [optional] 
**InvoicedObjectIdentifier** | Pointer to **string** | BT-18 Reference to a custom item, e.g. a phone number, person, contract or anything that would identify the invoice in the buyer&#39;s system BT-18 in XRechnung. | [optional] 
**AccountingReference** | Pointer to **string** | BT-19 This is a reference for accounting, like a cost center or similar. In German, this would be a \&quot;Kostenstelle\&quot;. BT-19 in XRechnung. | [optional] 

## Methods

### NewReferences

`func NewReferences() *References`

NewReferences instantiates a new References object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferencesWithDefaults

`func NewReferencesWithDefaults() *References`

NewReferencesWithDefaults instantiates a new References object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerReference

`func (o *References) GetBuyerReference() string`

GetBuyerReference returns the BuyerReference field if non-nil, zero value otherwise.

### GetBuyerReferenceOk

`func (o *References) GetBuyerReferenceOk() (*string, bool)`

GetBuyerReferenceOk returns a tuple with the BuyerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerReference

`func (o *References) SetBuyerReference(v string)`

SetBuyerReference sets BuyerReference field to given value.

### HasBuyerReference

`func (o *References) HasBuyerReference() bool`

HasBuyerReference returns a boolean if a field has been set.

### GetProjectReference

`func (o *References) GetProjectReference() string`

GetProjectReference returns the ProjectReference field if non-nil, zero value otherwise.

### GetProjectReferenceOk

`func (o *References) GetProjectReferenceOk() (*string, bool)`

GetProjectReferenceOk returns a tuple with the ProjectReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectReference

`func (o *References) SetProjectReference(v string)`

SetProjectReference sets ProjectReference field to given value.

### HasProjectReference

`func (o *References) HasProjectReference() bool`

HasProjectReference returns a boolean if a field has been set.

### GetContractReference

`func (o *References) GetContractReference() string`

GetContractReference returns the ContractReference field if non-nil, zero value otherwise.

### GetContractReferenceOk

`func (o *References) GetContractReferenceOk() (*string, bool)`

GetContractReferenceOk returns a tuple with the ContractReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractReference

`func (o *References) SetContractReference(v string)`

SetContractReference sets ContractReference field to given value.

### HasContractReference

`func (o *References) HasContractReference() bool`

HasContractReference returns a boolean if a field has been set.

### GetPurchaseOrderReference

`func (o *References) GetPurchaseOrderReference() string`

GetPurchaseOrderReference returns the PurchaseOrderReference field if non-nil, zero value otherwise.

### GetPurchaseOrderReferenceOk

`func (o *References) GetPurchaseOrderReferenceOk() (*string, bool)`

GetPurchaseOrderReferenceOk returns a tuple with the PurchaseOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderReference

`func (o *References) SetPurchaseOrderReference(v string)`

SetPurchaseOrderReference sets PurchaseOrderReference field to given value.

### HasPurchaseOrderReference

`func (o *References) HasPurchaseOrderReference() bool`

HasPurchaseOrderReference returns a boolean if a field has been set.

### GetSalesOrderReference

`func (o *References) GetSalesOrderReference() string`

GetSalesOrderReference returns the SalesOrderReference field if non-nil, zero value otherwise.

### GetSalesOrderReferenceOk

`func (o *References) GetSalesOrderReferenceOk() (*string, bool)`

GetSalesOrderReferenceOk returns a tuple with the SalesOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrderReference

`func (o *References) SetSalesOrderReference(v string)`

SetSalesOrderReference sets SalesOrderReference field to given value.

### HasSalesOrderReference

`func (o *References) HasSalesOrderReference() bool`

HasSalesOrderReference returns a boolean if a field has been set.

### GetReceivingAdviceReference

`func (o *References) GetReceivingAdviceReference() string`

GetReceivingAdviceReference returns the ReceivingAdviceReference field if non-nil, zero value otherwise.

### GetReceivingAdviceReferenceOk

`func (o *References) GetReceivingAdviceReferenceOk() (*string, bool)`

GetReceivingAdviceReferenceOk returns a tuple with the ReceivingAdviceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAdviceReference

`func (o *References) SetReceivingAdviceReference(v string)`

SetReceivingAdviceReference sets ReceivingAdviceReference field to given value.

### HasReceivingAdviceReference

`func (o *References) HasReceivingAdviceReference() bool`

HasReceivingAdviceReference returns a boolean if a field has been set.

### GetDispatchAdviceReference

`func (o *References) GetDispatchAdviceReference() string`

GetDispatchAdviceReference returns the DispatchAdviceReference field if non-nil, zero value otherwise.

### GetDispatchAdviceReferenceOk

`func (o *References) GetDispatchAdviceReferenceOk() (*string, bool)`

GetDispatchAdviceReferenceOk returns a tuple with the DispatchAdviceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchAdviceReference

`func (o *References) SetDispatchAdviceReference(v string)`

SetDispatchAdviceReference sets DispatchAdviceReference field to given value.

### HasDispatchAdviceReference

`func (o *References) HasDispatchAdviceReference() bool`

HasDispatchAdviceReference returns a boolean if a field has been set.

### GetTenderOrLotReference

`func (o *References) GetTenderOrLotReference() string`

GetTenderOrLotReference returns the TenderOrLotReference field if non-nil, zero value otherwise.

### GetTenderOrLotReferenceOk

`func (o *References) GetTenderOrLotReferenceOk() (*string, bool)`

GetTenderOrLotReferenceOk returns a tuple with the TenderOrLotReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderOrLotReference

`func (o *References) SetTenderOrLotReference(v string)`

SetTenderOrLotReference sets TenderOrLotReference field to given value.

### HasTenderOrLotReference

`func (o *References) HasTenderOrLotReference() bool`

HasTenderOrLotReference returns a boolean if a field has been set.

### GetInvoicedObjectIdentifier

`func (o *References) GetInvoicedObjectIdentifier() string`

GetInvoicedObjectIdentifier returns the InvoicedObjectIdentifier field if non-nil, zero value otherwise.

### GetInvoicedObjectIdentifierOk

`func (o *References) GetInvoicedObjectIdentifierOk() (*string, bool)`

GetInvoicedObjectIdentifierOk returns a tuple with the InvoicedObjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicedObjectIdentifier

`func (o *References) SetInvoicedObjectIdentifier(v string)`

SetInvoicedObjectIdentifier sets InvoicedObjectIdentifier field to given value.

### HasInvoicedObjectIdentifier

`func (o *References) HasInvoicedObjectIdentifier() bool`

HasInvoicedObjectIdentifier returns a boolean if a field has been set.

### GetAccountingReference

`func (o *References) GetAccountingReference() string`

GetAccountingReference returns the AccountingReference field if non-nil, zero value otherwise.

### GetAccountingReferenceOk

`func (o *References) GetAccountingReferenceOk() (*string, bool)`

GetAccountingReferenceOk returns a tuple with the AccountingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingReference

`func (o *References) SetAccountingReference(v string)`

SetAccountingReference sets AccountingReference field to given value.

### HasAccountingReference

`func (o *References) HasAccountingReference() bool`

HasAccountingReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


