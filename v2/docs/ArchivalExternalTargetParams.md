# ArchivalExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encryption** | [**EncryptionSettings**](EncryptionSettings.md) |  | 
**StorageType** | **NullableString** | Specifies the Storage type of the External Target. Nas option in archival_target_storage_type will soon be deprecated. Please use NAS instead. | 
**TargetBandwidthThrottlings** | Pointer to [**TargetBandwidthThrottlings**](TargetBandwidthThrottlings.md) |  | [optional] 
**AwsParams** | Pointer to [**ArchivalAwsExternalTargetParams**](ArchivalAwsExternalTargetParams.md) |  | [optional] 
**AzureParams** | Pointer to [**ArchivalAzureExternalTargetParams**](ArchivalAzureExternalTargetParams.md) |  | [optional] 
**GcpParams** | Pointer to [**ArchivalGcpExternalTargetParams**](ArchivalGcpExternalTargetParams.md) |  | [optional] 
**NasParams** | Pointer to [**ArchivalNasExternalTargetParams**](ArchivalNasExternalTargetParams.md) |  | [optional] 
**OracleParams** | Pointer to [**ArchivalOracleExternalTargetParams**](ArchivalOracleExternalTargetParams.md) |  | [optional] 
**QstarTapeParams** | Pointer to [**ArchivalQstarTapeExternalTargetParams**](ArchivalQstarTapeExternalTargetParams.md) |  | [optional] 
**S3CompParams** | Pointer to [**ArchivalS3CompExternalTargetParams**](ArchivalS3CompExternalTargetParams.md) |  | [optional] 

## Methods

### NewArchivalExternalTargetParams

`func NewArchivalExternalTargetParams(encryption EncryptionSettings, storageType NullableString, ) *ArchivalExternalTargetParams`

NewArchivalExternalTargetParams instantiates a new ArchivalExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalExternalTargetParamsWithDefaults

`func NewArchivalExternalTargetParamsWithDefaults() *ArchivalExternalTargetParams`

NewArchivalExternalTargetParamsWithDefaults instantiates a new ArchivalExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryption

`func (o *ArchivalExternalTargetParams) GetEncryption() EncryptionSettings`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *ArchivalExternalTargetParams) GetEncryptionOk() (*EncryptionSettings, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *ArchivalExternalTargetParams) SetEncryption(v EncryptionSettings)`

SetEncryption sets Encryption field to given value.


### GetStorageType

`func (o *ArchivalExternalTargetParams) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ArchivalExternalTargetParams) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ArchivalExternalTargetParams) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### SetStorageTypeNil

`func (o *ArchivalExternalTargetParams) SetStorageTypeNil(b bool)`

 SetStorageTypeNil sets the value for StorageType to be an explicit nil

### UnsetStorageType
`func (o *ArchivalExternalTargetParams) UnsetStorageType()`

UnsetStorageType ensures that no value is present for StorageType, not even an explicit nil
### GetTargetBandwidthThrottlings

`func (o *ArchivalExternalTargetParams) GetTargetBandwidthThrottlings() TargetBandwidthThrottlings`

GetTargetBandwidthThrottlings returns the TargetBandwidthThrottlings field if non-nil, zero value otherwise.

### GetTargetBandwidthThrottlingsOk

`func (o *ArchivalExternalTargetParams) GetTargetBandwidthThrottlingsOk() (*TargetBandwidthThrottlings, bool)`

GetTargetBandwidthThrottlingsOk returns a tuple with the TargetBandwidthThrottlings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBandwidthThrottlings

`func (o *ArchivalExternalTargetParams) SetTargetBandwidthThrottlings(v TargetBandwidthThrottlings)`

SetTargetBandwidthThrottlings sets TargetBandwidthThrottlings field to given value.

### HasTargetBandwidthThrottlings

`func (o *ArchivalExternalTargetParams) HasTargetBandwidthThrottlings() bool`

HasTargetBandwidthThrottlings returns a boolean if a field has been set.

### GetAwsParams

`func (o *ArchivalExternalTargetParams) GetAwsParams() ArchivalAwsExternalTargetParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *ArchivalExternalTargetParams) GetAwsParamsOk() (*ArchivalAwsExternalTargetParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *ArchivalExternalTargetParams) SetAwsParams(v ArchivalAwsExternalTargetParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *ArchivalExternalTargetParams) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *ArchivalExternalTargetParams) GetAzureParams() ArchivalAzureExternalTargetParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *ArchivalExternalTargetParams) GetAzureParamsOk() (*ArchivalAzureExternalTargetParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *ArchivalExternalTargetParams) SetAzureParams(v ArchivalAzureExternalTargetParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *ArchivalExternalTargetParams) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *ArchivalExternalTargetParams) GetGcpParams() ArchivalGcpExternalTargetParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *ArchivalExternalTargetParams) GetGcpParamsOk() (*ArchivalGcpExternalTargetParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *ArchivalExternalTargetParams) SetGcpParams(v ArchivalGcpExternalTargetParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *ArchivalExternalTargetParams) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetNasParams

`func (o *ArchivalExternalTargetParams) GetNasParams() ArchivalNasExternalTargetParams`

GetNasParams returns the NasParams field if non-nil, zero value otherwise.

### GetNasParamsOk

`func (o *ArchivalExternalTargetParams) GetNasParamsOk() (*ArchivalNasExternalTargetParams, bool)`

GetNasParamsOk returns a tuple with the NasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasParams

`func (o *ArchivalExternalTargetParams) SetNasParams(v ArchivalNasExternalTargetParams)`

SetNasParams sets NasParams field to given value.

### HasNasParams

`func (o *ArchivalExternalTargetParams) HasNasParams() bool`

HasNasParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *ArchivalExternalTargetParams) GetOracleParams() ArchivalOracleExternalTargetParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ArchivalExternalTargetParams) GetOracleParamsOk() (*ArchivalOracleExternalTargetParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ArchivalExternalTargetParams) SetOracleParams(v ArchivalOracleExternalTargetParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ArchivalExternalTargetParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetQstarTapeParams

`func (o *ArchivalExternalTargetParams) GetQstarTapeParams() ArchivalQstarTapeExternalTargetParams`

GetQstarTapeParams returns the QstarTapeParams field if non-nil, zero value otherwise.

### GetQstarTapeParamsOk

`func (o *ArchivalExternalTargetParams) GetQstarTapeParamsOk() (*ArchivalQstarTapeExternalTargetParams, bool)`

GetQstarTapeParamsOk returns a tuple with the QstarTapeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQstarTapeParams

`func (o *ArchivalExternalTargetParams) SetQstarTapeParams(v ArchivalQstarTapeExternalTargetParams)`

SetQstarTapeParams sets QstarTapeParams field to given value.

### HasQstarTapeParams

`func (o *ArchivalExternalTargetParams) HasQstarTapeParams() bool`

HasQstarTapeParams returns a boolean if a field has been set.

### GetS3CompParams

`func (o *ArchivalExternalTargetParams) GetS3CompParams() ArchivalS3CompExternalTargetParams`

GetS3CompParams returns the S3CompParams field if non-nil, zero value otherwise.

### GetS3CompParamsOk

`func (o *ArchivalExternalTargetParams) GetS3CompParamsOk() (*ArchivalS3CompExternalTargetParams, bool)`

GetS3CompParamsOk returns a tuple with the S3CompParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CompParams

`func (o *ArchivalExternalTargetParams) SetS3CompParams(v ArchivalS3CompExternalTargetParams)`

SetS3CompParams sets S3CompParams field to given value.

### HasS3CompParams

`func (o *ArchivalExternalTargetParams) HasS3CompParams() bool`

HasS3CompParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


