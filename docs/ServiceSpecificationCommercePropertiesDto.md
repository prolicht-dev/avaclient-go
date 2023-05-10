# ServiceSpecificationCommercePropertiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixedPriceDate** | Pointer to **time.Time** | The date until the price is valid or fixed. | [optional] 
**DeliveryVoucherDate** | Pointer to **time.Time** | The date of the delivery voucher note. | [optional] 
**DeliveryDate** | Pointer to **time.Time** | The actual date of delivery. | [optional] 
**InquiryNumber** | Pointer to **string** | The number of the inquiry, usually in a context of offer requests. | [optional] 
**OfferNumber** | Pointer to **string** | The number of the offer, usually in a context of an offer. | [optional] 
**OrderNumber** | Pointer to **string** | The order number, usually in the context of a grant or contract. | [optional] 
**OrderConfirmationNumber** | Pointer to **string** | The order confirmation number | [optional] 
**DeliveryNumber** | Pointer to **string** | The number of the delivery, e.g. on the delivery note voucher. | [optional] 
**CustomerReferenceNumber** | Pointer to **string** | The reference number of the customer / buyer. | [optional] 
**SupplierReferenceNumber** | Pointer to **string** | The reference number of the supplier / bidder. | [optional] 
**ShippingType** | Pointer to **string** | The type of shippment. | [optional] 
**InquiryType** | [**CommerceInquiryTypeDto**](CommerceInquiryTypeDto.md) |  | 

## Methods

### NewServiceSpecificationCommercePropertiesDto

`func NewServiceSpecificationCommercePropertiesDto(inquiryType CommerceInquiryTypeDto, ) *ServiceSpecificationCommercePropertiesDto`

NewServiceSpecificationCommercePropertiesDto instantiates a new ServiceSpecificationCommercePropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSpecificationCommercePropertiesDtoWithDefaults

`func NewServiceSpecificationCommercePropertiesDtoWithDefaults() *ServiceSpecificationCommercePropertiesDto`

NewServiceSpecificationCommercePropertiesDtoWithDefaults instantiates a new ServiceSpecificationCommercePropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixedPriceDate

`func (o *ServiceSpecificationCommercePropertiesDto) GetFixedPriceDate() time.Time`

GetFixedPriceDate returns the FixedPriceDate field if non-nil, zero value otherwise.

### GetFixedPriceDateOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetFixedPriceDateOk() (*time.Time, bool)`

GetFixedPriceDateOk returns a tuple with the FixedPriceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedPriceDate

`func (o *ServiceSpecificationCommercePropertiesDto) SetFixedPriceDate(v time.Time)`

SetFixedPriceDate sets FixedPriceDate field to given value.

### HasFixedPriceDate

`func (o *ServiceSpecificationCommercePropertiesDto) HasFixedPriceDate() bool`

HasFixedPriceDate returns a boolean if a field has been set.

### GetDeliveryVoucherDate

`func (o *ServiceSpecificationCommercePropertiesDto) GetDeliveryVoucherDate() time.Time`

GetDeliveryVoucherDate returns the DeliveryVoucherDate field if non-nil, zero value otherwise.

### GetDeliveryVoucherDateOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetDeliveryVoucherDateOk() (*time.Time, bool)`

GetDeliveryVoucherDateOk returns a tuple with the DeliveryVoucherDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryVoucherDate

`func (o *ServiceSpecificationCommercePropertiesDto) SetDeliveryVoucherDate(v time.Time)`

SetDeliveryVoucherDate sets DeliveryVoucherDate field to given value.

### HasDeliveryVoucherDate

`func (o *ServiceSpecificationCommercePropertiesDto) HasDeliveryVoucherDate() bool`

HasDeliveryVoucherDate returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *ServiceSpecificationCommercePropertiesDto) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ServiceSpecificationCommercePropertiesDto) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ServiceSpecificationCommercePropertiesDto) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetInquiryNumber

`func (o *ServiceSpecificationCommercePropertiesDto) GetInquiryNumber() string`

GetInquiryNumber returns the InquiryNumber field if non-nil, zero value otherwise.

### GetInquiryNumberOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetInquiryNumberOk() (*string, bool)`

GetInquiryNumberOk returns a tuple with the InquiryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquiryNumber

`func (o *ServiceSpecificationCommercePropertiesDto) SetInquiryNumber(v string)`

SetInquiryNumber sets InquiryNumber field to given value.

### HasInquiryNumber

`func (o *ServiceSpecificationCommercePropertiesDto) HasInquiryNumber() bool`

HasInquiryNumber returns a boolean if a field has been set.

### GetOfferNumber

`func (o *ServiceSpecificationCommercePropertiesDto) GetOfferNumber() string`

GetOfferNumber returns the OfferNumber field if non-nil, zero value otherwise.

### GetOfferNumberOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetOfferNumberOk() (*string, bool)`

GetOfferNumberOk returns a tuple with the OfferNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferNumber

`func (o *ServiceSpecificationCommercePropertiesDto) SetOfferNumber(v string)`

SetOfferNumber sets OfferNumber field to given value.

### HasOfferNumber

`func (o *ServiceSpecificationCommercePropertiesDto) HasOfferNumber() bool`

HasOfferNumber returns a boolean if a field has been set.

### GetOrderNumber

`func (o *ServiceSpecificationCommercePropertiesDto) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *ServiceSpecificationCommercePropertiesDto) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *ServiceSpecificationCommercePropertiesDto) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetOrderConfirmationNumber

`func (o *ServiceSpecificationCommercePropertiesDto) GetOrderConfirmationNumber() string`

GetOrderConfirmationNumber returns the OrderConfirmationNumber field if non-nil, zero value otherwise.

### GetOrderConfirmationNumberOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetOrderConfirmationNumberOk() (*string, bool)`

GetOrderConfirmationNumberOk returns a tuple with the OrderConfirmationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderConfirmationNumber

`func (o *ServiceSpecificationCommercePropertiesDto) SetOrderConfirmationNumber(v string)`

SetOrderConfirmationNumber sets OrderConfirmationNumber field to given value.

### HasOrderConfirmationNumber

`func (o *ServiceSpecificationCommercePropertiesDto) HasOrderConfirmationNumber() bool`

HasOrderConfirmationNumber returns a boolean if a field has been set.

### GetDeliveryNumber

`func (o *ServiceSpecificationCommercePropertiesDto) GetDeliveryNumber() string`

GetDeliveryNumber returns the DeliveryNumber field if non-nil, zero value otherwise.

### GetDeliveryNumberOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetDeliveryNumberOk() (*string, bool)`

GetDeliveryNumberOk returns a tuple with the DeliveryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryNumber

`func (o *ServiceSpecificationCommercePropertiesDto) SetDeliveryNumber(v string)`

SetDeliveryNumber sets DeliveryNumber field to given value.

### HasDeliveryNumber

`func (o *ServiceSpecificationCommercePropertiesDto) HasDeliveryNumber() bool`

HasDeliveryNumber returns a boolean if a field has been set.

### GetCustomerReferenceNumber

`func (o *ServiceSpecificationCommercePropertiesDto) GetCustomerReferenceNumber() string`

GetCustomerReferenceNumber returns the CustomerReferenceNumber field if non-nil, zero value otherwise.

### GetCustomerReferenceNumberOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetCustomerReferenceNumberOk() (*string, bool)`

GetCustomerReferenceNumberOk returns a tuple with the CustomerReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReferenceNumber

`func (o *ServiceSpecificationCommercePropertiesDto) SetCustomerReferenceNumber(v string)`

SetCustomerReferenceNumber sets CustomerReferenceNumber field to given value.

### HasCustomerReferenceNumber

`func (o *ServiceSpecificationCommercePropertiesDto) HasCustomerReferenceNumber() bool`

HasCustomerReferenceNumber returns a boolean if a field has been set.

### GetSupplierReferenceNumber

`func (o *ServiceSpecificationCommercePropertiesDto) GetSupplierReferenceNumber() string`

GetSupplierReferenceNumber returns the SupplierReferenceNumber field if non-nil, zero value otherwise.

### GetSupplierReferenceNumberOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetSupplierReferenceNumberOk() (*string, bool)`

GetSupplierReferenceNumberOk returns a tuple with the SupplierReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierReferenceNumber

`func (o *ServiceSpecificationCommercePropertiesDto) SetSupplierReferenceNumber(v string)`

SetSupplierReferenceNumber sets SupplierReferenceNumber field to given value.

### HasSupplierReferenceNumber

`func (o *ServiceSpecificationCommercePropertiesDto) HasSupplierReferenceNumber() bool`

HasSupplierReferenceNumber returns a boolean if a field has been set.

### GetShippingType

`func (o *ServiceSpecificationCommercePropertiesDto) GetShippingType() string`

GetShippingType returns the ShippingType field if non-nil, zero value otherwise.

### GetShippingTypeOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetShippingTypeOk() (*string, bool)`

GetShippingTypeOk returns a tuple with the ShippingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingType

`func (o *ServiceSpecificationCommercePropertiesDto) SetShippingType(v string)`

SetShippingType sets ShippingType field to given value.

### HasShippingType

`func (o *ServiceSpecificationCommercePropertiesDto) HasShippingType() bool`

HasShippingType returns a boolean if a field has been set.

### GetInquiryType

`func (o *ServiceSpecificationCommercePropertiesDto) GetInquiryType() CommerceInquiryTypeDto`

GetInquiryType returns the InquiryType field if non-nil, zero value otherwise.

### GetInquiryTypeOk

`func (o *ServiceSpecificationCommercePropertiesDto) GetInquiryTypeOk() (*CommerceInquiryTypeDto, bool)`

GetInquiryTypeOk returns a tuple with the InquiryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInquiryType

`func (o *ServiceSpecificationCommercePropertiesDto) SetInquiryType(v CommerceInquiryTypeDto)`

SetInquiryType sets InquiryType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


