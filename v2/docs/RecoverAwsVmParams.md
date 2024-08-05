# RecoverAwsVmParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsTargetParams** | Pointer to [**NullableRecoverAwsVmParamsAwsTargetParams**](RecoverAwsVmParamsAwsTargetParams.md) |  | [optional] 
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAwsVmParams

`func NewRecoverAwsVmParams(targetEnvironment string, ) *RecoverAwsVmParams`

NewRecoverAwsVmParams instantiates a new RecoverAwsVmParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsVmParamsWithDefaults

`func NewRecoverAwsVmParamsWithDefaults() *RecoverAwsVmParams`

NewRecoverAwsVmParamsWithDefaults instantiates a new RecoverAwsVmParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsTargetParams

`func (o *RecoverAwsVmParams) GetAwsTargetParams() RecoverAwsVmParamsAwsTargetParams`

GetAwsTargetParams returns the AwsTargetParams field if non-nil, zero value otherwise.

### GetAwsTargetParamsOk

`func (o *RecoverAwsVmParams) GetAwsTargetParamsOk() (*RecoverAwsVmParamsAwsTargetParams, bool)`

GetAwsTargetParamsOk returns a tuple with the AwsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetParams

`func (o *RecoverAwsVmParams) SetAwsTargetParams(v RecoverAwsVmParamsAwsTargetParams)`

SetAwsTargetParams sets AwsTargetParams field to given value.

### HasAwsTargetParams

`func (o *RecoverAwsVmParams) HasAwsTargetParams() bool`

HasAwsTargetParams returns a boolean if a field has been set.

### SetAwsTargetParamsNil

`func (o *RecoverAwsVmParams) SetAwsTargetParamsNil(b bool)`

 SetAwsTargetParamsNil sets the value for AwsTargetParams to be an explicit nil

### UnsetAwsTargetParams
`func (o *RecoverAwsVmParams) UnsetAwsTargetParams()`

UnsetAwsTargetParams ensures that no value is present for AwsTargetParams, not even an explicit nil
### GetRecoverProtectionGroupRunsParams

`func (o *RecoverAwsVmParams) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams`

GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field if non-nil, zero value otherwise.

### GetRecoverProtectionGroupRunsParamsOk

`func (o *RecoverAwsVmParams) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool)`

GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverProtectionGroupRunsParams

`func (o *RecoverAwsVmParams) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams)`

SetRecoverProtectionGroupRunsParams sets RecoverProtectionGroupRunsParams field to given value.

### HasRecoverProtectionGroupRunsParams

`func (o *RecoverAwsVmParams) HasRecoverProtectionGroupRunsParams() bool`

HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.

### SetRecoverProtectionGroupRunsParamsNil

`func (o *RecoverAwsVmParams) SetRecoverProtectionGroupRunsParamsNil(b bool)`

 SetRecoverProtectionGroupRunsParamsNil sets the value for RecoverProtectionGroupRunsParams to be an explicit nil

### UnsetRecoverProtectionGroupRunsParams
`func (o *RecoverAwsVmParams) UnsetRecoverProtectionGroupRunsParams()`

UnsetRecoverProtectionGroupRunsParams ensures that no value is present for RecoverProtectionGroupRunsParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAwsVmParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAwsVmParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAwsVmParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


