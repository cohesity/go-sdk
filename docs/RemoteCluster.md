# RemoteCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllEndpointsReachable** | Pointer to **NullableBool** | Specifies whether any endpoint (such as a Node) on the remote Cluster is reachable from this local Cluster. If true, a service running on the local Cluster can communicate directly with any of its peers running on the remote Cluster, without using a proxy. | [optional] 
**AutoRegisterTarget** | Pointer to **NullableBool** | Specifies whether the remote cluster needs to be kept in sync. This will be set to true by default. | [optional] 
**AutoRegistration** | Pointer to **NullableBool** | Specifies whether the remote registration has happened automatically (due to registration on the other site). Can&#39;t think of other states (other than manually &amp; automatically) so this isn&#39;t an enum. For a manual registration, this field will not be set. | [optional] 
**BandwidthLimit** | Pointer to [**BandwidthLimit**](BandwidthLimit.md) |  | [optional] 
**ClusterId** | Pointer to **NullableInt64** | Specifies the unique id of the remote Cluster. | [optional] 
**ClusterIncarnationId** | Pointer to **NullableInt64** | Specifies the unique incarnation id of the remote Cluster. This id is determined dynamically by contacting the remote Cluster. | [optional] 
**CompressionEnabled** | Pointer to **NullableBool** | Specifies whether to compress the outbound data when transferring the replication data over the network to the remote Cluster. | [optional] 
**Description** | Pointer to **NullableString** | Specifies any additional information if needed. | [optional] 
**EncryptionKey** | Pointer to **NullableString** | Specifies the encryption key used for encrypting the replication data from a local Cluster to a remote Cluster. If a key is not specified, replication traffic encryption is disabled. When Snapshots are replicated from a local Cluster to a remote Cluster, the encryption key specified on the local Cluster must be the same as the key specified on the remote Cluster. | [optional] 
**LocalIps** | Pointer to **[]string** | Array of Local IP Addresses.  Specifies the IP addresses of the interfaces in the local Cluster which will be used for communicating with the remote Cluster. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the remote cluster. This field is determined dynamically by contacting the remote cluster. | [optional] 
**NetworkInterface** | Pointer to **NullableString** | Specifies the name of the network interfaces to use for communicating with the remote Cluster. | [optional] 
**PurposeRemoteAccess** | Pointer to **NullableBool** | Whether the remote cluster will be used for remote access for SPOG. | [optional] 
**PurposeReplication** | Pointer to **NullableBool** | Whether the remote cluster will be used for replication. | [optional] 
**RemoteAccessCredentials** | Pointer to [**AccessTokenCredential**](AccessTokenCredential.md) |  | [optional] 
**RemoteIps** | Pointer to **[]string** | Array of Remote Node IP Addresses.  Specifies the IP addresses of the Nodes on the remote Cluster to connect with. These IP addresses can also be VIPS. Specifying hostnames is not supported. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the tenant Id of the organization that created this remote cluster configuration. | [optional] 
**UserName** | Pointer to **NullableString** | Specifies the Cohesity user name used to connect to the remote Cluster. | [optional] 
**ViewBoxPairInfo** | Pointer to [**[]ViewBoxPairInfo**](ViewBoxPairInfo.md) | Array of Storage Domain (View Box) Pairs.  Specifies pairings between Storage Domains (View Boxes) on the local Cluster with Storage Domains (View Boxes) on a remote Cluster that are used in replication. | [optional] 

## Methods

### NewRemoteCluster

`func NewRemoteCluster() *RemoteCluster`

NewRemoteCluster instantiates a new RemoteCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteClusterWithDefaults

`func NewRemoteClusterWithDefaults() *RemoteCluster`

NewRemoteClusterWithDefaults instantiates a new RemoteCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllEndpointsReachable

`func (o *RemoteCluster) GetAllEndpointsReachable() bool`

GetAllEndpointsReachable returns the AllEndpointsReachable field if non-nil, zero value otherwise.

### GetAllEndpointsReachableOk

`func (o *RemoteCluster) GetAllEndpointsReachableOk() (*bool, bool)`

GetAllEndpointsReachableOk returns a tuple with the AllEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEndpointsReachable

`func (o *RemoteCluster) SetAllEndpointsReachable(v bool)`

SetAllEndpointsReachable sets AllEndpointsReachable field to given value.

### HasAllEndpointsReachable

`func (o *RemoteCluster) HasAllEndpointsReachable() bool`

HasAllEndpointsReachable returns a boolean if a field has been set.

### SetAllEndpointsReachableNil

`func (o *RemoteCluster) SetAllEndpointsReachableNil(b bool)`

 SetAllEndpointsReachableNil sets the value for AllEndpointsReachable to be an explicit nil

### UnsetAllEndpointsReachable
`func (o *RemoteCluster) UnsetAllEndpointsReachable()`

UnsetAllEndpointsReachable ensures that no value is present for AllEndpointsReachable, not even an explicit nil
### GetAutoRegisterTarget

`func (o *RemoteCluster) GetAutoRegisterTarget() bool`

GetAutoRegisterTarget returns the AutoRegisterTarget field if non-nil, zero value otherwise.

### GetAutoRegisterTargetOk

`func (o *RemoteCluster) GetAutoRegisterTargetOk() (*bool, bool)`

GetAutoRegisterTargetOk returns a tuple with the AutoRegisterTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegisterTarget

`func (o *RemoteCluster) SetAutoRegisterTarget(v bool)`

SetAutoRegisterTarget sets AutoRegisterTarget field to given value.

### HasAutoRegisterTarget

`func (o *RemoteCluster) HasAutoRegisterTarget() bool`

HasAutoRegisterTarget returns a boolean if a field has been set.

### SetAutoRegisterTargetNil

`func (o *RemoteCluster) SetAutoRegisterTargetNil(b bool)`

 SetAutoRegisterTargetNil sets the value for AutoRegisterTarget to be an explicit nil

### UnsetAutoRegisterTarget
`func (o *RemoteCluster) UnsetAutoRegisterTarget()`

UnsetAutoRegisterTarget ensures that no value is present for AutoRegisterTarget, not even an explicit nil
### GetAutoRegistration

`func (o *RemoteCluster) GetAutoRegistration() bool`

GetAutoRegistration returns the AutoRegistration field if non-nil, zero value otherwise.

### GetAutoRegistrationOk

`func (o *RemoteCluster) GetAutoRegistrationOk() (*bool, bool)`

GetAutoRegistrationOk returns a tuple with the AutoRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegistration

`func (o *RemoteCluster) SetAutoRegistration(v bool)`

SetAutoRegistration sets AutoRegistration field to given value.

### HasAutoRegistration

`func (o *RemoteCluster) HasAutoRegistration() bool`

HasAutoRegistration returns a boolean if a field has been set.

### SetAutoRegistrationNil

`func (o *RemoteCluster) SetAutoRegistrationNil(b bool)`

 SetAutoRegistrationNil sets the value for AutoRegistration to be an explicit nil

### UnsetAutoRegistration
`func (o *RemoteCluster) UnsetAutoRegistration()`

UnsetAutoRegistration ensures that no value is present for AutoRegistration, not even an explicit nil
### GetBandwidthLimit

`func (o *RemoteCluster) GetBandwidthLimit() BandwidthLimit`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *RemoteCluster) GetBandwidthLimitOk() (*BandwidthLimit, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *RemoteCluster) SetBandwidthLimit(v BandwidthLimit)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *RemoteCluster) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.

### GetClusterId

`func (o *RemoteCluster) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RemoteCluster) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RemoteCluster) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *RemoteCluster) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *RemoteCluster) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *RemoteCluster) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *RemoteCluster) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *RemoteCluster) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *RemoteCluster) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *RemoteCluster) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *RemoteCluster) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *RemoteCluster) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetCompressionEnabled

`func (o *RemoteCluster) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *RemoteCluster) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *RemoteCluster) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *RemoteCluster) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### SetCompressionEnabledNil

`func (o *RemoteCluster) SetCompressionEnabledNil(b bool)`

 SetCompressionEnabledNil sets the value for CompressionEnabled to be an explicit nil

### UnsetCompressionEnabled
`func (o *RemoteCluster) UnsetCompressionEnabled()`

UnsetCompressionEnabled ensures that no value is present for CompressionEnabled, not even an explicit nil
### GetDescription

`func (o *RemoteCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RemoteCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RemoteCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RemoteCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RemoteCluster) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RemoteCluster) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEncryptionKey

`func (o *RemoteCluster) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *RemoteCluster) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *RemoteCluster) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *RemoteCluster) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *RemoteCluster) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *RemoteCluster) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetLocalIps

`func (o *RemoteCluster) GetLocalIps() []string`

GetLocalIps returns the LocalIps field if non-nil, zero value otherwise.

### GetLocalIpsOk

`func (o *RemoteCluster) GetLocalIpsOk() (*[]string, bool)`

GetLocalIpsOk returns a tuple with the LocalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIps

`func (o *RemoteCluster) SetLocalIps(v []string)`

SetLocalIps sets LocalIps field to given value.

### HasLocalIps

`func (o *RemoteCluster) HasLocalIps() bool`

HasLocalIps returns a boolean if a field has been set.

### SetLocalIpsNil

`func (o *RemoteCluster) SetLocalIpsNil(b bool)`

 SetLocalIpsNil sets the value for LocalIps to be an explicit nil

### UnsetLocalIps
`func (o *RemoteCluster) UnsetLocalIps()`

UnsetLocalIps ensures that no value is present for LocalIps, not even an explicit nil
### GetName

`func (o *RemoteCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RemoteCluster) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RemoteCluster) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetworkInterface

`func (o *RemoteCluster) GetNetworkInterface() string`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *RemoteCluster) GetNetworkInterfaceOk() (*string, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *RemoteCluster) SetNetworkInterface(v string)`

SetNetworkInterface sets NetworkInterface field to given value.

### HasNetworkInterface

`func (o *RemoteCluster) HasNetworkInterface() bool`

HasNetworkInterface returns a boolean if a field has been set.

### SetNetworkInterfaceNil

`func (o *RemoteCluster) SetNetworkInterfaceNil(b bool)`

 SetNetworkInterfaceNil sets the value for NetworkInterface to be an explicit nil

### UnsetNetworkInterface
`func (o *RemoteCluster) UnsetNetworkInterface()`

UnsetNetworkInterface ensures that no value is present for NetworkInterface, not even an explicit nil
### GetPurposeRemoteAccess

`func (o *RemoteCluster) GetPurposeRemoteAccess() bool`

GetPurposeRemoteAccess returns the PurposeRemoteAccess field if non-nil, zero value otherwise.

### GetPurposeRemoteAccessOk

`func (o *RemoteCluster) GetPurposeRemoteAccessOk() (*bool, bool)`

GetPurposeRemoteAccessOk returns a tuple with the PurposeRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeRemoteAccess

`func (o *RemoteCluster) SetPurposeRemoteAccess(v bool)`

SetPurposeRemoteAccess sets PurposeRemoteAccess field to given value.

### HasPurposeRemoteAccess

`func (o *RemoteCluster) HasPurposeRemoteAccess() bool`

HasPurposeRemoteAccess returns a boolean if a field has been set.

### SetPurposeRemoteAccessNil

`func (o *RemoteCluster) SetPurposeRemoteAccessNil(b bool)`

 SetPurposeRemoteAccessNil sets the value for PurposeRemoteAccess to be an explicit nil

### UnsetPurposeRemoteAccess
`func (o *RemoteCluster) UnsetPurposeRemoteAccess()`

UnsetPurposeRemoteAccess ensures that no value is present for PurposeRemoteAccess, not even an explicit nil
### GetPurposeReplication

`func (o *RemoteCluster) GetPurposeReplication() bool`

GetPurposeReplication returns the PurposeReplication field if non-nil, zero value otherwise.

### GetPurposeReplicationOk

`func (o *RemoteCluster) GetPurposeReplicationOk() (*bool, bool)`

GetPurposeReplicationOk returns a tuple with the PurposeReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeReplication

`func (o *RemoteCluster) SetPurposeReplication(v bool)`

SetPurposeReplication sets PurposeReplication field to given value.

### HasPurposeReplication

`func (o *RemoteCluster) HasPurposeReplication() bool`

HasPurposeReplication returns a boolean if a field has been set.

### SetPurposeReplicationNil

`func (o *RemoteCluster) SetPurposeReplicationNil(b bool)`

 SetPurposeReplicationNil sets the value for PurposeReplication to be an explicit nil

### UnsetPurposeReplication
`func (o *RemoteCluster) UnsetPurposeReplication()`

UnsetPurposeReplication ensures that no value is present for PurposeReplication, not even an explicit nil
### GetRemoteAccessCredentials

`func (o *RemoteCluster) GetRemoteAccessCredentials() AccessTokenCredential`

GetRemoteAccessCredentials returns the RemoteAccessCredentials field if non-nil, zero value otherwise.

### GetRemoteAccessCredentialsOk

`func (o *RemoteCluster) GetRemoteAccessCredentialsOk() (*AccessTokenCredential, bool)`

GetRemoteAccessCredentialsOk returns a tuple with the RemoteAccessCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessCredentials

`func (o *RemoteCluster) SetRemoteAccessCredentials(v AccessTokenCredential)`

SetRemoteAccessCredentials sets RemoteAccessCredentials field to given value.

### HasRemoteAccessCredentials

`func (o *RemoteCluster) HasRemoteAccessCredentials() bool`

HasRemoteAccessCredentials returns a boolean if a field has been set.

### GetRemoteIps

`func (o *RemoteCluster) GetRemoteIps() []string`

GetRemoteIps returns the RemoteIps field if non-nil, zero value otherwise.

### GetRemoteIpsOk

`func (o *RemoteCluster) GetRemoteIpsOk() (*[]string, bool)`

GetRemoteIpsOk returns a tuple with the RemoteIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIps

`func (o *RemoteCluster) SetRemoteIps(v []string)`

SetRemoteIps sets RemoteIps field to given value.

### HasRemoteIps

`func (o *RemoteCluster) HasRemoteIps() bool`

HasRemoteIps returns a boolean if a field has been set.

### SetRemoteIpsNil

`func (o *RemoteCluster) SetRemoteIpsNil(b bool)`

 SetRemoteIpsNil sets the value for RemoteIps to be an explicit nil

### UnsetRemoteIps
`func (o *RemoteCluster) UnsetRemoteIps()`

UnsetRemoteIps ensures that no value is present for RemoteIps, not even an explicit nil
### GetTenantId

`func (o *RemoteCluster) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RemoteCluster) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RemoteCluster) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *RemoteCluster) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *RemoteCluster) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *RemoteCluster) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUserName

`func (o *RemoteCluster) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *RemoteCluster) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *RemoteCluster) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *RemoteCluster) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *RemoteCluster) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *RemoteCluster) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetViewBoxPairInfo

`func (o *RemoteCluster) GetViewBoxPairInfo() []ViewBoxPairInfo`

GetViewBoxPairInfo returns the ViewBoxPairInfo field if non-nil, zero value otherwise.

### GetViewBoxPairInfoOk

`func (o *RemoteCluster) GetViewBoxPairInfoOk() (*[]ViewBoxPairInfo, bool)`

GetViewBoxPairInfoOk returns a tuple with the ViewBoxPairInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxPairInfo

`func (o *RemoteCluster) SetViewBoxPairInfo(v []ViewBoxPairInfo)`

SetViewBoxPairInfo sets ViewBoxPairInfo field to given value.

### HasViewBoxPairInfo

`func (o *RemoteCluster) HasViewBoxPairInfo() bool`

HasViewBoxPairInfo returns a boolean if a field has been set.

### SetViewBoxPairInfoNil

`func (o *RemoteCluster) SetViewBoxPairInfoNil(b bool)`

 SetViewBoxPairInfoNil sets the value for ViewBoxPairInfo to be an explicit nil

### UnsetViewBoxPairInfo
`func (o *RemoteCluster) UnsetViewBoxPairInfo()`

UnsetViewBoxPairInfo ensures that no value is present for ViewBoxPairInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


