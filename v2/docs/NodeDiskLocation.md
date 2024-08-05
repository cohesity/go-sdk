# NodeDiskLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskId** | Pointer to **int64** | The id of the disk where the data is exported. More information w.r.t the disk(mount_path etc) is found in cluster config. | [optional] 
**ExpectedMountPath** | Pointer to **string** | Denotes the mount path of the disk. NOTE: This is only expected as it might change during reboots. | [optional] 
**ExpectedNodeIp** | Pointer to **string** | Denotes the IP of node on which the disk is mounted. NOTE: This is only expected as it might change during reboots. | [optional] 

## Methods

### NewNodeDiskLocation

`func NewNodeDiskLocation() *NodeDiskLocation`

NewNodeDiskLocation instantiates a new NodeDiskLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDiskLocationWithDefaults

`func NewNodeDiskLocationWithDefaults() *NodeDiskLocation`

NewNodeDiskLocationWithDefaults instantiates a new NodeDiskLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskId

`func (o *NodeDiskLocation) GetDiskId() int64`

GetDiskId returns the DiskId field if non-nil, zero value otherwise.

### GetDiskIdOk

`func (o *NodeDiskLocation) GetDiskIdOk() (*int64, bool)`

GetDiskIdOk returns a tuple with the DiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskId

`func (o *NodeDiskLocation) SetDiskId(v int64)`

SetDiskId sets DiskId field to given value.

### HasDiskId

`func (o *NodeDiskLocation) HasDiskId() bool`

HasDiskId returns a boolean if a field has been set.

### GetExpectedMountPath

`func (o *NodeDiskLocation) GetExpectedMountPath() string`

GetExpectedMountPath returns the ExpectedMountPath field if non-nil, zero value otherwise.

### GetExpectedMountPathOk

`func (o *NodeDiskLocation) GetExpectedMountPathOk() (*string, bool)`

GetExpectedMountPathOk returns a tuple with the ExpectedMountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedMountPath

`func (o *NodeDiskLocation) SetExpectedMountPath(v string)`

SetExpectedMountPath sets ExpectedMountPath field to given value.

### HasExpectedMountPath

`func (o *NodeDiskLocation) HasExpectedMountPath() bool`

HasExpectedMountPath returns a boolean if a field has been set.

### GetExpectedNodeIp

`func (o *NodeDiskLocation) GetExpectedNodeIp() string`

GetExpectedNodeIp returns the ExpectedNodeIp field if non-nil, zero value otherwise.

### GetExpectedNodeIpOk

`func (o *NodeDiskLocation) GetExpectedNodeIpOk() (*string, bool)`

GetExpectedNodeIpOk returns a tuple with the ExpectedNodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedNodeIp

`func (o *NodeDiskLocation) SetExpectedNodeIp(v string)`

SetExpectedNodeIp sets ExpectedNodeIp field to given value.

### HasExpectedNodeIp

`func (o *NodeDiskLocation) HasExpectedNodeIp() bool`

HasExpectedNodeIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


