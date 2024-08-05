# ProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPolicy** | [**BackupPolicy**](BackupPolicy.md) |  | 
**BlackoutWindow** | Pointer to [**[]BlackoutWindow**](BlackoutWindow.md) | List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod. | [optional] 
**CascadedTargetsConfig** | Pointer to [**[]CascadedTargetConfiguration**](CascadedTargetConfiguration.md) | Specifies the configuration for cascaded replications. Using cascaded replication, replication cluster(Rx) can further replicate and archive the snapshot copies to further targets. Its recommended to create cascaded configuration where protection group will be created. | [optional] 
**DataLock** | Pointer to **NullableString** | This field is now deprecated. Please use the DataLockConfig in the backup retention. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the Protection Policy. | [optional] 
**ExtendedRetention** | Pointer to [**[]ExtendedRetentionPolicy**](ExtendedRetentionPolicy.md) | Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it. | [optional] 
**IsCBSEnabled** | Pointer to **NullableBool** | Specifies true if Calender Based Schedule is supported by client. Default value is assumed as false for this feature. | [optional] 
**LastModificationTimeUsecs** | Pointer to **NullableInt64** | Specifies the last time this Policy was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the policy was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error. | [optional] 
**Name** | **NullableString** | Specifies the name of the Protection Policy. | 
**RemoteTargetPolicy** | Pointer to [**TargetsConfiguration**](TargetsConfiguration.md) |  | [optional] 
**RetryOptions** | Pointer to [**RetryOptions**](RetryOptions.md) |  | [optional] 
**RpoPolicySettings** | Pointer to [**RpoPolicySettings**](RpoPolicySettings.md) |  | [optional] 
**SkipIntervalMins** | Pointer to **NullableInt32** | Specifies the period of time before skipping the execution of new group Runs if an existing queued group Run of the same Protection group has not started. For example if this field is set to 30 minutes and a group Run is scheduled to start at 5:00 AM every day but does not start due to conflicts (such as too many groups are running). If the new group Run does not start by 5:30AM, the Cohesity Cluster will skip the new group Run. If the original group Run completes before 5:30AM the next day, a new group Run is created and starts executing. This field is optional. | [optional] 
**Version** | Pointer to **NullableInt32** | Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases. | [optional] 

## Methods

### NewProtectionPolicy

`func NewProtectionPolicy(backupPolicy BackupPolicy, name NullableString, ) *ProtectionPolicy`

NewProtectionPolicy instantiates a new ProtectionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionPolicyWithDefaults

`func NewProtectionPolicyWithDefaults() *ProtectionPolicy`

NewProtectionPolicyWithDefaults instantiates a new ProtectionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicy

`func (o *ProtectionPolicy) GetBackupPolicy() BackupPolicy`

GetBackupPolicy returns the BackupPolicy field if non-nil, zero value otherwise.

### GetBackupPolicyOk

`func (o *ProtectionPolicy) GetBackupPolicyOk() (*BackupPolicy, bool)`

GetBackupPolicyOk returns a tuple with the BackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicy

`func (o *ProtectionPolicy) SetBackupPolicy(v BackupPolicy)`

SetBackupPolicy sets BackupPolicy field to given value.


### GetBlackoutWindow

`func (o *ProtectionPolicy) GetBlackoutWindow() []BlackoutWindow`

GetBlackoutWindow returns the BlackoutWindow field if non-nil, zero value otherwise.

### GetBlackoutWindowOk

`func (o *ProtectionPolicy) GetBlackoutWindowOk() (*[]BlackoutWindow, bool)`

GetBlackoutWindowOk returns a tuple with the BlackoutWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutWindow

`func (o *ProtectionPolicy) SetBlackoutWindow(v []BlackoutWindow)`

SetBlackoutWindow sets BlackoutWindow field to given value.

### HasBlackoutWindow

`func (o *ProtectionPolicy) HasBlackoutWindow() bool`

HasBlackoutWindow returns a boolean if a field has been set.

### SetBlackoutWindowNil

`func (o *ProtectionPolicy) SetBlackoutWindowNil(b bool)`

 SetBlackoutWindowNil sets the value for BlackoutWindow to be an explicit nil

### UnsetBlackoutWindow
`func (o *ProtectionPolicy) UnsetBlackoutWindow()`

UnsetBlackoutWindow ensures that no value is present for BlackoutWindow, not even an explicit nil
### GetCascadedTargetsConfig

`func (o *ProtectionPolicy) GetCascadedTargetsConfig() []CascadedTargetConfiguration`

GetCascadedTargetsConfig returns the CascadedTargetsConfig field if non-nil, zero value otherwise.

### GetCascadedTargetsConfigOk

`func (o *ProtectionPolicy) GetCascadedTargetsConfigOk() (*[]CascadedTargetConfiguration, bool)`

GetCascadedTargetsConfigOk returns a tuple with the CascadedTargetsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadedTargetsConfig

`func (o *ProtectionPolicy) SetCascadedTargetsConfig(v []CascadedTargetConfiguration)`

SetCascadedTargetsConfig sets CascadedTargetsConfig field to given value.

### HasCascadedTargetsConfig

`func (o *ProtectionPolicy) HasCascadedTargetsConfig() bool`

HasCascadedTargetsConfig returns a boolean if a field has been set.

### GetDataLock

`func (o *ProtectionPolicy) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *ProtectionPolicy) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *ProtectionPolicy) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *ProtectionPolicy) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *ProtectionPolicy) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *ProtectionPolicy) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
### GetDescription

`func (o *ProtectionPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtectionPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtectionPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProtectionPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProtectionPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProtectionPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedRetention

`func (o *ProtectionPolicy) GetExtendedRetention() []ExtendedRetentionPolicy`

GetExtendedRetention returns the ExtendedRetention field if non-nil, zero value otherwise.

### GetExtendedRetentionOk

`func (o *ProtectionPolicy) GetExtendedRetentionOk() (*[]ExtendedRetentionPolicy, bool)`

GetExtendedRetentionOk returns a tuple with the ExtendedRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetention

`func (o *ProtectionPolicy) SetExtendedRetention(v []ExtendedRetentionPolicy)`

SetExtendedRetention sets ExtendedRetention field to given value.

### HasExtendedRetention

`func (o *ProtectionPolicy) HasExtendedRetention() bool`

HasExtendedRetention returns a boolean if a field has been set.

### SetExtendedRetentionNil

`func (o *ProtectionPolicy) SetExtendedRetentionNil(b bool)`

 SetExtendedRetentionNil sets the value for ExtendedRetention to be an explicit nil

### UnsetExtendedRetention
`func (o *ProtectionPolicy) UnsetExtendedRetention()`

UnsetExtendedRetention ensures that no value is present for ExtendedRetention, not even an explicit nil
### GetIsCBSEnabled

`func (o *ProtectionPolicy) GetIsCBSEnabled() bool`

GetIsCBSEnabled returns the IsCBSEnabled field if non-nil, zero value otherwise.

### GetIsCBSEnabledOk

`func (o *ProtectionPolicy) GetIsCBSEnabledOk() (*bool, bool)`

GetIsCBSEnabledOk returns a tuple with the IsCBSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCBSEnabled

`func (o *ProtectionPolicy) SetIsCBSEnabled(v bool)`

SetIsCBSEnabled sets IsCBSEnabled field to given value.

### HasIsCBSEnabled

`func (o *ProtectionPolicy) HasIsCBSEnabled() bool`

HasIsCBSEnabled returns a boolean if a field has been set.

### SetIsCBSEnabledNil

`func (o *ProtectionPolicy) SetIsCBSEnabledNil(b bool)`

 SetIsCBSEnabledNil sets the value for IsCBSEnabled to be an explicit nil

### UnsetIsCBSEnabled
`func (o *ProtectionPolicy) UnsetIsCBSEnabled()`

UnsetIsCBSEnabled ensures that no value is present for IsCBSEnabled, not even an explicit nil
### GetLastModificationTimeUsecs

`func (o *ProtectionPolicy) GetLastModificationTimeUsecs() int64`

GetLastModificationTimeUsecs returns the LastModificationTimeUsecs field if non-nil, zero value otherwise.

### GetLastModificationTimeUsecsOk

`func (o *ProtectionPolicy) GetLastModificationTimeUsecsOk() (*int64, bool)`

GetLastModificationTimeUsecsOk returns a tuple with the LastModificationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTimeUsecs

`func (o *ProtectionPolicy) SetLastModificationTimeUsecs(v int64)`

SetLastModificationTimeUsecs sets LastModificationTimeUsecs field to given value.

### HasLastModificationTimeUsecs

`func (o *ProtectionPolicy) HasLastModificationTimeUsecs() bool`

HasLastModificationTimeUsecs returns a boolean if a field has been set.

### SetLastModificationTimeUsecsNil

`func (o *ProtectionPolicy) SetLastModificationTimeUsecsNil(b bool)`

 SetLastModificationTimeUsecsNil sets the value for LastModificationTimeUsecs to be an explicit nil

### UnsetLastModificationTimeUsecs
`func (o *ProtectionPolicy) UnsetLastModificationTimeUsecs()`

UnsetLastModificationTimeUsecs ensures that no value is present for LastModificationTimeUsecs, not even an explicit nil
### GetName

`func (o *ProtectionPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionPolicy) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProtectionPolicy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectionPolicy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRemoteTargetPolicy

`func (o *ProtectionPolicy) GetRemoteTargetPolicy() TargetsConfiguration`

GetRemoteTargetPolicy returns the RemoteTargetPolicy field if non-nil, zero value otherwise.

### GetRemoteTargetPolicyOk

`func (o *ProtectionPolicy) GetRemoteTargetPolicyOk() (*TargetsConfiguration, bool)`

GetRemoteTargetPolicyOk returns a tuple with the RemoteTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargetPolicy

`func (o *ProtectionPolicy) SetRemoteTargetPolicy(v TargetsConfiguration)`

SetRemoteTargetPolicy sets RemoteTargetPolicy field to given value.

### HasRemoteTargetPolicy

`func (o *ProtectionPolicy) HasRemoteTargetPolicy() bool`

HasRemoteTargetPolicy returns a boolean if a field has been set.

### GetRetryOptions

`func (o *ProtectionPolicy) GetRetryOptions() RetryOptions`

GetRetryOptions returns the RetryOptions field if non-nil, zero value otherwise.

### GetRetryOptionsOk

`func (o *ProtectionPolicy) GetRetryOptionsOk() (*RetryOptions, bool)`

GetRetryOptionsOk returns a tuple with the RetryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOptions

`func (o *ProtectionPolicy) SetRetryOptions(v RetryOptions)`

SetRetryOptions sets RetryOptions field to given value.

### HasRetryOptions

`func (o *ProtectionPolicy) HasRetryOptions() bool`

HasRetryOptions returns a boolean if a field has been set.

### GetRpoPolicySettings

`func (o *ProtectionPolicy) GetRpoPolicySettings() RpoPolicySettings`

GetRpoPolicySettings returns the RpoPolicySettings field if non-nil, zero value otherwise.

### GetRpoPolicySettingsOk

`func (o *ProtectionPolicy) GetRpoPolicySettingsOk() (*RpoPolicySettings, bool)`

GetRpoPolicySettingsOk returns a tuple with the RpoPolicySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoPolicySettings

`func (o *ProtectionPolicy) SetRpoPolicySettings(v RpoPolicySettings)`

SetRpoPolicySettings sets RpoPolicySettings field to given value.

### HasRpoPolicySettings

`func (o *ProtectionPolicy) HasRpoPolicySettings() bool`

HasRpoPolicySettings returns a boolean if a field has been set.

### GetSkipIntervalMins

`func (o *ProtectionPolicy) GetSkipIntervalMins() int32`

GetSkipIntervalMins returns the SkipIntervalMins field if non-nil, zero value otherwise.

### GetSkipIntervalMinsOk

`func (o *ProtectionPolicy) GetSkipIntervalMinsOk() (*int32, bool)`

GetSkipIntervalMinsOk returns a tuple with the SkipIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIntervalMins

`func (o *ProtectionPolicy) SetSkipIntervalMins(v int32)`

SetSkipIntervalMins sets SkipIntervalMins field to given value.

### HasSkipIntervalMins

`func (o *ProtectionPolicy) HasSkipIntervalMins() bool`

HasSkipIntervalMins returns a boolean if a field has been set.

### SetSkipIntervalMinsNil

`func (o *ProtectionPolicy) SetSkipIntervalMinsNil(b bool)`

 SetSkipIntervalMinsNil sets the value for SkipIntervalMins to be an explicit nil

### UnsetSkipIntervalMins
`func (o *ProtectionPolicy) UnsetSkipIntervalMins()`

UnsetSkipIntervalMins ensures that no value is present for SkipIntervalMins, not even an explicit nil
### GetVersion

`func (o *ProtectionPolicy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProtectionPolicy) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProtectionPolicy) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProtectionPolicy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ProtectionPolicy) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ProtectionPolicy) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


