# CloudDeployTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployVmsToCloudParams** | Pointer to [**DeployVMsToCloudParams**](DeployVMsToCloudParams.md) |  | [optional] 
**TargetEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**Type** | Pointer to **NullableInt32** | The type of the CloudDeploy target. | [optional] 

## Methods

### NewCloudDeployTarget

`func NewCloudDeployTarget() *CloudDeployTarget`

NewCloudDeployTarget instantiates a new CloudDeployTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDeployTargetWithDefaults

`func NewCloudDeployTargetWithDefaults() *CloudDeployTarget`

NewCloudDeployTargetWithDefaults instantiates a new CloudDeployTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployVmsToCloudParams

`func (o *CloudDeployTarget) GetDeployVmsToCloudParams() DeployVMsToCloudParams`

GetDeployVmsToCloudParams returns the DeployVmsToCloudParams field if non-nil, zero value otherwise.

### GetDeployVmsToCloudParamsOk

`func (o *CloudDeployTarget) GetDeployVmsToCloudParamsOk() (*DeployVMsToCloudParams, bool)`

GetDeployVmsToCloudParamsOk returns a tuple with the DeployVmsToCloudParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployVmsToCloudParams

`func (o *CloudDeployTarget) SetDeployVmsToCloudParams(v DeployVMsToCloudParams)`

SetDeployVmsToCloudParams sets DeployVmsToCloudParams field to given value.

### HasDeployVmsToCloudParams

`func (o *CloudDeployTarget) HasDeployVmsToCloudParams() bool`

HasDeployVmsToCloudParams returns a boolean if a field has been set.

### GetTargetEntity

`func (o *CloudDeployTarget) GetTargetEntity() EntityProto`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *CloudDeployTarget) GetTargetEntityOk() (*EntityProto, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *CloudDeployTarget) SetTargetEntity(v EntityProto)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *CloudDeployTarget) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.

### GetType

`func (o *CloudDeployTarget) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudDeployTarget) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudDeployTarget) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CloudDeployTarget) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CloudDeployTarget) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CloudDeployTarget) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


