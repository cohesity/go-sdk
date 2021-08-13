# McmPhysicalSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upgradability** | Pointer to **NullableString** | Specifies the upgradability of the agent software. | [optional] 
**UpgradeStatus** | Pointer to **NullableString** | Specifies the current upgrade status of the agent. | [optional] 
**UpgradeError** | Pointer to **NullableString** | Specifies the upgrade error if any for the agent. | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the operating system of the physical host. | [optional] 
**AgentsInfo** | Pointer to [**[]McmAgentInfo**](McmAgentInfo.md) | Specifies the information for agents related to source. | [optional] 

## Methods

### NewMcmPhysicalSourceInfo

`func NewMcmPhysicalSourceInfo() *McmPhysicalSourceInfo`

NewMcmPhysicalSourceInfo instantiates a new McmPhysicalSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmPhysicalSourceInfoWithDefaults

`func NewMcmPhysicalSourceInfoWithDefaults() *McmPhysicalSourceInfo`

NewMcmPhysicalSourceInfoWithDefaults instantiates a new McmPhysicalSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradability

`func (o *McmPhysicalSourceInfo) GetUpgradability() string`

GetUpgradability returns the Upgradability field if non-nil, zero value otherwise.

### GetUpgradabilityOk

`func (o *McmPhysicalSourceInfo) GetUpgradabilityOk() (*string, bool)`

GetUpgradabilityOk returns a tuple with the Upgradability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradability

`func (o *McmPhysicalSourceInfo) SetUpgradability(v string)`

SetUpgradability sets Upgradability field to given value.

### HasUpgradability

`func (o *McmPhysicalSourceInfo) HasUpgradability() bool`

HasUpgradability returns a boolean if a field has been set.

### SetUpgradabilityNil

`func (o *McmPhysicalSourceInfo) SetUpgradabilityNil(b bool)`

 SetUpgradabilityNil sets the value for Upgradability to be an explicit nil

### UnsetUpgradability
`func (o *McmPhysicalSourceInfo) UnsetUpgradability()`

UnsetUpgradability ensures that no value is present for Upgradability, not even an explicit nil
### GetUpgradeStatus

`func (o *McmPhysicalSourceInfo) GetUpgradeStatus() string`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *McmPhysicalSourceInfo) GetUpgradeStatusOk() (*string, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *McmPhysicalSourceInfo) SetUpgradeStatus(v string)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *McmPhysicalSourceInfo) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.

### SetUpgradeStatusNil

`func (o *McmPhysicalSourceInfo) SetUpgradeStatusNil(b bool)`

 SetUpgradeStatusNil sets the value for UpgradeStatus to be an explicit nil

### UnsetUpgradeStatus
`func (o *McmPhysicalSourceInfo) UnsetUpgradeStatus()`

UnsetUpgradeStatus ensures that no value is present for UpgradeStatus, not even an explicit nil
### GetUpgradeError

`func (o *McmPhysicalSourceInfo) GetUpgradeError() string`

GetUpgradeError returns the UpgradeError field if non-nil, zero value otherwise.

### GetUpgradeErrorOk

`func (o *McmPhysicalSourceInfo) GetUpgradeErrorOk() (*string, bool)`

GetUpgradeErrorOk returns a tuple with the UpgradeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeError

`func (o *McmPhysicalSourceInfo) SetUpgradeError(v string)`

SetUpgradeError sets UpgradeError field to given value.

### HasUpgradeError

`func (o *McmPhysicalSourceInfo) HasUpgradeError() bool`

HasUpgradeError returns a boolean if a field has been set.

### SetUpgradeErrorNil

`func (o *McmPhysicalSourceInfo) SetUpgradeErrorNil(b bool)`

 SetUpgradeErrorNil sets the value for UpgradeError to be an explicit nil

### UnsetUpgradeError
`func (o *McmPhysicalSourceInfo) UnsetUpgradeError()`

UnsetUpgradeError ensures that no value is present for UpgradeError, not even an explicit nil
### GetHostType

`func (o *McmPhysicalSourceInfo) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *McmPhysicalSourceInfo) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *McmPhysicalSourceInfo) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *McmPhysicalSourceInfo) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *McmPhysicalSourceInfo) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *McmPhysicalSourceInfo) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetAgentsInfo

`func (o *McmPhysicalSourceInfo) GetAgentsInfo() []McmAgentInfo`

GetAgentsInfo returns the AgentsInfo field if non-nil, zero value otherwise.

### GetAgentsInfoOk

`func (o *McmPhysicalSourceInfo) GetAgentsInfoOk() (*[]McmAgentInfo, bool)`

GetAgentsInfoOk returns a tuple with the AgentsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentsInfo

`func (o *McmPhysicalSourceInfo) SetAgentsInfo(v []McmAgentInfo)`

SetAgentsInfo sets AgentsInfo field to given value.

### HasAgentsInfo

`func (o *McmPhysicalSourceInfo) HasAgentsInfo() bool`

HasAgentsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


