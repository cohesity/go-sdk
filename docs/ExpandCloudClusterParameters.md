# ExpandCloudClusterParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIps** | **[]string** | Specifies the list of IPs of the new Nodes. | 

## Methods

### NewExpandCloudClusterParameters

`func NewExpandCloudClusterParameters(nodeIps []string, ) *ExpandCloudClusterParameters`

NewExpandCloudClusterParameters instantiates a new ExpandCloudClusterParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandCloudClusterParametersWithDefaults

`func NewExpandCloudClusterParametersWithDefaults() *ExpandCloudClusterParameters`

NewExpandCloudClusterParametersWithDefaults instantiates a new ExpandCloudClusterParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIps

`func (o *ExpandCloudClusterParameters) GetNodeIps() []string`

GetNodeIps returns the NodeIps field if non-nil, zero value otherwise.

### GetNodeIpsOk

`func (o *ExpandCloudClusterParameters) GetNodeIpsOk() (*[]string, bool)`

GetNodeIpsOk returns a tuple with the NodeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIps

`func (o *ExpandCloudClusterParameters) SetNodeIps(v []string)`

SetNodeIps sets NodeIps field to given value.


### SetNodeIpsNil

`func (o *ExpandCloudClusterParameters) SetNodeIpsNil(b bool)`

 SetNodeIpsNil sets the value for NodeIps to be an explicit nil

### UnsetNodeIps
`func (o *ExpandCloudClusterParameters) UnsetNodeIps()`

UnsetNodeIps ensures that no value is present for NodeIps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


