# KvmProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentError** | Pointer to **NullableString** | Specifies a message when the agent cannot be reached. | [optional] 
**AgentId** | Pointer to **NullableInt64** | Specifies the ID of the Agent with which this KVM entity is associated when the entity represents a Delegate host or KVM host. | [optional] 
**ClusterId** | Pointer to **NullableString** | Specifies the cluster ID for &#39;kCluster&#39; objects. | [optional] 
**DatacenterId** | Pointer to **NullableString** | Specifies the ID of the &#39;kDatacenter&#39; objects. | [optional] 
**Description** | Pointer to **NullableString** | Specifies a description about the Protection Source. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the KVM entity. | [optional] 
**NetworkId** | Pointer to **NullableString** | Specifies the network ID to which Vnic is attached. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the type of KVM Protection Source entities such as &#39;kDatacenter&#39;, &#39;kCluster&#39;, &#39;kVirtualMachine&#39;, etc. Specifies the type of an KVM source entity. &#39;kOVirtManager&#39; indicates the root entity registered with Cohesity cluster. &#39;kStandaloneHost&#39; indicates a host registered with Cohesity cluster. &#39;kDatacenter&#39; indicates a KVM datacenter managed by the OVirt manager. &#39;kCluster&#39; indicates a KVM cluster managed by the OVirt manager. &#39;kHost&#39; indicates a host within the KVM environment. &#39;kVirtualMachine&#39; indicates a virtual machine in the KVM enironment. &#39;kNetwork&#39; represents a network used by the virtual machine entity. &#39;kStorageDomain&#39; represents a storage domain in the KVM environment. &#39;kVNicProfile&#39; represents a VNic profile. | [optional] 
**Uuid** | Pointer to **NullableString** | Specifies the UUID of the Object. This is unique within the KVM environment. | [optional] 

## Methods

### NewKvmProtectionSource

`func NewKvmProtectionSource() *KvmProtectionSource`

NewKvmProtectionSource instantiates a new KvmProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmProtectionSourceWithDefaults

`func NewKvmProtectionSourceWithDefaults() *KvmProtectionSource`

NewKvmProtectionSourceWithDefaults instantiates a new KvmProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentError

`func (o *KvmProtectionSource) GetAgentError() string`

GetAgentError returns the AgentError field if non-nil, zero value otherwise.

### GetAgentErrorOk

`func (o *KvmProtectionSource) GetAgentErrorOk() (*string, bool)`

GetAgentErrorOk returns a tuple with the AgentError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentError

`func (o *KvmProtectionSource) SetAgentError(v string)`

SetAgentError sets AgentError field to given value.

### HasAgentError

`func (o *KvmProtectionSource) HasAgentError() bool`

HasAgentError returns a boolean if a field has been set.

### SetAgentErrorNil

`func (o *KvmProtectionSource) SetAgentErrorNil(b bool)`

 SetAgentErrorNil sets the value for AgentError to be an explicit nil

### UnsetAgentError
`func (o *KvmProtectionSource) UnsetAgentError()`

UnsetAgentError ensures that no value is present for AgentError, not even an explicit nil
### GetAgentId

`func (o *KvmProtectionSource) GetAgentId() int64`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *KvmProtectionSource) GetAgentIdOk() (*int64, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *KvmProtectionSource) SetAgentId(v int64)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *KvmProtectionSource) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *KvmProtectionSource) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *KvmProtectionSource) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetClusterId

`func (o *KvmProtectionSource) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *KvmProtectionSource) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *KvmProtectionSource) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *KvmProtectionSource) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *KvmProtectionSource) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *KvmProtectionSource) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetDatacenterId

`func (o *KvmProtectionSource) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *KvmProtectionSource) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *KvmProtectionSource) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *KvmProtectionSource) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### SetDatacenterIdNil

`func (o *KvmProtectionSource) SetDatacenterIdNil(b bool)`

 SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil

### UnsetDatacenterId
`func (o *KvmProtectionSource) UnsetDatacenterId()`

UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
### GetDescription

`func (o *KvmProtectionSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KvmProtectionSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KvmProtectionSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KvmProtectionSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KvmProtectionSource) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KvmProtectionSource) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *KvmProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KvmProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KvmProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KvmProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KvmProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KvmProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetworkId

`func (o *KvmProtectionSource) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *KvmProtectionSource) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *KvmProtectionSource) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *KvmProtectionSource) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### SetNetworkIdNil

`func (o *KvmProtectionSource) SetNetworkIdNil(b bool)`

 SetNetworkIdNil sets the value for NetworkId to be an explicit nil

### UnsetNetworkId
`func (o *KvmProtectionSource) UnsetNetworkId()`

UnsetNetworkId ensures that no value is present for NetworkId, not even an explicit nil
### GetType

`func (o *KvmProtectionSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KvmProtectionSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KvmProtectionSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KvmProtectionSource) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *KvmProtectionSource) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *KvmProtectionSource) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUuid

`func (o *KvmProtectionSource) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *KvmProtectionSource) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *KvmProtectionSource) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *KvmProtectionSource) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *KvmProtectionSource) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *KvmProtectionSource) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


