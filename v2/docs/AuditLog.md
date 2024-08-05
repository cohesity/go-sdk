# AuditLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the action type of this audit log. | [optional] 
**Details** | Pointer to **NullableString** | Specifies the change details of this audit log. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the domain of user who made this audit log. | [optional] 
**EntityName** | Pointer to **NullableString** | Specifies the entity name. | [optional] 
**EntityType** | Pointer to **NullableString** | Specifies the entity type. | [optional] 
**Ip** | Pointer to **NullableString** | Specifies the ip of user who made this audit log. | [optional] 
**IsImpersonation** | Pointer to **NullableBool** | Specifies if the action is made through impersonation. | [optional] 
**NewRecord** | Pointer to **NullableString** | Specifies the record after the action is invoked. This will be returned only if verbose audit is enabled.  | [optional] 
**OriginalTenantId** | Pointer to **NullableString** | Specifies the original tenant id who made this audit log. | [optional] 
**OriginalTenantName** | Pointer to **NullableString** | Specifies the original tenant name who made this audit log. | [optional] 
**PreviousRecord** | Pointer to **NullableString** | Specifies the record before the action is invoked. This will be returned only if verbose audit is enabled.  | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant id who made this audit log. | [optional] 
**TenantName** | Pointer to **NullableString** | Specifies the tenant name who made this audit log. | [optional] 
**TimestampUsecs** | Pointer to **NullableInt64** | Specifies a unix timestamp in micro seconds when the audit log was taken. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username who made this audit log. | [optional] 

## Methods

### NewAuditLog

`func NewAuditLog() *AuditLog`

NewAuditLog instantiates a new AuditLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogWithDefaults

`func NewAuditLogWithDefaults() *AuditLog`

NewAuditLogWithDefaults instantiates a new AuditLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuditLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *AuditLog) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *AuditLog) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDetails

`func (o *AuditLog) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AuditLog) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AuditLog) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AuditLog) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *AuditLog) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *AuditLog) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetDomain

`func (o *AuditLog) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AuditLog) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AuditLog) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AuditLog) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *AuditLog) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *AuditLog) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEntityName

`func (o *AuditLog) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *AuditLog) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *AuditLog) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *AuditLog) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### SetEntityNameNil

`func (o *AuditLog) SetEntityNameNil(b bool)`

 SetEntityNameNil sets the value for EntityName to be an explicit nil

### UnsetEntityName
`func (o *AuditLog) UnsetEntityName()`

UnsetEntityName ensures that no value is present for EntityName, not even an explicit nil
### GetEntityType

`func (o *AuditLog) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AuditLog) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AuditLog) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *AuditLog) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *AuditLog) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *AuditLog) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetIp

`func (o *AuditLog) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AuditLog) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AuditLog) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AuditLog) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *AuditLog) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *AuditLog) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetIsImpersonation

`func (o *AuditLog) GetIsImpersonation() bool`

GetIsImpersonation returns the IsImpersonation field if non-nil, zero value otherwise.

### GetIsImpersonationOk

`func (o *AuditLog) GetIsImpersonationOk() (*bool, bool)`

GetIsImpersonationOk returns a tuple with the IsImpersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImpersonation

`func (o *AuditLog) SetIsImpersonation(v bool)`

SetIsImpersonation sets IsImpersonation field to given value.

### HasIsImpersonation

`func (o *AuditLog) HasIsImpersonation() bool`

HasIsImpersonation returns a boolean if a field has been set.

### SetIsImpersonationNil

`func (o *AuditLog) SetIsImpersonationNil(b bool)`

 SetIsImpersonationNil sets the value for IsImpersonation to be an explicit nil

### UnsetIsImpersonation
`func (o *AuditLog) UnsetIsImpersonation()`

UnsetIsImpersonation ensures that no value is present for IsImpersonation, not even an explicit nil
### GetNewRecord

`func (o *AuditLog) GetNewRecord() string`

GetNewRecord returns the NewRecord field if non-nil, zero value otherwise.

### GetNewRecordOk

`func (o *AuditLog) GetNewRecordOk() (*string, bool)`

GetNewRecordOk returns a tuple with the NewRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRecord

`func (o *AuditLog) SetNewRecord(v string)`

SetNewRecord sets NewRecord field to given value.

### HasNewRecord

`func (o *AuditLog) HasNewRecord() bool`

HasNewRecord returns a boolean if a field has been set.

### SetNewRecordNil

`func (o *AuditLog) SetNewRecordNil(b bool)`

 SetNewRecordNil sets the value for NewRecord to be an explicit nil

### UnsetNewRecord
`func (o *AuditLog) UnsetNewRecord()`

UnsetNewRecord ensures that no value is present for NewRecord, not even an explicit nil
### GetOriginalTenantId

`func (o *AuditLog) GetOriginalTenantId() string`

GetOriginalTenantId returns the OriginalTenantId field if non-nil, zero value otherwise.

### GetOriginalTenantIdOk

`func (o *AuditLog) GetOriginalTenantIdOk() (*string, bool)`

GetOriginalTenantIdOk returns a tuple with the OriginalTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTenantId

`func (o *AuditLog) SetOriginalTenantId(v string)`

SetOriginalTenantId sets OriginalTenantId field to given value.

### HasOriginalTenantId

`func (o *AuditLog) HasOriginalTenantId() bool`

HasOriginalTenantId returns a boolean if a field has been set.

### SetOriginalTenantIdNil

`func (o *AuditLog) SetOriginalTenantIdNil(b bool)`

 SetOriginalTenantIdNil sets the value for OriginalTenantId to be an explicit nil

### UnsetOriginalTenantId
`func (o *AuditLog) UnsetOriginalTenantId()`

UnsetOriginalTenantId ensures that no value is present for OriginalTenantId, not even an explicit nil
### GetOriginalTenantName

`func (o *AuditLog) GetOriginalTenantName() string`

GetOriginalTenantName returns the OriginalTenantName field if non-nil, zero value otherwise.

### GetOriginalTenantNameOk

`func (o *AuditLog) GetOriginalTenantNameOk() (*string, bool)`

GetOriginalTenantNameOk returns a tuple with the OriginalTenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTenantName

`func (o *AuditLog) SetOriginalTenantName(v string)`

SetOriginalTenantName sets OriginalTenantName field to given value.

### HasOriginalTenantName

`func (o *AuditLog) HasOriginalTenantName() bool`

HasOriginalTenantName returns a boolean if a field has been set.

### SetOriginalTenantNameNil

`func (o *AuditLog) SetOriginalTenantNameNil(b bool)`

 SetOriginalTenantNameNil sets the value for OriginalTenantName to be an explicit nil

### UnsetOriginalTenantName
`func (o *AuditLog) UnsetOriginalTenantName()`

UnsetOriginalTenantName ensures that no value is present for OriginalTenantName, not even an explicit nil
### GetPreviousRecord

`func (o *AuditLog) GetPreviousRecord() string`

GetPreviousRecord returns the PreviousRecord field if non-nil, zero value otherwise.

### GetPreviousRecordOk

`func (o *AuditLog) GetPreviousRecordOk() (*string, bool)`

GetPreviousRecordOk returns a tuple with the PreviousRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRecord

`func (o *AuditLog) SetPreviousRecord(v string)`

SetPreviousRecord sets PreviousRecord field to given value.

### HasPreviousRecord

`func (o *AuditLog) HasPreviousRecord() bool`

HasPreviousRecord returns a boolean if a field has been set.

### SetPreviousRecordNil

`func (o *AuditLog) SetPreviousRecordNil(b bool)`

 SetPreviousRecordNil sets the value for PreviousRecord to be an explicit nil

### UnsetPreviousRecord
`func (o *AuditLog) UnsetPreviousRecord()`

UnsetPreviousRecord ensures that no value is present for PreviousRecord, not even an explicit nil
### GetTenantId

`func (o *AuditLog) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AuditLog) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AuditLog) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AuditLog) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AuditLog) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AuditLog) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantName

`func (o *AuditLog) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *AuditLog) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *AuditLog) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *AuditLog) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *AuditLog) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *AuditLog) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
### GetTimestampUsecs

`func (o *AuditLog) GetTimestampUsecs() int64`

GetTimestampUsecs returns the TimestampUsecs field if non-nil, zero value otherwise.

### GetTimestampUsecsOk

`func (o *AuditLog) GetTimestampUsecsOk() (*int64, bool)`

GetTimestampUsecsOk returns a tuple with the TimestampUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUsecs

`func (o *AuditLog) SetTimestampUsecs(v int64)`

SetTimestampUsecs sets TimestampUsecs field to given value.

### HasTimestampUsecs

`func (o *AuditLog) HasTimestampUsecs() bool`

HasTimestampUsecs returns a boolean if a field has been set.

### SetTimestampUsecsNil

`func (o *AuditLog) SetTimestampUsecsNil(b bool)`

 SetTimestampUsecsNil sets the value for TimestampUsecs to be an explicit nil

### UnsetTimestampUsecs
`func (o *AuditLog) UnsetTimestampUsecs()`

UnsetTimestampUsecs ensures that no value is present for TimestampUsecs, not even an explicit nil
### GetUsername

`func (o *AuditLog) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuditLog) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuditLog) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AuditLog) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AuditLog) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AuditLog) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


