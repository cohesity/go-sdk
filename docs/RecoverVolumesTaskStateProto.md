# RecoverVolumesTaskStateProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | Pointer to [**RecoverVolumesParams**](RecoverVolumesParams.md) |  | [optional] 
**TaskResultVec** | Pointer to [**[]RecoverVolumesTaskStateProtoTaskResult**](RecoverVolumesTaskStateProtoTaskResult.md) | Contains high-level per-volume information. This data is here because Iris cannot see into protobuf extensions yet needs to display per-subtask progress. | [optional] 

## Methods

### NewRecoverVolumesTaskStateProto

`func NewRecoverVolumesTaskStateProto() *RecoverVolumesTaskStateProto`

NewRecoverVolumesTaskStateProto instantiates a new RecoverVolumesTaskStateProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVolumesTaskStateProtoWithDefaults

`func NewRecoverVolumesTaskStateProtoWithDefaults() *RecoverVolumesTaskStateProto`

NewRecoverVolumesTaskStateProtoWithDefaults instantiates a new RecoverVolumesTaskStateProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *RecoverVolumesTaskStateProto) GetParams() RecoverVolumesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *RecoverVolumesTaskStateProto) GetParamsOk() (*RecoverVolumesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *RecoverVolumesTaskStateProto) SetParams(v RecoverVolumesParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *RecoverVolumesTaskStateProto) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetTaskResultVec

`func (o *RecoverVolumesTaskStateProto) GetTaskResultVec() []RecoverVolumesTaskStateProtoTaskResult`

GetTaskResultVec returns the TaskResultVec field if non-nil, zero value otherwise.

### GetTaskResultVecOk

`func (o *RecoverVolumesTaskStateProto) GetTaskResultVecOk() (*[]RecoverVolumesTaskStateProtoTaskResult, bool)`

GetTaskResultVecOk returns a tuple with the TaskResultVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskResultVec

`func (o *RecoverVolumesTaskStateProto) SetTaskResultVec(v []RecoverVolumesTaskStateProtoTaskResult)`

SetTaskResultVec sets TaskResultVec field to given value.

### HasTaskResultVec

`func (o *RecoverVolumesTaskStateProto) HasTaskResultVec() bool`

HasTaskResultVec returns a boolean if a field has been set.

### SetTaskResultVecNil

`func (o *RecoverVolumesTaskStateProto) SetTaskResultVecNil(b bool)`

 SetTaskResultVecNil sets the value for TaskResultVec to be an explicit nil

### UnsetTaskResultVec
`func (o *RecoverVolumesTaskStateProto) UnsetTaskResultVec()`

UnsetTaskResultVec ensures that no value is present for TaskResultVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


