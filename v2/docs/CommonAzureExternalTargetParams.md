# CommonAzureExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **NullableString** | Specifies the client id of the managed identity assigned to the cluster This is used only for clusters running as Azure VMs where authentication is done using AD. | [optional] 
**ContainerName** | **NullableString** | Specifies the container name of the external target. | 
**Region** | Pointer to **NullableString** | Specifies region of the External Target. This is only populated for FortKnox vaults. | [optional] 
**StorageAccessKey** | Pointer to **NullableString** | Specifies the storage access key of the external target. | [optional] 
**StorageAccountName** | **NullableString** | Specifies the storage account name of the external target. | 

## Methods

### NewCommonAzureExternalTargetParams

`func NewCommonAzureExternalTargetParams(containerName NullableString, storageAccountName NullableString, ) *CommonAzureExternalTargetParams`

NewCommonAzureExternalTargetParams instantiates a new CommonAzureExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonAzureExternalTargetParamsWithDefaults

`func NewCommonAzureExternalTargetParamsWithDefaults() *CommonAzureExternalTargetParams`

NewCommonAzureExternalTargetParamsWithDefaults instantiates a new CommonAzureExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CommonAzureExternalTargetParams) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CommonAzureExternalTargetParams) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CommonAzureExternalTargetParams) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CommonAzureExternalTargetParams) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *CommonAzureExternalTargetParams) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *CommonAzureExternalTargetParams) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetContainerName

`func (o *CommonAzureExternalTargetParams) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *CommonAzureExternalTargetParams) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *CommonAzureExternalTargetParams) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### SetContainerNameNil

`func (o *CommonAzureExternalTargetParams) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *CommonAzureExternalTargetParams) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetRegion

`func (o *CommonAzureExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CommonAzureExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CommonAzureExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CommonAzureExternalTargetParams) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CommonAzureExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CommonAzureExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStorageAccessKey

`func (o *CommonAzureExternalTargetParams) GetStorageAccessKey() string`

GetStorageAccessKey returns the StorageAccessKey field if non-nil, zero value otherwise.

### GetStorageAccessKeyOk

`func (o *CommonAzureExternalTargetParams) GetStorageAccessKeyOk() (*string, bool)`

GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccessKey

`func (o *CommonAzureExternalTargetParams) SetStorageAccessKey(v string)`

SetStorageAccessKey sets StorageAccessKey field to given value.

### HasStorageAccessKey

`func (o *CommonAzureExternalTargetParams) HasStorageAccessKey() bool`

HasStorageAccessKey returns a boolean if a field has been set.

### SetStorageAccessKeyNil

`func (o *CommonAzureExternalTargetParams) SetStorageAccessKeyNil(b bool)`

 SetStorageAccessKeyNil sets the value for StorageAccessKey to be an explicit nil

### UnsetStorageAccessKey
`func (o *CommonAzureExternalTargetParams) UnsetStorageAccessKey()`

UnsetStorageAccessKey ensures that no value is present for StorageAccessKey, not even an explicit nil
### GetStorageAccountName

`func (o *CommonAzureExternalTargetParams) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *CommonAzureExternalTargetParams) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *CommonAzureExternalTargetParams) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.


### SetStorageAccountNameNil

`func (o *CommonAzureExternalTargetParams) SetStorageAccountNameNil(b bool)`

 SetStorageAccountNameNil sets the value for StorageAccountName to be an explicit nil

### UnsetStorageAccountName
`func (o *CommonAzureExternalTargetParams) UnsetStorageAccountName()`

UnsetStorageAccountName ensures that no value is present for StorageAccountName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


