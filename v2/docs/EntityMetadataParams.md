# EntityMetadataParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsParams** | Pointer to [**AwsEntityMetadata**](AwsEntityMetadata.md) |  | [optional] 
**AzureParams** | Pointer to [**AzureEntityMetadata**](AzureEntityMetadata.md) |  | [optional] 
**EntityId** | **int64** | Specifies the entity id of the entity whose metadata is being updated. | 
**MaintenanceModeConfig** | Pointer to [**MaintenanceModeConfig**](MaintenanceModeConfig.md) |  | [optional] 

## Methods

### NewEntityMetadataParams

`func NewEntityMetadataParams(entityId int64, ) *EntityMetadataParams`

NewEntityMetadataParams instantiates a new EntityMetadataParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityMetadataParamsWithDefaults

`func NewEntityMetadataParamsWithDefaults() *EntityMetadataParams`

NewEntityMetadataParamsWithDefaults instantiates a new EntityMetadataParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsParams

`func (o *EntityMetadataParams) GetAwsParams() AwsEntityMetadata`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *EntityMetadataParams) GetAwsParamsOk() (*AwsEntityMetadata, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *EntityMetadataParams) SetAwsParams(v AwsEntityMetadata)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *EntityMetadataParams) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *EntityMetadataParams) GetAzureParams() AzureEntityMetadata`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *EntityMetadataParams) GetAzureParamsOk() (*AzureEntityMetadata, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *EntityMetadataParams) SetAzureParams(v AzureEntityMetadata)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *EntityMetadataParams) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetEntityId

`func (o *EntityMetadataParams) GetEntityId() int64`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityMetadataParams) GetEntityIdOk() (*int64, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityMetadataParams) SetEntityId(v int64)`

SetEntityId sets EntityId field to given value.


### GetMaintenanceModeConfig

`func (o *EntityMetadataParams) GetMaintenanceModeConfig() MaintenanceModeConfig`

GetMaintenanceModeConfig returns the MaintenanceModeConfig field if non-nil, zero value otherwise.

### GetMaintenanceModeConfigOk

`func (o *EntityMetadataParams) GetMaintenanceModeConfigOk() (*MaintenanceModeConfig, bool)`

GetMaintenanceModeConfigOk returns a tuple with the MaintenanceModeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceModeConfig

`func (o *EntityMetadataParams) SetMaintenanceModeConfig(v MaintenanceModeConfig)`

SetMaintenanceModeConfig sets MaintenanceModeConfig field to given value.

### HasMaintenanceModeConfig

`func (o *EntityMetadataParams) HasMaintenanceModeConfig() bool`

HasMaintenanceModeConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


