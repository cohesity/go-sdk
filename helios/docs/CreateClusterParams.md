# CreateClusterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the new cluster. | 
**Type** | **NullableString** | Specifies the type of the new cluster. | 
**EnableEncryption** | **NullableBool** | Specifies whether or not to enable encryption. If encryption is enabled, all data on the Cluster will be encrypted. This can only be enabled at Cluster creation time and cannot be disabled later. | 
**NetworkConfig** | [**ClusterCreateNetworkConfig**](ClusterCreateNetworkConfig.md) |  | 
**ProxyServerConfig** | Pointer to [**NullableClusterProxyServerConfig**](ClusterProxyServerConfig.md) |  | [optional] 
**PhysicalClusterParams** | Pointer to [**ClusterCreatePhysicalParams**](ClusterCreatePhysicalParams.md) |  | [optional] 
**VirtualClusterParams** | Pointer to [**ClusterCreateVirtualParams**](ClusterCreateVirtualParams.md) |  | [optional] 
**CloudClusterParams** | Pointer to [**NullableClusterCreateCloudParams**](ClusterCreateCloudParams.md) |  | [optional] 
**RigelClusterParams** | Pointer to [**NullableClusterCreateRigelParams**](ClusterCreateRigelParams.md) |  | [optional] 

## Methods

### NewCreateClusterParams

`func NewCreateClusterParams(name NullableString, type_ NullableString, enableEncryption NullableBool, networkConfig ClusterCreateNetworkConfig, ) *CreateClusterParams`

NewCreateClusterParams instantiates a new CreateClusterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterParamsWithDefaults

`func NewCreateClusterParamsWithDefaults() *CreateClusterParams`

NewCreateClusterParamsWithDefaults instantiates a new CreateClusterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClusterParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateClusterParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateClusterParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *CreateClusterParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateClusterParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateClusterParams) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CreateClusterParams) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CreateClusterParams) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEnableEncryption

`func (o *CreateClusterParams) GetEnableEncryption() bool`

GetEnableEncryption returns the EnableEncryption field if non-nil, zero value otherwise.

### GetEnableEncryptionOk

`func (o *CreateClusterParams) GetEnableEncryptionOk() (*bool, bool)`

GetEnableEncryptionOk returns a tuple with the EnableEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryption

`func (o *CreateClusterParams) SetEnableEncryption(v bool)`

SetEnableEncryption sets EnableEncryption field to given value.


### SetEnableEncryptionNil

`func (o *CreateClusterParams) SetEnableEncryptionNil(b bool)`

 SetEnableEncryptionNil sets the value for EnableEncryption to be an explicit nil

### UnsetEnableEncryption
`func (o *CreateClusterParams) UnsetEnableEncryption()`

UnsetEnableEncryption ensures that no value is present for EnableEncryption, not even an explicit nil
### GetNetworkConfig

`func (o *CreateClusterParams) GetNetworkConfig() ClusterCreateNetworkConfig`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *CreateClusterParams) GetNetworkConfigOk() (*ClusterCreateNetworkConfig, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *CreateClusterParams) SetNetworkConfig(v ClusterCreateNetworkConfig)`

SetNetworkConfig sets NetworkConfig field to given value.


### GetProxyServerConfig

`func (o *CreateClusterParams) GetProxyServerConfig() ClusterProxyServerConfig`

GetProxyServerConfig returns the ProxyServerConfig field if non-nil, zero value otherwise.

### GetProxyServerConfigOk

`func (o *CreateClusterParams) GetProxyServerConfigOk() (*ClusterProxyServerConfig, bool)`

GetProxyServerConfigOk returns a tuple with the ProxyServerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyServerConfig

`func (o *CreateClusterParams) SetProxyServerConfig(v ClusterProxyServerConfig)`

SetProxyServerConfig sets ProxyServerConfig field to given value.

### HasProxyServerConfig

`func (o *CreateClusterParams) HasProxyServerConfig() bool`

HasProxyServerConfig returns a boolean if a field has been set.

### SetProxyServerConfigNil

`func (o *CreateClusterParams) SetProxyServerConfigNil(b bool)`

 SetProxyServerConfigNil sets the value for ProxyServerConfig to be an explicit nil

### UnsetProxyServerConfig
`func (o *CreateClusterParams) UnsetProxyServerConfig()`

UnsetProxyServerConfig ensures that no value is present for ProxyServerConfig, not even an explicit nil
### GetPhysicalClusterParams

`func (o *CreateClusterParams) GetPhysicalClusterParams() ClusterCreatePhysicalParams`

GetPhysicalClusterParams returns the PhysicalClusterParams field if non-nil, zero value otherwise.

### GetPhysicalClusterParamsOk

`func (o *CreateClusterParams) GetPhysicalClusterParamsOk() (*ClusterCreatePhysicalParams, bool)`

GetPhysicalClusterParamsOk returns a tuple with the PhysicalClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalClusterParams

`func (o *CreateClusterParams) SetPhysicalClusterParams(v ClusterCreatePhysicalParams)`

SetPhysicalClusterParams sets PhysicalClusterParams field to given value.

### HasPhysicalClusterParams

`func (o *CreateClusterParams) HasPhysicalClusterParams() bool`

HasPhysicalClusterParams returns a boolean if a field has been set.

### GetVirtualClusterParams

`func (o *CreateClusterParams) GetVirtualClusterParams() ClusterCreateVirtualParams`

GetVirtualClusterParams returns the VirtualClusterParams field if non-nil, zero value otherwise.

### GetVirtualClusterParamsOk

`func (o *CreateClusterParams) GetVirtualClusterParamsOk() (*ClusterCreateVirtualParams, bool)`

GetVirtualClusterParamsOk returns a tuple with the VirtualClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualClusterParams

`func (o *CreateClusterParams) SetVirtualClusterParams(v ClusterCreateVirtualParams)`

SetVirtualClusterParams sets VirtualClusterParams field to given value.

### HasVirtualClusterParams

`func (o *CreateClusterParams) HasVirtualClusterParams() bool`

HasVirtualClusterParams returns a boolean if a field has been set.

### GetCloudClusterParams

`func (o *CreateClusterParams) GetCloudClusterParams() ClusterCreateCloudParams`

GetCloudClusterParams returns the CloudClusterParams field if non-nil, zero value otherwise.

### GetCloudClusterParamsOk

`func (o *CreateClusterParams) GetCloudClusterParamsOk() (*ClusterCreateCloudParams, bool)`

GetCloudClusterParamsOk returns a tuple with the CloudClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudClusterParams

`func (o *CreateClusterParams) SetCloudClusterParams(v ClusterCreateCloudParams)`

SetCloudClusterParams sets CloudClusterParams field to given value.

### HasCloudClusterParams

`func (o *CreateClusterParams) HasCloudClusterParams() bool`

HasCloudClusterParams returns a boolean if a field has been set.

### SetCloudClusterParamsNil

`func (o *CreateClusterParams) SetCloudClusterParamsNil(b bool)`

 SetCloudClusterParamsNil sets the value for CloudClusterParams to be an explicit nil

### UnsetCloudClusterParams
`func (o *CreateClusterParams) UnsetCloudClusterParams()`

UnsetCloudClusterParams ensures that no value is present for CloudClusterParams, not even an explicit nil
### GetRigelClusterParams

`func (o *CreateClusterParams) GetRigelClusterParams() ClusterCreateRigelParams`

GetRigelClusterParams returns the RigelClusterParams field if non-nil, zero value otherwise.

### GetRigelClusterParamsOk

`func (o *CreateClusterParams) GetRigelClusterParamsOk() (*ClusterCreateRigelParams, bool)`

GetRigelClusterParamsOk returns a tuple with the RigelClusterParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRigelClusterParams

`func (o *CreateClusterParams) SetRigelClusterParams(v ClusterCreateRigelParams)`

SetRigelClusterParams sets RigelClusterParams field to given value.

### HasRigelClusterParams

`func (o *CreateClusterParams) HasRigelClusterParams() bool`

HasRigelClusterParams returns a boolean if a field has been set.

### SetRigelClusterParamsNil

`func (o *CreateClusterParams) SetRigelClusterParamsNil(b bool)`

 SetRigelClusterParamsNil sets the value for RigelClusterParams to be an explicit nil

### UnsetRigelClusterParams
`func (o *CreateClusterParams) UnsetRigelClusterParams()`

UnsetRigelClusterParams ensures that no value is present for RigelClusterParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


