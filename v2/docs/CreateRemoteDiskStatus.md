# CreateRemoteDiskStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Specifies the error message when creating remote disk fails. | [optional] 
**MountPath** | Pointer to **NullableString** | Specifies the NFS mount path of the remote disk. | [optional] 
**NodeId** | Pointer to **NullableInt64** | Specifies the node id of the disk. If not specified, the disk will be evenly distributed across all the nodes. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the creating status of this disk. | [optional] 
**Tier** | Pointer to **NullableString** | Specifies the tier of the disk | [optional] 

## Methods

### NewCreateRemoteDiskStatus

`func NewCreateRemoteDiskStatus() *CreateRemoteDiskStatus`

NewCreateRemoteDiskStatus instantiates a new CreateRemoteDiskStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRemoteDiskStatusWithDefaults

`func NewCreateRemoteDiskStatusWithDefaults() *CreateRemoteDiskStatus`

NewCreateRemoteDiskStatusWithDefaults instantiates a new CreateRemoteDiskStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateRemoteDiskStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateRemoteDiskStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateRemoteDiskStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateRemoteDiskStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CreateRemoteDiskStatus) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CreateRemoteDiskStatus) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMountPath

`func (o *CreateRemoteDiskStatus) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *CreateRemoteDiskStatus) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *CreateRemoteDiskStatus) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *CreateRemoteDiskStatus) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *CreateRemoteDiskStatus) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *CreateRemoteDiskStatus) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetNodeId

`func (o *CreateRemoteDiskStatus) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *CreateRemoteDiskStatus) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *CreateRemoteDiskStatus) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *CreateRemoteDiskStatus) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *CreateRemoteDiskStatus) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *CreateRemoteDiskStatus) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetStatus

`func (o *CreateRemoteDiskStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateRemoteDiskStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateRemoteDiskStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateRemoteDiskStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CreateRemoteDiskStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CreateRemoteDiskStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTier

`func (o *CreateRemoteDiskStatus) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CreateRemoteDiskStatus) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CreateRemoteDiskStatus) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *CreateRemoteDiskStatus) HasTier() bool`

HasTier returns a boolean if a field has been set.

### SetTierNil

`func (o *CreateRemoteDiskStatus) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *CreateRemoteDiskStatus) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


