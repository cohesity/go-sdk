# AgentInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CbmrVersion** | Pointer to **NullableString** | Specifies the version if Cristie BMR product is installed on the host. | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the host type where the agent is running. This is only set for persistent agents. &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the agent&#39;s id. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the agent&#39;s name. | [optional] 
**OracleMultiNodeChannelSupported** | Pointer to **NullableBool** | Specifies whether oracle multi node multi channel is supported or not. | [optional] 
**RegistrationInfo** | Pointer to [**RegisteredSourceInfo**](RegisteredSourceInfo.md) |  | [optional] 
**SourceSideDedupEnabled** | Pointer to **NullableBool** | Specifies whether source side dedup is enabled or not. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the agent status. Specifies the status of the agent running on a physical source. &#39;kUnknown&#39; indicates the Agent is not known. No attempt to connect to the Agent has occurred. &#39;kUnreachable&#39; indicates the Agent is not reachable. &#39;kHealthy&#39; indicates the Agent is healthy. &#39;kDegraded&#39; indicates the Agent is running but in a degraded state. | [optional] 
**StatusMessage** | Pointer to **NullableString** | Specifies additional details about the agent status. | [optional] 
**Upgradability** | Pointer to **NullableString** | Specifies the upgradability of the agent running on the physical server. Specifies the upgradability of the agent running on the physical server. &#39;kUpgradable&#39; indicates the Agent can be upgraded to the agent software version on the cluster. &#39;kCurrent&#39; indicates the Agent is running the latest version. &#39;kUnknown&#39; indicates the Agent&#39;s version is not known. &#39;kNonUpgradableInvalidVersion&#39; indicates the Agent&#39;s version is invalid. &#39;kNonUpgradableAgentIsNewer&#39; indicates the Agent&#39;s version is newer than the agent software version the cluster. &#39;kNonUpgradableAgentIsOld&#39; indicates the Agent&#39;s version is too old that does not support upgrades. | [optional] 
**UpgradeStatus** | Pointer to **NullableString** | Specifies the status of the upgrade of the agent on a physical server. Specifies the status of the upgrade of the agent on a physical server. &#39;kIdle&#39; indicates there is no agent upgrade in progress. &#39;kAccepted&#39; indicates the Agent upgrade is accepted. &#39;kStarted&#39; indicates the Agent upgrade is in progress. &#39;kFinished&#39; indicates the Agent upgrade is completed. | [optional] 
**UpgradeStatusMessage** | Pointer to **NullableString** | Specifies detailed message about the agent upgrade failure. This field is not set for successful upgrade. | [optional] 
**Version** | Pointer to **NullableString** | Specifies the version of the Agent software. | [optional] 

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

### GetCbmrVersion

`func (o *AgentInformation) GetCbmrVersion() string`

GetCbmrVersion returns the CbmrVersion field if non-nil, zero value otherwise.

### GetCbmrVersionOk

`func (o *AgentInformation) GetCbmrVersionOk() (*string, bool)`

GetCbmrVersionOk returns a tuple with the CbmrVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbmrVersion

`func (o *AgentInformation) SetCbmrVersion(v string)`

SetCbmrVersion sets CbmrVersion field to given value.

### HasCbmrVersion

`func (o *AgentInformation) HasCbmrVersion() bool`

HasCbmrVersion returns a boolean if a field has been set.

### SetCbmrVersionNil

`func (o *AgentInformation) SetCbmrVersionNil(b bool)`

 SetCbmrVersionNil sets the value for CbmrVersion to be an explicit nil

### UnsetCbmrVersion
`func (o *AgentInformation) UnsetCbmrVersion()`

UnsetCbmrVersion ensures that no value is present for CbmrVersion, not even an explicit nil
### GetHostType

`func (o *AgentInformation) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *AgentInformation) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *AgentInformation) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *AgentInformation) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *AgentInformation) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *AgentInformation) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetId

`func (o *AgentInformation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentInformation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentInformation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AgentInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AgentInformation) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AgentInformation) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *AgentInformation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentInformation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentInformation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentInformation) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AgentInformation) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AgentInformation) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOracleMultiNodeChannelSupported

`func (o *AgentInformation) GetOracleMultiNodeChannelSupported() bool`

GetOracleMultiNodeChannelSupported returns the OracleMultiNodeChannelSupported field if non-nil, zero value otherwise.

### GetOracleMultiNodeChannelSupportedOk

`func (o *AgentInformation) GetOracleMultiNodeChannelSupportedOk() (*bool, bool)`

GetOracleMultiNodeChannelSupportedOk returns a tuple with the OracleMultiNodeChannelSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleMultiNodeChannelSupported

`func (o *AgentInformation) SetOracleMultiNodeChannelSupported(v bool)`

SetOracleMultiNodeChannelSupported sets OracleMultiNodeChannelSupported field to given value.

### HasOracleMultiNodeChannelSupported

`func (o *AgentInformation) HasOracleMultiNodeChannelSupported() bool`

HasOracleMultiNodeChannelSupported returns a boolean if a field has been set.

### SetOracleMultiNodeChannelSupportedNil

`func (o *AgentInformation) SetOracleMultiNodeChannelSupportedNil(b bool)`

 SetOracleMultiNodeChannelSupportedNil sets the value for OracleMultiNodeChannelSupported to be an explicit nil

### UnsetOracleMultiNodeChannelSupported
`func (o *AgentInformation) UnsetOracleMultiNodeChannelSupported()`

UnsetOracleMultiNodeChannelSupported ensures that no value is present for OracleMultiNodeChannelSupported, not even an explicit nil
### GetRegistrationInfo

`func (o *AgentInformation) GetRegistrationInfo() RegisteredSourceInfo`

GetRegistrationInfo returns the RegistrationInfo field if non-nil, zero value otherwise.

### GetRegistrationInfoOk

`func (o *AgentInformation) GetRegistrationInfoOk() (*RegisteredSourceInfo, bool)`

GetRegistrationInfoOk returns a tuple with the RegistrationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationInfo

`func (o *AgentInformation) SetRegistrationInfo(v RegisteredSourceInfo)`

SetRegistrationInfo sets RegistrationInfo field to given value.

### HasRegistrationInfo

`func (o *AgentInformation) HasRegistrationInfo() bool`

HasRegistrationInfo returns a boolean if a field has been set.

### GetSourceSideDedupEnabled

`func (o *AgentInformation) GetSourceSideDedupEnabled() bool`

GetSourceSideDedupEnabled returns the SourceSideDedupEnabled field if non-nil, zero value otherwise.

### GetSourceSideDedupEnabledOk

`func (o *AgentInformation) GetSourceSideDedupEnabledOk() (*bool, bool)`

GetSourceSideDedupEnabledOk returns a tuple with the SourceSideDedupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDedupEnabled

`func (o *AgentInformation) SetSourceSideDedupEnabled(v bool)`

SetSourceSideDedupEnabled sets SourceSideDedupEnabled field to given value.

### HasSourceSideDedupEnabled

`func (o *AgentInformation) HasSourceSideDedupEnabled() bool`

HasSourceSideDedupEnabled returns a boolean if a field has been set.

### SetSourceSideDedupEnabledNil

`func (o *AgentInformation) SetSourceSideDedupEnabledNil(b bool)`

 SetSourceSideDedupEnabledNil sets the value for SourceSideDedupEnabled to be an explicit nil

### UnsetSourceSideDedupEnabled
`func (o *AgentInformation) UnsetSourceSideDedupEnabled()`

UnsetSourceSideDedupEnabled ensures that no value is present for SourceSideDedupEnabled, not even an explicit nil
### GetStatus

`func (o *AgentInformation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentInformation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentInformation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentInformation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AgentInformation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AgentInformation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusMessage

`func (o *AgentInformation) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AgentInformation) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AgentInformation) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AgentInformation) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AgentInformation) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AgentInformation) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetUpgradability

`func (o *AgentInformation) GetUpgradability() string`

GetUpgradability returns the Upgradability field if non-nil, zero value otherwise.

### GetUpgradabilityOk

`func (o *AgentInformation) GetUpgradabilityOk() (*string, bool)`

GetUpgradabilityOk returns a tuple with the Upgradability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradability

`func (o *AgentInformation) SetUpgradability(v string)`

SetUpgradability sets Upgradability field to given value.

### HasUpgradability

`func (o *AgentInformation) HasUpgradability() bool`

HasUpgradability returns a boolean if a field has been set.

### SetUpgradabilityNil

`func (o *AgentInformation) SetUpgradabilityNil(b bool)`

 SetUpgradabilityNil sets the value for Upgradability to be an explicit nil

### UnsetUpgradability
`func (o *AgentInformation) UnsetUpgradability()`

UnsetUpgradability ensures that no value is present for Upgradability, not even an explicit nil
### GetUpgradeStatus

`func (o *AgentInformation) GetUpgradeStatus() string`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *AgentInformation) GetUpgradeStatusOk() (*string, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *AgentInformation) SetUpgradeStatus(v string)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *AgentInformation) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.

### SetUpgradeStatusNil

`func (o *AgentInformation) SetUpgradeStatusNil(b bool)`

 SetUpgradeStatusNil sets the value for UpgradeStatus to be an explicit nil

### UnsetUpgradeStatus
`func (o *AgentInformation) UnsetUpgradeStatus()`

UnsetUpgradeStatus ensures that no value is present for UpgradeStatus, not even an explicit nil
### GetUpgradeStatusMessage

`func (o *AgentInformation) GetUpgradeStatusMessage() string`

GetUpgradeStatusMessage returns the UpgradeStatusMessage field if non-nil, zero value otherwise.

### GetUpgradeStatusMessageOk

`func (o *AgentInformation) GetUpgradeStatusMessageOk() (*string, bool)`

GetUpgradeStatusMessageOk returns a tuple with the UpgradeStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatusMessage

`func (o *AgentInformation) SetUpgradeStatusMessage(v string)`

SetUpgradeStatusMessage sets UpgradeStatusMessage field to given value.

### HasUpgradeStatusMessage

`func (o *AgentInformation) HasUpgradeStatusMessage() bool`

HasUpgradeStatusMessage returns a boolean if a field has been set.

### SetUpgradeStatusMessageNil

`func (o *AgentInformation) SetUpgradeStatusMessageNil(b bool)`

 SetUpgradeStatusMessageNil sets the value for UpgradeStatusMessage to be an explicit nil

### UnsetUpgradeStatusMessage
`func (o *AgentInformation) UnsetUpgradeStatusMessage()`

UnsetUpgradeStatusMessage ensures that no value is present for UpgradeStatusMessage, not even an explicit nil
### GetVersion

`func (o *AgentInformation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgentInformation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgentInformation) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AgentInformation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *AgentInformation) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AgentInformation) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


