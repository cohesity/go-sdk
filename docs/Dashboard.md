# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogs** | Pointer to [**AuditLogsTile**](AuditLogsTile.md) |  | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Id of the cluster for which dashboard is given. | [optional] 
**Health** | Pointer to [**HealthTile**](HealthTile.md) |  | [optional] 
**Iops** | Pointer to [**IopsTile**](IopsTile.md) |  | [optional] 
**JobRuns** | Pointer to [**JobRunsTile**](JobRunsTile.md) |  | [optional] 
**ProtectedObjects** | Pointer to [**ProtectedObjectsTile**](ProtectedObjectsTile.md) |  | [optional] 
**Protection** | Pointer to [**ProtectionTile**](ProtectionTile.md) |  | [optional] 
**Recoveries** | Pointer to [**RecoveriesTile**](RecoveriesTile.md) |  | [optional] 
**StorageEfficiency** | Pointer to [**StorageEfficiencyTile**](StorageEfficiencyTile.md) |  | [optional] 
**Throughput** | Pointer to [**ThroughputTile**](ThroughputTile.md) |  | [optional] 

## Methods

### NewDashboard

`func NewDashboard() *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogs

`func (o *Dashboard) GetAuditLogs() AuditLogsTile`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *Dashboard) GetAuditLogsOk() (*AuditLogsTile, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *Dashboard) SetAuditLogs(v AuditLogsTile)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *Dashboard) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.

### GetClusterId

`func (o *Dashboard) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Dashboard) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Dashboard) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Dashboard) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *Dashboard) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *Dashboard) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetHealth

`func (o *Dashboard) GetHealth() HealthTile`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *Dashboard) GetHealthOk() (*HealthTile, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *Dashboard) SetHealth(v HealthTile)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *Dashboard) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetIops

`func (o *Dashboard) GetIops() IopsTile`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *Dashboard) GetIopsOk() (*IopsTile, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIops

`func (o *Dashboard) SetIops(v IopsTile)`

SetIops sets Iops field to given value.

### HasIops

`func (o *Dashboard) HasIops() bool`

HasIops returns a boolean if a field has been set.

### GetJobRuns

`func (o *Dashboard) GetJobRuns() JobRunsTile`

GetJobRuns returns the JobRuns field if non-nil, zero value otherwise.

### GetJobRunsOk

`func (o *Dashboard) GetJobRunsOk() (*JobRunsTile, bool)`

GetJobRunsOk returns a tuple with the JobRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRuns

`func (o *Dashboard) SetJobRuns(v JobRunsTile)`

SetJobRuns sets JobRuns field to given value.

### HasJobRuns

`func (o *Dashboard) HasJobRuns() bool`

HasJobRuns returns a boolean if a field has been set.

### GetProtectedObjects

`func (o *Dashboard) GetProtectedObjects() ProtectedObjectsTile`

GetProtectedObjects returns the ProtectedObjects field if non-nil, zero value otherwise.

### GetProtectedObjectsOk

`func (o *Dashboard) GetProtectedObjectsOk() (*ProtectedObjectsTile, bool)`

GetProtectedObjectsOk returns a tuple with the ProtectedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedObjects

`func (o *Dashboard) SetProtectedObjects(v ProtectedObjectsTile)`

SetProtectedObjects sets ProtectedObjects field to given value.

### HasProtectedObjects

`func (o *Dashboard) HasProtectedObjects() bool`

HasProtectedObjects returns a boolean if a field has been set.

### GetProtection

`func (o *Dashboard) GetProtection() ProtectionTile`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *Dashboard) GetProtectionOk() (*ProtectionTile, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *Dashboard) SetProtection(v ProtectionTile)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *Dashboard) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetRecoveries

`func (o *Dashboard) GetRecoveries() RecoveriesTile`

GetRecoveries returns the Recoveries field if non-nil, zero value otherwise.

### GetRecoveriesOk

`func (o *Dashboard) GetRecoveriesOk() (*RecoveriesTile, bool)`

GetRecoveriesOk returns a tuple with the Recoveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveries

`func (o *Dashboard) SetRecoveries(v RecoveriesTile)`

SetRecoveries sets Recoveries field to given value.

### HasRecoveries

`func (o *Dashboard) HasRecoveries() bool`

HasRecoveries returns a boolean if a field has been set.

### GetStorageEfficiency

`func (o *Dashboard) GetStorageEfficiency() StorageEfficiencyTile`

GetStorageEfficiency returns the StorageEfficiency field if non-nil, zero value otherwise.

### GetStorageEfficiencyOk

`func (o *Dashboard) GetStorageEfficiencyOk() (*StorageEfficiencyTile, bool)`

GetStorageEfficiencyOk returns a tuple with the StorageEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEfficiency

`func (o *Dashboard) SetStorageEfficiency(v StorageEfficiencyTile)`

SetStorageEfficiency sets StorageEfficiency field to given value.

### HasStorageEfficiency

`func (o *Dashboard) HasStorageEfficiency() bool`

HasStorageEfficiency returns a boolean if a field has been set.

### GetThroughput

`func (o *Dashboard) GetThroughput() ThroughputTile`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *Dashboard) GetThroughputOk() (*ThroughputTile, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *Dashboard) SetThroughput(v ThroughputTile)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *Dashboard) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


