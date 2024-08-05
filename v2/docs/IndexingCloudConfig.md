# IndexingCloudConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureEsConfig** | Pointer to [**AzureESConfigForIndexing**](AzureESConfigForIndexing.md) |  | [optional] 
**AzureStorageConfig** | Pointer to [**AzureStorageConfigForIndexing**](AzureStorageConfigForIndexing.md) |  | [optional] 
**EsConfig** | Pointer to [**ESConfigForIndexing**](ESConfigForIndexing.md) |  | [optional] 
**IsMigratedTenant** | Pointer to **NullableBool** | Whether this tenant is being migrated to this cluster or freshly onboarded. | [optional] 
**Region** | **NullableString** | Name of the cloud region. | 
**S3Config** | Pointer to [**S3ConfigForIndexing**](S3ConfigForIndexing.md) |  | [optional] 
**TenantId** | **NullableString** | Tenant ID to which this config belongs. | 

## Methods

### NewIndexingCloudConfig

`func NewIndexingCloudConfig(region NullableString, tenantId NullableString, ) *IndexingCloudConfig`

NewIndexingCloudConfig instantiates a new IndexingCloudConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexingCloudConfigWithDefaults

`func NewIndexingCloudConfigWithDefaults() *IndexingCloudConfig`

NewIndexingCloudConfigWithDefaults instantiates a new IndexingCloudConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureEsConfig

`func (o *IndexingCloudConfig) GetAzureEsConfig() AzureESConfigForIndexing`

GetAzureEsConfig returns the AzureEsConfig field if non-nil, zero value otherwise.

### GetAzureEsConfigOk

`func (o *IndexingCloudConfig) GetAzureEsConfigOk() (*AzureESConfigForIndexing, bool)`

GetAzureEsConfigOk returns a tuple with the AzureEsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEsConfig

`func (o *IndexingCloudConfig) SetAzureEsConfig(v AzureESConfigForIndexing)`

SetAzureEsConfig sets AzureEsConfig field to given value.

### HasAzureEsConfig

`func (o *IndexingCloudConfig) HasAzureEsConfig() bool`

HasAzureEsConfig returns a boolean if a field has been set.

### GetAzureStorageConfig

`func (o *IndexingCloudConfig) GetAzureStorageConfig() AzureStorageConfigForIndexing`

GetAzureStorageConfig returns the AzureStorageConfig field if non-nil, zero value otherwise.

### GetAzureStorageConfigOk

`func (o *IndexingCloudConfig) GetAzureStorageConfigOk() (*AzureStorageConfigForIndexing, bool)`

GetAzureStorageConfigOk returns a tuple with the AzureStorageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorageConfig

`func (o *IndexingCloudConfig) SetAzureStorageConfig(v AzureStorageConfigForIndexing)`

SetAzureStorageConfig sets AzureStorageConfig field to given value.

### HasAzureStorageConfig

`func (o *IndexingCloudConfig) HasAzureStorageConfig() bool`

HasAzureStorageConfig returns a boolean if a field has been set.

### GetEsConfig

`func (o *IndexingCloudConfig) GetEsConfig() ESConfigForIndexing`

GetEsConfig returns the EsConfig field if non-nil, zero value otherwise.

### GetEsConfigOk

`func (o *IndexingCloudConfig) GetEsConfigOk() (*ESConfigForIndexing, bool)`

GetEsConfigOk returns a tuple with the EsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsConfig

`func (o *IndexingCloudConfig) SetEsConfig(v ESConfigForIndexing)`

SetEsConfig sets EsConfig field to given value.

### HasEsConfig

`func (o *IndexingCloudConfig) HasEsConfig() bool`

HasEsConfig returns a boolean if a field has been set.

### GetIsMigratedTenant

`func (o *IndexingCloudConfig) GetIsMigratedTenant() bool`

GetIsMigratedTenant returns the IsMigratedTenant field if non-nil, zero value otherwise.

### GetIsMigratedTenantOk

`func (o *IndexingCloudConfig) GetIsMigratedTenantOk() (*bool, bool)`

GetIsMigratedTenantOk returns a tuple with the IsMigratedTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMigratedTenant

`func (o *IndexingCloudConfig) SetIsMigratedTenant(v bool)`

SetIsMigratedTenant sets IsMigratedTenant field to given value.

### HasIsMigratedTenant

`func (o *IndexingCloudConfig) HasIsMigratedTenant() bool`

HasIsMigratedTenant returns a boolean if a field has been set.

### SetIsMigratedTenantNil

`func (o *IndexingCloudConfig) SetIsMigratedTenantNil(b bool)`

 SetIsMigratedTenantNil sets the value for IsMigratedTenant to be an explicit nil

### UnsetIsMigratedTenant
`func (o *IndexingCloudConfig) UnsetIsMigratedTenant()`

UnsetIsMigratedTenant ensures that no value is present for IsMigratedTenant, not even an explicit nil
### GetRegion

`func (o *IndexingCloudConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IndexingCloudConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IndexingCloudConfig) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *IndexingCloudConfig) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *IndexingCloudConfig) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetS3Config

`func (o *IndexingCloudConfig) GetS3Config() S3ConfigForIndexing`

GetS3Config returns the S3Config field if non-nil, zero value otherwise.

### GetS3ConfigOk

`func (o *IndexingCloudConfig) GetS3ConfigOk() (*S3ConfigForIndexing, bool)`

GetS3ConfigOk returns a tuple with the S3Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Config

`func (o *IndexingCloudConfig) SetS3Config(v S3ConfigForIndexing)`

SetS3Config sets S3Config field to given value.

### HasS3Config

`func (o *IndexingCloudConfig) HasS3Config() bool`

HasS3Config returns a boolean if a field has been set.

### GetTenantId

`func (o *IndexingCloudConfig) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IndexingCloudConfig) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IndexingCloudConfig) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### SetTenantIdNil

`func (o *IndexingCloudConfig) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *IndexingCloudConfig) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


