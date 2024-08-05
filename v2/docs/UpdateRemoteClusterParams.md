# UpdateRemoteClusterParams

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
**NodeAddresses** | Pointer to **[]string** | Specifies the VIP or IP addresses of the Nodes on the Remote Cluster to connect with. Hostnames are not supported. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password for Cohesity user to use when connecting to the Remote Cluster. | [optional] 
**Username** | Pointer to **string** | Specifies the Cohesity user name used to connect to the Remote Cluster. | [optional] 

## Methods

### NewUpdateRemoteClusterParams

`func NewUpdateRemoteClusterParams() *UpdateRemoteClusterParams`

NewUpdateRemoteClusterParams instantiates a new UpdateRemoteClusterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRemoteClusterParamsWithDefaults

`func NewUpdateRemoteClusterParamsWithDefaults() *UpdateRemoteClusterParams`

NewUpdateRemoteClusterParamsWithDefaults instantiates a new UpdateRemoteClusterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRegisterTarget

`func (o *UpdateRemoteClusterParams) GetAutoRegisterTarget() bool`

GetAutoRegisterTarget returns the AutoRegisterTarget field if non-nil, zero value otherwise.

### GetAutoRegisterTargetOk

`func (o *UpdateRemoteClusterParams) GetAutoRegisterTargetOk() (*bool, bool)`

GetAutoRegisterTargetOk returns a tuple with the AutoRegisterTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegisterTarget

`func (o *UpdateRemoteClusterParams) SetAutoRegisterTarget(v bool)`

SetAutoRegisterTarget sets AutoRegisterTarget field to given value.

### HasAutoRegisterTarget

`func (o *UpdateRemoteClusterParams) HasAutoRegisterTarget() bool`

HasAutoRegisterTarget returns a boolean if a field has been set.

### SetAutoRegisterTargetNil

`func (o *UpdateRemoteClusterParams) SetAutoRegisterTargetNil(b bool)`

 SetAutoRegisterTargetNil sets the value for AutoRegisterTarget to be an explicit nil

### UnsetAutoRegisterTarget
`func (o *UpdateRemoteClusterParams) UnsetAutoRegisterTarget()`

UnsetAutoRegisterTarget ensures that no value is present for AutoRegisterTarget, not even an explicit nil
### GetClusterId

`func (o *UpdateRemoteClusterParams) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *UpdateRemoteClusterParams) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *UpdateRemoteClusterParams) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *UpdateRemoteClusterParams) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *UpdateRemoteClusterParams) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *UpdateRemoteClusterParams) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *UpdateRemoteClusterParams) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *UpdateRemoteClusterParams) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *UpdateRemoteClusterParams) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *UpdateRemoteClusterParams) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *UpdateRemoteClusterParams) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *UpdateRemoteClusterParams) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetClusterName

`func (o *UpdateRemoteClusterParams) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *UpdateRemoteClusterParams) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *UpdateRemoteClusterParams) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *UpdateRemoteClusterParams) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *UpdateRemoteClusterParams) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *UpdateRemoteClusterParams) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetDescription

`func (o *UpdateRemoteClusterParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRemoteClusterParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRemoteClusterParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateRemoteClusterParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateRemoteClusterParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateRemoteClusterParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEffectiveAesEncryptionMode

`func (o *UpdateRemoteClusterParams) GetEffectiveAesEncryptionMode() string`

GetEffectiveAesEncryptionMode returns the EffectiveAesEncryptionMode field if non-nil, zero value otherwise.

### GetEffectiveAesEncryptionModeOk

`func (o *UpdateRemoteClusterParams) GetEffectiveAesEncryptionModeOk() (*string, bool)`

GetEffectiveAesEncryptionModeOk returns a tuple with the EffectiveAesEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAesEncryptionMode

`func (o *UpdateRemoteClusterParams) SetEffectiveAesEncryptionMode(v string)`

SetEffectiveAesEncryptionMode sets EffectiveAesEncryptionMode field to given value.

### HasEffectiveAesEncryptionMode

`func (o *UpdateRemoteClusterParams) HasEffectiveAesEncryptionMode() bool`

HasEffectiveAesEncryptionMode returns a boolean if a field has been set.

### SetEffectiveAesEncryptionModeNil

`func (o *UpdateRemoteClusterParams) SetEffectiveAesEncryptionModeNil(b bool)`

 SetEffectiveAesEncryptionModeNil sets the value for EffectiveAesEncryptionMode to be an explicit nil

### UnsetEffectiveAesEncryptionMode
`func (o *UpdateRemoteClusterParams) UnsetEffectiveAesEncryptionMode()`

UnsetEffectiveAesEncryptionMode ensures that no value is present for EffectiveAesEncryptionMode, not even an explicit nil
### GetIsAutoRegistered

`func (o *UpdateRemoteClusterParams) GetIsAutoRegistered() bool`

GetIsAutoRegistered returns the IsAutoRegistered field if non-nil, zero value otherwise.

### GetIsAutoRegisteredOk

`func (o *UpdateRemoteClusterParams) GetIsAutoRegisteredOk() (*bool, bool)`

GetIsAutoRegisteredOk returns a tuple with the IsAutoRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRegistered

`func (o *UpdateRemoteClusterParams) SetIsAutoRegistered(v bool)`

SetIsAutoRegistered sets IsAutoRegistered field to given value.

### HasIsAutoRegistered

`func (o *UpdateRemoteClusterParams) HasIsAutoRegistered() bool`

HasIsAutoRegistered returns a boolean if a field has been set.

### SetIsAutoRegisteredNil

`func (o *UpdateRemoteClusterParams) SetIsAutoRegisteredNil(b bool)`

 SetIsAutoRegisteredNil sets the value for IsAutoRegistered to be an explicit nil

### UnsetIsAutoRegistered
`func (o *UpdateRemoteClusterParams) UnsetIsAutoRegistered()`

UnsetIsAutoRegistered ensures that no value is present for IsAutoRegistered, not even an explicit nil
### GetLocalAddresses

`func (o *UpdateRemoteClusterParams) GetLocalAddresses() []string`

GetLocalAddresses returns the LocalAddresses field if non-nil, zero value otherwise.

### GetLocalAddressesOk

`func (o *UpdateRemoteClusterParams) GetLocalAddressesOk() (*[]string, bool)`

GetLocalAddressesOk returns a tuple with the LocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddresses

`func (o *UpdateRemoteClusterParams) SetLocalAddresses(v []string)`

SetLocalAddresses sets LocalAddresses field to given value.

### HasLocalAddresses

`func (o *UpdateRemoteClusterParams) HasLocalAddresses() bool`

HasLocalAddresses returns a boolean if a field has been set.

### GetMultiTenancyEnabled

`func (o *UpdateRemoteClusterParams) GetMultiTenancyEnabled() bool`

GetMultiTenancyEnabled returns the MultiTenancyEnabled field if non-nil, zero value otherwise.

### GetMultiTenancyEnabledOk

`func (o *UpdateRemoteClusterParams) GetMultiTenancyEnabledOk() (*bool, bool)`

GetMultiTenancyEnabledOk returns a tuple with the MultiTenancyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTenancyEnabled

`func (o *UpdateRemoteClusterParams) SetMultiTenancyEnabled(v bool)`

SetMultiTenancyEnabled sets MultiTenancyEnabled field to given value.

### HasMultiTenancyEnabled

`func (o *UpdateRemoteClusterParams) HasMultiTenancyEnabled() bool`

HasMultiTenancyEnabled returns a boolean if a field has been set.

### SetMultiTenancyEnabledNil

`func (o *UpdateRemoteClusterParams) SetMultiTenancyEnabledNil(b bool)`

 SetMultiTenancyEnabledNil sets the value for MultiTenancyEnabled to be an explicit nil

### UnsetMultiTenancyEnabled
`func (o *UpdateRemoteClusterParams) UnsetMultiTenancyEnabled()`

UnsetMultiTenancyEnabled ensures that no value is present for MultiTenancyEnabled, not even an explicit nil
### GetNetworkInterface

`func (o *UpdateRemoteClusterParams) GetNetworkInterface() string`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *UpdateRemoteClusterParams) GetNetworkInterfaceOk() (*string, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *UpdateRemoteClusterParams) SetNetworkInterface(v string)`

SetNetworkInterface sets NetworkInterface field to given value.

### HasNetworkInterface

`func (o *UpdateRemoteClusterParams) HasNetworkInterface() bool`

HasNetworkInterface returns a boolean if a field has been set.

### SetNetworkInterfaceNil

`func (o *UpdateRemoteClusterParams) SetNetworkInterfaceNil(b bool)`

 SetNetworkInterfaceNil sets the value for NetworkInterface to be an explicit nil

### UnsetNetworkInterface
`func (o *UpdateRemoteClusterParams) UnsetNetworkInterface()`

UnsetNetworkInterface ensures that no value is present for NetworkInterface, not even an explicit nil
### GetPurpose

`func (o *UpdateRemoteClusterParams) GetPurpose() []string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *UpdateRemoteClusterParams) GetPurposeOk() (*[]string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *UpdateRemoteClusterParams) SetPurpose(v []string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *UpdateRemoteClusterParams) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *UpdateRemoteClusterParams) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *UpdateRemoteClusterParams) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetReplicationParams

`func (o *UpdateRemoteClusterParams) GetReplicationParams() RemoteClusterParamsReplicationParams`

GetReplicationParams returns the ReplicationParams field if non-nil, zero value otherwise.

### GetReplicationParamsOk

`func (o *UpdateRemoteClusterParams) GetReplicationParamsOk() (*RemoteClusterParamsReplicationParams, bool)`

GetReplicationParamsOk returns a tuple with the ReplicationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationParams

`func (o *UpdateRemoteClusterParams) SetReplicationParams(v RemoteClusterParamsReplicationParams)`

SetReplicationParams sets ReplicationParams field to given value.

### HasReplicationParams

`func (o *UpdateRemoteClusterParams) HasReplicationParams() bool`

HasReplicationParams returns a boolean if a field has been set.

### GetSupportedAesEncryptionMode

`func (o *UpdateRemoteClusterParams) GetSupportedAesEncryptionMode() string`

GetSupportedAesEncryptionMode returns the SupportedAesEncryptionMode field if non-nil, zero value otherwise.

### GetSupportedAesEncryptionModeOk

`func (o *UpdateRemoteClusterParams) GetSupportedAesEncryptionModeOk() (*string, bool)`

GetSupportedAesEncryptionModeOk returns a tuple with the SupportedAesEncryptionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAesEncryptionMode

`func (o *UpdateRemoteClusterParams) SetSupportedAesEncryptionMode(v string)`

SetSupportedAesEncryptionMode sets SupportedAesEncryptionMode field to given value.

### HasSupportedAesEncryptionMode

`func (o *UpdateRemoteClusterParams) HasSupportedAesEncryptionMode() bool`

HasSupportedAesEncryptionMode returns a boolean if a field has been set.

### SetSupportedAesEncryptionModeNil

`func (o *UpdateRemoteClusterParams) SetSupportedAesEncryptionModeNil(b bool)`

 SetSupportedAesEncryptionModeNil sets the value for SupportedAesEncryptionMode to be an explicit nil

### UnsetSupportedAesEncryptionMode
`func (o *UpdateRemoteClusterParams) UnsetSupportedAesEncryptionMode()`

UnsetSupportedAesEncryptionMode ensures that no value is present for SupportedAesEncryptionMode, not even an explicit nil
### GetTenantId

`func (o *UpdateRemoteClusterParams) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UpdateRemoteClusterParams) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UpdateRemoteClusterParams) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UpdateRemoteClusterParams) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UpdateRemoteClusterParams) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UpdateRemoteClusterParams) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTenantStorageDomainSharingEnabled

`func (o *UpdateRemoteClusterParams) GetTenantStorageDomainSharingEnabled() bool`

GetTenantStorageDomainSharingEnabled returns the TenantStorageDomainSharingEnabled field if non-nil, zero value otherwise.

### GetTenantStorageDomainSharingEnabledOk

`func (o *UpdateRemoteClusterParams) GetTenantStorageDomainSharingEnabledOk() (*bool, bool)`

GetTenantStorageDomainSharingEnabledOk returns a tuple with the TenantStorageDomainSharingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantStorageDomainSharingEnabled

`func (o *UpdateRemoteClusterParams) SetTenantStorageDomainSharingEnabled(v bool)`

SetTenantStorageDomainSharingEnabled sets TenantStorageDomainSharingEnabled field to given value.

### HasTenantStorageDomainSharingEnabled

`func (o *UpdateRemoteClusterParams) HasTenantStorageDomainSharingEnabled() bool`

HasTenantStorageDomainSharingEnabled returns a boolean if a field has been set.

### SetTenantStorageDomainSharingEnabledNil

`func (o *UpdateRemoteClusterParams) SetTenantStorageDomainSharingEnabledNil(b bool)`

 SetTenantStorageDomainSharingEnabledNil sets the value for TenantStorageDomainSharingEnabled to be an explicit nil

### UnsetTenantStorageDomainSharingEnabled
`func (o *UpdateRemoteClusterParams) UnsetTenantStorageDomainSharingEnabled()`

UnsetTenantStorageDomainSharingEnabled ensures that no value is present for TenantStorageDomainSharingEnabled, not even an explicit nil
### GetTlsEnabled

`func (o *UpdateRemoteClusterParams) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *UpdateRemoteClusterParams) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *UpdateRemoteClusterParams) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *UpdateRemoteClusterParams) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### SetTlsEnabledNil

`func (o *UpdateRemoteClusterParams) SetTlsEnabledNil(b bool)`

 SetTlsEnabledNil sets the value for TlsEnabled to be an explicit nil

### UnsetTlsEnabled
`func (o *UpdateRemoteClusterParams) UnsetTlsEnabled()`

UnsetTlsEnabled ensures that no value is present for TlsEnabled, not even an explicit nil
### GetNodeAddresses

`func (o *UpdateRemoteClusterParams) GetNodeAddresses() []string`

GetNodeAddresses returns the NodeAddresses field if non-nil, zero value otherwise.

### GetNodeAddressesOk

`func (o *UpdateRemoteClusterParams) GetNodeAddressesOk() (*[]string, bool)`

GetNodeAddressesOk returns a tuple with the NodeAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddresses

`func (o *UpdateRemoteClusterParams) SetNodeAddresses(v []string)`

SetNodeAddresses sets NodeAddresses field to given value.

### HasNodeAddresses

`func (o *UpdateRemoteClusterParams) HasNodeAddresses() bool`

HasNodeAddresses returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateRemoteClusterParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateRemoteClusterParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateRemoteClusterParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateRemoteClusterParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateRemoteClusterParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateRemoteClusterParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetUsername

`func (o *UpdateRemoteClusterParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateRemoteClusterParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateRemoteClusterParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateRemoteClusterParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


