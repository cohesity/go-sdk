# SchemaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableString** | Specifies the id of the entity represented as a string. | [optional] 
**Key** | Pointer to **NullableString** | Specifies the key which is public facing name for metric name. | [optional] 
**MetricName** | Pointer to **NullableString** | Specifies the Apollo schema metric name. | [optional] 
**SchemaName** | Pointer to **NullableString** | Specifies the name of entity schema such as &#39;ApolloViewBoxStats&#39;. | [optional] 

## Methods

### NewSchemaInfo

`func NewSchemaInfo() *SchemaInfo`

NewSchemaInfo instantiates a new SchemaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaInfoWithDefaults

`func NewSchemaInfoWithDefaults() *SchemaInfo`

NewSchemaInfoWithDefaults instantiates a new SchemaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *SchemaInfo) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SchemaInfo) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SchemaInfo) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SchemaInfo) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *SchemaInfo) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *SchemaInfo) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetKey

`func (o *SchemaInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SchemaInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SchemaInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SchemaInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *SchemaInfo) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SchemaInfo) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetMetricName

`func (o *SchemaInfo) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *SchemaInfo) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *SchemaInfo) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *SchemaInfo) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *SchemaInfo) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *SchemaInfo) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetSchemaName

`func (o *SchemaInfo) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *SchemaInfo) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *SchemaInfo) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *SchemaInfo) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### SetSchemaNameNil

`func (o *SchemaInfo) SetSchemaNameNil(b bool)`

 SetSchemaNameNil sets the value for SchemaName to be an explicit nil

### UnsetSchemaName
`func (o *SchemaInfo) UnsetSchemaName()`

UnsetSchemaName ensures that no value is present for SchemaName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


