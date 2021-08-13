# TieringExternalTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageType** | **NullableString** | Specifies the Storage type of the External Target. | 
**EncryptionLevel** | **NullableString** | Specifies the type of encryption for the Setting. | 
**AzureParams** | Pointer to [**TieringAzureExternalTargetParams**](TieringAzureExternalTargetParams.md) |  | [optional] 
**GcpParams** | Pointer to [**TieringGcpExternalTargetParams**](TieringGcpExternalTargetParams.md) |  | [optional] 
**AwsParams** | Pointer to [**TieringAwsExternalTargetParams**](TieringAwsExternalTargetParams.md) |  | [optional] 
**OracleParams** | Pointer to [**TieringOracleExternalTargetParams**](TieringOracleExternalTargetParams.md) |  | [optional] 
**S3CompParams** | Pointer to [**TieringS3CompExternalTargetParams**](TieringS3CompExternalTargetParams.md) |  | [optional] 

## Methods

### NewTieringExternalTargetParams

`func NewTieringExternalTargetParams(storageType NullableString, encryptionLevel NullableString, ) *TieringExternalTargetParams`

NewTieringExternalTargetParams instantiates a new TieringExternalTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieringExternalTargetParamsWithDefaults

`func NewTieringExternalTargetParamsWithDefaults() *TieringExternalTargetParams`

NewTieringExternalTargetParamsWithDefaults instantiates a new TieringExternalTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageType

`func (o *TieringExternalTargetParams) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *TieringExternalTargetParams) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *TieringExternalTargetParams) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.


### SetStorageTypeNil

`func (o *TieringExternalTargetParams) SetStorageTypeNil(b bool)`

 SetStorageTypeNil sets the value for StorageType to be an explicit nil

### UnsetStorageType
`func (o *TieringExternalTargetParams) UnsetStorageType()`

UnsetStorageType ensures that no value is present for StorageType, not even an explicit nil
### GetEncryptionLevel

`func (o *TieringExternalTargetParams) GetEncryptionLevel() string`

GetEncryptionLevel returns the EncryptionLevel field if non-nil, zero value otherwise.

### GetEncryptionLevelOk

`func (o *TieringExternalTargetParams) GetEncryptionLevelOk() (*string, bool)`

GetEncryptionLevelOk returns a tuple with the EncryptionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionLevel

`func (o *TieringExternalTargetParams) SetEncryptionLevel(v string)`

SetEncryptionLevel sets EncryptionLevel field to given value.


### SetEncryptionLevelNil

`func (o *TieringExternalTargetParams) SetEncryptionLevelNil(b bool)`

 SetEncryptionLevelNil sets the value for EncryptionLevel to be an explicit nil

### UnsetEncryptionLevel
`func (o *TieringExternalTargetParams) UnsetEncryptionLevel()`

UnsetEncryptionLevel ensures that no value is present for EncryptionLevel, not even an explicit nil
### GetAzureParams

`func (o *TieringExternalTargetParams) GetAzureParams() TieringAzureExternalTargetParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *TieringExternalTargetParams) GetAzureParamsOk() (*TieringAzureExternalTargetParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *TieringExternalTargetParams) SetAzureParams(v TieringAzureExternalTargetParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *TieringExternalTargetParams) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *TieringExternalTargetParams) GetGcpParams() TieringGcpExternalTargetParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *TieringExternalTargetParams) GetGcpParamsOk() (*TieringGcpExternalTargetParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *TieringExternalTargetParams) SetGcpParams(v TieringGcpExternalTargetParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *TieringExternalTargetParams) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *TieringExternalTargetParams) GetAwsParams() TieringAwsExternalTargetParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *TieringExternalTargetParams) GetAwsParamsOk() (*TieringAwsExternalTargetParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *TieringExternalTargetParams) SetAwsParams(v TieringAwsExternalTargetParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *TieringExternalTargetParams) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *TieringExternalTargetParams) GetOracleParams() TieringOracleExternalTargetParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *TieringExternalTargetParams) GetOracleParamsOk() (*TieringOracleExternalTargetParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *TieringExternalTargetParams) SetOracleParams(v TieringOracleExternalTargetParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *TieringExternalTargetParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetS3CompParams

`func (o *TieringExternalTargetParams) GetS3CompParams() TieringS3CompExternalTargetParams`

GetS3CompParams returns the S3CompParams field if non-nil, zero value otherwise.

### GetS3CompParamsOk

`func (o *TieringExternalTargetParams) GetS3CompParamsOk() (*TieringS3CompExternalTargetParams, bool)`

GetS3CompParamsOk returns a tuple with the S3CompParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CompParams

`func (o *TieringExternalTargetParams) SetS3CompParams(v TieringS3CompExternalTargetParams)`

SetS3CompParams sets S3CompParams field to given value.

### HasS3CompParams

`func (o *TieringExternalTargetParams) HasS3CompParams() bool`

HasS3CompParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


