# CommonTieringAzureExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **NullableString** | Specifies the client id of the managed identity assigned to the cluster This is used only for clusters running as Azure VMs where authentication is done using AD. | [optional] 
**ContainerName** | **NullableString** | Specifies the container name of the external target. | 
**Region** | Pointer to **NullableString** | Specifies region of the External Target. This is only populated for FortKnox vaults. | [optional] 
**StorageAccessKey** | Pointer to **NullableString** | Specifies the storage access key of the external target. | [optional] 
**StorageAccountName** | **NullableString** | Specifies the storage account name of the external target. | 
**StorageClass** | **NullableString** | Specifies the Azure External Target class. | 

## Methods

### NewCommonTieringAzureExternalTargetParams

`func NewCommonTieringAzureExternalTargetParams(containerName NullableString, storageAccountName NullableString, storageClass NullableString, ) *CommonTieringAzureExternalTargetParams`

NewCommonTieringAzureExternalTargetParams instantiates a new CommonTieringAzureExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTieringAzureExternalTargetParamsWithDefaults

`func NewCommonTieringAzureExternalTargetParamsWithDefaults() *CommonTieringAzureExternalTargetParams`

NewCommonTieringAzureExternalTargetParamsWithDefaults instantiates a new CommonTieringAzureExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CommonTieringAzureExternalTargetParams) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CommonTieringAzureExternalTargetParams) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CommonTieringAzureExternalTargetParams) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CommonTieringAzureExternalTargetParams) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *CommonTieringAzureExternalTargetParams) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *CommonTieringAzureExternalTargetParams) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetContainerName

`func (o *CommonTieringAzureExternalTargetParams) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *CommonTieringAzureExternalTargetParams) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *CommonTieringAzureExternalTargetParams) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### SetContainerNameNil

`func (o *CommonTieringAzureExternalTargetParams) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *CommonTieringAzureExternalTargetParams) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetRegion

`func (o *CommonTieringAzureExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CommonTieringAzureExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CommonTieringAzureExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CommonTieringAzureExternalTargetParams) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CommonTieringAzureExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CommonTieringAzureExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStorageAccessKey

`func (o *CommonTieringAzureExternalTargetParams) GetStorageAccessKey() string`

GetStorageAccessKey returns the StorageAccessKey field if non-nil, zero value otherwise.

### GetStorageAccessKeyOk

`func (o *CommonTieringAzureExternalTargetParams) GetStorageAccessKeyOk() (*string, bool)`

GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccessKey

`func (o *CommonTieringAzureExternalTargetParams) SetStorageAccessKey(v string)`

SetStorageAccessKey sets StorageAccessKey field to given value.

### HasStorageAccessKey

`func (o *CommonTieringAzureExternalTargetParams) HasStorageAccessKey() bool`

HasStorageAccessKey returns a boolean if a field has been set.

### SetStorageAccessKeyNil

`func (o *CommonTieringAzureExternalTargetParams) SetStorageAccessKeyNil(b bool)`

 SetStorageAccessKeyNil sets the value for StorageAccessKey to be an explicit nil

### UnsetStorageAccessKey
`func (o *CommonTieringAzureExternalTargetParams) UnsetStorageAccessKey()`

UnsetStorageAccessKey ensures that no value is present for StorageAccessKey, not even an explicit nil
### GetStorageAccountName

`func (o *CommonTieringAzureExternalTargetParams) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *CommonTieringAzureExternalTargetParams) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *CommonTieringAzureExternalTargetParams) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.


### SetStorageAccountNameNil

`func (o *CommonTieringAzureExternalTargetParams) SetStorageAccountNameNil(b bool)`

 SetStorageAccountNameNil sets the value for StorageAccountName to be an explicit nil

### UnsetStorageAccountName
`func (o *CommonTieringAzureExternalTargetParams) UnsetStorageAccountName()`

UnsetStorageAccountName ensures that no value is present for StorageAccountName, not even an explicit nil
### GetStorageClass

`func (o *CommonTieringAzureExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *CommonTieringAzureExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *CommonTieringAzureExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *CommonTieringAzureExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *CommonTieringAzureExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


