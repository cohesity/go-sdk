# CreatePhysicalClusterParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **NullableString** | Specifies the name of the new Cluster. | 
**EncryptionConfig** | Pointer to [**EncryptionConfiguration**](EncryptionConfiguration.md) |  | [optional] 
**IpPreference** | Pointer to **NullableInt32** | Specifies IP preference. | [optional] 
**IpmiConfig** | [**IpmiConfiguration**](IpmiConfiguration.md) |  | 
**MetadataFaultTolerance** | Pointer to **NullableInt32** | Specifies the metadata fault tolerance. | [optional] 
**NetworkConfig** | [**NetworkConfiguration**](NetworkConfiguration.md) |  | 
**NodeConfigs** | [**[]PhysicalNodeConfiguration**](PhysicalNodeConfiguration.md) | Specifies the configuration for the nodes in the new cluster. | 

## Methods

### NewCreatePhysicalClusterParameters

`func NewCreatePhysicalClusterParameters(clusterName NullableString, ipmiConfig IpmiConfiguration, networkConfig NetworkConfiguration, nodeConfigs []PhysicalNodeConfiguration, ) *CreatePhysicalClusterParameters`

NewCreatePhysicalClusterParameters instantiates a new CreatePhysicalClusterParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePhysicalClusterParametersWithDefaults

`func NewCreatePhysicalClusterParametersWithDefaults() *CreatePhysicalClusterParameters`

NewCreatePhysicalClusterParametersWithDefaults instantiates a new CreatePhysicalClusterParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *CreatePhysicalClusterParameters) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CreatePhysicalClusterParameters) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CreatePhysicalClusterParameters) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### SetClusterNameNil

`func (o *CreatePhysicalClusterParameters) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *CreatePhysicalClusterParameters) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetEncryptionConfig

`func (o *CreatePhysicalClusterParameters) GetEncryptionConfig() EncryptionConfiguration`

GetEncryptionConfig returns the EncryptionConfig field if non-nil, zero value otherwise.

### GetEncryptionConfigOk

`func (o *CreatePhysicalClusterParameters) GetEncryptionConfigOk() (*EncryptionConfiguration, bool)`

GetEncryptionConfigOk returns a tuple with the EncryptionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionConfig

`func (o *CreatePhysicalClusterParameters) SetEncryptionConfig(v EncryptionConfiguration)`

SetEncryptionConfig sets EncryptionConfig field to given value.

### HasEncryptionConfig

`func (o *CreatePhysicalClusterParameters) HasEncryptionConfig() bool`

HasEncryptionConfig returns a boolean if a field has been set.

### GetIpPreference

`func (o *CreatePhysicalClusterParameters) GetIpPreference() int32`

GetIpPreference returns the IpPreference field if non-nil, zero value otherwise.

### GetIpPreferenceOk

`func (o *CreatePhysicalClusterParameters) GetIpPreferenceOk() (*int32, bool)`

GetIpPreferenceOk returns a tuple with the IpPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPreference

`func (o *CreatePhysicalClusterParameters) SetIpPreference(v int32)`

SetIpPreference sets IpPreference field to given value.

### HasIpPreference

`func (o *CreatePhysicalClusterParameters) HasIpPreference() bool`

HasIpPreference returns a boolean if a field has been set.

### SetIpPreferenceNil

`func (o *CreatePhysicalClusterParameters) SetIpPreferenceNil(b bool)`

 SetIpPreferenceNil sets the value for IpPreference to be an explicit nil

### UnsetIpPreference
`func (o *CreatePhysicalClusterParameters) UnsetIpPreference()`

UnsetIpPreference ensures that no value is present for IpPreference, not even an explicit nil
### GetIpmiConfig

`func (o *CreatePhysicalClusterParameters) GetIpmiConfig() IpmiConfiguration`

GetIpmiConfig returns the IpmiConfig field if non-nil, zero value otherwise.

### GetIpmiConfigOk

`func (o *CreatePhysicalClusterParameters) GetIpmiConfigOk() (*IpmiConfiguration, bool)`

GetIpmiConfigOk returns a tuple with the IpmiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiConfig

`func (o *CreatePhysicalClusterParameters) SetIpmiConfig(v IpmiConfiguration)`

SetIpmiConfig sets IpmiConfig field to given value.


### GetMetadataFaultTolerance

`func (o *CreatePhysicalClusterParameters) GetMetadataFaultTolerance() int32`

GetMetadataFaultTolerance returns the MetadataFaultTolerance field if non-nil, zero value otherwise.

### GetMetadataFaultToleranceOk

`func (o *CreatePhysicalClusterParameters) GetMetadataFaultToleranceOk() (*int32, bool)`

GetMetadataFaultToleranceOk returns a tuple with the MetadataFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFaultTolerance

`func (o *CreatePhysicalClusterParameters) SetMetadataFaultTolerance(v int32)`

SetMetadataFaultTolerance sets MetadataFaultTolerance field to given value.

### HasMetadataFaultTolerance

`func (o *CreatePhysicalClusterParameters) HasMetadataFaultTolerance() bool`

HasMetadataFaultTolerance returns a boolean if a field has been set.

### SetMetadataFaultToleranceNil

`func (o *CreatePhysicalClusterParameters) SetMetadataFaultToleranceNil(b bool)`

 SetMetadataFaultToleranceNil sets the value for MetadataFaultTolerance to be an explicit nil

### UnsetMetadataFaultTolerance
`func (o *CreatePhysicalClusterParameters) UnsetMetadataFaultTolerance()`

UnsetMetadataFaultTolerance ensures that no value is present for MetadataFaultTolerance, not even an explicit nil
### GetNetworkConfig

`func (o *CreatePhysicalClusterParameters) GetNetworkConfig() NetworkConfiguration`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *CreatePhysicalClusterParameters) GetNetworkConfigOk() (*NetworkConfiguration, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *CreatePhysicalClusterParameters) SetNetworkConfig(v NetworkConfiguration)`

SetNetworkConfig sets NetworkConfig field to given value.


### GetNodeConfigs

`func (o *CreatePhysicalClusterParameters) GetNodeConfigs() []PhysicalNodeConfiguration`

GetNodeConfigs returns the NodeConfigs field if non-nil, zero value otherwise.

### GetNodeConfigsOk

`func (o *CreatePhysicalClusterParameters) GetNodeConfigsOk() (*[]PhysicalNodeConfiguration, bool)`

GetNodeConfigsOk returns a tuple with the NodeConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConfigs

`func (o *CreatePhysicalClusterParameters) SetNodeConfigs(v []PhysicalNodeConfiguration)`

SetNodeConfigs sets NodeConfigs field to given value.


### SetNodeConfigsNil

`func (o *CreatePhysicalClusterParameters) SetNodeConfigsNil(b bool)`

 SetNodeConfigsNil sets the value for NodeConfigs to be an explicit nil

### UnsetNodeConfigs
`func (o *CreatePhysicalClusterParameters) UnsetNodeConfigs()`

UnsetNodeConfigs ensures that no value is present for NodeConfigs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


