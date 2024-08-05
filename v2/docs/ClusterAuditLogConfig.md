# ClusterAuditLogConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **NullableBool** | Specifies if audit log is enabled. | 
**RetentionPeriodDays** | **NullableInt32** | Specifies the audit log retention period in days. Audit logs generated before the period of time specified by retentionPeriodDays are removed from the Cohesity Cluster. | 
**VerboseAudit** | Pointer to **NullableBool** | Specifies if the Cluster audit logging includes prev value and new value. | [optional] 

## Methods

### NewClusterAuditLogConfig

`func NewClusterAuditLogConfig(enabled NullableBool, retentionPeriodDays NullableInt32, ) *ClusterAuditLogConfig`

NewClusterAuditLogConfig instantiates a new ClusterAuditLogConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAuditLogConfigWithDefaults

`func NewClusterAuditLogConfigWithDefaults() *ClusterAuditLogConfig`

NewClusterAuditLogConfigWithDefaults instantiates a new ClusterAuditLogConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterAuditLogConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterAuditLogConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterAuditLogConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *ClusterAuditLogConfig) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ClusterAuditLogConfig) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetRetentionPeriodDays

`func (o *ClusterAuditLogConfig) GetRetentionPeriodDays() int32`

GetRetentionPeriodDays returns the RetentionPeriodDays field if non-nil, zero value otherwise.

### GetRetentionPeriodDaysOk

`func (o *ClusterAuditLogConfig) GetRetentionPeriodDaysOk() (*int32, bool)`

GetRetentionPeriodDaysOk returns a tuple with the RetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDays

`func (o *ClusterAuditLogConfig) SetRetentionPeriodDays(v int32)`

SetRetentionPeriodDays sets RetentionPeriodDays field to given value.


### SetRetentionPeriodDaysNil

`func (o *ClusterAuditLogConfig) SetRetentionPeriodDaysNil(b bool)`

 SetRetentionPeriodDaysNil sets the value for RetentionPeriodDays to be an explicit nil

### UnsetRetentionPeriodDays
`func (o *ClusterAuditLogConfig) UnsetRetentionPeriodDays()`

UnsetRetentionPeriodDays ensures that no value is present for RetentionPeriodDays, not even an explicit nil
### GetVerboseAudit

`func (o *ClusterAuditLogConfig) GetVerboseAudit() bool`

GetVerboseAudit returns the VerboseAudit field if non-nil, zero value otherwise.

### GetVerboseAuditOk

`func (o *ClusterAuditLogConfig) GetVerboseAuditOk() (*bool, bool)`

GetVerboseAuditOk returns a tuple with the VerboseAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseAudit

`func (o *ClusterAuditLogConfig) SetVerboseAudit(v bool)`

SetVerboseAudit sets VerboseAudit field to given value.

### HasVerboseAudit

`func (o *ClusterAuditLogConfig) HasVerboseAudit() bool`

HasVerboseAudit returns a boolean if a field has been set.

### SetVerboseAuditNil

`func (o *ClusterAuditLogConfig) SetVerboseAuditNil(b bool)`

 SetVerboseAuditNil sets the value for VerboseAudit to be an explicit nil

### UnsetVerboseAudit
`func (o *ClusterAuditLogConfig) UnsetVerboseAudit()`

UnsetVerboseAudit ensures that no value is present for VerboseAudit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


