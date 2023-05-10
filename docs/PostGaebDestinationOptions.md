# PostGaebDestinationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationGaebType** | [**DestinationGaebType**](DestinationGaebType.md) |  | 
**TargetExchangePhaseTransform** | [**DestinationGaebExchangePhase**](DestinationGaebExchangePhase.md) |  | 
**EnforceStrictOfferPhaseLongTextOutput** | **bool** | Defaults to false. If this is enabled, exported long texts to GAEB XML that use text additions will be strictly schema compliant. If this is not enabled, any text that is marked to contain a text addition is exported in full to ensure that incorrectly used text additions are still preserved in the export. | 
**ExportQuantityDetermination** | **bool** | Defaults to false. If this is enabled, quantities are exported in detail in GAEB XML targets via the &#39;QtyDeterm&#39; (Quantity Determination, or Quantity Take Off) fields. To control this, you can set custom quantity calculations in the &#39;QuantityComponents&#39; property of positions. Please see the entry for &#39;Quantity Determination&#39; in the Dangl.AVA HowTo documentation section. Please be advised that enabling this might export data that was not intended to be exported, like internal quantity calculation details, depending on what data you put in the &#39;QuantityComponents&#39; property. | 
**RemoveUnprintableCharactersFromTexts** | **bool** | If this is enabled, unprintable characters are removed from text elements. Otherwise, the conversion might fail in case some text content contains characters that are not allowed in XML output formats. Defaults to true. | 
**ForceIncludeDescriptions** | **bool** | If this is enabled, all description elements like texts and execution descriptions will be output to the result. This is mostly only applicable to GAEB exports to phase 84 - Offer, which does typically not include descriptions. | 
**TreatNullItemNumberSchemaAsInvalid** | **bool** | When exporting to GAEB, an item number schema is usually required. AVACloud will try to fix invalid item number schemas. With this setting, you can also tell AVACloud to treat a null value as invalid. Otherwise, null schemas will simply be ignored and not lead to any schema being generated. It is recommended to enable this option, but it is disabled by default for compatibility reasons. | 

## Methods

### NewPostGaebDestinationOptions

`func NewPostGaebDestinationOptions(destinationGaebType DestinationGaebType, targetExchangePhaseTransform DestinationGaebExchangePhase, enforceStrictOfferPhaseLongTextOutput bool, exportQuantityDetermination bool, removeUnprintableCharactersFromTexts bool, forceIncludeDescriptions bool, treatNullItemNumberSchemaAsInvalid bool, ) *PostGaebDestinationOptions`

NewPostGaebDestinationOptions instantiates a new PostGaebDestinationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGaebDestinationOptionsWithDefaults

`func NewPostGaebDestinationOptionsWithDefaults() *PostGaebDestinationOptions`

NewPostGaebDestinationOptionsWithDefaults instantiates a new PostGaebDestinationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationGaebType

`func (o *PostGaebDestinationOptions) GetDestinationGaebType() DestinationGaebType`

GetDestinationGaebType returns the DestinationGaebType field if non-nil, zero value otherwise.

### GetDestinationGaebTypeOk

`func (o *PostGaebDestinationOptions) GetDestinationGaebTypeOk() (*DestinationGaebType, bool)`

GetDestinationGaebTypeOk returns a tuple with the DestinationGaebType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGaebType

`func (o *PostGaebDestinationOptions) SetDestinationGaebType(v DestinationGaebType)`

SetDestinationGaebType sets DestinationGaebType field to given value.


### GetTargetExchangePhaseTransform

`func (o *PostGaebDestinationOptions) GetTargetExchangePhaseTransform() DestinationGaebExchangePhase`

GetTargetExchangePhaseTransform returns the TargetExchangePhaseTransform field if non-nil, zero value otherwise.

### GetTargetExchangePhaseTransformOk

`func (o *PostGaebDestinationOptions) GetTargetExchangePhaseTransformOk() (*DestinationGaebExchangePhase, bool)`

GetTargetExchangePhaseTransformOk returns a tuple with the TargetExchangePhaseTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetExchangePhaseTransform

`func (o *PostGaebDestinationOptions) SetTargetExchangePhaseTransform(v DestinationGaebExchangePhase)`

SetTargetExchangePhaseTransform sets TargetExchangePhaseTransform field to given value.


### GetEnforceStrictOfferPhaseLongTextOutput

`func (o *PostGaebDestinationOptions) GetEnforceStrictOfferPhaseLongTextOutput() bool`

GetEnforceStrictOfferPhaseLongTextOutput returns the EnforceStrictOfferPhaseLongTextOutput field if non-nil, zero value otherwise.

### GetEnforceStrictOfferPhaseLongTextOutputOk

`func (o *PostGaebDestinationOptions) GetEnforceStrictOfferPhaseLongTextOutputOk() (*bool, bool)`

GetEnforceStrictOfferPhaseLongTextOutputOk returns a tuple with the EnforceStrictOfferPhaseLongTextOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceStrictOfferPhaseLongTextOutput

`func (o *PostGaebDestinationOptions) SetEnforceStrictOfferPhaseLongTextOutput(v bool)`

SetEnforceStrictOfferPhaseLongTextOutput sets EnforceStrictOfferPhaseLongTextOutput field to given value.


### GetExportQuantityDetermination

`func (o *PostGaebDestinationOptions) GetExportQuantityDetermination() bool`

GetExportQuantityDetermination returns the ExportQuantityDetermination field if non-nil, zero value otherwise.

### GetExportQuantityDeterminationOk

`func (o *PostGaebDestinationOptions) GetExportQuantityDeterminationOk() (*bool, bool)`

GetExportQuantityDeterminationOk returns a tuple with the ExportQuantityDetermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportQuantityDetermination

`func (o *PostGaebDestinationOptions) SetExportQuantityDetermination(v bool)`

SetExportQuantityDetermination sets ExportQuantityDetermination field to given value.


### GetRemoveUnprintableCharactersFromTexts

`func (o *PostGaebDestinationOptions) GetRemoveUnprintableCharactersFromTexts() bool`

GetRemoveUnprintableCharactersFromTexts returns the RemoveUnprintableCharactersFromTexts field if non-nil, zero value otherwise.

### GetRemoveUnprintableCharactersFromTextsOk

`func (o *PostGaebDestinationOptions) GetRemoveUnprintableCharactersFromTextsOk() (*bool, bool)`

GetRemoveUnprintableCharactersFromTextsOk returns a tuple with the RemoveUnprintableCharactersFromTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveUnprintableCharactersFromTexts

`func (o *PostGaebDestinationOptions) SetRemoveUnprintableCharactersFromTexts(v bool)`

SetRemoveUnprintableCharactersFromTexts sets RemoveUnprintableCharactersFromTexts field to given value.


### GetForceIncludeDescriptions

`func (o *PostGaebDestinationOptions) GetForceIncludeDescriptions() bool`

GetForceIncludeDescriptions returns the ForceIncludeDescriptions field if non-nil, zero value otherwise.

### GetForceIncludeDescriptionsOk

`func (o *PostGaebDestinationOptions) GetForceIncludeDescriptionsOk() (*bool, bool)`

GetForceIncludeDescriptionsOk returns a tuple with the ForceIncludeDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceIncludeDescriptions

`func (o *PostGaebDestinationOptions) SetForceIncludeDescriptions(v bool)`

SetForceIncludeDescriptions sets ForceIncludeDescriptions field to given value.


### GetTreatNullItemNumberSchemaAsInvalid

`func (o *PostGaebDestinationOptions) GetTreatNullItemNumberSchemaAsInvalid() bool`

GetTreatNullItemNumberSchemaAsInvalid returns the TreatNullItemNumberSchemaAsInvalid field if non-nil, zero value otherwise.

### GetTreatNullItemNumberSchemaAsInvalidOk

`func (o *PostGaebDestinationOptions) GetTreatNullItemNumberSchemaAsInvalidOk() (*bool, bool)`

GetTreatNullItemNumberSchemaAsInvalidOk returns a tuple with the TreatNullItemNumberSchemaAsInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatNullItemNumberSchemaAsInvalid

`func (o *PostGaebDestinationOptions) SetTreatNullItemNumberSchemaAsInvalid(v bool)`

SetTreatNullItemNumberSchemaAsInvalid sets TreatNullItemNumberSchemaAsInvalid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


