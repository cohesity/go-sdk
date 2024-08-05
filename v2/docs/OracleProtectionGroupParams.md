# OracleProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullAutoKillTimeoutSecs** | Pointer to **NullableInt64** | Time in seconds after which the full backup of the database in given backup job should be auto-killed. | [optional] 
**IncrAutoKillTimeoutSecs** | Pointer to **NullableInt64** | Time in seconds after which the incremental backup of the database in given backup job should be auto-killed. | [optional] 
**LogAutoKillTimeoutSecs** | Pointer to **NullableInt64** | Time in seconds after which the log backup of the database in given backup job should be auto-killed. | [optional] 
**Objects** | [**[]OracleProtectionGroupObjectParams**](OracleProtectionGroupObjectParams.md) | Specifies the list of object ids to be protected. | 
**PersistMountpoints** | Pointer to **NullableBool** | Specifies whether the mountpoints created while backing up Oracle DBs should be persisted. Defaults to true if value is null to handle the backward compatibility for the upgrade case. | [optional] [default to true]
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**VlanParams** | Pointer to [**VlanParams**](VlanParams.md) |  | [optional] 

## Methods

### NewOracleProtectionGroupParams

`func NewOracleProtectionGroupParams(objects []OracleProtectionGroupObjectParams, ) *OracleProtectionGroupParams`

NewOracleProtectionGroupParams instantiates a new OracleProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleProtectionGroupParamsWithDefaults

`func NewOracleProtectionGroupParamsWithDefaults() *OracleProtectionGroupParams`

NewOracleProtectionGroupParamsWithDefaults instantiates a new OracleProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullAutoKillTimeoutSecs

`func (o *OracleProtectionGroupParams) GetFullAutoKillTimeoutSecs() int64`

GetFullAutoKillTimeoutSecs returns the FullAutoKillTimeoutSecs field if non-nil, zero value otherwise.

### GetFullAutoKillTimeoutSecsOk

`func (o *OracleProtectionGroupParams) GetFullAutoKillTimeoutSecsOk() (*int64, bool)`

GetFullAutoKillTimeoutSecsOk returns a tuple with the FullAutoKillTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAutoKillTimeoutSecs

`func (o *OracleProtectionGroupParams) SetFullAutoKillTimeoutSecs(v int64)`

SetFullAutoKillTimeoutSecs sets FullAutoKillTimeoutSecs field to given value.

### HasFullAutoKillTimeoutSecs

`func (o *OracleProtectionGroupParams) HasFullAutoKillTimeoutSecs() bool`

HasFullAutoKillTimeoutSecs returns a boolean if a field has been set.

### SetFullAutoKillTimeoutSecsNil

`func (o *OracleProtectionGroupParams) SetFullAutoKillTimeoutSecsNil(b bool)`

 SetFullAutoKillTimeoutSecsNil sets the value for FullAutoKillTimeoutSecs to be an explicit nil

### UnsetFullAutoKillTimeoutSecs
`func (o *OracleProtectionGroupParams) UnsetFullAutoKillTimeoutSecs()`

UnsetFullAutoKillTimeoutSecs ensures that no value is present for FullAutoKillTimeoutSecs, not even an explicit nil
### GetIncrAutoKillTimeoutSecs

`func (o *OracleProtectionGroupParams) GetIncrAutoKillTimeoutSecs() int64`

GetIncrAutoKillTimeoutSecs returns the IncrAutoKillTimeoutSecs field if non-nil, zero value otherwise.

### GetIncrAutoKillTimeoutSecsOk

`func (o *OracleProtectionGroupParams) GetIncrAutoKillTimeoutSecsOk() (*int64, bool)`

GetIncrAutoKillTimeoutSecsOk returns a tuple with the IncrAutoKillTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrAutoKillTimeoutSecs

`func (o *OracleProtectionGroupParams) SetIncrAutoKillTimeoutSecs(v int64)`

SetIncrAutoKillTimeoutSecs sets IncrAutoKillTimeoutSecs field to given value.

### HasIncrAutoKillTimeoutSecs

`func (o *OracleProtectionGroupParams) HasIncrAutoKillTimeoutSecs() bool`

HasIncrAutoKillTimeoutSecs returns a boolean if a field has been set.

### SetIncrAutoKillTimeoutSecsNil

`func (o *OracleProtectionGroupParams) SetIncrAutoKillTimeoutSecsNil(b bool)`

 SetIncrAutoKillTimeoutSecsNil sets the value for IncrAutoKillTimeoutSecs to be an explicit nil

### UnsetIncrAutoKillTimeoutSecs
`func (o *OracleProtectionGroupParams) UnsetIncrAutoKillTimeoutSecs()`

UnsetIncrAutoKillTimeoutSecs ensures that no value is present for IncrAutoKillTimeoutSecs, not even an explicit nil
### GetLogAutoKillTimeoutSecs

`func (o *OracleProtectionGroupParams) GetLogAutoKillTimeoutSecs() int64`

GetLogAutoKillTimeoutSecs returns the LogAutoKillTimeoutSecs field if non-nil, zero value otherwise.

### GetLogAutoKillTimeoutSecsOk

`func (o *OracleProtectionGroupParams) GetLogAutoKillTimeoutSecsOk() (*int64, bool)`

GetLogAutoKillTimeoutSecsOk returns a tuple with the LogAutoKillTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAutoKillTimeoutSecs

`func (o *OracleProtectionGroupParams) SetLogAutoKillTimeoutSecs(v int64)`

SetLogAutoKillTimeoutSecs sets LogAutoKillTimeoutSecs field to given value.

### HasLogAutoKillTimeoutSecs

`func (o *OracleProtectionGroupParams) HasLogAutoKillTimeoutSecs() bool`

HasLogAutoKillTimeoutSecs returns a boolean if a field has been set.

### SetLogAutoKillTimeoutSecsNil

`func (o *OracleProtectionGroupParams) SetLogAutoKillTimeoutSecsNil(b bool)`

 SetLogAutoKillTimeoutSecsNil sets the value for LogAutoKillTimeoutSecs to be an explicit nil

### UnsetLogAutoKillTimeoutSecs
`func (o *OracleProtectionGroupParams) UnsetLogAutoKillTimeoutSecs()`

UnsetLogAutoKillTimeoutSecs ensures that no value is present for LogAutoKillTimeoutSecs, not even an explicit nil
### GetObjects

`func (o *OracleProtectionGroupParams) GetObjects() []OracleProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *OracleProtectionGroupParams) GetObjectsOk() (*[]OracleProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *OracleProtectionGroupParams) SetObjects(v []OracleProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *OracleProtectionGroupParams) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *OracleProtectionGroupParams) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPersistMountpoints

`func (o *OracleProtectionGroupParams) GetPersistMountpoints() bool`

GetPersistMountpoints returns the PersistMountpoints field if non-nil, zero value otherwise.

### GetPersistMountpointsOk

`func (o *OracleProtectionGroupParams) GetPersistMountpointsOk() (*bool, bool)`

GetPersistMountpointsOk returns a tuple with the PersistMountpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMountpoints

`func (o *OracleProtectionGroupParams) SetPersistMountpoints(v bool)`

SetPersistMountpoints sets PersistMountpoints field to given value.

### HasPersistMountpoints

`func (o *OracleProtectionGroupParams) HasPersistMountpoints() bool`

HasPersistMountpoints returns a boolean if a field has been set.

### SetPersistMountpointsNil

`func (o *OracleProtectionGroupParams) SetPersistMountpointsNil(b bool)`

 SetPersistMountpointsNil sets the value for PersistMountpoints to be an explicit nil

### UnsetPersistMountpoints
`func (o *OracleProtectionGroupParams) UnsetPersistMountpoints()`

UnsetPersistMountpoints ensures that no value is present for PersistMountpoints, not even an explicit nil
### GetPrePostScript

`func (o *OracleProtectionGroupParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *OracleProtectionGroupParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *OracleProtectionGroupParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *OracleProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetVlanParams

`func (o *OracleProtectionGroupParams) GetVlanParams() VlanParams`

GetVlanParams returns the VlanParams field if non-nil, zero value otherwise.

### GetVlanParamsOk

`func (o *OracleProtectionGroupParams) GetVlanParamsOk() (*VlanParams, bool)`

GetVlanParamsOk returns a tuple with the VlanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParams

`func (o *OracleProtectionGroupParams) SetVlanParams(v VlanParams)`

SetVlanParams sets VlanParams field to given value.

### HasVlanParams

`func (o *OracleProtectionGroupParams) HasVlanParams() bool`

HasVlanParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


