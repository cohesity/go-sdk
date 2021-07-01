# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the group was created/added. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description of the group. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the domain of the group. | [optional] 
**LastUpdatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the epoch time in milliseconds when the group was last modified. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the group. | [optional] 
**Restricted** | Pointer to **NullableBool** | Whether the group is a restricted group. Users belonging to a restricted group can only view objects they have permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Array of Roles.  Specifies the Cohesity roles to associate with the group such as &#39;Admin&#39;, &#39;Ops&#39; or &#39;View&#39;. The Cohesity roles determine privileges on the Cohesity Cluster for all the users in this group. | [optional] 
**Sid** | Pointer to **NullableString** | Specifies the unique Security ID (SID) of the group. | [optional] 
**SmbPrincipals** | Pointer to [**[]SmbPrincipal**](SmbPrincipal.md) | Specifies the SMB principals. Principals will be added to this group only if IsSmbPrincipalOnly set to true. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the tenants to which the group belongs to. If not specified, session user&#39;s tenant id is assumed. | [optional] 
**Usernames** | Pointer to **[]string** | Specifies the username of users who are members of the group. This field is used only for local groups. | [optional] 
**Users** | Pointer to **[]string** | Specifies the SID of users who are members of the group. This field is used only for local groups. | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimeMsecs

`func (o *Group) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *Group) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *Group) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *Group) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *Group) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *Group) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Group) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Group) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomain

`func (o *Group) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Group) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Group) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Group) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *Group) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *Group) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetLastUpdatedTimeMsecs

`func (o *Group) GetLastUpdatedTimeMsecs() int64`

GetLastUpdatedTimeMsecs returns the LastUpdatedTimeMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimeMsecsOk

`func (o *Group) GetLastUpdatedTimeMsecsOk() (*int64, bool)`

GetLastUpdatedTimeMsecsOk returns a tuple with the LastUpdatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimeMsecs

`func (o *Group) SetLastUpdatedTimeMsecs(v int64)`

SetLastUpdatedTimeMsecs sets LastUpdatedTimeMsecs field to given value.

### HasLastUpdatedTimeMsecs

`func (o *Group) HasLastUpdatedTimeMsecs() bool`

HasLastUpdatedTimeMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimeMsecsNil

`func (o *Group) SetLastUpdatedTimeMsecsNil(b bool)`

 SetLastUpdatedTimeMsecsNil sets the value for LastUpdatedTimeMsecs to be an explicit nil

### UnsetLastUpdatedTimeMsecs
`func (o *Group) UnsetLastUpdatedTimeMsecs()`

UnsetLastUpdatedTimeMsecs ensures that no value is present for LastUpdatedTimeMsecs, not even an explicit nil
### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Group) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Group) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRestricted

`func (o *Group) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *Group) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *Group) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *Group) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *Group) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *Group) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *Group) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Group) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Group) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Group) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *Group) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *Group) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSid

`func (o *Group) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *Group) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *Group) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *Group) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *Group) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *Group) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSmbPrincipals

`func (o *Group) GetSmbPrincipals() []SmbPrincipal`

GetSmbPrincipals returns the SmbPrincipals field if non-nil, zero value otherwise.

### GetSmbPrincipalsOk

`func (o *Group) GetSmbPrincipalsOk() (*[]SmbPrincipal, bool)`

GetSmbPrincipalsOk returns a tuple with the SmbPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPrincipals

`func (o *Group) SetSmbPrincipals(v []SmbPrincipal)`

SetSmbPrincipals sets SmbPrincipals field to given value.

### HasSmbPrincipals

`func (o *Group) HasSmbPrincipals() bool`

HasSmbPrincipals returns a boolean if a field has been set.

### SetSmbPrincipalsNil

`func (o *Group) SetSmbPrincipalsNil(b bool)`

 SetSmbPrincipalsNil sets the value for SmbPrincipals to be an explicit nil

### UnsetSmbPrincipals
`func (o *Group) UnsetSmbPrincipals()`

UnsetSmbPrincipals ensures that no value is present for SmbPrincipals, not even an explicit nil
### GetTenantIds

`func (o *Group) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *Group) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *Group) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *Group) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *Group) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *Group) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil
### GetUsernames

`func (o *Group) GetUsernames() []string`

GetUsernames returns the Usernames field if non-nil, zero value otherwise.

### GetUsernamesOk

`func (o *Group) GetUsernamesOk() (*[]string, bool)`

GetUsernamesOk returns a tuple with the Usernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernames

`func (o *Group) SetUsernames(v []string)`

SetUsernames sets Usernames field to given value.

### HasUsernames

`func (o *Group) HasUsernames() bool`

HasUsernames returns a boolean if a field has been set.

### SetUsernamesNil

`func (o *Group) SetUsernamesNil(b bool)`

 SetUsernamesNil sets the value for Usernames to be an explicit nil

### UnsetUsernames
`func (o *Group) UnsetUsernames()`

UnsetUsernames ensures that no value is present for Usernames, not even an explicit nil
### GetUsers

`func (o *Group) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Group) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Group) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Group) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *Group) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *Group) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


