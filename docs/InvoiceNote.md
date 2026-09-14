# InvoiceNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortText** | Pointer to **string** | BT-21 Short summary / heading. Should be a text code from this list: https://unece.org/fileadmin/DAM/trade/untdid/d16b/tred/tred4451.htm Field BT-21 | [optional] 
**ShortTextType** | Pointer to [**ShortTextType**](ShortTextType.md) |  | [optional] 
**LongText** | Pointer to **string** | BT-22 Explanation of this note Field BT-22 | [optional] 

## Methods

### NewInvoiceNote

`func NewInvoiceNote() *InvoiceNote`

NewInvoiceNote instantiates a new InvoiceNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceNoteWithDefaults

`func NewInvoiceNoteWithDefaults() *InvoiceNote`

NewInvoiceNoteWithDefaults instantiates a new InvoiceNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortText

`func (o *InvoiceNote) GetShortText() string`

GetShortText returns the ShortText field if non-nil, zero value otherwise.

### GetShortTextOk

`func (o *InvoiceNote) GetShortTextOk() (*string, bool)`

GetShortTextOk returns a tuple with the ShortText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortText

`func (o *InvoiceNote) SetShortText(v string)`

SetShortText sets ShortText field to given value.

### HasShortText

`func (o *InvoiceNote) HasShortText() bool`

HasShortText returns a boolean if a field has been set.

### GetShortTextType

`func (o *InvoiceNote) GetShortTextType() ShortTextType`

GetShortTextType returns the ShortTextType field if non-nil, zero value otherwise.

### GetShortTextTypeOk

`func (o *InvoiceNote) GetShortTextTypeOk() (*ShortTextType, bool)`

GetShortTextTypeOk returns a tuple with the ShortTextType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortTextType

`func (o *InvoiceNote) SetShortTextType(v ShortTextType)`

SetShortTextType sets ShortTextType field to given value.

### HasShortTextType

`func (o *InvoiceNote) HasShortTextType() bool`

HasShortTextType returns a boolean if a field has been set.

### GetLongText

`func (o *InvoiceNote) GetLongText() string`

GetLongText returns the LongText field if non-nil, zero value otherwise.

### GetLongTextOk

`func (o *InvoiceNote) GetLongTextOk() (*string, bool)`

GetLongTextOk returns a tuple with the LongText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongText

`func (o *InvoiceNote) SetLongText(v string)`

SetLongText sets LongText field to given value.

### HasLongText

`func (o *InvoiceNote) HasLongText() bool`

HasLongText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


