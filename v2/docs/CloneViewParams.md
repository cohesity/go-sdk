# CloneViewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataLockExpiryUsecs** | Pointer to **NullableInt64** | DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If a view is marked as a DataLock view, only a Data Security Officer (a user having Data Security Privilege) can delete the view until the lock expiry time. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of the cloned View. | [optional] 
**IsReadOnly** | Pointer to **NullableBool** | Specifies if the view is a read only view. User will no longer be able to write to this view if this is set to true. | [optional] 
**Name** | **NullableString** | Specifies the name of the cloned View. | 
**NetgroupWhitelist** | Pointer to [**CloneViewParamsNetgroupWhitelist**](CloneViewParamsNetgroupWhitelist.md) |  | [optional] 
**ProtocolAccess** | Pointer to [**[]ViewProtocol**](ViewProtocol.md) | Specifies the supported Protocols for the View. | [optional] 
**Qos** | Pointer to [**CloneViewParamsQos**](CloneViewParamsQos.md) |  | [optional] 
**StoragePolicyOverride** | Pointer to [**CloneViewParamsStoragePolicyOverride**](CloneViewParamsStoragePolicyOverride.md) |  | [optional] 
**SubnetWhitelist** | Pointer to [**[]Subnet**](Subnet.md) | Array of Subnets. Specifies a list of Subnets with IP addresses that have permissions to access the View. (Overrides or extends the Subnets specified at the global Cohesity Cluster level.) | [optional] 

## Methods

### NewCloneViewParams

`func NewCloneViewParams(name NullableString, ) *CloneViewParams`

NewCloneViewParams instantiates a new CloneViewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneViewParamsWithDefaults

`func NewCloneViewParamsWithDefaults() *CloneViewParams`

NewCloneViewParamsWithDefaults instantiates a new CloneViewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataLockExpiryUsecs

`func (o *CloneViewParams) GetDataLockExpiryUsecs() int64`

GetDataLockExpiryUsecs returns the DataLockExpiryUsecs field if non-nil, zero value otherwise.

### GetDataLockExpiryUsecsOk

`func (o *CloneViewParams) GetDataLockExpiryUsecsOk() (*int64, bool)`

GetDataLockExpiryUsecsOk returns a tuple with the DataLockExpiryUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockExpiryUsecs

`func (o *CloneViewParams) SetDataLockExpiryUsecs(v int64)`

SetDataLockExpiryUsecs sets DataLockExpiryUsecs field to given value.

### HasDataLockExpiryUsecs

`func (o *CloneViewParams) HasDataLockExpiryUsecs() bool`

HasDataLockExpiryUsecs returns a boolean if a field has been set.

### SetDataLockExpiryUsecsNil

`func (o *CloneViewParams) SetDataLockExpiryUsecsNil(b bool)`

 SetDataLockExpiryUsecsNil sets the value for DataLockExpiryUsecs to be an explicit nil

### UnsetDataLockExpiryUsecs
`func (o *CloneViewParams) UnsetDataLockExpiryUsecs()`

UnsetDataLockExpiryUsecs ensures that no value is present for DataLockExpiryUsecs, not even an explicit nil
### GetDescription

`func (o *CloneViewParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloneViewParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloneViewParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloneViewParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CloneViewParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CloneViewParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsReadOnly

`func (o *CloneViewParams) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *CloneViewParams) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *CloneViewParams) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *CloneViewParams) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnlyNil

`func (o *CloneViewParams) SetIsReadOnlyNil(b bool)`

 SetIsReadOnlyNil sets the value for IsReadOnly to be an explicit nil

### UnsetIsReadOnly
`func (o *CloneViewParams) UnsetIsReadOnly()`

UnsetIsReadOnly ensures that no value is present for IsReadOnly, not even an explicit nil
### GetName

`func (o *CloneViewParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneViewParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneViewParams) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CloneViewParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CloneViewParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNetgroupWhitelist

`func (o *CloneViewParams) GetNetgroupWhitelist() CloneViewParamsNetgroupWhitelist`

GetNetgroupWhitelist returns the NetgroupWhitelist field if non-nil, zero value otherwise.

### GetNetgroupWhitelistOk

`func (o *CloneViewParams) GetNetgroupWhitelistOk() (*CloneViewParamsNetgroupWhitelist, bool)`

GetNetgroupWhitelistOk returns a tuple with the NetgroupWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetgroupWhitelist

`func (o *CloneViewParams) SetNetgroupWhitelist(v CloneViewParamsNetgroupWhitelist)`

SetNetgroupWhitelist sets NetgroupWhitelist field to given value.

### HasNetgroupWhitelist

`func (o *CloneViewParams) HasNetgroupWhitelist() bool`

HasNetgroupWhitelist returns a boolean if a field has been set.

### GetProtocolAccess

`func (o *CloneViewParams) GetProtocolAccess() []ViewProtocol`

GetProtocolAccess returns the ProtocolAccess field if non-nil, zero value otherwise.

### GetProtocolAccessOk

`func (o *CloneViewParams) GetProtocolAccessOk() (*[]ViewProtocol, bool)`

GetProtocolAccessOk returns a tuple with the ProtocolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAccess

`func (o *CloneViewParams) SetProtocolAccess(v []ViewProtocol)`

SetProtocolAccess sets ProtocolAccess field to given value.

### HasProtocolAccess

`func (o *CloneViewParams) HasProtocolAccess() bool`

HasProtocolAccess returns a boolean if a field has been set.

### SetProtocolAccessNil

`func (o *CloneViewParams) SetProtocolAccessNil(b bool)`

 SetProtocolAccessNil sets the value for ProtocolAccess to be an explicit nil

### UnsetProtocolAccess
`func (o *CloneViewParams) UnsetProtocolAccess()`

UnsetProtocolAccess ensures that no value is present for ProtocolAccess, not even an explicit nil
### GetQos

`func (o *CloneViewParams) GetQos() CloneViewParamsQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *CloneViewParams) GetQosOk() (*CloneViewParamsQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *CloneViewParams) SetQos(v CloneViewParamsQos)`

SetQos sets Qos field to given value.

### HasQos

`func (o *CloneViewParams) HasQos() bool`

HasQos returns a boolean if a field has been set.

### GetStoragePolicyOverride

`func (o *CloneViewParams) GetStoragePolicyOverride() CloneViewParamsStoragePolicyOverride`

GetStoragePolicyOverride returns the StoragePolicyOverride field if non-nil, zero value otherwise.

### GetStoragePolicyOverrideOk

`func (o *CloneViewParams) GetStoragePolicyOverrideOk() (*CloneViewParamsStoragePolicyOverride, bool)`

GetStoragePolicyOverrideOk returns a tuple with the StoragePolicyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyOverride

`func (o *CloneViewParams) SetStoragePolicyOverride(v CloneViewParamsStoragePolicyOverride)`

SetStoragePolicyOverride sets StoragePolicyOverride field to given value.

### HasStoragePolicyOverride

`func (o *CloneViewParams) HasStoragePolicyOverride() bool`

HasStoragePolicyOverride returns a boolean if a field has been set.

### GetSubnetWhitelist

`func (o *CloneViewParams) GetSubnetWhitelist() []Subnet`

GetSubnetWhitelist returns the SubnetWhitelist field if non-nil, zero value otherwise.

### GetSubnetWhitelistOk

`func (o *CloneViewParams) GetSubnetWhitelistOk() (*[]Subnet, bool)`

GetSubnetWhitelistOk returns a tuple with the SubnetWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetWhitelist

`func (o *CloneViewParams) SetSubnetWhitelist(v []Subnet)`

SetSubnetWhitelist sets SubnetWhitelist field to given value.

### HasSubnetWhitelist

`func (o *CloneViewParams) HasSubnetWhitelist() bool`

HasSubnetWhitelist returns a boolean if a field has been set.

### SetSubnetWhitelistNil

`func (o *CloneViewParams) SetSubnetWhitelistNil(b bool)`

 SetSubnetWhitelistNil sets the value for SubnetWhitelist to be an explicit nil

### UnsetSubnetWhitelist
`func (o *CloneViewParams) UnsetSubnetWhitelist()`

UnsetSubnetWhitelist ensures that no value is present for SubnetWhitelist, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


