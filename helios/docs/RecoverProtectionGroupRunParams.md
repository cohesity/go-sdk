# RecoverProtectionGroupRunParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionGroupRunId** | **NullableString** | Specifies the Protection Group Run id from which to recover VMs. All the VM&#39;s that are successfully protected by this Run will be recovered. | 
**ProtectionGroupInstanceId** | **NullableInt64** | Specifies the Protection Group Instance id. | 
**ArchivalTargetId** | Pointer to **NullableInt64** | Specifies the archival target id. If specified and Protection Group run has an archival snapshot then VMs are recovered from the specified archival snapshot. If not specified (default), VMs are recovered from local snapshot. | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the local Protection Group id. In case of recovering a replication Run, this field should be provided with local Protection Group id. | [optional] 

## Methods

### NewRecoverProtectionGroupRunParams

`func NewRecoverProtectionGroupRunParams(protectionGroupRunId NullableString, protectionGroupInstanceId NullableInt64, ) *RecoverProtectionGroupRunParams`

NewRecoverProtectionGroupRunParams instantiates a new RecoverProtectionGroupRunParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverProtectionGroupRunParamsWithDefaults

`func NewRecoverProtectionGroupRunParamsWithDefaults() *RecoverProtectionGroupRunParams`

NewRecoverProtectionGroupRunParamsWithDefaults instantiates a new RecoverProtectionGroupRunParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionGroupRunId

`func (o *RecoverProtectionGroupRunParams) GetProtectionGroupRunId() string`

GetProtectionGroupRunId returns the ProtectionGroupRunId field if non-nil, zero value otherwise.

### GetProtectionGroupRunIdOk

`func (o *RecoverProtectionGroupRunParams) GetProtectionGroupRunIdOk() (*string, bool)`

GetProtectionGroupRunIdOk returns a tuple with the ProtectionGroupRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupRunId

`func (o *RecoverProtectionGroupRunParams) SetProtectionGroupRunId(v string)`

SetProtectionGroupRunId sets ProtectionGroupRunId field to given value.


### SetProtectionGroupRunIdNil

`func (o *RecoverProtectionGroupRunParams) SetProtectionGroupRunIdNil(b bool)`

 SetProtectionGroupRunIdNil sets the value for ProtectionGroupRunId to be an explicit nil

### UnsetProtectionGroupRunId
`func (o *RecoverProtectionGroupRunParams) UnsetProtectionGroupRunId()`

UnsetProtectionGroupRunId ensures that no value is present for ProtectionGroupRunId, not even an explicit nil
### GetProtectionGroupInstanceId

`func (o *RecoverProtectionGroupRunParams) GetProtectionGroupInstanceId() int64`

GetProtectionGroupInstanceId returns the ProtectionGroupInstanceId field if non-nil, zero value otherwise.

### GetProtectionGroupInstanceIdOk

`func (o *RecoverProtectionGroupRunParams) GetProtectionGroupInstanceIdOk() (*int64, bool)`

GetProtectionGroupInstanceIdOk returns a tuple with the ProtectionGroupInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupInstanceId

`func (o *RecoverProtectionGroupRunParams) SetProtectionGroupInstanceId(v int64)`

SetProtectionGroupInstanceId sets ProtectionGroupInstanceId field to given value.


### SetProtectionGroupInstanceIdNil

`func (o *RecoverProtectionGroupRunParams) SetProtectionGroupInstanceIdNil(b bool)`

 SetProtectionGroupInstanceIdNil sets the value for ProtectionGroupInstanceId to be an explicit nil

### UnsetProtectionGroupInstanceId
`func (o *RecoverProtectionGroupRunParams) UnsetProtectionGroupInstanceId()`

UnsetProtectionGroupInstanceId ensures that no value is present for ProtectionGroupInstanceId, not even an explicit nil
### GetArchivalTargetId

`func (o *RecoverProtectionGroupRunParams) GetArchivalTargetId() int64`

GetArchivalTargetId returns the ArchivalTargetId field if non-nil, zero value otherwise.

### GetArchivalTargetIdOk

`func (o *RecoverProtectionGroupRunParams) GetArchivalTargetIdOk() (*int64, bool)`

GetArchivalTargetIdOk returns a tuple with the ArchivalTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTargetId

`func (o *RecoverProtectionGroupRunParams) SetArchivalTargetId(v int64)`

SetArchivalTargetId sets ArchivalTargetId field to given value.

### HasArchivalTargetId

`func (o *RecoverProtectionGroupRunParams) HasArchivalTargetId() bool`

HasArchivalTargetId returns a boolean if a field has been set.

### SetArchivalTargetIdNil

`func (o *RecoverProtectionGroupRunParams) SetArchivalTargetIdNil(b bool)`

 SetArchivalTargetIdNil sets the value for ArchivalTargetId to be an explicit nil

### UnsetArchivalTargetId
`func (o *RecoverProtectionGroupRunParams) UnsetArchivalTargetId()`

UnsetArchivalTargetId ensures that no value is present for ArchivalTargetId, not even an explicit nil
### GetProtectionGroupId

`func (o *RecoverProtectionGroupRunParams) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *RecoverProtectionGroupRunParams) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *RecoverProtectionGroupRunParams) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *RecoverProtectionGroupRunParams) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *RecoverProtectionGroupRunParams) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *RecoverProtectionGroupRunParams) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


