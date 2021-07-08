# ACLProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonAcl** | Pointer to [**CommonACLProto**](CommonACLProto.md) |  | [optional] 
**GrantVec** | Pointer to [**[]ACLProtoGrant**](ACLProtoGrant.md) |  | [optional] 
**KeystoneAcl** | Pointer to [**KeystoneACLProto**](KeystoneACLProto.md) |  | [optional] 
**SwiftReadAcl** | Pointer to **NullableString** | Swift ACL strings. | [optional] 
**SwiftWriteAcl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewACLProto

`func NewACLProto() *ACLProto`

NewACLProto instantiates a new ACLProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLProtoWithDefaults

`func NewACLProtoWithDefaults() *ACLProto`

NewACLProtoWithDefaults instantiates a new ACLProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonAcl

`func (o *ACLProto) GetCommonAcl() CommonACLProto`

GetCommonAcl returns the CommonAcl field if non-nil, zero value otherwise.

### GetCommonAclOk

`func (o *ACLProto) GetCommonAclOk() (*CommonACLProto, bool)`

GetCommonAclOk returns a tuple with the CommonAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonAcl

`func (o *ACLProto) SetCommonAcl(v CommonACLProto)`

SetCommonAcl sets CommonAcl field to given value.

### HasCommonAcl

`func (o *ACLProto) HasCommonAcl() bool`

HasCommonAcl returns a boolean if a field has been set.

### GetGrantVec

`func (o *ACLProto) GetGrantVec() []ACLProtoGrant`

GetGrantVec returns the GrantVec field if non-nil, zero value otherwise.

### GetGrantVecOk

`func (o *ACLProto) GetGrantVecOk() (*[]ACLProtoGrant, bool)`

GetGrantVecOk returns a tuple with the GrantVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantVec

`func (o *ACLProto) SetGrantVec(v []ACLProtoGrant)`

SetGrantVec sets GrantVec field to given value.

### HasGrantVec

`func (o *ACLProto) HasGrantVec() bool`

HasGrantVec returns a boolean if a field has been set.

### SetGrantVecNil

`func (o *ACLProto) SetGrantVecNil(b bool)`

 SetGrantVecNil sets the value for GrantVec to be an explicit nil

### UnsetGrantVec
`func (o *ACLProto) UnsetGrantVec()`

UnsetGrantVec ensures that no value is present for GrantVec, not even an explicit nil
### GetKeystoneAcl

`func (o *ACLProto) GetKeystoneAcl() KeystoneACLProto`

GetKeystoneAcl returns the KeystoneAcl field if non-nil, zero value otherwise.

### GetKeystoneAclOk

`func (o *ACLProto) GetKeystoneAclOk() (*KeystoneACLProto, bool)`

GetKeystoneAclOk returns a tuple with the KeystoneAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoneAcl

`func (o *ACLProto) SetKeystoneAcl(v KeystoneACLProto)`

SetKeystoneAcl sets KeystoneAcl field to given value.

### HasKeystoneAcl

`func (o *ACLProto) HasKeystoneAcl() bool`

HasKeystoneAcl returns a boolean if a field has been set.

### GetSwiftReadAcl

`func (o *ACLProto) GetSwiftReadAcl() string`

GetSwiftReadAcl returns the SwiftReadAcl field if non-nil, zero value otherwise.

### GetSwiftReadAclOk

`func (o *ACLProto) GetSwiftReadAclOk() (*string, bool)`

GetSwiftReadAclOk returns a tuple with the SwiftReadAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftReadAcl

`func (o *ACLProto) SetSwiftReadAcl(v string)`

SetSwiftReadAcl sets SwiftReadAcl field to given value.

### HasSwiftReadAcl

`func (o *ACLProto) HasSwiftReadAcl() bool`

HasSwiftReadAcl returns a boolean if a field has been set.

### SetSwiftReadAclNil

`func (o *ACLProto) SetSwiftReadAclNil(b bool)`

 SetSwiftReadAclNil sets the value for SwiftReadAcl to be an explicit nil

### UnsetSwiftReadAcl
`func (o *ACLProto) UnsetSwiftReadAcl()`

UnsetSwiftReadAcl ensures that no value is present for SwiftReadAcl, not even an explicit nil
### GetSwiftWriteAcl

`func (o *ACLProto) GetSwiftWriteAcl() string`

GetSwiftWriteAcl returns the SwiftWriteAcl field if non-nil, zero value otherwise.

### GetSwiftWriteAclOk

`func (o *ACLProto) GetSwiftWriteAclOk() (*string, bool)`

GetSwiftWriteAclOk returns a tuple with the SwiftWriteAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftWriteAcl

`func (o *ACLProto) SetSwiftWriteAcl(v string)`

SetSwiftWriteAcl sets SwiftWriteAcl field to given value.

### HasSwiftWriteAcl

`func (o *ACLProto) HasSwiftWriteAcl() bool`

HasSwiftWriteAcl returns a boolean if a field has been set.

### SetSwiftWriteAclNil

`func (o *ACLProto) SetSwiftWriteAclNil(b bool)`

 SetSwiftWriteAclNil sets the value for SwiftWriteAcl to be an explicit nil

### UnsetSwiftWriteAcl
`func (o *ACLProto) UnsetSwiftWriteAcl()`

UnsetSwiftWriteAcl ensures that no value is present for SwiftWriteAcl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


