# RemoteClusterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewRemoteClusterParams

`func NewRemoteClusterParams() *RemoteClusterParams`

NewRemoteClusterParams instantiates a new RemoteClusterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteClusterParamsWithDefaults

`func NewRemoteClusterParamsWithDefaults() *RemoteClusterParams`

NewRemoteClusterParamsWithDefaults instantiates a new RemoteClusterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRegisterTarget

`func (o *RemoteClusterParams) GetAutoRegisterTarget() bool`

GetAutoRegisterTarget returns the AutoRegisterTarget field if non-nil, zero value otherwise.

### GetAutoRegisterTargetOk

`func (o *RemoteClusterParams) GetAutoRegisterTargetOk() (*bool, bool)`

GetAutoRegisterTargetOk returns a tuple with the AutoRegisterTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegisterTarget

`func (o *RemoteClusterParams) SetAutoRegisterTarget(v bool)`

SetAutoRegisterTarget sets AutoRegisterTarget field to given value.

### HasAutoRegisterTarget

`func (o *RemoteClusterParams) HasAutoRegisterTarget() bool`

HasAutoRegisterTarget returns a boolean if a field has been set.

### SetAutoRegisterTargetNil

`func (o *RemoteClusterParams) SetAutoRegisterTargetNil(b bool)`

 SetAutoRegisterTargetNil sets the value for AutoRegisterTarget to be an explicit nil

### UnsetAutoRegisterTarget
`func (o *RemoteClusterParams) UnsetAutoRegisterTarget()`

UnsetAutoRegisterTarget ensures that no value is present for AutoRegisterTarget, not even an explicit nil
### GetClusterId

`func (o *RemoteClusterParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RemoteClusterParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RemoteClusterParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *RemoteClusterParams) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *RemoteClusterParams) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *RemoteClusterParams) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *RemoteClusterParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *RemoteClusterParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *RemoteClusterParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *RemoteClusterParams) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *RemoteClusterParams) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *RemoteClusterParams) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetClusterName

`func (o *RemoteClusterParams) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *RemoteClusterParams) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *RemoteClusterParams) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *RemoteClusterParams) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *RemoteClusterParams) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *RemoteClusterParams) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetDescription

`func (o *RemoteClusterParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RemoteClusterParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RemoteClusterParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RemoteClusterParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RemoteClusterParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RemoteClusterParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEffectiveAesEncryptionMode

`func (o *RemoteClusterParams) GetEffectiveAesEncryptionMode() string`

GetEffectiveAesEncryptionMode returns the EffectiveAesEncryptionMode field if non-nil, zero value otherwise.

### GetEffectiveAesEncryptionModeOk

`func (o *RemoteClusterParams) GetEffectiveAesEncryptionModeOk() (*string, bool)`

GetEffectiveAesEncryptionModeOk returns a tuple with the EffectiveAesEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAesEncryptionMode

`func (o *RemoteClusterParams) SetEffectiveAesEncryptionMode(v string)`

SetEffectiveAesEncryptionMode sets EffectiveAesEncryptionMode field to given value.

### HasEffectiveAesEncryptionMode

`func (o *RemoteClusterParams) HasEffectiveAesEncryptionMode() bool`

HasEffectiveAesEncryptionMode returns a boolean if a field has been set.

### SetEffectiveAesEncryptionModeNil

`func (o *RemoteClusterParams) SetEffectiveAesEncryptionModeNil(b bool)`

 SetEffectiveAesEncryptionModeNil sets the value for EffectiveAesEncryptionMode to be an explicit nil

### UnsetEffectiveAesEncryptionMode
`func (o *RemoteClusterParams) UnsetEffectiveAesEncryptionMode()`

UnsetEffectiveAesEncryptionMode ensures that no value is present for EffectiveAesEncryptionMode, not even an explicit nil
### GetIsAutoRegistered

`func (o *RemoteClusterParams) GetIsAutoRegistered() bool`

GetIsAutoRegistered returns the IsAutoRegistered field if non-nil, zero value otherwise.

### GetIsAutoRegisteredOk

`func (o *RemoteClusterParams) GetIsAutoRegisteredOk() (*bool, bool)`

GetIsAutoRegisteredOk returns a tuple with the IsAutoRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRegistered

`func (o *RemoteClusterParams) SetIsAutoRegistered(v bool)`

SetIsAutoRegistered sets IsAutoRegistered field to given value.

### HasIsAutoRegistered

`func (o *RemoteClusterParams) HasIsAutoRegistered() bool`

HasIsAutoRegistered returns a boolean if a field has been set.

### SetIsAutoRegisteredNil

`func (o *RemoteClusterParams) SetIsAutoRegisteredNil(b bool)`

 SetIsAutoRegisteredNil sets the value for IsAutoRegistered to be an explicit nil

### UnsetIsAutoRegistered
`func (o *RemoteClusterParams) UnsetIsAutoRegistered()`

UnsetIsAutoRegistered ensures that no value is present for IsAutoRegistered, not even an explicit nil
### GetLocalAddresses

`func (o *RemoteClusterParams) GetLocalAddresses() []string`

GetLocalAddresses returns the LocalAddresses field if non-nil, zero value otherwise.

### GetLocalAddressesOk

`func (o *RemoteClusterParams) GetLocalAddressesOk() (*[]string, bool)`

GetLocalAddressesOk returns a tuple with the LocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddresses

`func (o *RemoteClusterParams) SetLocalAddresses(v []string)`

SetLocalAddresses sets LocalAddresses field to given value.

### HasLocalAddresses

`func (o *RemoteClusterParams) HasLocalAddresses() bool`

HasLocalAddresses returns a boolean if a field has been set.

### GetMultiTenancyEnabled

`func (o *RemoteClusterParams) GetMultiTenancyEnabled() bool`

GetMultiTenancyEnabled returns the MultiTenancyEnabled field if non-nil, zero value otherwise.

### GetMultiTenancyEnabledOk

`func (o *RemoteClusterParams) GetMultiTenancyEnabledOk() (*bool, bool)`

GetMultiTenancyEnabledOk returns a tuple with the MultiTenancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenancyEnabled

`func (o *RemoteClusterParams) SetMultiTenancyEnabled(v bool)`

SetMultiTenancyEnabled sets MultiTenancyEnabled field to given value.

### HasMultiTenancyEnabled

`func (o *RemoteClusterParams) HasMultiTenancyEnabled() bool`

HasMultiTenancyEnabled returns a boolean if a field has been set.

### SetMultiTenancyEnabledNil

`func (o *RemoteClusterParams) SetMultiTenancyEnabledNil(b bool)`

 SetMultiTenancyEnabledNil sets the value for MultiTenancyEnabled to be an explicit nil

### UnsetMultiTenancyEnabled
`func (o *RemoteClusterParams) UnsetMultiTenancyEnabled()`

UnsetMultiTenancyEnabled ensures that no value is present for MultiTenancyEnabled, not even an explicit nil
### GetNetworkInterface

`func (o *RemoteClusterParams) GetNetworkInterface() string`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *RemoteClusterParams) GetNetworkInterfaceOk() (*string, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *RemoteClusterParams) SetNetworkInterface(v string)`

SetNetworkInterface sets NetworkInterface field to given value.

### HasNetworkInterface

`func (o *RemoteClusterParams) HasNetworkInterface() bool`

HasNetworkInterface returns a boolean if a field has been set.

### SetNetworkInterfaceNil

`func (o *RemoteClusterParams) SetNetworkInterfaceNil(b bool)`

 SetNetworkInterfaceNil sets the value for NetworkInterface to be an explicit nil

### UnsetNetworkInterface
`func (o *RemoteClusterParams) UnsetNetworkInterface()`

UnsetNetworkInterface ensures that no value is present for NetworkInterface, not even an explicit nil
### GetPurpose

`func (o *RemoteClusterParams) GetPurpose() []string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *RemoteClusterParams) GetPurposeOk() (*[]string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *RemoteClusterParams) SetPurpose(v []string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *RemoteClusterParams) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *RemoteClusterParams) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *RemoteClusterParams) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetReplicationParams

`func (o *RemoteClusterParams) GetReplicationParams() RemoteClusterParamsReplicationParams`

GetReplicationParams returns the ReplicationParams field if non-nil, zero value otherwise.

### GetReplicationParamsOk

`func (o *RemoteClusterParams) GetReplicationParamsOk() (*RemoteClusterParamsReplicationParams, bool)`

GetReplicationParamsOk returns a tuple with the ReplicationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationParams

`func (o *RemoteClusterParams) SetReplicationParams(v RemoteClusterParamsReplicationParams)`

SetReplicationParams sets ReplicationParams field to given value.

### HasReplicationParams

`func (o *RemoteClusterParams) HasReplicationParams() bool`

HasReplicationParams returns a boolean if a field has been set.

### GetSupportedAesEncryptionMode

`func (o *RemoteClusterParams) GetSupportedAesEncryptionMode() string`

GetSupportedAesEncryptionMode returns the SupportedAesEncryptionMode field if non-nil, zero value otherwise.

### GetSupportedAesEncryptionModeOk

`func (o *RemoteClusterParams) GetSupportedAesEncryptionModeOk() (*string, bool)`

GetSupportedAesEncryptionModeOk returns a tuple with the SupportedAesEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAesEncryptionMode

`func (o *RemoteClusterParams) SetSupportedAesEncryptionMode(v string)`

SetSupportedAesEncryptionMode sets SupportedAesEncryptionMode field to given value.

### HasSupportedAesEncryptionMode

`func (o *RemoteClusterParams) HasSupportedAesEncryptionMode() bool`

HasSupportedAesEncryptionMode returns a boolean if a field has been set.

### SetSupportedAesEncryptionModeNil

`func (o *RemoteClusterParams) SetSupportedAesEncryptionModeNil(b bool)`

 SetSupportedAesEncryptionModeNil sets the value for SupportedAesEncryptionMode to be an explicit nil

### UnsetSupportedAesEncryptionMode
`func (o *RemoteClusterParams) UnsetSupportedAesEncryptionMode()`

UnsetSupportedAesEncryptionMode ensures that no value is present for SupportedAesEncryptionMode, not even an explicit nil
### GetTenantId

`func (o *RemoteClusterParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RemoteClusterParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RemoteClusterParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *RemoteClusterParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *RemoteClusterParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *RemoteClusterParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantStorageDomainSharingEnabled

`func (o *RemoteClusterParams) GetTenantStorageDomainSharingEnabled() bool`

GetTenantStorageDomainSharingEnabled returns the TenantStorageDomainSharingEnabled field if non-nil, zero value otherwise.

### GetTenantStorageDomainSharingEnabledOk

`func (o *RemoteClusterParams) GetTenantStorageDomainSharingEnabledOk() (*bool, bool)`

GetTenantStorageDomainSharingEnabledOk returns a tuple with the TenantStorageDomainSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantStorageDomainSharingEnabled

`func (o *RemoteClusterParams) SetTenantStorageDomainSharingEnabled(v bool)`

SetTenantStorageDomainSharingEnabled sets TenantStorageDomainSharingEnabled field to given value.

### HasTenantStorageDomainSharingEnabled

`func (o *RemoteClusterParams) HasTenantStorageDomainSharingEnabled() bool`

HasTenantStorageDomainSharingEnabled returns a boolean if a field has been set.

### SetTenantStorageDomainSharingEnabledNil

`func (o *RemoteClusterParams) SetTenantStorageDomainSharingEnabledNil(b bool)`

 SetTenantStorageDomainSharingEnabledNil sets the value for TenantStorageDomainSharingEnabled to be an explicit nil

### UnsetTenantStorageDomainSharingEnabled
`func (o *RemoteClusterParams) UnsetTenantStorageDomainSharingEnabled()`

UnsetTenantStorageDomainSharingEnabled ensures that no value is present for TenantStorageDomainSharingEnabled, not even an explicit nil
### GetTlsEnabled

`func (o *RemoteClusterParams) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *RemoteClusterParams) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *RemoteClusterParams) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *RemoteClusterParams) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### SetTlsEnabledNil

`func (o *RemoteClusterParams) SetTlsEnabledNil(b bool)`

 SetTlsEnabledNil sets the value for TlsEnabled to be an explicit nil

### UnsetTlsEnabled
`func (o *RemoteClusterParams) UnsetTlsEnabled()`

UnsetTlsEnabled ensures that no value is present for TlsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


