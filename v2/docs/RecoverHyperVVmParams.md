# RecoverHyperVVmParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervTargetParams** | Pointer to [**NullableRecoverHyperVVmParamsHypervTargetParams**](RecoverHyperVVmParamsHypervTargetParams.md) |  | [optional] 
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverHyperVVmParams

`func NewRecoverHyperVVmParams(targetEnvironment string, ) *RecoverHyperVVmParams`

NewRecoverHyperVVmParams instantiates a new RecoverHyperVVmParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverHyperVVmParamsWithDefaults

`func NewRecoverHyperVVmParamsWithDefaults() *RecoverHyperVVmParams`

NewRecoverHyperVVmParamsWithDefaults instantiates a new RecoverHyperVVmParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHypervTargetParams

`func (o *RecoverHyperVVmParams) GetHypervTargetParams() RecoverHyperVVmParamsHypervTargetParams`

GetHypervTargetParams returns the HypervTargetParams field if non-nil, zero value otherwise.

### GetHypervTargetParamsOk

`func (o *RecoverHyperVVmParams) GetHypervTargetParamsOk() (*RecoverHyperVVmParamsHypervTargetParams, bool)`

GetHypervTargetParamsOk returns a tuple with the HypervTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervTargetParams

`func (o *RecoverHyperVVmParams) SetHypervTargetParams(v RecoverHyperVVmParamsHypervTargetParams)`

SetHypervTargetParams sets HypervTargetParams field to given value.

### HasHypervTargetParams

`func (o *RecoverHyperVVmParams) HasHypervTargetParams() bool`

HasHypervTargetParams returns a boolean if a field has been set.

### SetHypervTargetParamsNil

`func (o *RecoverHyperVVmParams) SetHypervTargetParamsNil(b bool)`

 SetHypervTargetParamsNil sets the value for HypervTargetParams to be an explicit nil

### UnsetHypervTargetParams
`func (o *RecoverHyperVVmParams) UnsetHypervTargetParams()`

UnsetHypervTargetParams ensures that no value is present for HypervTargetParams, not even an explicit nil
### GetRecoverProtectionGroupRunsParams

`func (o *RecoverHyperVVmParams) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams`

GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field if non-nil, zero value otherwise.

### GetRecoverProtectionGroupRunsParamsOk

`func (o *RecoverHyperVVmParams) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool)`

GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverProtectionGroupRunsParams

`func (o *RecoverHyperVVmParams) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams)`

SetRecoverProtectionGroupRunsParams sets RecoverProtectionGroupRunsParams field to given value.

### HasRecoverProtectionGroupRunsParams

`func (o *RecoverHyperVVmParams) HasRecoverProtectionGroupRunsParams() bool`

HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.

### SetRecoverProtectionGroupRunsParamsNil

`func (o *RecoverHyperVVmParams) SetRecoverProtectionGroupRunsParamsNil(b bool)`

 SetRecoverProtectionGroupRunsParamsNil sets the value for RecoverProtectionGroupRunsParams to be an explicit nil

### UnsetRecoverProtectionGroupRunsParams
`func (o *RecoverHyperVVmParams) UnsetRecoverProtectionGroupRunsParams()`

UnsetRecoverProtectionGroupRunsParams ensures that no value is present for RecoverProtectionGroupRunsParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverHyperVVmParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverHyperVVmParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverHyperVVmParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


