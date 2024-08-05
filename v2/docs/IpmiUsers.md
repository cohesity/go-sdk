# IpmiUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **NullableInt32** | Specifies the channel through which the IPMI interface communicates on the network. | [optional] 
**MaxUserId** | Pointer to **NullableInt32** | Specifies the maximum value of user id among all the users. | [optional] 
**UserList** | Pointer to [**[]IpmiUserInfo**](IpmiUserInfo.md) | Specifies the list of ipmi users with their permissions. | [optional] 

## Methods

### NewIpmiUsers

`func NewIpmiUsers() *IpmiUsers`

NewIpmiUsers instantiates a new IpmiUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiUsersWithDefaults

`func NewIpmiUsersWithDefaults() *IpmiUsers`

NewIpmiUsersWithDefaults instantiates a new IpmiUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *IpmiUsers) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *IpmiUsers) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *IpmiUsers) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *IpmiUsers) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *IpmiUsers) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *IpmiUsers) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetMaxUserId

`func (o *IpmiUsers) GetMaxUserId() int32`

GetMaxUserId returns the MaxUserId field if non-nil, zero value otherwise.

### GetMaxUserIdOk

`func (o *IpmiUsers) GetMaxUserIdOk() (*int32, bool)`

GetMaxUserIdOk returns a tuple with the MaxUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUserId

`func (o *IpmiUsers) SetMaxUserId(v int32)`

SetMaxUserId sets MaxUserId field to given value.

### HasMaxUserId

`func (o *IpmiUsers) HasMaxUserId() bool`

HasMaxUserId returns a boolean if a field has been set.

### SetMaxUserIdNil

`func (o *IpmiUsers) SetMaxUserIdNil(b bool)`

 SetMaxUserIdNil sets the value for MaxUserId to be an explicit nil

### UnsetMaxUserId
`func (o *IpmiUsers) UnsetMaxUserId()`

UnsetMaxUserId ensures that no value is present for MaxUserId, not even an explicit nil
### GetUserList

`func (o *IpmiUsers) GetUserList() []IpmiUserInfo`

GetUserList returns the UserList field if non-nil, zero value otherwise.

### GetUserListOk

`func (o *IpmiUsers) GetUserListOk() (*[]IpmiUserInfo, bool)`

GetUserListOk returns a tuple with the UserList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserList

`func (o *IpmiUsers) SetUserList(v []IpmiUserInfo)`

SetUserList sets UserList field to given value.

### HasUserList

`func (o *IpmiUsers) HasUserList() bool`

HasUserList returns a boolean if a field has been set.

### SetUserListNil

`func (o *IpmiUsers) SetUserListNil(b bool)`

 SetUserListNil sets the value for UserList to be an explicit nil

### UnsetUserList
`func (o *IpmiUsers) UnsetUserList()`

UnsetUserList ensures that no value is present for UserList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


