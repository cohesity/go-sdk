# CommonArchivalAzureExternalTargetParams

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

## Methods

### NewCommonArchivalAzureExternalTargetParams

`func NewCommonArchivalAzureExternalTargetParams(containerName NullableString, storageAccountName NullableString, storageClass NullableString, ) *CommonArchivalAzureExternalTargetParams`

NewCommonArchivalAzureExternalTargetParams instantiates a new CommonArchivalAzureExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonArchivalAzureExternalTargetParamsWithDefaults

`func NewCommonArchivalAzureExternalTargetParamsWithDefaults() *CommonArchivalAzureExternalTargetParams`

NewCommonArchivalAzureExternalTargetParamsWithDefaults instantiates a new CommonArchivalAzureExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *CommonArchivalAzureExternalTargetParams) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CommonArchivalAzureExternalTargetParams) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CommonArchivalAzureExternalTargetParams) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CommonArchivalAzureExternalTargetParams) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *CommonArchivalAzureExternalTargetParams) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *CommonArchivalAzureExternalTargetParams) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetContainerName

`func (o *CommonArchivalAzureExternalTargetParams) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *CommonArchivalAzureExternalTargetParams) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *CommonArchivalAzureExternalTargetParams) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### SetContainerNameNil

`func (o *CommonArchivalAzureExternalTargetParams) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *CommonArchivalAzureExternalTargetParams) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetRegion

`func (o *CommonArchivalAzureExternalTargetParams) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CommonArchivalAzureExternalTargetParams) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CommonArchivalAzureExternalTargetParams) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CommonArchivalAzureExternalTargetParams) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CommonArchivalAzureExternalTargetParams) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CommonArchivalAzureExternalTargetParams) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStorageAccessKey

`func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccessKey() string`

GetStorageAccessKey returns the StorageAccessKey field if non-nil, zero value otherwise.

### GetStorageAccessKeyOk

`func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccessKeyOk() (*string, bool)`

GetStorageAccessKeyOk returns a tuple with the StorageAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccessKey

`func (o *CommonArchivalAzureExternalTargetParams) SetStorageAccessKey(v string)`

SetStorageAccessKey sets StorageAccessKey field to given value.

### HasStorageAccessKey

`func (o *CommonArchivalAzureExternalTargetParams) HasStorageAccessKey() bool`

HasStorageAccessKey returns a boolean if a field has been set.

### SetStorageAccessKeyNil

`func (o *CommonArchivalAzureExternalTargetParams) SetStorageAccessKeyNil(b bool)`

 SetStorageAccessKeyNil sets the value for StorageAccessKey to be an explicit nil

### UnsetStorageAccessKey
`func (o *CommonArchivalAzureExternalTargetParams) UnsetStorageAccessKey()`

UnsetStorageAccessKey ensures that no value is present for StorageAccessKey, not even an explicit nil
### GetStorageAccountName

`func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *CommonArchivalAzureExternalTargetParams) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *CommonArchivalAzureExternalTargetParams) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.


### SetStorageAccountNameNil

`func (o *CommonArchivalAzureExternalTargetParams) SetStorageAccountNameNil(b bool)`

 SetStorageAccountNameNil sets the value for StorageAccountName to be an explicit nil

### UnsetStorageAccountName
`func (o *CommonArchivalAzureExternalTargetParams) UnsetStorageAccountName()`

UnsetStorageAccountName ensures that no value is present for StorageAccountName, not even an explicit nil
### GetIsForeverIncrementalArchivalEnabled

`func (o *CommonArchivalAzureExternalTargetParams) GetIsForeverIncrementalArchivalEnabled() bool`

GetIsForeverIncrementalArchivalEnabled returns the IsForeverIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsForeverIncrementalArchivalEnabledOk

`func (o *CommonArchivalAzureExternalTargetParams) GetIsForeverIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsForeverIncrementalArchivalEnabledOk returns a tuple with the IsForeverIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForeverIncrementalArchivalEnabled

`func (o *CommonArchivalAzureExternalTargetParams) SetIsForeverIncrementalArchivalEnabled(v bool)`

SetIsForeverIncrementalArchivalEnabled sets IsForeverIncrementalArchivalEnabled field to given value.

### HasIsForeverIncrementalArchivalEnabled

`func (o *CommonArchivalAzureExternalTargetParams) HasIsForeverIncrementalArchivalEnabled() bool`

HasIsForeverIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsForeverIncrementalArchivalEnabledNil

`func (o *CommonArchivalAzureExternalTargetParams) SetIsForeverIncrementalArchivalEnabledNil(b bool)`

 SetIsForeverIncrementalArchivalEnabledNil sets the value for IsForeverIncrementalArchivalEnabled to be an explicit nil

### UnsetIsForeverIncrementalArchivalEnabled
`func (o *CommonArchivalAzureExternalTargetParams) UnsetIsForeverIncrementalArchivalEnabled()`

UnsetIsForeverIncrementalArchivalEnabled ensures that no value is present for IsForeverIncrementalArchivalEnabled, not even an explicit nil
### GetIsIncrementalArchivalEnabled

`func (o *CommonArchivalAzureExternalTargetParams) GetIsIncrementalArchivalEnabled() bool`

GetIsIncrementalArchivalEnabled returns the IsIncrementalArchivalEnabled field if non-nil, zero value otherwise.

### GetIsIncrementalArchivalEnabledOk

`func (o *CommonArchivalAzureExternalTargetParams) GetIsIncrementalArchivalEnabledOk() (*bool, bool)`

GetIsIncrementalArchivalEnabledOk returns a tuple with the IsIncrementalArchivalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncrementalArchivalEnabled

`func (o *CommonArchivalAzureExternalTargetParams) SetIsIncrementalArchivalEnabled(v bool)`

SetIsIncrementalArchivalEnabled sets IsIncrementalArchivalEnabled field to given value.

### HasIsIncrementalArchivalEnabled

`func (o *CommonArchivalAzureExternalTargetParams) HasIsIncrementalArchivalEnabled() bool`

HasIsIncrementalArchivalEnabled returns a boolean if a field has been set.

### SetIsIncrementalArchivalEnabledNil

`func (o *CommonArchivalAzureExternalTargetParams) SetIsIncrementalArchivalEnabledNil(b bool)`

 SetIsIncrementalArchivalEnabledNil sets the value for IsIncrementalArchivalEnabled to be an explicit nil

### UnsetIsIncrementalArchivalEnabled
`func (o *CommonArchivalAzureExternalTargetParams) UnsetIsIncrementalArchivalEnabled()`

UnsetIsIncrementalArchivalEnabled ensures that no value is present for IsIncrementalArchivalEnabled, not even an explicit nil
### GetIsWormEnabled

`func (o *CommonArchivalAzureExternalTargetParams) GetIsWormEnabled() bool`

GetIsWormEnabled returns the IsWormEnabled field if non-nil, zero value otherwise.

### GetIsWormEnabledOk

`func (o *CommonArchivalAzureExternalTargetParams) GetIsWormEnabledOk() (*bool, bool)`

GetIsWormEnabledOk returns a tuple with the IsWormEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWormEnabled

`func (o *CommonArchivalAzureExternalTargetParams) SetIsWormEnabled(v bool)`

SetIsWormEnabled sets IsWormEnabled field to given value.

### HasIsWormEnabled

`func (o *CommonArchivalAzureExternalTargetParams) HasIsWormEnabled() bool`

HasIsWormEnabled returns a boolean if a field has been set.

### SetIsWormEnabledNil

`func (o *CommonArchivalAzureExternalTargetParams) SetIsWormEnabledNil(b bool)`

 SetIsWormEnabledNil sets the value for IsWormEnabled to be an explicit nil

### UnsetIsWormEnabled
`func (o *CommonArchivalAzureExternalTargetParams) UnsetIsWormEnabled()`

UnsetIsWormEnabled ensures that no value is present for IsWormEnabled, not even an explicit nil
### GetSourceSideDeduplication

`func (o *CommonArchivalAzureExternalTargetParams) GetSourceSideDeduplication() bool`

GetSourceSideDeduplication returns the SourceSideDeduplication field if non-nil, zero value otherwise.

### GetSourceSideDeduplicationOk

`func (o *CommonArchivalAzureExternalTargetParams) GetSourceSideDeduplicationOk() (*bool, bool)`

GetSourceSideDeduplicationOk returns a tuple with the SourceSideDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDeduplication

`func (o *CommonArchivalAzureExternalTargetParams) SetSourceSideDeduplication(v bool)`

SetSourceSideDeduplication sets SourceSideDeduplication field to given value.

### HasSourceSideDeduplication

`func (o *CommonArchivalAzureExternalTargetParams) HasSourceSideDeduplication() bool`

HasSourceSideDeduplication returns a boolean if a field has been set.

### SetSourceSideDeduplicationNil

`func (o *CommonArchivalAzureExternalTargetParams) SetSourceSideDeduplicationNil(b bool)`

 SetSourceSideDeduplicationNil sets the value for SourceSideDeduplication to be an explicit nil

### UnsetSourceSideDeduplication
`func (o *CommonArchivalAzureExternalTargetParams) UnsetSourceSideDeduplication()`

UnsetSourceSideDeduplication ensures that no value is present for SourceSideDeduplication, not even an explicit nil
### GetStorageClass

`func (o *CommonArchivalAzureExternalTargetParams) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *CommonArchivalAzureExternalTargetParams) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *CommonArchivalAzureExternalTargetParams) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.


### SetStorageClassNil

`func (o *CommonArchivalAzureExternalTargetParams) SetStorageClassNil(b bool)`

 SetStorageClassNil sets the value for StorageClass to be an explicit nil

### UnsetStorageClass
`func (o *CommonArchivalAzureExternalTargetParams) UnsetStorageClass()`

UnsetStorageClass ensures that no value is present for StorageClass, not even an explicit nil
### GetWormSpecificTargetParams

`func (o *CommonArchivalAzureExternalTargetParams) GetWormSpecificTargetParams() WormSpecificTargetParams`

GetWormSpecificTargetParams returns the WormSpecificTargetParams field if non-nil, zero value otherwise.

### GetWormSpecificTargetParamsOk

`func (o *CommonArchivalAzureExternalTargetParams) GetWormSpecificTargetParamsOk() (*WormSpecificTargetParams, bool)`

GetWormSpecificTargetParamsOk returns a tuple with the WormSpecificTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormSpecificTargetParams

`func (o *CommonArchivalAzureExternalTargetParams) SetWormSpecificTargetParams(v WormSpecificTargetParams)`

SetWormSpecificTargetParams sets WormSpecificTargetParams field to given value.

### HasWormSpecificTargetParams

`func (o *CommonArchivalAzureExternalTargetParams) HasWormSpecificTargetParams() bool`

HasWormSpecificTargetParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


