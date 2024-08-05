# Workload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthStatus** | Pointer to **string** | Kubernetes Workload Cluster Overall Health Status | [optional] 
**Nodes** | Pointer to [**[]NodeDef**](NodeDef.md) | Array of Kubernetes Workload Cluster Nodes | [optional] 
**State** | Pointer to **string** | Kubernetes Workload Cluster Overall State | [optional] 

## Methods

### NewWorkload

`func NewWorkload() *Workload`

NewWorkload instantiates a new Workload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWithDefaults

`func NewWorkloadWithDefaults() *Workload`

NewWorkloadWithDefaults instantiates a new Workload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthStatus

`func (o *Workload) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *Workload) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *Workload) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *Workload) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetNodes

`func (o *Workload) GetNodes() []NodeDef`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *Workload) GetNodesOk() (*[]NodeDef, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *Workload) SetNodes(v []NodeDef)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *Workload) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetState

`func (o *Workload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Workload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Workload) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Workload) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


