# ProtectionPolicyResponse

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
**Id** | Pointer to **NullableString** | Specifies a unique Policy id assigned by the Cohesity Cluster. | [optional] 
**IsReplicated** | Pointer to **NullableBool** | This field is set to true when policy is the replicated policy. | [optional] 
**IsUsable** | Pointer to **NullableBool** | This field is set to true if the linked policy which is internally created from a policy templates qualifies as usable to create more policies on the cluster. If the linked policy is partially filled and can not create a working policy then this field will be set to false. In case of normal policy created on the cluster, this field wont be populated. | [optional] 
**NumProtectedObjects** | Pointer to **NullableInt64** | Specifies the number of protected objects using the protection policy. | [optional] 
**NumProtectionGroups** | Pointer to **NullableInt64** | Specifies the number of protection groups using the protection policy. | [optional] 
**TemplateId** | Pointer to **NullableString** | Specifies the parent policy template id to which the policy is linked to. This field is set only when policy is created from template. | [optional] 

## Methods

### NewProtectionPolicyResponse

`func NewProtectionPolicyResponse(backupPolicy BackupPolicy, name NullableString, ) *ProtectionPolicyResponse`

NewProtectionPolicyResponse instantiates a new ProtectionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionPolicyResponseWithDefaults

`func NewProtectionPolicyResponseWithDefaults() *ProtectionPolicyResponse`

NewProtectionPolicyResponseWithDefaults instantiates a new ProtectionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicy

`func (o *ProtectionPolicyResponse) GetBackupPolicy() BackupPolicy`

GetBackupPolicy returns the BackupPolicy field if non-nil, zero value otherwise.

### GetBackupPolicyOk

`func (o *ProtectionPolicyResponse) GetBackupPolicyOk() (*BackupPolicy, bool)`

GetBackupPolicyOk returns a tuple with the BackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicy

`func (o *ProtectionPolicyResponse) SetBackupPolicy(v BackupPolicy)`

SetBackupPolicy sets BackupPolicy field to given value.


### GetBlackoutWindow

`func (o *ProtectionPolicyResponse) GetBlackoutWindow() []BlackoutWindow`

GetBlackoutWindow returns the BlackoutWindow field if non-nil, zero value otherwise.

### GetBlackoutWindowOk

`func (o *ProtectionPolicyResponse) GetBlackoutWindowOk() (*[]BlackoutWindow, bool)`

GetBlackoutWindowOk returns a tuple with the BlackoutWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutWindow

`func (o *ProtectionPolicyResponse) SetBlackoutWindow(v []BlackoutWindow)`

SetBlackoutWindow sets BlackoutWindow field to given value.

### HasBlackoutWindow

`func (o *ProtectionPolicyResponse) HasBlackoutWindow() bool`

HasBlackoutWindow returns a boolean if a field has been set.

### SetBlackoutWindowNil

`func (o *ProtectionPolicyResponse) SetBlackoutWindowNil(b bool)`

 SetBlackoutWindowNil sets the value for BlackoutWindow to be an explicit nil

### UnsetBlackoutWindow
`func (o *ProtectionPolicyResponse) UnsetBlackoutWindow()`

UnsetBlackoutWindow ensures that no value is present for BlackoutWindow, not even an explicit nil
### GetCascadedTargetsConfig

`func (o *ProtectionPolicyResponse) GetCascadedTargetsConfig() []CascadedTargetConfiguration`

GetCascadedTargetsConfig returns the CascadedTargetsConfig field if non-nil, zero value otherwise.

### GetCascadedTargetsConfigOk

`func (o *ProtectionPolicyResponse) GetCascadedTargetsConfigOk() (*[]CascadedTargetConfiguration, bool)`

GetCascadedTargetsConfigOk returns a tuple with the CascadedTargetsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadedTargetsConfig

`func (o *ProtectionPolicyResponse) SetCascadedTargetsConfig(v []CascadedTargetConfiguration)`

SetCascadedTargetsConfig sets CascadedTargetsConfig field to given value.

### HasCascadedTargetsConfig

`func (o *ProtectionPolicyResponse) HasCascadedTargetsConfig() bool`

HasCascadedTargetsConfig returns a boolean if a field has been set.

### GetDataLock

`func (o *ProtectionPolicyResponse) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *ProtectionPolicyResponse) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *ProtectionPolicyResponse) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *ProtectionPolicyResponse) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *ProtectionPolicyResponse) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *ProtectionPolicyResponse) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
### GetDescription

`func (o *ProtectionPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProtectionPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProtectionPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProtectionPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProtectionPolicyResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProtectionPolicyResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedRetention

`func (o *ProtectionPolicyResponse) GetExtendedRetention() []ExtendedRetentionPolicy`

GetExtendedRetention returns the ExtendedRetention field if non-nil, zero value otherwise.

### GetExtendedRetentionOk

`func (o *ProtectionPolicyResponse) GetExtendedRetentionOk() (*[]ExtendedRetentionPolicy, bool)`

GetExtendedRetentionOk returns a tuple with the ExtendedRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetention

`func (o *ProtectionPolicyResponse) SetExtendedRetention(v []ExtendedRetentionPolicy)`

SetExtendedRetention sets ExtendedRetention field to given value.

### HasExtendedRetention

`func (o *ProtectionPolicyResponse) HasExtendedRetention() bool`

HasExtendedRetention returns a boolean if a field has been set.

### SetExtendedRetentionNil

`func (o *ProtectionPolicyResponse) SetExtendedRetentionNil(b bool)`

 SetExtendedRetentionNil sets the value for ExtendedRetention to be an explicit nil

### UnsetExtendedRetention
`func (o *ProtectionPolicyResponse) UnsetExtendedRetention()`

UnsetExtendedRetention ensures that no value is present for ExtendedRetention, not even an explicit nil
### GetIsCBSEnabled

`func (o *ProtectionPolicyResponse) GetIsCBSEnabled() bool`

GetIsCBSEnabled returns the IsCBSEnabled field if non-nil, zero value otherwise.

### GetIsCBSEnabledOk

`func (o *ProtectionPolicyResponse) GetIsCBSEnabledOk() (*bool, bool)`

GetIsCBSEnabledOk returns a tuple with the IsCBSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCBSEnabled

`func (o *ProtectionPolicyResponse) SetIsCBSEnabled(v bool)`

SetIsCBSEnabled sets IsCBSEnabled field to given value.

### HasIsCBSEnabled

`func (o *ProtectionPolicyResponse) HasIsCBSEnabled() bool`

HasIsCBSEnabled returns a boolean if a field has been set.

### SetIsCBSEnabledNil

`func (o *ProtectionPolicyResponse) SetIsCBSEnabledNil(b bool)`

 SetIsCBSEnabledNil sets the value for IsCBSEnabled to be an explicit nil

### UnsetIsCBSEnabled
`func (o *ProtectionPolicyResponse) UnsetIsCBSEnabled()`

UnsetIsCBSEnabled ensures that no value is present for IsCBSEnabled, not even an explicit nil
### GetLastModificationTimeUsecs

`func (o *ProtectionPolicyResponse) GetLastModificationTimeUsecs() int64`

GetLastModificationTimeUsecs returns the LastModificationTimeUsecs field if non-nil, zero value otherwise.

### GetLastModificationTimeUsecsOk

`func (o *ProtectionPolicyResponse) GetLastModificationTimeUsecsOk() (*int64, bool)`

GetLastModificationTimeUsecsOk returns a tuple with the LastModificationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTimeUsecs

`func (o *ProtectionPolicyResponse) SetLastModificationTimeUsecs(v int64)`

SetLastModificationTimeUsecs sets LastModificationTimeUsecs field to given value.

### HasLastModificationTimeUsecs

`func (o *ProtectionPolicyResponse) HasLastModificationTimeUsecs() bool`

HasLastModificationTimeUsecs returns a boolean if a field has been set.

### SetLastModificationTimeUsecsNil

`func (o *ProtectionPolicyResponse) SetLastModificationTimeUsecsNil(b bool)`

 SetLastModificationTimeUsecsNil sets the value for LastModificationTimeUsecs to be an explicit nil

### UnsetLastModificationTimeUsecs
`func (o *ProtectionPolicyResponse) UnsetLastModificationTimeUsecs()`

UnsetLastModificationTimeUsecs ensures that no value is present for LastModificationTimeUsecs, not even an explicit nil
### GetName

`func (o *ProtectionPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionPolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProtectionPolicyResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectionPolicyResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRemoteTargetPolicy

`func (o *ProtectionPolicyResponse) GetRemoteTargetPolicy() TargetsConfiguration`

GetRemoteTargetPolicy returns the RemoteTargetPolicy field if non-nil, zero value otherwise.

### GetRemoteTargetPolicyOk

`func (o *ProtectionPolicyResponse) GetRemoteTargetPolicyOk() (*TargetsConfiguration, bool)`

GetRemoteTargetPolicyOk returns a tuple with the RemoteTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargetPolicy

`func (o *ProtectionPolicyResponse) SetRemoteTargetPolicy(v TargetsConfiguration)`

SetRemoteTargetPolicy sets RemoteTargetPolicy field to given value.

### HasRemoteTargetPolicy

`func (o *ProtectionPolicyResponse) HasRemoteTargetPolicy() bool`

HasRemoteTargetPolicy returns a boolean if a field has been set.

### GetRetryOptions

`func (o *ProtectionPolicyResponse) GetRetryOptions() RetryOptions`

GetRetryOptions returns the RetryOptions field if non-nil, zero value otherwise.

### GetRetryOptionsOk

`func (o *ProtectionPolicyResponse) GetRetryOptionsOk() (*RetryOptions, bool)`

GetRetryOptionsOk returns a tuple with the RetryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOptions

`func (o *ProtectionPolicyResponse) SetRetryOptions(v RetryOptions)`

SetRetryOptions sets RetryOptions field to given value.

### HasRetryOptions

`func (o *ProtectionPolicyResponse) HasRetryOptions() bool`

HasRetryOptions returns a boolean if a field has been set.

### GetRpoPolicySettings

`func (o *ProtectionPolicyResponse) GetRpoPolicySettings() RpoPolicySettings`

GetRpoPolicySettings returns the RpoPolicySettings field if non-nil, zero value otherwise.

### GetRpoPolicySettingsOk

`func (o *ProtectionPolicyResponse) GetRpoPolicySettingsOk() (*RpoPolicySettings, bool)`

GetRpoPolicySettingsOk returns a tuple with the RpoPolicySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoPolicySettings

`func (o *ProtectionPolicyResponse) SetRpoPolicySettings(v RpoPolicySettings)`

SetRpoPolicySettings sets RpoPolicySettings field to given value.

### HasRpoPolicySettings

`func (o *ProtectionPolicyResponse) HasRpoPolicySettings() bool`

HasRpoPolicySettings returns a boolean if a field has been set.

### GetSkipIntervalMins

`func (o *ProtectionPolicyResponse) GetSkipIntervalMins() int32`

GetSkipIntervalMins returns the SkipIntervalMins field if non-nil, zero value otherwise.

### GetSkipIntervalMinsOk

`func (o *ProtectionPolicyResponse) GetSkipIntervalMinsOk() (*int32, bool)`

GetSkipIntervalMinsOk returns a tuple with the SkipIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIntervalMins

`func (o *ProtectionPolicyResponse) SetSkipIntervalMins(v int32)`

SetSkipIntervalMins sets SkipIntervalMins field to given value.

### HasSkipIntervalMins

`func (o *ProtectionPolicyResponse) HasSkipIntervalMins() bool`

HasSkipIntervalMins returns a boolean if a field has been set.

### SetSkipIntervalMinsNil

`func (o *ProtectionPolicyResponse) SetSkipIntervalMinsNil(b bool)`

 SetSkipIntervalMinsNil sets the value for SkipIntervalMins to be an explicit nil

### UnsetSkipIntervalMins
`func (o *ProtectionPolicyResponse) UnsetSkipIntervalMins()`

UnsetSkipIntervalMins ensures that no value is present for SkipIntervalMins, not even an explicit nil
### GetVersion

`func (o *ProtectionPolicyResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProtectionPolicyResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProtectionPolicyResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProtectionPolicyResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ProtectionPolicyResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ProtectionPolicyResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetId

`func (o *ProtectionPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectionPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectionPolicyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectionPolicyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectionPolicyResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectionPolicyResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsReplicated

`func (o *ProtectionPolicyResponse) GetIsReplicated() bool`

GetIsReplicated returns the IsReplicated field if non-nil, zero value otherwise.

### GetIsReplicatedOk

`func (o *ProtectionPolicyResponse) GetIsReplicatedOk() (*bool, bool)`

GetIsReplicatedOk returns a tuple with the IsReplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplicated

`func (o *ProtectionPolicyResponse) SetIsReplicated(v bool)`

SetIsReplicated sets IsReplicated field to given value.

### HasIsReplicated

`func (o *ProtectionPolicyResponse) HasIsReplicated() bool`

HasIsReplicated returns a boolean if a field has been set.

### SetIsReplicatedNil

`func (o *ProtectionPolicyResponse) SetIsReplicatedNil(b bool)`

 SetIsReplicatedNil sets the value for IsReplicated to be an explicit nil

### UnsetIsReplicated
`func (o *ProtectionPolicyResponse) UnsetIsReplicated()`

UnsetIsReplicated ensures that no value is present for IsReplicated, not even an explicit nil
### GetIsUsable

`func (o *ProtectionPolicyResponse) GetIsUsable() bool`

GetIsUsable returns the IsUsable field if non-nil, zero value otherwise.

### GetIsUsableOk

`func (o *ProtectionPolicyResponse) GetIsUsableOk() (*bool, bool)`

GetIsUsableOk returns a tuple with the IsUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsable

`func (o *ProtectionPolicyResponse) SetIsUsable(v bool)`

SetIsUsable sets IsUsable field to given value.

### HasIsUsable

`func (o *ProtectionPolicyResponse) HasIsUsable() bool`

HasIsUsable returns a boolean if a field has been set.

### SetIsUsableNil

`func (o *ProtectionPolicyResponse) SetIsUsableNil(b bool)`

 SetIsUsableNil sets the value for IsUsable to be an explicit nil

### UnsetIsUsable
`func (o *ProtectionPolicyResponse) UnsetIsUsable()`

UnsetIsUsable ensures that no value is present for IsUsable, not even an explicit nil
### GetNumProtectedObjects

`func (o *ProtectionPolicyResponse) GetNumProtectedObjects() int64`

GetNumProtectedObjects returns the NumProtectedObjects field if non-nil, zero value otherwise.

### GetNumProtectedObjectsOk

`func (o *ProtectionPolicyResponse) GetNumProtectedObjectsOk() (*int64, bool)`

GetNumProtectedObjectsOk returns a tuple with the NumProtectedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProtectedObjects

`func (o *ProtectionPolicyResponse) SetNumProtectedObjects(v int64)`

SetNumProtectedObjects sets NumProtectedObjects field to given value.

### HasNumProtectedObjects

`func (o *ProtectionPolicyResponse) HasNumProtectedObjects() bool`

HasNumProtectedObjects returns a boolean if a field has been set.

### SetNumProtectedObjectsNil

`func (o *ProtectionPolicyResponse) SetNumProtectedObjectsNil(b bool)`

 SetNumProtectedObjectsNil sets the value for NumProtectedObjects to be an explicit nil

### UnsetNumProtectedObjects
`func (o *ProtectionPolicyResponse) UnsetNumProtectedObjects()`

UnsetNumProtectedObjects ensures that no value is present for NumProtectedObjects, not even an explicit nil
### GetNumProtectionGroups

`func (o *ProtectionPolicyResponse) GetNumProtectionGroups() int64`

GetNumProtectionGroups returns the NumProtectionGroups field if non-nil, zero value otherwise.

### GetNumProtectionGroupsOk

`func (o *ProtectionPolicyResponse) GetNumProtectionGroupsOk() (*int64, bool)`

GetNumProtectionGroupsOk returns a tuple with the NumProtectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumProtectionGroups

`func (o *ProtectionPolicyResponse) SetNumProtectionGroups(v int64)`

SetNumProtectionGroups sets NumProtectionGroups field to given value.

### HasNumProtectionGroups

`func (o *ProtectionPolicyResponse) HasNumProtectionGroups() bool`

HasNumProtectionGroups returns a boolean if a field has been set.

### SetNumProtectionGroupsNil

`func (o *ProtectionPolicyResponse) SetNumProtectionGroupsNil(b bool)`

 SetNumProtectionGroupsNil sets the value for NumProtectionGroups to be an explicit nil

### UnsetNumProtectionGroups
`func (o *ProtectionPolicyResponse) UnsetNumProtectionGroups()`

UnsetNumProtectionGroups ensures that no value is present for NumProtectionGroups, not even an explicit nil
### GetTemplateId

`func (o *ProtectionPolicyResponse) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ProtectionPolicyResponse) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ProtectionPolicyResponse) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ProtectionPolicyResponse) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *ProtectionPolicyResponse) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *ProtectionPolicyResponse) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


