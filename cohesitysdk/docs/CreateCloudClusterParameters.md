# CreateCloudClusterParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **NullableString** | Specifies the name of the new Cluster. | 
**EncryptionConfig** | Pointer to [**EncryptionConfiguration**](EncryptionConfiguration.md) |  | [optional] 
**IpPreference** | Pointer to **NullableInt32** | Specifies IP preference. | [optional] 
**MetadataFaultTolerance** | Pointer to **NullableInt32** | Specifies the metadata fault tolerance. | [optional] 
**NetworkConfig** | [**CloudNetworkConfiguration**](CloudNetworkConfiguration.md) |  | 
**NodeIps** | **[]string** | Specifies the configuration for the nodes in the new cluster. | 

## Methods

### NewCreateCloudClusterParameters

`func NewCreateCloudClusterParameters(clusterName NullableString, networkConfig CloudNetworkConfiguration, nodeIps []string, ) *CreateCloudClusterParameters`

NewCreateCloudClusterParameters instantiates a new CreateCloudClusterParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCloudClusterParametersWithDefaults

`func NewCreateCloudClusterParametersWithDefaults() *CreateCloudClusterParameters`

NewCreateCloudClusterParametersWithDefaults instantiates a new CreateCloudClusterParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *CreateCloudClusterParameters) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CreateCloudClusterParameters) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CreateCloudClusterParameters) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### SetClusterNameNil

`func (o *CreateCloudClusterParameters) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *CreateCloudClusterParameters) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetEncryptionConfig

`func (o *CreateCloudClusterParameters) GetEncryptionConfig() EncryptionConfiguration`

GetEncryptionConfig returns the EncryptionConfig field if non-nil, zero value otherwise.

### GetEncryptionConfigOk

`func (o *CreateCloudClusterParameters) GetEncryptionConfigOk() (*EncryptionConfiguration, bool)`

GetEncryptionConfigOk returns a tuple with the EncryptionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionConfig

`func (o *CreateCloudClusterParameters) SetEncryptionConfig(v EncryptionConfiguration)`

SetEncryptionConfig sets EncryptionConfig field to given value.

### HasEncryptionConfig

`func (o *CreateCloudClusterParameters) HasEncryptionConfig() bool`

HasEncryptionConfig returns a boolean if a field has been set.

### GetIpPreference

`func (o *CreateCloudClusterParameters) GetIpPreference() int32`

GetIpPreference returns the IpPreference field if non-nil, zero value otherwise.

### GetIpPreferenceOk

`func (o *CreateCloudClusterParameters) GetIpPreferenceOk() (*int32, bool)`

GetIpPreferenceOk returns a tuple with the IpPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPreference

`func (o *CreateCloudClusterParameters) SetIpPreference(v int32)`

SetIpPreference sets IpPreference field to given value.

### HasIpPreference

`func (o *CreateCloudClusterParameters) HasIpPreference() bool`

HasIpPreference returns a boolean if a field has been set.

### SetIpPreferenceNil

`func (o *CreateCloudClusterParameters) SetIpPreferenceNil(b bool)`

 SetIpPreferenceNil sets the value for IpPreference to be an explicit nil

### UnsetIpPreference
`func (o *CreateCloudClusterParameters) UnsetIpPreference()`

UnsetIpPreference ensures that no value is present for IpPreference, not even an explicit nil
### GetMetadataFaultTolerance

`func (o *CreateCloudClusterParameters) GetMetadataFaultTolerance() int32`

GetMetadataFaultTolerance returns the MetadataFaultTolerance field if non-nil, zero value otherwise.

### GetMetadataFaultToleranceOk

`func (o *CreateCloudClusterParameters) GetMetadataFaultToleranceOk() (*int32, bool)`

GetMetadataFaultToleranceOk returns a tuple with the MetadataFaultTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFaultTolerance

`func (o *CreateCloudClusterParameters) SetMetadataFaultTolerance(v int32)`

SetMetadataFaultTolerance sets MetadataFaultTolerance field to given value.

### HasMetadataFaultTolerance

`func (o *CreateCloudClusterParameters) HasMetadataFaultTolerance() bool`

HasMetadataFaultTolerance returns a boolean if a field has been set.

### SetMetadataFaultToleranceNil

`func (o *CreateCloudClusterParameters) SetMetadataFaultToleranceNil(b bool)`

 SetMetadataFaultToleranceNil sets the value for MetadataFaultTolerance to be an explicit nil

### UnsetMetadataFaultTolerance
`func (o *CreateCloudClusterParameters) UnsetMetadataFaultTolerance()`

UnsetMetadataFaultTolerance ensures that no value is present for MetadataFaultTolerance, not even an explicit nil
### GetNetworkConfig

`func (o *CreateCloudClusterParameters) GetNetworkConfig() CloudNetworkConfiguration`

GetNetworkConfig returns the NetworkConfig field if non-nil, zero value otherwise.

### GetNetworkConfigOk

`func (o *CreateCloudClusterParameters) GetNetworkConfigOk() (*CloudNetworkConfiguration, bool)`

GetNetworkConfigOk returns a tuple with the NetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfig

`func (o *CreateCloudClusterParameters) SetNetworkConfig(v CloudNetworkConfiguration)`

SetNetworkConfig sets NetworkConfig field to given value.


### GetNodeIps

`func (o *CreateCloudClusterParameters) GetNodeIps() []string`

GetNodeIps returns the NodeIps field if non-nil, zero value otherwise.

### GetNodeIpsOk

`func (o *CreateCloudClusterParameters) GetNodeIpsOk() (*[]string, bool)`

GetNodeIpsOk returns a tuple with the NodeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIps

`func (o *CreateCloudClusterParameters) SetNodeIps(v []string)`

SetNodeIps sets NodeIps field to given value.


### SetNodeIpsNil

`func (o *CreateCloudClusterParameters) SetNodeIpsNil(b bool)`

 SetNodeIpsNil sets the value for NodeIps to be an explicit nil

### UnsetNodeIps
`func (o *CreateCloudClusterParameters) UnsetNodeIps()`

UnsetNodeIps ensures that no value is present for NodeIps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


