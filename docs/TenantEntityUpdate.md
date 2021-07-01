# TenantEntityUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityIds** | Pointer to **[]int64** | Specifies the EntityIds for respective tenant. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantEntityUpdate

`func NewTenantEntityUpdate() *TenantEntityUpdate`

NewTenantEntityUpdate instantiates a new TenantEntityUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantEntityUpdateWithDefaults

`func NewTenantEntityUpdateWithDefaults() *TenantEntityUpdate`

NewTenantEntityUpdateWithDefaults instantiates a new TenantEntityUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityIds

`func (o *TenantEntityUpdate) GetEntityIds() []int64`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *TenantEntityUpdate) GetEntityIdsOk() (*[]int64, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *TenantEntityUpdate) SetEntityIds(v []int64)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *TenantEntityUpdate) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### SetEntityIdsNil

`func (o *TenantEntityUpdate) SetEntityIdsNil(b bool)`

 SetEntityIdsNil sets the value for EntityIds to be an explicit nil

### UnsetEntityIds
`func (o *TenantEntityUpdate) UnsetEntityIds()`

UnsetEntityIds ensures that no value is present for EntityIds, not even an explicit nil
### GetTenantId

`func (o *TenantEntityUpdate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantEntityUpdate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantEntityUpdate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantEntityUpdate) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantEntityUpdate) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantEntityUpdate) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


