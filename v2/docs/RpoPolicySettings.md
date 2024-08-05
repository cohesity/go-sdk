# RpoPolicySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertingPolicy** | Pointer to [**ProtectionGroupAlertingPolicy**](ProtectionGroupAlertingPolicy.md) |  | [optional] 
**BackupQosPrincipal** | Pointer to **NullableString** | Specifies whether the data will be written to HDD or SSD. | [optional] 
**EnvBackupParams** | Pointer to [**EnvironmentTypeJobParams**](EnvironmentTypeJobParams.md) |  | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**StorageDomainId** | Pointer to **NullableInt64** | Specifies the Storage Domain to which data will be written | [optional] 

## Methods

### NewRpoPolicySettings

`func NewRpoPolicySettings() *RpoPolicySettings`

NewRpoPolicySettings instantiates a new RpoPolicySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpoPolicySettingsWithDefaults

`func NewRpoPolicySettingsWithDefaults() *RpoPolicySettings`

NewRpoPolicySettingsWithDefaults instantiates a new RpoPolicySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertingPolicy

`func (o *RpoPolicySettings) GetAlertingPolicy() ProtectionGroupAlertingPolicy`

GetAlertingPolicy returns the AlertingPolicy field if non-nil, zero value otherwise.

### GetAlertingPolicyOk

`func (o *RpoPolicySettings) GetAlertingPolicyOk() (*ProtectionGroupAlertingPolicy, bool)`

GetAlertingPolicyOk returns a tuple with the AlertingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingPolicy

`func (o *RpoPolicySettings) SetAlertingPolicy(v ProtectionGroupAlertingPolicy)`

SetAlertingPolicy sets AlertingPolicy field to given value.

### HasAlertingPolicy

`func (o *RpoPolicySettings) HasAlertingPolicy() bool`

HasAlertingPolicy returns a boolean if a field has been set.

### GetBackupQosPrincipal

`func (o *RpoPolicySettings) GetBackupQosPrincipal() string`

GetBackupQosPrincipal returns the BackupQosPrincipal field if non-nil, zero value otherwise.

### GetBackupQosPrincipalOk

`func (o *RpoPolicySettings) GetBackupQosPrincipalOk() (*string, bool)`

GetBackupQosPrincipalOk returns a tuple with the BackupQosPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupQosPrincipal

`func (o *RpoPolicySettings) SetBackupQosPrincipal(v string)`

SetBackupQosPrincipal sets BackupQosPrincipal field to given value.

### HasBackupQosPrincipal

`func (o *RpoPolicySettings) HasBackupQosPrincipal() bool`

HasBackupQosPrincipal returns a boolean if a field has been set.

### SetBackupQosPrincipalNil

`func (o *RpoPolicySettings) SetBackupQosPrincipalNil(b bool)`

 SetBackupQosPrincipalNil sets the value for BackupQosPrincipal to be an explicit nil

### UnsetBackupQosPrincipal
`func (o *RpoPolicySettings) UnsetBackupQosPrincipal()`

UnsetBackupQosPrincipal ensures that no value is present for BackupQosPrincipal, not even an explicit nil
### GetEnvBackupParams

`func (o *RpoPolicySettings) GetEnvBackupParams() EnvironmentTypeJobParams`

GetEnvBackupParams returns the EnvBackupParams field if non-nil, zero value otherwise.

### GetEnvBackupParamsOk

`func (o *RpoPolicySettings) GetEnvBackupParamsOk() (*EnvironmentTypeJobParams, bool)`

GetEnvBackupParamsOk returns a tuple with the EnvBackupParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvBackupParams

`func (o *RpoPolicySettings) SetEnvBackupParams(v EnvironmentTypeJobParams)`

SetEnvBackupParams sets EnvBackupParams field to given value.

### HasEnvBackupParams

`func (o *RpoPolicySettings) HasEnvBackupParams() bool`

HasEnvBackupParams returns a boolean if a field has been set.

### GetIndexingPolicy

`func (o *RpoPolicySettings) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *RpoPolicySettings) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *RpoPolicySettings) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *RpoPolicySettings) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetStorageDomainId

`func (o *RpoPolicySettings) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *RpoPolicySettings) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *RpoPolicySettings) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.

### HasStorageDomainId

`func (o *RpoPolicySettings) HasStorageDomainId() bool`

HasStorageDomainId returns a boolean if a field has been set.

### SetStorageDomainIdNil

`func (o *RpoPolicySettings) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *RpoPolicySettings) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


