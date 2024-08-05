# AzureStorageConfigForIndexing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **NullableString** | The container to be used for the tenant. | 
**StorageAccountName** | **NullableString** | The storage account to be used for the tenant. | 

## Methods

### NewAzureStorageConfigForIndexing

`func NewAzureStorageConfigForIndexing(containerName NullableString, storageAccountName NullableString, ) *AzureStorageConfigForIndexing`

NewAzureStorageConfigForIndexing instantiates a new AzureStorageConfigForIndexing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageConfigForIndexingWithDefaults

`func NewAzureStorageConfigForIndexingWithDefaults() *AzureStorageConfigForIndexing`

NewAzureStorageConfigForIndexingWithDefaults instantiates a new AzureStorageConfigForIndexing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *AzureStorageConfigForIndexing) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *AzureStorageConfigForIndexing) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *AzureStorageConfigForIndexing) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### SetContainerNameNil

`func (o *AzureStorageConfigForIndexing) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *AzureStorageConfigForIndexing) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetStorageAccountName

`func (o *AzureStorageConfigForIndexing) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *AzureStorageConfigForIndexing) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *AzureStorageConfigForIndexing) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.


### SetStorageAccountNameNil

`func (o *AzureStorageConfigForIndexing) SetStorageAccountNameNil(b bool)`

 SetStorageAccountNameNil sets the value for StorageAccountName to be an explicit nil

### UnsetStorageAccountName
`func (o *AzureStorageConfigForIndexing) UnsetStorageAccountName()`

UnsetStorageAccountName ensures that no value is present for StorageAccountName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


