# StatsGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumer** | Pointer to [**Consumer**](Consumer.md) |  | [optional] 
**EntityId** | Pointer to **NullableString** | Specifies the entity id of the group. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the id of the group. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the group. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the id of the organization (tenant) with respect to this group. | [optional] 
**TenantName** | Pointer to **NullableString** | Specifies the name of the organization (tenant) with respect to this group. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Specifies the id of the view box (storage domain) with respect to this group. | [optional] 
**ViewBoxName** | Pointer to **NullableString** | Specifies the name of the view box (storage domain) with respect to this group. | [optional] 

## Methods

### NewStatsGroup

`func NewStatsGroup() *StatsGroup`

NewStatsGroup instantiates a new StatsGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsGroupWithDefaults

`func NewStatsGroupWithDefaults() *StatsGroup`

NewStatsGroupWithDefaults instantiates a new StatsGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumer

`func (o *StatsGroup) GetConsumer() Consumer`

GetConsumer returns the Consumer field if non-nil, zero value otherwise.

### GetConsumerOk

`func (o *StatsGroup) GetConsumerOk() (*Consumer, bool)`

GetConsumerOk returns a tuple with the Consumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumer

`func (o *StatsGroup) SetConsumer(v Consumer)`

SetConsumer sets Consumer field to given value.

### HasConsumer

`func (o *StatsGroup) HasConsumer() bool`

HasConsumer returns a boolean if a field has been set.

### GetEntityId

`func (o *StatsGroup) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *StatsGroup) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *StatsGroup) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *StatsGroup) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *StatsGroup) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *StatsGroup) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetId

`func (o *StatsGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StatsGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StatsGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StatsGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StatsGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StatsGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *StatsGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatsGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatsGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatsGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StatsGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StatsGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantId

`func (o *StatsGroup) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *StatsGroup) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *StatsGroup) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *StatsGroup) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *StatsGroup) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *StatsGroup) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantName

`func (o *StatsGroup) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *StatsGroup) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *StatsGroup) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *StatsGroup) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### SetTenantNameNil

`func (o *StatsGroup) SetTenantNameNil(b bool)`

 SetTenantNameNil sets the value for TenantName to be an explicit nil

### UnsetTenantName
`func (o *StatsGroup) UnsetTenantName()`

UnsetTenantName ensures that no value is present for TenantName, not even an explicit nil
### GetViewBoxId

`func (o *StatsGroup) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *StatsGroup) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *StatsGroup) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *StatsGroup) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *StatsGroup) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *StatsGroup) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil
### GetViewBoxName

`func (o *StatsGroup) GetViewBoxName() string`

GetViewBoxName returns the ViewBoxName field if non-nil, zero value otherwise.

### GetViewBoxNameOk

`func (o *StatsGroup) GetViewBoxNameOk() (*string, bool)`

GetViewBoxNameOk returns a tuple with the ViewBoxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxName

`func (o *StatsGroup) SetViewBoxName(v string)`

SetViewBoxName sets ViewBoxName field to given value.

### HasViewBoxName

`func (o *StatsGroup) HasViewBoxName() bool`

HasViewBoxName returns a boolean if a field has been set.

### SetViewBoxNameNil

`func (o *StatsGroup) SetViewBoxNameNil(b bool)`

 SetViewBoxNameNil sets the value for ViewBoxName to be an explicit nil

### UnsetViewBoxName
`func (o *StatsGroup) UnsetViewBoxName()`

UnsetViewBoxName ensures that no value is present for ViewBoxName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


