# BgpTimers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepAlive** | Pointer to **NullableInt32** | Keep alive interval in seconds. | [optional] 
**HoldTime** | Pointer to **NullableInt32** | Hold time in seconds. | [optional] 

## Methods

### NewBgpTimers

`func NewBgpTimers() *BgpTimers`

NewBgpTimers instantiates a new BgpTimers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpTimersWithDefaults

`func NewBgpTimersWithDefaults() *BgpTimers`

NewBgpTimersWithDefaults instantiates a new BgpTimers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeepAlive

`func (o *BgpTimers) GetKeepAlive() int32`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *BgpTimers) GetKeepAliveOk() (*int32, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *BgpTimers) SetKeepAlive(v int32)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *BgpTimers) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### SetKeepAliveNil

`func (o *BgpTimers) SetKeepAliveNil(b bool)`

 SetKeepAliveNil sets the value for KeepAlive to be an explicit nil

### UnsetKeepAlive
`func (o *BgpTimers) UnsetKeepAlive()`

UnsetKeepAlive ensures that no value is present for KeepAlive, not even an explicit nil
### GetHoldTime

`func (o *BgpTimers) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *BgpTimers) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *BgpTimers) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *BgpTimers) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### SetHoldTimeNil

`func (o *BgpTimers) SetHoldTimeNil(b bool)`

 SetHoldTimeNil sets the value for HoldTime to be an explicit nil

### UnsetHoldTime
`func (o *BgpTimers) UnsetHoldTime()`

UnsetHoldTime ensures that no value is present for HoldTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


