# RecoverExperimentalAdapterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overwrite** | Pointer to **NullableBool** | Set to true to overwrite an existing object at the destination. If set to false, and the same object exists at the destination, then recovery will fail for that object. | [optional] 
**RecoverTo** | Pointer to **NullableInt64** | Specifies the &#39;Source Registration ID&#39; of the source where the objects are to be recovered. If this is not specified, the recovery job will recover to the original location. | [optional] 
**RestoreType** | Pointer to **NullableString** | Specifies the type of experimental adapter restore. | [optional] 
**Snapshots** | [**[]RecoverExperimentalAdapterSnapshotParams**](RecoverExperimentalAdapterSnapshotParams.md) | Specifies the local snapshot ids and other details of the objects to be recovered. | 
**WorkflowParams** | Pointer to **NullableString** | Workflow parameters to be supplied to all backup workflow tasks. This specifies task configuration such as the time to wait before returning the workflow result, subtasks configuration such as number of subtasks to generate and the depth of the subtask tree, etc. | [optional] 

## Methods

### NewRecoverExperimentalAdapterParams

`func NewRecoverExperimentalAdapterParams(snapshots []RecoverExperimentalAdapterSnapshotParams, ) *RecoverExperimentalAdapterParams`

NewRecoverExperimentalAdapterParams instantiates a new RecoverExperimentalAdapterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverExperimentalAdapterParamsWithDefaults

`func NewRecoverExperimentalAdapterParamsWithDefaults() *RecoverExperimentalAdapterParams`

NewRecoverExperimentalAdapterParamsWithDefaults instantiates a new RecoverExperimentalAdapterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverwrite

`func (o *RecoverExperimentalAdapterParams) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RecoverExperimentalAdapterParams) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RecoverExperimentalAdapterParams) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RecoverExperimentalAdapterParams) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *RecoverExperimentalAdapterParams) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *RecoverExperimentalAdapterParams) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetRecoverTo

`func (o *RecoverExperimentalAdapterParams) GetRecoverTo() int64`

GetRecoverTo returns the RecoverTo field if non-nil, zero value otherwise.

### GetRecoverToOk

`func (o *RecoverExperimentalAdapterParams) GetRecoverToOk() (*int64, bool)`

GetRecoverToOk returns a tuple with the RecoverTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTo

`func (o *RecoverExperimentalAdapterParams) SetRecoverTo(v int64)`

SetRecoverTo sets RecoverTo field to given value.

### HasRecoverTo

`func (o *RecoverExperimentalAdapterParams) HasRecoverTo() bool`

HasRecoverTo returns a boolean if a field has been set.

### SetRecoverToNil

`func (o *RecoverExperimentalAdapterParams) SetRecoverToNil(b bool)`

 SetRecoverToNil sets the value for RecoverTo to be an explicit nil

### UnsetRecoverTo
`func (o *RecoverExperimentalAdapterParams) UnsetRecoverTo()`

UnsetRecoverTo ensures that no value is present for RecoverTo, not even an explicit nil
### GetRestoreType

`func (o *RecoverExperimentalAdapterParams) GetRestoreType() string`

GetRestoreType returns the RestoreType field if non-nil, zero value otherwise.

### GetRestoreTypeOk

`func (o *RecoverExperimentalAdapterParams) GetRestoreTypeOk() (*string, bool)`

GetRestoreTypeOk returns a tuple with the RestoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreType

`func (o *RecoverExperimentalAdapterParams) SetRestoreType(v string)`

SetRestoreType sets RestoreType field to given value.

### HasRestoreType

`func (o *RecoverExperimentalAdapterParams) HasRestoreType() bool`

HasRestoreType returns a boolean if a field has been set.

### SetRestoreTypeNil

`func (o *RecoverExperimentalAdapterParams) SetRestoreTypeNil(b bool)`

 SetRestoreTypeNil sets the value for RestoreType to be an explicit nil

### UnsetRestoreType
`func (o *RecoverExperimentalAdapterParams) UnsetRestoreType()`

UnsetRestoreType ensures that no value is present for RestoreType, not even an explicit nil
### GetSnapshots

`func (o *RecoverExperimentalAdapterParams) GetSnapshots() []RecoverExperimentalAdapterSnapshotParams`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *RecoverExperimentalAdapterParams) GetSnapshotsOk() (*[]RecoverExperimentalAdapterSnapshotParams, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *RecoverExperimentalAdapterParams) SetSnapshots(v []RecoverExperimentalAdapterSnapshotParams)`

SetSnapshots sets Snapshots field to given value.


### SetSnapshotsNil

`func (o *RecoverExperimentalAdapterParams) SetSnapshotsNil(b bool)`

 SetSnapshotsNil sets the value for Snapshots to be an explicit nil

### UnsetSnapshots
`func (o *RecoverExperimentalAdapterParams) UnsetSnapshots()`

UnsetSnapshots ensures that no value is present for Snapshots, not even an explicit nil
### GetWorkflowParams

`func (o *RecoverExperimentalAdapterParams) GetWorkflowParams() string`

GetWorkflowParams returns the WorkflowParams field if non-nil, zero value otherwise.

### GetWorkflowParamsOk

`func (o *RecoverExperimentalAdapterParams) GetWorkflowParamsOk() (*string, bool)`

GetWorkflowParamsOk returns a tuple with the WorkflowParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowParams

`func (o *RecoverExperimentalAdapterParams) SetWorkflowParams(v string)`

SetWorkflowParams sets WorkflowParams field to given value.

### HasWorkflowParams

`func (o *RecoverExperimentalAdapterParams) HasWorkflowParams() bool`

HasWorkflowParams returns a boolean if a field has been set.

### SetWorkflowParamsNil

`func (o *RecoverExperimentalAdapterParams) SetWorkflowParamsNil(b bool)`

 SetWorkflowParamsNil sets the value for WorkflowParams to be an explicit nil

### UnsetWorkflowParams
`func (o *RecoverExperimentalAdapterParams) UnsetWorkflowParams()`

UnsetWorkflowParams ensures that no value is present for WorkflowParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


