# DeployVMsToAWSParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**KeyPairName** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**NetworkSecurityGroups** | Pointer to [**[]EntityProto**](EntityProto.md) | Names of the network security groups within the above VPC. At least one entry should be present. | [optional] 
**ProxyVmSubnet** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**ProxyVmVpc** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**RdsParams** | Pointer to [**DeployDBInstancesToRDSParams**](DeployDBInstancesToRDSParams.md) |  | [optional] 
**Region** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**Subnet** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**Vpc** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 

## Methods

### NewDeployVMsToAWSParams

`func NewDeployVMsToAWSParams() *DeployVMsToAWSParams`

NewDeployVMsToAWSParams instantiates a new DeployVMsToAWSParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployVMsToAWSParamsWithDefaults

`func NewDeployVMsToAWSParamsWithDefaults() *DeployVMsToAWSParams`

NewDeployVMsToAWSParamsWithDefaults instantiates a new DeployVMsToAWSParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *DeployVMsToAWSParams) GetInstanceType() EntityProto`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *DeployVMsToAWSParams) GetInstanceTypeOk() (*EntityProto, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *DeployVMsToAWSParams) SetInstanceType(v EntityProto)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *DeployVMsToAWSParams) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetKeyPairName

`func (o *DeployVMsToAWSParams) GetKeyPairName() EntityProto`

GetKeyPairName returns the KeyPairName field if non-nil, zero value otherwise.

### GetKeyPairNameOk

`func (o *DeployVMsToAWSParams) GetKeyPairNameOk() (*EntityProto, bool)`

GetKeyPairNameOk returns a tuple with the KeyPairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairName

`func (o *DeployVMsToAWSParams) SetKeyPairName(v EntityProto)`

SetKeyPairName sets KeyPairName field to given value.

### HasKeyPairName

`func (o *DeployVMsToAWSParams) HasKeyPairName() bool`

HasKeyPairName returns a boolean if a field has been set.

### GetNetworkSecurityGroups

`func (o *DeployVMsToAWSParams) GetNetworkSecurityGroups() []EntityProto`

GetNetworkSecurityGroups returns the NetworkSecurityGroups field if non-nil, zero value otherwise.

### GetNetworkSecurityGroupsOk

`func (o *DeployVMsToAWSParams) GetNetworkSecurityGroupsOk() (*[]EntityProto, bool)`

GetNetworkSecurityGroupsOk returns a tuple with the NetworkSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSecurityGroups

`func (o *DeployVMsToAWSParams) SetNetworkSecurityGroups(v []EntityProto)`

SetNetworkSecurityGroups sets NetworkSecurityGroups field to given value.

### HasNetworkSecurityGroups

`func (o *DeployVMsToAWSParams) HasNetworkSecurityGroups() bool`

HasNetworkSecurityGroups returns a boolean if a field has been set.

### SetNetworkSecurityGroupsNil

`func (o *DeployVMsToAWSParams) SetNetworkSecurityGroupsNil(b bool)`

 SetNetworkSecurityGroupsNil sets the value for NetworkSecurityGroups to be an explicit nil

### UnsetNetworkSecurityGroups
`func (o *DeployVMsToAWSParams) UnsetNetworkSecurityGroups()`

UnsetNetworkSecurityGroups ensures that no value is present for NetworkSecurityGroups, not even an explicit nil
### GetProxyVmSubnet

`func (o *DeployVMsToAWSParams) GetProxyVmSubnet() EntityProto`

GetProxyVmSubnet returns the ProxyVmSubnet field if non-nil, zero value otherwise.

### GetProxyVmSubnetOk

`func (o *DeployVMsToAWSParams) GetProxyVmSubnetOk() (*EntityProto, bool)`

GetProxyVmSubnetOk returns a tuple with the ProxyVmSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyVmSubnet

`func (o *DeployVMsToAWSParams) SetProxyVmSubnet(v EntityProto)`

SetProxyVmSubnet sets ProxyVmSubnet field to given value.

### HasProxyVmSubnet

`func (o *DeployVMsToAWSParams) HasProxyVmSubnet() bool`

HasProxyVmSubnet returns a boolean if a field has been set.

### GetProxyVmVpc

`func (o *DeployVMsToAWSParams) GetProxyVmVpc() EntityProto`

GetProxyVmVpc returns the ProxyVmVpc field if non-nil, zero value otherwise.

### GetProxyVmVpcOk

`func (o *DeployVMsToAWSParams) GetProxyVmVpcOk() (*EntityProto, bool)`

GetProxyVmVpcOk returns a tuple with the ProxyVmVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyVmVpc

`func (o *DeployVMsToAWSParams) SetProxyVmVpc(v EntityProto)`

SetProxyVmVpc sets ProxyVmVpc field to given value.

### HasProxyVmVpc

`func (o *DeployVMsToAWSParams) HasProxyVmVpc() bool`

HasProxyVmVpc returns a boolean if a field has been set.

### GetRdsParams

`func (o *DeployVMsToAWSParams) GetRdsParams() DeployDBInstancesToRDSParams`

GetRdsParams returns the RdsParams field if non-nil, zero value otherwise.

### GetRdsParamsOk

`func (o *DeployVMsToAWSParams) GetRdsParamsOk() (*DeployDBInstancesToRDSParams, bool)`

GetRdsParamsOk returns a tuple with the RdsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsParams

`func (o *DeployVMsToAWSParams) SetRdsParams(v DeployDBInstancesToRDSParams)`

SetRdsParams sets RdsParams field to given value.

### HasRdsParams

`func (o *DeployVMsToAWSParams) HasRdsParams() bool`

HasRdsParams returns a boolean if a field has been set.

### GetRegion

`func (o *DeployVMsToAWSParams) GetRegion() EntityProto`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeployVMsToAWSParams) GetRegionOk() (*EntityProto, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeployVMsToAWSParams) SetRegion(v EntityProto)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DeployVMsToAWSParams) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSubnet

`func (o *DeployVMsToAWSParams) GetSubnet() EntityProto`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *DeployVMsToAWSParams) GetSubnetOk() (*EntityProto, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *DeployVMsToAWSParams) SetSubnet(v EntityProto)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *DeployVMsToAWSParams) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetVpc

`func (o *DeployVMsToAWSParams) GetVpc() EntityProto`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *DeployVMsToAWSParams) GetVpcOk() (*EntityProto, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *DeployVMsToAWSParams) SetVpc(v EntityProto)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *DeployVMsToAWSParams) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


