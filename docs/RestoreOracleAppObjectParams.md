# RestoreOracleAppObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateLocationParams** | Pointer to [**RestoreOracleAppObjectParamsAlternateLocationParams**](RestoreOracleAppObjectParamsAlternateLocationParams.md) |  | [optional] 
**NoOpenMode** | Pointer to **NullableBool** | If set to true, the recovered database will not be opened. | [optional] 
**OracleCloneAppViewParamsVec** | Pointer to [**[]CloneAppViewParams**](CloneAppViewParams.md) | Following field contains information related to view expose workflow. Ex mountpoint identifier specified by User from UI. | [optional] 
**OracleTargetParams** | Pointer to [**OracleSourceParams**](OracleSourceParams.md) |  | [optional] 
**ParallelOpEnabled** | Pointer to **NullableBool** | If set to true, parallel backups/restores/clones are enabled on same host. | [optional] 
**RestoreTimeSecs** | Pointer to **NullableInt64** | The time to which the Oracle database needs to be restored. This allows for granular recovery of Oracle databases. If this is not set, the Oracle database will be recovered to the full/incremental snapshot (specified in the owner&#39;s restore object in AppOwnerRestoreInfo). This is only applicable if restoring to the original Oracle instance. | [optional] 
**ShellEnvironmentVec** | Pointer to [**[]RestoreOracleAppObjectParamsKeyValuePair**](RestoreOracleAppObjectParamsKeyValuePair.md) |  | [optional] 

## Methods

### NewRestoreOracleAppObjectParams

`func NewRestoreOracleAppObjectParams() *RestoreOracleAppObjectParams`

NewRestoreOracleAppObjectParams instantiates a new RestoreOracleAppObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreOracleAppObjectParamsWithDefaults

`func NewRestoreOracleAppObjectParamsWithDefaults() *RestoreOracleAppObjectParams`

NewRestoreOracleAppObjectParamsWithDefaults instantiates a new RestoreOracleAppObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateLocationParams

`func (o *RestoreOracleAppObjectParams) GetAlternateLocationParams() RestoreOracleAppObjectParamsAlternateLocationParams`

GetAlternateLocationParams returns the AlternateLocationParams field if non-nil, zero value otherwise.

### GetAlternateLocationParamsOk

`func (o *RestoreOracleAppObjectParams) GetAlternateLocationParamsOk() (*RestoreOracleAppObjectParamsAlternateLocationParams, bool)`

GetAlternateLocationParamsOk returns a tuple with the AlternateLocationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateLocationParams

`func (o *RestoreOracleAppObjectParams) SetAlternateLocationParams(v RestoreOracleAppObjectParamsAlternateLocationParams)`

SetAlternateLocationParams sets AlternateLocationParams field to given value.

### HasAlternateLocationParams

`func (o *RestoreOracleAppObjectParams) HasAlternateLocationParams() bool`

HasAlternateLocationParams returns a boolean if a field has been set.

### GetNoOpenMode

`func (o *RestoreOracleAppObjectParams) GetNoOpenMode() bool`

GetNoOpenMode returns the NoOpenMode field if non-nil, zero value otherwise.

### GetNoOpenModeOk

`func (o *RestoreOracleAppObjectParams) GetNoOpenModeOk() (*bool, bool)`

GetNoOpenModeOk returns a tuple with the NoOpenMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOpenMode

`func (o *RestoreOracleAppObjectParams) SetNoOpenMode(v bool)`

SetNoOpenMode sets NoOpenMode field to given value.

### HasNoOpenMode

`func (o *RestoreOracleAppObjectParams) HasNoOpenMode() bool`

HasNoOpenMode returns a boolean if a field has been set.

### SetNoOpenModeNil

`func (o *RestoreOracleAppObjectParams) SetNoOpenModeNil(b bool)`

 SetNoOpenModeNil sets the value for NoOpenMode to be an explicit nil

### UnsetNoOpenMode
`func (o *RestoreOracleAppObjectParams) UnsetNoOpenMode()`

UnsetNoOpenMode ensures that no value is present for NoOpenMode, not even an explicit nil
### GetOracleCloneAppViewParamsVec

`func (o *RestoreOracleAppObjectParams) GetOracleCloneAppViewParamsVec() []CloneAppViewParams`

GetOracleCloneAppViewParamsVec returns the OracleCloneAppViewParamsVec field if non-nil, zero value otherwise.

### GetOracleCloneAppViewParamsVecOk

`func (o *RestoreOracleAppObjectParams) GetOracleCloneAppViewParamsVecOk() (*[]CloneAppViewParams, bool)`

GetOracleCloneAppViewParamsVecOk returns a tuple with the OracleCloneAppViewParamsVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleCloneAppViewParamsVec

`func (o *RestoreOracleAppObjectParams) SetOracleCloneAppViewParamsVec(v []CloneAppViewParams)`

SetOracleCloneAppViewParamsVec sets OracleCloneAppViewParamsVec field to given value.

### HasOracleCloneAppViewParamsVec

`func (o *RestoreOracleAppObjectParams) HasOracleCloneAppViewParamsVec() bool`

HasOracleCloneAppViewParamsVec returns a boolean if a field has been set.

### SetOracleCloneAppViewParamsVecNil

`func (o *RestoreOracleAppObjectParams) SetOracleCloneAppViewParamsVecNil(b bool)`

 SetOracleCloneAppViewParamsVecNil sets the value for OracleCloneAppViewParamsVec to be an explicit nil

### UnsetOracleCloneAppViewParamsVec
`func (o *RestoreOracleAppObjectParams) UnsetOracleCloneAppViewParamsVec()`

UnsetOracleCloneAppViewParamsVec ensures that no value is present for OracleCloneAppViewParamsVec, not even an explicit nil
### GetOracleTargetParams

`func (o *RestoreOracleAppObjectParams) GetOracleTargetParams() OracleSourceParams`

GetOracleTargetParams returns the OracleTargetParams field if non-nil, zero value otherwise.

### GetOracleTargetParamsOk

`func (o *RestoreOracleAppObjectParams) GetOracleTargetParamsOk() (*OracleSourceParams, bool)`

GetOracleTargetParamsOk returns a tuple with the OracleTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleTargetParams

`func (o *RestoreOracleAppObjectParams) SetOracleTargetParams(v OracleSourceParams)`

SetOracleTargetParams sets OracleTargetParams field to given value.

### HasOracleTargetParams

`func (o *RestoreOracleAppObjectParams) HasOracleTargetParams() bool`

HasOracleTargetParams returns a boolean if a field has been set.

### GetParallelOpEnabled

`func (o *RestoreOracleAppObjectParams) GetParallelOpEnabled() bool`

GetParallelOpEnabled returns the ParallelOpEnabled field if non-nil, zero value otherwise.

### GetParallelOpEnabledOk

`func (o *RestoreOracleAppObjectParams) GetParallelOpEnabledOk() (*bool, bool)`

GetParallelOpEnabledOk returns a tuple with the ParallelOpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelOpEnabled

`func (o *RestoreOracleAppObjectParams) SetParallelOpEnabled(v bool)`

SetParallelOpEnabled sets ParallelOpEnabled field to given value.

### HasParallelOpEnabled

`func (o *RestoreOracleAppObjectParams) HasParallelOpEnabled() bool`

HasParallelOpEnabled returns a boolean if a field has been set.

### SetParallelOpEnabledNil

`func (o *RestoreOracleAppObjectParams) SetParallelOpEnabledNil(b bool)`

 SetParallelOpEnabledNil sets the value for ParallelOpEnabled to be an explicit nil

### UnsetParallelOpEnabled
`func (o *RestoreOracleAppObjectParams) UnsetParallelOpEnabled()`

UnsetParallelOpEnabled ensures that no value is present for ParallelOpEnabled, not even an explicit nil
### GetRestoreTimeSecs

`func (o *RestoreOracleAppObjectParams) GetRestoreTimeSecs() int64`

GetRestoreTimeSecs returns the RestoreTimeSecs field if non-nil, zero value otherwise.

### GetRestoreTimeSecsOk

`func (o *RestoreOracleAppObjectParams) GetRestoreTimeSecsOk() (*int64, bool)`

GetRestoreTimeSecsOk returns a tuple with the RestoreTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTimeSecs

`func (o *RestoreOracleAppObjectParams) SetRestoreTimeSecs(v int64)`

SetRestoreTimeSecs sets RestoreTimeSecs field to given value.

### HasRestoreTimeSecs

`func (o *RestoreOracleAppObjectParams) HasRestoreTimeSecs() bool`

HasRestoreTimeSecs returns a boolean if a field has been set.

### SetRestoreTimeSecsNil

`func (o *RestoreOracleAppObjectParams) SetRestoreTimeSecsNil(b bool)`

 SetRestoreTimeSecsNil sets the value for RestoreTimeSecs to be an explicit nil

### UnsetRestoreTimeSecs
`func (o *RestoreOracleAppObjectParams) UnsetRestoreTimeSecs()`

UnsetRestoreTimeSecs ensures that no value is present for RestoreTimeSecs, not even an explicit nil
### GetShellEnvironmentVec

`func (o *RestoreOracleAppObjectParams) GetShellEnvironmentVec() []RestoreOracleAppObjectParamsKeyValuePair`

GetShellEnvironmentVec returns the ShellEnvironmentVec field if non-nil, zero value otherwise.

### GetShellEnvironmentVecOk

`func (o *RestoreOracleAppObjectParams) GetShellEnvironmentVecOk() (*[]RestoreOracleAppObjectParamsKeyValuePair, bool)`

GetShellEnvironmentVecOk returns a tuple with the ShellEnvironmentVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellEnvironmentVec

`func (o *RestoreOracleAppObjectParams) SetShellEnvironmentVec(v []RestoreOracleAppObjectParamsKeyValuePair)`

SetShellEnvironmentVec sets ShellEnvironmentVec field to given value.

### HasShellEnvironmentVec

`func (o *RestoreOracleAppObjectParams) HasShellEnvironmentVec() bool`

HasShellEnvironmentVec returns a boolean if a field has been set.

### SetShellEnvironmentVecNil

`func (o *RestoreOracleAppObjectParams) SetShellEnvironmentVecNil(b bool)`

 SetShellEnvironmentVecNil sets the value for ShellEnvironmentVec to be an explicit nil

### UnsetShellEnvironmentVec
`func (o *RestoreOracleAppObjectParams) UnsetShellEnvironmentVec()`

UnsetShellEnvironmentVec ensures that no value is present for ShellEnvironmentVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


