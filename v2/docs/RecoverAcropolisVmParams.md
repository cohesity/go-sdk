# RecoverAcropolisVmParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcropolisTargetParams** | Pointer to [**NullableRecoverAcropolisVmParamsAcropolisTargetParams**](RecoverAcropolisVmParamsAcropolisTargetParams.md) |  | [optional] 
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAcropolisVmParams

`func NewRecoverAcropolisVmParams(targetEnvironment string, ) *RecoverAcropolisVmParams`

NewRecoverAcropolisVmParams instantiates a new RecoverAcropolisVmParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAcropolisVmParamsWithDefaults

`func NewRecoverAcropolisVmParamsWithDefaults() *RecoverAcropolisVmParams`

NewRecoverAcropolisVmParamsWithDefaults instantiates a new RecoverAcropolisVmParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcropolisTargetParams

`func (o *RecoverAcropolisVmParams) GetAcropolisTargetParams() RecoverAcropolisVmParamsAcropolisTargetParams`

GetAcropolisTargetParams returns the AcropolisTargetParams field if non-nil, zero value otherwise.

### GetAcropolisTargetParamsOk

`func (o *RecoverAcropolisVmParams) GetAcropolisTargetParamsOk() (*RecoverAcropolisVmParamsAcropolisTargetParams, bool)`

GetAcropolisTargetParamsOk returns a tuple with the AcropolisTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisTargetParams

`func (o *RecoverAcropolisVmParams) SetAcropolisTargetParams(v RecoverAcropolisVmParamsAcropolisTargetParams)`

SetAcropolisTargetParams sets AcropolisTargetParams field to given value.

### HasAcropolisTargetParams

`func (o *RecoverAcropolisVmParams) HasAcropolisTargetParams() bool`

HasAcropolisTargetParams returns a boolean if a field has been set.

### SetAcropolisTargetParamsNil

`func (o *RecoverAcropolisVmParams) SetAcropolisTargetParamsNil(b bool)`

 SetAcropolisTargetParamsNil sets the value for AcropolisTargetParams to be an explicit nil

### UnsetAcropolisTargetParams
`func (o *RecoverAcropolisVmParams) UnsetAcropolisTargetParams()`

UnsetAcropolisTargetParams ensures that no value is present for AcropolisTargetParams, not even an explicit nil
### GetRecoverProtectionGroupRunsParams

`func (o *RecoverAcropolisVmParams) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams`

GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field if non-nil, zero value otherwise.

### GetRecoverProtectionGroupRunsParamsOk

`func (o *RecoverAcropolisVmParams) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool)`

GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverProtectionGroupRunsParams

`func (o *RecoverAcropolisVmParams) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams)`

SetRecoverProtectionGroupRunsParams sets RecoverProtectionGroupRunsParams field to given value.

### HasRecoverProtectionGroupRunsParams

`func (o *RecoverAcropolisVmParams) HasRecoverProtectionGroupRunsParams() bool`

HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.

### SetRecoverProtectionGroupRunsParamsNil

`func (o *RecoverAcropolisVmParams) SetRecoverProtectionGroupRunsParamsNil(b bool)`

 SetRecoverProtectionGroupRunsParamsNil sets the value for RecoverProtectionGroupRunsParams to be an explicit nil

### UnsetRecoverProtectionGroupRunsParams
`func (o *RecoverAcropolisVmParams) UnsetRecoverProtectionGroupRunsParams()`

UnsetRecoverProtectionGroupRunsParams ensures that no value is present for RecoverProtectionGroupRunsParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAcropolisVmParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAcropolisVmParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAcropolisVmParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


