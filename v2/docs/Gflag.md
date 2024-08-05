# Gflag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clear** | Pointer to **NullableBool** | If Clear is set to true, the GFlag is removed | [optional] [default to false]
**Name** | Pointer to **string** | Specifies the name of the gflag. | [optional] 
**ProductModel** | Pointer to **string** | Specifies product model this gflag set on. | [optional] 
**Reason** | Pointer to **NullableString** | Specifies the reason for setting the gflag. | [optional] 
**Timestamp** | Pointer to **NullableInt64** | Specifies timestamp when gflag was set. | [optional] 
**Value** | Pointer to **string** | Specifies the value of the gflag. | [optional] 

## Methods

### NewGflag

`func NewGflag() *Gflag`

NewGflag instantiates a new Gflag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGflagWithDefaults

`func NewGflagWithDefaults() *Gflag`

NewGflagWithDefaults instantiates a new Gflag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClear

`func (o *Gflag) GetClear() bool`

GetClear returns the Clear field if non-nil, zero value otherwise.

### GetClearOk

`func (o *Gflag) GetClearOk() (*bool, bool)`

GetClearOk returns a tuple with the Clear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClear

`func (o *Gflag) SetClear(v bool)`

SetClear sets Clear field to given value.

### HasClear

`func (o *Gflag) HasClear() bool`

HasClear returns a boolean if a field has been set.

### SetClearNil

`func (o *Gflag) SetClearNil(b bool)`

 SetClearNil sets the value for Clear to be an explicit nil

### UnsetClear
`func (o *Gflag) UnsetClear()`

UnsetClear ensures that no value is present for Clear, not even an explicit nil
### GetName

`func (o *Gflag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gflag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gflag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Gflag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductModel

`func (o *Gflag) GetProductModel() string`

GetProductModel returns the ProductModel field if non-nil, zero value otherwise.

### GetProductModelOk

`func (o *Gflag) GetProductModelOk() (*string, bool)`

GetProductModelOk returns a tuple with the ProductModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductModel

`func (o *Gflag) SetProductModel(v string)`

SetProductModel sets ProductModel field to given value.

### HasProductModel

`func (o *Gflag) HasProductModel() bool`

HasProductModel returns a boolean if a field has been set.

### GetReason

`func (o *Gflag) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Gflag) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Gflag) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Gflag) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *Gflag) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *Gflag) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetTimestamp

`func (o *Gflag) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Gflag) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Gflag) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Gflag) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *Gflag) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Gflag) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetValue

`func (o *Gflag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Gflag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Gflag) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Gflag) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


