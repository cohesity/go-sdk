# CreateClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **NullableInt64** | Specifies the ID of the new Cluster. | [optional] 
**ClusterName** | Pointer to **NullableString** | Specifies the name of the new Cluster. | [optional] 
**ClusterSwVersion** | Pointer to **NullableString** | Specifies the software version of the new Cluster. | [optional] 
**HealthyNodes** | Pointer to [**[]NodeStatus**](NodeStatus.md) | Specifies the status of the Nodes in the Cluster. All Nodes that are accepted to the Cluster are appended to this list. | [optional] 
**IncarnationId** | Pointer to **NullableInt64** | Specifies the Incarnation ID of the new Cluster. | [optional] 
**Message** | Pointer to **NullableString** | Specifies an optional message field. | [optional] 
**UnhealthyNodes** | Pointer to [**[]NodeStatus**](NodeStatus.md) | Specifies the status of the Nodes in the Cluster. All Nodes that are not accepted to the Cluster are appended to this list. | [optional] 

## Methods

### NewCreateClusterResult

`func NewCreateClusterResult() *CreateClusterResult`

NewCreateClusterResult instantiates a new CreateClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterResultWithDefaults

`func NewCreateClusterResultWithDefaults() *CreateClusterResult`

NewCreateClusterResultWithDefaults instantiates a new CreateClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *CreateClusterResult) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *CreateClusterResult) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *CreateClusterResult) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *CreateClusterResult) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *CreateClusterResult) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *CreateClusterResult) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterName

`func (o *CreateClusterResult) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *CreateClusterResult) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *CreateClusterResult) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *CreateClusterResult) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *CreateClusterResult) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *CreateClusterResult) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetClusterSwVersion

`func (o *CreateClusterResult) GetClusterSwVersion() string`

GetClusterSwVersion returns the ClusterSwVersion field if non-nil, zero value otherwise.

### GetClusterSwVersionOk

`func (o *CreateClusterResult) GetClusterSwVersionOk() (*string, bool)`

GetClusterSwVersionOk returns a tuple with the ClusterSwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSwVersion

`func (o *CreateClusterResult) SetClusterSwVersion(v string)`

SetClusterSwVersion sets ClusterSwVersion field to given value.

### HasClusterSwVersion

`func (o *CreateClusterResult) HasClusterSwVersion() bool`

HasClusterSwVersion returns a boolean if a field has been set.

### SetClusterSwVersionNil

`func (o *CreateClusterResult) SetClusterSwVersionNil(b bool)`

 SetClusterSwVersionNil sets the value for ClusterSwVersion to be an explicit nil

### UnsetClusterSwVersion
`func (o *CreateClusterResult) UnsetClusterSwVersion()`

UnsetClusterSwVersion ensures that no value is present for ClusterSwVersion, not even an explicit nil
### GetHealthyNodes

`func (o *CreateClusterResult) GetHealthyNodes() []NodeStatus`

GetHealthyNodes returns the HealthyNodes field if non-nil, zero value otherwise.

### GetHealthyNodesOk

`func (o *CreateClusterResult) GetHealthyNodesOk() (*[]NodeStatus, bool)`

GetHealthyNodesOk returns a tuple with the HealthyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyNodes

`func (o *CreateClusterResult) SetHealthyNodes(v []NodeStatus)`

SetHealthyNodes sets HealthyNodes field to given value.

### HasHealthyNodes

`func (o *CreateClusterResult) HasHealthyNodes() bool`

HasHealthyNodes returns a boolean if a field has been set.

### SetHealthyNodesNil

`func (o *CreateClusterResult) SetHealthyNodesNil(b bool)`

 SetHealthyNodesNil sets the value for HealthyNodes to be an explicit nil

### UnsetHealthyNodes
`func (o *CreateClusterResult) UnsetHealthyNodes()`

UnsetHealthyNodes ensures that no value is present for HealthyNodes, not even an explicit nil
### GetIncarnationId

`func (o *CreateClusterResult) GetIncarnationId() int64`

GetIncarnationId returns the IncarnationId field if non-nil, zero value otherwise.

### GetIncarnationIdOk

`func (o *CreateClusterResult) GetIncarnationIdOk() (*int64, bool)`

GetIncarnationIdOk returns a tuple with the IncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncarnationId

`func (o *CreateClusterResult) SetIncarnationId(v int64)`

SetIncarnationId sets IncarnationId field to given value.

### HasIncarnationId

`func (o *CreateClusterResult) HasIncarnationId() bool`

HasIncarnationId returns a boolean if a field has been set.

### SetIncarnationIdNil

`func (o *CreateClusterResult) SetIncarnationIdNil(b bool)`

 SetIncarnationIdNil sets the value for IncarnationId to be an explicit nil

### UnsetIncarnationId
`func (o *CreateClusterResult) UnsetIncarnationId()`

UnsetIncarnationId ensures that no value is present for IncarnationId, not even an explicit nil
### GetMessage

`func (o *CreateClusterResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateClusterResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateClusterResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateClusterResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *CreateClusterResult) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *CreateClusterResult) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetUnhealthyNodes

`func (o *CreateClusterResult) GetUnhealthyNodes() []NodeStatus`

GetUnhealthyNodes returns the UnhealthyNodes field if non-nil, zero value otherwise.

### GetUnhealthyNodesOk

`func (o *CreateClusterResult) GetUnhealthyNodesOk() (*[]NodeStatus, bool)`

GetUnhealthyNodesOk returns a tuple with the UnhealthyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyNodes

`func (o *CreateClusterResult) SetUnhealthyNodes(v []NodeStatus)`

SetUnhealthyNodes sets UnhealthyNodes field to given value.

### HasUnhealthyNodes

`func (o *CreateClusterResult) HasUnhealthyNodes() bool`

HasUnhealthyNodes returns a boolean if a field has been set.

### SetUnhealthyNodesNil

`func (o *CreateClusterResult) SetUnhealthyNodesNil(b bool)`

 SetUnhealthyNodesNil sets the value for UnhealthyNodes to be an explicit nil

### UnsetUnhealthyNodes
`func (o *CreateClusterResult) UnsetUnhealthyNodes()`

UnsetUnhealthyNodes ensures that no value is present for UnhealthyNodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


