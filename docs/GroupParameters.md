# GroupParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Specifies a description of the group. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the domain of the group. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the group. | [optional] 
**Restricted** | Pointer to **NullableBool** | Whether the group is a restricted group. Users belonging to a restricted group can only view objects they have permissions to. | [optional] 
**Roles** | Pointer to **[]string** | Array of Roles.  Specifies the Cohesity roles to associate with the group such as &#39;Admin&#39;, &#39;Ops&#39; or &#39;View&#39;. The Cohesity roles determine privileges on the Cohesity Cluster for all the users in this group. | [optional] 
**SmbPrincipals** | Pointer to [**[]SmbPrincipal**](SmbPrincipal.md) | Specifies the SMB principals. Principals will be added to this group only if IsSmbPrincipalOnly set to true. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the tenants to which the group belongs to. If not specified, session user&#39;s tenant id is assumed. | [optional] 
**Users** | Pointer to **[]string** | Specifies the SID of users who are members of the group. This field is used only for local groups. | [optional] 

## Methods

### NewGroupParameters

`func NewGroupParameters() *GroupParameters`

NewGroupParameters instantiates a new GroupParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupParametersWithDefaults

`func NewGroupParametersWithDefaults() *GroupParameters`

NewGroupParametersWithDefaults instantiates a new GroupParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GroupParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GroupParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomain

`func (o *GroupParameters) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GroupParameters) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GroupParameters) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GroupParameters) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *GroupParameters) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *GroupParameters) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetName

`func (o *GroupParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GroupParameters) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GroupParameters) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRestricted

`func (o *GroupParameters) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *GroupParameters) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *GroupParameters) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *GroupParameters) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *GroupParameters) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *GroupParameters) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetRoles

`func (o *GroupParameters) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupParameters) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupParameters) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GroupParameters) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *GroupParameters) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *GroupParameters) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSmbPrincipals

`func (o *GroupParameters) GetSmbPrincipals() []SmbPrincipal`

GetSmbPrincipals returns the SmbPrincipals field if non-nil, zero value otherwise.

### GetSmbPrincipalsOk

`func (o *GroupParameters) GetSmbPrincipalsOk() (*[]SmbPrincipal, bool)`

GetSmbPrincipalsOk returns a tuple with the SmbPrincipals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbPrincipals

`func (o *GroupParameters) SetSmbPrincipals(v []SmbPrincipal)`

SetSmbPrincipals sets SmbPrincipals field to given value.

### HasSmbPrincipals

`func (o *GroupParameters) HasSmbPrincipals() bool`

HasSmbPrincipals returns a boolean if a field has been set.

### SetSmbPrincipalsNil

`func (o *GroupParameters) SetSmbPrincipalsNil(b bool)`

 SetSmbPrincipalsNil sets the value for SmbPrincipals to be an explicit nil

### UnsetSmbPrincipals
`func (o *GroupParameters) UnsetSmbPrincipals()`

UnsetSmbPrincipals ensures that no value is present for SmbPrincipals, not even an explicit nil
### GetTenantIds

`func (o *GroupParameters) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *GroupParameters) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *GroupParameters) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *GroupParameters) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *GroupParameters) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *GroupParameters) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil
### GetUsers

`func (o *GroupParameters) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupParameters) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupParameters) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupParameters) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *GroupParameters) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *GroupParameters) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


