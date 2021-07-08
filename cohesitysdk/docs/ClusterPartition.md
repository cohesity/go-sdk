# ClusterPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | Pointer to **NullableString** | Specifies that hostname that resolves to one or more Virtual IP Addresses (VIPs). | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies a unique identifier for the Cluster Partition. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the Cluster Partition. | [optional] 
**NodeIds** | Pointer to **[]int64** | Array of Node Ids.  Specifies a list of Node Ids that assigned to the Cluster Partition. | [optional] 
**Vips** | Pointer to **[]string** | Array of VIPs.  Specifies a list of Virtual IP Addresses (VIPs) that route network traffic to the Cluster Partition. | [optional] 
**VlanIps** | Pointer to **[]string** | Array of VLAN IPs.  Specifies a list of VLAN IP Addresses that route network traffic within certain VLANs to the Cluster Partition. | [optional] 
**Vlans** | Pointer to [**[]Vlan**](Vlan.md) | Array of VLANs.  Specifies a list of VLANs for the Cluster Partition. | [optional] 

## Methods

### NewClusterPartition

`func NewClusterPartition() *ClusterPartition`

NewClusterPartition instantiates a new ClusterPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPartitionWithDefaults

`func NewClusterPartitionWithDefaults() *ClusterPartition`

NewClusterPartitionWithDefaults instantiates a new ClusterPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *ClusterPartition) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ClusterPartition) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ClusterPartition) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ClusterPartition) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *ClusterPartition) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *ClusterPartition) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetId

`func (o *ClusterPartition) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterPartition) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterPartition) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterPartition) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ClusterPartition) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ClusterPartition) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ClusterPartition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterPartition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterPartition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterPartition) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClusterPartition) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClusterPartition) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodeIds

`func (o *ClusterPartition) GetNodeIds() []int64`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *ClusterPartition) GetNodeIdsOk() (*[]int64, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *ClusterPartition) SetNodeIds(v []int64)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *ClusterPartition) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### SetNodeIdsNil

`func (o *ClusterPartition) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *ClusterPartition) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetVips

`func (o *ClusterPartition) GetVips() []string`

GetVips returns the Vips field if non-nil, zero value otherwise.

### GetVipsOk

`func (o *ClusterPartition) GetVipsOk() (*[]string, bool)`

GetVipsOk returns a tuple with the Vips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVips

`func (o *ClusterPartition) SetVips(v []string)`

SetVips sets Vips field to given value.

### HasVips

`func (o *ClusterPartition) HasVips() bool`

HasVips returns a boolean if a field has been set.

### SetVipsNil

`func (o *ClusterPartition) SetVipsNil(b bool)`

 SetVipsNil sets the value for Vips to be an explicit nil

### UnsetVips
`func (o *ClusterPartition) UnsetVips()`

UnsetVips ensures that no value is present for Vips, not even an explicit nil
### GetVlanIps

`func (o *ClusterPartition) GetVlanIps() []string`

GetVlanIps returns the VlanIps field if non-nil, zero value otherwise.

### GetVlanIpsOk

`func (o *ClusterPartition) GetVlanIpsOk() (*[]string, bool)`

GetVlanIpsOk returns a tuple with the VlanIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIps

`func (o *ClusterPartition) SetVlanIps(v []string)`

SetVlanIps sets VlanIps field to given value.

### HasVlanIps

`func (o *ClusterPartition) HasVlanIps() bool`

HasVlanIps returns a boolean if a field has been set.

### SetVlanIpsNil

`func (o *ClusterPartition) SetVlanIpsNil(b bool)`

 SetVlanIpsNil sets the value for VlanIps to be an explicit nil

### UnsetVlanIps
`func (o *ClusterPartition) UnsetVlanIps()`

UnsetVlanIps ensures that no value is present for VlanIps, not even an explicit nil
### GetVlans

`func (o *ClusterPartition) GetVlans() []Vlan`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *ClusterPartition) GetVlansOk() (*[]Vlan, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *ClusterPartition) SetVlans(v []Vlan)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *ClusterPartition) HasVlans() bool`

HasVlans returns a boolean if a field has been set.

### SetVlansNil

`func (o *ClusterPartition) SetVlansNil(b bool)`

 SetVlansNil sets the value for Vlans to be an explicit nil

### UnsetVlans
`func (o *ClusterPartition) UnsetVlans()`

UnsetVlans ensures that no value is present for Vlans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


