# InterfaceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RxBytes** | Pointer to **NullableInt64** | Total bytes received over the interface. | [optional] 
**RxDropped** | Pointer to **NullableInt64** | Number of packets received but not processed. | [optional] 
**RxErrors** | Pointer to **NullableInt64** | Total number of bad packets received. | [optional] 
**RxPkts** | Pointer to **NullableInt64** | Total number of packets received over the interface. | [optional] 
**TxBytes** | Pointer to **NullableInt64** | Total bytes transmitted over the interface. | [optional] 
**TxDropped** | Pointer to **NullableInt64** | Number of packets dropped on their way to transmission. | [optional] 
**TxErrors** | Pointer to **NullableInt64** | Total number of transmit problems. | [optional] 
**TxPkts** | Pointer to **NullableInt64** | Total number of packets transmitted over the interface. | [optional] 

## Methods

### NewInterfaceStats

`func NewInterfaceStats() *InterfaceStats`

NewInterfaceStats instantiates a new InterfaceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceStatsWithDefaults

`func NewInterfaceStatsWithDefaults() *InterfaceStats`

NewInterfaceStatsWithDefaults instantiates a new InterfaceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRxBytes

`func (o *InterfaceStats) GetRxBytes() int64`

GetRxBytes returns the RxBytes field if non-nil, zero value otherwise.

### GetRxBytesOk

`func (o *InterfaceStats) GetRxBytesOk() (*int64, bool)`

GetRxBytesOk returns a tuple with the RxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxBytes

`func (o *InterfaceStats) SetRxBytes(v int64)`

SetRxBytes sets RxBytes field to given value.

### HasRxBytes

`func (o *InterfaceStats) HasRxBytes() bool`

HasRxBytes returns a boolean if a field has been set.

### SetRxBytesNil

`func (o *InterfaceStats) SetRxBytesNil(b bool)`

 SetRxBytesNil sets the value for RxBytes to be an explicit nil

### UnsetRxBytes
`func (o *InterfaceStats) UnsetRxBytes()`

UnsetRxBytes ensures that no value is present for RxBytes, not even an explicit nil
### GetRxDropped

`func (o *InterfaceStats) GetRxDropped() int64`

GetRxDropped returns the RxDropped field if non-nil, zero value otherwise.

### GetRxDroppedOk

`func (o *InterfaceStats) GetRxDroppedOk() (*int64, bool)`

GetRxDroppedOk returns a tuple with the RxDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDropped

`func (o *InterfaceStats) SetRxDropped(v int64)`

SetRxDropped sets RxDropped field to given value.

### HasRxDropped

`func (o *InterfaceStats) HasRxDropped() bool`

HasRxDropped returns a boolean if a field has been set.

### SetRxDroppedNil

`func (o *InterfaceStats) SetRxDroppedNil(b bool)`

 SetRxDroppedNil sets the value for RxDropped to be an explicit nil

### UnsetRxDropped
`func (o *InterfaceStats) UnsetRxDropped()`

UnsetRxDropped ensures that no value is present for RxDropped, not even an explicit nil
### GetRxErrors

`func (o *InterfaceStats) GetRxErrors() int64`

GetRxErrors returns the RxErrors field if non-nil, zero value otherwise.

### GetRxErrorsOk

`func (o *InterfaceStats) GetRxErrorsOk() (*int64, bool)`

GetRxErrorsOk returns a tuple with the RxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxErrors

`func (o *InterfaceStats) SetRxErrors(v int64)`

SetRxErrors sets RxErrors field to given value.

### HasRxErrors

`func (o *InterfaceStats) HasRxErrors() bool`

HasRxErrors returns a boolean if a field has been set.

### SetRxErrorsNil

`func (o *InterfaceStats) SetRxErrorsNil(b bool)`

 SetRxErrorsNil sets the value for RxErrors to be an explicit nil

### UnsetRxErrors
`func (o *InterfaceStats) UnsetRxErrors()`

UnsetRxErrors ensures that no value is present for RxErrors, not even an explicit nil
### GetRxPkts

`func (o *InterfaceStats) GetRxPkts() int64`

GetRxPkts returns the RxPkts field if non-nil, zero value otherwise.

### GetRxPktsOk

`func (o *InterfaceStats) GetRxPktsOk() (*int64, bool)`

GetRxPktsOk returns a tuple with the RxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPkts

`func (o *InterfaceStats) SetRxPkts(v int64)`

SetRxPkts sets RxPkts field to given value.

### HasRxPkts

`func (o *InterfaceStats) HasRxPkts() bool`

HasRxPkts returns a boolean if a field has been set.

### SetRxPktsNil

`func (o *InterfaceStats) SetRxPktsNil(b bool)`

 SetRxPktsNil sets the value for RxPkts to be an explicit nil

### UnsetRxPkts
`func (o *InterfaceStats) UnsetRxPkts()`

UnsetRxPkts ensures that no value is present for RxPkts, not even an explicit nil
### GetTxBytes

`func (o *InterfaceStats) GetTxBytes() int64`

GetTxBytes returns the TxBytes field if non-nil, zero value otherwise.

### GetTxBytesOk

`func (o *InterfaceStats) GetTxBytesOk() (*int64, bool)`

GetTxBytesOk returns a tuple with the TxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxBytes

`func (o *InterfaceStats) SetTxBytes(v int64)`

SetTxBytes sets TxBytes field to given value.

### HasTxBytes

`func (o *InterfaceStats) HasTxBytes() bool`

HasTxBytes returns a boolean if a field has been set.

### SetTxBytesNil

`func (o *InterfaceStats) SetTxBytesNil(b bool)`

 SetTxBytesNil sets the value for TxBytes to be an explicit nil

### UnsetTxBytes
`func (o *InterfaceStats) UnsetTxBytes()`

UnsetTxBytes ensures that no value is present for TxBytes, not even an explicit nil
### GetTxDropped

`func (o *InterfaceStats) GetTxDropped() int64`

GetTxDropped returns the TxDropped field if non-nil, zero value otherwise.

### GetTxDroppedOk

`func (o *InterfaceStats) GetTxDroppedOk() (*int64, bool)`

GetTxDroppedOk returns a tuple with the TxDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxDropped

`func (o *InterfaceStats) SetTxDropped(v int64)`

SetTxDropped sets TxDropped field to given value.

### HasTxDropped

`func (o *InterfaceStats) HasTxDropped() bool`

HasTxDropped returns a boolean if a field has been set.

### SetTxDroppedNil

`func (o *InterfaceStats) SetTxDroppedNil(b bool)`

 SetTxDroppedNil sets the value for TxDropped to be an explicit nil

### UnsetTxDropped
`func (o *InterfaceStats) UnsetTxDropped()`

UnsetTxDropped ensures that no value is present for TxDropped, not even an explicit nil
### GetTxErrors

`func (o *InterfaceStats) GetTxErrors() int64`

GetTxErrors returns the TxErrors field if non-nil, zero value otherwise.

### GetTxErrorsOk

`func (o *InterfaceStats) GetTxErrorsOk() (*int64, bool)`

GetTxErrorsOk returns a tuple with the TxErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxErrors

`func (o *InterfaceStats) SetTxErrors(v int64)`

SetTxErrors sets TxErrors field to given value.

### HasTxErrors

`func (o *InterfaceStats) HasTxErrors() bool`

HasTxErrors returns a boolean if a field has been set.

### SetTxErrorsNil

`func (o *InterfaceStats) SetTxErrorsNil(b bool)`

 SetTxErrorsNil sets the value for TxErrors to be an explicit nil

### UnsetTxErrors
`func (o *InterfaceStats) UnsetTxErrors()`

UnsetTxErrors ensures that no value is present for TxErrors, not even an explicit nil
### GetTxPkts

`func (o *InterfaceStats) GetTxPkts() int64`

GetTxPkts returns the TxPkts field if non-nil, zero value otherwise.

### GetTxPktsOk

`func (o *InterfaceStats) GetTxPktsOk() (*int64, bool)`

GetTxPktsOk returns a tuple with the TxPkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPkts

`func (o *InterfaceStats) SetTxPkts(v int64)`

SetTxPkts sets TxPkts field to given value.

### HasTxPkts

`func (o *InterfaceStats) HasTxPkts() bool`

HasTxPkts returns a boolean if a field has been set.

### SetTxPktsNil

`func (o *InterfaceStats) SetTxPktsNil(b bool)`

 SetTxPktsNil sets the value for TxPkts to be an explicit nil

### UnsetTxPkts
`func (o *InterfaceStats) UnsetTxPkts()`

UnsetTxPkts ensures that no value is present for TxPkts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


