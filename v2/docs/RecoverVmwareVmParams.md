# RecoverVmwareVmParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**RestoreObjectCustomizations** | Pointer to [**[]RestoreObjectCustomization**](RestoreObjectCustomization.md) | Specifies the customization for the VMs being restored. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VmwareTargetParams** | Pointer to [**NullableRecoverVmwareVmParamsVmwareTargetParams**](RecoverVmwareVmParamsVmwareTargetParams.md) |  | [optional] 

## Methods

### NewRecoverVmwareVmParams

`func NewRecoverVmwareVmParams(targetEnvironment string, ) *RecoverVmwareVmParams`

NewRecoverVmwareVmParams instantiates a new RecoverVmwareVmParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareVmParamsWithDefaults

`func NewRecoverVmwareVmParamsWithDefaults() *RecoverVmwareVmParams`

NewRecoverVmwareVmParamsWithDefaults instantiates a new RecoverVmwareVmParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverProtectionGroupRunsParams

`func (o *RecoverVmwareVmParams) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams`

GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field if non-nil, zero value otherwise.

### GetRecoverProtectionGroupRunsParamsOk

`func (o *RecoverVmwareVmParams) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool)`

GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverProtectionGroupRunsParams

`func (o *RecoverVmwareVmParams) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams)`

SetRecoverProtectionGroupRunsParams sets RecoverProtectionGroupRunsParams field to given value.

### HasRecoverProtectionGroupRunsParams

`func (o *RecoverVmwareVmParams) HasRecoverProtectionGroupRunsParams() bool`

HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.

### SetRecoverProtectionGroupRunsParamsNil

`func (o *RecoverVmwareVmParams) SetRecoverProtectionGroupRunsParamsNil(b bool)`

 SetRecoverProtectionGroupRunsParamsNil sets the value for RecoverProtectionGroupRunsParams to be an explicit nil

### UnsetRecoverProtectionGroupRunsParams
`func (o *RecoverVmwareVmParams) UnsetRecoverProtectionGroupRunsParams()`

UnsetRecoverProtectionGroupRunsParams ensures that no value is present for RecoverProtectionGroupRunsParams, not even an explicit nil
### GetRestoreObjectCustomizations

`func (o *RecoverVmwareVmParams) GetRestoreObjectCustomizations() []RestoreObjectCustomization`

GetRestoreObjectCustomizations returns the RestoreObjectCustomizations field if non-nil, zero value otherwise.

### GetRestoreObjectCustomizationsOk

`func (o *RecoverVmwareVmParams) GetRestoreObjectCustomizationsOk() (*[]RestoreObjectCustomization, bool)`

GetRestoreObjectCustomizationsOk returns a tuple with the RestoreObjectCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreObjectCustomizations

`func (o *RecoverVmwareVmParams) SetRestoreObjectCustomizations(v []RestoreObjectCustomization)`

SetRestoreObjectCustomizations sets RestoreObjectCustomizations field to given value.

### HasRestoreObjectCustomizations

`func (o *RecoverVmwareVmParams) HasRestoreObjectCustomizations() bool`

HasRestoreObjectCustomizations returns a boolean if a field has been set.

### SetRestoreObjectCustomizationsNil

`func (o *RecoverVmwareVmParams) SetRestoreObjectCustomizationsNil(b bool)`

 SetRestoreObjectCustomizationsNil sets the value for RestoreObjectCustomizations to be an explicit nil

### UnsetRestoreObjectCustomizations
`func (o *RecoverVmwareVmParams) UnsetRestoreObjectCustomizations()`

UnsetRestoreObjectCustomizations ensures that no value is present for RestoreObjectCustomizations, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverVmwareVmParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverVmwareVmParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverVmwareVmParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetVmwareTargetParams

`func (o *RecoverVmwareVmParams) GetVmwareTargetParams() RecoverVmwareVmParamsVmwareTargetParams`

GetVmwareTargetParams returns the VmwareTargetParams field if non-nil, zero value otherwise.

### GetVmwareTargetParamsOk

`func (o *RecoverVmwareVmParams) GetVmwareTargetParamsOk() (*RecoverVmwareVmParamsVmwareTargetParams, bool)`

GetVmwareTargetParamsOk returns a tuple with the VmwareTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareTargetParams

`func (o *RecoverVmwareVmParams) SetVmwareTargetParams(v RecoverVmwareVmParamsVmwareTargetParams)`

SetVmwareTargetParams sets VmwareTargetParams field to given value.

### HasVmwareTargetParams

`func (o *RecoverVmwareVmParams) HasVmwareTargetParams() bool`

HasVmwareTargetParams returns a boolean if a field has been set.

### SetVmwareTargetParamsNil

`func (o *RecoverVmwareVmParams) SetVmwareTargetParamsNil(b bool)`

 SetVmwareTargetParamsNil sets the value for VmwareTargetParams to be an explicit nil

### UnsetVmwareTargetParams
`func (o *RecoverVmwareVmParams) UnsetVmwareTargetParams()`

UnsetVmwareTargetParams ensures that no value is present for VmwareTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


