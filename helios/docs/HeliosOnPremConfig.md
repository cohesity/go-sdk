# HeliosOnPremConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Name of the new Helios OnPrem VM. | [optional] 
**Nodes** | Pointer to [**[]HeliosOnPremVMNode**](HeliosOnPremVMNode.md) | Specifies the Nodes present in this Cluster. | [optional] 
**KubernetesSubnetCidr** | Pointer to **NullableString** | Subnet to use for setting up the Kubernetes cluster&#39;s internal network on which Cohesity Helios will run. | [optional] 
**NetworkConfig** | Pointer to [**ClusterCreateNetworkConfig**](ClusterCreateNetworkConfig.md) |  | [optional] 
**ProxyServerConfig** | Pointer to [**NullableClusterProxyServerConfig**](ClusterProxyServerConfig.md) |  | [optional] 

## Methods

### NewHeliosOnPremConfig

`func NewHeliosOnPremConfig() *HeliosOnPremConfig`

NewHeliosOnPremConfig instantiates a new HeliosOnPremConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosOnPremConfigWithDefaults

`func NewHeliosOnPremConfigWithDefaults() *HeliosOnPremConfig`

NewHeliosOnPremConfigWithDefaults instantiates a new HeliosOnPremConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HeliosOnPremConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosOnPremConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosOnPremConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeliosOnPremConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeliosOnPremConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosOnPremConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNodes

`func (o *HeliosOnPremConfig) GetNodes() []HeliosOnPremVMNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *HeliosOnPremConfig) GetNodesOk() (*[]HeliosOnPremVMNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *HeliosOnPremConfig) SetNodes(v []HeliosOnPremVMNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *HeliosOnPremConfig) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetKubernetesSubnetCidr

`func (o *HeliosOnPremConfig) GetKubernetesSubnetCidr() string`

GetKubernetesSubnetCidr returns the KubernetesSubnetCidr field if non-nil, zero value otherwise.

### GetKubernetesSubnetCidrOk

`func (o *HeliosOnPremConfig) GetKubernetesSubnetCidrOk() (*string, bool)`

GetKubernetesSubnetCidrOk returns a tuple with the KubernetesSubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesSubnetCidr

`func (o *HeliosOnPremConfig) SetKubernetesSubnetCidr(v string)`

SetKubernetesSubnetCidr sets KubernetesSubnetCidr field to given value.

### HasKubernetesSubnetCidr

`func (o *HeliosOnPremConfig) HasKubernetesSubnetCidr() bool`

HasKubernetesSubnetCidr returns a boolean if a field has been set.

### SetKubernetesSubnetCidrNil

`func (o *HeliosOnPremConfig) SetKubernetesSubnetCidrNil(b bool)`

 SetKubernetesSubnetCidrNil sets the value for KubernetesSubnetCidr to be an explicit nil

### UnsetKubernetesSubnetCidr
`func (o *HeliosOnPremConfig) UnsetKubernetesSubnetCidr()`

UnsetKubernetesSubnetCidr ensures that no value is present for KubernetesSubnetCidr, not even an explicit nil
### GetNetworkConfig

`func (o *HeliosOnPremConfig) GetNetworkConfig() ClusterCreateNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *HeliosOnPremConfig) GetNetworkConfigOk() (*ClusterCreateNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *HeliosOnPremConfig) SetNetworkConfig(v ClusterCreateNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.

### HasNetworkConfig

`func (o *HeliosOnPremConfig) HasNetworkConfig() bool`

HasNetworkConfig returns a boolean if a field has been set.

### GetProxyServerConfig

`func (o *HeliosOnPremConfig) GetProxyServerConfig() ClusterProxyServerConfig`

GetProxyServerConfig returns the ProxyServerConfig field if non-nil, zero value otherwise.

### GetProxyServerConfigOk

`func (o *HeliosOnPremConfig) GetProxyServerConfigOk() (*ClusterProxyServerConfig, bool)`

GetProxyServerConfigOk returns a tuple with the ProxyServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyServerConfig

`func (o *HeliosOnPremConfig) SetProxyServerConfig(v ClusterProxyServerConfig)`

SetProxyServerConfig sets ProxyServerConfig field to given value.

### HasProxyServerConfig

`func (o *HeliosOnPremConfig) HasProxyServerConfig() bool`

HasProxyServerConfig returns a boolean if a field has been set.

### SetProxyServerConfigNil

`func (o *HeliosOnPremConfig) SetProxyServerConfigNil(b bool)`

 SetProxyServerConfigNil sets the value for ProxyServerConfig to be an explicit nil

### UnsetProxyServerConfig
`func (o *HeliosOnPremConfig) UnsetProxyServerConfig()`

UnsetProxyServerConfig ensures that no value is present for ProxyServerConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


