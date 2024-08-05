# NodeHardwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **NullableString** | Cpu provides the information regarding the CPU. | [optional] 
**MemorySizeBytes** | Pointer to **NullableInt64** | MemorySizeBytes provides the memory size in bytes. | [optional] 
**Network** | Pointer to **NullableString** | Network provides the information regarding the network cards. | [optional] 

## Methods

### NewNodeHardwareInfo

`func NewNodeHardwareInfo() *NodeHardwareInfo`

NewNodeHardwareInfo instantiates a new NodeHardwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeHardwareInfoWithDefaults

`func NewNodeHardwareInfoWithDefaults() *NodeHardwareInfo`

NewNodeHardwareInfoWithDefaults instantiates a new NodeHardwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *NodeHardwareInfo) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeHardwareInfo) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeHardwareInfo) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *NodeHardwareInfo) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *NodeHardwareInfo) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *NodeHardwareInfo) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemorySizeBytes

`func (o *NodeHardwareInfo) GetMemorySizeBytes() int64`

GetMemorySizeBytes returns the MemorySizeBytes field if non-nil, zero value otherwise.

### GetMemorySizeBytesOk

`func (o *NodeHardwareInfo) GetMemorySizeBytesOk() (*int64, bool)`

GetMemorySizeBytesOk returns a tuple with the MemorySizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeBytes

`func (o *NodeHardwareInfo) SetMemorySizeBytes(v int64)`

SetMemorySizeBytes sets MemorySizeBytes field to given value.

### HasMemorySizeBytes

`func (o *NodeHardwareInfo) HasMemorySizeBytes() bool`

HasMemorySizeBytes returns a boolean if a field has been set.

### SetMemorySizeBytesNil

`func (o *NodeHardwareInfo) SetMemorySizeBytesNil(b bool)`

 SetMemorySizeBytesNil sets the value for MemorySizeBytes to be an explicit nil

### UnsetMemorySizeBytes
`func (o *NodeHardwareInfo) UnsetMemorySizeBytes()`

UnsetMemorySizeBytes ensures that no value is present for MemorySizeBytes, not even an explicit nil
### GetNetwork

`func (o *NodeHardwareInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NodeHardwareInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NodeHardwareInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NodeHardwareInfo) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### SetNetworkNil

`func (o *NodeHardwareInfo) SetNetworkNil(b bool)`

 SetNetworkNil sets the value for Network to be an explicit nil

### UnsetNetwork
`func (o *NodeHardwareInfo) UnsetNetwork()`

UnsetNetwork ensures that no value is present for Network, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


