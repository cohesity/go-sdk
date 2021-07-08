# NodeSystemDiskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicePath** | Pointer to **NullableString** | DevicePath is the device path of the disk. | [optional] 
**Id** | Pointer to **NullableInt64** | Id is the id of the disk. | [optional] 
**Offline** | Pointer to **NullableBool** | Offline specifies whether a disk is marked offline. | [optional] 

## Methods

### NewNodeSystemDiskInfo

`func NewNodeSystemDiskInfo() *NodeSystemDiskInfo`

NewNodeSystemDiskInfo instantiates a new NodeSystemDiskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeSystemDiskInfoWithDefaults

`func NewNodeSystemDiskInfoWithDefaults() *NodeSystemDiskInfo`

NewNodeSystemDiskInfoWithDefaults instantiates a new NodeSystemDiskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicePath

`func (o *NodeSystemDiskInfo) GetDevicePath() string`

GetDevicePath returns the DevicePath field if non-nil, zero value otherwise.

### GetDevicePathOk

`func (o *NodeSystemDiskInfo) GetDevicePathOk() (*string, bool)`

GetDevicePathOk returns a tuple with the DevicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePath

`func (o *NodeSystemDiskInfo) SetDevicePath(v string)`

SetDevicePath sets DevicePath field to given value.

### HasDevicePath

`func (o *NodeSystemDiskInfo) HasDevicePath() bool`

HasDevicePath returns a boolean if a field has been set.

### SetDevicePathNil

`func (o *NodeSystemDiskInfo) SetDevicePathNil(b bool)`

 SetDevicePathNil sets the value for DevicePath to be an explicit nil

### UnsetDevicePath
`func (o *NodeSystemDiskInfo) UnsetDevicePath()`

UnsetDevicePath ensures that no value is present for DevicePath, not even an explicit nil
### GetId

`func (o *NodeSystemDiskInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeSystemDiskInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeSystemDiskInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NodeSystemDiskInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NodeSystemDiskInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NodeSystemDiskInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOffline

`func (o *NodeSystemDiskInfo) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *NodeSystemDiskInfo) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *NodeSystemDiskInfo) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *NodeSystemDiskInfo) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### SetOfflineNil

`func (o *NodeSystemDiskInfo) SetOfflineNil(b bool)`

 SetOfflineNil sets the value for Offline to be an explicit nil

### UnsetOffline
`func (o *NodeSystemDiskInfo) UnsetOffline()`

UnsetOffline ensures that no value is present for Offline, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


