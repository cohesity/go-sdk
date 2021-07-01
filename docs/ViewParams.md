# ViewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSubnetWhitelistVec** | Pointer to [**[]ClusterConfigProtoSubnet**](ClusterConfigProtoSubnet.md) | List of external client subnets from where requests will be received for the new view. | [optional] 
**DisableNfsAccess** | Pointer to **NullableBool** | Whether to disable NFS access in the new view. | [optional] 
**ProtocolAccessInfo** | Pointer to [**ViewIdMappingProtoProtocolAccessInfo**](ViewIdMappingProtoProtocolAccessInfo.md) |  | [optional] 
**QosMappingVec** | Pointer to [**[]ClusterConfigProtoQoSMapping**](ClusterConfigProtoQoSMapping.md) | The qos mappings (if any) for the new view. | [optional] 
**StoragePolicyOverride** | Pointer to [**ClusterConfigProtoStoragePolicyOverride**](ClusterConfigProtoStoragePolicyOverride.md) |  | [optional] 
**ViewDescription** | Pointer to **NullableString** | The description to be applied to the new view. | [optional] 
**WormLockExpiryUsecs** | Pointer to **NullableInt64** | This value &#39;worm_lock_expiry_usecs&#39; if specified will be set on the cloned view. This guarantees that the cloned view cannot be removed till the specified timestamp has reached. NOTE: If this is specified the clone view will be marked as immutable. | [optional] 

## Methods

### NewViewParams

`func NewViewParams() *ViewParams`

NewViewParams instantiates a new ViewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewParamsWithDefaults

`func NewViewParamsWithDefaults() *ViewParams`

NewViewParamsWithDefaults instantiates a new ViewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSubnetWhitelistVec

`func (o *ViewParams) GetClientSubnetWhitelistVec() []ClusterConfigProtoSubnet`

GetClientSubnetWhitelistVec returns the ClientSubnetWhitelistVec field if non-nil, zero value otherwise.

### GetClientSubnetWhitelistVecOk

`func (o *ViewParams) GetClientSubnetWhitelistVecOk() (*[]ClusterConfigProtoSubnet, bool)`

GetClientSubnetWhitelistVecOk returns a tuple with the ClientSubnetWhitelistVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSubnetWhitelistVec

`func (o *ViewParams) SetClientSubnetWhitelistVec(v []ClusterConfigProtoSubnet)`

SetClientSubnetWhitelistVec sets ClientSubnetWhitelistVec field to given value.

### HasClientSubnetWhitelistVec

`func (o *ViewParams) HasClientSubnetWhitelistVec() bool`

HasClientSubnetWhitelistVec returns a boolean if a field has been set.

### SetClientSubnetWhitelistVecNil

`func (o *ViewParams) SetClientSubnetWhitelistVecNil(b bool)`

 SetClientSubnetWhitelistVecNil sets the value for ClientSubnetWhitelistVec to be an explicit nil

### UnsetClientSubnetWhitelistVec
`func (o *ViewParams) UnsetClientSubnetWhitelistVec()`

UnsetClientSubnetWhitelistVec ensures that no value is present for ClientSubnetWhitelistVec, not even an explicit nil
### GetDisableNfsAccess

`func (o *ViewParams) GetDisableNfsAccess() bool`

GetDisableNfsAccess returns the DisableNfsAccess field if non-nil, zero value otherwise.

### GetDisableNfsAccessOk

`func (o *ViewParams) GetDisableNfsAccessOk() (*bool, bool)`

GetDisableNfsAccessOk returns a tuple with the DisableNfsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNfsAccess

`func (o *ViewParams) SetDisableNfsAccess(v bool)`

SetDisableNfsAccess sets DisableNfsAccess field to given value.

### HasDisableNfsAccess

`func (o *ViewParams) HasDisableNfsAccess() bool`

HasDisableNfsAccess returns a boolean if a field has been set.

### SetDisableNfsAccessNil

`func (o *ViewParams) SetDisableNfsAccessNil(b bool)`

 SetDisableNfsAccessNil sets the value for DisableNfsAccess to be an explicit nil

### UnsetDisableNfsAccess
`func (o *ViewParams) UnsetDisableNfsAccess()`

UnsetDisableNfsAccess ensures that no value is present for DisableNfsAccess, not even an explicit nil
### GetProtocolAccessInfo

`func (o *ViewParams) GetProtocolAccessInfo() ViewIdMappingProtoProtocolAccessInfo`

GetProtocolAccessInfo returns the ProtocolAccessInfo field if non-nil, zero value otherwise.

### GetProtocolAccessInfoOk

`func (o *ViewParams) GetProtocolAccessInfoOk() (*ViewIdMappingProtoProtocolAccessInfo, bool)`

GetProtocolAccessInfoOk returns a tuple with the ProtocolAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAccessInfo

`func (o *ViewParams) SetProtocolAccessInfo(v ViewIdMappingProtoProtocolAccessInfo)`

SetProtocolAccessInfo sets ProtocolAccessInfo field to given value.

### HasProtocolAccessInfo

`func (o *ViewParams) HasProtocolAccessInfo() bool`

HasProtocolAccessInfo returns a boolean if a field has been set.

### GetQosMappingVec

`func (o *ViewParams) GetQosMappingVec() []ClusterConfigProtoQoSMapping`

GetQosMappingVec returns the QosMappingVec field if non-nil, zero value otherwise.

### GetQosMappingVecOk

`func (o *ViewParams) GetQosMappingVecOk() (*[]ClusterConfigProtoQoSMapping, bool)`

GetQosMappingVecOk returns a tuple with the QosMappingVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMappingVec

`func (o *ViewParams) SetQosMappingVec(v []ClusterConfigProtoQoSMapping)`

SetQosMappingVec sets QosMappingVec field to given value.

### HasQosMappingVec

`func (o *ViewParams) HasQosMappingVec() bool`

HasQosMappingVec returns a boolean if a field has been set.

### SetQosMappingVecNil

`func (o *ViewParams) SetQosMappingVecNil(b bool)`

 SetQosMappingVecNil sets the value for QosMappingVec to be an explicit nil

### UnsetQosMappingVec
`func (o *ViewParams) UnsetQosMappingVec()`

UnsetQosMappingVec ensures that no value is present for QosMappingVec, not even an explicit nil
### GetStoragePolicyOverride

`func (o *ViewParams) GetStoragePolicyOverride() ClusterConfigProtoStoragePolicyOverride`

GetStoragePolicyOverride returns the StoragePolicyOverride field if non-nil, zero value otherwise.

### GetStoragePolicyOverrideOk

`func (o *ViewParams) GetStoragePolicyOverrideOk() (*ClusterConfigProtoStoragePolicyOverride, bool)`

GetStoragePolicyOverrideOk returns a tuple with the StoragePolicyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicyOverride

`func (o *ViewParams) SetStoragePolicyOverride(v ClusterConfigProtoStoragePolicyOverride)`

SetStoragePolicyOverride sets StoragePolicyOverride field to given value.

### HasStoragePolicyOverride

`func (o *ViewParams) HasStoragePolicyOverride() bool`

HasStoragePolicyOverride returns a boolean if a field has been set.

### GetViewDescription

`func (o *ViewParams) GetViewDescription() string`

GetViewDescription returns the ViewDescription field if non-nil, zero value otherwise.

### GetViewDescriptionOk

`func (o *ViewParams) GetViewDescriptionOk() (*string, bool)`

GetViewDescriptionOk returns a tuple with the ViewDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewDescription

`func (o *ViewParams) SetViewDescription(v string)`

SetViewDescription sets ViewDescription field to given value.

### HasViewDescription

`func (o *ViewParams) HasViewDescription() bool`

HasViewDescription returns a boolean if a field has been set.

### SetViewDescriptionNil

`func (o *ViewParams) SetViewDescriptionNil(b bool)`

 SetViewDescriptionNil sets the value for ViewDescription to be an explicit nil

### UnsetViewDescription
`func (o *ViewParams) UnsetViewDescription()`

UnsetViewDescription ensures that no value is present for ViewDescription, not even an explicit nil
### GetWormLockExpiryUsecs

`func (o *ViewParams) GetWormLockExpiryUsecs() int64`

GetWormLockExpiryUsecs returns the WormLockExpiryUsecs field if non-nil, zero value otherwise.

### GetWormLockExpiryUsecsOk

`func (o *ViewParams) GetWormLockExpiryUsecsOk() (*int64, bool)`

GetWormLockExpiryUsecsOk returns a tuple with the WormLockExpiryUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormLockExpiryUsecs

`func (o *ViewParams) SetWormLockExpiryUsecs(v int64)`

SetWormLockExpiryUsecs sets WormLockExpiryUsecs field to given value.

### HasWormLockExpiryUsecs

`func (o *ViewParams) HasWormLockExpiryUsecs() bool`

HasWormLockExpiryUsecs returns a boolean if a field has been set.

### SetWormLockExpiryUsecsNil

`func (o *ViewParams) SetWormLockExpiryUsecsNil(b bool)`

 SetWormLockExpiryUsecsNil sets the value for WormLockExpiryUsecs to be an explicit nil

### UnsetWormLockExpiryUsecs
`func (o *ViewParams) UnsetWormLockExpiryUsecs()`

UnsetWormLockExpiryUsecs ensures that no value is present for WormLockExpiryUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


