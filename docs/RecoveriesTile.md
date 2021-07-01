# RecoveriesTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastMonthNumRecoveries** | Pointer to **NullableInt32** | Number of Recoveries in the last 30 days. | [optional] 
**LastMonthRecoveriesByType** | Pointer to [**[]RestoreCountByObjectType**](RestoreCountByObjectType.md) | Recoveries by Type in the last month. | [optional] 
**LastMonthRecoverySizeBytes** | Pointer to **NullableInt64** | Bytes recovered in the last 30 days. | [optional] 
**RecoveryNumRunning** | Pointer to **NullableInt32** | Number of recoveries that are currently running. | [optional] 

## Methods

### NewRecoveriesTile

`func NewRecoveriesTile() *RecoveriesTile`

NewRecoveriesTile instantiates a new RecoveriesTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveriesTileWithDefaults

`func NewRecoveriesTileWithDefaults() *RecoveriesTile`

NewRecoveriesTileWithDefaults instantiates a new RecoveriesTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastMonthNumRecoveries

`func (o *RecoveriesTile) GetLastMonthNumRecoveries() int32`

GetLastMonthNumRecoveries returns the LastMonthNumRecoveries field if non-nil, zero value otherwise.

### GetLastMonthNumRecoveriesOk

`func (o *RecoveriesTile) GetLastMonthNumRecoveriesOk() (*int32, bool)`

GetLastMonthNumRecoveriesOk returns a tuple with the LastMonthNumRecoveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMonthNumRecoveries

`func (o *RecoveriesTile) SetLastMonthNumRecoveries(v int32)`

SetLastMonthNumRecoveries sets LastMonthNumRecoveries field to given value.

### HasLastMonthNumRecoveries

`func (o *RecoveriesTile) HasLastMonthNumRecoveries() bool`

HasLastMonthNumRecoveries returns a boolean if a field has been set.

### SetLastMonthNumRecoveriesNil

`func (o *RecoveriesTile) SetLastMonthNumRecoveriesNil(b bool)`

 SetLastMonthNumRecoveriesNil sets the value for LastMonthNumRecoveries to be an explicit nil

### UnsetLastMonthNumRecoveries
`func (o *RecoveriesTile) UnsetLastMonthNumRecoveries()`

UnsetLastMonthNumRecoveries ensures that no value is present for LastMonthNumRecoveries, not even an explicit nil
### GetLastMonthRecoveriesByType

`func (o *RecoveriesTile) GetLastMonthRecoveriesByType() []RestoreCountByObjectType`

GetLastMonthRecoveriesByType returns the LastMonthRecoveriesByType field if non-nil, zero value otherwise.

### GetLastMonthRecoveriesByTypeOk

`func (o *RecoveriesTile) GetLastMonthRecoveriesByTypeOk() (*[]RestoreCountByObjectType, bool)`

GetLastMonthRecoveriesByTypeOk returns a tuple with the LastMonthRecoveriesByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMonthRecoveriesByType

`func (o *RecoveriesTile) SetLastMonthRecoveriesByType(v []RestoreCountByObjectType)`

SetLastMonthRecoveriesByType sets LastMonthRecoveriesByType field to given value.

### HasLastMonthRecoveriesByType

`func (o *RecoveriesTile) HasLastMonthRecoveriesByType() bool`

HasLastMonthRecoveriesByType returns a boolean if a field has been set.

### SetLastMonthRecoveriesByTypeNil

`func (o *RecoveriesTile) SetLastMonthRecoveriesByTypeNil(b bool)`

 SetLastMonthRecoveriesByTypeNil sets the value for LastMonthRecoveriesByType to be an explicit nil

### UnsetLastMonthRecoveriesByType
`func (o *RecoveriesTile) UnsetLastMonthRecoveriesByType()`

UnsetLastMonthRecoveriesByType ensures that no value is present for LastMonthRecoveriesByType, not even an explicit nil
### GetLastMonthRecoverySizeBytes

`func (o *RecoveriesTile) GetLastMonthRecoverySizeBytes() int64`

GetLastMonthRecoverySizeBytes returns the LastMonthRecoverySizeBytes field if non-nil, zero value otherwise.

### GetLastMonthRecoverySizeBytesOk

`func (o *RecoveriesTile) GetLastMonthRecoverySizeBytesOk() (*int64, bool)`

GetLastMonthRecoverySizeBytesOk returns a tuple with the LastMonthRecoverySizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMonthRecoverySizeBytes

`func (o *RecoveriesTile) SetLastMonthRecoverySizeBytes(v int64)`

SetLastMonthRecoverySizeBytes sets LastMonthRecoverySizeBytes field to given value.

### HasLastMonthRecoverySizeBytes

`func (o *RecoveriesTile) HasLastMonthRecoverySizeBytes() bool`

HasLastMonthRecoverySizeBytes returns a boolean if a field has been set.

### SetLastMonthRecoverySizeBytesNil

`func (o *RecoveriesTile) SetLastMonthRecoverySizeBytesNil(b bool)`

 SetLastMonthRecoverySizeBytesNil sets the value for LastMonthRecoverySizeBytes to be an explicit nil

### UnsetLastMonthRecoverySizeBytes
`func (o *RecoveriesTile) UnsetLastMonthRecoverySizeBytes()`

UnsetLastMonthRecoverySizeBytes ensures that no value is present for LastMonthRecoverySizeBytes, not even an explicit nil
### GetRecoveryNumRunning

`func (o *RecoveriesTile) GetRecoveryNumRunning() int32`

GetRecoveryNumRunning returns the RecoveryNumRunning field if non-nil, zero value otherwise.

### GetRecoveryNumRunningOk

`func (o *RecoveriesTile) GetRecoveryNumRunningOk() (*int32, bool)`

GetRecoveryNumRunningOk returns a tuple with the RecoveryNumRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryNumRunning

`func (o *RecoveriesTile) SetRecoveryNumRunning(v int32)`

SetRecoveryNumRunning sets RecoveryNumRunning field to given value.

### HasRecoveryNumRunning

`func (o *RecoveriesTile) HasRecoveryNumRunning() bool`

HasRecoveryNumRunning returns a boolean if a field has been set.

### SetRecoveryNumRunningNil

`func (o *RecoveriesTile) SetRecoveryNumRunningNil(b bool)`

 SetRecoveryNumRunningNil sets the value for RecoveryNumRunning to be an explicit nil

### UnsetRecoveryNumRunning
`func (o *RecoveriesTile) UnsetRecoveryNumRunning()`

UnsetRecoveryNumRunning ensures that no value is present for RecoveryNumRunning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


