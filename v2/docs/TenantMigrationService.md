# TenantMigrationService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]TenantMigrationServiceAction**](TenantMigrationServiceAction.md) | List of Actions. | [optional] 
**Service** | Pointer to **NullableString** | Specifies the cluster service on which this action needs to be performed. | [optional] 

## Methods

### NewTenantMigrationService

`func NewTenantMigrationService() *TenantMigrationService`

NewTenantMigrationService instantiates a new TenantMigrationService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantMigrationServiceWithDefaults

`func NewTenantMigrationServiceWithDefaults() *TenantMigrationService`

NewTenantMigrationServiceWithDefaults instantiates a new TenantMigrationService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *TenantMigrationService) GetActions() []TenantMigrationServiceAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TenantMigrationService) GetActionsOk() (*[]TenantMigrationServiceAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TenantMigrationService) SetActions(v []TenantMigrationServiceAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TenantMigrationService) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetService

`func (o *TenantMigrationService) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *TenantMigrationService) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *TenantMigrationService) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *TenantMigrationService) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *TenantMigrationService) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *TenantMigrationService) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


