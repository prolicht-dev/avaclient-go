# AvaProjectWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | [**ProjectDto**](ProjectDto.md) |  | 
**InvoiceNumber** | Pointer to **string** | The number of the invoice | [optional] 
**Documents** | Pointer to [**[]SupportingDocument**](SupportingDocument.md) | Attached documents to the invoice | [optional] 
**InvoicePeriod** | Pointer to [**InvoicePeriod**](InvoicePeriod.md) |  | [optional] 
**PaymentInstructions** | Pointer to [**PaymentInstructions**](PaymentInstructions.md) |  | [optional] 
**InvoiceType** | [**InvoiceType**](InvoiceType.md) |  | 
**IsEuReverseChargeInvoice** | **bool** | If this is set to true, the invoice will be marked as a reverse charge invoice in all VAT rates. You still need to ensure that the actual tax rates are set to zero in all elements | 
**PaymentDueDate** | Pointer to **time.Time** | Until when the payment is due | [optional] 
**PaymentTerms** | Pointer to **string** | A text description of payment terms, e.g. cash discounts or payment conditions | [optional] 
**PreviousInvoices** | Pointer to [**[]PreviousInvoice**](PreviousInvoice.md) | In case of partial of final invoices, it&#39;s possible to reference the previous invoice numbers | [optional] 
**AlreadyPaidTotalGross** | **float32** | If parts of this invoice (or of previous invoices) have already been paid, then this property can be used to specify that some parts of the invoice are already paid and this will be deducted from the to be paid amount. | 
**References** | Pointer to [**References**](References.md) |  | [optional] 
**Allowances** | Pointer to [**[]InvoiceAllowance**](InvoiceAllowance.md) | Optional document level allowances. For example, those can be used to supply global discounts e.g. for security deposits. When using this property, please pay attention that a deduction factor on the exported ServiceSpecification will also be exported as an allowance, in this case the total price of the invoice might be different from the total price of the ServiceSpecification. | [optional] 
**Charges** | Pointer to [**[]InvoiceCharge**](InvoiceCharge.md) | Optional document level charges, e.g. for shipping or handling. When using this property, please pay attention that the charges given here will be applied to the already calculated total price of the ServiceSpecification, in this case the total price of the invoice might be different from the total price of the ServiceSpecification. | [optional] 
**DeliveryInformation** | Pointer to [**DeliveryInformation**](DeliveryInformation.md) |  | [optional] 
**SellerTaxRepresentative** | Pointer to [**Organization**](Organization.md) |  | [optional] 

## Methods

### NewAvaProjectWrapper

`func NewAvaProjectWrapper(project ProjectDto, invoiceType InvoiceType, isEuReverseChargeInvoice bool, alreadyPaidTotalGross float32, ) *AvaProjectWrapper`

NewAvaProjectWrapper instantiates a new AvaProjectWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvaProjectWrapperWithDefaults

`func NewAvaProjectWrapperWithDefaults() *AvaProjectWrapper`

NewAvaProjectWrapperWithDefaults instantiates a new AvaProjectWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *AvaProjectWrapper) GetProject() ProjectDto`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AvaProjectWrapper) GetProjectOk() (*ProjectDto, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AvaProjectWrapper) SetProject(v ProjectDto)`

SetProject sets Project field to given value.


### GetInvoiceNumber

`func (o *AvaProjectWrapper) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *AvaProjectWrapper) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *AvaProjectWrapper) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *AvaProjectWrapper) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetDocuments

`func (o *AvaProjectWrapper) GetDocuments() []SupportingDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *AvaProjectWrapper) GetDocumentsOk() (*[]SupportingDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *AvaProjectWrapper) SetDocuments(v []SupportingDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *AvaProjectWrapper) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetInvoicePeriod

`func (o *AvaProjectWrapper) GetInvoicePeriod() InvoicePeriod`

GetInvoicePeriod returns the InvoicePeriod field if non-nil, zero value otherwise.

### GetInvoicePeriodOk

`func (o *AvaProjectWrapper) GetInvoicePeriodOk() (*InvoicePeriod, bool)`

GetInvoicePeriodOk returns a tuple with the InvoicePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePeriod

`func (o *AvaProjectWrapper) SetInvoicePeriod(v InvoicePeriod)`

SetInvoicePeriod sets InvoicePeriod field to given value.

### HasInvoicePeriod

`func (o *AvaProjectWrapper) HasInvoicePeriod() bool`

HasInvoicePeriod returns a boolean if a field has been set.

### GetPaymentInstructions

`func (o *AvaProjectWrapper) GetPaymentInstructions() PaymentInstructions`

GetPaymentInstructions returns the PaymentInstructions field if non-nil, zero value otherwise.

### GetPaymentInstructionsOk

`func (o *AvaProjectWrapper) GetPaymentInstructionsOk() (*PaymentInstructions, bool)`

GetPaymentInstructionsOk returns a tuple with the PaymentInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructions

`func (o *AvaProjectWrapper) SetPaymentInstructions(v PaymentInstructions)`

SetPaymentInstructions sets PaymentInstructions field to given value.

### HasPaymentInstructions

`func (o *AvaProjectWrapper) HasPaymentInstructions() bool`

HasPaymentInstructions returns a boolean if a field has been set.

### GetInvoiceType

`func (o *AvaProjectWrapper) GetInvoiceType() InvoiceType`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *AvaProjectWrapper) GetInvoiceTypeOk() (*InvoiceType, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *AvaProjectWrapper) SetInvoiceType(v InvoiceType)`

SetInvoiceType sets InvoiceType field to given value.


### GetIsEuReverseChargeInvoice

`func (o *AvaProjectWrapper) GetIsEuReverseChargeInvoice() bool`

GetIsEuReverseChargeInvoice returns the IsEuReverseChargeInvoice field if non-nil, zero value otherwise.

### GetIsEuReverseChargeInvoiceOk

`func (o *AvaProjectWrapper) GetIsEuReverseChargeInvoiceOk() (*bool, bool)`

GetIsEuReverseChargeInvoiceOk returns a tuple with the IsEuReverseChargeInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEuReverseChargeInvoice

`func (o *AvaProjectWrapper) SetIsEuReverseChargeInvoice(v bool)`

SetIsEuReverseChargeInvoice sets IsEuReverseChargeInvoice field to given value.


### GetPaymentDueDate

`func (o *AvaProjectWrapper) GetPaymentDueDate() time.Time`

GetPaymentDueDate returns the PaymentDueDate field if non-nil, zero value otherwise.

### GetPaymentDueDateOk

`func (o *AvaProjectWrapper) GetPaymentDueDateOk() (*time.Time, bool)`

GetPaymentDueDateOk returns a tuple with the PaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDueDate

`func (o *AvaProjectWrapper) SetPaymentDueDate(v time.Time)`

SetPaymentDueDate sets PaymentDueDate field to given value.

### HasPaymentDueDate

`func (o *AvaProjectWrapper) HasPaymentDueDate() bool`

HasPaymentDueDate returns a boolean if a field has been set.

### GetPaymentTerms

`func (o *AvaProjectWrapper) GetPaymentTerms() string`

GetPaymentTerms returns the PaymentTerms field if non-nil, zero value otherwise.

### GetPaymentTermsOk

`func (o *AvaProjectWrapper) GetPaymentTermsOk() (*string, bool)`

GetPaymentTermsOk returns a tuple with the PaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTerms

`func (o *AvaProjectWrapper) SetPaymentTerms(v string)`

SetPaymentTerms sets PaymentTerms field to given value.

### HasPaymentTerms

`func (o *AvaProjectWrapper) HasPaymentTerms() bool`

HasPaymentTerms returns a boolean if a field has been set.

### GetPreviousInvoices

`func (o *AvaProjectWrapper) GetPreviousInvoices() []PreviousInvoice`

GetPreviousInvoices returns the PreviousInvoices field if non-nil, zero value otherwise.

### GetPreviousInvoicesOk

`func (o *AvaProjectWrapper) GetPreviousInvoicesOk() (*[]PreviousInvoice, bool)`

GetPreviousInvoicesOk returns a tuple with the PreviousInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousInvoices

`func (o *AvaProjectWrapper) SetPreviousInvoices(v []PreviousInvoice)`

SetPreviousInvoices sets PreviousInvoices field to given value.

### HasPreviousInvoices

`func (o *AvaProjectWrapper) HasPreviousInvoices() bool`

HasPreviousInvoices returns a boolean if a field has been set.

### GetAlreadyPaidTotalGross

`func (o *AvaProjectWrapper) GetAlreadyPaidTotalGross() float32`

GetAlreadyPaidTotalGross returns the AlreadyPaidTotalGross field if non-nil, zero value otherwise.

### GetAlreadyPaidTotalGrossOk

`func (o *AvaProjectWrapper) GetAlreadyPaidTotalGrossOk() (*float32, bool)`

GetAlreadyPaidTotalGrossOk returns a tuple with the AlreadyPaidTotalGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyPaidTotalGross

`func (o *AvaProjectWrapper) SetAlreadyPaidTotalGross(v float32)`

SetAlreadyPaidTotalGross sets AlreadyPaidTotalGross field to given value.


### GetReferences

`func (o *AvaProjectWrapper) GetReferences() References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AvaProjectWrapper) GetReferencesOk() (*References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AvaProjectWrapper) SetReferences(v References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AvaProjectWrapper) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetAllowances

`func (o *AvaProjectWrapper) GetAllowances() []InvoiceAllowance`

GetAllowances returns the Allowances field if non-nil, zero value otherwise.

### GetAllowancesOk

`func (o *AvaProjectWrapper) GetAllowancesOk() (*[]InvoiceAllowance, bool)`

GetAllowancesOk returns a tuple with the Allowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowances

`func (o *AvaProjectWrapper) SetAllowances(v []InvoiceAllowance)`

SetAllowances sets Allowances field to given value.

### HasAllowances

`func (o *AvaProjectWrapper) HasAllowances() bool`

HasAllowances returns a boolean if a field has been set.

### GetCharges

`func (o *AvaProjectWrapper) GetCharges() []InvoiceCharge`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *AvaProjectWrapper) GetChargesOk() (*[]InvoiceCharge, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *AvaProjectWrapper) SetCharges(v []InvoiceCharge)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *AvaProjectWrapper) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetDeliveryInformation

`func (o *AvaProjectWrapper) GetDeliveryInformation() DeliveryInformation`

GetDeliveryInformation returns the DeliveryInformation field if non-nil, zero value otherwise.

### GetDeliveryInformationOk

`func (o *AvaProjectWrapper) GetDeliveryInformationOk() (*DeliveryInformation, bool)`

GetDeliveryInformationOk returns a tuple with the DeliveryInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryInformation

`func (o *AvaProjectWrapper) SetDeliveryInformation(v DeliveryInformation)`

SetDeliveryInformation sets DeliveryInformation field to given value.

### HasDeliveryInformation

`func (o *AvaProjectWrapper) HasDeliveryInformation() bool`

HasDeliveryInformation returns a boolean if a field has been set.

### GetSellerTaxRepresentative

`func (o *AvaProjectWrapper) GetSellerTaxRepresentative() Organization`

GetSellerTaxRepresentative returns the SellerTaxRepresentative field if non-nil, zero value otherwise.

### GetSellerTaxRepresentativeOk

`func (o *AvaProjectWrapper) GetSellerTaxRepresentativeOk() (*Organization, bool)`

GetSellerTaxRepresentativeOk returns a tuple with the SellerTaxRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerTaxRepresentative

`func (o *AvaProjectWrapper) SetSellerTaxRepresentative(v Organization)`

SetSellerTaxRepresentative sets SellerTaxRepresentative field to given value.

### HasSellerTaxRepresentative

`func (o *AvaProjectWrapper) HasSellerTaxRepresentative() bool`

HasSellerTaxRepresentative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


