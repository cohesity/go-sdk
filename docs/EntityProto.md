# EntityProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeVec** | Pointer to [**[]KeyValuePair**](KeyValuePair.md) | Array of Attributes.  List of attributes of an entity. | [optional] 
**EntityId** | Pointer to [**EntityIdentifier**](EntityIdentifier.md) |  | [optional] 
**LatestMetricVec** | Pointer to [**[]MetricValue**](MetricValue.md) | Array of Metric Statistics.  List of the latest statistics for all metrics defined in the schema that this entity belongs to. If statistics for a metric is not available, then that data point is not returned. | [optional] 

## Methods

### NewEntityProto

`func NewEntityProto() *EntityProto`

NewEntityProto instantiates a new EntityProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityProtoWithDefaults

`func NewEntityProtoWithDefaults() *EntityProto`

NewEntityProtoWithDefaults instantiates a new EntityProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeVec

`func (o *EntityProto) GetAttributeVec() []KeyValuePair`

GetAttributeVec returns the AttributeVec field if non-nil, zero value otherwise.

### GetAttributeVecOk

`func (o *EntityProto) GetAttributeVecOk() (*[]KeyValuePair, bool)`

GetAttributeVecOk returns a tuple with the AttributeVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeVec

`func (o *EntityProto) SetAttributeVec(v []KeyValuePair)`

SetAttributeVec sets AttributeVec field to given value.

### HasAttributeVec

`func (o *EntityProto) HasAttributeVec() bool`

HasAttributeVec returns a boolean if a field has been set.

### SetAttributeVecNil

`func (o *EntityProto) SetAttributeVecNil(b bool)`

 SetAttributeVecNil sets the value for AttributeVec to be an explicit nil

### UnsetAttributeVec
`func (o *EntityProto) UnsetAttributeVec()`

UnsetAttributeVec ensures that no value is present for AttributeVec, not even an explicit nil
### GetEntityId

`func (o *EntityProto) GetEntityId() EntityIdentifier`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityProto) GetEntityIdOk() (*EntityIdentifier, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityProto) SetEntityId(v EntityIdentifier)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *EntityProto) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetLatestMetricVec

`func (o *EntityProto) GetLatestMetricVec() []MetricValue`

GetLatestMetricVec returns the LatestMetricVec field if non-nil, zero value otherwise.

### GetLatestMetricVecOk

`func (o *EntityProto) GetLatestMetricVecOk() (*[]MetricValue, bool)`

GetLatestMetricVecOk returns a tuple with the LatestMetricVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMetricVec

`func (o *EntityProto) SetLatestMetricVec(v []MetricValue)`

SetLatestMetricVec sets LatestMetricVec field to given value.

### HasLatestMetricVec

`func (o *EntityProto) HasLatestMetricVec() bool`

HasLatestMetricVec returns a boolean if a field has been set.

### SetLatestMetricVecNil

`func (o *EntityProto) SetLatestMetricVecNil(b bool)`

 SetLatestMetricVecNil sets the value for LatestMetricVec to be an explicit nil

### UnsetLatestMetricVec
`func (o *EntityProto) UnsetLatestMetricVec()`

UnsetLatestMetricVec ensures that no value is present for LatestMetricVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


