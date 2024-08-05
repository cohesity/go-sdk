# RecoverKvmParamsRecoverVmParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KvmTargetParams** | Pointer to [**NullableRecoverKvmVmParamsKvmTargetParams**](RecoverKvmVmParamsKvmTargetParams.md) |  | [optional] 
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverKvmParamsRecoverVmParams

`func NewRecoverKvmParamsRecoverVmParams(targetEnvironment string, ) *RecoverKvmParamsRecoverVmParams`

NewRecoverKvmParamsRecoverVmParams instantiates a new RecoverKvmParamsRecoverVmParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKvmParamsRecoverVmParamsWithDefaults

`func NewRecoverKvmParamsRecoverVmParamsWithDefaults() *RecoverKvmParamsRecoverVmParams`

NewRecoverKvmParamsRecoverVmParamsWithDefaults instantiates a new RecoverKvmParamsRecoverVmParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKvmTargetParams

`func (o *RecoverKvmParamsRecoverVmParams) GetKvmTargetParams() RecoverKvmVmParamsKvmTargetParams`

GetKvmTargetParams returns the KvmTargetParams field if non-nil, zero value otherwise.

### GetKvmTargetParamsOk

`func (o *RecoverKvmParamsRecoverVmParams) GetKvmTargetParamsOk() (*RecoverKvmVmParamsKvmTargetParams, bool)`

GetKvmTargetParamsOk returns a tuple with the KvmTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmTargetParams

`func (o *RecoverKvmParamsRecoverVmParams) SetKvmTargetParams(v RecoverKvmVmParamsKvmTargetParams)`

SetKvmTargetParams sets KvmTargetParams field to given value.

### HasKvmTargetParams

`func (o *RecoverKvmParamsRecoverVmParams) HasKvmTargetParams() bool`

HasKvmTargetParams returns a boolean if a field has been set.

### SetKvmTargetParamsNil

`func (o *RecoverKvmParamsRecoverVmParams) SetKvmTargetParamsNil(b bool)`

 SetKvmTargetParamsNil sets the value for KvmTargetParams to be an explicit nil

### UnsetKvmTargetParams
`func (o *RecoverKvmParamsRecoverVmParams) UnsetKvmTargetParams()`

UnsetKvmTargetParams ensures that no value is present for KvmTargetParams, not even an explicit nil
### GetRecoverProtectionGroupRunsParams

`func (o *RecoverKvmParamsRecoverVmParams) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams`

GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field if non-nil, zero value otherwise.

### GetRecoverProtectionGroupRunsParamsOk

`func (o *RecoverKvmParamsRecoverVmParams) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool)`

GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverProtectionGroupRunsParams

`func (o *RecoverKvmParamsRecoverVmParams) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams)`

SetRecoverProtectionGroupRunsParams sets RecoverProtectionGroupRunsParams field to given value.

### HasRecoverProtectionGroupRunsParams

`func (o *RecoverKvmParamsRecoverVmParams) HasRecoverProtectionGroupRunsParams() bool`

HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.

### SetRecoverProtectionGroupRunsParamsNil

`func (o *RecoverKvmParamsRecoverVmParams) SetRecoverProtectionGroupRunsParamsNil(b bool)`

 SetRecoverProtectionGroupRunsParamsNil sets the value for RecoverProtectionGroupRunsParams to be an explicit nil

### UnsetRecoverProtectionGroupRunsParams
`func (o *RecoverKvmParamsRecoverVmParams) UnsetRecoverProtectionGroupRunsParams()`

UnsetRecoverProtectionGroupRunsParams ensures that no value is present for RecoverProtectionGroupRunsParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverKvmParamsRecoverVmParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverKvmParamsRecoverVmParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverKvmParamsRecoverVmParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


