# AlertMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertDocumentList** | Pointer to [**[]AlertDocument**](AlertDocument.md) | Specifies alert documentation one per each language supported. | [optional] 
**AlertTypeBucket** | Pointer to **NullableString** | Specifies the Alert type bucket. Specifies the Alert type bucket. kSoftware - Alerts which are related to Cohesity services. kHardware - Alerts related to hardware on which Cohesity software is running. kService - Alerts related to other external services. kOther - Alerts not of one of above categories. | [optional] 
**AlertTypeId** | Pointer to **NullableInt32** | Specifies unique id for the alert type. | [optional] 
**Category** | Pointer to **NullableString** | Specifies category of the alert type. Specifies the category of an Alert. kDisk - Alerts that are related to Disk. kNode - Alerts that are related to Node. kCluster - Alerts that are related to Cluster. kNodeHealth - Alerts that are related to Node Health. kClusterHealth - Alerts that are related to Cluster Health. kBackupRestore - Alerts that are related to Backup/Restore. kEncryption - Alerts that are related to Encryption. kArchivalRestore - Alerts that are related to Archival/Restore. kRemoteReplication - Alerts that are related to Remote Replication. kQuota - Alerts that are related to Quota. kLicense - Alerts that are related to License. kHeliosProActiveWellness - Alerts that are related to Helios ProActive Wellness. kHeliosAnalyticsJobs - Alerts that are related to Helios Analytics Jobs. kHeliosSignatureJobs - Alerts that are related to Helios Signature Jobs. kSecurity - Alerts that are related to Security. kAppsInfra - Alerts that are related to applications infra. kAntivirus - Alerts that are related to antivirus. kArchivalCopy - Alerts that are related to archival copies. | [optional] 
**DedupIntervalSeconds** | Pointer to **NullableInt32** | Specifies dedup interval in seconds. If the same alert is raised multiple times by any client in this duration, only one of them will be reported. | [optional] 
**DedupUntilResolved** | Pointer to **NullableBool** | Specifies if the alerts are to be deduped until the current one (if any) is resolved. | [optional] 
**HideAlertFromUser** | Pointer to **NullableBool** | Specifies whether to show the alert in the iris UI and CLI. | [optional] 
**IgnoreDuplicateOccurrences** | Pointer to **NullableBool** | Specifies whether to ignore duplicate occurrences completely. | [optional] 
**PrimaryKeyList** | Pointer to **[]string** | Specifies properties that serve as primary keys. | [optional] 
**PropertyList** | Pointer to **[]string** | Specifies list of properties that the client is supposed to provide when alert of this type is raised. | [optional] 
**SendSupportNotification** | Pointer to **NullableBool** | Specifies whether to send support notification for the alert. | [optional] 
**SnmpNotification** | Pointer to **NullableBool** | Specifies whether an SNMP notification is sent when an alert is raised. | [optional] 
**SyslogNotification** | Pointer to **NullableBool** | Specifies whether an syslog notification is sent when an alert is raised. | [optional] 
**Version** | Pointer to **NullableInt32** | Specifies version of the metadata. | [optional] 

## Methods

### NewAlertMetadata

`func NewAlertMetadata() *AlertMetadata`

NewAlertMetadata instantiates a new AlertMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertMetadataWithDefaults

`func NewAlertMetadataWithDefaults() *AlertMetadata`

NewAlertMetadataWithDefaults instantiates a new AlertMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertDocumentList

`func (o *AlertMetadata) GetAlertDocumentList() []AlertDocument`

GetAlertDocumentList returns the AlertDocumentList field if non-nil, zero value otherwise.

### GetAlertDocumentListOk

`func (o *AlertMetadata) GetAlertDocumentListOk() (*[]AlertDocument, bool)`

GetAlertDocumentListOk returns a tuple with the AlertDocumentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDocumentList

`func (o *AlertMetadata) SetAlertDocumentList(v []AlertDocument)`

SetAlertDocumentList sets AlertDocumentList field to given value.

### HasAlertDocumentList

`func (o *AlertMetadata) HasAlertDocumentList() bool`

HasAlertDocumentList returns a boolean if a field has been set.

### SetAlertDocumentListNil

`func (o *AlertMetadata) SetAlertDocumentListNil(b bool)`

 SetAlertDocumentListNil sets the value for AlertDocumentList to be an explicit nil

### UnsetAlertDocumentList
`func (o *AlertMetadata) UnsetAlertDocumentList()`

UnsetAlertDocumentList ensures that no value is present for AlertDocumentList, not even an explicit nil
### GetAlertTypeBucket

`func (o *AlertMetadata) GetAlertTypeBucket() string`

GetAlertTypeBucket returns the AlertTypeBucket field if non-nil, zero value otherwise.

### GetAlertTypeBucketOk

`func (o *AlertMetadata) GetAlertTypeBucketOk() (*string, bool)`

GetAlertTypeBucketOk returns a tuple with the AlertTypeBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeBucket

`func (o *AlertMetadata) SetAlertTypeBucket(v string)`

SetAlertTypeBucket sets AlertTypeBucket field to given value.

### HasAlertTypeBucket

`func (o *AlertMetadata) HasAlertTypeBucket() bool`

HasAlertTypeBucket returns a boolean if a field has been set.

### SetAlertTypeBucketNil

`func (o *AlertMetadata) SetAlertTypeBucketNil(b bool)`

 SetAlertTypeBucketNil sets the value for AlertTypeBucket to be an explicit nil

### UnsetAlertTypeBucket
`func (o *AlertMetadata) UnsetAlertTypeBucket()`

UnsetAlertTypeBucket ensures that no value is present for AlertTypeBucket, not even an explicit nil
### GetAlertTypeId

`func (o *AlertMetadata) GetAlertTypeId() int32`

GetAlertTypeId returns the AlertTypeId field if non-nil, zero value otherwise.

### GetAlertTypeIdOk

`func (o *AlertMetadata) GetAlertTypeIdOk() (*int32, bool)`

GetAlertTypeIdOk returns a tuple with the AlertTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypeId

`func (o *AlertMetadata) SetAlertTypeId(v int32)`

SetAlertTypeId sets AlertTypeId field to given value.

### HasAlertTypeId

`func (o *AlertMetadata) HasAlertTypeId() bool`

HasAlertTypeId returns a boolean if a field has been set.

### SetAlertTypeIdNil

`func (o *AlertMetadata) SetAlertTypeIdNil(b bool)`

 SetAlertTypeIdNil sets the value for AlertTypeId to be an explicit nil

### UnsetAlertTypeId
`func (o *AlertMetadata) UnsetAlertTypeId()`

UnsetAlertTypeId ensures that no value is present for AlertTypeId, not even an explicit nil
### GetCategory

`func (o *AlertMetadata) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AlertMetadata) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AlertMetadata) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AlertMetadata) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *AlertMetadata) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *AlertMetadata) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDedupIntervalSeconds

`func (o *AlertMetadata) GetDedupIntervalSeconds() int32`

GetDedupIntervalSeconds returns the DedupIntervalSeconds field if non-nil, zero value otherwise.

### GetDedupIntervalSecondsOk

`func (o *AlertMetadata) GetDedupIntervalSecondsOk() (*int32, bool)`

GetDedupIntervalSecondsOk returns a tuple with the DedupIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupIntervalSeconds

`func (o *AlertMetadata) SetDedupIntervalSeconds(v int32)`

SetDedupIntervalSeconds sets DedupIntervalSeconds field to given value.

### HasDedupIntervalSeconds

`func (o *AlertMetadata) HasDedupIntervalSeconds() bool`

HasDedupIntervalSeconds returns a boolean if a field has been set.

### SetDedupIntervalSecondsNil

`func (o *AlertMetadata) SetDedupIntervalSecondsNil(b bool)`

 SetDedupIntervalSecondsNil sets the value for DedupIntervalSeconds to be an explicit nil

### UnsetDedupIntervalSeconds
`func (o *AlertMetadata) UnsetDedupIntervalSeconds()`

UnsetDedupIntervalSeconds ensures that no value is present for DedupIntervalSeconds, not even an explicit nil
### GetDedupUntilResolved

`func (o *AlertMetadata) GetDedupUntilResolved() bool`

GetDedupUntilResolved returns the DedupUntilResolved field if non-nil, zero value otherwise.

### GetDedupUntilResolvedOk

`func (o *AlertMetadata) GetDedupUntilResolvedOk() (*bool, bool)`

GetDedupUntilResolvedOk returns a tuple with the DedupUntilResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupUntilResolved

`func (o *AlertMetadata) SetDedupUntilResolved(v bool)`

SetDedupUntilResolved sets DedupUntilResolved field to given value.

### HasDedupUntilResolved

`func (o *AlertMetadata) HasDedupUntilResolved() bool`

HasDedupUntilResolved returns a boolean if a field has been set.

### SetDedupUntilResolvedNil

`func (o *AlertMetadata) SetDedupUntilResolvedNil(b bool)`

 SetDedupUntilResolvedNil sets the value for DedupUntilResolved to be an explicit nil

### UnsetDedupUntilResolved
`func (o *AlertMetadata) UnsetDedupUntilResolved()`

UnsetDedupUntilResolved ensures that no value is present for DedupUntilResolved, not even an explicit nil
### GetHideAlertFromUser

`func (o *AlertMetadata) GetHideAlertFromUser() bool`

GetHideAlertFromUser returns the HideAlertFromUser field if non-nil, zero value otherwise.

### GetHideAlertFromUserOk

`func (o *AlertMetadata) GetHideAlertFromUserOk() (*bool, bool)`

GetHideAlertFromUserOk returns a tuple with the HideAlertFromUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAlertFromUser

`func (o *AlertMetadata) SetHideAlertFromUser(v bool)`

SetHideAlertFromUser sets HideAlertFromUser field to given value.

### HasHideAlertFromUser

`func (o *AlertMetadata) HasHideAlertFromUser() bool`

HasHideAlertFromUser returns a boolean if a field has been set.

### SetHideAlertFromUserNil

`func (o *AlertMetadata) SetHideAlertFromUserNil(b bool)`

 SetHideAlertFromUserNil sets the value for HideAlertFromUser to be an explicit nil

### UnsetHideAlertFromUser
`func (o *AlertMetadata) UnsetHideAlertFromUser()`

UnsetHideAlertFromUser ensures that no value is present for HideAlertFromUser, not even an explicit nil
### GetIgnoreDuplicateOccurrences

`func (o *AlertMetadata) GetIgnoreDuplicateOccurrences() bool`

GetIgnoreDuplicateOccurrences returns the IgnoreDuplicateOccurrences field if non-nil, zero value otherwise.

### GetIgnoreDuplicateOccurrencesOk

`func (o *AlertMetadata) GetIgnoreDuplicateOccurrencesOk() (*bool, bool)`

GetIgnoreDuplicateOccurrencesOk returns a tuple with the IgnoreDuplicateOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDuplicateOccurrences

`func (o *AlertMetadata) SetIgnoreDuplicateOccurrences(v bool)`

SetIgnoreDuplicateOccurrences sets IgnoreDuplicateOccurrences field to given value.

### HasIgnoreDuplicateOccurrences

`func (o *AlertMetadata) HasIgnoreDuplicateOccurrences() bool`

HasIgnoreDuplicateOccurrences returns a boolean if a field has been set.

### SetIgnoreDuplicateOccurrencesNil

`func (o *AlertMetadata) SetIgnoreDuplicateOccurrencesNil(b bool)`

 SetIgnoreDuplicateOccurrencesNil sets the value for IgnoreDuplicateOccurrences to be an explicit nil

### UnsetIgnoreDuplicateOccurrences
`func (o *AlertMetadata) UnsetIgnoreDuplicateOccurrences()`

UnsetIgnoreDuplicateOccurrences ensures that no value is present for IgnoreDuplicateOccurrences, not even an explicit nil
### GetPrimaryKeyList

`func (o *AlertMetadata) GetPrimaryKeyList() []string`

GetPrimaryKeyList returns the PrimaryKeyList field if non-nil, zero value otherwise.

### GetPrimaryKeyListOk

`func (o *AlertMetadata) GetPrimaryKeyListOk() (*[]string, bool)`

GetPrimaryKeyListOk returns a tuple with the PrimaryKeyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeyList

`func (o *AlertMetadata) SetPrimaryKeyList(v []string)`

SetPrimaryKeyList sets PrimaryKeyList field to given value.

### HasPrimaryKeyList

`func (o *AlertMetadata) HasPrimaryKeyList() bool`

HasPrimaryKeyList returns a boolean if a field has been set.

### SetPrimaryKeyListNil

`func (o *AlertMetadata) SetPrimaryKeyListNil(b bool)`

 SetPrimaryKeyListNil sets the value for PrimaryKeyList to be an explicit nil

### UnsetPrimaryKeyList
`func (o *AlertMetadata) UnsetPrimaryKeyList()`

UnsetPrimaryKeyList ensures that no value is present for PrimaryKeyList, not even an explicit nil
### GetPropertyList

`func (o *AlertMetadata) GetPropertyList() []string`

GetPropertyList returns the PropertyList field if non-nil, zero value otherwise.

### GetPropertyListOk

`func (o *AlertMetadata) GetPropertyListOk() (*[]string, bool)`

GetPropertyListOk returns a tuple with the PropertyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyList

`func (o *AlertMetadata) SetPropertyList(v []string)`

SetPropertyList sets PropertyList field to given value.

### HasPropertyList

`func (o *AlertMetadata) HasPropertyList() bool`

HasPropertyList returns a boolean if a field has been set.

### SetPropertyListNil

`func (o *AlertMetadata) SetPropertyListNil(b bool)`

 SetPropertyListNil sets the value for PropertyList to be an explicit nil

### UnsetPropertyList
`func (o *AlertMetadata) UnsetPropertyList()`

UnsetPropertyList ensures that no value is present for PropertyList, not even an explicit nil
### GetSendSupportNotification

`func (o *AlertMetadata) GetSendSupportNotification() bool`

GetSendSupportNotification returns the SendSupportNotification field if non-nil, zero value otherwise.

### GetSendSupportNotificationOk

`func (o *AlertMetadata) GetSendSupportNotificationOk() (*bool, bool)`

GetSendSupportNotificationOk returns a tuple with the SendSupportNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendSupportNotification

`func (o *AlertMetadata) SetSendSupportNotification(v bool)`

SetSendSupportNotification sets SendSupportNotification field to given value.

### HasSendSupportNotification

`func (o *AlertMetadata) HasSendSupportNotification() bool`

HasSendSupportNotification returns a boolean if a field has been set.

### SetSendSupportNotificationNil

`func (o *AlertMetadata) SetSendSupportNotificationNil(b bool)`

 SetSendSupportNotificationNil sets the value for SendSupportNotification to be an explicit nil

### UnsetSendSupportNotification
`func (o *AlertMetadata) UnsetSendSupportNotification()`

UnsetSendSupportNotification ensures that no value is present for SendSupportNotification, not even an explicit nil
### GetSnmpNotification

`func (o *AlertMetadata) GetSnmpNotification() bool`

GetSnmpNotification returns the SnmpNotification field if non-nil, zero value otherwise.

### GetSnmpNotificationOk

`func (o *AlertMetadata) GetSnmpNotificationOk() (*bool, bool)`

GetSnmpNotificationOk returns a tuple with the SnmpNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpNotification

`func (o *AlertMetadata) SetSnmpNotification(v bool)`

SetSnmpNotification sets SnmpNotification field to given value.

### HasSnmpNotification

`func (o *AlertMetadata) HasSnmpNotification() bool`

HasSnmpNotification returns a boolean if a field has been set.

### SetSnmpNotificationNil

`func (o *AlertMetadata) SetSnmpNotificationNil(b bool)`

 SetSnmpNotificationNil sets the value for SnmpNotification to be an explicit nil

### UnsetSnmpNotification
`func (o *AlertMetadata) UnsetSnmpNotification()`

UnsetSnmpNotification ensures that no value is present for SnmpNotification, not even an explicit nil
### GetSyslogNotification

`func (o *AlertMetadata) GetSyslogNotification() bool`

GetSyslogNotification returns the SyslogNotification field if non-nil, zero value otherwise.

### GetSyslogNotificationOk

`func (o *AlertMetadata) GetSyslogNotificationOk() (*bool, bool)`

GetSyslogNotificationOk returns a tuple with the SyslogNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogNotification

`func (o *AlertMetadata) SetSyslogNotification(v bool)`

SetSyslogNotification sets SyslogNotification field to given value.

### HasSyslogNotification

`func (o *AlertMetadata) HasSyslogNotification() bool`

HasSyslogNotification returns a boolean if a field has been set.

### SetSyslogNotificationNil

`func (o *AlertMetadata) SetSyslogNotificationNil(b bool)`

 SetSyslogNotificationNil sets the value for SyslogNotification to be an explicit nil

### UnsetSyslogNotification
`func (o *AlertMetadata) UnsetSyslogNotification()`

UnsetSyslogNotification ensures that no value is present for SyslogNotification, not even an explicit nil
### GetVersion

`func (o *AlertMetadata) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AlertMetadata) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AlertMetadata) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AlertMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *AlertMetadata) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AlertMetadata) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


