# WorkloadSubType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableString** | Specifies the Id of an Entity. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of an Entity. | [optional] 
**Entities** | Pointer to [**[]EntityIdentifier**](EntityIdentifier.md) | Specifies the entities part of Workload schema. | [optional] 
**Schema** | Pointer to **NullableString** | Specifies the Schema Name of Workload. | [optional] 

## Methods

### NewWorkloadSubType

`func NewWorkloadSubType() *WorkloadSubType`

NewWorkloadSubType instantiates a new WorkloadSubType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadSubTypeWithDefaults

`func NewWorkloadSubTypeWithDefaults() *WorkloadSubType`

NewWorkloadSubTypeWithDefaults instantiates a new WorkloadSubType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *WorkloadSubType) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *WorkloadSubType) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *WorkloadSubType) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *WorkloadSubType) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *WorkloadSubType) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *WorkloadSubType) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetName

`func (o *WorkloadSubType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadSubType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadSubType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadSubType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WorkloadSubType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WorkloadSubType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEntities

`func (o *WorkloadSubType) GetEntities() []EntityIdentifier`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *WorkloadSubType) GetEntitiesOk() (*[]EntityIdentifier, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *WorkloadSubType) SetEntities(v []EntityIdentifier)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *WorkloadSubType) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### SetEntitiesNil

`func (o *WorkloadSubType) SetEntitiesNil(b bool)`

 SetEntitiesNil sets the value for Entities to be an explicit nil

### UnsetEntities
`func (o *WorkloadSubType) UnsetEntities()`

UnsetEntities ensures that no value is present for Entities, not even an explicit nil
### GetSchema

`func (o *WorkloadSubType) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *WorkloadSubType) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *WorkloadSubType) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *WorkloadSubType) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *WorkloadSubType) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *WorkloadSubType) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


