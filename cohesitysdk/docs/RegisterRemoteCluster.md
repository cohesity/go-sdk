# RegisterRemoteCluster

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
**Name** | Pointer to **NullableString** | Specifies the name of the remote cluster. This field is determined dynamically by contacting the remote cluster. | [optional] 
**NetworkInterface** | Pointer to **NullableString** | Specifies the name of the network interfaces to use for communicating with the remote Cluster. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password for Cohesity user to use when connecting to the remote Cluster. | [optional] 
**PurposeRemoteAccess** | Pointer to **NullableBool** | Whether the remote cluster will be used for remote access for SPOG. | [optional] 
**PurposeReplication** | Pointer to **NullableBool** | Whether the remote cluster will be used for replication. | [optional] 
**RemoteAccessCredentials** | Pointer to [**AccessTokenCredential**](AccessTokenCredential.md) |  | [optional] 
**RemoteIps** | Pointer to **[]string** | Array of Remote Node IP Addresses.  Specifies the IP addresses of the Nodes on the remote Cluster to connect with. These IP addresses can also be VIPS. Specifying hostnames is not supported. | [optional] 
**RemoteIrisPorts** | Pointer to **[]int64** | Array of Ports.  Specifies the ports to use when connecting to the Nodes of the remote Cluster. | [optional] 
**UserName** | Pointer to **NullableString** | Specifies the Cohesity user name used to connect to the remote Cluster. | [optional] 
**ValidateOnly** | Pointer to **NullableBool** | Whether to only validate the credentials without saving the information. | [optional] 
**ViewBoxPairInfo** | Pointer to [**[]ViewBoxPairInfo**](ViewBoxPairInfo.md) | Array of Storage Domain (View Box) Pairs.  Specifies pairings between Storage Domains (View Boxes) on the local Cluster with Storage Domains (View Boxes) on a remote Cluster that are used in replication. | [optional] 

## Methods

### NewRegisterRemoteCluster

`func NewRegisterRemoteCluster() *RegisterRemoteCluster`

NewRegisterRemoteCluster instantiates a new RegisterRemoteCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRemoteClusterWithDefaults

`func NewRegisterRemoteClusterWithDefaults() *RegisterRemoteCluster`

NewRegisterRemoteClusterWithDefaults instantiates a new RegisterRemoteCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllEndpointsReachable

`func (o *RegisterRemoteCluster) GetAllEndpointsReachable() bool`

GetAllEndpointsReachable returns the AllEndpointsReachable field if non-nil, zero value otherwise.

### GetAllEndpointsReachableOk

`func (o *RegisterRemoteCluster) GetAllEndpointsReachableOk() (*bool, bool)`

GetAllEndpointsReachableOk returns a tuple with the AllEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEndpointsReachable

`func (o *RegisterRemoteCluster) SetAllEndpointsReachable(v bool)`

SetAllEndpointsReachable sets AllEndpointsReachable field to given value.

### HasAllEndpointsReachable

`func (o *RegisterRemoteCluster) HasAllEndpointsReachable() bool`

HasAllEndpointsReachable returns a boolean if a field has been set.

### SetAllEndpointsReachableNil

`func (o *RegisterRemoteCluster) SetAllEndpointsReachableNil(b bool)`

 SetAllEndpointsReachableNil sets the value for AllEndpointsReachable to be an explicit nil

### UnsetAllEndpointsReachable
`func (o *RegisterRemoteCluster) UnsetAllEndpointsReachable()`

UnsetAllEndpointsReachable ensures that no value is present for AllEndpointsReachable, not even an explicit nil
### GetAutoRegisterTarget

`func (o *RegisterRemoteCluster) GetAutoRegisterTarget() bool`

GetAutoRegisterTarget returns the AutoRegisterTarget field if non-nil, zero value otherwise.

### GetAutoRegisterTargetOk

`func (o *RegisterRemoteCluster) GetAutoRegisterTargetOk() (*bool, bool)`

GetAutoRegisterTargetOk returns a tuple with the AutoRegisterTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegisterTarget

`func (o *RegisterRemoteCluster) SetAutoRegisterTarget(v bool)`

SetAutoRegisterTarget sets AutoRegisterTarget field to given value.

### HasAutoRegisterTarget

`func (o *RegisterRemoteCluster) HasAutoRegisterTarget() bool`

HasAutoRegisterTarget returns a boolean if a field has been set.

### SetAutoRegisterTargetNil

`func (o *RegisterRemoteCluster) SetAutoRegisterTargetNil(b bool)`

 SetAutoRegisterTargetNil sets the value for AutoRegisterTarget to be an explicit nil

### UnsetAutoRegisterTarget
`func (o *RegisterRemoteCluster) UnsetAutoRegisterTarget()`

UnsetAutoRegisterTarget ensures that no value is present for AutoRegisterTarget, not even an explicit nil
### GetAutoRegistration

`func (o *RegisterRemoteCluster) GetAutoRegistration() bool`

GetAutoRegistration returns the AutoRegistration field if non-nil, zero value otherwise.

### GetAutoRegistrationOk

`func (o *RegisterRemoteCluster) GetAutoRegistrationOk() (*bool, bool)`

GetAutoRegistrationOk returns a tuple with the AutoRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegistration

`func (o *RegisterRemoteCluster) SetAutoRegistration(v bool)`

SetAutoRegistration sets AutoRegistration field to given value.

### HasAutoRegistration

`func (o *RegisterRemoteCluster) HasAutoRegistration() bool`

HasAutoRegistration returns a boolean if a field has been set.

### SetAutoRegistrationNil

`func (o *RegisterRemoteCluster) SetAutoRegistrationNil(b bool)`

 SetAutoRegistrationNil sets the value for AutoRegistration to be an explicit nil

### UnsetAutoRegistration
`func (o *RegisterRemoteCluster) UnsetAutoRegistration()`

UnsetAutoRegistration ensures that no value is present for AutoRegistration, not even an explicit nil
### GetBandwidthLimit

`func (o *RegisterRemoteCluster) GetBandwidthLimit() BandwidthLimit`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *RegisterRemoteCluster) GetBandwidthLimitOk() (*BandwidthLimit, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *RegisterRemoteCluster) SetBandwidthLimit(v BandwidthLimit)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *RegisterRemoteCluster) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.

### GetClusterId

`func (o *RegisterRemoteCluster) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RegisterRemoteCluster) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RegisterRemoteCluster) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *RegisterRemoteCluster) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *RegisterRemoteCluster) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *RegisterRemoteCluster) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterIncarnationId

`func (o *RegisterRemoteCluster) GetClusterIncarnationId() int64`

GetClusterIncarnationId returns the ClusterIncarnationId field if non-nil, zero value otherwise.

### GetClusterIncarnationIdOk

`func (o *RegisterRemoteCluster) GetClusterIncarnationIdOk() (*int64, bool)`

GetClusterIncarnationIdOk returns a tuple with the ClusterIncarnationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIncarnationId

`func (o *RegisterRemoteCluster) SetClusterIncarnationId(v int64)`

SetClusterIncarnationId sets ClusterIncarnationId field to given value.

### HasClusterIncarnationId

`func (o *RegisterRemoteCluster) HasClusterIncarnationId() bool`

HasClusterIncarnationId returns a boolean if a field has been set.

### SetClusterIncarnationIdNil

`func (o *RegisterRemoteCluster) SetClusterIncarnationIdNil(b bool)`

 SetClusterIncarnationIdNil sets the value for ClusterIncarnationId to be an explicit nil

### UnsetClusterIncarnationId
`func (o *RegisterRemoteCluster) UnsetClusterIncarnationId()`

UnsetClusterIncarnationId ensures that no value is present for ClusterIncarnationId, not even an explicit nil
### GetCompressionEnabled

`func (o *RegisterRemoteCluster) GetCompressionEnabled() bool`

GetCompressionEnabled returns the CompressionEnabled field if non-nil, zero value otherwise.

### GetCompressionEnabledOk

`func (o *RegisterRemoteCluster) GetCompressionEnabledOk() (*bool, bool)`

GetCompressionEnabledOk returns a tuple with the CompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEnabled

`func (o *RegisterRemoteCluster) SetCompressionEnabled(v bool)`

SetCompressionEnabled sets CompressionEnabled field to given value.

### HasCompressionEnabled

`func (o *RegisterRemoteCluster) HasCompressionEnabled() bool`

HasCompressionEnabled returns a boolean if a field has been set.

### SetCompressionEnabledNil

`func (o *RegisterRemoteCluster) SetCompressionEnabledNil(b bool)`

 SetCompressionEnabledNil sets the value for CompressionEnabled to be an explicit nil

### UnsetCompressionEnabled
`func (o *RegisterRemoteCluster) UnsetCompressionEnabled()`

UnsetCompressionEnabled ensures that no value is present for CompressionEnabled, not even an explicit nil
### GetDescription

`func (o *RegisterRemoteCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegisterRemoteCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegisterRemoteCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegisterRemoteCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RegisterRemoteCluster) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RegisterRemoteCluster) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEncryptionKey

`func (o *RegisterRemoteCluster) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *RegisterRemoteCluster) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *RegisterRemoteCluster) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *RegisterRemoteCluster) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### SetEncryptionKeyNil

`func (o *RegisterRemoteCluster) SetEncryptionKeyNil(b bool)`

 SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil

### UnsetEncryptionKey
`func (o *RegisterRemoteCluster) UnsetEncryptionKey()`

UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
### GetName

`func (o *RegisterRemoteCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterRemoteCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterRemoteCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegisterRemoteCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RegisterRemoteCluster) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RegisterRemoteCluster) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetworkInterface

`func (o *RegisterRemoteCluster) GetNetworkInterface() string`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *RegisterRemoteCluster) GetNetworkInterfaceOk() (*string, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *RegisterRemoteCluster) SetNetworkInterface(v string)`

SetNetworkInterface sets NetworkInterface field to given value.

### HasNetworkInterface

`func (o *RegisterRemoteCluster) HasNetworkInterface() bool`

HasNetworkInterface returns a boolean if a field has been set.

### SetNetworkInterfaceNil

`func (o *RegisterRemoteCluster) SetNetworkInterfaceNil(b bool)`

 SetNetworkInterfaceNil sets the value for NetworkInterface to be an explicit nil

### UnsetNetworkInterface
`func (o *RegisterRemoteCluster) UnsetNetworkInterface()`

UnsetNetworkInterface ensures that no value is present for NetworkInterface, not even an explicit nil
### GetPassword

`func (o *RegisterRemoteCluster) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterRemoteCluster) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterRemoteCluster) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegisterRemoteCluster) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RegisterRemoteCluster) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RegisterRemoteCluster) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPurposeRemoteAccess

`func (o *RegisterRemoteCluster) GetPurposeRemoteAccess() bool`

GetPurposeRemoteAccess returns the PurposeRemoteAccess field if non-nil, zero value otherwise.

### GetPurposeRemoteAccessOk

`func (o *RegisterRemoteCluster) GetPurposeRemoteAccessOk() (*bool, bool)`

GetPurposeRemoteAccessOk returns a tuple with the PurposeRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeRemoteAccess

`func (o *RegisterRemoteCluster) SetPurposeRemoteAccess(v bool)`

SetPurposeRemoteAccess sets PurposeRemoteAccess field to given value.

### HasPurposeRemoteAccess

`func (o *RegisterRemoteCluster) HasPurposeRemoteAccess() bool`

HasPurposeRemoteAccess returns a boolean if a field has been set.

### SetPurposeRemoteAccessNil

`func (o *RegisterRemoteCluster) SetPurposeRemoteAccessNil(b bool)`

 SetPurposeRemoteAccessNil sets the value for PurposeRemoteAccess to be an explicit nil

### UnsetPurposeRemoteAccess
`func (o *RegisterRemoteCluster) UnsetPurposeRemoteAccess()`

UnsetPurposeRemoteAccess ensures that no value is present for PurposeRemoteAccess, not even an explicit nil
### GetPurposeReplication

`func (o *RegisterRemoteCluster) GetPurposeReplication() bool`

GetPurposeReplication returns the PurposeReplication field if non-nil, zero value otherwise.

### GetPurposeReplicationOk

`func (o *RegisterRemoteCluster) GetPurposeReplicationOk() (*bool, bool)`

GetPurposeReplicationOk returns a tuple with the PurposeReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeReplication

`func (o *RegisterRemoteCluster) SetPurposeReplication(v bool)`

SetPurposeReplication sets PurposeReplication field to given value.

### HasPurposeReplication

`func (o *RegisterRemoteCluster) HasPurposeReplication() bool`

HasPurposeReplication returns a boolean if a field has been set.

### SetPurposeReplicationNil

`func (o *RegisterRemoteCluster) SetPurposeReplicationNil(b bool)`

 SetPurposeReplicationNil sets the value for PurposeReplication to be an explicit nil

### UnsetPurposeReplication
`func (o *RegisterRemoteCluster) UnsetPurposeReplication()`

UnsetPurposeReplication ensures that no value is present for PurposeReplication, not even an explicit nil
### GetRemoteAccessCredentials

`func (o *RegisterRemoteCluster) GetRemoteAccessCredentials() AccessTokenCredential`

GetRemoteAccessCredentials returns the RemoteAccessCredentials field if non-nil, zero value otherwise.

### GetRemoteAccessCredentialsOk

`func (o *RegisterRemoteCluster) GetRemoteAccessCredentialsOk() (*AccessTokenCredential, bool)`

GetRemoteAccessCredentialsOk returns a tuple with the RemoteAccessCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAccessCredentials

`func (o *RegisterRemoteCluster) SetRemoteAccessCredentials(v AccessTokenCredential)`

SetRemoteAccessCredentials sets RemoteAccessCredentials field to given value.

### HasRemoteAccessCredentials

`func (o *RegisterRemoteCluster) HasRemoteAccessCredentials() bool`

HasRemoteAccessCredentials returns a boolean if a field has been set.

### GetRemoteIps

`func (o *RegisterRemoteCluster) GetRemoteIps() []string`

GetRemoteIps returns the RemoteIps field if non-nil, zero value otherwise.

### GetRemoteIpsOk

`func (o *RegisterRemoteCluster) GetRemoteIpsOk() (*[]string, bool)`

GetRemoteIpsOk returns a tuple with the RemoteIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIps

`func (o *RegisterRemoteCluster) SetRemoteIps(v []string)`

SetRemoteIps sets RemoteIps field to given value.

### HasRemoteIps

`func (o *RegisterRemoteCluster) HasRemoteIps() bool`

HasRemoteIps returns a boolean if a field has been set.

### SetRemoteIpsNil

`func (o *RegisterRemoteCluster) SetRemoteIpsNil(b bool)`

 SetRemoteIpsNil sets the value for RemoteIps to be an explicit nil

### UnsetRemoteIps
`func (o *RegisterRemoteCluster) UnsetRemoteIps()`

UnsetRemoteIps ensures that no value is present for RemoteIps, not even an explicit nil
### GetRemoteIrisPorts

`func (o *RegisterRemoteCluster) GetRemoteIrisPorts() []int64`

GetRemoteIrisPorts returns the RemoteIrisPorts field if non-nil, zero value otherwise.

### GetRemoteIrisPortsOk

`func (o *RegisterRemoteCluster) GetRemoteIrisPortsOk() (*[]int64, bool)`

GetRemoteIrisPortsOk returns a tuple with the RemoteIrisPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIrisPorts

`func (o *RegisterRemoteCluster) SetRemoteIrisPorts(v []int64)`

SetRemoteIrisPorts sets RemoteIrisPorts field to given value.

### HasRemoteIrisPorts

`func (o *RegisterRemoteCluster) HasRemoteIrisPorts() bool`

HasRemoteIrisPorts returns a boolean if a field has been set.

### SetRemoteIrisPortsNil

`func (o *RegisterRemoteCluster) SetRemoteIrisPortsNil(b bool)`

 SetRemoteIrisPortsNil sets the value for RemoteIrisPorts to be an explicit nil

### UnsetRemoteIrisPorts
`func (o *RegisterRemoteCluster) UnsetRemoteIrisPorts()`

UnsetRemoteIrisPorts ensures that no value is present for RemoteIrisPorts, not even an explicit nil
### GetUserName

`func (o *RegisterRemoteCluster) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *RegisterRemoteCluster) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *RegisterRemoteCluster) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *RegisterRemoteCluster) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *RegisterRemoteCluster) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *RegisterRemoteCluster) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetValidateOnly

`func (o *RegisterRemoteCluster) GetValidateOnly() bool`

GetValidateOnly returns the ValidateOnly field if non-nil, zero value otherwise.

### GetValidateOnlyOk

`func (o *RegisterRemoteCluster) GetValidateOnlyOk() (*bool, bool)`

GetValidateOnlyOk returns a tuple with the ValidateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOnly

`func (o *RegisterRemoteCluster) SetValidateOnly(v bool)`

SetValidateOnly sets ValidateOnly field to given value.

### HasValidateOnly

`func (o *RegisterRemoteCluster) HasValidateOnly() bool`

HasValidateOnly returns a boolean if a field has been set.

### SetValidateOnlyNil

`func (o *RegisterRemoteCluster) SetValidateOnlyNil(b bool)`

 SetValidateOnlyNil sets the value for ValidateOnly to be an explicit nil

### UnsetValidateOnly
`func (o *RegisterRemoteCluster) UnsetValidateOnly()`

UnsetValidateOnly ensures that no value is present for ValidateOnly, not even an explicit nil
### GetViewBoxPairInfo

`func (o *RegisterRemoteCluster) GetViewBoxPairInfo() []ViewBoxPairInfo`

GetViewBoxPairInfo returns the ViewBoxPairInfo field if non-nil, zero value otherwise.

### GetViewBoxPairInfoOk

`func (o *RegisterRemoteCluster) GetViewBoxPairInfoOk() (*[]ViewBoxPairInfo, bool)`

GetViewBoxPairInfoOk returns a tuple with the ViewBoxPairInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxPairInfo

`func (o *RegisterRemoteCluster) SetViewBoxPairInfo(v []ViewBoxPairInfo)`

SetViewBoxPairInfo sets ViewBoxPairInfo field to given value.

### HasViewBoxPairInfo

`func (o *RegisterRemoteCluster) HasViewBoxPairInfo() bool`

HasViewBoxPairInfo returns a boolean if a field has been set.

### SetViewBoxPairInfoNil

`func (o *RegisterRemoteCluster) SetViewBoxPairInfoNil(b bool)`

 SetViewBoxPairInfoNil sets the value for ViewBoxPairInfo to be an explicit nil

### UnsetViewBoxPairInfo
`func (o *RegisterRemoteCluster) UnsetViewBoxPairInfo()`

UnsetViewBoxPairInfo ensures that no value is present for ViewBoxPairInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


