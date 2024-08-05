# ClusterCreatePhysicalParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]ClusterCreateNodeParams**](ClusterCreateNodeParams.md) |  | [optional] 

## Methods

### NewClusterCreatePhysicalParams

`func NewClusterCreatePhysicalParams() *ClusterCreatePhysicalParams`

NewClusterCreatePhysicalParams instantiates a new ClusterCreatePhysicalParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreatePhysicalParamsWithDefaults

`func NewClusterCreatePhysicalParamsWithDefaults() *ClusterCreatePhysicalParams`

NewClusterCreatePhysicalParamsWithDefaults instantiates a new ClusterCreatePhysicalParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ClusterCreatePhysicalParams) GetNodes() []ClusterCreateNodeParams`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterCreatePhysicalParams) GetNodesOk() (*[]ClusterCreateNodeParams, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterCreatePhysicalParams) SetNodes(v []ClusterCreateNodeParams)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterCreatePhysicalParams) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


