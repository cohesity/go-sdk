# WorkloadStatsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableString** | Specifies the Id of an Entity. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of an Entity. | [optional] 
**Entities** | Pointer to [**[]EntityIdentifier**](EntityIdentifier.md) | Specifies the entities part of Workload schema. | [optional] 
**Schema** | Pointer to **NullableString** | Specifies the Schema Name of Workload. | [optional] 
**SubTypes** | Pointer to [**[]WorkloadSubType**](WorkloadSubType.md) | Specifies the Workload Sub-Types. | [optional] 

## Methods

### NewWorkloadStatsSchema

`func NewWorkloadStatsSchema() *WorkloadStatsSchema`

NewWorkloadStatsSchema instantiates a new WorkloadStatsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadStatsSchemaWithDefaults

`func NewWorkloadStatsSchemaWithDefaults() *WorkloadStatsSchema`

NewWorkloadStatsSchemaWithDefaults instantiates a new WorkloadStatsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *WorkloadStatsSchema) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *WorkloadStatsSchema) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *WorkloadStatsSchema) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *WorkloadStatsSchema) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *WorkloadStatsSchema) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *WorkloadStatsSchema) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetName

`func (o *WorkloadStatsSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadStatsSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadStatsSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadStatsSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkloadStatsSchema) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkloadStatsSchema) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEntities

`func (o *WorkloadStatsSchema) GetEntities() []EntityIdentifier`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *WorkloadStatsSchema) GetEntitiesOk() (*[]EntityIdentifier, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *WorkloadStatsSchema) SetEntities(v []EntityIdentifier)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *WorkloadStatsSchema) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### SetEntitiesNil

`func (o *WorkloadStatsSchema) SetEntitiesNil(b bool)`

 SetEntitiesNil sets the value for Entities to be an explicit nil

### UnsetEntities
`func (o *WorkloadStatsSchema) UnsetEntities()`

UnsetEntities ensures that no value is present for Entities, not even an explicit nil
### GetSchema

`func (o *WorkloadStatsSchema) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *WorkloadStatsSchema) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *WorkloadStatsSchema) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *WorkloadStatsSchema) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *WorkloadStatsSchema) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *WorkloadStatsSchema) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetSubTypes

`func (o *WorkloadStatsSchema) GetSubTypes() []WorkloadSubType`

GetSubTypes returns the SubTypes field if non-nil, zero value otherwise.

### GetSubTypesOk

`func (o *WorkloadStatsSchema) GetSubTypesOk() (*[]WorkloadSubType, bool)`

GetSubTypesOk returns a tuple with the SubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTypes

`func (o *WorkloadStatsSchema) SetSubTypes(v []WorkloadSubType)`

SetSubTypes sets SubTypes field to given value.

### HasSubTypes

`func (o *WorkloadStatsSchema) HasSubTypes() bool`

HasSubTypes returns a boolean if a field has been set.

### SetSubTypesNil

`func (o *WorkloadStatsSchema) SetSubTypesNil(b bool)`

 SetSubTypesNil sets the value for SubTypes to be an explicit nil

### UnsetSubTypes
`func (o *WorkloadStatsSchema) UnsetSubTypes()`

UnsetSubTypes ensures that no value is present for SubTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


