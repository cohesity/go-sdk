# AuditLogConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **NullableBool** | Specifies if audit log is enabled. | 
**RetentionPeriodDays** | **NullableInt32** | Specifies the audit log retention period in days. Audit logs generated before the period of time specified by retentionPeriodDays are removed from the Cohesity Cluster. | 

## Methods

### NewAuditLogConfig

`func NewAuditLogConfig(enabled NullableBool, retentionPeriodDays NullableInt32, ) *AuditLogConfig`

NewAuditLogConfig instantiates a new AuditLogConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogConfigWithDefaults

`func NewAuditLogConfigWithDefaults() *AuditLogConfig`

NewAuditLogConfigWithDefaults instantiates a new AuditLogConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AuditLogConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuditLogConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuditLogConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *AuditLogConfig) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AuditLogConfig) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetRetentionPeriodDays

`func (o *AuditLogConfig) GetRetentionPeriodDays() int32`

GetRetentionPeriodDays returns the RetentionPeriodDays field if non-nil, zero value otherwise.

### GetRetentionPeriodDaysOk

`func (o *AuditLogConfig) GetRetentionPeriodDaysOk() (*int32, bool)`

GetRetentionPeriodDaysOk returns a tuple with the RetentionPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDays

`func (o *AuditLogConfig) SetRetentionPeriodDays(v int32)`

SetRetentionPeriodDays sets RetentionPeriodDays field to given value.


### SetRetentionPeriodDaysNil

`func (o *AuditLogConfig) SetRetentionPeriodDaysNil(b bool)`

 SetRetentionPeriodDaysNil sets the value for RetentionPeriodDays to be an explicit nil

### UnsetRetentionPeriodDays
`func (o *AuditLogConfig) UnsetRetentionPeriodDays()`

UnsetRetentionPeriodDays ensures that no value is present for RetentionPeriodDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


