# HeliosPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupPolicy** | Pointer to [**HeliosBackupPolicy**](HeliosBackupPolicy.md) |  | [optional] 
**BlackoutWindow** | Pointer to [**[]HeliosBlackoutWindow**](HeliosBlackoutWindow.md) | List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod. | [optional] 
**ClusterIdentifier** | Pointer to **NullableString** | Specifies the cluster to which this policy belongs. This required is only for type OnPremPolicy. The format is clusterId:clusterIncarnationId. | [optional] 
**DataLock** | Pointer to **NullableString** | This field is now deprecated. Please use the DataLockConfig in the backup retention. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the Protection Policy. | [optional] 
**ExtendedRetention** | Pointer to [**[]HeliosExtendedRetentionPolicy**](HeliosExtendedRetentionPolicy.md) | Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it. | [optional] 
**IsCBSEnabled** | Pointer to **NullableBool** | Specifies true if Calender Based Schedule is supported by client. Default value is assumed as false for this feature. | [optional] 
**Name** | **NullableString** | Specifies the name of the Protection Policy. | 
**RemoteTargetPolicy** | Pointer to [**HeliosTargetsConfiguration**](HeliosTargetsConfiguration.md) |  | [optional] 
**RetryOptions** | Pointer to [**HeliosRetryOptions**](HeliosRetryOptions.md) |  | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the tenants which have access to this object. | [optional] [readonly] 
**Type** | **NullableString** | Specifies the type of the Protection Policy to be created on Helios. | 
**Version** | Pointer to **NullableInt32** | Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases. | [optional] 

## Methods

### NewHeliosPolicyRequest

`func NewHeliosPolicyRequest(name NullableString, type_ NullableString, ) *HeliosPolicyRequest`

NewHeliosPolicyRequest instantiates a new HeliosPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosPolicyRequestWithDefaults

`func NewHeliosPolicyRequestWithDefaults() *HeliosPolicyRequest`

NewHeliosPolicyRequestWithDefaults instantiates a new HeliosPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicy

`func (o *HeliosPolicyRequest) GetBackupPolicy() HeliosBackupPolicy`

GetBackupPolicy returns the BackupPolicy field if non-nil, zero value otherwise.

### GetBackupPolicyOk

`func (o *HeliosPolicyRequest) GetBackupPolicyOk() (*HeliosBackupPolicy, bool)`

GetBackupPolicyOk returns a tuple with the BackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicy

`func (o *HeliosPolicyRequest) SetBackupPolicy(v HeliosBackupPolicy)`

SetBackupPolicy sets BackupPolicy field to given value.

### HasBackupPolicy

`func (o *HeliosPolicyRequest) HasBackupPolicy() bool`

HasBackupPolicy returns a boolean if a field has been set.

### GetBlackoutWindow

`func (o *HeliosPolicyRequest) GetBlackoutWindow() []HeliosBlackoutWindow`

GetBlackoutWindow returns the BlackoutWindow field if non-nil, zero value otherwise.

### GetBlackoutWindowOk

`func (o *HeliosPolicyRequest) GetBlackoutWindowOk() (*[]HeliosBlackoutWindow, bool)`

GetBlackoutWindowOk returns a tuple with the BlackoutWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutWindow

`func (o *HeliosPolicyRequest) SetBlackoutWindow(v []HeliosBlackoutWindow)`

SetBlackoutWindow sets BlackoutWindow field to given value.

### HasBlackoutWindow

`func (o *HeliosPolicyRequest) HasBlackoutWindow() bool`

HasBlackoutWindow returns a boolean if a field has been set.

### SetBlackoutWindowNil

`func (o *HeliosPolicyRequest) SetBlackoutWindowNil(b bool)`

 SetBlackoutWindowNil sets the value for BlackoutWindow to be an explicit nil

### UnsetBlackoutWindow
`func (o *HeliosPolicyRequest) UnsetBlackoutWindow()`

UnsetBlackoutWindow ensures that no value is present for BlackoutWindow, not even an explicit nil
### GetClusterIdentifier

`func (o *HeliosPolicyRequest) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosPolicyRequest) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosPolicyRequest) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosPolicyRequest) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosPolicyRequest) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosPolicyRequest) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetDataLock

`func (o *HeliosPolicyRequest) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *HeliosPolicyRequest) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *HeliosPolicyRequest) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *HeliosPolicyRequest) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *HeliosPolicyRequest) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *HeliosPolicyRequest) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
### GetDescription

`func (o *HeliosPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HeliosPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HeliosPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HeliosPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HeliosPolicyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HeliosPolicyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedRetention

`func (o *HeliosPolicyRequest) GetExtendedRetention() []HeliosExtendedRetentionPolicy`

GetExtendedRetention returns the ExtendedRetention field if non-nil, zero value otherwise.

### GetExtendedRetentionOk

`func (o *HeliosPolicyRequest) GetExtendedRetentionOk() (*[]HeliosExtendedRetentionPolicy, bool)`

GetExtendedRetentionOk returns a tuple with the ExtendedRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetention

`func (o *HeliosPolicyRequest) SetExtendedRetention(v []HeliosExtendedRetentionPolicy)`

SetExtendedRetention sets ExtendedRetention field to given value.

### HasExtendedRetention

`func (o *HeliosPolicyRequest) HasExtendedRetention() bool`

HasExtendedRetention returns a boolean if a field has been set.

### SetExtendedRetentionNil

`func (o *HeliosPolicyRequest) SetExtendedRetentionNil(b bool)`

 SetExtendedRetentionNil sets the value for ExtendedRetention to be an explicit nil

### UnsetExtendedRetention
`func (o *HeliosPolicyRequest) UnsetExtendedRetention()`

UnsetExtendedRetention ensures that no value is present for ExtendedRetention, not even an explicit nil
### GetIsCBSEnabled

`func (o *HeliosPolicyRequest) GetIsCBSEnabled() bool`

GetIsCBSEnabled returns the IsCBSEnabled field if non-nil, zero value otherwise.

### GetIsCBSEnabledOk

`func (o *HeliosPolicyRequest) GetIsCBSEnabledOk() (*bool, bool)`

GetIsCBSEnabledOk returns a tuple with the IsCBSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCBSEnabled

`func (o *HeliosPolicyRequest) SetIsCBSEnabled(v bool)`

SetIsCBSEnabled sets IsCBSEnabled field to given value.

### HasIsCBSEnabled

`func (o *HeliosPolicyRequest) HasIsCBSEnabled() bool`

HasIsCBSEnabled returns a boolean if a field has been set.

### SetIsCBSEnabledNil

`func (o *HeliosPolicyRequest) SetIsCBSEnabledNil(b bool)`

 SetIsCBSEnabledNil sets the value for IsCBSEnabled to be an explicit nil

### UnsetIsCBSEnabled
`func (o *HeliosPolicyRequest) UnsetIsCBSEnabled()`

UnsetIsCBSEnabled ensures that no value is present for IsCBSEnabled, not even an explicit nil
### GetName

`func (o *HeliosPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *HeliosPolicyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosPolicyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRemoteTargetPolicy

`func (o *HeliosPolicyRequest) GetRemoteTargetPolicy() HeliosTargetsConfiguration`

GetRemoteTargetPolicy returns the RemoteTargetPolicy field if non-nil, zero value otherwise.

### GetRemoteTargetPolicyOk

`func (o *HeliosPolicyRequest) GetRemoteTargetPolicyOk() (*HeliosTargetsConfiguration, bool)`

GetRemoteTargetPolicyOk returns a tuple with the RemoteTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargetPolicy

`func (o *HeliosPolicyRequest) SetRemoteTargetPolicy(v HeliosTargetsConfiguration)`

SetRemoteTargetPolicy sets RemoteTargetPolicy field to given value.

### HasRemoteTargetPolicy

`func (o *HeliosPolicyRequest) HasRemoteTargetPolicy() bool`

HasRemoteTargetPolicy returns a boolean if a field has been set.

### GetRetryOptions

`func (o *HeliosPolicyRequest) GetRetryOptions() HeliosRetryOptions`

GetRetryOptions returns the RetryOptions field if non-nil, zero value otherwise.

### GetRetryOptionsOk

`func (o *HeliosPolicyRequest) GetRetryOptionsOk() (*HeliosRetryOptions, bool)`

GetRetryOptionsOk returns a tuple with the RetryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOptions

`func (o *HeliosPolicyRequest) SetRetryOptions(v HeliosRetryOptions)`

SetRetryOptions sets RetryOptions field to given value.

### HasRetryOptions

`func (o *HeliosPolicyRequest) HasRetryOptions() bool`

HasRetryOptions returns a boolean if a field has been set.

### GetTenantIds

`func (o *HeliosPolicyRequest) GetTenantIds() []*string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *HeliosPolicyRequest) GetTenantIdsOk() (*[]*string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *HeliosPolicyRequest) SetTenantIds(v []*string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *HeliosPolicyRequest) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetType

`func (o *HeliosPolicyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeliosPolicyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeliosPolicyRequest) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *HeliosPolicyRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HeliosPolicyRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVersion

`func (o *HeliosPolicyRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HeliosPolicyRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HeliosPolicyRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HeliosPolicyRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HeliosPolicyRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HeliosPolicyRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


