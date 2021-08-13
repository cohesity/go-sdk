# ArchivalExternalTargetParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureParams** | Pointer to [**ArchivalAzureExternalTargetParams**](ArchivalAzureExternalTargetParams.md) |  | [optional] 
**GcpParams** | Pointer to [**ArchivalGcpExternalTargetParams**](ArchivalGcpExternalTargetParams.md) |  | [optional] 
**AwsParams** | Pointer to [**ArchivalAwsExternalTargetParams**](ArchivalAwsExternalTargetParams.md) |  | [optional] 
**OracleParams** | Pointer to [**ArchivalOracleExternalTargetParams**](ArchivalOracleExternalTargetParams.md) |  | [optional] 
**NasParams** | Pointer to [**ArchivalNasExternalTargetParams**](ArchivalNasExternalTargetParams.md) |  | [optional] 
**QstarTapeParams** | Pointer to [**ArchivalQstarTapeExternalTargetParams**](ArchivalQstarTapeExternalTargetParams.md) |  | [optional] 
**S3CompParams** | Pointer to [**ArchivalS3CompExternalTargetParams**](ArchivalS3CompExternalTargetParams.md) |  | [optional] 

## Methods

### NewArchivalExternalTargetParamsAllOf

`func NewArchivalExternalTargetParamsAllOf() *ArchivalExternalTargetParamsAllOf`

NewArchivalExternalTargetParamsAllOf instantiates a new ArchivalExternalTargetParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalExternalTargetParamsAllOfWithDefaults

`func NewArchivalExternalTargetParamsAllOfWithDefaults() *ArchivalExternalTargetParamsAllOf`

NewArchivalExternalTargetParamsAllOfWithDefaults instantiates a new ArchivalExternalTargetParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureParams

`func (o *ArchivalExternalTargetParamsAllOf) GetAzureParams() ArchivalAzureExternalTargetParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *ArchivalExternalTargetParamsAllOf) GetAzureParamsOk() (*ArchivalAzureExternalTargetParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *ArchivalExternalTargetParamsAllOf) SetAzureParams(v ArchivalAzureExternalTargetParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *ArchivalExternalTargetParamsAllOf) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetGcpParams

`func (o *ArchivalExternalTargetParamsAllOf) GetGcpParams() ArchivalGcpExternalTargetParams`

GetGcpParams returns the GcpParams field if non-nil, zero value otherwise.

### GetGcpParamsOk

`func (o *ArchivalExternalTargetParamsAllOf) GetGcpParamsOk() (*ArchivalGcpExternalTargetParams, bool)`

GetGcpParamsOk returns a tuple with the GcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpParams

`func (o *ArchivalExternalTargetParamsAllOf) SetGcpParams(v ArchivalGcpExternalTargetParams)`

SetGcpParams sets GcpParams field to given value.

### HasGcpParams

`func (o *ArchivalExternalTargetParamsAllOf) HasGcpParams() bool`

HasGcpParams returns a boolean if a field has been set.

### GetAwsParams

`func (o *ArchivalExternalTargetParamsAllOf) GetAwsParams() ArchivalAwsExternalTargetParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *ArchivalExternalTargetParamsAllOf) GetAwsParamsOk() (*ArchivalAwsExternalTargetParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *ArchivalExternalTargetParamsAllOf) SetAwsParams(v ArchivalAwsExternalTargetParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *ArchivalExternalTargetParamsAllOf) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetOracleParams

`func (o *ArchivalExternalTargetParamsAllOf) GetOracleParams() ArchivalOracleExternalTargetParams`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *ArchivalExternalTargetParamsAllOf) GetOracleParamsOk() (*ArchivalOracleExternalTargetParams, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *ArchivalExternalTargetParamsAllOf) SetOracleParams(v ArchivalOracleExternalTargetParams)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *ArchivalExternalTargetParamsAllOf) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.

### GetNasParams

`func (o *ArchivalExternalTargetParamsAllOf) GetNasParams() ArchivalNasExternalTargetParams`

GetNasParams returns the NasParams field if non-nil, zero value otherwise.

### GetNasParamsOk

`func (o *ArchivalExternalTargetParamsAllOf) GetNasParamsOk() (*ArchivalNasExternalTargetParams, bool)`

GetNasParamsOk returns a tuple with the NasParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasParams

`func (o *ArchivalExternalTargetParamsAllOf) SetNasParams(v ArchivalNasExternalTargetParams)`

SetNasParams sets NasParams field to given value.

### HasNasParams

`func (o *ArchivalExternalTargetParamsAllOf) HasNasParams() bool`

HasNasParams returns a boolean if a field has been set.

### GetQstarTapeParams

`func (o *ArchivalExternalTargetParamsAllOf) GetQstarTapeParams() ArchivalQstarTapeExternalTargetParams`

GetQstarTapeParams returns the QstarTapeParams field if non-nil, zero value otherwise.

### GetQstarTapeParamsOk

`func (o *ArchivalExternalTargetParamsAllOf) GetQstarTapeParamsOk() (*ArchivalQstarTapeExternalTargetParams, bool)`

GetQstarTapeParamsOk returns a tuple with the QstarTapeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQstarTapeParams

`func (o *ArchivalExternalTargetParamsAllOf) SetQstarTapeParams(v ArchivalQstarTapeExternalTargetParams)`

SetQstarTapeParams sets QstarTapeParams field to given value.

### HasQstarTapeParams

`func (o *ArchivalExternalTargetParamsAllOf) HasQstarTapeParams() bool`

HasQstarTapeParams returns a boolean if a field has been set.

### GetS3CompParams

`func (o *ArchivalExternalTargetParamsAllOf) GetS3CompParams() ArchivalS3CompExternalTargetParams`

GetS3CompParams returns the S3CompParams field if non-nil, zero value otherwise.

### GetS3CompParamsOk

`func (o *ArchivalExternalTargetParamsAllOf) GetS3CompParamsOk() (*ArchivalS3CompExternalTargetParams, bool)`

GetS3CompParamsOk returns a tuple with the S3CompParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CompParams

`func (o *ArchivalExternalTargetParamsAllOf) SetS3CompParams(v ArchivalS3CompExternalTargetParams)`

SetS3CompParams sets S3CompParams field to given value.

### HasS3CompParams

`func (o *ArchivalExternalTargetParamsAllOf) HasS3CompParams() bool`

HasS3CompParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


