# HeliosPolicyResponse

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
**Id** | Pointer to **NullableString** | Specifies a unique policy id assigned by the Helios. | [optional] 
**NumLinkedPolicies** | Pointer to **NullableInt64** | In case of global policy response, specifies the number of policies linked to this global policy on the cluster. | [optional] 
**NumObjectProtections** | Pointer to **NullableInt64** | Specifies the number of object protections using the protection policy. | [optional] 

## Methods

### NewHeliosPolicyResponse

`func NewHeliosPolicyResponse(name NullableString, type_ NullableString, ) *HeliosPolicyResponse`

NewHeliosPolicyResponse instantiates a new HeliosPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosPolicyResponseWithDefaults

`func NewHeliosPolicyResponseWithDefaults() *HeliosPolicyResponse`

NewHeliosPolicyResponseWithDefaults instantiates a new HeliosPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupPolicy

`func (o *HeliosPolicyResponse) GetBackupPolicy() HeliosBackupPolicy`

GetBackupPolicy returns the BackupPolicy field if non-nil, zero value otherwise.

### GetBackupPolicyOk

`func (o *HeliosPolicyResponse) GetBackupPolicyOk() (*HeliosBackupPolicy, bool)`

GetBackupPolicyOk returns a tuple with the BackupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicy

`func (o *HeliosPolicyResponse) SetBackupPolicy(v HeliosBackupPolicy)`

SetBackupPolicy sets BackupPolicy field to given value.

### HasBackupPolicy

`func (o *HeliosPolicyResponse) HasBackupPolicy() bool`

HasBackupPolicy returns a boolean if a field has been set.

### GetBlackoutWindow

`func (o *HeliosPolicyResponse) GetBlackoutWindow() []HeliosBlackoutWindow`

GetBlackoutWindow returns the BlackoutWindow field if non-nil, zero value otherwise.

### GetBlackoutWindowOk

`func (o *HeliosPolicyResponse) GetBlackoutWindowOk() (*[]HeliosBlackoutWindow, bool)`

GetBlackoutWindowOk returns a tuple with the BlackoutWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutWindow

`func (o *HeliosPolicyResponse) SetBlackoutWindow(v []HeliosBlackoutWindow)`

SetBlackoutWindow sets BlackoutWindow field to given value.

### HasBlackoutWindow

`func (o *HeliosPolicyResponse) HasBlackoutWindow() bool`

HasBlackoutWindow returns a boolean if a field has been set.

### SetBlackoutWindowNil

`func (o *HeliosPolicyResponse) SetBlackoutWindowNil(b bool)`

 SetBlackoutWindowNil sets the value for BlackoutWindow to be an explicit nil

### UnsetBlackoutWindow
`func (o *HeliosPolicyResponse) UnsetBlackoutWindow()`

UnsetBlackoutWindow ensures that no value is present for BlackoutWindow, not even an explicit nil
### GetClusterIdentifier

`func (o *HeliosPolicyResponse) GetClusterIdentifier() string`

GetClusterIdentifier returns the ClusterIdentifier field if non-nil, zero value otherwise.

### GetClusterIdentifierOk

`func (o *HeliosPolicyResponse) GetClusterIdentifierOk() (*string, bool)`

GetClusterIdentifierOk returns a tuple with the ClusterIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIdentifier

`func (o *HeliosPolicyResponse) SetClusterIdentifier(v string)`

SetClusterIdentifier sets ClusterIdentifier field to given value.

### HasClusterIdentifier

`func (o *HeliosPolicyResponse) HasClusterIdentifier() bool`

HasClusterIdentifier returns a boolean if a field has been set.

### SetClusterIdentifierNil

`func (o *HeliosPolicyResponse) SetClusterIdentifierNil(b bool)`

 SetClusterIdentifierNil sets the value for ClusterIdentifier to be an explicit nil

### UnsetClusterIdentifier
`func (o *HeliosPolicyResponse) UnsetClusterIdentifier()`

UnsetClusterIdentifier ensures that no value is present for ClusterIdentifier, not even an explicit nil
### GetDataLock

`func (o *HeliosPolicyResponse) GetDataLock() string`

GetDataLock returns the DataLock field if non-nil, zero value otherwise.

### GetDataLockOk

`func (o *HeliosPolicyResponse) GetDataLockOk() (*string, bool)`

GetDataLockOk returns a tuple with the DataLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLock

`func (o *HeliosPolicyResponse) SetDataLock(v string)`

SetDataLock sets DataLock field to given value.

### HasDataLock

`func (o *HeliosPolicyResponse) HasDataLock() bool`

HasDataLock returns a boolean if a field has been set.

### SetDataLockNil

`func (o *HeliosPolicyResponse) SetDataLockNil(b bool)`

 SetDataLockNil sets the value for DataLock to be an explicit nil

### UnsetDataLock
`func (o *HeliosPolicyResponse) UnsetDataLock()`

UnsetDataLock ensures that no value is present for DataLock, not even an explicit nil
### GetDescription

`func (o *HeliosPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HeliosPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HeliosPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HeliosPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HeliosPolicyResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HeliosPolicyResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExtendedRetention

`func (o *HeliosPolicyResponse) GetExtendedRetention() []HeliosExtendedRetentionPolicy`

GetExtendedRetention returns the ExtendedRetention field if non-nil, zero value otherwise.

### GetExtendedRetentionOk

`func (o *HeliosPolicyResponse) GetExtendedRetentionOk() (*[]HeliosExtendedRetentionPolicy, bool)`

GetExtendedRetentionOk returns a tuple with the ExtendedRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedRetention

`func (o *HeliosPolicyResponse) SetExtendedRetention(v []HeliosExtendedRetentionPolicy)`

SetExtendedRetention sets ExtendedRetention field to given value.

### HasExtendedRetention

`func (o *HeliosPolicyResponse) HasExtendedRetention() bool`

HasExtendedRetention returns a boolean if a field has been set.

### SetExtendedRetentionNil

`func (o *HeliosPolicyResponse) SetExtendedRetentionNil(b bool)`

 SetExtendedRetentionNil sets the value for ExtendedRetention to be an explicit nil

### UnsetExtendedRetention
`func (o *HeliosPolicyResponse) UnsetExtendedRetention()`

UnsetExtendedRetention ensures that no value is present for ExtendedRetention, not even an explicit nil
### GetIsCBSEnabled

`func (o *HeliosPolicyResponse) GetIsCBSEnabled() bool`

GetIsCBSEnabled returns the IsCBSEnabled field if non-nil, zero value otherwise.

### GetIsCBSEnabledOk

`func (o *HeliosPolicyResponse) GetIsCBSEnabledOk() (*bool, bool)`

GetIsCBSEnabledOk returns a tuple with the IsCBSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCBSEnabled

`func (o *HeliosPolicyResponse) SetIsCBSEnabled(v bool)`

SetIsCBSEnabled sets IsCBSEnabled field to given value.

### HasIsCBSEnabled

`func (o *HeliosPolicyResponse) HasIsCBSEnabled() bool`

HasIsCBSEnabled returns a boolean if a field has been set.

### SetIsCBSEnabledNil

`func (o *HeliosPolicyResponse) SetIsCBSEnabledNil(b bool)`

 SetIsCBSEnabledNil sets the value for IsCBSEnabled to be an explicit nil

### UnsetIsCBSEnabled
`func (o *HeliosPolicyResponse) UnsetIsCBSEnabled()`

UnsetIsCBSEnabled ensures that no value is present for IsCBSEnabled, not even an explicit nil
### GetName

`func (o *HeliosPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosPolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *HeliosPolicyResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosPolicyResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRemoteTargetPolicy

`func (o *HeliosPolicyResponse) GetRemoteTargetPolicy() HeliosTargetsConfiguration`

GetRemoteTargetPolicy returns the RemoteTargetPolicy field if non-nil, zero value otherwise.

### GetRemoteTargetPolicyOk

`func (o *HeliosPolicyResponse) GetRemoteTargetPolicyOk() (*HeliosTargetsConfiguration, bool)`

GetRemoteTargetPolicyOk returns a tuple with the RemoteTargetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargetPolicy

`func (o *HeliosPolicyResponse) SetRemoteTargetPolicy(v HeliosTargetsConfiguration)`

SetRemoteTargetPolicy sets RemoteTargetPolicy field to given value.

### HasRemoteTargetPolicy

`func (o *HeliosPolicyResponse) HasRemoteTargetPolicy() bool`

HasRemoteTargetPolicy returns a boolean if a field has been set.

### GetRetryOptions

`func (o *HeliosPolicyResponse) GetRetryOptions() HeliosRetryOptions`

GetRetryOptions returns the RetryOptions field if non-nil, zero value otherwise.

### GetRetryOptionsOk

`func (o *HeliosPolicyResponse) GetRetryOptionsOk() (*HeliosRetryOptions, bool)`

GetRetryOptionsOk returns a tuple with the RetryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOptions

`func (o *HeliosPolicyResponse) SetRetryOptions(v HeliosRetryOptions)`

SetRetryOptions sets RetryOptions field to given value.

### HasRetryOptions

`func (o *HeliosPolicyResponse) HasRetryOptions() bool`

HasRetryOptions returns a boolean if a field has been set.

### GetTenantIds

`func (o *HeliosPolicyResponse) GetTenantIds() []*string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *HeliosPolicyResponse) GetTenantIdsOk() (*[]*string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *HeliosPolicyResponse) SetTenantIds(v []*string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *HeliosPolicyResponse) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### GetType

`func (o *HeliosPolicyResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HeliosPolicyResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HeliosPolicyResponse) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *HeliosPolicyResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HeliosPolicyResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVersion

`func (o *HeliosPolicyResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HeliosPolicyResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HeliosPolicyResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HeliosPolicyResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HeliosPolicyResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HeliosPolicyResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetId

`func (o *HeliosPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeliosPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeliosPolicyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HeliosPolicyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HeliosPolicyResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HeliosPolicyResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNumLinkedPolicies

`func (o *HeliosPolicyResponse) GetNumLinkedPolicies() int64`

GetNumLinkedPolicies returns the NumLinkedPolicies field if non-nil, zero value otherwise.

### GetNumLinkedPoliciesOk

`func (o *HeliosPolicyResponse) GetNumLinkedPoliciesOk() (*int64, bool)`

GetNumLinkedPoliciesOk returns a tuple with the NumLinkedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumLinkedPolicies

`func (o *HeliosPolicyResponse) SetNumLinkedPolicies(v int64)`

SetNumLinkedPolicies sets NumLinkedPolicies field to given value.

### HasNumLinkedPolicies

`func (o *HeliosPolicyResponse) HasNumLinkedPolicies() bool`

HasNumLinkedPolicies returns a boolean if a field has been set.

### SetNumLinkedPoliciesNil

`func (o *HeliosPolicyResponse) SetNumLinkedPoliciesNil(b bool)`

 SetNumLinkedPoliciesNil sets the value for NumLinkedPolicies to be an explicit nil

### UnsetNumLinkedPolicies
`func (o *HeliosPolicyResponse) UnsetNumLinkedPolicies()`

UnsetNumLinkedPolicies ensures that no value is present for NumLinkedPolicies, not even an explicit nil
### GetNumObjectProtections

`func (o *HeliosPolicyResponse) GetNumObjectProtections() int64`

GetNumObjectProtections returns the NumObjectProtections field if non-nil, zero value otherwise.

### GetNumObjectProtectionsOk

`func (o *HeliosPolicyResponse) GetNumObjectProtectionsOk() (*int64, bool)`

GetNumObjectProtectionsOk returns a tuple with the NumObjectProtections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumObjectProtections

`func (o *HeliosPolicyResponse) SetNumObjectProtections(v int64)`

SetNumObjectProtections sets NumObjectProtections field to given value.

### HasNumObjectProtections

`func (o *HeliosPolicyResponse) HasNumObjectProtections() bool`

HasNumObjectProtections returns a boolean if a field has been set.

### SetNumObjectProtectionsNil

`func (o *HeliosPolicyResponse) SetNumObjectProtectionsNil(b bool)`

 SetNumObjectProtectionsNil sets the value for NumObjectProtections to be an explicit nil

### UnsetNumObjectProtections
`func (o *HeliosPolicyResponse) UnsetNumObjectProtections()`

UnsetNumObjectProtections ensures that no value is present for NumObjectProtections, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


