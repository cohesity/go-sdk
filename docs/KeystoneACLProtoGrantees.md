# KeystoneACLProtoGrantees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllUsers** | Pointer to **NullableBool** | This field indicates if all users are granted ACL permission. | [optional] 
**ProjectIdVec** | Pointer to **[]string** | This field holds a list of Keystone project ids whose users are granted ACL permission. | [optional] 
**ProjectUsersMap** | Pointer to [**[]KeystoneACLProtoGranteesProjectUsersMapEntry**](KeystoneACLProtoGranteesProjectUsersMapEntry.md) |  | [optional] 
**RoleNameVec** | Pointer to **[]string** | This field holds a list of Keystone roles for which any Keystone user with one (or more) of the roles on the project containing the swift container holding this KeystoneACLProto is granted ACL permission. | [optional] 
**UserIdVec** | Pointer to **[]string** | This field holds a list of keystone user ids who are granted ACL permission. | [optional] 

## Methods

### NewKeystoneACLProtoGrantees

`func NewKeystoneACLProtoGrantees() *KeystoneACLProtoGrantees`

NewKeystoneACLProtoGrantees instantiates a new KeystoneACLProtoGrantees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeystoneACLProtoGranteesWithDefaults

`func NewKeystoneACLProtoGranteesWithDefaults() *KeystoneACLProtoGrantees`

NewKeystoneACLProtoGranteesWithDefaults instantiates a new KeystoneACLProtoGrantees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllUsers

`func (o *KeystoneACLProtoGrantees) GetAllUsers() bool`

GetAllUsers returns the AllUsers field if non-nil, zero value otherwise.

### GetAllUsersOk

`func (o *KeystoneACLProtoGrantees) GetAllUsersOk() (*bool, bool)`

GetAllUsersOk returns a tuple with the AllUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllUsers

`func (o *KeystoneACLProtoGrantees) SetAllUsers(v bool)`

SetAllUsers sets AllUsers field to given value.

### HasAllUsers

`func (o *KeystoneACLProtoGrantees) HasAllUsers() bool`

HasAllUsers returns a boolean if a field has been set.

### SetAllUsersNil

`func (o *KeystoneACLProtoGrantees) SetAllUsersNil(b bool)`

 SetAllUsersNil sets the value for AllUsers to be an explicit nil

### UnsetAllUsers
`func (o *KeystoneACLProtoGrantees) UnsetAllUsers()`

UnsetAllUsers ensures that no value is present for AllUsers, not even an explicit nil
### GetProjectIdVec

`func (o *KeystoneACLProtoGrantees) GetProjectIdVec() []string`

GetProjectIdVec returns the ProjectIdVec field if non-nil, zero value otherwise.

### GetProjectIdVecOk

`func (o *KeystoneACLProtoGrantees) GetProjectIdVecOk() (*[]string, bool)`

GetProjectIdVecOk returns a tuple with the ProjectIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdVec

`func (o *KeystoneACLProtoGrantees) SetProjectIdVec(v []string)`

SetProjectIdVec sets ProjectIdVec field to given value.

### HasProjectIdVec

`func (o *KeystoneACLProtoGrantees) HasProjectIdVec() bool`

HasProjectIdVec returns a boolean if a field has been set.

### SetProjectIdVecNil

`func (o *KeystoneACLProtoGrantees) SetProjectIdVecNil(b bool)`

 SetProjectIdVecNil sets the value for ProjectIdVec to be an explicit nil

### UnsetProjectIdVec
`func (o *KeystoneACLProtoGrantees) UnsetProjectIdVec()`

UnsetProjectIdVec ensures that no value is present for ProjectIdVec, not even an explicit nil
### GetProjectUsersMap

`func (o *KeystoneACLProtoGrantees) GetProjectUsersMap() []KeystoneACLProtoGranteesProjectUsersMapEntry`

GetProjectUsersMap returns the ProjectUsersMap field if non-nil, zero value otherwise.

### GetProjectUsersMapOk

`func (o *KeystoneACLProtoGrantees) GetProjectUsersMapOk() (*[]KeystoneACLProtoGranteesProjectUsersMapEntry, bool)`

GetProjectUsersMapOk returns a tuple with the ProjectUsersMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUsersMap

`func (o *KeystoneACLProtoGrantees) SetProjectUsersMap(v []KeystoneACLProtoGranteesProjectUsersMapEntry)`

SetProjectUsersMap sets ProjectUsersMap field to given value.

### HasProjectUsersMap

`func (o *KeystoneACLProtoGrantees) HasProjectUsersMap() bool`

HasProjectUsersMap returns a boolean if a field has been set.

### SetProjectUsersMapNil

`func (o *KeystoneACLProtoGrantees) SetProjectUsersMapNil(b bool)`

 SetProjectUsersMapNil sets the value for ProjectUsersMap to be an explicit nil

### UnsetProjectUsersMap
`func (o *KeystoneACLProtoGrantees) UnsetProjectUsersMap()`

UnsetProjectUsersMap ensures that no value is present for ProjectUsersMap, not even an explicit nil
### GetRoleNameVec

`func (o *KeystoneACLProtoGrantees) GetRoleNameVec() []string`

GetRoleNameVec returns the RoleNameVec field if non-nil, zero value otherwise.

### GetRoleNameVecOk

`func (o *KeystoneACLProtoGrantees) GetRoleNameVecOk() (*[]string, bool)`

GetRoleNameVecOk returns a tuple with the RoleNameVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNameVec

`func (o *KeystoneACLProtoGrantees) SetRoleNameVec(v []string)`

SetRoleNameVec sets RoleNameVec field to given value.

### HasRoleNameVec

`func (o *KeystoneACLProtoGrantees) HasRoleNameVec() bool`

HasRoleNameVec returns a boolean if a field has been set.

### SetRoleNameVecNil

`func (o *KeystoneACLProtoGrantees) SetRoleNameVecNil(b bool)`

 SetRoleNameVecNil sets the value for RoleNameVec to be an explicit nil

### UnsetRoleNameVec
`func (o *KeystoneACLProtoGrantees) UnsetRoleNameVec()`

UnsetRoleNameVec ensures that no value is present for RoleNameVec, not even an explicit nil
### GetUserIdVec

`func (o *KeystoneACLProtoGrantees) GetUserIdVec() []string`

GetUserIdVec returns the UserIdVec field if non-nil, zero value otherwise.

### GetUserIdVecOk

`func (o *KeystoneACLProtoGrantees) GetUserIdVecOk() (*[]string, bool)`

GetUserIdVecOk returns a tuple with the UserIdVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdVec

`func (o *KeystoneACLProtoGrantees) SetUserIdVec(v []string)`

SetUserIdVec sets UserIdVec field to given value.

### HasUserIdVec

`func (o *KeystoneACLProtoGrantees) HasUserIdVec() bool`

HasUserIdVec returns a boolean if a field has been set.

### SetUserIdVecNil

`func (o *KeystoneACLProtoGrantees) SetUserIdVecNil(b bool)`

 SetUserIdVecNil sets the value for UserIdVec to be an explicit nil

### UnsetUserIdVec
`func (o *KeystoneACLProtoGrantees) UnsetUserIdVec()`

UnsetUserIdVec ensures that no value is present for UserIdVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


