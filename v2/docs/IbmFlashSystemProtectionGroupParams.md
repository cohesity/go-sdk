# IbmFlashSystemProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSafeGuardedCopySnapshot** | Pointer to **NullableBool** | Specifies whether the safeguarded copy snapshots are allowed or not | [optional] 
**Objects** | [**[]IbmFlashSystemProtectionGroupObjectParams**](IbmFlashSystemProtectionGroupObjectParams.md) | Specifies the objects to be included in the Protection Group. | 
**PrePostScript** | Pointer to [**HostBasedBackupScriptParams**](HostBasedBackupScriptParams.md) |  | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the id of the parent of the objects. | [optional] [readonly] 
**SourceName** | Pointer to **NullableString** | Specifies the name of the parent of the objects. | [optional] [readonly] 

## Methods

### NewIbmFlashSystemProtectionGroupParams

`func NewIbmFlashSystemProtectionGroupParams(objects []IbmFlashSystemProtectionGroupObjectParams, ) *IbmFlashSystemProtectionGroupParams`

NewIbmFlashSystemProtectionGroupParams instantiates a new IbmFlashSystemProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbmFlashSystemProtectionGroupParamsWithDefaults

`func NewIbmFlashSystemProtectionGroupParamsWithDefaults() *IbmFlashSystemProtectionGroupParams`

NewIbmFlashSystemProtectionGroupParamsWithDefaults instantiates a new IbmFlashSystemProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSafeGuardedCopySnapshot

`func (o *IbmFlashSystemProtectionGroupParams) GetIsSafeGuardedCopySnapshot() bool`

GetIsSafeGuardedCopySnapshot returns the IsSafeGuardedCopySnapshot field if non-nil, zero value otherwise.

### GetIsSafeGuardedCopySnapshotOk

`func (o *IbmFlashSystemProtectionGroupParams) GetIsSafeGuardedCopySnapshotOk() (*bool, bool)`

GetIsSafeGuardedCopySnapshotOk returns a tuple with the IsSafeGuardedCopySnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafeGuardedCopySnapshot

`func (o *IbmFlashSystemProtectionGroupParams) SetIsSafeGuardedCopySnapshot(v bool)`

SetIsSafeGuardedCopySnapshot sets IsSafeGuardedCopySnapshot field to given value.

### HasIsSafeGuardedCopySnapshot

`func (o *IbmFlashSystemProtectionGroupParams) HasIsSafeGuardedCopySnapshot() bool`

HasIsSafeGuardedCopySnapshot returns a boolean if a field has been set.

### SetIsSafeGuardedCopySnapshotNil

`func (o *IbmFlashSystemProtectionGroupParams) SetIsSafeGuardedCopySnapshotNil(b bool)`

 SetIsSafeGuardedCopySnapshotNil sets the value for IsSafeGuardedCopySnapshot to be an explicit nil

### UnsetIsSafeGuardedCopySnapshot
`func (o *IbmFlashSystemProtectionGroupParams) UnsetIsSafeGuardedCopySnapshot()`

UnsetIsSafeGuardedCopySnapshot ensures that no value is present for IsSafeGuardedCopySnapshot, not even an explicit nil
### GetObjects

`func (o *IbmFlashSystemProtectionGroupParams) GetObjects() []IbmFlashSystemProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *IbmFlashSystemProtectionGroupParams) GetObjectsOk() (*[]IbmFlashSystemProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *IbmFlashSystemProtectionGroupParams) SetObjects(v []IbmFlashSystemProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### GetPrePostScript

`func (o *IbmFlashSystemProtectionGroupParams) GetPrePostScript() HostBasedBackupScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *IbmFlashSystemProtectionGroupParams) GetPrePostScriptOk() (*HostBasedBackupScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *IbmFlashSystemProtectionGroupParams) SetPrePostScript(v HostBasedBackupScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *IbmFlashSystemProtectionGroupParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetSourceId

`func (o *IbmFlashSystemProtectionGroupParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *IbmFlashSystemProtectionGroupParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *IbmFlashSystemProtectionGroupParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *IbmFlashSystemProtectionGroupParams) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *IbmFlashSystemProtectionGroupParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *IbmFlashSystemProtectionGroupParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *IbmFlashSystemProtectionGroupParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *IbmFlashSystemProtectionGroupParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *IbmFlashSystemProtectionGroupParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *IbmFlashSystemProtectionGroupParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *IbmFlashSystemProtectionGroupParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *IbmFlashSystemProtectionGroupParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


