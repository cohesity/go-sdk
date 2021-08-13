# TdmCloneTaskResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment of the TDM Clone task. | 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the ID of an existing protection group, which should start protecting this clone. Specifying this implies that the clone is eligible for automated snapshots based on the policy configuration. If this is specified, policyId should also be specified and protectionGroupName should not be specified. | [optional] 
**ProtectionGroupName** | Pointer to **NullableString** | Specifies the name of a new protection group, which should be created to protect this clone. Specifying this implies that the clone is eligible for automated snapshots based on the policy configuration. If this is specified, policyId should also be specified and protectionGroupId should not be specified. | [optional] 
**PolicyId** | Pointer to **NullableString** | Specifies the ID of the policy, which should be used to protect this clone. This is useful for automatic snapshots. This must be specified if either of protectionGroupId and protectionGroupName is specified. | [optional] 
**Snapshot** | Pointer to [**NullableObjectSnapshot**](ObjectSnapshot.md) | Specifies the details of the snapshot used for cloning. | [optional] 
**Parent** | Pointer to [**NullableObjectSummary**](ObjectSummary.md) | Specifies the details of the parent object of the clone. | [optional] 
**Target** | Pointer to [**NullableObjectSummary**](ObjectSummary.md) | Specifies the details of the target, where the clone is created. | [optional] 
**View** | Pointer to [**NullableViewParams**](ViewParams.md) | Specifies the details of the view, which is used for the clone. | [optional] 
**OracleParams** | Pointer to [**OracleCloneTask**](OracleCloneTask.md) |  | [optional] 

## Methods

### NewTdmCloneTaskResponseParams

`func NewTdmCloneTaskResponseParams(environment NullableString, ) *TdmCloneTaskResponseParams`

NewTdmCloneTaskResponseParams instantiates a new TdmCloneTaskResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmCloneTaskResponseParamsWithDefaults

`func NewTdmCloneTaskResponseParamsWithDefaults() *TdmCloneTaskResponseParams`

NewTdmCloneTaskResponseParamsWithDefaults instantiates a new TdmCloneTaskResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *TdmCloneTaskResponseParams) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TdmCloneTaskResponseParams) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TdmCloneTaskResponseParams) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *TdmCloneTaskResponseParams) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *TdmCloneTaskResponseParams) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetProtectionGroupId

`func (o *TdmCloneTaskResponseParams) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *TdmCloneTaskResponseParams) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *TdmCloneTaskResponseParams) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *TdmCloneTaskResponseParams) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *TdmCloneTaskResponseParams) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *TdmCloneTaskResponseParams) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil
### GetProtectionGroupName

`func (o *TdmCloneTaskResponseParams) GetProtectionGroupName() string`

GetProtectionGroupName returns the ProtectionGroupName field if non-nil, zero value otherwise.

### GetProtectionGroupNameOk

`func (o *TdmCloneTaskResponseParams) GetProtectionGroupNameOk() (*string, bool)`

GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupName

`func (o *TdmCloneTaskResponseParams) SetProtectionGroupName(v string)`

SetProtectionGroupName sets ProtectionGroupName field to given value.

### HasProtectionGroupName

`func (o *TdmCloneTaskResponseParams) HasProtectionGroupName() bool`

HasProtectionGroupName returns a boolean if a field has been set.

### SetProtectionGroupNameNil

`func (o *TdmCloneTaskResponseParams) SetProtectionGroupNameNil(b bool)`

 SetProtectionGroupNameNil sets the value for ProtectionGroupName to be an explicit nil

### UnsetProtectionGroupName
`func (o *TdmCloneTaskResponseParams) UnsetProtectionGroupName()`

UnsetProtectionGroupName ensures that no value is present for ProtectionGroupName, not even an explicit nil
### GetPolicyId

`func (o *TdmCloneTaskResponseParams) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *TdmCloneTaskResponseParams) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *TdmCloneTaskResponseParams) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *TdmCloneTaskResponseParams) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *TdmCloneTaskResponseParams) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *TdmCloneTaskResponseParams) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetSnapshot

`func (o *TdmCloneTaskResponseParams) GetSnapshot() ObjectSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *TdmCloneTaskResponseParams) GetSnapshotOk() (*ObjectSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *TdmCloneTaskResponseParams) SetSnapshot(v ObjectSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *TdmCloneTaskResponseParams) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### SetSnapshotNil

`func (o *TdmCloneTaskResponseParams) SetSnapshotNil(b bool)`

 SetSnapshotNil sets the value for Snapshot to be an explicit nil

### UnsetSnapshot
`func (o *TdmCloneTaskResponseParams) UnsetSnapshot()`

UnsetSnapshot ensures that no value is present for Snapshot, not even an explicit nil
### GetParent

`func (o *TdmCloneTaskResponseParams) GetParent() ObjectSummary`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TdmCloneTaskResponseParams) GetParentOk() (*ObjectSummary, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TdmCloneTaskResponseParams) SetParent(v ObjectSummary)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TdmCloneTaskResponseParams) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *TdmCloneTaskResponseParams) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *TdmCloneTaskResponseParams) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetTarget

`func (o *TdmCloneTaskResponseParams) GetTarget() ObjectSummary`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TdmCloneTaskResponseParams) GetTargetOk() (*ObjectSummary, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TdmCloneTaskResponseParams) SetTarget(v ObjectSummary)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TdmCloneTaskResponseParams) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *TdmCloneTaskResponseParams) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *TdmCloneTaskResponseParams) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetView

`func (o *TdmCloneTaskResponseParams) GetView() ViewParams`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *TdmCloneTaskResponseParams) GetViewOk() (*ViewParams, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *TdmCloneTaskResponseParams) SetView(v ViewParams)`

SetView sets View field to given value.

### HasView

`func (o *TdmCloneTaskResponseParams) HasView() bool`

HasView returns a boolean if a field has been set.

### SetViewNil

`func (o *TdmCloneTaskResponseParams) SetViewNil(b bool)`

 SetViewNil sets the value for View to be an explicit nil

### UnsetView
`func (o *TdmCloneTaskResponseParams) UnsetView()`

UnsetView ensures that no value is present for View, not even an explicit nil
### GetOracleParams

`func (o *TdmCloneTaskResponseParams) GetOracleParams() OracleCloneTask`

GetOracleParams returns the OracleParams field if non-nil, zero value otherwise.

### GetOracleParamsOk

`func (o *TdmCloneTaskResponseParams) GetOracleParamsOk() (*OracleCloneTask, bool)`

GetOracleParamsOk returns a tuple with the OracleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleParams

`func (o *TdmCloneTaskResponseParams) SetOracleParams(v OracleCloneTask)`

SetOracleParams sets OracleParams field to given value.

### HasOracleParams

`func (o *TdmCloneTaskResponseParams) HasOracleParams() bool`

HasOracleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


