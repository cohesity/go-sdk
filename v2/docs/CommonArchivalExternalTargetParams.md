# CommonArchivalExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encryption** | [**EncryptionSettings**](EncryptionSettings.md) |  | 
**StorageType** | **NullableString** | Specifies the Storage type of the External Target. Nas option in archival_target_storage_type will soon be deprecated. Please use NAS instead. | 
**TargetBandwidthThrottlings** | Pointer to [**TargetBandwidthThrottlings**](TargetBandwidthThrottlings.md) |  | [optional] 

## Methods

### NewCommonArchivalExternalTargetParams

`func NewCommonArchivalExternalTargetParams(encryption EncryptionSettings, storageType NullableString, ) *CommonArchivalExternalTargetParams`

NewCommonArchivalExternalTargetParams instantiates a new CommonArchivalExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonArchivalExternalTargetParamsWithDefaults

`func NewCommonArchivalExternalTargetParamsWithDefaults() *CommonArchivalExternalTargetParams`

NewCommonArchivalExternalTargetParamsWithDefaults instantiates a new CommonArchivalExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryption

`func (o *CommonArchivalExternalTargetParams) GetEncryption() EncryptionSettings`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *CommonArchivalExternalTargetParams) GetEncryptionOk() (*EncryptionSettings, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *CommonArchivalExternalTargetParams) SetEncryption(v EncryptionSettings)`

SetEncryption sets Encryption field to given value.


### GetStorageType

`func (o *CommonArchivalExternalTargetParams) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *CommonArchivalExternalTargetParams) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *CommonArchivalExternalTargetParams) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### SetStorageTypeNil

`func (o *CommonArchivalExternalTargetParams) SetStorageTypeNil(b bool)`

 SetStorageTypeNil sets the value for StorageType to be an explicit nil

### UnsetStorageType
`func (o *CommonArchivalExternalTargetParams) UnsetStorageType()`

UnsetStorageType ensures that no value is present for StorageType, not even an explicit nil
### GetTargetBandwidthThrottlings

`func (o *CommonArchivalExternalTargetParams) GetTargetBandwidthThrottlings() TargetBandwidthThrottlings`

GetTargetBandwidthThrottlings returns the TargetBandwidthThrottlings field if non-nil, zero value otherwise.

### GetTargetBandwidthThrottlingsOk

`func (o *CommonArchivalExternalTargetParams) GetTargetBandwidthThrottlingsOk() (*TargetBandwidthThrottlings, bool)`

GetTargetBandwidthThrottlingsOk returns a tuple with the TargetBandwidthThrottlings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBandwidthThrottlings

`func (o *CommonArchivalExternalTargetParams) SetTargetBandwidthThrottlings(v TargetBandwidthThrottlings)`

SetTargetBandwidthThrottlings sets TargetBandwidthThrottlings field to given value.

### HasTargetBandwidthThrottlings

`func (o *CommonArchivalExternalTargetParams) HasTargetBandwidthThrottlings() bool`

HasTargetBandwidthThrottlings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


