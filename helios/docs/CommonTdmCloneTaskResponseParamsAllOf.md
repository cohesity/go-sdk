# CommonTdmCloneTaskResponseParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshot** | Pointer to [**NullableObjectSnapshot**](ObjectSnapshot.md) | Specifies the details of the snapshot used for cloning. | [optional] 
**Parent** | Pointer to [**NullableObjectSummary**](ObjectSummary.md) | Specifies the details of the parent object of the clone. | [optional] 
**Target** | Pointer to [**NullableObjectSummary**](ObjectSummary.md) | Specifies the details of the target, where the clone is created. | [optional] 
**View** | Pointer to [**NullableViewParams**](ViewParams.md) | Specifies the details of the view, which is used for the clone. | [optional] 

## Methods

### NewCommonTdmCloneTaskResponseParamsAllOf

`func NewCommonTdmCloneTaskResponseParamsAllOf() *CommonTdmCloneTaskResponseParamsAllOf`

NewCommonTdmCloneTaskResponseParamsAllOf instantiates a new CommonTdmCloneTaskResponseParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonTdmCloneTaskResponseParamsAllOfWithDefaults

`func NewCommonTdmCloneTaskResponseParamsAllOfWithDefaults() *CommonTdmCloneTaskResponseParamsAllOf`

NewCommonTdmCloneTaskResponseParamsAllOfWithDefaults instantiates a new CommonTdmCloneTaskResponseParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshot

`func (o *CommonTdmCloneTaskResponseParamsAllOf) GetSnapshot() ObjectSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CommonTdmCloneTaskResponseParamsAllOf) GetSnapshotOk() (*ObjectSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CommonTdmCloneTaskResponseParamsAllOf) SetSnapshot(v ObjectSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CommonTdmCloneTaskResponseParamsAllOf) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### SetSnapshotNil

`func (o *CommonTdmCloneTaskResponseParamsAllOf) SetSnapshotNil(b bool)`

 SetSnapshotNil sets the value for Snapshot to be an explicit nil

### UnsetSnapshot
`func (o *CommonTdmCloneTaskResponseParamsAllOf) UnsetSnapshot()`

UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
### GetParent

`func (o *CommonTdmCloneTaskResponseParamsAllOf) GetParent() ObjectSummary`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CommonTdmCloneTaskResponseParamsAllOf) GetParentOk() (*ObjectSummary, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CommonTdmCloneTaskResponseParamsAllOf) SetParent(v ObjectSummary)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CommonTdmCloneTaskResponseParamsAllOf) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *CommonTdmCloneTaskResponseParamsAllOf) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *CommonTdmCloneTaskResponseParamsAllOf) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetTarget

`func (o *CommonTdmCloneTaskResponseParamsAllOf) GetTarget() ObjectSummary`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CommonTdmCloneTaskResponseParamsAllOf) GetTargetOk() (*ObjectSummary, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CommonTdmCloneTaskResponseParamsAllOf) SetTarget(v ObjectSummary)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CommonTdmCloneTaskResponseParamsAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *CommonTdmCloneTaskResponseParamsAllOf) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *CommonTdmCloneTaskResponseParamsAllOf) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetView

`func (o *CommonTdmCloneTaskResponseParamsAllOf) GetView() ViewParams`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *CommonTdmCloneTaskResponseParamsAllOf) GetViewOk() (*ViewParams, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *CommonTdmCloneTaskResponseParamsAllOf) SetView(v ViewParams)`

SetView sets View field to given value.

### HasView

`func (o *CommonTdmCloneTaskResponseParamsAllOf) HasView() bool`

HasView returns a boolean if a field has been set.

### SetViewNil

`func (o *CommonTdmCloneTaskResponseParamsAllOf) SetViewNil(b bool)`

 SetViewNil sets the value for View to be an explicit nil

### UnsetView
`func (o *CommonTdmCloneTaskResponseParamsAllOf) UnsetView()`

UnsetView ensures that no value is present for View, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


