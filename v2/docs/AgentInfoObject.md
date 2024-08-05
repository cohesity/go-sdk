# AgentInfoObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the upgrade for an agent completed as a Unix epoch Timestamp (in microseconds). | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the source where the agent is installed. | [optional] 
**PreviousSoftwareVersion** | Pointer to **NullableString** | Specifies the software version of the agent before upgrade. | [optional] 
**StartTimeUsecs** | Pointer to **NullableInt64** | Specifies the time when the upgrade for an agent started as a Unix epoch Timestamp (in microseconds). | [optional] 
**Status** | Pointer to **NullableString** | Specifies the upgrade status of the agent.&lt;br&gt; &#39;Scheduled&#39; indicates that upgrade for the agent is yet to start.&lt;br&gt; &#39;Started&#39; indicates that upgrade for the agent is started.&lt;br&gt; &#39;Succeeded&#39; indicates that agent was upgraded successfully.&lt;br&gt; &#39;Failed&#39; indicates that upgrade for the agent has failed.&lt;br&gt; &#39;Skipped&#39; indicates that upgrade for the agent was skipped. | [optional] 

## Methods

### NewAgentInfoObject

`func NewAgentInfoObject() *AgentInfoObject`

NewAgentInfoObject instantiates a new AgentInfoObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentInfoObjectWithDefaults

`func NewAgentInfoObjectWithDefaults() *AgentInfoObject`

NewAgentInfoObjectWithDefaults instantiates a new AgentInfoObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeUsecs

`func (o *AgentInfoObject) GetEndTimeUsecs() int64`

GetEndTimeUsecs returns the EndTimeUsecs field if non-nil, zero value otherwise.

### GetEndTimeUsecsOk

`func (o *AgentInfoObject) GetEndTimeUsecsOk() (*int64, bool)`

GetEndTimeUsecsOk returns a tuple with the EndTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUsecs

`func (o *AgentInfoObject) SetEndTimeUsecs(v int64)`

SetEndTimeUsecs sets EndTimeUsecs field to given value.

### HasEndTimeUsecs

`func (o *AgentInfoObject) HasEndTimeUsecs() bool`

HasEndTimeUsecs returns a boolean if a field has been set.

### SetEndTimeUsecsNil

`func (o *AgentInfoObject) SetEndTimeUsecsNil(b bool)`

 SetEndTimeUsecsNil sets the value for EndTimeUsecs to be an explicit nil

### UnsetEndTimeUsecs
`func (o *AgentInfoObject) UnsetEndTimeUsecs()`

UnsetEndTimeUsecs ensures that no value is present for EndTimeUsecs, not even an explicit nil
### GetError

`func (o *AgentInfoObject) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AgentInfoObject) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AgentInfoObject) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *AgentInfoObject) HasError() bool`

HasError returns a boolean if a field has been set.

### GetName

`func (o *AgentInfoObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentInfoObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentInfoObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentInfoObject) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AgentInfoObject) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AgentInfoObject) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPreviousSoftwareVersion

`func (o *AgentInfoObject) GetPreviousSoftwareVersion() string`

GetPreviousSoftwareVersion returns the PreviousSoftwareVersion field if non-nil, zero value otherwise.

### GetPreviousSoftwareVersionOk

`func (o *AgentInfoObject) GetPreviousSoftwareVersionOk() (*string, bool)`

GetPreviousSoftwareVersionOk returns a tuple with the PreviousSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSoftwareVersion

`func (o *AgentInfoObject) SetPreviousSoftwareVersion(v string)`

SetPreviousSoftwareVersion sets PreviousSoftwareVersion field to given value.

### HasPreviousSoftwareVersion

`func (o *AgentInfoObject) HasPreviousSoftwareVersion() bool`

HasPreviousSoftwareVersion returns a boolean if a field has been set.

### SetPreviousSoftwareVersionNil

`func (o *AgentInfoObject) SetPreviousSoftwareVersionNil(b bool)`

 SetPreviousSoftwareVersionNil sets the value for PreviousSoftwareVersion to be an explicit nil

### UnsetPreviousSoftwareVersion
`func (o *AgentInfoObject) UnsetPreviousSoftwareVersion()`

UnsetPreviousSoftwareVersion ensures that no value is present for PreviousSoftwareVersion, not even an explicit nil
### GetStartTimeUsecs

`func (o *AgentInfoObject) GetStartTimeUsecs() int64`

GetStartTimeUsecs returns the StartTimeUsecs field if non-nil, zero value otherwise.

### GetStartTimeUsecsOk

`func (o *AgentInfoObject) GetStartTimeUsecsOk() (*int64, bool)`

GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUsecs

`func (o *AgentInfoObject) SetStartTimeUsecs(v int64)`

SetStartTimeUsecs sets StartTimeUsecs field to given value.

### HasStartTimeUsecs

`func (o *AgentInfoObject) HasStartTimeUsecs() bool`

HasStartTimeUsecs returns a boolean if a field has been set.

### SetStartTimeUsecsNil

`func (o *AgentInfoObject) SetStartTimeUsecsNil(b bool)`

 SetStartTimeUsecsNil sets the value for StartTimeUsecs to be an explicit nil

### UnsetStartTimeUsecs
`func (o *AgentInfoObject) UnsetStartTimeUsecs()`

UnsetStartTimeUsecs ensures that no value is present for StartTimeUsecs, not even an explicit nil
### GetStatus

`func (o *AgentInfoObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentInfoObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentInfoObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentInfoObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AgentInfoObject) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AgentInfoObject) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


