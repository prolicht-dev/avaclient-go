# PostOenormDestinationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationOenormType** | [**DestinationOenormType**](DestinationOenormType.md) |  | 
**TryRepairProjectStructure** | **bool** | Defaults to false. If this is enabled, the converter will try to ensure that the project structure can be mapped to Oenorm. It might introduce additional group levels to ensure a compatible target | 
**SkipTryEnforceSchemaCompliantXmlOutput** | **bool** | If this option is enabled, AVACloud will not attempt to force a schema-compliant Xml output for ÖNorm targets that are Xml based. By default, AVACloud will try to add required fields, even if no data is present, with sensible defaults. This behavior can be disabled with this option. | 
**RemoveUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 

## Methods

### NewPostOenormDestinationOptions

`func NewPostOenormDestinationOptions(destinationOenormType DestinationOenormType, tryRepairProjectStructure bool, skipTryEnforceSchemaCompliantXmlOutput bool, removeUnprintableCharactersFromTexts bool, ) *PostOenormDestinationOptions`

NewPostOenormDestinationOptions instantiates a new PostOenormDestinationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostOenormDestinationOptionsWithDefaults

`func NewPostOenormDestinationOptionsWithDefaults() *PostOenormDestinationOptions`

NewPostOenormDestinationOptionsWithDefaults instantiates a new PostOenormDestinationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationOenormType

`func (o *PostOenormDestinationOptions) GetDestinationOenormType() DestinationOenormType`

GetDestinationOenormType returns the DestinationOenormType field if non-nil, zero value otherwise.

### GetDestinationOenormTypeOk

`func (o *PostOenormDestinationOptions) GetDestinationOenormTypeOk() (*DestinationOenormType, bool)`

GetDestinationOenormTypeOk returns a tuple with the DestinationOenormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationOenormType

`func (o *PostOenormDestinationOptions) SetDestinationOenormType(v DestinationOenormType)`

SetDestinationOenormType sets DestinationOenormType field to given value.


### GetTryRepairProjectStructure

`func (o *PostOenormDestinationOptions) GetTryRepairProjectStructure() bool`

GetTryRepairProjectStructure returns the TryRepairProjectStructure field if non-nil, zero value otherwise.

### GetTryRepairProjectStructureOk

`func (o *PostOenormDestinationOptions) GetTryRepairProjectStructureOk() (*bool, bool)`

GetTryRepairProjectStructureOk returns a tuple with the TryRepairProjectStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryRepairProjectStructure

`func (o *PostOenormDestinationOptions) SetTryRepairProjectStructure(v bool)`

SetTryRepairProjectStructure sets TryRepairProjectStructure field to given value.


### GetSkipTryEnforceSchemaCompliantXmlOutput

`func (o *PostOenormDestinationOptions) GetSkipTryEnforceSchemaCompliantXmlOutput() bool`

GetSkipTryEnforceSchemaCompliantXmlOutput returns the SkipTryEnforceSchemaCompliantXmlOutput field if non-nil, zero value otherwise.

### GetSkipTryEnforceSchemaCompliantXmlOutputOk

`func (o *PostOenormDestinationOptions) GetSkipTryEnforceSchemaCompliantXmlOutputOk() (*bool, bool)`

GetSkipTryEnforceSchemaCompliantXmlOutputOk returns a tuple with the SkipTryEnforceSchemaCompliantXmlOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipTryEnforceSchemaCompliantXmlOutput

`func (o *PostOenormDestinationOptions) SetSkipTryEnforceSchemaCompliantXmlOutput(v bool)`

SetSkipTryEnforceSchemaCompliantXmlOutput sets SkipTryEnforceSchemaCompliantXmlOutput field to given value.


### GetRemoveUnprintableCharactersFromTexts

`func (o *PostOenormDestinationOptions) GetRemoveUnprintableCharactersFromTexts() bool`

GetRemoveUnprintableCharactersFromTexts returns the RemoveUnprintableCharactersFromTexts field if non-nil, zero value otherwise.

### GetRemoveUnprintableCharactersFromTextsOk

`func (o *PostOenormDestinationOptions) GetRemoveUnprintableCharactersFromTextsOk() (*bool, bool)`

GetRemoveUnprintableCharactersFromTextsOk returns a tuple with the RemoveUnprintableCharactersFromTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveUnprintableCharactersFromTexts

`func (o *PostOenormDestinationOptions) SetRemoveUnprintableCharactersFromTexts(v bool)`

SetRemoveUnprintableCharactersFromTexts sets RemoveUnprintableCharactersFromTexts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


