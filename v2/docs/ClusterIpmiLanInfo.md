# ClusterIpmiLanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIpmiGateway** | Pointer to **NullableString** | Specifies the gateway for the given cluster ipmi lan. | [optional] 
**ClusterIpmiSubnetMask** | Pointer to **NullableString** | Specifies the subnet mask for the given cluster ipmi lan. | [optional] 
**NodeIpmiEntries** | Pointer to [**[]NodeIpmiInfoEntry**](NodeIpmiInfoEntry.md) | Specifies the list of node ipmi info for all the nodes in the cluster. | [optional] 

## Methods

### NewClusterIpmiLanInfo

`func NewClusterIpmiLanInfo() *ClusterIpmiLanInfo`

NewClusterIpmiLanInfo instantiates a new ClusterIpmiLanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterIpmiLanInfoWithDefaults

`func NewClusterIpmiLanInfoWithDefaults() *ClusterIpmiLanInfo`

NewClusterIpmiLanInfoWithDefaults instantiates a new ClusterIpmiLanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIpmiGateway

`func (o *ClusterIpmiLanInfo) GetClusterIpmiGateway() string`

GetClusterIpmiGateway returns the ClusterIpmiGateway field if non-nil, zero value otherwise.

### GetClusterIpmiGatewayOk

`func (o *ClusterIpmiLanInfo) GetClusterIpmiGatewayOk() (*string, bool)`

GetClusterIpmiGatewayOk returns a tuple with the ClusterIpmiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIpmiGateway

`func (o *ClusterIpmiLanInfo) SetClusterIpmiGateway(v string)`

SetClusterIpmiGateway sets ClusterIpmiGateway field to given value.

### HasClusterIpmiGateway

`func (o *ClusterIpmiLanInfo) HasClusterIpmiGateway() bool`

HasClusterIpmiGateway returns a boolean if a field has been set.

### SetClusterIpmiGatewayNil

`func (o *ClusterIpmiLanInfo) SetClusterIpmiGatewayNil(b bool)`

 SetClusterIpmiGatewayNil sets the value for ClusterIpmiGateway to be an explicit nil

### UnsetClusterIpmiGateway
`func (o *ClusterIpmiLanInfo) UnsetClusterIpmiGateway()`

UnsetClusterIpmiGateway ensures that no value is present for ClusterIpmiGateway, not even an explicit nil
### GetClusterIpmiSubnetMask

`func (o *ClusterIpmiLanInfo) GetClusterIpmiSubnetMask() string`

GetClusterIpmiSubnetMask returns the ClusterIpmiSubnetMask field if non-nil, zero value otherwise.

### GetClusterIpmiSubnetMaskOk

`func (o *ClusterIpmiLanInfo) GetClusterIpmiSubnetMaskOk() (*string, bool)`

GetClusterIpmiSubnetMaskOk returns a tuple with the ClusterIpmiSubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIpmiSubnetMask

`func (o *ClusterIpmiLanInfo) SetClusterIpmiSubnetMask(v string)`

SetClusterIpmiSubnetMask sets ClusterIpmiSubnetMask field to given value.

### HasClusterIpmiSubnetMask

`func (o *ClusterIpmiLanInfo) HasClusterIpmiSubnetMask() bool`

HasClusterIpmiSubnetMask returns a boolean if a field has been set.

### SetClusterIpmiSubnetMaskNil

`func (o *ClusterIpmiLanInfo) SetClusterIpmiSubnetMaskNil(b bool)`

 SetClusterIpmiSubnetMaskNil sets the value for ClusterIpmiSubnetMask to be an explicit nil

### UnsetClusterIpmiSubnetMask
`func (o *ClusterIpmiLanInfo) UnsetClusterIpmiSubnetMask()`

UnsetClusterIpmiSubnetMask ensures that no value is present for ClusterIpmiSubnetMask, not even an explicit nil
### GetNodeIpmiEntries

`func (o *ClusterIpmiLanInfo) GetNodeIpmiEntries() []NodeIpmiInfoEntry`

GetNodeIpmiEntries returns the NodeIpmiEntries field if non-nil, zero value otherwise.

### GetNodeIpmiEntriesOk

`func (o *ClusterIpmiLanInfo) GetNodeIpmiEntriesOk() (*[]NodeIpmiInfoEntry, bool)`

GetNodeIpmiEntriesOk returns a tuple with the NodeIpmiEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpmiEntries

`func (o *ClusterIpmiLanInfo) SetNodeIpmiEntries(v []NodeIpmiInfoEntry)`

SetNodeIpmiEntries sets NodeIpmiEntries field to given value.

### HasNodeIpmiEntries

`func (o *ClusterIpmiLanInfo) HasNodeIpmiEntries() bool`

HasNodeIpmiEntries returns a boolean if a field has been set.

### SetNodeIpmiEntriesNil

`func (o *ClusterIpmiLanInfo) SetNodeIpmiEntriesNil(b bool)`

 SetNodeIpmiEntriesNil sets the value for NodeIpmiEntries to be an explicit nil

### UnsetNodeIpmiEntries
`func (o *ClusterIpmiLanInfo) UnsetNodeIpmiEntries()`

UnsetNodeIpmiEntries ensures that no value is present for NodeIpmiEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


