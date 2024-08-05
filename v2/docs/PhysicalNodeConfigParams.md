# PhysicalNodeConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Specifies the node ID for this node. | 
**Ip** | **string** | Specifies the IP address for the node. | 
**IpmiIp** | Pointer to **string** | Specifies IPMI IP for the node. | [optional] 
**IsComputeNode** | Pointer to **bool** | Specifies whether to use the node for compute only. | [optional] 

## Methods

### NewPhysicalNodeConfigParams

`func NewPhysicalNodeConfigParams(id int64, ip string, ) *PhysicalNodeConfigParams`

NewPhysicalNodeConfigParams instantiates a new PhysicalNodeConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalNodeConfigParamsWithDefaults

`func NewPhysicalNodeConfigParamsWithDefaults() *PhysicalNodeConfigParams`

NewPhysicalNodeConfigParamsWithDefaults instantiates a new PhysicalNodeConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhysicalNodeConfigParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhysicalNodeConfigParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhysicalNodeConfigParams) SetId(v int64)`

SetId sets Id field to given value.


### GetIp

`func (o *PhysicalNodeConfigParams) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PhysicalNodeConfigParams) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PhysicalNodeConfigParams) SetIp(v string)`

SetIp sets Ip field to given value.


### GetIpmiIp

`func (o *PhysicalNodeConfigParams) GetIpmiIp() string`

GetIpmiIp returns the IpmiIp field if non-nil, zero value otherwise.

### GetIpmiIpOk

`func (o *PhysicalNodeConfigParams) GetIpmiIpOk() (*string, bool)`

GetIpmiIpOk returns a tuple with the IpmiIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiIp

`func (o *PhysicalNodeConfigParams) SetIpmiIp(v string)`

SetIpmiIp sets IpmiIp field to given value.

### HasIpmiIp

`func (o *PhysicalNodeConfigParams) HasIpmiIp() bool`

HasIpmiIp returns a boolean if a field has been set.

### GetIsComputeNode

`func (o *PhysicalNodeConfigParams) GetIsComputeNode() bool`

GetIsComputeNode returns the IsComputeNode field if non-nil, zero value otherwise.

### GetIsComputeNodeOk

`func (o *PhysicalNodeConfigParams) GetIsComputeNodeOk() (*bool, bool)`

GetIsComputeNodeOk returns a tuple with the IsComputeNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComputeNode

`func (o *PhysicalNodeConfigParams) SetIsComputeNode(v bool)`

SetIsComputeNode sets IsComputeNode field to given value.

### HasIsComputeNode

`func (o *PhysicalNodeConfigParams) HasIsComputeNode() bool`

HasIsComputeNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


