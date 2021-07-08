# OracleDBConfigRedoLogGroupConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupMemberVec** | Pointer to **[]string** | List of members of this redo log group. | [optional] 
**MemberPrefix** | Pointer to **NullableString** | Log member name prefix. | [optional] 
**NumGroups** | Pointer to **NullableInt32** | Number of redo log groups. | [optional] 
**SizeMb** | Pointer to **NullableInt32** | Size of the member in MB. | [optional] 

## Methods

### NewOracleDBConfigRedoLogGroupConf

`func NewOracleDBConfigRedoLogGroupConf() *OracleDBConfigRedoLogGroupConf`

NewOracleDBConfigRedoLogGroupConf instantiates a new OracleDBConfigRedoLogGroupConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleDBConfigRedoLogGroupConfWithDefaults

`func NewOracleDBConfigRedoLogGroupConfWithDefaults() *OracleDBConfigRedoLogGroupConf`

NewOracleDBConfigRedoLogGroupConfWithDefaults instantiates a new OracleDBConfigRedoLogGroupConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupMemberVec

`func (o *OracleDBConfigRedoLogGroupConf) GetGroupMemberVec() []string`

GetGroupMemberVec returns the GroupMemberVec field if non-nil, zero value otherwise.

### GetGroupMemberVecOk

`func (o *OracleDBConfigRedoLogGroupConf) GetGroupMemberVecOk() (*[]string, bool)`

GetGroupMemberVecOk returns a tuple with the GroupMemberVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberVec

`func (o *OracleDBConfigRedoLogGroupConf) SetGroupMemberVec(v []string)`

SetGroupMemberVec sets GroupMemberVec field to given value.

### HasGroupMemberVec

`func (o *OracleDBConfigRedoLogGroupConf) HasGroupMemberVec() bool`

HasGroupMemberVec returns a boolean if a field has been set.

### SetGroupMemberVecNil

`func (o *OracleDBConfigRedoLogGroupConf) SetGroupMemberVecNil(b bool)`

 SetGroupMemberVecNil sets the value for GroupMemberVec to be an explicit nil

### UnsetGroupMemberVec
`func (o *OracleDBConfigRedoLogGroupConf) UnsetGroupMemberVec()`

UnsetGroupMemberVec ensures that no value is present for GroupMemberVec, not even an explicit nil
### GetMemberPrefix

`func (o *OracleDBConfigRedoLogGroupConf) GetMemberPrefix() string`

GetMemberPrefix returns the MemberPrefix field if non-nil, zero value otherwise.

### GetMemberPrefixOk

`func (o *OracleDBConfigRedoLogGroupConf) GetMemberPrefixOk() (*string, bool)`

GetMemberPrefixOk returns a tuple with the MemberPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPrefix

`func (o *OracleDBConfigRedoLogGroupConf) SetMemberPrefix(v string)`

SetMemberPrefix sets MemberPrefix field to given value.

### HasMemberPrefix

`func (o *OracleDBConfigRedoLogGroupConf) HasMemberPrefix() bool`

HasMemberPrefix returns a boolean if a field has been set.

### SetMemberPrefixNil

`func (o *OracleDBConfigRedoLogGroupConf) SetMemberPrefixNil(b bool)`

 SetMemberPrefixNil sets the value for MemberPrefix to be an explicit nil

### UnsetMemberPrefix
`func (o *OracleDBConfigRedoLogGroupConf) UnsetMemberPrefix()`

UnsetMemberPrefix ensures that no value is present for MemberPrefix, not even an explicit nil
### GetNumGroups

`func (o *OracleDBConfigRedoLogGroupConf) GetNumGroups() int32`

GetNumGroups returns the NumGroups field if non-nil, zero value otherwise.

### GetNumGroupsOk

`func (o *OracleDBConfigRedoLogGroupConf) GetNumGroupsOk() (*int32, bool)`

GetNumGroupsOk returns a tuple with the NumGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGroups

`func (o *OracleDBConfigRedoLogGroupConf) SetNumGroups(v int32)`

SetNumGroups sets NumGroups field to given value.

### HasNumGroups

`func (o *OracleDBConfigRedoLogGroupConf) HasNumGroups() bool`

HasNumGroups returns a boolean if a field has been set.

### SetNumGroupsNil

`func (o *OracleDBConfigRedoLogGroupConf) SetNumGroupsNil(b bool)`

 SetNumGroupsNil sets the value for NumGroups to be an explicit nil

### UnsetNumGroups
`func (o *OracleDBConfigRedoLogGroupConf) UnsetNumGroups()`

UnsetNumGroups ensures that no value is present for NumGroups, not even an explicit nil
### GetSizeMb

`func (o *OracleDBConfigRedoLogGroupConf) GetSizeMb() int32`

GetSizeMb returns the SizeMb field if non-nil, zero value otherwise.

### GetSizeMbOk

`func (o *OracleDBConfigRedoLogGroupConf) GetSizeMbOk() (*int32, bool)`

GetSizeMbOk returns a tuple with the SizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMb

`func (o *OracleDBConfigRedoLogGroupConf) SetSizeMb(v int32)`

SetSizeMb sets SizeMb field to given value.

### HasSizeMb

`func (o *OracleDBConfigRedoLogGroupConf) HasSizeMb() bool`

HasSizeMb returns a boolean if a field has been set.

### SetSizeMbNil

`func (o *OracleDBConfigRedoLogGroupConf) SetSizeMbNil(b bool)`

 SetSizeMbNil sets the value for SizeMb to be an explicit nil

### UnsetSizeMb
`func (o *OracleDBConfigRedoLogGroupConf) UnsetSizeMb()`

UnsetSizeMb ensures that no value is present for SizeMb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


