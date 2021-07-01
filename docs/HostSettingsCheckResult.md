# HostSettingsCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckType** | Pointer to **NullableString** | Specifies the type of the check internally performed. Specifies the type of the host check performed internally. &#39;kIsAgentPortAccessible&#39; indicates the check for agent port access. &#39;kIsAgentRunning&#39; indicates the status for the Cohesity agent service. &#39;kIsSQLWriterRunning&#39; indicates the status for SQLWriter service. &#39;kAreSQLInstancesRunning&#39; indicates the run status for all the SQL instances in the host. &#39;kCheckServiceLoginsConfig&#39; checks the privileges and sysadmin status of the logins used by the SQL instance services, Cohesity agent service and the SQLWriter service. &#39;kCheckSQLFCIVIP&#39; checks whether the SQL FCI is registered with a valid VIP or FQDN. &#39;kCheckSQLDiskSpace&#39; checks whether volumes containing SQL DBs have at least 10% free space. | [optional] 
**ResultType** | Pointer to **NullableString** | Specifies the type of the result returned after performing the internal host check. Specifies the type of the host check result performed internally. &#39;kPass&#39; indicates that the respective check was successful. &#39;kFail&#39; indicates that the respective check failed as some mandatory setting is not met &#39;kWarning&#39; indicates that the respective check has warning as certain non-mandatory setting is not met. | [optional] 
**UserMessage** | Pointer to **NullableString** | Specifies a descriptive message for failed/warning types. | [optional] 

## Methods

### NewHostSettingsCheckResult

`func NewHostSettingsCheckResult() *HostSettingsCheckResult`

NewHostSettingsCheckResult instantiates a new HostSettingsCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostSettingsCheckResultWithDefaults

`func NewHostSettingsCheckResultWithDefaults() *HostSettingsCheckResult`

NewHostSettingsCheckResultWithDefaults instantiates a new HostSettingsCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckType

`func (o *HostSettingsCheckResult) GetCheckType() string`

GetCheckType returns the CheckType field if non-nil, zero value otherwise.

### GetCheckTypeOk

`func (o *HostSettingsCheckResult) GetCheckTypeOk() (*string, bool)`

GetCheckTypeOk returns a tuple with the CheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckType

`func (o *HostSettingsCheckResult) SetCheckType(v string)`

SetCheckType sets CheckType field to given value.

### HasCheckType

`func (o *HostSettingsCheckResult) HasCheckType() bool`

HasCheckType returns a boolean if a field has been set.

### SetCheckTypeNil

`func (o *HostSettingsCheckResult) SetCheckTypeNil(b bool)`

 SetCheckTypeNil sets the value for CheckType to be an explicit nil

### UnsetCheckType
`func (o *HostSettingsCheckResult) UnsetCheckType()`

UnsetCheckType ensures that no value is present for CheckType, not even an explicit nil
### GetResultType

`func (o *HostSettingsCheckResult) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *HostSettingsCheckResult) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *HostSettingsCheckResult) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *HostSettingsCheckResult) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *HostSettingsCheckResult) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *HostSettingsCheckResult) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetUserMessage

`func (o *HostSettingsCheckResult) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *HostSettingsCheckResult) GetUserMessageOk() (*string, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *HostSettingsCheckResult) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### HasUserMessage

`func (o *HostSettingsCheckResult) HasUserMessage() bool`

HasUserMessage returns a boolean if a field has been set.

### SetUserMessageNil

`func (o *HostSettingsCheckResult) SetUserMessageNil(b bool)`

 SetUserMessageNil sets the value for UserMessage to be an explicit nil

### UnsetUserMessage
`func (o *HostSettingsCheckResult) UnsetUserMessage()`

UnsetUserMessage ensures that no value is present for UserMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


