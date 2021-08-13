# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **NullableString** | Specifies the name of schema metric. | [optional] 
**SchemaName** | Pointer to **NullableString** | Specifies the name of the schema. | [optional] 
**EntityId** | Pointer to **NullableString** | Specifies the id of the entity. | [optional] 

## Methods

### NewSchema

`func NewSchema() *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *Schema) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *Schema) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *Schema) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *Schema) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *Schema) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *Schema) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetSchemaName

`func (o *Schema) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *Schema) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *Schema) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *Schema) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### SetSchemaNameNil

`func (o *Schema) SetSchemaNameNil(b bool)`

 SetSchemaNameNil sets the value for SchemaName to be an explicit nil

### UnsetSchemaName
`func (o *Schema) UnsetSchemaName()`

UnsetSchemaName ensures that no value is present for SchemaName, not even an explicit nil
### GetEntityId

`func (o *Schema) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Schema) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Schema) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Schema) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *Schema) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *Schema) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


