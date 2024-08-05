# PolicyTemplateResponse

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
**IsUsable** | Pointer to **NullableBool** | This field is set to true if this policy template qualifies to create more policies. If the template is partially filled and can not create a working policy then this field will be set to false. | [optional] 
**NumLinkedPolicies** | Pointer to **NullableInt64** | Specifies the number of policies linked to this policy template. Only applicable in case of policy template. | [optional] 

## Methods

### NewPolicyTemplateResponse

`func NewPolicyTemplateResponse(backupPolicy BackupPolicy, name NullableString, ) *PolicyTemplateResponse`

NewPolicyTemplateResponse instantiates a new PolicyTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateResponseWithDefaults

`func NewPolicyTemplateResponseWithDefaults() *PolicyTemplateResponse`

NewPolicyTemplateResponseWithDefaults instantiates a new PolicyTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicy

`func (o *PolicyTemplateResponse) GetBackupPolicy() BackupPolicy`

GetBackupPolicy returns the BackupPolicy field if non-nil, zero value otherwise.

### GetBackupPolicyOk

`func (o *PolicyTemplateResponse) GetBackupPolicyOk() (*BackupPolicy, bool)`

GetBackupPolicyOk returns a tuple with the BackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicy

`func (o *PolicyTemplateResponse) SetBackupPolicy(v BackupPolicy)`

SetBackupPolicy sets BackupPolicy field to given value.


### GetBlackoutWindow

`func (o *PolicyTemplateResponse) GetBlackoutWindow() []BlackoutWindow`

GetBlackoutWindow returns the BlackoutWindow field if non-nil, zero value otherwise.

### GetBlackoutWindowOk

`func (o *PolicyTemplateResponse) GetBlackoutWindowOk() (*[]BlackoutWindow, bool)`

GetBlackoutWindowOk returns a tuple with the BlackoutWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutWindow

`func (o *PolicyTemplateResponse) SetBlackoutWindow(v []BlackoutWindow)`

SetBlackoutWindow sets BlackoutWindow field to given value.

### HasBlackoutWindow

`func (o *PolicyTemplateResponse) HasBlackoutWindow() bool`

HasBlackoutWindow returns a boolean if a field has been set.

### SetBlackoutWindowNil

`func (o *PolicyTemplateResponse) SetBlackoutWindowNil(b bool)`

 SetBlackoutWindowNil sets the value for BlackoutWindow to be an explicit nil

### UnsetBlackoutWindow
`func (o *PolicyTemplateResponse) UnsetBlackoutWindow()`

UnsetBlackoutWindow ensures that no value is present for BlackoutWindow, not even an explicit nil
### GetCascadedTargetsConfig

`func (o *PolicyTemplateResponse) GetCascadedTargetsConfig() []CascadedTargetConfiguration`

GetCascadedTargetsConfig returns the CascadedTargetsConfig field if non-nil, zero value otherwise.

### GetCascadedTargetsConfigOk

`func (o *PolicyTemplateResponse) GetCascadedTargetsConfigOk() (*[]CascadedTargetConfiguration, bool)`

GetCascadedTargetsConfigOk returns a tuple with the CascadedTargetsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascadedTargetsConfig

`func (o *PolicyTemplateResponse) SetCascadedTargetsConfig(v []CascadedTargetConfiguration)`

SetCascadedTargetsConfig sets CascadedTargetsConfig field to given value.

### HasCascadedTargetsConfig

`func (o *PolicyTemplateResponse) HasCascadedTargetsConfig() bool`

HasCascadedTargetsConfig returns a boolean if a field has been set.

### GetDataLock

`func (o *PolicyTemplateResponse) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *PolicyTemplateResponse) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *PolicyTemplateResponse) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *PolicyTemplateResponse) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *PolicyTemplateResponse) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *PolicyTemplateResponse) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
### GetDescription

`func (o *PolicyTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyTemplateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyTemplateResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyTemplateResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedRetention

`func (o *PolicyTemplateResponse) GetExtendedRetention() []ExtendedRetentionPolicy`

GetExtendedRetention returns the ExtendedRetention field if non-nil, zero value otherwise.

### GetExtendedRetentionOk

`func (o *PolicyTemplateResponse) GetExtendedRetentionOk() (*[]ExtendedRetentionPolicy, bool)`

GetExtendedRetentionOk returns a tuple with the ExtendedRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetention

`func (o *PolicyTemplateResponse) SetExtendedRetention(v []ExtendedRetentionPolicy)`

SetExtendedRetention sets ExtendedRetention field to given value.

### HasExtendedRetention

`func (o *PolicyTemplateResponse) HasExtendedRetention() bool`

HasExtendedRetention returns a boolean if a field has been set.

### SetExtendedRetentionNil

`func (o *PolicyTemplateResponse) SetExtendedRetentionNil(b bool)`

 SetExtendedRetentionNil sets the value for ExtendedRetention to be an explicit nil

### UnsetExtendedRetention
`func (o *PolicyTemplateResponse) UnsetExtendedRetention()`

UnsetExtendedRetention ensures that no value is present for ExtendedRetention, not even an explicit nil
### GetIsCBSEnabled

`func (o *PolicyTemplateResponse) GetIsCBSEnabled() bool`

GetIsCBSEnabled returns the IsCBSEnabled field if non-nil, zero value otherwise.

### GetIsCBSEnabledOk

`func (o *PolicyTemplateResponse) GetIsCBSEnabledOk() (*bool, bool)`

GetIsCBSEnabledOk returns a tuple with the IsCBSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCBSEnabled

`func (o *PolicyTemplateResponse) SetIsCBSEnabled(v bool)`

SetIsCBSEnabled sets IsCBSEnabled field to given value.

### HasIsCBSEnabled

`func (o *PolicyTemplateResponse) HasIsCBSEnabled() bool`

HasIsCBSEnabled returns a boolean if a field has been set.

### SetIsCBSEnabledNil

`func (o *PolicyTemplateResponse) SetIsCBSEnabledNil(b bool)`

 SetIsCBSEnabledNil sets the value for IsCBSEnabled to be an explicit nil

### UnsetIsCBSEnabled
`func (o *PolicyTemplateResponse) UnsetIsCBSEnabled()`

UnsetIsCBSEnabled ensures that no value is present for IsCBSEnabled, not even an explicit nil
### GetLastModificationTimeUsecs

`func (o *PolicyTemplateResponse) GetLastModificationTimeUsecs() int64`

GetLastModificationTimeUsecs returns the LastModificationTimeUsecs field if non-nil, zero value otherwise.

### GetLastModificationTimeUsecsOk

`func (o *PolicyTemplateResponse) GetLastModificationTimeUsecsOk() (*int64, bool)`

GetLastModificationTimeUsecsOk returns a tuple with the LastModificationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTimeUsecs

`func (o *PolicyTemplateResponse) SetLastModificationTimeUsecs(v int64)`

SetLastModificationTimeUsecs sets LastModificationTimeUsecs field to given value.

### HasLastModificationTimeUsecs

`func (o *PolicyTemplateResponse) HasLastModificationTimeUsecs() bool`

HasLastModificationTimeUsecs returns a boolean if a field has been set.

### SetLastModificationTimeUsecsNil

`func (o *PolicyTemplateResponse) SetLastModificationTimeUsecsNil(b bool)`

 SetLastModificationTimeUsecsNil sets the value for LastModificationTimeUsecs to be an explicit nil

### UnsetLastModificationTimeUsecs
`func (o *PolicyTemplateResponse) UnsetLastModificationTimeUsecs()`

UnsetLastModificationTimeUsecs ensures that no value is present for LastModificationTimeUsecs, not even an explicit nil
### GetName

`func (o *PolicyTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PolicyTemplateResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PolicyTemplateResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRemoteTargetPolicy

`func (o *PolicyTemplateResponse) GetRemoteTargetPolicy() TargetsConfiguration`

GetRemoteTargetPolicy returns the RemoteTargetPolicy field if non-nil, zero value otherwise.

### GetRemoteTargetPolicyOk

`func (o *PolicyTemplateResponse) GetRemoteTargetPolicyOk() (*TargetsConfiguration, bool)`

GetRemoteTargetPolicyOk returns a tuple with the RemoteTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargetPolicy

`func (o *PolicyTemplateResponse) SetRemoteTargetPolicy(v TargetsConfiguration)`

SetRemoteTargetPolicy sets RemoteTargetPolicy field to given value.

### HasRemoteTargetPolicy

`func (o *PolicyTemplateResponse) HasRemoteTargetPolicy() bool`

HasRemoteTargetPolicy returns a boolean if a field has been set.

### GetRetryOptions

`func (o *PolicyTemplateResponse) GetRetryOptions() RetryOptions`

GetRetryOptions returns the RetryOptions field if non-nil, zero value otherwise.

### GetRetryOptionsOk

`func (o *PolicyTemplateResponse) GetRetryOptionsOk() (*RetryOptions, bool)`

GetRetryOptionsOk returns a tuple with the RetryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOptions

`func (o *PolicyTemplateResponse) SetRetryOptions(v RetryOptions)`

SetRetryOptions sets RetryOptions field to given value.

### HasRetryOptions

`func (o *PolicyTemplateResponse) HasRetryOptions() bool`

HasRetryOptions returns a boolean if a field has been set.

### GetRpoPolicySettings

`func (o *PolicyTemplateResponse) GetRpoPolicySettings() RpoPolicySettings`

GetRpoPolicySettings returns the RpoPolicySettings field if non-nil, zero value otherwise.

### GetRpoPolicySettingsOk

`func (o *PolicyTemplateResponse) GetRpoPolicySettingsOk() (*RpoPolicySettings, bool)`

GetRpoPolicySettingsOk returns a tuple with the RpoPolicySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoPolicySettings

`func (o *PolicyTemplateResponse) SetRpoPolicySettings(v RpoPolicySettings)`

SetRpoPolicySettings sets RpoPolicySettings field to given value.

### HasRpoPolicySettings

`func (o *PolicyTemplateResponse) HasRpoPolicySettings() bool`

HasRpoPolicySettings returns a boolean if a field has been set.

### GetSkipIntervalMins

`func (o *PolicyTemplateResponse) GetSkipIntervalMins() int32`

GetSkipIntervalMins returns the SkipIntervalMins field if non-nil, zero value otherwise.

### GetSkipIntervalMinsOk

`func (o *PolicyTemplateResponse) GetSkipIntervalMinsOk() (*int32, bool)`

GetSkipIntervalMinsOk returns a tuple with the SkipIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIntervalMins

`func (o *PolicyTemplateResponse) SetSkipIntervalMins(v int32)`

SetSkipIntervalMins sets SkipIntervalMins field to given value.

### HasSkipIntervalMins

`func (o *PolicyTemplateResponse) HasSkipIntervalMins() bool`

HasSkipIntervalMins returns a boolean if a field has been set.

### SetSkipIntervalMinsNil

`func (o *PolicyTemplateResponse) SetSkipIntervalMinsNil(b bool)`

 SetSkipIntervalMinsNil sets the value for SkipIntervalMins to be an explicit nil

### UnsetSkipIntervalMins
`func (o *PolicyTemplateResponse) UnsetSkipIntervalMins()`

UnsetSkipIntervalMins ensures that no value is present for SkipIntervalMins, not even an explicit nil
### GetVersion

`func (o *PolicyTemplateResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PolicyTemplateResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PolicyTemplateResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PolicyTemplateResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PolicyTemplateResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PolicyTemplateResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetId

`func (o *PolicyTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyTemplateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyTemplateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PolicyTemplateResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PolicyTemplateResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsUsable

`func (o *PolicyTemplateResponse) GetIsUsable() bool`

GetIsUsable returns the IsUsable field if non-nil, zero value otherwise.

### GetIsUsableOk

`func (o *PolicyTemplateResponse) GetIsUsableOk() (*bool, bool)`

GetIsUsableOk returns a tuple with the IsUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsable

`func (o *PolicyTemplateResponse) SetIsUsable(v bool)`

SetIsUsable sets IsUsable field to given value.

### HasIsUsable

`func (o *PolicyTemplateResponse) HasIsUsable() bool`

HasIsUsable returns a boolean if a field has been set.

### SetIsUsableNil

`func (o *PolicyTemplateResponse) SetIsUsableNil(b bool)`

 SetIsUsableNil sets the value for IsUsable to be an explicit nil

### UnsetIsUsable
`func (o *PolicyTemplateResponse) UnsetIsUsable()`

UnsetIsUsable ensures that no value is present for IsUsable, not even an explicit nil
### GetNumLinkedPolicies

`func (o *PolicyTemplateResponse) GetNumLinkedPolicies() int64`

GetNumLinkedPolicies returns the NumLinkedPolicies field if non-nil, zero value otherwise.

### GetNumLinkedPoliciesOk

`func (o *PolicyTemplateResponse) GetNumLinkedPoliciesOk() (*int64, bool)`

GetNumLinkedPoliciesOk returns a tuple with the NumLinkedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLinkedPolicies

`func (o *PolicyTemplateResponse) SetNumLinkedPolicies(v int64)`

SetNumLinkedPolicies sets NumLinkedPolicies field to given value.

### HasNumLinkedPolicies

`func (o *PolicyTemplateResponse) HasNumLinkedPolicies() bool`

HasNumLinkedPolicies returns a boolean if a field has been set.

### SetNumLinkedPoliciesNil

`func (o *PolicyTemplateResponse) SetNumLinkedPoliciesNil(b bool)`

 SetNumLinkedPoliciesNil sets the value for NumLinkedPolicies to be an explicit nil

### UnsetNumLinkedPolicies
`func (o *PolicyTemplateResponse) UnsetNumLinkedPolicies()`

UnsetNumLinkedPolicies ensures that no value is present for NumLinkedPolicies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


