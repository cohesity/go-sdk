# ClusterAuditLogConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **NullableBool** | Specifies if the Cluster audit logging is enabled on the Cohesity Cluster. If &#39;true&#39;, Cluster audit logging is enabled. Otherwise, it is disabled. | 
**RetentionPeriodDays** | **NullableInt32** | Specifies the number of days to keep (retain) the Cluster audit logs. Audit logs generated before the period of time specified by retentionPeriodDays are removed from the Cohesity Cluster. | 

## Methods

### NewClusterAuditLogConfiguration

`func NewClusterAuditLogConfiguration(enabled NullableBool, retentionPeriodDays NullableInt32, ) *ClusterAuditLogConfiguration`

NewClusterAuditLogConfiguration instantiates a new ClusterAuditLogConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAuditLogConfigurationWithDefaults

`func NewClusterAuditLogConfigurationWithDefaults() *ClusterAuditLogConfiguration`

NewClusterAuditLogConfigurationWithDefaults instantiates a new ClusterAuditLogConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterAuditLogConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterAuditLogConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterAuditLogConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *ClusterAuditLogConfiguration) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ClusterAuditLogConfiguration) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetRetentionPeriodDays

`func (o *ClusterAuditLogConfiguration) GetRetentionPeriodDays() int32`

GetRetentionPeriodDays returns the RetentionPeriodDays field if non-nil, zero value otherwise.

### GetRetentionPeriodDaysOk

`func (o *ClusterAuditLogConfiguration) GetRetentionPeriodDaysOk() (*int32, bool)`

GetRetentionPeriodDaysOk returns a tuple with the RetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDays

`func (o *ClusterAuditLogConfiguration) SetRetentionPeriodDays(v int32)`

SetRetentionPeriodDays sets RetentionPeriodDays field to given value.


### SetRetentionPeriodDaysNil

`func (o *ClusterAuditLogConfiguration) SetRetentionPeriodDaysNil(b bool)`

 SetRetentionPeriodDaysNil sets the value for RetentionPeriodDays to be an explicit nil

### UnsetRetentionPeriodDays
`func (o *ClusterAuditLogConfiguration) UnsetRetentionPeriodDays()`

UnsetRetentionPeriodDays ensures that no value is present for RetentionPeriodDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


