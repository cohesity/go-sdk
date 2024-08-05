# CreateTdmTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **NullableString** | Specifies the TDM Task action. | 
**Name** | Pointer to **NullableString** | Specifies the name of the task. | [optional] 
**CloneParams** | Pointer to [**TdmCloneTaskRequestParams**](TdmCloneTaskRequestParams.md) |  | [optional] 
**RefreshParams** | Pointer to [**TdmRefreshTaskRequestParams**](TdmRefreshTaskRequestParams.md) |  | [optional] 
**SnapshotParams** | Pointer to [**TdmSnapshotTaskParams**](TdmSnapshotTaskParams.md) |  | [optional] 
**TeardownParams** | Pointer to [**TdmTeardownTaskRequestParams**](TdmTeardownTaskRequestParams.md) |  | [optional] 

## Methods

### NewCreateTdmTaskRequest

`func NewCreateTdmTaskRequest(action NullableString, ) *CreateTdmTaskRequest`

NewCreateTdmTaskRequest instantiates a new CreateTdmTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTdmTaskRequestWithDefaults

`func NewCreateTdmTaskRequestWithDefaults() *CreateTdmTaskRequest`

NewCreateTdmTaskRequestWithDefaults instantiates a new CreateTdmTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CreateTdmTaskRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateTdmTaskRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateTdmTaskRequest) SetAction(v string)`

SetAction sets Action field to given value.


### SetActionNil

`func (o *CreateTdmTaskRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CreateTdmTaskRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetName

`func (o *CreateTdmTaskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTdmTaskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTdmTaskRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTdmTaskRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTdmTaskRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTdmTaskRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCloneParams

`func (o *CreateTdmTaskRequest) GetCloneParams() TdmCloneTaskRequestParams`

GetCloneParams returns the CloneParams field if non-nil, zero value otherwise.

### GetCloneParamsOk

`func (o *CreateTdmTaskRequest) GetCloneParamsOk() (*TdmCloneTaskRequestParams, bool)`

GetCloneParamsOk returns a tuple with the CloneParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneParams

`func (o *CreateTdmTaskRequest) SetCloneParams(v TdmCloneTaskRequestParams)`

SetCloneParams sets CloneParams field to given value.

### HasCloneParams

`func (o *CreateTdmTaskRequest) HasCloneParams() bool`

HasCloneParams returns a boolean if a field has been set.

### GetRefreshParams

`func (o *CreateTdmTaskRequest) GetRefreshParams() TdmRefreshTaskRequestParams`

GetRefreshParams returns the RefreshParams field if non-nil, zero value otherwise.

### GetRefreshParamsOk

`func (o *CreateTdmTaskRequest) GetRefreshParamsOk() (*TdmRefreshTaskRequestParams, bool)`

GetRefreshParamsOk returns a tuple with the RefreshParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshParams

`func (o *CreateTdmTaskRequest) SetRefreshParams(v TdmRefreshTaskRequestParams)`

SetRefreshParams sets RefreshParams field to given value.

### HasRefreshParams

`func (o *CreateTdmTaskRequest) HasRefreshParams() bool`

HasRefreshParams returns a boolean if a field has been set.

### GetSnapshotParams

`func (o *CreateTdmTaskRequest) GetSnapshotParams() TdmSnapshotTaskParams`

GetSnapshotParams returns the SnapshotParams field if non-nil, zero value otherwise.

### GetSnapshotParamsOk

`func (o *CreateTdmTaskRequest) GetSnapshotParamsOk() (*TdmSnapshotTaskParams, bool)`

GetSnapshotParamsOk returns a tuple with the SnapshotParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotParams

`func (o *CreateTdmTaskRequest) SetSnapshotParams(v TdmSnapshotTaskParams)`

SetSnapshotParams sets SnapshotParams field to given value.

### HasSnapshotParams

`func (o *CreateTdmTaskRequest) HasSnapshotParams() bool`

HasSnapshotParams returns a boolean if a field has been set.

### GetTeardownParams

`func (o *CreateTdmTaskRequest) GetTeardownParams() TdmTeardownTaskRequestParams`

GetTeardownParams returns the TeardownParams field if non-nil, zero value otherwise.

### GetTeardownParamsOk

`func (o *CreateTdmTaskRequest) GetTeardownParamsOk() (*TdmTeardownTaskRequestParams, bool)`

GetTeardownParamsOk returns a tuple with the TeardownParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeardownParams

`func (o *CreateTdmTaskRequest) SetTeardownParams(v TdmTeardownTaskRequestParams)`

SetTeardownParams sets TeardownParams field to given value.

### HasTeardownParams

`func (o *CreateTdmTaskRequest) HasTeardownParams() bool`

HasTeardownParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


