# AgentDeploymentStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompactVersion** | Pointer to **NullableString** | Specifies the compact version of Cohesity agent. For example, 6.0.1. | [optional] 
**HealthStatus** | Pointer to **NullableString** | Specifies the health status of the Cohesity agent. Specifies the status of the agent running on a physical source. &#39;kUnknown&#39; indicates the Agent is not known. No attempt to connect to the Agent has occurred. &#39;kUnreachable&#39; indicates the Agent is not reachable. &#39;kHealthy&#39; indicates the Agent is healthy. &#39;kDegraded&#39; indicates the Agent is running but in a degraded state. | [optional] 
**HostIp** | Pointer to **NullableString** | Specifies the IP of the host on which the agent is installed. | [optional] 
**HostOsType** | Pointer to **NullableString** | Specifies the host type on which the agent is installed. &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**LastUpgradeStatus** | Pointer to **NullableString** | Specifies the status of the last upgrade attempt. Specifies the status of the upgrade of the agent on a physical server. &#39;kIdle&#39; indicates there is no agent upgrade in progress. &#39;kAccepted&#39; indicates the Agent upgrade is accepted. &#39;kStarted&#39; indicates the Agent upgrade is in progress. &#39;kFinished&#39; indicates the Agent upgrade is completed. | [optional] 
**Upgradability** | Pointer to **NullableString** | Specifies the upgradability of the agent running on the server. Specifies the upgradability of the agent running on the physical server. &#39;kUpgradable&#39; indicates the Agent can be upgraded to the agent software version on the cluster. &#39;kCurrent&#39; indicates the Agent is running the latest version. &#39;kUnknown&#39; indicates the Agent&#39;s version is not known. &#39;kNonUpgradableInvalidVersion&#39; indicates the Agent&#39;s version is invalid. &#39;kNonUpgradableAgentIsNewer&#39; indicates the Agent&#39;s version is newer than the agent software version the cluster. &#39;kNonUpgradableAgentIsOld&#39; indicates the Agent&#39;s version is too old that does not support upgrades. | [optional] 
**UpgradeStatusMessage** | Pointer to **NullableString** | Specifies detailed message about the agent upgrade failure. This field is not set for successful upgrade. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the Cohesity software version of the agent. For example: 6.0.1-release-YYYYMMDD_&lt;hash&gt; | [optional] 

## Methods

### NewAgentDeploymentStatusResponse

`func NewAgentDeploymentStatusResponse() *AgentDeploymentStatusResponse`

NewAgentDeploymentStatusResponse instantiates a new AgentDeploymentStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentDeploymentStatusResponseWithDefaults

`func NewAgentDeploymentStatusResponseWithDefaults() *AgentDeploymentStatusResponse`

NewAgentDeploymentStatusResponseWithDefaults instantiates a new AgentDeploymentStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompactVersion

`func (o *AgentDeploymentStatusResponse) GetCompactVersion() string`

GetCompactVersion returns the CompactVersion field if non-nil, zero value otherwise.

### GetCompactVersionOk

`func (o *AgentDeploymentStatusResponse) GetCompactVersionOk() (*string, bool)`

GetCompactVersionOk returns a tuple with the CompactVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompactVersion

`func (o *AgentDeploymentStatusResponse) SetCompactVersion(v string)`

SetCompactVersion sets CompactVersion field to given value.

### HasCompactVersion

`func (o *AgentDeploymentStatusResponse) HasCompactVersion() bool`

HasCompactVersion returns a boolean if a field has been set.

### SetCompactVersionNil

`func (o *AgentDeploymentStatusResponse) SetCompactVersionNil(b bool)`

 SetCompactVersionNil sets the value for CompactVersion to be an explicit nil

### UnsetCompactVersion
`func (o *AgentDeploymentStatusResponse) UnsetCompactVersion()`

UnsetCompactVersion ensures that no value is present for CompactVersion, not even an explicit nil
### GetHealthStatus

`func (o *AgentDeploymentStatusResponse) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *AgentDeploymentStatusResponse) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *AgentDeploymentStatusResponse) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *AgentDeploymentStatusResponse) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### SetHealthStatusNil

`func (o *AgentDeploymentStatusResponse) SetHealthStatusNil(b bool)`

 SetHealthStatusNil sets the value for HealthStatus to be an explicit nil

### UnsetHealthStatus
`func (o *AgentDeploymentStatusResponse) UnsetHealthStatus()`

UnsetHealthStatus ensures that no value is present for HealthStatus, not even an explicit nil
### GetHostIp

`func (o *AgentDeploymentStatusResponse) GetHostIp() string`

GetHostIp returns the HostIp field if non-nil, zero value otherwise.

### GetHostIpOk

`func (o *AgentDeploymentStatusResponse) GetHostIpOk() (*string, bool)`

GetHostIpOk returns a tuple with the HostIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIp

`func (o *AgentDeploymentStatusResponse) SetHostIp(v string)`

SetHostIp sets HostIp field to given value.

### HasHostIp

`func (o *AgentDeploymentStatusResponse) HasHostIp() bool`

HasHostIp returns a boolean if a field has been set.

### SetHostIpNil

`func (o *AgentDeploymentStatusResponse) SetHostIpNil(b bool)`

 SetHostIpNil sets the value for HostIp to be an explicit nil

### UnsetHostIp
`func (o *AgentDeploymentStatusResponse) UnsetHostIp()`

UnsetHostIp ensures that no value is present for HostIp, not even an explicit nil
### GetHostOsType

`func (o *AgentDeploymentStatusResponse) GetHostOsType() string`

GetHostOsType returns the HostOsType field if non-nil, zero value otherwise.

### GetHostOsTypeOk

`func (o *AgentDeploymentStatusResponse) GetHostOsTypeOk() (*string, bool)`

GetHostOsTypeOk returns a tuple with the HostOsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOsType

`func (o *AgentDeploymentStatusResponse) SetHostOsType(v string)`

SetHostOsType sets HostOsType field to given value.

### HasHostOsType

`func (o *AgentDeploymentStatusResponse) HasHostOsType() bool`

HasHostOsType returns a boolean if a field has been set.

### SetHostOsTypeNil

`func (o *AgentDeploymentStatusResponse) SetHostOsTypeNil(b bool)`

 SetHostOsTypeNil sets the value for HostOsType to be an explicit nil

### UnsetHostOsType
`func (o *AgentDeploymentStatusResponse) UnsetHostOsType()`

UnsetHostOsType ensures that no value is present for HostOsType, not even an explicit nil
### GetLastUpgradeStatus

`func (o *AgentDeploymentStatusResponse) GetLastUpgradeStatus() string`

GetLastUpgradeStatus returns the LastUpgradeStatus field if non-nil, zero value otherwise.

### GetLastUpgradeStatusOk

`func (o *AgentDeploymentStatusResponse) GetLastUpgradeStatusOk() (*string, bool)`

GetLastUpgradeStatusOk returns a tuple with the LastUpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgradeStatus

`func (o *AgentDeploymentStatusResponse) SetLastUpgradeStatus(v string)`

SetLastUpgradeStatus sets LastUpgradeStatus field to given value.

### HasLastUpgradeStatus

`func (o *AgentDeploymentStatusResponse) HasLastUpgradeStatus() bool`

HasLastUpgradeStatus returns a boolean if a field has been set.

### SetLastUpgradeStatusNil

`func (o *AgentDeploymentStatusResponse) SetLastUpgradeStatusNil(b bool)`

 SetLastUpgradeStatusNil sets the value for LastUpgradeStatus to be an explicit nil

### UnsetLastUpgradeStatus
`func (o *AgentDeploymentStatusResponse) UnsetLastUpgradeStatus()`

UnsetLastUpgradeStatus ensures that no value is present for LastUpgradeStatus, not even an explicit nil
### GetUpgradability

`func (o *AgentDeploymentStatusResponse) GetUpgradability() string`

GetUpgradability returns the Upgradability field if non-nil, zero value otherwise.

### GetUpgradabilityOk

`func (o *AgentDeploymentStatusResponse) GetUpgradabilityOk() (*string, bool)`

GetUpgradabilityOk returns a tuple with the Upgradability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradability

`func (o *AgentDeploymentStatusResponse) SetUpgradability(v string)`

SetUpgradability sets Upgradability field to given value.

### HasUpgradability

`func (o *AgentDeploymentStatusResponse) HasUpgradability() bool`

HasUpgradability returns a boolean if a field has been set.

### SetUpgradabilityNil

`func (o *AgentDeploymentStatusResponse) SetUpgradabilityNil(b bool)`

 SetUpgradabilityNil sets the value for Upgradability to be an explicit nil

### UnsetUpgradability
`func (o *AgentDeploymentStatusResponse) UnsetUpgradability()`

UnsetUpgradability ensures that no value is present for Upgradability, not even an explicit nil
### GetUpgradeStatusMessage

`func (o *AgentDeploymentStatusResponse) GetUpgradeStatusMessage() string`

GetUpgradeStatusMessage returns the UpgradeStatusMessage field if non-nil, zero value otherwise.

### GetUpgradeStatusMessageOk

`func (o *AgentDeploymentStatusResponse) GetUpgradeStatusMessageOk() (*string, bool)`

GetUpgradeStatusMessageOk returns a tuple with the UpgradeStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatusMessage

`func (o *AgentDeploymentStatusResponse) SetUpgradeStatusMessage(v string)`

SetUpgradeStatusMessage sets UpgradeStatusMessage field to given value.

### HasUpgradeStatusMessage

`func (o *AgentDeploymentStatusResponse) HasUpgradeStatusMessage() bool`

HasUpgradeStatusMessage returns a boolean if a field has been set.

### SetUpgradeStatusMessageNil

`func (o *AgentDeploymentStatusResponse) SetUpgradeStatusMessageNil(b bool)`

 SetUpgradeStatusMessageNil sets the value for UpgradeStatusMessage to be an explicit nil

### UnsetUpgradeStatusMessage
`func (o *AgentDeploymentStatusResponse) UnsetUpgradeStatusMessage()`

UnsetUpgradeStatusMessage ensures that no value is present for UpgradeStatusMessage, not even an explicit nil
### GetVersion

`func (o *AgentDeploymentStatusResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgentDeploymentStatusResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgentDeploymentStatusResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AgentDeploymentStatusResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *AgentDeploymentStatusResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AgentDeploymentStatusResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


