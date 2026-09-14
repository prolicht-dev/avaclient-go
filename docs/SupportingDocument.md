# SupportingDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentReference** | Pointer to **string** | BT-122 This is typically an id or a name | [optional] 
**Description** | Pointer to **string** | BT-123 A description of the document | [optional] 
**ExternalDocumentUrl** | Pointer to **string** | BT-124 A link to the document in case this is an external reference | [optional] 
**DocumentBase64** | Pointer to **string** | BT-125 The document in base64 encoded form as binary data. This is ignored if you set PlaceholderId. | [optional] 
**FileName** | Pointer to **string** | BT-125 This should be present if DocumentBase64 is set. | [optional] 
**DocumentMimeType** | [**AttachmentMimeType**](AttachmentMimeType.md) |  | 
**PlaceholderId** | Pointer to **string** | This is placeholder id to use instead of DocumentBase64 property for lazy-loading of documents. The use case of this is to avoid having the documents be processed in memory when they might be handled by the calling code. For example, in the AVACloud webservice, this placeholder can be used to avoid having to send the Base64 encoded document to the server and then downloading it again in the result. If this is given, the generated XML file will not be valid, since it will include the literal Guid value as a placeholder. You can use the &#39;xmlExportCallback&#39; parameter in the XRechnungWriter to replace the placeholder with the actual Base64 contents of the file in this library, or just use a string replace method if you&#39;re working with the XML in a different context. | [optional] 

## Methods

### NewSupportingDocument

`func NewSupportingDocument(documentMimeType AttachmentMimeType, ) *SupportingDocument`

NewSupportingDocument instantiates a new SupportingDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportingDocumentWithDefaults

`func NewSupportingDocumentWithDefaults() *SupportingDocument`

NewSupportingDocumentWithDefaults instantiates a new SupportingDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentReference

`func (o *SupportingDocument) GetDocumentReference() string`

GetDocumentReference returns the DocumentReference field if non-nil, zero value otherwise.

### GetDocumentReferenceOk

`func (o *SupportingDocument) GetDocumentReferenceOk() (*string, bool)`

GetDocumentReferenceOk returns a tuple with the DocumentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReference

`func (o *SupportingDocument) SetDocumentReference(v string)`

SetDocumentReference sets DocumentReference field to given value.

### HasDocumentReference

`func (o *SupportingDocument) HasDocumentReference() bool`

HasDocumentReference returns a boolean if a field has been set.

### GetDescription

`func (o *SupportingDocument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SupportingDocument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SupportingDocument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SupportingDocument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalDocumentUrl

`func (o *SupportingDocument) GetExternalDocumentUrl() string`

GetExternalDocumentUrl returns the ExternalDocumentUrl field if non-nil, zero value otherwise.

### GetExternalDocumentUrlOk

`func (o *SupportingDocument) GetExternalDocumentUrlOk() (*string, bool)`

GetExternalDocumentUrlOk returns a tuple with the ExternalDocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentUrl

`func (o *SupportingDocument) SetExternalDocumentUrl(v string)`

SetExternalDocumentUrl sets ExternalDocumentUrl field to given value.

### HasExternalDocumentUrl

`func (o *SupportingDocument) HasExternalDocumentUrl() bool`

HasExternalDocumentUrl returns a boolean if a field has been set.

### GetDocumentBase64

`func (o *SupportingDocument) GetDocumentBase64() string`

GetDocumentBase64 returns the DocumentBase64 field if non-nil, zero value otherwise.

### GetDocumentBase64Ok

`func (o *SupportingDocument) GetDocumentBase64Ok() (*string, bool)`

GetDocumentBase64Ok returns a tuple with the DocumentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentBase64

`func (o *SupportingDocument) SetDocumentBase64(v string)`

SetDocumentBase64 sets DocumentBase64 field to given value.

### HasDocumentBase64

`func (o *SupportingDocument) HasDocumentBase64() bool`

HasDocumentBase64 returns a boolean if a field has been set.

### GetFileName

`func (o *SupportingDocument) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SupportingDocument) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SupportingDocument) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *SupportingDocument) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetDocumentMimeType

`func (o *SupportingDocument) GetDocumentMimeType() AttachmentMimeType`

GetDocumentMimeType returns the DocumentMimeType field if non-nil, zero value otherwise.

### GetDocumentMimeTypeOk

`func (o *SupportingDocument) GetDocumentMimeTypeOk() (*AttachmentMimeType, bool)`

GetDocumentMimeTypeOk returns a tuple with the DocumentMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMimeType

`func (o *SupportingDocument) SetDocumentMimeType(v AttachmentMimeType)`

SetDocumentMimeType sets DocumentMimeType field to given value.


### GetPlaceholderId

`func (o *SupportingDocument) GetPlaceholderId() string`

GetPlaceholderId returns the PlaceholderId field if non-nil, zero value otherwise.

### GetPlaceholderIdOk

`func (o *SupportingDocument) GetPlaceholderIdOk() (*string, bool)`

GetPlaceholderIdOk returns a tuple with the PlaceholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholderId

`func (o *SupportingDocument) SetPlaceholderId(v string)`

SetPlaceholderId sets PlaceholderId field to given value.

### HasPlaceholderId

`func (o *SupportingDocument) HasPlaceholderId() bool`

HasPlaceholderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


