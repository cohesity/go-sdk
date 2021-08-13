# TdmRefreshTaskResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment of the TDM refresh task. | 
**Parent** | Pointer to [**NullableObjectSummary**](ObjectSummary.md) | Specifies the details of the parent object of the clone. | [optional] 
**Target** | Pointer to [**NullableObjectSummary**](ObjectSummary.md) | Specifies the details of the target, where the clone is created. | [optional] 
**Snapshot** | Pointer to [**NullableObjectSnapshot**](ObjectSnapshot.md) | Specifies the details of the snapshot used for refresh. | [optional] 
**View** | Pointer to [**NullableViewParams**](ViewParams.md) | Specifies the details of the view, which is used for the clone. | [optional] 
**OracleParams** | Pointer to [**OracleCloneRefreshTask**](OracleCloneRefreshTask.md) |  | [optional] 

## Methods

### NewTdmRefreshTaskResponseParams

`func NewTdmRefreshTaskResponseParams(environment NullableString, ) *TdmRefreshTaskResponseParams`

NewTdmRefreshTaskResponseParams instantiates a new TdmRefreshTaskResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmRefreshTaskResponseParamsWithDefaults

`func NewTdmRefreshTaskResponseParamsWithDefaults() *TdmRefreshTaskResponseParams`

NewTdmRefreshTaskResponseParamsWithDefaults instantiates a new TdmRefreshTaskResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *TdmRefreshTaskResponseParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TdmRefreshTaskResponseParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TdmRefreshTaskResponseParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *TdmRefreshTaskResponseParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *TdmRefreshTaskResponseParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetParent

`func (o *TdmRefreshTaskResponseParams) GetParent() ObjectSummary`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TdmRefreshTaskResponseParams) GetParentOk() (*ObjectSummary, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TdmRefreshTaskResponseParams) SetParent(v ObjectSummary)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TdmRefreshTaskResponseParams) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *TdmRefreshTaskResponseParams) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *TdmRefreshTaskResponseParams) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetTarget

`func (o *TdmRefreshTaskResponseParams) GetTarget() ObjectSummary`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TdmRefreshTaskResponseParams) GetTargetOk() (*ObjectSummary, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TdmRefreshTaskResponseParams) SetTarget(v ObjectSummary)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TdmRefreshTaskResponseParams) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *TdmRefreshTaskResponseParams) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *TdmRefreshTaskResponseParams) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetSnapshot

`func (o *TdmRefreshTaskResponseParams) GetSnapshot() ObjectSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *TdmRefreshTaskResponseParams) GetSnapshotOk() (*ObjectSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *TdmRefreshTaskResponseParams) SetSnapshot(v ObjectSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *TdmRefreshTaskResponseParams) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### SetSnapshotNil

`func (o *TdmRefreshTaskResponseParams) SetSnapshotNil(b bool)`

 SetSnapshotNil sets the value for Snapshot to be an explicit nil

### UnsetSnapshot
`func (o *TdmRefreshTaskResponseParams) UnsetSnapshot()`

UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
### GetView

`func (o *TdmRefreshTaskResponseParams) GetView() ViewParams`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *TdmRefreshTaskResponseParams) GetViewOk() (*ViewParams, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *TdmRefreshTaskResponseParams) SetView(v ViewParams)`

SetView sets View field to given value.

### HasView

`func (o *TdmRefreshTaskResponseParams) HasView() bool`

HasView returns a boolean if a field has been set.

### SetViewNil

`func (o *TdmRefreshTaskResponseParams) SetViewNil(b bool)`

 SetViewNil sets the value for View to be an explicit nil

### UnsetView
`func (o *TdmRefreshTaskResponseParams) UnsetView()`

UnsetView ensures that no value is present for View, not even an explicit nil
### GetOracleParams

`func (o *TdmRefreshTaskResponseParams) GetOracleParams() OracleCloneRefreshTask`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *TdmRefreshTaskResponseParams) GetOracleParamsOk() (*OracleCloneRefreshTask, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *TdmRefreshTaskResponseParams) SetOracleParams(v OracleCloneRefreshTask)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *TdmRefreshTaskResponseParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


