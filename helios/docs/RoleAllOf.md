# RoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **NullableString** | Specifies the Role label. | [optional] 
**CreatedTimestampMsecs** | Pointer to **NullableInt64** | Specifies the timestamp when the Role is created in milliseconds. | [optional] 
**LastUpdatedTimestampMsecs** | Pointer to **NullableInt64** | Specifies the timestamp when the Role is last updated in milliseconds. | [optional] 
**IsUserCreatedRole** | Pointer to **NullableBool** | Specifies if the Role is created by user. | [optional] 
**TenantIds** | Pointer to **[]string** | Specifies the list of tenant ids who have access to this Role. | [optional] 

## Methods

### NewRoleAllOf

`func NewRoleAllOf() *RoleAllOf`

NewRoleAllOf instantiates a new RoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAllOfWithDefaults

`func NewRoleAllOfWithDefaults() *RoleAllOf`

NewRoleAllOfWithDefaults instantiates a new RoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *RoleAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RoleAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RoleAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RoleAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *RoleAllOf) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *RoleAllOf) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetCreatedTimestampMsecs

`func (o *RoleAllOf) GetCreatedTimestampMsecs() int64`

GetCreatedTimestampMsecs returns the CreatedTimestampMsecs field if non-nil, zero value otherwise.

### GetCreatedTimestampMsecsOk

`func (o *RoleAllOf) GetCreatedTimestampMsecsOk() (*int64, bool)`

GetCreatedTimestampMsecsOk returns a tuple with the CreatedTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestampMsecs

`func (o *RoleAllOf) SetCreatedTimestampMsecs(v int64)`

SetCreatedTimestampMsecs sets CreatedTimestampMsecs field to given value.

### HasCreatedTimestampMsecs

`func (o *RoleAllOf) HasCreatedTimestampMsecs() bool`

HasCreatedTimestampMsecs returns a boolean if a field has been set.

### SetCreatedTimestampMsecsNil

`func (o *RoleAllOf) SetCreatedTimestampMsecsNil(b bool)`

 SetCreatedTimestampMsecsNil sets the value for CreatedTimestampMsecs to be an explicit nil

### UnsetCreatedTimestampMsecs
`func (o *RoleAllOf) UnsetCreatedTimestampMsecs()`

UnsetCreatedTimestampMsecs ensures that no value is present for CreatedTimestampMsecs, not even an explicit nil
### GetLastUpdatedTimestampMsecs

`func (o *RoleAllOf) GetLastUpdatedTimestampMsecs() int64`

GetLastUpdatedTimestampMsecs returns the LastUpdatedTimestampMsecs field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampMsecsOk

`func (o *RoleAllOf) GetLastUpdatedTimestampMsecsOk() (*int64, bool)`

GetLastUpdatedTimestampMsecsOk returns a tuple with the LastUpdatedTimestampMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestampMsecs

`func (o *RoleAllOf) SetLastUpdatedTimestampMsecs(v int64)`

SetLastUpdatedTimestampMsecs sets LastUpdatedTimestampMsecs field to given value.

### HasLastUpdatedTimestampMsecs

`func (o *RoleAllOf) HasLastUpdatedTimestampMsecs() bool`

HasLastUpdatedTimestampMsecs returns a boolean if a field has been set.

### SetLastUpdatedTimestampMsecsNil

`func (o *RoleAllOf) SetLastUpdatedTimestampMsecsNil(b bool)`

 SetLastUpdatedTimestampMsecsNil sets the value for LastUpdatedTimestampMsecs to be an explicit nil

### UnsetLastUpdatedTimestampMsecs
`func (o *RoleAllOf) UnsetLastUpdatedTimestampMsecs()`

UnsetLastUpdatedTimestampMsecs ensures that no value is present for LastUpdatedTimestampMsecs, not even an explicit nil
### GetIsUserCreatedRole

`func (o *RoleAllOf) GetIsUserCreatedRole() bool`

GetIsUserCreatedRole returns the IsUserCreatedRole field if non-nil, zero value otherwise.

### GetIsUserCreatedRoleOk

`func (o *RoleAllOf) GetIsUserCreatedRoleOk() (*bool, bool)`

GetIsUserCreatedRoleOk returns a tuple with the IsUserCreatedRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserCreatedRole

`func (o *RoleAllOf) SetIsUserCreatedRole(v bool)`

SetIsUserCreatedRole sets IsUserCreatedRole field to given value.

### HasIsUserCreatedRole

`func (o *RoleAllOf) HasIsUserCreatedRole() bool`

HasIsUserCreatedRole returns a boolean if a field has been set.

### SetIsUserCreatedRoleNil

`func (o *RoleAllOf) SetIsUserCreatedRoleNil(b bool)`

 SetIsUserCreatedRoleNil sets the value for IsUserCreatedRole to be an explicit nil

### UnsetIsUserCreatedRole
`func (o *RoleAllOf) UnsetIsUserCreatedRole()`

UnsetIsUserCreatedRole ensures that no value is present for IsUserCreatedRole, not even an explicit nil
### GetTenantIds

`func (o *RoleAllOf) GetTenantIds() []string`

GetTenantIds returns the TenantIds field if non-nil, zero value otherwise.

### GetTenantIdsOk

`func (o *RoleAllOf) GetTenantIdsOk() (*[]string, bool)`

GetTenantIdsOk returns a tuple with the TenantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIds

`func (o *RoleAllOf) SetTenantIds(v []string)`

SetTenantIds sets TenantIds field to given value.

### HasTenantIds

`func (o *RoleAllOf) HasTenantIds() bool`

HasTenantIds returns a boolean if a field has been set.

### SetTenantIdsNil

`func (o *RoleAllOf) SetTenantIdsNil(b bool)`

 SetTenantIdsNil sets the value for TenantIds to be an explicit nil

### UnsetTenantIds
`func (o *RoleAllOf) UnsetTenantIds()`

UnsetTenantIds ensures that no value is present for TenantIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


