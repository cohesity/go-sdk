# AwsObjectProtectionUpdateRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionType** | Pointer to **string** | Specifies the AWS Protection Job type. | [optional] 
**NativeProtectionTypeParams** | Pointer to [**AwsNativeObjectProtectionParams**](AwsNativeObjectProtectionParams.md) |  | [optional] 
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


