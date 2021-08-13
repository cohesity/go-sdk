# ClusterInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]NodeInterfaces**](NodeInterfaces.md) | Specifies the list of nodes present in the cluster. If the request was sent to a free node, then this list will have exactly one entry. | [optional] 

## Methods

### NewClusterInterfaces

`func NewClusterInterfaces() *ClusterInterfaces`

NewClusterInterfaces instantiates a new ClusterInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInterfacesWithDefaults

`func NewClusterInterfacesWithDefaults() *ClusterInterfaces`

NewClusterInterfacesWithDefaults instantiates a new ClusterInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ClusterInterfaces) GetNodes() []NodeInterfaces`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterInterfaces) GetNodesOk() (*[]NodeInterfaces, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterInterfaces) SetNodes(v []NodeInterfaces)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterInterfaces) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


