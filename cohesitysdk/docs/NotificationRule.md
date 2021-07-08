# NotificationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertTypeList** | Pointer to **[]int32** | Specifies alert types this rule is applicable to. | [optional] 
**Categories** | Pointer to **[]string** | Specifies alert categories this rule is applicable to. Specifies the category of an Alert. kDisk - Alerts that are related to Disk. kNode - Alerts that are related to Node. kCluster - Alerts that are related to Cluster. kNodeHealth - Alerts that are related to Node Health. kClusterHealth - Alerts that are related to Cluster Health. kBackupRestore - Alerts that are related to Backup/Restore. kEncryption - Alerts that are related to Encryption. kArchivalRestore - Alerts that are related to Archival/Restore. kRemoteReplication - Alerts that are related to Remote Replication. kQuota - Alerts that are related to Quota. kLicense - Alerts that are related to License. kHeliosProActiveWellness - Alerts that are related to Helios ProActive Wellness. kHeliosAnalyticsJobs - Alerts that are related to Helios Analytics Jobs. kHeliosSignatureJobs - Alerts that are related to Helios Signature Jobs. kSecurity - Alerts that are related to Security. kAppsInfra - Alerts that are related to applications infra. kAntivirus - Alerts that are related to antivirus. kArchivalCopy - Alerts that are related to archival copies. | [optional] 
**EmailDeliveryTargets** | Pointer to [**[]EmailDeliveryTarget**](EmailDeliveryTarget.md) | Specifies email addresses to be notified when an alert matching this rule is generated. | [optional] 
**RuleId** | Pointer to **NullableInt64** | Specifies id of the alert delivery rule. | [optional] 
**RuleName** | Pointer to **NullableString** | Specifies name of the alert delivery rule. | [optional] 
**Severities** | Pointer to **[]string** | Specifies alert severity types this rule is applicable to. Specifies the severity level of an Alert. kCritical - Alerts whose severity type is Critical. kWarning - Alerts whose severity type is Warning. kInfo - Alerts whose severity type is Info. | [optional] 
**SnmpEnabled** | Pointer to **NullableBool** | Specifies whether SNMP notification to be invoked when an alert matching this rule is generated. | [optional] 
**SyslogEnabled** | Pointer to **NullableBool** | Specifies whether syslog notification to be invoked when an alert matching this rule is generated. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies tenant id this rule is applicable to. | [optional] 
**WebHookDeliveryTargets** | Pointer to [**[]WebHookDeliveryTarget**](WebHookDeliveryTarget.md) | Specifies external api urls to be invoked when an alert matching this rule is generated. | [optional] 

## Methods

### NewNotificationRule

`func NewNotificationRule() *NotificationRule`

NewNotificationRule instantiates a new NotificationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleWithDefaults

`func NewNotificationRuleWithDefaults() *NotificationRule`

NewNotificationRuleWithDefaults instantiates a new NotificationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertTypeList

`func (o *NotificationRule) GetAlertTypeList() []int32`

GetAlertTypeList returns the AlertTypeList field if non-nil, zero value otherwise.

### GetAlertTypeListOk

`func (o *NotificationRule) GetAlertTypeListOk() (*[]int32, bool)`

GetAlertTypeListOk returns a tuple with the AlertTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeList

`func (o *NotificationRule) SetAlertTypeList(v []int32)`

SetAlertTypeList sets AlertTypeList field to given value.

### HasAlertTypeList

`func (o *NotificationRule) HasAlertTypeList() bool`

HasAlertTypeList returns a boolean if a field has been set.

### SetAlertTypeListNil

`func (o *NotificationRule) SetAlertTypeListNil(b bool)`

 SetAlertTypeListNil sets the value for AlertTypeList to be an explicit nil

### UnsetAlertTypeList
`func (o *NotificationRule) UnsetAlertTypeList()`

UnsetAlertTypeList ensures that no value is present for AlertTypeList, not even an explicit nil
### GetCategories

`func (o *NotificationRule) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *NotificationRule) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *NotificationRule) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *NotificationRule) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *NotificationRule) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *NotificationRule) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetEmailDeliveryTargets

`func (o *NotificationRule) GetEmailDeliveryTargets() []EmailDeliveryTarget`

GetEmailDeliveryTargets returns the EmailDeliveryTargets field if non-nil, zero value otherwise.

### GetEmailDeliveryTargetsOk

`func (o *NotificationRule) GetEmailDeliveryTargetsOk() (*[]EmailDeliveryTarget, bool)`

GetEmailDeliveryTargetsOk returns a tuple with the EmailDeliveryTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDeliveryTargets

`func (o *NotificationRule) SetEmailDeliveryTargets(v []EmailDeliveryTarget)`

SetEmailDeliveryTargets sets EmailDeliveryTargets field to given value.

### HasEmailDeliveryTargets

`func (o *NotificationRule) HasEmailDeliveryTargets() bool`

HasEmailDeliveryTargets returns a boolean if a field has been set.

### SetEmailDeliveryTargetsNil

`func (o *NotificationRule) SetEmailDeliveryTargetsNil(b bool)`

 SetEmailDeliveryTargetsNil sets the value for EmailDeliveryTargets to be an explicit nil

### UnsetEmailDeliveryTargets
`func (o *NotificationRule) UnsetEmailDeliveryTargets()`

UnsetEmailDeliveryTargets ensures that no value is present for EmailDeliveryTargets, not even an explicit nil
### GetRuleId

`func (o *NotificationRule) GetRuleId() int64`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *NotificationRule) GetRuleIdOk() (*int64, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *NotificationRule) SetRuleId(v int64)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *NotificationRule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *NotificationRule) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *NotificationRule) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetRuleName

`func (o *NotificationRule) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *NotificationRule) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *NotificationRule) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *NotificationRule) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### SetRuleNameNil

`func (o *NotificationRule) SetRuleNameNil(b bool)`

 SetRuleNameNil sets the value for RuleName to be an explicit nil

### UnsetRuleName
`func (o *NotificationRule) UnsetRuleName()`

UnsetRuleName ensures that no value is present for RuleName, not even an explicit nil
### GetSeverities

`func (o *NotificationRule) GetSeverities() []string`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *NotificationRule) GetSeveritiesOk() (*[]string, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *NotificationRule) SetSeverities(v []string)`

SetSeverities sets Severities field to given value.

### HasSeverities

`func (o *NotificationRule) HasSeverities() bool`

HasSeverities returns a boolean if a field has been set.

### SetSeveritiesNil

`func (o *NotificationRule) SetSeveritiesNil(b bool)`

 SetSeveritiesNil sets the value for Severities to be an explicit nil

### UnsetSeverities
`func (o *NotificationRule) UnsetSeverities()`

UnsetSeverities ensures that no value is present for Severities, not even an explicit nil
### GetSnmpEnabled

`func (o *NotificationRule) GetSnmpEnabled() bool`

GetSnmpEnabled returns the SnmpEnabled field if non-nil, zero value otherwise.

### GetSnmpEnabledOk

`func (o *NotificationRule) GetSnmpEnabledOk() (*bool, bool)`

GetSnmpEnabledOk returns a tuple with the SnmpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpEnabled

`func (o *NotificationRule) SetSnmpEnabled(v bool)`

SetSnmpEnabled sets SnmpEnabled field to given value.

### HasSnmpEnabled

`func (o *NotificationRule) HasSnmpEnabled() bool`

HasSnmpEnabled returns a boolean if a field has been set.

### SetSnmpEnabledNil

`func (o *NotificationRule) SetSnmpEnabledNil(b bool)`

 SetSnmpEnabledNil sets the value for SnmpEnabled to be an explicit nil

### UnsetSnmpEnabled
`func (o *NotificationRule) UnsetSnmpEnabled()`

UnsetSnmpEnabled ensures that no value is present for SnmpEnabled, not even an explicit nil
### GetSyslogEnabled

`func (o *NotificationRule) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *NotificationRule) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *NotificationRule) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *NotificationRule) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.

### SetSyslogEnabledNil

`func (o *NotificationRule) SetSyslogEnabledNil(b bool)`

 SetSyslogEnabledNil sets the value for SyslogEnabled to be an explicit nil

### UnsetSyslogEnabled
`func (o *NotificationRule) UnsetSyslogEnabled()`

UnsetSyslogEnabled ensures that no value is present for SyslogEnabled, not even an explicit nil
### GetTenantId

`func (o *NotificationRule) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NotificationRule) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NotificationRule) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *NotificationRule) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *NotificationRule) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *NotificationRule) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetWebHookDeliveryTargets

`func (o *NotificationRule) GetWebHookDeliveryTargets() []WebHookDeliveryTarget`

GetWebHookDeliveryTargets returns the WebHookDeliveryTargets field if non-nil, zero value otherwise.

### GetWebHookDeliveryTargetsOk

`func (o *NotificationRule) GetWebHookDeliveryTargetsOk() (*[]WebHookDeliveryTarget, bool)`

GetWebHookDeliveryTargetsOk returns a tuple with the WebHookDeliveryTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookDeliveryTargets

`func (o *NotificationRule) SetWebHookDeliveryTargets(v []WebHookDeliveryTarget)`

SetWebHookDeliveryTargets sets WebHookDeliveryTargets field to given value.

### HasWebHookDeliveryTargets

`func (o *NotificationRule) HasWebHookDeliveryTargets() bool`

HasWebHookDeliveryTargets returns a boolean if a field has been set.

### SetWebHookDeliveryTargetsNil

`func (o *NotificationRule) SetWebHookDeliveryTargetsNil(b bool)`

 SetWebHookDeliveryTargetsNil sets the value for WebHookDeliveryTargets to be an explicit nil

### UnsetWebHookDeliveryTargets
`func (o *NotificationRule) UnsetWebHookDeliveryTargets()`

UnsetWebHookDeliveryTargets ensures that no value is present for WebHookDeliveryTargets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


