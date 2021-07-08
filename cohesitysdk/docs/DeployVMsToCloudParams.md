# DeployVMsToCloudParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployFleetParams** | Pointer to [**DeployFleetParams**](DeployFleetParams.md) |  | [optional] 
**DeployVmsToAwsParams** | Pointer to [**DeployVMsToAWSParams**](DeployVMsToAWSParams.md) |  | [optional] 
**DeployVmsToAzureParams** | Pointer to [**DeployVMsToAzureParams**](DeployVMsToAzureParams.md) |  | [optional] 
**DeployVmsToGcpParams** | Pointer to [**DeployVMsToGCPParams**](DeployVMsToGCPParams.md) |  | [optional] 
**ReplicateSnapshotsToAwsParams** | Pointer to [**ReplicateSnapshotsToAWSParams**](ReplicateSnapshotsToAWSParams.md) |  | [optional] 
**ReplicateSnapshotsToAzureParams** | Pointer to [**ReplicateSnapshotsToAzureParams**](ReplicateSnapshotsToAzureParams.md) |  | [optional] 

## Methods

### NewDeployVMsToCloudParams

`func NewDeployVMsToCloudParams() *DeployVMsToCloudParams`

NewDeployVMsToCloudParams instantiates a new DeployVMsToCloudParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployVMsToCloudParamsWithDefaults

`func NewDeployVMsToCloudParamsWithDefaults() *DeployVMsToCloudParams`

NewDeployVMsToCloudParamsWithDefaults instantiates a new DeployVMsToCloudParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployFleetParams

`func (o *DeployVMsToCloudParams) GetDeployFleetParams() DeployFleetParams`

GetDeployFleetParams returns the DeployFleetParams field if non-nil, zero value otherwise.

### GetDeployFleetParamsOk

`func (o *DeployVMsToCloudParams) GetDeployFleetParamsOk() (*DeployFleetParams, bool)`

GetDeployFleetParamsOk returns a tuple with the DeployFleetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployFleetParams

`func (o *DeployVMsToCloudParams) SetDeployFleetParams(v DeployFleetParams)`

SetDeployFleetParams sets DeployFleetParams field to given value.

### HasDeployFleetParams

`func (o *DeployVMsToCloudParams) HasDeployFleetParams() bool`

HasDeployFleetParams returns a boolean if a field has been set.

### GetDeployVmsToAwsParams

`func (o *DeployVMsToCloudParams) GetDeployVmsToAwsParams() DeployVMsToAWSParams`

GetDeployVmsToAwsParams returns the DeployVmsToAwsParams field if non-nil, zero value otherwise.

### GetDeployVmsToAwsParamsOk

`func (o *DeployVMsToCloudParams) GetDeployVmsToAwsParamsOk() (*DeployVMsToAWSParams, bool)`

GetDeployVmsToAwsParamsOk returns a tuple with the DeployVmsToAwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployVmsToAwsParams

`func (o *DeployVMsToCloudParams) SetDeployVmsToAwsParams(v DeployVMsToAWSParams)`

SetDeployVmsToAwsParams sets DeployVmsToAwsParams field to given value.

### HasDeployVmsToAwsParams

`func (o *DeployVMsToCloudParams) HasDeployVmsToAwsParams() bool`

HasDeployVmsToAwsParams returns a boolean if a field has been set.

### GetDeployVmsToAzureParams

`func (o *DeployVMsToCloudParams) GetDeployVmsToAzureParams() DeployVMsToAzureParams`

GetDeployVmsToAzureParams returns the DeployVmsToAzureParams field if non-nil, zero value otherwise.

### GetDeployVmsToAzureParamsOk

`func (o *DeployVMsToCloudParams) GetDeployVmsToAzureParamsOk() (*DeployVMsToAzureParams, bool)`

GetDeployVmsToAzureParamsOk returns a tuple with the DeployVmsToAzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployVmsToAzureParams

`func (o *DeployVMsToCloudParams) SetDeployVmsToAzureParams(v DeployVMsToAzureParams)`

SetDeployVmsToAzureParams sets DeployVmsToAzureParams field to given value.

### HasDeployVmsToAzureParams

`func (o *DeployVMsToCloudParams) HasDeployVmsToAzureParams() bool`

HasDeployVmsToAzureParams returns a boolean if a field has been set.

### GetDeployVmsToGcpParams

`func (o *DeployVMsToCloudParams) GetDeployVmsToGcpParams() DeployVMsToGCPParams`

GetDeployVmsToGcpParams returns the DeployVmsToGcpParams field if non-nil, zero value otherwise.

### GetDeployVmsToGcpParamsOk

`func (o *DeployVMsToCloudParams) GetDeployVmsToGcpParamsOk() (*DeployVMsToGCPParams, bool)`

GetDeployVmsToGcpParamsOk returns a tuple with the DeployVmsToGcpParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployVmsToGcpParams

`func (o *DeployVMsToCloudParams) SetDeployVmsToGcpParams(v DeployVMsToGCPParams)`

SetDeployVmsToGcpParams sets DeployVmsToGcpParams field to given value.

### HasDeployVmsToGcpParams

`func (o *DeployVMsToCloudParams) HasDeployVmsToGcpParams() bool`

HasDeployVmsToGcpParams returns a boolean if a field has been set.

### GetReplicateSnapshotsToAwsParams

`func (o *DeployVMsToCloudParams) GetReplicateSnapshotsToAwsParams() ReplicateSnapshotsToAWSParams`

GetReplicateSnapshotsToAwsParams returns the ReplicateSnapshotsToAwsParams field if non-nil, zero value otherwise.

### GetReplicateSnapshotsToAwsParamsOk

`func (o *DeployVMsToCloudParams) GetReplicateSnapshotsToAwsParamsOk() (*ReplicateSnapshotsToAWSParams, bool)`

GetReplicateSnapshotsToAwsParamsOk returns a tuple with the ReplicateSnapshotsToAwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateSnapshotsToAwsParams

`func (o *DeployVMsToCloudParams) SetReplicateSnapshotsToAwsParams(v ReplicateSnapshotsToAWSParams)`

SetReplicateSnapshotsToAwsParams sets ReplicateSnapshotsToAwsParams field to given value.

### HasReplicateSnapshotsToAwsParams

`func (o *DeployVMsToCloudParams) HasReplicateSnapshotsToAwsParams() bool`

HasReplicateSnapshotsToAwsParams returns a boolean if a field has been set.

### GetReplicateSnapshotsToAzureParams

`func (o *DeployVMsToCloudParams) GetReplicateSnapshotsToAzureParams() ReplicateSnapshotsToAzureParams`

GetReplicateSnapshotsToAzureParams returns the ReplicateSnapshotsToAzureParams field if non-nil, zero value otherwise.

### GetReplicateSnapshotsToAzureParamsOk

`func (o *DeployVMsToCloudParams) GetReplicateSnapshotsToAzureParamsOk() (*ReplicateSnapshotsToAzureParams, bool)`

GetReplicateSnapshotsToAzureParamsOk returns a tuple with the ReplicateSnapshotsToAzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateSnapshotsToAzureParams

`func (o *DeployVMsToCloudParams) SetReplicateSnapshotsToAzureParams(v ReplicateSnapshotsToAzureParams)`

SetReplicateSnapshotsToAzureParams sets ReplicateSnapshotsToAzureParams field to given value.

### HasReplicateSnapshotsToAzureParams

`func (o *DeployVMsToCloudParams) HasReplicateSnapshotsToAzureParams() bool`

HasReplicateSnapshotsToAzureParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


