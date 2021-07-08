# CreateVirtualClusterParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **NullableString** | Specifies the name of the new Cluster. | 
**EncryptionConfig** | Pointer to [**EncryptionConfiguration**](EncryptionConfiguration.md) |  | [optional] 
**IpPreference** | Pointer to **NullableInt32** | Specifies IP preference. | [optional] 
**MetadataFaultTolerance** | Pointer to **NullableInt32** | Specifies the metadata fault tolerance. | [optional] 
**NetworkConfig** | [**NetworkConfiguration**](NetworkConfiguration.md) |  | 
**NodeConfigs** | [**[]VirtualNodeConfiguration**](VirtualNodeConfiguration.md) | Specifies the configuration for the nodes in the new cluster. | 

## Methods

### NewCreateVirtualClusterParameters

`func NewCreateVirtualClusterParameters(clusterName NullableString, networkConfig NetworkConfiguration, nodeConfigs []VirtualNodeConfiguration, ) *CreateVirtualClusterParameters`

NewCreateVirtualClusterParameters instantiates a new CreateVirtualClusterParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVirtualClusterParametersWithDefaults

`func NewCreateVirtualClusterParametersWithDefaults() *CreateVirtualClusterParameters`

NewCreateVirtualClusterParametersWithDefaults instantiates a new CreateVirtualClusterParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *CreateVirtualClusterParameters) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CreateVirtualClusterParameters) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CreateVirtualClusterParameters) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### SetClusterNameNil

`func (o *CreateVirtualClusterParameters) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *CreateVirtualClusterParameters) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetEncryptionConfig

`func (o *CreateVirtualClusterParameters) GetEncryptionConfig() EncryptionConfiguration`

GetEncryptionConfig returns the EncryptionConfig field if non-nil, zero value otherwise.

### GetEncryptionConfigOk

`func (o *CreateVirtualClusterParameters) GetEncryptionConfigOk() (*EncryptionConfiguration, bool)`

GetEncryptionConfigOk returns a tuple with the EncryptionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionConfig

`func (o *CreateVirtualClusterParameters) SetEncryptionConfig(v EncryptionConfiguration)`

SetEncryptionConfig sets EncryptionConfig field to given value.

### HasEncryptionConfig

`func (o *CreateVirtualClusterParameters) HasEncryptionConfig() bool`

HasEncryptionConfig returns a boolean if a field has been set.

### GetIpPreference

`func (o *CreateVirtualClusterParameters) GetIpPreference() int32`

GetIpPreference returns the IpPreference field if non-nil, zero value otherwise.

### GetIpPreferenceOk

`func (o *CreateVirtualClusterParameters) GetIpPreferenceOk() (*int32, bool)`

GetIpPreferenceOk returns a tuple with the IpPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPreference

`func (o *CreateVirtualClusterParameters) SetIpPreference(v int32)`

SetIpPreference sets IpPreference field to given value.

### HasIpPreference

`func (o *CreateVirtualClusterParameters) HasIpPreference() bool`

HasIpPreference returns a boolean if a field has been set.

### SetIpPreferenceNil

`func (o *CreateVirtualClusterParameters) SetIpPreferenceNil(b bool)`

 SetIpPreferenceNil sets the value for IpPreference to be an explicit nil

### UnsetIpPreference
`func (o *CreateVirtualClusterParameters) UnsetIpPreference()`

UnsetIpPreference ensures that no value is present for IpPreference, not even an explicit nil
### GetMetadataFaultTolerance

`func (o *CreateVirtualClusterParameters) GetMetadataFaultTolerance() int32`

GetMetadataFaultTolerance returns the MetadataFaultTolerance field if non-nil, zero value otherwise.

### GetMetadataFaultToleranceOk

`func (o *CreateVirtualClusterParameters) GetMetadataFaultToleranceOk() (*int32, bool)`

GetMetadataFaultToleranceOk returns a tuple with the MetadataFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFaultTolerance

`func (o *CreateVirtualClusterParameters) SetMetadataFaultTolerance(v int32)`

SetMetadataFaultTolerance sets MetadataFaultTolerance field to given value.

### HasMetadataFaultTolerance

`func (o *CreateVirtualClusterParameters) HasMetadataFaultTolerance() bool`

HasMetadataFaultTolerance returns a boolean if a field has been set.

### SetMetadataFaultToleranceNil

`func (o *CreateVirtualClusterParameters) SetMetadataFaultToleranceNil(b bool)`

 SetMetadataFaultToleranceNil sets the value for MetadataFaultTolerance to be an explicit nil

### UnsetMetadataFaultTolerance
`func (o *CreateVirtualClusterParameters) UnsetMetadataFaultTolerance()`

UnsetMetadataFaultTolerance ensures that no value is present for MetadataFaultTolerance, not even an explicit nil
### GetNetworkConfig

`func (o *CreateVirtualClusterParameters) GetNetworkConfig() NetworkConfiguration`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *CreateVirtualClusterParameters) GetNetworkConfigOk() (*NetworkConfiguration, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *CreateVirtualClusterParameters) SetNetworkConfig(v NetworkConfiguration)`

SetNetworkConfig sets NetworkConfig field to given value.


### GetNodeConfigs

`func (o *CreateVirtualClusterParameters) GetNodeConfigs() []VirtualNodeConfiguration`

GetNodeConfigs returns the NodeConfigs field if non-nil, zero value otherwise.

### GetNodeConfigsOk

`func (o *CreateVirtualClusterParameters) GetNodeConfigsOk() (*[]VirtualNodeConfiguration, bool)`

GetNodeConfigsOk returns a tuple with the NodeConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConfigs

`func (o *CreateVirtualClusterParameters) SetNodeConfigs(v []VirtualNodeConfiguration)`

SetNodeConfigs sets NodeConfigs field to given value.


### SetNodeConfigsNil

`func (o *CreateVirtualClusterParameters) SetNodeConfigsNil(b bool)`

 SetNodeConfigsNil sets the value for NodeConfigs to be an explicit nil

### UnsetNodeConfigs
`func (o *CreateVirtualClusterParameters) UnsetNodeConfigs()`

UnsetNodeConfigs ensures that no value is present for NodeConfigs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


