# ClusterConfigProtoQoSMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalId** | Pointer to **NullableInt64** | Principal id of the QoS principal to which qos_context maps to. | [optional] 
**QosContext** | Pointer to [**ClusterConfigProtoQoSMappingQoSContext**](ClusterConfigProtoQoSMappingQoSContext.md) |  | [optional] 

## Methods

### NewClusterConfigProtoQoSMapping

`func NewClusterConfigProtoQoSMapping() *ClusterConfigProtoQoSMapping`

NewClusterConfigProtoQoSMapping instantiates a new ClusterConfigProtoQoSMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigProtoQoSMappingWithDefaults

`func NewClusterConfigProtoQoSMappingWithDefaults() *ClusterConfigProtoQoSMapping`

NewClusterConfigProtoQoSMappingWithDefaults instantiates a new ClusterConfigProtoQoSMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalId

`func (o *ClusterConfigProtoQoSMapping) GetPrincipalId() int64`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *ClusterConfigProtoQoSMapping) GetPrincipalIdOk() (*int64, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *ClusterConfigProtoQoSMapping) SetPrincipalId(v int64)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *ClusterConfigProtoQoSMapping) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *ClusterConfigProtoQoSMapping) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *ClusterConfigProtoQoSMapping) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetQosContext

`func (o *ClusterConfigProtoQoSMapping) GetQosContext() ClusterConfigProtoQoSMappingQoSContext`

GetQosContext returns the QosContext field if non-nil, zero value otherwise.

### GetQosContextOk

`func (o *ClusterConfigProtoQoSMapping) GetQosContextOk() (*ClusterConfigProtoQoSMappingQoSContext, bool)`

GetQosContextOk returns a tuple with the QosContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosContext

`func (o *ClusterConfigProtoQoSMapping) SetQosContext(v ClusterConfigProtoQoSMappingQoSContext)`

SetQosContext sets QosContext field to given value.

### HasQosContext

`func (o *ClusterConfigProtoQoSMapping) HasQosContext() bool`

HasQosContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


