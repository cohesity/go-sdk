# RecoverGcpVmParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcpTargetParams** | Pointer to [**NullableRecoverGcpVmParamsGcpTargetParams**](RecoverGcpVmParamsGcpTargetParams.md) |  | [optional] 
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverGcpVmParams

`func NewRecoverGcpVmParams(targetEnvironment string, ) *RecoverGcpVmParams`

NewRecoverGcpVmParams instantiates a new RecoverGcpVmParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGcpVmParamsWithDefaults

`func NewRecoverGcpVmParamsWithDefaults() *RecoverGcpVmParams`

NewRecoverGcpVmParamsWithDefaults instantiates a new RecoverGcpVmParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcpTargetParams

`func (o *RecoverGcpVmParams) GetGcpTargetParams() RecoverGcpVmParamsGcpTargetParams`

GetGcpTargetParams returns the GcpTargetParams field if non-nil, zero value otherwise.

### GetGcpTargetParamsOk

`func (o *RecoverGcpVmParams) GetGcpTargetParamsOk() (*RecoverGcpVmParamsGcpTargetParams, bool)`

GetGcpTargetParamsOk returns a tuple with the GcpTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTargetParams

`func (o *RecoverGcpVmParams) SetGcpTargetParams(v RecoverGcpVmParamsGcpTargetParams)`

SetGcpTargetParams sets GcpTargetParams field to given value.

### HasGcpTargetParams

`func (o *RecoverGcpVmParams) HasGcpTargetParams() bool`

HasGcpTargetParams returns a boolean if a field has been set.

### SetGcpTargetParamsNil

`func (o *RecoverGcpVmParams) SetGcpTargetParamsNil(b bool)`

 SetGcpTargetParamsNil sets the value for GcpTargetParams to be an explicit nil

### UnsetGcpTargetParams
`func (o *RecoverGcpVmParams) UnsetGcpTargetParams()`

UnsetGcpTargetParams ensures that no value is present for GcpTargetParams, not even an explicit nil
### GetRecoverProtectionGroupRunsParams

`func (o *RecoverGcpVmParams) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams`

GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field if non-nil, zero value otherwise.

### GetRecoverProtectionGroupRunsParamsOk

`func (o *RecoverGcpVmParams) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool)`

GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverProtectionGroupRunsParams

`func (o *RecoverGcpVmParams) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams)`

SetRecoverProtectionGroupRunsParams sets RecoverProtectionGroupRunsParams field to given value.

### HasRecoverProtectionGroupRunsParams

`func (o *RecoverGcpVmParams) HasRecoverProtectionGroupRunsParams() bool`

HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.

### SetRecoverProtectionGroupRunsParamsNil

`func (o *RecoverGcpVmParams) SetRecoverProtectionGroupRunsParamsNil(b bool)`

 SetRecoverProtectionGroupRunsParamsNil sets the value for RecoverProtectionGroupRunsParams to be an explicit nil

### UnsetRecoverProtectionGroupRunsParams
`func (o *RecoverGcpVmParams) UnsetRecoverProtectionGroupRunsParams()`

UnsetRecoverProtectionGroupRunsParams ensures that no value is present for RecoverProtectionGroupRunsParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverGcpVmParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverGcpVmParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverGcpVmParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


