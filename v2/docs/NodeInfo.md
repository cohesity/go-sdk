# NodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChassisModel** | Pointer to **string** | Chassis model. | [optional] 
**ChassisSerial** | Pointer to **string** | Chassis serial number programmed by manufacturer. | [optional] 
**ClusterId** | Pointer to **int64** | Specifies the Id of the cluster to which the node belongs. | [optional] 
**CohesityChassisSerial** | Pointer to **string** | Chassis serial number programmed by cohesity software. | [optional] 
**CohesityNodeSerial** | Pointer to **string** | Node serial number programmed by cohesity software. | [optional] 
**Cpu** | Pointer to **int32** | Number of CPUs | [optional] 
**Hostname** | Pointer to **string** | Host name of the node reported by the kernel. | [optional] 
**IncarnationId** | Pointer to **int64** | Specifies the cluster incarnation Id. | [optional] 
**InterfaceList** | Pointer to [**[]EndPoint**](EndPoint.md) | List of interfaces in node. | [optional] 
**IpmiIp** | Pointer to **string** | Ipmi IpAddress | [optional] 
**IsNodeReachable** | Pointer to **bool** | Specifies whether the node is reachable or not | [optional] 
**NodeId** | Pointer to **int64** | Specifies the Id of the node. | [optional] 
**NodeModel** | Pointer to **string** | Node model. | [optional] 
**NodeSerial** | Pointer to **string** | Node serial number programmed by manufacturer. | [optional] 
**ProductModel** | Pointer to **string** | Product Model | [optional] 
**ServicesVersionInfo** | Pointer to [**[]ServiceVersionInfo**](ServiceVersionInfo.md) | Specifies the version information of the cohesity services. | [optional] 
**SlotNumber** | Pointer to **string** | Slot number of the node in the chassis. | [optional] 
**SoftwareVersion** | Pointer to **string** | Version of the Cohesity software running on the node. | [optional] 
**SystemMemoryBytes** | Pointer to **int64** | System Memory in bytes | [optional] 

## Methods

### NewNodeInfo

`func NewNodeInfo() *NodeInfo`

NewNodeInfo instantiates a new NodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeInfoWithDefaults

`func NewNodeInfoWithDefaults() *NodeInfo`

NewNodeInfoWithDefaults instantiates a new NodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassisModel

`func (o *NodeInfo) GetChassisModel() string`

GetChassisModel returns the ChassisModel field if non-nil, zero value otherwise.

### GetChassisModelOk

`func (o *NodeInfo) GetChassisModelOk() (*string, bool)`

GetChassisModelOk returns a tuple with the ChassisModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisModel

`func (o *NodeInfo) SetChassisModel(v string)`

SetChassisModel sets ChassisModel field to given value.

### HasChassisModel

`func (o *NodeInfo) HasChassisModel() bool`

HasChassisModel returns a boolean if a field has been set.

### GetChassisSerial

`func (o *NodeInfo) GetChassisSerial() string`

GetChassisSerial returns the ChassisSerial field if non-nil, zero value otherwise.

### GetChassisSerialOk

`func (o *NodeInfo) GetChassisSerialOk() (*string, bool)`

GetChassisSerialOk returns a tuple with the ChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerial

`func (o *NodeInfo) SetChassisSerial(v string)`

SetChassisSerial sets ChassisSerial field to given value.

### HasChassisSerial

`func (o *NodeInfo) HasChassisSerial() bool`

HasChassisSerial returns a boolean if a field has been set.

### GetClusterId

`func (o *NodeInfo) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NodeInfo) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NodeInfo) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *NodeInfo) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetCohesityChassisSerial

`func (o *NodeInfo) GetCohesityChassisSerial() string`

GetCohesityChassisSerial returns the CohesityChassisSerial field if non-nil, zero value otherwise.

### GetCohesityChassisSerialOk

`func (o *NodeInfo) GetCohesityChassisSerialOk() (*string, bool)`

GetCohesityChassisSerialOk returns a tuple with the CohesityChassisSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohesityChassisSerial

`func (o *NodeInfo) SetCohesityChassisSerial(v string)`

SetCohesityChassisSerial sets CohesityChassisSerial field to given value.

### HasCohesityChassisSerial

`func (o *NodeInfo) HasCohesityChassisSerial() bool`

HasCohesityChassisSerial returns a boolean if a field has been set.

### GetCohesityNodeSerial

`func (o *NodeInfo) GetCohesityNodeSerial() string`

GetCohesityNodeSerial returns the CohesityNodeSerial field if non-nil, zero value otherwise.

### GetCohesityNodeSerialOk

`func (o *NodeInfo) GetCohesityNodeSerialOk() (*string, bool)`

GetCohesityNodeSerialOk returns a tuple with the CohesityNodeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCohesityNodeSerial

`func (o *NodeInfo) SetCohesityNodeSerial(v string)`

SetCohesityNodeSerial sets CohesityNodeSerial field to given value.

### HasCohesityNodeSerial

`func (o *NodeInfo) HasCohesityNodeSerial() bool`

HasCohesityNodeSerial returns a boolean if a field has been set.

### GetCpu

`func (o *NodeInfo) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeInfo) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeInfo) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *NodeInfo) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetHostname

`func (o *NodeInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NodeInfo) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIncarnationId

`func (o *NodeInfo) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *NodeInfo) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *NodeInfo) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *NodeInfo) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### GetInterfaceList

`func (o *NodeInfo) GetInterfaceList() []EndPoint`

GetInterfaceList returns the InterfaceList field if non-nil, zero value otherwise.

### GetInterfaceListOk

`func (o *NodeInfo) GetInterfaceListOk() (*[]EndPoint, bool)`

GetInterfaceListOk returns a tuple with the InterfaceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceList

`func (o *NodeInfo) SetInterfaceList(v []EndPoint)`

SetInterfaceList sets InterfaceList field to given value.

### HasInterfaceList

`func (o *NodeInfo) HasInterfaceList() bool`

HasInterfaceList returns a boolean if a field has been set.

### GetIpmiIp

`func (o *NodeInfo) GetIpmiIp() string`

GetIpmiIp returns the IpmiIp field if non-nil, zero value otherwise.

### GetIpmiIpOk

`func (o *NodeInfo) GetIpmiIpOk() (*string, bool)`

GetIpmiIpOk returns a tuple with the IpmiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiIp

`func (o *NodeInfo) SetIpmiIp(v string)`

SetIpmiIp sets IpmiIp field to given value.

### HasIpmiIp

`func (o *NodeInfo) HasIpmiIp() bool`

HasIpmiIp returns a boolean if a field has been set.

### GetIsNodeReachable

`func (o *NodeInfo) GetIsNodeReachable() bool`

GetIsNodeReachable returns the IsNodeReachable field if non-nil, zero value otherwise.

### GetIsNodeReachableOk

`func (o *NodeInfo) GetIsNodeReachableOk() (*bool, bool)`

GetIsNodeReachableOk returns a tuple with the IsNodeReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNodeReachable

`func (o *NodeInfo) SetIsNodeReachable(v bool)`

SetIsNodeReachable sets IsNodeReachable field to given value.

### HasIsNodeReachable

`func (o *NodeInfo) HasIsNodeReachable() bool`

HasIsNodeReachable returns a boolean if a field has been set.

### GetNodeId

`func (o *NodeInfo) GetNodeId() int64`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeInfo) GetNodeIdOk() (*int64, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeInfo) SetNodeId(v int64)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeModel

`func (o *NodeInfo) GetNodeModel() string`

GetNodeModel returns the NodeModel field if non-nil, zero value otherwise.

### GetNodeModelOk

`func (o *NodeInfo) GetNodeModelOk() (*string, bool)`

GetNodeModelOk returns a tuple with the NodeModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeModel

`func (o *NodeInfo) SetNodeModel(v string)`

SetNodeModel sets NodeModel field to given value.

### HasNodeModel

`func (o *NodeInfo) HasNodeModel() bool`

HasNodeModel returns a boolean if a field has been set.

### GetNodeSerial

`func (o *NodeInfo) GetNodeSerial() string`

GetNodeSerial returns the NodeSerial field if non-nil, zero value otherwise.

### GetNodeSerialOk

`func (o *NodeInfo) GetNodeSerialOk() (*string, bool)`

GetNodeSerialOk returns a tuple with the NodeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSerial

`func (o *NodeInfo) SetNodeSerial(v string)`

SetNodeSerial sets NodeSerial field to given value.

### HasNodeSerial

`func (o *NodeInfo) HasNodeSerial() bool`

HasNodeSerial returns a boolean if a field has been set.

### GetProductModel

`func (o *NodeInfo) GetProductModel() string`

GetProductModel returns the ProductModel field if non-nil, zero value otherwise.

### GetProductModelOk

`func (o *NodeInfo) GetProductModelOk() (*string, bool)`

GetProductModelOk returns a tuple with the ProductModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductModel

`func (o *NodeInfo) SetProductModel(v string)`

SetProductModel sets ProductModel field to given value.

### HasProductModel

`func (o *NodeInfo) HasProductModel() bool`

HasProductModel returns a boolean if a field has been set.

### GetServicesVersionInfo

`func (o *NodeInfo) GetServicesVersionInfo() []ServiceVersionInfo`

GetServicesVersionInfo returns the ServicesVersionInfo field if non-nil, zero value otherwise.

### GetServicesVersionInfoOk

`func (o *NodeInfo) GetServicesVersionInfoOk() (*[]ServiceVersionInfo, bool)`

GetServicesVersionInfoOk returns a tuple with the ServicesVersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesVersionInfo

`func (o *NodeInfo) SetServicesVersionInfo(v []ServiceVersionInfo)`

SetServicesVersionInfo sets ServicesVersionInfo field to given value.

### HasServicesVersionInfo

`func (o *NodeInfo) HasServicesVersionInfo() bool`

HasServicesVersionInfo returns a boolean if a field has been set.

### GetSlotNumber

`func (o *NodeInfo) GetSlotNumber() string`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *NodeInfo) GetSlotNumberOk() (*string, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *NodeInfo) SetSlotNumber(v string)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *NodeInfo) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *NodeInfo) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *NodeInfo) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *NodeInfo) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *NodeInfo) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetSystemMemoryBytes

`func (o *NodeInfo) GetSystemMemoryBytes() int64`

GetSystemMemoryBytes returns the SystemMemoryBytes field if non-nil, zero value otherwise.

### GetSystemMemoryBytesOk

`func (o *NodeInfo) GetSystemMemoryBytesOk() (*int64, bool)`

GetSystemMemoryBytesOk returns a tuple with the SystemMemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMemoryBytes

`func (o *NodeInfo) SetSystemMemoryBytes(v int64)`

SetSystemMemoryBytes sets SystemMemoryBytes field to given value.

### HasSystemMemoryBytes

`func (o *NodeInfo) HasSystemMemoryBytes() bool`

HasSystemMemoryBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


