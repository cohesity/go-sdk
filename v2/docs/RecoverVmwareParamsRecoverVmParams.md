# RecoverVmwareParamsRecoverVmParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**RestoreObjectCustomizations** | Pointer to [**[]RestoreObjectCustomization**](RestoreObjectCustomization.md) | Specifies the customization for the VMs being restored. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VmwareTargetParams** | Pointer to [**NullableRecoverVmwareVmParamsVmwareTargetParams**](RecoverVmwareVmParamsVmwareTargetParams.md) |  | [optional] 

## Methods

### NewRecoverVmwareParamsRecoverVmParams

`func NewRecoverVmwareParamsRecoverVmParams(targetEnvironment string, ) *RecoverVmwareParamsRecoverVmParams`

NewRecoverVmwareParamsRecoverVmParams instantiates a new RecoverVmwareParamsRecoverVmParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVmwareParamsRecoverVmParamsWithDefaults

`func NewRecoverVmwareParamsRecoverVmParamsWithDefaults() *RecoverVmwareParamsRecoverVmParams`

NewRecoverVmwareParamsRecoverVmParamsWithDefaults instantiates a new RecoverVmwareParamsRecoverVmParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverProtectionGroupRunsParams

`func (o *RecoverVmwareParamsRecoverVmParams) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams`

GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field if non-nil, zero value otherwise.

### GetRecoverProtectionGroupRunsParamsOk

`func (o *RecoverVmwareParamsRecoverVmParams) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool)`

GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverProtectionGroupRunsParams

`func (o *RecoverVmwareParamsRecoverVmParams) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams)`

SetRecoverProtectionGroupRunsParams sets RecoverProtectionGroupRunsParams field to given value.

### HasRecoverProtectionGroupRunsParams

`func (o *RecoverVmwareParamsRecoverVmParams) HasRecoverProtectionGroupRunsParams() bool`

HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.

### SetRecoverProtectionGroupRunsParamsNil

`func (o *RecoverVmwareParamsRecoverVmParams) SetRecoverProtectionGroupRunsParamsNil(b bool)`

 SetRecoverProtectionGroupRunsParamsNil sets the value for RecoverProtectionGroupRunsParams to be an explicit nil

### UnsetRecoverProtectionGroupRunsParams
`func (o *RecoverVmwareParamsRecoverVmParams) UnsetRecoverProtectionGroupRunsParams()`

UnsetRecoverProtectionGroupRunsParams ensures that no value is present for RecoverProtectionGroupRunsParams, not even an explicit nil
### GetRestoreObjectCustomizations

`func (o *RecoverVmwareParamsRecoverVmParams) GetRestoreObjectCustomizations() []RestoreObjectCustomization`

GetRestoreObjectCustomizations returns the RestoreObjectCustomizations field if non-nil, zero value otherwise.

### GetRestoreObjectCustomizationsOk

`func (o *RecoverVmwareParamsRecoverVmParams) GetRestoreObjectCustomizationsOk() (*[]RestoreObjectCustomization, bool)`

GetRestoreObjectCustomizationsOk returns a tuple with the RestoreObjectCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreObjectCustomizations

`func (o *RecoverVmwareParamsRecoverVmParams) SetRestoreObjectCustomizations(v []RestoreObjectCustomization)`

SetRestoreObjectCustomizations sets RestoreObjectCustomizations field to given value.

### HasRestoreObjectCustomizations

`func (o *RecoverVmwareParamsRecoverVmParams) HasRestoreObjectCustomizations() bool`

HasRestoreObjectCustomizations returns a boolean if a field has been set.

### SetRestoreObjectCustomizationsNil

`func (o *RecoverVmwareParamsRecoverVmParams) SetRestoreObjectCustomizationsNil(b bool)`

 SetRestoreObjectCustomizationsNil sets the value for RestoreObjectCustomizations to be an explicit nil

### UnsetRestoreObjectCustomizations
`func (o *RecoverVmwareParamsRecoverVmParams) UnsetRestoreObjectCustomizations()`

UnsetRestoreObjectCustomizations ensures that no value is present for RestoreObjectCustomizations, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverVmwareParamsRecoverVmParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverVmwareParamsRecoverVmParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverVmwareParamsRecoverVmParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetVmwareTargetParams

`func (o *RecoverVmwareParamsRecoverVmParams) GetVmwareTargetParams() RecoverVmwareVmParamsVmwareTargetParams`

GetVmwareTargetParams returns the VmwareTargetParams field if non-nil, zero value otherwise.

### GetVmwareTargetParamsOk

`func (o *RecoverVmwareParamsRecoverVmParams) GetVmwareTargetParamsOk() (*RecoverVmwareVmParamsVmwareTargetParams, bool)`

GetVmwareTargetParamsOk returns a tuple with the VmwareTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareTargetParams

`func (o *RecoverVmwareParamsRecoverVmParams) SetVmwareTargetParams(v RecoverVmwareVmParamsVmwareTargetParams)`

SetVmwareTargetParams sets VmwareTargetParams field to given value.

### HasVmwareTargetParams

`func (o *RecoverVmwareParamsRecoverVmParams) HasVmwareTargetParams() bool`

HasVmwareTargetParams returns a boolean if a field has been set.

### SetVmwareTargetParamsNil

`func (o *RecoverVmwareParamsRecoverVmParams) SetVmwareTargetParamsNil(b bool)`

 SetVmwareTargetParamsNil sets the value for VmwareTargetParams to be an explicit nil

### UnsetVmwareTargetParams
`func (o *RecoverVmwareParamsRecoverVmParams) UnsetVmwareTargetParams()`

UnsetVmwareTargetParams ensures that no value is present for VmwareTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


