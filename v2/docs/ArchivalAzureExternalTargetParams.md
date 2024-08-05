# ArchivalAzureExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **NullableString** | Specifies the client id of the managed identity assigned to the cluster This is used only for clusters running as Azure VMs where authentication is done using AD. | [optional] 
**ContainerName** | **NullableString** | Specifies the container name of the external target. | 
**Region** | Pointer to **NullableString** | Specifies region of the External Target. This is only populated for FortKnox vaults. | [optional] 
**StorageAccessKey** | Pointer to **NullableString** | Specifies the storage access key of the external target. | [optional] 
**StorageAccountName** | **NullableString** | Specifies the storage account name of the external target. | 
**IsForeverIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Forever Incremental Archival setting is enabled or not. | [optional] 
**IsIncrementalArchivalEnabled** | Pointer to **NullableBool** | Specifies if Incremental Archival setting is enabled or not. | [optional] 
**IsWormEnabled** | Pointer to **NullableBool** | Specifies whether write once read many (WORM) protection is enabled for the Azure container or not. | [optional] 
**SourceSideDeduplication** | Pointer to **NullableBool** | Specifies the Source Side Deduplication setting for the Azure external target | [optional] 
**StorageClass** | **NullableString** | Specifies the Azure External Target storage class. | 
**WormSpecificTargetParams** | Pointer to [**WormSpecificTargetParams**](WormSpecificTargetParams.md) |  | [optional] 
**ArchiveBlobParams** | Pointer to [**AzureArchiveBlobParams**](AzureArchiveBlobParams.md) |  | [optional] 
**CoolBlobParams** | Pointer to [**AzureCoolBlobParams**](AzureCoolBlobParams.md) |  | [optional] 
**HotBlobParams** | Pointer to [**AzureHotBlobParams**](AzureHotBlobParams.md) |  | [optional] 

## Methods

### NewArchivalAzureExternalTargetParams

`func NewArchivalAzureExternalTargetParams(containerName NullableString, storageAccountName NullableString, storageClass NullableString, ) *ArchivalAzureExternalTargetParams`

NewArchivalAzureExternalTargetParams instantiates a new ArchivalAzureExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalAzureExternalTargetParamsWithDefaults

`func NewArchivalAzureExternalTargetParamsWithDefaults() *ArchivalAzureExternalTargetParams`

NewArchivalAzureExternalTargetParamsWithDefaults instantiates a new ArchivalAzureExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ArchivalAzureExternalTargetParams) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ArchivalAzureExternalTargetParams) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ArchivalAzureExternalTargetParams) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ArchivalAzureExternalTargetParams) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *ArchivalAzureExternalTargetParams) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *ArchivalAzureExternalTargetParams) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetContainerName

`func (o *ArchivalAzureExternalTargetParams) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ArchivalAzureExternalTargetParams) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ArchivalAzureExternalTargetParams) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### SetContainerNameNil

`func (o *ArchivalAzureExternalTargetParams) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *ArchivalAzureExternalTargetParams) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetRegion

`func (o *ArchivalAzureExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ArchivalAzureExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ArchivalAzureExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ArchivalAzureExternalTargetParams) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ArchivalAzureExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ArchivalAzureExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStorageAccessKey

`func (o *ArchivalAzureExternalTargetParams) GetStorageAccessKey() string`

GetStorageAccessKey returns the StorageAccessKey field if non-nil, zero value otherwise.

### GetStorageAccessKeyOk

`func (o *ArchivalAzureExternalTargetParams) GetStorageAccessKeyOk() (*string, bool)`

GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccessKey

`func (o *ArchivalAzureExternalTargetParams) SetStorageAccessKey(v string)`

SetStorageAccessKey sets StorageAccessKey field to given value.

### HasStorageAccessKey

`func (o *ArchivalAzureExternalTargetParams) HasStorageAccessKey() bool`

HasStorageAccessKey returns a boolean if a field has been set.

### SetStorageAccessKeyNil

`func (o *ArchivalAzureExternalTargetParams) SetStorageAccessKeyNil(b bool)`

 SetStorageAccessKeyNil sets the value for StorageAccessKey to be an explicit nil

### UnsetStorageAccessKey
`func (o *ArchivalAzureExternalTargetParams) UnsetStorageAccessKey()`

UnsetStorageAccessKey ensures that no value is present for StorageAccessKey, not even an explicit nil
### GetStorageAccountName

`func (o *ArchivalAzureExternalTargetParams) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *ArchivalAzureExternalTargetParams) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *ArchivalAzureExternalTargetParams) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.


### SetStorageAccountNameNil

`func (o *ArchivalAzureExternalTargetParams) SetStorageAccountNameNil(b bool)`

 SetStorageAccountNameNil sets the value for StorageAccountName to be an explicit nil

### UnsetStorageAccountName
`func (o *ArchivalAzureExternalTargetParams) UnsetStorageAccountName()`

UnsetStorageAccountName ensures that no value is present for StorageAccountName, not even an explicit nil
### GetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalAzureExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool`

GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsForeverIncrementalArchivalEnabledOk

`func (o *ArchivalAzureExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncrementalArchivalEnabled

`func (o *ArchivalAzureExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool)`

SetIsForeverIncrementalArchivalEnabled sets IsForeverIncrementalArchivalEnabled field to given value.

### HasIsForeverIncrementalArchivalEnabled

`func (o *ArchivalAzureExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool`

HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsForeverIncrementalArchivalEnabledNil

`func (o *ArchivalAzureExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil(b bool)`

 SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil

### UnsetIsForeverIncrementalArchivalEnabled
`func (o *ArchivalAzureExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled()`

UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil
### GetIsIncrementalArchivalEnabled

`func (o *ArchivalAzureExternalTargetParams) GetIsIncrementalArchivalEnabled() bool`

GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsIncrementalArchivalEnabledOk

`func (o *ArchivalAzureExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncrementalArchivalEnabled

`func (o *ArchivalAzureExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool)`

SetIsIncrementalArchivalEnabled sets IsIncrementalArchivalEnabled field to given value.

### HasIsIncrementalArchivalEnabled

`func (o *ArchivalAzureExternalTargetParams) HasIsIncrementalArchivalEnabled() bool`

HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsIncrementalArchivalEnabledNil

`func (o *ArchivalAzureExternalTargetParams) SetIsIncrementalArchivalEnabledNil(b bool)`

 SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil

### UnsetIsIncrementalArchivalEnabled
`func (o *ArchivalAzureExternalTargetParams) UnsetIsIncrementalArchivalEnabled()`

UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
### GetIsWormEnabled

`func (o *ArchivalAzureExternalTargetParams) GetIsWormEnabled() bool`

GetIsWormEnabled returns the IsWormEnabled field if non-nil, zero value otherwise.

### GetIsWormEnabledOk

`func (o *ArchivalAzureExternalTargetParams) GetIsWormEnabledOk() (*bool, bool)`

GetIsWormEnabledOk returns a tuple with the IsWormEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWormEnabled

`func (o *ArchivalAzureExternalTargetParams) SetIsWormEnabled(v bool)`

SetIsWormEnabled sets IsWormEnabled field to given value.

### HasIsWormEnabled

`func (o *ArchivalAzureExternalTargetParams) HasIsWormEnabled() bool`

HasIsWormEnabled returns a boolean if a field has been set.

### SetIsWormEnabledNil

`func (o *ArchivalAzureExternalTargetParams) SetIsWormEnabledNil(b bool)`

 SetIsWormEnabledNil sets the value for IsWormEnabled to be an explicit nil

### UnsetIsWormEnabled
`func (o *ArchivalAzureExternalTargetParams) UnsetIsWormEnabled()`

UnsetIsWormEnabled ensures that no value is present for IsWormEnabled, not even an explicit nil
### GetSourceSideDeduplication

`func (o *ArchivalAzureExternalTargetParams) GetSourceSideDeduplication() bool`

GetSourceSideDeduplication returns the SourceSideDeduplication field if non-nil, zero value otherwise.

### GetSourceSideDeduplicationOk

`func (o *ArchivalAzureExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool)`

GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDeduplication

`func (o *ArchivalAzureExternalTargetParams) SetSourceSideDeduplication(v bool)`

SetSourceSideDeduplication sets SourceSideDeduplication field to given value.

### HasSourceSideDeduplication

`func (o *ArchivalAzureExternalTargetParams) HasSourceSideDeduplication() bool`

HasSourceSideDeduplication returns a boolean if a field has been set.

### SetSourceSideDeduplicationNil

`func (o *ArchivalAzureExternalTargetParams) SetSourceSideDeduplicationNil(b bool)`

 SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil

### UnsetSourceSideDeduplication
`func (o *ArchivalAzureExternalTargetParams) UnsetSourceSideDeduplication()`

UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil
### GetStorageClass

`func (o *ArchivalAzureExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *ArchivalAzureExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *ArchivalAzureExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *ArchivalAzureExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *ArchivalAzureExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetWormSpecificTargetParams

`func (o *ArchivalAzureExternalTargetParams) GetWormSpecificTargetParams() WormSpecificTargetParams`

GetWormSpecificTargetParams returns the WormSpecificTargetParams field if non-nil, zero value otherwise.

### GetWormSpecificTargetParamsOk

`func (o *ArchivalAzureExternalTargetParams) GetWormSpecificTargetParamsOk() (*WormSpecificTargetParams, bool)`

GetWormSpecificTargetParamsOk returns a tuple with the WormSpecificTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormSpecificTargetParams

`func (o *ArchivalAzureExternalTargetParams) SetWormSpecificTargetParams(v WormSpecificTargetParams)`

SetWormSpecificTargetParams sets WormSpecificTargetParams field to given value.

### HasWormSpecificTargetParams

`func (o *ArchivalAzureExternalTargetParams) HasWormSpecificTargetParams() bool`

HasWormSpecificTargetParams returns a boolean if a field has been set.

### GetArchiveBlobParams

`func (o *ArchivalAzureExternalTargetParams) GetArchiveBlobParams() AzureArchiveBlobParams`

GetArchiveBlobParams returns the ArchiveBlobParams field if non-nil, zero value otherwise.

### GetArchiveBlobParamsOk

`func (o *ArchivalAzureExternalTargetParams) GetArchiveBlobParamsOk() (*AzureArchiveBlobParams, bool)`

GetArchiveBlobParamsOk returns a tuple with the ArchiveBlobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBlobParams

`func (o *ArchivalAzureExternalTargetParams) SetArchiveBlobParams(v AzureArchiveBlobParams)`

SetArchiveBlobParams sets ArchiveBlobParams field to given value.

### HasArchiveBlobParams

`func (o *ArchivalAzureExternalTargetParams) HasArchiveBlobParams() bool`

HasArchiveBlobParams returns a boolean if a field has been set.

### GetCoolBlobParams

`func (o *ArchivalAzureExternalTargetParams) GetCoolBlobParams() AzureCoolBlobParams`

GetCoolBlobParams returns the CoolBlobParams field if non-nil, zero value otherwise.

### GetCoolBlobParamsOk

`func (o *ArchivalAzureExternalTargetParams) GetCoolBlobParamsOk() (*AzureCoolBlobParams, bool)`

GetCoolBlobParamsOk returns a tuple with the CoolBlobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolBlobParams

`func (o *ArchivalAzureExternalTargetParams) SetCoolBlobParams(v AzureCoolBlobParams)`

SetCoolBlobParams sets CoolBlobParams field to given value.

### HasCoolBlobParams

`func (o *ArchivalAzureExternalTargetParams) HasCoolBlobParams() bool`

HasCoolBlobParams returns a boolean if a field has been set.

### GetHotBlobParams

`func (o *ArchivalAzureExternalTargetParams) GetHotBlobParams() AzureHotBlobParams`

GetHotBlobParams returns the HotBlobParams field if non-nil, zero value otherwise.

### GetHotBlobParamsOk

`func (o *ArchivalAzureExternalTargetParams) GetHotBlobParamsOk() (*AzureHotBlobParams, bool)`

GetHotBlobParamsOk returns a tuple with the HotBlobParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotBlobParams

`func (o *ArchivalAzureExternalTargetParams) SetHotBlobParams(v AzureHotBlobParams)`

SetHotBlobParams sets HotBlobParams field to given value.

### HasHotBlobParams

`func (o *ArchivalAzureExternalTargetParams) HasHotBlobParams() bool`

HasHotBlobParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


