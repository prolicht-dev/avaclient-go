# FacturXData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuidelineId** | Pointer to **string** | The guideline id from the source CII file, if present. | [optional] 
**Profile** | [**FacturXProfile**](FacturXProfile.md) |  | 
**RoutingId** | Pointer to **string** | Factur-X specific routing ID, \&quot;Leitweg-ID\&quot; in German. This is typically required in public administration invoicing in Germany. | [optional] 
**FacturXPaymentTerms** | Pointer to [**[]FacturXPaymentTerm**](FacturXPaymentTerm.md) | The field BT-20 in XRechnung can only hold one payment term, but Factur-X / ZUGFeRD allows multiple payment terms. If the invoice was originally created in Factur-X / ZUGFeRD format, then this property holds all additional payment terms except the first one. Similar, on export, this property is used to export additional payment terms to Factur-X / ZUGFeRD. Factur-X / ZUGFeRD also allows a more detailled representation of payment terms, so they are given in a special object instead of just a string. If those are present, the regular PaymentTerms will be ignored on export to the CII / Factur-X format, as the information in this object is more detailled. | [optional] 

## Methods

### NewFacturXData

`func NewFacturXData(profile FacturXProfile, ) *FacturXData`

NewFacturXData instantiates a new FacturXData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacturXDataWithDefaults

`func NewFacturXDataWithDefaults() *FacturXData`

NewFacturXDataWithDefaults instantiates a new FacturXData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuidelineId

`func (o *FacturXData) GetGuidelineId() string`

GetGuidelineId returns the GuidelineId field if non-nil, zero value otherwise.

### GetGuidelineIdOk

`func (o *FacturXData) GetGuidelineIdOk() (*string, bool)`

GetGuidelineIdOk returns a tuple with the GuidelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidelineId

`func (o *FacturXData) SetGuidelineId(v string)`

SetGuidelineId sets GuidelineId field to given value.

### HasGuidelineId

`func (o *FacturXData) HasGuidelineId() bool`

HasGuidelineId returns a boolean if a field has been set.

### GetProfile

`func (o *FacturXData) GetProfile() FacturXProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FacturXData) GetProfileOk() (*FacturXProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FacturXData) SetProfile(v FacturXProfile)`

SetProfile sets Profile field to given value.


### GetRoutingId

`func (o *FacturXData) GetRoutingId() string`

GetRoutingId returns the RoutingId field if non-nil, zero value otherwise.

### GetRoutingIdOk

`func (o *FacturXData) GetRoutingIdOk() (*string, bool)`

GetRoutingIdOk returns a tuple with the RoutingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingId

`func (o *FacturXData) SetRoutingId(v string)`

SetRoutingId sets RoutingId field to given value.

### HasRoutingId

`func (o *FacturXData) HasRoutingId() bool`

HasRoutingId returns a boolean if a field has been set.

### GetFacturXPaymentTerms

`func (o *FacturXData) GetFacturXPaymentTerms() []FacturXPaymentTerm`

GetFacturXPaymentTerms returns the FacturXPaymentTerms field if non-nil, zero value otherwise.

### GetFacturXPaymentTermsOk

`func (o *FacturXData) GetFacturXPaymentTermsOk() (*[]FacturXPaymentTerm, bool)`

GetFacturXPaymentTermsOk returns a tuple with the FacturXPaymentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacturXPaymentTerms

`func (o *FacturXData) SetFacturXPaymentTerms(v []FacturXPaymentTerm)`

SetFacturXPaymentTerms sets FacturXPaymentTerms field to given value.

### HasFacturXPaymentTerms

`func (o *FacturXData) HasFacturXPaymentTerms() bool`

HasFacturXPaymentTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


