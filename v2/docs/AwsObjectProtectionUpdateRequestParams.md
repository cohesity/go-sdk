# AwsObjectProtectionUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionType** | Pointer to **string** | Specifies the AWS Protection Job type. | [optional] 
**AuroraSnapshotManagerProtectionTypeParams** | Pointer to [**AwsAuroraSnapshotManagerObjectProtectionParams**](AwsAuroraSnapshotManagerObjectProtectionParams.md) |  | [optional] 
**DynamoDBProtectionTypeParams** | Pointer to [**AwsDynamoDBProtectionParams**](AwsDynamoDBProtectionParams.md) |  | [optional] 
**NativeProtectionTypeParams** | Pointer to [**AwsNativeObjectProtectionParams**](AwsNativeObjectProtectionParams.md) |  | [optional] 
**RdsPostgresProtectionTypeParams** | Pointer to [**AwsRdsPostgresProtectionParams**](AwsRdsPostgresProtectionParams.md) |  | [optional] 
**RdsSnapshotManagerProtectionTypeParams** | Pointer to [**AwsRdsSnapshotManagerObjectProtectionParams**](AwsRdsSnapshotManagerObjectProtectionParams.md) |  | [optional] 
**S3ProtectionTypeParams** | Pointer to [**AwsS3ProtectionParams**](AwsS3ProtectionParams.md) |  | [optional] 
**SnapshotManagerProtectionTypeParams** | Pointer to [**AwsSnapshotManagerObjectProtectionParams**](AwsSnapshotManagerObjectProtectionParams.md) |  | [optional] 

## Methods

### NewAwsObjectProtectionUpdateRequestParams

`func NewAwsObjectProtectionUpdateRequestParams() *AwsObjectProtectionUpdateRequestParams`

NewAwsObjectProtectionUpdateRequestParams instantiates a new AwsObjectProtectionUpdateRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsObjectProtectionUpdateRequestParamsWithDefaults

`func NewAwsObjectProtectionUpdateRequestParamsWithDefaults() *AwsObjectProtectionUpdateRequestParams`

NewAwsObjectProtectionUpdateRequestParamsWithDefaults instantiates a new AwsObjectProtectionUpdateRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionType

`func (o *AwsObjectProtectionUpdateRequestParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *AwsObjectProtectionUpdateRequestParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *AwsObjectProtectionUpdateRequestParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *AwsObjectProtectionUpdateRequestParams) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### GetAuroraSnapshotManagerProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) GetAuroraSnapshotManagerProtectionTypeParams() AwsAuroraSnapshotManagerObjectProtectionParams`

GetAuroraSnapshotManagerProtectionTypeParams returns the AuroraSnapshotManagerProtectionTypeParams field if non-nil, zero value otherwise.

### GetAuroraSnapshotManagerProtectionTypeParamsOk

`func (o *AwsObjectProtectionUpdateRequestParams) GetAuroraSnapshotManagerProtectionTypeParamsOk() (*AwsAuroraSnapshotManagerObjectProtectionParams, bool)`

GetAuroraSnapshotManagerProtectionTypeParamsOk returns a tuple with the AuroraSnapshotManagerProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuroraSnapshotManagerProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) SetAuroraSnapshotManagerProtectionTypeParams(v AwsAuroraSnapshotManagerObjectProtectionParams)`

SetAuroraSnapshotManagerProtectionTypeParams sets AuroraSnapshotManagerProtectionTypeParams field to given value.

### HasAuroraSnapshotManagerProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) HasAuroraSnapshotManagerProtectionTypeParams() bool`

HasAuroraSnapshotManagerProtectionTypeParams returns a boolean if a field has been set.

### GetDynamoDBProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) GetDynamoDBProtectionTypeParams() AwsDynamoDBProtectionParams`

GetDynamoDBProtectionTypeParams returns the DynamoDBProtectionTypeParams field if non-nil, zero value otherwise.

### GetDynamoDBProtectionTypeParamsOk

`func (o *AwsObjectProtectionUpdateRequestParams) GetDynamoDBProtectionTypeParamsOk() (*AwsDynamoDBProtectionParams, bool)`

GetDynamoDBProtectionTypeParamsOk returns a tuple with the DynamoDBProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamoDBProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) SetDynamoDBProtectionTypeParams(v AwsDynamoDBProtectionParams)`

SetDynamoDBProtectionTypeParams sets DynamoDBProtectionTypeParams field to given value.

### HasDynamoDBProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) HasDynamoDBProtectionTypeParams() bool`

HasDynamoDBProtectionTypeParams returns a boolean if a field has been set.

### GetNativeProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) GetNativeProtectionTypeParams() AwsNativeObjectProtectionParams`

GetNativeProtectionTypeParams returns the NativeProtectionTypeParams field if non-nil, zero value otherwise.

### GetNativeProtectionTypeParamsOk

`func (o *AwsObjectProtectionUpdateRequestParams) GetNativeProtectionTypeParamsOk() (*AwsNativeObjectProtectionParams, bool)`

GetNativeProtectionTypeParamsOk returns a tuple with the NativeProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) SetNativeProtectionTypeParams(v AwsNativeObjectProtectionParams)`

SetNativeProtectionTypeParams sets NativeProtectionTypeParams field to given value.

### HasNativeProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) HasNativeProtectionTypeParams() bool`

HasNativeProtectionTypeParams returns a boolean if a field has been set.

### GetRdsPostgresProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) GetRdsPostgresProtectionTypeParams() AwsRdsPostgresProtectionParams`

GetRdsPostgresProtectionTypeParams returns the RdsPostgresProtectionTypeParams field if non-nil, zero value otherwise.

### GetRdsPostgresProtectionTypeParamsOk

`func (o *AwsObjectProtectionUpdateRequestParams) GetRdsPostgresProtectionTypeParamsOk() (*AwsRdsPostgresProtectionParams, bool)`

GetRdsPostgresProtectionTypeParamsOk returns a tuple with the RdsPostgresProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsPostgresProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) SetRdsPostgresProtectionTypeParams(v AwsRdsPostgresProtectionParams)`

SetRdsPostgresProtectionTypeParams sets RdsPostgresProtectionTypeParams field to given value.

### HasRdsPostgresProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) HasRdsPostgresProtectionTypeParams() bool`

HasRdsPostgresProtectionTypeParams returns a boolean if a field has been set.

### GetRdsSnapshotManagerProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) GetRdsSnapshotManagerProtectionTypeParams() AwsRdsSnapshotManagerObjectProtectionParams`

GetRdsSnapshotManagerProtectionTypeParams returns the RdsSnapshotManagerProtectionTypeParams field if non-nil, zero value otherwise.

### GetRdsSnapshotManagerProtectionTypeParamsOk

`func (o *AwsObjectProtectionUpdateRequestParams) GetRdsSnapshotManagerProtectionTypeParamsOk() (*AwsRdsSnapshotManagerObjectProtectionParams, bool)`

GetRdsSnapshotManagerProtectionTypeParamsOk returns a tuple with the RdsSnapshotManagerProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSnapshotManagerProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) SetRdsSnapshotManagerProtectionTypeParams(v AwsRdsSnapshotManagerObjectProtectionParams)`

SetRdsSnapshotManagerProtectionTypeParams sets RdsSnapshotManagerProtectionTypeParams field to given value.

### HasRdsSnapshotManagerProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) HasRdsSnapshotManagerProtectionTypeParams() bool`

HasRdsSnapshotManagerProtectionTypeParams returns a boolean if a field has been set.

### GetS3ProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) GetS3ProtectionTypeParams() AwsS3ProtectionParams`

GetS3ProtectionTypeParams returns the S3ProtectionTypeParams field if non-nil, zero value otherwise.

### GetS3ProtectionTypeParamsOk

`func (o *AwsObjectProtectionUpdateRequestParams) GetS3ProtectionTypeParamsOk() (*AwsS3ProtectionParams, bool)`

GetS3ProtectionTypeParamsOk returns a tuple with the S3ProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) SetS3ProtectionTypeParams(v AwsS3ProtectionParams)`

SetS3ProtectionTypeParams sets S3ProtectionTypeParams field to given value.

### HasS3ProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) HasS3ProtectionTypeParams() bool`

HasS3ProtectionTypeParams returns a boolean if a field has been set.

### GetSnapshotManagerProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) GetSnapshotManagerProtectionTypeParams() AwsSnapshotManagerObjectProtectionParams`

GetSnapshotManagerProtectionTypeParams returns the SnapshotManagerProtectionTypeParams field if non-nil, zero value otherwise.

### GetSnapshotManagerProtectionTypeParamsOk

`func (o *AwsObjectProtectionUpdateRequestParams) GetSnapshotManagerProtectionTypeParamsOk() (*AwsSnapshotManagerObjectProtectionParams, bool)`

GetSnapshotManagerProtectionTypeParamsOk returns a tuple with the SnapshotManagerProtectionTypeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotManagerProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) SetSnapshotManagerProtectionTypeParams(v AwsSnapshotManagerObjectProtectionParams)`

SetSnapshotManagerProtectionTypeParams sets SnapshotManagerProtectionTypeParams field to given value.

### HasSnapshotManagerProtectionTypeParams

`func (o *AwsObjectProtectionUpdateRequestParams) HasSnapshotManagerProtectionTypeParams() bool`

HasSnapshotManagerProtectionTypeParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


