# PostXRechnungPdfEmbedd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | [**Invoice**](Invoice.md) |  | 
**PdfFileBase64** | **string** | The PDF file as Base64 string | 

## Methods

### NewPostXRechnungPdfEmbedd

`func NewPostXRechnungPdfEmbedd(invoice Invoice, pdfFileBase64 string, ) *PostXRechnungPdfEmbedd`

NewPostXRechnungPdfEmbedd instantiates a new PostXRechnungPdfEmbedd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostXRechnungPdfEmbeddWithDefaults

`func NewPostXRechnungPdfEmbeddWithDefaults() *PostXRechnungPdfEmbedd`

NewPostXRechnungPdfEmbeddWithDefaults instantiates a new PostXRechnungPdfEmbedd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *PostXRechnungPdfEmbedd) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *PostXRechnungPdfEmbedd) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *PostXRechnungPdfEmbedd) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.


### GetPdfFileBase64

`func (o *PostXRechnungPdfEmbedd) GetPdfFileBase64() string`

GetPdfFileBase64 returns the PdfFileBase64 field if non-nil, zero value otherwise.

### GetPdfFileBase64Ok

`func (o *PostXRechnungPdfEmbedd) GetPdfFileBase64Ok() (*string, bool)`

GetPdfFileBase64Ok returns a tuple with the PdfFileBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfFileBase64

`func (o *PostXRechnungPdfEmbedd) SetPdfFileBase64(v string)`

SetPdfFileBase64 sets PdfFileBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


