# RegisterRemoteClusterParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeAddresses** | **[]string** | Specifies the VIP or IP addresses of the Nodes on the Remote Cluster to connect with. Hostnames are not supported. | 
**Password** | **string** | Specifies the password for Cohesity user to use when connecting to the Remote Cluster. | 
**Username** | **string** | Specifies the Cohesity user name used to connect to the Remote Cluster. | 
**AutoRegisterTarget** | Pointer to **NullableBool** | Specifies if the Tx clusters should be automatically registered at the Rx site. | [optional] [default to false]
**ClusterId** | Pointer to **NullableInt64** | Specifies the Remote Cluster id. | [optional] [readonly] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the Remote Cluster incarnation id. | [optional] [readonly] 
**ClusterName** | Pointer to **NullableString** | Specifies the Remote Cluster name. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Specifies any additional information if needed. | [optional] 
**EffectiveAesEncryptionMode** | Pointer to **NullableString** | Specifies the effective AES Encryption mode negotiated between local and the remote cluster. | [optional] 
**IsAutoRegistered** | Pointer to **NullableBool** | Specifies if the Remote Cluster was registered automatically or manually. | [optional] [readonly] 
**LocalAddresses** | Pointer to **[]string** | Specifies the IP addresses of the interfaces in the local Cluster which will be used for communicating with the remote Cluster. | [optional] [readonly] 
**MultiTenancyEnabled** | Pointer to **NullableBool** | Specifies if the Remote Cluster has Multi-Tenancy enabled. | [optional] 
**NetworkInterface** | Pointer to **NullableString** | Specifies the name of the network interfaces to use for communicating with the Remote Cluster. | [optional] 
**Purpose** | Pointer to **[]string** | Specifies the purpose for which the remote cluster is being registered. | [optional] 
**ReplicationParams** | Pointer to [**RemoteClusterParamsReplicationParams**](RemoteClusterParamsReplicationParams.md) |  | [optional] 
**SupportedAesEncryptionMode** | Pointer to **NullableString** | Specifies the AES Encryption mode of the remote cluster. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant Id of the Remote Cluster. | [optional] [readonly] 
**TenantStorageDomainSharingEnabled** | Pointer to **NullableBool** | Specifies if Tenant Storage Domain sharing is enabled on the Remote Cluster. | [optional] 
**TlsEnabled** | Pointer to **NullableBool** | Specifies if TLS is enabled on the Remote Cluster. | [optional] 

## Methods

### NewRegisterRemoteClusterParameters

`func NewRegisterRemoteClusterParameters(nodeAddresses []string, password string, username string, ) *RegisterRemoteClusterParameters`

NewRegisterRemoteClusterParameters instantiates a new RegisterRemoteClusterParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRemoteClusterParametersWithDefaults

`func NewRegisterRemoteClusterParametersWithDefaults() *RegisterRemoteClusterParameters`

NewRegisterRemoteClusterParametersWithDefaults instantiates a new RegisterRemoteClusterParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeAddresses

`func (o *RegisterRemoteClusterParameters) GetNodeAddresses() []string`

GetNodeAddresses returns the NodeAddresses field if non-nil, zero value otherwise.

### GetNodeAddressesOk

`func (o *RegisterRemoteClusterParameters) GetNodeAddressesOk() (*[]string, bool)`

GetNodeAddressesOk returns a tuple with the NodeAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddresses

`func (o *RegisterRemoteClusterParameters) SetNodeAddresses(v []string)`

SetNodeAddresses sets NodeAddresses field to given value.


### GetPassword

`func (o *RegisterRemoteClusterParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterRemoteClusterParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterRemoteClusterParameters) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *RegisterRemoteClusterParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterRemoteClusterParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterRemoteClusterParameters) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAutoRegisterTarget

`func (o *RegisterRemoteClusterParameters) GetAutoRegisterTarget() bool`

GetAutoRegisterTarget returns the AutoRegisterTarget field if non-nil, zero value otherwise.

### GetAutoRegisterTargetOk

`func (o *RegisterRemoteClusterParameters) GetAutoRegisterTargetOk() (*bool, bool)`

GetAutoRegisterTargetOk returns a tuple with the AutoRegisterTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegisterTarget

`func (o *RegisterRemoteClusterParameters) SetAutoRegisterTarget(v bool)`

SetAutoRegisterTarget sets AutoRegisterTarget field to given value.

### HasAutoRegisterTarget

`func (o *RegisterRemoteClusterParameters) HasAutoRegisterTarget() bool`

HasAutoRegisterTarget returns a boolean if a field has been set.

### SetAutoRegisterTargetNil

`func (o *RegisterRemoteClusterParameters) SetAutoRegisterTargetNil(b bool)`

 SetAutoRegisterTargetNil sets the value for AutoRegisterTarget to be an explicit nil

### UnsetAutoRegisterTarget
`func (o *RegisterRemoteClusterParameters) UnsetAutoRegisterTarget()`

UnsetAutoRegisterTarget ensures that no value is present for AutoRegisterTarget, not even an explicit nil
### GetClusterId

`func (o *RegisterRemoteClusterParameters) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RegisterRemoteClusterParameters) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RegisterRemoteClusterParameters) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *RegisterRemoteClusterParameters) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *RegisterRemoteClusterParameters) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *RegisterRemoteClusterParameters) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *RegisterRemoteClusterParameters) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *RegisterRemoteClusterParameters) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *RegisterRemoteClusterParameters) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *RegisterRemoteClusterParameters) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *RegisterRemoteClusterParameters) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *RegisterRemoteClusterParameters) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetClusterName

`func (o *RegisterRemoteClusterParameters) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *RegisterRemoteClusterParameters) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *RegisterRemoteClusterParameters) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *RegisterRemoteClusterParameters) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *RegisterRemoteClusterParameters) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *RegisterRemoteClusterParameters) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetDescription

`func (o *RegisterRemoteClusterParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegisterRemoteClusterParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegisterRemoteClusterParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegisterRemoteClusterParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RegisterRemoteClusterParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RegisterRemoteClusterParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEffectiveAesEncryptionMode

`func (o *RegisterRemoteClusterParameters) GetEffectiveAesEncryptionMode() string`

GetEffectiveAesEncryptionMode returns the EffectiveAesEncryptionMode field if non-nil, zero value otherwise.

### GetEffectiveAesEncryptionModeOk

`func (o *RegisterRemoteClusterParameters) GetEffectiveAesEncryptionModeOk() (*string, bool)`

GetEffectiveAesEncryptionModeOk returns a tuple with the EffectiveAesEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAesEncryptionMode

`func (o *RegisterRemoteClusterParameters) SetEffectiveAesEncryptionMode(v string)`

SetEffectiveAesEncryptionMode sets EffectiveAesEncryptionMode field to given value.

### HasEffectiveAesEncryptionMode

`func (o *RegisterRemoteClusterParameters) HasEffectiveAesEncryptionMode() bool`

HasEffectiveAesEncryptionMode returns a boolean if a field has been set.

### SetEffectiveAesEncryptionModeNil

`func (o *RegisterRemoteClusterParameters) SetEffectiveAesEncryptionModeNil(b bool)`

 SetEffectiveAesEncryptionModeNil sets the value for EffectiveAesEncryptionMode to be an explicit nil

### UnsetEffectiveAesEncryptionMode
`func (o *RegisterRemoteClusterParameters) UnsetEffectiveAesEncryptionMode()`

UnsetEffectiveAesEncryptionMode ensures that no value is present for EffectiveAesEncryptionMode, not even an explicit nil
### GetIsAutoRegistered

`func (o *RegisterRemoteClusterParameters) GetIsAutoRegistered() bool`

GetIsAutoRegistered returns the IsAutoRegistered field if non-nil, zero value otherwise.

### GetIsAutoRegisteredOk

`func (o *RegisterRemoteClusterParameters) GetIsAutoRegisteredOk() (*bool, bool)`

GetIsAutoRegisteredOk returns a tuple with the IsAutoRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRegistered

`func (o *RegisterRemoteClusterParameters) SetIsAutoRegistered(v bool)`

SetIsAutoRegistered sets IsAutoRegistered field to given value.

### HasIsAutoRegistered

`func (o *RegisterRemoteClusterParameters) HasIsAutoRegistered() bool`

HasIsAutoRegistered returns a boolean if a field has been set.

### SetIsAutoRegisteredNil

`func (o *RegisterRemoteClusterParameters) SetIsAutoRegisteredNil(b bool)`

 SetIsAutoRegisteredNil sets the value for IsAutoRegistered to be an explicit nil

### UnsetIsAutoRegistered
`func (o *RegisterRemoteClusterParameters) UnsetIsAutoRegistered()`

UnsetIsAutoRegistered ensures that no value is present for IsAutoRegistered, not even an explicit nil
### GetLocalAddresses

`func (o *RegisterRemoteClusterParameters) GetLocalAddresses() []string`

GetLocalAddresses returns the LocalAddresses field if non-nil, zero value otherwise.

### GetLocalAddressesOk

`func (o *RegisterRemoteClusterParameters) GetLocalAddressesOk() (*[]string, bool)`

GetLocalAddressesOk returns a tuple with the LocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddresses

`func (o *RegisterRemoteClusterParameters) SetLocalAddresses(v []string)`

SetLocalAddresses sets LocalAddresses field to given value.

### HasLocalAddresses

`func (o *RegisterRemoteClusterParameters) HasLocalAddresses() bool`

HasLocalAddresses returns a boolean if a field has been set.

### GetMultiTenancyEnabled

`func (o *RegisterRemoteClusterParameters) GetMultiTenancyEnabled() bool`

GetMultiTenancyEnabled returns the MultiTenancyEnabled field if non-nil, zero value otherwise.

### GetMultiTenancyEnabledOk

`func (o *RegisterRemoteClusterParameters) GetMultiTenancyEnabledOk() (*bool, bool)`

GetMultiTenancyEnabledOk returns a tuple with the MultiTenancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenancyEnabled

`func (o *RegisterRemoteClusterParameters) SetMultiTenancyEnabled(v bool)`

SetMultiTenancyEnabled sets MultiTenancyEnabled field to given value.

### HasMultiTenancyEnabled

`func (o *RegisterRemoteClusterParameters) HasMultiTenancyEnabled() bool`

HasMultiTenancyEnabled returns a boolean if a field has been set.

### SetMultiTenancyEnabledNil

`func (o *RegisterRemoteClusterParameters) SetMultiTenancyEnabledNil(b bool)`

 SetMultiTenancyEnabledNil sets the value for MultiTenancyEnabled to be an explicit nil

### UnsetMultiTenancyEnabled
`func (o *RegisterRemoteClusterParameters) UnsetMultiTenancyEnabled()`

UnsetMultiTenancyEnabled ensures that no value is present for MultiTenancyEnabled, not even an explicit nil
### GetNetworkInterface

`func (o *RegisterRemoteClusterParameters) GetNetworkInterface() string`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *RegisterRemoteClusterParameters) GetNetworkInterfaceOk() (*string, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *RegisterRemoteClusterParameters) SetNetworkInterface(v string)`

SetNetworkInterface sets NetworkInterface field to given value.

### HasNetworkInterface

`func (o *RegisterRemoteClusterParameters) HasNetworkInterface() bool`

HasNetworkInterface returns a boolean if a field has been set.

### SetNetworkInterfaceNil

`func (o *RegisterRemoteClusterParameters) SetNetworkInterfaceNil(b bool)`

 SetNetworkInterfaceNil sets the value for NetworkInterface to be an explicit nil

### UnsetNetworkInterface
`func (o *RegisterRemoteClusterParameters) UnsetNetworkInterface()`

UnsetNetworkInterface ensures that no value is present for NetworkInterface, not even an explicit nil
### GetPurpose

`func (o *RegisterRemoteClusterParameters) GetPurpose() []string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *RegisterRemoteClusterParameters) GetPurposeOk() (*[]string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *RegisterRemoteClusterParameters) SetPurpose(v []string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *RegisterRemoteClusterParameters) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *RegisterRemoteClusterParameters) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *RegisterRemoteClusterParameters) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetReplicationParams

`func (o *RegisterRemoteClusterParameters) GetReplicationParams() RemoteClusterParamsReplicationParams`

GetReplicationParams returns the ReplicationParams field if non-nil, zero value otherwise.

### GetReplicationParamsOk

`func (o *RegisterRemoteClusterParameters) GetReplicationParamsOk() (*RemoteClusterParamsReplicationParams, bool)`

GetReplicationParamsOk returns a tuple with the ReplicationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationParams

`func (o *RegisterRemoteClusterParameters) SetReplicationParams(v RemoteClusterParamsReplicationParams)`

SetReplicationParams sets ReplicationParams field to given value.

### HasReplicationParams

`func (o *RegisterRemoteClusterParameters) HasReplicationParams() bool`

HasReplicationParams returns a boolean if a field has been set.

### GetSupportedAesEncryptionMode

`func (o *RegisterRemoteClusterParameters) GetSupportedAesEncryptionMode() string`

GetSupportedAesEncryptionMode returns the SupportedAesEncryptionMode field if non-nil, zero value otherwise.

### GetSupportedAesEncryptionModeOk

`func (o *RegisterRemoteClusterParameters) GetSupportedAesEncryptionModeOk() (*string, bool)`

GetSupportedAesEncryptionModeOk returns a tuple with the SupportedAesEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAesEncryptionMode

`func (o *RegisterRemoteClusterParameters) SetSupportedAesEncryptionMode(v string)`

SetSupportedAesEncryptionMode sets SupportedAesEncryptionMode field to given value.

### HasSupportedAesEncryptionMode

`func (o *RegisterRemoteClusterParameters) HasSupportedAesEncryptionMode() bool`

HasSupportedAesEncryptionMode returns a boolean if a field has been set.

### SetSupportedAesEncryptionModeNil

`func (o *RegisterRemoteClusterParameters) SetSupportedAesEncryptionModeNil(b bool)`

 SetSupportedAesEncryptionModeNil sets the value for SupportedAesEncryptionMode to be an explicit nil

### UnsetSupportedAesEncryptionMode
`func (o *RegisterRemoteClusterParameters) UnsetSupportedAesEncryptionMode()`

UnsetSupportedAesEncryptionMode ensures that no value is present for SupportedAesEncryptionMode, not even an explicit nil
### GetTenantId

`func (o *RegisterRemoteClusterParameters) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RegisterRemoteClusterParameters) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RegisterRemoteClusterParameters) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *RegisterRemoteClusterParameters) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *RegisterRemoteClusterParameters) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *RegisterRemoteClusterParameters) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantStorageDomainSharingEnabled

`func (o *RegisterRemoteClusterParameters) GetTenantStorageDomainSharingEnabled() bool`

GetTenantStorageDomainSharingEnabled returns the TenantStorageDomainSharingEnabled field if non-nil, zero value otherwise.

### GetTenantStorageDomainSharingEnabledOk

`func (o *RegisterRemoteClusterParameters) GetTenantStorageDomainSharingEnabledOk() (*bool, bool)`

GetTenantStorageDomainSharingEnabledOk returns a tuple with the TenantStorageDomainSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantStorageDomainSharingEnabled

`func (o *RegisterRemoteClusterParameters) SetTenantStorageDomainSharingEnabled(v bool)`

SetTenantStorageDomainSharingEnabled sets TenantStorageDomainSharingEnabled field to given value.

### HasTenantStorageDomainSharingEnabled

`func (o *RegisterRemoteClusterParameters) HasTenantStorageDomainSharingEnabled() bool`

HasTenantStorageDomainSharingEnabled returns a boolean if a field has been set.

### SetTenantStorageDomainSharingEnabledNil

`func (o *RegisterRemoteClusterParameters) SetTenantStorageDomainSharingEnabledNil(b bool)`

 SetTenantStorageDomainSharingEnabledNil sets the value for TenantStorageDomainSharingEnabled to be an explicit nil

### UnsetTenantStorageDomainSharingEnabled
`func (o *RegisterRemoteClusterParameters) UnsetTenantStorageDomainSharingEnabled()`

UnsetTenantStorageDomainSharingEnabled ensures that no value is present for TenantStorageDomainSharingEnabled, not even an explicit nil
### GetTlsEnabled

`func (o *RegisterRemoteClusterParameters) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *RegisterRemoteClusterParameters) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *RegisterRemoteClusterParameters) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *RegisterRemoteClusterParameters) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### SetTlsEnabledNil

`func (o *RegisterRemoteClusterParameters) SetTlsEnabledNil(b bool)`

 SetTlsEnabledNil sets the value for TlsEnabled to be an explicit nil

### UnsetTlsEnabled
`func (o *RegisterRemoteClusterParameters) UnsetTlsEnabled()`

UnsetTlsEnabled ensures that no value is present for TlsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


