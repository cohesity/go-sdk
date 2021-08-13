# CreateGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the Group name. | 
**Domain** | **NullableString** | Specifies the Group domain. | 
**Description** | Pointer to **NullableString** | Specifies the description of the Group. | [optional] 
**Roles** | Pointer to **[]string** | Specifies the Roles of the Group. | [optional] 
**Users** | Pointer to **[]string** | Specifies a list of Users who are member of this Group. | [optional] 
**Restricted** | Pointer to **NullableBool** | Specifies whether the Group is restricted. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies a list of tenant ids who can access this group. | [optional] 

## Methods

### NewCreateGroupParams

`func NewCreateGroupParams(name NullableString, domain NullableString, ) *CreateGroupParams`

NewCreateGroupParams instantiates a new CreateGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupParamsWithDefaults

`func NewCreateGroupParamsWithDefaults() *CreateGroupParams`

NewCreateGroupParamsWithDefaults instantiates a new CreateGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGroupParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateGroupParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateGroupParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDomain

`func (o *CreateGroupParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateGroupParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateGroupParams) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *CreateGroupParams) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *CreateGroupParams) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetDescription

`func (o *CreateGroupParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateGroupParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateGroupParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRoles

`func (o *CreateGroupParams) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CreateGroupParams) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CreateGroupParams) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CreateGroupParams) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *CreateGroupParams) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *CreateGroupParams) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetUsers

`func (o *CreateGroupParams) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CreateGroupParams) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CreateGroupParams) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CreateGroupParams) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *CreateGroupParams) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *CreateGroupParams) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetRestricted

`func (o *CreateGroupParams) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *CreateGroupParams) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *CreateGroupParams) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *CreateGroupParams) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### SetRestrictedNil

`func (o *CreateGroupParams) SetRestrictedNil(b bool)`

 SetRestrictedNil sets the value for Restricted to be an explicit nil

### UnsetRestricted
`func (o *CreateGroupParams) UnsetRestricted()`

UnsetRestricted ensures that no value is present for Restricted, not even an explicit nil
### GetTenantIds

`func (o *CreateGroupParams) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *CreateGroupParams) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *CreateGroupParams) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *CreateGroupParams) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *CreateGroupParams) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *CreateGroupParams) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


