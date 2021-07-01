# ClusterAuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the action that caused the log to be generated. | [optional] 
**ClusterInfo** | Pointer to **NullableString** | Specifies the information of the cluster. | [optional] 
**Details** | Pointer to **NullableString** | Specifies more information about the action. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the domain of the user who caused the action that generated the log. | [optional] 
**EntityId** | Pointer to **NullableString** | Specifies the id of the entity (object) that the action is invoked on. | [optional] 
**EntityName** | Pointer to **NullableString** | Specifies the entity (object) name that the action is invoked on. For example, if a Job called BackupEng is paused, this field returns BackupEng. | [optional] 
**EntityType** | Pointer to **NullableString** | Specifies the type of the entity (object) that the action is invoked on. For example, if a Job called BackupEng is paused, this field returns &#39;Protection Job&#39;. | [optional] 
**HumanTimestamp** | Pointer to **NullableString** | Specifies the time when the log was generated. The time is specified using a human readable timestamp. | [optional] 
**Impersonation** | Pointer to **NullableBool** | Specifies if the log was generated during impersonation. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the IP address of the user making this action. | [optional] 
**NewRecord** | Pointer to **NullableString** | Specifies the record after the action is invoked. | [optional] 
**OriginalTenant** | Pointer to [**Tenant**](Tenant.md) |  | [optional] 
**PreviousRecord** | Pointer to **NullableString** | Specifies the record before the action is invoked. | [optional] 
**Tenant** | Pointer to [**Tenant**](Tenant.md) |  | [optional] 
**TimestampUsecs** | Pointer to **NullableInt64** | Specifies the time when the log was generated. The time is specified using a Unix epoch Timestamp (in microseconds). | [optional] 
**UserName** | Pointer to **NullableString** | Specifies the user who caused the action that generated the log. | [optional] 

## Methods

### NewClusterAuditLog

`func NewClusterAuditLog() *ClusterAuditLog`

NewClusterAuditLog instantiates a new ClusterAuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAuditLogWithDefaults

`func NewClusterAuditLogWithDefaults() *ClusterAuditLog`

NewClusterAuditLogWithDefaults instantiates a new ClusterAuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ClusterAuditLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ClusterAuditLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ClusterAuditLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ClusterAuditLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *ClusterAuditLog) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *ClusterAuditLog) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetClusterInfo

`func (o *ClusterAuditLog) GetClusterInfo() string`

GetClusterInfo returns the ClusterInfo field if non-nil, zero value otherwise.

### GetClusterInfoOk

`func (o *ClusterAuditLog) GetClusterInfoOk() (*string, bool)`

GetClusterInfoOk returns a tuple with the ClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInfo

`func (o *ClusterAuditLog) SetClusterInfo(v string)`

SetClusterInfo sets ClusterInfo field to given value.

### HasClusterInfo

`func (o *ClusterAuditLog) HasClusterInfo() bool`

HasClusterInfo returns a boolean if a field has been set.

### SetClusterInfoNil

`func (o *ClusterAuditLog) SetClusterInfoNil(b bool)`

 SetClusterInfoNil sets the value for ClusterInfo to be an explicit nil

### UnsetClusterInfo
`func (o *ClusterAuditLog) UnsetClusterInfo()`

UnsetClusterInfo ensures that no value is present for ClusterInfo, not even an explicit nil
### GetDetails

`func (o *ClusterAuditLog) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ClusterAuditLog) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ClusterAuditLog) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ClusterAuditLog) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ClusterAuditLog) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ClusterAuditLog) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetDomain

`func (o *ClusterAuditLog) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ClusterAuditLog) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ClusterAuditLog) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ClusterAuditLog) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *ClusterAuditLog) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *ClusterAuditLog) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEntityId

`func (o *ClusterAuditLog) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ClusterAuditLog) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ClusterAuditLog) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ClusterAuditLog) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *ClusterAuditLog) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *ClusterAuditLog) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetEntityName

`func (o *ClusterAuditLog) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *ClusterAuditLog) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *ClusterAuditLog) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *ClusterAuditLog) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### SetEntityNameNil

`func (o *ClusterAuditLog) SetEntityNameNil(b bool)`

 SetEntityNameNil sets the value for EntityName to be an explicit nil

### UnsetEntityName
`func (o *ClusterAuditLog) UnsetEntityName()`

UnsetEntityName ensures that no value is present for EntityName, not even an explicit nil
### GetEntityType

`func (o *ClusterAuditLog) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ClusterAuditLog) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ClusterAuditLog) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ClusterAuditLog) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *ClusterAuditLog) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *ClusterAuditLog) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetHumanTimestamp

`func (o *ClusterAuditLog) GetHumanTimestamp() string`

GetHumanTimestamp returns the HumanTimestamp field if non-nil, zero value otherwise.

### GetHumanTimestampOk

`func (o *ClusterAuditLog) GetHumanTimestampOk() (*string, bool)`

GetHumanTimestampOk returns a tuple with the HumanTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanTimestamp

`func (o *ClusterAuditLog) SetHumanTimestamp(v string)`

SetHumanTimestamp sets HumanTimestamp field to given value.

### HasHumanTimestamp

`func (o *ClusterAuditLog) HasHumanTimestamp() bool`

HasHumanTimestamp returns a boolean if a field has been set.

### SetHumanTimestampNil

`func (o *ClusterAuditLog) SetHumanTimestampNil(b bool)`

 SetHumanTimestampNil sets the value for HumanTimestamp to be an explicit nil

### UnsetHumanTimestamp
`func (o *ClusterAuditLog) UnsetHumanTimestamp()`

UnsetHumanTimestamp ensures that no value is present for HumanTimestamp, not even an explicit nil
### GetImpersonation

`func (o *ClusterAuditLog) GetImpersonation() bool`

GetImpersonation returns the Impersonation field if non-nil, zero value otherwise.

### GetImpersonationOk

`func (o *ClusterAuditLog) GetImpersonationOk() (*bool, bool)`

GetImpersonationOk returns a tuple with the Impersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonation

`func (o *ClusterAuditLog) SetImpersonation(v bool)`

SetImpersonation sets Impersonation field to given value.

### HasImpersonation

`func (o *ClusterAuditLog) HasImpersonation() bool`

HasImpersonation returns a boolean if a field has been set.

### SetImpersonationNil

`func (o *ClusterAuditLog) SetImpersonationNil(b bool)`

 SetImpersonationNil sets the value for Impersonation to be an explicit nil

### UnsetImpersonation
`func (o *ClusterAuditLog) UnsetImpersonation()`

UnsetImpersonation ensures that no value is present for Impersonation, not even an explicit nil
### GetIp

`func (o *ClusterAuditLog) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClusterAuditLog) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClusterAuditLog) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClusterAuditLog) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *ClusterAuditLog) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ClusterAuditLog) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetNewRecord

`func (o *ClusterAuditLog) GetNewRecord() string`

GetNewRecord returns the NewRecord field if non-nil, zero value otherwise.

### GetNewRecordOk

`func (o *ClusterAuditLog) GetNewRecordOk() (*string, bool)`

GetNewRecordOk returns a tuple with the NewRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRecord

`func (o *ClusterAuditLog) SetNewRecord(v string)`

SetNewRecord sets NewRecord field to given value.

### HasNewRecord

`func (o *ClusterAuditLog) HasNewRecord() bool`

HasNewRecord returns a boolean if a field has been set.

### SetNewRecordNil

`func (o *ClusterAuditLog) SetNewRecordNil(b bool)`

 SetNewRecordNil sets the value for NewRecord to be an explicit nil

### UnsetNewRecord
`func (o *ClusterAuditLog) UnsetNewRecord()`

UnsetNewRecord ensures that no value is present for NewRecord, not even an explicit nil
### GetOriginalTenant

`func (o *ClusterAuditLog) GetOriginalTenant() Tenant`

GetOriginalTenant returns the OriginalTenant field if non-nil, zero value otherwise.

### GetOriginalTenantOk

`func (o *ClusterAuditLog) GetOriginalTenantOk() (*Tenant, bool)`

GetOriginalTenantOk returns a tuple with the OriginalTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTenant

`func (o *ClusterAuditLog) SetOriginalTenant(v Tenant)`

SetOriginalTenant sets OriginalTenant field to given value.

### HasOriginalTenant

`func (o *ClusterAuditLog) HasOriginalTenant() bool`

HasOriginalTenant returns a boolean if a field has been set.

### GetPreviousRecord

`func (o *ClusterAuditLog) GetPreviousRecord() string`

GetPreviousRecord returns the PreviousRecord field if non-nil, zero value otherwise.

### GetPreviousRecordOk

`func (o *ClusterAuditLog) GetPreviousRecordOk() (*string, bool)`

GetPreviousRecordOk returns a tuple with the PreviousRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRecord

`func (o *ClusterAuditLog) SetPreviousRecord(v string)`

SetPreviousRecord sets PreviousRecord field to given value.

### HasPreviousRecord

`func (o *ClusterAuditLog) HasPreviousRecord() bool`

HasPreviousRecord returns a boolean if a field has been set.

### SetPreviousRecordNil

`func (o *ClusterAuditLog) SetPreviousRecordNil(b bool)`

 SetPreviousRecordNil sets the value for PreviousRecord to be an explicit nil

### UnsetPreviousRecord
`func (o *ClusterAuditLog) UnsetPreviousRecord()`

UnsetPreviousRecord ensures that no value is present for PreviousRecord, not even an explicit nil
### GetTenant

`func (o *ClusterAuditLog) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ClusterAuditLog) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ClusterAuditLog) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ClusterAuditLog) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetTimestampUsecs

`func (o *ClusterAuditLog) GetTimestampUsecs() int64`

GetTimestampUsecs returns the TimestampUsecs field if non-nil, zero value otherwise.

### GetTimestampUsecsOk

`func (o *ClusterAuditLog) GetTimestampUsecsOk() (*int64, bool)`

GetTimestampUsecsOk returns a tuple with the TimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUsecs

`func (o *ClusterAuditLog) SetTimestampUsecs(v int64)`

SetTimestampUsecs sets TimestampUsecs field to given value.

### HasTimestampUsecs

`func (o *ClusterAuditLog) HasTimestampUsecs() bool`

HasTimestampUsecs returns a boolean if a field has been set.

### SetTimestampUsecsNil

`func (o *ClusterAuditLog) SetTimestampUsecsNil(b bool)`

 SetTimestampUsecsNil sets the value for TimestampUsecs to be an explicit nil

### UnsetTimestampUsecs
`func (o *ClusterAuditLog) UnsetTimestampUsecs()`

UnsetTimestampUsecs ensures that no value is present for TimestampUsecs, not even an explicit nil
### GetUserName

`func (o *ClusterAuditLog) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ClusterAuditLog) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ClusterAuditLog) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ClusterAuditLog) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ClusterAuditLog) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ClusterAuditLog) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


