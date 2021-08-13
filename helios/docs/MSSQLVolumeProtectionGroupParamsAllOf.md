# MSSQLVolumeProtectionGroupParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]MSSQLVolumeProtectionGroupObjectParams**](MSSQLVolumeProtectionGroupObjectParams.md) | Specifies the list of object ids to be protected. | 
**IncrementalBackupAfterRestart** | Pointer to **NullableBool** | Specifies whether or to perform incremental backups the first time after a server restarts. By default, a full backup will be performed. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**BackupDbVolumesOnly** | Pointer to **NullableBool** | Specifies whether to only backup volumes on which the specified databases reside. If not specified (default), all the volumes of the host will be protected. | [optional] 
**AdditionalHostParams** | Pointer to [**[]MSSQLVolumeProtectionGroupHostParams**](MSSQLVolumeProtectionGroupHostParams.md) | Specifies settings which are to be applied to specific host containers in this protection group. | [optional] 

## Methods

### NewMSSQLVolumeProtectionGroupParamsAllOf

`func NewMSSQLVolumeProtectionGroupParamsAllOf(objects []MSSQLVolumeProtectionGroupObjectParams, ) *MSSQLVolumeProtectionGroupParamsAllOf`

NewMSSQLVolumeProtectionGroupParamsAllOf instantiates a new MSSQLVolumeProtectionGroupParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSSQLVolumeProtectionGroupParamsAllOfWithDefaults

`func NewMSSQLVolumeProtectionGroupParamsAllOfWithDefaults() *MSSQLVolumeProtectionGroupParamsAllOf`

NewMSSQLVolumeProtectionGroupParamsAllOfWithDefaults instantiates a new MSSQLVolumeProtectionGroupParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetObjects() []MSSQLVolumeProtectionGroupObjectParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetObjectsOk() (*[]MSSQLVolumeProtectionGroupObjectParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetObjects(v []MSSQLVolumeProtectionGroupObjectParams)`

SetObjects sets Objects field to given value.


### SetObjectsNil

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *MSSQLVolumeProtectionGroupParamsAllOf) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetIncrementalBackupAfterRestart

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetIncrementalBackupAfterRestart() bool`

GetIncrementalBackupAfterRestart returns the IncrementalBackupAfterRestart field if non-nil, zero value otherwise.

### GetIncrementalBackupAfterRestartOk

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetIncrementalBackupAfterRestartOk() (*bool, bool)`

GetIncrementalBackupAfterRestartOk returns a tuple with the IncrementalBackupAfterRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupAfterRestart

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetIncrementalBackupAfterRestart(v bool)`

SetIncrementalBackupAfterRestart sets IncrementalBackupAfterRestart field to given value.

### HasIncrementalBackupAfterRestart

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) HasIncrementalBackupAfterRestart() bool`

HasIncrementalBackupAfterRestart returns a boolean if a field has been set.

### SetIncrementalBackupAfterRestartNil

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetIncrementalBackupAfterRestartNil(b bool)`

 SetIncrementalBackupAfterRestartNil sets the value for IncrementalBackupAfterRestart to be an explicit nil

### UnsetIncrementalBackupAfterRestart
`func (o *MSSQLVolumeProtectionGroupParamsAllOf) UnsetIncrementalBackupAfterRestart()`

UnsetIncrementalBackupAfterRestart ensures that no value is present for IncrementalBackupAfterRestart, not even an explicit nil
### GetIndexingPolicy

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetBackupDbVolumesOnly

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetBackupDbVolumesOnly() bool`

GetBackupDbVolumesOnly returns the BackupDbVolumesOnly field if non-nil, zero value otherwise.

### GetBackupDbVolumesOnlyOk

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetBackupDbVolumesOnlyOk() (*bool, bool)`

GetBackupDbVolumesOnlyOk returns a tuple with the BackupDbVolumesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDbVolumesOnly

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetBackupDbVolumesOnly(v bool)`

SetBackupDbVolumesOnly sets BackupDbVolumesOnly field to given value.

### HasBackupDbVolumesOnly

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) HasBackupDbVolumesOnly() bool`

HasBackupDbVolumesOnly returns a boolean if a field has been set.

### SetBackupDbVolumesOnlyNil

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetBackupDbVolumesOnlyNil(b bool)`

 SetBackupDbVolumesOnlyNil sets the value for BackupDbVolumesOnly to be an explicit nil

### UnsetBackupDbVolumesOnly
`func (o *MSSQLVolumeProtectionGroupParamsAllOf) UnsetBackupDbVolumesOnly()`

UnsetBackupDbVolumesOnly ensures that no value is present for BackupDbVolumesOnly, not even an explicit nil
### GetAdditionalHostParams

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetAdditionalHostParams() []MSSQLVolumeProtectionGroupHostParams`

GetAdditionalHostParams returns the AdditionalHostParams field if non-nil, zero value otherwise.

### GetAdditionalHostParamsOk

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) GetAdditionalHostParamsOk() (*[]MSSQLVolumeProtectionGroupHostParams, bool)`

GetAdditionalHostParamsOk returns a tuple with the AdditionalHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHostParams

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) SetAdditionalHostParams(v []MSSQLVolumeProtectionGroupHostParams)`

SetAdditionalHostParams sets AdditionalHostParams field to given value.

### HasAdditionalHostParams

`func (o *MSSQLVolumeProtectionGroupParamsAllOf) HasAdditionalHostParams() bool`

HasAdditionalHostParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


