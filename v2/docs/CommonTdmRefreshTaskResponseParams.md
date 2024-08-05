# CommonTdmRefreshTaskResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment of the TDM refresh task. | 
**Parent** | Pointer to [**NullableObjectSummary**](ObjectSummary.md) | Specifies the details of the parent object of the clone. | [optional] 
**Snapshot** | Pointer to [**NullableObjectSnapshot**](ObjectSnapshot.md) | Specifies the details of the snapshot used for refresh. | [optional] 
**Target** | Pointer to [**NullableObjectSummary**](ObjectSummary.md) | Specifies the details of the target, where the clone is created. | [optional] 
**View** | Pointer to [**NullableViewParams**](ViewParams.md) | Specifies the details of the view, which is used for the clone. | [optional] 

## Methods

### NewCommonTdmRefreshTaskResponseParams

`func NewCommonTdmRefreshTaskResponseParams(environment NullableString, ) *CommonTdmRefreshTaskResponseParams`

NewCommonTdmRefreshTaskResponseParams instantiates a new CommonTdmRefreshTaskResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTdmRefreshTaskResponseParamsWithDefaults

`func NewCommonTdmRefreshTaskResponseParamsWithDefaults() *CommonTdmRefreshTaskResponseParams`

NewCommonTdmRefreshTaskResponseParamsWithDefaults instantiates a new CommonTdmRefreshTaskResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *CommonTdmRefreshTaskResponseParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonTdmRefreshTaskResponseParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonTdmRefreshTaskResponseParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *CommonTdmRefreshTaskResponseParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonTdmRefreshTaskResponseParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetParent

`func (o *CommonTdmRefreshTaskResponseParams) GetParent() ObjectSummary`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CommonTdmRefreshTaskResponseParams) GetParentOk() (*ObjectSummary, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CommonTdmRefreshTaskResponseParams) SetParent(v ObjectSummary)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CommonTdmRefreshTaskResponseParams) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *CommonTdmRefreshTaskResponseParams) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *CommonTdmRefreshTaskResponseParams) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetSnapshot

`func (o *CommonTdmRefreshTaskResponseParams) GetSnapshot() ObjectSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CommonTdmRefreshTaskResponseParams) GetSnapshotOk() (*ObjectSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CommonTdmRefreshTaskResponseParams) SetSnapshot(v ObjectSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CommonTdmRefreshTaskResponseParams) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### SetSnapshotNil

`func (o *CommonTdmRefreshTaskResponseParams) SetSnapshotNil(b bool)`

 SetSnapshotNil sets the value for Snapshot to be an explicit nil

### UnsetSnapshot
`func (o *CommonTdmRefreshTaskResponseParams) UnsetSnapshot()`

UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
### GetTarget

`func (o *CommonTdmRefreshTaskResponseParams) GetTarget() ObjectSummary`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CommonTdmRefreshTaskResponseParams) GetTargetOk() (*ObjectSummary, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CommonTdmRefreshTaskResponseParams) SetTarget(v ObjectSummary)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CommonTdmRefreshTaskResponseParams) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *CommonTdmRefreshTaskResponseParams) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *CommonTdmRefreshTaskResponseParams) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetView

`func (o *CommonTdmRefreshTaskResponseParams) GetView() ViewParams`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *CommonTdmRefreshTaskResponseParams) GetViewOk() (*ViewParams, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *CommonTdmRefreshTaskResponseParams) SetView(v ViewParams)`

SetView sets View field to given value.

### HasView

`func (o *CommonTdmRefreshTaskResponseParams) HasView() bool`

HasView returns a boolean if a field has been set.

### SetViewNil

`func (o *CommonTdmRefreshTaskResponseParams) SetViewNil(b bool)`

 SetViewNil sets the value for View to be an explicit nil

### UnsetView
`func (o *CommonTdmRefreshTaskResponseParams) UnsetView()`

UnsetView ensures that no value is present for View, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


