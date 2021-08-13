# PermissionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | Pointer to **NullableInt64** | Specifies the id of the object. | [optional] 
**Users** | Pointer to [**[]User**](User.md) | Specifies the list of users which has the permissions to the object. | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) | Specifies the list of user groups which has permissions to the object. | [optional] 
**Tenant** | Pointer to [**Tenant**](Tenant.md) |  | [optional] 

## Methods

### NewPermissionInfo

`func NewPermissionInfo() *PermissionInfo`

NewPermissionInfo instantiates a new PermissionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionInfoWithDefaults

`func NewPermissionInfoWithDefaults() *PermissionInfo`

NewPermissionInfoWithDefaults instantiates a new PermissionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *PermissionInfo) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PermissionInfo) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PermissionInfo) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PermissionInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *PermissionInfo) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *PermissionInfo) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetUsers

`func (o *PermissionInfo) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PermissionInfo) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PermissionInfo) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PermissionInfo) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *PermissionInfo) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *PermissionInfo) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetGroups

`func (o *PermissionInfo) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PermissionInfo) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PermissionInfo) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PermissionInfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *PermissionInfo) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *PermissionInfo) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetTenant

`func (o *PermissionInfo) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PermissionInfo) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PermissionInfo) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PermissionInfo) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


