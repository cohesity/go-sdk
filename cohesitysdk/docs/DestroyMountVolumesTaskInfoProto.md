# DestroyMountVolumesTaskInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**ErrorProto**](ErrorProto.md) |  | [optional] 
**Finished** | Pointer to **NullableBool** | This will be set to true if the task is complete on the slave. | [optional] 
**HostName** | Pointer to **NullableString** | This is the host name of the ESXi host. It is used if magneto_vmware_use_fqdn_for_guest_file_operations is set. | [optional] 
**MountVolumesInfoProto** | Pointer to [**MountVolumesInfoProto**](MountVolumesInfoProto.md) |  | [optional] 
**SlaveTaskStartTimeUsecs** | Pointer to **NullableInt64** | This is the timestamp at which the slave task started. | [optional] 
**TargetEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**UseExistingAgent** | Pointer to **NullableBool** | This will be set to true in two cases: 1. If persistent agent was used for IVM. 2. If user chose ephemeral agent during IVM but the host already had persistent agent installed. | [optional] 

## Methods

### NewDestroyMountVolumesTaskInfoProto

`func NewDestroyMountVolumesTaskInfoProto() *DestroyMountVolumesTaskInfoProto`

NewDestroyMountVolumesTaskInfoProto instantiates a new DestroyMountVolumesTaskInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestroyMountVolumesTaskInfoProtoWithDefaults

`func NewDestroyMountVolumesTaskInfoProtoWithDefaults() *DestroyMountVolumesTaskInfoProto`

NewDestroyMountVolumesTaskInfoProtoWithDefaults instantiates a new DestroyMountVolumesTaskInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *DestroyMountVolumesTaskInfoProto) GetError() ErrorProto`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DestroyMountVolumesTaskInfoProto) GetErrorOk() (*ErrorProto, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DestroyMountVolumesTaskInfoProto) SetError(v ErrorProto)`

SetError sets Error field to given value.

### HasError

`func (o *DestroyMountVolumesTaskInfoProto) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinished

`func (o *DestroyMountVolumesTaskInfoProto) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *DestroyMountVolumesTaskInfoProto) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *DestroyMountVolumesTaskInfoProto) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *DestroyMountVolumesTaskInfoProto) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *DestroyMountVolumesTaskInfoProto) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *DestroyMountVolumesTaskInfoProto) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetHostName

`func (o *DestroyMountVolumesTaskInfoProto) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *DestroyMountVolumesTaskInfoProto) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *DestroyMountVolumesTaskInfoProto) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *DestroyMountVolumesTaskInfoProto) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *DestroyMountVolumesTaskInfoProto) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *DestroyMountVolumesTaskInfoProto) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetMountVolumesInfoProto

`func (o *DestroyMountVolumesTaskInfoProto) GetMountVolumesInfoProto() MountVolumesInfoProto`

GetMountVolumesInfoProto returns the MountVolumesInfoProto field if non-nil, zero value otherwise.

### GetMountVolumesInfoProtoOk

`func (o *DestroyMountVolumesTaskInfoProto) GetMountVolumesInfoProtoOk() (*MountVolumesInfoProto, bool)`

GetMountVolumesInfoProtoOk returns a tuple with the MountVolumesInfoProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountVolumesInfoProto

`func (o *DestroyMountVolumesTaskInfoProto) SetMountVolumesInfoProto(v MountVolumesInfoProto)`

SetMountVolumesInfoProto sets MountVolumesInfoProto field to given value.

### HasMountVolumesInfoProto

`func (o *DestroyMountVolumesTaskInfoProto) HasMountVolumesInfoProto() bool`

HasMountVolumesInfoProto returns a boolean if a field has been set.

### GetSlaveTaskStartTimeUsecs

`func (o *DestroyMountVolumesTaskInfoProto) GetSlaveTaskStartTimeUsecs() int64`

GetSlaveTaskStartTimeUsecs returns the SlaveTaskStartTimeUsecs field if non-nil, zero value otherwise.

### GetSlaveTaskStartTimeUsecsOk

`func (o *DestroyMountVolumesTaskInfoProto) GetSlaveTaskStartTimeUsecsOk() (*int64, bool)`

GetSlaveTaskStartTimeUsecsOk returns a tuple with the SlaveTaskStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveTaskStartTimeUsecs

`func (o *DestroyMountVolumesTaskInfoProto) SetSlaveTaskStartTimeUsecs(v int64)`

SetSlaveTaskStartTimeUsecs sets SlaveTaskStartTimeUsecs field to given value.

### HasSlaveTaskStartTimeUsecs

`func (o *DestroyMountVolumesTaskInfoProto) HasSlaveTaskStartTimeUsecs() bool`

HasSlaveTaskStartTimeUsecs returns a boolean if a field has been set.

### SetSlaveTaskStartTimeUsecsNil

`func (o *DestroyMountVolumesTaskInfoProto) SetSlaveTaskStartTimeUsecsNil(b bool)`

 SetSlaveTaskStartTimeUsecsNil sets the value for SlaveTaskStartTimeUsecs to be an explicit nil

### UnsetSlaveTaskStartTimeUsecs
`func (o *DestroyMountVolumesTaskInfoProto) UnsetSlaveTaskStartTimeUsecs()`

UnsetSlaveTaskStartTimeUsecs ensures that no value is present for SlaveTaskStartTimeUsecs, not even an explicit nil
### GetTargetEntity

`func (o *DestroyMountVolumesTaskInfoProto) GetTargetEntity() EntityProto`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *DestroyMountVolumesTaskInfoProto) GetTargetEntityOk() (*EntityProto, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *DestroyMountVolumesTaskInfoProto) SetTargetEntity(v EntityProto)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *DestroyMountVolumesTaskInfoProto) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.

### GetUseExistingAgent

`func (o *DestroyMountVolumesTaskInfoProto) GetUseExistingAgent() bool`

GetUseExistingAgent returns the UseExistingAgent field if non-nil, zero value otherwise.

### GetUseExistingAgentOk

`func (o *DestroyMountVolumesTaskInfoProto) GetUseExistingAgentOk() (*bool, bool)`

GetUseExistingAgentOk returns a tuple with the UseExistingAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingAgent

`func (o *DestroyMountVolumesTaskInfoProto) SetUseExistingAgent(v bool)`

SetUseExistingAgent sets UseExistingAgent field to given value.

### HasUseExistingAgent

`func (o *DestroyMountVolumesTaskInfoProto) HasUseExistingAgent() bool`

HasUseExistingAgent returns a boolean if a field has been set.

### SetUseExistingAgentNil

`func (o *DestroyMountVolumesTaskInfoProto) SetUseExistingAgentNil(b bool)`

 SetUseExistingAgentNil sets the value for UseExistingAgent to be an explicit nil

### UnsetUseExistingAgent
`func (o *DestroyMountVolumesTaskInfoProto) UnsetUseExistingAgent()`

UnsetUseExistingAgent ensures that no value is present for UseExistingAgent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


