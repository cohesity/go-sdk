# ProtectdObjectsActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Specifies the action type to be performed on object getting protected. Based on selected action, provide the action params. | 
**ObjectActionKey** | Pointer to **NullableString** | Specifies the object action key for any action on the given object. | [optional] 
**PauseParams** | Pointer to [**ProtectedObjectPauseActionParams**](ProtectedObjectPauseActionParams.md) |  | [optional] 
**ResumeParams** | Pointer to [**ProtectedObjectResumeActionParams**](ProtectedObjectResumeActionParams.md) |  | [optional] 
**RunNowParams** | Pointer to [**ProtectedObjectRunNowActionParams**](ProtectedObjectRunNowActionParams.md) |  | [optional] 
**SnapshotBackendTypes** | Pointer to **[]string** | Specifies the protections type on which action to be performed. This is used when an object is protected by multiple protection types. If not specified action will be performed on all protection types. | [optional] 
**UnProtectParams** | Pointer to [**ProtectedObjectUnProtectActionParams**](ProtectedObjectUnProtectActionParams.md) |  | [optional] 

## Methods

### NewProtectdObjectsActionRequest

`func NewProtectdObjectsActionRequest(action string, ) *ProtectdObjectsActionRequest`

NewProtectdObjectsActionRequest instantiates a new ProtectdObjectsActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectdObjectsActionRequestWithDefaults

`func NewProtectdObjectsActionRequestWithDefaults() *ProtectdObjectsActionRequest`

NewProtectdObjectsActionRequestWithDefaults instantiates a new ProtectdObjectsActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ProtectdObjectsActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProtectdObjectsActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProtectdObjectsActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetObjectActionKey

`func (o *ProtectdObjectsActionRequest) GetObjectActionKey() string`

GetObjectActionKey returns the ObjectActionKey field if non-nil, zero value otherwise.

### GetObjectActionKeyOk

`func (o *ProtectdObjectsActionRequest) GetObjectActionKeyOk() (*string, bool)`

GetObjectActionKeyOk returns a tuple with the ObjectActionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectActionKey

`func (o *ProtectdObjectsActionRequest) SetObjectActionKey(v string)`

SetObjectActionKey sets ObjectActionKey field to given value.

### HasObjectActionKey

`func (o *ProtectdObjectsActionRequest) HasObjectActionKey() bool`

HasObjectActionKey returns a boolean if a field has been set.

### SetObjectActionKeyNil

`func (o *ProtectdObjectsActionRequest) SetObjectActionKeyNil(b bool)`

 SetObjectActionKeyNil sets the value for ObjectActionKey to be an explicit nil

### UnsetObjectActionKey
`func (o *ProtectdObjectsActionRequest) UnsetObjectActionKey()`

UnsetObjectActionKey ensures that no value is present for ObjectActionKey, not even an explicit nil
### GetPauseParams

`func (o *ProtectdObjectsActionRequest) GetPauseParams() ProtectedObjectPauseActionParams`

GetPauseParams returns the PauseParams field if non-nil, zero value otherwise.

### GetPauseParamsOk

`func (o *ProtectdObjectsActionRequest) GetPauseParamsOk() (*ProtectedObjectPauseActionParams, bool)`

GetPauseParamsOk returns a tuple with the PauseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseParams

`func (o *ProtectdObjectsActionRequest) SetPauseParams(v ProtectedObjectPauseActionParams)`

SetPauseParams sets PauseParams field to given value.

### HasPauseParams

`func (o *ProtectdObjectsActionRequest) HasPauseParams() bool`

HasPauseParams returns a boolean if a field has been set.

### GetResumeParams

`func (o *ProtectdObjectsActionRequest) GetResumeParams() ProtectedObjectResumeActionParams`

GetResumeParams returns the ResumeParams field if non-nil, zero value otherwise.

### GetResumeParamsOk

`func (o *ProtectdObjectsActionRequest) GetResumeParamsOk() (*ProtectedObjectResumeActionParams, bool)`

GetResumeParamsOk returns a tuple with the ResumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeParams

`func (o *ProtectdObjectsActionRequest) SetResumeParams(v ProtectedObjectResumeActionParams)`

SetResumeParams sets ResumeParams field to given value.

### HasResumeParams

`func (o *ProtectdObjectsActionRequest) HasResumeParams() bool`

HasResumeParams returns a boolean if a field has been set.

### GetRunNowParams

`func (o *ProtectdObjectsActionRequest) GetRunNowParams() ProtectedObjectRunNowActionParams`

GetRunNowParams returns the RunNowParams field if non-nil, zero value otherwise.

### GetRunNowParamsOk

`func (o *ProtectdObjectsActionRequest) GetRunNowParamsOk() (*ProtectedObjectRunNowActionParams, bool)`

GetRunNowParamsOk returns a tuple with the RunNowParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNowParams

`func (o *ProtectdObjectsActionRequest) SetRunNowParams(v ProtectedObjectRunNowActionParams)`

SetRunNowParams sets RunNowParams field to given value.

### HasRunNowParams

`func (o *ProtectdObjectsActionRequest) HasRunNowParams() bool`

HasRunNowParams returns a boolean if a field has been set.

### GetSnapshotBackendTypes

`func (o *ProtectdObjectsActionRequest) GetSnapshotBackendTypes() []string`

GetSnapshotBackendTypes returns the SnapshotBackendTypes field if non-nil, zero value otherwise.

### GetSnapshotBackendTypesOk

`func (o *ProtectdObjectsActionRequest) GetSnapshotBackendTypesOk() (*[]string, bool)`

GetSnapshotBackendTypesOk returns a tuple with the SnapshotBackendTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotBackendTypes

`func (o *ProtectdObjectsActionRequest) SetSnapshotBackendTypes(v []string)`

SetSnapshotBackendTypes sets SnapshotBackendTypes field to given value.

### HasSnapshotBackendTypes

`func (o *ProtectdObjectsActionRequest) HasSnapshotBackendTypes() bool`

HasSnapshotBackendTypes returns a boolean if a field has been set.

### SetSnapshotBackendTypesNil

`func (o *ProtectdObjectsActionRequest) SetSnapshotBackendTypesNil(b bool)`

 SetSnapshotBackendTypesNil sets the value for SnapshotBackendTypes to be an explicit nil

### UnsetSnapshotBackendTypes
`func (o *ProtectdObjectsActionRequest) UnsetSnapshotBackendTypes()`

UnsetSnapshotBackendTypes ensures that no value is present for SnapshotBackendTypes, not even an explicit nil
### GetUnProtectParams

`func (o *ProtectdObjectsActionRequest) GetUnProtectParams() ProtectedObjectUnProtectActionParams`

GetUnProtectParams returns the UnProtectParams field if non-nil, zero value otherwise.

### GetUnProtectParamsOk

`func (o *ProtectdObjectsActionRequest) GetUnProtectParamsOk() (*ProtectedObjectUnProtectActionParams, bool)`

GetUnProtectParamsOk returns a tuple with the UnProtectParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnProtectParams

`func (o *ProtectdObjectsActionRequest) SetUnProtectParams(v ProtectedObjectUnProtectActionParams)`

SetUnProtectParams sets UnProtectParams field to given value.

### HasUnProtectParams

`func (o *ProtectdObjectsActionRequest) HasUnProtectParams() bool`

HasUnProtectParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


