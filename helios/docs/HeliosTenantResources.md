# HeliosTenantResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **NullableString** | Specifies the new tenant id associated with the api key. | [optional] 
**Resources** | Pointer to [**[]HeliosTenantClusterResources**](HeliosTenantClusterResources.md) | THe list of assigned resources grouped by cluster Id. | [optional] 

## Methods

### NewHeliosTenantResources

`func NewHeliosTenantResources() *HeliosTenantResources`

NewHeliosTenantResources instantiates a new HeliosTenantResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosTenantResourcesWithDefaults

`func NewHeliosTenantResourcesWithDefaults() *HeliosTenantResources`

NewHeliosTenantResourcesWithDefaults instantiates a new HeliosTenantResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *HeliosTenantResources) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *HeliosTenantResources) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *HeliosTenantResources) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *HeliosTenantResources) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *HeliosTenantResources) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *HeliosTenantResources) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetResources

`func (o *HeliosTenantResources) GetResources() []HeliosTenantClusterResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *HeliosTenantResources) GetResourcesOk() (*[]HeliosTenantClusterResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *HeliosTenantResources) SetResources(v []HeliosTenantClusterResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *HeliosTenantResources) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *HeliosTenantResources) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *HeliosTenantResources) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


