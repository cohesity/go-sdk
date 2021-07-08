# Gflag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies name of the gflag. | 
**ProductModel** | Pointer to **NullableString** | Specifies product model this gflag set on. | [optional] 
**Reason** | Pointer to **NullableString** | Specifies reason for setting the gflag. | [optional] 
**Timestamp** | Pointer to **NullableInt64** | Specifies timestamp when gflag was set. | [optional] [readonly] 
**Value** | Pointer to **NullableString** | Specifies value of the gflag. | [optional] 

## Methods

### NewGflag

`func NewGflag(name NullableString, ) *Gflag`

NewGflag instantiates a new Gflag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGflagWithDefaults

`func NewGflagWithDefaults() *Gflag`

NewGflagWithDefaults instantiates a new Gflag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### SetNameNil

`func (o *Gflag) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Gflag) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### SetProductModelNil

`func (o *Gflag) SetProductModelNil(b bool)`

 SetProductModelNil sets the value for ProductModel to be an explicit nil

### UnsetProductModel
`func (o *Gflag) UnsetProductModel()`

UnsetProductModel ensures that no value is present for ProductModel, not even an explicit nil
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

### SetValueNil

`func (o *Gflag) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Gflag) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


