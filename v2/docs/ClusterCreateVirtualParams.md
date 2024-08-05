# ClusterCreateVirtualParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]ClusterCreateNodeParams**](ClusterCreateNodeParams.md) |  | [optional] 

## Methods

### NewClusterCreateVirtualParams

`func NewClusterCreateVirtualParams() *ClusterCreateVirtualParams`

NewClusterCreateVirtualParams instantiates a new ClusterCreateVirtualParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateVirtualParamsWithDefaults

`func NewClusterCreateVirtualParamsWithDefaults() *ClusterCreateVirtualParams`

NewClusterCreateVirtualParamsWithDefaults instantiates a new ClusterCreateVirtualParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ClusterCreateVirtualParams) GetNodes() []ClusterCreateNodeParams`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterCreateVirtualParams) GetNodesOk() (*[]ClusterCreateNodeParams, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterCreateVirtualParams) SetNodes(v []ClusterCreateNodeParams)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterCreateVirtualParams) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


