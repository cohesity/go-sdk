# IpmiUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallIn** | Pointer to **NullableBool** | Specifies whether the user is allowed to initiate IPMI communication. | [optional] 
**Id** | Pointer to **NullableInt32** | Specifies the id for the ipmi user. | [optional] 
**IpmiMsg** | Pointer to **NullableBool** | Specifies whether IPMI messaging is enabled for this user. | [optional] 
**LinkAuth** | Pointer to **NullableBool** | Specifies whether link-level authentication is required for this user. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the ipmi user. | [optional] 
**PrivilegeLevel** | Pointer to **NullableInt32** | Specifies the privilege level assigned for this user. | [optional] 

## Methods

### NewIpmiUserInfo

`func NewIpmiUserInfo() *IpmiUserInfo`

NewIpmiUserInfo instantiates a new IpmiUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiUserInfoWithDefaults

`func NewIpmiUserInfoWithDefaults() *IpmiUserInfo`

NewIpmiUserInfoWithDefaults instantiates a new IpmiUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallIn

`func (o *IpmiUserInfo) GetCallIn() bool`

GetCallIn returns the CallIn field if non-nil, zero value otherwise.

### GetCallInOk

`func (o *IpmiUserInfo) GetCallInOk() (*bool, bool)`

GetCallInOk returns a tuple with the CallIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallIn

`func (o *IpmiUserInfo) SetCallIn(v bool)`

SetCallIn sets CallIn field to given value.

### HasCallIn

`func (o *IpmiUserInfo) HasCallIn() bool`

HasCallIn returns a boolean if a field has been set.

### SetCallInNil

`func (o *IpmiUserInfo) SetCallInNil(b bool)`

 SetCallInNil sets the value for CallIn to be an explicit nil

### UnsetCallIn
`func (o *IpmiUserInfo) UnsetCallIn()`

UnsetCallIn ensures that no value is present for CallIn, not even an explicit nil
### GetId

`func (o *IpmiUserInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpmiUserInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpmiUserInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IpmiUserInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IpmiUserInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IpmiUserInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIpmiMsg

`func (o *IpmiUserInfo) GetIpmiMsg() bool`

GetIpmiMsg returns the IpmiMsg field if non-nil, zero value otherwise.

### GetIpmiMsgOk

`func (o *IpmiUserInfo) GetIpmiMsgOk() (*bool, bool)`

GetIpmiMsgOk returns a tuple with the IpmiMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiMsg

`func (o *IpmiUserInfo) SetIpmiMsg(v bool)`

SetIpmiMsg sets IpmiMsg field to given value.

### HasIpmiMsg

`func (o *IpmiUserInfo) HasIpmiMsg() bool`

HasIpmiMsg returns a boolean if a field has been set.

### SetIpmiMsgNil

`func (o *IpmiUserInfo) SetIpmiMsgNil(b bool)`

 SetIpmiMsgNil sets the value for IpmiMsg to be an explicit nil

### UnsetIpmiMsg
`func (o *IpmiUserInfo) UnsetIpmiMsg()`

UnsetIpmiMsg ensures that no value is present for IpmiMsg, not even an explicit nil
### GetLinkAuth

`func (o *IpmiUserInfo) GetLinkAuth() bool`

GetLinkAuth returns the LinkAuth field if non-nil, zero value otherwise.

### GetLinkAuthOk

`func (o *IpmiUserInfo) GetLinkAuthOk() (*bool, bool)`

GetLinkAuthOk returns a tuple with the LinkAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAuth

`func (o *IpmiUserInfo) SetLinkAuth(v bool)`

SetLinkAuth sets LinkAuth field to given value.

### HasLinkAuth

`func (o *IpmiUserInfo) HasLinkAuth() bool`

HasLinkAuth returns a boolean if a field has been set.

### SetLinkAuthNil

`func (o *IpmiUserInfo) SetLinkAuthNil(b bool)`

 SetLinkAuthNil sets the value for LinkAuth to be an explicit nil

### UnsetLinkAuth
`func (o *IpmiUserInfo) UnsetLinkAuth()`

UnsetLinkAuth ensures that no value is present for LinkAuth, not even an explicit nil
### GetName

`func (o *IpmiUserInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpmiUserInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpmiUserInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpmiUserInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IpmiUserInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IpmiUserInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPrivilegeLevel

`func (o *IpmiUserInfo) GetPrivilegeLevel() int32`

GetPrivilegeLevel returns the PrivilegeLevel field if non-nil, zero value otherwise.

### GetPrivilegeLevelOk

`func (o *IpmiUserInfo) GetPrivilegeLevelOk() (*int32, bool)`

GetPrivilegeLevelOk returns a tuple with the PrivilegeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeLevel

`func (o *IpmiUserInfo) SetPrivilegeLevel(v int32)`

SetPrivilegeLevel sets PrivilegeLevel field to given value.

### HasPrivilegeLevel

`func (o *IpmiUserInfo) HasPrivilegeLevel() bool`

HasPrivilegeLevel returns a boolean if a field has been set.

### SetPrivilegeLevelNil

`func (o *IpmiUserInfo) SetPrivilegeLevelNil(b bool)`

 SetPrivilegeLevelNil sets the value for PrivilegeLevel to be an explicit nil

### UnsetPrivilegeLevel
`func (o *IpmiUserInfo) UnsetPrivilegeLevel()`

UnsetPrivilegeLevel ensures that no value is present for PrivilegeLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


