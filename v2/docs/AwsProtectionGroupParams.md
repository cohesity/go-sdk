# AwsProtectionGroupParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentProtectionTypeParams** | Pointer to [**AwsAgentProtectionGroupParams**](AwsAgentProtectionGroupParams.md) |  | [optional] 
**AuroraProtectionTypeParams** | Pointer to [**AwsAuroraProtectionGroupParams**](AwsAuroraProtectionGroupParams.md) |  | [optional] 
**NativeProtectionTypeParams** | Pointer to [**AwsNativeProtectionGroupParams**](AwsNativeProtectionGroupParams.md) |  | [optional] 
**ProtectionType** | **string** | Specifies the AWS Protection Group type. | 
**RdsPostgresProtectionTypeParams** | Pointer to [**AwsRdsPostgresProtectionGroupParams**](AwsRdsPostgresProtectionGroupParams.md) |  | [optional] 
**RdsProtectionTypeParams** | Pointer to [**AwsRdsProtectionGroupParams**](AwsRdsProtectionGroupParams.md) |  | [optional] 
**S3ProtectionTypeParams** | Pointer to [**AwsS3ProtectionGroupParams**](AwsS3ProtectionGroupParams.md) |  | [optional] 
**SnapshotManagerProtectionTypeParams** | Pointer to [**AwsSnapshotManagerProtectionGroupParams**](AwsSnapshotManagerProtectionGroupParams.md) |  | [optional] 

## Methods

### NewAwsProtectionGroupParams

`func NewAwsProtectionGroupParams(protectionType string, ) *AwsProtectionGroupParams`

NewAwsProtectionGroupParams instantiates a new AwsProtectionGroupParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsProtectionGroupParamsWithDefaults

`func NewAwsProtectionGroupParamsWithDefaults() *AwsProtectionGroupParams`

NewAwsProtectionGroupParamsWithDefaults instantiates a new AwsProtectionGroupParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentProtectionTypeParams

`func (o *AwsProtectionGroupParams) GetAgentProtectionTypeParams() AwsAgentProtectionGroupParams`

GetAgentProtectionTypeParams returns the AgentProtectionTypeParams field if non-nil, zero value otherwise.

### GetAgentProtectionTypeParamsOk

`func (o *AwsProtectionGroupParams) GetAgentProtectionTypeParamsOk() (*AwsAgentProtectionGroupParams, bool)`

GetAgentProtectionTypeParamsOk returns a tuple with the AgentProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProtectionTypeParams

`func (o *AwsProtectionGroupParams) SetAgentProtectionTypeParams(v AwsAgentProtectionGroupParams)`

SetAgentProtectionTypeParams sets AgentProtectionTypeParams field to given value.

### HasAgentProtectionTypeParams

`func (o *AwsProtectionGroupParams) HasAgentProtectionTypeParams() bool`

HasAgentProtectionTypeParams returns a boolean if a field has been set.

### GetAuroraProtectionTypeParams

`func (o *AwsProtectionGroupParams) GetAuroraProtectionTypeParams() AwsAuroraProtectionGroupParams`

GetAuroraProtectionTypeParams returns the AuroraProtectionTypeParams field if non-nil, zero value otherwise.

### GetAuroraProtectionTypeParamsOk

`func (o *AwsProtectionGroupParams) GetAuroraProtectionTypeParamsOk() (*AwsAuroraProtectionGroupParams, bool)`

GetAuroraProtectionTypeParamsOk returns a tuple with the AuroraProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuroraProtectionTypeParams

`func (o *AwsProtectionGroupParams) SetAuroraProtectionTypeParams(v AwsAuroraProtectionGroupParams)`

SetAuroraProtectionTypeParams sets AuroraProtectionTypeParams field to given value.

### HasAuroraProtectionTypeParams

`func (o *AwsProtectionGroupParams) HasAuroraProtectionTypeParams() bool`

HasAuroraProtectionTypeParams returns a boolean if a field has been set.

### GetNativeProtectionTypeParams

`func (o *AwsProtectionGroupParams) GetNativeProtectionTypeParams() AwsNativeProtectionGroupParams`

GetNativeProtectionTypeParams returns the NativeProtectionTypeParams field if non-nil, zero value otherwise.

### GetNativeProtectionTypeParamsOk

`func (o *AwsProtectionGroupParams) GetNativeProtectionTypeParamsOk() (*AwsNativeProtectionGroupParams, bool)`

GetNativeProtectionTypeParamsOk returns a tuple with the NativeProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeProtectionTypeParams

`func (o *AwsProtectionGroupParams) SetNativeProtectionTypeParams(v AwsNativeProtectionGroupParams)`

SetNativeProtectionTypeParams sets NativeProtectionTypeParams field to given value.

### HasNativeProtectionTypeParams

`func (o *AwsProtectionGroupParams) HasNativeProtectionTypeParams() bool`

HasNativeProtectionTypeParams returns a boolean if a field has been set.

### GetProtectionType

`func (o *AwsProtectionGroupParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *AwsProtectionGroupParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *AwsProtectionGroupParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.


### GetRdsPostgresProtectionTypeParams

`func (o *AwsProtectionGroupParams) GetRdsPostgresProtectionTypeParams() AwsRdsPostgresProtectionGroupParams`

GetRdsPostgresProtectionTypeParams returns the RdsPostgresProtectionTypeParams field if non-nil, zero value otherwise.

### GetRdsPostgresProtectionTypeParamsOk

`func (o *AwsProtectionGroupParams) GetRdsPostgresProtectionTypeParamsOk() (*AwsRdsPostgresProtectionGroupParams, bool)`

GetRdsPostgresProtectionTypeParamsOk returns a tuple with the RdsPostgresProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsPostgresProtectionTypeParams

`func (o *AwsProtectionGroupParams) SetRdsPostgresProtectionTypeParams(v AwsRdsPostgresProtectionGroupParams)`

SetRdsPostgresProtectionTypeParams sets RdsPostgresProtectionTypeParams field to given value.

### HasRdsPostgresProtectionTypeParams

`func (o *AwsProtectionGroupParams) HasRdsPostgresProtectionTypeParams() bool`

HasRdsPostgresProtectionTypeParams returns a boolean if a field has been set.

### GetRdsProtectionTypeParams

`func (o *AwsProtectionGroupParams) GetRdsProtectionTypeParams() AwsRdsProtectionGroupParams`

GetRdsProtectionTypeParams returns the RdsProtectionTypeParams field if non-nil, zero value otherwise.

### GetRdsProtectionTypeParamsOk

`func (o *AwsProtectionGroupParams) GetRdsProtectionTypeParamsOk() (*AwsRdsProtectionGroupParams, bool)`

GetRdsProtectionTypeParamsOk returns a tuple with the RdsProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsProtectionTypeParams

`func (o *AwsProtectionGroupParams) SetRdsProtectionTypeParams(v AwsRdsProtectionGroupParams)`

SetRdsProtectionTypeParams sets RdsProtectionTypeParams field to given value.

### HasRdsProtectionTypeParams

`func (o *AwsProtectionGroupParams) HasRdsProtectionTypeParams() bool`

HasRdsProtectionTypeParams returns a boolean if a field has been set.

### GetS3ProtectionTypeParams

`func (o *AwsProtectionGroupParams) GetS3ProtectionTypeParams() AwsS3ProtectionGroupParams`

GetS3ProtectionTypeParams returns the S3ProtectionTypeParams field if non-nil, zero value otherwise.

### GetS3ProtectionTypeParamsOk

`func (o *AwsProtectionGroupParams) GetS3ProtectionTypeParamsOk() (*AwsS3ProtectionGroupParams, bool)`

GetS3ProtectionTypeParamsOk returns a tuple with the S3ProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ProtectionTypeParams

`func (o *AwsProtectionGroupParams) SetS3ProtectionTypeParams(v AwsS3ProtectionGroupParams)`

SetS3ProtectionTypeParams sets S3ProtectionTypeParams field to given value.

### HasS3ProtectionTypeParams

`func (o *AwsProtectionGroupParams) HasS3ProtectionTypeParams() bool`

HasS3ProtectionTypeParams returns a boolean if a field has been set.

### GetSnapshotManagerProtectionTypeParams

`func (o *AwsProtectionGroupParams) GetSnapshotManagerProtectionTypeParams() AwsSnapshotManagerProtectionGroupParams`

GetSnapshotManagerProtectionTypeParams returns the SnapshotManagerProtectionTypeParams field if non-nil, zero value otherwise.

### GetSnapshotManagerProtectionTypeParamsOk

`func (o *AwsProtectionGroupParams) GetSnapshotManagerProtectionTypeParamsOk() (*AwsSnapshotManagerProtectionGroupParams, bool)`

GetSnapshotManagerProtectionTypeParamsOk returns a tuple with the SnapshotManagerProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotManagerProtectionTypeParams

`func (o *AwsProtectionGroupParams) SetSnapshotManagerProtectionTypeParams(v AwsSnapshotManagerProtectionGroupParams)`

SetSnapshotManagerProtectionTypeParams sets SnapshotManagerProtectionTypeParams field to given value.

### HasSnapshotManagerProtectionTypeParams

`func (o *AwsProtectionGroupParams) HasSnapshotManagerProtectionTypeParams() bool`

HasSnapshotManagerProtectionTypeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


