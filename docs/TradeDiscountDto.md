# TradeDiscountDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Elements GUID identifier. | 
**Deadline** | **int32** | The amount of days for how long a declared trade discount applies. Must be a positive number, negative values will be ignored and not set. | 
**Rate** | **float32** | The rate of the trade discount. | 

## Methods

### NewTradeDiscountDto

`func NewTradeDiscountDto(id string, deadline int32, rate float32, ) *TradeDiscountDto`

NewTradeDiscountDto instantiates a new TradeDiscountDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeDiscountDtoWithDefaults

`func NewTradeDiscountDtoWithDefaults() *TradeDiscountDto`

NewTradeDiscountDtoWithDefaults instantiates a new TradeDiscountDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TradeDiscountDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TradeDiscountDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TradeDiscountDto) SetId(v string)`

SetId sets Id field to given value.


### GetDeadline

`func (o *TradeDiscountDto) GetDeadline() int32`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *TradeDiscountDto) GetDeadlineOk() (*int32, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *TradeDiscountDto) SetDeadline(v int32)`

SetDeadline sets Deadline field to given value.


### GetRate

`func (o *TradeDiscountDto) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TradeDiscountDto) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TradeDiscountDto) SetRate(v float32)`

SetRate sets Rate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


