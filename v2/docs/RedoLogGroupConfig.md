# RedoLogGroupConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupMembers** | Pointer to **[]string** | Specifies list of members of this redo log group. | [optional] 
**MemberPrefix** | Pointer to **NullableString** | Specifies Log member name prefix. | [optional] 
**NumGroups** | Pointer to **NullableInt32** | Specifies no. of redo log groups. | [optional] 
**SizeMBytes** | Pointer to **NullableInt32** | Specifies Size of the member in MB. | [optional] 

## Methods

### NewRedoLogGroupConfig

`func NewRedoLogGroupConfig() *RedoLogGroupConfig`

NewRedoLogGroupConfig instantiates a new RedoLogGroupConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedoLogGroupConfigWithDefaults

`func NewRedoLogGroupConfigWithDefaults() *RedoLogGroupConfig`

NewRedoLogGroupConfigWithDefaults instantiates a new RedoLogGroupConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupMembers

`func (o *RedoLogGroupConfig) GetGroupMembers() []string`

GetGroupMembers returns the GroupMembers field if non-nil, zero value otherwise.

### GetGroupMembersOk

`func (o *RedoLogGroupConfig) GetGroupMembersOk() (*[]string, bool)`

GetGroupMembersOk returns a tuple with the GroupMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembers

`func (o *RedoLogGroupConfig) SetGroupMembers(v []string)`

SetGroupMembers sets GroupMembers field to given value.

### HasGroupMembers

`func (o *RedoLogGroupConfig) HasGroupMembers() bool`

HasGroupMembers returns a boolean if a field has been set.

### SetGroupMembersNil

`func (o *RedoLogGroupConfig) SetGroupMembersNil(b bool)`

 SetGroupMembersNil sets the value for GroupMembers to be an explicit nil

### UnsetGroupMembers
`func (o *RedoLogGroupConfig) UnsetGroupMembers()`

UnsetGroupMembers ensures that no value is present for GroupMembers, not even an explicit nil
### GetMemberPrefix

`func (o *RedoLogGroupConfig) GetMemberPrefix() string`

GetMemberPrefix returns the MemberPrefix field if non-nil, zero value otherwise.

### GetMemberPrefixOk

`func (o *RedoLogGroupConfig) GetMemberPrefixOk() (*string, bool)`

GetMemberPrefixOk returns a tuple with the MemberPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPrefix

`func (o *RedoLogGroupConfig) SetMemberPrefix(v string)`

SetMemberPrefix sets MemberPrefix field to given value.

### HasMemberPrefix

`func (o *RedoLogGroupConfig) HasMemberPrefix() bool`

HasMemberPrefix returns a boolean if a field has been set.

### SetMemberPrefixNil

`func (o *RedoLogGroupConfig) SetMemberPrefixNil(b bool)`

 SetMemberPrefixNil sets the value for MemberPrefix to be an explicit nil

### UnsetMemberPrefix
`func (o *RedoLogGroupConfig) UnsetMemberPrefix()`

UnsetMemberPrefix ensures that no value is present for MemberPrefix, not even an explicit nil
### GetNumGroups

`func (o *RedoLogGroupConfig) GetNumGroups() int32`

GetNumGroups returns the NumGroups field if non-nil, zero value otherwise.

### GetNumGroupsOk

`func (o *RedoLogGroupConfig) GetNumGroupsOk() (*int32, bool)`

GetNumGroupsOk returns a tuple with the NumGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGroups

`func (o *RedoLogGroupConfig) SetNumGroups(v int32)`

SetNumGroups sets NumGroups field to given value.

### HasNumGroups

`func (o *RedoLogGroupConfig) HasNumGroups() bool`

HasNumGroups returns a boolean if a field has been set.

### SetNumGroupsNil

`func (o *RedoLogGroupConfig) SetNumGroupsNil(b bool)`

 SetNumGroupsNil sets the value for NumGroups to be an explicit nil

### UnsetNumGroups
`func (o *RedoLogGroupConfig) UnsetNumGroups()`

UnsetNumGroups ensures that no value is present for NumGroups, not even an explicit nil
### GetSizeMBytes

`func (o *RedoLogGroupConfig) GetSizeMBytes() int32`

GetSizeMBytes returns the SizeMBytes field if non-nil, zero value otherwise.

### GetSizeMBytesOk

`func (o *RedoLogGroupConfig) GetSizeMBytesOk() (*int32, bool)`

GetSizeMBytesOk returns a tuple with the SizeMBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMBytes

`func (o *RedoLogGroupConfig) SetSizeMBytes(v int32)`

SetSizeMBytes sets SizeMBytes field to given value.

### HasSizeMBytes

`func (o *RedoLogGroupConfig) HasSizeMBytes() bool`

HasSizeMBytes returns a boolean if a field has been set.

### SetSizeMBytesNil

`func (o *RedoLogGroupConfig) SetSizeMBytesNil(b bool)`

 SetSizeMBytesNil sets the value for SizeMBytes to be an explicit nil

### UnsetSizeMBytes
`func (o *RedoLogGroupConfig) UnsetSizeMBytes()`

UnsetSizeMBytes ensures that no value is present for SizeMBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


