# EntityPermissionInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableInt64** | Specifies the entity id. | [optional] 
**Groups** | Pointer to [**[]GroupInfo**](GroupInfo.md) | Specifies groups that have access to entity in case of restricted user. | [optional] 
**IsInferred** | Pointer to **NullableBool** | Specifies whether the Entity Permission Information is inferred or not. For example, SQL application hosted over vCenter will have inferred entity permission information. | [optional] 
**Tenant** | Pointer to [**TenantInfo**](TenantInfo.md) |  | [optional] 
**Users** | Pointer to [**[]UserInfo**](UserInfo.md) | Specifies users that have access to entity in case of restricted user. | [optional] 

## Methods

### NewEntityPermissionInformation

`func NewEntityPermissionInformation() *EntityPermissionInformation`

NewEntityPermissionInformation instantiates a new EntityPermissionInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPermissionInformationWithDefaults

`func NewEntityPermissionInformationWithDefaults() *EntityPermissionInformation`

NewEntityPermissionInformationWithDefaults instantiates a new EntityPermissionInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *EntityPermissionInformation) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityPermissionInformation) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityPermissionInformation) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *EntityPermissionInformation) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *EntityPermissionInformation) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *EntityPermissionInformation) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetGroups

`func (o *EntityPermissionInformation) GetGroups() []GroupInfo`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EntityPermissionInformation) GetGroupsOk() (*[]GroupInfo, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EntityPermissionInformation) SetGroups(v []GroupInfo)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EntityPermissionInformation) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *EntityPermissionInformation) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *EntityPermissionInformation) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetIsInferred

`func (o *EntityPermissionInformation) GetIsInferred() bool`

GetIsInferred returns the IsInferred field if non-nil, zero value otherwise.

### GetIsInferredOk

`func (o *EntityPermissionInformation) GetIsInferredOk() (*bool, bool)`

GetIsInferredOk returns a tuple with the IsInferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInferred

`func (o *EntityPermissionInformation) SetIsInferred(v bool)`

SetIsInferred sets IsInferred field to given value.

### HasIsInferred

`func (o *EntityPermissionInformation) HasIsInferred() bool`

HasIsInferred returns a boolean if a field has been set.

### SetIsInferredNil

`func (o *EntityPermissionInformation) SetIsInferredNil(b bool)`

 SetIsInferredNil sets the value for IsInferred to be an explicit nil

### UnsetIsInferred
`func (o *EntityPermissionInformation) UnsetIsInferred()`

UnsetIsInferred ensures that no value is present for IsInferred, not even an explicit nil
### GetTenant

`func (o *EntityPermissionInformation) GetTenant() TenantInfo`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *EntityPermissionInformation) GetTenantOk() (*TenantInfo, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *EntityPermissionInformation) SetTenant(v TenantInfo)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *EntityPermissionInformation) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUsers

`func (o *EntityPermissionInformation) GetUsers() []UserInfo`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *EntityPermissionInformation) GetUsersOk() (*[]UserInfo, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *EntityPermissionInformation) SetUsers(v []UserInfo)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *EntityPermissionInformation) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *EntityPermissionInformation) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *EntityPermissionInformation) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


