# HeliosProtectionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the Protection Policy. | 
**Type** | **NullableString** | Specifies the type of the Protection Policy to be created on Helios. | 
**ClusterIdentifier** | Pointer to **NullableString** | Specifies the cluster to which this policy belongs. This required is only for type OnPremPolicy. The format is clusterId:clusterIncarnationId. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the tenants which have access to this object. | [optional] [readonly] 
**BackupPolicy** | Pointer to [**HeliosBackupPolicy**](HeliosBackupPolicy.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the Protection Policy. | [optional] 
**BlackoutWindow** | Pointer to [**[]HeliosBlackoutWindow**](HeliosBlackoutWindow.md) | List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod. | [optional] 
**ExtendedRetention** | Pointer to [**[]HeliosExtendedRetentionPolicy**](HeliosExtendedRetentionPolicy.md) | Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it. | [optional] 
**RemoteTargetPolicy** | Pointer to [**HeliosTargetsConfiguration**](HeliosTargetsConfiguration.md) |  | [optional] 
**RetryOptions** | Pointer to [**HeliosRetryOptions**](HeliosRetryOptions.md) |  | [optional] 
**DataLock** | Pointer to **NullableString** | This field is now deprecated. Please use the DataLockConfig in the backup retention. | [optional] 

## Methods

### NewHeliosProtectionPolicy

`func NewHeliosProtectionPolicy(name NullableString, type_ NullableString, ) *HeliosProtectionPolicy`

NewHeliosProtectionPolicy instantiates a new HeliosProtectionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosProtectionPolicyWithDefaults

`func NewHeliosProtectionPolicyWithDefaults() *HeliosProtectionPolicy`

NewHeliosProtectionPolicyWithDefaults instantiates a new HeliosProtectionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HeliosProtectionPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosProtectionPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosProtectionPolicy) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *HeliosProtectionPolicy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosProtectionPolicy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *HeliosProtectionPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeliosProtectionPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeliosProtectionPolicy) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *HeliosProtectionPolicy) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HeliosProtectionPolicy) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetClusterIdentifier

`func (o *HeliosProtectionPolicy) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosProtectionPolicy) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosProtectionPolicy) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosProtectionPolicy) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosProtectionPolicy) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosProtectionPolicy) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetTenantIds

`func (o *HeliosProtectionPolicy) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *HeliosProtectionPolicy) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *HeliosProtectionPolicy) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *HeliosProtectionPolicy) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetBackupPolicy

`func (o *HeliosProtectionPolicy) GetBackupPolicy() HeliosBackupPolicy`

GetBackupPolicy returns the BackupPolicy field if non-nil, zero value otherwise.

### GetBackupPolicyOk

`func (o *HeliosProtectionPolicy) GetBackupPolicyOk() (*HeliosBackupPolicy, bool)`

GetBackupPolicyOk returns a tuple with the BackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicy

`func (o *HeliosProtectionPolicy) SetBackupPolicy(v HeliosBackupPolicy)`

SetBackupPolicy sets BackupPolicy field to given value.

### HasBackupPolicy

`func (o *HeliosProtectionPolicy) HasBackupPolicy() bool`

HasBackupPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *HeliosProtectionPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HeliosProtectionPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HeliosProtectionPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HeliosProtectionPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HeliosProtectionPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HeliosProtectionPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBlackoutWindow

`func (o *HeliosProtectionPolicy) GetBlackoutWindow() []HeliosBlackoutWindow`

GetBlackoutWindow returns the BlackoutWindow field if non-nil, zero value otherwise.

### GetBlackoutWindowOk

`func (o *HeliosProtectionPolicy) GetBlackoutWindowOk() (*[]HeliosBlackoutWindow, bool)`

GetBlackoutWindowOk returns a tuple with the BlackoutWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutWindow

`func (o *HeliosProtectionPolicy) SetBlackoutWindow(v []HeliosBlackoutWindow)`

SetBlackoutWindow sets BlackoutWindow field to given value.

### HasBlackoutWindow

`func (o *HeliosProtectionPolicy) HasBlackoutWindow() bool`

HasBlackoutWindow returns a boolean if a field has been set.

### SetBlackoutWindowNil

`func (o *HeliosProtectionPolicy) SetBlackoutWindowNil(b bool)`

 SetBlackoutWindowNil sets the value for BlackoutWindow to be an explicit nil

### UnsetBlackoutWindow
`func (o *HeliosProtectionPolicy) UnsetBlackoutWindow()`

UnsetBlackoutWindow ensures that no value is present for BlackoutWindow, not even an explicit nil
### GetExtendedRetention

`func (o *HeliosProtectionPolicy) GetExtendedRetention() []HeliosExtendedRetentionPolicy`

GetExtendedRetention returns the ExtendedRetention field if non-nil, zero value otherwise.

### GetExtendedRetentionOk

`func (o *HeliosProtectionPolicy) GetExtendedRetentionOk() (*[]HeliosExtendedRetentionPolicy, bool)`

GetExtendedRetentionOk returns a tuple with the ExtendedRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetention

`func (o *HeliosProtectionPolicy) SetExtendedRetention(v []HeliosExtendedRetentionPolicy)`

SetExtendedRetention sets ExtendedRetention field to given value.

### HasExtendedRetention

`func (o *HeliosProtectionPolicy) HasExtendedRetention() bool`

HasExtendedRetention returns a boolean if a field has been set.

### SetExtendedRetentionNil

`func (o *HeliosProtectionPolicy) SetExtendedRetentionNil(b bool)`

 SetExtendedRetentionNil sets the value for ExtendedRetention to be an explicit nil

### UnsetExtendedRetention
`func (o *HeliosProtectionPolicy) UnsetExtendedRetention()`

UnsetExtendedRetention ensures that no value is present for ExtendedRetention, not even an explicit nil
### GetRemoteTargetPolicy

`func (o *HeliosProtectionPolicy) GetRemoteTargetPolicy() HeliosTargetsConfiguration`

GetRemoteTargetPolicy returns the RemoteTargetPolicy field if non-nil, zero value otherwise.

### GetRemoteTargetPolicyOk

`func (o *HeliosProtectionPolicy) GetRemoteTargetPolicyOk() (*HeliosTargetsConfiguration, bool)`

GetRemoteTargetPolicyOk returns a tuple with the RemoteTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargetPolicy

`func (o *HeliosProtectionPolicy) SetRemoteTargetPolicy(v HeliosTargetsConfiguration)`

SetRemoteTargetPolicy sets RemoteTargetPolicy field to given value.

### HasRemoteTargetPolicy

`func (o *HeliosProtectionPolicy) HasRemoteTargetPolicy() bool`

HasRemoteTargetPolicy returns a boolean if a field has been set.

### GetRetryOptions

`func (o *HeliosProtectionPolicy) GetRetryOptions() HeliosRetryOptions`

GetRetryOptions returns the RetryOptions field if non-nil, zero value otherwise.

### GetRetryOptionsOk

`func (o *HeliosProtectionPolicy) GetRetryOptionsOk() (*HeliosRetryOptions, bool)`

GetRetryOptionsOk returns a tuple with the RetryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOptions

`func (o *HeliosProtectionPolicy) SetRetryOptions(v HeliosRetryOptions)`

SetRetryOptions sets RetryOptions field to given value.

### HasRetryOptions

`func (o *HeliosProtectionPolicy) HasRetryOptions() bool`

HasRetryOptions returns a boolean if a field has been set.

### GetDataLock

`func (o *HeliosProtectionPolicy) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *HeliosProtectionPolicy) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *HeliosProtectionPolicy) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *HeliosProtectionPolicy) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *HeliosProtectionPolicy) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *HeliosProtectionPolicy) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


