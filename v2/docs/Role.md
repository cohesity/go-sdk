# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the Role name. | 
**Description** | Pointer to **NullableString** | Specifies the description message for the Role. | [optional] 
**Privileges** | **[]string** | Specifies the list of Privileges of the Role. | 
**CreatedTimestampMsecs** | Pointer to **NullableInt64** | Specifies the timestamp when the Role is created in milliseconds. | [optional] 
**IsUserCreatedRole** | Pointer to **NullableBool** | Specifies if the Role is created by user. | [optional] 
**Label** | Pointer to **NullableString** | Specifies the Role label. | [optional] 
**LastUpdatedTimestampMsecs** | Pointer to **NullableInt64** | Specifies the timestamp when the Role is last updated in milliseconds. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the list of tenant ids who have access to this Role. | [optional] 

## Methods

### NewRole

`func NewRole(name NullableString, privileges []string, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Role) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Role) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Role) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Role) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrivileges

`func (o *Role) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *Role) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *Role) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.


### SetPrivilegesNil

`func (o *Role) SetPrivilegesNil(b bool)`

 SetPrivilegesNil sets the value for Privileges to be an explicit nil

### UnsetPrivileges
`func (o *Role) UnsetPrivileges()`

UnsetPrivileges ensures that no value is present for Privileges, not even an explicit nil
### GetCreatedTimestampMsecs

`func (o *Role) GetCreatedTimestampMsecs() int64`

GetCreatedTimestampMsecs returns the CreatedTimestampMsecs field if non-nil, zero value otherwise.

### GetCreatedTimestampMsecsOk

`func (o *Role) GetCreatedTimestampMsecsOk() (*int64, bool)`

GetCreatedTimestampMsecsOk returns a tuple with the CreatedTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestampMsecs

`func (o *Role) SetCreatedTimestampMsecs(v int64)`

SetCreatedTimestampMsecs sets CreatedTimestampMsecs field to given value.

### HasCreatedTimestampMsecs

`func (o *Role) HasCreatedTimestampMsecs() bool`

HasCreatedTimestampMsecs returns a boolean if a field has been set.

### SetCreatedTimestampMsecsNil

`func (o *Role) SetCreatedTimestampMsecsNil(b bool)`

 SetCreatedTimestampMsecsNil sets the value for CreatedTimestampMsecs to be an explicit nil

### UnsetCreatedTimestampMsecs
`func (o *Role) UnsetCreatedTimestampMsecs()`

UnsetCreatedTimestampMsecs ensures that no value is present for CreatedTimestampMsecs, not even an explicit nil
### GetIsUserCreatedRole

`func (o *Role) GetIsUserCreatedRole() bool`

GetIsUserCreatedRole returns the IsUserCreatedRole field if non-nil, zero value otherwise.

### GetIsUserCreatedRoleOk

`func (o *Role) GetIsUserCreatedRoleOk() (*bool, bool)`

GetIsUserCreatedRoleOk returns a tuple with the IsUserCreatedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserCreatedRole

`func (o *Role) SetIsUserCreatedRole(v bool)`

SetIsUserCreatedRole sets IsUserCreatedRole field to given value.

### HasIsUserCreatedRole

`func (o *Role) HasIsUserCreatedRole() bool`

HasIsUserCreatedRole returns a boolean if a field has been set.

### SetIsUserCreatedRoleNil

`func (o *Role) SetIsUserCreatedRoleNil(b bool)`

 SetIsUserCreatedRoleNil sets the value for IsUserCreatedRole to be an explicit nil

### UnsetIsUserCreatedRole
`func (o *Role) UnsetIsUserCreatedRole()`

UnsetIsUserCreatedRole ensures that no value is present for IsUserCreatedRole, not even an explicit nil
### GetLabel

`func (o *Role) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Role) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Role) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Role) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *Role) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *Role) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetLastUpdatedTimestampMsecs

`func (o *Role) GetLastUpdatedTimestampMsecs() int64`

GetLastUpdatedTimestampMsecs returns the LastUpdatedTimestampMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampMsecsOk

`func (o *Role) GetLastUpdatedTimestampMsecsOk() (*int64, bool)`

GetLastUpdatedTimestampMsecsOk returns a tuple with the LastUpdatedTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestampMsecs

`func (o *Role) SetLastUpdatedTimestampMsecs(v int64)`

SetLastUpdatedTimestampMsecs sets LastUpdatedTimestampMsecs field to given value.

### HasLastUpdatedTimestampMsecs

`func (o *Role) HasLastUpdatedTimestampMsecs() bool`

HasLastUpdatedTimestampMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimestampMsecsNil

`func (o *Role) SetLastUpdatedTimestampMsecsNil(b bool)`

 SetLastUpdatedTimestampMsecsNil sets the value for LastUpdatedTimestampMsecs to be an explicit nil

### UnsetLastUpdatedTimestampMsecs
`func (o *Role) UnsetLastUpdatedTimestampMsecs()`

UnsetLastUpdatedTimestampMsecs ensures that no value is present for LastUpdatedTimestampMsecs, not even an explicit nil
### GetTenantIds

`func (o *Role) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *Role) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *Role) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *Role) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *Role) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *Role) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


