# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNumber** | Pointer to **string** | BT-1 The number of the invoice | [optional] 
**InvoiceDate** | Pointer to **time.Time** | BT-2 The date of the invoice | [optional] 
**InvoiceType** | Pointer to [**InvoiceType**](InvoiceType.md) |  | [optional] 
**CurrencyCode** | Pointer to **string** | BT-5 The three letter ISO currency code | [optional] 
**PaymentDueDate** | Pointer to **time.Time** | BT-9 Until when the payment is due | [optional] 
**References** | Pointer to [**References**](References.md) |  | [optional] 
**PaymentTerms** | Pointer to **string** | BT-20 A text description of payment terms, e.g. cash discounts or payment conditions | [optional] 
**ParsedPaymentTerms** | Pointer to [**[]PaymentTerm**](PaymentTerm.md) | If the PaymentTerms are present and in a correct format, this read-only property can be used to get the actual parsed payment terms in a deserialized and structured format. This will be null if PaymentTerms are not present or not in the correct format | [optional] 
**PreviousInvoices** | Pointer to [**[]PreviousInvoice**](PreviousInvoice.md) | BG-3 In case of partial or final invoices, it&#39;s possible to reference the previous invoice numbers | [optional] 
**Buyer** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Seller** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**SellerTaxRepresentative** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Payee** | Pointer to [**Payee**](Payee.md) |  | [optional] 
**InvoiceNotes** | Pointer to [**[]InvoiceNote**](InvoiceNote.md) | BG-1 If present, this is an unstructured text that can be used to give additional information about the invoice. Field BG-1 | [optional] 
**InvoicePeriod** | Pointer to [**InvoicePeriod**](InvoicePeriod.md) |  | [optional] 
**DeliveryInformation** | Pointer to [**DeliveryInformation**](DeliveryInformation.md) |  | [optional] 
**Totals** | Pointer to [**InvoiceTotals**](InvoiceTotals.md) |  | [optional] 
**PaymentInstructions** | Pointer to [**PaymentInstructions**](PaymentInstructions.md) |  | [optional] 
**Documents** | Pointer to [**[]SupportingDocument**](SupportingDocument.md) | BG-24 Attached documents to the invoice | [optional] 
**VatBreakdown** | Pointer to [**[]VatBreakdown**](VatBreakdown.md) | BG-23 This is used to list VAT by category, e.g. for invoices where parts are in different VAT categories | [optional] 
**Allowances** | Pointer to [**[]InvoiceAllowance**](InvoiceAllowance.md) | BG-20 Allowances / deductions on the invoice | [optional] 
**Charges** | Pointer to [**[]InvoiceCharge**](InvoiceCharge.md) | BG-21 Charges / additional costs on the invoice | [optional] 
**LineItems** | Pointer to [**[]InvoiceLineItem**](InvoiceLineItem.md) | BG-25 Single line items in the invoice | [optional] 
**SourceType** | [**SourceType**](SourceType.md) |  | 
**FacturXData** | Pointer to [**FacturXData**](FacturXData.md) |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice(sourceType SourceType, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNumber

`func (o *Invoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *Invoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *Invoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *Invoice) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *Invoice) GetInvoiceDate() time.Time`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *Invoice) GetInvoiceDateOk() (*time.Time, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *Invoice) SetInvoiceDate(v time.Time)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *Invoice) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetInvoiceType

`func (o *Invoice) GetInvoiceType() InvoiceType`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *Invoice) GetInvoiceTypeOk() (*InvoiceType, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *Invoice) SetInvoiceType(v InvoiceType)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *Invoice) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *Invoice) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Invoice) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Invoice) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *Invoice) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetPaymentDueDate

`func (o *Invoice) GetPaymentDueDate() time.Time`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *Invoice) GetPaymentDueDateOk() (*time.Time, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *Invoice) SetPaymentDueDate(v time.Time)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *Invoice) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### GetReferences

`func (o *Invoice) GetReferences() References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Invoice) GetReferencesOk() (*References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Invoice) SetReferences(v References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Invoice) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *Invoice) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *Invoice) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *Invoice) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *Invoice) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetParsedPaymentTerms

`func (o *Invoice) GetParsedPaymentTerms() []PaymentTerm`

GetParsedPaymentTerms returns the ParsedPaymentTerms field if non-nil, zero value otherwise.

### GetParsedPaymentTermsOk

`func (o *Invoice) GetParsedPaymentTermsOk() (*[]PaymentTerm, bool)`

GetParsedPaymentTermsOk returns a tuple with the ParsedPaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedPaymentTerms

`func (o *Invoice) SetParsedPaymentTerms(v []PaymentTerm)`

SetParsedPaymentTerms sets ParsedPaymentTerms field to given value.

### HasParsedPaymentTerms

`func (o *Invoice) HasParsedPaymentTerms() bool`

HasParsedPaymentTerms returns a boolean if a field has been set.

### GetPreviousInvoices

`func (o *Invoice) GetPreviousInvoices() []PreviousInvoice`

GetPreviousInvoices returns the PreviousInvoices field if non-nil, zero value otherwise.

### GetPreviousInvoicesOk

`func (o *Invoice) GetPreviousInvoicesOk() (*[]PreviousInvoice, bool)`

GetPreviousInvoicesOk returns a tuple with the PreviousInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousInvoices

`func (o *Invoice) SetPreviousInvoices(v []PreviousInvoice)`

SetPreviousInvoices sets PreviousInvoices field to given value.

### HasPreviousInvoices

`func (o *Invoice) HasPreviousInvoices() bool`

HasPreviousInvoices returns a boolean if a field has been set.

### GetBuyer

`func (o *Invoice) GetBuyer() Organization`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *Invoice) GetBuyerOk() (*Organization, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *Invoice) SetBuyer(v Organization)`

SetBuyer sets Buyer field to given value.

### HasBuyer

`func (o *Invoice) HasBuyer() bool`

HasBuyer returns a boolean if a field has been set.

### GetSeller

`func (o *Invoice) GetSeller() Organization`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *Invoice) GetSellerOk() (*Organization, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *Invoice) SetSeller(v Organization)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *Invoice) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetSellerTaxRepresentative

`func (o *Invoice) GetSellerTaxRepresentative() Organization`

GetSellerTaxRepresentative returns the SellerTaxRepresentative field if non-nil, zero value otherwise.

### GetSellerTaxRepresentativeOk

`func (o *Invoice) GetSellerTaxRepresentativeOk() (*Organization, bool)`

GetSellerTaxRepresentativeOk returns a tuple with the SellerTaxRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerTaxRepresentative

`func (o *Invoice) SetSellerTaxRepresentative(v Organization)`

SetSellerTaxRepresentative sets SellerTaxRepresentative field to given value.

### HasSellerTaxRepresentative

`func (o *Invoice) HasSellerTaxRepresentative() bool`

HasSellerTaxRepresentative returns a boolean if a field has been set.

### GetPayee

`func (o *Invoice) GetPayee() Payee`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *Invoice) GetPayeeOk() (*Payee, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *Invoice) SetPayee(v Payee)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *Invoice) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### GetInvoiceNotes

`func (o *Invoice) GetInvoiceNotes() []InvoiceNote`

GetInvoiceNotes returns the InvoiceNotes field if non-nil, zero value otherwise.

### GetInvoiceNotesOk

`func (o *Invoice) GetInvoiceNotesOk() (*[]InvoiceNote, bool)`

GetInvoiceNotesOk returns a tuple with the InvoiceNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNotes

`func (o *Invoice) SetInvoiceNotes(v []InvoiceNote)`

SetInvoiceNotes sets InvoiceNotes field to given value.

### HasInvoiceNotes

`func (o *Invoice) HasInvoiceNotes() bool`

HasInvoiceNotes returns a boolean if a field has been set.

### GetInvoicePeriod

`func (o *Invoice) GetInvoicePeriod() InvoicePeriod`

GetInvoicePeriod returns the InvoicePeriod field if non-nil, zero value otherwise.

### GetInvoicePeriodOk

`func (o *Invoice) GetInvoicePeriodOk() (*InvoicePeriod, bool)`

GetInvoicePeriodOk returns a tuple with the InvoicePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePeriod

`func (o *Invoice) SetInvoicePeriod(v InvoicePeriod)`

SetInvoicePeriod sets InvoicePeriod field to given value.

### HasInvoicePeriod

`func (o *Invoice) HasInvoicePeriod() bool`

HasInvoicePeriod returns a boolean if a field has been set.

### GetDeliveryInformation

`func (o *Invoice) GetDeliveryInformation() DeliveryInformation`

GetDeliveryInformation returns the DeliveryInformation field if non-nil, zero value otherwise.

### GetDeliveryInformationOk

`func (o *Invoice) GetDeliveryInformationOk() (*DeliveryInformation, bool)`

GetDeliveryInformationOk returns a tuple with the DeliveryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryInformation

`func (o *Invoice) SetDeliveryInformation(v DeliveryInformation)`

SetDeliveryInformation sets DeliveryInformation field to given value.

### HasDeliveryInformation

`func (o *Invoice) HasDeliveryInformation() bool`

HasDeliveryInformation returns a boolean if a field has been set.

### GetTotals

`func (o *Invoice) GetTotals() InvoiceTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Invoice) GetTotalsOk() (*InvoiceTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Invoice) SetTotals(v InvoiceTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *Invoice) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetPaymentInstructions

`func (o *Invoice) GetPaymentInstructions() PaymentInstructions`

GetPaymentInstructions returns the PaymentInstructions field if non-nil, zero value otherwise.

### GetPaymentInstructionsOk

`func (o *Invoice) GetPaymentInstructionsOk() (*PaymentInstructions, bool)`

GetPaymentInstructionsOk returns a tuple with the PaymentInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructions

`func (o *Invoice) SetPaymentInstructions(v PaymentInstructions)`

SetPaymentInstructions sets PaymentInstructions field to given value.

### HasPaymentInstructions

`func (o *Invoice) HasPaymentInstructions() bool`

HasPaymentInstructions returns a boolean if a field has been set.

### GetDocuments

`func (o *Invoice) GetDocuments() []SupportingDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Invoice) GetDocumentsOk() (*[]SupportingDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Invoice) SetDocuments(v []SupportingDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Invoice) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetVatBreakdown

`func (o *Invoice) GetVatBreakdown() []VatBreakdown`

GetVatBreakdown returns the VatBreakdown field if non-nil, zero value otherwise.

### GetVatBreakdownOk

`func (o *Invoice) GetVatBreakdownOk() (*[]VatBreakdown, bool)`

GetVatBreakdownOk returns a tuple with the VatBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatBreakdown

`func (o *Invoice) SetVatBreakdown(v []VatBreakdown)`

SetVatBreakdown sets VatBreakdown field to given value.

### HasVatBreakdown

`func (o *Invoice) HasVatBreakdown() bool`

HasVatBreakdown returns a boolean if a field has been set.

### GetAllowances

`func (o *Invoice) GetAllowances() []InvoiceAllowance`

GetAllowances returns the Allowances field if non-nil, zero value otherwise.

### GetAllowancesOk

`func (o *Invoice) GetAllowancesOk() (*[]InvoiceAllowance, bool)`

GetAllowancesOk returns a tuple with the Allowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowances

`func (o *Invoice) SetAllowances(v []InvoiceAllowance)`

SetAllowances sets Allowances field to given value.

### HasAllowances

`func (o *Invoice) HasAllowances() bool`

HasAllowances returns a boolean if a field has been set.

### GetCharges

`func (o *Invoice) GetCharges() []InvoiceCharge`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *Invoice) GetChargesOk() (*[]InvoiceCharge, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *Invoice) SetCharges(v []InvoiceCharge)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *Invoice) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetLineItems

`func (o *Invoice) GetLineItems() []InvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Invoice) GetLineItemsOk() (*[]InvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Invoice) SetLineItems(v []InvoiceLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *Invoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetSourceType

`func (o *Invoice) GetSourceType() SourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Invoice) GetSourceTypeOk() (*SourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Invoice) SetSourceType(v SourceType)`

SetSourceType sets SourceType field to given value.


### GetFacturXData

`func (o *Invoice) GetFacturXData() FacturXData`

GetFacturXData returns the FacturXData field if non-nil, zero value otherwise.

### GetFacturXDataOk

`func (o *Invoice) GetFacturXDataOk() (*FacturXData, bool)`

GetFacturXDataOk returns a tuple with the FacturXData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacturXData

`func (o *Invoice) SetFacturXData(v FacturXData)`

SetFacturXData sets FacturXData field to given value.

### HasFacturXData

`func (o *Invoice) HasFacturXData() bool`

HasFacturXData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


