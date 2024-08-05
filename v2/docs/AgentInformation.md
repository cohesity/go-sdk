# AgentInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentSwVersion** | Pointer to **NullableString** | Specifies the software version of the agent | [optional] 
**ConnectionStatus** | Pointer to **NullableString** | Specifies the status of agent connection. | [optional] 
**HostSettingChecks** | Pointer to [**[]HostSettingCheck**](HostSettingCheck.md) | Specifies the list of host checks and its results. | [optional] 
**LastFetchedTimeInUsecs** | Pointer to **NullableInt64** | Specifies the time in usecs when the last agent info was fetched. | [optional] 
**SupportStatus** | Pointer to **NullableString** | Specifies the whether agent version is compatible with cluster version ro use various features. | [optional] 

## Methods

### NewAgentInformation

`func NewAgentInformation() *AgentInformation`

NewAgentInformation instantiates a new AgentInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentInformationWithDefaults

`func NewAgentInformationWithDefaults() *AgentInformation`

NewAgentInformationWithDefaults instantiates a new AgentInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentSwVersion

`func (o *AgentInformation) GetAgentSwVersion() string`

GetAgentSwVersion returns the AgentSwVersion field if non-nil, zero value otherwise.

### GetAgentSwVersionOk

`func (o *AgentInformation) GetAgentSwVersionOk() (*string, bool)`

GetAgentSwVersionOk returns a tuple with the AgentSwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSwVersion

`func (o *AgentInformation) SetAgentSwVersion(v string)`

SetAgentSwVersion sets AgentSwVersion field to given value.

### HasAgentSwVersion

`func (o *AgentInformation) HasAgentSwVersion() bool`

HasAgentSwVersion returns a boolean if a field has been set.

### SetAgentSwVersionNil

`func (o *AgentInformation) SetAgentSwVersionNil(b bool)`

 SetAgentSwVersionNil sets the value for AgentSwVersion to be an explicit nil

### UnsetAgentSwVersion
`func (o *AgentInformation) UnsetAgentSwVersion()`

UnsetAgentSwVersion ensures that no value is present for AgentSwVersion, not even an explicit nil
### GetConnectionStatus

`func (o *AgentInformation) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *AgentInformation) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *AgentInformation) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *AgentInformation) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### SetConnectionStatusNil

`func (o *AgentInformation) SetConnectionStatusNil(b bool)`

 SetConnectionStatusNil sets the value for ConnectionStatus to be an explicit nil

### UnsetConnectionStatus
`func (o *AgentInformation) UnsetConnectionStatus()`

UnsetConnectionStatus ensures that no value is present for ConnectionStatus, not even an explicit nil
### GetHostSettingChecks

`func (o *AgentInformation) GetHostSettingChecks() []HostSettingCheck`

GetHostSettingChecks returns the HostSettingChecks field if non-nil, zero value otherwise.

### GetHostSettingChecksOk

`func (o *AgentInformation) GetHostSettingChecksOk() (*[]HostSettingCheck, bool)`

GetHostSettingChecksOk returns a tuple with the HostSettingChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostSettingChecks

`func (o *AgentInformation) SetHostSettingChecks(v []HostSettingCheck)`

SetHostSettingChecks sets HostSettingChecks field to given value.

### HasHostSettingChecks

`func (o *AgentInformation) HasHostSettingChecks() bool`

HasHostSettingChecks returns a boolean if a field has been set.

### SetHostSettingChecksNil

`func (o *AgentInformation) SetHostSettingChecksNil(b bool)`

 SetHostSettingChecksNil sets the value for HostSettingChecks to be an explicit nil

### UnsetHostSettingChecks
`func (o *AgentInformation) UnsetHostSettingChecks()`

UnsetHostSettingChecks ensures that no value is present for HostSettingChecks, not even an explicit nil
### GetLastFetchedTimeInUsecs

`func (o *AgentInformation) GetLastFetchedTimeInUsecs() int64`

GetLastFetchedTimeInUsecs returns the LastFetchedTimeInUsecs field if non-nil, zero value otherwise.

### GetLastFetchedTimeInUsecsOk

`func (o *AgentInformation) GetLastFetchedTimeInUsecsOk() (*int64, bool)`

GetLastFetchedTimeInUsecsOk returns a tuple with the LastFetchedTimeInUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFetchedTimeInUsecs

`func (o *AgentInformation) SetLastFetchedTimeInUsecs(v int64)`

SetLastFetchedTimeInUsecs sets LastFetchedTimeInUsecs field to given value.

### HasLastFetchedTimeInUsecs

`func (o *AgentInformation) HasLastFetchedTimeInUsecs() bool`

HasLastFetchedTimeInUsecs returns a boolean if a field has been set.

### SetLastFetchedTimeInUsecsNil

`func (o *AgentInformation) SetLastFetchedTimeInUsecsNil(b bool)`

 SetLastFetchedTimeInUsecsNil sets the value for LastFetchedTimeInUsecs to be an explicit nil

### UnsetLastFetchedTimeInUsecs
`func (o *AgentInformation) UnsetLastFetchedTimeInUsecs()`

UnsetLastFetchedTimeInUsecs ensures that no value is present for LastFetchedTimeInUsecs, not even an explicit nil
### GetSupportStatus

`func (o *AgentInformation) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *AgentInformation) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *AgentInformation) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *AgentInformation) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.

### SetSupportStatusNil

`func (o *AgentInformation) SetSupportStatusNil(b bool)`

 SetSupportStatusNil sets the value for SupportStatus to be an explicit nil

### UnsetSupportStatus
`func (o *AgentInformation) UnsetSupportStatus()`

UnsetSupportStatus ensures that no value is present for SupportStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


