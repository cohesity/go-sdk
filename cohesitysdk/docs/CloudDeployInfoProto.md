# CloudDeployInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudDeployEntityVec** | Pointer to [**[]CloudDeployInfoProtoCloudDeployEntity**](CloudDeployInfoProtoCloudDeployEntity.md) | Contains the file paths and the information of the entities deployed to cloud. | [optional] 
**IsIncremental** | Pointer to **NullableBool** | Whether this Cloud deploy info is for incremental cloudspin. | [optional] 
**RestoreInfo** | Pointer to [**RestoreInfoProto**](RestoreInfoProto.md) |  | [optional] 
**TargetType** | Pointer to **NullableInt32** | Specifies the target type for the task. The field is only valid if the task has got a permit. | [optional] 
**TotalBytesTransferredToSource** | Pointer to **NullableInt64** | Total bytes transferred to source. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment this cloud deploy info pertains to. | [optional] 

## Methods

### NewCloudDeployInfoProto

`func NewCloudDeployInfoProto() *CloudDeployInfoProto`

NewCloudDeployInfoProto instantiates a new CloudDeployInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDeployInfoProtoWithDefaults

`func NewCloudDeployInfoProtoWithDefaults() *CloudDeployInfoProto`

NewCloudDeployInfoProtoWithDefaults instantiates a new CloudDeployInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudDeployEntityVec

`func (o *CloudDeployInfoProto) GetCloudDeployEntityVec() []CloudDeployInfoProtoCloudDeployEntity`

GetCloudDeployEntityVec returns the CloudDeployEntityVec field if non-nil, zero value otherwise.

### GetCloudDeployEntityVecOk

`func (o *CloudDeployInfoProto) GetCloudDeployEntityVecOk() (*[]CloudDeployInfoProtoCloudDeployEntity, bool)`

GetCloudDeployEntityVecOk returns a tuple with the CloudDeployEntityVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudDeployEntityVec

`func (o *CloudDeployInfoProto) SetCloudDeployEntityVec(v []CloudDeployInfoProtoCloudDeployEntity)`

SetCloudDeployEntityVec sets CloudDeployEntityVec field to given value.

### HasCloudDeployEntityVec

`func (o *CloudDeployInfoProto) HasCloudDeployEntityVec() bool`

HasCloudDeployEntityVec returns a boolean if a field has been set.

### SetCloudDeployEntityVecNil

`func (o *CloudDeployInfoProto) SetCloudDeployEntityVecNil(b bool)`

 SetCloudDeployEntityVecNil sets the value for CloudDeployEntityVec to be an explicit nil

### UnsetCloudDeployEntityVec
`func (o *CloudDeployInfoProto) UnsetCloudDeployEntityVec()`

UnsetCloudDeployEntityVec ensures that no value is present for CloudDeployEntityVec, not even an explicit nil
### GetIsIncremental

`func (o *CloudDeployInfoProto) GetIsIncremental() bool`

GetIsIncremental returns the IsIncremental field if non-nil, zero value otherwise.

### GetIsIncrementalOk

`func (o *CloudDeployInfoProto) GetIsIncrementalOk() (*bool, bool)`

GetIsIncrementalOk returns a tuple with the IsIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncremental

`func (o *CloudDeployInfoProto) SetIsIncremental(v bool)`

SetIsIncremental sets IsIncremental field to given value.

### HasIsIncremental

`func (o *CloudDeployInfoProto) HasIsIncremental() bool`

HasIsIncremental returns a boolean if a field has been set.

### SetIsIncrementalNil

`func (o *CloudDeployInfoProto) SetIsIncrementalNil(b bool)`

 SetIsIncrementalNil sets the value for IsIncremental to be an explicit nil

### UnsetIsIncremental
`func (o *CloudDeployInfoProto) UnsetIsIncremental()`

UnsetIsIncremental ensures that no value is present for IsIncremental, not even an explicit nil
### GetRestoreInfo

`func (o *CloudDeployInfoProto) GetRestoreInfo() RestoreInfoProto`

GetRestoreInfo returns the RestoreInfo field if non-nil, zero value otherwise.

### GetRestoreInfoOk

`func (o *CloudDeployInfoProto) GetRestoreInfoOk() (*RestoreInfoProto, bool)`

GetRestoreInfoOk returns a tuple with the RestoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreInfo

`func (o *CloudDeployInfoProto) SetRestoreInfo(v RestoreInfoProto)`

SetRestoreInfo sets RestoreInfo field to given value.

### HasRestoreInfo

`func (o *CloudDeployInfoProto) HasRestoreInfo() bool`

HasRestoreInfo returns a boolean if a field has been set.

### GetTargetType

`func (o *CloudDeployInfoProto) GetTargetType() int32`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CloudDeployInfoProto) GetTargetTypeOk() (*int32, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CloudDeployInfoProto) SetTargetType(v int32)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CloudDeployInfoProto) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *CloudDeployInfoProto) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *CloudDeployInfoProto) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTotalBytesTransferredToSource

`func (o *CloudDeployInfoProto) GetTotalBytesTransferredToSource() int64`

GetTotalBytesTransferredToSource returns the TotalBytesTransferredToSource field if non-nil, zero value otherwise.

### GetTotalBytesTransferredToSourceOk

`func (o *CloudDeployInfoProto) GetTotalBytesTransferredToSourceOk() (*int64, bool)`

GetTotalBytesTransferredToSourceOk returns a tuple with the TotalBytesTransferredToSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesTransferredToSource

`func (o *CloudDeployInfoProto) SetTotalBytesTransferredToSource(v int64)`

SetTotalBytesTransferredToSource sets TotalBytesTransferredToSource field to given value.

### HasTotalBytesTransferredToSource

`func (o *CloudDeployInfoProto) HasTotalBytesTransferredToSource() bool`

HasTotalBytesTransferredToSource returns a boolean if a field has been set.

### SetTotalBytesTransferredToSourceNil

`func (o *CloudDeployInfoProto) SetTotalBytesTransferredToSourceNil(b bool)`

 SetTotalBytesTransferredToSourceNil sets the value for TotalBytesTransferredToSource to be an explicit nil

### UnsetTotalBytesTransferredToSource
`func (o *CloudDeployInfoProto) UnsetTotalBytesTransferredToSource()`

UnsetTotalBytesTransferredToSource ensures that no value is present for TotalBytesTransferredToSource, not even an explicit nil
### GetType

`func (o *CloudDeployInfoProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudDeployInfoProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudDeployInfoProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CloudDeployInfoProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CloudDeployInfoProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CloudDeployInfoProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


