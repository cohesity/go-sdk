# RemoteDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the disk id. | [optional] [readonly] 
**MountPath** | **NullableString** | Specifies the NFS mount path of the remote disk. | 
**NodeId** | Pointer to **NullableInt64** | Specifies the node id of the disk. If not specified, the disk will be evenly distributed across all the nodes. | [optional] 
**Tier** | **NullableString** | Specifies the tier of the disk | 
**CapacityBytes** | Pointer to **NullableInt64** | Specifies the logical capacity of the disk in bytes. | [optional] [readonly] 
**UsedCapacityBytes** | Pointer to **NullableInt64** | Specifies the logical used capacity of remote disk in bytes. | [optional] [readonly] 
**Status** | Pointer to **NullableString** | Specifies the status of a remote disk. | [optional] [readonly] 
**FileSystemName** | Pointer to **NullableString** | Specifies the name of filesystem on remote storage. | [optional] 
**DataVip** | Pointer to **NullableString** | Specifies the data vip used to mount the filesystem. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies ip address of the node that this remote disk is mounted on. | [optional] 
**UsedCapacityBytesPhysical** | Pointer to **NullableInt64** | Specifies the physical used capacity of remote disk in bytes. | [optional] [readonly] 
**CapacityBytesPhysical** | Pointer to **NullableInt64** | Specifies the physical capacity of the disk in bytes. | [optional] [readonly] 

## Methods

### NewRemoteDisk

`func NewRemoteDisk(mountPath NullableString, tier NullableString, ) *RemoteDisk`

NewRemoteDisk instantiates a new RemoteDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteDiskWithDefaults

`func NewRemoteDiskWithDefaults() *RemoteDisk`

NewRemoteDiskWithDefaults instantiates a new RemoteDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteDisk) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteDisk) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteDisk) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteDisk) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RemoteDisk) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RemoteDisk) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMountPath

`func (o *RemoteDisk) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *RemoteDisk) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *RemoteDisk) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.


### SetMountPathNil

`func (o *RemoteDisk) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *RemoteDisk) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetNodeId

`func (o *RemoteDisk) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *RemoteDisk) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *RemoteDisk) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *RemoteDisk) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *RemoteDisk) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *RemoteDisk) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetTier

`func (o *RemoteDisk) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *RemoteDisk) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *RemoteDisk) SetTier(v string)`

SetTier sets Tier field to given value.


### SetTierNil

`func (o *RemoteDisk) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *RemoteDisk) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetCapacityBytes

`func (o *RemoteDisk) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *RemoteDisk) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *RemoteDisk) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.

### HasCapacityBytes

`func (o *RemoteDisk) HasCapacityBytes() bool`

HasCapacityBytes returns a boolean if a field has been set.

### SetCapacityBytesNil

`func (o *RemoteDisk) SetCapacityBytesNil(b bool)`

 SetCapacityBytesNil sets the value for CapacityBytes to be an explicit nil

### UnsetCapacityBytes
`func (o *RemoteDisk) UnsetCapacityBytes()`

UnsetCapacityBytes ensures that no value is present for CapacityBytes, not even an explicit nil
### GetUsedCapacityBytes

`func (o *RemoteDisk) GetUsedCapacityBytes() int64`

GetUsedCapacityBytes returns the UsedCapacityBytes field if non-nil, zero value otherwise.

### GetUsedCapacityBytesOk

`func (o *RemoteDisk) GetUsedCapacityBytesOk() (*int64, bool)`

GetUsedCapacityBytesOk returns a tuple with the UsedCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacityBytes

`func (o *RemoteDisk) SetUsedCapacityBytes(v int64)`

SetUsedCapacityBytes sets UsedCapacityBytes field to given value.

### HasUsedCapacityBytes

`func (o *RemoteDisk) HasUsedCapacityBytes() bool`

HasUsedCapacityBytes returns a boolean if a field has been set.

### SetUsedCapacityBytesNil

`func (o *RemoteDisk) SetUsedCapacityBytesNil(b bool)`

 SetUsedCapacityBytesNil sets the value for UsedCapacityBytes to be an explicit nil

### UnsetUsedCapacityBytes
`func (o *RemoteDisk) UnsetUsedCapacityBytes()`

UnsetUsedCapacityBytes ensures that no value is present for UsedCapacityBytes, not even an explicit nil
### GetStatus

`func (o *RemoteDisk) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RemoteDisk) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RemoteDisk) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RemoteDisk) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RemoteDisk) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RemoteDisk) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFileSystemName

`func (o *RemoteDisk) GetFileSystemName() string`

GetFileSystemName returns the FileSystemName field if non-nil, zero value otherwise.

### GetFileSystemNameOk

`func (o *RemoteDisk) GetFileSystemNameOk() (*string, bool)`

GetFileSystemNameOk returns a tuple with the FileSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystemName

`func (o *RemoteDisk) SetFileSystemName(v string)`

SetFileSystemName sets FileSystemName field to given value.

### HasFileSystemName

`func (o *RemoteDisk) HasFileSystemName() bool`

HasFileSystemName returns a boolean if a field has been set.

### SetFileSystemNameNil

`func (o *RemoteDisk) SetFileSystemNameNil(b bool)`

 SetFileSystemNameNil sets the value for FileSystemName to be an explicit nil

### UnsetFileSystemName
`func (o *RemoteDisk) UnsetFileSystemName()`

UnsetFileSystemName ensures that no value is present for FileSystemName, not even an explicit nil
### GetDataVip

`func (o *RemoteDisk) GetDataVip() string`

GetDataVip returns the DataVip field if non-nil, zero value otherwise.

### GetDataVipOk

`func (o *RemoteDisk) GetDataVipOk() (*string, bool)`

GetDataVipOk returns a tuple with the DataVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVip

`func (o *RemoteDisk) SetDataVip(v string)`

SetDataVip sets DataVip field to given value.

### HasDataVip

`func (o *RemoteDisk) HasDataVip() bool`

HasDataVip returns a boolean if a field has been set.

### SetDataVipNil

`func (o *RemoteDisk) SetDataVipNil(b bool)`

 SetDataVipNil sets the value for DataVip to be an explicit nil

### UnsetDataVip
`func (o *RemoteDisk) UnsetDataVip()`

UnsetDataVip ensures that no value is present for DataVip, not even an explicit nil
### GetNodeIp

`func (o *RemoteDisk) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *RemoteDisk) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *RemoteDisk) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *RemoteDisk) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *RemoteDisk) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *RemoteDisk) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetUsedCapacityBytesPhysical

`func (o *RemoteDisk) GetUsedCapacityBytesPhysical() int64`

GetUsedCapacityBytesPhysical returns the UsedCapacityBytesPhysical field if non-nil, zero value otherwise.

### GetUsedCapacityBytesPhysicalOk

`func (o *RemoteDisk) GetUsedCapacityBytesPhysicalOk() (*int64, bool)`

GetUsedCapacityBytesPhysicalOk returns a tuple with the UsedCapacityBytesPhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCapacityBytesPhysical

`func (o *RemoteDisk) SetUsedCapacityBytesPhysical(v int64)`

SetUsedCapacityBytesPhysical sets UsedCapacityBytesPhysical field to given value.

### HasUsedCapacityBytesPhysical

`func (o *RemoteDisk) HasUsedCapacityBytesPhysical() bool`

HasUsedCapacityBytesPhysical returns a boolean if a field has been set.

### SetUsedCapacityBytesPhysicalNil

`func (o *RemoteDisk) SetUsedCapacityBytesPhysicalNil(b bool)`

 SetUsedCapacityBytesPhysicalNil sets the value for UsedCapacityBytesPhysical to be an explicit nil

### UnsetUsedCapacityBytesPhysical
`func (o *RemoteDisk) UnsetUsedCapacityBytesPhysical()`

UnsetUsedCapacityBytesPhysical ensures that no value is present for UsedCapacityBytesPhysical, not even an explicit nil
### GetCapacityBytesPhysical

`func (o *RemoteDisk) GetCapacityBytesPhysical() int64`

GetCapacityBytesPhysical returns the CapacityBytesPhysical field if non-nil, zero value otherwise.

### GetCapacityBytesPhysicalOk

`func (o *RemoteDisk) GetCapacityBytesPhysicalOk() (*int64, bool)`

GetCapacityBytesPhysicalOk returns a tuple with the CapacityBytesPhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytesPhysical

`func (o *RemoteDisk) SetCapacityBytesPhysical(v int64)`

SetCapacityBytesPhysical sets CapacityBytesPhysical field to given value.

### HasCapacityBytesPhysical

`func (o *RemoteDisk) HasCapacityBytesPhysical() bool`

HasCapacityBytesPhysical returns a boolean if a field has been set.

### SetCapacityBytesPhysicalNil

`func (o *RemoteDisk) SetCapacityBytesPhysicalNil(b bool)`

 SetCapacityBytesPhysicalNil sets the value for CapacityBytesPhysical to be an explicit nil

### UnsetCapacityBytesPhysical
`func (o *RemoteDisk) UnsetCapacityBytesPhysical()`

UnsetCapacityBytesPhysical ensures that no value is present for CapacityBytesPhysical, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


